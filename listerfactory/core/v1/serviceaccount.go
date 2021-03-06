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

var _ v1.ServiceAccountLister = &serviceAccountLister{}

var _ v1.ServiceAccountNamespaceLister = &serviceAccountNamespaceLister{}

// serviceAccountLister implements the ServiceAccountLister interface.
type serviceAccountLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewServiceAccountLister returns a new ServiceAccountLister.
func NewServiceAccountLister(client kubernetes.Interface) v1.ServiceAccountLister {
	return NewFilteredServiceAccountLister(client, nil)
}

func NewFilteredServiceAccountLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.ServiceAccountLister {
	return &serviceAccountLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all ServiceAccounts in the indexer.
func (s *serviceAccountLister) List(selector labels.Selector) (ret []*core_v1.ServiceAccount, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoreV1().ServiceAccounts(meta_v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// ServiceAccounts returns an object that can list and get ServiceAccounts.
func (s *serviceAccountLister) ServiceAccounts(namespace string) v1.ServiceAccountNamespaceLister {
	return serviceAccountNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// serviceAccountNamespaceLister implements the ServiceAccountNamespaceLister
// interface.
type serviceAccountNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all ServiceAccounts in the indexer for a given namespace.
func (s serviceAccountNamespaceLister) List(selector labels.Selector) (ret []*core_v1.ServiceAccount, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CoreV1().ServiceAccounts(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the ServiceAccount from the indexer for a given namespace and name.
func (s serviceAccountNamespaceLister) Get(name string) (*core_v1.ServiceAccount, error) {
	return s.client.CoreV1().ServiceAccounts(s.namespace).Get(name, meta_v1.GetOptions{})
}
