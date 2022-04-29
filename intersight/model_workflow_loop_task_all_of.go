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

// WorkflowLoopTaskAllOf Definition of the list of properties defined in 'workflow.LoopTask', excluding properties defined in parent classes.
type WorkflowLoopTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When tasks are run in parallel and the count is large, the actual number of task run in parallel can be controlled by this property. If count is 100 and numberOfBatches is 5 then 20 tasks are run in parallel 5 times. Parallel batch size must be less than the count. In cases where count is dynamic and depends on input given during workflow execution, if that count is less than batch then empty batches might get created which do not have any tasks under them.
	NumberOfBatches *int64 `json:"NumberOfBatches,omitempty"`
	// This field is deprecated. Always set to true for parallel loop.
	// Deprecated
	Parallel             *bool `json:"Parallel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLoopTaskAllOf WorkflowLoopTaskAllOf

// NewWorkflowLoopTaskAllOf instantiates a new WorkflowLoopTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLoopTaskAllOf(classId string, objectType string) *WorkflowLoopTaskAllOf {
	this := WorkflowLoopTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var numberOfBatches int64 = 1
	this.NumberOfBatches = &numberOfBatches
	var parallel bool = true
	this.Parallel = &parallel
	return &this
}

// NewWorkflowLoopTaskAllOfWithDefaults instantiates a new WorkflowLoopTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLoopTaskAllOfWithDefaults() *WorkflowLoopTaskAllOf {
	this := WorkflowLoopTaskAllOf{}
	var classId string = "workflow.LoopTask"
	this.ClassId = classId
	var objectType string = "workflow.LoopTask"
	this.ObjectType = objectType
	var numberOfBatches int64 = 1
	this.NumberOfBatches = &numberOfBatches
	var parallel bool = true
	this.Parallel = &parallel
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowLoopTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowLoopTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowLoopTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowLoopTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumberOfBatches returns the NumberOfBatches field value if set, zero value otherwise.
func (o *WorkflowLoopTaskAllOf) GetNumberOfBatches() int64 {
	if o == nil || o.NumberOfBatches == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfBatches
}

// GetNumberOfBatchesOk returns a tuple with the NumberOfBatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLoopTaskAllOf) GetNumberOfBatchesOk() (*int64, bool) {
	if o == nil || o.NumberOfBatches == nil {
		return nil, false
	}
	return o.NumberOfBatches, true
}

// HasNumberOfBatches returns a boolean if a field has been set.
func (o *WorkflowLoopTaskAllOf) HasNumberOfBatches() bool {
	if o != nil && o.NumberOfBatches != nil {
		return true
	}

	return false
}

// SetNumberOfBatches gets a reference to the given int64 and assigns it to the NumberOfBatches field.
func (o *WorkflowLoopTaskAllOf) SetNumberOfBatches(v int64) {
	o.NumberOfBatches = &v
}

// GetParallel returns the Parallel field value if set, zero value otherwise.
// Deprecated
func (o *WorkflowLoopTaskAllOf) GetParallel() bool {
	if o == nil || o.Parallel == nil {
		var ret bool
		return ret
	}
	return *o.Parallel
}

// GetParallelOk returns a tuple with the Parallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *WorkflowLoopTaskAllOf) GetParallelOk() (*bool, bool) {
	if o == nil || o.Parallel == nil {
		return nil, false
	}
	return o.Parallel, true
}

// HasParallel returns a boolean if a field has been set.
func (o *WorkflowLoopTaskAllOf) HasParallel() bool {
	if o != nil && o.Parallel != nil {
		return true
	}

	return false
}

// SetParallel gets a reference to the given bool and assigns it to the Parallel field.
// Deprecated
func (o *WorkflowLoopTaskAllOf) SetParallel(v bool) {
	o.Parallel = &v
}

func (o WorkflowLoopTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NumberOfBatches != nil {
		toSerialize["NumberOfBatches"] = o.NumberOfBatches
	}
	if o.Parallel != nil {
		toSerialize["Parallel"] = o.Parallel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowLoopTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowLoopTaskAllOf := _WorkflowLoopTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowLoopTaskAllOf); err == nil {
		*o = WorkflowLoopTaskAllOf(varWorkflowLoopTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NumberOfBatches")
		delete(additionalProperties, "Parallel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLoopTaskAllOf struct {
	value *WorkflowLoopTaskAllOf
	isSet bool
}

func (v NullableWorkflowLoopTaskAllOf) Get() *WorkflowLoopTaskAllOf {
	return v.value
}

func (v *NullableWorkflowLoopTaskAllOf) Set(val *WorkflowLoopTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLoopTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLoopTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLoopTaskAllOf(val *WorkflowLoopTaskAllOf) *NullableWorkflowLoopTaskAllOf {
	return &NullableWorkflowLoopTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowLoopTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLoopTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
