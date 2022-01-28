/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareExcludeComponentPidListTypeAllOf Definition of the list of properties defined in 'firmware.ExcludeComponentPidListType', excluding properties defined in parent classes.
type FirmwareExcludeComponentPidListTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string   `json:"ObjectType"`
	ExcludeLocalDiskList         []string `json:"ExcludeLocalDiskList,omitempty"`
	ExcludeStorageControllerList []string `json:"ExcludeStorageControllerList,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _FirmwareExcludeComponentPidListTypeAllOf FirmwareExcludeComponentPidListTypeAllOf

// NewFirmwareExcludeComponentPidListTypeAllOf instantiates a new FirmwareExcludeComponentPidListTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareExcludeComponentPidListTypeAllOf(classId string, objectType string) *FirmwareExcludeComponentPidListTypeAllOf {
	this := FirmwareExcludeComponentPidListTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareExcludeComponentPidListTypeAllOfWithDefaults instantiates a new FirmwareExcludeComponentPidListTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareExcludeComponentPidListTypeAllOfWithDefaults() *FirmwareExcludeComponentPidListTypeAllOf {
	this := FirmwareExcludeComponentPidListTypeAllOf{}
	var classId string = "firmware.ExcludeComponentPidListType"
	this.ClassId = classId
	var objectType string = "firmware.ExcludeComponentPidListType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareExcludeComponentPidListTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareExcludeComponentPidListTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExcludeLocalDiskList returns the ExcludeLocalDiskList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeLocalDiskList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeLocalDiskList
}

// GetExcludeLocalDiskListOk returns a tuple with the ExcludeLocalDiskList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeLocalDiskListOk() (*[]string, bool) {
	if o == nil || o.ExcludeLocalDiskList == nil {
		return nil, false
	}
	return &o.ExcludeLocalDiskList, true
}

// HasExcludeLocalDiskList returns a boolean if a field has been set.
func (o *FirmwareExcludeComponentPidListTypeAllOf) HasExcludeLocalDiskList() bool {
	if o != nil && o.ExcludeLocalDiskList != nil {
		return true
	}

	return false
}

// SetExcludeLocalDiskList gets a reference to the given []string and assigns it to the ExcludeLocalDiskList field.
func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeLocalDiskList(v []string) {
	o.ExcludeLocalDiskList = v
}

// GetExcludeStorageControllerList returns the ExcludeStorageControllerList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeStorageControllerList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeStorageControllerList
}

// GetExcludeStorageControllerListOk returns a tuple with the ExcludeStorageControllerList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareExcludeComponentPidListTypeAllOf) GetExcludeStorageControllerListOk() (*[]string, bool) {
	if o == nil || o.ExcludeStorageControllerList == nil {
		return nil, false
	}
	return &o.ExcludeStorageControllerList, true
}

// HasExcludeStorageControllerList returns a boolean if a field has been set.
func (o *FirmwareExcludeComponentPidListTypeAllOf) HasExcludeStorageControllerList() bool {
	if o != nil && o.ExcludeStorageControllerList != nil {
		return true
	}

	return false
}

// SetExcludeStorageControllerList gets a reference to the given []string and assigns it to the ExcludeStorageControllerList field.
func (o *FirmwareExcludeComponentPidListTypeAllOf) SetExcludeStorageControllerList(v []string) {
	o.ExcludeStorageControllerList = v
}

func (o FirmwareExcludeComponentPidListTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExcludeLocalDiskList != nil {
		toSerialize["ExcludeLocalDiskList"] = o.ExcludeLocalDiskList
	}
	if o.ExcludeStorageControllerList != nil {
		toSerialize["ExcludeStorageControllerList"] = o.ExcludeStorageControllerList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareExcludeComponentPidListTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareExcludeComponentPidListTypeAllOf := _FirmwareExcludeComponentPidListTypeAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareExcludeComponentPidListTypeAllOf); err == nil {
		*o = FirmwareExcludeComponentPidListTypeAllOf(varFirmwareExcludeComponentPidListTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeLocalDiskList")
		delete(additionalProperties, "ExcludeStorageControllerList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareExcludeComponentPidListTypeAllOf struct {
	value *FirmwareExcludeComponentPidListTypeAllOf
	isSet bool
}

func (v NullableFirmwareExcludeComponentPidListTypeAllOf) Get() *FirmwareExcludeComponentPidListTypeAllOf {
	return v.value
}

func (v *NullableFirmwareExcludeComponentPidListTypeAllOf) Set(val *FirmwareExcludeComponentPidListTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareExcludeComponentPidListTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareExcludeComponentPidListTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareExcludeComponentPidListTypeAllOf(val *FirmwareExcludeComponentPidListTypeAllOf) *NullableFirmwareExcludeComponentPidListTypeAllOf {
	return &NullableFirmwareExcludeComponentPidListTypeAllOf{value: val, isSet: true}
}

func (v NullableFirmwareExcludeComponentPidListTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareExcludeComponentPidListTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
