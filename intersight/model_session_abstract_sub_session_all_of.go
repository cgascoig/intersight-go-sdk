/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SessionAbstractSubSessionAllOf Definition of the list of properties defined in 'session.AbstractSubSession', excluding properties defined in parent classes.
type SessionAbstractSubSessionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of target on which session is initiated.
	TargetName           *string                             `json:"TargetName,omitempty"`
	Session              *SessionAbstractSessionRelationship `json:"Session,omitempty"`
	Target               *MoBaseMoRelationship               `json:"Target,omitempty"`
	User                 *IamUserRelationship                `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SessionAbstractSubSessionAllOf SessionAbstractSubSessionAllOf

// NewSessionAbstractSubSessionAllOf instantiates a new SessionAbstractSubSessionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionAbstractSubSessionAllOf(classId string, objectType string) *SessionAbstractSubSessionAllOf {
	this := SessionAbstractSubSessionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSessionAbstractSubSessionAllOfWithDefaults instantiates a new SessionAbstractSubSessionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionAbstractSubSessionAllOfWithDefaults() *SessionAbstractSubSessionAllOf {
	this := SessionAbstractSubSessionAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *SessionAbstractSubSessionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SessionAbstractSubSessionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SessionAbstractSubSessionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SessionAbstractSubSessionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *SessionAbstractSubSessionAllOf) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *SessionAbstractSubSessionAllOf) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *SessionAbstractSubSessionAllOf) SetTargetName(v string) {
	o.TargetName = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *SessionAbstractSubSessionAllOf) GetSession() SessionAbstractSessionRelationship {
	if o == nil || o.Session == nil {
		var ret SessionAbstractSessionRelationship
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetSessionOk() (*SessionAbstractSessionRelationship, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *SessionAbstractSubSessionAllOf) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given SessionAbstractSessionRelationship and assigns it to the Session field.
func (o *SessionAbstractSubSessionAllOf) SetSession(v SessionAbstractSessionRelationship) {
	o.Session = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SessionAbstractSubSessionAllOf) GetTarget() MoBaseMoRelationship {
	if o == nil || o.Target == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetTargetOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SessionAbstractSubSessionAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given MoBaseMoRelationship and assigns it to the Target field.
func (o *SessionAbstractSubSessionAllOf) SetTarget(v MoBaseMoRelationship) {
	o.Target = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SessionAbstractSubSessionAllOf) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionAbstractSubSessionAllOf) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SessionAbstractSubSessionAllOf) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *SessionAbstractSubSessionAllOf) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o SessionAbstractSubSessionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.Session != nil {
		toSerialize["Session"] = o.Session
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SessionAbstractSubSessionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSessionAbstractSubSessionAllOf := _SessionAbstractSubSessionAllOf{}

	if err = json.Unmarshal(bytes, &varSessionAbstractSubSessionAllOf); err == nil {
		*o = SessionAbstractSubSessionAllOf(varSessionAbstractSubSessionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "Session")
		delete(additionalProperties, "Target")
		delete(additionalProperties, "User")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSessionAbstractSubSessionAllOf struct {
	value *SessionAbstractSubSessionAllOf
	isSet bool
}

func (v NullableSessionAbstractSubSessionAllOf) Get() *SessionAbstractSubSessionAllOf {
	return v.value
}

func (v *NullableSessionAbstractSubSessionAllOf) Set(val *SessionAbstractSubSessionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionAbstractSubSessionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionAbstractSubSessionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionAbstractSubSessionAllOf(val *SessionAbstractSubSessionAllOf) *NullableSessionAbstractSubSessionAllOf {
	return &NullableSessionAbstractSubSessionAllOf{value: val, isSet: true}
}

func (v NullableSessionAbstractSubSessionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionAbstractSubSessionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
