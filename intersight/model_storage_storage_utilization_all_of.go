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

// StorageStorageUtilizationAllOf Definition of the list of properties defined in 'storage.StorageUtilization', excluding properties defined in parent classes.
type StorageStorageUtilizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5.0 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules.
	DataReduction *float32 `json:"DataReduction,omitempty"`
	// Physical space occupied by the snapshots, represented in bytes.
	Snapshot *int64 `json:"Snapshot,omitempty"`
	// Percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed.
	ThinProvisioned *float32 `json:"ThinProvisioned,omitempty"`
	// Ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10.0 means that for every 10 MB of provisioned space, 1 MB is stored on the array's flash modules.
	TotalReduction *float32 `json:"TotalReduction,omitempty"`
	// Physical space occupied by volume data, excluding shared, array metadata and snapshots. Size id represented in bytes.
	Volume               *int64 `json:"Volume,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageStorageUtilizationAllOf StorageStorageUtilizationAllOf

// NewStorageStorageUtilizationAllOf instantiates a new StorageStorageUtilizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStorageUtilizationAllOf(classId string, objectType string) *StorageStorageUtilizationAllOf {
	this := StorageStorageUtilizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageStorageUtilizationAllOfWithDefaults instantiates a new StorageStorageUtilizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStorageUtilizationAllOfWithDefaults() *StorageStorageUtilizationAllOf {
	this := StorageStorageUtilizationAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStorageUtilizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStorageUtilizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageStorageUtilizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStorageUtilizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReduction returns the DataReduction field value if set, zero value otherwise.
func (o *StorageStorageUtilizationAllOf) GetDataReduction() float32 {
	if o == nil || o.DataReduction == nil {
		var ret float32
		return ret
	}
	return *o.DataReduction
}

// GetDataReductionOk returns a tuple with the DataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetDataReductionOk() (*float32, bool) {
	if o == nil || o.DataReduction == nil {
		return nil, false
	}
	return o.DataReduction, true
}

// HasDataReduction returns a boolean if a field has been set.
func (o *StorageStorageUtilizationAllOf) HasDataReduction() bool {
	if o != nil && o.DataReduction != nil {
		return true
	}

	return false
}

// SetDataReduction gets a reference to the given float32 and assigns it to the DataReduction field.
func (o *StorageStorageUtilizationAllOf) SetDataReduction(v float32) {
	o.DataReduction = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *StorageStorageUtilizationAllOf) GetSnapshot() int64 {
	if o == nil || o.Snapshot == nil {
		var ret int64
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetSnapshotOk() (*int64, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *StorageStorageUtilizationAllOf) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given int64 and assigns it to the Snapshot field.
func (o *StorageStorageUtilizationAllOf) SetSnapshot(v int64) {
	o.Snapshot = &v
}

// GetThinProvisioned returns the ThinProvisioned field value if set, zero value otherwise.
func (o *StorageStorageUtilizationAllOf) GetThinProvisioned() float32 {
	if o == nil || o.ThinProvisioned == nil {
		var ret float32
		return ret
	}
	return *o.ThinProvisioned
}

// GetThinProvisionedOk returns a tuple with the ThinProvisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetThinProvisionedOk() (*float32, bool) {
	if o == nil || o.ThinProvisioned == nil {
		return nil, false
	}
	return o.ThinProvisioned, true
}

// HasThinProvisioned returns a boolean if a field has been set.
func (o *StorageStorageUtilizationAllOf) HasThinProvisioned() bool {
	if o != nil && o.ThinProvisioned != nil {
		return true
	}

	return false
}

// SetThinProvisioned gets a reference to the given float32 and assigns it to the ThinProvisioned field.
func (o *StorageStorageUtilizationAllOf) SetThinProvisioned(v float32) {
	o.ThinProvisioned = &v
}

// GetTotalReduction returns the TotalReduction field value if set, zero value otherwise.
func (o *StorageStorageUtilizationAllOf) GetTotalReduction() float32 {
	if o == nil || o.TotalReduction == nil {
		var ret float32
		return ret
	}
	return *o.TotalReduction
}

// GetTotalReductionOk returns a tuple with the TotalReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetTotalReductionOk() (*float32, bool) {
	if o == nil || o.TotalReduction == nil {
		return nil, false
	}
	return o.TotalReduction, true
}

// HasTotalReduction returns a boolean if a field has been set.
func (o *StorageStorageUtilizationAllOf) HasTotalReduction() bool {
	if o != nil && o.TotalReduction != nil {
		return true
	}

	return false
}

// SetTotalReduction gets a reference to the given float32 and assigns it to the TotalReduction field.
func (o *StorageStorageUtilizationAllOf) SetTotalReduction(v float32) {
	o.TotalReduction = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageStorageUtilizationAllOf) GetVolume() int64 {
	if o == nil || o.Volume == nil {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilizationAllOf) GetVolumeOk() (*int64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageStorageUtilizationAllOf) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *StorageStorageUtilizationAllOf) SetVolume(v int64) {
	o.Volume = &v
}

func (o StorageStorageUtilizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DataReduction != nil {
		toSerialize["DataReduction"] = o.DataReduction
	}
	if o.Snapshot != nil {
		toSerialize["Snapshot"] = o.Snapshot
	}
	if o.ThinProvisioned != nil {
		toSerialize["ThinProvisioned"] = o.ThinProvisioned
	}
	if o.TotalReduction != nil {
		toSerialize["TotalReduction"] = o.TotalReduction
	}
	if o.Volume != nil {
		toSerialize["Volume"] = o.Volume
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageStorageUtilizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageStorageUtilizationAllOf := _StorageStorageUtilizationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageStorageUtilizationAllOf); err == nil {
		*o = StorageStorageUtilizationAllOf(varStorageStorageUtilizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataReduction")
		delete(additionalProperties, "Snapshot")
		delete(additionalProperties, "ThinProvisioned")
		delete(additionalProperties, "TotalReduction")
		delete(additionalProperties, "Volume")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageStorageUtilizationAllOf struct {
	value *StorageStorageUtilizationAllOf
	isSet bool
}

func (v NullableStorageStorageUtilizationAllOf) Get() *StorageStorageUtilizationAllOf {
	return v.value
}

func (v *NullableStorageStorageUtilizationAllOf) Set(val *StorageStorageUtilizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStorageUtilizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStorageUtilizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStorageUtilizationAllOf(val *StorageStorageUtilizationAllOf) *NullableStorageStorageUtilizationAllOf {
	return &NullableStorageStorageUtilizationAllOf{value: val, isSet: true}
}

func (v NullableStorageStorageUtilizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStorageUtilizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
