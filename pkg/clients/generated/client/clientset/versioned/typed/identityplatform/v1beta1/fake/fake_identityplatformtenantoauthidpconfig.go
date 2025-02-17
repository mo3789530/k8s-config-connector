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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityPlatformTenantOAuthIDPConfigs implements IdentityPlatformTenantOAuthIDPConfigInterface
type FakeIdentityPlatformTenantOAuthIDPConfigs struct {
	Fake *FakeIdentityplatformV1beta1
	ns   string
}

var identityplatformtenantoauthidpconfigsResource = schema.GroupVersionResource{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1beta1", Resource: "identityplatformtenantoauthidpconfigs"}

var identityplatformtenantoauthidpconfigsKind = schema.GroupVersionKind{Group: "identityplatform.cnrm.cloud.google.com", Version: "v1beta1", Kind: "IdentityPlatformTenantOAuthIDPConfig"}

// Get takes name of the identityPlatformTenantOAuthIDPConfig, and returns the corresponding identityPlatformTenantOAuthIDPConfig object, and an error if there is any.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identityplatformtenantoauthidpconfigsResource, c.ns, name), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfig), err
}

// List takes label and field selectors, and returns the list of IdentityPlatformTenantOAuthIDPConfigs that match those selectors.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identityplatformtenantoauthidpconfigsResource, identityplatformtenantoauthidpconfigsKind, c.ns, opts), &v1beta1.IdentityPlatformTenantOAuthIDPConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IdentityPlatformTenantOAuthIDPConfigList{ListMeta: obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityPlatformTenantOAuthIDPConfigs.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identityplatformtenantoauthidpconfigsResource, c.ns, opts))

}

// Create takes the representation of a identityPlatformTenantOAuthIDPConfig and creates it.  Returns the server's representation of the identityPlatformTenantOAuthIDPConfig, and an error, if there is any.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Create(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.CreateOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identityplatformtenantoauthidpconfigsResource, c.ns, identityPlatformTenantOAuthIDPConfig), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfig), err
}

// Update takes the representation of a identityPlatformTenantOAuthIDPConfig and updates it. Returns the server's representation of the identityPlatformTenantOAuthIDPConfig, and an error, if there is any.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Update(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identityplatformtenantoauthidpconfigsResource, c.ns, identityPlatformTenantOAuthIDPConfig), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) UpdateStatus(ctx context.Context, identityPlatformTenantOAuthIDPConfig *v1beta1.IdentityPlatformTenantOAuthIDPConfig, opts v1.UpdateOptions) (*v1beta1.IdentityPlatformTenantOAuthIDPConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(identityplatformtenantoauthidpconfigsResource, "status", c.ns, identityPlatformTenantOAuthIDPConfig), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfig), err
}

// Delete takes name of the identityPlatformTenantOAuthIDPConfig and deletes it. Returns an error if one occurs.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(identityplatformtenantoauthidpconfigsResource, c.ns, name, opts), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identityplatformtenantoauthidpconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IdentityPlatformTenantOAuthIDPConfigList{})
	return err
}

// Patch applies the patch and returns the patched identityPlatformTenantOAuthIDPConfig.
func (c *FakeIdentityPlatformTenantOAuthIDPConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IdentityPlatformTenantOAuthIDPConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identityplatformtenantoauthidpconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.IdentityPlatformTenantOAuthIDPConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityPlatformTenantOAuthIDPConfig), err
}
