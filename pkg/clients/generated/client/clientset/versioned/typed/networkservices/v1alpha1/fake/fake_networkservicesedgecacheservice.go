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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesEdgeCacheServices implements NetworkServicesEdgeCacheServiceInterface
type FakeNetworkServicesEdgeCacheServices struct {
	Fake *FakeNetworkservicesV1alpha1
	ns   string
}

var networkservicesedgecacheservicesResource = schema.GroupVersionResource{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "networkservicesedgecacheservices"}

var networkservicesedgecacheservicesKind = schema.GroupVersionKind{Group: "networkservices.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "NetworkServicesEdgeCacheService"}

// Get takes name of the networkServicesEdgeCacheService, and returns the corresponding networkServicesEdgeCacheService object, and an error if there is any.
func (c *FakeNetworkServicesEdgeCacheServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkServicesEdgeCacheService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicesedgecacheservicesResource, c.ns, name), &v1alpha1.NetworkServicesEdgeCacheService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheService), err
}

// List takes label and field selectors, and returns the list of NetworkServicesEdgeCacheServices that match those selectors.
func (c *FakeNetworkServicesEdgeCacheServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkServicesEdgeCacheServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicesedgecacheservicesResource, networkservicesedgecacheservicesKind, c.ns, opts), &v1alpha1.NetworkServicesEdgeCacheServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkServicesEdgeCacheServiceList{ListMeta: obj.(*v1alpha1.NetworkServicesEdgeCacheServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkServicesEdgeCacheServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesEdgeCacheServices.
func (c *FakeNetworkServicesEdgeCacheServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicesedgecacheservicesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesEdgeCacheService and creates it.  Returns the server's representation of the networkServicesEdgeCacheService, and an error, if there is any.
func (c *FakeNetworkServicesEdgeCacheServices) Create(ctx context.Context, networkServicesEdgeCacheService *v1alpha1.NetworkServicesEdgeCacheService, opts v1.CreateOptions) (result *v1alpha1.NetworkServicesEdgeCacheService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicesedgecacheservicesResource, c.ns, networkServicesEdgeCacheService), &v1alpha1.NetworkServicesEdgeCacheService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheService), err
}

// Update takes the representation of a networkServicesEdgeCacheService and updates it. Returns the server's representation of the networkServicesEdgeCacheService, and an error, if there is any.
func (c *FakeNetworkServicesEdgeCacheServices) Update(ctx context.Context, networkServicesEdgeCacheService *v1alpha1.NetworkServicesEdgeCacheService, opts v1.UpdateOptions) (result *v1alpha1.NetworkServicesEdgeCacheService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicesedgecacheservicesResource, c.ns, networkServicesEdgeCacheService), &v1alpha1.NetworkServicesEdgeCacheService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesEdgeCacheServices) UpdateStatus(ctx context.Context, networkServicesEdgeCacheService *v1alpha1.NetworkServicesEdgeCacheService, opts v1.UpdateOptions) (*v1alpha1.NetworkServicesEdgeCacheService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicesedgecacheservicesResource, "status", c.ns, networkServicesEdgeCacheService), &v1alpha1.NetworkServicesEdgeCacheService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheService), err
}

// Delete takes name of the networkServicesEdgeCacheService and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesEdgeCacheServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicesedgecacheservicesResource, c.ns, name, opts), &v1alpha1.NetworkServicesEdgeCacheService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesEdgeCacheServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicesedgecacheservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkServicesEdgeCacheServiceList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesEdgeCacheService.
func (c *FakeNetworkServicesEdgeCacheServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkServicesEdgeCacheService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicesedgecacheservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkServicesEdgeCacheService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkServicesEdgeCacheService), err
}
