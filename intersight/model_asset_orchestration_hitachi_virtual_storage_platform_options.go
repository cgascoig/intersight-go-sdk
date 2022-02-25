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

// AssetOrchestrationHitachiVirtualStoragePlatformOptions Captures configuration specific to the Hitachi Virtual Storage Platform target for the Orchestration service.
type AssetOrchestrationHitachiVirtualStoragePlatformOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The DNS hostname or IP address of the Hitachi Ops Center API Configuration Manager used to manage the Hitachi Virtual Storage Platform.
	OpsCenterManagementAddress *string `json:"OpsCenterManagementAddress,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _AssetOrchestrationHitachiVirtualStoragePlatformOptions AssetOrchestrationHitachiVirtualStoragePlatformOptions

// NewAssetOrchestrationHitachiVirtualStoragePlatformOptions instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationHitachiVirtualStoragePlatformOptions(classId string, objectType string) *AssetOrchestrationHitachiVirtualStoragePlatformOptions {
	this := AssetOrchestrationHitachiVirtualStoragePlatformOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithDefaults instantiates a new AssetOrchestrationHitachiVirtualStoragePlatformOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithDefaults() *AssetOrchestrationHitachiVirtualStoragePlatformOptions {
	this := AssetOrchestrationHitachiVirtualStoragePlatformOptions{}
	var classId string = "asset.OrchestrationHitachiVirtualStoragePlatformOptions"
	this.ClassId = classId
	var objectType string = "asset.OrchestrationHitachiVirtualStoragePlatformOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOpsCenterManagementAddress returns the OpsCenterManagementAddress field value if set, zero value otherwise.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetOpsCenterManagementAddress() string {
	if o == nil || o.OpsCenterManagementAddress == nil {
		var ret string
		return ret
	}
	return *o.OpsCenterManagementAddress
}

// GetOpsCenterManagementAddressOk returns a tuple with the OpsCenterManagementAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) GetOpsCenterManagementAddressOk() (*string, bool) {
	if o == nil || o.OpsCenterManagementAddress == nil {
		return nil, false
	}
	return o.OpsCenterManagementAddress, true
}

// HasOpsCenterManagementAddress returns a boolean if a field has been set.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) HasOpsCenterManagementAddress() bool {
	if o != nil && o.OpsCenterManagementAddress != nil {
		return true
	}

	return false
}

// SetOpsCenterManagementAddress gets a reference to the given string and assigns it to the OpsCenterManagementAddress field.
func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) SetOpsCenterManagementAddress(v string) {
	o.OpsCenterManagementAddress = &v
}

func (o AssetOrchestrationHitachiVirtualStoragePlatformOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OpsCenterManagementAddress != nil {
		toSerialize["OpsCenterManagementAddress"] = o.OpsCenterManagementAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOrchestrationHitachiVirtualStoragePlatformOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The DNS hostname or IP address of the Hitachi Ops Center API Configuration Manager used to manage the Hitachi Virtual Storage Platform.
		OpsCenterManagementAddress *string `json:"OpsCenterManagementAddress,omitempty"`
	}

	varAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct := AssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetOrchestrationHitachiVirtualStoragePlatformOptions := _AssetOrchestrationHitachiVirtualStoragePlatformOptions{}
		varAssetOrchestrationHitachiVirtualStoragePlatformOptions.ClassId = varAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct.ClassId
		varAssetOrchestrationHitachiVirtualStoragePlatformOptions.ObjectType = varAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct.ObjectType
		varAssetOrchestrationHitachiVirtualStoragePlatformOptions.OpsCenterManagementAddress = varAssetOrchestrationHitachiVirtualStoragePlatformOptionsWithoutEmbeddedStruct.OpsCenterManagementAddress
		*o = AssetOrchestrationHitachiVirtualStoragePlatformOptions(varAssetOrchestrationHitachiVirtualStoragePlatformOptions)
	} else {
		return err
	}

	varAssetOrchestrationHitachiVirtualStoragePlatformOptions := _AssetOrchestrationHitachiVirtualStoragePlatformOptions{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationHitachiVirtualStoragePlatformOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetOrchestrationHitachiVirtualStoragePlatformOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OpsCenterManagementAddress")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions struct {
	value *AssetOrchestrationHitachiVirtualStoragePlatformOptions
	isSet bool
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) Get() *AssetOrchestrationHitachiVirtualStoragePlatformOptions {
	return v.value
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) Set(val *AssetOrchestrationHitachiVirtualStoragePlatformOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationHitachiVirtualStoragePlatformOptions(val *AssetOrchestrationHitachiVirtualStoragePlatformOptions) *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions {
	return &NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions{value: val, isSet: true}
}

func (v NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationHitachiVirtualStoragePlatformOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
