package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ProjectSpec is the project specification
type ProjectSpec struct {
	Replicas int `json:"replicas"`
}

// Project is the project
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectSpec `json:"spec"`
}

// ProjectList is the project list
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
}
