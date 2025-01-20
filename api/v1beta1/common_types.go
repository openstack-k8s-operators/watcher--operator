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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/util"
	corev1 "k8s.io/api/core/v1"
)

// Container image fall-back defaults
const (
	WatcherAPIContainerImage            = "quay.io/podified-master-centos9/openstack-watcher-api:current-podified"
	WatcherDecisionEngineContainerImage = "quay.io/podified-master-centos9/openstack-watcher-decision-engine:current-podified"
	WatcherApplierContainerImage        = "quay.io/podified-master-centos9/openstack-watcher-applier:current-podified"
)

// WatcherCommon defines a spec based reusable for all the CRDs
type WatcherCommon struct {

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=watcher
	// ServiceUser - optional username used for this service to register in keystone
	ServiceUser string `json:"serviceUser"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default={service: WatcherPassword,}
	// PasswordSelectors - Selectors to identify the ServiceUser password from the Secret
	PasswordSelectors PasswordSelector `json:"passwordSelectors"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=memcached
	// MemcachedInstance is the name of the Memcached CR that all watcher service will use.
	MemcachedInstance string `json:"memcachedInstance"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=false
	// PreserveJobs - do not delete jobs after they finished e.g. to check logs
	PreserveJobs bool `json:"preserveJobs"`

	// +kubebuilder:validation:Optional
	// NodeSelector to target subset of worker nodes running this component. Setting here overrides
	// any global NodeSelector settings within the Watcher CR.
	NodeSelector *map[string]string `json:"nodeSelector,omitempty"`
}

// WatcherTemplate defines the fields used in the top level CR
type WatcherTemplate struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	WatcherCommon `json:",inline"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default=rabbitmq
	// RabbitMQ instance name
	// Needed to request a transportURL that is created and used in Watcher
	RabbitMqClusterName *string `json:"rabbitMqClusterName"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=osp-secret
	// Secret containing all passwords / keys needed
	Secret string `json:"secret"`

	// +kubebuilder:validation:Required
	// MariaDB instance name
	// Required to use the mariadb-operator instance to create the DB and user
	DatabaseInstance *string `json:"databaseInstance"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=watcher
	// DatabaseAccount - MariaDBAccount CR name used for watcher DB, defaults to watcher
	DatabaseAccount string `json:"databaseAccount"`

	// +kubebuilder:validation:Required
	// +kubebuilder:default={replicas:1}
	// APIServiceTemplate - define the watcher-api service
	APIServiceTemplate WatcherAPITemplate `json:"apiServiceTemplate"`
}

// PasswordSelector to identify the DB and AdminUser password from the Secret
type PasswordSelector struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default="WatcherPassword"
	// Service - Selector to get the watcher service user password from the Secret
	Service string `json:"service"`
}

// WatcherSubCrsCommon
type WatcherSubCrsCommon struct {
	// +kubebuilder:validation:Optional
	// The service specific Container Image URL (will be set to environmental default if empty)
	ContainerImage string `json:"containerImage"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=0
	// Replicas of Watcher service to run
	Replicas *int32 `json:"replicas"`

	// +kubebuilder:validation:Optional
	// Resources - Compute Resources required by this service (Limits/Requests).
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Required
	// ServiceAccount - service account name used internally to provide
	// Watcher services the default SA name
	ServiceAccount string `json:"serviceAccount"`
}

// WatcherSubCrsTemplate define de common part of the input parameters specified by the user to
// create a 2nd CR via higher level CRDs.
type WatcherSubCrsTemplate struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=0
	// Replicas of Watcher service to run
	Replicas *int32 `json:"replicas"`

	// +kubebuilder:validation:Optional
	// Resources - Compute Resources required by this service (Limits/Requests).
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// NodeSelector to target subset of worker nodes running this component. Setting here overrides
	// any global NodeSelector settings within the Watcher CR.
	NodeSelector *map[string]string `json:"nodeSelector,omitempty"`
}

// MetalLBConfig to configure the MetalLB loadbalancer service
type MetalLBConfig struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// IPAddressPool expose VIP via MetalLB on the IPAddressPool
	IPAddressPool string `json:"ipAddressPool"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// SharedIP if true, VIP/VIPs get shared with multiple services
	SharedIP bool `json:"sharedIP"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	// SharedIPKey specifies the sharing key which gets set as the annotation on the LoadBalancer service.
	// Services which share the same VIP must have the same SharedIPKey. Defaults to the IPAddressPool if
	// SharedIP is true, but no SharedIPKey specified.
	SharedIPKey string `json:"sharedIPKey"`

	// +kubebuilder:validation:Optional
	// LoadBalancerIPs, request given IPs from the pool if available. Using a list to allow dual stack (IPv4/IPv6) support
	LoadBalancerIPs []string `json:"loadBalancerIPs"`
}

type WatcherImages struct {
	// +kubebuilder:validation:Required
	// APIContainerImageURL
	APIContainerImageURL string `json:"apiContainerImageURL"`

	// +kubebuilder:validation:Required
	// DecisionEngineContainerImageURL
	DecisionEngineContainerImageURL string `json:"decisionengineContainerImageURL"`

	// +kubebuilder:validation:Required
	// ApplierContainerImageURL
	ApplierContainerImageURL string `json:"applierContainerImageURL"`
}

func (r *WatcherImages) Default(defaults WatcherDefaults) {
	if r.APIContainerImageURL == "" {
		r.APIContainerImageURL = defaults.APIContainerImageURL
	}
	if r.DecisionEngineContainerImageURL == "" {
		r.DecisionEngineContainerImageURL = defaults.DecisionEngineContainerImageURL
	}
	if r.ApplierContainerImageURL == "" {
		r.ApplierContainerImageURL = defaults.ApplierContainerImageURL
	}
}

// SetupDefaults - initializes any CRD field defaults based on environment variables (the defaulting mechanism itself is implemented via webhooks)

func SetupDefaults() {
	// Acquire environmental defaults and initialize Watcher defaults with them
	watcherDefaults := WatcherDefaults{
		APIContainerImageURL:            util.GetEnvVar("WATCHER_API_IMAGE_URL_DEFAULT", WatcherAPIContainerImage),
		ApplierContainerImageURL:        util.GetEnvVar("WATCHER_APPLIER_IMAGE_URL_DEFAULT", WatcherApplierContainerImage),
		DecisionEngineContainerImageURL: util.GetEnvVar("WATCHER_DECISION_ENGINE_IMAGE_URL_DEFAULT", WatcherDecisionEngineContainerImage),
	}
	SetupWatcherDefaults(watcherDefaults)
}
