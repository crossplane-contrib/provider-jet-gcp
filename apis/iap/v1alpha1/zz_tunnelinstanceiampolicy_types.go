/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TunnelInstanceIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TunnelInstanceIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	Instance *string `json:"instance" tf:"instance,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// TunnelInstanceIAMPolicySpec defines the desired state of TunnelInstanceIAMPolicy
type TunnelInstanceIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TunnelInstanceIAMPolicyParameters `json:"forProvider"`
}

// TunnelInstanceIAMPolicyStatus defines the observed state of TunnelInstanceIAMPolicy.
type TunnelInstanceIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TunnelInstanceIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelInstanceIAMPolicy is the Schema for the TunnelInstanceIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type TunnelInstanceIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TunnelInstanceIAMPolicySpec   `json:"spec"`
	Status            TunnelInstanceIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TunnelInstanceIAMPolicyList contains a list of TunnelInstanceIAMPolicys
type TunnelInstanceIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TunnelInstanceIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	TunnelInstanceIAMPolicy_Kind             = "TunnelInstanceIAMPolicy"
	TunnelInstanceIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TunnelInstanceIAMPolicy_Kind}.String()
	TunnelInstanceIAMPolicy_KindAPIVersion   = TunnelInstanceIAMPolicy_Kind + "." + CRDGroupVersion.String()
	TunnelInstanceIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TunnelInstanceIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TunnelInstanceIAMPolicy{}, &TunnelInstanceIAMPolicyList{})
}