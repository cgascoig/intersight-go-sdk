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
)

// BootPchStorageDeviceAllOf Definition of the list of properties defined in 'boot.PchStorageDevice', excluding properties defined in parent classes.
type BootPchStorageDeviceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	ComputePhysical      *ComputePhysicalRelationship         `json:"ComputePhysical,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootPchStorageDeviceAllOf BootPchStorageDeviceAllOf

// NewBootPchStorageDeviceAllOf instantiates a new BootPchStorageDeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootPchStorageDeviceAllOf(classId string, objectType string) *BootPchStorageDeviceAllOf {
	this := BootPchStorageDeviceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootPchStorageDeviceAllOfWithDefaults instantiates a new BootPchStorageDeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootPchStorageDeviceAllOfWithDefaults() *BootPchStorageDeviceAllOf {
	this := BootPchStorageDeviceAllOf{}
	var classId string = "boot.PchStorageDevice"
	this.ClassId = classId
	var objectType string = "boot.PchStorageDevice"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootPchStorageDeviceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootPchStorageDeviceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootPchStorageDeviceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootPchStorageDeviceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootPchStorageDeviceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootPchStorageDeviceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComputePhysical returns the ComputePhysical field value if set, zero value otherwise.
func (o *BootPchStorageDeviceAllOf) GetComputePhysical() ComputePhysicalRelationship {
	if o == nil || o.ComputePhysical == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.ComputePhysical
}

// GetComputePhysicalOk returns a tuple with the ComputePhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPchStorageDeviceAllOf) GetComputePhysicalOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.ComputePhysical == nil {
		return nil, false
	}
	return o.ComputePhysical, true
}

// HasComputePhysical returns a boolean if a field has been set.
func (o *BootPchStorageDeviceAllOf) HasComputePhysical() bool {
	if o != nil && o.ComputePhysical != nil {
		return true
	}

	return false
}

// SetComputePhysical gets a reference to the given ComputePhysicalRelationship and assigns it to the ComputePhysical field.
func (o *BootPchStorageDeviceAllOf) SetComputePhysical(v ComputePhysicalRelationship) {
	o.ComputePhysical = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *BootPchStorageDeviceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPchStorageDeviceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BootPchStorageDeviceAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BootPchStorageDeviceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *BootPchStorageDeviceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootPchStorageDeviceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BootPchStorageDeviceAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BootPchStorageDeviceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o BootPchStorageDeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ComputePhysical != nil {
		toSerialize["ComputePhysical"] = o.ComputePhysical
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

func (o *BootPchStorageDeviceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootPchStorageDeviceAllOf := _BootPchStorageDeviceAllOf{}

	if err = json.Unmarshal(bytes, &varBootPchStorageDeviceAllOf); err == nil {
		*o = BootPchStorageDeviceAllOf(varBootPchStorageDeviceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ComputePhysical")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootPchStorageDeviceAllOf struct {
	value *BootPchStorageDeviceAllOf
	isSet bool
}

func (v NullableBootPchStorageDeviceAllOf) Get() *BootPchStorageDeviceAllOf {
	return v.value
}

func (v *NullableBootPchStorageDeviceAllOf) Set(val *BootPchStorageDeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootPchStorageDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootPchStorageDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootPchStorageDeviceAllOf(val *BootPchStorageDeviceAllOf) *NullableBootPchStorageDeviceAllOf {
	return &NullableBootPchStorageDeviceAllOf{value: val, isSet: true}
}

func (v NullableBootPchStorageDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootPchStorageDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
