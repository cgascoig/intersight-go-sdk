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
)

// StorageNetAppClusterEventAllOf Definition of the list of properties defined in 'storage.NetAppClusterEvent', excluding properties defined in parent classes.
type StorageNetAppClusterEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                            `json:"ObjectType"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterEventAllOf StorageNetAppClusterEventAllOf

// NewStorageNetAppClusterEventAllOf instantiates a new StorageNetAppClusterEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterEventAllOf(classId string, objectType string) *StorageNetAppClusterEventAllOf {
	this := StorageNetAppClusterEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterEventAllOfWithDefaults instantiates a new StorageNetAppClusterEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterEventAllOfWithDefaults() *StorageNetAppClusterEventAllOf {
	this := StorageNetAppClusterEventAllOf{}
	var classId string = "storage.NetAppClusterEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppClusterEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppClusterEventAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterEventAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppClusterEventAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppClusterEventAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppClusterEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppClusterEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppClusterEventAllOf := _StorageNetAppClusterEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppClusterEventAllOf); err == nil {
		*o = StorageNetAppClusterEventAllOf(varStorageNetAppClusterEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppClusterEventAllOf struct {
	value *StorageNetAppClusterEventAllOf
	isSet bool
}

func (v NullableStorageNetAppClusterEventAllOf) Get() *StorageNetAppClusterEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppClusterEventAllOf) Set(val *StorageNetAppClusterEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterEventAllOf(val *StorageNetAppClusterEventAllOf) *NullableStorageNetAppClusterEventAllOf {
	return &NullableStorageNetAppClusterEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
