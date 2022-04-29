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

// HyperflexDrive A Hyperflex drive entity attached to a node in a Hyperflex cluster.
type HyperflexDrive struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Provisioned capacity of the drive in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Disk inventory state. Should be one of values defined in enum. * `UNKNOWN` - The use state of the disk is unknown. * `USED` - The use state of the disk is used. * `NOTUSED` - The use state of the disk is unused. * `EMPTY` - The use state of the disk is empty.
	DiskUseState *string `json:"DiskUseState,omitempty"`
	// Host Name.
	HostName *string `json:"HostName,omitempty"`
	// The unique identifier of the drive's host.
	HostUuid *string `json:"HostUuid,omitempty"`
	// The model number of the disk or SSD.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// The unique identifier of the Hyperflex Node to which the disk is attached.
	NodeUuid *string `json:"NodeUuid,omitempty"`
	// Device path for the drive.
	Path *string `json:"Path,omitempty"`
	// Disk Protocol - SATA, NVME, SAS. * `Unknown` - Disk protocol is unknown. * `SAS` - Serial Attached SCSI protocol (SAS) used in disk. * `NVMe` - Non-volatile memory express (NVMe) protocol used in disk. * `SATA` - Serial Advanced Technology Attachment (SATA) used in disk.
	Protocol *string `json:"Protocol,omitempty"`
	// Serial number of the Hyperflex drive.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The SCSI slot number of the drive.
	SlotNumber *string `json:"SlotNumber,omitempty"`
	// Disk inventory state as determined by storfs inventory module. Should be one of values defined in enum. * `UNKNOWN` - The state of the disk is unknown. * `CLAIMED` - The state of the disk is claimed by storfs and has a valid storfs format. * `AVAILABLE` - The disk is available but not claimed by storfs. * `IGNORED` - The disk ash been ignored by storfs. * `BLACKLISTED` - The disk has been blacklisted by storfs. * `SECUREERASED` - The disk has been secure erased. * `BLOCKED` - The disk has been blocked by storfs.
	Status *string `json:"Status,omitempty"`
	// Type of disk - UNKNOWN, HDD, SSD. * `Unknown` - Default unknown disk type. * `SSD` - Storage disk with Solid state disk. * `HDD` - Storage disk with Hard disk drive. * `NVRAM` - Storage disk with Non-volatile random-access memory type. * `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA). * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FC` - Storage disk with Fiber Channel. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SAS` - Storage disk with serial attached SCSI. * `VMDISK` - Virtual machine Data Disk.
	Type *string `json:"Type,omitempty"`
	// Specifies what the disk is used for. Should be one of values defined in enum. * `UNKNOWN` - The usage of the disk is unknown. * `PERSISTENCE` - The usage of the disk is for generic data storage. * `SYSTEM` - The usage of the disk is related to system storage. * `CACHING` - The usage of the disk is for caching.
	Usage *string `json:"Usage,omitempty"`
	// Used Capacity of the drive in Bytes.
	UsedCapacity *int64 `json:"UsedCapacity,omitempty"`
	// The unique identifier of the Hyperflex drive.
	Uuid *string `json:"Uuid,omitempty"`
	// The firmware version of the Hyperflex drive.
	Version              *string                          `json:"Version,omitempty"`
	LocatorLed           *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	Node                 *HyperflexNodeRelationship       `json:"Node,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexDrive HyperflexDrive

// NewHyperflexDrive instantiates a new HyperflexDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexDrive(classId string, objectType string) *HyperflexDrive {
	this := HyperflexDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexDriveWithDefaults instantiates a new HyperflexDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexDriveWithDefaults() *HyperflexDrive {
	this := HyperflexDrive{}
	var classId string = "hyperflex.Drive"
	this.ClassId = classId
	var objectType string = "hyperflex.Drive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *HyperflexDrive) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *HyperflexDrive) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *HyperflexDrive) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetDiskUseState returns the DiskUseState field value if set, zero value otherwise.
func (o *HyperflexDrive) GetDiskUseState() string {
	if o == nil || o.DiskUseState == nil {
		var ret string
		return ret
	}
	return *o.DiskUseState
}

// GetDiskUseStateOk returns a tuple with the DiskUseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetDiskUseStateOk() (*string, bool) {
	if o == nil || o.DiskUseState == nil {
		return nil, false
	}
	return o.DiskUseState, true
}

// HasDiskUseState returns a boolean if a field has been set.
func (o *HyperflexDrive) HasDiskUseState() bool {
	if o != nil && o.DiskUseState != nil {
		return true
	}

	return false
}

// SetDiskUseState gets a reference to the given string and assigns it to the DiskUseState field.
func (o *HyperflexDrive) SetDiskUseState(v string) {
	o.DiskUseState = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexDrive) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexDrive) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexDrive) SetHostName(v string) {
	o.HostName = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *HyperflexDrive) GetHostUuid() string {
	if o == nil || o.HostUuid == nil {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetHostUuidOk() (*string, bool) {
	if o == nil || o.HostUuid == nil {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *HyperflexDrive) HasHostUuid() bool {
	if o != nil && o.HostUuid != nil {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *HyperflexDrive) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *HyperflexDrive) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *HyperflexDrive) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *HyperflexDrive) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetNodeUuid returns the NodeUuid field value if set, zero value otherwise.
func (o *HyperflexDrive) GetNodeUuid() string {
	if o == nil || o.NodeUuid == nil {
		var ret string
		return ret
	}
	return *o.NodeUuid
}

// GetNodeUuidOk returns a tuple with the NodeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetNodeUuidOk() (*string, bool) {
	if o == nil || o.NodeUuid == nil {
		return nil, false
	}
	return o.NodeUuid, true
}

// HasNodeUuid returns a boolean if a field has been set.
func (o *HyperflexDrive) HasNodeUuid() bool {
	if o != nil && o.NodeUuid != nil {
		return true
	}

	return false
}

// SetNodeUuid gets a reference to the given string and assigns it to the NodeUuid field.
func (o *HyperflexDrive) SetNodeUuid(v string) {
	o.NodeUuid = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HyperflexDrive) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HyperflexDrive) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HyperflexDrive) SetPath(v string) {
	o.Path = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *HyperflexDrive) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *HyperflexDrive) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *HyperflexDrive) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexDrive) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexDrive) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexDrive) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSlotNumber returns the SlotNumber field value if set, zero value otherwise.
func (o *HyperflexDrive) GetSlotNumber() string {
	if o == nil || o.SlotNumber == nil {
		var ret string
		return ret
	}
	return *o.SlotNumber
}

// GetSlotNumberOk returns a tuple with the SlotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetSlotNumberOk() (*string, bool) {
	if o == nil || o.SlotNumber == nil {
		return nil, false
	}
	return o.SlotNumber, true
}

// HasSlotNumber returns a boolean if a field has been set.
func (o *HyperflexDrive) HasSlotNumber() bool {
	if o != nil && o.SlotNumber != nil {
		return true
	}

	return false
}

// SetSlotNumber gets a reference to the given string and assigns it to the SlotNumber field.
func (o *HyperflexDrive) SetSlotNumber(v string) {
	o.SlotNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexDrive) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexDrive) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexDrive) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HyperflexDrive) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HyperflexDrive) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HyperflexDrive) SetType(v string) {
	o.Type = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *HyperflexDrive) GetUsage() string {
	if o == nil || o.Usage == nil {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetUsageOk() (*string, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *HyperflexDrive) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *HyperflexDrive) SetUsage(v string) {
	o.Usage = &v
}

// GetUsedCapacity returns the UsedCapacity field value if set, zero value otherwise.
func (o *HyperflexDrive) GetUsedCapacity() int64 {
	if o == nil || o.UsedCapacity == nil {
		var ret int64
		return ret
	}
	return *o.UsedCapacity
}

// GetUsedCapacityOk returns a tuple with the UsedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetUsedCapacityOk() (*int64, bool) {
	if o == nil || o.UsedCapacity == nil {
		return nil, false
	}
	return o.UsedCapacity, true
}

// HasUsedCapacity returns a boolean if a field has been set.
func (o *HyperflexDrive) HasUsedCapacity() bool {
	if o != nil && o.UsedCapacity != nil {
		return true
	}

	return false
}

// SetUsedCapacity gets a reference to the given int64 and assigns it to the UsedCapacity field.
func (o *HyperflexDrive) SetUsedCapacity(v int64) {
	o.UsedCapacity = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexDrive) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexDrive) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexDrive) SetUuid(v string) {
	o.Uuid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexDrive) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexDrive) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexDrive) SetVersion(v string) {
	o.Version = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *HyperflexDrive) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *HyperflexDrive) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *HyperflexDrive) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *HyperflexDrive) GetNode() HyperflexNodeRelationship {
	if o == nil || o.Node == nil {
		var ret HyperflexNodeRelationship
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexDrive) GetNodeOk() (*HyperflexNodeRelationship, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *HyperflexDrive) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given HyperflexNodeRelationship and assigns it to the Node field.
func (o *HyperflexDrive) SetNode(v HyperflexNodeRelationship) {
	o.Node = &v
}

func (o HyperflexDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.DiskUseState != nil {
		toSerialize["DiskUseState"] = o.DiskUseState
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.HostUuid != nil {
		toSerialize["HostUuid"] = o.HostUuid
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.NodeUuid != nil {
		toSerialize["NodeUuid"] = o.NodeUuid
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.SlotNumber != nil {
		toSerialize["SlotNumber"] = o.SlotNumber
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.UsedCapacity != nil {
		toSerialize["UsedCapacity"] = o.UsedCapacity
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.Node != nil {
		toSerialize["Node"] = o.Node
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexDrive) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Provisioned capacity of the drive in bytes.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Disk inventory state. Should be one of values defined in enum. * `UNKNOWN` - The use state of the disk is unknown. * `USED` - The use state of the disk is used. * `NOTUSED` - The use state of the disk is unused. * `EMPTY` - The use state of the disk is empty.
		DiskUseState *string `json:"DiskUseState,omitempty"`
		// Host Name.
		HostName *string `json:"HostName,omitempty"`
		// The unique identifier of the drive's host.
		HostUuid *string `json:"HostUuid,omitempty"`
		// The model number of the disk or SSD.
		ModelNumber *string `json:"ModelNumber,omitempty"`
		// The unique identifier of the Hyperflex Node to which the disk is attached.
		NodeUuid *string `json:"NodeUuid,omitempty"`
		// Device path for the drive.
		Path *string `json:"Path,omitempty"`
		// Disk Protocol - SATA, NVME, SAS. * `Unknown` - Disk protocol is unknown. * `SAS` - Serial Attached SCSI protocol (SAS) used in disk. * `NVMe` - Non-volatile memory express (NVMe) protocol used in disk. * `SATA` - Serial Advanced Technology Attachment (SATA) used in disk.
		Protocol *string `json:"Protocol,omitempty"`
		// Serial number of the Hyperflex drive.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The SCSI slot number of the drive.
		SlotNumber *string `json:"SlotNumber,omitempty"`
		// Disk inventory state as determined by storfs inventory module. Should be one of values defined in enum. * `UNKNOWN` - The state of the disk is unknown. * `CLAIMED` - The state of the disk is claimed by storfs and has a valid storfs format. * `AVAILABLE` - The disk is available but not claimed by storfs. * `IGNORED` - The disk ash been ignored by storfs. * `BLACKLISTED` - The disk has been blacklisted by storfs. * `SECUREERASED` - The disk has been secure erased. * `BLOCKED` - The disk has been blocked by storfs.
		Status *string `json:"Status,omitempty"`
		// Type of disk - UNKNOWN, HDD, SSD. * `Unknown` - Default unknown disk type. * `SSD` - Storage disk with Solid state disk. * `HDD` - Storage disk with Hard disk drive. * `NVRAM` - Storage disk with Non-volatile random-access memory type. * `SATA` - Disk drive implementation with Serial Advanced Technology Attachment (SATA). * `BSAS` - Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf. * `FC` - Storage disk with Fiber Channel. * `FSAS` - Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives. * `LUN` - Logical Unit Number refers to a logical disk. * `MSATA` - SATA disk in multi-disk carrier storage shelf. * `SAS` - Storage disk with serial attached SCSI. * `VMDISK` - Virtual machine Data Disk.
		Type *string `json:"Type,omitempty"`
		// Specifies what the disk is used for. Should be one of values defined in enum. * `UNKNOWN` - The usage of the disk is unknown. * `PERSISTENCE` - The usage of the disk is for generic data storage. * `SYSTEM` - The usage of the disk is related to system storage. * `CACHING` - The usage of the disk is for caching.
		Usage *string `json:"Usage,omitempty"`
		// Used Capacity of the drive in Bytes.
		UsedCapacity *int64 `json:"UsedCapacity,omitempty"`
		// The unique identifier of the Hyperflex drive.
		Uuid *string `json:"Uuid,omitempty"`
		// The firmware version of the Hyperflex drive.
		Version    *string                          `json:"Version,omitempty"`
		LocatorLed *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
		Node       *HyperflexNodeRelationship       `json:"Node,omitempty"`
	}

	varHyperflexDriveWithoutEmbeddedStruct := HyperflexDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexDriveWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexDrive := _HyperflexDrive{}
		varHyperflexDrive.ClassId = varHyperflexDriveWithoutEmbeddedStruct.ClassId
		varHyperflexDrive.ObjectType = varHyperflexDriveWithoutEmbeddedStruct.ObjectType
		varHyperflexDrive.Capacity = varHyperflexDriveWithoutEmbeddedStruct.Capacity
		varHyperflexDrive.DiskUseState = varHyperflexDriveWithoutEmbeddedStruct.DiskUseState
		varHyperflexDrive.HostName = varHyperflexDriveWithoutEmbeddedStruct.HostName
		varHyperflexDrive.HostUuid = varHyperflexDriveWithoutEmbeddedStruct.HostUuid
		varHyperflexDrive.ModelNumber = varHyperflexDriveWithoutEmbeddedStruct.ModelNumber
		varHyperflexDrive.NodeUuid = varHyperflexDriveWithoutEmbeddedStruct.NodeUuid
		varHyperflexDrive.Path = varHyperflexDriveWithoutEmbeddedStruct.Path
		varHyperflexDrive.Protocol = varHyperflexDriveWithoutEmbeddedStruct.Protocol
		varHyperflexDrive.SerialNumber = varHyperflexDriveWithoutEmbeddedStruct.SerialNumber
		varHyperflexDrive.SlotNumber = varHyperflexDriveWithoutEmbeddedStruct.SlotNumber
		varHyperflexDrive.Status = varHyperflexDriveWithoutEmbeddedStruct.Status
		varHyperflexDrive.Type = varHyperflexDriveWithoutEmbeddedStruct.Type
		varHyperflexDrive.Usage = varHyperflexDriveWithoutEmbeddedStruct.Usage
		varHyperflexDrive.UsedCapacity = varHyperflexDriveWithoutEmbeddedStruct.UsedCapacity
		varHyperflexDrive.Uuid = varHyperflexDriveWithoutEmbeddedStruct.Uuid
		varHyperflexDrive.Version = varHyperflexDriveWithoutEmbeddedStruct.Version
		varHyperflexDrive.LocatorLed = varHyperflexDriveWithoutEmbeddedStruct.LocatorLed
		varHyperflexDrive.Node = varHyperflexDriveWithoutEmbeddedStruct.Node
		*o = HyperflexDrive(varHyperflexDrive)
	} else {
		return err
	}

	varHyperflexDrive := _HyperflexDrive{}

	err = json.Unmarshal(bytes, &varHyperflexDrive)
	if err == nil {
		o.MoBaseMo = varHyperflexDrive.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "DiskUseState")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "HostUuid")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "NodeUuid")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "SlotNumber")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "UsedCapacity")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "Node")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexDrive struct {
	value *HyperflexDrive
	isSet bool
}

func (v NullableHyperflexDrive) Get() *HyperflexDrive {
	return v.value
}

func (v *NullableHyperflexDrive) Set(val *HyperflexDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexDrive(val *HyperflexDrive) *NullableHyperflexDrive {
	return &NullableHyperflexDrive{value: val, isSet: true}
}

func (v NullableHyperflexDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
