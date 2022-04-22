/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageStorageUtilization Storage space utilized by Pure storage entity.
type StorageStorageUtilization struct {
	StorageBaseCapacity
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

type _StorageStorageUtilization StorageStorageUtilization

// NewStorageStorageUtilization instantiates a new StorageStorageUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStorageUtilization(classId string, objectType string) *StorageStorageUtilization {
	this := StorageStorageUtilization{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageStorageUtilizationWithDefaults instantiates a new StorageStorageUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStorageUtilizationWithDefaults() *StorageStorageUtilization {
	this := StorageStorageUtilization{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStorageUtilization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStorageUtilization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageStorageUtilization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStorageUtilization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDataReduction returns the DataReduction field value if set, zero value otherwise.
func (o *StorageStorageUtilization) GetDataReduction() float32 {
	if o == nil || o.DataReduction == nil {
		var ret float32
		return ret
	}
	return *o.DataReduction
}

// GetDataReductionOk returns a tuple with the DataReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetDataReductionOk() (*float32, bool) {
	if o == nil || o.DataReduction == nil {
		return nil, false
	}
	return o.DataReduction, true
}

// HasDataReduction returns a boolean if a field has been set.
func (o *StorageStorageUtilization) HasDataReduction() bool {
	if o != nil && o.DataReduction != nil {
		return true
	}

	return false
}

// SetDataReduction gets a reference to the given float32 and assigns it to the DataReduction field.
func (o *StorageStorageUtilization) SetDataReduction(v float32) {
	o.DataReduction = &v
}

// GetSnapshot returns the Snapshot field value if set, zero value otherwise.
func (o *StorageStorageUtilization) GetSnapshot() int64 {
	if o == nil || o.Snapshot == nil {
		var ret int64
		return ret
	}
	return *o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetSnapshotOk() (*int64, bool) {
	if o == nil || o.Snapshot == nil {
		return nil, false
	}
	return o.Snapshot, true
}

// HasSnapshot returns a boolean if a field has been set.
func (o *StorageStorageUtilization) HasSnapshot() bool {
	if o != nil && o.Snapshot != nil {
		return true
	}

	return false
}

// SetSnapshot gets a reference to the given int64 and assigns it to the Snapshot field.
func (o *StorageStorageUtilization) SetSnapshot(v int64) {
	o.Snapshot = &v
}

// GetThinProvisioned returns the ThinProvisioned field value if set, zero value otherwise.
func (o *StorageStorageUtilization) GetThinProvisioned() float32 {
	if o == nil || o.ThinProvisioned == nil {
		var ret float32
		return ret
	}
	return *o.ThinProvisioned
}

// GetThinProvisionedOk returns a tuple with the ThinProvisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetThinProvisionedOk() (*float32, bool) {
	if o == nil || o.ThinProvisioned == nil {
		return nil, false
	}
	return o.ThinProvisioned, true
}

// HasThinProvisioned returns a boolean if a field has been set.
func (o *StorageStorageUtilization) HasThinProvisioned() bool {
	if o != nil && o.ThinProvisioned != nil {
		return true
	}

	return false
}

// SetThinProvisioned gets a reference to the given float32 and assigns it to the ThinProvisioned field.
func (o *StorageStorageUtilization) SetThinProvisioned(v float32) {
	o.ThinProvisioned = &v
}

// GetTotalReduction returns the TotalReduction field value if set, zero value otherwise.
func (o *StorageStorageUtilization) GetTotalReduction() float32 {
	if o == nil || o.TotalReduction == nil {
		var ret float32
		return ret
	}
	return *o.TotalReduction
}

// GetTotalReductionOk returns a tuple with the TotalReduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetTotalReductionOk() (*float32, bool) {
	if o == nil || o.TotalReduction == nil {
		return nil, false
	}
	return o.TotalReduction, true
}

// HasTotalReduction returns a boolean if a field has been set.
func (o *StorageStorageUtilization) HasTotalReduction() bool {
	if o != nil && o.TotalReduction != nil {
		return true
	}

	return false
}

// SetTotalReduction gets a reference to the given float32 and assigns it to the TotalReduction field.
func (o *StorageStorageUtilization) SetTotalReduction(v float32) {
	o.TotalReduction = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *StorageStorageUtilization) GetVolume() int64 {
	if o == nil || o.Volume == nil {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStorageUtilization) GetVolumeOk() (*int64, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *StorageStorageUtilization) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *StorageStorageUtilization) SetVolume(v int64) {
	o.Volume = &v
}

func (o StorageStorageUtilization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseCapacity, errStorageBaseCapacity := json.Marshal(o.StorageBaseCapacity)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
	errStorageBaseCapacity = json.Unmarshal([]byte(serializedStorageBaseCapacity), &toSerialize)
	if errStorageBaseCapacity != nil {
		return []byte{}, errStorageBaseCapacity
	}
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

func (o *StorageStorageUtilization) UnmarshalJSON(bytes []byte) (err error) {
	type StorageStorageUtilizationWithoutEmbeddedStruct struct {
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
		Volume *int64 `json:"Volume,omitempty"`
	}

	varStorageStorageUtilizationWithoutEmbeddedStruct := StorageStorageUtilizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageStorageUtilizationWithoutEmbeddedStruct)
	if err == nil {
		varStorageStorageUtilization := _StorageStorageUtilization{}
		varStorageStorageUtilization.ClassId = varStorageStorageUtilizationWithoutEmbeddedStruct.ClassId
		varStorageStorageUtilization.ObjectType = varStorageStorageUtilizationWithoutEmbeddedStruct.ObjectType
		varStorageStorageUtilization.DataReduction = varStorageStorageUtilizationWithoutEmbeddedStruct.DataReduction
		varStorageStorageUtilization.Snapshot = varStorageStorageUtilizationWithoutEmbeddedStruct.Snapshot
		varStorageStorageUtilization.ThinProvisioned = varStorageStorageUtilizationWithoutEmbeddedStruct.ThinProvisioned
		varStorageStorageUtilization.TotalReduction = varStorageStorageUtilizationWithoutEmbeddedStruct.TotalReduction
		varStorageStorageUtilization.Volume = varStorageStorageUtilizationWithoutEmbeddedStruct.Volume
		*o = StorageStorageUtilization(varStorageStorageUtilization)
	} else {
		return err
	}

	varStorageStorageUtilization := _StorageStorageUtilization{}

	err = json.Unmarshal(bytes, &varStorageStorageUtilization)
	if err == nil {
		o.StorageBaseCapacity = varStorageStorageUtilization.StorageBaseCapacity
	} else {
		return err
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

		// remove fields from embedded structs
		reflectStorageBaseCapacity := reflect.ValueOf(o.StorageBaseCapacity)
		for i := 0; i < reflectStorageBaseCapacity.Type().NumField(); i++ {
			t := reflectStorageBaseCapacity.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageStorageUtilization struct {
	value *StorageStorageUtilization
	isSet bool
}

func (v NullableStorageStorageUtilization) Get() *StorageStorageUtilization {
	return v.value
}

func (v *NullableStorageStorageUtilization) Set(val *StorageStorageUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStorageUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStorageUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStorageUtilization(val *StorageStorageUtilization) *NullableStorageStorageUtilization {
	return &NullableStorageStorageUtilization{value: val, isSet: true}
}

func (v NullableStorageStorageUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStorageUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
