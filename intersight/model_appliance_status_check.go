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

// ApplianceStatusCheck Status check on an Intersight Appliance entity.
type ApplianceStatusCheck struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier of the status check.
	Code *string `json:"Code,omitempty"`
	// Result of this status check. * `OK` - Result of the check is OK. * `Warning` - Result of the check is Warning. * `Critical` - Result of the check is Critical. * `Info` - Result of the check is low.
	Result               *string `json:"Result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceStatusCheck ApplianceStatusCheck

// NewApplianceStatusCheck instantiates a new ApplianceStatusCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceStatusCheck(classId string, objectType string) *ApplianceStatusCheck {
	this := ApplianceStatusCheck{}
	this.ClassId = classId
	this.ObjectType = objectType
	var result string = "OK"
	this.Result = &result
	return &this
}

// NewApplianceStatusCheckWithDefaults instantiates a new ApplianceStatusCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceStatusCheckWithDefaults() *ApplianceStatusCheck {
	this := ApplianceStatusCheck{}
	var classId string = "appliance.StatusCheck"
	this.ClassId = classId
	var objectType string = "appliance.StatusCheck"
	this.ObjectType = objectType
	var result string = "OK"
	this.Result = &result
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceStatusCheck) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheck) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceStatusCheck) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceStatusCheck) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheck) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceStatusCheck) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApplianceStatusCheck) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheck) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApplianceStatusCheck) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApplianceStatusCheck) SetCode(v string) {
	o.Code = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ApplianceStatusCheck) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceStatusCheck) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ApplianceStatusCheck) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ApplianceStatusCheck) SetResult(v string) {
	o.Result = &v
}

func (o ApplianceStatusCheck) MarshalJSON() ([]byte, error) {
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
	if o.Code != nil {
		toSerialize["Code"] = o.Code
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceStatusCheck) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceStatusCheckWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identifier of the status check.
		Code *string `json:"Code,omitempty"`
		// Result of this status check. * `OK` - Result of the check is OK. * `Warning` - Result of the check is Warning. * `Critical` - Result of the check is Critical. * `Info` - Result of the check is low.
		Result *string `json:"Result,omitempty"`
	}

	varApplianceStatusCheckWithoutEmbeddedStruct := ApplianceStatusCheckWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceStatusCheckWithoutEmbeddedStruct)
	if err == nil {
		varApplianceStatusCheck := _ApplianceStatusCheck{}
		varApplianceStatusCheck.ClassId = varApplianceStatusCheckWithoutEmbeddedStruct.ClassId
		varApplianceStatusCheck.ObjectType = varApplianceStatusCheckWithoutEmbeddedStruct.ObjectType
		varApplianceStatusCheck.Code = varApplianceStatusCheckWithoutEmbeddedStruct.Code
		varApplianceStatusCheck.Result = varApplianceStatusCheckWithoutEmbeddedStruct.Result
		*o = ApplianceStatusCheck(varApplianceStatusCheck)
	} else {
		return err
	}

	varApplianceStatusCheck := _ApplianceStatusCheck{}

	err = json.Unmarshal(bytes, &varApplianceStatusCheck)
	if err == nil {
		o.MoBaseComplexType = varApplianceStatusCheck.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Code")
		delete(additionalProperties, "Result")

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

type NullableApplianceStatusCheck struct {
	value *ApplianceStatusCheck
	isSet bool
}

func (v NullableApplianceStatusCheck) Get() *ApplianceStatusCheck {
	return v.value
}

func (v *NullableApplianceStatusCheck) Set(val *ApplianceStatusCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceStatusCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceStatusCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceStatusCheck(val *ApplianceStatusCheck) *NullableApplianceStatusCheck {
	return &NullableApplianceStatusCheck{value: val, isSet: true}
}

func (v NullableApplianceStatusCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceStatusCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
