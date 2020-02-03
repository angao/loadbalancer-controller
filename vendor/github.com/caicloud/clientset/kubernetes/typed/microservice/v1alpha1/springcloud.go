/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/microservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpringcloudsGetter has a method to return a SpringcloudInterface.
// A group's client should implement this interface.
type SpringcloudsGetter interface {
	Springclouds(namespace string) SpringcloudInterface
}

// SpringcloudInterface has methods to work with Springcloud resources.
type SpringcloudInterface interface {
	Create(*v1alpha1.Springcloud) (*v1alpha1.Springcloud, error)
	Update(*v1alpha1.Springcloud) (*v1alpha1.Springcloud, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Springcloud, error)
	List(opts v1.ListOptions) (*v1alpha1.SpringcloudList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Springcloud, err error)
	SpringcloudExpansion
}

// springclouds implements SpringcloudInterface
type springclouds struct {
	client rest.Interface
	ns     string
}

// newSpringclouds returns a Springclouds
func newSpringclouds(c *MicroserviceV1alpha1Client, namespace string) *springclouds {
	return &springclouds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the springcloud, and returns the corresponding springcloud object, and an error if there is any.
func (c *springclouds) Get(name string, options v1.GetOptions) (result *v1alpha1.Springcloud, err error) {
	result = &v1alpha1.Springcloud{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("springclouds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Springclouds that match those selectors.
func (c *springclouds) List(opts v1.ListOptions) (result *v1alpha1.SpringcloudList, err error) {
	result = &v1alpha1.SpringcloudList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("springclouds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested springclouds.
func (c *springclouds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("springclouds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a springcloud and creates it.  Returns the server's representation of the springcloud, and an error, if there is any.
func (c *springclouds) Create(springcloud *v1alpha1.Springcloud) (result *v1alpha1.Springcloud, err error) {
	result = &v1alpha1.Springcloud{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("springclouds").
		Body(springcloud).
		Do().
		Into(result)
	return
}

// Update takes the representation of a springcloud and updates it. Returns the server's representation of the springcloud, and an error, if there is any.
func (c *springclouds) Update(springcloud *v1alpha1.Springcloud) (result *v1alpha1.Springcloud, err error) {
	result = &v1alpha1.Springcloud{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("springclouds").
		Name(springcloud.Name).
		Body(springcloud).
		Do().
		Into(result)
	return
}

// Delete takes name of the springcloud and deletes it. Returns an error if one occurs.
func (c *springclouds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("springclouds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *springclouds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("springclouds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched springcloud.
func (c *springclouds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Springcloud, err error) {
	result = &v1alpha1.Springcloud{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("springclouds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
