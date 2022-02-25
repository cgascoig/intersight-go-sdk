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

// UuidpoolPool Pool represents a collection of UUID items that can be allocated to server profiles.
type UuidpoolPool struct {
	PoolAbstractPool
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx.
	Prefix           *string             `json:"Prefix,omitempty"`
	UuidSuffixBlocks []UuidpoolUuidBlock `json:"UuidSuffixBlocks,omitempty"`
	// An array of relationships to uuidpoolBlock resources.
	BlockHeads           []UuidpoolBlockRelationship           `json:"BlockHeads,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolPool UuidpoolPool

// NewUuidpoolPool instantiates a new UuidpoolPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolPool(classId string, objectType string) *UuidpoolPool {
	this := UuidpoolPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	var assignmentOrder string = "sequential"
	this.AssignmentOrder = &assignmentOrder
	return &this
}

// NewUuidpoolPoolWithDefaults instantiates a new UuidpoolPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolPoolWithDefaults() *UuidpoolPool {
	this := UuidpoolPool{}
	var classId string = "uuidpool.Pool"
	this.ClassId = classId
	var objectType string = "uuidpool.Pool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UuidpoolPool) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UuidpoolPool) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UuidpoolPool) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUuidSuffixBlocks returns the UuidSuffixBlocks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetUuidSuffixBlocks() []UuidpoolUuidBlock {
	if o == nil {
		var ret []UuidpoolUuidBlock
		return ret
	}
	return o.UuidSuffixBlocks
}

// GetUuidSuffixBlocksOk returns a tuple with the UuidSuffixBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetUuidSuffixBlocksOk() (*[]UuidpoolUuidBlock, bool) {
	if o == nil || o.UuidSuffixBlocks == nil {
		return nil, false
	}
	return &o.UuidSuffixBlocks, true
}

// HasUuidSuffixBlocks returns a boolean if a field has been set.
func (o *UuidpoolPool) HasUuidSuffixBlocks() bool {
	if o != nil && o.UuidSuffixBlocks != nil {
		return true
	}

	return false
}

// SetUuidSuffixBlocks gets a reference to the given []UuidpoolUuidBlock and assigns it to the UuidSuffixBlocks field.
func (o *UuidpoolPool) SetUuidSuffixBlocks(v []UuidpoolUuidBlock) {
	o.UuidSuffixBlocks = v
}

// GetBlockHeads returns the BlockHeads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolPool) GetBlockHeads() []UuidpoolBlockRelationship {
	if o == nil {
		var ret []UuidpoolBlockRelationship
		return ret
	}
	return o.BlockHeads
}

// GetBlockHeadsOk returns a tuple with the BlockHeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolPool) GetBlockHeadsOk() (*[]UuidpoolBlockRelationship, bool) {
	if o == nil || o.BlockHeads == nil {
		return nil, false
	}
	return &o.BlockHeads, true
}

// HasBlockHeads returns a boolean if a field has been set.
func (o *UuidpoolPool) HasBlockHeads() bool {
	if o != nil && o.BlockHeads != nil {
		return true
	}

	return false
}

// SetBlockHeads gets a reference to the given []UuidpoolBlockRelationship and assigns it to the BlockHeads field.
func (o *UuidpoolPool) SetBlockHeads(v []UuidpoolBlockRelationship) {
	o.BlockHeads = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UuidpoolPool) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolPool) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UuidpoolPool) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *UuidpoolPool) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o UuidpoolPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractPool, errPoolAbstractPool := json.Marshal(o.PoolAbstractPool)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	errPoolAbstractPool = json.Unmarshal([]byte(serializedPoolAbstractPool), &toSerialize)
	if errPoolAbstractPool != nil {
		return []byte{}, errPoolAbstractPool
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.UuidSuffixBlocks != nil {
		toSerialize["UuidSuffixBlocks"] = o.UuidSuffixBlocks
	}
	if o.BlockHeads != nil {
		toSerialize["BlockHeads"] = o.BlockHeads
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UuidpoolPool) UnmarshalJSON(bytes []byte) (err error) {
	type UuidpoolPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx.
		Prefix           *string             `json:"Prefix,omitempty"`
		UuidSuffixBlocks []UuidpoolUuidBlock `json:"UuidSuffixBlocks,omitempty"`
		// An array of relationships to uuidpoolBlock resources.
		BlockHeads   []UuidpoolBlockRelationship           `json:"BlockHeads,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varUuidpoolPoolWithoutEmbeddedStruct := UuidpoolPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varUuidpoolPoolWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolPool := _UuidpoolPool{}
		varUuidpoolPool.ClassId = varUuidpoolPoolWithoutEmbeddedStruct.ClassId
		varUuidpoolPool.ObjectType = varUuidpoolPoolWithoutEmbeddedStruct.ObjectType
		varUuidpoolPool.Prefix = varUuidpoolPoolWithoutEmbeddedStruct.Prefix
		varUuidpoolPool.UuidSuffixBlocks = varUuidpoolPoolWithoutEmbeddedStruct.UuidSuffixBlocks
		varUuidpoolPool.BlockHeads = varUuidpoolPoolWithoutEmbeddedStruct.BlockHeads
		varUuidpoolPool.Organization = varUuidpoolPoolWithoutEmbeddedStruct.Organization
		*o = UuidpoolPool(varUuidpoolPool)
	} else {
		return err
	}

	varUuidpoolPool := _UuidpoolPool{}

	err = json.Unmarshal(bytes, &varUuidpoolPool)
	if err == nil {
		o.PoolAbstractPool = varUuidpoolPool.PoolAbstractPool
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "UuidSuffixBlocks")
		delete(additionalProperties, "BlockHeads")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPoolAbstractPool := reflect.ValueOf(o.PoolAbstractPool)
		for i := 0; i < reflectPoolAbstractPool.Type().NumField(); i++ {
			t := reflectPoolAbstractPool.Type().Field(i)

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

type NullableUuidpoolPool struct {
	value *UuidpoolPool
	isSet bool
}

func (v NullableUuidpoolPool) Get() *UuidpoolPool {
	return v.value
}

func (v *NullableUuidpoolPool) Set(val *UuidpoolPool) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolPool) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolPool(val *UuidpoolPool) *NullableUuidpoolPool {
	return &NullableUuidpoolPool{value: val, isSet: true}
}

func (v NullableUuidpoolPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
