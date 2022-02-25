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

// WorkflowPrimitiveDataProperty Capture all the properties for primitive data type.
type WorkflowPrimitiveDataProperty struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                        `json:"ObjectType"`
	Constraints       NullableWorkflowConstraints   `json:"Constraints,omitempty"`
	InventorySelector []WorkflowMoReferenceProperty `json:"InventorySelector,omitempty"`
	// Intersight supports secure properties as task input/output. The values of these properties are encrypted and stored in Intersight. This flag marks the property to be secure when it is set to true.
	Secure *bool `json:"Secure,omitempty"`
	// Specify the enum type for primitive data type. * `string` - Enum to specify a string data type. * `integer` - Enum to specify an integer32 data type. * `float` - Enum to specify a float64 data type. * `boolean` - Enum to specify a boolean data type. * `json` - Enum to specify a json data type. * `enum` - Enum to specify a enum data type which is a list of pre-defined strings.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPrimitiveDataProperty WorkflowPrimitiveDataProperty

// NewWorkflowPrimitiveDataProperty instantiates a new WorkflowPrimitiveDataProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPrimitiveDataProperty(classId string, objectType string) *WorkflowPrimitiveDataProperty {
	this := WorkflowPrimitiveDataProperty{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// NewWorkflowPrimitiveDataPropertyWithDefaults instantiates a new WorkflowPrimitiveDataProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPrimitiveDataPropertyWithDefaults() *WorkflowPrimitiveDataProperty {
	this := WorkflowPrimitiveDataProperty{}
	var classId string = "workflow.PrimitiveDataProperty"
	this.ClassId = classId
	var objectType string = "workflow.PrimitiveDataProperty"
	this.ObjectType = objectType
	var type_ string = "string"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPrimitiveDataProperty) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataProperty) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPrimitiveDataProperty) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPrimitiveDataProperty) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataProperty) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPrimitiveDataProperty) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPrimitiveDataProperty) GetConstraints() WorkflowConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret WorkflowConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPrimitiveDataProperty) GetConstraintsOk() (*WorkflowConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataProperty) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowConstraints and assigns it to the Constraints field.
func (o *WorkflowPrimitiveDataProperty) SetConstraints(v WorkflowConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowPrimitiveDataProperty) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowPrimitiveDataProperty) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetInventorySelector returns the InventorySelector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPrimitiveDataProperty) GetInventorySelector() []WorkflowMoReferenceProperty {
	if o == nil {
		var ret []WorkflowMoReferenceProperty
		return ret
	}
	return o.InventorySelector
}

// GetInventorySelectorOk returns a tuple with the InventorySelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPrimitiveDataProperty) GetInventorySelectorOk() (*[]WorkflowMoReferenceProperty, bool) {
	if o == nil || o.InventorySelector == nil {
		return nil, false
	}
	return &o.InventorySelector, true
}

// HasInventorySelector returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataProperty) HasInventorySelector() bool {
	if o != nil && o.InventorySelector != nil {
		return true
	}

	return false
}

// SetInventorySelector gets a reference to the given []WorkflowMoReferenceProperty and assigns it to the InventorySelector field.
func (o *WorkflowPrimitiveDataProperty) SetInventorySelector(v []WorkflowMoReferenceProperty) {
	o.InventorySelector = v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataProperty) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataProperty) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataProperty) HasSecure() bool {
	if o != nil && o.Secure != nil {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *WorkflowPrimitiveDataProperty) SetSecure(v bool) {
	o.Secure = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowPrimitiveDataProperty) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPrimitiveDataProperty) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowPrimitiveDataProperty) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowPrimitiveDataProperty) SetType(v string) {
	o.Type = &v
}

func (o WorkflowPrimitiveDataProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if o.InventorySelector != nil {
		toSerialize["InventorySelector"] = o.InventorySelector
	}
	if o.Secure != nil {
		toSerialize["Secure"] = o.Secure
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPrimitiveDataProperty) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPrimitiveDataPropertyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                        `json:"ObjectType"`
		Constraints       NullableWorkflowConstraints   `json:"Constraints,omitempty"`
		InventorySelector []WorkflowMoReferenceProperty `json:"InventorySelector,omitempty"`
		// Intersight supports secure properties as task input/output. The values of these properties are encrypted and stored in Intersight. This flag marks the property to be secure when it is set to true.
		Secure *bool `json:"Secure,omitempty"`
		// Specify the enum type for primitive data type. * `string` - Enum to specify a string data type. * `integer` - Enum to specify an integer32 data type. * `float` - Enum to specify a float64 data type. * `boolean` - Enum to specify a boolean data type. * `json` - Enum to specify a json data type. * `enum` - Enum to specify a enum data type which is a list of pre-defined strings.
		Type *string `json:"Type,omitempty"`
	}

	varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct := WorkflowPrimitiveDataPropertyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPrimitiveDataProperty := _WorkflowPrimitiveDataProperty{}
		varWorkflowPrimitiveDataProperty.ClassId = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.ClassId
		varWorkflowPrimitiveDataProperty.ObjectType = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.ObjectType
		varWorkflowPrimitiveDataProperty.Constraints = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.Constraints
		varWorkflowPrimitiveDataProperty.InventorySelector = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.InventorySelector
		varWorkflowPrimitiveDataProperty.Secure = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.Secure
		varWorkflowPrimitiveDataProperty.Type = varWorkflowPrimitiveDataPropertyWithoutEmbeddedStruct.Type
		*o = WorkflowPrimitiveDataProperty(varWorkflowPrimitiveDataProperty)
	} else {
		return err
	}

	varWorkflowPrimitiveDataProperty := _WorkflowPrimitiveDataProperty{}

	err = json.Unmarshal(bytes, &varWorkflowPrimitiveDataProperty)
	if err == nil {
		o.MoBaseComplexType = varWorkflowPrimitiveDataProperty.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "InventorySelector")
		delete(additionalProperties, "Secure")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowPrimitiveDataProperty struct {
	value *WorkflowPrimitiveDataProperty
	isSet bool
}

func (v NullableWorkflowPrimitiveDataProperty) Get() *WorkflowPrimitiveDataProperty {
	return v.value
}

func (v *NullableWorkflowPrimitiveDataProperty) Set(val *WorkflowPrimitiveDataProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPrimitiveDataProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPrimitiveDataProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPrimitiveDataProperty(val *WorkflowPrimitiveDataProperty) *NullableWorkflowPrimitiveDataProperty {
	return &NullableWorkflowPrimitiveDataProperty{value: val, isSet: true}
}

func (v NullableWorkflowPrimitiveDataProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPrimitiveDataProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
