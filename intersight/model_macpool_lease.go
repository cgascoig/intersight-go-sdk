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

// MacpoolLease Lease represents a single MAC address that is part of the universe, allocated either from a pool or through static assignment.
type MacpoolLease struct {
	PoolAbstractLease
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MAC address allocated for pool-based allocation.
	MacAddress           *string                        `json:"MacAddress,omitempty"`
	AssignedToEntity     *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
	Pool                 *MacpoolPoolRelationship       `json:"Pool,omitempty"`
	PoolMember           *MacpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
	Universe             *MacpoolUniverseRelationship   `json:"Universe,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolLease MacpoolLease

// NewMacpoolLease instantiates a new MacpoolLease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolLease(classId string, objectType string) *MacpoolLease {
	this := MacpoolLease{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allocationType string = "dynamic"
	this.AllocationType = &allocationType
	return &this
}

// NewMacpoolLeaseWithDefaults instantiates a new MacpoolLease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolLeaseWithDefaults() *MacpoolLease {
	this := MacpoolLease{}
	var classId string = "macpool.Lease"
	this.ClassId = classId
	var objectType string = "macpool.Lease"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolLease) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolLease) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolLease) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolLease) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *MacpoolLease) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *MacpoolLease) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *MacpoolLease) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise.
func (o *MacpoolLease) GetAssignedToEntity() MoBaseMoRelationship {
	if o == nil || o.AssignedToEntity == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetAssignedToEntityOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *MacpoolLease) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *MacpoolLease) SetAssignedToEntity(v MoBaseMoRelationship) {
	o.AssignedToEntity = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *MacpoolLease) GetPool() MacpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolLease) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given MacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolLease) SetPool(v MacpoolPoolRelationship) {
	o.Pool = &v
}

// GetPoolMember returns the PoolMember field value if set, zero value otherwise.
func (o *MacpoolLease) GetPoolMember() MacpoolPoolMemberRelationship {
	if o == nil || o.PoolMember == nil {
		var ret MacpoolPoolMemberRelationship
		return ret
	}
	return *o.PoolMember
}

// GetPoolMemberOk returns a tuple with the PoolMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetPoolMemberOk() (*MacpoolPoolMemberRelationship, bool) {
	if o == nil || o.PoolMember == nil {
		return nil, false
	}
	return o.PoolMember, true
}

// HasPoolMember returns a boolean if a field has been set.
func (o *MacpoolLease) HasPoolMember() bool {
	if o != nil && o.PoolMember != nil {
		return true
	}

	return false
}

// SetPoolMember gets a reference to the given MacpoolPoolMemberRelationship and assigns it to the PoolMember field.
func (o *MacpoolLease) SetPoolMember(v MacpoolPoolMemberRelationship) {
	o.PoolMember = &v
}

// GetUniverse returns the Universe field value if set, zero value otherwise.
func (o *MacpoolLease) GetUniverse() MacpoolUniverseRelationship {
	if o == nil || o.Universe == nil {
		var ret MacpoolUniverseRelationship
		return ret
	}
	return *o.Universe
}

// GetUniverseOk returns a tuple with the Universe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolLease) GetUniverseOk() (*MacpoolUniverseRelationship, bool) {
	if o == nil || o.Universe == nil {
		return nil, false
	}
	return o.Universe, true
}

// HasUniverse returns a boolean if a field has been set.
func (o *MacpoolLease) HasUniverse() bool {
	if o != nil && o.Universe != nil {
		return true
	}

	return false
}

// SetUniverse gets a reference to the given MacpoolUniverseRelationship and assigns it to the Universe field.
func (o *MacpoolLease) SetUniverse(v MacpoolUniverseRelationship) {
	o.Universe = &v
}

func (o MacpoolLease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractLease, errPoolAbstractLease := json.Marshal(o.PoolAbstractLease)
	if errPoolAbstractLease != nil {
		return []byte{}, errPoolAbstractLease
	}
	errPoolAbstractLease = json.Unmarshal([]byte(serializedPoolAbstractLease), &toSerialize)
	if errPoolAbstractLease != nil {
		return []byte{}, errPoolAbstractLease
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.PoolMember != nil {
		toSerialize["PoolMember"] = o.PoolMember
	}
	if o.Universe != nil {
		toSerialize["Universe"] = o.Universe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MacpoolLease) UnmarshalJSON(bytes []byte) (err error) {
	type MacpoolLeaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// MAC address allocated for pool-based allocation.
		MacAddress       *string                        `json:"MacAddress,omitempty"`
		AssignedToEntity *MoBaseMoRelationship          `json:"AssignedToEntity,omitempty"`
		Pool             *MacpoolPoolRelationship       `json:"Pool,omitempty"`
		PoolMember       *MacpoolPoolMemberRelationship `json:"PoolMember,omitempty"`
		Universe         *MacpoolUniverseRelationship   `json:"Universe,omitempty"`
	}

	varMacpoolLeaseWithoutEmbeddedStruct := MacpoolLeaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMacpoolLeaseWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolLease := _MacpoolLease{}
		varMacpoolLease.ClassId = varMacpoolLeaseWithoutEmbeddedStruct.ClassId
		varMacpoolLease.ObjectType = varMacpoolLeaseWithoutEmbeddedStruct.ObjectType
		varMacpoolLease.MacAddress = varMacpoolLeaseWithoutEmbeddedStruct.MacAddress
		varMacpoolLease.AssignedToEntity = varMacpoolLeaseWithoutEmbeddedStruct.AssignedToEntity
		varMacpoolLease.Pool = varMacpoolLeaseWithoutEmbeddedStruct.Pool
		varMacpoolLease.PoolMember = varMacpoolLeaseWithoutEmbeddedStruct.PoolMember
		varMacpoolLease.Universe = varMacpoolLeaseWithoutEmbeddedStruct.Universe
		*o = MacpoolLease(varMacpoolLease)
	} else {
		return err
	}

	varMacpoolLease := _MacpoolLease{}

	err = json.Unmarshal(bytes, &varMacpoolLease)
	if err == nil {
		o.PoolAbstractLease = varMacpoolLease.PoolAbstractLease
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "PoolMember")
		delete(additionalProperties, "Universe")

		// remove fields from embedded structs
		reflectPoolAbstractLease := reflect.ValueOf(o.PoolAbstractLease)
		for i := 0; i < reflectPoolAbstractLease.Type().NumField(); i++ {
			t := reflectPoolAbstractLease.Type().Field(i)

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

type NullableMacpoolLease struct {
	value *MacpoolLease
	isSet bool
}

func (v NullableMacpoolLease) Get() *MacpoolLease {
	return v.value
}

func (v *NullableMacpoolLease) Set(val *MacpoolLease) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolLease) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolLease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolLease(val *MacpoolLease) *NullableMacpoolLease {
	return &NullableMacpoolLease{value: val, isSet: true}
}

func (v NullableMacpoolLease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolLease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
