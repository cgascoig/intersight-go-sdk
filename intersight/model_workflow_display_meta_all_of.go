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

// WorkflowDisplayMetaAllOf Definition of the list of properties defined in 'workflow.DisplayMeta', excluding properties defined in parent classes.
type WorkflowDisplayMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Inventory selector specified for primitive data property should be used in Intersight User Interface.
	InventorySelector *bool `json:"InventorySelector,omitempty"`
	// Specify the widget type for data display. * `None` - Display none of the widget types. * `Radio` - Display the widget as a radio button. * `Dropdown` - Display the widget as a dropdown. * `GridSelector` - Display the widget as a selector. * `DrawerSelector` - Display the widget as a selector.
	WidgetType           *string `json:"WidgetType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowDisplayMetaAllOf WorkflowDisplayMetaAllOf

// NewWorkflowDisplayMetaAllOf instantiates a new WorkflowDisplayMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDisplayMetaAllOf(classId string, objectType string) *WorkflowDisplayMetaAllOf {
	this := WorkflowDisplayMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var inventorySelector bool = true
	this.InventorySelector = &inventorySelector
	var widgetType string = "None"
	this.WidgetType = &widgetType
	return &this
}

// NewWorkflowDisplayMetaAllOfWithDefaults instantiates a new WorkflowDisplayMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDisplayMetaAllOfWithDefaults() *WorkflowDisplayMetaAllOf {
	this := WorkflowDisplayMetaAllOf{}
	var classId string = "workflow.DisplayMeta"
	this.ClassId = classId
	var objectType string = "workflow.DisplayMeta"
	this.ObjectType = objectType
	var inventorySelector bool = true
	this.InventorySelector = &inventorySelector
	var widgetType string = "None"
	this.WidgetType = &widgetType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowDisplayMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDisplayMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowDisplayMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowDisplayMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowDisplayMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowDisplayMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInventorySelector returns the InventorySelector field value if set, zero value otherwise.
func (o *WorkflowDisplayMetaAllOf) GetInventorySelector() bool {
	if o == nil || o.InventorySelector == nil {
		var ret bool
		return ret
	}
	return *o.InventorySelector
}

// GetInventorySelectorOk returns a tuple with the InventorySelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDisplayMetaAllOf) GetInventorySelectorOk() (*bool, bool) {
	if o == nil || o.InventorySelector == nil {
		return nil, false
	}
	return o.InventorySelector, true
}

// HasInventorySelector returns a boolean if a field has been set.
func (o *WorkflowDisplayMetaAllOf) HasInventorySelector() bool {
	if o != nil && o.InventorySelector != nil {
		return true
	}

	return false
}

// SetInventorySelector gets a reference to the given bool and assigns it to the InventorySelector field.
func (o *WorkflowDisplayMetaAllOf) SetInventorySelector(v bool) {
	o.InventorySelector = &v
}

// GetWidgetType returns the WidgetType field value if set, zero value otherwise.
func (o *WorkflowDisplayMetaAllOf) GetWidgetType() string {
	if o == nil || o.WidgetType == nil {
		var ret string
		return ret
	}
	return *o.WidgetType
}

// GetWidgetTypeOk returns a tuple with the WidgetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDisplayMetaAllOf) GetWidgetTypeOk() (*string, bool) {
	if o == nil || o.WidgetType == nil {
		return nil, false
	}
	return o.WidgetType, true
}

// HasWidgetType returns a boolean if a field has been set.
func (o *WorkflowDisplayMetaAllOf) HasWidgetType() bool {
	if o != nil && o.WidgetType != nil {
		return true
	}

	return false
}

// SetWidgetType gets a reference to the given string and assigns it to the WidgetType field.
func (o *WorkflowDisplayMetaAllOf) SetWidgetType(v string) {
	o.WidgetType = &v
}

func (o WorkflowDisplayMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InventorySelector != nil {
		toSerialize["InventorySelector"] = o.InventorySelector
	}
	if o.WidgetType != nil {
		toSerialize["WidgetType"] = o.WidgetType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowDisplayMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowDisplayMetaAllOf := _WorkflowDisplayMetaAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowDisplayMetaAllOf); err == nil {
		*o = WorkflowDisplayMetaAllOf(varWorkflowDisplayMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InventorySelector")
		delete(additionalProperties, "WidgetType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowDisplayMetaAllOf struct {
	value *WorkflowDisplayMetaAllOf
	isSet bool
}

func (v NullableWorkflowDisplayMetaAllOf) Get() *WorkflowDisplayMetaAllOf {
	return v.value
}

func (v *NullableWorkflowDisplayMetaAllOf) Set(val *WorkflowDisplayMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDisplayMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDisplayMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDisplayMetaAllOf(val *WorkflowDisplayMetaAllOf) *NullableWorkflowDisplayMetaAllOf {
	return &NullableWorkflowDisplayMetaAllOf{value: val, isSet: true}
}

func (v NullableWorkflowDisplayMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDisplayMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
