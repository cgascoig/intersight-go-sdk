/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexRpoStatus Status for timeliness of replication job in relation to the schedule interval.
type HyperflexRpoStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Actual end time for the snapshot.
	Actual *int64 `json:"Actual,omitempty"`
	// Expected end time for the snapshot.
	Expected *int64 `json:"Expected,omitempty"`
	// A flag to determine if snapshot delivery is delayed.
	RpoExceeded          *bool `json:"RpoExceeded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexRpoStatus HyperflexRpoStatus

// NewHyperflexRpoStatus instantiates a new HyperflexRpoStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexRpoStatus(classId string, objectType string) *HyperflexRpoStatus {
	this := HyperflexRpoStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexRpoStatusWithDefaults instantiates a new HyperflexRpoStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexRpoStatusWithDefaults() *HyperflexRpoStatus {
	this := HyperflexRpoStatus{}
	var classId string = "hyperflex.RpoStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.RpoStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexRpoStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexRpoStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexRpoStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexRpoStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *HyperflexRpoStatus) GetActual() int64 {
	if o == nil || o.Actual == nil {
		var ret int64
		return ret
	}
	return *o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatus) GetActualOk() (*int64, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *HyperflexRpoStatus) HasActual() bool {
	if o != nil && o.Actual != nil {
		return true
	}

	return false
}

// SetActual gets a reference to the given int64 and assigns it to the Actual field.
func (o *HyperflexRpoStatus) SetActual(v int64) {
	o.Actual = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *HyperflexRpoStatus) GetExpected() int64 {
	if o == nil || o.Expected == nil {
		var ret int64
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatus) GetExpectedOk() (*int64, bool) {
	if o == nil || o.Expected == nil {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *HyperflexRpoStatus) HasExpected() bool {
	if o != nil && o.Expected != nil {
		return true
	}

	return false
}

// SetExpected gets a reference to the given int64 and assigns it to the Expected field.
func (o *HyperflexRpoStatus) SetExpected(v int64) {
	o.Expected = &v
}

// GetRpoExceeded returns the RpoExceeded field value if set, zero value otherwise.
func (o *HyperflexRpoStatus) GetRpoExceeded() bool {
	if o == nil || o.RpoExceeded == nil {
		var ret bool
		return ret
	}
	return *o.RpoExceeded
}

// GetRpoExceededOk returns a tuple with the RpoExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatus) GetRpoExceededOk() (*bool, bool) {
	if o == nil || o.RpoExceeded == nil {
		return nil, false
	}
	return o.RpoExceeded, true
}

// HasRpoExceeded returns a boolean if a field has been set.
func (o *HyperflexRpoStatus) HasRpoExceeded() bool {
	if o != nil && o.RpoExceeded != nil {
		return true
	}

	return false
}

// SetRpoExceeded gets a reference to the given bool and assigns it to the RpoExceeded field.
func (o *HyperflexRpoStatus) SetRpoExceeded(v bool) {
	o.RpoExceeded = &v
}

func (o HyperflexRpoStatus) MarshalJSON() ([]byte, error) {
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
	if o.Actual != nil {
		toSerialize["Actual"] = o.Actual
	}
	if o.Expected != nil {
		toSerialize["Expected"] = o.Expected
	}
	if o.RpoExceeded != nil {
		toSerialize["RpoExceeded"] = o.RpoExceeded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexRpoStatus) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexRpoStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Actual end time for the snapshot.
		Actual *int64 `json:"Actual,omitempty"`
		// Expected end time for the snapshot.
		Expected *int64 `json:"Expected,omitempty"`
		// A flag to determine if snapshot delivery is delayed.
		RpoExceeded *bool `json:"RpoExceeded,omitempty"`
	}

	varHyperflexRpoStatusWithoutEmbeddedStruct := HyperflexRpoStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexRpoStatusWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexRpoStatus := _HyperflexRpoStatus{}
		varHyperflexRpoStatus.ClassId = varHyperflexRpoStatusWithoutEmbeddedStruct.ClassId
		varHyperflexRpoStatus.ObjectType = varHyperflexRpoStatusWithoutEmbeddedStruct.ObjectType
		varHyperflexRpoStatus.Actual = varHyperflexRpoStatusWithoutEmbeddedStruct.Actual
		varHyperflexRpoStatus.Expected = varHyperflexRpoStatusWithoutEmbeddedStruct.Expected
		varHyperflexRpoStatus.RpoExceeded = varHyperflexRpoStatusWithoutEmbeddedStruct.RpoExceeded
		*o = HyperflexRpoStatus(varHyperflexRpoStatus)
	} else {
		return err
	}

	varHyperflexRpoStatus := _HyperflexRpoStatus{}

	err = json.Unmarshal(bytes, &varHyperflexRpoStatus)
	if err == nil {
		o.MoBaseComplexType = varHyperflexRpoStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actual")
		delete(additionalProperties, "Expected")
		delete(additionalProperties, "RpoExceeded")

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

type NullableHyperflexRpoStatus struct {
	value *HyperflexRpoStatus
	isSet bool
}

func (v NullableHyperflexRpoStatus) Get() *HyperflexRpoStatus {
	return v.value
}

func (v *NullableHyperflexRpoStatus) Set(val *HyperflexRpoStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexRpoStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexRpoStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexRpoStatus(val *HyperflexRpoStatus) *NullableHyperflexRpoStatus {
	return &NullableHyperflexRpoStatus{value: val, isSet: true}
}

func (v NullableHyperflexRpoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexRpoStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
