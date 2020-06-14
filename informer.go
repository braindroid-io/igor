package main

import (
	"time"

	"github.com/braindroid-io/igor/api/types/v1alpha1"
	client_v1alpha1 "github.com/braindroid-io/igor/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet client_v1alpha1.V1Alpha1Interface) cache.Store {
	websiteStore, websiteController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.WebSites("").List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.WebSites("").Watch(lo)
			},
		},
		&v1alpha1.WebSite{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go websiteController.Run(wait.NeverStop)
	return websiteStore
}
