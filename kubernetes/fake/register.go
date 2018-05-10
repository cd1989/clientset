/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	apiextensionsv1beta1 "github.com/caicloud/clientset/pkg/apis/apiextensions/v1beta1"
	apiregistrationv1 "github.com/caicloud/clientset/pkg/apis/apiregistration/v1"
	cnetworkingv1alpha1 "github.com/caicloud/clientset/pkg/apis/cnetworking/v1alpha1"
	configv1alpha1 "github.com/caicloud/clientset/pkg/apis/config/v1alpha1"
	loadbalancev1alpha2 "github.com/caicloud/clientset/pkg/apis/loadbalance/v1alpha2"
	releasev1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	resourcev1alpha1 "github.com/caicloud/clientset/pkg/apis/resource/v1alpha1"
	resourcev1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	tenantv1alpha1 "github.com/caicloud/clientset/pkg/apis/tenant/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	fake "k8s.io/client-go/kubernetes/fake"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	fake.AddToScheme(scheme)
	AddToScheme(scheme)
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
func AddToScheme(scheme *runtime.Scheme) {
	apiextensionsv1beta1.AddToScheme(scheme)
	apiregistrationv1.AddToScheme(scheme)
	cnetworkingv1alpha1.AddToScheme(scheme)
	configv1alpha1.AddToScheme(scheme)
	loadbalancev1alpha2.AddToScheme(scheme)
	releasev1alpha1.AddToScheme(scheme)
	resourcev1alpha1.AddToScheme(scheme)
	resourcev1beta1.AddToScheme(scheme)
	tenantv1alpha1.AddToScheme(scheme)
}
