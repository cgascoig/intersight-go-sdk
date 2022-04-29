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
	"time"
)

// TerminalAuditLog Audit log of remote terminal user sessions.
type TerminalAuditLog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time the terminal was closed. If terminal has not closed, value is zero time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// The time the terminal session was opened.
	StartTime            *time.Time                         `json:"StartTime,omitempty"`
	DeviceRegistration   *AssetDeviceConnectionRelationship `json:"DeviceRegistration,omitempty"`
	User                 *IamUserRelationship               `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TerminalAuditLog TerminalAuditLog

// NewTerminalAuditLog instantiates a new TerminalAuditLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalAuditLog(classId string, objectType string) *TerminalAuditLog {
	this := TerminalAuditLog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTerminalAuditLogWithDefaults instantiates a new TerminalAuditLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalAuditLogWithDefaults() *TerminalAuditLog {
	this := TerminalAuditLog{}
	var classId string = "terminal.AuditLog"
	this.ClassId = classId
	var objectType string = "terminal.AuditLog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TerminalAuditLog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TerminalAuditLog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TerminalAuditLog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TerminalAuditLog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TerminalAuditLog) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TerminalAuditLog) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TerminalAuditLog) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TerminalAuditLog) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TerminalAuditLog) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TerminalAuditLog) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *TerminalAuditLog) GetDeviceRegistration() AssetDeviceConnectionRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceConnectionRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetDeviceRegistrationOk() (*AssetDeviceConnectionRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *TerminalAuditLog) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceConnectionRelationship and assigns it to the DeviceRegistration field.
func (o *TerminalAuditLog) SetDeviceRegistration(v AssetDeviceConnectionRelationship) {
	o.DeviceRegistration = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TerminalAuditLog) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalAuditLog) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TerminalAuditLog) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *TerminalAuditLog) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o TerminalAuditLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TerminalAuditLog) UnmarshalJSON(bytes []byte) (err error) {
	type TerminalAuditLogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time the terminal was closed. If terminal has not closed, value is zero time.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// The time the terminal session was opened.
		StartTime          *time.Time                         `json:"StartTime,omitempty"`
		DeviceRegistration *AssetDeviceConnectionRelationship `json:"DeviceRegistration,omitempty"`
		User               *IamUserRelationship               `json:"User,omitempty"`
	}

	varTerminalAuditLogWithoutEmbeddedStruct := TerminalAuditLogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTerminalAuditLogWithoutEmbeddedStruct)
	if err == nil {
		varTerminalAuditLog := _TerminalAuditLog{}
		varTerminalAuditLog.ClassId = varTerminalAuditLogWithoutEmbeddedStruct.ClassId
		varTerminalAuditLog.ObjectType = varTerminalAuditLogWithoutEmbeddedStruct.ObjectType
		varTerminalAuditLog.EndTime = varTerminalAuditLogWithoutEmbeddedStruct.EndTime
		varTerminalAuditLog.StartTime = varTerminalAuditLogWithoutEmbeddedStruct.StartTime
		varTerminalAuditLog.DeviceRegistration = varTerminalAuditLogWithoutEmbeddedStruct.DeviceRegistration
		varTerminalAuditLog.User = varTerminalAuditLogWithoutEmbeddedStruct.User
		*o = TerminalAuditLog(varTerminalAuditLog)
	} else {
		return err
	}

	varTerminalAuditLog := _TerminalAuditLog{}

	err = json.Unmarshal(bytes, &varTerminalAuditLog)
	if err == nil {
		o.MoBaseMo = varTerminalAuditLog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "User")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableTerminalAuditLog struct {
	value *TerminalAuditLog
	isSet bool
}

func (v NullableTerminalAuditLog) Get() *TerminalAuditLog {
	return v.value
}

func (v *NullableTerminalAuditLog) Set(val *TerminalAuditLog) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalAuditLog) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalAuditLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalAuditLog(val *TerminalAuditLog) *NullableTerminalAuditLog {
	return &NullableTerminalAuditLog{value: val, isSet: true}
}

func (v NullableTerminalAuditLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalAuditLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
