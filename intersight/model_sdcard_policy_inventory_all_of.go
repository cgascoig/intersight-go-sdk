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

// SdcardPolicyInventoryAllOf Definition of the list of properties defined in 'sdcard.PolicyInventory', excluding properties defined in parent classes.
type SdcardPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                `json:"ObjectType"`
	Partitions           []SdcardPartition     `json:"Partitions,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdcardPolicyInventoryAllOf SdcardPolicyInventoryAllOf

// NewSdcardPolicyInventoryAllOf instantiates a new SdcardPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdcardPolicyInventoryAllOf(classId string, objectType string) *SdcardPolicyInventoryAllOf {
	this := SdcardPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSdcardPolicyInventoryAllOfWithDefaults instantiates a new SdcardPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdcardPolicyInventoryAllOfWithDefaults() *SdcardPolicyInventoryAllOf {
	this := SdcardPolicyInventoryAllOf{}
	var classId string = "sdcard.PolicyInventory"
	this.ClassId = classId
	var objectType string = "sdcard.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdcardPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdcardPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdcardPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdcardPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdcardPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdcardPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdcardPolicyInventoryAllOf) GetPartitions() []SdcardPartition {
	if o == nil {
		var ret []SdcardPartition
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdcardPolicyInventoryAllOf) GetPartitionsOk() (*[]SdcardPartition, bool) {
	if o == nil || o.Partitions == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *SdcardPolicyInventoryAllOf) HasPartitions() bool {
	if o != nil && o.Partitions != nil {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []SdcardPartition and assigns it to the Partitions field.
func (o *SdcardPolicyInventoryAllOf) SetPartitions(v []SdcardPartition) {
	o.Partitions = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *SdcardPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdcardPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *SdcardPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *SdcardPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o SdcardPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Partitions != nil {
		toSerialize["Partitions"] = o.Partitions
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdcardPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSdcardPolicyInventoryAllOf := _SdcardPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varSdcardPolicyInventoryAllOf); err == nil {
		*o = SdcardPolicyInventoryAllOf(varSdcardPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Partitions")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSdcardPolicyInventoryAllOf struct {
	value *SdcardPolicyInventoryAllOf
	isSet bool
}

func (v NullableSdcardPolicyInventoryAllOf) Get() *SdcardPolicyInventoryAllOf {
	return v.value
}

func (v *NullableSdcardPolicyInventoryAllOf) Set(val *SdcardPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSdcardPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSdcardPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdcardPolicyInventoryAllOf(val *SdcardPolicyInventoryAllOf) *NullableSdcardPolicyInventoryAllOf {
	return &NullableSdcardPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableSdcardPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdcardPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
