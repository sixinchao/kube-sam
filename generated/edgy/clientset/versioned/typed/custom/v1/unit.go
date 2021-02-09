/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "builder/api/custom/v1"
	scheme "builder/generated/edgy/clientset/versioned/scheme"
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UnitsGetter has a method to return a UnitInterface.
// A group's client should implement this interface.
type UnitsGetter interface {
	Units(namespace string) UnitInterface
}

// UnitInterface has methods to work with Unit resources.
type UnitInterface interface {
	Create(ctx context.Context, unit *v1.Unit, opts metav1.CreateOptions) (*v1.Unit, error)
	Update(ctx context.Context, unit *v1.Unit, opts metav1.UpdateOptions) (*v1.Unit, error)
	UpdateStatus(ctx context.Context, unit *v1.Unit, opts metav1.UpdateOptions) (*v1.Unit, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Unit, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UnitList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Unit, err error)
	UnitExpansion
}

// units implements UnitInterface
type units struct {
	client rest.Interface
	ns     string
}

// newUnits returns a Units
func newUnits(c *CustomV1Client, namespace string) *units {
	return &units{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the unit, and returns the corresponding unit object, and an error if there is any.
func (c *units) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Unit, err error) {
	result = &v1.Unit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("units").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Units that match those selectors.
func (c *units) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UnitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UnitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("units").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested units.
func (c *units) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("units").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a unit and creates it.  Returns the server's representation of the unit, and an error, if there is any.
func (c *units) Create(ctx context.Context, unit *v1.Unit, opts metav1.CreateOptions) (result *v1.Unit, err error) {
	result = &v1.Unit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("units").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unit).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a unit and updates it. Returns the server's representation of the unit, and an error, if there is any.
func (c *units) Update(ctx context.Context, unit *v1.Unit, opts metav1.UpdateOptions) (result *v1.Unit, err error) {
	result = &v1.Unit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("units").
		Name(unit.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unit).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *units) UpdateStatus(ctx context.Context, unit *v1.Unit, opts metav1.UpdateOptions) (result *v1.Unit, err error) {
	result = &v1.Unit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("units").
		Name(unit.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(unit).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the unit and deletes it. Returns an error if one occurs.
func (c *units) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("units").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *units) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("units").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched unit.
func (c *units) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Unit, err error) {
	result = &v1.Unit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("units").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
