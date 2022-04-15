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

// StorageVirtualDriveExtensionAllOf Definition of the list of properties defined in 'storage.VirtualDriveExtension', excluding properties defined in parent classes.
type StorageVirtualDriveExtensionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The ability to boot from the virtual drive.
	Bootable *string `json:"Bootable,omitempty"`
	// The container id of the virtual drive.
	ContainerId *int64 `json:"ContainerId,omitempty"`
	// The state of the virtual drive.
	DriveState *string `json:"DriveState,omitempty"`
	// The name of the Virtual drive.
	Name *string `json:"Name,omitempty"`
	// The operational device id of the virtual drive.
	OperDeviceId *string `json:"OperDeviceId,omitempty"`
	// The UUID assigned to the virtual drive.
	Uuid *string `json:"Uuid,omitempty"`
	// The UUID value of the vendor assigned to the virtual drive.
	VendorUuid *string `json:"VendorUuid,omitempty"`
	// The distinguished name of the virtual drive for which the extended data is provided.
	VirtualDriveDn *string `json:"VirtualDriveDn,omitempty"`
	// The Id of the virtual drive.
	VirtualDriveId       *string                              `json:"VirtualDriveId,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    *StorageControllerRelationship       `json:"StorageController,omitempty"`
	VirtualDrive         *StorageVirtualDriveRelationship     `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveExtensionAllOf StorageVirtualDriveExtensionAllOf

// NewStorageVirtualDriveExtensionAllOf instantiates a new StorageVirtualDriveExtensionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveExtensionAllOf(classId string, objectType string) *StorageVirtualDriveExtensionAllOf {
	this := StorageVirtualDriveExtensionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveExtensionAllOfWithDefaults instantiates a new StorageVirtualDriveExtensionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveExtensionAllOfWithDefaults() *StorageVirtualDriveExtensionAllOf {
	this := StorageVirtualDriveExtensionAllOf{}
	var classId string = "storage.VirtualDriveExtension"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveExtension"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveExtensionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveExtensionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveExtensionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveExtensionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *StorageVirtualDriveExtensionAllOf) SetBootable(v string) {
	o.Bootable = &v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetContainerId() int64 {
	if o == nil || o.ContainerId == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetContainerIdOk() (*int64, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasContainerId() bool {
	if o != nil && o.ContainerId != nil {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *StorageVirtualDriveExtensionAllOf) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetDriveState returns the DriveState field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetDriveState() string {
	if o == nil || o.DriveState == nil {
		var ret string
		return ret
	}
	return *o.DriveState
}

// GetDriveStateOk returns a tuple with the DriveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetDriveStateOk() (*string, bool) {
	if o == nil || o.DriveState == nil {
		return nil, false
	}
	return o.DriveState, true
}

// HasDriveState returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasDriveState() bool {
	if o != nil && o.DriveState != nil {
		return true
	}

	return false
}

// SetDriveState gets a reference to the given string and assigns it to the DriveState field.
func (o *StorageVirtualDriveExtensionAllOf) SetDriveState(v string) {
	o.DriveState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveExtensionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperDeviceId returns the OperDeviceId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetOperDeviceId() string {
	if o == nil || o.OperDeviceId == nil {
		var ret string
		return ret
	}
	return *o.OperDeviceId
}

// GetOperDeviceIdOk returns a tuple with the OperDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetOperDeviceIdOk() (*string, bool) {
	if o == nil || o.OperDeviceId == nil {
		return nil, false
	}
	return o.OperDeviceId, true
}

// HasOperDeviceId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasOperDeviceId() bool {
	if o != nil && o.OperDeviceId != nil {
		return true
	}

	return false
}

// SetOperDeviceId gets a reference to the given string and assigns it to the OperDeviceId field.
func (o *StorageVirtualDriveExtensionAllOf) SetOperDeviceId(v string) {
	o.OperDeviceId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageVirtualDriveExtensionAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetVendorUuid returns the VendorUuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetVendorUuid() string {
	if o == nil || o.VendorUuid == nil {
		var ret string
		return ret
	}
	return *o.VendorUuid
}

// GetVendorUuidOk returns a tuple with the VendorUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetVendorUuidOk() (*string, bool) {
	if o == nil || o.VendorUuid == nil {
		return nil, false
	}
	return o.VendorUuid, true
}

// HasVendorUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasVendorUuid() bool {
	if o != nil && o.VendorUuid != nil {
		return true
	}

	return false
}

// SetVendorUuid gets a reference to the given string and assigns it to the VendorUuid field.
func (o *StorageVirtualDriveExtensionAllOf) SetVendorUuid(v string) {
	o.VendorUuid = &v
}

// GetVirtualDriveDn returns the VirtualDriveDn field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveDn() string {
	if o == nil || o.VirtualDriveDn == nil {
		var ret string
		return ret
	}
	return *o.VirtualDriveDn
}

// GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveDnOk() (*string, bool) {
	if o == nil || o.VirtualDriveDn == nil {
		return nil, false
	}
	return o.VirtualDriveDn, true
}

// HasVirtualDriveDn returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDriveDn() bool {
	if o != nil && o.VirtualDriveDn != nil {
		return true
	}

	return false
}

// SetVirtualDriveDn gets a reference to the given string and assigns it to the VirtualDriveDn field.
func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDriveDn(v string) {
	o.VirtualDriveDn = &v
}

// GetVirtualDriveId returns the VirtualDriveId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveId() string {
	if o == nil || o.VirtualDriveId == nil {
		var ret string
		return ret
	}
	return *o.VirtualDriveId
}

// GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveIdOk() (*string, bool) {
	if o == nil || o.VirtualDriveId == nil {
		return nil, false
	}
	return o.VirtualDriveId, true
}

// HasVirtualDriveId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDriveId() bool {
	if o != nil && o.VirtualDriveId != nil {
		return true
	}

	return false
}

// SetVirtualDriveId gets a reference to the given string and assigns it to the VirtualDriveId field.
func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDriveId(v string) {
	o.VirtualDriveId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVirtualDriveExtensionAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVirtualDriveExtensionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageVirtualDriveExtensionAllOf) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || o.VirtualDrive == nil {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtensionAllOf) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtensionAllOf) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveExtensionAllOf) SetVirtualDrive(v StorageVirtualDriveRelationship) {
	o.VirtualDrive = &v
}

func (o StorageVirtualDriveExtensionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.ContainerId != nil {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if o.DriveState != nil {
		toSerialize["DriveState"] = o.DriveState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperDeviceId != nil {
		toSerialize["OperDeviceId"] = o.OperDeviceId
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VendorUuid != nil {
		toSerialize["VendorUuid"] = o.VendorUuid
	}
	if o.VirtualDriveDn != nil {
		toSerialize["VirtualDriveDn"] = o.VirtualDriveDn
	}
	if o.VirtualDriveId != nil {
		toSerialize["VirtualDriveId"] = o.VirtualDriveId
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveExtensionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageVirtualDriveExtensionAllOf := _StorageVirtualDriveExtensionAllOf{}

	if err = json.Unmarshal(bytes, &varStorageVirtualDriveExtensionAllOf); err == nil {
		*o = StorageVirtualDriveExtensionAllOf(varStorageVirtualDriveExtensionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "DriveState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperDeviceId")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VendorUuid")
		delete(additionalProperties, "VirtualDriveDn")
		delete(additionalProperties, "VirtualDriveId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "VirtualDrive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageVirtualDriveExtensionAllOf struct {
	value *StorageVirtualDriveExtensionAllOf
	isSet bool
}

func (v NullableStorageVirtualDriveExtensionAllOf) Get() *StorageVirtualDriveExtensionAllOf {
	return v.value
}

func (v *NullableStorageVirtualDriveExtensionAllOf) Set(val *StorageVirtualDriveExtensionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveExtensionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveExtensionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveExtensionAllOf(val *StorageVirtualDriveExtensionAllOf) *NullableStorageVirtualDriveExtensionAllOf {
	return &NullableStorageVirtualDriveExtensionAllOf{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveExtensionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveExtensionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
