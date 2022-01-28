/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationVmwareVirtualDiskAllOf Definition of the list of properties defined in 'virtualization.VmwareVirtualDisk', excluding properties defined in parent classes.
type VirtualizationVmwareVirtualDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Compatibility mode of the raw disk mapping (RDM). * `notApplicable` - Value specified for any disk which is not of raw device mapping type. * `physicalMode` - A disk device backed by a physical compatibility mode raw disk mapping cannot use disk modes, and commands are passed straight through to the LUN indicated by the raw disk mapping. * `virtualMode` - A disk device backed by a virtual compatibility mode raw disk mapping can use disk modes.
	CompatibilityMode *string `json:"CompatibilityMode,omitempty"`
	// Key of the controller on which the disk is created.
	ControllerKey *int64 `json:"ControllerKey,omitempty"`
	// Host-specific device the LUN is being accessed through. If the target LUN is not available on the host then it is empty. For example, this could happen if it has accidentally been masked out.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Persistence mode of the virtual disk. For RDM disks, it is only supported when the raw disk mapping is using virtual compatibility mode. * `persistent` - Changes are immediately and permanently written to the virtual disk. * `independent_persistent` - Changes are immediately and permanently written to the virtual disk and not affected by snapshots. * `independent_nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off and not affected by snapshots. * `nonpersistent` - Changes to virtual disk are made to a redo log and discarded at power off. * `undoable` - Changes to virtual disk are made to a redo log and has the option to commit or undo. * `append` - Changes are appended to the redo log and can be revoked by removing the undo log.
	DiskMode *string `json:"DiskMode,omitempty"`
	// Specifies whether the virtual disk is a RDM or a Hard disk. * `flatDisk` - Specifies that it is a flat disk. * `rdmDisk` - Specifies that it is a raw device mapping disk.
	DiskType *string `json:"DiskType,omitempty"`
	// The internally assigned key of this disk. This entity is not manipulated by users.
	Key *int64 `json:"Key,omitempty"`
	// The utilization of a virtual machine will not exceed this limit, even if there are available resources. Used to ensure a consistent performance of virtual machines independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). The unit is number of I/O per second.
	Limit *int64 `json:"Limit,omitempty"`
	// Unique identifier of the LUN accessed by the raw disk mapping (RDM).
	LunUuid *string `json:"LunUuid,omitempty"`
	// Serial ID of the storage device.
	Serial *string                                `json:"Serial,omitempty"`
	Shares NullableVirtualizationVmwareSharesInfo `json:"Shares,omitempty"`
	// Sharing mode of the virtual disk. * `sharingNone` - The virtual disk is not shared. * `sharingMultiWriter` - The virtual disk is shared between multiple virtual machines.
	Sharing *string `json:"Sharing,omitempty"`
	// Allocation type for the virtual disk. * `notApplicable` - Value specified for a disk for which storage allocation type is not applicable. * `thin` - A thin provisioned disk consumes only the space that it needs for its initial operrations, and grows with time according to demand. It is the fastest method to create a virtual disk because it creates a disk with just the header information. * `lazyZeroedThick` - A lazy zeroed thick disk has all space allocated at the time of its creation. Data remaining on the physical device is not erased during creation, but is zeroed out on demand later on first write from the virtual machine. * `eagerZeroedThick` - An eager zeroed thick disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. Such disks may take longer time during creation compared to other disk formats.
	StorageAllocationType *string `json:"StorageAllocationType,omitempty"`
	// Unit number of the disk on its controller.
	UnitNumber *int64 `json:"UnitNumber,omitempty"`
	// UUID assigned by vCenter to every disk.
	Uuid *string `json:"Uuid,omitempty"`
	// Identity of the virtual disk object as the first class entity.
	VdiskId *string `json:"VdiskId,omitempty"`
	// Vendor of the storage device.
	Vendor *string `json:"Vendor,omitempty"`
	// Path of the virtual disk.
	VirtualDiskPath *string `json:"VirtualDiskPath,omitempty"`
	// Identity of the virtual machine where the virtual disk is created.
	VmIdentity           *string                                         `json:"VmIdentity,omitempty"`
	Datastore            *VirtualizationVmwareDatastoreRelationship      `json:"Datastore,omitempty"`
	VirtualMachine       *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualDiskAllOf VirtualizationVmwareVirtualDiskAllOf

// NewVirtualizationVmwareVirtualDiskAllOf instantiates a new VirtualizationVmwareVirtualDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualDiskAllOf(classId string, objectType string) *VirtualizationVmwareVirtualDiskAllOf {
	this := VirtualizationVmwareVirtualDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var compatibilityMode string = "notApplicable"
	this.CompatibilityMode = &compatibilityMode
	var diskMode string = "persistent"
	this.DiskMode = &diskMode
	var diskType string = "flatDisk"
	this.DiskType = &diskType
	var sharing string = "sharingNone"
	this.Sharing = &sharing
	var storageAllocationType string = "notApplicable"
	this.StorageAllocationType = &storageAllocationType
	return &this
}

// NewVirtualizationVmwareVirtualDiskAllOfWithDefaults instantiates a new VirtualizationVmwareVirtualDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualDiskAllOfWithDefaults() *VirtualizationVmwareVirtualDiskAllOf {
	this := VirtualizationVmwareVirtualDiskAllOf{}
	var classId string = "virtualization.VmwareVirtualDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualDisk"
	this.ObjectType = objectType
	var compatibilityMode string = "notApplicable"
	this.CompatibilityMode = &compatibilityMode
	var diskMode string = "persistent"
	this.DiskMode = &diskMode
	var diskType string = "flatDisk"
	this.DiskType = &diskType
	var sharing string = "sharingNone"
	this.Sharing = &sharing
	var storageAllocationType string = "notApplicable"
	this.StorageAllocationType = &storageAllocationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVirtualDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompatibilityMode returns the CompatibilityMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetCompatibilityMode() string {
	if o == nil || o.CompatibilityMode == nil {
		var ret string
		return ret
	}
	return *o.CompatibilityMode
}

// GetCompatibilityModeOk returns a tuple with the CompatibilityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetCompatibilityModeOk() (*string, bool) {
	if o == nil || o.CompatibilityMode == nil {
		return nil, false
	}
	return o.CompatibilityMode, true
}

// HasCompatibilityMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasCompatibilityMode() bool {
	if o != nil && o.CompatibilityMode != nil {
		return true
	}

	return false
}

// SetCompatibilityMode gets a reference to the given string and assigns it to the CompatibilityMode field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetCompatibilityMode(v string) {
	o.CompatibilityMode = &v
}

// GetControllerKey returns the ControllerKey field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetControllerKey() int64 {
	if o == nil || o.ControllerKey == nil {
		var ret int64
		return ret
	}
	return *o.ControllerKey
}

// GetControllerKeyOk returns a tuple with the ControllerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetControllerKeyOk() (*int64, bool) {
	if o == nil || o.ControllerKey == nil {
		return nil, false
	}
	return o.ControllerKey, true
}

// HasControllerKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasControllerKey() bool {
	if o != nil && o.ControllerKey != nil {
		return true
	}

	return false
}

// SetControllerKey gets a reference to the given int64 and assigns it to the ControllerKey field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetControllerKey(v int64) {
	o.ControllerKey = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDiskMode returns the DiskMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDiskMode() string {
	if o == nil || o.DiskMode == nil {
		var ret string
		return ret
	}
	return *o.DiskMode
}

// GetDiskModeOk returns a tuple with the DiskMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDiskModeOk() (*string, bool) {
	if o == nil || o.DiskMode == nil {
		return nil, false
	}
	return o.DiskMode, true
}

// HasDiskMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasDiskMode() bool {
	if o != nil && o.DiskMode != nil {
		return true
	}

	return false
}

// SetDiskMode gets a reference to the given string and assigns it to the DiskMode field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetDiskMode(v string) {
	o.DiskMode = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetDiskType(v string) {
	o.DiskType = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetKey() int64 {
	if o == nil || o.Key == nil {
		var ret int64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetKeyOk() (*int64, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given int64 and assigns it to the Key field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetKey(v int64) {
	o.Key = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetLimit(v int64) {
	o.Limit = &v
}

// GetLunUuid returns the LunUuid field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetLunUuid() string {
	if o == nil || o.LunUuid == nil {
		var ret string
		return ret
	}
	return *o.LunUuid
}

// GetLunUuidOk returns a tuple with the LunUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetLunUuidOk() (*string, bool) {
	if o == nil || o.LunUuid == nil {
		return nil, false
	}
	return o.LunUuid, true
}

// HasLunUuid returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasLunUuid() bool {
	if o != nil && o.LunUuid != nil {
		return true
	}

	return false
}

// SetLunUuid gets a reference to the given string and assigns it to the LunUuid field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetLunUuid(v string) {
	o.LunUuid = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetShares returns the Shares field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareVirtualDiskAllOf) GetShares() VirtualizationVmwareSharesInfo {
	if o == nil || o.Shares.Get() == nil {
		var ret VirtualizationVmwareSharesInfo
		return ret
	}
	return *o.Shares.Get()
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareVirtualDiskAllOf) GetSharesOk() (*VirtualizationVmwareSharesInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shares.Get(), o.Shares.IsSet()
}

// HasShares returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasShares() bool {
	if o != nil && o.Shares.IsSet() {
		return true
	}

	return false
}

// SetShares gets a reference to the given NullableVirtualizationVmwareSharesInfo and assigns it to the Shares field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetShares(v VirtualizationVmwareSharesInfo) {
	o.Shares.Set(&v)
}

// SetSharesNil sets the value for Shares to be an explicit nil
func (o *VirtualizationVmwareVirtualDiskAllOf) SetSharesNil() {
	o.Shares.Set(nil)
}

// UnsetShares ensures that no value is present for Shares, not even an explicit nil
func (o *VirtualizationVmwareVirtualDiskAllOf) UnsetShares() {
	o.Shares.Unset()
}

// GetSharing returns the Sharing field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetSharing() string {
	if o == nil || o.Sharing == nil {
		var ret string
		return ret
	}
	return *o.Sharing
}

// GetSharingOk returns a tuple with the Sharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetSharingOk() (*string, bool) {
	if o == nil || o.Sharing == nil {
		return nil, false
	}
	return o.Sharing, true
}

// HasSharing returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasSharing() bool {
	if o != nil && o.Sharing != nil {
		return true
	}

	return false
}

// SetSharing gets a reference to the given string and assigns it to the Sharing field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetSharing(v string) {
	o.Sharing = &v
}

// GetStorageAllocationType returns the StorageAllocationType field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetStorageAllocationType() string {
	if o == nil || o.StorageAllocationType == nil {
		var ret string
		return ret
	}
	return *o.StorageAllocationType
}

// GetStorageAllocationTypeOk returns a tuple with the StorageAllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetStorageAllocationTypeOk() (*string, bool) {
	if o == nil || o.StorageAllocationType == nil {
		return nil, false
	}
	return o.StorageAllocationType, true
}

// HasStorageAllocationType returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasStorageAllocationType() bool {
	if o != nil && o.StorageAllocationType != nil {
		return true
	}

	return false
}

// SetStorageAllocationType gets a reference to the given string and assigns it to the StorageAllocationType field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetStorageAllocationType(v string) {
	o.StorageAllocationType = &v
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetUnitNumber() int64 {
	if o == nil || o.UnitNumber == nil {
		var ret int64
		return ret
	}
	return *o.UnitNumber
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetUnitNumberOk() (*int64, bool) {
	if o == nil || o.UnitNumber == nil {
		return nil, false
	}
	return o.UnitNumber, true
}

// HasUnitNumber returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasUnitNumber() bool {
	if o != nil && o.UnitNumber != nil {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given int64 and assigns it to the UnitNumber field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetUnitNumber(v int64) {
	o.UnitNumber = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVdiskId returns the VdiskId field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVdiskId() string {
	if o == nil || o.VdiskId == nil {
		var ret string
		return ret
	}
	return *o.VdiskId
}

// GetVdiskIdOk returns a tuple with the VdiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVdiskIdOk() (*string, bool) {
	if o == nil || o.VdiskId == nil {
		return nil, false
	}
	return o.VdiskId, true
}

// HasVdiskId returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasVdiskId() bool {
	if o != nil && o.VdiskId != nil {
		return true
	}

	return false
}

// SetVdiskId gets a reference to the given string and assigns it to the VdiskId field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetVdiskId(v string) {
	o.VdiskId = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetVirtualDiskPath returns the VirtualDiskPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVirtualDiskPath() string {
	if o == nil || o.VirtualDiskPath == nil {
		var ret string
		return ret
	}
	return *o.VirtualDiskPath
}

// GetVirtualDiskPathOk returns a tuple with the VirtualDiskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVirtualDiskPathOk() (*string, bool) {
	if o == nil || o.VirtualDiskPath == nil {
		return nil, false
	}
	return o.VirtualDiskPath, true
}

// HasVirtualDiskPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasVirtualDiskPath() bool {
	if o != nil && o.VirtualDiskPath != nil {
		return true
	}

	return false
}

// SetVirtualDiskPath gets a reference to the given string and assigns it to the VirtualDiskPath field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetVirtualDiskPath(v string) {
	o.VirtualDiskPath = &v
}

// GetVmIdentity returns the VmIdentity field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVmIdentity() string {
	if o == nil || o.VmIdentity == nil {
		var ret string
		return ret
	}
	return *o.VmIdentity
}

// GetVmIdentityOk returns a tuple with the VmIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVmIdentityOk() (*string, bool) {
	if o == nil || o.VmIdentity == nil {
		return nil, false
	}
	return o.VmIdentity, true
}

// HasVmIdentity returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasVmIdentity() bool {
	if o != nil && o.VmIdentity != nil {
		return true
	}

	return false
}

// SetVmIdentity gets a reference to the given string and assigns it to the VmIdentity field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetVmIdentity(v string) {
	o.VmIdentity = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDatastore() VirtualizationVmwareDatastoreRelationship {
	if o == nil || o.Datastore == nil {
		var ret VirtualizationVmwareDatastoreRelationship
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetDatastoreOk() (*VirtualizationVmwareDatastoreRelationship, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given VirtualizationVmwareDatastoreRelationship and assigns it to the Datastore field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetDatastore(v VirtualizationVmwareDatastoreRelationship) {
	o.Datastore = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVmwareVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualDiskAllOf) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVmwareVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VirtualizationVmwareVirtualDiskAllOf) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VirtualizationVmwareVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompatibilityMode != nil {
		toSerialize["CompatibilityMode"] = o.CompatibilityMode
	}
	if o.ControllerKey != nil {
		toSerialize["ControllerKey"] = o.ControllerKey
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DiskMode != nil {
		toSerialize["DiskMode"] = o.DiskMode
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Limit != nil {
		toSerialize["Limit"] = o.Limit
	}
	if o.LunUuid != nil {
		toSerialize["LunUuid"] = o.LunUuid
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Shares.IsSet() {
		toSerialize["Shares"] = o.Shares.Get()
	}
	if o.Sharing != nil {
		toSerialize["Sharing"] = o.Sharing
	}
	if o.StorageAllocationType != nil {
		toSerialize["StorageAllocationType"] = o.StorageAllocationType
	}
	if o.UnitNumber != nil {
		toSerialize["UnitNumber"] = o.UnitNumber
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VdiskId != nil {
		toSerialize["VdiskId"] = o.VdiskId
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.VirtualDiskPath != nil {
		toSerialize["VirtualDiskPath"] = o.VirtualDiskPath
	}
	if o.VmIdentity != nil {
		toSerialize["VmIdentity"] = o.VmIdentity
	}
	if o.Datastore != nil {
		toSerialize["Datastore"] = o.Datastore
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVirtualDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareVirtualDiskAllOf := _VirtualizationVmwareVirtualDiskAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualDiskAllOf); err == nil {
		*o = VirtualizationVmwareVirtualDiskAllOf(varVirtualizationVmwareVirtualDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompatibilityMode")
		delete(additionalProperties, "ControllerKey")
		delete(additionalProperties, "DeviceName")
		delete(additionalProperties, "DiskMode")
		delete(additionalProperties, "DiskType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Limit")
		delete(additionalProperties, "LunUuid")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Shares")
		delete(additionalProperties, "Sharing")
		delete(additionalProperties, "StorageAllocationType")
		delete(additionalProperties, "UnitNumber")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VdiskId")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "VirtualDiskPath")
		delete(additionalProperties, "VmIdentity")
		delete(additionalProperties, "Datastore")
		delete(additionalProperties, "VirtualMachine")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareVirtualDiskAllOf struct {
	value *VirtualizationVmwareVirtualDiskAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualDiskAllOf) Get() *VirtualizationVmwareVirtualDiskAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualDiskAllOf) Set(val *VirtualizationVmwareVirtualDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualDiskAllOf(val *VirtualizationVmwareVirtualDiskAllOf) *NullableVirtualizationVmwareVirtualDiskAllOf {
	return &NullableVirtualizationVmwareVirtualDiskAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
