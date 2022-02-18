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
)

// AccessConfigurationTypeAllOf Definition of the list of properties defined in 'access.ConfigurationType', excluding properties defined in parent classes.
type AccessConfigurationTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This flag enables the use of In-Band configuration for end-point access.
	ConfigureInband *bool `json:"ConfigureInband,omitempty"`
	// This flag enables the use of Out-Of-Band configuration for end-point access.
	ConfigureOutOfBand   *bool `json:"ConfigureOutOfBand,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessConfigurationTypeAllOf AccessConfigurationTypeAllOf

// NewAccessConfigurationTypeAllOf instantiates a new AccessConfigurationTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessConfigurationTypeAllOf(classId string, objectType string) *AccessConfigurationTypeAllOf {
	this := AccessConfigurationTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configureInband bool = true
	this.ConfigureInband = &configureInband
	var configureOutOfBand bool = false
	this.ConfigureOutOfBand = &configureOutOfBand
	return &this
}

// NewAccessConfigurationTypeAllOfWithDefaults instantiates a new AccessConfigurationTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessConfigurationTypeAllOfWithDefaults() *AccessConfigurationTypeAllOf {
	this := AccessConfigurationTypeAllOf{}
	var classId string = "access.ConfigurationType"
	this.ClassId = classId
	var objectType string = "access.ConfigurationType"
	this.ObjectType = objectType
	var configureInband bool = true
	this.ConfigureInband = &configureInband
	var configureOutOfBand bool = false
	this.ConfigureOutOfBand = &configureOutOfBand
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessConfigurationTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessConfigurationTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessConfigurationTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AccessConfigurationTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessConfigurationTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessConfigurationTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigureInband returns the ConfigureInband field value if set, zero value otherwise.
func (o *AccessConfigurationTypeAllOf) GetConfigureInband() bool {
	if o == nil || o.ConfigureInband == nil {
		var ret bool
		return ret
	}
	return *o.ConfigureInband
}

// GetConfigureInbandOk returns a tuple with the ConfigureInband field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessConfigurationTypeAllOf) GetConfigureInbandOk() (*bool, bool) {
	if o == nil || o.ConfigureInband == nil {
		return nil, false
	}
	return o.ConfigureInband, true
}

// HasConfigureInband returns a boolean if a field has been set.
func (o *AccessConfigurationTypeAllOf) HasConfigureInband() bool {
	if o != nil && o.ConfigureInband != nil {
		return true
	}

	return false
}

// SetConfigureInband gets a reference to the given bool and assigns it to the ConfigureInband field.
func (o *AccessConfigurationTypeAllOf) SetConfigureInband(v bool) {
	o.ConfigureInband = &v
}

// GetConfigureOutOfBand returns the ConfigureOutOfBand field value if set, zero value otherwise.
func (o *AccessConfigurationTypeAllOf) GetConfigureOutOfBand() bool {
	if o == nil || o.ConfigureOutOfBand == nil {
		var ret bool
		return ret
	}
	return *o.ConfigureOutOfBand
}

// GetConfigureOutOfBandOk returns a tuple with the ConfigureOutOfBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessConfigurationTypeAllOf) GetConfigureOutOfBandOk() (*bool, bool) {
	if o == nil || o.ConfigureOutOfBand == nil {
		return nil, false
	}
	return o.ConfigureOutOfBand, true
}

// HasConfigureOutOfBand returns a boolean if a field has been set.
func (o *AccessConfigurationTypeAllOf) HasConfigureOutOfBand() bool {
	if o != nil && o.ConfigureOutOfBand != nil {
		return true
	}

	return false
}

// SetConfigureOutOfBand gets a reference to the given bool and assigns it to the ConfigureOutOfBand field.
func (o *AccessConfigurationTypeAllOf) SetConfigureOutOfBand(v bool) {
	o.ConfigureOutOfBand = &v
}

func (o AccessConfigurationTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigureInband != nil {
		toSerialize["ConfigureInband"] = o.ConfigureInband
	}
	if o.ConfigureOutOfBand != nil {
		toSerialize["ConfigureOutOfBand"] = o.ConfigureOutOfBand
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessConfigurationTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessConfigurationTypeAllOf := _AccessConfigurationTypeAllOf{}

	if err = json.Unmarshal(bytes, &varAccessConfigurationTypeAllOf); err == nil {
		*o = AccessConfigurationTypeAllOf(varAccessConfigurationTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigureInband")
		delete(additionalProperties, "ConfigureOutOfBand")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessConfigurationTypeAllOf struct {
	value *AccessConfigurationTypeAllOf
	isSet bool
}

func (v NullableAccessConfigurationTypeAllOf) Get() *AccessConfigurationTypeAllOf {
	return v.value
}

func (v *NullableAccessConfigurationTypeAllOf) Set(val *AccessConfigurationTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessConfigurationTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessConfigurationTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessConfigurationTypeAllOf(val *AccessConfigurationTypeAllOf) *NullableAccessConfigurationTypeAllOf {
	return &NullableAccessConfigurationTypeAllOf{value: val, isSet: true}
}

func (v NullableAccessConfigurationTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessConfigurationTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
