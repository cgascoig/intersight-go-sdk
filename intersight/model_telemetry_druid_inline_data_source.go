/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidInlineDataSource Inline datasources allow you to query a small amount of data that is embedded in the query itself. They are useful when you want to write a query on a small amount of data without loading it first. They are also useful as inputs into a join. Druid also uses them internally to handle subqueries that need to be inlined on the Broker. There are two fields in an inline datasource, an array of columnNames and an array of rows. Each row is an array that must be exactly as long as the list of columnNames. The first element in each row corresponds to the first column in columnNames, and so on.
type TelemetryDruidInlineDataSource struct {
	// The type of data source.
	Type string `json:"type"`
	// the column names.
	ColumnNames []string `json:"columnNames"`
	// an array of rows.
	Rows                 [][]string `json:"rows"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidInlineDataSource TelemetryDruidInlineDataSource

// NewTelemetryDruidInlineDataSource instantiates a new TelemetryDruidInlineDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidInlineDataSource(type_ string, columnNames []string, rows [][]string) *TelemetryDruidInlineDataSource {
	this := TelemetryDruidInlineDataSource{}
	this.Type = type_
	this.ColumnNames = columnNames
	this.Rows = rows
	return &this
}

// NewTelemetryDruidInlineDataSourceWithDefaults instantiates a new TelemetryDruidInlineDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidInlineDataSourceWithDefaults() *TelemetryDruidInlineDataSource {
	this := TelemetryDruidInlineDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidInlineDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInlineDataSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidInlineDataSource) SetType(v string) {
	o.Type = v
}

// GetColumnNames returns the ColumnNames field value
func (o *TelemetryDruidInlineDataSource) GetColumnNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ColumnNames
}

// GetColumnNamesOk returns a tuple with the ColumnNames field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInlineDataSource) GetColumnNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnNames, true
}

// SetColumnNames sets field value
func (o *TelemetryDruidInlineDataSource) SetColumnNames(v []string) {
	o.ColumnNames = v
}

// GetRows returns the Rows field value
func (o *TelemetryDruidInlineDataSource) GetRows() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidInlineDataSource) GetRowsOk() (*[][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value
func (o *TelemetryDruidInlineDataSource) SetRows(v [][]string) {
	o.Rows = v
}

func (o TelemetryDruidInlineDataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["columnNames"] = o.ColumnNames
	}
	if true {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidInlineDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidInlineDataSource := _TelemetryDruidInlineDataSource{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidInlineDataSource); err == nil {
		*o = TelemetryDruidInlineDataSource(varTelemetryDruidInlineDataSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "columnNames")
		delete(additionalProperties, "rows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidInlineDataSource struct {
	value *TelemetryDruidInlineDataSource
	isSet bool
}

func (v NullableTelemetryDruidInlineDataSource) Get() *TelemetryDruidInlineDataSource {
	return v.value
}

func (v *NullableTelemetryDruidInlineDataSource) Set(val *TelemetryDruidInlineDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidInlineDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidInlineDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidInlineDataSource(val *TelemetryDruidInlineDataSource) *NullableTelemetryDruidInlineDataSource {
	return &NullableTelemetryDruidInlineDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidInlineDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidInlineDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
