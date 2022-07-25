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

type ConsentStoreIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConsentStoreIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConsentStoreID *string `json:"consentStoreId" tf:"consent_store_id,omitempty"`

	// +kubebuilder:validation:Required
	Dataset *string `json:"dataset" tf:"dataset,omitempty"`

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`
}

// ConsentStoreIAMPolicySpec defines the desired state of ConsentStoreIAMPolicy
type ConsentStoreIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsentStoreIAMPolicyParameters `json:"forProvider"`
}

// ConsentStoreIAMPolicyStatus defines the observed state of ConsentStoreIAMPolicy.
type ConsentStoreIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsentStoreIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConsentStoreIAMPolicy is the Schema for the ConsentStoreIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ConsentStoreIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsentStoreIAMPolicySpec   `json:"spec"`
	Status            ConsentStoreIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsentStoreIAMPolicyList contains a list of ConsentStoreIAMPolicys
type ConsentStoreIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsentStoreIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	ConsentStoreIAMPolicy_Kind             = "ConsentStoreIAMPolicy"
	ConsentStoreIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConsentStoreIAMPolicy_Kind}.String()
	ConsentStoreIAMPolicy_KindAPIVersion   = ConsentStoreIAMPolicy_Kind + "." + CRDGroupVersion.String()
	ConsentStoreIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ConsentStoreIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ConsentStoreIAMPolicy{}, &ConsentStoreIAMPolicyList{})
}