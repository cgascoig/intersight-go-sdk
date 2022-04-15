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

// VirtualizationVmwareDiscoveryProtocolAllOf Definition of the list of properties defined in 'virtualization.VmwareDiscoveryProtocol', excluding properties defined in parent classes.
type VirtualizationVmwareDiscoveryProtocolAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operational mode of the ESXI hosts connected to the distributed virtual switch.
	Operation *string `json:"Operation,omitempty"`
	// Discovery protocol type enabled on the distributed virtual switch.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDiscoveryProtocolAllOf VirtualizationVmwareDiscoveryProtocolAllOf

// NewVirtualizationVmwareDiscoveryProtocolAllOf instantiates a new VirtualizationVmwareDiscoveryProtocolAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDiscoveryProtocolAllOf(classId string, objectType string) *VirtualizationVmwareDiscoveryProtocolAllOf {
	this := VirtualizationVmwareDiscoveryProtocolAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareDiscoveryProtocolAllOfWithDefaults instantiates a new VirtualizationVmwareDiscoveryProtocolAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDiscoveryProtocolAllOfWithDefaults() *VirtualizationVmwareDiscoveryProtocolAllOf {
	this := VirtualizationVmwareDiscoveryProtocolAllOf{}
	var classId string = "virtualization.VmwareDiscoveryProtocol"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDiscoveryProtocol"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetOperation(v string) {
	o.Operation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationVmwareDiscoveryProtocolAllOf) SetType(v string) {
	o.Type = &v
}

func (o VirtualizationVmwareDiscoveryProtocolAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Operation != nil {
		toSerialize["Operation"] = o.Operation
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareDiscoveryProtocolAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmwareDiscoveryProtocolAllOf := _VirtualizationVmwareDiscoveryProtocolAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmwareDiscoveryProtocolAllOf); err == nil {
		*o = VirtualizationVmwareDiscoveryProtocolAllOf(varVirtualizationVmwareDiscoveryProtocolAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Operation")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmwareDiscoveryProtocolAllOf struct {
	value *VirtualizationVmwareDiscoveryProtocolAllOf
	isSet bool
}

func (v NullableVirtualizationVmwareDiscoveryProtocolAllOf) Get() *VirtualizationVmwareDiscoveryProtocolAllOf {
	return v.value
}

func (v *NullableVirtualizationVmwareDiscoveryProtocolAllOf) Set(val *VirtualizationVmwareDiscoveryProtocolAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDiscoveryProtocolAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDiscoveryProtocolAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDiscoveryProtocolAllOf(val *VirtualizationVmwareDiscoveryProtocolAllOf) *NullableVirtualizationVmwareDiscoveryProtocolAllOf {
	return &NullableVirtualizationVmwareDiscoveryProtocolAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDiscoveryProtocolAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDiscoveryProtocolAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
