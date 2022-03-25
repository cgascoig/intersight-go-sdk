/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppLunEventAllOf Definition of the list of properties defined in 'storage.NetAppLunEvent', excluding properties defined in parent classes.
type StorageNetAppLunEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                        `json:"ObjectType"`
	Lun                  *StorageNetAppLunRelationship `json:"Lun,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLunEventAllOf StorageNetAppLunEventAllOf

// NewStorageNetAppLunEventAllOf instantiates a new StorageNetAppLunEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLunEventAllOf(classId string, objectType string) *StorageNetAppLunEventAllOf {
	this := StorageNetAppLunEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLunEventAllOfWithDefaults instantiates a new StorageNetAppLunEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLunEventAllOfWithDefaults() *StorageNetAppLunEventAllOf {
	this := StorageNetAppLunEventAllOf{}
	var classId string = "storage.NetAppLunEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppLunEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLunEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLunEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLunEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLunEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLun returns the Lun field value if set, zero value otherwise.
func (o *StorageNetAppLunEventAllOf) GetLun() StorageNetAppLunRelationship {
	if o == nil || o.Lun == nil {
		var ret StorageNetAppLunRelationship
		return ret
	}
	return *o.Lun
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLunEventAllOf) GetLunOk() (*StorageNetAppLunRelationship, bool) {
	if o == nil || o.Lun == nil {
		return nil, false
	}
	return o.Lun, true
}

// HasLun returns a boolean if a field has been set.
func (o *StorageNetAppLunEventAllOf) HasLun() bool {
	if o != nil && o.Lun != nil {
		return true
	}

	return false
}

// SetLun gets a reference to the given StorageNetAppLunRelationship and assigns it to the Lun field.
func (o *StorageNetAppLunEventAllOf) SetLun(v StorageNetAppLunRelationship) {
	o.Lun = &v
}

func (o StorageNetAppLunEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Lun != nil {
		toSerialize["Lun"] = o.Lun
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppLunEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppLunEventAllOf := _StorageNetAppLunEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppLunEventAllOf); err == nil {
		*o = StorageNetAppLunEventAllOf(varStorageNetAppLunEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Lun")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppLunEventAllOf struct {
	value *StorageNetAppLunEventAllOf
	isSet bool
}

func (v NullableStorageNetAppLunEventAllOf) Get() *StorageNetAppLunEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppLunEventAllOf) Set(val *StorageNetAppLunEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLunEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLunEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLunEventAllOf(val *StorageNetAppLunEventAllOf) *NullableStorageNetAppLunEventAllOf {
	return &NullableStorageNetAppLunEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppLunEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLunEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
