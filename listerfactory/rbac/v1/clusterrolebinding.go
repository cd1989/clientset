/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	rbac_v1 "k8s.io/api/rbac/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/rbac/v1"
)

var _ v1.ClusterRoleBindingLister = &clusterRoleBindingLister{}

// clusterRoleBindingLister implements the ClusterRoleBindingLister interface.
type clusterRoleBindingLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterRoleBindingLister returns a new ClusterRoleBindingLister.
func NewClusterRoleBindingLister(client kubernetes.Interface) v1.ClusterRoleBindingLister {
	return NewFilteredClusterRoleBindingLister(client, nil)
}

func NewFilteredClusterRoleBindingLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.ClusterRoleBindingLister {
	return &clusterRoleBindingLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all ClusterRoleBindings in the indexer.
func (s *clusterRoleBindingLister) List(selector labels.Selector) (ret []*rbac_v1.ClusterRoleBinding, err error) {
	listopt := meta_v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.RbacV1().ClusterRoleBindings().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the ClusterRoleBinding from the index for a given name.
func (s *clusterRoleBindingLister) Get(name string) (*rbac_v1.ClusterRoleBinding, error) {
	return s.client.RbacV1().ClusterRoleBindings().Get(name, meta_v1.GetOptions{})
}
