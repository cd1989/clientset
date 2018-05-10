/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/caicloud/clientset/kubernetes"
	apiextensionsv1beta1 "github.com/caicloud/clientset/kubernetes/typed/apiextensions/v1beta1"
	fakeapiextensionsv1beta1 "github.com/caicloud/clientset/kubernetes/typed/apiextensions/v1beta1/fake"
	apiregistrationv1 "github.com/caicloud/clientset/kubernetes/typed/apiregistration/v1"
	fakeapiregistrationv1 "github.com/caicloud/clientset/kubernetes/typed/apiregistration/v1/fake"
	cnetworkingv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/cnetworking/v1alpha1"
	fakecnetworkingv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/cnetworking/v1alpha1/fake"
	configv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/config/v1alpha1"
	fakeconfigv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/config/v1alpha1/fake"
	loadbalancev1alpha2 "github.com/caicloud/clientset/kubernetes/typed/loadbalance/v1alpha2"
	fakeloadbalancev1alpha2 "github.com/caicloud/clientset/kubernetes/typed/loadbalance/v1alpha2/fake"
	releasev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/release/v1alpha1"
	fakereleasev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/release/v1alpha1/fake"
	resourcev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1alpha1"
	fakeresourcev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1alpha1/fake"
	resourcev1beta1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1beta1"
	fakeresourcev1beta1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1beta1/fake"
	tenantv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/tenant/v1alpha1"
	faketenantv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/tenant/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	fake "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return &Clientset{&fake.Clientset{Fake: fakePtr}, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	*fake.Clientset
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// ApiextensionsV1beta1 retrieves the ApiextensionsV1beta1Client
func (c *Clientset) ApiextensionsV1beta1() apiextensionsv1beta1.ApiextensionsV1beta1Interface {
	return &fakeapiextensionsv1beta1.FakeApiextensionsV1beta1{Fake: &c.Fake}
}

// Apiextensions retrieves the ApiextensionsV1beta1Client
func (c *Clientset) Apiextensions() apiextensionsv1beta1.ApiextensionsV1beta1Interface {
	return &fakeapiextensionsv1beta1.FakeApiextensionsV1beta1{Fake: &c.Fake}
}

// ApiregistrationV1 retrieves the ApiregistrationV1Client
func (c *Clientset) ApiregistrationV1() apiregistrationv1.ApiregistrationV1Interface {
	return &fakeapiregistrationv1.FakeApiregistrationV1{Fake: &c.Fake}
}

// Apiregistration retrieves the ApiregistrationV1Client
func (c *Clientset) Apiregistration() apiregistrationv1.ApiregistrationV1Interface {
	return &fakeapiregistrationv1.FakeApiregistrationV1{Fake: &c.Fake}
}

// CnetworkingV1alpha1 retrieves the CnetworkingV1alpha1Client
func (c *Clientset) CnetworkingV1alpha1() cnetworkingv1alpha1.CnetworkingV1alpha1Interface {
	return &fakecnetworkingv1alpha1.FakeCnetworkingV1alpha1{Fake: &c.Fake}
}

// Cnetworking retrieves the CnetworkingV1alpha1Client
func (c *Clientset) Cnetworking() cnetworkingv1alpha1.CnetworkingV1alpha1Interface {
	return &fakecnetworkingv1alpha1.FakeCnetworkingV1alpha1{Fake: &c.Fake}
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return &fakeconfigv1alpha1.FakeConfigV1alpha1{Fake: &c.Fake}
}

// Config retrieves the ConfigV1alpha1Client
func (c *Clientset) Config() configv1alpha1.ConfigV1alpha1Interface {
	return &fakeconfigv1alpha1.FakeConfigV1alpha1{Fake: &c.Fake}
}

// LoadbalanceV1alpha2 retrieves the LoadbalanceV1alpha2Client
func (c *Clientset) LoadbalanceV1alpha2() loadbalancev1alpha2.LoadbalanceV1alpha2Interface {
	return &fakeloadbalancev1alpha2.FakeLoadbalanceV1alpha2{Fake: &c.Fake}
}

// Loadbalance retrieves the LoadbalanceV1alpha2Client
func (c *Clientset) Loadbalance() loadbalancev1alpha2.LoadbalanceV1alpha2Interface {
	return &fakeloadbalancev1alpha2.FakeLoadbalanceV1alpha2{Fake: &c.Fake}
}

// ReleaseV1alpha1 retrieves the ReleaseV1alpha1Client
func (c *Clientset) ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface {
	return &fakereleasev1alpha1.FakeReleaseV1alpha1{Fake: &c.Fake}
}

// Release retrieves the ReleaseV1alpha1Client
func (c *Clientset) Release() releasev1alpha1.ReleaseV1alpha1Interface {
	return &fakereleasev1alpha1.FakeReleaseV1alpha1{Fake: &c.Fake}
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return &fakeresourcev1alpha1.FakeResourceV1alpha1{Fake: &c.Fake}
}

// ResourceV1beta1 retrieves the ResourceV1beta1Client
func (c *Clientset) ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface {
	return &fakeresourcev1beta1.FakeResourceV1beta1{Fake: &c.Fake}
}

// Resource retrieves the ResourceV1beta1Client
func (c *Clientset) Resource() resourcev1beta1.ResourceV1beta1Interface {
	return &fakeresourcev1beta1.FakeResourceV1beta1{Fake: &c.Fake}
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return &faketenantv1alpha1.FakeTenantV1alpha1{Fake: &c.Fake}
}

// Tenant retrieves the TenantV1alpha1Client
func (c *Clientset) Tenant() tenantv1alpha1.TenantV1alpha1Interface {
	return &faketenantv1alpha1.FakeTenantV1alpha1{Fake: &c.Fake}
}
