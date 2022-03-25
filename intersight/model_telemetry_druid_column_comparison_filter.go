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

// TelemetryDruidColumnComparisonFilter The column comparison filter compares dimensions to each other.
type TelemetryDruidColumnComparisonFilter struct {
	// The filter type.
	Type string `json:"type"`
	// All filters except the \"spatial\" filter support extraction functions. An extraction function is defined by setting the \"extractionFn\" field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \"bar_1\".
	ExtractionFn *map[string]interface{} `json:"extractionFn,omitempty"`
	// A list of DimensionSpecs, making it possible to apply an extraction function if needed.
	Dimensions           []TelemetryDruidDimensionSpec `json:"dimensions"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidColumnComparisonFilter TelemetryDruidColumnComparisonFilter

// NewTelemetryDruidColumnComparisonFilter instantiates a new TelemetryDruidColumnComparisonFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidColumnComparisonFilter(type_ string, dimensions []TelemetryDruidDimensionSpec) *TelemetryDruidColumnComparisonFilter {
	this := TelemetryDruidColumnComparisonFilter{}
	this.Type = type_
	this.Dimensions = dimensions
	return &this
}

// NewTelemetryDruidColumnComparisonFilterWithDefaults instantiates a new TelemetryDruidColumnComparisonFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidColumnComparisonFilterWithDefaults() *TelemetryDruidColumnComparisonFilter {
	this := TelemetryDruidColumnComparisonFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidColumnComparisonFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidColumnComparisonFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidColumnComparisonFilter) SetType(v string) {
	o.Type = v
}

// GetExtractionFn returns the ExtractionFn field value if set, zero value otherwise.
func (o *TelemetryDruidColumnComparisonFilter) GetExtractionFn() map[string]interface{} {
	if o == nil || o.ExtractionFn == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ExtractionFn
}

// GetExtractionFnOk returns a tuple with the ExtractionFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidColumnComparisonFilter) GetExtractionFnOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtractionFn == nil {
		return nil, false
	}
	return o.ExtractionFn, true
}

// HasExtractionFn returns a boolean if a field has been set.
func (o *TelemetryDruidColumnComparisonFilter) HasExtractionFn() bool {
	if o != nil && o.ExtractionFn != nil {
		return true
	}

	return false
}

// SetExtractionFn gets a reference to the given map[string]interface{} and assigns it to the ExtractionFn field.
func (o *TelemetryDruidColumnComparisonFilter) SetExtractionFn(v map[string]interface{}) {
	o.ExtractionFn = &v
}

// GetDimensions returns the Dimensions field value
func (o *TelemetryDruidColumnComparisonFilter) GetDimensions() []TelemetryDruidDimensionSpec {
	if o == nil {
		var ret []TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidColumnComparisonFilter) GetDimensionsOk() (*[]TelemetryDruidDimensionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *TelemetryDruidColumnComparisonFilter) SetDimensions(v []TelemetryDruidDimensionSpec) {
	o.Dimensions = v
}

func (o TelemetryDruidColumnComparisonFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ExtractionFn != nil {
		toSerialize["extractionFn"] = o.ExtractionFn
	}
	if true {
		toSerialize["dimensions"] = o.Dimensions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidColumnComparisonFilter) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidColumnComparisonFilter := _TelemetryDruidColumnComparisonFilter{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidColumnComparisonFilter); err == nil {
		*o = TelemetryDruidColumnComparisonFilter(varTelemetryDruidColumnComparisonFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "extractionFn")
		delete(additionalProperties, "dimensions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidColumnComparisonFilter struct {
	value *TelemetryDruidColumnComparisonFilter
	isSet bool
}

func (v NullableTelemetryDruidColumnComparisonFilter) Get() *TelemetryDruidColumnComparisonFilter {
	return v.value
}

func (v *NullableTelemetryDruidColumnComparisonFilter) Set(val *TelemetryDruidColumnComparisonFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidColumnComparisonFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidColumnComparisonFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidColumnComparisonFilter(val *TelemetryDruidColumnComparisonFilter) *NullableTelemetryDruidColumnComparisonFilter {
	return &NullableTelemetryDruidColumnComparisonFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidColumnComparisonFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidColumnComparisonFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
