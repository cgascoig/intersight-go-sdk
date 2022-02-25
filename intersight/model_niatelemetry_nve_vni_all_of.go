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
)

// NiatelemetryNveVniAllOf Definition of the list of properties defined in 'niatelemetry.NveVni', excluding properties defined in parent classes.
type NiatelemetryNveVniAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of cp vni count.
	CpVniCount *int64 `json:"CpVniCount,omitempty"`
	// Return value of cp vni down count.
	CpVniDown *int64 `json:"CpVniDown,omitempty"`
	// Return value of cp vni up count.
	CpVniUp *int64 `json:"CpVniUp,omitempty"`
	// Return value of dp vni count.
	DpVniCount *int64 `json:"DpVniCount,omitempty"`
	// Return value of cp vni down count.
	DpVniDown *int64 `json:"DpVniDown,omitempty"`
	// Return value of cp vni up count.
	DpVniUp              *int64 `json:"DpVniUp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNveVniAllOf NiatelemetryNveVniAllOf

// NewNiatelemetryNveVniAllOf instantiates a new NiatelemetryNveVniAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNveVniAllOf(classId string, objectType string) *NiatelemetryNveVniAllOf {
	this := NiatelemetryNveVniAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNveVniAllOfWithDefaults instantiates a new NiatelemetryNveVniAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNveVniAllOfWithDefaults() *NiatelemetryNveVniAllOf {
	this := NiatelemetryNveVniAllOf{}
	var classId string = "niatelemetry.NveVni"
	this.ClassId = classId
	var objectType string = "niatelemetry.NveVni"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNveVniAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNveVniAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNveVniAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNveVniAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpVniCount returns the CpVniCount field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetCpVniCount() int64 {
	if o == nil || o.CpVniCount == nil {
		var ret int64
		return ret
	}
	return *o.CpVniCount
}

// GetCpVniCountOk returns a tuple with the CpVniCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetCpVniCountOk() (*int64, bool) {
	if o == nil || o.CpVniCount == nil {
		return nil, false
	}
	return o.CpVniCount, true
}

// HasCpVniCount returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasCpVniCount() bool {
	if o != nil && o.CpVniCount != nil {
		return true
	}

	return false
}

// SetCpVniCount gets a reference to the given int64 and assigns it to the CpVniCount field.
func (o *NiatelemetryNveVniAllOf) SetCpVniCount(v int64) {
	o.CpVniCount = &v
}

// GetCpVniDown returns the CpVniDown field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetCpVniDown() int64 {
	if o == nil || o.CpVniDown == nil {
		var ret int64
		return ret
	}
	return *o.CpVniDown
}

// GetCpVniDownOk returns a tuple with the CpVniDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetCpVniDownOk() (*int64, bool) {
	if o == nil || o.CpVniDown == nil {
		return nil, false
	}
	return o.CpVniDown, true
}

// HasCpVniDown returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasCpVniDown() bool {
	if o != nil && o.CpVniDown != nil {
		return true
	}

	return false
}

// SetCpVniDown gets a reference to the given int64 and assigns it to the CpVniDown field.
func (o *NiatelemetryNveVniAllOf) SetCpVniDown(v int64) {
	o.CpVniDown = &v
}

// GetCpVniUp returns the CpVniUp field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetCpVniUp() int64 {
	if o == nil || o.CpVniUp == nil {
		var ret int64
		return ret
	}
	return *o.CpVniUp
}

// GetCpVniUpOk returns a tuple with the CpVniUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetCpVniUpOk() (*int64, bool) {
	if o == nil || o.CpVniUp == nil {
		return nil, false
	}
	return o.CpVniUp, true
}

// HasCpVniUp returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasCpVniUp() bool {
	if o != nil && o.CpVniUp != nil {
		return true
	}

	return false
}

// SetCpVniUp gets a reference to the given int64 and assigns it to the CpVniUp field.
func (o *NiatelemetryNveVniAllOf) SetCpVniUp(v int64) {
	o.CpVniUp = &v
}

// GetDpVniCount returns the DpVniCount field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetDpVniCount() int64 {
	if o == nil || o.DpVniCount == nil {
		var ret int64
		return ret
	}
	return *o.DpVniCount
}

// GetDpVniCountOk returns a tuple with the DpVniCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetDpVniCountOk() (*int64, bool) {
	if o == nil || o.DpVniCount == nil {
		return nil, false
	}
	return o.DpVniCount, true
}

// HasDpVniCount returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasDpVniCount() bool {
	if o != nil && o.DpVniCount != nil {
		return true
	}

	return false
}

// SetDpVniCount gets a reference to the given int64 and assigns it to the DpVniCount field.
func (o *NiatelemetryNveVniAllOf) SetDpVniCount(v int64) {
	o.DpVniCount = &v
}

// GetDpVniDown returns the DpVniDown field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetDpVniDown() int64 {
	if o == nil || o.DpVniDown == nil {
		var ret int64
		return ret
	}
	return *o.DpVniDown
}

// GetDpVniDownOk returns a tuple with the DpVniDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetDpVniDownOk() (*int64, bool) {
	if o == nil || o.DpVniDown == nil {
		return nil, false
	}
	return o.DpVniDown, true
}

// HasDpVniDown returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasDpVniDown() bool {
	if o != nil && o.DpVniDown != nil {
		return true
	}

	return false
}

// SetDpVniDown gets a reference to the given int64 and assigns it to the DpVniDown field.
func (o *NiatelemetryNveVniAllOf) SetDpVniDown(v int64) {
	o.DpVniDown = &v
}

// GetDpVniUp returns the DpVniUp field value if set, zero value otherwise.
func (o *NiatelemetryNveVniAllOf) GetDpVniUp() int64 {
	if o == nil || o.DpVniUp == nil {
		var ret int64
		return ret
	}
	return *o.DpVniUp
}

// GetDpVniUpOk returns a tuple with the DpVniUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNveVniAllOf) GetDpVniUpOk() (*int64, bool) {
	if o == nil || o.DpVniUp == nil {
		return nil, false
	}
	return o.DpVniUp, true
}

// HasDpVniUp returns a boolean if a field has been set.
func (o *NiatelemetryNveVniAllOf) HasDpVniUp() bool {
	if o != nil && o.DpVniUp != nil {
		return true
	}

	return false
}

// SetDpVniUp gets a reference to the given int64 and assigns it to the DpVniUp field.
func (o *NiatelemetryNveVniAllOf) SetDpVniUp(v int64) {
	o.DpVniUp = &v
}

func (o NiatelemetryNveVniAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpVniCount != nil {
		toSerialize["CpVniCount"] = o.CpVniCount
	}
	if o.CpVniDown != nil {
		toSerialize["CpVniDown"] = o.CpVniDown
	}
	if o.CpVniUp != nil {
		toSerialize["CpVniUp"] = o.CpVniUp
	}
	if o.DpVniCount != nil {
		toSerialize["DpVniCount"] = o.DpVniCount
	}
	if o.DpVniDown != nil {
		toSerialize["DpVniDown"] = o.DpVniDown
	}
	if o.DpVniUp != nil {
		toSerialize["DpVniUp"] = o.DpVniUp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNveVniAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNveVniAllOf := _NiatelemetryNveVniAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNveVniAllOf); err == nil {
		*o = NiatelemetryNveVniAllOf(varNiatelemetryNveVniAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpVniCount")
		delete(additionalProperties, "CpVniDown")
		delete(additionalProperties, "CpVniUp")
		delete(additionalProperties, "DpVniCount")
		delete(additionalProperties, "DpVniDown")
		delete(additionalProperties, "DpVniUp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNveVniAllOf struct {
	value *NiatelemetryNveVniAllOf
	isSet bool
}

func (v NullableNiatelemetryNveVniAllOf) Get() *NiatelemetryNveVniAllOf {
	return v.value
}

func (v *NullableNiatelemetryNveVniAllOf) Set(val *NiatelemetryNveVniAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNveVniAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNveVniAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNveVniAllOf(val *NiatelemetryNveVniAllOf) *NullableNiatelemetryNveVniAllOf {
	return &NullableNiatelemetryNveVniAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNveVniAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNveVniAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
