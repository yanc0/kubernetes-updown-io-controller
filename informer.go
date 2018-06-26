package main

import (
	"time"

	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
	client_v1alpha1 "github.com/yanc0/kubernetes-updown-io-controller/clientset/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func watchResources(clientSet client_v1alpha1.UpdownV1Alpha1Interface, namespace string) cache.Store {
	checkStore, checkController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				return clientSet.Checks(namespace).List(lo)
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				return clientSet.Checks(namespace).Watch(lo)
			},
		},
		&v1alpha1.Check{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go checkController.Run(wait.NeverStop)
	return checkStore
}
