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
	"reflect"
	"strings"
)

// ConnectorFileMessage Message carries file operations to perform on the platforms file system. Cloud services can send message to open and write to files on the connector platforms file system. Writes to a file can be buffered across many 'FileContent' messages, the file plugin will append to an open file as it receives file content until a close message is received. If any operation fails (such as a file write returns error) an error will be returned to the cloud service and a best effort to close and remove the file will be made (if the file was previously opened).
type ConnectorFileMessage struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Message type carrying the file operation to perform. * `OpenFile` - Open a file for write to on the platforms file system. Must be the first message sent to the plugin when writing files, attempting to write to or close a file that has not been opened will return error. If file does not exist file will be created, if it does exist file will be truncated. If the files parent directory does not exist an error will be returned. * `FileContent` - Carries content to write to an open file. The complete file content can be buffered across many FileContent messages, with each subsequent message received appending to the end of the file. If the total size of FileContent messages exceeds the file size limit of 5MB the plugin will return error and remove the opened file. * `CloseFile` - Close the file. Close message must be sent at the end of file transfer. If no close is sent and no input is received after a timeout of 10 minutes the file will be removed. After a CloseFile is received no further FileContent messages will be accepted and will return error. * `DeleteFile` - Delete a file from the filesystem. If the file does not exist operation is a no-op and no error is returned. File must be within one of the platforms allowed writable directories. * `ListDir` - Open a directory and return a list of files in the directory. Does not recursively walk subdirectories, returns only the top level files in the directory. Output will be a list of json encoded FileInfos. * `` - Open a file and read the contents into a response message. The response message body is the entire contents of the file with no encoding or compression. If the file does not exist an error will be returned. If the size of the file exceeds 500KB an error will be returned without reading any file contents. * `PopulateRemoteTemplateFile` - Populates data driven template file with input values to generate textual output. Inputs are - the path of the template file on the device and values to populate. An error will be returned if the file does not exist or if there is an error while executing the templates. * `GetFreeSpace` - Check the disk space availability of folder and returns the free disk space in percentage. Output will be a list of encoded FileInfos.
	MsgType *string `json:"MsgType,omitempty"`
	// The absolute path of the file to open on the platforms file system. Must be a sub-directory of a directory defined within the platform configurations WriteableDirectories. The file system device to write to must also have sufficient free space to write to (<75% full). Must be set for each message that is sent.
	Path *string `json:"Path,omitempty"`
	// The stream of bytes to write to file when message type is FileContent. Ignored for OpenFile and CloseFile messages. This field holds the input values (key-value pairs) to be populated in the template file when message type is PopulateTemplateFile.
	Stream               *string `json:"Stream,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorFileMessage ConnectorFileMessage

// NewConnectorFileMessage instantiates a new ConnectorFileMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorFileMessage(classId string, objectType string) *ConnectorFileMessage {
	this := ConnectorFileMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	var msgType string = "OpenFile"
	this.MsgType = &msgType
	return &this
}

// NewConnectorFileMessageWithDefaults instantiates a new ConnectorFileMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorFileMessageWithDefaults() *ConnectorFileMessage {
	this := ConnectorFileMessage{}
	var classId string = "connector.FileMessage"
	this.ClassId = classId
	var objectType string = "connector.FileMessage"
	this.ObjectType = objectType
	var msgType string = "OpenFile"
	this.MsgType = &msgType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorFileMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorFileMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorFileMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorFileMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorFileMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorFileMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *ConnectorFileMessage) GetMsgType() string {
	if o == nil || o.MsgType == nil {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileMessage) GetMsgTypeOk() (*string, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *ConnectorFileMessage) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *ConnectorFileMessage) SetMsgType(v string) {
	o.MsgType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ConnectorFileMessage) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileMessage) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ConnectorFileMessage) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ConnectorFileMessage) SetPath(v string) {
	o.Path = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ConnectorFileMessage) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorFileMessage) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ConnectorFileMessage) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *ConnectorFileMessage) SetStream(v string) {
	o.Stream = &v
}

func (o ConnectorFileMessage) MarshalJSON() ([]byte, error) {
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
	if o.MsgType != nil {
		toSerialize["MsgType"] = o.MsgType
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Stream != nil {
		toSerialize["Stream"] = o.Stream
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorFileMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorFileMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Message type carrying the file operation to perform. * `OpenFile` - Open a file for write to on the platforms file system. Must be the first message sent to the plugin when writing files, attempting to write to or close a file that has not been opened will return error. If file does not exist file will be created, if it does exist file will be truncated. If the files parent directory does not exist an error will be returned. * `FileContent` - Carries content to write to an open file. The complete file content can be buffered across many FileContent messages, with each subsequent message received appending to the end of the file. If the total size of FileContent messages exceeds the file size limit of 5MB the plugin will return error and remove the opened file. * `CloseFile` - Close the file. Close message must be sent at the end of file transfer. If no close is sent and no input is received after a timeout of 10 minutes the file will be removed. After a CloseFile is received no further FileContent messages will be accepted and will return error. * `DeleteFile` - Delete a file from the filesystem. If the file does not exist operation is a no-op and no error is returned. File must be within one of the platforms allowed writable directories. * `ListDir` - Open a directory and return a list of files in the directory. Does not recursively walk subdirectories, returns only the top level files in the directory. Output will be a list of json encoded FileInfos. * `` - Open a file and read the contents into a response message. The response message body is the entire contents of the file with no encoding or compression. If the file does not exist an error will be returned. If the size of the file exceeds 500KB an error will be returned without reading any file contents. * `PopulateRemoteTemplateFile` - Populates data driven template file with input values to generate textual output. Inputs are - the path of the template file on the device and values to populate. An error will be returned if the file does not exist or if there is an error while executing the templates. * `GetFreeSpace` - Check the disk space availability of folder and returns the free disk space in percentage. Output will be a list of encoded FileInfos.
		MsgType *string `json:"MsgType,omitempty"`
		// The absolute path of the file to open on the platforms file system. Must be a sub-directory of a directory defined within the platform configurations WriteableDirectories. The file system device to write to must also have sufficient free space to write to (<75% full). Must be set for each message that is sent.
		Path *string `json:"Path,omitempty"`
		// The stream of bytes to write to file when message type is FileContent. Ignored for OpenFile and CloseFile messages. This field holds the input values (key-value pairs) to be populated in the template file when message type is PopulateTemplateFile.
		Stream *string `json:"Stream,omitempty"`
	}

	varConnectorFileMessageWithoutEmbeddedStruct := ConnectorFileMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorFileMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorFileMessage := _ConnectorFileMessage{}
		varConnectorFileMessage.ClassId = varConnectorFileMessageWithoutEmbeddedStruct.ClassId
		varConnectorFileMessage.ObjectType = varConnectorFileMessageWithoutEmbeddedStruct.ObjectType
		varConnectorFileMessage.MsgType = varConnectorFileMessageWithoutEmbeddedStruct.MsgType
		varConnectorFileMessage.Path = varConnectorFileMessageWithoutEmbeddedStruct.Path
		varConnectorFileMessage.Stream = varConnectorFileMessageWithoutEmbeddedStruct.Stream
		*o = ConnectorFileMessage(varConnectorFileMessage)
	} else {
		return err
	}

	varConnectorFileMessage := _ConnectorFileMessage{}

	err = json.Unmarshal(bytes, &varConnectorFileMessage)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorFileMessage.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MsgType")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Stream")

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

type NullableConnectorFileMessage struct {
	value *ConnectorFileMessage
	isSet bool
}

func (v NullableConnectorFileMessage) Get() *ConnectorFileMessage {
	return v.value
}

func (v *NullableConnectorFileMessage) Set(val *ConnectorFileMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorFileMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorFileMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorFileMessage(val *ConnectorFileMessage) *NullableConnectorFileMessage {
	return &NullableConnectorFileMessage{value: val, isSet: true}
}

func (v NullableConnectorFileMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorFileMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
