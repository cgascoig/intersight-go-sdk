/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageNetAppBaseEvent An event is a notification that is generated automatically when a predefined condition occurs or when an object crosses a threshold.
type StorageNetAppBaseEvent struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// A message describing the cause for the event.
	Cause *string `json:"Cause,omitempty"`
	// Unique identifier of the cluster across the datacenter.
	ClusterUuid *string `json:"ClusterUuid,omitempty"`
	// The current state of the event. * `unknown` - Default unknown current state. * `new` - The current state of the event is new. * `acknowledged` - The current state of the event is acknowledged. * `resolved` - The current state of the event is resolved. * `obsolete` - The current state of the event is obsolete.
	CurrentState *string `json:"CurrentState,omitempty"`
	// Time since the event was created, in ISO8601 standard.
	Duration *string `json:"Duration,omitempty"`
	// Impact area of the event (availability, capacity, configuration, performance, protection, or security). * `unknown` - Default unknown impact area. * `availability` - The impact area of the event is availability. * `capacity` - The impact area of the event is capacity. * `configuration` - The impact area of the event is configuration. * `performance` - The impact area of the event is performance. * `protection` - The impact area of the event is protection. * `security` - The impact area of the event is security.
	ImpactArea *string `json:"ImpactArea,omitempty"`
	// Impact level of the event (event, risk, incident, or upgrade). * `unknown` - Default unknown impact level. * `event` - The impact level of the event is event. * `risk` - The impact level of the event is risk. * `incident` - The impact level of the event is incident. * `upgrade` - The impact level of the event is upgrade.
	ImpactLevel *string `json:"ImpactLevel,omitempty"`
	// The full name of the source of the event.
	ImpactResourceName *string `json:"ImpactResourceName,omitempty"`
	// Type of resource with which the event is associated. * `unknown` - Default unknown resource type. * `aggregate` - The type of resource impacted by the event is an aggregate. * `cluster` - The type of resource impacted by the event is a cluster. * `cluster_node` - The type of resource impacted by the event is a node. * `disk` - The type of resource impacted by the event is a disk. * `fcp_lif` - The type of resource impacted by the event is a FC interface. * `fcp_port` - The type of resource impacted by the event is a FC port. * `lun` - The type of resource impacted by the event is a lun. * `network_lif` - The type of resource impacted by the event is an ethernet interface. * `network_port` - The type of resource impacted by the event is an ethernet port. * `volume` - The type of resource impacted by the event is a volume. * `vserver` - The type of resource impacted by the event is a storage VM.
	ImpactResourceType *string `json:"ImpactResourceType,omitempty"`
	// The unique identifier of the impacted resource.
	ImpactResourceUuid *string `json:"ImpactResourceUuid,omitempty"`
	// The name of the event that occurred.
	Name *string `json:"Name,omitempty"`
	// Unique identifier of the node across the cluster.
	NodeUuid *string `json:"NodeUuid,omitempty"`
	// The severity of the event. * `unknown` - Default unknown severity. * `normal` - The severity of the event is normal. * `information` - The severity of the event is information. * `warning` - The severity of the event is warning. * `error` - The severity of the event is error. * `critical` - The severity of the event is critical.
	Severity *string `json:"Severity,omitempty"`
	// Unique identifier of the storage VM.
	SvmUuid *string `json:"SvmUuid,omitempty"`
	// Unique identifier of the event.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppBaseEvent StorageNetAppBaseEvent

// NewStorageNetAppBaseEvent instantiates a new StorageNetAppBaseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppBaseEvent(classId string, objectType string) *StorageNetAppBaseEvent {
	this := StorageNetAppBaseEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppBaseEventWithDefaults instantiates a new StorageNetAppBaseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppBaseEventWithDefaults() *StorageNetAppBaseEvent {
	this := StorageNetAppBaseEvent{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppBaseEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppBaseEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppBaseEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppBaseEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetCause() string {
	if o == nil || o.Cause == nil {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetCauseOk() (*string, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *StorageNetAppBaseEvent) SetCause(v string) {
	o.Cause = &v
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetClusterUuid() string {
	if o == nil || o.ClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetClusterUuidOk() (*string, bool) {
	if o == nil || o.ClusterUuid == nil {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasClusterUuid() bool {
	if o != nil && o.ClusterUuid != nil {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppBaseEvent) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetCurrentState() string {
	if o == nil || o.CurrentState == nil {
		var ret string
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetCurrentStateOk() (*string, bool) {
	if o == nil || o.CurrentState == nil {
		return nil, false
	}
	return o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasCurrentState() bool {
	if o != nil && o.CurrentState != nil {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.
func (o *StorageNetAppBaseEvent) SetCurrentState(v string) {
	o.CurrentState = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *StorageNetAppBaseEvent) SetDuration(v string) {
	o.Duration = &v
}

// GetImpactArea returns the ImpactArea field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetImpactArea() string {
	if o == nil || o.ImpactArea == nil {
		var ret string
		return ret
	}
	return *o.ImpactArea
}

// GetImpactAreaOk returns a tuple with the ImpactArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetImpactAreaOk() (*string, bool) {
	if o == nil || o.ImpactArea == nil {
		return nil, false
	}
	return o.ImpactArea, true
}

// HasImpactArea returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasImpactArea() bool {
	if o != nil && o.ImpactArea != nil {
		return true
	}

	return false
}

// SetImpactArea gets a reference to the given string and assigns it to the ImpactArea field.
func (o *StorageNetAppBaseEvent) SetImpactArea(v string) {
	o.ImpactArea = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetImpactLevel() string {
	if o == nil || o.ImpactLevel == nil {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetImpactLevelOk() (*string, bool) {
	if o == nil || o.ImpactLevel == nil {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasImpactLevel() bool {
	if o != nil && o.ImpactLevel != nil {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *StorageNetAppBaseEvent) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetImpactResourceName returns the ImpactResourceName field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetImpactResourceName() string {
	if o == nil || o.ImpactResourceName == nil {
		var ret string
		return ret
	}
	return *o.ImpactResourceName
}

// GetImpactResourceNameOk returns a tuple with the ImpactResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetImpactResourceNameOk() (*string, bool) {
	if o == nil || o.ImpactResourceName == nil {
		return nil, false
	}
	return o.ImpactResourceName, true
}

// HasImpactResourceName returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasImpactResourceName() bool {
	if o != nil && o.ImpactResourceName != nil {
		return true
	}

	return false
}

// SetImpactResourceName gets a reference to the given string and assigns it to the ImpactResourceName field.
func (o *StorageNetAppBaseEvent) SetImpactResourceName(v string) {
	o.ImpactResourceName = &v
}

// GetImpactResourceType returns the ImpactResourceType field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetImpactResourceType() string {
	if o == nil || o.ImpactResourceType == nil {
		var ret string
		return ret
	}
	return *o.ImpactResourceType
}

// GetImpactResourceTypeOk returns a tuple with the ImpactResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetImpactResourceTypeOk() (*string, bool) {
	if o == nil || o.ImpactResourceType == nil {
		return nil, false
	}
	return o.ImpactResourceType, true
}

// HasImpactResourceType returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasImpactResourceType() bool {
	if o != nil && o.ImpactResourceType != nil {
		return true
	}

	return false
}

// SetImpactResourceType gets a reference to the given string and assigns it to the ImpactResourceType field.
func (o *StorageNetAppBaseEvent) SetImpactResourceType(v string) {
	o.ImpactResourceType = &v
}

// GetImpactResourceUuid returns the ImpactResourceUuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetImpactResourceUuid() string {
	if o == nil || o.ImpactResourceUuid == nil {
		var ret string
		return ret
	}
	return *o.ImpactResourceUuid
}

// GetImpactResourceUuidOk returns a tuple with the ImpactResourceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetImpactResourceUuidOk() (*string, bool) {
	if o == nil || o.ImpactResourceUuid == nil {
		return nil, false
	}
	return o.ImpactResourceUuid, true
}

// HasImpactResourceUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasImpactResourceUuid() bool {
	if o != nil && o.ImpactResourceUuid != nil {
		return true
	}

	return false
}

// SetImpactResourceUuid gets a reference to the given string and assigns it to the ImpactResourceUuid field.
func (o *StorageNetAppBaseEvent) SetImpactResourceUuid(v string) {
	o.ImpactResourceUuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppBaseEvent) SetName(v string) {
	o.Name = &v
}

// GetNodeUuid returns the NodeUuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetNodeUuid() string {
	if o == nil || o.NodeUuid == nil {
		var ret string
		return ret
	}
	return *o.NodeUuid
}

// GetNodeUuidOk returns a tuple with the NodeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetNodeUuidOk() (*string, bool) {
	if o == nil || o.NodeUuid == nil {
		return nil, false
	}
	return o.NodeUuid, true
}

// HasNodeUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasNodeUuid() bool {
	if o != nil && o.NodeUuid != nil {
		return true
	}

	return false
}

// SetNodeUuid gets a reference to the given string and assigns it to the NodeUuid field.
func (o *StorageNetAppBaseEvent) SetNodeUuid(v string) {
	o.NodeUuid = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *StorageNetAppBaseEvent) SetSeverity(v string) {
	o.Severity = &v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetSvmUuid() string {
	if o == nil || o.SvmUuid == nil {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetSvmUuidOk() (*string, bool) {
	if o == nil || o.SvmUuid == nil {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasSvmUuid() bool {
	if o != nil && o.SvmUuid != nil {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppBaseEvent) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseEvent) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseEvent) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseEvent) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppBaseEvent) SetUuid(v string) {
	o.Uuid = &v
}

func (o StorageNetAppBaseEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cause != nil {
		toSerialize["Cause"] = o.Cause
	}
	if o.ClusterUuid != nil {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if o.CurrentState != nil {
		toSerialize["CurrentState"] = o.CurrentState
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.ImpactArea != nil {
		toSerialize["ImpactArea"] = o.ImpactArea
	}
	if o.ImpactLevel != nil {
		toSerialize["ImpactLevel"] = o.ImpactLevel
	}
	if o.ImpactResourceName != nil {
		toSerialize["ImpactResourceName"] = o.ImpactResourceName
	}
	if o.ImpactResourceType != nil {
		toSerialize["ImpactResourceType"] = o.ImpactResourceType
	}
	if o.ImpactResourceUuid != nil {
		toSerialize["ImpactResourceUuid"] = o.ImpactResourceUuid
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NodeUuid != nil {
		toSerialize["NodeUuid"] = o.NodeUuid
	}
	if o.Severity != nil {
		toSerialize["Severity"] = o.Severity
	}
	if o.SvmUuid != nil {
		toSerialize["SvmUuid"] = o.SvmUuid
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppBaseEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppBaseEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// A message describing the cause for the event.
		Cause *string `json:"Cause,omitempty"`
		// Unique identifier of the cluster across the datacenter.
		ClusterUuid *string `json:"ClusterUuid,omitempty"`
		// The current state of the event. * `unknown` - Default unknown current state. * `new` - The current state of the event is new. * `acknowledged` - The current state of the event is acknowledged. * `resolved` - The current state of the event is resolved. * `obsolete` - The current state of the event is obsolete.
		CurrentState *string `json:"CurrentState,omitempty"`
		// Time since the event was created, in ISO8601 standard.
		Duration *string `json:"Duration,omitempty"`
		// Impact area of the event (availability, capacity, configuration, performance, protection, or security). * `unknown` - Default unknown impact area. * `availability` - The impact area of the event is availability. * `capacity` - The impact area of the event is capacity. * `configuration` - The impact area of the event is configuration. * `performance` - The impact area of the event is performance. * `protection` - The impact area of the event is protection. * `security` - The impact area of the event is security.
		ImpactArea *string `json:"ImpactArea,omitempty"`
		// Impact level of the event (event, risk, incident, or upgrade). * `unknown` - Default unknown impact level. * `event` - The impact level of the event is event. * `risk` - The impact level of the event is risk. * `incident` - The impact level of the event is incident. * `upgrade` - The impact level of the event is upgrade.
		ImpactLevel *string `json:"ImpactLevel,omitempty"`
		// The full name of the source of the event.
		ImpactResourceName *string `json:"ImpactResourceName,omitempty"`
		// Type of resource with which the event is associated. * `unknown` - Default unknown resource type. * `aggregate` - The type of resource impacted by the event is an aggregate. * `cluster` - The type of resource impacted by the event is a cluster. * `cluster_node` - The type of resource impacted by the event is a node. * `disk` - The type of resource impacted by the event is a disk. * `fcp_lif` - The type of resource impacted by the event is a FC interface. * `fcp_port` - The type of resource impacted by the event is a FC port. * `lun` - The type of resource impacted by the event is a lun. * `network_lif` - The type of resource impacted by the event is an ethernet interface. * `network_port` - The type of resource impacted by the event is an ethernet port. * `volume` - The type of resource impacted by the event is a volume. * `vserver` - The type of resource impacted by the event is a storage VM.
		ImpactResourceType *string `json:"ImpactResourceType,omitempty"`
		// The unique identifier of the impacted resource.
		ImpactResourceUuid *string `json:"ImpactResourceUuid,omitempty"`
		// The name of the event that occurred.
		Name *string `json:"Name,omitempty"`
		// Unique identifier of the node across the cluster.
		NodeUuid *string `json:"NodeUuid,omitempty"`
		// The severity of the event. * `unknown` - Default unknown severity. * `normal` - The severity of the event is normal. * `information` - The severity of the event is information. * `warning` - The severity of the event is warning. * `error` - The severity of the event is error. * `critical` - The severity of the event is critical.
		Severity *string `json:"Severity,omitempty"`
		// Unique identifier of the storage VM.
		SvmUuid *string `json:"SvmUuid,omitempty"`
		// Unique identifier of the event.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varStorageNetAppBaseEventWithoutEmbeddedStruct := StorageNetAppBaseEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppBaseEvent := _StorageNetAppBaseEvent{}
		varStorageNetAppBaseEvent.ClassId = varStorageNetAppBaseEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppBaseEvent.ObjectType = varStorageNetAppBaseEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppBaseEvent.Cause = varStorageNetAppBaseEventWithoutEmbeddedStruct.Cause
		varStorageNetAppBaseEvent.ClusterUuid = varStorageNetAppBaseEventWithoutEmbeddedStruct.ClusterUuid
		varStorageNetAppBaseEvent.CurrentState = varStorageNetAppBaseEventWithoutEmbeddedStruct.CurrentState
		varStorageNetAppBaseEvent.Duration = varStorageNetAppBaseEventWithoutEmbeddedStruct.Duration
		varStorageNetAppBaseEvent.ImpactArea = varStorageNetAppBaseEventWithoutEmbeddedStruct.ImpactArea
		varStorageNetAppBaseEvent.ImpactLevel = varStorageNetAppBaseEventWithoutEmbeddedStruct.ImpactLevel
		varStorageNetAppBaseEvent.ImpactResourceName = varStorageNetAppBaseEventWithoutEmbeddedStruct.ImpactResourceName
		varStorageNetAppBaseEvent.ImpactResourceType = varStorageNetAppBaseEventWithoutEmbeddedStruct.ImpactResourceType
		varStorageNetAppBaseEvent.ImpactResourceUuid = varStorageNetAppBaseEventWithoutEmbeddedStruct.ImpactResourceUuid
		varStorageNetAppBaseEvent.Name = varStorageNetAppBaseEventWithoutEmbeddedStruct.Name
		varStorageNetAppBaseEvent.NodeUuid = varStorageNetAppBaseEventWithoutEmbeddedStruct.NodeUuid
		varStorageNetAppBaseEvent.Severity = varStorageNetAppBaseEventWithoutEmbeddedStruct.Severity
		varStorageNetAppBaseEvent.SvmUuid = varStorageNetAppBaseEventWithoutEmbeddedStruct.SvmUuid
		varStorageNetAppBaseEvent.Uuid = varStorageNetAppBaseEventWithoutEmbeddedStruct.Uuid
		*o = StorageNetAppBaseEvent(varStorageNetAppBaseEvent)
	} else {
		return err
	}

	varStorageNetAppBaseEvent := _StorageNetAppBaseEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseEvent)
	if err == nil {
		o.MoBaseMo = varStorageNetAppBaseEvent.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cause")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "CurrentState")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "ImpactArea")
		delete(additionalProperties, "ImpactLevel")
		delete(additionalProperties, "ImpactResourceName")
		delete(additionalProperties, "ImpactResourceType")
		delete(additionalProperties, "ImpactResourceUuid")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NodeUuid")
		delete(additionalProperties, "Severity")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "Uuid")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppBaseEvent struct {
	value *StorageNetAppBaseEvent
	isSet bool
}

func (v NullableStorageNetAppBaseEvent) Get() *StorageNetAppBaseEvent {
	return v.value
}

func (v *NullableStorageNetAppBaseEvent) Set(val *StorageNetAppBaseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseEvent(val *StorageNetAppBaseEvent) *NullableStorageNetAppBaseEvent {
	return &NullableStorageNetAppBaseEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
