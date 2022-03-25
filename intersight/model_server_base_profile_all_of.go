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
)

// ServerBaseProfileAllOf Definition of the list of properties defined in 'server.BaseProfile', excluding properties defined in parent classes.
type ServerBaseProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// UUID address allocation type selected to assign an UUID address for the server. * `NONE` - The user did not assign any UUID address. * `STATIC` - The user assigns a static UUID address. * `POOL` - The user selects a pool from which the address will be leased.
	UuidAddressType      *string                   `json:"UuidAddressType,omitempty"`
	UuidPool             *UuidpoolPoolRelationship `json:"UuidPool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerBaseProfileAllOf ServerBaseProfileAllOf

// NewServerBaseProfileAllOf instantiates a new ServerBaseProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerBaseProfileAllOf(classId string, objectType string) *ServerBaseProfileAllOf {
	this := ServerBaseProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var uuidAddressType string = "NONE"
	this.UuidAddressType = &uuidAddressType
	return &this
}

// NewServerBaseProfileAllOfWithDefaults instantiates a new ServerBaseProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerBaseProfileAllOfWithDefaults() *ServerBaseProfileAllOf {
	this := ServerBaseProfileAllOf{}
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	var uuidAddressType string = "NONE"
	this.UuidAddressType = &uuidAddressType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerBaseProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerBaseProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerBaseProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerBaseProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerBaseProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerBaseProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ServerBaseProfileAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfileAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ServerBaseProfileAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ServerBaseProfileAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetUuidAddressType returns the UuidAddressType field value if set, zero value otherwise.
func (o *ServerBaseProfileAllOf) GetUuidAddressType() string {
	if o == nil || o.UuidAddressType == nil {
		var ret string
		return ret
	}
	return *o.UuidAddressType
}

// GetUuidAddressTypeOk returns a tuple with the UuidAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfileAllOf) GetUuidAddressTypeOk() (*string, bool) {
	if o == nil || o.UuidAddressType == nil {
		return nil, false
	}
	return o.UuidAddressType, true
}

// HasUuidAddressType returns a boolean if a field has been set.
func (o *ServerBaseProfileAllOf) HasUuidAddressType() bool {
	if o != nil && o.UuidAddressType != nil {
		return true
	}

	return false
}

// SetUuidAddressType gets a reference to the given string and assigns it to the UuidAddressType field.
func (o *ServerBaseProfileAllOf) SetUuidAddressType(v string) {
	o.UuidAddressType = &v
}

// GetUuidPool returns the UuidPool field value if set, zero value otherwise.
func (o *ServerBaseProfileAllOf) GetUuidPool() UuidpoolPoolRelationship {
	if o == nil || o.UuidPool == nil {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.UuidPool
}

// GetUuidPoolOk returns a tuple with the UuidPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerBaseProfileAllOf) GetUuidPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil || o.UuidPool == nil {
		return nil, false
	}
	return o.UuidPool, true
}

// HasUuidPool returns a boolean if a field has been set.
func (o *ServerBaseProfileAllOf) HasUuidPool() bool {
	if o != nil && o.UuidPool != nil {
		return true
	}

	return false
}

// SetUuidPool gets a reference to the given UuidpoolPoolRelationship and assigns it to the UuidPool field.
func (o *ServerBaseProfileAllOf) SetUuidPool(v UuidpoolPoolRelationship) {
	o.UuidPool = &v
}

func (o ServerBaseProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.UuidAddressType != nil {
		toSerialize["UuidAddressType"] = o.UuidAddressType
	}
	if o.UuidPool != nil {
		toSerialize["UuidPool"] = o.UuidPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerBaseProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerBaseProfileAllOf := _ServerBaseProfileAllOf{}

	if err = json.Unmarshal(bytes, &varServerBaseProfileAllOf); err == nil {
		*o = ServerBaseProfileAllOf(varServerBaseProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "UuidAddressType")
		delete(additionalProperties, "UuidPool")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerBaseProfileAllOf struct {
	value *ServerBaseProfileAllOf
	isSet bool
}

func (v NullableServerBaseProfileAllOf) Get() *ServerBaseProfileAllOf {
	return v.value
}

func (v *NullableServerBaseProfileAllOf) Set(val *ServerBaseProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerBaseProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerBaseProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerBaseProfileAllOf(val *ServerBaseProfileAllOf) *NullableServerBaseProfileAllOf {
	return &NullableServerBaseProfileAllOf{value: val, isSet: true}
}

func (v NullableServerBaseProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerBaseProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
