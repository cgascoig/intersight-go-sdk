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
)

// WorkflowMoReferenceArrayItemAllOf Definition of the list of properties defined in 'workflow.MoReferenceArrayItem', excluding properties defined in parent classes.
type WorkflowMoReferenceArrayItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                        `json:"ObjectType"`
	Properties           []WorkflowMoReferenceProperty `json:"Properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMoReferenceArrayItemAllOf WorkflowMoReferenceArrayItemAllOf

// NewWorkflowMoReferenceArrayItemAllOf instantiates a new WorkflowMoReferenceArrayItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoReferenceArrayItemAllOf(classId string, objectType string) *WorkflowMoReferenceArrayItemAllOf {
	this := WorkflowMoReferenceArrayItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMoReferenceArrayItemAllOfWithDefaults instantiates a new WorkflowMoReferenceArrayItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoReferenceArrayItemAllOfWithDefaults() *WorkflowMoReferenceArrayItemAllOf {
	this := WorkflowMoReferenceArrayItemAllOf{}
	var classId string = "workflow.MoReferenceArrayItem"
	this.ClassId = classId
	var objectType string = "workflow.MoReferenceArrayItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoReferenceArrayItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceArrayItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoReferenceArrayItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoReferenceArrayItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferenceArrayItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoReferenceArrayItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferenceArrayItemAllOf) GetProperties() []WorkflowMoReferenceProperty {
	if o == nil {
		var ret []WorkflowMoReferenceProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferenceArrayItemAllOf) GetPropertiesOk() (*[]WorkflowMoReferenceProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowMoReferenceArrayItemAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []WorkflowMoReferenceProperty and assigns it to the Properties field.
func (o *WorkflowMoReferenceArrayItemAllOf) SetProperties(v []WorkflowMoReferenceProperty) {
	o.Properties = v
}

func (o WorkflowMoReferenceArrayItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMoReferenceArrayItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMoReferenceArrayItemAllOf := _WorkflowMoReferenceArrayItemAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowMoReferenceArrayItemAllOf); err == nil {
		*o = WorkflowMoReferenceArrayItemAllOf(varWorkflowMoReferenceArrayItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowMoReferenceArrayItemAllOf struct {
	value *WorkflowMoReferenceArrayItemAllOf
	isSet bool
}

func (v NullableWorkflowMoReferenceArrayItemAllOf) Get() *WorkflowMoReferenceArrayItemAllOf {
	return v.value
}

func (v *NullableWorkflowMoReferenceArrayItemAllOf) Set(val *WorkflowMoReferenceArrayItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoReferenceArrayItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoReferenceArrayItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoReferenceArrayItemAllOf(val *WorkflowMoReferenceArrayItemAllOf) *NullableWorkflowMoReferenceArrayItemAllOf {
	return &NullableWorkflowMoReferenceArrayItemAllOf{value: val, isSet: true}
}

func (v NullableWorkflowMoReferenceArrayItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoReferenceArrayItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
