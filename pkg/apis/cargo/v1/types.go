package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cargo struct {
	// Metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`

	// Metadata for the particular object, including name, namespace, labels, etc
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Custom resource specification
	Spec CargoSpec `json:"spec"`
}

type CargoSpec struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CargoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cargo `json:"items""`
}
