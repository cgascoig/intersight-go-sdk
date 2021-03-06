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

// StorageNetAppNonDataIpInterfaceEventAllOf Definition of the list of properties defined in 'storage.NetAppNonDataIpInterfaceEvent', excluding properties defined in parent classes.
type StorageNetAppNonDataIpInterfaceEventAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                       `json:"ObjectType"`
	IpInterface          *StorageNetAppNonDataIpInterfaceRelationship `json:"IpInterface,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNonDataIpInterfaceEventAllOf StorageNetAppNonDataIpInterfaceEventAllOf

// NewStorageNetAppNonDataIpInterfaceEventAllOf instantiates a new StorageNetAppNonDataIpInterfaceEventAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNonDataIpInterfaceEventAllOf(classId string, objectType string) *StorageNetAppNonDataIpInterfaceEventAllOf {
	this := StorageNetAppNonDataIpInterfaceEventAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNonDataIpInterfaceEventAllOfWithDefaults instantiates a new StorageNetAppNonDataIpInterfaceEventAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNonDataIpInterfaceEventAllOfWithDefaults() *StorageNetAppNonDataIpInterfaceEventAllOf {
	this := StorageNetAppNonDataIpInterfaceEventAllOf{}
	var classId string = "storage.NetAppNonDataIpInterfaceEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppNonDataIpInterfaceEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpInterface returns the IpInterface field value if set, zero value otherwise.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetIpInterface() StorageNetAppNonDataIpInterfaceRelationship {
	if o == nil || o.IpInterface == nil {
		var ret StorageNetAppNonDataIpInterfaceRelationship
		return ret
	}
	return *o.IpInterface
}

// GetIpInterfaceOk returns a tuple with the IpInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) GetIpInterfaceOk() (*StorageNetAppNonDataIpInterfaceRelationship, bool) {
	if o == nil || o.IpInterface == nil {
		return nil, false
	}
	return o.IpInterface, true
}

// HasIpInterface returns a boolean if a field has been set.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) HasIpInterface() bool {
	if o != nil && o.IpInterface != nil {
		return true
	}

	return false
}

// SetIpInterface gets a reference to the given StorageNetAppNonDataIpInterfaceRelationship and assigns it to the IpInterface field.
func (o *StorageNetAppNonDataIpInterfaceEventAllOf) SetIpInterface(v StorageNetAppNonDataIpInterfaceRelationship) {
	o.IpInterface = &v
}

func (o StorageNetAppNonDataIpInterfaceEventAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpInterface != nil {
		toSerialize["IpInterface"] = o.IpInterface
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNonDataIpInterfaceEventAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNonDataIpInterfaceEventAllOf := _StorageNetAppNonDataIpInterfaceEventAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNonDataIpInterfaceEventAllOf); err == nil {
		*o = StorageNetAppNonDataIpInterfaceEventAllOf(varStorageNetAppNonDataIpInterfaceEventAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpInterface")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNonDataIpInterfaceEventAllOf struct {
	value *StorageNetAppNonDataIpInterfaceEventAllOf
	isSet bool
}

func (v NullableStorageNetAppNonDataIpInterfaceEventAllOf) Get() *StorageNetAppNonDataIpInterfaceEventAllOf {
	return v.value
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventAllOf) Set(val *StorageNetAppNonDataIpInterfaceEventAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNonDataIpInterfaceEventAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNonDataIpInterfaceEventAllOf(val *StorageNetAppNonDataIpInterfaceEventAllOf) *NullableStorageNetAppNonDataIpInterfaceEventAllOf {
	return &NullableStorageNetAppNonDataIpInterfaceEventAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNonDataIpInterfaceEventAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNonDataIpInterfaceEventAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
