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

// RecommendationCapacityRunwayAllOf Definition of the list of properties defined in 'recommendation.CapacityRunway', excluding properties defined in parent classes.
type RecommendationCapacityRunwayAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Additional capacity is the capacity which is needed more after exhausing all hardware on current cluster.
	AdditionalCapacity *int64 `json:"AdditionalCapacity,omitempty"`
	// Number of months in future for which recommendation is provided for.
	Period *int64 `json:"Period,omitempty"`
	// This represents the new runway, that is the number of days remaining before the cluster's storage utilization reaches the recommended capacity limit after the recommended hardware is added.
	Runway *int64 `json:"Runway,omitempty"`
	// Total capacity of the cluster after the recommended hardware is added.
	TotalCapacity *int64 `json:"TotalCapacity,omitempty"`
	// Unit for the new capacity. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes.
	Unit             *string                       `json:"Unit,omitempty"`
	ForecastInstance *ForecastInstanceRelationship `json:"ForecastInstance,omitempty"`
	// An array of relationships to recommendationPhysicalItem resources.
	PhysicalItem         []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationCapacityRunwayAllOf RecommendationCapacityRunwayAllOf

// NewRecommendationCapacityRunwayAllOf instantiates a new RecommendationCapacityRunwayAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationCapacityRunwayAllOf(classId string, objectType string) *RecommendationCapacityRunwayAllOf {
	this := RecommendationCapacityRunwayAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationCapacityRunwayAllOfWithDefaults instantiates a new RecommendationCapacityRunwayAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationCapacityRunwayAllOfWithDefaults() *RecommendationCapacityRunwayAllOf {
	this := RecommendationCapacityRunwayAllOf{}
	var classId string = "recommendation.CapacityRunway"
	this.ClassId = classId
	var objectType string = "recommendation.CapacityRunway"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationCapacityRunwayAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationCapacityRunwayAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationCapacityRunwayAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationCapacityRunwayAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalCapacity returns the AdditionalCapacity field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetAdditionalCapacity() int64 {
	if o == nil || o.AdditionalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.AdditionalCapacity
}

// GetAdditionalCapacityOk returns a tuple with the AdditionalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetAdditionalCapacityOk() (*int64, bool) {
	if o == nil || o.AdditionalCapacity == nil {
		return nil, false
	}
	return o.AdditionalCapacity, true
}

// HasAdditionalCapacity returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasAdditionalCapacity() bool {
	if o != nil && o.AdditionalCapacity != nil {
		return true
	}

	return false
}

// SetAdditionalCapacity gets a reference to the given int64 and assigns it to the AdditionalCapacity field.
func (o *RecommendationCapacityRunwayAllOf) SetAdditionalCapacity(v int64) {
	o.AdditionalCapacity = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetPeriod() int64 {
	if o == nil || o.Period == nil {
		var ret int64
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetPeriodOk() (*int64, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int64 and assigns it to the Period field.
func (o *RecommendationCapacityRunwayAllOf) SetPeriod(v int64) {
	o.Period = &v
}

// GetRunway returns the Runway field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetRunway() int64 {
	if o == nil || o.Runway == nil {
		var ret int64
		return ret
	}
	return *o.Runway
}

// GetRunwayOk returns a tuple with the Runway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetRunwayOk() (*int64, bool) {
	if o == nil || o.Runway == nil {
		return nil, false
	}
	return o.Runway, true
}

// HasRunway returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasRunway() bool {
	if o != nil && o.Runway != nil {
		return true
	}

	return false
}

// SetRunway gets a reference to the given int64 and assigns it to the Runway field.
func (o *RecommendationCapacityRunwayAllOf) SetRunway(v int64) {
	o.Runway = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetTotalCapacity() int64 {
	if o == nil || o.TotalCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetTotalCapacityOk() (*int64, bool) {
	if o == nil || o.TotalCapacity == nil {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasTotalCapacity() bool {
	if o != nil && o.TotalCapacity != nil {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given int64 and assigns it to the TotalCapacity field.
func (o *RecommendationCapacityRunwayAllOf) SetTotalCapacity(v int64) {
	o.TotalCapacity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *RecommendationCapacityRunwayAllOf) SetUnit(v string) {
	o.Unit = &v
}

// GetForecastInstance returns the ForecastInstance field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetForecastInstance() ForecastInstanceRelationship {
	if o == nil || o.ForecastInstance == nil {
		var ret ForecastInstanceRelationship
		return ret
	}
	return *o.ForecastInstance
}

// GetForecastInstanceOk returns a tuple with the ForecastInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetForecastInstanceOk() (*ForecastInstanceRelationship, bool) {
	if o == nil || o.ForecastInstance == nil {
		return nil, false
	}
	return o.ForecastInstance, true
}

// HasForecastInstance returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasForecastInstance() bool {
	if o != nil && o.ForecastInstance != nil {
		return true
	}

	return false
}

// SetForecastInstance gets a reference to the given ForecastInstanceRelationship and assigns it to the ForecastInstance field.
func (o *RecommendationCapacityRunwayAllOf) SetForecastInstance(v ForecastInstanceRelationship) {
	o.ForecastInstance = &v
}

// GetPhysicalItem returns the PhysicalItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationCapacityRunwayAllOf) GetPhysicalItem() []RecommendationPhysicalItemRelationship {
	if o == nil {
		var ret []RecommendationPhysicalItemRelationship
		return ret
	}
	return o.PhysicalItem
}

// GetPhysicalItemOk returns a tuple with the PhysicalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationCapacityRunwayAllOf) GetPhysicalItemOk() (*[]RecommendationPhysicalItemRelationship, bool) {
	if o == nil || o.PhysicalItem == nil {
		return nil, false
	}
	return &o.PhysicalItem, true
}

// HasPhysicalItem returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasPhysicalItem() bool {
	if o != nil && o.PhysicalItem != nil {
		return true
	}

	return false
}

// SetPhysicalItem gets a reference to the given []RecommendationPhysicalItemRelationship and assigns it to the PhysicalItem field.
func (o *RecommendationCapacityRunwayAllOf) SetPhysicalItem(v []RecommendationPhysicalItemRelationship) {
	o.PhysicalItem = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecommendationCapacityRunwayAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationCapacityRunwayAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecommendationCapacityRunwayAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecommendationCapacityRunwayAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o RecommendationCapacityRunwayAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdditionalCapacity != nil {
		toSerialize["AdditionalCapacity"] = o.AdditionalCapacity
	}
	if o.Period != nil {
		toSerialize["Period"] = o.Period
	}
	if o.Runway != nil {
		toSerialize["Runway"] = o.Runway
	}
	if o.TotalCapacity != nil {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}
	if o.ForecastInstance != nil {
		toSerialize["ForecastInstance"] = o.ForecastInstance
	}
	if o.PhysicalItem != nil {
		toSerialize["PhysicalItem"] = o.PhysicalItem
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationCapacityRunwayAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationCapacityRunwayAllOf := _RecommendationCapacityRunwayAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationCapacityRunwayAllOf); err == nil {
		*o = RecommendationCapacityRunwayAllOf(varRecommendationCapacityRunwayAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalCapacity")
		delete(additionalProperties, "Period")
		delete(additionalProperties, "Runway")
		delete(additionalProperties, "TotalCapacity")
		delete(additionalProperties, "Unit")
		delete(additionalProperties, "ForecastInstance")
		delete(additionalProperties, "PhysicalItem")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationCapacityRunwayAllOf struct {
	value *RecommendationCapacityRunwayAllOf
	isSet bool
}

func (v NullableRecommendationCapacityRunwayAllOf) Get() *RecommendationCapacityRunwayAllOf {
	return v.value
}

func (v *NullableRecommendationCapacityRunwayAllOf) Set(val *RecommendationCapacityRunwayAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationCapacityRunwayAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationCapacityRunwayAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationCapacityRunwayAllOf(val *RecommendationCapacityRunwayAllOf) *NullableRecommendationCapacityRunwayAllOf {
	return &NullableRecommendationCapacityRunwayAllOf{value: val, isSet: true}
}

func (v NullableRecommendationCapacityRunwayAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationCapacityRunwayAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
