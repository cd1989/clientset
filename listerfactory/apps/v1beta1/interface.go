/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/apps/v1beta1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // ControllerRevisions returns a ControllerRevisionLister
	ControllerRevisions() v1beta1.ControllerRevisionLister
	// Deployments returns a DeploymentLister
	Deployments() v1beta1.DeploymentLister
	// StatefulSets returns a StatefulSetLister
	StatefulSets() v1beta1.StatefulSetLister
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

// ControllerRevisions returns a ControllerRevisionLister.
func (v *version) ControllerRevisions() v1beta1.ControllerRevisionLister {
	return &controllerRevisionLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// ControllerRevisions returns a ControllerRevisionLister.
func (v *infromerVersion) ControllerRevisions() v1beta1.ControllerRevisionLister {
	return v.factory.Apps().V1beta1().ControllerRevisions().Lister()
}

// Deployments returns a DeploymentLister.
func (v *version) Deployments() v1beta1.DeploymentLister {
	return &deploymentLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentLister.
func (v *infromerVersion) Deployments() v1beta1.DeploymentLister {
	return v.factory.Apps().V1beta1().Deployments().Lister()
}

// StatefulSets returns a StatefulSetLister.
func (v *version) StatefulSets() v1beta1.StatefulSetLister {
	return &statefulSetLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// StatefulSets returns a StatefulSetLister.
func (v *infromerVersion) StatefulSets() v1beta1.StatefulSetLister {
	return v.factory.Apps().V1beta1().StatefulSets().Lister()
}
