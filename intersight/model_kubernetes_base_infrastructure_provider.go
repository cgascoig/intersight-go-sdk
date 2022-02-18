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

// KubernetesBaseInfrastructureProvider Infrastructure backend to use for providing compute and storage resources.
type KubernetesBaseInfrastructureProvider struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Description for the infrastructure provider.
	Description *string `json:"Description,omitempty"`
	// Name of an infrastructure provider.
	Name                 *string                                 `json:"Name,omitempty"`
	NodeGroup            *KubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesBaseInfrastructureProvider KubernetesBaseInfrastructureProvider

// NewKubernetesBaseInfrastructureProvider instantiates a new KubernetesBaseInfrastructureProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesBaseInfrastructureProvider(classId string, objectType string) *KubernetesBaseInfrastructureProvider {
	this := KubernetesBaseInfrastructureProvider{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesBaseInfrastructureProviderWithDefaults instantiates a new KubernetesBaseInfrastructureProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesBaseInfrastructureProviderWithDefaults() *KubernetesBaseInfrastructureProvider {
	this := KubernetesBaseInfrastructureProvider{}
	var classId string = "kubernetes.VirtualMachineInfrastructureProvider"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInfrastructureProvider"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesBaseInfrastructureProvider) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseInfrastructureProvider) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesBaseInfrastructureProvider) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesBaseInfrastructureProvider) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesBaseInfrastructureProvider) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesBaseInfrastructureProvider) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KubernetesBaseInfrastructureProvider) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseInfrastructureProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KubernetesBaseInfrastructureProvider) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KubernetesBaseInfrastructureProvider) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesBaseInfrastructureProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseInfrastructureProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesBaseInfrastructureProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesBaseInfrastructureProvider) SetName(v string) {
	o.Name = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise.
func (o *KubernetesBaseInfrastructureProvider) GetNodeGroup() KubernetesNodeGroupProfileRelationship {
	if o == nil || o.NodeGroup == nil {
		var ret KubernetesNodeGroupProfileRelationship
		return ret
	}
	return *o.NodeGroup
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesBaseInfrastructureProvider) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.NodeGroup == nil {
		return nil, false
	}
	return o.NodeGroup, true
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *KubernetesBaseInfrastructureProvider) HasNodeGroup() bool {
	if o != nil && o.NodeGroup != nil {
		return true
	}

	return false
}

// SetNodeGroup gets a reference to the given KubernetesNodeGroupProfileRelationship and assigns it to the NodeGroup field.
func (o *KubernetesBaseInfrastructureProvider) SetNodeGroup(v KubernetesNodeGroupProfileRelationship) {
	o.NodeGroup = &v
}

func (o KubernetesBaseInfrastructureProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NodeGroup != nil {
		toSerialize["NodeGroup"] = o.NodeGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesBaseInfrastructureProvider) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesBaseInfrastructureProviderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Description for the infrastructure provider.
		Description *string `json:"Description,omitempty"`
		// Name of an infrastructure provider.
		Name      *string                                 `json:"Name,omitempty"`
		NodeGroup *KubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
	}

	varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct := KubernetesBaseInfrastructureProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesBaseInfrastructureProvider := _KubernetesBaseInfrastructureProvider{}
		varKubernetesBaseInfrastructureProvider.ClassId = varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct.ClassId
		varKubernetesBaseInfrastructureProvider.ObjectType = varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct.ObjectType
		varKubernetesBaseInfrastructureProvider.Description = varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct.Description
		varKubernetesBaseInfrastructureProvider.Name = varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct.Name
		varKubernetesBaseInfrastructureProvider.NodeGroup = varKubernetesBaseInfrastructureProviderWithoutEmbeddedStruct.NodeGroup
		*o = KubernetesBaseInfrastructureProvider(varKubernetesBaseInfrastructureProvider)
	} else {
		return err
	}

	varKubernetesBaseInfrastructureProvider := _KubernetesBaseInfrastructureProvider{}

	err = json.Unmarshal(bytes, &varKubernetesBaseInfrastructureProvider)
	if err == nil {
		o.MoBaseMo = varKubernetesBaseInfrastructureProvider.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NodeGroup")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableKubernetesBaseInfrastructureProvider struct {
	value *KubernetesBaseInfrastructureProvider
	isSet bool
}

func (v NullableKubernetesBaseInfrastructureProvider) Get() *KubernetesBaseInfrastructureProvider {
	return v.value
}

func (v *NullableKubernetesBaseInfrastructureProvider) Set(val *KubernetesBaseInfrastructureProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesBaseInfrastructureProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesBaseInfrastructureProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesBaseInfrastructureProvider(val *KubernetesBaseInfrastructureProvider) *NullableKubernetesBaseInfrastructureProvider {
	return &NullableKubernetesBaseInfrastructureProvider{value: val, isSet: true}
}

func (v NullableKubernetesBaseInfrastructureProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesBaseInfrastructureProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
