/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	certificates_v1beta1 "k8s.io/api/certificates/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/certificates/v1beta1"
)

var _ v1beta1.CertificateSigningRequestLister = &certificateSigningRequestLister{}

// certificateSigningRequestLister implements the CertificateSigningRequestLister interface.
type certificateSigningRequestLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCertificateSigningRequestLister returns a new CertificateSigningRequestLister.
func NewCertificateSigningRequestLister(client kubernetes.Interface) v1beta1.CertificateSigningRequestLister {
	return NewFilteredCertificateSigningRequestLister(client, nil)
}

func NewFilteredCertificateSigningRequestLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.CertificateSigningRequestLister {
	return &certificateSigningRequestLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all CertificateSigningRequests in the indexer.
func (s *certificateSigningRequestLister) List(selector labels.Selector) (ret []*certificates_v1beta1.CertificateSigningRequest, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.CertificatesV1beta1().CertificateSigningRequests().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the CertificateSigningRequest from the index for a given name.
func (s *certificateSigningRequestLister) Get(name string) (*certificates_v1beta1.CertificateSigningRequest, error) {
	return s.client.CertificatesV1beta1().CertificateSigningRequests().Get(name, v1.GetOptions{})
}
