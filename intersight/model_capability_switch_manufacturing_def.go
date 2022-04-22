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

// CapabilitySwitchManufacturingDef Switch/Fabric-Interconnect manufacturing def properties.
type CapabilitySwitchManufacturingDef struct {
	CapabilitySwitchCapabilityDef
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Caption for Switch/Fabric-Interconnect.
	Caption *string `json:"Caption,omitempty"`
	// Description for Switch/Fabric-Interconnect.
	Description *string `json:"Description,omitempty"`
	// Part Number for Switch/Fabric-Interconnect.
	PartNumber *string `json:"PartNumber,omitempty"`
	// Product Name for Switch/Fabric-Interconnect.
	ProductName          *string `json:"ProductName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchManufacturingDef CapabilitySwitchManufacturingDef

// NewCapabilitySwitchManufacturingDef instantiates a new CapabilitySwitchManufacturingDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchManufacturingDef(classId string, objectType string) *CapabilitySwitchManufacturingDef {
	this := CapabilitySwitchManufacturingDef{}
	this.ClassId = classId
	this.ObjectType = objectType
	var pid string = "UCS-FI-6454"
	this.Pid = &pid
	return &this
}

// NewCapabilitySwitchManufacturingDefWithDefaults instantiates a new CapabilitySwitchManufacturingDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchManufacturingDefWithDefaults() *CapabilitySwitchManufacturingDef {
	this := CapabilitySwitchManufacturingDef{}
	var classId string = "capability.SwitchManufacturingDef"
	this.ClassId = classId
	var objectType string = "capability.SwitchManufacturingDef"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchManufacturingDef) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchManufacturingDef) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchManufacturingDef) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchManufacturingDef) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *CapabilitySwitchManufacturingDef) SetCaption(v string) {
	o.Caption = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilitySwitchManufacturingDef) SetDescription(v string) {
	o.Description = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *CapabilitySwitchManufacturingDef) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *CapabilitySwitchManufacturingDef) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchManufacturingDef) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *CapabilitySwitchManufacturingDef) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *CapabilitySwitchManufacturingDef) SetProductName(v string) {
	o.ProductName = &v
}

func (o CapabilitySwitchManufacturingDef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilitySwitchCapabilityDef, errCapabilitySwitchCapabilityDef := json.Marshal(o.CapabilitySwitchCapabilityDef)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	errCapabilitySwitchCapabilityDef = json.Unmarshal([]byte(serializedCapabilitySwitchCapabilityDef), &toSerialize)
	if errCapabilitySwitchCapabilityDef != nil {
		return []byte{}, errCapabilitySwitchCapabilityDef
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Caption != nil {
		toSerialize["Caption"] = o.Caption
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchManufacturingDef) UnmarshalJSON(bytes []byte) (err error) {
	type CapabilitySwitchManufacturingDefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Caption for Switch/Fabric-Interconnect.
		Caption *string `json:"Caption,omitempty"`
		// Description for Switch/Fabric-Interconnect.
		Description *string `json:"Description,omitempty"`
		// Part Number for Switch/Fabric-Interconnect.
		PartNumber *string `json:"PartNumber,omitempty"`
		// Product Name for Switch/Fabric-Interconnect.
		ProductName *string `json:"ProductName,omitempty"`
	}

	varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct := CapabilitySwitchManufacturingDefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct)
	if err == nil {
		varCapabilitySwitchManufacturingDef := _CapabilitySwitchManufacturingDef{}
		varCapabilitySwitchManufacturingDef.ClassId = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ClassId
		varCapabilitySwitchManufacturingDef.ObjectType = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ObjectType
		varCapabilitySwitchManufacturingDef.Caption = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.Caption
		varCapabilitySwitchManufacturingDef.Description = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.Description
		varCapabilitySwitchManufacturingDef.PartNumber = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.PartNumber
		varCapabilitySwitchManufacturingDef.ProductName = varCapabilitySwitchManufacturingDefWithoutEmbeddedStruct.ProductName
		*o = CapabilitySwitchManufacturingDef(varCapabilitySwitchManufacturingDef)
	} else {
		return err
	}

	varCapabilitySwitchManufacturingDef := _CapabilitySwitchManufacturingDef{}

	err = json.Unmarshal(bytes, &varCapabilitySwitchManufacturingDef)
	if err == nil {
		o.CapabilitySwitchCapabilityDef = varCapabilitySwitchManufacturingDef.CapabilitySwitchCapabilityDef
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Caption")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "ProductName")

		// remove fields from embedded structs
		reflectCapabilitySwitchCapabilityDef := reflect.ValueOf(o.CapabilitySwitchCapabilityDef)
		for i := 0; i < reflectCapabilitySwitchCapabilityDef.Type().NumField(); i++ {
			t := reflectCapabilitySwitchCapabilityDef.Type().Field(i)

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

type NullableCapabilitySwitchManufacturingDef struct {
	value *CapabilitySwitchManufacturingDef
	isSet bool
}

func (v NullableCapabilitySwitchManufacturingDef) Get() *CapabilitySwitchManufacturingDef {
	return v.value
}

func (v *NullableCapabilitySwitchManufacturingDef) Set(val *CapabilitySwitchManufacturingDef) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchManufacturingDef) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchManufacturingDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchManufacturingDef(val *CapabilitySwitchManufacturingDef) *NullableCapabilitySwitchManufacturingDef {
	return &NullableCapabilitySwitchManufacturingDef{value: val, isSet: true}
}

func (v NullableCapabilitySwitchManufacturingDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchManufacturingDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
