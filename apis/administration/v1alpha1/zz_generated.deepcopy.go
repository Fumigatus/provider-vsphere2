//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicense) DeepCopyInto(out *VSphereLicense) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicense.
func (in *VSphereLicense) DeepCopy() *VSphereLicense {
	if in == nil {
		return nil
	}
	out := new(VSphereLicense)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereLicense) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicenseList) DeepCopyInto(out *VSphereLicenseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VSphereLicense, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicenseList.
func (in *VSphereLicenseList) DeepCopy() *VSphereLicenseList {
	if in == nil {
		return nil
	}
	out := new(VSphereLicenseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VSphereLicenseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicenseObservation) DeepCopyInto(out *VSphereLicenseObservation) {
	*out = *in
	if in.EditionKey != nil {
		in, out := &in.EditionKey, &out.EditionKey
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LicenseKey != nil {
		in, out := &in.LicenseKey, &out.LicenseKey
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Total != nil {
		in, out := &in.Total, &out.Total
		*out = new(float64)
		**out = **in
	}
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicenseObservation.
func (in *VSphereLicenseObservation) DeepCopy() *VSphereLicenseObservation {
	if in == nil {
		return nil
	}
	out := new(VSphereLicenseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicenseParameters) DeepCopyInto(out *VSphereLicenseParameters) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LicenseKey != nil {
		in, out := &in.LicenseKey, &out.LicenseKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicenseParameters.
func (in *VSphereLicenseParameters) DeepCopy() *VSphereLicenseParameters {
	if in == nil {
		return nil
	}
	out := new(VSphereLicenseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicenseSpec) DeepCopyInto(out *VSphereLicenseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicenseSpec.
func (in *VSphereLicenseSpec) DeepCopy() *VSphereLicenseSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereLicenseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereLicenseStatus) DeepCopyInto(out *VSphereLicenseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereLicenseStatus.
func (in *VSphereLicenseStatus) DeepCopy() *VSphereLicenseStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereLicenseStatus)
	in.DeepCopyInto(out)
	return out
}
