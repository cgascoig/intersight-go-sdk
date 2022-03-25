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

// TelemetryDruidScanResult Raw Druid row result. Currently only supports 'list' format.
type TelemetryDruidScanResult struct {
	// The identifier for the row(s)' segement
	SegmentId *string `json:"segmentId,omitempty"`
	// A list of columns returned in the row(s)
	Columns *[]string `json:"columns,omitempty"`
	// Row results
	Events               *[]map[string]interface{} `json:"events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidScanResult TelemetryDruidScanResult

// NewTelemetryDruidScanResult instantiates a new TelemetryDruidScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidScanResult() *TelemetryDruidScanResult {
	this := TelemetryDruidScanResult{}
	return &this
}

// NewTelemetryDruidScanResultWithDefaults instantiates a new TelemetryDruidScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidScanResultWithDefaults() *TelemetryDruidScanResult {
	this := TelemetryDruidScanResult{}
	return &this
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *TelemetryDruidScanResult) GetSegmentId() string {
	if o == nil || o.SegmentId == nil {
		var ret string
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanResult) GetSegmentIdOk() (*string, bool) {
	if o == nil || o.SegmentId == nil {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *TelemetryDruidScanResult) HasSegmentId() bool {
	if o != nil && o.SegmentId != nil {
		return true
	}

	return false
}

// SetSegmentId gets a reference to the given string and assigns it to the SegmentId field.
func (o *TelemetryDruidScanResult) SetSegmentId(v string) {
	o.SegmentId = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *TelemetryDruidScanResult) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanResult) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *TelemetryDruidScanResult) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *TelemetryDruidScanResult) SetColumns(v []string) {
	o.Columns = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *TelemetryDruidScanResult) GetEvents() []map[string]interface{} {
	if o == nil || o.Events == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidScanResult) GetEventsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *TelemetryDruidScanResult) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []map[string]interface{} and assigns it to the Events field.
func (o *TelemetryDruidScanResult) SetEvents(v []map[string]interface{}) {
	o.Events = &v
}

func (o TelemetryDruidScanResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SegmentId != nil {
		toSerialize["segmentId"] = o.SegmentId
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidScanResult) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidScanResult := _TelemetryDruidScanResult{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidScanResult); err == nil {
		*o = TelemetryDruidScanResult(varTelemetryDruidScanResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "segmentId")
		delete(additionalProperties, "columns")
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidScanResult struct {
	value *TelemetryDruidScanResult
	isSet bool
}

func (v NullableTelemetryDruidScanResult) Get() *TelemetryDruidScanResult {
	return v.value
}

func (v *NullableTelemetryDruidScanResult) Set(val *TelemetryDruidScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidScanResult(val *TelemetryDruidScanResult) *NullableTelemetryDruidScanResult {
	return &NullableTelemetryDruidScanResult{value: val, isSet: true}
}

func (v NullableTelemetryDruidScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
