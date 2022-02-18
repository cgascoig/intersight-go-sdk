/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowTargetPropertyAllOf Definition of the list of properties defined in 'workflow.TargetProperty', excluding properties defined in parent classes.
type WorkflowTargetPropertyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A singleton value which will contain the path to connector object from the selected object.
	ConnectorAttribute   *string  `json:"ConnectorAttribute,omitempty"`
	ConstraintAttributes []string `json:"ConstraintAttributes,omitempty"`
	DisplayAttributes    []string `json:"DisplayAttributes,omitempty"`
	// Field to hold an Intersight API along with an optional filter to narrow down the search options for target device.
	Selector             *string                          `json:"Selector,omitempty"`
	SelectorProperty     NullableWorkflowSelectorProperty `json:"SelectorProperty,omitempty"`
	SupportedObjects     []string                         `json:"SupportedObjects,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetPropertyAllOf WorkflowTargetPropertyAllOf

// NewWorkflowTargetPropertyAllOf instantiates a new WorkflowTargetPropertyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetPropertyAllOf(classId string, objectType string) *WorkflowTargetPropertyAllOf {
	this := WorkflowTargetPropertyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowTargetPropertyAllOfWithDefaults instantiates a new WorkflowTargetPropertyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetPropertyAllOfWithDefaults() *WorkflowTargetPropertyAllOf {
	this := WorkflowTargetPropertyAllOf{}
	var classId string = "workflow.TargetProperty"
	this.ClassId = classId
	var objectType string = "workflow.TargetProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTargetPropertyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetPropertyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTargetPropertyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTargetPropertyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetPropertyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTargetPropertyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectorAttribute returns the ConnectorAttribute field value if set, zero value otherwise.
func (o *WorkflowTargetPropertyAllOf) GetConnectorAttribute() string {
	if o == nil || o.ConnectorAttribute == nil {
		var ret string
		return ret
	}
	return *o.ConnectorAttribute
}

// GetConnectorAttributeOk returns a tuple with the ConnectorAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetPropertyAllOf) GetConnectorAttributeOk() (*string, bool) {
	if o == nil || o.ConnectorAttribute == nil {
		return nil, false
	}
	return o.ConnectorAttribute, true
}

// HasConnectorAttribute returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasConnectorAttribute() bool {
	if o != nil && o.ConnectorAttribute != nil {
		return true
	}

	return false
}

// SetConnectorAttribute gets a reference to the given string and assigns it to the ConnectorAttribute field.
func (o *WorkflowTargetPropertyAllOf) SetConnectorAttribute(v string) {
	o.ConnectorAttribute = &v
}

// GetConstraintAttributes returns the ConstraintAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetPropertyAllOf) GetConstraintAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConstraintAttributes
}

// GetConstraintAttributesOk returns a tuple with the ConstraintAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetPropertyAllOf) GetConstraintAttributesOk() (*[]string, bool) {
	if o == nil || o.ConstraintAttributes == nil {
		return nil, false
	}
	return &o.ConstraintAttributes, true
}

// HasConstraintAttributes returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasConstraintAttributes() bool {
	if o != nil && o.ConstraintAttributes != nil {
		return true
	}

	return false
}

// SetConstraintAttributes gets a reference to the given []string and assigns it to the ConstraintAttributes field.
func (o *WorkflowTargetPropertyAllOf) SetConstraintAttributes(v []string) {
	o.ConstraintAttributes = v
}

// GetDisplayAttributes returns the DisplayAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetPropertyAllOf) GetDisplayAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisplayAttributes
}

// GetDisplayAttributesOk returns a tuple with the DisplayAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetPropertyAllOf) GetDisplayAttributesOk() (*[]string, bool) {
	if o == nil || o.DisplayAttributes == nil {
		return nil, false
	}
	return &o.DisplayAttributes, true
}

// HasDisplayAttributes returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasDisplayAttributes() bool {
	if o != nil && o.DisplayAttributes != nil {
		return true
	}

	return false
}

// SetDisplayAttributes gets a reference to the given []string and assigns it to the DisplayAttributes field.
func (o *WorkflowTargetPropertyAllOf) SetDisplayAttributes(v []string) {
	o.DisplayAttributes = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *WorkflowTargetPropertyAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetPropertyAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *WorkflowTargetPropertyAllOf) SetSelector(v string) {
	o.Selector = &v
}

// GetSelectorProperty returns the SelectorProperty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetPropertyAllOf) GetSelectorProperty() WorkflowSelectorProperty {
	if o == nil || o.SelectorProperty.Get() == nil {
		var ret WorkflowSelectorProperty
		return ret
	}
	return *o.SelectorProperty.Get()
}

// GetSelectorPropertyOk returns a tuple with the SelectorProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetPropertyAllOf) GetSelectorPropertyOk() (*WorkflowSelectorProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectorProperty.Get(), o.SelectorProperty.IsSet()
}

// HasSelectorProperty returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasSelectorProperty() bool {
	if o != nil && o.SelectorProperty.IsSet() {
		return true
	}

	return false
}

// SetSelectorProperty gets a reference to the given NullableWorkflowSelectorProperty and assigns it to the SelectorProperty field.
func (o *WorkflowTargetPropertyAllOf) SetSelectorProperty(v WorkflowSelectorProperty) {
	o.SelectorProperty.Set(&v)
}

// SetSelectorPropertyNil sets the value for SelectorProperty to be an explicit nil
func (o *WorkflowTargetPropertyAllOf) SetSelectorPropertyNil() {
	o.SelectorProperty.Set(nil)
}

// UnsetSelectorProperty ensures that no value is present for SelectorProperty, not even an explicit nil
func (o *WorkflowTargetPropertyAllOf) UnsetSelectorProperty() {
	o.SelectorProperty.Unset()
}

// GetSupportedObjects returns the SupportedObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowTargetPropertyAllOf) GetSupportedObjects() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedObjects
}

// GetSupportedObjectsOk returns a tuple with the SupportedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowTargetPropertyAllOf) GetSupportedObjectsOk() (*[]string, bool) {
	if o == nil || o.SupportedObjects == nil {
		return nil, false
	}
	return &o.SupportedObjects, true
}

// HasSupportedObjects returns a boolean if a field has been set.
func (o *WorkflowTargetPropertyAllOf) HasSupportedObjects() bool {
	if o != nil && o.SupportedObjects != nil {
		return true
	}

	return false
}

// SetSupportedObjects gets a reference to the given []string and assigns it to the SupportedObjects field.
func (o *WorkflowTargetPropertyAllOf) SetSupportedObjects(v []string) {
	o.SupportedObjects = v
}

func (o WorkflowTargetPropertyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectorAttribute != nil {
		toSerialize["ConnectorAttribute"] = o.ConnectorAttribute
	}
	if o.ConstraintAttributes != nil {
		toSerialize["ConstraintAttributes"] = o.ConstraintAttributes
	}
	if o.DisplayAttributes != nil {
		toSerialize["DisplayAttributes"] = o.DisplayAttributes
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}
	if o.SelectorProperty.IsSet() {
		toSerialize["SelectorProperty"] = o.SelectorProperty.Get()
	}
	if o.SupportedObjects != nil {
		toSerialize["SupportedObjects"] = o.SupportedObjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetPropertyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetPropertyAllOf := _WorkflowTargetPropertyAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetPropertyAllOf); err == nil {
		*o = WorkflowTargetPropertyAllOf(varWorkflowTargetPropertyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectorAttribute")
		delete(additionalProperties, "ConstraintAttributes")
		delete(additionalProperties, "DisplayAttributes")
		delete(additionalProperties, "Selector")
		delete(additionalProperties, "SelectorProperty")
		delete(additionalProperties, "SupportedObjects")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTargetPropertyAllOf struct {
	value *WorkflowTargetPropertyAllOf
	isSet bool
}

func (v NullableWorkflowTargetPropertyAllOf) Get() *WorkflowTargetPropertyAllOf {
	return v.value
}

func (v *NullableWorkflowTargetPropertyAllOf) Set(val *WorkflowTargetPropertyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetPropertyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetPropertyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetPropertyAllOf(val *WorkflowTargetPropertyAllOf) *NullableWorkflowTargetPropertyAllOf {
	return &NullableWorkflowTargetPropertyAllOf{value: val, isSet: true}
}

func (v NullableWorkflowTargetPropertyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetPropertyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
