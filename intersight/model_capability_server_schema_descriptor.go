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
	"reflect"
	"strings"
)

// CapabilityServerSchemaDescriptor Descriptor that identifies the server's redfish locatorled using cimc firmware info.
type CapabilityServerSchemaDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Redfish property name for the server.
	LocatorLedName *string `json:"LocatorLedName,omitempty"`
	// Redfish version information for the server.
	RedfishSchema        *string `json:"RedfishSchema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityServerSchemaDescriptor CapabilityServerSchemaDescriptor

// NewCapabilityServerSchemaDescriptor instantiates a new CapabilityServerSchemaDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityServerSchemaDescriptor(classId string, objectType string) *CapabilityServerSchemaDescriptor {
	this := CapabilityServerSchemaDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityServerSchemaDescriptorWithDefaults instantiates a new CapabilityServerSchemaDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityServerSchemaDescriptorWithDefaults() *CapabilityServerSchemaDescriptor {
	this := CapabilityServerSchemaDescriptor{}
	var classId string = "capability.ServerSchemaDescriptor"
	this.ClassId = classId
	var objectType string = "capability.ServerSchemaDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityServerSchemaDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerSchemaDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityServerSchemaDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityServerSchemaDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityServerSchemaDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityServerSchemaDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocatorLedName returns the LocatorLedName field value if set, zero value otherwise.
func (o *CapabilityServerSchemaDescriptor) GetLocatorLedName() string {
	if o == nil || o.LocatorLedName == nil {
		var ret string
		return ret
	}
	return *o.LocatorLedName
}

// GetLocatorLedNameOk returns a tuple with the LocatorLedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerSchemaDescriptor) GetLocatorLedNameOk() (*string, bool) {
	if o == nil || o.LocatorLedName == nil {
		return nil, false
	}
	return o.LocatorLedName, true
}

// HasLocatorLedName returns a boolean if a field has been set.
func (o *CapabilityServerSchemaDescriptor) HasLocatorLedName() bool {
	if o != nil && o.LocatorLedName != nil {
		return true
	}

	return false
}

// SetLocatorLedName gets a reference to the given string and assigns it to the LocatorLedName field.
func (o *CapabilityServerSchemaDescriptor) SetLocatorLedName(v string) {
	o.LocatorLedName = &v
}

// GetRedfishSchema returns the RedfishSchema field value if set, zero value otherwise.
func (o *CapabilityServerSchemaDescriptor) GetRedfishSchema() string {
	if o == nil || o.RedfishSchema == nil {
		var ret string
		return ret
	}
	return *o.RedfishSchema
}

// GetRedfishSchemaOk returns a tuple with the RedfishSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityServerSchemaDescriptor) GetRedfishSchemaOk() (*string, bool) {
	if o == nil || o.RedfishSchema == nil {
		return nil, false
	}
	return o.RedfishSchema, true
}

// HasRedfishSchema returns a boolean if a field has been set.
func (o *CapabilityServerSchemaDescriptor) HasRedfishSchema() bool {
	if o != nil && o.RedfishSchema != nil {
		return true
	}

	return false
}

// SetRedfishSchema gets a reference to the given string and assigns it to the RedfishSchema field.
func (o *CapabilityServerSchemaDescriptor) SetRedfishSchema(v string) {
	o.RedfishSchema = &v
}

func (o CapabilityServerSchemaDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return []byte{}, errCapabilityHardwareDescriptor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocatorLedName != nil {
		toSerialize["LocatorLedName"] = o.LocatorLedName
	}
	if o.RedfishSchema != nil {
		toSerialize["RedfishSchema"] = o.RedfishSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityServerSchemaDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilityServerSchemaDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Redfish property name for the server.
		LocatorLedName *string `json:"LocatorLedName,omitempty"`
		// Redfish version information for the server.
		RedfishSchema *string `json:"RedfishSchema,omitempty"`
	}

	varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct := CapabilityServerSchemaDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityServerSchemaDescriptor := _CapabilityServerSchemaDescriptor{}
		varCapabilityServerSchemaDescriptor.ClassId = varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityServerSchemaDescriptor.ObjectType = varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityServerSchemaDescriptor.LocatorLedName = varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct.LocatorLedName
		varCapabilityServerSchemaDescriptor.RedfishSchema = varCapabilityServerSchemaDescriptorWithoutEmbeddedStruct.RedfishSchema
		*o = CapabilityServerSchemaDescriptor(varCapabilityServerSchemaDescriptor)
	} else {
		return err
	}

	varCapabilityServerSchemaDescriptor := _CapabilityServerSchemaDescriptor{}

	err = json.Unmarshal(bytes, &varCapabilityServerSchemaDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilityServerSchemaDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocatorLedName")
		delete(additionalProperties, "RedfishSchema")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilityServerSchemaDescriptor struct {
	value *CapabilityServerSchemaDescriptor
	isSet bool
}

func (v NullableCapabilityServerSchemaDescriptor) Get() *CapabilityServerSchemaDescriptor {
	return v.value
}

func (v *NullableCapabilityServerSchemaDescriptor) Set(val *CapabilityServerSchemaDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityServerSchemaDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityServerSchemaDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityServerSchemaDescriptor(val *CapabilityServerSchemaDescriptor) *NullableCapabilityServerSchemaDescriptor {
	return &NullableCapabilityServerSchemaDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityServerSchemaDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityServerSchemaDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
