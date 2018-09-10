/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/core/v1"
)

var _ v1.ComponentStatusLister = &componentStatusLister{}

// componentStatusLister implements the ComponentStatusLister interface.
type componentStatusLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewComponentStatusLister returns a new ComponentStatusLister.
func NewComponentStatusLister(client kubernetes.Interface) v1.ComponentStatusLister {
	return NewFilteredComponentStatusLister(client, nil)
}

func NewFilteredComponentStatusLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.ComponentStatusLister {
	return &componentStatusLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all ComponentStatuses in the indexer.
func (s *componentStatusLister) List(selector labels.Selector) (ret []*core_v1.ComponentStatus, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoreV1().ComponentStatuses().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the ComponentStatus from the index for a given name.
func (s *componentStatusLister) Get(name string) (*core_v1.ComponentStatus, error) {
	return s.client.CoreV1().ComponentStatuses().Get(name, meta_v1.GetOptions{})
}