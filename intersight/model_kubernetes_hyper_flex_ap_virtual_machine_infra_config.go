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

// KubernetesHyperFlexApVirtualMachineInfraConfig Infrastructure provider allocation configuration for HyperFlex Application platform virtual machine Kubernetes nodes.
type KubernetesHyperFlexApVirtualMachineInfraConfig struct {
	KubernetesBaseVirtualMachineInfraConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Disk mode to use for volumes. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	DiskMode             *string `json:"DiskMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesHyperFlexApVirtualMachineInfraConfig KubernetesHyperFlexApVirtualMachineInfraConfig

// NewKubernetesHyperFlexApVirtualMachineInfraConfig instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesHyperFlexApVirtualMachineInfraConfig(classId string, objectType string) *KubernetesHyperFlexApVirtualMachineInfraConfig {
	this := KubernetesHyperFlexApVirtualMachineInfraConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var diskMode string = "Block"
	this.DiskMode = &diskMode
	return &this
}

// NewKubernetesHyperFlexApVirtualMachineInfraConfigWithDefaults instantiates a new KubernetesHyperFlexApVirtualMachineInfraConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesHyperFlexApVirtualMachineInfraConfigWithDefaults() *KubernetesHyperFlexApVirtualMachineInfraConfig {
	this := KubernetesHyperFlexApVirtualMachineInfraConfig{}
	var classId string = "kubernetes.HyperFlexApVirtualMachineInfraConfig"
	this.ClassId = classId
	var objectType string = "kubernetes.HyperFlexApVirtualMachineInfraConfig"
	this.ObjectType = objectType
	var diskMode string = "Block"
	this.DiskMode = &diskMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskMode returns the DiskMode field value if set, zero value otherwise.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetDiskMode() string {
	if o == nil || o.DiskMode == nil {
		var ret string
		return ret
	}
	return *o.DiskMode
}

// GetDiskModeOk returns a tuple with the DiskMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) GetDiskModeOk() (*string, bool) {
	if o == nil || o.DiskMode == nil {
		return nil, false
	}
	return o.DiskMode, true
}

// HasDiskMode returns a boolean if a field has been set.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) HasDiskMode() bool {
	if o != nil && o.DiskMode != nil {
		return true
	}

	return false
}

// SetDiskMode gets a reference to the given string and assigns it to the DiskMode field.
func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) SetDiskMode(v string) {
	o.DiskMode = &v
}

func (o KubernetesHyperFlexApVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesBaseVirtualMachineInfraConfig, errKubernetesBaseVirtualMachineInfraConfig := json.Marshal(o.KubernetesBaseVirtualMachineInfraConfig)
	if errKubernetesBaseVirtualMachineInfraConfig != nil {
		return []byte{}, errKubernetesBaseVirtualMachineInfraConfig
	}
	errKubernetesBaseVirtualMachineInfraConfig = json.Unmarshal([]byte(serializedKubernetesBaseVirtualMachineInfraConfig), &toSerialize)
	if errKubernetesBaseVirtualMachineInfraConfig != nil {
		return []byte{}, errKubernetesBaseVirtualMachineInfraConfig
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiskMode != nil {
		toSerialize["DiskMode"] = o.DiskMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesHyperFlexApVirtualMachineInfraConfig) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Disk mode to use for volumes. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
		DiskMode *string `json:"DiskMode,omitempty"`
	}

	varKubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct := KubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesHyperFlexApVirtualMachineInfraConfig := _KubernetesHyperFlexApVirtualMachineInfraConfig{}
		varKubernetesHyperFlexApVirtualMachineInfraConfig.ClassId = varKubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct.ClassId
		varKubernetesHyperFlexApVirtualMachineInfraConfig.ObjectType = varKubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct.ObjectType
		varKubernetesHyperFlexApVirtualMachineInfraConfig.DiskMode = varKubernetesHyperFlexApVirtualMachineInfraConfigWithoutEmbeddedStruct.DiskMode
		*o = KubernetesHyperFlexApVirtualMachineInfraConfig(varKubernetesHyperFlexApVirtualMachineInfraConfig)
	} else {
		return err
	}

	varKubernetesHyperFlexApVirtualMachineInfraConfig := _KubernetesHyperFlexApVirtualMachineInfraConfig{}

	err = json.Unmarshal(bytes, &varKubernetesHyperFlexApVirtualMachineInfraConfig)
	if err == nil {
		o.KubernetesBaseVirtualMachineInfraConfig = varKubernetesHyperFlexApVirtualMachineInfraConfig.KubernetesBaseVirtualMachineInfraConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskMode")

		// remove fields from embedded structs
		reflectKubernetesBaseVirtualMachineInfraConfig := reflect.ValueOf(o.KubernetesBaseVirtualMachineInfraConfig)
		for i := 0; i < reflectKubernetesBaseVirtualMachineInfraConfig.Type().NumField(); i++ {
			t := reflectKubernetesBaseVirtualMachineInfraConfig.Type().Field(i)

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

type NullableKubernetesHyperFlexApVirtualMachineInfraConfig struct {
	value *KubernetesHyperFlexApVirtualMachineInfraConfig
	isSet bool
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfig) Get() *KubernetesHyperFlexApVirtualMachineInfraConfig {
	return v.value
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfig) Set(val *KubernetesHyperFlexApVirtualMachineInfraConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesHyperFlexApVirtualMachineInfraConfig(val *KubernetesHyperFlexApVirtualMachineInfraConfig) *NullableKubernetesHyperFlexApVirtualMachineInfraConfig {
	return &NullableKubernetesHyperFlexApVirtualMachineInfraConfig{value: val, isSet: true}
}

func (v NullableKubernetesHyperFlexApVirtualMachineInfraConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesHyperFlexApVirtualMachineInfraConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
