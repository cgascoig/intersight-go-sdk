/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricPortChannelRole Configuration object sent by user to apply on a port-channel.
type FabricPortChannelRole struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Unique Identifier of the port-channel, local to this switch.
	PcId                 *int64                        `json:"PcId,omitempty"`
	Ports                []FabricPortIdentifier        `json:"Ports,omitempty"`
	PortPolicy           *FabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPortChannelRole FabricPortChannelRole

// NewFabricPortChannelRole instantiates a new FabricPortChannelRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortChannelRole(classId string, objectType string) *FabricPortChannelRole {
	this := FabricPortChannelRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricPortChannelRoleWithDefaults instantiates a new FabricPortChannelRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortChannelRoleWithDefaults() *FabricPortChannelRole {
	this := FabricPortChannelRole{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPortChannelRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPortChannelRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPortChannelRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPortChannelRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricPortChannelRole) GetPcId() int64 {
	if o == nil || o.PcId == nil {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetPcIdOk() (*int64, bool) {
	if o == nil || o.PcId == nil {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPcId() bool {
	if o != nil && o.PcId != nil {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricPortChannelRole) SetPcId(v int64) {
	o.PcId = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricPortChannelRole) GetPorts() []FabricPortIdentifier {
	if o == nil {
		var ret []FabricPortIdentifier
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricPortChannelRole) GetPortsOk() (*[]FabricPortIdentifier, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return &o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []FabricPortIdentifier and assigns it to the Ports field.
func (o *FabricPortChannelRole) SetPorts(v []FabricPortIdentifier) {
	o.Ports = v
}

// GetPortPolicy returns the PortPolicy field value if set, zero value otherwise.
func (o *FabricPortChannelRole) GetPortPolicy() FabricPortPolicyRelationship {
	if o == nil || o.PortPolicy == nil {
		var ret FabricPortPolicyRelationship
		return ret
	}
	return *o.PortPolicy
}

// GetPortPolicyOk returns a tuple with the PortPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortChannelRole) GetPortPolicyOk() (*FabricPortPolicyRelationship, bool) {
	if o == nil || o.PortPolicy == nil {
		return nil, false
	}
	return o.PortPolicy, true
}

// HasPortPolicy returns a boolean if a field has been set.
func (o *FabricPortChannelRole) HasPortPolicy() bool {
	if o != nil && o.PortPolicy != nil {
		return true
	}

	return false
}

// SetPortPolicy gets a reference to the given FabricPortPolicyRelationship and assigns it to the PortPolicy field.
func (o *FabricPortChannelRole) SetPortPolicy(v FabricPortPolicyRelationship) {
	o.PortPolicy = &v
}

func (o FabricPortChannelRole) MarshalJSON() ([]byte, error) {
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
	if o.PcId != nil {
		toSerialize["PcId"] = o.PcId
	}
	if o.Ports != nil {
		toSerialize["Ports"] = o.Ports
	}
	if o.PortPolicy != nil {
		toSerialize["PortPolicy"] = o.PortPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPortChannelRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricPortChannelRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Unique Identifier of the port-channel, local to this switch.
		PcId       *int64                        `json:"PcId,omitempty"`
		Ports      []FabricPortIdentifier        `json:"Ports,omitempty"`
		PortPolicy *FabricPortPolicyRelationship `json:"PortPolicy,omitempty"`
	}

	varFabricPortChannelRoleWithoutEmbeddedStruct := FabricPortChannelRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricPortChannelRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricPortChannelRole := _FabricPortChannelRole{}
		varFabricPortChannelRole.ClassId = varFabricPortChannelRoleWithoutEmbeddedStruct.ClassId
		varFabricPortChannelRole.ObjectType = varFabricPortChannelRoleWithoutEmbeddedStruct.ObjectType
		varFabricPortChannelRole.PcId = varFabricPortChannelRoleWithoutEmbeddedStruct.PcId
		varFabricPortChannelRole.Ports = varFabricPortChannelRoleWithoutEmbeddedStruct.Ports
		varFabricPortChannelRole.PortPolicy = varFabricPortChannelRoleWithoutEmbeddedStruct.PortPolicy
		*o = FabricPortChannelRole(varFabricPortChannelRole)
	} else {
		return err
	}

	varFabricPortChannelRole := _FabricPortChannelRole{}

	err = json.Unmarshal(bytes, &varFabricPortChannelRole)
	if err == nil {
		o.MoBaseMo = varFabricPortChannelRole.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PcId")
		delete(additionalProperties, "Ports")
		delete(additionalProperties, "PortPolicy")

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

type NullableFabricPortChannelRole struct {
	value *FabricPortChannelRole
	isSet bool
}

func (v NullableFabricPortChannelRole) Get() *FabricPortChannelRole {
	return v.value
}

func (v *NullableFabricPortChannelRole) Set(val *FabricPortChannelRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortChannelRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortChannelRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortChannelRole(val *FabricPortChannelRole) *NullableFabricPortChannelRole {
	return &NullableFabricPortChannelRole{value: val, isSet: true}
}

func (v NullableFabricPortChannelRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortChannelRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
