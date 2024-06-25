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

type PublicationObservation struct {
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	PublishURL *string `json:"publishUrl,omitempty" tf:"publish_url,omitempty"`

	Published *bool `json:"published,omitempty" tf:"published,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type PublicationParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// +kubebuilder:validation:Optional
	Published *bool `json:"published,omitempty" tf:"published,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SubscriptionObservation struct {
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	AutomaticSync *bool `json:"automaticSync,omitempty" tf:"automatic_sync,omitempty"`

	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	SubscriptionURL *string `json:"subscriptionUrl,omitempty" tf:"subscription_url,omitempty"`

	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type SubscriptionParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationMethod *string `json:"authenticationMethod,omitempty" tf:"authentication_method,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticSync *bool `json:"automaticSync,omitempty" tf:"automatic_sync,omitempty"`

	// +kubebuilder:validation:Optional
	OnDemand *bool `json:"onDemand,omitempty" tf:"on_demand,omitempty"`

	// +kubebuilder:validation:Optional
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionURL *string `json:"subscriptionUrl,omitempty" tf:"subscription_url,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type VSphereContentLibraryObservation struct {

	// Optional description of the content library.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the content library.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Publication configuration for content library.
	Publication []PublicationObservation `json:"publication,omitempty" tf:"publication,omitempty"`

	// The name of the content library.
	StorageBacking []*string `json:"storageBacking,omitempty" tf:"storage_backing,omitempty"`

	// Publication configuration for content library.
	Subscription []SubscriptionObservation `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

type VSphereContentLibraryParameters struct {

	// Optional description of the content library.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the content library.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Publication configuration for content library.
	// +kubebuilder:validation:Optional
	Publication []PublicationParameters `json:"publication,omitempty" tf:"publication,omitempty"`

	// The name of the content library.
	// +kubebuilder:validation:Optional
	StorageBacking []*string `json:"storageBacking,omitempty" tf:"storage_backing,omitempty"`

	// Publication configuration for content library.
	// +kubebuilder:validation:Optional
	Subscription []SubscriptionParameters `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

// VSphereContentLibrarySpec defines the desired state of VSphereContentLibrary
type VSphereContentLibrarySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereContentLibraryParameters `json:"forProvider"`
}

// VSphereContentLibraryStatus defines the observed state of VSphereContentLibrary.
type VSphereContentLibraryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereContentLibraryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereContentLibrary is the Schema for the VSphereContentLibrarys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereContentLibrary struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.storageBacking)",message="storageBacking is a required parameter"
	Spec   VSphereContentLibrarySpec   `json:"spec"`
	Status VSphereContentLibraryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereContentLibraryList contains a list of VSphereContentLibrarys
type VSphereContentLibraryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereContentLibrary `json:"items"`
}

// Repository type metadata.
var (
	VSphereContentLibrary_Kind             = "VSphereContentLibrary"
	VSphereContentLibrary_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereContentLibrary_Kind}.String()
	VSphereContentLibrary_KindAPIVersion   = VSphereContentLibrary_Kind + "." + CRDGroupVersion.String()
	VSphereContentLibrary_GroupVersionKind = CRDGroupVersion.WithKind(VSphereContentLibrary_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereContentLibrary{}, &VSphereContentLibraryList{})
}
