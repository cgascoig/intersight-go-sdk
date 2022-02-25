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

// OsVmwareParametersAllOf Definition of the list of properties defined in 'os.VmwareParameters', excluding properties defined in parent classes.
type OsVmwareParametersAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the VLAN ID in which the ESXi host is turned on. Valid values ranges between 1 – 4095.
	Vlanid               *int64 `json:"Vlanid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsVmwareParametersAllOf OsVmwareParametersAllOf

// NewOsVmwareParametersAllOf instantiates a new OsVmwareParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsVmwareParametersAllOf(classId string, objectType string) *OsVmwareParametersAllOf {
	this := OsVmwareParametersAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsVmwareParametersAllOfWithDefaults instantiates a new OsVmwareParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsVmwareParametersAllOfWithDefaults() *OsVmwareParametersAllOf {
	this := OsVmwareParametersAllOf{}
	var classId string = "os.VmwareParameters"
	this.ClassId = classId
	var objectType string = "os.VmwareParameters"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsVmwareParametersAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsVmwareParametersAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsVmwareParametersAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsVmwareParametersAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsVmwareParametersAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsVmwareParametersAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVlanid returns the Vlanid field value if set, zero value otherwise.
func (o *OsVmwareParametersAllOf) GetVlanid() int64 {
	if o == nil || o.Vlanid == nil {
		var ret int64
		return ret
	}
	return *o.Vlanid
}

// GetVlanidOk returns a tuple with the Vlanid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVmwareParametersAllOf) GetVlanidOk() (*int64, bool) {
	if o == nil || o.Vlanid == nil {
		return nil, false
	}
	return o.Vlanid, true
}

// HasVlanid returns a boolean if a field has been set.
func (o *OsVmwareParametersAllOf) HasVlanid() bool {
	if o != nil && o.Vlanid != nil {
		return true
	}

	return false
}

// SetVlanid gets a reference to the given int64 and assigns it to the Vlanid field.
func (o *OsVmwareParametersAllOf) SetVlanid(v int64) {
	o.Vlanid = &v
}

func (o OsVmwareParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Vlanid != nil {
		toSerialize["Vlanid"] = o.Vlanid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsVmwareParametersAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsVmwareParametersAllOf := _OsVmwareParametersAllOf{}

	if err = json.Unmarshal(bytes, &varOsVmwareParametersAllOf); err == nil {
		*o = OsVmwareParametersAllOf(varOsVmwareParametersAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Vlanid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsVmwareParametersAllOf struct {
	value *OsVmwareParametersAllOf
	isSet bool
}

func (v NullableOsVmwareParametersAllOf) Get() *OsVmwareParametersAllOf {
	return v.value
}

func (v *NullableOsVmwareParametersAllOf) Set(val *OsVmwareParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsVmwareParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsVmwareParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsVmwareParametersAllOf(val *OsVmwareParametersAllOf) *NullableOsVmwareParametersAllOf {
	return &NullableOsVmwareParametersAllOf{value: val, isSet: true}
}

func (v NullableOsVmwareParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsVmwareParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
