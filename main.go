package main

import (
	"flag"
	"fmt"
	"time"
	"strings"
	"os"

	"github.com/braindroid-io/igor/util/logger"
	"github.com/braindroid-io/igor/api/types/v1alpha1"
	clientV1alpha1 "github.com/braindroid-io/igor/clientset/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/apps/v1"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string
var log *logger.Logger

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func int32Ptr(i int) *int32 {
	intPtr := int32(i)
	return &intPtr
}

func printWebSite(ws *v1alpha1.WebSite) {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("UID: %s, Namespace: %s, Image: %s, Template: %s, Replicas: %d",
			ws.ObjectMeta.GetUID(),
			ws.ObjectMeta.GetNamespace(),
			ws.Spec.Image,
			ws.Spec.Template,
			ws.Spec.Replicas,
	));

	log.Info(b.String())
}

func getDeployments(ws *v1alpha1.WebSite, clientSet *kubernetes.Clientset) (*v1.DeploymentList, error) {
	api := clientSet.AppsV1().Deployments(ws.ObjectMeta.GetNamespace())

	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("website.lineage=%s", string(ws.ObjectMeta.GetUID())),
	}

	list, err := api.List(listOptions)

	return list, err
}

func createDeploymentObject(ws *v1alpha1.WebSite) (*v1.Deployment) {
	return &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: ws.ObjectMeta.Name,
			Labels: map[string]string{
				"app": "igor",
				"website.template": ws.Spec.Template,
				"website.lineage": string(ws.ObjectMeta.GetUID()),
			},
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(ws.Spec.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "igor",
					"website.lineage": string(ws.ObjectMeta.GetUID()),
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "igor",
						"website.template": ws.Spec.Template,
						"website.lineage": string(ws.ObjectMeta.GetUID()),
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  ws.ObjectMeta.Name,
							Image: ws.Spec.Image,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}

func createDeploymentForWebSite(ws *v1alpha1.WebSite, clientSet *kubernetes.Clientset) (*v1.Deployment, error) {
	api := clientSet.AppsV1().Deployments(ws.ObjectMeta.GetNamespace())

	deployment := createDeploymentObject(ws)

	result, err := api.Create(deployment)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func reconcileDeployments(webSites []interface{}, clientSet *kubernetes.Clientset) {
	var ws *v1alpha1.WebSite

	log.Info("Reconciling deployments for website definitions")

	for i := 0; i < len(webSites); i++ {
		ws = webSites[i].(*v1alpha1.WebSite)

		printWebSite(ws);

		deployments, err := getDeployments(ws, clientSet)
		if err != nil {
			printMessageAndBail("Oh nooo! We have an error getting deployments!", err)
		}

		if len(deployments.Items) == 0 {
			log.Info("No deployments found, masterrr...");

			deployment, err := createDeploymentForWebSite(ws, clientSet)
			if err != nil {
				printMessageAndBail("Oh nooo! We have an error creating a deployment!", err)
			}

			log.Info("Created deployment: %s", deployment.ObjectMeta.Name);

			return
		}
	}
}

func printMessageAndBail(msg string, err error) {
	log.Info(msg)
	log.Error(err.Error())
	os.Exit(1)
}

func main() {
	log.Info("Igor is starting...")

	var config *rest.Config
	var err error

	log.Info("What configuration do I have, masterrr?..")

	if kubeconfig == "" {
		log.Info("Ahhh, yes. In-cluster configuration...")
		config, err = rest.InClusterConfig()
	} else {
		log.Info("Yes, master... Using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		printMessageAndBail("Oh nooo! We have an error!", err)
	}

	v1alpha1.AddToScheme(scheme.Scheme)

	v1Alpha1ClientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		printMessageAndBail("Oh nooo! We have an error!", err)
	}

	store := WatchResources(v1Alpha1ClientSet)

	k8sClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		printMessageAndBail("Oh nooo! We have an error! Unable to get K8s clientset.", err)
	}

	for {
		websitesFromStore := store.List()

		if len(websitesFromStore) > 0 {
			log.Info("We have %d websites in the store...\n", len(websitesFromStore))

			reconcileDeployments(websitesFromStore, k8sClientSet);
		}

		time.Sleep(2 * time.Second)
	}
}
