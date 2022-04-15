/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
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

// SessionAbstractSession A base abstract class for all sessions.
type SessionAbstractSession struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The user agent IP address from which the session is launched.
	ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
	// The time at which the session ended.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Role of the user who launched the session.
	Role *string `json:"Role,omitempty"`
	// The status of the session. * `Active` - The session is currently active. * `Ended` - The session has ended normally. * `Terminated` - The session was terminated by an admin.
	Status *string `json:"Status,omitempty"`
	// User ID or E-mail Address of the user who launched the session.
	UserIdOrEmail        *string `json:"UserIdOrEmail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionAbstractSession SessionAbstractSession

// NewSessionAbstractSession instantiates a new SessionAbstractSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionAbstractSession(classId string, objectType string) *SessionAbstractSession {
	this := SessionAbstractSession{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewSessionAbstractSessionWithDefaults instantiates a new SessionAbstractSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionAbstractSessionWithDefaults() *SessionAbstractSession {
	this := SessionAbstractSession{}
	var status string = "Active"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *SessionAbstractSession) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SessionAbstractSession) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SessionAbstractSession) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SessionAbstractSession) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientIpAddress returns the ClientIpAddress field value if set, zero value otherwise.
func (o *SessionAbstractSession) GetClientIpAddress() string {
	if o == nil || o.ClientIpAddress == nil {
		var ret string
		return ret
	}
	return *o.ClientIpAddress
}

// GetClientIpAddressOk returns a tuple with the ClientIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetClientIpAddressOk() (*string, bool) {
	if o == nil || o.ClientIpAddress == nil {
		return nil, false
	}
	return o.ClientIpAddress, true
}

// HasClientIpAddress returns a boolean if a field has been set.
func (o *SessionAbstractSession) HasClientIpAddress() bool {
	if o != nil && o.ClientIpAddress != nil {
		return true
	}

	return false
}

// SetClientIpAddress gets a reference to the given string and assigns it to the ClientIpAddress field.
func (o *SessionAbstractSession) SetClientIpAddress(v string) {
	o.ClientIpAddress = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SessionAbstractSession) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SessionAbstractSession) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SessionAbstractSession) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SessionAbstractSession) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SessionAbstractSession) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SessionAbstractSession) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SessionAbstractSession) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SessionAbstractSession) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SessionAbstractSession) SetStatus(v string) {
	o.Status = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *SessionAbstractSession) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSession) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *SessionAbstractSession) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *SessionAbstractSession) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

func (o SessionAbstractSession) MarshalJSON() ([]byte, error) {
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
	if o.ClientIpAddress != nil {
		toSerialize["ClientIpAddress"] = o.ClientIpAddress
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionAbstractSession) UnmarshalJSON(bytes []byte) (err error) {
	type SessionAbstractSessionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The user agent IP address from which the session is launched.
		ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
		// The time at which the session ended.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Role of the user who launched the session.
		Role *string `json:"Role,omitempty"`
		// The status of the session. * `Active` - The session is currently active. * `Ended` - The session has ended normally. * `Terminated` - The session was terminated by an admin.
		Status *string `json:"Status,omitempty"`
		// User ID or E-mail Address of the user who launched the session.
		UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
	}

	varSessionAbstractSessionWithoutEmbeddedStruct := SessionAbstractSessionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSessionAbstractSessionWithoutEmbeddedStruct)
	if err == nil {
		varSessionAbstractSession := _SessionAbstractSession{}
		varSessionAbstractSession.ClassId = varSessionAbstractSessionWithoutEmbeddedStruct.ClassId
		varSessionAbstractSession.ObjectType = varSessionAbstractSessionWithoutEmbeddedStruct.ObjectType
		varSessionAbstractSession.ClientIpAddress = varSessionAbstractSessionWithoutEmbeddedStruct.ClientIpAddress
		varSessionAbstractSession.EndTime = varSessionAbstractSessionWithoutEmbeddedStruct.EndTime
		varSessionAbstractSession.Role = varSessionAbstractSessionWithoutEmbeddedStruct.Role
		varSessionAbstractSession.Status = varSessionAbstractSessionWithoutEmbeddedStruct.Status
		varSessionAbstractSession.UserIdOrEmail = varSessionAbstractSessionWithoutEmbeddedStruct.UserIdOrEmail
		*o = SessionAbstractSession(varSessionAbstractSession)
	} else {
		return err
	}

	varSessionAbstractSession := _SessionAbstractSession{}

	err = json.Unmarshal(bytes, &varSessionAbstractSession)
	if err == nil {
		o.MoBaseMo = varSessionAbstractSession.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientIpAddress")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UserIdOrEmail")

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

type NullableSessionAbstractSession struct {
	value *SessionAbstractSession
	isSet bool
}

func (v NullableSessionAbstractSession) Get() *SessionAbstractSession {
	return v.value
}

func (v *NullableSessionAbstractSession) Set(val *SessionAbstractSession) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionAbstractSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionAbstractSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionAbstractSession(val *SessionAbstractSession) *NullableSessionAbstractSession {
	return &NullableSessionAbstractSession{value: val, isSet: true}
}

func (v NullableSessionAbstractSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionAbstractSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
