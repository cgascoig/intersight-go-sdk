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

// WorkflowPendingDynamicWorkflowInfoAllOf Definition of the list of properties defined in 'workflow.PendingDynamicWorkflowInfo', excluding properties defined in parent classes.
type WorkflowPendingDynamicWorkflowInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The inputs of the workflow.
	Input interface{} `json:"Input,omitempty"`
	// A name for the pending dynamic workflow.
	Name            *string  `json:"Name,omitempty"`
	PendingServices []string `json:"PendingServices,omitempty"`
	// The src is workflow owner service.
	Src *string `json:"Src,omitempty"`
	// The current status of the PendingDynamicWorkflowInfo. * `GatheringTasks` - Dynamic workflow is gathering tasks before workflow can start execution. * `Waiting` - Dynamic workflow is in waiting state and not yet started execution. * `RateLimit` - Dynamic workflow is rate limited and hasn't started execution.
	Status *string `json:"Status,omitempty"`
	// When set to true workflow engine will wait for a duplicate to finish before starting a new one.
	WaitOnDuplicate         *bool                                   `json:"WaitOnDuplicate,omitempty"`
	WorkflowActionTaskLists []WorkflowDynamicWorkflowActionTaskList `json:"WorkflowActionTaskLists,omitempty"`
	WorkflowCtx             NullableWorkflowWorkflowCtx             `json:"WorkflowCtx,omitempty"`
	// This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup.
	WorkflowKey *string `json:"WorkflowKey,omitempty"`
	// The metadata of the workflow.
	WorkflowMeta         interface{}                       `json:"WorkflowMeta,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPendingDynamicWorkflowInfoAllOf WorkflowPendingDynamicWorkflowInfoAllOf

// NewWorkflowPendingDynamicWorkflowInfoAllOf instantiates a new WorkflowPendingDynamicWorkflowInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPendingDynamicWorkflowInfoAllOf(classId string, objectType string) *WorkflowPendingDynamicWorkflowInfoAllOf {
	this := WorkflowPendingDynamicWorkflowInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "GatheringTasks"
	this.Status = &status
	return &this
}

// NewWorkflowPendingDynamicWorkflowInfoAllOfWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPendingDynamicWorkflowInfoAllOfWithDefaults() *WorkflowPendingDynamicWorkflowInfoAllOf {
	this := WorkflowPendingDynamicWorkflowInfoAllOf{}
	var classId string = "workflow.PendingDynamicWorkflowInfo"
	this.ClassId = classId
	var objectType string = "workflow.PendingDynamicWorkflowInfo"
	this.ObjectType = objectType
	var status string = "GatheringTasks"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetInput(v interface{}) {
	o.Input = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetName(v string) {
	o.Name = &v
}

// GetPendingServices returns the PendingServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetPendingServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PendingServices
}

// GetPendingServicesOk returns a tuple with the PendingServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetPendingServicesOk() (*[]string, bool) {
	if o == nil || o.PendingServices == nil {
		return nil, false
	}
	return &o.PendingServices, true
}

// HasPendingServices returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasPendingServices() bool {
	if o != nil && o.PendingServices != nil {
		return true
	}

	return false
}

// SetPendingServices gets a reference to the given []string and assigns it to the PendingServices field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetPendingServices(v []string) {
	o.PendingServices = v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetSrc() string {
	if o == nil || o.Src == nil {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetSrcOk() (*string, bool) {
	if o == nil || o.Src == nil {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasSrc() bool {
	if o != nil && o.Src != nil {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetSrc(v string) {
	o.Src = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetWaitOnDuplicate returns the WaitOnDuplicate field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWaitOnDuplicate() bool {
	if o == nil || o.WaitOnDuplicate == nil {
		var ret bool
		return ret
	}
	return *o.WaitOnDuplicate
}

// GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWaitOnDuplicateOk() (*bool, bool) {
	if o == nil || o.WaitOnDuplicate == nil {
		return nil, false
	}
	return o.WaitOnDuplicate, true
}

// HasWaitOnDuplicate returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWaitOnDuplicate() bool {
	if o != nil && o.WaitOnDuplicate != nil {
		return true
	}

	return false
}

// SetWaitOnDuplicate gets a reference to the given bool and assigns it to the WaitOnDuplicate field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWaitOnDuplicate(v bool) {
	o.WaitOnDuplicate = &v
}

// GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList {
	if o == nil {
		var ret []WorkflowDynamicWorkflowActionTaskList
		return ret
	}
	return o.WorkflowActionTaskLists
}

// GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowActionTaskListsOk() (*[]WorkflowDynamicWorkflowActionTaskList, bool) {
	if o == nil || o.WorkflowActionTaskLists == nil {
		return nil, false
	}
	return &o.WorkflowActionTaskLists, true
}

// HasWorkflowActionTaskLists returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowActionTaskLists() bool {
	if o != nil && o.WorkflowActionTaskLists != nil {
		return true
	}

	return false
}

// SetWorkflowActionTaskLists gets a reference to the given []WorkflowDynamicWorkflowActionTaskList and assigns it to the WorkflowActionTaskLists field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList) {
	o.WorkflowActionTaskLists = v
}

// GetWorkflowCtx returns the WorkflowCtx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowCtx() WorkflowWorkflowCtx {
	if o == nil || o.WorkflowCtx.Get() == nil {
		var ret WorkflowWorkflowCtx
		return ret
	}
	return *o.WorkflowCtx.Get()
}

// GetWorkflowCtxOk returns a tuple with the WorkflowCtx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowCtx.Get(), o.WorkflowCtx.IsSet()
}

// HasWorkflowCtx returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowCtx() bool {
	if o != nil && o.WorkflowCtx.IsSet() {
		return true
	}

	return false
}

// SetWorkflowCtx gets a reference to the given NullableWorkflowWorkflowCtx and assigns it to the WorkflowCtx field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowCtx(v WorkflowWorkflowCtx) {
	o.WorkflowCtx.Set(&v)
}

// SetWorkflowCtxNil sets the value for WorkflowCtx to be an explicit nil
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowCtxNil() {
	o.WorkflowCtx.Set(nil)
}

// UnsetWorkflowCtx ensures that no value is present for WorkflowCtx, not even an explicit nil
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) UnsetWorkflowCtx() {
	o.WorkflowCtx.Unset()
}

// GetWorkflowKey returns the WorkflowKey field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowKey() string {
	if o == nil || o.WorkflowKey == nil {
		var ret string
		return ret
	}
	return *o.WorkflowKey
}

// GetWorkflowKeyOk returns a tuple with the WorkflowKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowKeyOk() (*string, bool) {
	if o == nil || o.WorkflowKey == nil {
		return nil, false
	}
	return o.WorkflowKey, true
}

// HasWorkflowKey returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowKey() bool {
	if o != nil && o.WorkflowKey != nil {
		return true
	}

	return false
}

// SetWorkflowKey gets a reference to the given string and assigns it to the WorkflowKey field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowKey(v string) {
	o.WorkflowKey = &v
}

// GetWorkflowMeta returns the WorkflowMeta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.WorkflowMeta
}

// GetWorkflowMetaOk returns a tuple with the WorkflowMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowMetaOk() (*interface{}, bool) {
	if o == nil || o.WorkflowMeta == nil {
		return nil, false
	}
	return &o.WorkflowMeta, true
}

// HasWorkflowMeta returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowMeta() bool {
	if o != nil && o.WorkflowMeta != nil {
		return true
	}

	return false
}

// SetWorkflowMeta gets a reference to the given interface{} and assigns it to the WorkflowMeta field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowMeta(v interface{}) {
	o.WorkflowMeta = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o WorkflowPendingDynamicWorkflowInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PendingServices != nil {
		toSerialize["PendingServices"] = o.PendingServices
	}
	if o.Src != nil {
		toSerialize["Src"] = o.Src
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.WaitOnDuplicate != nil {
		toSerialize["WaitOnDuplicate"] = o.WaitOnDuplicate
	}
	if o.WorkflowActionTaskLists != nil {
		toSerialize["WorkflowActionTaskLists"] = o.WorkflowActionTaskLists
	}
	if o.WorkflowCtx.IsSet() {
		toSerialize["WorkflowCtx"] = o.WorkflowCtx.Get()
	}
	if o.WorkflowKey != nil {
		toSerialize["WorkflowKey"] = o.WorkflowKey
	}
	if o.WorkflowMeta != nil {
		toSerialize["WorkflowMeta"] = o.WorkflowMeta
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPendingDynamicWorkflowInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowPendingDynamicWorkflowInfoAllOf := _WorkflowPendingDynamicWorkflowInfoAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowPendingDynamicWorkflowInfoAllOf); err == nil {
		*o = WorkflowPendingDynamicWorkflowInfoAllOf(varWorkflowPendingDynamicWorkflowInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PendingServices")
		delete(additionalProperties, "Src")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "WaitOnDuplicate")
		delete(additionalProperties, "WorkflowActionTaskLists")
		delete(additionalProperties, "WorkflowCtx")
		delete(additionalProperties, "WorkflowKey")
		delete(additionalProperties, "WorkflowMeta")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPendingDynamicWorkflowInfoAllOf struct {
	value *WorkflowPendingDynamicWorkflowInfoAllOf
	isSet bool
}

func (v NullableWorkflowPendingDynamicWorkflowInfoAllOf) Get() *WorkflowPendingDynamicWorkflowInfoAllOf {
	return v.value
}

func (v *NullableWorkflowPendingDynamicWorkflowInfoAllOf) Set(val *WorkflowPendingDynamicWorkflowInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPendingDynamicWorkflowInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPendingDynamicWorkflowInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPendingDynamicWorkflowInfoAllOf(val *WorkflowPendingDynamicWorkflowInfoAllOf) *NullableWorkflowPendingDynamicWorkflowInfoAllOf {
	return &NullableWorkflowPendingDynamicWorkflowInfoAllOf{value: val, isSet: true}
}

func (v NullableWorkflowPendingDynamicWorkflowInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPendingDynamicWorkflowInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
