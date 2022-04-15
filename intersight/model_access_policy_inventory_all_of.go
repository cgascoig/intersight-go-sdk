/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AccessPolicyInventoryAllOf Definition of the list of properties defined in 'access.PolicyInventory', excluding properties defined in parent classes.
type AccessPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                          `json:"ObjectType"`
	AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
	ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
	// VLAN to be used for server access over Inband network.
	InbandVlan           *int64                  `json:"InbandVlan,omitempty"`
	InbandIpPool         *IppoolPoolRelationship `json:"InbandIpPool,omitempty"`
	InbandVrf            *VrfVrfRelationship     `json:"InbandVrf,omitempty"`
	OutOfBandIpPool      *IppoolPoolRelationship `json:"OutOfBandIpPool,omitempty"`
	OutOfBandVrf         *VrfVrfRelationship     `json:"OutOfBandVrf,omitempty"`
	TargetMo             *MoBaseMoRelationship   `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyInventoryAllOf AccessPolicyInventoryAllOf

// NewAccessPolicyInventoryAllOf instantiates a new AccessPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyInventoryAllOf(classId string, objectType string) *AccessPolicyInventoryAllOf {
	this := AccessPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessPolicyInventoryAllOfWithDefaults instantiates a new AccessPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyInventoryAllOfWithDefaults() *AccessPolicyInventoryAllOf {
	this := AccessPolicyInventoryAllOf{}
	var classId string = "access.PolicyInventory"
	this.ClassId = classId
	var objectType string = "access.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AccessPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddressType returns the AddressType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventoryAllOf) GetAddressType() AccessAddressType {
	if o == nil || o.AddressType.Get() == nil {
		var ret AccessAddressType
		return ret
	}
	return *o.AddressType.Get()
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventoryAllOf) GetAddressTypeOk() (*AccessAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressType.Get(), o.AddressType.IsSet()
}

// HasAddressType returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasAddressType() bool {
	if o != nil && o.AddressType.IsSet() {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given NullableAccessAddressType and assigns it to the AddressType field.
func (o *AccessPolicyInventoryAllOf) SetAddressType(v AccessAddressType) {
	o.AddressType.Set(&v)
}

// SetAddressTypeNil sets the value for AddressType to be an explicit nil
func (o *AccessPolicyInventoryAllOf) SetAddressTypeNil() {
	o.AddressType.Set(nil)
}

// UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
func (o *AccessPolicyInventoryAllOf) UnsetAddressType() {
	o.AddressType.Unset()
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventoryAllOf) GetConfigurationType() AccessConfigurationType {
	if o == nil || o.ConfigurationType.Get() == nil {
		var ret AccessConfigurationType
		return ret
	}
	return *o.ConfigurationType.Get()
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventoryAllOf) GetConfigurationTypeOk() (*AccessConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationType.Get(), o.ConfigurationType.IsSet()
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType.IsSet() {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given NullableAccessConfigurationType and assigns it to the ConfigurationType field.
func (o *AccessPolicyInventoryAllOf) SetConfigurationType(v AccessConfigurationType) {
	o.ConfigurationType.Set(&v)
}

// SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil
func (o *AccessPolicyInventoryAllOf) SetConfigurationTypeNil() {
	o.ConfigurationType.Set(nil)
}

// UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
func (o *AccessPolicyInventoryAllOf) UnsetConfigurationType() {
	o.ConfigurationType.Unset()
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetInbandVlan() int64 {
	if o == nil || o.InbandVlan == nil {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetInbandVlanOk() (*int64, bool) {
	if o == nil || o.InbandVlan == nil {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasInbandVlan() bool {
	if o != nil && o.InbandVlan != nil {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *AccessPolicyInventoryAllOf) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetInbandIpPool returns the InbandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetInbandIpPool() IppoolPoolRelationship {
	if o == nil || o.InbandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InbandIpPool
}

// GetInbandIpPoolOk returns a tuple with the InbandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.InbandIpPool == nil {
		return nil, false
	}
	return o.InbandIpPool, true
}

// HasInbandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasInbandIpPool() bool {
	if o != nil && o.InbandIpPool != nil {
		return true
	}

	return false
}

// SetInbandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the InbandIpPool field.
func (o *AccessPolicyInventoryAllOf) SetInbandIpPool(v IppoolPoolRelationship) {
	o.InbandIpPool = &v
}

// GetInbandVrf returns the InbandVrf field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetInbandVrf() VrfVrfRelationship {
	if o == nil || o.InbandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.InbandVrf
}

// GetInbandVrfOk returns a tuple with the InbandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetInbandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.InbandVrf == nil {
		return nil, false
	}
	return o.InbandVrf, true
}

// HasInbandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasInbandVrf() bool {
	if o != nil && o.InbandVrf != nil {
		return true
	}

	return false
}

// SetInbandVrf gets a reference to the given VrfVrfRelationship and assigns it to the InbandVrf field.
func (o *AccessPolicyInventoryAllOf) SetInbandVrf(v VrfVrfRelationship) {
	o.InbandVrf = &v
}

// GetOutOfBandIpPool returns the OutOfBandIpPool field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetOutOfBandIpPool() IppoolPoolRelationship {
	if o == nil || o.OutOfBandIpPool == nil {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.OutOfBandIpPool
}

// GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil || o.OutOfBandIpPool == nil {
		return nil, false
	}
	return o.OutOfBandIpPool, true
}

// HasOutOfBandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasOutOfBandIpPool() bool {
	if o != nil && o.OutOfBandIpPool != nil {
		return true
	}

	return false
}

// SetOutOfBandIpPool gets a reference to the given IppoolPoolRelationship and assigns it to the OutOfBandIpPool field.
func (o *AccessPolicyInventoryAllOf) SetOutOfBandIpPool(v IppoolPoolRelationship) {
	o.OutOfBandIpPool = &v
}

// GetOutOfBandVrf returns the OutOfBandVrf field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetOutOfBandVrf() VrfVrfRelationship {
	if o == nil || o.OutOfBandVrf == nil {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.OutOfBandVrf
}

// GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil || o.OutOfBandVrf == nil {
		return nil, false
	}
	return o.OutOfBandVrf, true
}

// HasOutOfBandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasOutOfBandVrf() bool {
	if o != nil && o.OutOfBandVrf != nil {
		return true
	}

	return false
}

// SetOutOfBandVrf gets a reference to the given VrfVrfRelationship and assigns it to the OutOfBandVrf field.
func (o *AccessPolicyInventoryAllOf) SetOutOfBandVrf(v VrfVrfRelationship) {
	o.OutOfBandVrf = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *AccessPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *AccessPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *AccessPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o AccessPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AddressType.IsSet() {
		toSerialize["AddressType"] = o.AddressType.Get()
	}
	if o.ConfigurationType.IsSet() {
		toSerialize["ConfigurationType"] = o.ConfigurationType.Get()
	}
	if o.InbandVlan != nil {
		toSerialize["InbandVlan"] = o.InbandVlan
	}
	if o.InbandIpPool != nil {
		toSerialize["InbandIpPool"] = o.InbandIpPool
	}
	if o.InbandVrf != nil {
		toSerialize["InbandVrf"] = o.InbandVrf
	}
	if o.OutOfBandIpPool != nil {
		toSerialize["OutOfBandIpPool"] = o.OutOfBandIpPool
	}
	if o.OutOfBandVrf != nil {
		toSerialize["OutOfBandVrf"] = o.OutOfBandVrf
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyInventoryAllOf := _AccessPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varAccessPolicyInventoryAllOf); err == nil {
		*o = AccessPolicyInventoryAllOf(varAccessPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddressType")
		delete(additionalProperties, "ConfigurationType")
		delete(additionalProperties, "InbandVlan")
		delete(additionalProperties, "InbandIpPool")
		delete(additionalProperties, "InbandVrf")
		delete(additionalProperties, "OutOfBandIpPool")
		delete(additionalProperties, "OutOfBandVrf")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyInventoryAllOf struct {
	value *AccessPolicyInventoryAllOf
	isSet bool
}

func (v NullableAccessPolicyInventoryAllOf) Get() *AccessPolicyInventoryAllOf {
	return v.value
}

func (v *NullableAccessPolicyInventoryAllOf) Set(val *AccessPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyInventoryAllOf(val *AccessPolicyInventoryAllOf) *NullableAccessPolicyInventoryAllOf {
	return &NullableAccessPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableAccessPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
