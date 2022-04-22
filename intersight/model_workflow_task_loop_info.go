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
	"reflect"
	"strings"
)

// WorkflowTaskLoopInfo Keep track of loop information for every task executed inside the loop.
type WorkflowTaskLoopInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This specifies the count of iteration for the specific task executed inside the loop.
	Iteration *int64 `json:"Iteration,omitempty"`
	// Label of the loop task inside which this task is executed.
	LoopTaskLabel *string `json:"LoopTaskLabel,omitempty"`
	// Name of the loop task inside which this task is executed.
	LoopTaskName *string `json:"LoopTaskName,omitempty"`
	// This specifies the type of loop, Serial or Parallel. * `None` - The enum specifies the option as None which implies this is not a Loop type and this is the default value for loop type. * `Parallel` - The enum specifies the option as Parallel where the loop task type is parallel loop. * `Serial` - The enum specifies the option as Serial where the loop task type is serial loop.
	LoopType             *string `json:"LoopType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTaskLoopInfo WorkflowTaskLoopInfo

// NewWorkflowTaskLoopInfo instantiates a new WorkflowTaskLoopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTaskLoopInfo(classId string, objectType string) *WorkflowTaskLoopInfo {
	this := WorkflowTaskLoopInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var loopType string = "None"
	this.LoopType = &loopType
	return &this
}

// NewWorkflowTaskLoopInfoWithDefaults instantiates a new WorkflowTaskLoopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTaskLoopInfoWithDefaults() *WorkflowTaskLoopInfo {
	this := WorkflowTaskLoopInfo{}
	var classId string = "workflow.TaskLoopInfo"
	this.ClassId = classId
	var objectType string = "workflow.TaskLoopInfo"
	this.ObjectType = objectType
	var loopType string = "None"
	this.LoopType = &loopType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowTaskLoopInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowTaskLoopInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowTaskLoopInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowTaskLoopInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfo) GetIteration() int64 {
	if o == nil || o.Iteration == nil {
		var ret int64
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetIterationOk() (*int64, bool) {
	if o == nil || o.Iteration == nil {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfo) HasIteration() bool {
	if o != nil && o.Iteration != nil {
		return true
	}

	return false
}

// SetIteration gets a reference to the given int64 and assigns it to the Iteration field.
func (o *WorkflowTaskLoopInfo) SetIteration(v int64) {
	o.Iteration = &v
}

// GetLoopTaskLabel returns the LoopTaskLabel field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfo) GetLoopTaskLabel() string {
	if o == nil || o.LoopTaskLabel == nil {
		var ret string
		return ret
	}
	return *o.LoopTaskLabel
}

// GetLoopTaskLabelOk returns a tuple with the LoopTaskLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetLoopTaskLabelOk() (*string, bool) {
	if o == nil || o.LoopTaskLabel == nil {
		return nil, false
	}
	return o.LoopTaskLabel, true
}

// HasLoopTaskLabel returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfo) HasLoopTaskLabel() bool {
	if o != nil && o.LoopTaskLabel != nil {
		return true
	}

	return false
}

// SetLoopTaskLabel gets a reference to the given string and assigns it to the LoopTaskLabel field.
func (o *WorkflowTaskLoopInfo) SetLoopTaskLabel(v string) {
	o.LoopTaskLabel = &v
}

// GetLoopTaskName returns the LoopTaskName field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfo) GetLoopTaskName() string {
	if o == nil || o.LoopTaskName == nil {
		var ret string
		return ret
	}
	return *o.LoopTaskName
}

// GetLoopTaskNameOk returns a tuple with the LoopTaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetLoopTaskNameOk() (*string, bool) {
	if o == nil || o.LoopTaskName == nil {
		return nil, false
	}
	return o.LoopTaskName, true
}

// HasLoopTaskName returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfo) HasLoopTaskName() bool {
	if o != nil && o.LoopTaskName != nil {
		return true
	}

	return false
}

// SetLoopTaskName gets a reference to the given string and assigns it to the LoopTaskName field.
func (o *WorkflowTaskLoopInfo) SetLoopTaskName(v string) {
	o.LoopTaskName = &v
}

// GetLoopType returns the LoopType field value if set, zero value otherwise.
func (o *WorkflowTaskLoopInfo) GetLoopType() string {
	if o == nil || o.LoopType == nil {
		var ret string
		return ret
	}
	return *o.LoopType
}

// GetLoopTypeOk returns a tuple with the LoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTaskLoopInfo) GetLoopTypeOk() (*string, bool) {
	if o == nil || o.LoopType == nil {
		return nil, false
	}
	return o.LoopType, true
}

// HasLoopType returns a boolean if a field has been set.
func (o *WorkflowTaskLoopInfo) HasLoopType() bool {
	if o != nil && o.LoopType != nil {
		return true
	}

	return false
}

// SetLoopType gets a reference to the given string and assigns it to the LoopType field.
func (o *WorkflowTaskLoopInfo) SetLoopType(v string) {
	o.LoopType = &v
}

func (o WorkflowTaskLoopInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Iteration != nil {
		toSerialize["Iteration"] = o.Iteration
	}
	if o.LoopTaskLabel != nil {
		toSerialize["LoopTaskLabel"] = o.LoopTaskLabel
	}
	if o.LoopTaskName != nil {
		toSerialize["LoopTaskName"] = o.LoopTaskName
	}
	if o.LoopType != nil {
		toSerialize["LoopType"] = o.LoopType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTaskLoopInfo) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowTaskLoopInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This specifies the count of iteration for the specific task executed inside the loop.
		Iteration *int64 `json:"Iteration,omitempty"`
		// Label of the loop task inside which this task is executed.
		LoopTaskLabel *string `json:"LoopTaskLabel,omitempty"`
		// Name of the loop task inside which this task is executed.
		LoopTaskName *string `json:"LoopTaskName,omitempty"`
		// This specifies the type of loop, Serial or Parallel. * `None` - The enum specifies the option as None which implies this is not a Loop type and this is the default value for loop type. * `Parallel` - The enum specifies the option as Parallel where the loop task type is parallel loop. * `Serial` - The enum specifies the option as Serial where the loop task type is serial loop.
		LoopType *string `json:"LoopType,omitempty"`
	}

	varWorkflowTaskLoopInfoWithoutEmbeddedStruct := WorkflowTaskLoopInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowTaskLoopInfoWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowTaskLoopInfo := _WorkflowTaskLoopInfo{}
		varWorkflowTaskLoopInfo.ClassId = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.ClassId
		varWorkflowTaskLoopInfo.ObjectType = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.ObjectType
		varWorkflowTaskLoopInfo.Iteration = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.Iteration
		varWorkflowTaskLoopInfo.LoopTaskLabel = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.LoopTaskLabel
		varWorkflowTaskLoopInfo.LoopTaskName = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.LoopTaskName
		varWorkflowTaskLoopInfo.LoopType = varWorkflowTaskLoopInfoWithoutEmbeddedStruct.LoopType
		*o = WorkflowTaskLoopInfo(varWorkflowTaskLoopInfo)
	} else {
		return err
	}

	varWorkflowTaskLoopInfo := _WorkflowTaskLoopInfo{}

	err = json.Unmarshal(bytes, &varWorkflowTaskLoopInfo)
	if err == nil {
		o.MoBaseComplexType = varWorkflowTaskLoopInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Iteration")
		delete(additionalProperties, "LoopTaskLabel")
		delete(additionalProperties, "LoopTaskName")
		delete(additionalProperties, "LoopType")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTaskLoopInfo struct {
	value *WorkflowTaskLoopInfo
	isSet bool
}

func (v NullableWorkflowTaskLoopInfo) Get() *WorkflowTaskLoopInfo {
	return v.value
}

func (v *NullableWorkflowTaskLoopInfo) Set(val *WorkflowTaskLoopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTaskLoopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTaskLoopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTaskLoopInfo(val *WorkflowTaskLoopInfo) *NullableWorkflowTaskLoopInfo {
	return &NullableWorkflowTaskLoopInfo{value: val, isSet: true}
}

func (v NullableWorkflowTaskLoopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTaskLoopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
