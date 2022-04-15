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

// StorageHitachiDiskAllOf Definition of the list of properties defined in 'storage.HitachiDisk', excluding properties defined in parent classes.
type StorageHitachiDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Drive type code of the Hitachi Disk.
	DriveTypeCode *string `json:"DriveTypeCode,omitempty"`
	// Parity group number. When the drive does not belong to any parity group, an empty character string is output.
	ParityGroupId *string `json:"ParityGroupId,omitempty"`
	// Drive type of the Hitachi Disk. * `N/A` - Drive Type is not available. * `SAS` - SAS stands for Serial Attached SCSI. * `SSD(MLC)` - SSD (MLC) stands for Multiple Level Cell. * `SSD(FMC)` - SSD (FMC) stands for Flash Memory Compressed. * `SSD(FMD)` - SSD (FMD) stands for Flash Module Drive. * `SSD(SLC)` - SSD (SLC) stands for Single Level Cell. * `SSD` - SSD stands for Solid-State Drive. * `SSD(RI)` - SSD (RI) stands for Read Intensive. * `SCM` - SCM stands for Storage Class Memory.
	TypeDetail *string `json:"TypeDetail,omitempty"`
	// Purpose for which the drive is used.
	Usage                *string                                `json:"Usage,omitempty"`
	Array                *StorageHitachiArrayRelationship       `json:"Array,omitempty"`
	ParityGroup          *StorageHitachiParityGroupRelationship `json:"ParityGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiDiskAllOf StorageHitachiDiskAllOf

// NewStorageHitachiDiskAllOf instantiates a new StorageHitachiDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiDiskAllOf(classId string, objectType string) *StorageHitachiDiskAllOf {
	this := StorageHitachiDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiDiskAllOfWithDefaults instantiates a new StorageHitachiDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiDiskAllOfWithDefaults() *StorageHitachiDiskAllOf {
	this := StorageHitachiDiskAllOf{}
	var classId string = "storage.HitachiDisk"
	this.ClassId = classId
	var objectType string = "storage.HitachiDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveTypeCode returns the DriveTypeCode field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetDriveTypeCode() string {
	if o == nil || o.DriveTypeCode == nil {
		var ret string
		return ret
	}
	return *o.DriveTypeCode
}

// GetDriveTypeCodeOk returns a tuple with the DriveTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetDriveTypeCodeOk() (*string, bool) {
	if o == nil || o.DriveTypeCode == nil {
		return nil, false
	}
	return o.DriveTypeCode, true
}

// HasDriveTypeCode returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasDriveTypeCode() bool {
	if o != nil && o.DriveTypeCode != nil {
		return true
	}

	return false
}

// SetDriveTypeCode gets a reference to the given string and assigns it to the DriveTypeCode field.
func (o *StorageHitachiDiskAllOf) SetDriveTypeCode(v string) {
	o.DriveTypeCode = &v
}

// GetParityGroupId returns the ParityGroupId field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetParityGroupId() string {
	if o == nil || o.ParityGroupId == nil {
		var ret string
		return ret
	}
	return *o.ParityGroupId
}

// GetParityGroupIdOk returns a tuple with the ParityGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetParityGroupIdOk() (*string, bool) {
	if o == nil || o.ParityGroupId == nil {
		return nil, false
	}
	return o.ParityGroupId, true
}

// HasParityGroupId returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasParityGroupId() bool {
	if o != nil && o.ParityGroupId != nil {
		return true
	}

	return false
}

// SetParityGroupId gets a reference to the given string and assigns it to the ParityGroupId field.
func (o *StorageHitachiDiskAllOf) SetParityGroupId(v string) {
	o.ParityGroupId = &v
}

// GetTypeDetail returns the TypeDetail field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetTypeDetail() string {
	if o == nil || o.TypeDetail == nil {
		var ret string
		return ret
	}
	return *o.TypeDetail
}

// GetTypeDetailOk returns a tuple with the TypeDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetTypeDetailOk() (*string, bool) {
	if o == nil || o.TypeDetail == nil {
		return nil, false
	}
	return o.TypeDetail, true
}

// HasTypeDetail returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasTypeDetail() bool {
	if o != nil && o.TypeDetail != nil {
		return true
	}

	return false
}

// SetTypeDetail gets a reference to the given string and assigns it to the TypeDetail field.
func (o *StorageHitachiDiskAllOf) SetTypeDetail(v string) {
	o.TypeDetail = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetUsage() string {
	if o == nil || o.Usage == nil {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetUsageOk() (*string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *StorageHitachiDiskAllOf) SetUsage(v string) {
	o.Usage = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiDiskAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetParityGroup returns the ParityGroup field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetParityGroup() StorageHitachiParityGroupRelationship {
	if o == nil || o.ParityGroup == nil {
		var ret StorageHitachiParityGroupRelationship
		return ret
	}
	return *o.ParityGroup
}

// GetParityGroupOk returns a tuple with the ParityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetParityGroupOk() (*StorageHitachiParityGroupRelationship, bool) {
	if o == nil || o.ParityGroup == nil {
		return nil, false
	}
	return o.ParityGroup, true
}

// HasParityGroup returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasParityGroup() bool {
	if o != nil && o.ParityGroup != nil {
		return true
	}

	return false
}

// SetParityGroup gets a reference to the given StorageHitachiParityGroupRelationship and assigns it to the ParityGroup field.
func (o *StorageHitachiDiskAllOf) SetParityGroup(v StorageHitachiParityGroupRelationship) {
	o.ParityGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiDiskAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriveTypeCode != nil {
		toSerialize["DriveTypeCode"] = o.DriveTypeCode
	}
	if o.ParityGroupId != nil {
		toSerialize["ParityGroupId"] = o.ParityGroupId
	}
	if o.TypeDetail != nil {
		toSerialize["TypeDetail"] = o.TypeDetail
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.ParityGroup != nil {
		toSerialize["ParityGroup"] = o.ParityGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiDiskAllOf := _StorageHitachiDiskAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiDiskAllOf); err == nil {
		*o = StorageHitachiDiskAllOf(varStorageHitachiDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveTypeCode")
		delete(additionalProperties, "ParityGroupId")
		delete(additionalProperties, "TypeDetail")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "ParityGroup")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiDiskAllOf struct {
	value *StorageHitachiDiskAllOf
	isSet bool
}

func (v NullableStorageHitachiDiskAllOf) Get() *StorageHitachiDiskAllOf {
	return v.value
}

func (v *NullableStorageHitachiDiskAllOf) Set(val *StorageHitachiDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiDiskAllOf(val *StorageHitachiDiskAllOf) *NullableStorageHitachiDiskAllOf {
	return &NullableStorageHitachiDiskAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
