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
)

// BiosTokenSettingsAllOf Definition of the list of properties defined in 'bios.TokenSettings', excluding properties defined in parent classes.
type BiosTokenSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Value that represents if the BIOS configuration is active. Possible values are \"yes\" and \"no\".
	IsAssigned *string `json:"IsAssigned,omitempty"`
	// Parent server serial number.
	Serial *string `json:"Serial,omitempty"`
	// BIOS configuration name as found in dn. Possible values are \"ADDDC-Sparing\", \"Maximum-Performance\", \"Partial-Mirror-mode-1LM\", \"Mirror-Mode-1LM\", \"Mirroring\", \"Lockstep\", \"Sparing\", \"Platform-Default\".
	SettingsMoRn         *string                              `json:"SettingsMoRn,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosTokenSettingsAllOf BiosTokenSettingsAllOf

// NewBiosTokenSettingsAllOf instantiates a new BiosTokenSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosTokenSettingsAllOf(classId string, objectType string) *BiosTokenSettingsAllOf {
	this := BiosTokenSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosTokenSettingsAllOfWithDefaults instantiates a new BiosTokenSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosTokenSettingsAllOfWithDefaults() *BiosTokenSettingsAllOf {
	this := BiosTokenSettingsAllOf{}
	var classId string = "bios.TokenSettings"
	this.ClassId = classId
	var objectType string = "bios.TokenSettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosTokenSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosTokenSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BiosTokenSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosTokenSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsAssigned returns the IsAssigned field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetIsAssigned() string {
	if o == nil || o.IsAssigned == nil {
		var ret string
		return ret
	}
	return *o.IsAssigned
}

// GetIsAssignedOk returns a tuple with the IsAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetIsAssignedOk() (*string, bool) {
	if o == nil || o.IsAssigned == nil {
		return nil, false
	}
	return o.IsAssigned, true
}

// HasIsAssigned returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasIsAssigned() bool {
	if o != nil && o.IsAssigned != nil {
		return true
	}

	return false
}

// SetIsAssigned gets a reference to the given string and assigns it to the IsAssigned field.
func (o *BiosTokenSettingsAllOf) SetIsAssigned(v string) {
	o.IsAssigned = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *BiosTokenSettingsAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSettingsMoRn returns the SettingsMoRn field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetSettingsMoRn() string {
	if o == nil || o.SettingsMoRn == nil {
		var ret string
		return ret
	}
	return *o.SettingsMoRn
}

// GetSettingsMoRnOk returns a tuple with the SettingsMoRn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetSettingsMoRnOk() (*string, bool) {
	if o == nil || o.SettingsMoRn == nil {
		return nil, false
	}
	return o.SettingsMoRn, true
}

// HasSettingsMoRn returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasSettingsMoRn() bool {
	if o != nil && o.SettingsMoRn != nil {
		return true
	}

	return false
}

// SetSettingsMoRn gets a reference to the given string and assigns it to the SettingsMoRn field.
func (o *BiosTokenSettingsAllOf) SetSettingsMoRn(v string) {
	o.SettingsMoRn = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosTokenSettingsAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosTokenSettingsAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosTokenSettingsAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BiosTokenSettingsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosTokenSettingsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosTokenSettingsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosTokenSettingsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BiosTokenSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsAssigned != nil {
		toSerialize["IsAssigned"] = o.IsAssigned
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SettingsMoRn != nil {
		toSerialize["SettingsMoRn"] = o.SettingsMoRn
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
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

func (o *BiosTokenSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBiosTokenSettingsAllOf := _BiosTokenSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varBiosTokenSettingsAllOf); err == nil {
		*o = BiosTokenSettingsAllOf(varBiosTokenSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsAssigned")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SettingsMoRn")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBiosTokenSettingsAllOf struct {
	value *BiosTokenSettingsAllOf
	isSet bool
}

func (v NullableBiosTokenSettingsAllOf) Get() *BiosTokenSettingsAllOf {
	return v.value
}

func (v *NullableBiosTokenSettingsAllOf) Set(val *BiosTokenSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosTokenSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosTokenSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosTokenSettingsAllOf(val *BiosTokenSettingsAllOf) *NullableBiosTokenSettingsAllOf {
	return &NullableBiosTokenSettingsAllOf{value: val, isSet: true}
}

func (v NullableBiosTokenSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosTokenSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
