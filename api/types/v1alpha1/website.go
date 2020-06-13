package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type WebSiteSpec struct {
	Image string `json:"image"`
	Template string `json:"template"`
	Replicas int `json:"replicas"`
}

type WebSite struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WebSiteSpec `json:"spec"`
}

type WebSiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []WebSite `json:"items"`
}
