/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageTypes implements StorageTypeInterface
type FakeStorageTypes struct {
	Fake *FakeResourceV1beta1
}

var storagetypesResource = schema.GroupVersionResource{Group: "resource.caicloud.io", Version: "v1beta1", Resource: "storagetypes"}

var storagetypesKind = schema.GroupVersionKind{Group: "resource.caicloud.io", Version: "v1beta1", Kind: "StorageType"}

// Get takes name of the storageType, and returns the corresponding storageType object, and an error if there is any.
func (c *FakeStorageTypes) Get(name string, options v1.GetOptions) (result *v1beta1.StorageType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storagetypesResource, name), &v1beta1.StorageType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageType), err
}

// List takes label and field selectors, and returns the list of StorageTypes that match those selectors.
func (c *FakeStorageTypes) List(opts v1.ListOptions) (result *v1beta1.StorageTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storagetypesResource, storagetypesKind, opts), &v1beta1.StorageTypeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.StorageTypeList{}
	for _, item := range obj.(*v1beta1.StorageTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageTypes.
func (c *FakeStorageTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storagetypesResource, opts))
}

// Create takes the representation of a storageType and creates it.  Returns the server's representation of the storageType, and an error, if there is any.
func (c *FakeStorageTypes) Create(storageType *v1beta1.StorageType) (result *v1beta1.StorageType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storagetypesResource, storageType), &v1beta1.StorageType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageType), err
}

// Update takes the representation of a storageType and updates it. Returns the server's representation of the storageType, and an error, if there is any.
func (c *FakeStorageTypes) Update(storageType *v1beta1.StorageType) (result *v1beta1.StorageType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storagetypesResource, storageType), &v1beta1.StorageType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageType), err
}

// Delete takes name of the storageType and deletes it. Returns an error if one occurs.
func (c *FakeStorageTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storagetypesResource, name), &v1beta1.StorageType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storagetypesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.StorageTypeList{})
	return err
}

// Patch applies the patch and returns the patched storageType.
func (c *FakeStorageTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.StorageType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storagetypesResource, name, data, subresources...), &v1beta1.StorageType{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StorageType), err
}
