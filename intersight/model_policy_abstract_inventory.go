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

// PolicyAbstractInventory A base abstract class used to represent all inventoried objects from a profile associated resource.
type PolicyAbstractInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Device ID of the entity from where inventory is reported.
	DeviceMoId           *string `json:"DeviceMoId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractInventory PolicyAbstractInventory

// NewPolicyAbstractInventory instantiates a new PolicyAbstractInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractInventory(classId string, objectType string) *PolicyAbstractInventory {
	this := PolicyAbstractInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyAbstractInventoryWithDefaults instantiates a new PolicyAbstractInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractInventoryWithDefaults() *PolicyAbstractInventory {
	this := PolicyAbstractInventory{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceMoId returns the DeviceMoId field value if set, zero value otherwise.
func (o *PolicyAbstractInventory) GetDeviceMoId() string {
	if o == nil || o.DeviceMoId == nil {
		var ret string
		return ret
	}
	return *o.DeviceMoId
}

// GetDeviceMoIdOk returns a tuple with the DeviceMoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractInventory) GetDeviceMoIdOk() (*string, bool) {
	if o == nil || o.DeviceMoId == nil {
		return nil, false
	}
	return o.DeviceMoId, true
}

// HasDeviceMoId returns a boolean if a field has been set.
func (o *PolicyAbstractInventory) HasDeviceMoId() bool {
	if o != nil && o.DeviceMoId != nil {
		return true
	}

	return false
}

// SetDeviceMoId gets a reference to the given string and assigns it to the DeviceMoId field.
func (o *PolicyAbstractInventory) SetDeviceMoId(v string) {
	o.DeviceMoId = &v
}

func (o PolicyAbstractInventory) MarshalJSON() ([]byte, error) {
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
	if o.DeviceMoId != nil {
		toSerialize["DeviceMoId"] = o.DeviceMoId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractInventory) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyAbstractInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Device ID of the entity from where inventory is reported.
		DeviceMoId *string `json:"DeviceMoId,omitempty"`
	}

	varPolicyAbstractInventoryWithoutEmbeddedStruct := PolicyAbstractInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyAbstractInventoryWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractInventory := _PolicyAbstractInventory{}
		varPolicyAbstractInventory.ClassId = varPolicyAbstractInventoryWithoutEmbeddedStruct.ClassId
		varPolicyAbstractInventory.ObjectType = varPolicyAbstractInventoryWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractInventory.DeviceMoId = varPolicyAbstractInventoryWithoutEmbeddedStruct.DeviceMoId
		*o = PolicyAbstractInventory(varPolicyAbstractInventory)
	} else {
		return err
	}

	varPolicyAbstractInventory := _PolicyAbstractInventory{}

	err = json.Unmarshal(bytes, &varPolicyAbstractInventory)
	if err == nil {
		o.MoBaseMo = varPolicyAbstractInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceMoId")

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

type NullablePolicyAbstractInventory struct {
	value *PolicyAbstractInventory
	isSet bool
}

func (v NullablePolicyAbstractInventory) Get() *PolicyAbstractInventory {
	return v.value
}

func (v *NullablePolicyAbstractInventory) Set(val *PolicyAbstractInventory) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractInventory) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractInventory(val *PolicyAbstractInventory) *NullablePolicyAbstractInventory {
	return &NullablePolicyAbstractInventory{value: val, isSet: true}
}

func (v NullablePolicyAbstractInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
