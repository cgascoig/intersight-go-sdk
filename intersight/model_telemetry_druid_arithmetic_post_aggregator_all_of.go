/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidArithmeticPostAggregatorAllOf struct for TelemetryDruidArithmeticPostAggregatorAllOf
type TelemetryDruidArithmeticPostAggregatorAllOf struct {
	// Output name for the minimum/maximum timestamp value.
	Name *string `json:"name,omitempty"`
	// null
	Fn *string `json:"fn,omitempty"`
	// Arithmetic post-aggregators may specify an ordering, which defines the order of resulting values when sorting results. This can be useful for topN queries for instance. If no ordering (or null) is specified, the default floating point ordering is used. numericFirst ordering always returns finite values first, followed by NaN, and infinite values last.
	Ordering             *string `json:"ordering,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidArithmeticPostAggregatorAllOf TelemetryDruidArithmeticPostAggregatorAllOf

// NewTelemetryDruidArithmeticPostAggregatorAllOf instantiates a new TelemetryDruidArithmeticPostAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidArithmeticPostAggregatorAllOf() *TelemetryDruidArithmeticPostAggregatorAllOf {
	this := TelemetryDruidArithmeticPostAggregatorAllOf{}
	return &this
}

// NewTelemetryDruidArithmeticPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidArithmeticPostAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidArithmeticPostAggregatorAllOfWithDefaults() *TelemetryDruidArithmeticPostAggregatorAllOf {
	this := TelemetryDruidArithmeticPostAggregatorAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetFn returns the Fn field value if set, zero value otherwise.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFn() string {
	if o == nil || o.Fn == nil {
		var ret string
		return ret
	}
	return *o.Fn
}

// GetFnOk returns a tuple with the Fn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetFnOk() (*string, bool) {
	if o == nil || o.Fn == nil {
		return nil, false
	}
	return o.Fn, true
}

// HasFn returns a boolean if a field has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasFn() bool {
	if o != nil && o.Fn != nil {
		return true
	}

	return false
}

// SetFn gets a reference to the given string and assigns it to the Fn field.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetFn(v string) {
	o.Fn = &v
}

// GetOrdering returns the Ordering field value if set, zero value otherwise.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetOrdering() string {
	if o == nil || o.Ordering == nil {
		var ret string
		return ret
	}
	return *o.Ordering
}

// GetOrderingOk returns a tuple with the Ordering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) GetOrderingOk() (*string, bool) {
	if o == nil || o.Ordering == nil {
		return nil, false
	}
	return o.Ordering, true
}

// HasOrdering returns a boolean if a field has been set.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) HasOrdering() bool {
	if o != nil && o.Ordering != nil {
		return true
	}

	return false
}

// SetOrdering gets a reference to the given string and assigns it to the Ordering field.
func (o *TelemetryDruidArithmeticPostAggregatorAllOf) SetOrdering(v string) {
	o.Ordering = &v
}

func (o TelemetryDruidArithmeticPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Fn != nil {
		toSerialize["fn"] = o.Fn
	}
	if o.Ordering != nil {
		toSerialize["ordering"] = o.Ordering
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidArithmeticPostAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidArithmeticPostAggregatorAllOf := _TelemetryDruidArithmeticPostAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidArithmeticPostAggregatorAllOf); err == nil {
		*o = TelemetryDruidArithmeticPostAggregatorAllOf(varTelemetryDruidArithmeticPostAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "fn")
		delete(additionalProperties, "ordering")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidArithmeticPostAggregatorAllOf struct {
	value *TelemetryDruidArithmeticPostAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidArithmeticPostAggregatorAllOf) Get() *TelemetryDruidArithmeticPostAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidArithmeticPostAggregatorAllOf) Set(val *TelemetryDruidArithmeticPostAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidArithmeticPostAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidArithmeticPostAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidArithmeticPostAggregatorAllOf(val *TelemetryDruidArithmeticPostAggregatorAllOf) *NullableTelemetryDruidArithmeticPostAggregatorAllOf {
	return &NullableTelemetryDruidArithmeticPostAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidArithmeticPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidArithmeticPostAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
