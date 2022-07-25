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

type ProjectOrganizationPolicyBooleanPolicyObservation struct {
}

type ProjectOrganizationPolicyBooleanPolicyParameters struct {

	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	// +kubebuilder:validation:Required
	Enforced *bool `json:"enforced" tf:"enforced,omitempty"`
}

type ProjectOrganizationPolicyListPolicyAllowObservation struct {
}

type ProjectOrganizationPolicyListPolicyAllowParameters struct {

	// The policy allows or denies all values.
	// +kubebuilder:validation:Optional
	All *bool `json:"all,omitempty" tf:"all,omitempty"`

	// The policy can define specific values that are allowed or denied.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicyListPolicyDenyObservation struct {
}

type ProjectOrganizationPolicyListPolicyDenyParameters struct {

	// The policy allows or denies all values.
	// +kubebuilder:validation:Optional
	All *bool `json:"all,omitempty" tf:"all,omitempty"`

	// The policy can define specific values that are allowed or denied.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ProjectOrganizationPolicyListPolicyObservation struct {
}

type ProjectOrganizationPolicyListPolicyParameters struct {

	// One or the other must be set.
	// +kubebuilder:validation:Optional
	Allow []ProjectOrganizationPolicyListPolicyAllowParameters `json:"allow,omitempty" tf:"allow,omitempty"`

	// One or the other must be set.
	// +kubebuilder:validation:Optional
	Deny []ProjectOrganizationPolicyListPolicyDenyParameters `json:"deny,omitempty" tf:"deny,omitempty"`

	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	// +kubebuilder:validation:Optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent,omitempty"`

	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	// +kubebuilder:validation:Optional
	SuggestedValue *string `json:"suggestedValue,omitempty" tf:"suggested_value,omitempty"`
}

type ProjectOrganizationPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type ProjectOrganizationPolicyParameters struct {

	// A boolean policy is a constraint that is either enforced or not.
	// +kubebuilder:validation:Optional
	BooleanPolicy []ProjectOrganizationPolicyBooleanPolicyParameters `json:"booleanPolicy,omitempty" tf:"boolean_policy,omitempty"`

	// The name of the Constraint the Policy is configuring, for example, serviceuser.services.
	// +kubebuilder:validation:Required
	Constraint *string `json:"constraint" tf:"constraint,omitempty"`

	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values.
	// +kubebuilder:validation:Optional
	ListPolicy []ProjectOrganizationPolicyListPolicyParameters `json:"listPolicy,omitempty" tf:"list_policy,omitempty"`

	// The project ID.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// A restore policy is a constraint to restore the default policy.
	// +kubebuilder:validation:Optional
	RestorePolicy []ProjectOrganizationPolicyRestorePolicyParameters `json:"restorePolicy,omitempty" tf:"restore_policy,omitempty"`

	// Version of the Policy. Default version is 0.
	// +kubebuilder:validation:Optional
	Version *int64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ProjectOrganizationPolicyRestorePolicyObservation struct {
}

type ProjectOrganizationPolicyRestorePolicyParameters struct {

	// May only be set to true. If set, then the default Policy is restored.
	// +kubebuilder:validation:Required
	Default *bool `json:"default" tf:"default,omitempty"`
}

// ProjectOrganizationPolicySpec defines the desired state of ProjectOrganizationPolicy
type ProjectOrganizationPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectOrganizationPolicyParameters `json:"forProvider"`
}

// ProjectOrganizationPolicyStatus defines the observed state of ProjectOrganizationPolicy.
type ProjectOrganizationPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectOrganizationPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectOrganizationPolicy is the Schema for the ProjectOrganizationPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectOrganizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectOrganizationPolicySpec   `json:"spec"`
	Status            ProjectOrganizationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectOrganizationPolicyList contains a list of ProjectOrganizationPolicys
type ProjectOrganizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectOrganizationPolicy `json:"items"`
}

// Repository type metadata.
var (
	ProjectOrganizationPolicy_Kind             = "ProjectOrganizationPolicy"
	ProjectOrganizationPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectOrganizationPolicy_Kind}.String()
	ProjectOrganizationPolicy_KindAPIVersion   = ProjectOrganizationPolicy_Kind + "." + CRDGroupVersion.String()
	ProjectOrganizationPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ProjectOrganizationPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectOrganizationPolicy{}, &ProjectOrganizationPolicyList{})
}