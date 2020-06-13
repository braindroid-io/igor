package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/braindroid-io/igor/api/types/v1alpha1"
	clientV1alpha1 "github.com/braindroid-io/igor/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err)
	}

	v1alpha1.AddToScheme(scheme.Scheme)

	clientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	websites, err := clientSet.WebSites("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("websites found: %+v\n", projects)

	store := WatchResources(clientSet)

	for {
		websitesFromStore := store.List()
		fmt.Printf("websites in store: %d\n", len(websitesFromStore))

		time.Sleep(2 * time.Second)
	}
}