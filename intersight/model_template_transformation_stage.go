/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TemplateTransformationStage Contains the template to be executed and a unique name by which this transformation stage output can be accessed in further stages.
type TemplateTransformationStage struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The function to be executed.
	Function *string `json:"Function,omitempty"`
	// A collection of arguments for the function being executed.
	FunctionArguments interface{} `json:"FunctionArguments,omitempty"`
	// The unique name by which the output of this transformation stage can be accessed in further stages. Only alphanumeric characters are allowed.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateTransformationStage TemplateTransformationStage

// NewTemplateTransformationStage instantiates a new TemplateTransformationStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateTransformationStage(classId string, objectType string) *TemplateTransformationStage {
	this := TemplateTransformationStage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTemplateTransformationStageWithDefaults instantiates a new TemplateTransformationStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateTransformationStageWithDefaults() *TemplateTransformationStage {
	this := TemplateTransformationStage{}
	var classId string = "template.TransformationStage"
	this.ClassId = classId
	var objectType string = "template.TransformationStage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TemplateTransformationStage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TemplateTransformationStage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TemplateTransformationStage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TemplateTransformationStage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFunction returns the Function field value if set, zero value otherwise.
func (o *TemplateTransformationStage) GetFunction() string {
	if o == nil || o.Function == nil {
		var ret string
		return ret
	}
	return *o.Function
}

// GetFunctionOk returns a tuple with the Function field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStage) GetFunctionOk() (*string, bool) {
	if o == nil || o.Function == nil {
		return nil, false
	}
	return o.Function, true
}

// HasFunction returns a boolean if a field has been set.
func (o *TemplateTransformationStage) HasFunction() bool {
	if o != nil && o.Function != nil {
		return true
	}

	return false
}

// SetFunction gets a reference to the given string and assigns it to the Function field.
func (o *TemplateTransformationStage) SetFunction(v string) {
	o.Function = &v
}

// GetFunctionArguments returns the FunctionArguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateTransformationStage) GetFunctionArguments() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FunctionArguments
}

// GetFunctionArgumentsOk returns a tuple with the FunctionArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateTransformationStage) GetFunctionArgumentsOk() (*interface{}, bool) {
	if o == nil || o.FunctionArguments == nil {
		return nil, false
	}
	return &o.FunctionArguments, true
}

// HasFunctionArguments returns a boolean if a field has been set.
func (o *TemplateTransformationStage) HasFunctionArguments() bool {
	if o != nil && o.FunctionArguments != nil {
		return true
	}

	return false
}

// SetFunctionArguments gets a reference to the given interface{} and assigns it to the FunctionArguments field.
func (o *TemplateTransformationStage) SetFunctionArguments(v interface{}) {
	o.FunctionArguments = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateTransformationStage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateTransformationStage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateTransformationStage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateTransformationStage) SetName(v string) {
	o.Name = &v
}

func (o TemplateTransformationStage) MarshalJSON() ([]byte, error) {
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
	if o.Function != nil {
		toSerialize["Function"] = o.Function
	}
	if o.FunctionArguments != nil {
		toSerialize["FunctionArguments"] = o.FunctionArguments
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateTransformationStage) UnmarshalJSON(bytes []byte) (err error) {
	type TemplateTransformationStageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The function to be executed.
		Function *string `json:"Function,omitempty"`
		// A collection of arguments for the function being executed.
		FunctionArguments interface{} `json:"FunctionArguments,omitempty"`
		// The unique name by which the output of this transformation stage can be accessed in further stages. Only alphanumeric characters are allowed.
		Name *string `json:"Name,omitempty"`
	}

	varTemplateTransformationStageWithoutEmbeddedStruct := TemplateTransformationStageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTemplateTransformationStageWithoutEmbeddedStruct)
	if err == nil {
		varTemplateTransformationStage := _TemplateTransformationStage{}
		varTemplateTransformationStage.ClassId = varTemplateTransformationStageWithoutEmbeddedStruct.ClassId
		varTemplateTransformationStage.ObjectType = varTemplateTransformationStageWithoutEmbeddedStruct.ObjectType
		varTemplateTransformationStage.Function = varTemplateTransformationStageWithoutEmbeddedStruct.Function
		varTemplateTransformationStage.FunctionArguments = varTemplateTransformationStageWithoutEmbeddedStruct.FunctionArguments
		varTemplateTransformationStage.Name = varTemplateTransformationStageWithoutEmbeddedStruct.Name
		*o = TemplateTransformationStage(varTemplateTransformationStage)
	} else {
		return err
	}

	varTemplateTransformationStage := _TemplateTransformationStage{}

	err = json.Unmarshal(bytes, &varTemplateTransformationStage)
	if err == nil {
		o.MoBaseComplexType = varTemplateTransformationStage.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Function")
		delete(additionalProperties, "FunctionArguments")
		delete(additionalProperties, "Name")

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

type NullableTemplateTransformationStage struct {
	value *TemplateTransformationStage
	isSet bool
}

func (v NullableTemplateTransformationStage) Get() *TemplateTransformationStage {
	return v.value
}

func (v *NullableTemplateTransformationStage) Set(val *TemplateTransformationStage) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateTransformationStage) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateTransformationStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateTransformationStage(val *TemplateTransformationStage) *NullableTemplateTransformationStage {
	return &NullableTemplateTransformationStage{value: val, isSet: true}
}

func (v NullableTemplateTransformationStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateTransformationStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
