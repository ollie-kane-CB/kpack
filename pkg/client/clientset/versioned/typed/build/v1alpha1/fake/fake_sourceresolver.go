/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSourceResolvers implements SourceResolverInterface
type FakeSourceResolvers struct {
	Fake *FakeKpackV1alpha1
	ns   string
}

var sourceresolversResource = schema.GroupVersionResource{Group: "kpack.io", Version: "v1alpha1", Resource: "sourceresolvers"}

var sourceresolversKind = schema.GroupVersionKind{Group: "kpack.io", Version: "v1alpha1", Kind: "SourceResolver"}

// Get takes name of the sourceResolver, and returns the corresponding sourceResolver object, and an error if there is any.
func (c *FakeSourceResolvers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sourceresolversResource, c.ns, name), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// List takes label and field selectors, and returns the list of SourceResolvers that match those selectors.
func (c *FakeSourceResolvers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SourceResolverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sourceresolversResource, sourceresolversKind, c.ns, opts), &v1alpha1.SourceResolverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SourceResolverList{ListMeta: obj.(*v1alpha1.SourceResolverList).ListMeta}
	for _, item := range obj.(*v1alpha1.SourceResolverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sourceResolvers.
func (c *FakeSourceResolvers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sourceresolversResource, c.ns, opts))

}

// Create takes the representation of a sourceResolver and creates it.  Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *FakeSourceResolvers) Create(ctx context.Context, sourceResolver *v1alpha1.SourceResolver, opts v1.CreateOptions) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sourceresolversResource, c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// Update takes the representation of a sourceResolver and updates it. Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *FakeSourceResolvers) Update(ctx context.Context, sourceResolver *v1alpha1.SourceResolver, opts v1.UpdateOptions) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sourceresolversResource, c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSourceResolvers) UpdateStatus(ctx context.Context, sourceResolver *v1alpha1.SourceResolver, opts v1.UpdateOptions) (*v1alpha1.SourceResolver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sourceresolversResource, "status", c.ns, sourceResolver), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}

// Delete takes name of the sourceResolver and deletes it. Returns an error if one occurs.
func (c *FakeSourceResolvers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(sourceresolversResource, c.ns, name, opts), &v1alpha1.SourceResolver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSourceResolvers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sourceresolversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SourceResolverList{})
	return err
}

// Patch applies the patch and returns the patched sourceResolver.
func (c *FakeSourceResolvers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SourceResolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sourceresolversResource, c.ns, name, pt, data, subresources...), &v1alpha1.SourceResolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SourceResolver), err
}
