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
	"reflect"
	"strings"
)

// SshPolicyInventory Secure shell policy on the endpoint.
type SshPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of SSH service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// Port used for secure shell access.
	Port *int64 `json:"Port,omitempty"`
	// Number of seconds to wait before the system considers a SSH request to have timed out.
	Timeout              *int64                `json:"Timeout,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SshPolicyInventory SshPolicyInventory

// NewSshPolicyInventory instantiates a new SshPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshPolicyInventory(classId string, objectType string) *SshPolicyInventory {
	this := SshPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSshPolicyInventoryWithDefaults instantiates a new SshPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshPolicyInventoryWithDefaults() *SshPolicyInventory {
	this := SshPolicyInventory{}
	var classId string = "ssh.PolicyInventory"
	this.ClassId = classId
	var objectType string = "ssh.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SshPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SshPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SshPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SshPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SshPolicyInventory) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SshPolicyInventory) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SshPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SshPolicyInventory) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SshPolicyInventory) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *SshPolicyInventory) SetPort(v int64) {
	o.Port = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SshPolicyInventory) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SshPolicyInventory) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SshPolicyInventory) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *SshPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *SshPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *SshPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o SshPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SshPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type SshPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of SSH service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// Port used for secure shell access.
		Port *int64 `json:"Port,omitempty"`
		// Number of seconds to wait before the system considers a SSH request to have timed out.
		Timeout  *int64                `json:"Timeout,omitempty"`
		TargetMo *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varSshPolicyInventoryWithoutEmbeddedStruct := SshPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSshPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varSshPolicyInventory := _SshPolicyInventory{}
		varSshPolicyInventory.ClassId = varSshPolicyInventoryWithoutEmbeddedStruct.ClassId
		varSshPolicyInventory.ObjectType = varSshPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varSshPolicyInventory.Enabled = varSshPolicyInventoryWithoutEmbeddedStruct.Enabled
		varSshPolicyInventory.Port = varSshPolicyInventoryWithoutEmbeddedStruct.Port
		varSshPolicyInventory.Timeout = varSshPolicyInventoryWithoutEmbeddedStruct.Timeout
		varSshPolicyInventory.TargetMo = varSshPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = SshPolicyInventory(varSshPolicyInventory)
	} else {
		return err
	}

	varSshPolicyInventory := _SshPolicyInventory{}

	err = json.Unmarshal(bytes, &varSshPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varSshPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Timeout")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableSshPolicyInventory struct {
	value *SshPolicyInventory
	isSet bool
}

func (v NullableSshPolicyInventory) Get() *SshPolicyInventory {
	return v.value
}

func (v *NullableSshPolicyInventory) Set(val *SshPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableSshPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableSshPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshPolicyInventory(val *SshPolicyInventory) *NullableSshPolicyInventory {
	return &NullableSshPolicyInventory{value: val, isSet: true}
}

func (v NullableSshPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
