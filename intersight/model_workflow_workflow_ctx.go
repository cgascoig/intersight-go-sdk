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
	"reflect"
	"strings"
)

// WorkflowWorkflowCtx The workflow context contains initiator and target information along with workflow type and action information.
type WorkflowWorkflowCtx struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                           `json:"ObjectType"`
	InitiatorCtx  NullableWorkflowInitiatorContext `json:"InitiatorCtx,omitempty"`
	TargetCtxList []WorkflowTargetContext          `json:"TargetCtxList,omitempty"`
	// The name of workflowMeta of the workflow running.
	WorkflowMetaName *string `json:"WorkflowMetaName,omitempty"`
	// The subtype of the workflow.
	WorkflowSubtype *string `json:"WorkflowSubtype,omitempty"`
	// Type of the workflow being started. This can be any string for client services to distinguish workflow by type.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowCtx WorkflowWorkflowCtx

// NewWorkflowWorkflowCtx instantiates a new WorkflowWorkflowCtx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowCtx(classId string, objectType string) *WorkflowWorkflowCtx {
	this := WorkflowWorkflowCtx{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWorkflowCtxWithDefaults instantiates a new WorkflowWorkflowCtx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowCtxWithDefaults() *WorkflowWorkflowCtx {
	this := WorkflowWorkflowCtx{}
	var classId string = "workflow.WorkflowCtx"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowCtx"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowCtx) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtx) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowCtx) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowCtx) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtx) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowCtx) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInitiatorCtx returns the InitiatorCtx field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowCtx) GetInitiatorCtx() WorkflowInitiatorContext {
	if o == nil || o.InitiatorCtx.Get() == nil {
		var ret WorkflowInitiatorContext
		return ret
	}
	return *o.InitiatorCtx.Get()
}

// GetInitiatorCtxOk returns a tuple with the InitiatorCtx field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowCtx) GetInitiatorCtxOk() (*WorkflowInitiatorContext, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitiatorCtx.Get(), o.InitiatorCtx.IsSet()
}

// HasInitiatorCtx returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtx) HasInitiatorCtx() bool {
	if o != nil && o.InitiatorCtx.IsSet() {
		return true
	}

	return false
}

// SetInitiatorCtx gets a reference to the given NullableWorkflowInitiatorContext and assigns it to the InitiatorCtx field.
func (o *WorkflowWorkflowCtx) SetInitiatorCtx(v WorkflowInitiatorContext) {
	o.InitiatorCtx.Set(&v)
}

// SetInitiatorCtxNil sets the value for InitiatorCtx to be an explicit nil
func (o *WorkflowWorkflowCtx) SetInitiatorCtxNil() {
	o.InitiatorCtx.Set(nil)
}

// UnsetInitiatorCtx ensures that no value is present for InitiatorCtx, not even an explicit nil
func (o *WorkflowWorkflowCtx) UnsetInitiatorCtx() {
	o.InitiatorCtx.Unset()
}

// GetTargetCtxList returns the TargetCtxList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowCtx) GetTargetCtxList() []WorkflowTargetContext {
	if o == nil {
		var ret []WorkflowTargetContext
		return ret
	}
	return o.TargetCtxList
}

// GetTargetCtxListOk returns a tuple with the TargetCtxList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowCtx) GetTargetCtxListOk() (*[]WorkflowTargetContext, bool) {
	if o == nil || o.TargetCtxList == nil {
		return nil, false
	}
	return &o.TargetCtxList, true
}

// HasTargetCtxList returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtx) HasTargetCtxList() bool {
	if o != nil && o.TargetCtxList != nil {
		return true
	}

	return false
}

// SetTargetCtxList gets a reference to the given []WorkflowTargetContext and assigns it to the TargetCtxList field.
func (o *WorkflowWorkflowCtx) SetTargetCtxList(v []WorkflowTargetContext) {
	o.TargetCtxList = v
}

// GetWorkflowMetaName returns the WorkflowMetaName field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtx) GetWorkflowMetaName() string {
	if o == nil || o.WorkflowMetaName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowMetaName
}

// GetWorkflowMetaNameOk returns a tuple with the WorkflowMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtx) GetWorkflowMetaNameOk() (*string, bool) {
	if o == nil || o.WorkflowMetaName == nil {
		return nil, false
	}
	return o.WorkflowMetaName, true
}

// HasWorkflowMetaName returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtx) HasWorkflowMetaName() bool {
	if o != nil && o.WorkflowMetaName != nil {
		return true
	}

	return false
}

// SetWorkflowMetaName gets a reference to the given string and assigns it to the WorkflowMetaName field.
func (o *WorkflowWorkflowCtx) SetWorkflowMetaName(v string) {
	o.WorkflowMetaName = &v
}

// GetWorkflowSubtype returns the WorkflowSubtype field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtx) GetWorkflowSubtype() string {
	if o == nil || o.WorkflowSubtype == nil {
		var ret string
		return ret
	}
	return *o.WorkflowSubtype
}

// GetWorkflowSubtypeOk returns a tuple with the WorkflowSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtx) GetWorkflowSubtypeOk() (*string, bool) {
	if o == nil || o.WorkflowSubtype == nil {
		return nil, false
	}
	return o.WorkflowSubtype, true
}

// HasWorkflowSubtype returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtx) HasWorkflowSubtype() bool {
	if o != nil && o.WorkflowSubtype != nil {
		return true
	}

	return false
}

// SetWorkflowSubtype gets a reference to the given string and assigns it to the WorkflowSubtype field.
func (o *WorkflowWorkflowCtx) SetWorkflowSubtype(v string) {
	o.WorkflowSubtype = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowWorkflowCtx) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowCtx) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowWorkflowCtx) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowWorkflowCtx) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o WorkflowWorkflowCtx) MarshalJSON() ([]byte, error) {
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
	if o.InitiatorCtx.IsSet() {
		toSerialize["InitiatorCtx"] = o.InitiatorCtx.Get()
	}
	if o.TargetCtxList != nil {
		toSerialize["TargetCtxList"] = o.TargetCtxList
	}
	if o.WorkflowMetaName != nil {
		toSerialize["WorkflowMetaName"] = o.WorkflowMetaName
	}
	if o.WorkflowSubtype != nil {
		toSerialize["WorkflowSubtype"] = o.WorkflowSubtype
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowCtx) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkflowCtxWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                           `json:"ObjectType"`
		InitiatorCtx  NullableWorkflowInitiatorContext `json:"InitiatorCtx,omitempty"`
		TargetCtxList []WorkflowTargetContext          `json:"TargetCtxList,omitempty"`
		// The name of workflowMeta of the workflow running.
		WorkflowMetaName *string `json:"WorkflowMetaName,omitempty"`
		// The subtype of the workflow.
		WorkflowSubtype *string `json:"WorkflowSubtype,omitempty"`
		// Type of the workflow being started. This can be any string for client services to distinguish workflow by type.
		WorkflowType *string `json:"WorkflowType,omitempty"`
	}

	varWorkflowWorkflowCtxWithoutEmbeddedStruct := WorkflowWorkflowCtxWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowCtxWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkflowCtx := _WorkflowWorkflowCtx{}
		varWorkflowWorkflowCtx.ClassId = varWorkflowWorkflowCtxWithoutEmbeddedStruct.ClassId
		varWorkflowWorkflowCtx.ObjectType = varWorkflowWorkflowCtxWithoutEmbeddedStruct.ObjectType
		varWorkflowWorkflowCtx.InitiatorCtx = varWorkflowWorkflowCtxWithoutEmbeddedStruct.InitiatorCtx
		varWorkflowWorkflowCtx.TargetCtxList = varWorkflowWorkflowCtxWithoutEmbeddedStruct.TargetCtxList
		varWorkflowWorkflowCtx.WorkflowMetaName = varWorkflowWorkflowCtxWithoutEmbeddedStruct.WorkflowMetaName
		varWorkflowWorkflowCtx.WorkflowSubtype = varWorkflowWorkflowCtxWithoutEmbeddedStruct.WorkflowSubtype
		varWorkflowWorkflowCtx.WorkflowType = varWorkflowWorkflowCtxWithoutEmbeddedStruct.WorkflowType
		*o = WorkflowWorkflowCtx(varWorkflowWorkflowCtx)
	} else {
		return err
	}

	varWorkflowWorkflowCtx := _WorkflowWorkflowCtx{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowCtx)
	if err == nil {
		o.MoBaseComplexType = varWorkflowWorkflowCtx.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitiatorCtx")
		delete(additionalProperties, "TargetCtxList")
		delete(additionalProperties, "WorkflowMetaName")
		delete(additionalProperties, "WorkflowSubtype")
		delete(additionalProperties, "WorkflowType")

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

type NullableWorkflowWorkflowCtx struct {
	value *WorkflowWorkflowCtx
	isSet bool
}

func (v NullableWorkflowWorkflowCtx) Get() *WorkflowWorkflowCtx {
	return v.value
}

func (v *NullableWorkflowWorkflowCtx) Set(val *WorkflowWorkflowCtx) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowCtx) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowCtx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowCtx(val *WorkflowWorkflowCtx) *NullableWorkflowWorkflowCtx {
	return &NullableWorkflowWorkflowCtx{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowCtx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowCtx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
