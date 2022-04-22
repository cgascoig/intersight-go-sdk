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

// IamDomainNameInfoAllOf Definition of the list of properties defined in 'iam.DomainNameInfo', excluding properties defined in parent classes.
type IamDomainNameInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Regenerate TXT record and validate TXT record. * `generate` - Generate TXT record for domain name ownership validation. * `verify` - Verify TXT record for domain name ownership validation.
	Action *string `json:"Action,omitempty"`
	// Email domain name. When a user enters an email during login in the Intersight home page, the IdP is picked by matching this domain name with the email domain name for authentication.
	DomainName     *string                   `json:"DomainName,omitempty"`
	FailureDetails NullableIamFailureDetails `json:"FailureDetails,omitempty"`
	// Expiration time of TXT Record.
	RecordExpiryTime *time.Time `json:"RecordExpiryTime,omitempty"`
	// Status of Domain Ownership Verification. * `Pending` - Domain verification is pending. * `Failed` - Domain verification failed. Re-generate token and verify. * `Verified` - Domain verification succeeded. * `Expired` - TXT Record for Domain verification expired.
	Status *string `json:"Status,omitempty"`
	// Resource record used to verify Domain Ownership. Users need to verify the domain by adding the TXT Record in their DNS Host.
	TxtRecord            *string                 `json:"TxtRecord,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamDomainNameInfoAllOf IamDomainNameInfoAllOf

// NewIamDomainNameInfoAllOf instantiates a new IamDomainNameInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamDomainNameInfoAllOf(classId string, objectType string) *IamDomainNameInfoAllOf {
	this := IamDomainNameInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "generate"
	this.Action = &action
	return &this
}

// NewIamDomainNameInfoAllOfWithDefaults instantiates a new IamDomainNameInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamDomainNameInfoAllOfWithDefaults() *IamDomainNameInfoAllOf {
	this := IamDomainNameInfoAllOf{}
	var classId string = "iam.DomainNameInfo"
	this.ClassId = classId
	var objectType string = "iam.DomainNameInfo"
	this.ObjectType = objectType
	var action string = "generate"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamDomainNameInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamDomainNameInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamDomainNameInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamDomainNameInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *IamDomainNameInfoAllOf) SetAction(v string) {
	o.Action = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *IamDomainNameInfoAllOf) SetDomainName(v string) {
	o.DomainName = &v
}

// GetFailureDetails returns the FailureDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamDomainNameInfoAllOf) GetFailureDetails() IamFailureDetails {
	if o == nil || o.FailureDetails.Get() == nil {
		var ret IamFailureDetails
		return ret
	}
	return *o.FailureDetails.Get()
}

// GetFailureDetailsOk returns a tuple with the FailureDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamDomainNameInfoAllOf) GetFailureDetailsOk() (*IamFailureDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureDetails.Get(), o.FailureDetails.IsSet()
}

// HasFailureDetails returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasFailureDetails() bool {
	if o != nil && o.FailureDetails.IsSet() {
		return true
	}

	return false
}

// SetFailureDetails gets a reference to the given NullableIamFailureDetails and assigns it to the FailureDetails field.
func (o *IamDomainNameInfoAllOf) SetFailureDetails(v IamFailureDetails) {
	o.FailureDetails.Set(&v)
}

// SetFailureDetailsNil sets the value for FailureDetails to be an explicit nil
func (o *IamDomainNameInfoAllOf) SetFailureDetailsNil() {
	o.FailureDetails.Set(nil)
}

// UnsetFailureDetails ensures that no value is present for FailureDetails, not even an explicit nil
func (o *IamDomainNameInfoAllOf) UnsetFailureDetails() {
	o.FailureDetails.Unset()
}

// GetRecordExpiryTime returns the RecordExpiryTime field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetRecordExpiryTime() time.Time {
	if o == nil || o.RecordExpiryTime == nil {
		var ret time.Time
		return ret
	}
	return *o.RecordExpiryTime
}

// GetRecordExpiryTimeOk returns a tuple with the RecordExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetRecordExpiryTimeOk() (*time.Time, bool) {
	if o == nil || o.RecordExpiryTime == nil {
		return nil, false
	}
	return o.RecordExpiryTime, true
}

// HasRecordExpiryTime returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasRecordExpiryTime() bool {
	if o != nil && o.RecordExpiryTime != nil {
		return true
	}

	return false
}

// SetRecordExpiryTime gets a reference to the given time.Time and assigns it to the RecordExpiryTime field.
func (o *IamDomainNameInfoAllOf) SetRecordExpiryTime(v time.Time) {
	o.RecordExpiryTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IamDomainNameInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTxtRecord returns the TxtRecord field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetTxtRecord() string {
	if o == nil || o.TxtRecord == nil {
		var ret string
		return ret
	}
	return *o.TxtRecord
}

// GetTxtRecordOk returns a tuple with the TxtRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetTxtRecordOk() (*string, bool) {
	if o == nil || o.TxtRecord == nil {
		return nil, false
	}
	return o.TxtRecord, true
}

// HasTxtRecord returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasTxtRecord() bool {
	if o != nil && o.TxtRecord != nil {
		return true
	}

	return false
}

// SetTxtRecord gets a reference to the given string and assigns it to the TxtRecord field.
func (o *IamDomainNameInfoAllOf) SetTxtRecord(v string) {
	o.TxtRecord = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamDomainNameInfoAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainNameInfoAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamDomainNameInfoAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamDomainNameInfoAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamDomainNameInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.DomainName != nil {
		toSerialize["DomainName"] = o.DomainName
	}
	if o.FailureDetails.IsSet() {
		toSerialize["FailureDetails"] = o.FailureDetails.Get()
	}
	if o.RecordExpiryTime != nil {
		toSerialize["RecordExpiryTime"] = o.RecordExpiryTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TxtRecord != nil {
		toSerialize["TxtRecord"] = o.TxtRecord
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamDomainNameInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamDomainNameInfoAllOf := _IamDomainNameInfoAllOf{}

	if err = json.Unmarshal(bytes, &varIamDomainNameInfoAllOf); err == nil {
		*o = IamDomainNameInfoAllOf(varIamDomainNameInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "DomainName")
		delete(additionalProperties, "FailureDetails")
		delete(additionalProperties, "RecordExpiryTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TxtRecord")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamDomainNameInfoAllOf struct {
	value *IamDomainNameInfoAllOf
	isSet bool
}

func (v NullableIamDomainNameInfoAllOf) Get() *IamDomainNameInfoAllOf {
	return v.value
}

func (v *NullableIamDomainNameInfoAllOf) Set(val *IamDomainNameInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamDomainNameInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamDomainNameInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamDomainNameInfoAllOf(val *IamDomainNameInfoAllOf) *NullableIamDomainNameInfoAllOf {
	return &NullableIamDomainNameInfoAllOf{value: val, isSet: true}
}

func (v NullableIamDomainNameInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamDomainNameInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
