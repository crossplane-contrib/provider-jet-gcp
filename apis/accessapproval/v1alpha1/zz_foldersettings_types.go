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

type EnrolledServicesObservation struct {
}

type EnrolledServicesParameters struct {

	// The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
	// * all
	// * App Engine
	// * BigQuery
	// * Cloud Bigtable
	// * Cloud Key Management Service
	// * Compute Engine
	// * Cloud Dataflow
	// * Cloud Identity and Access Management
	// * Cloud Pub/Sub
	// * Cloud Storage
	// * Persistent Disk
	//
	// Note: These values are supported as input, but considered a legacy format:
	// * all
	// * appengine.googleapis.com
	// * bigquery.googleapis.com
	// * bigtable.googleapis.com
	// * cloudkms.googleapis.com
	// * compute.googleapis.com
	// * dataflow.googleapis.com
	// * iam.googleapis.com
	// * pubsub.googleapis.com
	// * storage.googleapis.com
	// +kubebuilder:validation:Required
	CloudProduct *string `json:"cloudProduct" tf:"cloud_product,omitempty"`

	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]
	// +kubebuilder:validation:Optional
	EnrollmentLevel *string `json:"enrollmentLevel,omitempty" tf:"enrollment_level,omitempty"`
}

type FolderSettingsObservation struct {
	EnrolledAncestor *bool `json:"enrolledAncestor,omitempty" tf:"enrolled_ancestor,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FolderSettingsParameters struct {

	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	//
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	// +kubebuilder:validation:Required
	EnrolledServices []EnrolledServicesParameters `json:"enrolledServices" tf:"enrolled_services,omitempty"`

	// ID of the folder of the access approval settings.
	// +kubebuilder:validation:Required
	FolderID *string `json:"folderId" tf:"folder_id,omitempty"`

	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	// +kubebuilder:validation:Optional
	NotificationEmails []*string `json:"notificationEmails,omitempty" tf:"notification_emails,omitempty"`
}

// FolderSettingsSpec defines the desired state of FolderSettings
type FolderSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderSettingsParameters `json:"forProvider"`
}

// FolderSettingsStatus defines the observed state of FolderSettings.
type FolderSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderSettings is the Schema for the FolderSettingss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FolderSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderSettingsSpec   `json:"spec"`
	Status            FolderSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderSettingsList contains a list of FolderSettingss
type FolderSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderSettings `json:"items"`
}

// Repository type metadata.
var (
	FolderSettings_Kind             = "FolderSettings"
	FolderSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderSettings_Kind}.String()
	FolderSettings_KindAPIVersion   = FolderSettings_Kind + "." + CRDGroupVersion.String()
	FolderSettings_GroupVersionKind = CRDGroupVersion.WithKind(FolderSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderSettings{}, &FolderSettingsList{})
}