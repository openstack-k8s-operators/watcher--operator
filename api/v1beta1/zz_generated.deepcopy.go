//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Watcher) DeepCopyInto(out *Watcher) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Watcher.
func (in *Watcher) DeepCopy() *Watcher {
	if in == nil {
		return nil
	}
	out := new(Watcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Watcher) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherAPI) DeepCopyInto(out *WatcherAPI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherAPI.
func (in *WatcherAPI) DeepCopy() *WatcherAPI {
	if in == nil {
		return nil
	}
	out := new(WatcherAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherAPI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherAPIList) DeepCopyInto(out *WatcherAPIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatcherAPI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherAPIList.
func (in *WatcherAPIList) DeepCopy() *WatcherAPIList {
	if in == nil {
		return nil
	}
	out := new(WatcherAPIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherAPIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherAPISpec) DeepCopyInto(out *WatcherAPISpec) {
	*out = *in
	out.WatcherCommon = in.WatcherCommon
	out.WatcherSubCrsCommon = in.WatcherSubCrsCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherAPISpec.
func (in *WatcherAPISpec) DeepCopy() *WatcherAPISpec {
	if in == nil {
		return nil
	}
	out := new(WatcherAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherAPIStatus) DeepCopyInto(out *WatcherAPIStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherAPIStatus.
func (in *WatcherAPIStatus) DeepCopy() *WatcherAPIStatus {
	if in == nil {
		return nil
	}
	out := new(WatcherAPIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherApplier) DeepCopyInto(out *WatcherApplier) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherApplier.
func (in *WatcherApplier) DeepCopy() *WatcherApplier {
	if in == nil {
		return nil
	}
	out := new(WatcherApplier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherApplier) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherApplierList) DeepCopyInto(out *WatcherApplierList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatcherApplier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherApplierList.
func (in *WatcherApplierList) DeepCopy() *WatcherApplierList {
	if in == nil {
		return nil
	}
	out := new(WatcherApplierList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherApplierList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherApplierSpec) DeepCopyInto(out *WatcherApplierSpec) {
	*out = *in
	out.WatcherSubCrsCommon = in.WatcherSubCrsCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherApplierSpec.
func (in *WatcherApplierSpec) DeepCopy() *WatcherApplierSpec {
	if in == nil {
		return nil
	}
	out := new(WatcherApplierSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherApplierStatus) DeepCopyInto(out *WatcherApplierStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherApplierStatus.
func (in *WatcherApplierStatus) DeepCopy() *WatcherApplierStatus {
	if in == nil {
		return nil
	}
	out := new(WatcherApplierStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherCommon) DeepCopyInto(out *WatcherCommon) {
	*out = *in
	out.PasswordSelectors = in.PasswordSelectors
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherCommon.
func (in *WatcherCommon) DeepCopy() *WatcherCommon {
	if in == nil {
		return nil
	}
	out := new(WatcherCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherDecisionEngine) DeepCopyInto(out *WatcherDecisionEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherDecisionEngine.
func (in *WatcherDecisionEngine) DeepCopy() *WatcherDecisionEngine {
	if in == nil {
		return nil
	}
	out := new(WatcherDecisionEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherDecisionEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherDecisionEngineList) DeepCopyInto(out *WatcherDecisionEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatcherDecisionEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherDecisionEngineList.
func (in *WatcherDecisionEngineList) DeepCopy() *WatcherDecisionEngineList {
	if in == nil {
		return nil
	}
	out := new(WatcherDecisionEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherDecisionEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherDecisionEngineSpec) DeepCopyInto(out *WatcherDecisionEngineSpec) {
	*out = *in
	out.WatcherSubCrsCommon = in.WatcherSubCrsCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherDecisionEngineSpec.
func (in *WatcherDecisionEngineSpec) DeepCopy() *WatcherDecisionEngineSpec {
	if in == nil {
		return nil
	}
	out := new(WatcherDecisionEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherDecisionEngineStatus) DeepCopyInto(out *WatcherDecisionEngineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherDecisionEngineStatus.
func (in *WatcherDecisionEngineStatus) DeepCopy() *WatcherDecisionEngineStatus {
	if in == nil {
		return nil
	}
	out := new(WatcherDecisionEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherDefaults) DeepCopyInto(out *WatcherDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherDefaults.
func (in *WatcherDefaults) DeepCopy() *WatcherDefaults {
	if in == nil {
		return nil
	}
	out := new(WatcherDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherImages) DeepCopyInto(out *WatcherImages) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherImages.
func (in *WatcherImages) DeepCopy() *WatcherImages {
	if in == nil {
		return nil
	}
	out := new(WatcherImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherList) DeepCopyInto(out *WatcherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Watcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherList.
func (in *WatcherList) DeepCopy() *WatcherList {
	if in == nil {
		return nil
	}
	out := new(WatcherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatcherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherSpec) DeepCopyInto(out *WatcherSpec) {
	*out = *in
	out.WatcherTemplate = in.WatcherTemplate
	out.WatcherImages = in.WatcherImages
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherSpec.
func (in *WatcherSpec) DeepCopy() *WatcherSpec {
	if in == nil {
		return nil
	}
	out := new(WatcherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherStatus) DeepCopyInto(out *WatcherStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherStatus.
func (in *WatcherStatus) DeepCopy() *WatcherStatus {
	if in == nil {
		return nil
	}
	out := new(WatcherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherSubCrsCommon) DeepCopyInto(out *WatcherSubCrsCommon) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherSubCrsCommon.
func (in *WatcherSubCrsCommon) DeepCopy() *WatcherSubCrsCommon {
	if in == nil {
		return nil
	}
	out := new(WatcherSubCrsCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatcherTemplate) DeepCopyInto(out *WatcherTemplate) {
	*out = *in
	out.WatcherCommon = in.WatcherCommon
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatcherTemplate.
func (in *WatcherTemplate) DeepCopy() *WatcherTemplate {
	if in == nil {
		return nil
	}
	out := new(WatcherTemplate)
	in.DeepCopyInto(out)
	return out
}
