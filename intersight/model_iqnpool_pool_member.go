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

// IqnpoolPoolMember PoolMember represents a single IQN address that is part of a pool.
type IqnpoolPoolMember struct {
	PoolAbstractPoolMember
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IQN Address of this pool member.
	IqnAddress           *string                   `json:"IqnAddress,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship     `json:"AssignedToEntity,omitempty"`
	BlockHead            *IqnpoolBlockRelationship `json:"BlockHead,omitempty"`
	Peer                 *IqnpoolLeaseRelationship `json:"Peer,omitempty"`
	Pool                 *IqnpoolPoolRelationship  `json:"Pool,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IqnpoolPoolMember IqnpoolPoolMember

// NewIqnpoolPoolMember instantiates a new IqnpoolPoolMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIqnpoolPoolMember(classId string, objectType string) *IqnpoolPoolMember {
	this := IqnpoolPoolMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assigned bool = false
	this.Assigned = &assigned
	return &this
}

// NewIqnpoolPoolMemberWithDefaults instantiates a new IqnpoolPoolMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIqnpoolPoolMemberWithDefaults() *IqnpoolPoolMember {
	this := IqnpoolPoolMember{}
	var classId string = "iqnpool.PoolMember"
	this.ClassId = classId
	var objectType string = "iqnpool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IqnpoolPoolMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IqnpoolPoolMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IqnpoolPoolMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IqnpoolPoolMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIqnAddress returns the IqnAddress field value if set, zero value otherwise.
func (o *IqnpoolPoolMember) GetIqnAddress() string {
	if o == nil || o.IqnAddress == nil {
		var ret string
		return ret
	}
	return *o.IqnAddress
}

// GetIqnAddressOk returns a tuple with the IqnAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetIqnAddressOk() (*string, bool) {
	if o == nil || o.IqnAddress == nil {
		return nil, false
	}
	return o.IqnAddress, true
}

// HasIqnAddress returns a boolean if a field has been set.
func (o *IqnpoolPoolMember) HasIqnAddress() bool {
	if o != nil && o.IqnAddress != nil {
		return true
	}

	return false
}

// SetIqnAddress gets a reference to the given string and assigns it to the IqnAddress field.
func (o *IqnpoolPoolMember) SetIqnAddress(v string) {
	o.IqnAddress = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *IqnpoolPoolMember) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IqnpoolPoolMember) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IqnpoolPoolMember) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetBlockHead returns the BlockHead field value if set, zero value otherwise.
func (o *IqnpoolPoolMember) GetBlockHead() IqnpoolBlockRelationship {
	if o == nil || o.BlockHead == nil {
		var ret IqnpoolBlockRelationship
		return ret
	}
	return *o.BlockHead
}

// GetBlockHeadOk returns a tuple with the BlockHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetBlockHeadOk() (*IqnpoolBlockRelationship, bool) {
	if o == nil || o.BlockHead == nil {
		return nil, false
	}
	return o.BlockHead, true
}

// HasBlockHead returns a boolean if a field has been set.
func (o *IqnpoolPoolMember) HasBlockHead() bool {
	if o != nil && o.BlockHead != nil {
		return true
	}

	return false
}

// SetBlockHead gets a reference to the given IqnpoolBlockRelationship and assigns it to the BlockHead field.
func (o *IqnpoolPoolMember) SetBlockHead(v IqnpoolBlockRelationship) {
	o.BlockHead = &v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *IqnpoolPoolMember) GetPeer() IqnpoolLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret IqnpoolLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetPeerOk() (*IqnpoolLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *IqnpoolPoolMember) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given IqnpoolLeaseRelationship and assigns it to the Peer field.
func (o *IqnpoolPoolMember) SetPeer(v IqnpoolLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *IqnpoolPoolMember) GetPool() IqnpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IqnpoolPoolMember) GetPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *IqnpoolPoolMember) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given IqnpoolPoolRelationship and assigns it to the Pool field.
func (o *IqnpoolPoolMember) SetPool(v IqnpoolPoolRelationship) {
	o.Pool = &v
}

func (o IqnpoolPoolMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPoolMember, errPoolAbstractPoolMember := json.Marshal(o.PoolAbstractPoolMember)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	errPoolAbstractPoolMember = json.Unmarshal([]byte(serializedPoolAbstractPoolMember), &toSerialize)
	if errPoolAbstractPoolMember != nil {
		return []byte{}, errPoolAbstractPoolMember
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IqnAddress != nil {
		toSerialize["IqnAddress"] = o.IqnAddress
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.BlockHead != nil {
		toSerialize["BlockHead"] = o.BlockHead
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IqnpoolPoolMember) UnmarshalJSON(bytes []byte) (err error) {
	type IqnpoolPoolMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IQN Address of this pool member.
		IqnAddress       *string                   `json:"IqnAddress,omitempty"`
		AssignedToEntity *MoBaseMoRelationship     `json:"AssignedToEntity,omitempty"`
		BlockHead        *IqnpoolBlockRelationship `json:"BlockHead,omitempty"`
		Peer             *IqnpoolLeaseRelationship `json:"Peer,omitempty"`
		Pool             *IqnpoolPoolRelationship  `json:"Pool,omitempty"`
	}

	varIqnpoolPoolMemberWithoutEmbeddedStruct := IqnpoolPoolMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIqnpoolPoolMemberWithoutEmbeddedStruct)
	if err == nil {
		varIqnpoolPoolMember := _IqnpoolPoolMember{}
		varIqnpoolPoolMember.ClassId = varIqnpoolPoolMemberWithoutEmbeddedStruct.ClassId
		varIqnpoolPoolMember.ObjectType = varIqnpoolPoolMemberWithoutEmbeddedStruct.ObjectType
		varIqnpoolPoolMember.IqnAddress = varIqnpoolPoolMemberWithoutEmbeddedStruct.IqnAddress
		varIqnpoolPoolMember.AssignedToEntity = varIqnpoolPoolMemberWithoutEmbeddedStruct.AssignedToEntity
		varIqnpoolPoolMember.BlockHead = varIqnpoolPoolMemberWithoutEmbeddedStruct.BlockHead
		varIqnpoolPoolMember.Peer = varIqnpoolPoolMemberWithoutEmbeddedStruct.Peer
		varIqnpoolPoolMember.Pool = varIqnpoolPoolMemberWithoutEmbeddedStruct.Pool
		*o = IqnpoolPoolMember(varIqnpoolPoolMember)
	} else {
		return err
	}

	varIqnpoolPoolMember := _IqnpoolPoolMember{}

	err = json.Unmarshal(bytes, &varIqnpoolPoolMember)
	if err == nil {
		o.PoolAbstractPoolMember = varIqnpoolPoolMember.PoolAbstractPoolMember
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IqnAddress")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "BlockHead")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")

		// remove fields from embedded structs
		reflectPoolAbstractPoolMember := reflect.ValueOf(o.PoolAbstractPoolMember)
		for i := 0; i < reflectPoolAbstractPoolMember.Type().NumField(); i++ {
			t := reflectPoolAbstractPoolMember.Type().Field(i)

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

type NullableIqnpoolPoolMember struct {
	value *IqnpoolPoolMember
	isSet bool
}

func (v NullableIqnpoolPoolMember) Get() *IqnpoolPoolMember {
	return v.value
}

func (v *NullableIqnpoolPoolMember) Set(val *IqnpoolPoolMember) {
	v.value = val
	v.isSet = true
}

func (v NullableIqnpoolPoolMember) IsSet() bool {
	return v.isSet
}

func (v *NullableIqnpoolPoolMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIqnpoolPoolMember(val *IqnpoolPoolMember) *NullableIqnpoolPoolMember {
	return &NullableIqnpoolPoolMember{value: val, isSet: true}
}

func (v NullableIqnpoolPoolMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIqnpoolPoolMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
