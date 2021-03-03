// +build !ignore_autogenerated

/*
Copyright 2019 The Authors.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborRobotAccount) DeepCopyInto(out *HarborRobotAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborRobotAccount.
func (in *HarborRobotAccount) DeepCopy() *HarborRobotAccount {
	if in == nil {
		return nil
	}
	out := new(HarborRobotAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HarborRobotAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborRobotAccountList) DeepCopyInto(out *HarborRobotAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HarborRobotAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborRobotAccountList.
func (in *HarborRobotAccountList) DeepCopy() *HarborRobotAccountList {
	if in == nil {
		return nil
	}
	out := new(HarborRobotAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HarborRobotAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborRobotAccountSpec) DeepCopyInto(out *HarborRobotAccountSpec) {
	*out = *in
	out.Credential = in.Credential
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborRobotAccountSpec.
func (in *HarborRobotAccountSpec) DeepCopy() *HarborRobotAccountSpec {
	if in == nil {
		return nil
	}
	out := new(HarborRobotAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborRobotAccountStatus) DeepCopyInto(out *HarborRobotAccountStatus) {
	*out = *in
	in.RefreshTime.DeepCopyInto(&out.RefreshTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RobotAccountStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborRobotAccountStatus.
func (in *HarborRobotAccountStatus) DeepCopy() *HarborRobotAccountStatus {
	if in == nil {
		return nil
	}
	out := new(HarborRobotAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborSync) DeepCopyInto(out *HarborSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborSync.
func (in *HarborSync) DeepCopy() *HarborSync {
	if in == nil {
		return nil
	}
	out := new(HarborSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HarborSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborSyncList) DeepCopyInto(out *HarborSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HarborSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborSyncList.
func (in *HarborSyncList) DeepCopy() *HarborSyncList {
	if in == nil {
		return nil
	}
	out := new(HarborSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HarborSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborSyncSpec) DeepCopyInto(out *HarborSyncSpec) {
	*out = *in
	if in.Mapping != nil {
		in, out := &in.Mapping, &out.Mapping
		*out = make([]ProjectMapping, len(*in))
		copy(*out, *in)
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = make([]WebhookConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborSyncSpec.
func (in *HarborSyncSpec) DeepCopy() *HarborSyncSpec {
	if in == nil {
		return nil
	}
	out := new(HarborSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborSyncStatus) DeepCopyInto(out *HarborSyncStatus) {
	*out = *in
	if in.ProjectList != nil {
		in, out := &in.ProjectList, &out.ProjectList
		*out = make([]ProjectStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]HarborSyncStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborSyncStatus.
func (in *HarborSyncStatus) DeepCopy() *HarborSyncStatus {
	if in == nil {
		return nil
	}
	out := new(HarborSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HarborSyncStatusCondition) DeepCopyInto(out *HarborSyncStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HarborSyncStatusCondition.
func (in *HarborSyncStatusCondition) DeepCopy() *HarborSyncStatusCondition {
	if in == nil {
		return nil
	}
	out := new(HarborSyncStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectMapping) DeepCopyInto(out *ProjectMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectMapping.
func (in *ProjectMapping) DeepCopy() *ProjectMapping {
	if in == nil {
		return nil
	}
	out := new(ProjectMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.LastRobotReconciliation.DeepCopyInto(&out.LastRobotReconciliation)
	if in.ManagedNamespaces != nil {
		in, out := &in.ManagedNamespaces, &out.ManagedNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotAccountCredential) DeepCopyInto(out *RobotAccountCredential) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotAccountCredential.
func (in *RobotAccountCredential) DeepCopy() *RobotAccountCredential {
	if in == nil {
		return nil
	}
	out := new(RobotAccountCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotAccountStatusCondition) DeepCopyInto(out *RobotAccountStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotAccountStatusCondition.
func (in *RobotAccountStatusCondition) DeepCopy() *RobotAccountStatusCondition {
	if in == nil {
		return nil
	}
	out := new(RobotAccountStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookUpdatePayload) DeepCopyInto(out *WebhookUpdatePayload) {
	*out = *in
	out.Credentials = in.Credentials
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookUpdatePayload.
func (in *WebhookUpdatePayload) DeepCopy() *WebhookUpdatePayload {
	if in == nil {
		return nil
	}
	out := new(WebhookUpdatePayload)
	in.DeepCopyInto(out)
	return out
}
