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
)

// StorageSpanAllOf Definition of the list of properties defined in 'storage.Span', excluding properties defined in parent classes.
type StorageSpanAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string  `json:"ObjectType"`
	Slots      []int64 `json:"Slots,omitempty"`
	// Unique identifier value of this span.
	SpanId    *int64                        `json:"SpanId,omitempty"`
	DiskGroup *StorageDiskGroupRelationship `json:"DiskGroup,omitempty"`
	// An array of relationships to storagePhysicalDisk resources.
	PhysicalDisks        []StoragePhysicalDiskRelationship    `json:"PhysicalDisks,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSpanAllOf StorageSpanAllOf

// NewStorageSpanAllOf instantiates a new StorageSpanAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpanAllOf(classId string, objectType string) *StorageSpanAllOf {
	this := StorageSpanAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSpanAllOfWithDefaults instantiates a new StorageSpanAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpanAllOfWithDefaults() *StorageSpanAllOf {
	this := StorageSpanAllOf{}
	var classId string = "storage.Span"
	this.ClassId = classId
	var objectType string = "storage.Span"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSpanAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSpanAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSpanAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageSpanAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSpanAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSpanAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSlots returns the Slots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageSpanAllOf) GetSlots() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageSpanAllOf) GetSlotsOk() (*[]int64, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return &o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *StorageSpanAllOf) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []int64 and assigns it to the Slots field.
func (o *StorageSpanAllOf) SetSlots(v []int64) {
	o.Slots = v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *StorageSpanAllOf) GetSpanId() int64 {
	if o == nil || o.SpanId == nil {
		var ret int64
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanAllOf) GetSpanIdOk() (*int64, bool) {
	if o == nil || o.SpanId == nil {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *StorageSpanAllOf) HasSpanId() bool {
	if o != nil && o.SpanId != nil {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given int64 and assigns it to the SpanId field.
func (o *StorageSpanAllOf) SetSpanId(v int64) {
	o.SpanId = &v
}

// GetDiskGroup returns the DiskGroup field value if set, zero value otherwise.
func (o *StorageSpanAllOf) GetDiskGroup() StorageDiskGroupRelationship {
	if o == nil || o.DiskGroup == nil {
		var ret StorageDiskGroupRelationship
		return ret
	}
	return *o.DiskGroup
}

// GetDiskGroupOk returns a tuple with the DiskGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanAllOf) GetDiskGroupOk() (*StorageDiskGroupRelationship, bool) {
	if o == nil || o.DiskGroup == nil {
		return nil, false
	}
	return o.DiskGroup, true
}

// HasDiskGroup returns a boolean if a field has been set.
func (o *StorageSpanAllOf) HasDiskGroup() bool {
	if o != nil && o.DiskGroup != nil {
		return true
	}

	return false
}

// SetDiskGroup gets a reference to the given StorageDiskGroupRelationship and assigns it to the DiskGroup field.
func (o *StorageSpanAllOf) SetDiskGroup(v StorageDiskGroupRelationship) {
	o.DiskGroup = &v
}

// GetPhysicalDisks returns the PhysicalDisks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageSpanAllOf) GetPhysicalDisks() []StoragePhysicalDiskRelationship {
	if o == nil {
		var ret []StoragePhysicalDiskRelationship
		return ret
	}
	return o.PhysicalDisks
}

// GetPhysicalDisksOk returns a tuple with the PhysicalDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageSpanAllOf) GetPhysicalDisksOk() (*[]StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.PhysicalDisks == nil {
		return nil, false
	}
	return &o.PhysicalDisks, true
}

// HasPhysicalDisks returns a boolean if a field has been set.
func (o *StorageSpanAllOf) HasPhysicalDisks() bool {
	if o != nil && o.PhysicalDisks != nil {
		return true
	}

	return false
}

// SetPhysicalDisks gets a reference to the given []StoragePhysicalDiskRelationship and assigns it to the PhysicalDisks field.
func (o *StorageSpanAllOf) SetPhysicalDisks(v []StoragePhysicalDiskRelationship) {
	o.PhysicalDisks = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageSpanAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpanAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageSpanAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageSpanAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageSpanAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Slots != nil {
		toSerialize["Slots"] = o.Slots
	}
	if o.SpanId != nil {
		toSerialize["SpanId"] = o.SpanId
	}
	if o.DiskGroup != nil {
		toSerialize["DiskGroup"] = o.DiskGroup
	}
	if o.PhysicalDisks != nil {
		toSerialize["PhysicalDisks"] = o.PhysicalDisks
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageSpanAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageSpanAllOf := _StorageSpanAllOf{}

	if err = json.Unmarshal(bytes, &varStorageSpanAllOf); err == nil {
		*o = StorageSpanAllOf(varStorageSpanAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Slots")
		delete(additionalProperties, "SpanId")
		delete(additionalProperties, "DiskGroup")
		delete(additionalProperties, "PhysicalDisks")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageSpanAllOf struct {
	value *StorageSpanAllOf
	isSet bool
}

func (v NullableStorageSpanAllOf) Get() *StorageSpanAllOf {
	return v.value
}

func (v *NullableStorageSpanAllOf) Set(val *StorageSpanAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpanAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpanAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpanAllOf(val *StorageSpanAllOf) *NullableStorageSpanAllOf {
	return &NullableStorageSpanAllOf{value: val, isSet: true}
}

func (v NullableStorageSpanAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpanAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
