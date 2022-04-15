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

// StorageNetAppVolume NetApp volume are data containers that enable you to partition and manage your data.
type StorageNetAppVolume struct {
	StorageBaseStorageContainer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink. * `off` - The volume will not grow or shrink in size in response to the amount of used space. * `grow` - The volume will automatically grow when used space in the volume is above the grow threshold. * `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space.
	AutosizeMode          *string                                 `json:"AutosizeMode,omitempty"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// The name of the Export Policy.
	ExportPolicyName *string `json:"ExportPolicyName,omitempty"`
	// Unique identifier of NetApp Volume across data center.
	Key *string `json:"Key,omitempty"`
	// The name of the Snapshot Policy.
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitempty"`
	// The UUID of the Snapshot Policy.
	SnapshotPolicyUuid *string `json:"SnapshotPolicyUuid,omitempty"`
	// The space that has been set aside as a reserve for Snapshot copy usage represented as a percent.
	SnapshotReservePercent *int64 `json:"SnapshotReservePercent,omitempty"`
	// The total space used by Snapshot copies in the volume represented in bytes.
	SnapshotUsed *float64 `json:"SnapshotUsed,omitempty"`
	// The current state of a NetApp volume. * `offline` - Read and write access to the volume is not allowed. * `online` - Read and write access to the volume is allowed. * `error` - Storage volume state of error type. * `mixed` - The constituents of a FlexGroup volume are not all in the same state.
	State *string `json:"State,omitempty"`
	// NetApp volume type. The volume type can be Read-write, Data-protection, or Load-sharing. * `data-protection` - Prevents modification of the data on the Volume. * `read-write` - Data on the Volume can be modified. * `load-sharing` - The volume type is Load Sharing DP.
	Type *string `json:"Type,omitempty"`
	// Universally unique identifier of a NetApp Volume.
	Uuid  *string                           `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppAggregate resources.
	DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
	// An array of relationships to storageNetAppVolumeEvent resources.
	Events               []StorageNetAppVolumeEventRelationship `json:"Events,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship    `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppVolume StorageNetAppVolume

// NewStorageNetAppVolume instantiates a new StorageNetAppVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppVolume(classId string, objectType string) *StorageNetAppVolume {
	this := StorageNetAppVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppVolumeWithDefaults instantiates a new StorageNetAppVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppVolumeWithDefaults() *StorageNetAppVolume {
	this := StorageNetAppVolume{}
	var classId string = "storage.NetAppVolume"
	this.ClassId = classId
	var objectType string = "storage.NetAppVolume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutosizeMode returns the AutosizeMode field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetAutosizeMode() string {
	if o == nil || o.AutosizeMode == nil {
		var ret string
		return ret
	}
	return *o.AutosizeMode
}

// GetAutosizeModeOk returns a tuple with the AutosizeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetAutosizeModeOk() (*string, bool) {
	if o == nil || o.AutosizeMode == nil {
		return nil, false
	}
	return o.AutosizeMode, true
}

// HasAutosizeMode returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasAutosizeMode() bool {
	if o != nil && o.AutosizeMode != nil {
		return true
	}

	return false
}

// SetAutosizeMode gets a reference to the given string and assigns it to the AutosizeMode field.
func (o *StorageNetAppVolume) SetAutosizeMode(v string) {
	o.AutosizeMode = &v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppVolume) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetExportPolicyName returns the ExportPolicyName field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetExportPolicyName() string {
	if o == nil || o.ExportPolicyName == nil {
		var ret string
		return ret
	}
	return *o.ExportPolicyName
}

// GetExportPolicyNameOk returns a tuple with the ExportPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetExportPolicyNameOk() (*string, bool) {
	if o == nil || o.ExportPolicyName == nil {
		return nil, false
	}
	return o.ExportPolicyName, true
}

// HasExportPolicyName returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasExportPolicyName() bool {
	if o != nil && o.ExportPolicyName != nil {
		return true
	}

	return false
}

// SetExportPolicyName gets a reference to the given string and assigns it to the ExportPolicyName field.
func (o *StorageNetAppVolume) SetExportPolicyName(v string) {
	o.ExportPolicyName = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppVolume) SetKey(v string) {
	o.Key = &v
}

// GetSnapshotPolicyName returns the SnapshotPolicyName field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetSnapshotPolicyName() string {
	if o == nil || o.SnapshotPolicyName == nil {
		var ret string
		return ret
	}
	return *o.SnapshotPolicyName
}

// GetSnapshotPolicyNameOk returns a tuple with the SnapshotPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetSnapshotPolicyNameOk() (*string, bool) {
	if o == nil || o.SnapshotPolicyName == nil {
		return nil, false
	}
	return o.SnapshotPolicyName, true
}

// HasSnapshotPolicyName returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasSnapshotPolicyName() bool {
	if o != nil && o.SnapshotPolicyName != nil {
		return true
	}

	return false
}

// SetSnapshotPolicyName gets a reference to the given string and assigns it to the SnapshotPolicyName field.
func (o *StorageNetAppVolume) SetSnapshotPolicyName(v string) {
	o.SnapshotPolicyName = &v
}

// GetSnapshotPolicyUuid returns the SnapshotPolicyUuid field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetSnapshotPolicyUuid() string {
	if o == nil || o.SnapshotPolicyUuid == nil {
		var ret string
		return ret
	}
	return *o.SnapshotPolicyUuid
}

// GetSnapshotPolicyUuidOk returns a tuple with the SnapshotPolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetSnapshotPolicyUuidOk() (*string, bool) {
	if o == nil || o.SnapshotPolicyUuid == nil {
		return nil, false
	}
	return o.SnapshotPolicyUuid, true
}

// HasSnapshotPolicyUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasSnapshotPolicyUuid() bool {
	if o != nil && o.SnapshotPolicyUuid != nil {
		return true
	}

	return false
}

// SetSnapshotPolicyUuid gets a reference to the given string and assigns it to the SnapshotPolicyUuid field.
func (o *StorageNetAppVolume) SetSnapshotPolicyUuid(v string) {
	o.SnapshotPolicyUuid = &v
}

// GetSnapshotReservePercent returns the SnapshotReservePercent field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetSnapshotReservePercent() int64 {
	if o == nil || o.SnapshotReservePercent == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotReservePercent
}

// GetSnapshotReservePercentOk returns a tuple with the SnapshotReservePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetSnapshotReservePercentOk() (*int64, bool) {
	if o == nil || o.SnapshotReservePercent == nil {
		return nil, false
	}
	return o.SnapshotReservePercent, true
}

// HasSnapshotReservePercent returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasSnapshotReservePercent() bool {
	if o != nil && o.SnapshotReservePercent != nil {
		return true
	}

	return false
}

// SetSnapshotReservePercent gets a reference to the given int64 and assigns it to the SnapshotReservePercent field.
func (o *StorageNetAppVolume) SetSnapshotReservePercent(v int64) {
	o.SnapshotReservePercent = &v
}

// GetSnapshotUsed returns the SnapshotUsed field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetSnapshotUsed() float64 {
	if o == nil || o.SnapshotUsed == nil {
		var ret float64
		return ret
	}
	return *o.SnapshotUsed
}

// GetSnapshotUsedOk returns a tuple with the SnapshotUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetSnapshotUsedOk() (*float64, bool) {
	if o == nil || o.SnapshotUsed == nil {
		return nil, false
	}
	return o.SnapshotUsed, true
}

// HasSnapshotUsed returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasSnapshotUsed() bool {
	if o != nil && o.SnapshotUsed != nil {
		return true
	}

	return false
}

// SetSnapshotUsed gets a reference to the given float64 and assigns it to the SnapshotUsed field.
func (o *StorageNetAppVolume) SetSnapshotUsed(v float64) {
	o.SnapshotUsed = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppVolume) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageNetAppVolume) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppVolume) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppVolume) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetDiskPool returns the DiskPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppVolume) GetDiskPool() []StorageNetAppAggregateRelationship {
	if o == nil {
		var ret []StorageNetAppAggregateRelationship
		return ret
	}
	return o.DiskPool
}

// GetDiskPoolOk returns a tuple with the DiskPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppVolume) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool) {
	if o == nil || o.DiskPool == nil {
		return nil, false
	}
	return &o.DiskPool, true
}

// HasDiskPool returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasDiskPool() bool {
	if o != nil && o.DiskPool != nil {
		return true
	}

	return false
}

// SetDiskPool gets a reference to the given []StorageNetAppAggregateRelationship and assigns it to the DiskPool field.
func (o *StorageNetAppVolume) SetDiskPool(v []StorageNetAppAggregateRelationship) {
	o.DiskPool = v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppVolume) GetEvents() []StorageNetAppVolumeEventRelationship {
	if o == nil {
		var ret []StorageNetAppVolumeEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppVolume) GetEventsOk() (*[]StorageNetAppVolumeEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppVolumeEventRelationship and assigns it to the Events field.
func (o *StorageNetAppVolume) SetEvents(v []StorageNetAppVolumeEventRelationship) {
	o.Events = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppVolume) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppVolume) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppVolume) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppVolume) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseStorageContainer, errStorageBaseStorageContainer := json.Marshal(o.StorageBaseStorageContainer)
	if errStorageBaseStorageContainer != nil {
		return []byte{}, errStorageBaseStorageContainer
	}
	errStorageBaseStorageContainer = json.Unmarshal([]byte(serializedStorageBaseStorageContainer), &toSerialize)
	if errStorageBaseStorageContainer != nil {
		return []byte{}, errStorageBaseStorageContainer
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutosizeMode != nil {
		toSerialize["AutosizeMode"] = o.AutosizeMode
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.ExportPolicyName != nil {
		toSerialize["ExportPolicyName"] = o.ExportPolicyName
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.SnapshotPolicyName != nil {
		toSerialize["SnapshotPolicyName"] = o.SnapshotPolicyName
	}
	if o.SnapshotPolicyUuid != nil {
		toSerialize["SnapshotPolicyUuid"] = o.SnapshotPolicyUuid
	}
	if o.SnapshotReservePercent != nil {
		toSerialize["SnapshotReservePercent"] = o.SnapshotReservePercent
	}
	if o.SnapshotUsed != nil {
		toSerialize["SnapshotUsed"] = o.SnapshotUsed
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.DiskPool != nil {
		toSerialize["DiskPool"] = o.DiskPool
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppVolume) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink. * `off` - The volume will not grow or shrink in size in response to the amount of used space. * `grow` - The volume will automatically grow when used space in the volume is above the grow threshold. * `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space.
		AutosizeMode          *string                                 `json:"AutosizeMode,omitempty"`
		AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
		// The name of the Export Policy.
		ExportPolicyName *string `json:"ExportPolicyName,omitempty"`
		// Unique identifier of NetApp Volume across data center.
		Key *string `json:"Key,omitempty"`
		// The name of the Snapshot Policy.
		SnapshotPolicyName *string `json:"SnapshotPolicyName,omitempty"`
		// The UUID of the Snapshot Policy.
		SnapshotPolicyUuid *string `json:"SnapshotPolicyUuid,omitempty"`
		// The space that has been set aside as a reserve for Snapshot copy usage represented as a percent.
		SnapshotReservePercent *int64 `json:"SnapshotReservePercent,omitempty"`
		// The total space used by Snapshot copies in the volume represented in bytes.
		SnapshotUsed *float64 `json:"SnapshotUsed,omitempty"`
		// The current state of a NetApp volume. * `offline` - Read and write access to the volume is not allowed. * `online` - Read and write access to the volume is allowed. * `error` - Storage volume state of error type. * `mixed` - The constituents of a FlexGroup volume are not all in the same state.
		State *string `json:"State,omitempty"`
		// NetApp volume type. The volume type can be Read-write, Data-protection, or Load-sharing. * `data-protection` - Prevents modification of the data on the Volume. * `read-write` - Data on the Volume can be modified. * `load-sharing` - The volume type is Load Sharing DP.
		Type *string `json:"Type,omitempty"`
		// Universally unique identifier of a NetApp Volume.
		Uuid  *string                           `json:"Uuid,omitempty"`
		Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
		// An array of relationships to storageNetAppAggregate resources.
		DiskPool []StorageNetAppAggregateRelationship `json:"DiskPool,omitempty"`
		// An array of relationships to storageNetAppVolumeEvent resources.
		Events []StorageNetAppVolumeEventRelationship `json:"Events,omitempty"`
		Tenant *StorageNetAppStorageVmRelationship    `json:"Tenant,omitempty"`
	}

	varStorageNetAppVolumeWithoutEmbeddedStruct := StorageNetAppVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppVolumeWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppVolume := _StorageNetAppVolume{}
		varStorageNetAppVolume.ClassId = varStorageNetAppVolumeWithoutEmbeddedStruct.ClassId
		varStorageNetAppVolume.ObjectType = varStorageNetAppVolumeWithoutEmbeddedStruct.ObjectType
		varStorageNetAppVolume.AutosizeMode = varStorageNetAppVolumeWithoutEmbeddedStruct.AutosizeMode
		varStorageNetAppVolume.AvgPerformanceMetrics = varStorageNetAppVolumeWithoutEmbeddedStruct.AvgPerformanceMetrics
		varStorageNetAppVolume.ExportPolicyName = varStorageNetAppVolumeWithoutEmbeddedStruct.ExportPolicyName
		varStorageNetAppVolume.Key = varStorageNetAppVolumeWithoutEmbeddedStruct.Key
		varStorageNetAppVolume.SnapshotPolicyName = varStorageNetAppVolumeWithoutEmbeddedStruct.SnapshotPolicyName
		varStorageNetAppVolume.SnapshotPolicyUuid = varStorageNetAppVolumeWithoutEmbeddedStruct.SnapshotPolicyUuid
		varStorageNetAppVolume.SnapshotReservePercent = varStorageNetAppVolumeWithoutEmbeddedStruct.SnapshotReservePercent
		varStorageNetAppVolume.SnapshotUsed = varStorageNetAppVolumeWithoutEmbeddedStruct.SnapshotUsed
		varStorageNetAppVolume.State = varStorageNetAppVolumeWithoutEmbeddedStruct.State
		varStorageNetAppVolume.Type = varStorageNetAppVolumeWithoutEmbeddedStruct.Type
		varStorageNetAppVolume.Uuid = varStorageNetAppVolumeWithoutEmbeddedStruct.Uuid
		varStorageNetAppVolume.Array = varStorageNetAppVolumeWithoutEmbeddedStruct.Array
		varStorageNetAppVolume.DiskPool = varStorageNetAppVolumeWithoutEmbeddedStruct.DiskPool
		varStorageNetAppVolume.Events = varStorageNetAppVolumeWithoutEmbeddedStruct.Events
		varStorageNetAppVolume.Tenant = varStorageNetAppVolumeWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppVolume(varStorageNetAppVolume)
	} else {
		return err
	}

	varStorageNetAppVolume := _StorageNetAppVolume{}

	err = json.Unmarshal(bytes, &varStorageNetAppVolume)
	if err == nil {
		o.StorageBaseStorageContainer = varStorageNetAppVolume.StorageBaseStorageContainer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutosizeMode")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "ExportPolicyName")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "SnapshotPolicyName")
		delete(additionalProperties, "SnapshotPolicyUuid")
		delete(additionalProperties, "SnapshotReservePercent")
		delete(additionalProperties, "SnapshotUsed")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "DiskPool")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageBaseStorageContainer := reflect.ValueOf(o.StorageBaseStorageContainer)
		for i := 0; i < reflectStorageBaseStorageContainer.Type().NumField(); i++ {
			t := reflectStorageBaseStorageContainer.Type().Field(i)

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

type NullableStorageNetAppVolume struct {
	value *StorageNetAppVolume
	isSet bool
}

func (v NullableStorageNetAppVolume) Get() *StorageNetAppVolume {
	return v.value
}

func (v *NullableStorageNetAppVolume) Set(val *StorageNetAppVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppVolume(val *StorageNetAppVolume) *NullableStorageNetAppVolume {
	return &NullableStorageNetAppVolume{value: val, isSet: true}
}

func (v NullableStorageNetAppVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
