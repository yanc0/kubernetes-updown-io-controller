package main

import (
	"flag"
	clientV1alpha1 "github.com/yanc0/kubernetes-updown-io-controller/clientset/v1alpha1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"log"
	"time"
	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
	"github.com/yanc0/kubernetes-updown-io-controller/updown"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig string
var updownAPIKey string

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.StringVar(&updownAPIKey, "apikey", "", "updown.io api key")
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

	if updownAPIKey == "" {
		panic("apiKey parameter is missing")
	}

	clientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	store := WatchResources(clientSet, "default")



	for {
		checks := make([]*v1alpha1.Check, 0)

		checksFromStore := store.List()
		for _, check := range checksFromStore {
			checks = append(checks, check.(*v1alpha1.Check))
		}
		updown.Sync(checks, updownAPIKey)
		time.Sleep(2 * time.Second)
	}
}
