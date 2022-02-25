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

// ConnectorSshMessage An SSH message for opening, closing, executing a command, fetching or writing a file to a remote server. Cloud services send this message to a connectors SSH plugin to open and execute operations on an SSH session.
type ConnectorSshMessage struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                  `json:"ObjectType"`
	ExpectPrompts []ConnectorExpectPrompt `json:"ExpectPrompts,omitempty"`
	// The operation to execute on a new or existing session.
	MsgType *int64 `json:"MsgType,omitempty"`
	// Unique id of session to route messages to.
	SessionId *string `json:"SessionId,omitempty"`
	// The regex of the secure shell prompt.
	ShellPrompt *string `json:"ShellPrompt,omitempty"`
	// Input to the SSH operation to be executed. e.g. file contents to write.
	Stream *string `json:"Stream,omitempty"`
	// The timeout for the ssh command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorSshMessage ConnectorSshMessage

// NewConnectorSshMessage instantiates a new ConnectorSshMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorSshMessage(classId string, objectType string) *ConnectorSshMessage {
	this := ConnectorSshMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorSshMessageWithDefaults instantiates a new ConnectorSshMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorSshMessageWithDefaults() *ConnectorSshMessage {
	this := ConnectorSshMessage{}
	var classId string = "connector.SshMessage"
	this.ClassId = classId
	var objectType string = "connector.SshMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorSshMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorSshMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorSshMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorSshMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExpectPrompts returns the ExpectPrompts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorSshMessage) GetExpectPrompts() []ConnectorExpectPrompt {
	if o == nil {
		var ret []ConnectorExpectPrompt
		return ret
	}
	return o.ExpectPrompts
}

// GetExpectPromptsOk returns a tuple with the ExpectPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorSshMessage) GetExpectPromptsOk() (*[]ConnectorExpectPrompt, bool) {
	if o == nil || o.ExpectPrompts == nil {
		return nil, false
	}
	return &o.ExpectPrompts, true
}

// HasExpectPrompts returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasExpectPrompts() bool {
	if o != nil && o.ExpectPrompts != nil {
		return true
	}

	return false
}

// SetExpectPrompts gets a reference to the given []ConnectorExpectPrompt and assigns it to the ExpectPrompts field.
func (o *ConnectorSshMessage) SetExpectPrompts(v []ConnectorExpectPrompt) {
	o.ExpectPrompts = v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *ConnectorSshMessage) GetMsgType() int64 {
	if o == nil || o.MsgType == nil {
		var ret int64
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetMsgTypeOk() (*int64, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given int64 and assigns it to the MsgType field.
func (o *ConnectorSshMessage) SetMsgType(v int64) {
	o.MsgType = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ConnectorSshMessage) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ConnectorSshMessage) SetSessionId(v string) {
	o.SessionId = &v
}

// GetShellPrompt returns the ShellPrompt field value if set, zero value otherwise.
func (o *ConnectorSshMessage) GetShellPrompt() string {
	if o == nil || o.ShellPrompt == nil {
		var ret string
		return ret
	}
	return *o.ShellPrompt
}

// GetShellPromptOk returns a tuple with the ShellPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetShellPromptOk() (*string, bool) {
	if o == nil || o.ShellPrompt == nil {
		return nil, false
	}
	return o.ShellPrompt, true
}

// HasShellPrompt returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasShellPrompt() bool {
	if o != nil && o.ShellPrompt != nil {
		return true
	}

	return false
}

// SetShellPrompt gets a reference to the given string and assigns it to the ShellPrompt field.
func (o *ConnectorSshMessage) SetShellPrompt(v string) {
	o.ShellPrompt = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ConnectorSshMessage) GetStream() string {
	if o == nil || o.Stream == nil {
		var ret string
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetStreamOk() (*string, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasStream() bool {
	if o != nil && o.Stream != nil {
		return true
	}

	return false
}

// SetStream gets a reference to the given string and assigns it to the Stream field.
func (o *ConnectorSshMessage) SetStream(v string) {
	o.Stream = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectorSshMessage) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSshMessage) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectorSshMessage) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ConnectorSshMessage) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o ConnectorSshMessage) MarshalJSON() ([]byte, error) {
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
	if o.ExpectPrompts != nil {
		toSerialize["ExpectPrompts"] = o.ExpectPrompts
	}
	if o.MsgType != nil {
		toSerialize["MsgType"] = o.MsgType
	}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.ShellPrompt != nil {
		toSerialize["ShellPrompt"] = o.ShellPrompt
	}
	if o.Stream != nil {
		toSerialize["Stream"] = o.Stream
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorSshMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorSshMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                  `json:"ObjectType"`
		ExpectPrompts []ConnectorExpectPrompt `json:"ExpectPrompts,omitempty"`
		// The operation to execute on a new or existing session.
		MsgType *int64 `json:"MsgType,omitempty"`
		// Unique id of session to route messages to.
		SessionId *string `json:"SessionId,omitempty"`
		// The regex of the secure shell prompt.
		ShellPrompt *string `json:"ShellPrompt,omitempty"`
		// Input to the SSH operation to be executed. e.g. file contents to write.
		Stream *string `json:"Stream,omitempty"`
		// The timeout for the ssh command to complete and exit after starting or receiving input. If timeout is not set a default of 10 minutes will be used.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varConnectorSshMessageWithoutEmbeddedStruct := ConnectorSshMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorSshMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorSshMessage := _ConnectorSshMessage{}
		varConnectorSshMessage.ClassId = varConnectorSshMessageWithoutEmbeddedStruct.ClassId
		varConnectorSshMessage.ObjectType = varConnectorSshMessageWithoutEmbeddedStruct.ObjectType
		varConnectorSshMessage.ExpectPrompts = varConnectorSshMessageWithoutEmbeddedStruct.ExpectPrompts
		varConnectorSshMessage.MsgType = varConnectorSshMessageWithoutEmbeddedStruct.MsgType
		varConnectorSshMessage.SessionId = varConnectorSshMessageWithoutEmbeddedStruct.SessionId
		varConnectorSshMessage.ShellPrompt = varConnectorSshMessageWithoutEmbeddedStruct.ShellPrompt
		varConnectorSshMessage.Stream = varConnectorSshMessageWithoutEmbeddedStruct.Stream
		varConnectorSshMessage.Timeout = varConnectorSshMessageWithoutEmbeddedStruct.Timeout
		*o = ConnectorSshMessage(varConnectorSshMessage)
	} else {
		return err
	}

	varConnectorSshMessage := _ConnectorSshMessage{}

	err = json.Unmarshal(bytes, &varConnectorSshMessage)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorSshMessage.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExpectPrompts")
		delete(additionalProperties, "MsgType")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "ShellPrompt")
		delete(additionalProperties, "Stream")
		delete(additionalProperties, "Timeout")

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

type NullableConnectorSshMessage struct {
	value *ConnectorSshMessage
	isSet bool
}

func (v NullableConnectorSshMessage) Get() *ConnectorSshMessage {
	return v.value
}

func (v *NullableConnectorSshMessage) Set(val *ConnectorSshMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorSshMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorSshMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorSshMessage(val *ConnectorSshMessage) *NullableConnectorSshMessage {
	return &NullableConnectorSshMessage{value: val, isSet: true}
}

func (v NullableConnectorSshMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorSshMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
