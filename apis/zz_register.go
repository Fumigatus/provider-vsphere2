/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/ankasoftco/provider-vsphere/apis/administration/v1alpha1"
	v1alpha1hostandclustermanagement "github.com/ankasoftco/provider-vsphere/apis/hostandclustermanagement/v1alpha1"
	v1alpha1inventory "github.com/ankasoftco/provider-vsphere/apis/inventory/v1alpha1"
	v1alpha1networking "github.com/ankasoftco/provider-vsphere/apis/networking/v1alpha1"
	v1alpha1security "github.com/ankasoftco/provider-vsphere/apis/security/v1alpha1"
	v1alpha1storage "github.com/ankasoftco/provider-vsphere/apis/storage/v1alpha1"
	v1alpha1apis "github.com/ankasoftco/provider-vsphere/apis/v1alpha1"
	v1beta1 "github.com/ankasoftco/provider-vsphere/apis/v1beta1"
	v1alpha1virtualmachine "github.com/ankasoftco/provider-vsphere/apis/virtualmachine/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1hostandclustermanagement.SchemeBuilder.AddToScheme,
		v1alpha1inventory.SchemeBuilder.AddToScheme,
		v1alpha1networking.SchemeBuilder.AddToScheme,
		v1alpha1security.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1virtualmachine.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
