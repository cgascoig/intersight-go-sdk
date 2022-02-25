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

// FabricPortIdentifierAllOf Definition of the list of properties defined in 'fabric.PortIdentifier', excluding properties defined in parent classes.
type FabricPortIdentifierAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Breakout port Identifier of the Switch Interface. When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused. When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment, e.g. the id of the port on the switch.
	AggregatePortId *int64 `json:"AggregatePortId,omitempty"`
	// Port Identifier of the Switch/FEX/Chassis Interface. When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment, e.g. the id of the port on the switch, FEX or chassis. When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable.
	PortId *int64 `json:"PortId,omitempty"`
	// Slot Identifier of the Switch/FEX/Chassis Interface.
	SlotId               *int64 `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricPortIdentifierAllOf FabricPortIdentifierAllOf

// NewFabricPortIdentifierAllOf instantiates a new FabricPortIdentifierAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricPortIdentifierAllOf(classId string, objectType string) *FabricPortIdentifierAllOf {
	this := FabricPortIdentifierAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricPortIdentifierAllOfWithDefaults instantiates a new FabricPortIdentifierAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricPortIdentifierAllOfWithDefaults() *FabricPortIdentifierAllOf {
	this := FabricPortIdentifierAllOf{}
	var classId string = "fabric.PortIdentifier"
	this.ClassId = classId
	var objectType string = "fabric.PortIdentifier"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricPortIdentifierAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifierAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricPortIdentifierAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricPortIdentifierAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifierAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricPortIdentifierAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAggregatePortId returns the AggregatePortId field value if set, zero value otherwise.
func (o *FabricPortIdentifierAllOf) GetAggregatePortId() int64 {
	if o == nil || o.AggregatePortId == nil {
		var ret int64
		return ret
	}
	return *o.AggregatePortId
}

// GetAggregatePortIdOk returns a tuple with the AggregatePortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifierAllOf) GetAggregatePortIdOk() (*int64, bool) {
	if o == nil || o.AggregatePortId == nil {
		return nil, false
	}
	return o.AggregatePortId, true
}

// HasAggregatePortId returns a boolean if a field has been set.
func (o *FabricPortIdentifierAllOf) HasAggregatePortId() bool {
	if o != nil && o.AggregatePortId != nil {
		return true
	}

	return false
}

// SetAggregatePortId gets a reference to the given int64 and assigns it to the AggregatePortId field.
func (o *FabricPortIdentifierAllOf) SetAggregatePortId(v int64) {
	o.AggregatePortId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *FabricPortIdentifierAllOf) GetPortId() int64 {
	if o == nil || o.PortId == nil {
		var ret int64
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifierAllOf) GetPortIdOk() (*int64, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *FabricPortIdentifierAllOf) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int64 and assigns it to the PortId field.
func (o *FabricPortIdentifierAllOf) SetPortId(v int64) {
	o.PortId = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *FabricPortIdentifierAllOf) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricPortIdentifierAllOf) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *FabricPortIdentifierAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *FabricPortIdentifierAllOf) SetSlotId(v int64) {
	o.SlotId = &v
}

func (o FabricPortIdentifierAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AggregatePortId != nil {
		toSerialize["AggregatePortId"] = o.AggregatePortId
	}
	if o.PortId != nil {
		toSerialize["PortId"] = o.PortId
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricPortIdentifierAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricPortIdentifierAllOf := _FabricPortIdentifierAllOf{}

	if err = json.Unmarshal(bytes, &varFabricPortIdentifierAllOf); err == nil {
		*o = FabricPortIdentifierAllOf(varFabricPortIdentifierAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AggregatePortId")
		delete(additionalProperties, "PortId")
		delete(additionalProperties, "SlotId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricPortIdentifierAllOf struct {
	value *FabricPortIdentifierAllOf
	isSet bool
}

func (v NullableFabricPortIdentifierAllOf) Get() *FabricPortIdentifierAllOf {
	return v.value
}

func (v *NullableFabricPortIdentifierAllOf) Set(val *FabricPortIdentifierAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortIdentifierAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortIdentifierAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortIdentifierAllOf(val *FabricPortIdentifierAllOf) *NullableFabricPortIdentifierAllOf {
	return &NullableFabricPortIdentifierAllOf{value: val, isSet: true}
}

func (v NullableFabricPortIdentifierAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortIdentifierAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
