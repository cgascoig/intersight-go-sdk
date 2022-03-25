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

// StorageNetAppEthernetPortVlan A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.
type StorageNetAppEthernetPortVlan struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                    `json:"ObjectType"`
	BasePort   NullableStorageNetAppPort `json:"BasePort,omitempty"`
	// The ID tag of the VLAN for this port.
	Tag                  *int64 `json:"Tag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPortVlan StorageNetAppEthernetPortVlan

// NewStorageNetAppEthernetPortVlan instantiates a new StorageNetAppEthernetPortVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPortVlan(classId string, objectType string) *StorageNetAppEthernetPortVlan {
	this := StorageNetAppEthernetPortVlan{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortVlanWithDefaults instantiates a new StorageNetAppEthernetPortVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortVlanWithDefaults() *StorageNetAppEthernetPortVlan {
	this := StorageNetAppEthernetPortVlan{}
	var classId string = "storage.NetAppEthernetPortVlan"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPortVlan"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPortVlan) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortVlan) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPortVlan) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPortVlan) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortVlan) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPortVlan) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBasePort returns the BasePort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortVlan) GetBasePort() StorageNetAppPort {
	if o == nil || o.BasePort.Get() == nil {
		var ret StorageNetAppPort
		return ret
	}
	return *o.BasePort.Get()
}

// GetBasePortOk returns a tuple with the BasePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortVlan) GetBasePortOk() (*StorageNetAppPort, bool) {
	if o == nil {
		return nil, false
	}
	return o.BasePort.Get(), o.BasePort.IsSet()
}

// HasBasePort returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortVlan) HasBasePort() bool {
	if o != nil && o.BasePort.IsSet() {
		return true
	}

	return false
}

// SetBasePort gets a reference to the given NullableStorageNetAppPort and assigns it to the BasePort field.
func (o *StorageNetAppEthernetPortVlan) SetBasePort(v StorageNetAppPort) {
	o.BasePort.Set(&v)
}

// SetBasePortNil sets the value for BasePort to be an explicit nil
func (o *StorageNetAppEthernetPortVlan) SetBasePortNil() {
	o.BasePort.Set(nil)
}

// UnsetBasePort ensures that no value is present for BasePort, not even an explicit nil
func (o *StorageNetAppEthernetPortVlan) UnsetBasePort() {
	o.BasePort.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortVlan) GetTag() int64 {
	if o == nil || o.Tag == nil {
		var ret int64
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortVlan) GetTagOk() (*int64, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortVlan) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given int64 and assigns it to the Tag field.
func (o *StorageNetAppEthernetPortVlan) SetTag(v int64) {
	o.Tag = &v
}

func (o StorageNetAppEthernetPortVlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BasePort.IsSet() {
		toSerialize["BasePort"] = o.BasePort.Get()
	}
	if o.Tag != nil {
		toSerialize["Tag"] = o.Tag
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppEthernetPortVlan) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppEthernetPortVlanWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                    `json:"ObjectType"`
		BasePort   NullableStorageNetAppPort `json:"BasePort,omitempty"`
		// The ID tag of the VLAN for this port.
		Tag *int64 `json:"Tag,omitempty"`
	}

	varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct := StorageNetAppEthernetPortVlanWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppEthernetPortVlan := _StorageNetAppEthernetPortVlan{}
		varStorageNetAppEthernetPortVlan.ClassId = varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct.ClassId
		varStorageNetAppEthernetPortVlan.ObjectType = varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct.ObjectType
		varStorageNetAppEthernetPortVlan.BasePort = varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct.BasePort
		varStorageNetAppEthernetPortVlan.Tag = varStorageNetAppEthernetPortVlanWithoutEmbeddedStruct.Tag
		*o = StorageNetAppEthernetPortVlan(varStorageNetAppEthernetPortVlan)
	} else {
		return err
	}

	varStorageNetAppEthernetPortVlan := _StorageNetAppEthernetPortVlan{}

	err = json.Unmarshal(bytes, &varStorageNetAppEthernetPortVlan)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppEthernetPortVlan.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BasePort")
		delete(additionalProperties, "Tag")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageNetAppEthernetPortVlan struct {
	value *StorageNetAppEthernetPortVlan
	isSet bool
}

func (v NullableStorageNetAppEthernetPortVlan) Get() *StorageNetAppEthernetPortVlan {
	return v.value
}

func (v *NullableStorageNetAppEthernetPortVlan) Set(val *StorageNetAppEthernetPortVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPortVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPortVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPortVlan(val *StorageNetAppEthernetPortVlan) *NullableStorageNetAppEthernetPortVlan {
	return &NullableStorageNetAppEthernetPortVlan{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPortVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPortVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
