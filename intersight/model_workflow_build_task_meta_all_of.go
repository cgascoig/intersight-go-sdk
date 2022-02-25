/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowBuildTaskMetaAllOf Definition of the list of properties defined in 'workflow.BuildTaskMeta', excluding properties defined in parent classes.
type WorkflowBuildTaskMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name for the BuildTaskMeta instance used to created a dynamic workflow.
	Name *string `json:"Name,omitempty"`
	// Microservice owner for the task in this workflow.
	Src *string `json:"Src,omitempty"`
	// Task list used to build the dynamic workflow.
	TaskList interface{} `json:"TaskList,omitempty"`
	// The type of the task within this workflow.
	TaskType *string `json:"TaskType,omitempty"`
	// The type for the dynamic workflow.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBuildTaskMetaAllOf WorkflowBuildTaskMetaAllOf

// NewWorkflowBuildTaskMetaAllOf instantiates a new WorkflowBuildTaskMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBuildTaskMetaAllOf(classId string, objectType string) *WorkflowBuildTaskMetaAllOf {
	this := WorkflowBuildTaskMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBuildTaskMetaAllOfWithDefaults instantiates a new WorkflowBuildTaskMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBuildTaskMetaAllOfWithDefaults() *WorkflowBuildTaskMetaAllOf {
	this := WorkflowBuildTaskMetaAllOf{}
	var classId string = "workflow.BuildTaskMeta"
	this.ClassId = classId
	var objectType string = "workflow.BuildTaskMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBuildTaskMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBuildTaskMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBuildTaskMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBuildTaskMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBuildTaskMetaAllOf) SetName(v string) {
	o.Name = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaAllOf) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaAllOf) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *WorkflowBuildTaskMetaAllOf) SetSrc(v string) {
	o.Src = &v
}

// GetTaskList returns the TaskList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBuildTaskMetaAllOf) GetTaskList() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaskList
}

// GetTaskListOk returns a tuple with the TaskList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBuildTaskMetaAllOf) GetTaskListOk() (*interface{}, bool) {
	if o == nil || o.TaskList == nil {
		return nil, false
	}
	return &o.TaskList, true
}

// HasTaskList returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaAllOf) HasTaskList() bool {
	if o != nil && o.TaskList != nil {
		return true
	}

	return false
}

// SetTaskList gets a reference to the given interface{} and assigns it to the TaskList field.
func (o *WorkflowBuildTaskMetaAllOf) SetTaskList(v interface{}) {
	o.TaskList = v
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaAllOf) GetTaskType() string {
	if o == nil || o.TaskType == nil {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetTaskTypeOk() (*string, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaAllOf) HasTaskType() bool {
	if o != nil && o.TaskType != nil {
		return true
	}

	return false
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *WorkflowBuildTaskMetaAllOf) SetTaskType(v string) {
	o.TaskType = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaAllOf) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaAllOf) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaAllOf) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowBuildTaskMetaAllOf) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o WorkflowBuildTaskMetaAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Src != nil {
		toSerialize["Src"] = o.Src
	}
	if o.TaskList != nil {
		toSerialize["TaskList"] = o.TaskList
	}
	if o.TaskType != nil {
		toSerialize["TaskType"] = o.TaskType
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBuildTaskMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowBuildTaskMetaAllOf := _WorkflowBuildTaskMetaAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowBuildTaskMetaAllOf); err == nil {
		*o = WorkflowBuildTaskMetaAllOf(varWorkflowBuildTaskMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Src")
		delete(additionalProperties, "TaskList")
		delete(additionalProperties, "TaskType")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowBuildTaskMetaAllOf struct {
	value *WorkflowBuildTaskMetaAllOf
	isSet bool
}

func (v NullableWorkflowBuildTaskMetaAllOf) Get() *WorkflowBuildTaskMetaAllOf {
	return v.value
}

func (v *NullableWorkflowBuildTaskMetaAllOf) Set(val *WorkflowBuildTaskMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBuildTaskMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBuildTaskMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBuildTaskMetaAllOf(val *WorkflowBuildTaskMetaAllOf) *NullableWorkflowBuildTaskMetaAllOf {
	return &NullableWorkflowBuildTaskMetaAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBuildTaskMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBuildTaskMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
