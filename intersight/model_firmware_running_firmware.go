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
	"reflect"
	"strings"
)

// FirmwareRunningFirmware Running Firmware on an endpoint.
type FirmwareRunningFirmware struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Kind of the firmware - boot-booloader/system/kernel.
	Component *string `json:"Component,omitempty"`
	// Package version which the firmware belongs to.
	PackageVersion *string `json:"PackageVersion,omitempty"`
	// The type of the firmware.
	Type *string `json:"Type,omitempty"`
	// The version of the firmware.
	Version              *string                           `json:"Version,omitempty"`
	BiosUnit             *BiosUnitRelationship             `json:"BiosUnit,omitempty"`
	GraphicsCard         *GraphicsCardRelationship         `json:"GraphicsCard,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
	ManagementController *ManagementControllerRelationship `json:"ManagementController,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements            []NetworkElementRelationship            `json:"NetworkElements,omitempty"`
	PciSwitch                  *PciSwitchRelationship                  `json:"PciSwitch,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	StorageController          *StorageControllerRelationship          `json:"StorageController,omitempty"`
	StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
	StoragePhysicalDisk        *StoragePhysicalDiskRelationship        `json:"StoragePhysicalDisk,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _FirmwareRunningFirmware FirmwareRunningFirmware

// NewFirmwareRunningFirmware instantiates a new FirmwareRunningFirmware object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareRunningFirmware(classId string, objectType string) *FirmwareRunningFirmware {
	this := FirmwareRunningFirmware{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareRunningFirmwareWithDefaults instantiates a new FirmwareRunningFirmware object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareRunningFirmwareWithDefaults() *FirmwareRunningFirmware {
	this := FirmwareRunningFirmware{}
	var classId string = "firmware.RunningFirmware"
	this.ClassId = classId
	var objectType string = "firmware.RunningFirmware"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareRunningFirmware) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareRunningFirmware) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareRunningFirmware) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareRunningFirmware) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *FirmwareRunningFirmware) SetComponent(v string) {
	o.Component = &v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetPackageVersion() string {
	if o == nil || o.PackageVersion == nil {
		var ret string
		return ret
	}
	return *o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetPackageVersionOk() (*string, bool) {
	if o == nil || o.PackageVersion == nil {
		return nil, false
	}
	return o.PackageVersion, true
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasPackageVersion() bool {
	if o != nil && o.PackageVersion != nil {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given string and assigns it to the PackageVersion field.
func (o *FirmwareRunningFirmware) SetPackageVersion(v string) {
	o.PackageVersion = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FirmwareRunningFirmware) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FirmwareRunningFirmware) SetVersion(v string) {
	o.Version = &v
}

// GetBiosUnit returns the BiosUnit field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetBiosUnit() BiosUnitRelationship {
	if o == nil || o.BiosUnit == nil {
		var ret BiosUnitRelationship
		return ret
	}
	return *o.BiosUnit
}

// GetBiosUnitOk returns a tuple with the BiosUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetBiosUnitOk() (*BiosUnitRelationship, bool) {
	if o == nil || o.BiosUnit == nil {
		return nil, false
	}
	return o.BiosUnit, true
}

// HasBiosUnit returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasBiosUnit() bool {
	if o != nil && o.BiosUnit != nil {
		return true
	}

	return false
}

// SetBiosUnit gets a reference to the given BiosUnitRelationship and assigns it to the BiosUnit field.
func (o *FirmwareRunningFirmware) SetBiosUnit(v BiosUnitRelationship) {
	o.BiosUnit = &v
}

// GetGraphicsCard returns the GraphicsCard field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetGraphicsCard() GraphicsCardRelationship {
	if o == nil || o.GraphicsCard == nil {
		var ret GraphicsCardRelationship
		return ret
	}
	return *o.GraphicsCard
}

// GetGraphicsCardOk returns a tuple with the GraphicsCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetGraphicsCardOk() (*GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCard == nil {
		return nil, false
	}
	return o.GraphicsCard, true
}

// HasGraphicsCard returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasGraphicsCard() bool {
	if o != nil && o.GraphicsCard != nil {
		return true
	}

	return false
}

// SetGraphicsCard gets a reference to the given GraphicsCardRelationship and assigns it to the GraphicsCard field.
func (o *FirmwareRunningFirmware) SetGraphicsCard(v GraphicsCardRelationship) {
	o.GraphicsCard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *FirmwareRunningFirmware) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetManagementController returns the ManagementController field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetManagementController() ManagementControllerRelationship {
	if o == nil || o.ManagementController == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementController
}

// GetManagementControllerOk returns a tuple with the ManagementController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetManagementControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.ManagementController == nil {
		return nil, false
	}
	return o.ManagementController, true
}

// HasManagementController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasManagementController() bool {
	if o != nil && o.ManagementController != nil {
		return true
	}

	return false
}

// SetManagementController gets a reference to the given ManagementControllerRelationship and assigns it to the ManagementController field.
func (o *FirmwareRunningFirmware) SetManagementController(v ManagementControllerRelationship) {
	o.ManagementController = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareRunningFirmware) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareRunningFirmware) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareRunningFirmware) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetPciSwitch returns the PciSwitch field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetPciSwitch() PciSwitchRelationship {
	if o == nil || o.PciSwitch == nil {
		var ret PciSwitchRelationship
		return ret
	}
	return *o.PciSwitch
}

// GetPciSwitchOk returns a tuple with the PciSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetPciSwitchOk() (*PciSwitchRelationship, bool) {
	if o == nil || o.PciSwitch == nil {
		return nil, false
	}
	return o.PciSwitch, true
}

// HasPciSwitch returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasPciSwitch() bool {
	if o != nil && o.PciSwitch != nil {
		return true
	}

	return false
}

// SetPciSwitch gets a reference to the given PciSwitchRelationship and assigns it to the PciSwitch field.
func (o *FirmwareRunningFirmware) SetPciSwitch(v PciSwitchRelationship) {
	o.PciSwitch = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FirmwareRunningFirmware) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *FirmwareRunningFirmware) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

// GetStorageFlexFlashController returns the StorageFlexFlashController field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetStorageFlexFlashController() StorageFlexFlashControllerRelationship {
	if o == nil || o.StorageFlexFlashController == nil {
		var ret StorageFlexFlashControllerRelationship
		return ret
	}
	return *o.StorageFlexFlashController
}

// GetStorageFlexFlashControllerOk returns a tuple with the StorageFlexFlashController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetStorageFlexFlashControllerOk() (*StorageFlexFlashControllerRelationship, bool) {
	if o == nil || o.StorageFlexFlashController == nil {
		return nil, false
	}
	return o.StorageFlexFlashController, true
}

// HasStorageFlexFlashController returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStorageFlexFlashController() bool {
	if o != nil && o.StorageFlexFlashController != nil {
		return true
	}

	return false
}

// SetStorageFlexFlashController gets a reference to the given StorageFlexFlashControllerRelationship and assigns it to the StorageFlexFlashController field.
func (o *FirmwareRunningFirmware) SetStorageFlexFlashController(v StorageFlexFlashControllerRelationship) {
	o.StorageFlexFlashController = &v
}

// GetStoragePhysicalDisk returns the StoragePhysicalDisk field value if set, zero value otherwise.
func (o *FirmwareRunningFirmware) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.StoragePhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.StoragePhysicalDisk
}

// GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareRunningFirmware) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.StoragePhysicalDisk == nil {
		return nil, false
	}
	return o.StoragePhysicalDisk, true
}

// HasStoragePhysicalDisk returns a boolean if a field has been set.
func (o *FirmwareRunningFirmware) HasStoragePhysicalDisk() bool {
	if o != nil && o.StoragePhysicalDisk != nil {
		return true
	}

	return false
}

// SetStoragePhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the StoragePhysicalDisk field.
func (o *FirmwareRunningFirmware) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.StoragePhysicalDisk = &v
}

func (o FirmwareRunningFirmware) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Component != nil {
		toSerialize["Component"] = o.Component
	}
	if o.PackageVersion != nil {
		toSerialize["PackageVersion"] = o.PackageVersion
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.BiosUnit != nil {
		toSerialize["BiosUnit"] = o.BiosUnit
	}
	if o.GraphicsCard != nil {
		toSerialize["GraphicsCard"] = o.GraphicsCard
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.ManagementController != nil {
		toSerialize["ManagementController"] = o.ManagementController
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.PciSwitch != nil {
		toSerialize["PciSwitch"] = o.PciSwitch
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}
	if o.StorageFlexFlashController != nil {
		toSerialize["StorageFlexFlashController"] = o.StorageFlexFlashController
	}
	if o.StoragePhysicalDisk != nil {
		toSerialize["StoragePhysicalDisk"] = o.StoragePhysicalDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareRunningFirmware) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareRunningFirmwareWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Kind of the firmware - boot-booloader/system/kernel.
		Component *string `json:"Component,omitempty"`
		// Package version which the firmware belongs to.
		PackageVersion *string `json:"PackageVersion,omitempty"`
		// The type of the firmware.
		Type *string `json:"Type,omitempty"`
		// The version of the firmware.
		Version              *string                           `json:"Version,omitempty"`
		BiosUnit             *BiosUnitRelationship             `json:"BiosUnit,omitempty"`
		GraphicsCard         *GraphicsCardRelationship         `json:"GraphicsCard,omitempty"`
		InventoryDeviceInfo  *InventoryDeviceInfoRelationship  `json:"InventoryDeviceInfo,omitempty"`
		ManagementController *ManagementControllerRelationship `json:"ManagementController,omitempty"`
		// An array of relationships to networkElement resources.
		NetworkElements            []NetworkElementRelationship            `json:"NetworkElements,omitempty"`
		PciSwitch                  *PciSwitchRelationship                  `json:"PciSwitch,omitempty"`
		RegisteredDevice           *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
		StorageController          *StorageControllerRelationship          `json:"StorageController,omitempty"`
		StorageFlexFlashController *StorageFlexFlashControllerRelationship `json:"StorageFlexFlashController,omitempty"`
		StoragePhysicalDisk        *StoragePhysicalDiskRelationship        `json:"StoragePhysicalDisk,omitempty"`
	}

	varFirmwareRunningFirmwareWithoutEmbeddedStruct := FirmwareRunningFirmwareWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareRunningFirmwareWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareRunningFirmware := _FirmwareRunningFirmware{}
		varFirmwareRunningFirmware.ClassId = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ClassId
		varFirmwareRunningFirmware.ObjectType = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ObjectType
		varFirmwareRunningFirmware.Component = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Component
		varFirmwareRunningFirmware.PackageVersion = varFirmwareRunningFirmwareWithoutEmbeddedStruct.PackageVersion
		varFirmwareRunningFirmware.Type = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Type
		varFirmwareRunningFirmware.Version = varFirmwareRunningFirmwareWithoutEmbeddedStruct.Version
		varFirmwareRunningFirmware.BiosUnit = varFirmwareRunningFirmwareWithoutEmbeddedStruct.BiosUnit
		varFirmwareRunningFirmware.GraphicsCard = varFirmwareRunningFirmwareWithoutEmbeddedStruct.GraphicsCard
		varFirmwareRunningFirmware.InventoryDeviceInfo = varFirmwareRunningFirmwareWithoutEmbeddedStruct.InventoryDeviceInfo
		varFirmwareRunningFirmware.ManagementController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.ManagementController
		varFirmwareRunningFirmware.NetworkElements = varFirmwareRunningFirmwareWithoutEmbeddedStruct.NetworkElements
		varFirmwareRunningFirmware.PciSwitch = varFirmwareRunningFirmwareWithoutEmbeddedStruct.PciSwitch
		varFirmwareRunningFirmware.RegisteredDevice = varFirmwareRunningFirmwareWithoutEmbeddedStruct.RegisteredDevice
		varFirmwareRunningFirmware.StorageController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StorageController
		varFirmwareRunningFirmware.StorageFlexFlashController = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StorageFlexFlashController
		varFirmwareRunningFirmware.StoragePhysicalDisk = varFirmwareRunningFirmwareWithoutEmbeddedStruct.StoragePhysicalDisk
		*o = FirmwareRunningFirmware(varFirmwareRunningFirmware)
	} else {
		return err
	}

	varFirmwareRunningFirmware := _FirmwareRunningFirmware{}

	err = json.Unmarshal(bytes, &varFirmwareRunningFirmware)
	if err == nil {
		o.InventoryBase = varFirmwareRunningFirmware.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Component")
		delete(additionalProperties, "PackageVersion")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "BiosUnit")
		delete(additionalProperties, "GraphicsCard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "ManagementController")
		delete(additionalProperties, "NetworkElements")
		delete(additionalProperties, "PciSwitch")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "StorageFlexFlashController")
		delete(additionalProperties, "StoragePhysicalDisk")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableFirmwareRunningFirmware struct {
	value *FirmwareRunningFirmware
	isSet bool
}

func (v NullableFirmwareRunningFirmware) Get() *FirmwareRunningFirmware {
	return v.value
}

func (v *NullableFirmwareRunningFirmware) Set(val *FirmwareRunningFirmware) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareRunningFirmware) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareRunningFirmware) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareRunningFirmware(val *FirmwareRunningFirmware) *NullableFirmwareRunningFirmware {
	return &NullableFirmwareRunningFirmware{value: val, isSet: true}
}

func (v NullableFirmwareRunningFirmware) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareRunningFirmware) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
