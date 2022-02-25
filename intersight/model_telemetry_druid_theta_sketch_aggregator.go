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

// TelemetryDruidThetaSketchAggregator Theta sketch aggregation provides a class of specialized algorithms, called streaming algorithms, or sketches that can produce results orders-of magnitude faster and with mathematically proven error bounds. For interactive queries there may not be other viable alternatives, and in the case of real-time analysis, sketches are the only known solution. In the analysis of big data there are often problem queries that don't scale because they require huge compute resources and time to generate exact results. Examples include count distinct, quantiles, most-frequent items, joins, matrix computations, and graph analysis. If approximate results are acceptable, theta sketch aggregation can be used. Logically speaking, a Theta sketch object can be thought of as a set data structure. Sketches are read and aggregated (set unioned) together. A theta sketch query returns the estimate of the number of unique entries in the sketch object. You can use post aggregators to do union, intersection or difference on sketch columns in the same row. You can use thetaSketch aggregator on columns which were not ingested using the same. It will return estimated cardinality of the column.
type TelemetryDruidThetaSketchAggregator struct {
	// The aggregator type.
	Type string `json:"type"`
	// A String for the output (result) name of the calculation.
	Name string `json:"name"`
	// A String for the name of the aggregator used at ingestion time.
	FieldName string `json:"fieldName"`
	// Must be a power of 2. Internally, size refers to the maximum number of entries sketch object will retain. Higher size means higher accuracy but more space to store sketches. Note that after you index with a particular size, druid will persist sketch in segments and you will use size greater or equal to that at query time. In general, We recommend just sticking to default size.
	Size                 *int32 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidThetaSketchAggregator TelemetryDruidThetaSketchAggregator

// NewTelemetryDruidThetaSketchAggregator instantiates a new TelemetryDruidThetaSketchAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidThetaSketchAggregator(type_ string, name string, fieldName string) *TelemetryDruidThetaSketchAggregator {
	this := TelemetryDruidThetaSketchAggregator{}
	this.Type = type_
	this.Name = name
	this.FieldName = fieldName
	var size int32 = 16384
	this.Size = &size
	return &this
}

// NewTelemetryDruidThetaSketchAggregatorWithDefaults instantiates a new TelemetryDruidThetaSketchAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidThetaSketchAggregatorWithDefaults() *TelemetryDruidThetaSketchAggregator {
	this := TelemetryDruidThetaSketchAggregator{}
	var size int32 = 16384
	this.Size = &size
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidThetaSketchAggregator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchAggregator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidThetaSketchAggregator) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *TelemetryDruidThetaSketchAggregator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchAggregator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TelemetryDruidThetaSketchAggregator) SetName(v string) {
	o.Name = v
}

// GetFieldName returns the FieldName field value
func (o *TelemetryDruidThetaSketchAggregator) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchAggregator) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *TelemetryDruidThetaSketchAggregator) SetFieldName(v string) {
	o.FieldName = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchAggregator) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchAggregator) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchAggregator) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TelemetryDruidThetaSketchAggregator) SetSize(v int32) {
	o.Size = &v
}

func (o TelemetryDruidThetaSketchAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fieldName"] = o.FieldName
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidThetaSketchAggregator) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidThetaSketchAggregator := _TelemetryDruidThetaSketchAggregator{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidThetaSketchAggregator); err == nil {
		*o = TelemetryDruidThetaSketchAggregator(varTelemetryDruidThetaSketchAggregator)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "fieldName")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidThetaSketchAggregator struct {
	value *TelemetryDruidThetaSketchAggregator
	isSet bool
}

func (v NullableTelemetryDruidThetaSketchAggregator) Get() *TelemetryDruidThetaSketchAggregator {
	return v.value
}

func (v *NullableTelemetryDruidThetaSketchAggregator) Set(val *TelemetryDruidThetaSketchAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidThetaSketchAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidThetaSketchAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidThetaSketchAggregator(val *TelemetryDruidThetaSketchAggregator) *NullableTelemetryDruidThetaSketchAggregator {
	return &NullableTelemetryDruidThetaSketchAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidThetaSketchAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidThetaSketchAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
