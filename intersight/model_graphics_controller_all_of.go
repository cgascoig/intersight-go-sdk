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
)

// GraphicsControllerAllOf Definition of the list of properties defined in 'graphics.Controller', excluding properties defined in parent classes.
type GraphicsControllerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the graphics controller.
	ControllerId *int64 `json:"ControllerId,omitempty"`
	// The PCI address of the graphics controller.
	PciAddr *string `json:"PciAddr,omitempty"`
	// The PCI slot information of the graphics controller.
	PciSlot              *string                              `json:"PciSlot,omitempty"`
	GraphicsCard         *GraphicsCardRelationship            `json:"GraphicsCard,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GraphicsControllerAllOf GraphicsControllerAllOf

// NewGraphicsControllerAllOf instantiates a new GraphicsControllerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphicsControllerAllOf(classId string, objectType string) *GraphicsControllerAllOf {
	this := GraphicsControllerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewGraphicsControllerAllOfWithDefaults instantiates a new GraphicsControllerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphicsControllerAllOfWithDefaults() *GraphicsControllerAllOf {
	this := GraphicsControllerAllOf{}
	var classId string = "graphics.Controller"
	this.ClassId = classId
	var objectType string = "graphics.Controller"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *GraphicsControllerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *GraphicsControllerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *GraphicsControllerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GraphicsControllerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetControllerId() int64 {
	if o == nil || o.ControllerId == nil {
		var ret int64
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetControllerIdOk() (*int64, bool) {
	if o == nil || o.ControllerId == nil {
		return nil, false
	}
	return o.ControllerId, true
}

// HasControllerId returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasControllerId() bool {
	if o != nil && o.ControllerId != nil {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given int64 and assigns it to the ControllerId field.
func (o *GraphicsControllerAllOf) SetControllerId(v int64) {
	o.ControllerId = &v
}

// GetPciAddr returns the PciAddr field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetPciAddr() string {
	if o == nil || o.PciAddr == nil {
		var ret string
		return ret
	}
	return *o.PciAddr
}

// GetPciAddrOk returns a tuple with the PciAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetPciAddrOk() (*string, bool) {
	if o == nil || o.PciAddr == nil {
		return nil, false
	}
	return o.PciAddr, true
}

// HasPciAddr returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasPciAddr() bool {
	if o != nil && o.PciAddr != nil {
		return true
	}

	return false
}

// SetPciAddr gets a reference to the given string and assigns it to the PciAddr field.
func (o *GraphicsControllerAllOf) SetPciAddr(v string) {
	o.PciAddr = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *GraphicsControllerAllOf) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetGraphicsCard returns the GraphicsCard field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetGraphicsCard() GraphicsCardRelationship {
	if o == nil || o.GraphicsCard == nil {
		var ret GraphicsCardRelationship
		return ret
	}
	return *o.GraphicsCard
}

// GetGraphicsCardOk returns a tuple with the GraphicsCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetGraphicsCardOk() (*GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCard == nil {
		return nil, false
	}
	return o.GraphicsCard, true
}

// HasGraphicsCard returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasGraphicsCard() bool {
	if o != nil && o.GraphicsCard != nil {
		return true
	}

	return false
}

// SetGraphicsCard gets a reference to the given GraphicsCardRelationship and assigns it to the GraphicsCard field.
func (o *GraphicsControllerAllOf) SetGraphicsCard(v GraphicsCardRelationship) {
	o.GraphicsCard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *GraphicsControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *GraphicsControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *GraphicsControllerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *GraphicsControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o GraphicsControllerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerId != nil {
		toSerialize["ControllerId"] = o.ControllerId
	}
	if o.PciAddr != nil {
		toSerialize["PciAddr"] = o.PciAddr
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.GraphicsCard != nil {
		toSerialize["GraphicsCard"] = o.GraphicsCard
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

func (o *GraphicsControllerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varGraphicsControllerAllOf := _GraphicsControllerAllOf{}

	if err = json.Unmarshal(bytes, &varGraphicsControllerAllOf); err == nil {
		*o = GraphicsControllerAllOf(varGraphicsControllerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerId")
		delete(additionalProperties, "PciAddr")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "GraphicsCard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGraphicsControllerAllOf struct {
	value *GraphicsControllerAllOf
	isSet bool
}

func (v NullableGraphicsControllerAllOf) Get() *GraphicsControllerAllOf {
	return v.value
}

func (v *NullableGraphicsControllerAllOf) Set(val *GraphicsControllerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphicsControllerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphicsControllerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphicsControllerAllOf(val *GraphicsControllerAllOf) *NullableGraphicsControllerAllOf {
	return &NullableGraphicsControllerAllOf{value: val, isSet: true}
}

func (v NullableGraphicsControllerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphicsControllerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
