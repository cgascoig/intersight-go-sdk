/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// StorageBaseSnapshot Reference marker for volume. It is a point-in-time copy of the volume.
type StorageBaseSnapshot struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Exact date and time at which snapshot was created.
	CreatedTime *time.Time `json:"CreatedTime,omitempty"`
	// Name of the snapshot which represents point-in-time copy of volume.
	Name *string `json:"Name,omitempty"`
	// Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume.
	ProtectionGroupName *string `json:"ProtectionGroupName,omitempty"`
	// Snapshot size represented in bytes.
	Size *int64 `json:"Size,omitempty"`
	// Source object on which the snapshot is created. It is the name of the originating volume.
	Source               *string `json:"Source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseSnapshot StorageBaseSnapshot

// NewStorageBaseSnapshot instantiates a new StorageBaseSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseSnapshot(classId string, objectType string) *StorageBaseSnapshot {
	this := StorageBaseSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseSnapshotWithDefaults instantiates a new StorageBaseSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseSnapshotWithDefaults() *StorageBaseSnapshot {
	this := StorageBaseSnapshot{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshot) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshot) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *StorageBaseSnapshot) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseSnapshot) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseSnapshot) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseSnapshot) SetName(v string) {
	o.Name = &v
}

// GetProtectionGroupName returns the ProtectionGroupName field value if set, zero value otherwise.
func (o *StorageBaseSnapshot) GetProtectionGroupName() string {
	if o == nil || o.ProtectionGroupName == nil {
		var ret string
		return ret
	}
	return *o.ProtectionGroupName
}

// GetProtectionGroupNameOk returns a tuple with the ProtectionGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetProtectionGroupNameOk() (*string, bool) {
	if o == nil || o.ProtectionGroupName == nil {
		return nil, false
	}
	return o.ProtectionGroupName, true
}

// HasProtectionGroupName returns a boolean if a field has been set.
func (o *StorageBaseSnapshot) HasProtectionGroupName() bool {
	if o != nil && o.ProtectionGroupName != nil {
		return true
	}

	return false
}

// SetProtectionGroupName gets a reference to the given string and assigns it to the ProtectionGroupName field.
func (o *StorageBaseSnapshot) SetProtectionGroupName(v string) {
	o.ProtectionGroupName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageBaseSnapshot) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageBaseSnapshot) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageBaseSnapshot) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageBaseSnapshot) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshot) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageBaseSnapshot) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageBaseSnapshot) SetSource(v string) {
	o.Source = &v
}

func (o StorageBaseSnapshot) MarshalJSON() ([]byte, error) {
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
	if o.CreatedTime != nil {
		toSerialize["CreatedTime"] = o.CreatedTime
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProtectionGroupName != nil {
		toSerialize["ProtectionGroupName"] = o.ProtectionGroupName
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Exact date and time at which snapshot was created.
		CreatedTime *time.Time `json:"CreatedTime,omitempty"`
		// Name of the snapshot which represents point-in-time copy of volume.
		Name *string `json:"Name,omitempty"`
		// Name of the protection group to which the snapshot belongs. Value is empty, if the snapshot is created directly on volume.
		ProtectionGroupName *string `json:"ProtectionGroupName,omitempty"`
		// Snapshot size represented in bytes.
		Size *int64 `json:"Size,omitempty"`
		// Source object on which the snapshot is created. It is the name of the originating volume.
		Source *string `json:"Source,omitempty"`
	}

	varStorageBaseSnapshotWithoutEmbeddedStruct := StorageBaseSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseSnapshot := _StorageBaseSnapshot{}
		varStorageBaseSnapshot.ClassId = varStorageBaseSnapshotWithoutEmbeddedStruct.ClassId
		varStorageBaseSnapshot.ObjectType = varStorageBaseSnapshotWithoutEmbeddedStruct.ObjectType
		varStorageBaseSnapshot.CreatedTime = varStorageBaseSnapshotWithoutEmbeddedStruct.CreatedTime
		varStorageBaseSnapshot.Name = varStorageBaseSnapshotWithoutEmbeddedStruct.Name
		varStorageBaseSnapshot.ProtectionGroupName = varStorageBaseSnapshotWithoutEmbeddedStruct.ProtectionGroupName
		varStorageBaseSnapshot.Size = varStorageBaseSnapshotWithoutEmbeddedStruct.Size
		varStorageBaseSnapshot.Source = varStorageBaseSnapshotWithoutEmbeddedStruct.Source
		*o = StorageBaseSnapshot(varStorageBaseSnapshot)
	} else {
		return err
	}

	varStorageBaseSnapshot := _StorageBaseSnapshot{}

	err = json.Unmarshal(bytes, &varStorageBaseSnapshot)
	if err == nil {
		o.MoBaseMo = varStorageBaseSnapshot.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProtectionGroupName")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Source")

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

type NullableStorageBaseSnapshot struct {
	value *StorageBaseSnapshot
	isSet bool
}

func (v NullableStorageBaseSnapshot) Get() *StorageBaseSnapshot {
	return v.value
}

func (v *NullableStorageBaseSnapshot) Set(val *StorageBaseSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseSnapshot(val *StorageBaseSnapshot) *NullableStorageBaseSnapshot {
	return &NullableStorageBaseSnapshot{value: val, isSet: true}
}

func (v NullableStorageBaseSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
