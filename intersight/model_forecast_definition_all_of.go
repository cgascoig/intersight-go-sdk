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

// ForecastDefinitionAllOf Definition of the list of properties defined in 'forecast.Definition', excluding properties defined in parent classes.
type ForecastDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Threshold above which user needs to be indicated through alarm/alert.
	AlertThresholdInPercentage *int64 `json:"AlertThresholdInPercentage,omitempty"`
	// Data source from where we get the data for the metrics to compute regression model. For example Druid.
	DataSource *string `json:"DataSource,omitempty"`
	// Metric for which forecast prediction is done. Metrics are defined in the catalog file. Currently its only HyperFlex cluster storage capacity usage.
	MetricName *string `json:"MetricName,omitempty"`
	// Minimum number of days of data required for computing forecast model.
	MinNumOfDaysOfData *int64 `json:"MinNumOfDaysOfData,omitempty"`
	// Number of days of data queried from the data source (example Druid ) which is used as input data for computing forecast model.
	NumOfDaysOfHistoricalData *int64 `json:"NumOfDaysOfHistoricalData,omitempty"`
	// The platform type for which we want to compute forecast. For example HyperFlex, NetworkElement.
	PlatformType         *string                      `json:"PlatformType,omitempty"`
	Catalog              *ForecastCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastDefinitionAllOf ForecastDefinitionAllOf

// NewForecastDefinitionAllOf instantiates a new ForecastDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastDefinitionAllOf(classId string, objectType string) *ForecastDefinitionAllOf {
	this := ForecastDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewForecastDefinitionAllOfWithDefaults instantiates a new ForecastDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastDefinitionAllOfWithDefaults() *ForecastDefinitionAllOf {
	this := ForecastDefinitionAllOf{}
	var classId string = "forecast.Definition"
	this.ClassId = classId
	var objectType string = "forecast.Definition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ForecastDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlertThresholdInPercentage returns the AlertThresholdInPercentage field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetAlertThresholdInPercentage() int64 {
	if o == nil || o.AlertThresholdInPercentage == nil {
		var ret int64
		return ret
	}
	return *o.AlertThresholdInPercentage
}

// GetAlertThresholdInPercentageOk returns a tuple with the AlertThresholdInPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetAlertThresholdInPercentageOk() (*int64, bool) {
	if o == nil || o.AlertThresholdInPercentage == nil {
		return nil, false
	}
	return o.AlertThresholdInPercentage, true
}

// HasAlertThresholdInPercentage returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasAlertThresholdInPercentage() bool {
	if o != nil && o.AlertThresholdInPercentage != nil {
		return true
	}

	return false
}

// SetAlertThresholdInPercentage gets a reference to the given int64 and assigns it to the AlertThresholdInPercentage field.
func (o *ForecastDefinitionAllOf) SetAlertThresholdInPercentage(v int64) {
	o.AlertThresholdInPercentage = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetDataSource() string {
	if o == nil || o.DataSource == nil {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetDataSourceOk() (*string, bool) {
	if o == nil || o.DataSource == nil {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasDataSource() bool {
	if o != nil && o.DataSource != nil {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *ForecastDefinitionAllOf) SetDataSource(v string) {
	o.DataSource = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasMetricName() bool {
	if o != nil && o.MetricName != nil {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ForecastDefinitionAllOf) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMinNumOfDaysOfData returns the MinNumOfDaysOfData field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetMinNumOfDaysOfData() int64 {
	if o == nil || o.MinNumOfDaysOfData == nil {
		var ret int64
		return ret
	}
	return *o.MinNumOfDaysOfData
}

// GetMinNumOfDaysOfDataOk returns a tuple with the MinNumOfDaysOfData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetMinNumOfDaysOfDataOk() (*int64, bool) {
	if o == nil || o.MinNumOfDaysOfData == nil {
		return nil, false
	}
	return o.MinNumOfDaysOfData, true
}

// HasMinNumOfDaysOfData returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasMinNumOfDaysOfData() bool {
	if o != nil && o.MinNumOfDaysOfData != nil {
		return true
	}

	return false
}

// SetMinNumOfDaysOfData gets a reference to the given int64 and assigns it to the MinNumOfDaysOfData field.
func (o *ForecastDefinitionAllOf) SetMinNumOfDaysOfData(v int64) {
	o.MinNumOfDaysOfData = &v
}

// GetNumOfDaysOfHistoricalData returns the NumOfDaysOfHistoricalData field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetNumOfDaysOfHistoricalData() int64 {
	if o == nil || o.NumOfDaysOfHistoricalData == nil {
		var ret int64
		return ret
	}
	return *o.NumOfDaysOfHistoricalData
}

// GetNumOfDaysOfHistoricalDataOk returns a tuple with the NumOfDaysOfHistoricalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetNumOfDaysOfHistoricalDataOk() (*int64, bool) {
	if o == nil || o.NumOfDaysOfHistoricalData == nil {
		return nil, false
	}
	return o.NumOfDaysOfHistoricalData, true
}

// HasNumOfDaysOfHistoricalData returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasNumOfDaysOfHistoricalData() bool {
	if o != nil && o.NumOfDaysOfHistoricalData != nil {
		return true
	}

	return false
}

// SetNumOfDaysOfHistoricalData gets a reference to the given int64 and assigns it to the NumOfDaysOfHistoricalData field.
func (o *ForecastDefinitionAllOf) SetNumOfDaysOfHistoricalData(v int64) {
	o.NumOfDaysOfHistoricalData = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *ForecastDefinitionAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *ForecastDefinitionAllOf) GetCatalog() ForecastCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret ForecastCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastDefinitionAllOf) GetCatalogOk() (*ForecastCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *ForecastDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given ForecastCatalogRelationship and assigns it to the Catalog field.
func (o *ForecastDefinitionAllOf) SetCatalog(v ForecastCatalogRelationship) {
	o.Catalog = &v
}

func (o ForecastDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlertThresholdInPercentage != nil {
		toSerialize["AlertThresholdInPercentage"] = o.AlertThresholdInPercentage
	}
	if o.DataSource != nil {
		toSerialize["DataSource"] = o.DataSource
	}
	if o.MetricName != nil {
		toSerialize["MetricName"] = o.MetricName
	}
	if o.MinNumOfDaysOfData != nil {
		toSerialize["MinNumOfDaysOfData"] = o.MinNumOfDaysOfData
	}
	if o.NumOfDaysOfHistoricalData != nil {
		toSerialize["NumOfDaysOfHistoricalData"] = o.NumOfDaysOfHistoricalData
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForecastDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varForecastDefinitionAllOf := _ForecastDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varForecastDefinitionAllOf); err == nil {
		*o = ForecastDefinitionAllOf(varForecastDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlertThresholdInPercentage")
		delete(additionalProperties, "DataSource")
		delete(additionalProperties, "MetricName")
		delete(additionalProperties, "MinNumOfDaysOfData")
		delete(additionalProperties, "NumOfDaysOfHistoricalData")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForecastDefinitionAllOf struct {
	value *ForecastDefinitionAllOf
	isSet bool
}

func (v NullableForecastDefinitionAllOf) Get() *ForecastDefinitionAllOf {
	return v.value
}

func (v *NullableForecastDefinitionAllOf) Set(val *ForecastDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastDefinitionAllOf(val *ForecastDefinitionAllOf) *NullableForecastDefinitionAllOf {
	return &NullableForecastDefinitionAllOf{value: val, isSet: true}
}

func (v NullableForecastDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
