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
	"reflect"
	"strings"
)

// StorageNetAppCluster NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.
type StorageNetAppCluster struct {
	StorageBaseArray
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                        `json:"ObjectType"`
	AutoSupport           NullableStorageNetAppAutoSupport              `json:"AutoSupport,omitempty"`
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
	// Indicates whether or not telnet is enabled on the cluster.
	TelnetEnabled *bool `json:"TelnetEnabled,omitempty"`
	// An array of relationships to storageNetAppClusterEvent resources.
	Events               []StorageNetAppClusterEventRelationship `json:"Events,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppCluster StorageNetAppCluster

// NewStorageNetAppCluster instantiates a new StorageNetAppCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppCluster(classId string, objectType string) *StorageNetAppCluster {
	this := StorageNetAppCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterWithDefaults instantiates a new StorageNetAppCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterWithDefaults() *StorageNetAppCluster {
	this := StorageNetAppCluster{}
	var classId string = "storage.NetAppCluster"
	this.ClassId = classId
	var objectType string = "storage.NetAppCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoSupport returns the AutoSupport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetAutoSupport() StorageNetAppAutoSupport {
	if o == nil || o.AutoSupport.Get() == nil {
		var ret StorageNetAppAutoSupport
		return ret
	}
	return *o.AutoSupport.Get()
}

// GetAutoSupportOk returns a tuple with the AutoSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetAutoSupportOk() (*StorageNetAppAutoSupport, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoSupport.Get(), o.AutoSupport.IsSet()
}

// HasAutoSupport returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasAutoSupport() bool {
	if o != nil && o.AutoSupport.IsSet() {
		return true
	}

	return false
}

// SetAutoSupport gets a reference to the given NullableStorageNetAppAutoSupport and assigns it to the AutoSupport field.
func (o *StorageNetAppCluster) SetAutoSupport(v StorageNetAppAutoSupport) {
	o.AutoSupport.Set(&v)
}

// SetAutoSupportNil sets the value for AutoSupport to be an explicit nil
func (o *StorageNetAppCluster) SetAutoSupportNil() {
	o.AutoSupport.Set(nil)
}

// UnsetAutoSupport ensures that no value is present for AutoSupport, not even an explicit nil
func (o *StorageNetAppCluster) UnsetAutoSupport() {
	o.AutoSupport.Unset()
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppCluster) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetClusterEfficiency returns the ClusterEfficiency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetClusterEfficiency() StorageNetAppStorageClusterEfficiency {
	if o == nil || o.ClusterEfficiency.Get() == nil {
		var ret StorageNetAppStorageClusterEfficiency
		return ret
	}
	return *o.ClusterEfficiency.Get()
}

// GetClusterEfficiencyOk returns a tuple with the ClusterEfficiency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetClusterEfficiencyOk() (*StorageNetAppStorageClusterEfficiency, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterEfficiency.Get(), o.ClusterEfficiency.IsSet()
}

// HasClusterEfficiency returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasClusterEfficiency() bool {
	if o != nil && o.ClusterEfficiency.IsSet() {
		return true
	}

	return false
}

// SetClusterEfficiency gets a reference to the given NullableStorageNetAppStorageClusterEfficiency and assigns it to the ClusterEfficiency field.
func (o *StorageNetAppCluster) SetClusterEfficiency(v StorageNetAppStorageClusterEfficiency) {
	o.ClusterEfficiency.Set(&v)
}

// SetClusterEfficiencyNil sets the value for ClusterEfficiency to be an explicit nil
func (o *StorageNetAppCluster) SetClusterEfficiencyNil() {
	o.ClusterEfficiency.Set(nil)
}

// UnsetClusterEfficiency ensures that no value is present for ClusterEfficiency, not even an explicit nil
func (o *StorageNetAppCluster) UnsetClusterEfficiency() {
	o.ClusterEfficiency.Unset()
}

// GetClusterHealthStatus returns the ClusterHealthStatus field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetClusterHealthStatus() string {
	if o == nil || o.ClusterHealthStatus == nil {
		var ret string
		return ret
	}
	return *o.ClusterHealthStatus
}

// GetClusterHealthStatusOk returns a tuple with the ClusterHealthStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetClusterHealthStatusOk() (*string, bool) {
	if o == nil || o.ClusterHealthStatus == nil {
		return nil, false
	}
	return o.ClusterHealthStatus, true
}

// HasClusterHealthStatus returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasClusterHealthStatus() bool {
	if o != nil && o.ClusterHealthStatus != nil {
		return true
	}

	return false
}

// SetClusterHealthStatus gets a reference to the given string and assigns it to the ClusterHealthStatus field.
func (o *StorageNetAppCluster) SetClusterHealthStatus(v string) {
	o.ClusterHealthStatus = &v
}

// GetDnsDomains returns the DnsDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetDnsDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DnsDomains
}

// GetDnsDomainsOk returns a tuple with the DnsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetDnsDomainsOk() (*[]string, bool) {
	if o == nil || o.DnsDomains == nil {
		return nil, false
	}
	return &o.DnsDomains, true
}

// HasDnsDomains returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasDnsDomains() bool {
	if o != nil && o.DnsDomains != nil {
		return true
	}

	return false
}

// SetDnsDomains gets a reference to the given []string and assigns it to the DnsDomains field.
func (o *StorageNetAppCluster) SetDnsDomains(v []string) {
	o.DnsDomains = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppCluster) SetKey(v string) {
	o.Key = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *StorageNetAppCluster) SetLocation(v string) {
	o.Location = &v
}

// GetManagementAddress returns the ManagementAddress field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetManagementAddress() string {
	if o == nil || o.ManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.ManagementAddress
}

// GetManagementAddressOk returns a tuple with the ManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetManagementAddressOk() (*string, bool) {
	if o == nil || o.ManagementAddress == nil {
		return nil, false
	}
	return o.ManagementAddress, true
}

// HasManagementAddress returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasManagementAddress() bool {
	if o != nil && o.ManagementAddress != nil {
		return true
	}

	return false
}

// SetManagementAddress gets a reference to the given string and assigns it to the ManagementAddress field.
func (o *StorageNetAppCluster) SetManagementAddress(v string) {
	o.ManagementAddress = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetNameServersOk() (*[]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return &o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *StorageNetAppCluster) SetNameServers(v []string) {
	o.NameServers = v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetNtpServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetNtpServersOk() (*[]string, bool) {
	if o == nil || o.NtpServers == nil {
		return nil, false
	}
	return &o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasNtpServers() bool {
	if o != nil && o.NtpServers != nil {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *StorageNetAppCluster) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetTelnetEnabled returns the TelnetEnabled field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetTelnetEnabled() bool {
	if o == nil || o.TelnetEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TelnetEnabled
}

// GetTelnetEnabledOk returns a tuple with the TelnetEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetTelnetEnabledOk() (*bool, bool) {
	if o == nil || o.TelnetEnabled == nil {
		return nil, false
	}
	return o.TelnetEnabled, true
}

// HasTelnetEnabled returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasTelnetEnabled() bool {
	if o != nil && o.TelnetEnabled != nil {
		return true
	}

	return false
}

// SetTelnetEnabled gets a reference to the given bool and assigns it to the TelnetEnabled field.
func (o *StorageNetAppCluster) SetTelnetEnabled(v bool) {
	o.TelnetEnabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppCluster) GetEvents() []StorageNetAppClusterEventRelationship {
	if o == nil {
		var ret []StorageNetAppClusterEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppCluster) GetEventsOk() (*[]StorageNetAppClusterEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppClusterEventRelationship and assigns it to the Events field.
func (o *StorageNetAppCluster) SetEvents(v []StorageNetAppClusterEventRelationship) {
	o.Events = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageNetAppCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageNetAppCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageNetAppCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageNetAppCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArray, errStorageBaseArray := json.Marshal(o.StorageBaseArray)
	if errStorageBaseArray != nil {
		return []byte{}, errStorageBaseArray
	}
	errStorageBaseArray = json.Unmarshal([]byte(serializedStorageBaseArray), &toSerialize)
	if errStorageBaseArray != nil {
		return []byte{}, errStorageBaseArray
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoSupport.IsSet() {
		toSerialize["AutoSupport"] = o.AutoSupport.Get()
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
	if o.TelnetEnabled != nil {
		toSerialize["TelnetEnabled"] = o.TelnetEnabled
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

func (o *StorageNetAppCluster) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType            string                                        `json:"ObjectType"`
		AutoSupport           NullableStorageNetAppAutoSupport              `json:"AutoSupport,omitempty"`
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
		// Indicates whether or not telnet is enabled on the cluster.
		TelnetEnabled *bool `json:"TelnetEnabled,omitempty"`
		// An array of relationships to storageNetAppClusterEvent resources.
		Events           []StorageNetAppClusterEventRelationship `json:"Events,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStorageNetAppClusterWithoutEmbeddedStruct := StorageNetAppClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppClusterWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppCluster := _StorageNetAppCluster{}
		varStorageNetAppCluster.ClassId = varStorageNetAppClusterWithoutEmbeddedStruct.ClassId
		varStorageNetAppCluster.ObjectType = varStorageNetAppClusterWithoutEmbeddedStruct.ObjectType
		varStorageNetAppCluster.AutoSupport = varStorageNetAppClusterWithoutEmbeddedStruct.AutoSupport
		varStorageNetAppCluster.AvgPerformanceMetrics = varStorageNetAppClusterWithoutEmbeddedStruct.AvgPerformanceMetrics
		varStorageNetAppCluster.ClusterEfficiency = varStorageNetAppClusterWithoutEmbeddedStruct.ClusterEfficiency
		varStorageNetAppCluster.ClusterHealthStatus = varStorageNetAppClusterWithoutEmbeddedStruct.ClusterHealthStatus
		varStorageNetAppCluster.DnsDomains = varStorageNetAppClusterWithoutEmbeddedStruct.DnsDomains
		varStorageNetAppCluster.Key = varStorageNetAppClusterWithoutEmbeddedStruct.Key
		varStorageNetAppCluster.Location = varStorageNetAppClusterWithoutEmbeddedStruct.Location
		varStorageNetAppCluster.ManagementAddress = varStorageNetAppClusterWithoutEmbeddedStruct.ManagementAddress
		varStorageNetAppCluster.NameServers = varStorageNetAppClusterWithoutEmbeddedStruct.NameServers
		varStorageNetAppCluster.NtpServers = varStorageNetAppClusterWithoutEmbeddedStruct.NtpServers
		varStorageNetAppCluster.TelnetEnabled = varStorageNetAppClusterWithoutEmbeddedStruct.TelnetEnabled
		varStorageNetAppCluster.Events = varStorageNetAppClusterWithoutEmbeddedStruct.Events
		varStorageNetAppCluster.RegisteredDevice = varStorageNetAppClusterWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageNetAppCluster(varStorageNetAppCluster)
	} else {
		return err
	}

	varStorageNetAppCluster := _StorageNetAppCluster{}

	err = json.Unmarshal(bytes, &varStorageNetAppCluster)
	if err == nil {
		o.StorageBaseArray = varStorageNetAppCluster.StorageBaseArray
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoSupport")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "ClusterEfficiency")
		delete(additionalProperties, "ClusterHealthStatus")
		delete(additionalProperties, "DnsDomains")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "ManagementAddress")
		delete(additionalProperties, "NameServers")
		delete(additionalProperties, "NtpServers")
		delete(additionalProperties, "TelnetEnabled")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseArray := reflect.ValueOf(o.StorageBaseArray)
		for i := 0; i < reflectStorageBaseArray.Type().NumField(); i++ {
			t := reflectStorageBaseArray.Type().Field(i)

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

type NullableStorageNetAppCluster struct {
	value *StorageNetAppCluster
	isSet bool
}

func (v NullableStorageNetAppCluster) Get() *StorageNetAppCluster {
	return v.value
}

func (v *NullableStorageNetAppCluster) Set(val *StorageNetAppCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppCluster(val *StorageNetAppCluster) *NullableStorageNetAppCluster {
	return &NullableStorageNetAppCluster{value: val, isSet: true}
}

func (v NullableStorageNetAppCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
