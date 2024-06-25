/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VSphereVirtualDiskObservation struct {
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	VmdkPath *string `json:"vmdkPath,omitempty" tf:"vmdk_path,omitempty"`
}

type VSphereVirtualDiskParameters struct {

	// +kubebuilder:validation:Optional
	AdapterType *string `json:"adapterType,omitempty" tf:"adapter_type,omitempty"`

	// +kubebuilder:validation:Optional
	CreateDirectories *bool `json:"createDirectories,omitempty" tf:"create_directories,omitempty"`

	// +kubebuilder:validation:Optional
	Datacenter *string `json:"datacenter,omitempty" tf:"datacenter,omitempty"`

	// +kubebuilder:validation:Optional
	Datastore *string `json:"datastore,omitempty" tf:"datastore,omitempty"`

	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	VmdkPath *string `json:"vmdkPath,omitempty" tf:"vmdk_path,omitempty"`
}

// VSphereVirtualDiskSpec defines the desired state of VSphereVirtualDisk
type VSphereVirtualDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereVirtualDiskParameters `json:"forProvider"`
}

// VSphereVirtualDiskStatus defines the observed state of VSphereVirtualDisk.
type VSphereVirtualDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereVirtualDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereVirtualDisk is the Schema for the VSphereVirtualDisks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereVirtualDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.datastore)",message="datastore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.size)",message="size is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vmdkPath)",message="vmdkPath is a required parameter"
	Spec   VSphereVirtualDiskSpec   `json:"spec"`
	Status VSphereVirtualDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereVirtualDiskList contains a list of VSphereVirtualDisks
type VSphereVirtualDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereVirtualDisk `json:"items"`
}

// Repository type metadata.
var (
	VSphereVirtualDisk_Kind             = "VSphereVirtualDisk"
	VSphereVirtualDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereVirtualDisk_Kind}.String()
	VSphereVirtualDisk_KindAPIVersion   = VSphereVirtualDisk_Kind + "." + CRDGroupVersion.String()
	VSphereVirtualDisk_GroupVersionKind = CRDGroupVersion.WithKind(VSphereVirtualDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereVirtualDisk{}, &VSphereVirtualDiskList{})
}
