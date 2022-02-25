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
)

// KubernetesVirtualMachineInfrastructureProviderAllOf Definition of the list of properties defined in 'kubernetes.VirtualMachineInfrastructureProvider', excluding properties defined in parent classes.
type KubernetesVirtualMachineInfrastructureProviderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                                 `json:"ObjectType"`
	InfraConfig          NullableKubernetesBaseVirtualMachineInfraConfig        `json:"InfraConfig,omitempty"`
	InfraConfigPolicy    *KubernetesVirtualMachineInfraConfigPolicyRelationship `json:"InfraConfigPolicy,omitempty"`
	InstanceType         *KubernetesVirtualMachineInstanceTypeRelationship      `json:"InstanceType,omitempty"`
	Target               *AssetDeviceRegistrationRelationship                   `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineInfrastructureProviderAllOf KubernetesVirtualMachineInfrastructureProviderAllOf

// NewKubernetesVirtualMachineInfrastructureProviderAllOf instantiates a new KubernetesVirtualMachineInfrastructureProviderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineInfrastructureProviderAllOf(classId string, objectType string) *KubernetesVirtualMachineInfrastructureProviderAllOf {
	this := KubernetesVirtualMachineInfrastructureProviderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVirtualMachineInfrastructureProviderAllOfWithDefaults instantiates a new KubernetesVirtualMachineInfrastructureProviderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineInfrastructureProviderAllOfWithDefaults() *KubernetesVirtualMachineInfrastructureProviderAllOf {
	this := KubernetesVirtualMachineInfrastructureProviderAllOf{}
	var classId string = "kubernetes.VirtualMachineInfrastructureProvider"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInfrastructureProvider"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInfraConfig returns the InfraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfig() KubernetesBaseVirtualMachineInfraConfig {
	if o == nil || o.InfraConfig.Get() == nil {
		var ret KubernetesBaseVirtualMachineInfraConfig
		return ret
	}
	return *o.InfraConfig.Get()
}

// GetInfraConfigOk returns a tuple with the InfraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigOk() (*KubernetesBaseVirtualMachineInfraConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.InfraConfig.Get(), o.InfraConfig.IsSet()
}

// HasInfraConfig returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInfraConfig() bool {
	if o != nil && o.InfraConfig.IsSet() {
		return true
	}

	return false
}

// SetInfraConfig gets a reference to the given NullableKubernetesBaseVirtualMachineInfraConfig and assigns it to the InfraConfig field.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfig(v KubernetesBaseVirtualMachineInfraConfig) {
	o.InfraConfig.Set(&v)
}

// SetInfraConfigNil sets the value for InfraConfig to be an explicit nil
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfigNil() {
	o.InfraConfig.Set(nil)
}

// UnsetInfraConfig ensures that no value is present for InfraConfig, not even an explicit nil
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) UnsetInfraConfig() {
	o.InfraConfig.Unset()
}

// GetInfraConfigPolicy returns the InfraConfigPolicy field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigPolicy() KubernetesVirtualMachineInfraConfigPolicyRelationship {
	if o == nil || o.InfraConfigPolicy == nil {
		var ret KubernetesVirtualMachineInfraConfigPolicyRelationship
		return ret
	}
	return *o.InfraConfigPolicy
}

// GetInfraConfigPolicyOk returns a tuple with the InfraConfigPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInfraConfigPolicyOk() (*KubernetesVirtualMachineInfraConfigPolicyRelationship, bool) {
	if o == nil || o.InfraConfigPolicy == nil {
		return nil, false
	}
	return o.InfraConfigPolicy, true
}

// HasInfraConfigPolicy returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInfraConfigPolicy() bool {
	if o != nil && o.InfraConfigPolicy != nil {
		return true
	}

	return false
}

// SetInfraConfigPolicy gets a reference to the given KubernetesVirtualMachineInfraConfigPolicyRelationship and assigns it to the InfraConfigPolicy field.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInfraConfigPolicy(v KubernetesVirtualMachineInfraConfigPolicyRelationship) {
	o.InfraConfigPolicy = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInstanceType() KubernetesVirtualMachineInstanceTypeRelationship {
	if o == nil || o.InstanceType == nil {
		var ret KubernetesVirtualMachineInstanceTypeRelationship
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetInstanceTypeOk() (*KubernetesVirtualMachineInstanceTypeRelationship, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given KubernetesVirtualMachineInstanceTypeRelationship and assigns it to the InstanceType field.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetInstanceType(v KubernetesVirtualMachineInstanceTypeRelationship) {
	o.InstanceType = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || o.Target == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target = &v
}

func (o KubernetesVirtualMachineInfrastructureProviderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InfraConfig.IsSet() {
		toSerialize["InfraConfig"] = o.InfraConfig.Get()
	}
	if o.InfraConfigPolicy != nil {
		toSerialize["InfraConfigPolicy"] = o.InfraConfigPolicy
	}
	if o.InstanceType != nil {
		toSerialize["InstanceType"] = o.InstanceType
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineInfrastructureProviderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesVirtualMachineInfrastructureProviderAllOf := _KubernetesVirtualMachineInfrastructureProviderAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInfrastructureProviderAllOf); err == nil {
		*o = KubernetesVirtualMachineInfrastructureProviderAllOf(varKubernetesVirtualMachineInfrastructureProviderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InfraConfig")
		delete(additionalProperties, "InfraConfigPolicy")
		delete(additionalProperties, "InstanceType")
		delete(additionalProperties, "Target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesVirtualMachineInfrastructureProviderAllOf struct {
	value *KubernetesVirtualMachineInfrastructureProviderAllOf
	isSet bool
}

func (v NullableKubernetesVirtualMachineInfrastructureProviderAllOf) Get() *KubernetesVirtualMachineInfrastructureProviderAllOf {
	return v.value
}

func (v *NullableKubernetesVirtualMachineInfrastructureProviderAllOf) Set(val *KubernetesVirtualMachineInfrastructureProviderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineInfrastructureProviderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineInfrastructureProviderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineInfrastructureProviderAllOf(val *KubernetesVirtualMachineInfrastructureProviderAllOf) *NullableKubernetesVirtualMachineInfrastructureProviderAllOf {
	return &NullableKubernetesVirtualMachineInfrastructureProviderAllOf{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineInfrastructureProviderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineInfrastructureProviderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
