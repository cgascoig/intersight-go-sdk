/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricVlanSettings VLAN configuration for the virtual interface.
type FabricVlanSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Allowed VLAN IDs of the virtual interface. A list of comma seperated VLAN ids and/or VLAN id ranges.
	AllowedVlans *string `json:"AllowedVlans,omitempty"`
	// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. If the native VLAN is not a part of the allowed VLANs, it will automatically be added to the list of allowed VLANs.
	NativeVlan           *int64 `json:"NativeVlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVlanSettings FabricVlanSettings

// NewFabricVlanSettings instantiates a new FabricVlanSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVlanSettings(classId string, objectType string) *FabricVlanSettings {
	this := FabricVlanSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var nativeVlan int64 = 1
	this.NativeVlan = &nativeVlan
	return &this
}

// NewFabricVlanSettingsWithDefaults instantiates a new FabricVlanSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVlanSettingsWithDefaults() *FabricVlanSettings {
	this := FabricVlanSettings{}
	var classId string = "fabric.VlanSettings"
	this.ClassId = classId
	var objectType string = "fabric.VlanSettings"
	this.ObjectType = objectType
	var nativeVlan int64 = 1
	this.NativeVlan = &nativeVlan
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVlanSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVlanSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVlanSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVlanSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVlanSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *FabricVlanSettings) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSettings) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *FabricVlanSettings) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *FabricVlanSettings) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetNativeVlan returns the NativeVlan field value if set, zero value otherwise.
func (o *FabricVlanSettings) GetNativeVlan() int64 {
	if o == nil || o.NativeVlan == nil {
		var ret int64
		return ret
	}
	return *o.NativeVlan
}

// GetNativeVlanOk returns a tuple with the NativeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVlanSettings) GetNativeVlanOk() (*int64, bool) {
	if o == nil || o.NativeVlan == nil {
		return nil, false
	}
	return o.NativeVlan, true
}

// HasNativeVlan returns a boolean if a field has been set.
func (o *FabricVlanSettings) HasNativeVlan() bool {
	if o != nil && o.NativeVlan != nil {
		return true
	}

	return false
}

// SetNativeVlan gets a reference to the given int64 and assigns it to the NativeVlan field.
func (o *FabricVlanSettings) SetNativeVlan(v int64) {
	o.NativeVlan = &v
}

func (o FabricVlanSettings) MarshalJSON() ([]byte, error) {
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
	if o.AllowedVlans != nil {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if o.NativeVlan != nil {
		toSerialize["NativeVlan"] = o.NativeVlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVlanSettings) UnmarshalJSON(bytes []byte) (err error) {
	type FabricVlanSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Allowed VLAN IDs of the virtual interface. A list of comma seperated VLAN ids and/or VLAN id ranges.
		AllowedVlans *string `json:"AllowedVlans,omitempty"`
		// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. If the native VLAN is not a part of the allowed VLANs, it will automatically be added to the list of allowed VLANs.
		NativeVlan *int64 `json:"NativeVlan,omitempty"`
	}

	varFabricVlanSettingsWithoutEmbeddedStruct := FabricVlanSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricVlanSettingsWithoutEmbeddedStruct)
	if err == nil {
		varFabricVlanSettings := _FabricVlanSettings{}
		varFabricVlanSettings.ClassId = varFabricVlanSettingsWithoutEmbeddedStruct.ClassId
		varFabricVlanSettings.ObjectType = varFabricVlanSettingsWithoutEmbeddedStruct.ObjectType
		varFabricVlanSettings.AllowedVlans = varFabricVlanSettingsWithoutEmbeddedStruct.AllowedVlans
		varFabricVlanSettings.NativeVlan = varFabricVlanSettingsWithoutEmbeddedStruct.NativeVlan
		*o = FabricVlanSettings(varFabricVlanSettings)
	} else {
		return err
	}

	varFabricVlanSettings := _FabricVlanSettings{}

	err = json.Unmarshal(bytes, &varFabricVlanSettings)
	if err == nil {
		o.MoBaseComplexType = varFabricVlanSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowedVlans")
		delete(additionalProperties, "NativeVlan")

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

type NullableFabricVlanSettings struct {
	value *FabricVlanSettings
	isSet bool
}

func (v NullableFabricVlanSettings) Get() *FabricVlanSettings {
	return v.value
}

func (v *NullableFabricVlanSettings) Set(val *FabricVlanSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVlanSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVlanSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVlanSettings(val *FabricVlanSettings) *NullableFabricVlanSettings {
	return &NullableFabricVlanSettings{value: val, isSet: true}
}

func (v NullableFabricVlanSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVlanSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
