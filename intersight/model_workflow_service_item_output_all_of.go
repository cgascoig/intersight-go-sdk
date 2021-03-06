/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowServiceItemOutputAllOf Definition of the list of properties defined in 'workflow.ServiceItemOutput', excluding properties defined in parent classes.
type WorkflowServiceItemOutputAllOf struct {
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

type _WorkflowServiceItemOutputAllOf WorkflowServiceItemOutputAllOf

// NewWorkflowServiceItemOutputAllOf instantiates a new WorkflowServiceItemOutputAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemOutputAllOf(classId string, objectType string) *WorkflowServiceItemOutputAllOf {
	this := WorkflowServiceItemOutputAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemOutputAllOfWithDefaults instantiates a new WorkflowServiceItemOutputAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemOutputAllOfWithDefaults() *WorkflowServiceItemOutputAllOf {
	this := WorkflowServiceItemOutputAllOf{}
	var classId string = "workflow.ServiceItemOutput"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemOutput"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemOutputAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutputAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemOutputAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemOutputAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutputAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemOutputAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemOutputAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutputAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutputAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemOutputAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemOutputAllOf) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemOutputAllOf) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutputAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowServiceItemOutputAllOf) SetOutput(v interface{}) {
	o.Output = v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowServiceItemOutputAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || o.ServiceItemInstance == nil {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemOutputAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil || o.ServiceItemInstance == nil {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemOutputAllOf) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance != nil {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given WorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemOutputAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance = &v
}

func (o WorkflowServiceItemOutputAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *WorkflowServiceItemOutputAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemOutputAllOf := _WorkflowServiceItemOutputAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemOutputAllOf); err == nil {
		*o = WorkflowServiceItemOutputAllOf(varWorkflowServiceItemOutputAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "ServiceItemInstance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemOutputAllOf struct {
	value *WorkflowServiceItemOutputAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemOutputAllOf) Get() *WorkflowServiceItemOutputAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemOutputAllOf) Set(val *WorkflowServiceItemOutputAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemOutputAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemOutputAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemOutputAllOf(val *WorkflowServiceItemOutputAllOf) *NullableWorkflowServiceItemOutputAllOf {
	return &NullableWorkflowServiceItemOutputAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemOutputAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemOutputAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
