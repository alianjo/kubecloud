package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ClusterID",type=string,JSONPath=`.status.vmID`
// +kubebuilder:printcolumn:name="Progress",type=string,JSONPath=`.status.progress`
type Vm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VmSpec        `json:"spec,omitempty"`
	Status KlsuterStatus `json:"status,omitempty"`
}

type KlsuterStatus struct {
	VmID     string `json:"vmID,omitempty"`
	Progress string `json:"progress,omitempty"`
}

type VmSpec struct {
	Name string `json:"name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Vm `json:"items,omitempty"`
}
