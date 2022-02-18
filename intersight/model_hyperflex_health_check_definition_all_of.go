/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// HyperflexHealthCheckDefinitionAllOf Definition of the list of properties defined in 'hyperflex.HealthCheckDefinition', excluding properties defined in parent classes.
type HyperflexHealthCheckDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Category that the health check belongs to.
	Category *string `json:"Category,omitempty"`
	// Static information detailing the common causes for the health check failure.
	CommonCauses *string `json:"CommonCauses,omitempty"`
	// Execution configuration fo the health check script.
	Configuration                *string                                `json:"Configuration,omitempty"`
	DefaultHealthCheckScriptInfo NullableHyperflexHealthCheckScriptInfo `json:"DefaultHealthCheckScriptInfo,omitempty"`
	// Description of the health check definition.
	Description            *string                          `json:"Description,omitempty"`
	HealthCheckScriptInfos []HyperflexHealthCheckScriptInfo `json:"HealthCheckScriptInfos,omitempty"`
	// Static information detailing the health impact of the health check failure.
	HealthImpact *string `json:"HealthImpact,omitempty"`
	// Internal name of the health check definition.
	InternalName *string `json:"InternalName,omitempty"`
	// Minimum HyperFlex version that the check is supported on. Defaults to version 3.0.1.
	MinimumHyperFlexVersion *string `json:"MinimumHyperFlexVersion,omitempty"`
	// Name of the health check definition.
	Name *string `json:"Name,omitempty"`
	// Static information containing additional reference for the health check.
	Reference *string `json:"Reference,omitempty"`
	// Static information detailing the possible remediation actions that can be taken to remedy the health check failure.
	Resolution *string `json:"Resolution,omitempty"`
	// Execution mode of the health script on the HyperFlex cluster. * `ON_DEMAND` - Execute the health check on-demand. * `SCHEDULED` - Execute the health check on a scheduled interval.
	ScriptExecutionMode *string `json:"ScriptExecutionMode,omitempty"`
	// Indicates if the script needs to be executed on HyperFlex compute nodes. | Typically, scripts are only executed on the storage Nodes.
	ScriptExecutionOnComputeNodes *bool `json:"ScriptExecutionOnComputeNodes,omitempty"`
	// Hypervisor type that the Health Check is supported on (All, if it is hypervisor agnostic). * `All` - The Health Check is hypervisor-agnostic. * `ESXi` - The Health Check is supported only on Vmware ESXi hypervisor of any version. * `` - The Health Check is supported only on Cisco HyperFlexAp platform. * `IWE` - The Health Check is supported only on Cisco IWE platform. * `HyperV` - The Health Check is supported only on Microsoft HyperV hypervisor.
	SupportedHypervisorType *string `json:"SupportedHypervisorType,omitempty"`
	// Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster. * `EXECUTE_ON_LEADER_NODE` - Execute the health check script only on the HyperFlex cluster's leader node. * `EXECUTE_ON_ALL_NODES` - Execute health check on all nodes and aggregate the results. * `EXECUTE_ON_ALL_NODES_AND_AGGREGATE` - Execute the health check on all Nodes and perform custom aggregation.
	TargetExecutionType *string `json:"TargetExecutionType,omitempty"`
	// Health check script execution timeout.
	Timeout                      *int64   `json:"Timeout,omitempty"`
	UnsupportedHyperFlexVersions []string `json:"UnsupportedHyperFlexVersions,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _HyperflexHealthCheckDefinitionAllOf HyperflexHealthCheckDefinitionAllOf

// NewHyperflexHealthCheckDefinitionAllOf instantiates a new HyperflexHealthCheckDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHealthCheckDefinitionAllOf(classId string, objectType string) *HyperflexHealthCheckDefinitionAllOf {
	this := HyperflexHealthCheckDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var minimumHyperFlexVersion string = "3"
	this.MinimumHyperFlexVersion = &minimumHyperFlexVersion
	var scriptExecutionMode string = "ON_DEMAND"
	this.ScriptExecutionMode = &scriptExecutionMode
	var targetExecutionType string = "EXECUTE_ON_LEADER_NODE"
	this.TargetExecutionType = &targetExecutionType
	return &this
}

// NewHyperflexHealthCheckDefinitionAllOfWithDefaults instantiates a new HyperflexHealthCheckDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHealthCheckDefinitionAllOfWithDefaults() *HyperflexHealthCheckDefinitionAllOf {
	this := HyperflexHealthCheckDefinitionAllOf{}
	var classId string = "hyperflex.HealthCheckDefinition"
	this.ClassId = classId
	var objectType string = "hyperflex.HealthCheckDefinition"
	this.ObjectType = objectType
	var minimumHyperFlexVersion string = "3"
	this.MinimumHyperFlexVersion = &minimumHyperFlexVersion
	var scriptExecutionMode string = "ON_DEMAND"
	this.ScriptExecutionMode = &scriptExecutionMode
	var targetExecutionType string = "EXECUTE_ON_LEADER_NODE"
	this.TargetExecutionType = &targetExecutionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHealthCheckDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHealthCheckDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHealthCheckDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHealthCheckDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetCommonCauses returns the CommonCauses field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetCommonCauses() string {
	if o == nil || o.CommonCauses == nil {
		var ret string
		return ret
	}
	return *o.CommonCauses
}

// GetCommonCausesOk returns a tuple with the CommonCauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetCommonCausesOk() (*string, bool) {
	if o == nil || o.CommonCauses == nil {
		return nil, false
	}
	return o.CommonCauses, true
}

// HasCommonCauses returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasCommonCauses() bool {
	if o != nil && o.CommonCauses != nil {
		return true
	}

	return false
}

// SetCommonCauses gets a reference to the given string and assigns it to the CommonCauses field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetCommonCauses(v string) {
	o.CommonCauses = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDefaultHealthCheckScriptInfo returns the DefaultHealthCheckScriptInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealthCheckDefinitionAllOf) GetDefaultHealthCheckScriptInfo() HyperflexHealthCheckScriptInfo {
	if o == nil || o.DefaultHealthCheckScriptInfo.Get() == nil {
		var ret HyperflexHealthCheckScriptInfo
		return ret
	}
	return *o.DefaultHealthCheckScriptInfo.Get()
}

// GetDefaultHealthCheckScriptInfoOk returns a tuple with the DefaultHealthCheckScriptInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealthCheckDefinitionAllOf) GetDefaultHealthCheckScriptInfoOk() (*HyperflexHealthCheckScriptInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultHealthCheckScriptInfo.Get(), o.DefaultHealthCheckScriptInfo.IsSet()
}

// HasDefaultHealthCheckScriptInfo returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasDefaultHealthCheckScriptInfo() bool {
	if o != nil && o.DefaultHealthCheckScriptInfo.IsSet() {
		return true
	}

	return false
}

// SetDefaultHealthCheckScriptInfo gets a reference to the given NullableHyperflexHealthCheckScriptInfo and assigns it to the DefaultHealthCheckScriptInfo field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetDefaultHealthCheckScriptInfo(v HyperflexHealthCheckScriptInfo) {
	o.DefaultHealthCheckScriptInfo.Set(&v)
}

// SetDefaultHealthCheckScriptInfoNil sets the value for DefaultHealthCheckScriptInfo to be an explicit nil
func (o *HyperflexHealthCheckDefinitionAllOf) SetDefaultHealthCheckScriptInfoNil() {
	o.DefaultHealthCheckScriptInfo.Set(nil)
}

// UnsetDefaultHealthCheckScriptInfo ensures that no value is present for DefaultHealthCheckScriptInfo, not even an explicit nil
func (o *HyperflexHealthCheckDefinitionAllOf) UnsetDefaultHealthCheckScriptInfo() {
	o.DefaultHealthCheckScriptInfo.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetHealthCheckScriptInfos returns the HealthCheckScriptInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealthCheckDefinitionAllOf) GetHealthCheckScriptInfos() []HyperflexHealthCheckScriptInfo {
	if o == nil {
		var ret []HyperflexHealthCheckScriptInfo
		return ret
	}
	return o.HealthCheckScriptInfos
}

// GetHealthCheckScriptInfosOk returns a tuple with the HealthCheckScriptInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealthCheckDefinitionAllOf) GetHealthCheckScriptInfosOk() (*[]HyperflexHealthCheckScriptInfo, bool) {
	if o == nil || o.HealthCheckScriptInfos == nil {
		return nil, false
	}
	return &o.HealthCheckScriptInfos, true
}

// HasHealthCheckScriptInfos returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasHealthCheckScriptInfos() bool {
	if o != nil && o.HealthCheckScriptInfos != nil {
		return true
	}

	return false
}

// SetHealthCheckScriptInfos gets a reference to the given []HyperflexHealthCheckScriptInfo and assigns it to the HealthCheckScriptInfos field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetHealthCheckScriptInfos(v []HyperflexHealthCheckScriptInfo) {
	o.HealthCheckScriptInfos = v
}

// GetHealthImpact returns the HealthImpact field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetHealthImpact() string {
	if o == nil || o.HealthImpact == nil {
		var ret string
		return ret
	}
	return *o.HealthImpact
}

// GetHealthImpactOk returns a tuple with the HealthImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetHealthImpactOk() (*string, bool) {
	if o == nil || o.HealthImpact == nil {
		return nil, false
	}
	return o.HealthImpact, true
}

// HasHealthImpact returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasHealthImpact() bool {
	if o != nil && o.HealthImpact != nil {
		return true
	}

	return false
}

// SetHealthImpact gets a reference to the given string and assigns it to the HealthImpact field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetHealthImpact(v string) {
	o.HealthImpact = &v
}

// GetInternalName returns the InternalName field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetInternalName() string {
	if o == nil || o.InternalName == nil {
		var ret string
		return ret
	}
	return *o.InternalName
}

// GetInternalNameOk returns a tuple with the InternalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetInternalNameOk() (*string, bool) {
	if o == nil || o.InternalName == nil {
		return nil, false
	}
	return o.InternalName, true
}

// HasInternalName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasInternalName() bool {
	if o != nil && o.InternalName != nil {
		return true
	}

	return false
}

// SetInternalName gets a reference to the given string and assigns it to the InternalName field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetInternalName(v string) {
	o.InternalName = &v
}

// GetMinimumHyperFlexVersion returns the MinimumHyperFlexVersion field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetMinimumHyperFlexVersion() string {
	if o == nil || o.MinimumHyperFlexVersion == nil {
		var ret string
		return ret
	}
	return *o.MinimumHyperFlexVersion
}

// GetMinimumHyperFlexVersionOk returns a tuple with the MinimumHyperFlexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetMinimumHyperFlexVersionOk() (*string, bool) {
	if o == nil || o.MinimumHyperFlexVersion == nil {
		return nil, false
	}
	return o.MinimumHyperFlexVersion, true
}

// HasMinimumHyperFlexVersion returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasMinimumHyperFlexVersion() bool {
	if o != nil && o.MinimumHyperFlexVersion != nil {
		return true
	}

	return false
}

// SetMinimumHyperFlexVersion gets a reference to the given string and assigns it to the MinimumHyperFlexVersion field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetMinimumHyperFlexVersion(v string) {
	o.MinimumHyperFlexVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetReference(v string) {
	o.Reference = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetResolution() string {
	if o == nil || o.Resolution == nil {
		var ret string
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetResolutionOk() (*string, bool) {
	if o == nil || o.Resolution == nil {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasResolution() bool {
	if o != nil && o.Resolution != nil {
		return true
	}

	return false
}

// SetResolution gets a reference to the given string and assigns it to the Resolution field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetResolution(v string) {
	o.Resolution = &v
}

// GetScriptExecutionMode returns the ScriptExecutionMode field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetScriptExecutionMode() string {
	if o == nil || o.ScriptExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ScriptExecutionMode
}

// GetScriptExecutionModeOk returns a tuple with the ScriptExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetScriptExecutionModeOk() (*string, bool) {
	if o == nil || o.ScriptExecutionMode == nil {
		return nil, false
	}
	return o.ScriptExecutionMode, true
}

// HasScriptExecutionMode returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasScriptExecutionMode() bool {
	if o != nil && o.ScriptExecutionMode != nil {
		return true
	}

	return false
}

// SetScriptExecutionMode gets a reference to the given string and assigns it to the ScriptExecutionMode field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetScriptExecutionMode(v string) {
	o.ScriptExecutionMode = &v
}

// GetScriptExecutionOnComputeNodes returns the ScriptExecutionOnComputeNodes field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetScriptExecutionOnComputeNodes() bool {
	if o == nil || o.ScriptExecutionOnComputeNodes == nil {
		var ret bool
		return ret
	}
	return *o.ScriptExecutionOnComputeNodes
}

// GetScriptExecutionOnComputeNodesOk returns a tuple with the ScriptExecutionOnComputeNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetScriptExecutionOnComputeNodesOk() (*bool, bool) {
	if o == nil || o.ScriptExecutionOnComputeNodes == nil {
		return nil, false
	}
	return o.ScriptExecutionOnComputeNodes, true
}

// HasScriptExecutionOnComputeNodes returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasScriptExecutionOnComputeNodes() bool {
	if o != nil && o.ScriptExecutionOnComputeNodes != nil {
		return true
	}

	return false
}

// SetScriptExecutionOnComputeNodes gets a reference to the given bool and assigns it to the ScriptExecutionOnComputeNodes field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetScriptExecutionOnComputeNodes(v bool) {
	o.ScriptExecutionOnComputeNodes = &v
}

// GetSupportedHypervisorType returns the SupportedHypervisorType field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetSupportedHypervisorType() string {
	if o == nil || o.SupportedHypervisorType == nil {
		var ret string
		return ret
	}
	return *o.SupportedHypervisorType
}

// GetSupportedHypervisorTypeOk returns a tuple with the SupportedHypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetSupportedHypervisorTypeOk() (*string, bool) {
	if o == nil || o.SupportedHypervisorType == nil {
		return nil, false
	}
	return o.SupportedHypervisorType, true
}

// HasSupportedHypervisorType returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasSupportedHypervisorType() bool {
	if o != nil && o.SupportedHypervisorType != nil {
		return true
	}

	return false
}

// SetSupportedHypervisorType gets a reference to the given string and assigns it to the SupportedHypervisorType field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetSupportedHypervisorType(v string) {
	o.SupportedHypervisorType = &v
}

// GetTargetExecutionType returns the TargetExecutionType field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetTargetExecutionType() string {
	if o == nil || o.TargetExecutionType == nil {
		var ret string
		return ret
	}
	return *o.TargetExecutionType
}

// GetTargetExecutionTypeOk returns a tuple with the TargetExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetTargetExecutionTypeOk() (*string, bool) {
	if o == nil || o.TargetExecutionType == nil {
		return nil, false
	}
	return o.TargetExecutionType, true
}

// HasTargetExecutionType returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasTargetExecutionType() bool {
	if o != nil && o.TargetExecutionType != nil {
		return true
	}

	return false
}

// SetTargetExecutionType gets a reference to the given string and assigns it to the TargetExecutionType field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetTargetExecutionType(v string) {
	o.TargetExecutionType = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *HyperflexHealthCheckDefinitionAllOf) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetUnsupportedHyperFlexVersions returns the UnsupportedHyperFlexVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHealthCheckDefinitionAllOf) GetUnsupportedHyperFlexVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UnsupportedHyperFlexVersions
}

// GetUnsupportedHyperFlexVersionsOk returns a tuple with the UnsupportedHyperFlexVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHealthCheckDefinitionAllOf) GetUnsupportedHyperFlexVersionsOk() (*[]string, bool) {
	if o == nil || o.UnsupportedHyperFlexVersions == nil {
		return nil, false
	}
	return &o.UnsupportedHyperFlexVersions, true
}

// HasUnsupportedHyperFlexVersions returns a boolean if a field has been set.
func (o *HyperflexHealthCheckDefinitionAllOf) HasUnsupportedHyperFlexVersions() bool {
	if o != nil && o.UnsupportedHyperFlexVersions != nil {
		return true
	}

	return false
}

// SetUnsupportedHyperFlexVersions gets a reference to the given []string and assigns it to the UnsupportedHyperFlexVersions field.
func (o *HyperflexHealthCheckDefinitionAllOf) SetUnsupportedHyperFlexVersions(v []string) {
	o.UnsupportedHyperFlexVersions = v
}

func (o HyperflexHealthCheckDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.CommonCauses != nil {
		toSerialize["CommonCauses"] = o.CommonCauses
	}
	if o.Configuration != nil {
		toSerialize["Configuration"] = o.Configuration
	}
	if o.DefaultHealthCheckScriptInfo.IsSet() {
		toSerialize["DefaultHealthCheckScriptInfo"] = o.DefaultHealthCheckScriptInfo.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.HealthCheckScriptInfos != nil {
		toSerialize["HealthCheckScriptInfos"] = o.HealthCheckScriptInfos
	}
	if o.HealthImpact != nil {
		toSerialize["HealthImpact"] = o.HealthImpact
	}
	if o.InternalName != nil {
		toSerialize["InternalName"] = o.InternalName
	}
	if o.MinimumHyperFlexVersion != nil {
		toSerialize["MinimumHyperFlexVersion"] = o.MinimumHyperFlexVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Reference != nil {
		toSerialize["Reference"] = o.Reference
	}
	if o.Resolution != nil {
		toSerialize["Resolution"] = o.Resolution
	}
	if o.ScriptExecutionMode != nil {
		toSerialize["ScriptExecutionMode"] = o.ScriptExecutionMode
	}
	if o.ScriptExecutionOnComputeNodes != nil {
		toSerialize["ScriptExecutionOnComputeNodes"] = o.ScriptExecutionOnComputeNodes
	}
	if o.SupportedHypervisorType != nil {
		toSerialize["SupportedHypervisorType"] = o.SupportedHypervisorType
	}
	if o.TargetExecutionType != nil {
		toSerialize["TargetExecutionType"] = o.TargetExecutionType
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.UnsupportedHyperFlexVersions != nil {
		toSerialize["UnsupportedHyperFlexVersions"] = o.UnsupportedHyperFlexVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHealthCheckDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHealthCheckDefinitionAllOf := _HyperflexHealthCheckDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHealthCheckDefinitionAllOf); err == nil {
		*o = HyperflexHealthCheckDefinitionAllOf(varHyperflexHealthCheckDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "CommonCauses")
		delete(additionalProperties, "Configuration")
		delete(additionalProperties, "DefaultHealthCheckScriptInfo")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "HealthCheckScriptInfos")
		delete(additionalProperties, "HealthImpact")
		delete(additionalProperties, "InternalName")
		delete(additionalProperties, "MinimumHyperFlexVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Reference")
		delete(additionalProperties, "Resolution")
		delete(additionalProperties, "ScriptExecutionMode")
		delete(additionalProperties, "ScriptExecutionOnComputeNodes")
		delete(additionalProperties, "SupportedHypervisorType")
		delete(additionalProperties, "TargetExecutionType")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "UnsupportedHyperFlexVersions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHealthCheckDefinitionAllOf struct {
	value *HyperflexHealthCheckDefinitionAllOf
	isSet bool
}

func (v NullableHyperflexHealthCheckDefinitionAllOf) Get() *HyperflexHealthCheckDefinitionAllOf {
	return v.value
}

func (v *NullableHyperflexHealthCheckDefinitionAllOf) Set(val *HyperflexHealthCheckDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHealthCheckDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHealthCheckDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHealthCheckDefinitionAllOf(val *HyperflexHealthCheckDefinitionAllOf) *NullableHyperflexHealthCheckDefinitionAllOf {
	return &NullableHyperflexHealthCheckDefinitionAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHealthCheckDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHealthCheckDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
