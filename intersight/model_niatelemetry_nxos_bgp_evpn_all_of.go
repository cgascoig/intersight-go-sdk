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

// NiatelemetryNxosBgpEvpnAllOf Definition of the list of properties defined in 'niatelemetry.NxosBgpEvpn', excluding properties defined in parent classes.
type NiatelemetryNxosBgpEvpnAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the EVPN mac count.
	NxosEvpnMacCount *string `json:"NxosEvpnMacCount,omitempty"`
	// Returns the BGP EVPN total networks.
	TotalNetworks *int64 `json:"TotalNetworks,omitempty"`
	// Returns the BGP EVPN total paths.
	TotalPaths           *int64 `json:"TotalPaths,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNxosBgpEvpnAllOf NiatelemetryNxosBgpEvpnAllOf

// NewNiatelemetryNxosBgpEvpnAllOf instantiates a new NiatelemetryNxosBgpEvpnAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNxosBgpEvpnAllOf(classId string, objectType string) *NiatelemetryNxosBgpEvpnAllOf {
	this := NiatelemetryNxosBgpEvpnAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNxosBgpEvpnAllOfWithDefaults instantiates a new NiatelemetryNxosBgpEvpnAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNxosBgpEvpnAllOfWithDefaults() *NiatelemetryNxosBgpEvpnAllOf {
	this := NiatelemetryNxosBgpEvpnAllOf{}
	var classId string = "niatelemetry.NxosBgpEvpn"
	this.ClassId = classId
	var objectType string = "niatelemetry.NxosBgpEvpn"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNxosBgpEvpnAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNxosBgpEvpnAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNxosBgpEvpnAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNxosBgpEvpnAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNxosEvpnMacCount returns the NxosEvpnMacCount field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetNxosEvpnMacCount() string {
	if o == nil || o.NxosEvpnMacCount == nil {
		var ret string
		return ret
	}
	return *o.NxosEvpnMacCount
}

// GetNxosEvpnMacCountOk returns a tuple with the NxosEvpnMacCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetNxosEvpnMacCountOk() (*string, bool) {
	if o == nil || o.NxosEvpnMacCount == nil {
		return nil, false
	}
	return o.NxosEvpnMacCount, true
}

// HasNxosEvpnMacCount returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) HasNxosEvpnMacCount() bool {
	if o != nil && o.NxosEvpnMacCount != nil {
		return true
	}

	return false
}

// SetNxosEvpnMacCount gets a reference to the given string and assigns it to the NxosEvpnMacCount field.
func (o *NiatelemetryNxosBgpEvpnAllOf) SetNxosEvpnMacCount(v string) {
	o.NxosEvpnMacCount = &v
}

// GetTotalNetworks returns the TotalNetworks field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetTotalNetworks() int64 {
	if o == nil || o.TotalNetworks == nil {
		var ret int64
		return ret
	}
	return *o.TotalNetworks
}

// GetTotalNetworksOk returns a tuple with the TotalNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetTotalNetworksOk() (*int64, bool) {
	if o == nil || o.TotalNetworks == nil {
		return nil, false
	}
	return o.TotalNetworks, true
}

// HasTotalNetworks returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) HasTotalNetworks() bool {
	if o != nil && o.TotalNetworks != nil {
		return true
	}

	return false
}

// SetTotalNetworks gets a reference to the given int64 and assigns it to the TotalNetworks field.
func (o *NiatelemetryNxosBgpEvpnAllOf) SetTotalNetworks(v int64) {
	o.TotalNetworks = &v
}

// GetTotalPaths returns the TotalPaths field value if set, zero value otherwise.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetTotalPaths() int64 {
	if o == nil || o.TotalPaths == nil {
		var ret int64
		return ret
	}
	return *o.TotalPaths
}

// GetTotalPathsOk returns a tuple with the TotalPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) GetTotalPathsOk() (*int64, bool) {
	if o == nil || o.TotalPaths == nil {
		return nil, false
	}
	return o.TotalPaths, true
}

// HasTotalPaths returns a boolean if a field has been set.
func (o *NiatelemetryNxosBgpEvpnAllOf) HasTotalPaths() bool {
	if o != nil && o.TotalPaths != nil {
		return true
	}

	return false
}

// SetTotalPaths gets a reference to the given int64 and assigns it to the TotalPaths field.
func (o *NiatelemetryNxosBgpEvpnAllOf) SetTotalPaths(v int64) {
	o.TotalPaths = &v
}

func (o NiatelemetryNxosBgpEvpnAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NxosEvpnMacCount != nil {
		toSerialize["NxosEvpnMacCount"] = o.NxosEvpnMacCount
	}
	if o.TotalNetworks != nil {
		toSerialize["TotalNetworks"] = o.TotalNetworks
	}
	if o.TotalPaths != nil {
		toSerialize["TotalPaths"] = o.TotalPaths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNxosBgpEvpnAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNxosBgpEvpnAllOf := _NiatelemetryNxosBgpEvpnAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNxosBgpEvpnAllOf); err == nil {
		*o = NiatelemetryNxosBgpEvpnAllOf(varNiatelemetryNxosBgpEvpnAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NxosEvpnMacCount")
		delete(additionalProperties, "TotalNetworks")
		delete(additionalProperties, "TotalPaths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNxosBgpEvpnAllOf struct {
	value *NiatelemetryNxosBgpEvpnAllOf
	isSet bool
}

func (v NullableNiatelemetryNxosBgpEvpnAllOf) Get() *NiatelemetryNxosBgpEvpnAllOf {
	return v.value
}

func (v *NullableNiatelemetryNxosBgpEvpnAllOf) Set(val *NiatelemetryNxosBgpEvpnAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNxosBgpEvpnAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNxosBgpEvpnAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNxosBgpEvpnAllOf(val *NiatelemetryNxosBgpEvpnAllOf) *NullableNiatelemetryNxosBgpEvpnAllOf {
	return &NullableNiatelemetryNxosBgpEvpnAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNxosBgpEvpnAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNxosBgpEvpnAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
