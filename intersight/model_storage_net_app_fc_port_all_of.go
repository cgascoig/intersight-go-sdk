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

// StorageNetAppFcPortAllOf Definition of the list of properties defined in 'storage.NetAppFcPort', excluding properties defined in parent classes.
type StorageNetAppFcPortAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of storage array port.
	PortStatus *string `json:"PortStatus,omitempty"`
	// State of the port available in storage array. * `Unknown` - Default unknown port state. * `StartUp` - The port in the storage array is booting up. * `LinkNotConnected` - The port has finished initialization, but a link with the fabric is not established. * `Online` - The port is initialized and a link with the fabric has been established. * `LinkDisconnected` - The link on this port is currently not established. * `OfflineUser` - The port is administratively disabled. * `OfflineSystem` - The port is set to offline by the system. This happens when the port encounters too many errors. * `NodeOffline` - The state information for the port cannot be retrieved. The node is offline or inaccessible.
	State *string `json:"State,omitempty"`
	// Universally unique identifier of the FC port.
	Uuid            *string                        `json:"Uuid,omitempty"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppFcPortEvent resources.
	Events               []StorageNetAppFcPortEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppFcPortAllOf StorageNetAppFcPortAllOf

// NewStorageNetAppFcPortAllOf instantiates a new StorageNetAppFcPortAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppFcPortAllOf(classId string, objectType string) *StorageNetAppFcPortAllOf {
	this := StorageNetAppFcPortAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppFcPortAllOfWithDefaults instantiates a new StorageNetAppFcPortAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppFcPortAllOfWithDefaults() *StorageNetAppFcPortAllOf {
	this := StorageNetAppFcPortAllOf{}
	var classId string = "storage.NetAppFcPort"
	this.ClassId = classId
	var objectType string = "storage.NetAppFcPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppFcPortAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppFcPortAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppFcPortAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppFcPortAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPortStatus returns the PortStatus field value if set, zero value otherwise.
func (o *StorageNetAppFcPortAllOf) GetPortStatus() string {
	if o == nil || o.PortStatus == nil {
		var ret string
		return ret
	}
	return *o.PortStatus
}

// GetPortStatusOk returns a tuple with the PortStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetPortStatusOk() (*string, bool) {
	if o == nil || o.PortStatus == nil {
		return nil, false
	}
	return o.PortStatus, true
}

// HasPortStatus returns a boolean if a field has been set.
func (o *StorageNetAppFcPortAllOf) HasPortStatus() bool {
	if o != nil && o.PortStatus != nil {
		return true
	}

	return false
}

// SetPortStatus gets a reference to the given string and assigns it to the PortStatus field.
func (o *StorageNetAppFcPortAllOf) SetPortStatus(v string) {
	o.PortStatus = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppFcPortAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppFcPortAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppFcPortAllOf) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppFcPortAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppFcPortAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppFcPortAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppFcPortAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppFcPortAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppFcPortAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppFcPortAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppFcPortAllOf) GetEvents() []StorageNetAppFcPortEventRelationship {
	if o == nil {
		var ret []StorageNetAppFcPortEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppFcPortAllOf) GetEventsOk() (*[]StorageNetAppFcPortEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppFcPortAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppFcPortEventRelationship and assigns it to the Events field.
func (o *StorageNetAppFcPortAllOf) SetEvents(v []StorageNetAppFcPortEventRelationship) {
	o.Events = v
}

func (o StorageNetAppFcPortAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PortStatus != nil {
		toSerialize["PortStatus"] = o.PortStatus
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppFcPortAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppFcPortAllOf := _StorageNetAppFcPortAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppFcPortAllOf); err == nil {
		*o = StorageNetAppFcPortAllOf(varStorageNetAppFcPortAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PortStatus")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppFcPortAllOf struct {
	value *StorageNetAppFcPortAllOf
	isSet bool
}

func (v NullableStorageNetAppFcPortAllOf) Get() *StorageNetAppFcPortAllOf {
	return v.value
}

func (v *NullableStorageNetAppFcPortAllOf) Set(val *StorageNetAppFcPortAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppFcPortAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppFcPortAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppFcPortAllOf(val *StorageNetAppFcPortAllOf) *NullableStorageNetAppFcPortAllOf {
	return &NullableStorageNetAppFcPortAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppFcPortAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppFcPortAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
