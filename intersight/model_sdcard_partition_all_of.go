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

// SdcardPartitionAllOf Definition of the list of properties defined in 'sdcard.Partition', excluding properties defined in parent classes.
type SdcardPartitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This specifies the type of the partition. Allowed values are OS, Utility. * `OS` - This partition contains virtual drives where user can install operating system. * `Utility` - This partition contains virtual drives for utilities such as SCU, HUU, Drivers and Diagnostics.
	Type                 *string              `json:"Type,omitempty"`
	VirtualDrives        []SdcardVirtualDrive `json:"VirtualDrives,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdcardPartitionAllOf SdcardPartitionAllOf

// NewSdcardPartitionAllOf instantiates a new SdcardPartitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardPartitionAllOf(classId string, objectType string) *SdcardPartitionAllOf {
	this := SdcardPartitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "OS"
	this.Type = &type_
	return &this
}

// NewSdcardPartitionAllOfWithDefaults instantiates a new SdcardPartitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardPartitionAllOfWithDefaults() *SdcardPartitionAllOf {
	this := SdcardPartitionAllOf{}
	var classId string = "sdcard.Partition"
	this.ClassId = classId
	var objectType string = "sdcard.Partition"
	this.ObjectType = objectType
	var type_ string = "OS"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdcardPartitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdcardPartitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdcardPartitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdcardPartitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdcardPartitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdcardPartitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SdcardPartitionAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPartitionAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SdcardPartitionAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SdcardPartitionAllOf) SetType(v string) {
	o.Type = &v
}

// GetVirtualDrives returns the VirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdcardPartitionAllOf) GetVirtualDrives() []SdcardVirtualDrive {
	if o == nil {
		var ret []SdcardVirtualDrive
		return ret
	}
	return o.VirtualDrives
}

// GetVirtualDrivesOk returns a tuple with the VirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdcardPartitionAllOf) GetVirtualDrivesOk() (*[]SdcardVirtualDrive, bool) {
	if o == nil || o.VirtualDrives == nil {
		return nil, false
	}
	return &o.VirtualDrives, true
}

// HasVirtualDrives returns a boolean if a field has been set.
func (o *SdcardPartitionAllOf) HasVirtualDrives() bool {
	if o != nil && o.VirtualDrives != nil {
		return true
	}

	return false
}

// SetVirtualDrives gets a reference to the given []SdcardVirtualDrive and assigns it to the VirtualDrives field.
func (o *SdcardPartitionAllOf) SetVirtualDrives(v []SdcardVirtualDrive) {
	o.VirtualDrives = v
}

func (o SdcardPartitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDrives != nil {
		toSerialize["VirtualDrives"] = o.VirtualDrives
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdcardPartitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdcardPartitionAllOf := _SdcardPartitionAllOf{}

	if err = json.Unmarshal(bytes, &varSdcardPartitionAllOf); err == nil {
		*o = SdcardPartitionAllOf(varSdcardPartitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDrives")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdcardPartitionAllOf struct {
	value *SdcardPartitionAllOf
	isSet bool
}

func (v NullableSdcardPartitionAllOf) Get() *SdcardPartitionAllOf {
	return v.value
}

func (v *NullableSdcardPartitionAllOf) Set(val *SdcardPartitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardPartitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardPartitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardPartitionAllOf(val *SdcardPartitionAllOf) *NullableSdcardPartitionAllOf {
	return &NullableSdcardPartitionAllOf{value: val, isSet: true}
}

func (v NullableSdcardPartitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardPartitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
