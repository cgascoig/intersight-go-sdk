/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageHitachiPool A pool entity in Hitachi storage array.
type StorageHitachiPool struct {
	StorageBaseDiskPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
	BlockingModeBlockade *string `json:"BlockingModeBlockade,omitempty"`
	// Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
	BlockingModeFull *string `json:"BlockingModeFull,omitempty"`
	// The depletion threshold set for the pool (%).
	DepletionThreshold *string `json:"DepletionThreshold,omitempty"`
	// Whether the pool is shrinking is output.
	IsShrinking *bool `json:"IsShrinking,omitempty"`
	// Performance monitoring execution mode (monitor mode). * `N/A` - Performance monitoring is not available. * `Period mode` - Period mode is the default setting. If Period mode is enabled, tier range values and page relocations are determined based solely on the monitoring data from the last complete cycle. * `Continuous mode` - When Continuous mode is enabled, the weighted average efficiency is calculated using the latest monitoring information and the collected monitoring information in the past cycles. Page relocations are determined using this weighted average efficiency.
	MonitoringMode *string `json:"MonitoringMode,omitempty"`
	// Status of monitor information.
	MonitoringStatus *string `json:"MonitoringStatus,omitempty"`
	// Execution mode for the pool. * `N/A` - Execution Mode is not available for the pool. * `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator). * `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC.
	PoolActionMode *string `json:"PoolActionMode,omitempty"`
	// Displays the status of the tier relocation processing.
	ProgressOfReplacing *string `json:"ProgressOfReplacing,omitempty"`
	// Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool.
	TotalReservedCapacity *int64 `json:"TotalReservedCapacity,omitempty"`
	// The warning threshold set for the pool (%).
	WarningThreshold     *int64                               `json:"WarningThreshold,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiPool StorageHitachiPool

// NewStorageHitachiPool instantiates a new StorageHitachiPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiPool(classId string, objectType string) *StorageHitachiPool {
	this := StorageHitachiPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiPoolWithDefaults instantiates a new StorageHitachiPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiPoolWithDefaults() *StorageHitachiPool {
	this := StorageHitachiPool{}
	var classId string = "storage.HitachiPool"
	this.ClassId = classId
	var objectType string = "storage.HitachiPool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBlockingModeBlockade returns the BlockingModeBlockade field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetBlockingModeBlockade() string {
	if o == nil || o.BlockingModeBlockade == nil {
		var ret string
		return ret
	}
	return *o.BlockingModeBlockade
}

// GetBlockingModeBlockadeOk returns a tuple with the BlockingModeBlockade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetBlockingModeBlockadeOk() (*string, bool) {
	if o == nil || o.BlockingModeBlockade == nil {
		return nil, false
	}
	return o.BlockingModeBlockade, true
}

// HasBlockingModeBlockade returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasBlockingModeBlockade() bool {
	if o != nil && o.BlockingModeBlockade != nil {
		return true
	}

	return false
}

// SetBlockingModeBlockade gets a reference to the given string and assigns it to the BlockingModeBlockade field.
func (o *StorageHitachiPool) SetBlockingModeBlockade(v string) {
	o.BlockingModeBlockade = &v
}

// GetBlockingModeFull returns the BlockingModeFull field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetBlockingModeFull() string {
	if o == nil || o.BlockingModeFull == nil {
		var ret string
		return ret
	}
	return *o.BlockingModeFull
}

// GetBlockingModeFullOk returns a tuple with the BlockingModeFull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetBlockingModeFullOk() (*string, bool) {
	if o == nil || o.BlockingModeFull == nil {
		return nil, false
	}
	return o.BlockingModeFull, true
}

// HasBlockingModeFull returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasBlockingModeFull() bool {
	if o != nil && o.BlockingModeFull != nil {
		return true
	}

	return false
}

// SetBlockingModeFull gets a reference to the given string and assigns it to the BlockingModeFull field.
func (o *StorageHitachiPool) SetBlockingModeFull(v string) {
	o.BlockingModeFull = &v
}

// GetDepletionThreshold returns the DepletionThreshold field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetDepletionThreshold() string {
	if o == nil || o.DepletionThreshold == nil {
		var ret string
		return ret
	}
	return *o.DepletionThreshold
}

// GetDepletionThresholdOk returns a tuple with the DepletionThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetDepletionThresholdOk() (*string, bool) {
	if o == nil || o.DepletionThreshold == nil {
		return nil, false
	}
	return o.DepletionThreshold, true
}

// HasDepletionThreshold returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasDepletionThreshold() bool {
	if o != nil && o.DepletionThreshold != nil {
		return true
	}

	return false
}

// SetDepletionThreshold gets a reference to the given string and assigns it to the DepletionThreshold field.
func (o *StorageHitachiPool) SetDepletionThreshold(v string) {
	o.DepletionThreshold = &v
}

// GetIsShrinking returns the IsShrinking field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetIsShrinking() bool {
	if o == nil || o.IsShrinking == nil {
		var ret bool
		return ret
	}
	return *o.IsShrinking
}

// GetIsShrinkingOk returns a tuple with the IsShrinking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetIsShrinkingOk() (*bool, bool) {
	if o == nil || o.IsShrinking == nil {
		return nil, false
	}
	return o.IsShrinking, true
}

// HasIsShrinking returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasIsShrinking() bool {
	if o != nil && o.IsShrinking != nil {
		return true
	}

	return false
}

// SetIsShrinking gets a reference to the given bool and assigns it to the IsShrinking field.
func (o *StorageHitachiPool) SetIsShrinking(v bool) {
	o.IsShrinking = &v
}

// GetMonitoringMode returns the MonitoringMode field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetMonitoringMode() string {
	if o == nil || o.MonitoringMode == nil {
		var ret string
		return ret
	}
	return *o.MonitoringMode
}

// GetMonitoringModeOk returns a tuple with the MonitoringMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetMonitoringModeOk() (*string, bool) {
	if o == nil || o.MonitoringMode == nil {
		return nil, false
	}
	return o.MonitoringMode, true
}

// HasMonitoringMode returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasMonitoringMode() bool {
	if o != nil && o.MonitoringMode != nil {
		return true
	}

	return false
}

// SetMonitoringMode gets a reference to the given string and assigns it to the MonitoringMode field.
func (o *StorageHitachiPool) SetMonitoringMode(v string) {
	o.MonitoringMode = &v
}

// GetMonitoringStatus returns the MonitoringStatus field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetMonitoringStatus() string {
	if o == nil || o.MonitoringStatus == nil {
		var ret string
		return ret
	}
	return *o.MonitoringStatus
}

// GetMonitoringStatusOk returns a tuple with the MonitoringStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetMonitoringStatusOk() (*string, bool) {
	if o == nil || o.MonitoringStatus == nil {
		return nil, false
	}
	return o.MonitoringStatus, true
}

// HasMonitoringStatus returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasMonitoringStatus() bool {
	if o != nil && o.MonitoringStatus != nil {
		return true
	}

	return false
}

// SetMonitoringStatus gets a reference to the given string and assigns it to the MonitoringStatus field.
func (o *StorageHitachiPool) SetMonitoringStatus(v string) {
	o.MonitoringStatus = &v
}

// GetPoolActionMode returns the PoolActionMode field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetPoolActionMode() string {
	if o == nil || o.PoolActionMode == nil {
		var ret string
		return ret
	}
	return *o.PoolActionMode
}

// GetPoolActionModeOk returns a tuple with the PoolActionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetPoolActionModeOk() (*string, bool) {
	if o == nil || o.PoolActionMode == nil {
		return nil, false
	}
	return o.PoolActionMode, true
}

// HasPoolActionMode returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasPoolActionMode() bool {
	if o != nil && o.PoolActionMode != nil {
		return true
	}

	return false
}

// SetPoolActionMode gets a reference to the given string and assigns it to the PoolActionMode field.
func (o *StorageHitachiPool) SetPoolActionMode(v string) {
	o.PoolActionMode = &v
}

// GetProgressOfReplacing returns the ProgressOfReplacing field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetProgressOfReplacing() string {
	if o == nil || o.ProgressOfReplacing == nil {
		var ret string
		return ret
	}
	return *o.ProgressOfReplacing
}

// GetProgressOfReplacingOk returns a tuple with the ProgressOfReplacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetProgressOfReplacingOk() (*string, bool) {
	if o == nil || o.ProgressOfReplacing == nil {
		return nil, false
	}
	return o.ProgressOfReplacing, true
}

// HasProgressOfReplacing returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasProgressOfReplacing() bool {
	if o != nil && o.ProgressOfReplacing != nil {
		return true
	}

	return false
}

// SetProgressOfReplacing gets a reference to the given string and assigns it to the ProgressOfReplacing field.
func (o *StorageHitachiPool) SetProgressOfReplacing(v string) {
	o.ProgressOfReplacing = &v
}

// GetTotalReservedCapacity returns the TotalReservedCapacity field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetTotalReservedCapacity() int64 {
	if o == nil || o.TotalReservedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.TotalReservedCapacity
}

// GetTotalReservedCapacityOk returns a tuple with the TotalReservedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetTotalReservedCapacityOk() (*int64, bool) {
	if o == nil || o.TotalReservedCapacity == nil {
		return nil, false
	}
	return o.TotalReservedCapacity, true
}

// HasTotalReservedCapacity returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasTotalReservedCapacity() bool {
	if o != nil && o.TotalReservedCapacity != nil {
		return true
	}

	return false
}

// SetTotalReservedCapacity gets a reference to the given int64 and assigns it to the TotalReservedCapacity field.
func (o *StorageHitachiPool) SetTotalReservedCapacity(v int64) {
	o.TotalReservedCapacity = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetWarningThreshold() int64 {
	if o == nil || o.WarningThreshold == nil {
		var ret int64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetWarningThresholdOk() (*int64, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasWarningThreshold() bool {
	if o != nil && o.WarningThreshold != nil {
		return true
	}

	return false
}

// SetWarningThreshold gets a reference to the given int64 and assigns it to the WarningThreshold field.
func (o *StorageHitachiPool) SetWarningThreshold(v int64) {
	o.WarningThreshold = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiPool) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiPool) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPool) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiPool) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiPool) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseDiskPool, errStorageBaseDiskPool := json.Marshal(o.StorageBaseDiskPool)
	if errStorageBaseDiskPool != nil {
		return []byte{}, errStorageBaseDiskPool
	}
	errStorageBaseDiskPool = json.Unmarshal([]byte(serializedStorageBaseDiskPool), &toSerialize)
	if errStorageBaseDiskPool != nil {
		return []byte{}, errStorageBaseDiskPool
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BlockingModeBlockade != nil {
		toSerialize["BlockingModeBlockade"] = o.BlockingModeBlockade
	}
	if o.BlockingModeFull != nil {
		toSerialize["BlockingModeFull"] = o.BlockingModeFull
	}
	if o.DepletionThreshold != nil {
		toSerialize["DepletionThreshold"] = o.DepletionThreshold
	}
	if o.IsShrinking != nil {
		toSerialize["IsShrinking"] = o.IsShrinking
	}
	if o.MonitoringMode != nil {
		toSerialize["MonitoringMode"] = o.MonitoringMode
	}
	if o.MonitoringStatus != nil {
		toSerialize["MonitoringStatus"] = o.MonitoringStatus
	}
	if o.PoolActionMode != nil {
		toSerialize["PoolActionMode"] = o.PoolActionMode
	}
	if o.ProgressOfReplacing != nil {
		toSerialize["ProgressOfReplacing"] = o.ProgressOfReplacing
	}
	if o.TotalReservedCapacity != nil {
		toSerialize["TotalReservedCapacity"] = o.TotalReservedCapacity
	}
	if o.WarningThreshold != nil {
		toSerialize["WarningThreshold"] = o.WarningThreshold
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiPool) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Setting the protection function for a virtual volume. When the DP pool is blockade, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
		BlockingModeBlockade *string `json:"BlockingModeBlockade,omitempty"`
		// Setting the protection function for a virtual volume. When the DP pool is full, whether the read and write operations can be performed for the DP volume that uses the target DP pool is output. Yes, read and write operations are not possible. No, read and write operations are possible. -, Thin Image pool or not available.
		BlockingModeFull *string `json:"BlockingModeFull,omitempty"`
		// The depletion threshold set for the pool (%).
		DepletionThreshold *string `json:"DepletionThreshold,omitempty"`
		// Whether the pool is shrinking is output.
		IsShrinking *bool `json:"IsShrinking,omitempty"`
		// Performance monitoring execution mode (monitor mode). * `N/A` - Performance monitoring is not available. * `Period mode` - Period mode is the default setting. If Period mode is enabled, tier range values and page relocations are determined based solely on the monitoring data from the last complete cycle. * `Continuous mode` - When Continuous mode is enabled, the weighted average efficiency is calculated using the latest monitoring information and the collected monitoring information in the past cycles. Page relocations are determined using this weighted average efficiency.
		MonitoringMode *string `json:"MonitoringMode,omitempty"`
		// Status of monitor information.
		MonitoringStatus *string `json:"MonitoringStatus,omitempty"`
		// Execution mode for the pool. * `N/A` - Execution Mode is not available for the pool. * `Auto` - The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator). * `Manual` - The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC.
		PoolActionMode *string `json:"PoolActionMode,omitempty"`
		// Displays the status of the tier relocation processing.
		ProgressOfReplacing *string `json:"ProgressOfReplacing,omitempty"`
		// Total capacity of the reserved page (bytes) of the DP volume that is related to the DP pool.
		TotalReservedCapacity *int64 `json:"TotalReservedCapacity,omitempty"`
		// The warning threshold set for the pool (%).
		WarningThreshold *int64                               `json:"WarningThreshold,omitempty"`
		Array            *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiPoolWithoutEmbeddedStruct := StorageHitachiPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiPoolWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiPool := _StorageHitachiPool{}
		varStorageHitachiPool.ClassId = varStorageHitachiPoolWithoutEmbeddedStruct.ClassId
		varStorageHitachiPool.ObjectType = varStorageHitachiPoolWithoutEmbeddedStruct.ObjectType
		varStorageHitachiPool.BlockingModeBlockade = varStorageHitachiPoolWithoutEmbeddedStruct.BlockingModeBlockade
		varStorageHitachiPool.BlockingModeFull = varStorageHitachiPoolWithoutEmbeddedStruct.BlockingModeFull
		varStorageHitachiPool.DepletionThreshold = varStorageHitachiPoolWithoutEmbeddedStruct.DepletionThreshold
		varStorageHitachiPool.IsShrinking = varStorageHitachiPoolWithoutEmbeddedStruct.IsShrinking
		varStorageHitachiPool.MonitoringMode = varStorageHitachiPoolWithoutEmbeddedStruct.MonitoringMode
		varStorageHitachiPool.MonitoringStatus = varStorageHitachiPoolWithoutEmbeddedStruct.MonitoringStatus
		varStorageHitachiPool.PoolActionMode = varStorageHitachiPoolWithoutEmbeddedStruct.PoolActionMode
		varStorageHitachiPool.ProgressOfReplacing = varStorageHitachiPoolWithoutEmbeddedStruct.ProgressOfReplacing
		varStorageHitachiPool.TotalReservedCapacity = varStorageHitachiPoolWithoutEmbeddedStruct.TotalReservedCapacity
		varStorageHitachiPool.WarningThreshold = varStorageHitachiPoolWithoutEmbeddedStruct.WarningThreshold
		varStorageHitachiPool.Array = varStorageHitachiPoolWithoutEmbeddedStruct.Array
		varStorageHitachiPool.RegisteredDevice = varStorageHitachiPoolWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiPool(varStorageHitachiPool)
	} else {
		return err
	}

	varStorageHitachiPool := _StorageHitachiPool{}

	err = json.Unmarshal(bytes, &varStorageHitachiPool)
	if err == nil {
		o.StorageBaseDiskPool = varStorageHitachiPool.StorageBaseDiskPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BlockingModeBlockade")
		delete(additionalProperties, "BlockingModeFull")
		delete(additionalProperties, "DepletionThreshold")
		delete(additionalProperties, "IsShrinking")
		delete(additionalProperties, "MonitoringMode")
		delete(additionalProperties, "MonitoringStatus")
		delete(additionalProperties, "PoolActionMode")
		delete(additionalProperties, "ProgressOfReplacing")
		delete(additionalProperties, "TotalReservedCapacity")
		delete(additionalProperties, "WarningThreshold")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseDiskPool := reflect.ValueOf(o.StorageBaseDiskPool)
		for i := 0; i < reflectStorageBaseDiskPool.Type().NumField(); i++ {
			t := reflectStorageBaseDiskPool.Type().Field(i)

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

type NullableStorageHitachiPool struct {
	value *StorageHitachiPool
	isSet bool
}

func (v NullableStorageHitachiPool) Get() *StorageHitachiPool {
	return v.value
}

func (v *NullableStorageHitachiPool) Set(val *StorageHitachiPool) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiPool) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiPool(val *StorageHitachiPool) *NullableStorageHitachiPool {
	return &NullableStorageHitachiPool{value: val, isSet: true}
}

func (v NullableStorageHitachiPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
