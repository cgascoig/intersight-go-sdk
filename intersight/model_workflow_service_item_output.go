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

// WorkflowServiceItemOutput Service item output which represents all the artifacts created or related to this service item instance.
type WorkflowServiceItemOutput struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Output name which is used in the output definition of the service item.
	Name *string `json:"Name,omitempty"`
	// Service item output for a service item instance and the format is specified by output definition of the service item definition.
	Output               interface{}                              `json:"Output,omitempty"`
	ServiceItemInstance  *WorkflowServiceItemInstanceRelationship `json:"ServiceItemInstance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowServiceItemOutput WorkflowServiceItemOutput

// NewWorkflowServiceItemOutput instantiates a new WorkflowServiceItemOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemOutput(classId string, objectType string) *WorkflowServiceItemOutput {
	this := WorkflowServiceItemOutput{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemOutputWithDefaults instantiates a new WorkflowServiceItemOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemOutputWithDefaults() *WorkflowServiceItemOutput {
	this := WorkflowServiceItemOutput{}
	var classId string = "workflow.ServiceItemOutput"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemOutput"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemOutput) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutput) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemOutput) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemOutput) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutput) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemOutput) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemOutput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutput) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemOutput) SetName(v string) {
	o.Name = &v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemOutput) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemOutput) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutput) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowServiceItemOutput) SetOutput(v interface{}) {
	o.Output = v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowServiceItemOutput) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || o.ServiceItemInstance == nil {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutput) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil || o.ServiceItemInstance == nil {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutput) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance != nil {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given WorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemOutput) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance = &v
}

func (o WorkflowServiceItemOutput) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.ServiceItemInstance != nil {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemOutput) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowServiceItemOutputWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Output name which is used in the output definition of the service item.
		Name *string `json:"Name,omitempty"`
		// Service item output for a service item instance and the format is specified by output definition of the service item definition.
		Output              interface{}                              `json:"Output,omitempty"`
		ServiceItemInstance *WorkflowServiceItemInstanceRelationship `json:"ServiceItemInstance,omitempty"`
	}

	varWorkflowServiceItemOutputWithoutEmbeddedStruct := WorkflowServiceItemOutputWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemOutputWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemOutput := _WorkflowServiceItemOutput{}
		varWorkflowServiceItemOutput.ClassId = varWorkflowServiceItemOutputWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemOutput.ObjectType = varWorkflowServiceItemOutputWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemOutput.Name = varWorkflowServiceItemOutputWithoutEmbeddedStruct.Name
		varWorkflowServiceItemOutput.Output = varWorkflowServiceItemOutputWithoutEmbeddedStruct.Output
		varWorkflowServiceItemOutput.ServiceItemInstance = varWorkflowServiceItemOutputWithoutEmbeddedStruct.ServiceItemInstance
		*o = WorkflowServiceItemOutput(varWorkflowServiceItemOutput)
	} else {
		return err
	}

	varWorkflowServiceItemOutput := _WorkflowServiceItemOutput{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemOutput)
	if err == nil {
		o.MoBaseMo = varWorkflowServiceItemOutput.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "ServiceItemInstance")

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

type NullableWorkflowServiceItemOutput struct {
	value *WorkflowServiceItemOutput
	isSet bool
}

func (v NullableWorkflowServiceItemOutput) Get() *WorkflowServiceItemOutput {
	return v.value
}

func (v *NullableWorkflowServiceItemOutput) Set(val *WorkflowServiceItemOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemOutput(val *WorkflowServiceItemOutput) *NullableWorkflowServiceItemOutput {
	return &NullableWorkflowServiceItemOutput{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
