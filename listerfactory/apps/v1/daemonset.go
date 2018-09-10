/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	apps_v1 "k8s.io/api/apps/v1"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/apps/v1"
)

var _ v1.DaemonSetLister = &daemonSetLister{}

var _ v1.DaemonSetNamespaceLister = &daemonSetNamespaceLister{}

// daemonSetLister implements the DaemonSetLister interface.
type daemonSetLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDaemonSetLister returns a new DaemonSetLister.
func NewDaemonSetLister(client kubernetes.Interface) v1.DaemonSetLister {
	return NewFilteredDaemonSetLister(client, nil)
}

func NewFilteredDaemonSetLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.DaemonSetLister {
	return &daemonSetLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all DaemonSets in the indexer.
func (s *daemonSetLister) List(selector labels.Selector) (ret []*apps_v1.DaemonSet, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AppsV1().DaemonSets(meta_v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

func (s *daemonSetLister) GetHistoryDaemonSets(*apps_v1.ControllerRevision) ([]*apps_v1.DaemonSet, error) {
	return nil, nil
}

func (s *daemonSetLister) GetPodDaemonSets(*core_v1.Pod) ([]*apps_v1.DaemonSet, error) {
	return nil, nil
}

// DaemonSets returns an object that can list and get DaemonSets.
func (s *daemonSetLister) DaemonSets(namespace string) v1.DaemonSetNamespaceLister {
	return daemonSetNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// daemonSetNamespaceLister implements the DaemonSetNamespaceLister
// interface.
type daemonSetNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all DaemonSets in the indexer for a given namespace.
func (s daemonSetNamespaceLister) List(selector labels.Selector) (ret []*apps_v1.DaemonSet, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.AppsV1().DaemonSets(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the DaemonSet from the indexer for a given namespace and name.
func (s daemonSetNamespaceLister) Get(name string) (*apps_v1.DaemonSet, error) {
	return s.client.AppsV1().DaemonSets(s.namespace).Get(name, meta_v1.GetOptions{})
}
