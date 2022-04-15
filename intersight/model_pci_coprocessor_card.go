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
	"reflect"
	"strings"
)

// PciCoprocessorCard PCIe Compression and Cryptographic CPU Offload Card.
type PciCoprocessorCard struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the coprocessor card.
	CardId *int64 `json:"CardId,omitempty"`
	// The PCI slot name for the coprocessor card.
	PciSlot              *string                              `json:"PciSlot,omitempty"`
	ComputeBoard         *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciCoprocessorCard PciCoprocessorCard

// NewPciCoprocessorCard instantiates a new PciCoprocessorCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciCoprocessorCard(classId string, objectType string) *PciCoprocessorCard {
	this := PciCoprocessorCard{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPciCoprocessorCardWithDefaults instantiates a new PciCoprocessorCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciCoprocessorCardWithDefaults() *PciCoprocessorCard {
	this := PciCoprocessorCard{}
	var classId string = "pci.CoprocessorCard"
	this.ClassId = classId
	var objectType string = "pci.CoprocessorCard"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PciCoprocessorCard) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PciCoprocessorCard) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PciCoprocessorCard) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PciCoprocessorCard) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *PciCoprocessorCard) GetCardId() int64 {
	if o == nil || o.CardId == nil {
		var ret int64
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetCardIdOk() (*int64, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *PciCoprocessorCard) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given int64 and assigns it to the CardId field.
func (o *PciCoprocessorCard) SetCardId(v int64) {
	o.CardId = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *PciCoprocessorCard) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *PciCoprocessorCard) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *PciCoprocessorCard) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *PciCoprocessorCard) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *PciCoprocessorCard) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *PciCoprocessorCard) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *PciCoprocessorCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciCoprocessorCard) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciCoprocessorCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PciCoprocessorCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciCoprocessorCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciCoprocessorCard) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciCoprocessorCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PciCoprocessorCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CardId != nil {
		toSerialize["CardId"] = o.CardId
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PciCoprocessorCard) UnmarshalJSON(bytes []byte) (err error) {
	type PciCoprocessorCardWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The id of the coprocessor card.
		CardId *int64 `json:"CardId,omitempty"`
		// The PCI slot name for the coprocessor card.
		PciSlot             *string                              `json:"PciSlot,omitempty"`
		ComputeBoard        *ComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varPciCoprocessorCardWithoutEmbeddedStruct := PciCoprocessorCardWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPciCoprocessorCardWithoutEmbeddedStruct)
	if err == nil {
		varPciCoprocessorCard := _PciCoprocessorCard{}
		varPciCoprocessorCard.ClassId = varPciCoprocessorCardWithoutEmbeddedStruct.ClassId
		varPciCoprocessorCard.ObjectType = varPciCoprocessorCardWithoutEmbeddedStruct.ObjectType
		varPciCoprocessorCard.CardId = varPciCoprocessorCardWithoutEmbeddedStruct.CardId
		varPciCoprocessorCard.PciSlot = varPciCoprocessorCardWithoutEmbeddedStruct.PciSlot
		varPciCoprocessorCard.ComputeBoard = varPciCoprocessorCardWithoutEmbeddedStruct.ComputeBoard
		varPciCoprocessorCard.InventoryDeviceInfo = varPciCoprocessorCardWithoutEmbeddedStruct.InventoryDeviceInfo
		varPciCoprocessorCard.RegisteredDevice = varPciCoprocessorCardWithoutEmbeddedStruct.RegisteredDevice
		*o = PciCoprocessorCard(varPciCoprocessorCard)
	} else {
		return err
	}

	varPciCoprocessorCard := _PciCoprocessorCard{}

	err = json.Unmarshal(bytes, &varPciCoprocessorCard)
	if err == nil {
		o.EquipmentBase = varPciCoprocessorCard.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CardId")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullablePciCoprocessorCard struct {
	value *PciCoprocessorCard
	isSet bool
}

func (v NullablePciCoprocessorCard) Get() *PciCoprocessorCard {
	return v.value
}

func (v *NullablePciCoprocessorCard) Set(val *PciCoprocessorCard) {
	v.value = val
	v.isSet = true
}

func (v NullablePciCoprocessorCard) IsSet() bool {
	return v.isSet
}

func (v *NullablePciCoprocessorCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciCoprocessorCard(val *PciCoprocessorCard) *NullablePciCoprocessorCard {
	return &NullablePciCoprocessorCard{value: val, isSet: true}
}

func (v NullablePciCoprocessorCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciCoprocessorCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
