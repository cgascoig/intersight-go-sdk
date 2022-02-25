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

// ConnectorStreamKeepalive Keepalive message sent between service and device over lifetime of stream. Cloud services will configure device to emit message periodically. On receipt cloud service will send a response keepalive to the device. If no input/output or keepalive message is received on either side within a configured timeout the stream will be closed.
type ConnectorStreamKeepalive struct {
	ConnectorStreamMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Stream has been paused and is not actively sending outputs to cloud. Device will pause streams when it does not receive an acknowledgment of messages after sending number messages cloud has configured in the stream acknowledge rate (set in StartStream.ackRate field on stream start). If cloud receives a keepalive messages with paused set it will attempt to acknowledge the latest messages processed from the device in order to unpause the stream.
	Paused               *bool `json:"Paused,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorStreamKeepalive ConnectorStreamKeepalive

// NewConnectorStreamKeepalive instantiates a new ConnectorStreamKeepalive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorStreamKeepalive(classId string, objectType string) *ConnectorStreamKeepalive {
	this := ConnectorStreamKeepalive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorStreamKeepaliveWithDefaults instantiates a new ConnectorStreamKeepalive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorStreamKeepaliveWithDefaults() *ConnectorStreamKeepalive {
	this := ConnectorStreamKeepalive{}
	var classId string = "connector.StreamKeepalive"
	this.ClassId = classId
	var objectType string = "connector.StreamKeepalive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorStreamKeepalive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamKeepalive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorStreamKeepalive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorStreamKeepalive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorStreamKeepalive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorStreamKeepalive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *ConnectorStreamKeepalive) GetPaused() bool {
	if o == nil || o.Paused == nil {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorStreamKeepalive) GetPausedOk() (*bool, bool) {
	if o == nil || o.Paused == nil {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *ConnectorStreamKeepalive) HasPaused() bool {
	if o != nil && o.Paused != nil {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *ConnectorStreamKeepalive) SetPaused(v bool) {
	o.Paused = &v
}

func (o ConnectorStreamKeepalive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorStreamMessage, errConnectorStreamMessage := json.Marshal(o.ConnectorStreamMessage)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	errConnectorStreamMessage = json.Unmarshal([]byte(serializedConnectorStreamMessage), &toSerialize)
	if errConnectorStreamMessage != nil {
		return []byte{}, errConnectorStreamMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Paused != nil {
		toSerialize["Paused"] = o.Paused
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorStreamKeepalive) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorStreamKeepaliveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Stream has been paused and is not actively sending outputs to cloud. Device will pause streams when it does not receive an acknowledgment of messages after sending number messages cloud has configured in the stream acknowledge rate (set in StartStream.ackRate field on stream start). If cloud receives a keepalive messages with paused set it will attempt to acknowledge the latest messages processed from the device in order to unpause the stream.
		Paused *bool `json:"Paused,omitempty"`
	}

	varConnectorStreamKeepaliveWithoutEmbeddedStruct := ConnectorStreamKeepaliveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorStreamKeepaliveWithoutEmbeddedStruct)
	if err == nil {
		varConnectorStreamKeepalive := _ConnectorStreamKeepalive{}
		varConnectorStreamKeepalive.ClassId = varConnectorStreamKeepaliveWithoutEmbeddedStruct.ClassId
		varConnectorStreamKeepalive.ObjectType = varConnectorStreamKeepaliveWithoutEmbeddedStruct.ObjectType
		varConnectorStreamKeepalive.Paused = varConnectorStreamKeepaliveWithoutEmbeddedStruct.Paused
		*o = ConnectorStreamKeepalive(varConnectorStreamKeepalive)
	} else {
		return err
	}

	varConnectorStreamKeepalive := _ConnectorStreamKeepalive{}

	err = json.Unmarshal(bytes, &varConnectorStreamKeepalive)
	if err == nil {
		o.ConnectorStreamMessage = varConnectorStreamKeepalive.ConnectorStreamMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Paused")

		// remove fields from embedded structs
		reflectConnectorStreamMessage := reflect.ValueOf(o.ConnectorStreamMessage)
		for i := 0; i < reflectConnectorStreamMessage.Type().NumField(); i++ {
			t := reflectConnectorStreamMessage.Type().Field(i)

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

type NullableConnectorStreamKeepalive struct {
	value *ConnectorStreamKeepalive
	isSet bool
}

func (v NullableConnectorStreamKeepalive) Get() *ConnectorStreamKeepalive {
	return v.value
}

func (v *NullableConnectorStreamKeepalive) Set(val *ConnectorStreamKeepalive) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorStreamKeepalive) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorStreamKeepalive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorStreamKeepalive(val *ConnectorStreamKeepalive) *NullableConnectorStreamKeepalive {
	return &NullableConnectorStreamKeepalive{value: val, isSet: true}
}

func (v NullableConnectorStreamKeepalive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorStreamKeepalive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
