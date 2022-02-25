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
	"reflect"
	"strings"
)

// FirmwareUnsupportedVersionUpgrade This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
type FirmwareUnsupportedVersionUpgrade struct {
	ConnectorDownloadStatus
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Workflow status of firmware upgrade. * `None` - Upgrade status is none when upgrade is in progress. * `Completed` - Upgrade completed successfully. * `Failed` - Upgrade status is failed when upgrade has failed.
	UpgradeStatus        *string                                `json:"UpgradeStatus,omitempty"`
	Device               *AssetDeviceRegistrationRelationship   `json:"Device,omitempty"`
	Distributable        *FirmwareDistributableRelationship     `json:"Distributable,omitempty"`
	PhysicalIdentity     *EquipmentPhysicalIdentityRelationship `json:"PhysicalIdentity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUnsupportedVersionUpgrade FirmwareUnsupportedVersionUpgrade

// NewFirmwareUnsupportedVersionUpgrade instantiates a new FirmwareUnsupportedVersionUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUnsupportedVersionUpgrade(classId string, objectType string) *FirmwareUnsupportedVersionUpgrade {
	this := FirmwareUnsupportedVersionUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var upgradeStatus string = "None"
	this.UpgradeStatus = &upgradeStatus
	return &this
}

// NewFirmwareUnsupportedVersionUpgradeWithDefaults instantiates a new FirmwareUnsupportedVersionUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUnsupportedVersionUpgradeWithDefaults() *FirmwareUnsupportedVersionUpgrade {
	this := FirmwareUnsupportedVersionUpgrade{}
	var classId string = "firmware.UnsupportedVersionUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.UnsupportedVersionUpgrade"
	this.ObjectType = objectType
	var upgradeStatus string = "None"
	this.UpgradeStatus = &upgradeStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUnsupportedVersionUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUnsupportedVersionUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUnsupportedVersionUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUnsupportedVersionUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUpgradeStatus returns the UpgradeStatus field value if set, zero value otherwise.
func (o *FirmwareUnsupportedVersionUpgrade) GetUpgradeStatus() string {
	if o == nil || o.UpgradeStatus == nil {
		var ret string
		return ret
	}
	return *o.UpgradeStatus
}

// GetUpgradeStatusOk returns a tuple with the UpgradeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetUpgradeStatusOk() (*string, bool) {
	if o == nil || o.UpgradeStatus == nil {
		return nil, false
	}
	return o.UpgradeStatus, true
}

// HasUpgradeStatus returns a boolean if a field has been set.
func (o *FirmwareUnsupportedVersionUpgrade) HasUpgradeStatus() bool {
	if o != nil && o.UpgradeStatus != nil {
		return true
	}

	return false
}

// SetUpgradeStatus gets a reference to the given string and assigns it to the UpgradeStatus field.
func (o *FirmwareUnsupportedVersionUpgrade) SetUpgradeStatus(v string) {
	o.UpgradeStatus = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *FirmwareUnsupportedVersionUpgrade) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareUnsupportedVersionUpgrade) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareUnsupportedVersionUpgrade) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetDistributable returns the Distributable field value if set, zero value otherwise.
func (o *FirmwareUnsupportedVersionUpgrade) GetDistributable() FirmwareDistributableRelationship {
	if o == nil || o.Distributable == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.Distributable
}

// GetDistributableOk returns a tuple with the Distributable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetDistributableOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.Distributable == nil {
		return nil, false
	}
	return o.Distributable, true
}

// HasDistributable returns a boolean if a field has been set.
func (o *FirmwareUnsupportedVersionUpgrade) HasDistributable() bool {
	if o != nil && o.Distributable != nil {
		return true
	}

	return false
}

// SetDistributable gets a reference to the given FirmwareDistributableRelationship and assigns it to the Distributable field.
func (o *FirmwareUnsupportedVersionUpgrade) SetDistributable(v FirmwareDistributableRelationship) {
	o.Distributable = &v
}

// GetPhysicalIdentity returns the PhysicalIdentity field value if set, zero value otherwise.
func (o *FirmwareUnsupportedVersionUpgrade) GetPhysicalIdentity() EquipmentPhysicalIdentityRelationship {
	if o == nil || o.PhysicalIdentity == nil {
		var ret EquipmentPhysicalIdentityRelationship
		return ret
	}
	return *o.PhysicalIdentity
}

// GetPhysicalIdentityOk returns a tuple with the PhysicalIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUnsupportedVersionUpgrade) GetPhysicalIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool) {
	if o == nil || o.PhysicalIdentity == nil {
		return nil, false
	}
	return o.PhysicalIdentity, true
}

// HasPhysicalIdentity returns a boolean if a field has been set.
func (o *FirmwareUnsupportedVersionUpgrade) HasPhysicalIdentity() bool {
	if o != nil && o.PhysicalIdentity != nil {
		return true
	}

	return false
}

// SetPhysicalIdentity gets a reference to the given EquipmentPhysicalIdentityRelationship and assigns it to the PhysicalIdentity field.
func (o *FirmwareUnsupportedVersionUpgrade) SetPhysicalIdentity(v EquipmentPhysicalIdentityRelationship) {
	o.PhysicalIdentity = &v
}

func (o FirmwareUnsupportedVersionUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorDownloadStatus, errConnectorDownloadStatus := json.Marshal(o.ConnectorDownloadStatus)
	if errConnectorDownloadStatus != nil {
		return []byte{}, errConnectorDownloadStatus
	}
	errConnectorDownloadStatus = json.Unmarshal([]byte(serializedConnectorDownloadStatus), &toSerialize)
	if errConnectorDownloadStatus != nil {
		return []byte{}, errConnectorDownloadStatus
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.UpgradeStatus != nil {
		toSerialize["UpgradeStatus"] = o.UpgradeStatus
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Distributable != nil {
		toSerialize["Distributable"] = o.Distributable
	}
	if o.PhysicalIdentity != nil {
		toSerialize["PhysicalIdentity"] = o.PhysicalIdentity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUnsupportedVersionUpgrade) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Workflow status of firmware upgrade. * `None` - Upgrade status is none when upgrade is in progress. * `Completed` - Upgrade completed successfully. * `Failed` - Upgrade status is failed when upgrade has failed.
		UpgradeStatus    *string                                `json:"UpgradeStatus,omitempty"`
		Device           *AssetDeviceRegistrationRelationship   `json:"Device,omitempty"`
		Distributable    *FirmwareDistributableRelationship     `json:"Distributable,omitempty"`
		PhysicalIdentity *EquipmentPhysicalIdentityRelationship `json:"PhysicalIdentity,omitempty"`
	}

	varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct := FirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUnsupportedVersionUpgrade := _FirmwareUnsupportedVersionUpgrade{}
		varFirmwareUnsupportedVersionUpgrade.ClassId = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.ClassId
		varFirmwareUnsupportedVersionUpgrade.ObjectType = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.ObjectType
		varFirmwareUnsupportedVersionUpgrade.UpgradeStatus = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.UpgradeStatus
		varFirmwareUnsupportedVersionUpgrade.Device = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.Device
		varFirmwareUnsupportedVersionUpgrade.Distributable = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.Distributable
		varFirmwareUnsupportedVersionUpgrade.PhysicalIdentity = varFirmwareUnsupportedVersionUpgradeWithoutEmbeddedStruct.PhysicalIdentity
		*o = FirmwareUnsupportedVersionUpgrade(varFirmwareUnsupportedVersionUpgrade)
	} else {
		return err
	}

	varFirmwareUnsupportedVersionUpgrade := _FirmwareUnsupportedVersionUpgrade{}

	err = json.Unmarshal(bytes, &varFirmwareUnsupportedVersionUpgrade)
	if err == nil {
		o.ConnectorDownloadStatus = varFirmwareUnsupportedVersionUpgrade.ConnectorDownloadStatus
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UpgradeStatus")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Distributable")
		delete(additionalProperties, "PhysicalIdentity")

		// remove fields from embedded structs
		reflectConnectorDownloadStatus := reflect.ValueOf(o.ConnectorDownloadStatus)
		for i := 0; i < reflectConnectorDownloadStatus.Type().NumField(); i++ {
			t := reflectConnectorDownloadStatus.Type().Field(i)

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

type NullableFirmwareUnsupportedVersionUpgrade struct {
	value *FirmwareUnsupportedVersionUpgrade
	isSet bool
}

func (v NullableFirmwareUnsupportedVersionUpgrade) Get() *FirmwareUnsupportedVersionUpgrade {
	return v.value
}

func (v *NullableFirmwareUnsupportedVersionUpgrade) Set(val *FirmwareUnsupportedVersionUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUnsupportedVersionUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUnsupportedVersionUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUnsupportedVersionUpgrade(val *FirmwareUnsupportedVersionUpgrade) *NullableFirmwareUnsupportedVersionUpgrade {
	return &NullableFirmwareUnsupportedVersionUpgrade{value: val, isSet: true}
}

func (v NullableFirmwareUnsupportedVersionUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUnsupportedVersionUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
