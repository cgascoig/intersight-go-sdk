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

// EquipmentLocatorLed Locator Led of an Equipment like Rack, Disk etc.
type EquipmentLocatorLed struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Color of the locatorled available on an equipment.
	Color *string `json:"Color,omitempty"`
	// Identifies the operational state of locatorled.
	OperState            *string                              `json:"OperState,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	EquipmentFex         *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StoragePhysicalDisk  *StoragePhysicalDiskRelationship     `json:"StoragePhysicalDisk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentLocatorLed EquipmentLocatorLed

// NewEquipmentLocatorLed instantiates a new EquipmentLocatorLed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentLocatorLed(classId string, objectType string) *EquipmentLocatorLed {
	this := EquipmentLocatorLed{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentLocatorLedWithDefaults instantiates a new EquipmentLocatorLed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentLocatorLedWithDefaults() *EquipmentLocatorLed {
	this := EquipmentLocatorLed{}
	var classId string = "equipment.LocatorLed"
	this.ClassId = classId
	var objectType string = "equipment.LocatorLed"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentLocatorLed) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentLocatorLed) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentLocatorLed) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentLocatorLed) SetObjectType(v string) {
	o.ObjectType = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *EquipmentLocatorLed) SetColor(v string) {
	o.Color = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentLocatorLed) SetOperState(v string) {
	o.OperState = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *EquipmentLocatorLed) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentLocatorLed) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentLocatorLed) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentLocatorLed) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentLocatorLed) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentLocatorLed) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStoragePhysicalDisk returns the StoragePhysicalDisk field value if set, zero value otherwise.
func (o *EquipmentLocatorLed) GetStoragePhysicalDisk() StoragePhysicalDiskRelationship {
	if o == nil || o.StoragePhysicalDisk == nil {
		var ret StoragePhysicalDiskRelationship
		return ret
	}
	return *o.StoragePhysicalDisk
}

// GetStoragePhysicalDiskOk returns a tuple with the StoragePhysicalDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentLocatorLed) GetStoragePhysicalDiskOk() (*StoragePhysicalDiskRelationship, bool) {
	if o == nil || o.StoragePhysicalDisk == nil {
		return nil, false
	}
	return o.StoragePhysicalDisk, true
}

// HasStoragePhysicalDisk returns a boolean if a field has been set.
func (o *EquipmentLocatorLed) HasStoragePhysicalDisk() bool {
	if o != nil && o.StoragePhysicalDisk != nil {
		return true
	}

	return false
}

// SetStoragePhysicalDisk gets a reference to the given StoragePhysicalDiskRelationship and assigns it to the StoragePhysicalDisk field.
func (o *EquipmentLocatorLed) SetStoragePhysicalDisk(v StoragePhysicalDiskRelationship) {
	o.StoragePhysicalDisk = &v
}

func (o EquipmentLocatorLed) MarshalJSON() ([]byte, error) {
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
	if o.Color != nil {
		toSerialize["Color"] = o.Color
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StoragePhysicalDisk != nil {
		toSerialize["StoragePhysicalDisk"] = o.StoragePhysicalDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentLocatorLed) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentLocatorLedWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Color of the locatorled available on an equipment.
		Color *string `json:"Color,omitempty"`
		// Identifies the operational state of locatorled.
		OperState           *string                              `json:"OperState,omitempty"`
		ComputeBlade        *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		EquipmentChassis    *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		EquipmentFex        *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StoragePhysicalDisk *StoragePhysicalDiskRelationship     `json:"StoragePhysicalDisk,omitempty"`
	}

	varEquipmentLocatorLedWithoutEmbeddedStruct := EquipmentLocatorLedWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentLocatorLedWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentLocatorLed := _EquipmentLocatorLed{}
		varEquipmentLocatorLed.ClassId = varEquipmentLocatorLedWithoutEmbeddedStruct.ClassId
		varEquipmentLocatorLed.ObjectType = varEquipmentLocatorLedWithoutEmbeddedStruct.ObjectType
		varEquipmentLocatorLed.Color = varEquipmentLocatorLedWithoutEmbeddedStruct.Color
		varEquipmentLocatorLed.OperState = varEquipmentLocatorLedWithoutEmbeddedStruct.OperState
		varEquipmentLocatorLed.ComputeBlade = varEquipmentLocatorLedWithoutEmbeddedStruct.ComputeBlade
		varEquipmentLocatorLed.ComputeRackUnit = varEquipmentLocatorLedWithoutEmbeddedStruct.ComputeRackUnit
		varEquipmentLocatorLed.EquipmentChassis = varEquipmentLocatorLedWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentLocatorLed.EquipmentFex = varEquipmentLocatorLedWithoutEmbeddedStruct.EquipmentFex
		varEquipmentLocatorLed.InventoryDeviceInfo = varEquipmentLocatorLedWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentLocatorLed.RegisteredDevice = varEquipmentLocatorLedWithoutEmbeddedStruct.RegisteredDevice
		varEquipmentLocatorLed.StoragePhysicalDisk = varEquipmentLocatorLedWithoutEmbeddedStruct.StoragePhysicalDisk
		*o = EquipmentLocatorLed(varEquipmentLocatorLed)
	} else {
		return err
	}

	varEquipmentLocatorLed := _EquipmentLocatorLed{}

	err = json.Unmarshal(bytes, &varEquipmentLocatorLed)
	if err == nil {
		o.InventoryBase = varEquipmentLocatorLed.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Color")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
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

type NullableEquipmentLocatorLed struct {
	value *EquipmentLocatorLed
	isSet bool
}

func (v NullableEquipmentLocatorLed) Get() *EquipmentLocatorLed {
	return v.value
}

func (v *NullableEquipmentLocatorLed) Set(val *EquipmentLocatorLed) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentLocatorLed) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentLocatorLed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentLocatorLed(val *EquipmentLocatorLed) *NullableEquipmentLocatorLed {
	return &NullableEquipmentLocatorLed{value: val, isSet: true}
}

func (v NullableEquipmentLocatorLed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentLocatorLed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
