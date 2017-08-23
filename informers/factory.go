/*
Copyright 2017 caicloud authors. All rights reserved.
*/

// This file was automatically generated by informer-gen

package informers

import (
	release "github.com/caicloud/clientset/informers/release"
	kubernetes "github.com/caicloud/clientset/kubernetes"
	informers "k8s.io/client-go/informers"
	time "time"
)

type sharedInformerFactory struct {
	informers.SharedInformerFactory
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory
func NewSharedInformerFactory(client kubernetes.Interface, defaultResync time.Duration) SharedInformerFactory {
	return &sharedInformerFactory{
		informers.NewSharedInformerFactory(client, defaultResync),
	}
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	informers.SharedInformerFactory

	Release() release.Interface
}

func (f *sharedInformerFactory) Release() release.Interface {
	return release.New(f)
}