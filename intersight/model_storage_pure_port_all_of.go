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

// StoragePurePortAllOf Definition of the list of properties defined in 'storage.PurePort', excluding properties defined in parent classes.
type StoragePurePortAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the port to which this port has failed over.
	Failover *string `json:"Failover,omitempty"`
	// Ip address of iSCSI portal configured on the port.
	Portal               *string                              `json:"Portal,omitempty"`
	Array                *StoragePureArrayRelationship        `json:"Array,omitempty"`
	Controller           *StoragePureControllerRelationship   `json:"Controller,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePurePortAllOf StoragePurePortAllOf

// NewStoragePurePortAllOf instantiates a new StoragePurePortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePurePortAllOf(classId string, objectType string) *StoragePurePortAllOf {
	this := StoragePurePortAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePurePortAllOfWithDefaults instantiates a new StoragePurePortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePurePortAllOfWithDefaults() *StoragePurePortAllOf {
	this := StoragePurePortAllOf{}
	var classId string = "storage.PurePort"
	this.ClassId = classId
	var objectType string = "storage.PurePort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePurePortAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePurePortAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePurePortAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePurePortAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFailover returns the Failover field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetFailover() string {
	if o == nil || o.Failover == nil {
		var ret string
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetFailoverOk() (*string, bool) {
	if o == nil || o.Failover == nil {
		return nil, false
	}
	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasFailover() bool {
	if o != nil && o.Failover != nil {
		return true
	}

	return false
}

// SetFailover gets a reference to the given string and assigns it to the Failover field.
func (o *StoragePurePortAllOf) SetFailover(v string) {
	o.Failover = &v
}

// GetPortal returns the Portal field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetPortal() string {
	if o == nil || o.Portal == nil {
		var ret string
		return ret
	}
	return *o.Portal
}

// GetPortalOk returns a tuple with the Portal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetPortalOk() (*string, bool) {
	if o == nil || o.Portal == nil {
		return nil, false
	}
	return o.Portal, true
}

// HasPortal returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasPortal() bool {
	if o != nil && o.Portal != nil {
		return true
	}

	return false
}

// SetPortal gets a reference to the given string and assigns it to the Portal field.
func (o *StoragePurePortAllOf) SetPortal(v string) {
	o.Portal = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePurePortAllOf) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetController() StoragePureControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret StoragePureControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetControllerOk() (*StoragePureControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given StoragePureControllerRelationship and assigns it to the Controller field.
func (o *StoragePurePortAllOf) SetController(v StoragePureControllerRelationship) {
	o.Controller = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePurePortAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePurePortAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePurePortAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePurePortAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePurePortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Failover != nil {
		toSerialize["Failover"] = o.Failover
	}
	if o.Portal != nil {
		toSerialize["Portal"] = o.Portal
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePurePortAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStoragePurePortAllOf := _StoragePurePortAllOf{}

	if err = json.Unmarshal(bytes, &varStoragePurePortAllOf); err == nil {
		*o = StoragePurePortAllOf(varStoragePurePortAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Failover")
		delete(additionalProperties, "Portal")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStoragePurePortAllOf struct {
	value *StoragePurePortAllOf
	isSet bool
}

func (v NullableStoragePurePortAllOf) Get() *StoragePurePortAllOf {
	return v.value
}

func (v *NullableStoragePurePortAllOf) Set(val *StoragePurePortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePurePortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePurePortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePurePortAllOf(val *StoragePurePortAllOf) *NullableStoragePurePortAllOf {
	return &NullableStoragePurePortAllOf{value: val, isSet: true}
}

func (v NullableStoragePurePortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePurePortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
