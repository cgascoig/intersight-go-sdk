/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// StorageNetAppIpInterfaceEvent An event where the impacted resource type is an ip interface.
type StorageNetAppIpInterfaceEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	IpInterface          *StorageNetAppIpInterfaceRelationship `json:"IpInterface,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppIpInterfaceEvent StorageNetAppIpInterfaceEvent

// NewStorageNetAppIpInterfaceEvent instantiates a new StorageNetAppIpInterfaceEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppIpInterfaceEvent(classId string, objectType string) *StorageNetAppIpInterfaceEvent {
	this := StorageNetAppIpInterfaceEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppIpInterfaceEventWithDefaults instantiates a new StorageNetAppIpInterfaceEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppIpInterfaceEventWithDefaults() *StorageNetAppIpInterfaceEvent {
	this := StorageNetAppIpInterfaceEvent{}
	var classId string = "storage.NetAppIpInterfaceEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppIpInterfaceEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppIpInterfaceEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterfaceEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppIpInterfaceEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppIpInterfaceEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterfaceEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppIpInterfaceEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpInterface returns the IpInterface field value if set, zero value otherwise.
func (o *StorageNetAppIpInterfaceEvent) GetIpInterface() StorageNetAppIpInterfaceRelationship {
	if o == nil || o.IpInterface == nil {
		var ret StorageNetAppIpInterfaceRelationship
		return ret
	}
	return *o.IpInterface
}

// GetIpInterfaceOk returns a tuple with the IpInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppIpInterfaceEvent) GetIpInterfaceOk() (*StorageNetAppIpInterfaceRelationship, bool) {
	if o == nil || o.IpInterface == nil {
		return nil, false
	}
	return o.IpInterface, true
}

// HasIpInterface returns a boolean if a field has been set.
func (o *StorageNetAppIpInterfaceEvent) HasIpInterface() bool {
	if o != nil && o.IpInterface != nil {
		return true
	}

	return false
}

// SetIpInterface gets a reference to the given StorageNetAppIpInterfaceRelationship and assigns it to the IpInterface field.
func (o *StorageNetAppIpInterfaceEvent) SetIpInterface(v StorageNetAppIpInterfaceRelationship) {
	o.IpInterface = &v
}

func (o StorageNetAppIpInterfaceEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseEvent, errStorageNetAppBaseEvent := json.Marshal(o.StorageNetAppBaseEvent)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
	errStorageNetAppBaseEvent = json.Unmarshal([]byte(serializedStorageNetAppBaseEvent), &toSerialize)
	if errStorageNetAppBaseEvent != nil {
		return []byte{}, errStorageNetAppBaseEvent
	}
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

func (o *StorageNetAppIpInterfaceEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppIpInterfaceEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                                `json:"ObjectType"`
		IpInterface *StorageNetAppIpInterfaceRelationship `json:"IpInterface,omitempty"`
	}

	varStorageNetAppIpInterfaceEventWithoutEmbeddedStruct := StorageNetAppIpInterfaceEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppIpInterfaceEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppIpInterfaceEvent := _StorageNetAppIpInterfaceEvent{}
		varStorageNetAppIpInterfaceEvent.ClassId = varStorageNetAppIpInterfaceEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppIpInterfaceEvent.ObjectType = varStorageNetAppIpInterfaceEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppIpInterfaceEvent.IpInterface = varStorageNetAppIpInterfaceEventWithoutEmbeddedStruct.IpInterface
		*o = StorageNetAppIpInterfaceEvent(varStorageNetAppIpInterfaceEvent)
	} else {
		return err
	}

	varStorageNetAppIpInterfaceEvent := _StorageNetAppIpInterfaceEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppIpInterfaceEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppIpInterfaceEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpInterface")

		// remove fields from embedded structs
		reflectStorageNetAppBaseEvent := reflect.ValueOf(o.StorageNetAppBaseEvent)
		for i := 0; i < reflectStorageNetAppBaseEvent.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseEvent.Type().Field(i)

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

type NullableStorageNetAppIpInterfaceEvent struct {
	value *StorageNetAppIpInterfaceEvent
	isSet bool
}

func (v NullableStorageNetAppIpInterfaceEvent) Get() *StorageNetAppIpInterfaceEvent {
	return v.value
}

func (v *NullableStorageNetAppIpInterfaceEvent) Set(val *StorageNetAppIpInterfaceEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppIpInterfaceEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppIpInterfaceEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppIpInterfaceEvent(val *StorageNetAppIpInterfaceEvent) *NullableStorageNetAppIpInterfaceEvent {
	return &NullableStorageNetAppIpInterfaceEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppIpInterfaceEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppIpInterfaceEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
