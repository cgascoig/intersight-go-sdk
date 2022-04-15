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

// EquipmentIoCardOperation Models the configurable properties of a iomodule in Intersight.
type EquipmentIoCardOperation struct {
	MoBaseMo
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

type _EquipmentIoCardOperation EquipmentIoCardOperation

// NewEquipmentIoCardOperation instantiates a new EquipmentIoCardOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardOperation(classId string, objectType string) *EquipmentIoCardOperation {
	this := EquipmentIoCardOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminPowerState string = "None"
	this.AdminPowerState = &adminPowerState
	return &this
}

// NewEquipmentIoCardOperationWithDefaults instantiates a new EquipmentIoCardOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardOperationWithDefaults() *EquipmentIoCardOperation {
	this := EquipmentIoCardOperation{}
	var classId string = "equipment.IoCardOperation"
	this.ClassId = classId
	var objectType string = "equipment.IoCardOperation"
	this.ObjectType = objectType
	var adminPowerState string = "None"
	this.AdminPowerState = &adminPowerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *EquipmentIoCardOperation) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *EquipmentIoCardOperation) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *EquipmentIoCardOperation) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentIoCardOperation) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentIoCardOperation) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentIoCardOperation) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentIoCardOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentIoCardOperation) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentIoCardOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetIoCard returns the IoCard field value if set, zero value otherwise.
func (o *EquipmentIoCardOperation) GetIoCard() EquipmentIoCardRelationship {
	if o == nil || o.IoCard == nil {
		var ret EquipmentIoCardRelationship
		return ret
	}
	return *o.IoCard
}

// GetIoCardOk returns a tuple with the IoCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardOperation) GetIoCardOk() (*EquipmentIoCardRelationship, bool) {
	if o == nil || o.IoCard == nil {
		return nil, false
	}
	return o.IoCard, true
}

// HasIoCard returns a boolean if a field has been set.
func (o *EquipmentIoCardOperation) HasIoCard() bool {
	if o != nil && o.IoCard != nil {
		return true
	}

	return false
}

// SetIoCard gets a reference to the given EquipmentIoCardRelationship and assigns it to the IoCard field.
func (o *EquipmentIoCardOperation) SetIoCard(v EquipmentIoCardRelationship) {
	o.IoCard = &v
}

func (o EquipmentIoCardOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
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

func (o *EquipmentIoCardOperation) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentIoCardOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// User configured power state of the iomodule. * `None` - Placeholder default value for iom power state property. * `Reboot` - IO Module reboot state property value.
		AdminPowerState *string `json:"AdminPowerState,omitempty"`
		// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis iomodule. Applying - This state denotes that the settings are being applied in the target chassis iomodule. Failed - This state denotes that the settings could not be applied in the target chassis iomodule. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState        *string                              `json:"ConfigState,omitempty"`
		DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
		IoCard             *EquipmentIoCardRelationship         `json:"IoCard,omitempty"`
	}

	varEquipmentIoCardOperationWithoutEmbeddedStruct := EquipmentIoCardOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentIoCardOperationWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCardOperation := _EquipmentIoCardOperation{}
		varEquipmentIoCardOperation.ClassId = varEquipmentIoCardOperationWithoutEmbeddedStruct.ClassId
		varEquipmentIoCardOperation.ObjectType = varEquipmentIoCardOperationWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCardOperation.AdminPowerState = varEquipmentIoCardOperationWithoutEmbeddedStruct.AdminPowerState
		varEquipmentIoCardOperation.ConfigState = varEquipmentIoCardOperationWithoutEmbeddedStruct.ConfigState
		varEquipmentIoCardOperation.DeviceRegistration = varEquipmentIoCardOperationWithoutEmbeddedStruct.DeviceRegistration
		varEquipmentIoCardOperation.IoCard = varEquipmentIoCardOperationWithoutEmbeddedStruct.IoCard
		*o = EquipmentIoCardOperation(varEquipmentIoCardOperation)
	} else {
		return err
	}

	varEquipmentIoCardOperation := _EquipmentIoCardOperation{}

	err = json.Unmarshal(bytes, &varEquipmentIoCardOperation)
	if err == nil {
		o.MoBaseMo = varEquipmentIoCardOperation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminPowerState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "DeviceRegistration")
		delete(additionalProperties, "IoCard")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableEquipmentIoCardOperation struct {
	value *EquipmentIoCardOperation
	isSet bool
}

func (v NullableEquipmentIoCardOperation) Get() *EquipmentIoCardOperation {
	return v.value
}

func (v *NullableEquipmentIoCardOperation) Set(val *EquipmentIoCardOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardOperation(val *EquipmentIoCardOperation) *NullableEquipmentIoCardOperation {
	return &NullableEquipmentIoCardOperation{value: val, isSet: true}
}

func (v NullableEquipmentIoCardOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
