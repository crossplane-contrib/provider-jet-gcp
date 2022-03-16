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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkObservation struct {
	GatewayIPv4 *string `json:"gatewayIpv4,omitempty" tf:"gateway_ipv4,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type NetworkParameters struct {

	// When set to 'true', the network is created in "auto subnet mode" and
	// it will create a subnet for each region automatically across the
	// '10.128.0.0/9' address range.
	//
	// When set to 'false', the network is created in "custom subnet mode" so
	// the user can explicitly connect subnetwork resources.
	// +kubebuilder:validation:Optional
	AutoCreateSubnetworks *bool `json:"autoCreateSubnetworks,omitempty" tf:"auto_create_subnetworks,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteDefaultRoutesOnCreate *bool `json:"deleteDefaultRoutesOnCreate,omitempty" tf:"delete_default_routes_on_create,omitempty"`

	// An optional description of this resource. The resource must be
	// recreated to modify this field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Maximum Transmission Unit in bytes. The minimum value for this field is 1460
	// and the maximum value is 1500 bytes.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The network-wide routing mode to use. If set to 'REGIONAL', this
	// network's cloud routers will only advertise routes with subnetworks
	// of this network in the same region as the router. If set to 'GLOBAL',
	// this network's cloud routers will advertise routes with all
	// subnetworks of this network, across regions. Possible values: ["REGIONAL", "GLOBAL"]
	// +kubebuilder:validation:Optional
	RoutingMode *string `json:"routingMode,omitempty" tf:"routing_mode,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Network is the Schema for the Networks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSpec   `json:"spec"`
	Status            NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
