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

// WorkflowServiceItemActionWorkflowDefinitionAllOf Definition of the list of properties defined in 'workflow.ServiceItemActionWorkflowDefinition', excluding properties defined in parent classes.
type WorkflowServiceItemActionWorkflowDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// The description of this workflow instance.
	Description *string `json:"Description,omitempty"`
	// Capture the mapping of ActionDefinition inputDefinition to workflow definition.
	InputParameters interface{} `json:"InputParameters,omitempty"`
	// A user defined label identifier of the workflow used for UI display.
	Label *string `json:"Label,omitempty"`
	// The name of the workflow, this name must be unique across all the workflow definition used within the action definitions.
	Name *string `json:"Name,omitempty"`
	// The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.
	Version *int64 `json:"Version,omitempty"`
	// The qualified name of workflow that should be executed.
	WorkflowDefinitionName *string `json:"WorkflowDefinitionName,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _WorkflowServiceItemActionWorkflowDefinitionAllOf WorkflowServiceItemActionWorkflowDefinitionAllOf

// NewWorkflowServiceItemActionWorkflowDefinitionAllOf instantiates a new WorkflowServiceItemActionWorkflowDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemActionWorkflowDefinitionAllOf(classId string, objectType string) *WorkflowServiceItemActionWorkflowDefinitionAllOf {
	this := WorkflowServiceItemActionWorkflowDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemActionWorkflowDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemActionWorkflowDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemActionWorkflowDefinitionAllOfWithDefaults() *WorkflowServiceItemActionWorkflowDefinitionAllOf {
	this := WorkflowServiceItemActionWorkflowDefinitionAllOf{}
	var classId string = "workflow.ServiceItemActionWorkflowDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemActionWorkflowDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetInputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetInputParametersOk() (*interface{}, bool) {
	if o == nil || o.InputParameters == nil {
		return nil, false
	}
	return &o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasInputParameters() bool {
	if o != nil && o.InputParameters != nil {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given interface{} and assigns it to the InputParameters field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetInputParameters(v interface{}) {
	o.InputParameters = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetWorkflowDefinitionName returns the WorkflowDefinitionName field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetWorkflowDefinitionName() string {
	if o == nil || o.WorkflowDefinitionName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowDefinitionName
}

// GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetWorkflowDefinitionNameOk() (*string, bool) {
	if o == nil || o.WorkflowDefinitionName == nil {
		return nil, false
	}
	return o.WorkflowDefinitionName, true
}

// HasWorkflowDefinitionName returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasWorkflowDefinitionName() bool {
	if o != nil && o.WorkflowDefinitionName != nil {
		return true
	}

	return false
}

// SetWorkflowDefinitionName gets a reference to the given string and assigns it to the WorkflowDefinitionName field.
func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetWorkflowDefinitionName(v string) {
	o.WorkflowDefinitionName = &v
}

func (o WorkflowServiceItemActionWorkflowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputParameters != nil {
		toSerialize["InputParameters"] = o.InputParameters
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WorkflowDefinitionName != nil {
		toSerialize["WorkflowDefinitionName"] = o.WorkflowDefinitionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemActionWorkflowDefinitionAllOf := _WorkflowServiceItemActionWorkflowDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemActionWorkflowDefinitionAllOf); err == nil {
		*o = WorkflowServiceItemActionWorkflowDefinitionAllOf(varWorkflowServiceItemActionWorkflowDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputParameters")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WorkflowDefinitionName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemActionWorkflowDefinitionAllOf struct {
	value *WorkflowServiceItemActionWorkflowDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) Get() *WorkflowServiceItemActionWorkflowDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) Set(val *WorkflowServiceItemActionWorkflowDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemActionWorkflowDefinitionAllOf(val *WorkflowServiceItemActionWorkflowDefinitionAllOf) *NullableWorkflowServiceItemActionWorkflowDefinitionAllOf {
	return &NullableWorkflowServiceItemActionWorkflowDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemActionWorkflowDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
