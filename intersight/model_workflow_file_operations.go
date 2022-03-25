/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowFileOperations This models a single File Operation request within a batch of requests that get executed within a single workflow task.
type WorkflowFileOperations struct {
	WorkflowApi
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                         `json:"ObjectType"`
	FileDownload NullableWorkflowFileDownloadOp `json:"FileDownload,omitempty"`
	FileTemplate NullableWorkflowFileTemplateOp `json:"FileTemplate,omitempty"`
	// File operation type to be executed on the connected endpoint. * `FileDownload` - The API is executed in a remote device connected to the Intersightthrough its device connector. This operation is to download the filefrom specified storage bucket to the specific path on the device. * `FileTemplatize` - Populates data driven template file with input values to generate textual output.Inputs - the path of the template file on the device and json values to populate.An error will be returned if the file does not exists or if there is an error whileexecuting the template.
	OperationType        *string `json:"OperationType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowFileOperations WorkflowFileOperations

// NewWorkflowFileOperations instantiates a new WorkflowFileOperations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFileOperations(classId string, objectType string) *WorkflowFileOperations {
	this := WorkflowFileOperations{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationType string = "FileDownload"
	this.OperationType = &operationType
	return &this
}

// NewWorkflowFileOperationsWithDefaults instantiates a new WorkflowFileOperations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFileOperationsWithDefaults() *WorkflowFileOperations {
	this := WorkflowFileOperations{}
	var classId string = "workflow.FileOperations"
	this.ClassId = classId
	var objectType string = "workflow.FileOperations"
	this.ObjectType = objectType
	var operationType string = "FileDownload"
	this.OperationType = &operationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowFileOperations) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileOperations) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowFileOperations) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowFileOperations) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileOperations) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowFileOperations) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileDownload returns the FileDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowFileOperations) GetFileDownload() WorkflowFileDownloadOp {
	if o == nil || o.FileDownload.Get() == nil {
		var ret WorkflowFileDownloadOp
		return ret
	}
	return *o.FileDownload.Get()
}

// GetFileDownloadOk returns a tuple with the FileDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowFileOperations) GetFileDownloadOk() (*WorkflowFileDownloadOp, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileDownload.Get(), o.FileDownload.IsSet()
}

// HasFileDownload returns a boolean if a field has been set.
func (o *WorkflowFileOperations) HasFileDownload() bool {
	if o != nil && o.FileDownload.IsSet() {
		return true
	}

	return false
}

// SetFileDownload gets a reference to the given NullableWorkflowFileDownloadOp and assigns it to the FileDownload field.
func (o *WorkflowFileOperations) SetFileDownload(v WorkflowFileDownloadOp) {
	o.FileDownload.Set(&v)
}

// SetFileDownloadNil sets the value for FileDownload to be an explicit nil
func (o *WorkflowFileOperations) SetFileDownloadNil() {
	o.FileDownload.Set(nil)
}

// UnsetFileDownload ensures that no value is present for FileDownload, not even an explicit nil
func (o *WorkflowFileOperations) UnsetFileDownload() {
	o.FileDownload.Unset()
}

// GetFileTemplate returns the FileTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowFileOperations) GetFileTemplate() WorkflowFileTemplateOp {
	if o == nil || o.FileTemplate.Get() == nil {
		var ret WorkflowFileTemplateOp
		return ret
	}
	return *o.FileTemplate.Get()
}

// GetFileTemplateOk returns a tuple with the FileTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowFileOperations) GetFileTemplateOk() (*WorkflowFileTemplateOp, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileTemplate.Get(), o.FileTemplate.IsSet()
}

// HasFileTemplate returns a boolean if a field has been set.
func (o *WorkflowFileOperations) HasFileTemplate() bool {
	if o != nil && o.FileTemplate.IsSet() {
		return true
	}

	return false
}

// SetFileTemplate gets a reference to the given NullableWorkflowFileTemplateOp and assigns it to the FileTemplate field.
func (o *WorkflowFileOperations) SetFileTemplate(v WorkflowFileTemplateOp) {
	o.FileTemplate.Set(&v)
}

// SetFileTemplateNil sets the value for FileTemplate to be an explicit nil
func (o *WorkflowFileOperations) SetFileTemplateNil() {
	o.FileTemplate.Set(nil)
}

// UnsetFileTemplate ensures that no value is present for FileTemplate, not even an explicit nil
func (o *WorkflowFileOperations) UnsetFileTemplate() {
	o.FileTemplate.Unset()
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *WorkflowFileOperations) GetOperationType() string {
	if o == nil || o.OperationType == nil {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileOperations) GetOperationTypeOk() (*string, bool) {
	if o == nil || o.OperationType == nil {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *WorkflowFileOperations) HasOperationType() bool {
	if o != nil && o.OperationType != nil {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *WorkflowFileOperations) SetOperationType(v string) {
	o.OperationType = &v
}

func (o WorkflowFileOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileDownload.IsSet() {
		toSerialize["FileDownload"] = o.FileDownload.Get()
	}
	if o.FileTemplate.IsSet() {
		toSerialize["FileTemplate"] = o.FileTemplate.Get()
	}
	if o.OperationType != nil {
		toSerialize["OperationType"] = o.OperationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowFileOperations) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowFileOperationsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                         `json:"ObjectType"`
		FileDownload NullableWorkflowFileDownloadOp `json:"FileDownload,omitempty"`
		FileTemplate NullableWorkflowFileTemplateOp `json:"FileTemplate,omitempty"`
		// File operation type to be executed on the connected endpoint. * `FileDownload` - The API is executed in a remote device connected to the Intersightthrough its device connector. This operation is to download the filefrom specified storage bucket to the specific path on the device. * `FileTemplatize` - Populates data driven template file with input values to generate textual output.Inputs - the path of the template file on the device and json values to populate.An error will be returned if the file does not exists or if there is an error whileexecuting the template.
		OperationType *string `json:"OperationType,omitempty"`
	}

	varWorkflowFileOperationsWithoutEmbeddedStruct := WorkflowFileOperationsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowFileOperationsWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowFileOperations := _WorkflowFileOperations{}
		varWorkflowFileOperations.ClassId = varWorkflowFileOperationsWithoutEmbeddedStruct.ClassId
		varWorkflowFileOperations.ObjectType = varWorkflowFileOperationsWithoutEmbeddedStruct.ObjectType
		varWorkflowFileOperations.FileDownload = varWorkflowFileOperationsWithoutEmbeddedStruct.FileDownload
		varWorkflowFileOperations.FileTemplate = varWorkflowFileOperationsWithoutEmbeddedStruct.FileTemplate
		varWorkflowFileOperations.OperationType = varWorkflowFileOperationsWithoutEmbeddedStruct.OperationType
		*o = WorkflowFileOperations(varWorkflowFileOperations)
	} else {
		return err
	}

	varWorkflowFileOperations := _WorkflowFileOperations{}

	err = json.Unmarshal(bytes, &varWorkflowFileOperations)
	if err == nil {
		o.WorkflowApi = varWorkflowFileOperations.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileDownload")
		delete(additionalProperties, "FileTemplate")
		delete(additionalProperties, "OperationType")

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowFileOperations struct {
	value *WorkflowFileOperations
	isSet bool
}

func (v NullableWorkflowFileOperations) Get() *WorkflowFileOperations {
	return v.value
}

func (v *NullableWorkflowFileOperations) Set(val *WorkflowFileOperations) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFileOperations) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFileOperations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFileOperations(val *WorkflowFileOperations) *NullableWorkflowFileOperations {
	return &NullableWorkflowFileOperations{value: val, isSet: true}
}

func (v NullableWorkflowFileOperations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFileOperations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
