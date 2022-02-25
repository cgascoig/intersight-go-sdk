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
	"reflect"
	"strings"
)

// ComputeRackUnitIdentity Identity object that uniquely represents a server object under a DR.
type ComputeRackUnitIdentity struct {
	EquipmentPhysicalIdentity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Serial Identifier of an adapter participating in SWM.
	AdapterSerial        *string `json:"AdapterSerial,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeRackUnitIdentity ComputeRackUnitIdentity

// NewComputeRackUnitIdentity instantiates a new ComputeRackUnitIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeRackUnitIdentity(classId string, objectType string) *ComputeRackUnitIdentity {
	this := ComputeRackUnitIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewComputeRackUnitIdentityWithDefaults instantiates a new ComputeRackUnitIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeRackUnitIdentityWithDefaults() *ComputeRackUnitIdentity {
	this := ComputeRackUnitIdentity{}
	var classId string = "compute.RackUnitIdentity"
	this.ClassId = classId
	var objectType string = "compute.RackUnitIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeRackUnitIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeRackUnitIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeRackUnitIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeRackUnitIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdapterSerial returns the AdapterSerial field value if set, zero value otherwise.
func (o *ComputeRackUnitIdentity) GetAdapterSerial() string {
	if o == nil || o.AdapterSerial == nil {
		var ret string
		return ret
	}
	return *o.AdapterSerial
}

// GetAdapterSerialOk returns a tuple with the AdapterSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeRackUnitIdentity) GetAdapterSerialOk() (*string, bool) {
	if o == nil || o.AdapterSerial == nil {
		return nil, false
	}
	return o.AdapterSerial, true
}

// HasAdapterSerial returns a boolean if a field has been set.
func (o *ComputeRackUnitIdentity) HasAdapterSerial() bool {
	if o != nil && o.AdapterSerial != nil {
		return true
	}

	return false
}

// SetAdapterSerial gets a reference to the given string and assigns it to the AdapterSerial field.
func (o *ComputeRackUnitIdentity) SetAdapterSerial(v string) {
	o.AdapterSerial = &v
}

func (o ComputeRackUnitIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentPhysicalIdentity, errEquipmentPhysicalIdentity := json.Marshal(o.EquipmentPhysicalIdentity)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	errEquipmentPhysicalIdentity = json.Unmarshal([]byte(serializedEquipmentPhysicalIdentity), &toSerialize)
	if errEquipmentPhysicalIdentity != nil {
		return []byte{}, errEquipmentPhysicalIdentity
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdapterSerial != nil {
		toSerialize["AdapterSerial"] = o.AdapterSerial
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeRackUnitIdentity) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeRackUnitIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Serial Identifier of an adapter participating in SWM.
		AdapterSerial *string `json:"AdapterSerial,omitempty"`
	}

	varComputeRackUnitIdentityWithoutEmbeddedStruct := ComputeRackUnitIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeRackUnitIdentityWithoutEmbeddedStruct)
	if err == nil {
		varComputeRackUnitIdentity := _ComputeRackUnitIdentity{}
		varComputeRackUnitIdentity.ClassId = varComputeRackUnitIdentityWithoutEmbeddedStruct.ClassId
		varComputeRackUnitIdentity.ObjectType = varComputeRackUnitIdentityWithoutEmbeddedStruct.ObjectType
		varComputeRackUnitIdentity.AdapterSerial = varComputeRackUnitIdentityWithoutEmbeddedStruct.AdapterSerial
		*o = ComputeRackUnitIdentity(varComputeRackUnitIdentity)
	} else {
		return err
	}

	varComputeRackUnitIdentity := _ComputeRackUnitIdentity{}

	err = json.Unmarshal(bytes, &varComputeRackUnitIdentity)
	if err == nil {
		o.EquipmentPhysicalIdentity = varComputeRackUnitIdentity.EquipmentPhysicalIdentity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterSerial")

		// remove fields from embedded structs
		reflectEquipmentPhysicalIdentity := reflect.ValueOf(o.EquipmentPhysicalIdentity)
		for i := 0; i < reflectEquipmentPhysicalIdentity.Type().NumField(); i++ {
			t := reflectEquipmentPhysicalIdentity.Type().Field(i)

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

type NullableComputeRackUnitIdentity struct {
	value *ComputeRackUnitIdentity
	isSet bool
}

func (v NullableComputeRackUnitIdentity) Get() *ComputeRackUnitIdentity {
	return v.value
}

func (v *NullableComputeRackUnitIdentity) Set(val *ComputeRackUnitIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeRackUnitIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeRackUnitIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeRackUnitIdentity(val *ComputeRackUnitIdentity) *NullableComputeRackUnitIdentity {
	return &NullableComputeRackUnitIdentity{value: val, isSet: true}
}

func (v NullableComputeRackUnitIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeRackUnitIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
