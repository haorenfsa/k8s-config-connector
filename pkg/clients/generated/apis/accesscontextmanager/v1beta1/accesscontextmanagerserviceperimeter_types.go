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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceperimeterEgressFrom struct {
	// +optional
	Identities []ServiceperimeterIdentities `json:"identities,omitempty"`

	/* Specifies the type of identities that are allowed access to outside the
	perimeter. If left unspecified, then members of 'identities' field will
	be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]. */
	// +optional
	IdentityType *string `json:"identityType,omitempty"`
}

type ServiceperimeterEgressPolicies struct {
	/* Defines conditions on the source of a request causing this 'EgressPolicy' to apply. */
	// +optional
	EgressFrom *ServiceperimeterEgressFrom `json:"egressFrom,omitempty"`

	/* Defines the conditions on the 'ApiOperation' and destination resources that
	cause this 'EgressPolicy' to apply. */
	// +optional
	EgressTo *ServiceperimeterEgressTo `json:"egressTo,omitempty"`
}

type ServiceperimeterEgressTo struct {
	/* A list of external resources that are allowed to be accessed. A request
	matches if it contains an external resource in this list (Example:
	s3://bucket/path). Currently '*' is not allowed. */
	// +optional
	ExternalResources []string `json:"externalResources,omitempty"`

	/* A list of 'ApiOperations' that this egress rule applies to. A request matches
	if it contains an operation/service in this list. */
	// +optional
	Operations []ServiceperimeterOperations `json:"operations,omitempty"`

	// +optional
	Resources []ServiceperimeterResources `json:"resources,omitempty"`
}

type ServiceperimeterIdentities struct {
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	// +optional
	User *string `json:"user,omitempty"`
}

type ServiceperimeterIngressFrom struct {
	// +optional
	Identities []ServiceperimeterIdentities `json:"identities,omitempty"`

	/* Specifies the type of identities that are allowed access from outside the
	perimeter. If left unspecified, then members of 'identities' field will be
	allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]. */
	// +optional
	IdentityType *string `json:"identityType,omitempty"`

	/* Sources that this 'IngressPolicy' authorizes access from. */
	// +optional
	Sources []ServiceperimeterSources `json:"sources,omitempty"`
}

type ServiceperimeterIngressPolicies struct {
	/* Defines the conditions on the source of a request causing this 'IngressPolicy'
	to apply. */
	// +optional
	IngressFrom *ServiceperimeterIngressFrom `json:"ingressFrom,omitempty"`

	/* Defines the conditions on the 'ApiOperation' and request destination that cause
	this 'IngressPolicy' to apply. */
	// +optional
	IngressTo *ServiceperimeterIngressTo `json:"ingressTo,omitempty"`
}

type ServiceperimeterIngressTo struct {
	/* A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'
	are allowed to perform in this 'ServicePerimeter'. */
	// +optional
	Operations []ServiceperimeterOperations `json:"operations,omitempty"`

	// +optional
	Resources []ServiceperimeterResources `json:"resources,omitempty"`
}

type ServiceperimeterMethodSelectors struct {
	/* Value for method should be a valid method name for the corresponding
	serviceName in 'ApiOperation'. If '*' used as value for 'method', then
	ALL methods and permissions are allowed. */
	// +optional
	Method *string `json:"method,omitempty"`

	/* Value for permission should be a valid Cloud IAM permission for the
	corresponding 'serviceName' in 'ApiOperation'. */
	// +optional
	Permission *string `json:"permission,omitempty"`
}

type ServiceperimeterOperations struct {
	/* API methods or permissions to allow. Method or permission must belong to
	the service specified by serviceName field. A single 'MethodSelector' entry
	with '*' specified for the method field will allow all methods AND
	permissions for the service specified in 'serviceName'. */
	// +optional
	MethodSelectors []ServiceperimeterMethodSelectors `json:"methodSelectors,omitempty"`

	/* The name of the API whose methods or permissions the 'IngressPolicy' or
	'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName'
	field set to '*' will allow all methods AND permissions for all services. */
	// +optional
	ServiceName *string `json:"serviceName,omitempty"`
}

type ServiceperimeterResources struct {
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`
}

type ServiceperimeterSources struct {
	/* An AccessLevel resource name that allow resources within the
	ServicePerimeters to be accessed from the internet. AccessLevels
	listed must be in the same policy as this ServicePerimeter.
	Referencing a nonexistent AccessLevel will cause an error. If no
	AccessLevel names are listed, resources within the perimeter can
	only be accessed via Google Cloud calls with request origins within
	the perimeter. */
	// +optional
	AccessLevelRef *v1alpha1.ResourceRef `json:"accessLevelRef,omitempty"`

	/* (Optional) A Google Cloud resource that is allowed to ingress the
	perimeter. Requests from these resources will be allowed to access
	perimeter data. Currently only projects are allowed. Format
	"projects/{project_number}" The project may be in any Google Cloud
	organization, not just the organization that the perimeter is defined in. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`
}

type ServiceperimeterSpec struct {
	// +optional
	AccessLevels []v1alpha1.ResourceRef `json:"accessLevels,omitempty"`

	/* List of EgressPolicies to apply to the perimeter. A perimeter may
	have multiple EgressPolicies, each of which is evaluated separately.
	Access is granted if any EgressPolicy grants it. Must be empty for
	a perimeter bridge. */
	// +optional
	EgressPolicies []ServiceperimeterEgressPolicies `json:"egressPolicies,omitempty"`

	/* List of 'IngressPolicies' to apply to the perimeter. A perimeter may
	have multiple 'IngressPolicies', each of which is evaluated
	separately. Access is granted if any 'Ingress Policy' grants it.
	Must be empty for a perimeter bridge. */
	// +optional
	IngressPolicies []ServiceperimeterIngressPolicies `json:"ingressPolicies,omitempty"`

	// +optional
	Resources []ServiceperimeterResources `json:"resources,omitempty"`

	/* GCP services that are subject to the Service Perimeter
	restrictions. Must contain a list of services. For example, if
	'storage.googleapis.com' is specified, access to the storage
	buckets inside the perimeter must meet the perimeter's access
	restrictions. */
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	/* Specifies how APIs are allowed to communicate within the Service
	Perimeter. */
	// +optional
	VpcAccessibleServices *ServiceperimeterVpcAccessibleServices `json:"vpcAccessibleServices,omitempty"`
}

type ServiceperimeterStatus struct {
	// +optional
	AccessLevels []v1alpha1.ResourceRef `json:"accessLevels,omitempty"`

	/* List of EgressPolicies to apply to the perimeter. A perimeter may
	have multiple EgressPolicies, each of which is evaluated separately.
	Access is granted if any EgressPolicy grants it. Must be empty for
	a perimeter bridge. */
	// +optional
	EgressPolicies []ServiceperimeterEgressPolicies `json:"egressPolicies,omitempty"`

	/* List of 'IngressPolicies' to apply to the perimeter. A perimeter may
	have multiple 'IngressPolicies', each of which is evaluated
	separately. Access is granted if any 'Ingress Policy' grants it.
	Must be empty for a perimeter bridge. */
	// +optional
	IngressPolicies []ServiceperimeterIngressPolicies `json:"ingressPolicies,omitempty"`

	// +optional
	Resources []ServiceperimeterResources `json:"resources,omitempty"`

	/* GCP services that are subject to the Service Perimeter
	restrictions. Must contain a list of services. For example, if
	'storage.googleapis.com' is specified, access to the storage
	buckets inside the perimeter must meet the perimeter's access
	restrictions. */
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty"`

	/* Specifies how APIs are allowed to communicate within the Service
	Perimeter. */
	// +optional
	VpcAccessibleServices *ServiceperimeterVpcAccessibleServices `json:"vpcAccessibleServices,omitempty"`
}

type ServiceperimeterVpcAccessibleServices struct {
	/* The list of APIs usable within the Service Perimeter.
	Must be empty unless 'enableRestriction' is True. */
	// +optional
	AllowedServices []string `json:"allowedServices,omitempty"`

	/* Whether to restrict API calls within the Service Perimeter to the
	list of APIs specified in 'allowedServices'. */
	// +optional
	EnableRestriction *bool `json:"enableRestriction,omitempty"`
}

type AccessContextManagerServicePerimeterSpec struct {
	/* The AccessContextManagerAccessPolicy this
	AccessContextManagerServicePerimeter lives in. */
	AccessPolicyRef v1alpha1.ResourceRef `json:"accessPolicyRef"`

	/* Description of the ServicePerimeter and its use. Does not affect
	behavior. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Specifies the type of the Perimeter. There are two types: regular and
	bridge. Regular Service Perimeter contains resources, access levels,
	and restricted services. Every resource can be in at most
	ONE regular Service Perimeter.

	In addition to being in a regular service perimeter, a resource can also
	be in zero or more perimeter bridges. A perimeter bridge only contains
	resources. Cross project operations are permitted if all effected
	resources share some perimeter (whether bridge or regular). Perimeter
	Bridge does not contain access levels or services: those are governed
	entirely by the regular perimeter that resource is in.

	Perimeter Bridges are typically useful when building more complex
	topologies with many independent perimeters that need to share some data
	with a common perimeter, but should not be able to share data among
	themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"]. */
	// +optional
	PerimeterType *string `json:"perimeterType,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Proposed (or dry run) ServicePerimeter configuration.
	This configuration allows to specify and test ServicePerimeter configuration
	without enforcing actual access restrictions. Only allowed to be set when
	the 'useExplicitDryRunSpec' flag is set. */
	// +optional
	Spec *ServiceperimeterSpec `json:"spec,omitempty"`

	/* ServicePerimeter configuration. Specifies sets of resources,
	restricted services and access levels that determine
	perimeter content and boundaries. */
	// +optional
	Status *ServiceperimeterStatus `json:"status,omitempty"`

	/* Human readable title. Must be unique within the Policy. */
	Title string `json:"title"`

	/* Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	for all Service Perimeters, and that spec is identical to the status for those
	Service Perimeters. When this flag is set, it inhibits the generation of the
	implicit spec, thereby allowing the user to explicitly provide a
	configuration ("spec") to use in a dry-run version of the Service Perimeter.
	This allows the user to test changes to the enforced config ("status") without
	actually enforcing them. This testing is done through analyzing the differences
	between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	bet set to True if any of the fields in the spec are set to non-default values. */
	// +optional
	UseExplicitDryRunSpec *bool `json:"useExplicitDryRunSpec,omitempty"`
}

type AccessContextManagerServicePerimeterStatus struct {
	/* Conditions represent the latest available observations of the
	   AccessContextManagerServicePerimeter's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Time the AccessPolicy was created in UTC. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Time the AccessPolicy was updated in UTC. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerServicePerimeter is the Schema for the accesscontextmanager API
// +k8s:openapi-gen=true
type AccessContextManagerServicePerimeter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerServicePerimeterSpec   `json:"spec,omitempty"`
	Status AccessContextManagerServicePerimeterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerServicePerimeterList contains a list of AccessContextManagerServicePerimeter
type AccessContextManagerServicePerimeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerServicePerimeter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerServicePerimeter{}, &AccessContextManagerServicePerimeterList{})
}
