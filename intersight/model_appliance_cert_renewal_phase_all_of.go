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
	"time"
)

// ApplianceCertRenewalPhaseAllOf Definition of the list of properties defined in 'appliance.CertRenewalPhase', excluding properties defined in parent classes.
type ApplianceCertRenewalPhaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// End date of the cert renewal phase.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Indicates if the cert renewal phase has failed or not.
	Failed *bool `json:"Failed,omitempty"`
	// Status message set during the cert renewal phase.
	Message *string `json:"Message,omitempty"`
	// Name of the cert renewal phase phase. * `Init` - New certificate detected, cleanup the old process if any running. * `ScheduleCertificateAddOperation` - Certificate Add Operation Schedulled. * `WaitForCertCollectionByEndpoint` - Monitor cert collection by endpoint. * `Success` - Certificate Renewal Task Success. * `Fail` - Certificate Renewal Task Fail. * `Cancel` - Certificate Renewal Task Cancel.
	Name *string `json:"Name,omitempty"`
	// Start date of the cert renewal phase.
	StartTime            *time.Time `json:"StartTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceCertRenewalPhaseAllOf ApplianceCertRenewalPhaseAllOf

// NewApplianceCertRenewalPhaseAllOf instantiates a new ApplianceCertRenewalPhaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceCertRenewalPhaseAllOf(classId string, objectType string) *ApplianceCertRenewalPhaseAllOf {
	this := ApplianceCertRenewalPhaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceCertRenewalPhaseAllOfWithDefaults instantiates a new ApplianceCertRenewalPhaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceCertRenewalPhaseAllOfWithDefaults() *ApplianceCertRenewalPhaseAllOf {
	this := ApplianceCertRenewalPhaseAllOf{}
	var classId string = "appliance.CertRenewalPhase"
	this.ClassId = classId
	var objectType string = "appliance.CertRenewalPhase"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceCertRenewalPhaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceCertRenewalPhaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceCertRenewalPhaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceCertRenewalPhaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhaseAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhaseAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplianceCertRenewalPhaseAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhaseAllOf) GetFailed() bool {
	if o == nil || o.Failed == nil {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetFailedOk() (*bool, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhaseAllOf) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *ApplianceCertRenewalPhaseAllOf) SetFailed(v bool) {
	o.Failed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhaseAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhaseAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceCertRenewalPhaseAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhaseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhaseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplianceCertRenewalPhaseAllOf) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplianceCertRenewalPhaseAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceCertRenewalPhaseAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplianceCertRenewalPhaseAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplianceCertRenewalPhaseAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o ApplianceCertRenewalPhaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Failed != nil {
		toSerialize["Failed"] = o.Failed
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceCertRenewalPhaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceCertRenewalPhaseAllOf := _ApplianceCertRenewalPhaseAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceCertRenewalPhaseAllOf); err == nil {
		*o = ApplianceCertRenewalPhaseAllOf(varApplianceCertRenewalPhaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Failed")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StartTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceCertRenewalPhaseAllOf struct {
	value *ApplianceCertRenewalPhaseAllOf
	isSet bool
}

func (v NullableApplianceCertRenewalPhaseAllOf) Get() *ApplianceCertRenewalPhaseAllOf {
	return v.value
}

func (v *NullableApplianceCertRenewalPhaseAllOf) Set(val *ApplianceCertRenewalPhaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceCertRenewalPhaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceCertRenewalPhaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceCertRenewalPhaseAllOf(val *ApplianceCertRenewalPhaseAllOf) *NullableApplianceCertRenewalPhaseAllOf {
	return &NullableApplianceCertRenewalPhaseAllOf{value: val, isSet: true}
}

func (v NullableApplianceCertRenewalPhaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceCertRenewalPhaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
