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

// HyperflexReplicationPlatDatastorePairAllOf Definition of the list of properties defined in 'hyperflex.ReplicationPlatDatastorePair', excluding properties defined in parent classes.
type HyperflexReplicationPlatDatastorePairAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                    `json:"ObjectType"`
	Ads        NullableHyperflexReplicationPlatDatastore `json:"Ads,omitempty"`
	// Boolean representing if this is a backup only pair.
	BackupOnly *bool                                     `json:"BackupOnly,omitempty"`
	Bds        NullableHyperflexReplicationPlatDatastore `json:"Bds,omitempty"`
	// Boolean representing if this datastore pairing has quiesce snapshots enabled.
	Quiesce *bool `json:"Quiesce,omitempty"`
	// The replication interval for this pair in minutes.
	ReplicationIntervalInMinutes *int64                                    `json:"ReplicationIntervalInMinutes,omitempty"`
	Sourceds                     NullableHyperflexReplicationPlatDatastore `json:"Sourceds,omitempty"`
	// HyperFlex datastore pair is used for storage only.
	StorageOnly          *bool `json:"StorageOnly,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexReplicationPlatDatastorePairAllOf HyperflexReplicationPlatDatastorePairAllOf

// NewHyperflexReplicationPlatDatastorePairAllOf instantiates a new HyperflexReplicationPlatDatastorePairAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationPlatDatastorePairAllOf(classId string, objectType string) *HyperflexReplicationPlatDatastorePairAllOf {
	this := HyperflexReplicationPlatDatastorePairAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationPlatDatastorePairAllOfWithDefaults instantiates a new HyperflexReplicationPlatDatastorePairAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationPlatDatastorePairAllOfWithDefaults() *HyperflexReplicationPlatDatastorePairAllOf {
	this := HyperflexReplicationPlatDatastorePairAllOf{}
	var classId string = "hyperflex.ReplicationPlatDatastorePair"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationPlatDatastorePair"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAds returns the Ads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetAds() HyperflexReplicationPlatDatastore {
	if o == nil || o.Ads.Get() == nil {
		var ret HyperflexReplicationPlatDatastore
		return ret
	}
	return *o.Ads.Get()
}

// GetAdsOk returns a tuple with the Ads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetAdsOk() (*HyperflexReplicationPlatDatastore, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ads.Get(), o.Ads.IsSet()
}

// HasAds returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasAds() bool {
	if o != nil && o.Ads.IsSet() {
		return true
	}

	return false
}

// SetAds gets a reference to the given NullableHyperflexReplicationPlatDatastore and assigns it to the Ads field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetAds(v HyperflexReplicationPlatDatastore) {
	o.Ads.Set(&v)
}

// SetAdsNil sets the value for Ads to be an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetAdsNil() {
	o.Ads.Set(nil)
}

// UnsetAds ensures that no value is present for Ads, not even an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetAds() {
	o.Ads.Unset()
}

// GetBackupOnly returns the BackupOnly field value if set, zero value otherwise.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBackupOnly() bool {
	if o == nil || o.BackupOnly == nil {
		var ret bool
		return ret
	}
	return *o.BackupOnly
}

// GetBackupOnlyOk returns a tuple with the BackupOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBackupOnlyOk() (*bool, bool) {
	if o == nil || o.BackupOnly == nil {
		return nil, false
	}
	return o.BackupOnly, true
}

// HasBackupOnly returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasBackupOnly() bool {
	if o != nil && o.BackupOnly != nil {
		return true
	}

	return false
}

// SetBackupOnly gets a reference to the given bool and assigns it to the BackupOnly field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBackupOnly(v bool) {
	o.BackupOnly = &v
}

// GetBds returns the Bds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBds() HyperflexReplicationPlatDatastore {
	if o == nil || o.Bds.Get() == nil {
		var ret HyperflexReplicationPlatDatastore
		return ret
	}
	return *o.Bds.Get()
}

// GetBdsOk returns a tuple with the Bds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetBdsOk() (*HyperflexReplicationPlatDatastore, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bds.Get(), o.Bds.IsSet()
}

// HasBds returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasBds() bool {
	if o != nil && o.Bds.IsSet() {
		return true
	}

	return false
}

// SetBds gets a reference to the given NullableHyperflexReplicationPlatDatastore and assigns it to the Bds field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBds(v HyperflexReplicationPlatDatastore) {
	o.Bds.Set(&v)
}

// SetBdsNil sets the value for Bds to be an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetBdsNil() {
	o.Bds.Set(nil)
}

// UnsetBds ensures that no value is present for Bds, not even an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetBds() {
	o.Bds.Unset()
}

// GetQuiesce returns the Quiesce field value if set, zero value otherwise.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetQuiesce() bool {
	if o == nil || o.Quiesce == nil {
		var ret bool
		return ret
	}
	return *o.Quiesce
}

// GetQuiesceOk returns a tuple with the Quiesce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetQuiesceOk() (*bool, bool) {
	if o == nil || o.Quiesce == nil {
		return nil, false
	}
	return o.Quiesce, true
}

// HasQuiesce returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasQuiesce() bool {
	if o != nil && o.Quiesce != nil {
		return true
	}

	return false
}

// SetQuiesce gets a reference to the given bool and assigns it to the Quiesce field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetQuiesce(v bool) {
	o.Quiesce = &v
}

// GetReplicationIntervalInMinutes returns the ReplicationIntervalInMinutes field value if set, zero value otherwise.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetReplicationIntervalInMinutes() int64 {
	if o == nil || o.ReplicationIntervalInMinutes == nil {
		var ret int64
		return ret
	}
	return *o.ReplicationIntervalInMinutes
}

// GetReplicationIntervalInMinutesOk returns a tuple with the ReplicationIntervalInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetReplicationIntervalInMinutesOk() (*int64, bool) {
	if o == nil || o.ReplicationIntervalInMinutes == nil {
		return nil, false
	}
	return o.ReplicationIntervalInMinutes, true
}

// HasReplicationIntervalInMinutes returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasReplicationIntervalInMinutes() bool {
	if o != nil && o.ReplicationIntervalInMinutes != nil {
		return true
	}

	return false
}

// SetReplicationIntervalInMinutes gets a reference to the given int64 and assigns it to the ReplicationIntervalInMinutes field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetReplicationIntervalInMinutes(v int64) {
	o.ReplicationIntervalInMinutes = &v
}

// GetSourceds returns the Sourceds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetSourceds() HyperflexReplicationPlatDatastore {
	if o == nil || o.Sourceds.Get() == nil {
		var ret HyperflexReplicationPlatDatastore
		return ret
	}
	return *o.Sourceds.Get()
}

// GetSourcedsOk returns a tuple with the Sourceds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetSourcedsOk() (*HyperflexReplicationPlatDatastore, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sourceds.Get(), o.Sourceds.IsSet()
}

// HasSourceds returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasSourceds() bool {
	if o != nil && o.Sourceds.IsSet() {
		return true
	}

	return false
}

// SetSourceds gets a reference to the given NullableHyperflexReplicationPlatDatastore and assigns it to the Sourceds field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetSourceds(v HyperflexReplicationPlatDatastore) {
	o.Sourceds.Set(&v)
}

// SetSourcedsNil sets the value for Sourceds to be an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetSourcedsNil() {
	o.Sourceds.Set(nil)
}

// UnsetSourceds ensures that no value is present for Sourceds, not even an explicit nil
func (o *HyperflexReplicationPlatDatastorePairAllOf) UnsetSourceds() {
	o.Sourceds.Unset()
}

// GetStorageOnly returns the StorageOnly field value if set, zero value otherwise.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetStorageOnly() bool {
	if o == nil || o.StorageOnly == nil {
		var ret bool
		return ret
	}
	return *o.StorageOnly
}

// GetStorageOnlyOk returns a tuple with the StorageOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) GetStorageOnlyOk() (*bool, bool) {
	if o == nil || o.StorageOnly == nil {
		return nil, false
	}
	return o.StorageOnly, true
}

// HasStorageOnly returns a boolean if a field has been set.
func (o *HyperflexReplicationPlatDatastorePairAllOf) HasStorageOnly() bool {
	if o != nil && o.StorageOnly != nil {
		return true
	}

	return false
}

// SetStorageOnly gets a reference to the given bool and assigns it to the StorageOnly field.
func (o *HyperflexReplicationPlatDatastorePairAllOf) SetStorageOnly(v bool) {
	o.StorageOnly = &v
}

func (o HyperflexReplicationPlatDatastorePairAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Ads.IsSet() {
		toSerialize["Ads"] = o.Ads.Get()
	}
	if o.BackupOnly != nil {
		toSerialize["BackupOnly"] = o.BackupOnly
	}
	if o.Bds.IsSet() {
		toSerialize["Bds"] = o.Bds.Get()
	}
	if o.Quiesce != nil {
		toSerialize["Quiesce"] = o.Quiesce
	}
	if o.ReplicationIntervalInMinutes != nil {
		toSerialize["ReplicationIntervalInMinutes"] = o.ReplicationIntervalInMinutes
	}
	if o.Sourceds.IsSet() {
		toSerialize["Sourceds"] = o.Sourceds.Get()
	}
	if o.StorageOnly != nil {
		toSerialize["StorageOnly"] = o.StorageOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexReplicationPlatDatastorePairAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexReplicationPlatDatastorePairAllOf := _HyperflexReplicationPlatDatastorePairAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexReplicationPlatDatastorePairAllOf); err == nil {
		*o = HyperflexReplicationPlatDatastorePairAllOf(varHyperflexReplicationPlatDatastorePairAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ads")
		delete(additionalProperties, "BackupOnly")
		delete(additionalProperties, "Bds")
		delete(additionalProperties, "Quiesce")
		delete(additionalProperties, "ReplicationIntervalInMinutes")
		delete(additionalProperties, "Sourceds")
		delete(additionalProperties, "StorageOnly")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexReplicationPlatDatastorePairAllOf struct {
	value *HyperflexReplicationPlatDatastorePairAllOf
	isSet bool
}

func (v NullableHyperflexReplicationPlatDatastorePairAllOf) Get() *HyperflexReplicationPlatDatastorePairAllOf {
	return v.value
}

func (v *NullableHyperflexReplicationPlatDatastorePairAllOf) Set(val *HyperflexReplicationPlatDatastorePairAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationPlatDatastorePairAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationPlatDatastorePairAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationPlatDatastorePairAllOf(val *HyperflexReplicationPlatDatastorePairAllOf) *NullableHyperflexReplicationPlatDatastorePairAllOf {
	return &NullableHyperflexReplicationPlatDatastorePairAllOf{value: val, isSet: true}
}

func (v NullableHyperflexReplicationPlatDatastorePairAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationPlatDatastorePairAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
