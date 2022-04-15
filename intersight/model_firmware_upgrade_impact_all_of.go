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

// FirmwareUpgradeImpactAllOf Definition of the list of properties defined in 'firmware.UpgradeImpact', excluding properties defined in parent classes.
type FirmwareUpgradeImpactAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to equipmentChassis resources.
	Chassis []EquipmentChassisRelationship `json:"Chassis,omitempty"`
	// An array of relationships to assetDeviceRegistration resources.
	Device        []AssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Distributable *FirmwareDistributableRelationship    `json:"Distributable,omitempty"`
	// An array of relationships to networkElement resources.
	NetworkElements []NetworkElementRelationship           `json:"NetworkElements,omitempty"`
	Release         *SoftwarerepositoryReleaseRelationship `json:"Release,omitempty"`
	// An array of relationships to computePhysical resources.
	Server               []ComputePhysicalRelationship `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeImpactAllOf FirmwareUpgradeImpactAllOf

// NewFirmwareUpgradeImpactAllOf instantiates a new FirmwareUpgradeImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeImpactAllOf(classId string, objectType string) *FirmwareUpgradeImpactAllOf {
	this := FirmwareUpgradeImpactAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareUpgradeImpactAllOfWithDefaults instantiates a new FirmwareUpgradeImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeImpactAllOfWithDefaults() *FirmwareUpgradeImpactAllOf {
	this := FirmwareUpgradeImpactAllOf{}
	var classId string = "firmware.UpgradeImpact"
	this.ClassId = classId
	var objectType string = "firmware.UpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeImpactAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeImpactAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeImpactAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeImpactAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassis returns the Chassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpactAllOf) GetChassis() []EquipmentChassisRelationship {
	if o == nil {
		var ret []EquipmentChassisRelationship
		return ret
	}
	return o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpactAllOf) GetChassisOk() (*[]EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return &o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given []EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *FirmwareUpgradeImpactAllOf) SetChassis(v []EquipmentChassisRelationship) {
	o.Chassis = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpactAllOf) GetDevice() []AssetDeviceRegistrationRelationship {
	if o == nil {
		var ret []AssetDeviceRegistrationRelationship
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpactAllOf) GetDeviceOk() (*[]AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return &o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given []AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareUpgradeImpactAllOf) SetDevice(v []AssetDeviceRegistrationRelationship) {
	o.Device = v
}

// GetDistributable returns the Distributable field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactAllOf) GetDistributable() FirmwareDistributableRelationship {
	if o == nil || o.Distributable == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.Distributable
}

// GetDistributableOk returns a tuple with the Distributable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactAllOf) GetDistributableOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.Distributable == nil {
		return nil, false
	}
	return o.Distributable, true
}

// HasDistributable returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasDistributable() bool {
	if o != nil && o.Distributable != nil {
		return true
	}

	return false
}

// SetDistributable gets a reference to the given FirmwareDistributableRelationship and assigns it to the Distributable field.
func (o *FirmwareUpgradeImpactAllOf) SetDistributable(v FirmwareDistributableRelationship) {
	o.Distributable = &v
}

// GetNetworkElements returns the NetworkElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpactAllOf) GetNetworkElements() []NetworkElementRelationship {
	if o == nil {
		var ret []NetworkElementRelationship
		return ret
	}
	return o.NetworkElements
}

// GetNetworkElementsOk returns a tuple with the NetworkElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpactAllOf) GetNetworkElementsOk() (*[]NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElements == nil {
		return nil, false
	}
	return &o.NetworkElements, true
}

// HasNetworkElements returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasNetworkElements() bool {
	if o != nil && o.NetworkElements != nil {
		return true
	}

	return false
}

// SetNetworkElements gets a reference to the given []NetworkElementRelationship and assigns it to the NetworkElements field.
func (o *FirmwareUpgradeImpactAllOf) SetNetworkElements(v []NetworkElementRelationship) {
	o.NetworkElements = v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *FirmwareUpgradeImpactAllOf) GetRelease() SoftwarerepositoryReleaseRelationship {
	if o == nil || o.Release == nil {
		var ret SoftwarerepositoryReleaseRelationship
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeImpactAllOf) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given SoftwarerepositoryReleaseRelationship and assigns it to the Release field.
func (o *FirmwareUpgradeImpactAllOf) SetRelease(v SoftwarerepositoryReleaseRelationship) {
	o.Release = &v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeImpactAllOf) GetServer() []ComputePhysicalRelationship {
	if o == nil {
		var ret []ComputePhysicalRelationship
		return ret
	}
	return o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeImpactAllOf) GetServerOk() (*[]ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return &o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *FirmwareUpgradeImpactAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given []ComputePhysicalRelationship and assigns it to the Server field.
func (o *FirmwareUpgradeImpactAllOf) SetServer(v []ComputePhysicalRelationship) {
	o.Server = v
}

func (o FirmwareUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Distributable != nil {
		toSerialize["Distributable"] = o.Distributable
	}
	if o.NetworkElements != nil {
		toSerialize["NetworkElements"] = o.NetworkElements
	}
	if o.Release != nil {
		toSerialize["Release"] = o.Release
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeImpactAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareUpgradeImpactAllOf := _FirmwareUpgradeImpactAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareUpgradeImpactAllOf); err == nil {
		*o = FirmwareUpgradeImpactAllOf(varFirmwareUpgradeImpactAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Distributable")
		delete(additionalProperties, "NetworkElements")
		delete(additionalProperties, "Release")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareUpgradeImpactAllOf struct {
	value *FirmwareUpgradeImpactAllOf
	isSet bool
}

func (v NullableFirmwareUpgradeImpactAllOf) Get() *FirmwareUpgradeImpactAllOf {
	return v.value
}

func (v *NullableFirmwareUpgradeImpactAllOf) Set(val *FirmwareUpgradeImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeImpactAllOf(val *FirmwareUpgradeImpactAllOf) *NullableFirmwareUpgradeImpactAllOf {
	return &NullableFirmwareUpgradeImpactAllOf{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
