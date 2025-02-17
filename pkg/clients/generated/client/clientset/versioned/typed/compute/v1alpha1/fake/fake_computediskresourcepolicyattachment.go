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

// FakeComputeDiskResourcePolicyAttachments implements ComputeDiskResourcePolicyAttachmentInterface
type FakeComputeDiskResourcePolicyAttachments struct {
	Fake *FakeComputeV1alpha1
	ns   string
}

var computediskresourcepolicyattachmentsResource = schema.GroupVersionResource{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "computediskresourcepolicyattachments"}

var computediskresourcepolicyattachmentsKind = schema.GroupVersionKind{Group: "compute.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "ComputeDiskResourcePolicyAttachment"}

// Get takes name of the computeDiskResourcePolicyAttachment, and returns the corresponding computeDiskResourcePolicyAttachment object, and an error if there is any.
func (c *FakeComputeDiskResourcePolicyAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computediskresourcepolicyattachmentsResource, c.ns, name), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeDiskResourcePolicyAttachment), err
}

// List takes label and field selectors, and returns the list of ComputeDiskResourcePolicyAttachments that match those selectors.
func (c *FakeComputeDiskResourcePolicyAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computediskresourcepolicyattachmentsResource, computediskresourcepolicyattachmentsKind, c.ns, opts), &v1alpha1.ComputeDiskResourcePolicyAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeDiskResourcePolicyAttachmentList{ListMeta: obj.(*v1alpha1.ComputeDiskResourcePolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeDiskResourcePolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeDiskResourcePolicyAttachments.
func (c *FakeComputeDiskResourcePolicyAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computediskresourcepolicyattachmentsResource, c.ns, opts))

}

// Create takes the representation of a computeDiskResourcePolicyAttachment and creates it.  Returns the server's representation of the computeDiskResourcePolicyAttachment, and an error, if there is any.
func (c *FakeComputeDiskResourcePolicyAttachments) Create(ctx context.Context, computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment, opts v1.CreateOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computediskresourcepolicyattachmentsResource, c.ns, computeDiskResourcePolicyAttachment), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeDiskResourcePolicyAttachment), err
}

// Update takes the representation of a computeDiskResourcePolicyAttachment and updates it. Returns the server's representation of the computeDiskResourcePolicyAttachment, and an error, if there is any.
func (c *FakeComputeDiskResourcePolicyAttachments) Update(ctx context.Context, computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment, opts v1.UpdateOptions) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computediskresourcepolicyattachmentsResource, c.ns, computeDiskResourcePolicyAttachment), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeDiskResourcePolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeDiskResourcePolicyAttachments) UpdateStatus(ctx context.Context, computeDiskResourcePolicyAttachment *v1alpha1.ComputeDiskResourcePolicyAttachment, opts v1.UpdateOptions) (*v1alpha1.ComputeDiskResourcePolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computediskresourcepolicyattachmentsResource, "status", c.ns, computeDiskResourcePolicyAttachment), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeDiskResourcePolicyAttachment), err
}

// Delete takes name of the computeDiskResourcePolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeComputeDiskResourcePolicyAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computediskresourcepolicyattachmentsResource, c.ns, name, opts), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeDiskResourcePolicyAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computediskresourcepolicyattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeDiskResourcePolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched computeDiskResourcePolicyAttachment.
func (c *FakeComputeDiskResourcePolicyAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeDiskResourcePolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computediskresourcepolicyattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeDiskResourcePolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeDiskResourcePolicyAttachment), err
}
