/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppDiskEventAllOf Definition of the list of properties defined in 'storage.NetAppDiskEvent', excluding properties defined in parent classes.
type StorageNetAppDiskEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Disk                 *StorageNetAppBaseDiskRelationship `json:"Disk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppDiskEventAllOf StorageNetAppDiskEventAllOf

// NewStorageNetAppDiskEventAllOf instantiates a new StorageNetAppDiskEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppDiskEventAllOf(classId string, objectType string) *StorageNetAppDiskEventAllOf {
	this := StorageNetAppDiskEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppDiskEventAllOfWithDefaults instantiates a new StorageNetAppDiskEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppDiskEventAllOfWithDefaults() *StorageNetAppDiskEventAllOf {
	this := StorageNetAppDiskEventAllOf{}
	var classId string = "storage.NetAppDiskEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppDiskEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppDiskEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDiskEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppDiskEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppDiskEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDiskEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppDiskEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *StorageNetAppDiskEventAllOf) GetDisk() StorageNetAppBaseDiskRelationship {
	if o == nil || o.Disk == nil {
		var ret StorageNetAppBaseDiskRelationship
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDiskEventAllOf) GetDiskOk() (*StorageNetAppBaseDiskRelationship, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *StorageNetAppDiskEventAllOf) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given StorageNetAppBaseDiskRelationship and assigns it to the Disk field.
func (o *StorageNetAppDiskEventAllOf) SetDisk(v StorageNetAppBaseDiskRelationship) {
	o.Disk = &v
}

func (o StorageNetAppDiskEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Disk != nil {
		toSerialize["Disk"] = o.Disk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppDiskEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppDiskEventAllOf := _StorageNetAppDiskEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppDiskEventAllOf); err == nil {
		*o = StorageNetAppDiskEventAllOf(varStorageNetAppDiskEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Disk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppDiskEventAllOf struct {
	value *StorageNetAppDiskEventAllOf
	isSet bool
}

func (v NullableStorageNetAppDiskEventAllOf) Get() *StorageNetAppDiskEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppDiskEventAllOf) Set(val *StorageNetAppDiskEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDiskEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDiskEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDiskEventAllOf(val *StorageNetAppDiskEventAllOf) *NullableStorageNetAppDiskEventAllOf {
	return &NullableStorageNetAppDiskEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppDiskEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDiskEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
