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

// EquipmentChassisOperation Models the configurable properties of Chassis.
type EquipmentChassisOperation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured state of the locator LED for the Chassis. * `None` - No operation action for the Locator Led of an equipment. * `TurnOn` - Turn on the Locator Led of an equipment. * `TurnOff` - Turn off the Locator Led of an equipment.
	AdminLocatorLedAction *string `json:"AdminLocatorLedAction,omitempty"`
	// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis. Applying - This state denotes that the settings are being applied in the target chassis. Failed - This state denotes that the settings could not be applied in the target chassis. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState          *string                              `json:"ConfigState,omitempty"`
	Chassis              *EquipmentChassisRelationship        `json:"Chassis,omitempty"`
	DeviceRegistration   *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentChassisOperation EquipmentChassisOperation

// NewEquipmentChassisOperation instantiates a new EquipmentChassisOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentChassisOperation(classId string, objectType string) *EquipmentChassisOperation {
	this := EquipmentChassisOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	return &this
}

// NewEquipmentChassisOperationWithDefaults instantiates a new EquipmentChassisOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentChassisOperationWithDefaults() *EquipmentChassisOperation {
	this := EquipmentChassisOperation{}
	var classId string = "equipment.ChassisOperation"
	this.ClassId = classId
	var objectType string = "equipment.ChassisOperation"
	this.ObjectType = objectType
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentChassisOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentChassisOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentChassisOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentChassisOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminLocatorLedAction returns the AdminLocatorLedAction field value if set, zero value otherwise.
func (o *EquipmentChassisOperation) GetAdminLocatorLedAction() string {
	if o == nil || o.AdminLocatorLedAction == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedAction
}

// GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetAdminLocatorLedActionOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedAction == nil {
		return nil, false
	}
	return o.AdminLocatorLedAction, true
}

// HasAdminLocatorLedAction returns a boolean if a field has been set.
func (o *EquipmentChassisOperation) HasAdminLocatorLedAction() bool {
	if o != nil && o.AdminLocatorLedAction != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedAction gets a reference to the given string and assigns it to the AdminLocatorLedAction field.
func (o *EquipmentChassisOperation) SetAdminLocatorLedAction(v string) {
	o.AdminLocatorLedAction = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentChassisOperation) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentChassisOperation) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentChassisOperation) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *EquipmentChassisOperation) GetChassis() EquipmentChassisRelationship {
	if o == nil || o.Chassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *EquipmentChassisOperation) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *EquipmentChassisOperation) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentChassisOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentChassisOperation) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentChassisOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

func (o EquipmentChassisOperation) MarshalJSON() ([]byte, error) {
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
	if o.AdminLocatorLedAction != nil {
		toSerialize["AdminLocatorLedAction"] = o.AdminLocatorLedAction
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentChassisOperation) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentChassisOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// User configured state of the locator LED for the Chassis. * `None` - No operation action for the Locator Led of an equipment. * `TurnOn` - Turn on the Locator Led of an equipment. * `TurnOff` - Turn off the Locator Led of an equipment.
		AdminLocatorLedAction *string `json:"AdminLocatorLedAction,omitempty"`
		// The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target chassis. Applying - This state denotes that the settings are being applied in the target chassis. Failed - This state denotes that the settings could not be applied in the target chassis. * `None` - Nil value when no action has been triggered by the user. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState        *string                              `json:"ConfigState,omitempty"`
		Chassis            *EquipmentChassisRelationship        `json:"Chassis,omitempty"`
		DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	}

	varEquipmentChassisOperationWithoutEmbeddedStruct := EquipmentChassisOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentChassisOperationWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentChassisOperation := _EquipmentChassisOperation{}
		varEquipmentChassisOperation.ClassId = varEquipmentChassisOperationWithoutEmbeddedStruct.ClassId
		varEquipmentChassisOperation.ObjectType = varEquipmentChassisOperationWithoutEmbeddedStruct.ObjectType
		varEquipmentChassisOperation.AdminLocatorLedAction = varEquipmentChassisOperationWithoutEmbeddedStruct.AdminLocatorLedAction
		varEquipmentChassisOperation.ConfigState = varEquipmentChassisOperationWithoutEmbeddedStruct.ConfigState
		varEquipmentChassisOperation.Chassis = varEquipmentChassisOperationWithoutEmbeddedStruct.Chassis
		varEquipmentChassisOperation.DeviceRegistration = varEquipmentChassisOperationWithoutEmbeddedStruct.DeviceRegistration
		*o = EquipmentChassisOperation(varEquipmentChassisOperation)
	} else {
		return err
	}

	varEquipmentChassisOperation := _EquipmentChassisOperation{}

	err = json.Unmarshal(bytes, &varEquipmentChassisOperation)
	if err == nil {
		o.MoBaseMo = varEquipmentChassisOperation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminLocatorLedAction")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "DeviceRegistration")

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

type NullableEquipmentChassisOperation struct {
	value *EquipmentChassisOperation
	isSet bool
}

func (v NullableEquipmentChassisOperation) Get() *EquipmentChassisOperation {
	return v.value
}

func (v *NullableEquipmentChassisOperation) Set(val *EquipmentChassisOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentChassisOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentChassisOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentChassisOperation(val *EquipmentChassisOperation) *NullableEquipmentChassisOperation {
	return &NullableEquipmentChassisOperation{value: val, isSet: true}
}

func (v NullableEquipmentChassisOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentChassisOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
