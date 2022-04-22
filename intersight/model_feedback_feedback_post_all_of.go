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
)

// FeedbackFeedbackPostAllOf Definition of the list of properties defined in 'feedback.FeedbackPost', excluding properties defined in parent classes.
type FeedbackFeedbackPostAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	FeedbackData         NullableFeedbackFeedbackData `json:"FeedbackData,omitempty"`
	Account              *IamAccountRelationship      `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeedbackFeedbackPostAllOf FeedbackFeedbackPostAllOf

// NewFeedbackFeedbackPostAllOf instantiates a new FeedbackFeedbackPostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFeedbackPostAllOf(classId string, objectType string) *FeedbackFeedbackPostAllOf {
	this := FeedbackFeedbackPostAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFeedbackFeedbackPostAllOfWithDefaults instantiates a new FeedbackFeedbackPostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFeedbackPostAllOfWithDefaults() *FeedbackFeedbackPostAllOf {
	this := FeedbackFeedbackPostAllOf{}
	var classId string = "feedback.FeedbackPost"
	this.ClassId = classId
	var objectType string = "feedback.FeedbackPost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FeedbackFeedbackPostAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPostAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FeedbackFeedbackPostAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FeedbackFeedbackPostAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPostAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FeedbackFeedbackPostAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeedbackData returns the FeedbackData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackFeedbackPostAllOf) GetFeedbackData() FeedbackFeedbackData {
	if o == nil || o.FeedbackData.Get() == nil {
		var ret FeedbackFeedbackData
		return ret
	}
	return *o.FeedbackData.Get()
}

// GetFeedbackDataOk returns a tuple with the FeedbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackFeedbackPostAllOf) GetFeedbackDataOk() (*FeedbackFeedbackData, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeedbackData.Get(), o.FeedbackData.IsSet()
}

// HasFeedbackData returns a boolean if a field has been set.
func (o *FeedbackFeedbackPostAllOf) HasFeedbackData() bool {
	if o != nil && o.FeedbackData.IsSet() {
		return true
	}

	return false
}

// SetFeedbackData gets a reference to the given NullableFeedbackFeedbackData and assigns it to the FeedbackData field.
func (o *FeedbackFeedbackPostAllOf) SetFeedbackData(v FeedbackFeedbackData) {
	o.FeedbackData.Set(&v)
}

// SetFeedbackDataNil sets the value for FeedbackData to be an explicit nil
func (o *FeedbackFeedbackPostAllOf) SetFeedbackDataNil() {
	o.FeedbackData.Set(nil)
}

// UnsetFeedbackData ensures that no value is present for FeedbackData, not even an explicit nil
func (o *FeedbackFeedbackPostAllOf) UnsetFeedbackData() {
	o.FeedbackData.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *FeedbackFeedbackPostAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFeedbackPostAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *FeedbackFeedbackPostAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *FeedbackFeedbackPostAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o FeedbackFeedbackPostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FeedbackData.IsSet() {
		toSerialize["FeedbackData"] = o.FeedbackData.Get()
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeedbackFeedbackPostAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFeedbackFeedbackPostAllOf := _FeedbackFeedbackPostAllOf{}

	if err = json.Unmarshal(bytes, &varFeedbackFeedbackPostAllOf); err == nil {
		*o = FeedbackFeedbackPostAllOf(varFeedbackFeedbackPostAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FeedbackData")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeedbackFeedbackPostAllOf struct {
	value *FeedbackFeedbackPostAllOf
	isSet bool
}

func (v NullableFeedbackFeedbackPostAllOf) Get() *FeedbackFeedbackPostAllOf {
	return v.value
}

func (v *NullableFeedbackFeedbackPostAllOf) Set(val *FeedbackFeedbackPostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFeedbackPostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFeedbackPostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFeedbackPostAllOf(val *FeedbackFeedbackPostAllOf) *NullableFeedbackFeedbackPostAllOf {
	return &NullableFeedbackFeedbackPostAllOf{value: val, isSet: true}
}

func (v NullableFeedbackFeedbackPostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFeedbackPostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
