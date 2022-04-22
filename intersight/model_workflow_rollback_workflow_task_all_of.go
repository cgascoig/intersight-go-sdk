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
)

// WorkflowRollbackWorkflowTaskAllOf Definition of the list of properties defined in 'workflow.RollbackWorkflowTask', excluding properties defined in parent classes.
type WorkflowRollbackWorkflowTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the rollback task.
	Description *string `json:"Description,omitempty"`
	// Name of TaskInfo that needs to be rolled back.
	Name *string `json:"Name,omitempty"`
	// Reference name of TaskInfo that need to be rolled back.
	RefName *string `json:"RefName,omitempty"`
	// Status of the rollback operation for the task.
	RollbackCompleted *bool `json:"RollbackCompleted,omitempty"`
	// Name of TaskInfo that performs the rollback operation.
	RollbackTaskName *string `json:"RollbackTaskName,omitempty"`
	// Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed. * `NotStarted` - Status of rollback task when it is not started rollback. * `NotSupported` - Status of task when it is not supporting rollback. * `Completed` - Status of rollback task once execution is successful. * `Failed` - Status of rollback task when it is failed. * `Disabled` - Status of rollback task when rollback is disabled.
	Status *string `json:"Status,omitempty"`
	// Moid of TaskInfo that supports rollback operation.
	TaskInfoMoid *string `json:"TaskInfoMoid,omitempty"`
	// Path of rollback task if it is inside sub-workflow.
	TaskPath             *string `json:"TaskPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowRollbackWorkflowTaskAllOf WorkflowRollbackWorkflowTaskAllOf

// NewWorkflowRollbackWorkflowTaskAllOf instantiates a new WorkflowRollbackWorkflowTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRollbackWorkflowTaskAllOf(classId string, objectType string) *WorkflowRollbackWorkflowTaskAllOf {
	this := WorkflowRollbackWorkflowTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowRollbackWorkflowTaskAllOfWithDefaults instantiates a new WorkflowRollbackWorkflowTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRollbackWorkflowTaskAllOfWithDefaults() *WorkflowRollbackWorkflowTaskAllOf {
	this := WorkflowRollbackWorkflowTaskAllOf{}
	var classId string = "workflow.RollbackWorkflowTask"
	this.ClassId = classId
	var objectType string = "workflow.RollbackWorkflowTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowRollbackWorkflowTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowRollbackWorkflowTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowRollbackWorkflowTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowRollbackWorkflowTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetName(v string) {
	o.Name = &v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRefName() string {
	if o == nil || o.RefName == nil {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRefNameOk() (*string, bool) {
	if o == nil || o.RefName == nil {
		return nil, false
	}
	return o.RefName, true
}

// HasRefName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasRefName() bool {
	if o != nil && o.RefName != nil {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetRefName(v string) {
	o.RefName = &v
}

// GetRollbackCompleted returns the RollbackCompleted field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackCompleted() bool {
	if o == nil || o.RollbackCompleted == nil {
		var ret bool
		return ret
	}
	return *o.RollbackCompleted
}

// GetRollbackCompletedOk returns a tuple with the RollbackCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackCompletedOk() (*bool, bool) {
	if o == nil || o.RollbackCompleted == nil {
		return nil, false
	}
	return o.RollbackCompleted, true
}

// HasRollbackCompleted returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasRollbackCompleted() bool {
	if o != nil && o.RollbackCompleted != nil {
		return true
	}

	return false
}

// SetRollbackCompleted gets a reference to the given bool and assigns it to the RollbackCompleted field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetRollbackCompleted(v bool) {
	o.RollbackCompleted = &v
}

// GetRollbackTaskName returns the RollbackTaskName field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackTaskName() string {
	if o == nil || o.RollbackTaskName == nil {
		var ret string
		return ret
	}
	return *o.RollbackTaskName
}

// GetRollbackTaskNameOk returns a tuple with the RollbackTaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetRollbackTaskNameOk() (*string, bool) {
	if o == nil || o.RollbackTaskName == nil {
		return nil, false
	}
	return o.RollbackTaskName, true
}

// HasRollbackTaskName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasRollbackTaskName() bool {
	if o != nil && o.RollbackTaskName != nil {
		return true
	}

	return false
}

// SetRollbackTaskName gets a reference to the given string and assigns it to the RollbackTaskName field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetRollbackTaskName(v string) {
	o.RollbackTaskName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTaskInfoMoid returns the TaskInfoMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskInfoMoid() string {
	if o == nil || o.TaskInfoMoid == nil {
		var ret string
		return ret
	}
	return *o.TaskInfoMoid
}

// GetTaskInfoMoidOk returns a tuple with the TaskInfoMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskInfoMoidOk() (*string, bool) {
	if o == nil || o.TaskInfoMoid == nil {
		return nil, false
	}
	return o.TaskInfoMoid, true
}

// HasTaskInfoMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasTaskInfoMoid() bool {
	if o != nil && o.TaskInfoMoid != nil {
		return true
	}

	return false
}

// SetTaskInfoMoid gets a reference to the given string and assigns it to the TaskInfoMoid field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetTaskInfoMoid(v string) {
	o.TaskInfoMoid = &v
}

// GetTaskPath returns the TaskPath field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskPath() string {
	if o == nil || o.TaskPath == nil {
		var ret string
		return ret
	}
	return *o.TaskPath
}

// GetTaskPathOk returns a tuple with the TaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) GetTaskPathOk() (*string, bool) {
	if o == nil || o.TaskPath == nil {
		return nil, false
	}
	return o.TaskPath, true
}

// HasTaskPath returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTaskAllOf) HasTaskPath() bool {
	if o != nil && o.TaskPath != nil {
		return true
	}

	return false
}

// SetTaskPath gets a reference to the given string and assigns it to the TaskPath field.
func (o *WorkflowRollbackWorkflowTaskAllOf) SetTaskPath(v string) {
	o.TaskPath = &v
}

func (o WorkflowRollbackWorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RefName != nil {
		toSerialize["RefName"] = o.RefName
	}
	if o.RollbackCompleted != nil {
		toSerialize["RollbackCompleted"] = o.RollbackCompleted
	}
	if o.RollbackTaskName != nil {
		toSerialize["RollbackTaskName"] = o.RollbackTaskName
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TaskInfoMoid != nil {
		toSerialize["TaskInfoMoid"] = o.TaskInfoMoid
	}
	if o.TaskPath != nil {
		toSerialize["TaskPath"] = o.TaskPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowRollbackWorkflowTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowRollbackWorkflowTaskAllOf := _WorkflowRollbackWorkflowTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowRollbackWorkflowTaskAllOf); err == nil {
		*o = WorkflowRollbackWorkflowTaskAllOf(varWorkflowRollbackWorkflowTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RefName")
		delete(additionalProperties, "RollbackCompleted")
		delete(additionalProperties, "RollbackTaskName")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskInfoMoid")
		delete(additionalProperties, "TaskPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowRollbackWorkflowTaskAllOf struct {
	value *WorkflowRollbackWorkflowTaskAllOf
	isSet bool
}

func (v NullableWorkflowRollbackWorkflowTaskAllOf) Get() *WorkflowRollbackWorkflowTaskAllOf {
	return v.value
}

func (v *NullableWorkflowRollbackWorkflowTaskAllOf) Set(val *WorkflowRollbackWorkflowTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRollbackWorkflowTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRollbackWorkflowTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRollbackWorkflowTaskAllOf(val *WorkflowRollbackWorkflowTaskAllOf) *NullableWorkflowRollbackWorkflowTaskAllOf {
	return &NullableWorkflowRollbackWorkflowTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowRollbackWorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRollbackWorkflowTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
