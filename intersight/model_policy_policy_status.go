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
	"reflect"
	"strings"
)

// PolicyPolicyStatus The status descriptor for any policy attach/detach to/from a profile/template. It contains the policy details like moid and type and the status progress like validating, ok or errored. In case of an error, the error message is populated in reason.
type PolicyPolicyStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The object id of the policy being attached.
	Moid *string `json:"Moid,omitempty"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the policy attach/detach was successful or not. Values  -- ok, errored, validating. * `ok` - The policy attach/detach is successful. * `error` - The policy cannot be attached/detached due to an error. * `validating` - The policy preconfig validation is in progress.
	Status *string `json:"Status,omitempty"`
	// The object type of the policy being attached.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyPolicyStatus PolicyPolicyStatus

// NewPolicyPolicyStatus instantiates a new PolicyPolicyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyPolicyStatus(classId string, objectType string) *PolicyPolicyStatus {
	this := PolicyPolicyStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewPolicyPolicyStatusWithDefaults instantiates a new PolicyPolicyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyPolicyStatusWithDefaults() *PolicyPolicyStatus {
	this := PolicyPolicyStatus{}
	var classId string = "policy.PolicyStatus"
	this.ClassId = classId
	var objectType string = "policy.PolicyStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyPolicyStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyPolicyStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyPolicyStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyPolicyStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *PolicyPolicyStatus) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *PolicyPolicyStatus) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *PolicyPolicyStatus) SetMoid(v string) {
	o.Moid = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PolicyPolicyStatus) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PolicyPolicyStatus) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PolicyPolicyStatus) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PolicyPolicyStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PolicyPolicyStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PolicyPolicyStatus) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyPolicyStatus) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatus) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyPolicyStatus) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyPolicyStatus) SetType(v string) {
	o.Type = &v
}

func (o PolicyPolicyStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyPolicyStatus) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyPolicyStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The object id of the policy being attached.
		Moid *string `json:"Moid,omitempty"`
		// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
		Reason *string `json:"Reason,omitempty"`
		// Indicates if the policy attach/detach was successful or not. Values  -- ok, errored, validating. * `ok` - The policy attach/detach is successful. * `error` - The policy cannot be attached/detached due to an error. * `validating` - The policy preconfig validation is in progress.
		Status *string `json:"Status,omitempty"`
		// The object type of the policy being attached.
		Type *string `json:"Type,omitempty"`
	}

	varPolicyPolicyStatusWithoutEmbeddedStruct := PolicyPolicyStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyPolicyStatusWithoutEmbeddedStruct)
	if err == nil {
		varPolicyPolicyStatus := _PolicyPolicyStatus{}
		varPolicyPolicyStatus.ClassId = varPolicyPolicyStatusWithoutEmbeddedStruct.ClassId
		varPolicyPolicyStatus.ObjectType = varPolicyPolicyStatusWithoutEmbeddedStruct.ObjectType
		varPolicyPolicyStatus.Moid = varPolicyPolicyStatusWithoutEmbeddedStruct.Moid
		varPolicyPolicyStatus.Reason = varPolicyPolicyStatusWithoutEmbeddedStruct.Reason
		varPolicyPolicyStatus.Status = varPolicyPolicyStatusWithoutEmbeddedStruct.Status
		varPolicyPolicyStatus.Type = varPolicyPolicyStatusWithoutEmbeddedStruct.Type
		*o = PolicyPolicyStatus(varPolicyPolicyStatus)
	} else {
		return err
	}

	varPolicyPolicyStatus := _PolicyPolicyStatus{}

	err = json.Unmarshal(bytes, &varPolicyPolicyStatus)
	if err == nil {
		o.MoBaseComplexType = varPolicyPolicyStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullablePolicyPolicyStatus struct {
	value *PolicyPolicyStatus
	isSet bool
}

func (v NullablePolicyPolicyStatus) Get() *PolicyPolicyStatus {
	return v.value
}

func (v *NullablePolicyPolicyStatus) Set(val *PolicyPolicyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyPolicyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyPolicyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyPolicyStatus(val *PolicyPolicyStatus) *NullablePolicyPolicyStatus {
	return &NullablePolicyPolicyStatus{value: val, isSet: true}
}

func (v NullablePolicyPolicyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyPolicyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
