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

// StorageFlexUtilVirtualDrive Storage Flex Util Virtual Drive.
type StorageFlexUtilVirtualDrive struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the Flex Util virtual drive.
	DriveStatus *string `json:"DriveStatus,omitempty"`
	// Type of virtual drive managed by flex util controller.
	DriveType *string `json:"DriveType,omitempty"`
	// Disk Partition Id of virtual drive managed by flex util controller.
	PartitionId *string `json:"PartitionId,omitempty"`
	// Partition name of the Flex Util virtual drive.
	PartitionName *string `json:"PartitionName,omitempty"`
	// The resident image on the flex util virtual Drive.
	ResidentImage *string `json:"ResidentImage,omitempty"`
	// Size of the Flex Util virtual drive.
	Size *string `json:"Size,omitempty"`
	// Virtual drive on the Flex Util controller.
	VirtualDrive              *string                                `json:"VirtualDrive,omitempty"`
	InventoryDeviceInfo       *InventoryDeviceInfoRelationship       `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
	StorageFlexUtilController *StorageFlexUtilControllerRelationship `json:"StorageFlexUtilController,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _StorageFlexUtilVirtualDrive StorageFlexUtilVirtualDrive

// NewStorageFlexUtilVirtualDrive instantiates a new StorageFlexUtilVirtualDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexUtilVirtualDrive(classId string, objectType string) *StorageFlexUtilVirtualDrive {
	this := StorageFlexUtilVirtualDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexUtilVirtualDriveWithDefaults instantiates a new StorageFlexUtilVirtualDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexUtilVirtualDriveWithDefaults() *StorageFlexUtilVirtualDrive {
	this := StorageFlexUtilVirtualDrive{}
	var classId string = "storage.FlexUtilVirtualDrive"
	this.ClassId = classId
	var objectType string = "storage.FlexUtilVirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexUtilVirtualDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexUtilVirtualDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexUtilVirtualDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexUtilVirtualDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveStatus returns the DriveStatus field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetDriveStatus() string {
	if o == nil || o.DriveStatus == nil {
		var ret string
		return ret
	}
	return *o.DriveStatus
}

// GetDriveStatusOk returns a tuple with the DriveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetDriveStatusOk() (*string, bool) {
	if o == nil || o.DriveStatus == nil {
		return nil, false
	}
	return o.DriveStatus, true
}

// HasDriveStatus returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasDriveStatus() bool {
	if o != nil && o.DriveStatus != nil {
		return true
	}

	return false
}

// SetDriveStatus gets a reference to the given string and assigns it to the DriveStatus field.
func (o *StorageFlexUtilVirtualDrive) SetDriveStatus(v string) {
	o.DriveStatus = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *StorageFlexUtilVirtualDrive) SetDriveType(v string) {
	o.DriveType = &v
}

// GetPartitionId returns the PartitionId field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetPartitionId() string {
	if o == nil || o.PartitionId == nil {
		var ret string
		return ret
	}
	return *o.PartitionId
}

// GetPartitionIdOk returns a tuple with the PartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetPartitionIdOk() (*string, bool) {
	if o == nil || o.PartitionId == nil {
		return nil, false
	}
	return o.PartitionId, true
}

// HasPartitionId returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasPartitionId() bool {
	if o != nil && o.PartitionId != nil {
		return true
	}

	return false
}

// SetPartitionId gets a reference to the given string and assigns it to the PartitionId field.
func (o *StorageFlexUtilVirtualDrive) SetPartitionId(v string) {
	o.PartitionId = &v
}

// GetPartitionName returns the PartitionName field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetPartitionName() string {
	if o == nil || o.PartitionName == nil {
		var ret string
		return ret
	}
	return *o.PartitionName
}

// GetPartitionNameOk returns a tuple with the PartitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetPartitionNameOk() (*string, bool) {
	if o == nil || o.PartitionName == nil {
		return nil, false
	}
	return o.PartitionName, true
}

// HasPartitionName returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasPartitionName() bool {
	if o != nil && o.PartitionName != nil {
		return true
	}

	return false
}

// SetPartitionName gets a reference to the given string and assigns it to the PartitionName field.
func (o *StorageFlexUtilVirtualDrive) SetPartitionName(v string) {
	o.PartitionName = &v
}

// GetResidentImage returns the ResidentImage field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetResidentImage() string {
	if o == nil || o.ResidentImage == nil {
		var ret string
		return ret
	}
	return *o.ResidentImage
}

// GetResidentImageOk returns a tuple with the ResidentImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetResidentImageOk() (*string, bool) {
	if o == nil || o.ResidentImage == nil {
		return nil, false
	}
	return o.ResidentImage, true
}

// HasResidentImage returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasResidentImage() bool {
	if o != nil && o.ResidentImage != nil {
		return true
	}

	return false
}

// SetResidentImage gets a reference to the given string and assigns it to the ResidentImage field.
func (o *StorageFlexUtilVirtualDrive) SetResidentImage(v string) {
	o.ResidentImage = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StorageFlexUtilVirtualDrive) SetSize(v string) {
	o.Size = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetVirtualDrive() string {
	if o == nil || o.VirtualDrive == nil {
		var ret string
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetVirtualDriveOk() (*string, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given string and assigns it to the VirtualDrive field.
func (o *StorageFlexUtilVirtualDrive) SetVirtualDrive(v string) {
	o.VirtualDrive = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexUtilVirtualDrive) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexUtilVirtualDrive) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageFlexUtilController returns the StorageFlexUtilController field value if set, zero value otherwise.
func (o *StorageFlexUtilVirtualDrive) GetStorageFlexUtilController() StorageFlexUtilControllerRelationship {
	if o == nil || o.StorageFlexUtilController == nil {
		var ret StorageFlexUtilControllerRelationship
		return ret
	}
	return *o.StorageFlexUtilController
}

// GetStorageFlexUtilControllerOk returns a tuple with the StorageFlexUtilController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilVirtualDrive) GetStorageFlexUtilControllerOk() (*StorageFlexUtilControllerRelationship, bool) {
	if o == nil || o.StorageFlexUtilController == nil {
		return nil, false
	}
	return o.StorageFlexUtilController, true
}

// HasStorageFlexUtilController returns a boolean if a field has been set.
func (o *StorageFlexUtilVirtualDrive) HasStorageFlexUtilController() bool {
	if o != nil && o.StorageFlexUtilController != nil {
		return true
	}

	return false
}

// SetStorageFlexUtilController gets a reference to the given StorageFlexUtilControllerRelationship and assigns it to the StorageFlexUtilController field.
func (o *StorageFlexUtilVirtualDrive) SetStorageFlexUtilController(v StorageFlexUtilControllerRelationship) {
	o.StorageFlexUtilController = &v
}

func (o StorageFlexUtilVirtualDrive) MarshalJSON() ([]byte, error) {
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
	if o.DriveStatus != nil {
		toSerialize["DriveStatus"] = o.DriveStatus
	}
	if o.DriveType != nil {
		toSerialize["DriveType"] = o.DriveType
	}
	if o.PartitionId != nil {
		toSerialize["PartitionId"] = o.PartitionId
	}
	if o.PartitionName != nil {
		toSerialize["PartitionName"] = o.PartitionName
	}
	if o.ResidentImage != nil {
		toSerialize["ResidentImage"] = o.ResidentImage
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageFlexUtilController != nil {
		toSerialize["StorageFlexUtilController"] = o.StorageFlexUtilController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexUtilVirtualDrive) UnmarshalJSON(bytes []byte) (err error) {
	type StorageFlexUtilVirtualDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the Flex Util virtual drive.
		DriveStatus *string `json:"DriveStatus,omitempty"`
		// Type of virtual drive managed by flex util controller.
		DriveType *string `json:"DriveType,omitempty"`
		// Disk Partition Id of virtual drive managed by flex util controller.
		PartitionId *string `json:"PartitionId,omitempty"`
		// Partition name of the Flex Util virtual drive.
		PartitionName *string `json:"PartitionName,omitempty"`
		// The resident image on the flex util virtual Drive.
		ResidentImage *string `json:"ResidentImage,omitempty"`
		// Size of the Flex Util virtual drive.
		Size *string `json:"Size,omitempty"`
		// Virtual drive on the Flex Util controller.
		VirtualDrive              *string                                `json:"VirtualDrive,omitempty"`
		InventoryDeviceInfo       *InventoryDeviceInfoRelationship       `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice          *AssetDeviceRegistrationRelationship   `json:"RegisteredDevice,omitempty"`
		StorageFlexUtilController *StorageFlexUtilControllerRelationship `json:"StorageFlexUtilController,omitempty"`
	}

	varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct := StorageFlexUtilVirtualDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexUtilVirtualDrive := _StorageFlexUtilVirtualDrive{}
		varStorageFlexUtilVirtualDrive.ClassId = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.ClassId
		varStorageFlexUtilVirtualDrive.ObjectType = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.ObjectType
		varStorageFlexUtilVirtualDrive.DriveStatus = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.DriveStatus
		varStorageFlexUtilVirtualDrive.DriveType = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.DriveType
		varStorageFlexUtilVirtualDrive.PartitionId = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.PartitionId
		varStorageFlexUtilVirtualDrive.PartitionName = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.PartitionName
		varStorageFlexUtilVirtualDrive.ResidentImage = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.ResidentImage
		varStorageFlexUtilVirtualDrive.Size = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.Size
		varStorageFlexUtilVirtualDrive.VirtualDrive = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.VirtualDrive
		varStorageFlexUtilVirtualDrive.InventoryDeviceInfo = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexUtilVirtualDrive.RegisteredDevice = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexUtilVirtualDrive.StorageFlexUtilController = varStorageFlexUtilVirtualDriveWithoutEmbeddedStruct.StorageFlexUtilController
		*o = StorageFlexUtilVirtualDrive(varStorageFlexUtilVirtualDrive)
	} else {
		return err
	}

	varStorageFlexUtilVirtualDrive := _StorageFlexUtilVirtualDrive{}

	err = json.Unmarshal(bytes, &varStorageFlexUtilVirtualDrive)
	if err == nil {
		o.InventoryBase = varStorageFlexUtilVirtualDrive.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveStatus")
		delete(additionalProperties, "DriveType")
		delete(additionalProperties, "PartitionId")
		delete(additionalProperties, "PartitionName")
		delete(additionalProperties, "ResidentImage")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "VirtualDrive")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageFlexUtilController")

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

type NullableStorageFlexUtilVirtualDrive struct {
	value *StorageFlexUtilVirtualDrive
	isSet bool
}

func (v NullableStorageFlexUtilVirtualDrive) Get() *StorageFlexUtilVirtualDrive {
	return v.value
}

func (v *NullableStorageFlexUtilVirtualDrive) Set(val *StorageFlexUtilVirtualDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexUtilVirtualDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexUtilVirtualDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexUtilVirtualDrive(val *StorageFlexUtilVirtualDrive) *NullableStorageFlexUtilVirtualDrive {
	return &NullableStorageFlexUtilVirtualDrive{value: val, isSet: true}
}

func (v NullableStorageFlexUtilVirtualDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexUtilVirtualDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
