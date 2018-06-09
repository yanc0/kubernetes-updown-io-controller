package v1alpha1

import (
	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type CheckInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.CheckList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Check, error)
	Create(*v1alpha1.Check) (*v1alpha1.Check, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type checkClient struct {
	restClient rest.Interface
	ns         string
}

func (c *checkClient) List(opts metav1.ListOptions) (*v1alpha1.CheckList, error) {
	result := v1alpha1.CheckList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("checks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *checkClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Check, error) {
	result := v1alpha1.Check{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("checks").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *checkClient) Create(Check *v1alpha1.Check) (*v1alpha1.Check, error) {
	result := v1alpha1.Check{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("checks").
		Body(Check).
		Do().
		Into(&result)

	return &result, err
}

func (c *checkClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("checks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
