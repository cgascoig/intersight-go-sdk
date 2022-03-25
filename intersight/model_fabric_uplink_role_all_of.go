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
)

// FabricUplinkRoleAllOf Definition of the list of properties defined in 'fabric.UplinkRole', excluding properties defined in parent classes.
type FabricUplinkRoleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to fabricEthNetworkGroupPolicy resources.
	EthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship `json:"EthNetworkGroupPolicy,omitempty"`
	FlowControlPolicy     *FabricFlowControlPolicyRelationship      `json:"FlowControlPolicy,omitempty"`
	LinkControlPolicy     *FabricLinkControlPolicyRelationship      `json:"LinkControlPolicy,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _FabricUplinkRoleAllOf FabricUplinkRoleAllOf

// NewFabricUplinkRoleAllOf instantiates a new FabricUplinkRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUplinkRoleAllOf(classId string, objectType string) *FabricUplinkRoleAllOf {
	this := FabricUplinkRoleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricUplinkRoleAllOfWithDefaults instantiates a new FabricUplinkRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUplinkRoleAllOfWithDefaults() *FabricUplinkRoleAllOf {
	this := FabricUplinkRoleAllOf{}
	var classId string = "fabric.UplinkRole"
	this.ClassId = classId
	var objectType string = "fabric.UplinkRole"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUplinkRoleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRoleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUplinkRoleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricUplinkRoleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkRoleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUplinkRoleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkRoleAllOf) GetEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship {
	if o == nil {
		var ret []FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return o.EthNetworkGroupPolicy
}

// GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkRoleAllOf) GetEthNetworkGroupPolicyOk() (*[]FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || o.EthNetworkGroupPolicy == nil {
		return nil, false
	}
	return &o.EthNetworkGroupPolicy, true
}

// HasEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRoleAllOf) HasEthNetworkGroupPolicy() bool {
	if o != nil && o.EthNetworkGroupPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkGroupPolicy gets a reference to the given []FabricEthNetworkGroupPolicyRelationship and assigns it to the EthNetworkGroupPolicy field.
func (o *FabricUplinkRoleAllOf) SetEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship) {
	o.EthNetworkGroupPolicy = v
}

// GetFlowControlPolicy returns the FlowControlPolicy field value if set, zero value otherwise.
func (o *FabricUplinkRoleAllOf) GetFlowControlPolicy() FabricFlowControlPolicyRelationship {
	if o == nil || o.FlowControlPolicy == nil {
		var ret FabricFlowControlPolicyRelationship
		return ret
	}
	return *o.FlowControlPolicy
}

// GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkRoleAllOf) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool) {
	if o == nil || o.FlowControlPolicy == nil {
		return nil, false
	}
	return o.FlowControlPolicy, true
}

// HasFlowControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRoleAllOf) HasFlowControlPolicy() bool {
	if o != nil && o.FlowControlPolicy != nil {
		return true
	}

	return false
}

// SetFlowControlPolicy gets a reference to the given FabricFlowControlPolicyRelationship and assigns it to the FlowControlPolicy field.
func (o *FabricUplinkRoleAllOf) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship) {
	o.FlowControlPolicy = &v
}

// GetLinkControlPolicy returns the LinkControlPolicy field value if set, zero value otherwise.
func (o *FabricUplinkRoleAllOf) GetLinkControlPolicy() FabricLinkControlPolicyRelationship {
	if o == nil || o.LinkControlPolicy == nil {
		var ret FabricLinkControlPolicyRelationship
		return ret
	}
	return *o.LinkControlPolicy
}

// GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkRoleAllOf) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool) {
	if o == nil || o.LinkControlPolicy == nil {
		return nil, false
	}
	return o.LinkControlPolicy, true
}

// HasLinkControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkRoleAllOf) HasLinkControlPolicy() bool {
	if o != nil && o.LinkControlPolicy != nil {
		return true
	}

	return false
}

// SetLinkControlPolicy gets a reference to the given FabricLinkControlPolicyRelationship and assigns it to the LinkControlPolicy field.
func (o *FabricUplinkRoleAllOf) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship) {
	o.LinkControlPolicy = &v
}

func (o FabricUplinkRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EthNetworkGroupPolicy != nil {
		toSerialize["EthNetworkGroupPolicy"] = o.EthNetworkGroupPolicy
	}
	if o.FlowControlPolicy != nil {
		toSerialize["FlowControlPolicy"] = o.FlowControlPolicy
	}
	if o.LinkControlPolicy != nil {
		toSerialize["LinkControlPolicy"] = o.LinkControlPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricUplinkRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricUplinkRoleAllOf := _FabricUplinkRoleAllOf{}

	if err = json.Unmarshal(bytes, &varFabricUplinkRoleAllOf); err == nil {
		*o = FabricUplinkRoleAllOf(varFabricUplinkRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EthNetworkGroupPolicy")
		delete(additionalProperties, "FlowControlPolicy")
		delete(additionalProperties, "LinkControlPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricUplinkRoleAllOf struct {
	value *FabricUplinkRoleAllOf
	isSet bool
}

func (v NullableFabricUplinkRoleAllOf) Get() *FabricUplinkRoleAllOf {
	return v.value
}

func (v *NullableFabricUplinkRoleAllOf) Set(val *FabricUplinkRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUplinkRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUplinkRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUplinkRoleAllOf(val *FabricUplinkRoleAllOf) *NullableFabricUplinkRoleAllOf {
	return &NullableFabricUplinkRoleAllOf{value: val, isSet: true}
}

func (v NullableFabricUplinkRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUplinkRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
