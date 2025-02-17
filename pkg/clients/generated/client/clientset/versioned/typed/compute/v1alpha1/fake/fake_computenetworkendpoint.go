// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeNetworkEndpoints implements ComputeNetworkEndpointInterface
type FakeComputeNetworkEndpoints struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computenetworkendpointsResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "computenetworkendpoints"}

var computenetworkendpointsKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeNetworkEndpoint"}

// Get takes name of the computeNetworkEndpoint, and returns the corresponding computeNetworkEndpoint object, and an error if there is any.
func (c *FakeComputeNetworkEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeNetworkEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computenetworkendpointsResource, c.ns, name), &v1alpha1.ComputeNetworkEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkEndpoint), err
}

// List takes label and field selectors, and returns the list of ComputeNetworkEndpoints that match those selectors.
func (c *FakeComputeNetworkEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeNetworkEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computenetworkendpointsResource, computenetworkendpointsKind, c.ns, opts), &v1alpha1.ComputeNetworkEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeNetworkEndpointList{ListMeta: obj.(*v1alpha1.ComputeNetworkEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeNetworkEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeNetworkEndpoints.
func (c *FakeComputeNetworkEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computenetworkendpointsResource, c.ns, opts))

}

// Create takes the representation of a computeNetworkEndpoint and creates it.  Returns the server's representation of the computeNetworkEndpoint, and an error, if there is any.
func (c *FakeComputeNetworkEndpoints) Create(ctx context.Context, computeNetworkEndpoint *v1alpha1.ComputeNetworkEndpoint, opts v1.CreateOptions) (result *v1alpha1.ComputeNetworkEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computenetworkendpointsResource, c.ns, computeNetworkEndpoint), &v1alpha1.ComputeNetworkEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkEndpoint), err
}

// Update takes the representation of a computeNetworkEndpoint and updates it. Returns the server's representation of the computeNetworkEndpoint, and an error, if there is any.
func (c *FakeComputeNetworkEndpoints) Update(ctx context.Context, computeNetworkEndpoint *v1alpha1.ComputeNetworkEndpoint, opts v1.UpdateOptions) (result *v1alpha1.ComputeNetworkEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computenetworkendpointsResource, c.ns, computeNetworkEndpoint), &v1alpha1.ComputeNetworkEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeNetworkEndpoints) UpdateStatus(ctx context.Context, computeNetworkEndpoint *v1alpha1.ComputeNetworkEndpoint, opts v1.UpdateOptions) (*v1alpha1.ComputeNetworkEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computenetworkendpointsResource, "status", c.ns, computeNetworkEndpoint), &v1alpha1.ComputeNetworkEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkEndpoint), err
}

// Delete takes name of the computeNetworkEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeComputeNetworkEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computenetworkendpointsResource, c.ns, name, opts), &v1alpha1.ComputeNetworkEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeNetworkEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computenetworkendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeNetworkEndpointList{})
	return err
}

// Patch applies the patch and returns the patched computeNetworkEndpoint.
func (c *FakeComputeNetworkEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeNetworkEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computenetworkendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeNetworkEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeNetworkEndpoint), err
}
