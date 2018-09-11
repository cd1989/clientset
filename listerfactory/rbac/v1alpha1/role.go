/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	rbacv1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/rbac/v1alpha1"
)

var _ v1alpha1.RoleLister = &roleLister{}

var _ v1alpha1.RoleNamespaceLister = &roleNamespaceLister{}

// roleLister implements the RoleLister interface.
type roleLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRoleLister returns a new RoleLister.
func NewRoleLister(client kubernetes.Interface) v1alpha1.RoleLister {
	return NewFilteredRoleLister(client, nil)
}

func NewFilteredRoleLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1alpha1.RoleLister {
	return &roleLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all Roles in the indexer.
func (s *roleLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.RbacV1alpha1().Roles(v1.NamespaceAll).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Roles returns an object that can list and get Roles.
func (s *roleLister) Roles(namespace string) v1alpha1.RoleNamespaceLister {
	return roleNamespaceLister{client: s.client, tweakListOptions: s.tweakListOptions, namespace: namespace}
}

// roleNamespaceLister implements the RoleNamespaceLister
// interface.
type roleNamespaceLister struct {
	client           kubernetes.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// List lists all Roles in the indexer for a given namespace.
func (s roleNamespaceLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.RbacV1alpha1().Roles(s.namespace).List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the Role from the indexer for a given namespace and name.
func (s roleNamespaceLister) Get(name string) (*rbacv1alpha1.Role, error) {
	return s.client.RbacV1alpha1().Roles(s.namespace).Get(name, v1.GetOptions{})
}
