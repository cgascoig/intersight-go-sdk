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

// NiatelemetryVniStatus Stores information related to vni details.
type NiatelemetryVniStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the vni id of the vni.
	Vni *string `json:"Vni,omitempty"`
	// Returns the vni state of the vni.
	VniState *string `json:"VniState,omitempty"`
	// Returns the vni type of the vni.
	VniType              *string `json:"VniType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryVniStatus NiatelemetryVniStatus

// NewNiatelemetryVniStatus instantiates a new NiatelemetryVniStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryVniStatus(classId string, objectType string) *NiatelemetryVniStatus {
	this := NiatelemetryVniStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryVniStatusWithDefaults instantiates a new NiatelemetryVniStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryVniStatusWithDefaults() *NiatelemetryVniStatus {
	this := NiatelemetryVniStatus{}
	var classId string = "niatelemetry.VniStatus"
	this.ClassId = classId
	var objectType string = "niatelemetry.VniStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryVniStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryVniStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryVniStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryVniStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryVniStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryVniStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVni returns the Vni field value if set, zero value otherwise.
func (o *NiatelemetryVniStatus) GetVni() string {
	if o == nil || o.Vni == nil {
		var ret string
		return ret
	}
	return *o.Vni
}

// GetVniOk returns a tuple with the Vni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVniStatus) GetVniOk() (*string, bool) {
	if o == nil || o.Vni == nil {
		return nil, false
	}
	return o.Vni, true
}

// HasVni returns a boolean if a field has been set.
func (o *NiatelemetryVniStatus) HasVni() bool {
	if o != nil && o.Vni != nil {
		return true
	}

	return false
}

// SetVni gets a reference to the given string and assigns it to the Vni field.
func (o *NiatelemetryVniStatus) SetVni(v string) {
	o.Vni = &v
}

// GetVniState returns the VniState field value if set, zero value otherwise.
func (o *NiatelemetryVniStatus) GetVniState() string {
	if o == nil || o.VniState == nil {
		var ret string
		return ret
	}
	return *o.VniState
}

// GetVniStateOk returns a tuple with the VniState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVniStatus) GetVniStateOk() (*string, bool) {
	if o == nil || o.VniState == nil {
		return nil, false
	}
	return o.VniState, true
}

// HasVniState returns a boolean if a field has been set.
func (o *NiatelemetryVniStatus) HasVniState() bool {
	if o != nil && o.VniState != nil {
		return true
	}

	return false
}

// SetVniState gets a reference to the given string and assigns it to the VniState field.
func (o *NiatelemetryVniStatus) SetVniState(v string) {
	o.VniState = &v
}

// GetVniType returns the VniType field value if set, zero value otherwise.
func (o *NiatelemetryVniStatus) GetVniType() string {
	if o == nil || o.VniType == nil {
		var ret string
		return ret
	}
	return *o.VniType
}

// GetVniTypeOk returns a tuple with the VniType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryVniStatus) GetVniTypeOk() (*string, bool) {
	if o == nil || o.VniType == nil {
		return nil, false
	}
	return o.VniType, true
}

// HasVniType returns a boolean if a field has been set.
func (o *NiatelemetryVniStatus) HasVniType() bool {
	if o != nil && o.VniType != nil {
		return true
	}

	return false
}

// SetVniType gets a reference to the given string and assigns it to the VniType field.
func (o *NiatelemetryVniStatus) SetVniType(v string) {
	o.VniType = &v
}

func (o NiatelemetryVniStatus) MarshalJSON() ([]byte, error) {
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
	if o.Vni != nil {
		toSerialize["Vni"] = o.Vni
	}
	if o.VniState != nil {
		toSerialize["VniState"] = o.VniState
	}
	if o.VniType != nil {
		toSerialize["VniType"] = o.VniType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryVniStatus) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryVniStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the vni id of the vni.
		Vni *string `json:"Vni,omitempty"`
		// Returns the vni state of the vni.
		VniState *string `json:"VniState,omitempty"`
		// Returns the vni type of the vni.
		VniType *string `json:"VniType,omitempty"`
	}

	varNiatelemetryVniStatusWithoutEmbeddedStruct := NiatelemetryVniStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryVniStatusWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryVniStatus := _NiatelemetryVniStatus{}
		varNiatelemetryVniStatus.ClassId = varNiatelemetryVniStatusWithoutEmbeddedStruct.ClassId
		varNiatelemetryVniStatus.ObjectType = varNiatelemetryVniStatusWithoutEmbeddedStruct.ObjectType
		varNiatelemetryVniStatus.Vni = varNiatelemetryVniStatusWithoutEmbeddedStruct.Vni
		varNiatelemetryVniStatus.VniState = varNiatelemetryVniStatusWithoutEmbeddedStruct.VniState
		varNiatelemetryVniStatus.VniType = varNiatelemetryVniStatusWithoutEmbeddedStruct.VniType
		*o = NiatelemetryVniStatus(varNiatelemetryVniStatus)
	} else {
		return err
	}

	varNiatelemetryVniStatus := _NiatelemetryVniStatus{}

	err = json.Unmarshal(bytes, &varNiatelemetryVniStatus)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryVniStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Vni")
		delete(additionalProperties, "VniState")
		delete(additionalProperties, "VniType")

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

type NullableNiatelemetryVniStatus struct {
	value *NiatelemetryVniStatus
	isSet bool
}

func (v NullableNiatelemetryVniStatus) Get() *NiatelemetryVniStatus {
	return v.value
}

func (v *NullableNiatelemetryVniStatus) Set(val *NiatelemetryVniStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryVniStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryVniStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryVniStatus(val *NiatelemetryVniStatus) *NullableNiatelemetryVniStatus {
	return &NullableNiatelemetryVniStatus{value: val, isSet: true}
}

func (v NullableNiatelemetryVniStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryVniStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
