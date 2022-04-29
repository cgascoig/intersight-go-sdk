/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppDataIpInterfaceAllOf Definition of the list of properties defined in 'storage.NetAppDataIpInterface', excluding properties defined in parent classes.
type StorageNetAppDataIpInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                         `json:"ObjectType"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppDataIpInterfaceEvent resources.
	Events               []StorageNetAppDataIpInterfaceEventRelationship `json:"Events,omitempty"`
	NetAppEthernetPort   *StorageNetAppEthernetPortRelationship          `json:"NetAppEthernetPort,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship             `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppDataIpInterfaceAllOf StorageNetAppDataIpInterfaceAllOf

// NewStorageNetAppDataIpInterfaceAllOf instantiates a new StorageNetAppDataIpInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppDataIpInterfaceAllOf(classId string, objectType string) *StorageNetAppDataIpInterfaceAllOf {
	this := StorageNetAppDataIpInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppDataIpInterfaceAllOfWithDefaults instantiates a new StorageNetAppDataIpInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppDataIpInterfaceAllOfWithDefaults() *StorageNetAppDataIpInterfaceAllOf {
	this := StorageNetAppDataIpInterfaceAllOf{}
	var classId string = "storage.NetAppDataIpInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppDataIpInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppDataIpInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppDataIpInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppDataIpInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppDataIpInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppDataIpInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppDataIpInterfaceAllOf) GetEvents() []StorageNetAppDataIpInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppDataIpInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppDataIpInterfaceAllOf) GetEventsOk() (*[]StorageNetAppDataIpInterfaceEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppDataIpInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppDataIpInterfaceAllOf) SetEvents(v []StorageNetAppDataIpInterfaceEventRelationship) {
	o.Events = v
}

// GetNetAppEthernetPort returns the NetAppEthernetPort field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterfaceAllOf) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || o.NetAppEthernetPort == nil {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.NetAppEthernetPort
}

// GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil || o.NetAppEthernetPort == nil {
		return nil, false
	}
	return o.NetAppEthernetPort, true
}

// HasNetAppEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) HasNetAppEthernetPort() bool {
	if o != nil && o.NetAppEthernetPort != nil {
		return true
	}

	return false
}

// SetNetAppEthernetPort gets a reference to the given StorageNetAppEthernetPortRelationship and assigns it to the NetAppEthernetPort field.
func (o *StorageNetAppDataIpInterfaceAllOf) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.NetAppEthernetPort = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterfaceAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterfaceAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppDataIpInterfaceAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppDataIpInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.NetAppEthernetPort != nil {
		toSerialize["NetAppEthernetPort"] = o.NetAppEthernetPort
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppDataIpInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppDataIpInterfaceAllOf := _StorageNetAppDataIpInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppDataIpInterfaceAllOf); err == nil {
		*o = StorageNetAppDataIpInterfaceAllOf(varStorageNetAppDataIpInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "NetAppEthernetPort")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppDataIpInterfaceAllOf struct {
	value *StorageNetAppDataIpInterfaceAllOf
	isSet bool
}

func (v NullableStorageNetAppDataIpInterfaceAllOf) Get() *StorageNetAppDataIpInterfaceAllOf {
	return v.value
}

func (v *NullableStorageNetAppDataIpInterfaceAllOf) Set(val *StorageNetAppDataIpInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDataIpInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDataIpInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDataIpInterfaceAllOf(val *StorageNetAppDataIpInterfaceAllOf) *NullableStorageNetAppDataIpInterfaceAllOf {
	return &NullableStorageNetAppDataIpInterfaceAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppDataIpInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDataIpInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
