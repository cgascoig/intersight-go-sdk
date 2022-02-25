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

// StorageBaseSnapshotScheduleAllOf Definition of the list of properties defined in 'storage.BaseSnapshotSchedule', excluding properties defined in parent classes.
type StorageBaseSnapshotScheduleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Snapshot frequency. It is an interval at which snapshot is set to trigger on source array. Examples:     PT2H Snapshot is generated every 2 hours.     P4D, Snapshot is scheduled for every 4 days.     PT2H34M56.123S is 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	Frequency *string `json:"Frequency,omitempty"`
	// Name of the snapshot schedule.
	Name *string `json:"Name,omitempty"`
	// Duration to keep the snapshots on the source array. Once this period expires, system deletes the snapshot automatically from the source array. Examples: P200D,  200 days. PT2H34M56.123S, 2 hours, 34 minutes, 56 seconds and 123 milliseconds.
	RetentionTime *string `json:"RetentionTime,omitempty"`
	// Preferred time of the day to capture the snapshot. It is applicable only if the frequency is set for a day or more. Format: hh:mm:ss Example: 08:30:00, Snapshot is set for 08:30 AM.
	SnapshotTime         *string `json:"SnapshotTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseSnapshotScheduleAllOf StorageBaseSnapshotScheduleAllOf

// NewStorageBaseSnapshotScheduleAllOf instantiates a new StorageBaseSnapshotScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseSnapshotScheduleAllOf(classId string, objectType string) *StorageBaseSnapshotScheduleAllOf {
	this := StorageBaseSnapshotScheduleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseSnapshotScheduleAllOfWithDefaults instantiates a new StorageBaseSnapshotScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseSnapshotScheduleAllOfWithDefaults() *StorageBaseSnapshotScheduleAllOf {
	this := StorageBaseSnapshotScheduleAllOf{}
	var classId string = "storage.PureSnapshotSchedule"
	this.ClassId = classId
	var objectType string = "storage.PureSnapshotSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseSnapshotScheduleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseSnapshotScheduleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseSnapshotScheduleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseSnapshotScheduleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *StorageBaseSnapshotScheduleAllOf) GetFrequency() string {
	if o == nil || o.Frequency == nil {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetFrequencyOk() (*string, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *StorageBaseSnapshotScheduleAllOf) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *StorageBaseSnapshotScheduleAllOf) SetFrequency(v string) {
	o.Frequency = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseSnapshotScheduleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseSnapshotScheduleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseSnapshotScheduleAllOf) SetName(v string) {
	o.Name = &v
}

// GetRetentionTime returns the RetentionTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshotScheduleAllOf) GetRetentionTime() string {
	if o == nil || o.RetentionTime == nil {
		var ret string
		return ret
	}
	return *o.RetentionTime
}

// GetRetentionTimeOk returns a tuple with the RetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetRetentionTimeOk() (*string, bool) {
	if o == nil || o.RetentionTime == nil {
		return nil, false
	}
	return o.RetentionTime, true
}

// HasRetentionTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshotScheduleAllOf) HasRetentionTime() bool {
	if o != nil && o.RetentionTime != nil {
		return true
	}

	return false
}

// SetRetentionTime gets a reference to the given string and assigns it to the RetentionTime field.
func (o *StorageBaseSnapshotScheduleAllOf) SetRetentionTime(v string) {
	o.RetentionTime = &v
}

// GetSnapshotTime returns the SnapshotTime field value if set, zero value otherwise.
func (o *StorageBaseSnapshotScheduleAllOf) GetSnapshotTime() string {
	if o == nil || o.SnapshotTime == nil {
		var ret string
		return ret
	}
	return *o.SnapshotTime
}

// GetSnapshotTimeOk returns a tuple with the SnapshotTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseSnapshotScheduleAllOf) GetSnapshotTimeOk() (*string, bool) {
	if o == nil || o.SnapshotTime == nil {
		return nil, false
	}
	return o.SnapshotTime, true
}

// HasSnapshotTime returns a boolean if a field has been set.
func (o *StorageBaseSnapshotScheduleAllOf) HasSnapshotTime() bool {
	if o != nil && o.SnapshotTime != nil {
		return true
	}

	return false
}

// SetSnapshotTime gets a reference to the given string and assigns it to the SnapshotTime field.
func (o *StorageBaseSnapshotScheduleAllOf) SetSnapshotTime(v string) {
	o.SnapshotTime = &v
}

func (o StorageBaseSnapshotScheduleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Frequency != nil {
		toSerialize["Frequency"] = o.Frequency
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.RetentionTime != nil {
		toSerialize["RetentionTime"] = o.RetentionTime
	}
	if o.SnapshotTime != nil {
		toSerialize["SnapshotTime"] = o.SnapshotTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseSnapshotScheduleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageBaseSnapshotScheduleAllOf := _StorageBaseSnapshotScheduleAllOf{}

	if err = json.Unmarshal(bytes, &varStorageBaseSnapshotScheduleAllOf); err == nil {
		*o = StorageBaseSnapshotScheduleAllOf(varStorageBaseSnapshotScheduleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Frequency")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RetentionTime")
		delete(additionalProperties, "SnapshotTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageBaseSnapshotScheduleAllOf struct {
	value *StorageBaseSnapshotScheduleAllOf
	isSet bool
}

func (v NullableStorageBaseSnapshotScheduleAllOf) Get() *StorageBaseSnapshotScheduleAllOf {
	return v.value
}

func (v *NullableStorageBaseSnapshotScheduleAllOf) Set(val *StorageBaseSnapshotScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseSnapshotScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseSnapshotScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseSnapshotScheduleAllOf(val *StorageBaseSnapshotScheduleAllOf) *NullableStorageBaseSnapshotScheduleAllOf {
	return &NullableStorageBaseSnapshotScheduleAllOf{value: val, isSet: true}
}

func (v NullableStorageBaseSnapshotScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseSnapshotScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
