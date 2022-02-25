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

// VirtualizationVirtualDiskAllOf Definition of the list of properties defined in 'virtualization.VirtualDisk', excluding properties defined in parent classes.
type VirtualizationVirtualDiskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Billing rate for this resource.
	BillingUnitId *string `json:"BillingUnitId,omitempty"`
	// Disk capacity to be provided with units example - 10Gi.
	Capacity *string `json:"Capacity,omitempty"`
	// Flag to indicate whether the configuration is created from inventory object.
	Discovered *bool `json:"Discovered,omitempty"`
	// Action to perform on the disk example resize, shrink, defragment etc.
	DiskAction *string `json:"DiskAction,omitempty"`
	// Encryption key used if volume is encrypted.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// Encryption method used to encrypt the volume.
	EncryptionType *string `json:"EncryptionType,omitempty"`
	// File mode of the disk  example - Filesystem, Block. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	Mode *string `json:"Mode,omitempty"`
	// Name of the storage disk. Name must be unique within a Datastore.
	Name *string `json:"Name,omitempty"`
	// Base64 encoded CA certificates of the https source to check against.
	SourceCerts *string `json:"SourceCerts,omitempty"`
	// Source disk from which the content is copied.
	SourceDiskToClone *string `json:"SourceDiskToClone,omitempty"`
	// Image path used to import on the created disk.
	SourceFilePath       *string                                    `json:"SourceFilePath,omitempty"`
	VolumeIopsInfo       NullableCloudVolumeIopsInfo                `json:"VolumeIopsInfo,omitempty"`
	Zone                 NullableCloudAvailabilityZone              `json:"Zone,omitempty"`
	Cluster              *VirtualizationBaseClusterRelationship     `json:"Cluster,omitempty"`
	Inventory            *VirtualizationBaseVirtualDiskRelationship `json:"Inventory,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship       `json:"RegisteredDevice,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship          `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVirtualDiskAllOf VirtualizationVirtualDiskAllOf

// NewVirtualizationVirtualDiskAllOf instantiates a new VirtualizationVirtualDiskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVirtualDiskAllOf(classId string, objectType string) *VirtualizationVirtualDiskAllOf {
	this := VirtualizationVirtualDiskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// NewVirtualizationVirtualDiskAllOfWithDefaults instantiates a new VirtualizationVirtualDiskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVirtualDiskAllOfWithDefaults() *VirtualizationVirtualDiskAllOf {
	this := VirtualizationVirtualDiskAllOf{}
	var classId string = "virtualization.VirtualDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VirtualDisk"
	this.ObjectType = objectType
	var mode string = "Block"
	this.Mode = &mode
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVirtualDiskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVirtualDiskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVirtualDiskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVirtualDiskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingUnitId returns the BillingUnitId field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetBillingUnitId() string {
	if o == nil || o.BillingUnitId == nil {
		var ret string
		return ret
	}
	return *o.BillingUnitId
}

// GetBillingUnitIdOk returns a tuple with the BillingUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetBillingUnitIdOk() (*string, bool) {
	if o == nil || o.BillingUnitId == nil {
		return nil, false
	}
	return o.BillingUnitId, true
}

// HasBillingUnitId returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasBillingUnitId() bool {
	if o != nil && o.BillingUnitId != nil {
		return true
	}

	return false
}

// SetBillingUnitId gets a reference to the given string and assigns it to the BillingUnitId field.
func (o *VirtualizationVirtualDiskAllOf) SetBillingUnitId(v string) {
	o.BillingUnitId = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *VirtualizationVirtualDiskAllOf) SetCapacity(v string) {
	o.Capacity = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *VirtualizationVirtualDiskAllOf) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetDiskAction returns the DiskAction field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetDiskAction() string {
	if o == nil || o.DiskAction == nil {
		var ret string
		return ret
	}
	return *o.DiskAction
}

// GetDiskActionOk returns a tuple with the DiskAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetDiskActionOk() (*string, bool) {
	if o == nil || o.DiskAction == nil {
		return nil, false
	}
	return o.DiskAction, true
}

// HasDiskAction returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasDiskAction() bool {
	if o != nil && o.DiskAction != nil {
		return true
	}

	return false
}

// SetDiskAction gets a reference to the given string and assigns it to the DiskAction field.
func (o *VirtualizationVirtualDiskAllOf) SetDiskAction(v string) {
	o.DiskAction = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *VirtualizationVirtualDiskAllOf) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetEncryptionType returns the EncryptionType field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetEncryptionType() string {
	if o == nil || o.EncryptionType == nil {
		var ret string
		return ret
	}
	return *o.EncryptionType
}

// GetEncryptionTypeOk returns a tuple with the EncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetEncryptionTypeOk() (*string, bool) {
	if o == nil || o.EncryptionType == nil {
		return nil, false
	}
	return o.EncryptionType, true
}

// HasEncryptionType returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasEncryptionType() bool {
	if o != nil && o.EncryptionType != nil {
		return true
	}

	return false
}

// SetEncryptionType gets a reference to the given string and assigns it to the EncryptionType field.
func (o *VirtualizationVirtualDiskAllOf) SetEncryptionType(v string) {
	o.EncryptionType = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VirtualizationVirtualDiskAllOf) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVirtualDiskAllOf) SetName(v string) {
	o.Name = &v
}

// GetSourceCerts returns the SourceCerts field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetSourceCerts() string {
	if o == nil || o.SourceCerts == nil {
		var ret string
		return ret
	}
	return *o.SourceCerts
}

// GetSourceCertsOk returns a tuple with the SourceCerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetSourceCertsOk() (*string, bool) {
	if o == nil || o.SourceCerts == nil {
		return nil, false
	}
	return o.SourceCerts, true
}

// HasSourceCerts returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasSourceCerts() bool {
	if o != nil && o.SourceCerts != nil {
		return true
	}

	return false
}

// SetSourceCerts gets a reference to the given string and assigns it to the SourceCerts field.
func (o *VirtualizationVirtualDiskAllOf) SetSourceCerts(v string) {
	o.SourceCerts = &v
}

// GetSourceDiskToClone returns the SourceDiskToClone field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToClone() string {
	if o == nil || o.SourceDiskToClone == nil {
		var ret string
		return ret
	}
	return *o.SourceDiskToClone
}

// GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetSourceDiskToCloneOk() (*string, bool) {
	if o == nil || o.SourceDiskToClone == nil {
		return nil, false
	}
	return o.SourceDiskToClone, true
}

// HasSourceDiskToClone returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasSourceDiskToClone() bool {
	if o != nil && o.SourceDiskToClone != nil {
		return true
	}

	return false
}

// SetSourceDiskToClone gets a reference to the given string and assigns it to the SourceDiskToClone field.
func (o *VirtualizationVirtualDiskAllOf) SetSourceDiskToClone(v string) {
	o.SourceDiskToClone = &v
}

// GetSourceFilePath returns the SourceFilePath field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePath() string {
	if o == nil || o.SourceFilePath == nil {
		var ret string
		return ret
	}
	return *o.SourceFilePath
}

// GetSourceFilePathOk returns a tuple with the SourceFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetSourceFilePathOk() (*string, bool) {
	if o == nil || o.SourceFilePath == nil {
		return nil, false
	}
	return o.SourceFilePath, true
}

// HasSourceFilePath returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasSourceFilePath() bool {
	if o != nil && o.SourceFilePath != nil {
		return true
	}

	return false
}

// SetSourceFilePath gets a reference to the given string and assigns it to the SourceFilePath field.
func (o *VirtualizationVirtualDiskAllOf) SetSourceFilePath(v string) {
	o.SourceFilePath = &v
}

// GetVolumeIopsInfo returns the VolumeIopsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVirtualDiskAllOf) GetVolumeIopsInfo() CloudVolumeIopsInfo {
	if o == nil || o.VolumeIopsInfo.Get() == nil {
		var ret CloudVolumeIopsInfo
		return ret
	}
	return *o.VolumeIopsInfo.Get()
}

// GetVolumeIopsInfoOk returns a tuple with the VolumeIopsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVirtualDiskAllOf) GetVolumeIopsInfoOk() (*CloudVolumeIopsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeIopsInfo.Get(), o.VolumeIopsInfo.IsSet()
}

// HasVolumeIopsInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasVolumeIopsInfo() bool {
	if o != nil && o.VolumeIopsInfo.IsSet() {
		return true
	}

	return false
}

// SetVolumeIopsInfo gets a reference to the given NullableCloudVolumeIopsInfo and assigns it to the VolumeIopsInfo field.
func (o *VirtualizationVirtualDiskAllOf) SetVolumeIopsInfo(v CloudVolumeIopsInfo) {
	o.VolumeIopsInfo.Set(&v)
}

// SetVolumeIopsInfoNil sets the value for VolumeIopsInfo to be an explicit nil
func (o *VirtualizationVirtualDiskAllOf) SetVolumeIopsInfoNil() {
	o.VolumeIopsInfo.Set(nil)
}

// UnsetVolumeIopsInfo ensures that no value is present for VolumeIopsInfo, not even an explicit nil
func (o *VirtualizationVirtualDiskAllOf) UnsetVolumeIopsInfo() {
	o.VolumeIopsInfo.Unset()
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVirtualDiskAllOf) GetZone() CloudAvailabilityZone {
	if o == nil || o.Zone.Get() == nil {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVirtualDiskAllOf) GetZoneOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableCloudAvailabilityZone and assigns it to the Zone field.
func (o *VirtualizationVirtualDiskAllOf) SetZone(v CloudAvailabilityZone) {
	o.Zone.Set(&v)
}

// SetZoneNil sets the value for Zone to be an explicit nil
func (o *VirtualizationVirtualDiskAllOf) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *VirtualizationVirtualDiskAllOf) UnsetZone() {
	o.Zone.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetCluster() VirtualizationBaseClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationBaseClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationBaseClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVirtualDiskAllOf) SetCluster(v VirtualizationBaseClusterRelationship) {
	o.Cluster = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetInventory() VirtualizationBaseVirtualDiskRelationship {
	if o == nil || o.Inventory == nil {
		var ret VirtualizationBaseVirtualDiskRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetInventoryOk() (*VirtualizationBaseVirtualDiskRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given VirtualizationBaseVirtualDiskRelationship and assigns it to the Inventory field.
func (o *VirtualizationVirtualDiskAllOf) SetInventory(v VirtualizationBaseVirtualDiskRelationship) {
	o.Inventory = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVirtualDiskAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVirtualDiskAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *VirtualizationVirtualDiskAllOf) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *VirtualizationVirtualDiskAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o VirtualizationVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingUnitId != nil {
		toSerialize["BillingUnitId"] = o.BillingUnitId
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.DiskAction != nil {
		toSerialize["DiskAction"] = o.DiskAction
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.EncryptionType != nil {
		toSerialize["EncryptionType"] = o.EncryptionType
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SourceCerts != nil {
		toSerialize["SourceCerts"] = o.SourceCerts
	}
	if o.SourceDiskToClone != nil {
		toSerialize["SourceDiskToClone"] = o.SourceDiskToClone
	}
	if o.SourceFilePath != nil {
		toSerialize["SourceFilePath"] = o.SourceFilePath
	}
	if o.VolumeIopsInfo.IsSet() {
		toSerialize["VolumeIopsInfo"] = o.VolumeIopsInfo.Get()
	}
	if o.Zone.IsSet() {
		toSerialize["Zone"] = o.Zone.Get()
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVirtualDiskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVirtualDiskAllOf := _VirtualizationVirtualDiskAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVirtualDiskAllOf); err == nil {
		*o = VirtualizationVirtualDiskAllOf(varVirtualizationVirtualDiskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnitId")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Discovered")
		delete(additionalProperties, "DiskAction")
		delete(additionalProperties, "EncryptionKey")
		delete(additionalProperties, "EncryptionType")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SourceCerts")
		delete(additionalProperties, "SourceDiskToClone")
		delete(additionalProperties, "SourceFilePath")
		delete(additionalProperties, "VolumeIopsInfo")
		delete(additionalProperties, "Zone")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Inventory")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "WorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVirtualDiskAllOf struct {
	value *VirtualizationVirtualDiskAllOf
	isSet bool
}

func (v NullableVirtualizationVirtualDiskAllOf) Get() *VirtualizationVirtualDiskAllOf {
	return v.value
}

func (v *NullableVirtualizationVirtualDiskAllOf) Set(val *VirtualizationVirtualDiskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVirtualDiskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVirtualDiskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVirtualDiskAllOf(val *VirtualizationVirtualDiskAllOf) *NullableVirtualizationVirtualDiskAllOf {
	return &NullableVirtualizationVirtualDiskAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVirtualDiskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVirtualDiskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
