/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StoragePureReplicationScheduleAllOf Definition of the list of properties defined in 'storage.PureReplicationSchedule', excluding properties defined in parent classes.
type StoragePureReplicationScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total number of snapshots per day to be available on target above and over the specified retention time. PureStorage FlashArray maintains all created snapshot until retention period. Daily limit is applied only on the snapshots once retention time is expired. In case of, daily limit is less than the number of snapshot available on source, system select snapshots evenly spaced out throughout the day.
	DailyLimit                   *int64                           `json:"DailyLimit,omitempty"`
	ReplicationBlackoutIntervals []StoragePureReplicationBlackout `json:"ReplicationBlackoutIntervals,omitempty"`
	// Duration to keep the daily limit snapshots on target array. StorageArray deletes the snapshots permanently from the targets beyond this period.
	SnapshotExpiryTime   *string                                 `json:"SnapshotExpiryTime,omitempty"`
	Array                *StoragePureArrayRelationship           `json:"Array,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureReplicationScheduleAllOf StoragePureReplicationScheduleAllOf

// NewStoragePureReplicationScheduleAllOf instantiates a new StoragePureReplicationScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureReplicationScheduleAllOf(classId string, objectType string) *StoragePureReplicationScheduleAllOf {
	this := StoragePureReplicationScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureReplicationScheduleAllOfWithDefaults instantiates a new StoragePureReplicationScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureReplicationScheduleAllOfWithDefaults() *StoragePureReplicationScheduleAllOf {
	this := StoragePureReplicationScheduleAllOf{}
	var classId string = "storage.PureReplicationSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureReplicationSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureReplicationScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureReplicationScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureReplicationScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureReplicationScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDailyLimit returns the DailyLimit field value if set, zero value otherwise.
func (o *StoragePureReplicationScheduleAllOf) GetDailyLimit() int64 {
	if o == nil || o.DailyLimit == nil {
		var ret int64
		return ret
	}
	return *o.DailyLimit
}

// GetDailyLimitOk returns a tuple with the DailyLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetDailyLimitOk() (*int64, bool) {
	if o == nil || o.DailyLimit == nil {
		return nil, false
	}
	return o.DailyLimit, true
}

// HasDailyLimit returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasDailyLimit() bool {
	if o != nil && o.DailyLimit != nil {
		return true
	}

	return false
}

// SetDailyLimit gets a reference to the given int64 and assigns it to the DailyLimit field.
func (o *StoragePureReplicationScheduleAllOf) SetDailyLimit(v int64) {
	o.DailyLimit = &v
}

// GetReplicationBlackoutIntervals returns the ReplicationBlackoutIntervals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureReplicationScheduleAllOf) GetReplicationBlackoutIntervals() []StoragePureReplicationBlackout {
	if o == nil {
		var ret []StoragePureReplicationBlackout
		return ret
	}
	return o.ReplicationBlackoutIntervals
}

// GetReplicationBlackoutIntervalsOk returns a tuple with the ReplicationBlackoutIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureReplicationScheduleAllOf) GetReplicationBlackoutIntervalsOk() (*[]StoragePureReplicationBlackout, bool) {
	if o == nil || o.ReplicationBlackoutIntervals == nil {
		return nil, false
	}
	return &o.ReplicationBlackoutIntervals, true
}

// HasReplicationBlackoutIntervals returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasReplicationBlackoutIntervals() bool {
	if o != nil && o.ReplicationBlackoutIntervals != nil {
		return true
	}

	return false
}

// SetReplicationBlackoutIntervals gets a reference to the given []StoragePureReplicationBlackout and assigns it to the ReplicationBlackoutIntervals field.
func (o *StoragePureReplicationScheduleAllOf) SetReplicationBlackoutIntervals(v []StoragePureReplicationBlackout) {
	o.ReplicationBlackoutIntervals = v
}

// GetSnapshotExpiryTime returns the SnapshotExpiryTime field value if set, zero value otherwise.
func (o *StoragePureReplicationScheduleAllOf) GetSnapshotExpiryTime() string {
	if o == nil || o.SnapshotExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.SnapshotExpiryTime
}

// GetSnapshotExpiryTimeOk returns a tuple with the SnapshotExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetSnapshotExpiryTimeOk() (*string, bool) {
	if o == nil || o.SnapshotExpiryTime == nil {
		return nil, false
	}
	return o.SnapshotExpiryTime, true
}

// HasSnapshotExpiryTime returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasSnapshotExpiryTime() bool {
	if o != nil && o.SnapshotExpiryTime != nil {
		return true
	}

	return false
}

// SetSnapshotExpiryTime gets a reference to the given string and assigns it to the SnapshotExpiryTime field.
func (o *StoragePureReplicationScheduleAllOf) SetSnapshotExpiryTime(v string) {
	o.SnapshotExpiryTime = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureReplicationScheduleAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureReplicationScheduleAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureReplicationScheduleAllOf) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureReplicationScheduleAllOf) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureReplicationScheduleAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureReplicationScheduleAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureReplicationScheduleAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureReplicationScheduleAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DailyLimit != nil {
		toSerialize["DailyLimit"] = o.DailyLimit
	}
	if o.ReplicationBlackoutIntervals != nil {
		toSerialize["ReplicationBlackoutIntervals"] = o.ReplicationBlackoutIntervals
	}
	if o.SnapshotExpiryTime != nil {
		toSerialize["SnapshotExpiryTime"] = o.SnapshotExpiryTime
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureReplicationScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePureReplicationScheduleAllOf := _StoragePureReplicationScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePureReplicationScheduleAllOf); err == nil {
		*o = StoragePureReplicationScheduleAllOf(varStoragePureReplicationScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DailyLimit")
		delete(additionalProperties, "ReplicationBlackoutIntervals")
		delete(additionalProperties, "SnapshotExpiryTime")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePureReplicationScheduleAllOf struct {
	value *StoragePureReplicationScheduleAllOf
	isSet bool
}

func (v NullableStoragePureReplicationScheduleAllOf) Get() *StoragePureReplicationScheduleAllOf {
	return v.value
}

func (v *NullableStoragePureReplicationScheduleAllOf) Set(val *StoragePureReplicationScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureReplicationScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureReplicationScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureReplicationScheduleAllOf(val *StoragePureReplicationScheduleAllOf) *NullableStoragePureReplicationScheduleAllOf {
	return &NullableStoragePureReplicationScheduleAllOf{value: val, isSet: true}
}

func (v NullableStoragePureReplicationScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureReplicationScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
