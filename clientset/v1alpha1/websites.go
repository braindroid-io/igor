package v1alpha1

import (
	"github.com/braindroid-io/igor/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type WebSiteInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.WebSiteList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.WebSite, error)
	Create(*v1alpha1.WebSite) (*v1alpha1.WebSite, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type websiteClient struct {
	restClient rest.Interface
	ns         string
}

func (c *websiteClient) List(opts metav1.ListOptions) (*v1alpha1.WebSiteList, error) {
	result := v1alpha1.WebSiteList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("websites").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *websiteClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.WebSite, error) {
	result := v1alpha1.WebSite{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("websites").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *websiteClient) Create(website *v1alpha1.WebSite) (*v1alpha1.WebSite, error) {
	result := v1alpha1.WebSite{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("websites").
		Body(website).
		Do().
		Into(&result)

	return &result, err
}

func (c *websiteClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("websites").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
