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
	"reflect"
	"strings"
)

// WorkflowFileTransfer Message to transfer a file from Intersight connected device to remote server.
type WorkflowFileTransfer struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Destination file path on the target server.
	DestinationFilePath *string `json:"DestinationFilePath,omitempty"`
	// File permission to set on the transferred file.
	FileMode *int64 `json:"FileMode,omitempty"`
	// Source file path on the Intersight connected device.
	SourceFilePath       *string `json:"SourceFilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowFileTransfer WorkflowFileTransfer

// NewWorkflowFileTransfer instantiates a new WorkflowFileTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFileTransfer(classId string, objectType string) *WorkflowFileTransfer {
	this := WorkflowFileTransfer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowFileTransferWithDefaults instantiates a new WorkflowFileTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFileTransferWithDefaults() *WorkflowFileTransfer {
	this := WorkflowFileTransfer{}
	var classId string = "workflow.FileTransfer"
	this.ClassId = classId
	var objectType string = "workflow.FileTransfer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowFileTransfer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowFileTransfer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowFileTransfer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowFileTransfer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDestinationFilePath returns the DestinationFilePath field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetDestinationFilePath() string {
	if o == nil || o.DestinationFilePath == nil {
		var ret string
		return ret
	}
	return *o.DestinationFilePath
}

// GetDestinationFilePathOk returns a tuple with the DestinationFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetDestinationFilePathOk() (*string, bool) {
	if o == nil || o.DestinationFilePath == nil {
		return nil, false
	}
	return o.DestinationFilePath, true
}

// HasDestinationFilePath returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasDestinationFilePath() bool {
	if o != nil && o.DestinationFilePath != nil {
		return true
	}

	return false
}

// SetDestinationFilePath gets a reference to the given string and assigns it to the DestinationFilePath field.
func (o *WorkflowFileTransfer) SetDestinationFilePath(v string) {
	o.DestinationFilePath = &v
}

// GetFileMode returns the FileMode field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetFileMode() int64 {
	if o == nil || o.FileMode == nil {
		var ret int64
		return ret
	}
	return *o.FileMode
}

// GetFileModeOk returns a tuple with the FileMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetFileModeOk() (*int64, bool) {
	if o == nil || o.FileMode == nil {
		return nil, false
	}
	return o.FileMode, true
}

// HasFileMode returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasFileMode() bool {
	if o != nil && o.FileMode != nil {
		return true
	}

	return false
}

// SetFileMode gets a reference to the given int64 and assigns it to the FileMode field.
func (o *WorkflowFileTransfer) SetFileMode(v int64) {
	o.FileMode = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *WorkflowFileTransfer) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowFileTransfer) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *WorkflowFileTransfer) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *WorkflowFileTransfer) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

func (o WorkflowFileTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DestinationFilePath != nil {
		toSerialize["DestinationFilePath"] = o.DestinationFilePath
	}
	if o.FileMode != nil {
		toSerialize["FileMode"] = o.FileMode
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowFileTransfer) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowFileTransferWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Destination file path on the target server.
		DestinationFilePath *string `json:"DestinationFilePath,omitempty"`
		// File permission to set on the transferred file.
		FileMode *int64 `json:"FileMode,omitempty"`
		// Source file path on the Intersight connected device.
		SourceFilePath *string `json:"SourceFilePath,omitempty"`
	}

	varWorkflowFileTransferWithoutEmbeddedStruct := WorkflowFileTransferWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowFileTransferWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowFileTransfer := _WorkflowFileTransfer{}
		varWorkflowFileTransfer.ClassId = varWorkflowFileTransferWithoutEmbeddedStruct.ClassId
		varWorkflowFileTransfer.ObjectType = varWorkflowFileTransferWithoutEmbeddedStruct.ObjectType
		varWorkflowFileTransfer.DestinationFilePath = varWorkflowFileTransferWithoutEmbeddedStruct.DestinationFilePath
		varWorkflowFileTransfer.FileMode = varWorkflowFileTransferWithoutEmbeddedStruct.FileMode
		varWorkflowFileTransfer.SourceFilePath = varWorkflowFileTransferWithoutEmbeddedStruct.SourceFilePath
		*o = WorkflowFileTransfer(varWorkflowFileTransfer)
	} else {
		return err
	}

	varWorkflowFileTransfer := _WorkflowFileTransfer{}

	err = json.Unmarshal(bytes, &varWorkflowFileTransfer)
	if err == nil {
		o.ConnectorBaseMessage = varWorkflowFileTransfer.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestinationFilePath")
		delete(additionalProperties, "FileMode")
		delete(additionalProperties, "SourceFilePath")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableWorkflowFileTransfer struct {
	value *WorkflowFileTransfer
	isSet bool
}

func (v NullableWorkflowFileTransfer) Get() *WorkflowFileTransfer {
	return v.value
}

func (v *NullableWorkflowFileTransfer) Set(val *WorkflowFileTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFileTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFileTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFileTransfer(val *WorkflowFileTransfer) *NullableWorkflowFileTransfer {
	return &NullableWorkflowFileTransfer{value: val, isSet: true}
}

func (v NullableWorkflowFileTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFileTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
