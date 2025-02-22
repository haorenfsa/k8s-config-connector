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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudiot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudIOTDevices implements CloudIOTDeviceInterface
type FakeCloudIOTDevices struct {
	Fake *FakeCloudiotV1alpha1
	ns   string
}

var cloudiotdevicesResource = schema.GroupVersionResource{Group: "cloudiot.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "cloudiotdevices"}

var cloudiotdevicesKind = schema.GroupVersionKind{Group: "cloudiot.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "CloudIOTDevice"}

// Get takes name of the cloudIOTDevice, and returns the corresponding cloudIOTDevice object, and an error if there is any.
func (c *FakeCloudIOTDevices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudIOTDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudiotdevicesResource, c.ns, name), &v1alpha1.CloudIOTDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIOTDevice), err
}

// List takes label and field selectors, and returns the list of CloudIOTDevices that match those selectors.
func (c *FakeCloudIOTDevices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudIOTDeviceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudiotdevicesResource, cloudiotdevicesKind, c.ns, opts), &v1alpha1.CloudIOTDeviceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudIOTDeviceList{ListMeta: obj.(*v1alpha1.CloudIOTDeviceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudIOTDeviceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudIOTDevices.
func (c *FakeCloudIOTDevices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudiotdevicesResource, c.ns, opts))

}

// Create takes the representation of a cloudIOTDevice and creates it.  Returns the server's representation of the cloudIOTDevice, and an error, if there is any.
func (c *FakeCloudIOTDevices) Create(ctx context.Context, cloudIOTDevice *v1alpha1.CloudIOTDevice, opts v1.CreateOptions) (result *v1alpha1.CloudIOTDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudiotdevicesResource, c.ns, cloudIOTDevice), &v1alpha1.CloudIOTDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIOTDevice), err
}

// Update takes the representation of a cloudIOTDevice and updates it. Returns the server's representation of the cloudIOTDevice, and an error, if there is any.
func (c *FakeCloudIOTDevices) Update(ctx context.Context, cloudIOTDevice *v1alpha1.CloudIOTDevice, opts v1.UpdateOptions) (result *v1alpha1.CloudIOTDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudiotdevicesResource, c.ns, cloudIOTDevice), &v1alpha1.CloudIOTDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIOTDevice), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudIOTDevices) UpdateStatus(ctx context.Context, cloudIOTDevice *v1alpha1.CloudIOTDevice, opts v1.UpdateOptions) (*v1alpha1.CloudIOTDevice, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudiotdevicesResource, "status", c.ns, cloudIOTDevice), &v1alpha1.CloudIOTDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIOTDevice), err
}

// Delete takes name of the cloudIOTDevice and deletes it. Returns an error if one occurs.
func (c *FakeCloudIOTDevices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudiotdevicesResource, c.ns, name, opts), &v1alpha1.CloudIOTDevice{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudIOTDevices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudiotdevicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudIOTDeviceList{})
	return err
}

// Patch applies the patch and returns the patched cloudIOTDevice.
func (c *FakeCloudIOTDevices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudIOTDevice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudiotdevicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudIOTDevice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIOTDevice), err
}
