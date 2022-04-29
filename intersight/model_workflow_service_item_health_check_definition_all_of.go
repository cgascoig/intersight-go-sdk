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

// WorkflowServiceItemHealthCheckDefinitionAllOf Definition of the list of properties defined in 'workflow.ServiceItemHealthCheckDefinition', excluding properties defined in parent classes.
type WorkflowServiceItemHealthCheckDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Category that the health check belongs to.
	Category *string `json:"Category,omitempty"`
	// Static information detailing the common cause for the health check failure. It also gives possible remediation actions that can be taken to remedy the health check failure.
	CommonCauseAndResolution *string `json:"CommonCauseAndResolution,omitempty"`
	// Description of the health check definition.
	Description *string `json:"Description,omitempty"`
	// Execution mode of the health check on service item. * `OnDemand` - Execute the health check on-demand. * `Periodic` - Execute the health check on a periodic basis.
	ExecutionMode       *string                                             `json:"ExecutionMode,omitempty"`
	HealthCheckWorkflow NullableWorkflowServiceItemActionWorkflowDefinition `json:"HealthCheckWorkflow,omitempty"`
	// Label for the health check definition that is displayed on UI.
	Label *string `json:"Label,omitempty"`
	// Name of the health check definition.
	Name                  *string                                    `json:"Name,omitempty"`
	ServiceItemDefinition *WorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemHealthCheckDefinitionAllOf WorkflowServiceItemHealthCheckDefinitionAllOf

// NewWorkflowServiceItemHealthCheckDefinitionAllOf instantiates a new WorkflowServiceItemHealthCheckDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemHealthCheckDefinitionAllOf(classId string, objectType string) *WorkflowServiceItemHealthCheckDefinitionAllOf {
	this := WorkflowServiceItemHealthCheckDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// NewWorkflowServiceItemHealthCheckDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemHealthCheckDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemHealthCheckDefinitionAllOfWithDefaults() *WorkflowServiceItemHealthCheckDefinitionAllOf {
	this := WorkflowServiceItemHealthCheckDefinitionAllOf{}
	var classId string = "workflow.ServiceItemHealthCheckDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemHealthCheckDefinition"
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetCommonCauseAndResolution returns the CommonCauseAndResolution field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCommonCauseAndResolution() string {
	if o == nil || o.CommonCauseAndResolution == nil {
		var ret string
		return ret
	}
	return *o.CommonCauseAndResolution
}

// GetCommonCauseAndResolutionOk returns a tuple with the CommonCauseAndResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetCommonCauseAndResolutionOk() (*string, bool) {
	if o == nil || o.CommonCauseAndResolution == nil {
		return nil, false
	}
	return o.CommonCauseAndResolution, true
}

// HasCommonCauseAndResolution returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasCommonCauseAndResolution() bool {
	if o != nil && o.CommonCauseAndResolution != nil {
		return true
	}

	return false
}

// SetCommonCauseAndResolution gets a reference to the given string and assigns it to the CommonCauseAndResolution field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetCommonCauseAndResolution(v string) {
	o.CommonCauseAndResolution = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetExecutionMode() string {
	if o == nil || o.ExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetExecutionModeOk() (*string, bool) {
	if o == nil || o.ExecutionMode == nil {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasExecutionMode() bool {
	if o != nil && o.ExecutionMode != nil {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetHealthCheckWorkflow returns the HealthCheckWorkflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetHealthCheckWorkflow() WorkflowServiceItemActionWorkflowDefinition {
	if o == nil || o.HealthCheckWorkflow.Get() == nil {
		var ret WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return *o.HealthCheckWorkflow.Get()
}

// GetHealthCheckWorkflowOk returns a tuple with the HealthCheckWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetHealthCheckWorkflowOk() (*WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthCheckWorkflow.Get(), o.HealthCheckWorkflow.IsSet()
}

// HasHealthCheckWorkflow returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasHealthCheckWorkflow() bool {
	if o != nil && o.HealthCheckWorkflow.IsSet() {
		return true
	}

	return false
}

// SetHealthCheckWorkflow gets a reference to the given NullableWorkflowServiceItemActionWorkflowDefinition and assigns it to the HealthCheckWorkflow field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetHealthCheckWorkflow(v WorkflowServiceItemActionWorkflowDefinition) {
	o.HealthCheckWorkflow.Set(&v)
}

// SetHealthCheckWorkflowNil sets the value for HealthCheckWorkflow to be an explicit nil
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetHealthCheckWorkflowNil() {
	o.HealthCheckWorkflow.Set(nil)
}

// UnsetHealthCheckWorkflow ensures that no value is present for HealthCheckWorkflow, not even an explicit nil
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) UnsetHealthCheckWorkflow() {
	o.HealthCheckWorkflow.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || o.ServiceItemDefinition == nil {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemDefinition == nil {
		return nil, false
	}
	return o.ServiceItemDefinition, true
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given WorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition = &v
}

func (o WorkflowServiceItemHealthCheckDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.CommonCauseAndResolution != nil {
		toSerialize["CommonCauseAndResolution"] = o.CommonCauseAndResolution
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ExecutionMode != nil {
		toSerialize["ExecutionMode"] = o.ExecutionMode
	}
	if o.HealthCheckWorkflow.IsSet() {
		toSerialize["HealthCheckWorkflow"] = o.HealthCheckWorkflow.Get()
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServiceItemDefinition != nil {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemHealthCheckDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemHealthCheckDefinitionAllOf := _WorkflowServiceItemHealthCheckDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemHealthCheckDefinitionAllOf); err == nil {
		*o = WorkflowServiceItemHealthCheckDefinitionAllOf(varWorkflowServiceItemHealthCheckDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "CommonCauseAndResolution")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ExecutionMode")
		delete(additionalProperties, "HealthCheckWorkflow")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServiceItemDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemHealthCheckDefinitionAllOf struct {
	value *WorkflowServiceItemHealthCheckDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemHealthCheckDefinitionAllOf) Get() *WorkflowServiceItemHealthCheckDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemHealthCheckDefinitionAllOf) Set(val *WorkflowServiceItemHealthCheckDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemHealthCheckDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemHealthCheckDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemHealthCheckDefinitionAllOf(val *WorkflowServiceItemHealthCheckDefinitionAllOf) *NullableWorkflowServiceItemHealthCheckDefinitionAllOf {
	return &NullableWorkflowServiceItemHealthCheckDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemHealthCheckDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemHealthCheckDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
