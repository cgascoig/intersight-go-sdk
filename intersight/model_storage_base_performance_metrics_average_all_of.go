/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageBasePerformanceMetricsAverageAllOf Definition of the list of properties defined in 'storage.BasePerformanceMetricsAverage', excluding properties defined in parent classes.
type StorageBasePerformanceMetricsAverageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Rate of I/O operations observed at the storage object.
	Iops *float64 `json:"Iops,omitempty"`
	// Latency observed at the storage object.
	Latency *float64 `json:"Latency,omitempty"`
	// Duration of periodic aggregation, in hours.
	Period *int64 `json:"Period,omitempty"`
	// Throughput observed at the storage object.
	Throughput           *float64 `json:"Throughput,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBasePerformanceMetricsAverageAllOf StorageBasePerformanceMetricsAverageAllOf

// NewStorageBasePerformanceMetricsAverageAllOf instantiates a new StorageBasePerformanceMetricsAverageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBasePerformanceMetricsAverageAllOf(classId string, objectType string) *StorageBasePerformanceMetricsAverageAllOf {
	this := StorageBasePerformanceMetricsAverageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBasePerformanceMetricsAverageAllOfWithDefaults instantiates a new StorageBasePerformanceMetricsAverageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBasePerformanceMetricsAverageAllOfWithDefaults() *StorageBasePerformanceMetricsAverageAllOf {
	this := StorageBasePerformanceMetricsAverageAllOf{}
	var classId string = "storage.NetAppPerformanceMetricsAverage"
	this.ClassId = classId
	var objectType string = "storage.NetAppPerformanceMetricsAverage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBasePerformanceMetricsAverageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBasePerformanceMetricsAverageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBasePerformanceMetricsAverageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBasePerformanceMetricsAverageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetIops() float64 {
	if o == nil || o.Iops == nil {
		var ret float64
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetIopsOk() (*float64, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given float64 and assigns it to the Iops field.
func (o *StorageBasePerformanceMetricsAverageAllOf) SetIops(v float64) {
	o.Iops = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetLatency() float64 {
	if o == nil || o.Latency == nil {
		var ret float64
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetLatencyOk() (*float64, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) HasLatency() bool {
	if o != nil && o.Latency != nil {
		return true
	}

	return false
}

// SetLatency gets a reference to the given float64 and assigns it to the Latency field.
func (o *StorageBasePerformanceMetricsAverageAllOf) SetLatency(v float64) {
	o.Latency = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetPeriod() int64 {
	if o == nil || o.Period == nil {
		var ret int64
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetPeriodOk() (*int64, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int64 and assigns it to the Period field.
func (o *StorageBasePerformanceMetricsAverageAllOf) SetPeriod(v int64) {
	o.Period = &v
}

// GetThroughput returns the Throughput field value if set, zero value otherwise.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetThroughput() float64 {
	if o == nil || o.Throughput == nil {
		var ret float64
		return ret
	}
	return *o.Throughput
}

// GetThroughputOk returns a tuple with the Throughput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) GetThroughputOk() (*float64, bool) {
	if o == nil || o.Throughput == nil {
		return nil, false
	}
	return o.Throughput, true
}

// HasThroughput returns a boolean if a field has been set.
func (o *StorageBasePerformanceMetricsAverageAllOf) HasThroughput() bool {
	if o != nil && o.Throughput != nil {
		return true
	}

	return false
}

// SetThroughput gets a reference to the given float64 and assigns it to the Throughput field.
func (o *StorageBasePerformanceMetricsAverageAllOf) SetThroughput(v float64) {
	o.Throughput = &v
}

func (o StorageBasePerformanceMetricsAverageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.Latency != nil {
		toSerialize["Latency"] = o.Latency
	}
	if o.Period != nil {
		toSerialize["Period"] = o.Period
	}
	if o.Throughput != nil {
		toSerialize["Throughput"] = o.Throughput
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBasePerformanceMetricsAverageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBasePerformanceMetricsAverageAllOf := _StorageBasePerformanceMetricsAverageAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBasePerformanceMetricsAverageAllOf); err == nil {
		*o = StorageBasePerformanceMetricsAverageAllOf(varStorageBasePerformanceMetricsAverageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Iops")
		delete(additionalProperties, "Latency")
		delete(additionalProperties, "Period")
		delete(additionalProperties, "Throughput")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBasePerformanceMetricsAverageAllOf struct {
	value *StorageBasePerformanceMetricsAverageAllOf
	isSet bool
}

func (v NullableStorageBasePerformanceMetricsAverageAllOf) Get() *StorageBasePerformanceMetricsAverageAllOf {
	return v.value
}

func (v *NullableStorageBasePerformanceMetricsAverageAllOf) Set(val *StorageBasePerformanceMetricsAverageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBasePerformanceMetricsAverageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBasePerformanceMetricsAverageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBasePerformanceMetricsAverageAllOf(val *StorageBasePerformanceMetricsAverageAllOf) *NullableStorageBasePerformanceMetricsAverageAllOf {
	return &NullableStorageBasePerformanceMetricsAverageAllOf{value: val, isSet: true}
}

func (v NullableStorageBasePerformanceMetricsAverageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBasePerformanceMetricsAverageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
