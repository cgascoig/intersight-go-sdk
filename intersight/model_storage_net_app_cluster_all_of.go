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
)

// StorageNetAppClusterAllOf Definition of the list of properties defined in 'storage.NetAppCluster', excluding properties defined in parent classes.
type StorageNetAppClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                        `json:"ObjectType"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage       `json:"AvgPerformanceMetrics,omitempty"`
	ClusterEfficiency     NullableStorageNetAppStorageClusterEfficiency `json:"ClusterEfficiency,omitempty"`
	// The health status of the cluster. Possible states are ok, ok-with-suppressed, degraded, and unreachable. * `Unreachable` - Cluster status is unreachable. * `OK` - Cluster status is either ok or ok-with-suppressed. * `Degraded` - Cluster status is degraded.
	ClusterHealthStatus *string  `json:"ClusterHealthStatus,omitempty"`
	DnsDomains          []string `json:"DnsDomains,omitempty"`
	// Unique identifier of NetApp Cluster across data center.
	Key *string `json:"Key,omitempty"`
	// Location of the storage controller.
	Location *string `json:"Location,omitempty"`
	// FQDN or IP Address of Storage Cluster.
	ManagementAddress *string  `json:"ManagementAddress,omitempty"`
	NameServers       []string `json:"NameServers,omitempty"`
	NtpServers        []string `json:"NtpServers,omitempty"`
	// An array of relationships to storageNetAppClusterEvent resources.
	Events               []StorageNetAppClusterEventRelationship `json:"Events,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterAllOf StorageNetAppClusterAllOf

// NewStorageNetAppClusterAllOf instantiates a new StorageNetAppClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterAllOf(classId string, objectType string) *StorageNetAppClusterAllOf {
	this := StorageNetAppClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterAllOfWithDefaults instantiates a new StorageNetAppClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterAllOfWithDefaults() *StorageNetAppClusterAllOf {
	this := StorageNetAppClusterAllOf{}
	var classId string = "storage.NetAppCluster"
	this.ClassId = classId
	var objectType string = "storage.NetAppCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppClusterAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetClusterEfficiency returns the ClusterEfficiency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterAllOf) GetClusterEfficiency() StorageNetAppStorageClusterEfficiency {
	if o == nil || o.ClusterEfficiency.Get() == nil {
		var ret StorageNetAppStorageClusterEfficiency
		return ret
	}
	return *o.ClusterEfficiency.Get()
}

// GetClusterEfficiencyOk returns a tuple with the ClusterEfficiency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterAllOf) GetClusterEfficiencyOk() (*StorageNetAppStorageClusterEfficiency, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterEfficiency.Get(), o.ClusterEfficiency.IsSet()
}

// HasClusterEfficiency returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasClusterEfficiency() bool {
	if o != nil && o.ClusterEfficiency.IsSet() {
		return true
	}

	return false
}

// SetClusterEfficiency gets a reference to the given NullableStorageNetAppStorageClusterEfficiency and assigns it to the ClusterEfficiency field.
func (o *StorageNetAppClusterAllOf) SetClusterEfficiency(v StorageNetAppStorageClusterEfficiency) {
	o.ClusterEfficiency.Set(&v)
}

// SetClusterEfficiencyNil sets the value for ClusterEfficiency to be an explicit nil
func (o *StorageNetAppClusterAllOf) SetClusterEfficiencyNil() {
	o.ClusterEfficiency.Set(nil)
}

// UnsetClusterEfficiency ensures that no value is present for ClusterEfficiency, not even an explicit nil
func (o *StorageNetAppClusterAllOf) UnsetClusterEfficiency() {
	o.ClusterEfficiency.Unset()
}

// GetClusterHealthStatus returns the ClusterHealthStatus field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetClusterHealthStatus() string {
	if o == nil || o.ClusterHealthStatus == nil {
		var ret string
		return ret
	}
	return *o.ClusterHealthStatus
}

// GetClusterHealthStatusOk returns a tuple with the ClusterHealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetClusterHealthStatusOk() (*string, bool) {
	if o == nil || o.ClusterHealthStatus == nil {
		return nil, false
	}
	return o.ClusterHealthStatus, true
}

// HasClusterHealthStatus returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasClusterHealthStatus() bool {
	if o != nil && o.ClusterHealthStatus != nil {
		return true
	}

	return false
}

// SetClusterHealthStatus gets a reference to the given string and assigns it to the ClusterHealthStatus field.
func (o *StorageNetAppClusterAllOf) SetClusterHealthStatus(v string) {
	o.ClusterHealthStatus = &v
}

// GetDnsDomains returns the DnsDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterAllOf) GetDnsDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsDomains
}

// GetDnsDomainsOk returns a tuple with the DnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterAllOf) GetDnsDomainsOk() (*[]string, bool) {
	if o == nil || o.DnsDomains == nil {
		return nil, false
	}
	return &o.DnsDomains, true
}

// HasDnsDomains returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasDnsDomains() bool {
	if o != nil && o.DnsDomains != nil {
		return true
	}

	return false
}

// SetDnsDomains gets a reference to the given []string and assigns it to the DnsDomains field.
func (o *StorageNetAppClusterAllOf) SetDnsDomains(v []string) {
	o.DnsDomains = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppClusterAllOf) SetKey(v string) {
	o.Key = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *StorageNetAppClusterAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *StorageNetAppClusterAllOf) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterAllOf) GetNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterAllOf) GetNameServersOk() (*[]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return &o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *StorageNetAppClusterAllOf) SetNameServers(v []string) {
	o.NameServers = v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterAllOf) GetNtpServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterAllOf) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return &o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *StorageNetAppClusterAllOf) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterAllOf) GetEvents() []StorageNetAppClusterEventRelationship {
	if o == nil {
		var ret []StorageNetAppClusterEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterAllOf) GetEventsOk() (*[]StorageNetAppClusterEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppClusterEventRelationship and assigns it to the Events field.
func (o *StorageNetAppClusterAllOf) SetEvents(v []StorageNetAppClusterEventRelationship) {
	o.Events = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageNetAppClusterAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageNetAppClusterAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageNetAppClusterAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageNetAppClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.ClusterEfficiency.IsSet() {
		toSerialize["ClusterEfficiency"] = o.ClusterEfficiency.Get()
	}
	if o.ClusterHealthStatus != nil {
		toSerialize["ClusterHealthStatus"] = o.ClusterHealthStatus
	}
	if o.DnsDomains != nil {
		toSerialize["DnsDomains"] = o.DnsDomains
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.ManagementAddress != nil {
		toSerialize["ManagementAddress"] = o.ManagementAddress
	}
	if o.NameServers != nil {
		toSerialize["NameServers"] = o.NameServers
	}
	if o.NtpServers != nil {
		toSerialize["NtpServers"] = o.NtpServers
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppClusterAllOf := _StorageNetAppClusterAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppClusterAllOf); err == nil {
		*o = StorageNetAppClusterAllOf(varStorageNetAppClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "ClusterEfficiency")
		delete(additionalProperties, "ClusterHealthStatus")
		delete(additionalProperties, "DnsDomains")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "NameServers")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppClusterAllOf struct {
	value *StorageNetAppClusterAllOf
	isSet bool
}

func (v NullableStorageNetAppClusterAllOf) Get() *StorageNetAppClusterAllOf {
	return v.value
}

func (v *NullableStorageNetAppClusterAllOf) Set(val *StorageNetAppClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterAllOf(val *StorageNetAppClusterAllOf) *NullableStorageNetAppClusterAllOf {
	return &NullableStorageNetAppClusterAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
