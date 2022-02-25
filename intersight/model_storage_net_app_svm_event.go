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

// StorageNetAppSvmEvent An event where the impacted resource type is a storage vm.
type StorageNetAppSvmEvent struct {
	StorageNetAppBaseEvent
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                              `json:"ObjectType"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppSvmEvent StorageNetAppSvmEvent

// NewStorageNetAppSvmEvent instantiates a new StorageNetAppSvmEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppSvmEvent(classId string, objectType string) *StorageNetAppSvmEvent {
	this := StorageNetAppSvmEvent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppSvmEventWithDefaults instantiates a new StorageNetAppSvmEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppSvmEventWithDefaults() *StorageNetAppSvmEvent {
	this := StorageNetAppSvmEvent{}
	var classId string = "storage.NetAppSvmEvent"
	this.ClassId = classId
	var objectType string = "storage.NetAppSvmEvent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppSvmEvent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEvent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppSvmEvent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppSvmEvent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEvent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppSvmEvent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppSvmEvent) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppSvmEvent) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppSvmEvent) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppSvmEvent) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppSvmEvent) MarshalJSON() ([]byte, error) {
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
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppSvmEvent) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppSvmEventWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                              `json:"ObjectType"`
		Tenant     *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	}

	varStorageNetAppSvmEventWithoutEmbeddedStruct := StorageNetAppSvmEventWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppSvmEventWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppSvmEvent := _StorageNetAppSvmEvent{}
		varStorageNetAppSvmEvent.ClassId = varStorageNetAppSvmEventWithoutEmbeddedStruct.ClassId
		varStorageNetAppSvmEvent.ObjectType = varStorageNetAppSvmEventWithoutEmbeddedStruct.ObjectType
		varStorageNetAppSvmEvent.Tenant = varStorageNetAppSvmEventWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppSvmEvent(varStorageNetAppSvmEvent)
	} else {
		return err
	}

	varStorageNetAppSvmEvent := _StorageNetAppSvmEvent{}

	err = json.Unmarshal(bytes, &varStorageNetAppSvmEvent)
	if err == nil {
		o.StorageNetAppBaseEvent = varStorageNetAppSvmEvent.StorageNetAppBaseEvent
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Tenant")

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

type NullableStorageNetAppSvmEvent struct {
	value *StorageNetAppSvmEvent
	isSet bool
}

func (v NullableStorageNetAppSvmEvent) Get() *StorageNetAppSvmEvent {
	return v.value
}

func (v *NullableStorageNetAppSvmEvent) Set(val *StorageNetAppSvmEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppSvmEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppSvmEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppSvmEvent(val *StorageNetAppSvmEvent) *NullableStorageNetAppSvmEvent {
	return &NullableStorageNetAppSvmEvent{value: val, isSet: true}
}

func (v NullableStorageNetAppSvmEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppSvmEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
