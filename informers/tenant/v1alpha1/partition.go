/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/tenant/v1alpha1"
	tenantv1alpha1 "github.com/caicloud/clientset/pkg/apis/tenant/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// PartitionInformer provides access to a shared informer and lister for
// Partitions.
type PartitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PartitionLister
}

type partitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPartitionInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPartitionInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPartitionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPartitionInformer constructs a new informer for Partition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPartitionInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha1().Partitions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TenantV1alpha1().Partitions().Watch(options)
			},
		},
		&tenantv1alpha1.Partition{},
		resyncPeriod,
		indexers,
	)
}

func (f *partitionInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPartitionInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *partitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tenantv1alpha1.Partition{}, f.defaultInformer)
}

func (f *partitionInformer) Lister() v1alpha1.PartitionLister {
	return v1alpha1.NewPartitionLister(f.Informer().GetIndexer())
}
