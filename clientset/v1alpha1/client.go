package v1alpha1

import (
	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type UpdownV1Alpha1Interface interface {
	Checks(namespace string) CheckInterface
}

type UpdownV1Alpha1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*UpdownV1Alpha1Client, error) {
	v1alpha1.AddToScheme(scheme.Scheme)
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &UpdownV1Alpha1Client{restClient: client}, nil
}

func (c *UpdownV1Alpha1Client) Checks(namespace string) CheckInterface {
	return &checkClient{
		restClient: c.restClient,
		ns:         namespace,
	}
}
