/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HclConstraint A key value pair, determining a compatibility constraint for the matching version combinations.
type HclConstraint struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name or key of the applicable compatibility constraint.
	ConstraintName *string `json:"ConstraintName,omitempty"`
	// Value of the applicable compatibility constraint. Could either be a string value or a regex.
	ConstraintValue      *string `json:"ConstraintValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclConstraint HclConstraint

// NewHclConstraint instantiates a new HclConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclConstraint(classId string, objectType string) *HclConstraint {
	this := HclConstraint{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclConstraintWithDefaults instantiates a new HclConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclConstraintWithDefaults() *HclConstraint {
	this := HclConstraint{}
	var classId string = "hcl.Constraint"
	this.ClassId = classId
	var objectType string = "hcl.Constraint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclConstraint) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclConstraint) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclConstraint) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclConstraint) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclConstraint) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclConstraint) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConstraintName returns the ConstraintName field value if set, zero value otherwise.
func (o *HclConstraint) GetConstraintName() string {
	if o == nil || o.ConstraintName == nil {
		var ret string
		return ret
	}
	return *o.ConstraintName
}

// GetConstraintNameOk returns a tuple with the ConstraintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclConstraint) GetConstraintNameOk() (*string, bool) {
	if o == nil || o.ConstraintName == nil {
		return nil, false
	}
	return o.ConstraintName, true
}

// HasConstraintName returns a boolean if a field has been set.
func (o *HclConstraint) HasConstraintName() bool {
	if o != nil && o.ConstraintName != nil {
		return true
	}

	return false
}

// SetConstraintName gets a reference to the given string and assigns it to the ConstraintName field.
func (o *HclConstraint) SetConstraintName(v string) {
	o.ConstraintName = &v
}

// GetConstraintValue returns the ConstraintValue field value if set, zero value otherwise.
func (o *HclConstraint) GetConstraintValue() string {
	if o == nil || o.ConstraintValue == nil {
		var ret string
		return ret
	}
	return *o.ConstraintValue
}

// GetConstraintValueOk returns a tuple with the ConstraintValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclConstraint) GetConstraintValueOk() (*string, bool) {
	if o == nil || o.ConstraintValue == nil {
		return nil, false
	}
	return o.ConstraintValue, true
}

// HasConstraintValue returns a boolean if a field has been set.
func (o *HclConstraint) HasConstraintValue() bool {
	if o != nil && o.ConstraintValue != nil {
		return true
	}

	return false
}

// SetConstraintValue gets a reference to the given string and assigns it to the ConstraintValue field.
func (o *HclConstraint) SetConstraintValue(v string) {
	o.ConstraintValue = &v
}

func (o HclConstraint) MarshalJSON() ([]byte, error) {
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
	if o.ConstraintName != nil {
		toSerialize["ConstraintName"] = o.ConstraintName
	}
	if o.ConstraintValue != nil {
		toSerialize["ConstraintValue"] = o.ConstraintValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclConstraint) UnmarshalJSON(bytes []byte) (err error) {
	type HclConstraintWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name or key of the applicable compatibility constraint.
		ConstraintName *string `json:"ConstraintName,omitempty"`
		// Value of the applicable compatibility constraint. Could either be a string value or a regex.
		ConstraintValue *string `json:"ConstraintValue,omitempty"`
	}

	varHclConstraintWithoutEmbeddedStruct := HclConstraintWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclConstraintWithoutEmbeddedStruct)
	if err == nil {
		varHclConstraint := _HclConstraint{}
		varHclConstraint.ClassId = varHclConstraintWithoutEmbeddedStruct.ClassId
		varHclConstraint.ObjectType = varHclConstraintWithoutEmbeddedStruct.ObjectType
		varHclConstraint.ConstraintName = varHclConstraintWithoutEmbeddedStruct.ConstraintName
		varHclConstraint.ConstraintValue = varHclConstraintWithoutEmbeddedStruct.ConstraintValue
		*o = HclConstraint(varHclConstraint)
	} else {
		return err
	}

	varHclConstraint := _HclConstraint{}

	err = json.Unmarshal(bytes, &varHclConstraint)
	if err == nil {
		o.MoBaseComplexType = varHclConstraint.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConstraintName")
		delete(additionalProperties, "ConstraintValue")

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

type NullableHclConstraint struct {
	value *HclConstraint
	isSet bool
}

func (v NullableHclConstraint) Get() *HclConstraint {
	return v.value
}

func (v *NullableHclConstraint) Set(val *HclConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableHclConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableHclConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclConstraint(val *HclConstraint) *NullableHclConstraint {
	return &NullableHclConstraint{value: val, isSet: true}
}

func (v NullableHclConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
