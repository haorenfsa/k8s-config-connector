//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowAgent) DeepCopyInto(out *DialogflowAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowAgent.
func (in *DialogflowAgent) DeepCopy() *DialogflowAgent {
	if in == nil {
		return nil
	}
	out := new(DialogflowAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowAgentList) DeepCopyInto(out *DialogflowAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DialogflowAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowAgentList.
func (in *DialogflowAgentList) DeepCopy() *DialogflowAgentList {
	if in == nil {
		return nil
	}
	out := new(DialogflowAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowAgentSpec) DeepCopyInto(out *DialogflowAgentSpec) {
	*out = *in
	if in.ApiVersion != nil {
		in, out := &in.ApiVersion, &out.ApiVersion
		*out = new(string)
		**out = **in
	}
	if in.AvatarUri != nil {
		in, out := &in.AvatarUri, &out.AvatarUri
		*out = new(string)
		**out = **in
	}
	if in.ClassificationThreshold != nil {
		in, out := &in.ClassificationThreshold, &out.ClassificationThreshold
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EnableLogging != nil {
		in, out := &in.EnableLogging, &out.EnableLogging
		*out = new(bool)
		**out = **in
	}
	if in.MatchMode != nil {
		in, out := &in.MatchMode, &out.MatchMode
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.SupportedLanguageCodes != nil {
		in, out := &in.SupportedLanguageCodes, &out.SupportedLanguageCodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowAgentSpec.
func (in *DialogflowAgentSpec) DeepCopy() *DialogflowAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DialogflowAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowAgentStatus) DeepCopyInto(out *DialogflowAgentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.AvatarUriBackend != nil {
		in, out := &in.AvatarUriBackend, &out.AvatarUriBackend
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowAgentStatus.
func (in *DialogflowAgentStatus) DeepCopy() *DialogflowAgentStatus {
	if in == nil {
		return nil
	}
	out := new(DialogflowAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowEntityType) DeepCopyInto(out *DialogflowEntityType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowEntityType.
func (in *DialogflowEntityType) DeepCopy() *DialogflowEntityType {
	if in == nil {
		return nil
	}
	out := new(DialogflowEntityType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowEntityType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowEntityTypeList) DeepCopyInto(out *DialogflowEntityTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DialogflowEntityType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowEntityTypeList.
func (in *DialogflowEntityTypeList) DeepCopy() *DialogflowEntityTypeList {
	if in == nil {
		return nil
	}
	out := new(DialogflowEntityTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowEntityTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowEntityTypeSpec) DeepCopyInto(out *DialogflowEntityTypeSpec) {
	*out = *in
	if in.EnableFuzzyExtraction != nil {
		in, out := &in.EnableFuzzyExtraction, &out.EnableFuzzyExtraction
		*out = new(bool)
		**out = **in
	}
	if in.Entities != nil {
		in, out := &in.Entities, &out.Entities
		*out = make([]EntitytypeEntities, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowEntityTypeSpec.
func (in *DialogflowEntityTypeSpec) DeepCopy() *DialogflowEntityTypeSpec {
	if in == nil {
		return nil
	}
	out := new(DialogflowEntityTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowEntityTypeStatus) DeepCopyInto(out *DialogflowEntityTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowEntityTypeStatus.
func (in *DialogflowEntityTypeStatus) DeepCopy() *DialogflowEntityTypeStatus {
	if in == nil {
		return nil
	}
	out := new(DialogflowEntityTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowFulfillment) DeepCopyInto(out *DialogflowFulfillment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowFulfillment.
func (in *DialogflowFulfillment) DeepCopy() *DialogflowFulfillment {
	if in == nil {
		return nil
	}
	out := new(DialogflowFulfillment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowFulfillment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowFulfillmentList) DeepCopyInto(out *DialogflowFulfillmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DialogflowFulfillment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowFulfillmentList.
func (in *DialogflowFulfillmentList) DeepCopy() *DialogflowFulfillmentList {
	if in == nil {
		return nil
	}
	out := new(DialogflowFulfillmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowFulfillmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowFulfillmentSpec) DeepCopyInto(out *DialogflowFulfillmentSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]FulfillmentFeatures, len(*in))
		copy(*out, *in)
	}
	if in.GenericWebService != nil {
		in, out := &in.GenericWebService, &out.GenericWebService
		*out = new(FulfillmentGenericWebService)
		(*in).DeepCopyInto(*out)
	}
	out.ProjectRef = in.ProjectRef
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowFulfillmentSpec.
func (in *DialogflowFulfillmentSpec) DeepCopy() *DialogflowFulfillmentSpec {
	if in == nil {
		return nil
	}
	out := new(DialogflowFulfillmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowFulfillmentStatus) DeepCopyInto(out *DialogflowFulfillmentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowFulfillmentStatus.
func (in *DialogflowFulfillmentStatus) DeepCopy() *DialogflowFulfillmentStatus {
	if in == nil {
		return nil
	}
	out := new(DialogflowFulfillmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowIntent) DeepCopyInto(out *DialogflowIntent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowIntent.
func (in *DialogflowIntent) DeepCopy() *DialogflowIntent {
	if in == nil {
		return nil
	}
	out := new(DialogflowIntent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowIntent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowIntentList) DeepCopyInto(out *DialogflowIntentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DialogflowIntent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowIntentList.
func (in *DialogflowIntentList) DeepCopy() *DialogflowIntentList {
	if in == nil {
		return nil
	}
	out := new(DialogflowIntentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DialogflowIntentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowIntentSpec) DeepCopyInto(out *DialogflowIntentSpec) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.DefaultResponsePlatforms != nil {
		in, out := &in.DefaultResponsePlatforms, &out.DefaultResponsePlatforms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InputContextNames != nil {
		in, out := &in.InputContextNames, &out.InputContextNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IsFallback != nil {
		in, out := &in.IsFallback, &out.IsFallback
		*out = new(bool)
		**out = **in
	}
	if in.MlDisabled != nil {
		in, out := &in.MlDisabled, &out.MlDisabled
		*out = new(bool)
		**out = **in
	}
	if in.ParentFollowupIntentName != nil {
		in, out := &in.ParentFollowupIntentName, &out.ParentFollowupIntentName
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int)
		**out = **in
	}
	out.ProjectRef = in.ProjectRef
	if in.ResetContexts != nil {
		in, out := &in.ResetContexts, &out.ResetContexts
		*out = new(bool)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.WebhookState != nil {
		in, out := &in.WebhookState, &out.WebhookState
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowIntentSpec.
func (in *DialogflowIntentSpec) DeepCopy() *DialogflowIntentSpec {
	if in == nil {
		return nil
	}
	out := new(DialogflowIntentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DialogflowIntentStatus) DeepCopyInto(out *DialogflowIntentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.FollowupIntentInfo != nil {
		in, out := &in.FollowupIntentInfo, &out.FollowupIntentInfo
		*out = make([]IntentFollowupIntentInfoStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int)
		**out = **in
	}
	if in.RootFollowupIntentName != nil {
		in, out := &in.RootFollowupIntentName, &out.RootFollowupIntentName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DialogflowIntentStatus.
func (in *DialogflowIntentStatus) DeepCopy() *DialogflowIntentStatus {
	if in == nil {
		return nil
	}
	out := new(DialogflowIntentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntitytypeEntities) DeepCopyInto(out *EntitytypeEntities) {
	*out = *in
	if in.Synonyms != nil {
		in, out := &in.Synonyms, &out.Synonyms
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntitytypeEntities.
func (in *EntitytypeEntities) DeepCopy() *EntitytypeEntities {
	if in == nil {
		return nil
	}
	out := new(EntitytypeEntities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulfillmentFeatures) DeepCopyInto(out *FulfillmentFeatures) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulfillmentFeatures.
func (in *FulfillmentFeatures) DeepCopy() *FulfillmentFeatures {
	if in == nil {
		return nil
	}
	out := new(FulfillmentFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FulfillmentGenericWebService) DeepCopyInto(out *FulfillmentGenericWebService) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FulfillmentGenericWebService.
func (in *FulfillmentGenericWebService) DeepCopy() *FulfillmentGenericWebService {
	if in == nil {
		return nil
	}
	out := new(FulfillmentGenericWebService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IntentFollowupIntentInfoStatus) DeepCopyInto(out *IntentFollowupIntentInfoStatus) {
	*out = *in
	if in.FollowupIntentName != nil {
		in, out := &in.FollowupIntentName, &out.FollowupIntentName
		*out = new(string)
		**out = **in
	}
	if in.ParentFollowupIntentName != nil {
		in, out := &in.ParentFollowupIntentName, &out.ParentFollowupIntentName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IntentFollowupIntentInfoStatus.
func (in *IntentFollowupIntentInfoStatus) DeepCopy() *IntentFollowupIntentInfoStatus {
	if in == nil {
		return nil
	}
	out := new(IntentFollowupIntentInfoStatus)
	in.DeepCopyInto(out)
	return out
}
