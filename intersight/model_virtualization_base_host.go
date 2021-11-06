/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationBaseHost Common attributes of any host associated to a hypervisor manager. Serves as a base class for all concrete host types A Host belongs to a datacenter and optionally to a cluster, and runs virtual machines on it. A host is basically a hardware platform that runs the VMs. Depending on the capacity of the host, it can support 100s of virtual machines.
type VirtualizationBaseHost struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType   string                        `json:"ObjectType"`
	CpuInfo      NullableVirtualizationCpuInfo `json:"CpuInfo,omitempty"`
	HardwareInfo NullableInfraHardwareInfo     `json:"HardwareInfo,omitempty"`
	// Identifies the broad type of the underlying hypervisor. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference).
	Identity *string `json:"Identity,omitempty"`
	// Is this host in maintenance mode. Set to true or false.
	MaintenanceMode *bool                                `json:"MaintenanceMode,omitempty"`
	MemoryCapacity  NullableVirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
	// Commercial model information about this hardware.
	Model *string `json:"Model,omitempty"`
	// Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations.
	Name              *string                               `json:"Name,omitempty"`
	ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	ProductInfo       NullableVirtualizationProductInfo     `json:"ProductInfo,omitempty"`
	// Serial number of this host (internally generated).
	Serial *string `json:"Serial,omitempty"`
	// Host health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
	Status *string `json:"Status,omitempty"`
	// The uptime of the host, stored as Duration (from w3c).
	UpTime *string `json:"UpTime,omitempty"`
	// Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated.
	Uuid *string `json:"Uuid,omitempty"`
	// Commercial vendor details of this hardware.
	Vendor               *string `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseHost VirtualizationBaseHost

// NewVirtualizationBaseHost instantiates a new VirtualizationBaseHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseHost(classId string, objectType string) *VirtualizationBaseHost {
	this := VirtualizationBaseHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationBaseHostWithDefaults instantiates a new VirtualizationBaseHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseHostWithDefaults() *VirtualizationBaseHost {
	this := VirtualizationBaseHost{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuInfo returns the CpuInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseHost) GetCpuInfo() VirtualizationCpuInfo {
	if o == nil || o.CpuInfo.Get() == nil {
		var ret VirtualizationCpuInfo
		return ret
	}
	return *o.CpuInfo.Get()
}

// GetCpuInfoOk returns a tuple with the CpuInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseHost) GetCpuInfoOk() (*VirtualizationCpuInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuInfo.Get(), o.CpuInfo.IsSet()
}

// HasCpuInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasCpuInfo() bool {
	if o != nil && o.CpuInfo.IsSet() {
		return true
	}

	return false
}

// SetCpuInfo gets a reference to the given NullableVirtualizationCpuInfo and assigns it to the CpuInfo field.
func (o *VirtualizationBaseHost) SetCpuInfo(v VirtualizationCpuInfo) {
	o.CpuInfo.Set(&v)
}

// SetCpuInfoNil sets the value for CpuInfo to be an explicit nil
func (o *VirtualizationBaseHost) SetCpuInfoNil() {
	o.CpuInfo.Set(nil)
}

// UnsetCpuInfo ensures that no value is present for CpuInfo, not even an explicit nil
func (o *VirtualizationBaseHost) UnsetCpuInfo() {
	o.CpuInfo.Unset()
}

// GetHardwareInfo returns the HardwareInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseHost) GetHardwareInfo() InfraHardwareInfo {
	if o == nil || o.HardwareInfo.Get() == nil {
		var ret InfraHardwareInfo
		return ret
	}
	return *o.HardwareInfo.Get()
}

// GetHardwareInfoOk returns a tuple with the HardwareInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseHost) GetHardwareInfoOk() (*InfraHardwareInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardwareInfo.Get(), o.HardwareInfo.IsSet()
}

// HasHardwareInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasHardwareInfo() bool {
	if o != nil && o.HardwareInfo.IsSet() {
		return true
	}

	return false
}

// SetHardwareInfo gets a reference to the given NullableInfraHardwareInfo and assigns it to the HardwareInfo field.
func (o *VirtualizationBaseHost) SetHardwareInfo(v InfraHardwareInfo) {
	o.HardwareInfo.Set(&v)
}

// SetHardwareInfoNil sets the value for HardwareInfo to be an explicit nil
func (o *VirtualizationBaseHost) SetHardwareInfoNil() {
	o.HardwareInfo.Set(nil)
}

// UnsetHardwareInfo ensures that no value is present for HardwareInfo, not even an explicit nil
func (o *VirtualizationBaseHost) UnsetHardwareInfo() {
	o.HardwareInfo.Unset()
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseHost) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseHost) SetIdentity(v string) {
	o.Identity = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetMaintenanceMode() bool {
	if o == nil || o.MaintenanceMode == nil {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || o.MaintenanceMode == nil {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasMaintenanceMode() bool {
	if o != nil && o.MaintenanceMode != nil {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *VirtualizationBaseHost) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseHost) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity.Get() == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity.Get()
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseHost) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryCapacity.Get(), o.MemoryCapacity.IsSet()
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity.IsSet() {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given NullableVirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseHost) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity.Set(&v)
}

// SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil
func (o *VirtualizationBaseHost) SetMemoryCapacityNil() {
	o.MemoryCapacity.Set(nil)
}

// UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
func (o *VirtualizationBaseHost) UnsetMemoryCapacity() {
	o.MemoryCapacity.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *VirtualizationBaseHost) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBaseHost) SetName(v string) {
	o.Name = &v
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseHost) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity.Get() == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity.Get()
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseHost) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessorCapacity.Get(), o.ProcessorCapacity.IsSet()
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity.IsSet() {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given NullableVirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseHost) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity.Set(&v)
}

// SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil
func (o *VirtualizationBaseHost) SetProcessorCapacityNil() {
	o.ProcessorCapacity.Set(nil)
}

// UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
func (o *VirtualizationBaseHost) UnsetProcessorCapacity() {
	o.ProcessorCapacity.Unset()
}

// GetProductInfo returns the ProductInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseHost) GetProductInfo() VirtualizationProductInfo {
	if o == nil || o.ProductInfo.Get() == nil {
		var ret VirtualizationProductInfo
		return ret
	}
	return *o.ProductInfo.Get()
}

// GetProductInfoOk returns a tuple with the ProductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseHost) GetProductInfoOk() (*VirtualizationProductInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductInfo.Get(), o.ProductInfo.IsSet()
}

// HasProductInfo returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasProductInfo() bool {
	if o != nil && o.ProductInfo.IsSet() {
		return true
	}

	return false
}

// SetProductInfo gets a reference to the given NullableVirtualizationProductInfo and assigns it to the ProductInfo field.
func (o *VirtualizationBaseHost) SetProductInfo(v VirtualizationProductInfo) {
	o.ProductInfo.Set(&v)
}

// SetProductInfoNil sets the value for ProductInfo to be an explicit nil
func (o *VirtualizationBaseHost) SetProductInfoNil() {
	o.ProductInfo.Set(nil)
}

// UnsetProductInfo ensures that no value is present for ProductInfo, not even an explicit nil
func (o *VirtualizationBaseHost) UnsetProductInfo() {
	o.ProductInfo.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *VirtualizationBaseHost) SetSerial(v string) {
	o.Serial = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationBaseHost) SetStatus(v string) {
	o.Status = &v
}

// GetUpTime returns the UpTime field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetUpTime() string {
	if o == nil || o.UpTime == nil {
		var ret string
		return ret
	}
	return *o.UpTime
}

// GetUpTimeOk returns a tuple with the UpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetUpTimeOk() (*string, bool) {
	if o == nil || o.UpTime == nil {
		return nil, false
	}
	return o.UpTime, true
}

// HasUpTime returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasUpTime() bool {
	if o != nil && o.UpTime != nil {
		return true
	}

	return false
}

// SetUpTime gets a reference to the given string and assigns it to the UpTime field.
func (o *VirtualizationBaseHost) SetUpTime(v string) {
	o.UpTime = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBaseHost) SetUuid(v string) {
	o.Uuid = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *VirtualizationBaseHost) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseHost) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *VirtualizationBaseHost) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *VirtualizationBaseHost) SetVendor(v string) {
	o.Vendor = &v
}

func (o VirtualizationBaseHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return []byte{}, errVirtualizationBaseSourceDevice
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuInfo.IsSet() {
		toSerialize["CpuInfo"] = o.CpuInfo.Get()
	}
	if o.HardwareInfo.IsSet() {
		toSerialize["HardwareInfo"] = o.HardwareInfo.Get()
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.MaintenanceMode != nil {
		toSerialize["MaintenanceMode"] = o.MaintenanceMode
	}
	if o.MemoryCapacity.IsSet() {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity.Get()
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ProcessorCapacity.IsSet() {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity.Get()
	}
	if o.ProductInfo.IsSet() {
		toSerialize["ProductInfo"] = o.ProductInfo.Get()
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UpTime != nil {
		toSerialize["UpTime"] = o.UpTime
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseHost) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType   string                        `json:"ObjectType"`
		CpuInfo      NullableVirtualizationCpuInfo `json:"CpuInfo,omitempty"`
		HardwareInfo NullableInfraHardwareInfo     `json:"HardwareInfo,omitempty"`
		// Identifies the broad type of the underlying hypervisor. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference).
		Identity *string `json:"Identity,omitempty"`
		// Is this host in maintenance mode. Set to true or false.
		MaintenanceMode *bool                                `json:"MaintenanceMode,omitempty"`
		MemoryCapacity  NullableVirtualizationMemoryCapacity `json:"MemoryCapacity,omitempty"`
		// Commercial model information about this hardware.
		Model *string `json:"Model,omitempty"`
		// Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations.
		Name              *string                               `json:"Name,omitempty"`
		ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
		ProductInfo       NullableVirtualizationProductInfo     `json:"ProductInfo,omitempty"`
		// Serial number of this host (internally generated).
		Serial *string `json:"Serial,omitempty"`
		// Host health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
		Status *string `json:"Status,omitempty"`
		// The uptime of the host, stored as Duration (from w3c).
		UpTime *string `json:"UpTime,omitempty"`
		// Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated.
		Uuid *string `json:"Uuid,omitempty"`
		// Commercial vendor details of this hardware.
		Vendor *string `json:"Vendor,omitempty"`
	}

	varVirtualizationBaseHostWithoutEmbeddedStruct := VirtualizationBaseHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseHostWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseHost := _VirtualizationBaseHost{}
		varVirtualizationBaseHost.ClassId = varVirtualizationBaseHostWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseHost.ObjectType = varVirtualizationBaseHostWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseHost.CpuInfo = varVirtualizationBaseHostWithoutEmbeddedStruct.CpuInfo
		varVirtualizationBaseHost.HardwareInfo = varVirtualizationBaseHostWithoutEmbeddedStruct.HardwareInfo
		varVirtualizationBaseHost.HypervisorType = varVirtualizationBaseHostWithoutEmbeddedStruct.HypervisorType
		varVirtualizationBaseHost.Identity = varVirtualizationBaseHostWithoutEmbeddedStruct.Identity
		varVirtualizationBaseHost.MaintenanceMode = varVirtualizationBaseHostWithoutEmbeddedStruct.MaintenanceMode
		varVirtualizationBaseHost.MemoryCapacity = varVirtualizationBaseHostWithoutEmbeddedStruct.MemoryCapacity
		varVirtualizationBaseHost.Model = varVirtualizationBaseHostWithoutEmbeddedStruct.Model
		varVirtualizationBaseHost.Name = varVirtualizationBaseHostWithoutEmbeddedStruct.Name
		varVirtualizationBaseHost.ProcessorCapacity = varVirtualizationBaseHostWithoutEmbeddedStruct.ProcessorCapacity
		varVirtualizationBaseHost.ProductInfo = varVirtualizationBaseHostWithoutEmbeddedStruct.ProductInfo
		varVirtualizationBaseHost.Serial = varVirtualizationBaseHostWithoutEmbeddedStruct.Serial
		varVirtualizationBaseHost.Status = varVirtualizationBaseHostWithoutEmbeddedStruct.Status
		varVirtualizationBaseHost.UpTime = varVirtualizationBaseHostWithoutEmbeddedStruct.UpTime
		varVirtualizationBaseHost.Uuid = varVirtualizationBaseHostWithoutEmbeddedStruct.Uuid
		varVirtualizationBaseHost.Vendor = varVirtualizationBaseHostWithoutEmbeddedStruct.Vendor
		*o = VirtualizationBaseHost(varVirtualizationBaseHost)
	} else {
		return err
	}

	varVirtualizationBaseHost := _VirtualizationBaseHost{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseHost)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBaseHost.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuInfo")
		delete(additionalProperties, "HardwareInfo")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "MaintenanceMode")
		delete(additionalProperties, "MemoryCapacity")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ProcessorCapacity")
		delete(additionalProperties, "ProductInfo")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UpTime")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBaseHost struct {
	value *VirtualizationBaseHost
	isSet bool
}

func (v NullableVirtualizationBaseHost) Get() *VirtualizationBaseHost {
	return v.value
}

func (v *NullableVirtualizationBaseHost) Set(val *VirtualizationBaseHost) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseHost) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseHost(val *VirtualizationBaseHost) *NullableVirtualizationBaseHost {
	return &NullableVirtualizationBaseHost{value: val, isSet: true}
}

func (v NullableVirtualizationBaseHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
