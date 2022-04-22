/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentIoCardAllOf Definition of the list of properties defined in 'equipment.IoCard', excluding properties defined in parent classes.
type EquipmentIoCardAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Switch Id to which the IOM is connected to. The value can be A or B.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// IOM device connector support.
	DcSupported       *bool              `json:"DcSupported,omitempty"`
	InbandIpAddresses []ComputeIpAddress `json:"InbandIpAddresses,omitempty"`
	// Location of IOM within a chassis. The value can be left or right.
	Side             *string                       `json:"Side,omitempty"`
	EquipmentChassis *EquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
	EquipmentFex     *EquipmentFexRelationship     `json:"EquipmentFex,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	FanModules                 []EquipmentFanModuleRelationship     `json:"FanModules,omitempty"`
	InventoryDeviceInfo        *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	PhysicalDeviceRegistration *AssetDeviceRegistrationRelationship `json:"PhysicalDeviceRegistration,omitempty"`
	RegisteredDevice           *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _EquipmentIoCardAllOf EquipmentIoCardAllOf

// NewEquipmentIoCardAllOf instantiates a new EquipmentIoCardAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardAllOf(classId string, objectType string) *EquipmentIoCardAllOf {
	this := EquipmentIoCardAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardAllOfWithDefaults instantiates a new EquipmentIoCardAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardAllOfWithDefaults() *EquipmentIoCardAllOf {
	this := EquipmentIoCardAllOf{}
	var classId string = "equipment.IoCard"
	this.ClassId = classId
	var objectType string = "equipment.IoCard"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetConnectionPath() string {
	if o == nil || o.ConnectionPath == nil {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetConnectionPathOk() (*string, bool) {
	if o == nil || o.ConnectionPath == nil {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasConnectionPath() bool {
	if o != nil && o.ConnectionPath != nil {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentIoCardAllOf) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetDcSupported returns the DcSupported field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetDcSupported() bool {
	if o == nil || o.DcSupported == nil {
		var ret bool
		return ret
	}
	return *o.DcSupported
}

// GetDcSupportedOk returns a tuple with the DcSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetDcSupportedOk() (*bool, bool) {
	if o == nil || o.DcSupported == nil {
		return nil, false
	}
	return o.DcSupported, true
}

// HasDcSupported returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasDcSupported() bool {
	if o != nil && o.DcSupported != nil {
		return true
	}

	return false
}

// SetDcSupported gets a reference to the given bool and assigns it to the DcSupported field.
func (o *EquipmentIoCardAllOf) SetDcSupported(v bool) {
	o.DcSupported = &v
}

// GetInbandIpAddresses returns the InbandIpAddresses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardAllOf) GetInbandIpAddresses() []ComputeIpAddress {
	if o == nil {
		var ret []ComputeIpAddress
		return ret
	}
	return o.InbandIpAddresses
}

// GetInbandIpAddressesOk returns a tuple with the InbandIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardAllOf) GetInbandIpAddressesOk() (*[]ComputeIpAddress, bool) {
	if o == nil || o.InbandIpAddresses == nil {
		return nil, false
	}
	return &o.InbandIpAddresses, true
}

// HasInbandIpAddresses returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasInbandIpAddresses() bool {
	if o != nil && o.InbandIpAddresses != nil {
		return true
	}

	return false
}

// SetInbandIpAddresses gets a reference to the given []ComputeIpAddress and assigns it to the InbandIpAddresses field.
func (o *EquipmentIoCardAllOf) SetInbandIpAddresses(v []ComputeIpAddress) {
	o.InbandIpAddresses = v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetSide() string {
	if o == nil || o.Side == nil {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetSideOk() (*string, bool) {
	if o == nil || o.Side == nil {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasSide() bool {
	if o != nil && o.Side != nil {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *EquipmentIoCardAllOf) SetSide(v string) {
	o.Side = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentIoCardAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentIoCardAllOf) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetFanModules returns the FanModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardAllOf) GetFanModules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.FanModules
}

// GetFanModulesOk returns a tuple with the FanModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardAllOf) GetFanModulesOk() (*[]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.FanModules == nil {
		return nil, false
	}
	return &o.FanModules, true
}

// HasFanModules returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasFanModules() bool {
	if o != nil && o.FanModules != nil {
		return true
	}

	return false
}

// SetFanModules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the FanModules field.
func (o *EquipmentIoCardAllOf) SetFanModules(v []EquipmentFanModuleRelationship) {
	o.FanModules = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentIoCardAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPhysicalDeviceRegistration returns the PhysicalDeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetPhysicalDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.PhysicalDeviceRegistration
}

// GetPhysicalDeviceRegistrationOk returns a tuple with the PhysicalDeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetPhysicalDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.PhysicalDeviceRegistration == nil {
		return nil, false
	}
	return o.PhysicalDeviceRegistration, true
}

// HasPhysicalDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasPhysicalDeviceRegistration() bool {
	if o != nil && o.PhysicalDeviceRegistration != nil {
		return true
	}

	return false
}

// SetPhysicalDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the PhysicalDeviceRegistration field.
func (o *EquipmentIoCardAllOf) SetPhysicalDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.PhysicalDeviceRegistration = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIoCardAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIoCardAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIoCardAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIoCardAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionPath != nil {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if o.DcSupported != nil {
		toSerialize["DcSupported"] = o.DcSupported
	}
	if o.InbandIpAddresses != nil {
		toSerialize["InbandIpAddresses"] = o.InbandIpAddresses
	}
	if o.Side != nil {
		toSerialize["Side"] = o.Side
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.FanModules != nil {
		toSerialize["FanModules"] = o.FanModules
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PhysicalDeviceRegistration != nil {
		toSerialize["PhysicalDeviceRegistration"] = o.PhysicalDeviceRegistration
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCardAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoCardAllOf := _EquipmentIoCardAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoCardAllOf); err == nil {
		*o = EquipmentIoCardAllOf(varEquipmentIoCardAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "DcSupported")
		delete(additionalProperties, "InbandIpAddresses")
		delete(additionalProperties, "Side")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "FanModules")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PhysicalDeviceRegistration")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoCardAllOf struct {
	value *EquipmentIoCardAllOf
	isSet bool
}

func (v NullableEquipmentIoCardAllOf) Get() *EquipmentIoCardAllOf {
	return v.value
}

func (v *NullableEquipmentIoCardAllOf) Set(val *EquipmentIoCardAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardAllOf(val *EquipmentIoCardAllOf) *NullableEquipmentIoCardAllOf {
	return &NullableEquipmentIoCardAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoCardAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
