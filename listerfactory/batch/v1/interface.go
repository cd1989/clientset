/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/batch/v1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // Jobs returns a JobLister
	Jobs() v1.JobLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// Jobs returns a JobLister.
func (v *version) Jobs() v1.JobLister {
	return &jobLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// Jobs returns a JobLister.
func (v *infromerVersion) Jobs() v1.JobLister {
	return v.factory.Batch().V1().Jobs().Lister()
}
