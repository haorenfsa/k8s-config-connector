// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AccessContextManagerServicePerimeterResourceSpec struct {
	PerimeterNameRef v1alpha1.ResourceRef `json:"perimeterNameRef"`

	/* Immutable. A GCP resource that is inside of the service perimeter.
	Currently only projects are allowed.
	Format: projects/{project_number}. */
	Resource string `json:"resource"`
}

type AccessContextManagerServicePerimeterResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   AccessContextManagerServicePerimeterResource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerServicePerimeterResource is the Schema for the accesscontextmanager API
// +k8s:openapi-gen=true
type AccessContextManagerServicePerimeterResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerServicePerimeterResourceSpec   `json:"spec,omitempty"`
	Status AccessContextManagerServicePerimeterResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerServicePerimeterResourceList contains a list of AccessContextManagerServicePerimeterResource
type AccessContextManagerServicePerimeterResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerServicePerimeterResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerServicePerimeterResource{}, &AccessContextManagerServicePerimeterResourceList{})
}
