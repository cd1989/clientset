/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// Configs returns a ConfigInformer.
	Configs() ConfigInformer
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// StorageServices returns a StorageServiceInformer.
	StorageServices() StorageServiceInformer
	// StorageTypes returns a StorageTypeInformer.
	StorageTypes() StorageTypeInformer
	// Tags returns a TagInformer.
	Tags() TagInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Configs returns a ConfigInformer.
func (v *version) Configs() ConfigInformer {
	return &configInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StorageServices returns a StorageServiceInformer.
func (v *version) StorageServices() StorageServiceInformer {
	return &storageServiceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StorageTypes returns a StorageTypeInformer.
func (v *version) StorageTypes() StorageTypeInformer {
	return &storageTypeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Tags returns a TagInformer.
func (v *version) Tags() TagInformer {
	return &tagInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
