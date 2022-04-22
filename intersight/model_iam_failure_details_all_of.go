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
	"time"
)

// IamFailureDetailsAllOf Definition of the list of properties defined in 'iam.FailureDetails', excluding properties defined in parent classes.
type IamFailureDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reason for the failure during verification.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Timestamp of the failure during verification.
	FailureTime          *time.Time `json:"FailureTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamFailureDetailsAllOf IamFailureDetailsAllOf

// NewIamFailureDetailsAllOf instantiates a new IamFailureDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamFailureDetailsAllOf(classId string, objectType string) *IamFailureDetailsAllOf {
	this := IamFailureDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamFailureDetailsAllOfWithDefaults instantiates a new IamFailureDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamFailureDetailsAllOfWithDefaults() *IamFailureDetailsAllOf {
	this := IamFailureDetailsAllOf{}
	var classId string = "iam.FailureDetails"
	this.ClassId = classId
	var objectType string = "iam.FailureDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamFailureDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamFailureDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamFailureDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamFailureDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamFailureDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamFailureDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *IamFailureDetailsAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamFailureDetailsAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *IamFailureDetailsAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *IamFailureDetailsAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetFailureTime returns the FailureTime field value if set, zero value otherwise.
func (o *IamFailureDetailsAllOf) GetFailureTime() time.Time {
	if o == nil || o.FailureTime == nil {
		var ret time.Time
		return ret
	}
	return *o.FailureTime
}

// GetFailureTimeOk returns a tuple with the FailureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamFailureDetailsAllOf) GetFailureTimeOk() (*time.Time, bool) {
	if o == nil || o.FailureTime == nil {
		return nil, false
	}
	return o.FailureTime, true
}

// HasFailureTime returns a boolean if a field has been set.
func (o *IamFailureDetailsAllOf) HasFailureTime() bool {
	if o != nil && o.FailureTime != nil {
		return true
	}

	return false
}

// SetFailureTime gets a reference to the given time.Time and assigns it to the FailureTime field.
func (o *IamFailureDetailsAllOf) SetFailureTime(v time.Time) {
	o.FailureTime = &v
}

func (o IamFailureDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.FailureTime != nil {
		toSerialize["FailureTime"] = o.FailureTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamFailureDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamFailureDetailsAllOf := _IamFailureDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varIamFailureDetailsAllOf); err == nil {
		*o = IamFailureDetailsAllOf(varIamFailureDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "FailureTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamFailureDetailsAllOf struct {
	value *IamFailureDetailsAllOf
	isSet bool
}

func (v NullableIamFailureDetailsAllOf) Get() *IamFailureDetailsAllOf {
	return v.value
}

func (v *NullableIamFailureDetailsAllOf) Set(val *IamFailureDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamFailureDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamFailureDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamFailureDetailsAllOf(val *IamFailureDetailsAllOf) *NullableIamFailureDetailsAllOf {
	return &NullableIamFailureDetailsAllOf{value: val, isSet: true}
}

func (v NullableIamFailureDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamFailureDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
