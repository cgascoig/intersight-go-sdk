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

// StorageNetAppStorageUtilizationAllOf Definition of the list of properties defined in 'storage.NetAppStorageUtilization', excluding properties defined in parent classes.
type StorageNetAppStorageUtilizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Total logical used capacity, represented in bytes.
	LogicalUsed *int64 `json:"LogicalUsed,omitempty"`
	// Total savings capacity, represented in bytes.
	Savings              *int64 `json:"Savings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppStorageUtilizationAllOf StorageNetAppStorageUtilizationAllOf

// NewStorageNetAppStorageUtilizationAllOf instantiates a new StorageNetAppStorageUtilizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppStorageUtilizationAllOf(classId string, objectType string) *StorageNetAppStorageUtilizationAllOf {
	this := StorageNetAppStorageUtilizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppStorageUtilizationAllOfWithDefaults instantiates a new StorageNetAppStorageUtilizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppStorageUtilizationAllOfWithDefaults() *StorageNetAppStorageUtilizationAllOf {
	this := StorageNetAppStorageUtilizationAllOf{}
	var classId string = "storage.NetAppStorageUtilization"
	this.ClassId = classId
	var objectType string = "storage.NetAppStorageUtilization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppStorageUtilizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppStorageUtilizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppStorageUtilizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppStorageUtilizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLogicalUsed returns the LogicalUsed field value if set, zero value otherwise.
func (o *StorageNetAppStorageUtilizationAllOf) GetLogicalUsed() int64 {
	if o == nil || o.LogicalUsed == nil {
		var ret int64
		return ret
	}
	return *o.LogicalUsed
}

// GetLogicalUsedOk returns a tuple with the LogicalUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilizationAllOf) GetLogicalUsedOk() (*int64, bool) {
	if o == nil || o.LogicalUsed == nil {
		return nil, false
	}
	return o.LogicalUsed, true
}

// HasLogicalUsed returns a boolean if a field has been set.
func (o *StorageNetAppStorageUtilizationAllOf) HasLogicalUsed() bool {
	if o != nil && o.LogicalUsed != nil {
		return true
	}

	return false
}

// SetLogicalUsed gets a reference to the given int64 and assigns it to the LogicalUsed field.
func (o *StorageNetAppStorageUtilizationAllOf) SetLogicalUsed(v int64) {
	o.LogicalUsed = &v
}

// GetSavings returns the Savings field value if set, zero value otherwise.
func (o *StorageNetAppStorageUtilizationAllOf) GetSavings() int64 {
	if o == nil || o.Savings == nil {
		var ret int64
		return ret
	}
	return *o.Savings
}

// GetSavingsOk returns a tuple with the Savings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppStorageUtilizationAllOf) GetSavingsOk() (*int64, bool) {
	if o == nil || o.Savings == nil {
		return nil, false
	}
	return o.Savings, true
}

// HasSavings returns a boolean if a field has been set.
func (o *StorageNetAppStorageUtilizationAllOf) HasSavings() bool {
	if o != nil && o.Savings != nil {
		return true
	}

	return false
}

// SetSavings gets a reference to the given int64 and assigns it to the Savings field.
func (o *StorageNetAppStorageUtilizationAllOf) SetSavings(v int64) {
	o.Savings = &v
}

func (o StorageNetAppStorageUtilizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LogicalUsed != nil {
		toSerialize["LogicalUsed"] = o.LogicalUsed
	}
	if o.Savings != nil {
		toSerialize["Savings"] = o.Savings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppStorageUtilizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppStorageUtilizationAllOf := _StorageNetAppStorageUtilizationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppStorageUtilizationAllOf); err == nil {
		*o = StorageNetAppStorageUtilizationAllOf(varStorageNetAppStorageUtilizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LogicalUsed")
		delete(additionalProperties, "Savings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppStorageUtilizationAllOf struct {
	value *StorageNetAppStorageUtilizationAllOf
	isSet bool
}

func (v NullableStorageNetAppStorageUtilizationAllOf) Get() *StorageNetAppStorageUtilizationAllOf {
	return v.value
}

func (v *NullableStorageNetAppStorageUtilizationAllOf) Set(val *StorageNetAppStorageUtilizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppStorageUtilizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppStorageUtilizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppStorageUtilizationAllOf(val *StorageNetAppStorageUtilizationAllOf) *NullableStorageNetAppStorageUtilizationAllOf {
	return &NullableStorageNetAppStorageUtilizationAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppStorageUtilizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppStorageUtilizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
