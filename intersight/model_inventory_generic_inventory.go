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

// InventoryGenericInventory Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.
type InventoryGenericInventory struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Key of inventory data for Generic Inventory data set.
	Key *string `json:"Key,omitempty"`
	// Type of inventory data for Generic Inventory data set.
	Type *string `json:"Type,omitempty"`
	// Value of inventory data for Generic Inventory data set.
	Value                           *string                                      `json:"Value,omitempty"`
	InventoryDeviceInfo             *InventoryDeviceInfoRelationship             `json:"InventoryDeviceInfo,omitempty"`
	InventoryGenericInventoryHolder *InventoryGenericInventoryHolderRelationship `json:"InventoryGenericInventoryHolder,omitempty"`
	RegisteredDevice                *AssetDeviceRegistrationRelationship         `json:"RegisteredDevice,omitempty"`
	AdditionalProperties            map[string]interface{}
}

type _InventoryGenericInventory InventoryGenericInventory

// NewInventoryGenericInventory instantiates a new InventoryGenericInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryGenericInventory(classId string, objectType string) *InventoryGenericInventory {
	this := InventoryGenericInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInventoryGenericInventoryWithDefaults instantiates a new InventoryGenericInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryGenericInventoryWithDefaults() *InventoryGenericInventory {
	this := InventoryGenericInventory{}
	var classId string = "inventory.GenericInventory"
	this.ClassId = classId
	var objectType string = "inventory.GenericInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *InventoryGenericInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InventoryGenericInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InventoryGenericInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InventoryGenericInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InventoryGenericInventory) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InventoryGenericInventory) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *InventoryGenericInventory) SetValue(v string) {
	o.Value = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *InventoryGenericInventory) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetInventoryGenericInventoryHolder returns the InventoryGenericInventoryHolder field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetInventoryGenericInventoryHolder() InventoryGenericInventoryHolderRelationship {
	if o == nil || o.InventoryGenericInventoryHolder == nil {
		var ret InventoryGenericInventoryHolderRelationship
		return ret
	}
	return *o.InventoryGenericInventoryHolder
}

// GetInventoryGenericInventoryHolderOk returns a tuple with the InventoryGenericInventoryHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetInventoryGenericInventoryHolderOk() (*InventoryGenericInventoryHolderRelationship, bool) {
	if o == nil || o.InventoryGenericInventoryHolder == nil {
		return nil, false
	}
	return o.InventoryGenericInventoryHolder, true
}

// HasInventoryGenericInventoryHolder returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasInventoryGenericInventoryHolder() bool {
	if o != nil && o.InventoryGenericInventoryHolder != nil {
		return true
	}

	return false
}

// SetInventoryGenericInventoryHolder gets a reference to the given InventoryGenericInventoryHolderRelationship and assigns it to the InventoryGenericInventoryHolder field.
func (o *InventoryGenericInventory) SetInventoryGenericInventoryHolder(v InventoryGenericInventoryHolderRelationship) {
	o.InventoryGenericInventoryHolder = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *InventoryGenericInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryGenericInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *InventoryGenericInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *InventoryGenericInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o InventoryGenericInventory) MarshalJSON() ([]byte, error) {
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
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.InventoryGenericInventoryHolder != nil {
		toSerialize["InventoryGenericInventoryHolder"] = o.InventoryGenericInventoryHolder
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InventoryGenericInventory) UnmarshalJSON(bytes []byte) (err error) {
	type InventoryGenericInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Key of inventory data for Generic Inventory data set.
		Key *string `json:"Key,omitempty"`
		// Type of inventory data for Generic Inventory data set.
		Type *string `json:"Type,omitempty"`
		// Value of inventory data for Generic Inventory data set.
		Value                           *string                                      `json:"Value,omitempty"`
		InventoryDeviceInfo             *InventoryDeviceInfoRelationship             `json:"InventoryDeviceInfo,omitempty"`
		InventoryGenericInventoryHolder *InventoryGenericInventoryHolderRelationship `json:"InventoryGenericInventoryHolder,omitempty"`
		RegisteredDevice                *AssetDeviceRegistrationRelationship         `json:"RegisteredDevice,omitempty"`
	}

	varInventoryGenericInventoryWithoutEmbeddedStruct := InventoryGenericInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varInventoryGenericInventoryWithoutEmbeddedStruct)
	if err == nil {
		varInventoryGenericInventory := _InventoryGenericInventory{}
		varInventoryGenericInventory.ClassId = varInventoryGenericInventoryWithoutEmbeddedStruct.ClassId
		varInventoryGenericInventory.ObjectType = varInventoryGenericInventoryWithoutEmbeddedStruct.ObjectType
		varInventoryGenericInventory.Key = varInventoryGenericInventoryWithoutEmbeddedStruct.Key
		varInventoryGenericInventory.Type = varInventoryGenericInventoryWithoutEmbeddedStruct.Type
		varInventoryGenericInventory.Value = varInventoryGenericInventoryWithoutEmbeddedStruct.Value
		varInventoryGenericInventory.InventoryDeviceInfo = varInventoryGenericInventoryWithoutEmbeddedStruct.InventoryDeviceInfo
		varInventoryGenericInventory.InventoryGenericInventoryHolder = varInventoryGenericInventoryWithoutEmbeddedStruct.InventoryGenericInventoryHolder
		varInventoryGenericInventory.RegisteredDevice = varInventoryGenericInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = InventoryGenericInventory(varInventoryGenericInventory)
	} else {
		return err
	}

	varInventoryGenericInventory := _InventoryGenericInventory{}

	err = json.Unmarshal(bytes, &varInventoryGenericInventory)
	if err == nil {
		o.InventoryBase = varInventoryGenericInventory.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Value")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "InventoryGenericInventoryHolder")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableInventoryGenericInventory struct {
	value *InventoryGenericInventory
	isSet bool
}

func (v NullableInventoryGenericInventory) Get() *InventoryGenericInventory {
	return v.value
}

func (v *NullableInventoryGenericInventory) Set(val *InventoryGenericInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryGenericInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryGenericInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryGenericInventory(val *InventoryGenericInventory) *NullableInventoryGenericInventory {
	return &NullableInventoryGenericInventory{value: val, isSet: true}
}

func (v NullableInventoryGenericInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryGenericInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
