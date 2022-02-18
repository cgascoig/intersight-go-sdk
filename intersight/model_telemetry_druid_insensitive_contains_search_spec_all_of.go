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

// TelemetryDruidInsensitiveContainsSearchSpecAllOf struct for TelemetryDruidInsensitiveContainsSearchSpecAllOf
type TelemetryDruidInsensitiveContainsSearchSpecAllOf struct {
	// The value to match.  If any part of a dimension value contains the value specified in this search query spec, regardless of case, a \"match\" occurs.
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInsensitiveContainsSearchSpecAllOf TelemetryDruidInsensitiveContainsSearchSpecAllOf

// NewTelemetryDruidInsensitiveContainsSearchSpecAllOf instantiates a new TelemetryDruidInsensitiveContainsSearchSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInsensitiveContainsSearchSpecAllOf(value string) *TelemetryDruidInsensitiveContainsSearchSpecAllOf {
	this := TelemetryDruidInsensitiveContainsSearchSpecAllOf{}
	this.Value = value
	return &this
}

// NewTelemetryDruidInsensitiveContainsSearchSpecAllOfWithDefaults instantiates a new TelemetryDruidInsensitiveContainsSearchSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInsensitiveContainsSearchSpecAllOfWithDefaults() *TelemetryDruidInsensitiveContainsSearchSpecAllOf {
	this := TelemetryDruidInsensitiveContainsSearchSpecAllOf{}
	return &this
}

// GetValue returns the Value field value
func (o *TelemetryDruidInsensitiveContainsSearchSpecAllOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInsensitiveContainsSearchSpecAllOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidInsensitiveContainsSearchSpecAllOf) SetValue(v string) {
	o.Value = v
}

func (o TelemetryDruidInsensitiveContainsSearchSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidInsensitiveContainsSearchSpecAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidInsensitiveContainsSearchSpecAllOf := _TelemetryDruidInsensitiveContainsSearchSpecAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidInsensitiveContainsSearchSpecAllOf); err == nil {
		*o = TelemetryDruidInsensitiveContainsSearchSpecAllOf(varTelemetryDruidInsensitiveContainsSearchSpecAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf struct {
	value *TelemetryDruidInsensitiveContainsSearchSpecAllOf
	isSet bool
}

func (v NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) Get() *TelemetryDruidInsensitiveContainsSearchSpecAllOf {
	return v.value
}

func (v *NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) Set(val *TelemetryDruidInsensitiveContainsSearchSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInsensitiveContainsSearchSpecAllOf(val *TelemetryDruidInsensitiveContainsSearchSpecAllOf) *NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf {
	return &NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInsensitiveContainsSearchSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
