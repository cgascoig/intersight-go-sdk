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

// WorkflowTemplateParserAllOf Definition of the list of properties defined in 'workflow.TemplateParser', excluding properties defined in parent classes.
type WorkflowTemplateParserAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                      `json:"ObjectType"`
	Placeholders []WorkflowPrimitiveDataType `json:"Placeholders,omitempty"`
	// The content of the entire template file. The content can either be a static or dynamic file that can contain placeholders. The placeholders are expected to conform to the golang template syntax i.e. it must be provided inside {{ }}.
	TemplateContent      *string `json:"TemplateContent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTemplateParserAllOf WorkflowTemplateParserAllOf

// NewWorkflowTemplateParserAllOf instantiates a new WorkflowTemplateParserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTemplateParserAllOf(classId string, objectType string) *WorkflowTemplateParserAllOf {
	this := WorkflowTemplateParserAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTemplateParserAllOfWithDefaults instantiates a new WorkflowTemplateParserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTemplateParserAllOfWithDefaults() *WorkflowTemplateParserAllOf {
	this := WorkflowTemplateParserAllOf{}
	var classId string = "workflow.TemplateParser"
	this.ClassId = classId
	var objectType string = "workflow.TemplateParser"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTemplateParserAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateParserAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTemplateParserAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTemplateParserAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateParserAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTemplateParserAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTemplateParserAllOf) GetPlaceholders() []WorkflowPrimitiveDataType {
	if o == nil {
		var ret []WorkflowPrimitiveDataType
		return ret
	}
	return o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTemplateParserAllOf) GetPlaceholdersOk() (*[]WorkflowPrimitiveDataType, bool) {
	if o == nil || o.Placeholders == nil {
		return nil, false
	}
	return &o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *WorkflowTemplateParserAllOf) HasPlaceholders() bool {
	if o != nil && o.Placeholders != nil {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given []WorkflowPrimitiveDataType and assigns it to the Placeholders field.
func (o *WorkflowTemplateParserAllOf) SetPlaceholders(v []WorkflowPrimitiveDataType) {
	o.Placeholders = v
}

// GetTemplateContent returns the TemplateContent field value if set, zero value otherwise.
func (o *WorkflowTemplateParserAllOf) GetTemplateContent() string {
	if o == nil || o.TemplateContent == nil {
		var ret string
		return ret
	}
	return *o.TemplateContent
}

// GetTemplateContentOk returns a tuple with the TemplateContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTemplateParserAllOf) GetTemplateContentOk() (*string, bool) {
	if o == nil || o.TemplateContent == nil {
		return nil, false
	}
	return o.TemplateContent, true
}

// HasTemplateContent returns a boolean if a field has been set.
func (o *WorkflowTemplateParserAllOf) HasTemplateContent() bool {
	if o != nil && o.TemplateContent != nil {
		return true
	}

	return false
}

// SetTemplateContent gets a reference to the given string and assigns it to the TemplateContent field.
func (o *WorkflowTemplateParserAllOf) SetTemplateContent(v string) {
	o.TemplateContent = &v
}

func (o WorkflowTemplateParserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Placeholders != nil {
		toSerialize["Placeholders"] = o.Placeholders
	}
	if o.TemplateContent != nil {
		toSerialize["TemplateContent"] = o.TemplateContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTemplateParserAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTemplateParserAllOf := _WorkflowTemplateParserAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTemplateParserAllOf); err == nil {
		*o = WorkflowTemplateParserAllOf(varWorkflowTemplateParserAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Placeholders")
		delete(additionalProperties, "TemplateContent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTemplateParserAllOf struct {
	value *WorkflowTemplateParserAllOf
	isSet bool
}

func (v NullableWorkflowTemplateParserAllOf) Get() *WorkflowTemplateParserAllOf {
	return v.value
}

func (v *NullableWorkflowTemplateParserAllOf) Set(val *WorkflowTemplateParserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTemplateParserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTemplateParserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTemplateParserAllOf(val *WorkflowTemplateParserAllOf) *NullableWorkflowTemplateParserAllOf {
	return &NullableWorkflowTemplateParserAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTemplateParserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTemplateParserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
