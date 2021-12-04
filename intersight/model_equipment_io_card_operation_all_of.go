/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentIoCardOperationAllOf Definition of the list of properties defined in 'equipment.IoCardOperation', excluding properties defined in parent classes.
type EquipmentIoCardOperationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured power state of the iomodule. * `None` - Placeholder default value for iom power state property. * `Reboot` - IO Module reboot state property value.
	AdminPowerState *string `json:"AdminPowerState,omitempty"`
	// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis iomodule. Applying - This state denotes that the settings are being applied in the target chassis iomodule. Failed - This state denotes that the settings could not be applied in the target chassis iomodule. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState          *string                              `json:"ConfigState,omitempty"`
	DeviceRegistration   *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	IoCard               *EquipmentIoCardRelationship         `json:"IoCard,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardOperationAllOf EquipmentIoCardOperationAllOf

// NewEquipmentIoCardOperationAllOf instantiates a new EquipmentIoCardOperationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardOperationAllOf(classId string, objectType string) *EquipmentIoCardOperationAllOf {
	this := EquipmentIoCardOperationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminPowerState string = "None"
	this.AdminPowerState = &adminPowerState
	return &this
}

// NewEquipmentIoCardOperationAllOfWithDefaults instantiates a new EquipmentIoCardOperationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardOperationAllOfWithDefaults() *EquipmentIoCardOperationAllOf {
	this := EquipmentIoCardOperationAllOf{}
	var classId string = "equipment.IoCardOperation"
	this.ClassId = classId
	var objectType string = "equipment.IoCardOperation"
	this.ObjectType = objectType
	var adminPowerState string = "None"
	this.AdminPowerState = &adminPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardOperationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardOperationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardOperationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardOperationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationAllOf) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationAllOf) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *EquipmentIoCardOperationAllOf) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentIoCardOperationAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationAllOf) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentIoCardOperationAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetIoCard returns the IoCard field value if set, zero value otherwise.
func (o *EquipmentIoCardOperationAllOf) GetIoCard() EquipmentIoCardRelationship {
	if o == nil || o.IoCard == nil {
		var ret EquipmentIoCardRelationship
		return ret
	}
	return *o.IoCard
}

// GetIoCardOk returns a tuple with the IoCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperationAllOf) GetIoCardOk() (*EquipmentIoCardRelationship, bool) {
	if o == nil || o.IoCard == nil {
		return nil, false
	}
	return o.IoCard, true
}

// HasIoCard returns a boolean if a field has been set.
func (o *EquipmentIoCardOperationAllOf) HasIoCard() bool {
	if o != nil && o.IoCard != nil {
		return true
	}

	return false
}

// SetIoCard gets a reference to the given EquipmentIoCardRelationship and assigns it to the IoCard field.
func (o *EquipmentIoCardOperationAllOf) SetIoCard(v EquipmentIoCardRelationship) {
	o.IoCard = &v
}

func (o EquipmentIoCardOperationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminPowerState != nil {
		toSerialize["AdminPowerState"] = o.AdminPowerState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.IoCard != nil {
		toSerialize["IoCard"] = o.IoCard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCardOperationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoCardOperationAllOf := _EquipmentIoCardOperationAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoCardOperationAllOf); err == nil {
		*o = EquipmentIoCardOperationAllOf(varEquipmentIoCardOperationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminPowerState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "IoCard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoCardOperationAllOf struct {
	value *EquipmentIoCardOperationAllOf
	isSet bool
}

func (v NullableEquipmentIoCardOperationAllOf) Get() *EquipmentIoCardOperationAllOf {
	return v.value
}

func (v *NullableEquipmentIoCardOperationAllOf) Set(val *EquipmentIoCardOperationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardOperationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardOperationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardOperationAllOf(val *EquipmentIoCardOperationAllOf) *NullableEquipmentIoCardOperationAllOf {
	return &NullableEquipmentIoCardOperationAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoCardOperationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardOperationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
