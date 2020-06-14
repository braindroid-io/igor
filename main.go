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
	log.Printf("Igor is starting...")

	var config *rest.Config
	var err error

	log.Printf("What configuration do I have, masterrr?..")

	if kubeconfig == "" {
		log.Printf("Ahhh, yes. In-cluster configuration...")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("Yes, master... Using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		log.Printf("Oh nooo! We have an error!")
		log.Panic(err.Error())
	}

	v1alpha1.AddToScheme(scheme.Scheme)

	clientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		log.Printf("Oh nooo! We have an error!")
		log.Panic(err.Error())
	}

	websites, err := clientSet.WebSites("default").List(metav1.ListOptions{})
	if err != nil {
		log.Printf("Oh nooo! We have an error!")
		log.Panic(err.Error())
	}

	fmt.Printf("Websites found: %+v\n", websites)

	store := WatchResources(clientSet)

	for {
		websitesFromStore := store.List()
		fmt.Printf("We have %d websites in store...\n", len(websitesFromStore))
		time.Sleep(2 * time.Second)
	}
}
