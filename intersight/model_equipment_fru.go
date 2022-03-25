/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentFru Managed object for all equipments which contains the previous vendor /model / serial before insertion/replacement/removal.
type EquipmentFru struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the action performed on a component. * `None` - No action performed on the FRU. * `Inserted` - A new FRU is inserted or added. * `Removed` - The previous FRU is removed. * `Replaced` - The previous FRU is replaced with a new FRU. * `ReplacedWithAlarm` - The previous FRU is replaced with a new FRU and a alarm is raised.
	Action               *string                              `json:"Action,omitempty"`
	CurrentFru           *EquipmentBaseRelationship           `json:"CurrentFru,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentFru EquipmentFru

// NewEquipmentFru instantiates a new EquipmentFru object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFru(classId string, objectType string) *EquipmentFru {
	this := EquipmentFru{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewEquipmentFruWithDefaults instantiates a new EquipmentFru object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFruWithDefaults() *EquipmentFru {
	this := EquipmentFru{}
	var classId string = "equipment.Fru"
	this.ClassId = classId
	var objectType string = "equipment.Fru"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentFru) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentFru) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentFru) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentFru) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *EquipmentFru) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *EquipmentFru) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *EquipmentFru) SetAction(v string) {
	o.Action = &v
}

// GetCurrentFru returns the CurrentFru field value if set, zero value otherwise.
func (o *EquipmentFru) GetCurrentFru() EquipmentBaseRelationship {
	if o == nil || o.CurrentFru == nil {
		var ret EquipmentBaseRelationship
		return ret
	}
	return *o.CurrentFru
}

// GetCurrentFruOk returns a tuple with the CurrentFru field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetCurrentFruOk() (*EquipmentBaseRelationship, bool) {
	if o == nil || o.CurrentFru == nil {
		return nil, false
	}
	return o.CurrentFru, true
}

// HasCurrentFru returns a boolean if a field has been set.
func (o *EquipmentFru) HasCurrentFru() bool {
	if o != nil && o.CurrentFru != nil {
		return true
	}

	return false
}

// SetCurrentFru gets a reference to the given EquipmentBaseRelationship and assigns it to the CurrentFru field.
func (o *EquipmentFru) SetCurrentFru(v EquipmentBaseRelationship) {
	o.CurrentFru = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentFru) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentFru) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentFru) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentFru) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFru) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentFru) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentFru) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentFru) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.CurrentFru != nil {
		toSerialize["CurrentFru"] = o.CurrentFru
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

func (o *EquipmentFru) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentFruWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This field identifies the action performed on a component. * `None` - No action performed on the FRU. * `Inserted` - A new FRU is inserted or added. * `Removed` - The previous FRU is removed. * `Replaced` - The previous FRU is replaced with a new FRU. * `ReplacedWithAlarm` - The previous FRU is replaced with a new FRU and a alarm is raised.
		Action              *string                              `json:"Action,omitempty"`
		CurrentFru          *EquipmentBaseRelationship           `json:"CurrentFru,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentFruWithoutEmbeddedStruct := EquipmentFruWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentFruWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentFru := _EquipmentFru{}
		varEquipmentFru.ClassId = varEquipmentFruWithoutEmbeddedStruct.ClassId
		varEquipmentFru.ObjectType = varEquipmentFruWithoutEmbeddedStruct.ObjectType
		varEquipmentFru.Action = varEquipmentFruWithoutEmbeddedStruct.Action
		varEquipmentFru.CurrentFru = varEquipmentFruWithoutEmbeddedStruct.CurrentFru
		varEquipmentFru.InventoryDeviceInfo = varEquipmentFruWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentFru.RegisteredDevice = varEquipmentFruWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentFru(varEquipmentFru)
	} else {
		return err
	}

	varEquipmentFru := _EquipmentFru{}

	err = json.Unmarshal(bytes, &varEquipmentFru)
	if err == nil {
		o.EquipmentBase = varEquipmentFru.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "CurrentFru")
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

type NullableEquipmentFru struct {
	value *EquipmentFru
	isSet bool
}

func (v NullableEquipmentFru) Get() *EquipmentFru {
	return v.value
}

func (v *NullableEquipmentFru) Set(val *EquipmentFru) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFru) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFru) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFru(val *EquipmentFru) *NullableEquipmentFru {
	return &NullableEquipmentFru{value: val, isSet: true}
}

func (v NullableEquipmentFru) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFru) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
