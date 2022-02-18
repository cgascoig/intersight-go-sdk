/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// AssetTargetSignatureAllOf Definition of the list of properties defined in 'asset.TargetSignature', excluding properties defined in parent classes.
type AssetTargetSignatureAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the associated target.
	TargetId             *string `json:"TargetId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetTargetSignatureAllOf AssetTargetSignatureAllOf

// NewAssetTargetSignatureAllOf instantiates a new AssetTargetSignatureAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetTargetSignatureAllOf(classId string, objectType string) *AssetTargetSignatureAllOf {
	this := AssetTargetSignatureAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetTargetSignatureAllOfWithDefaults instantiates a new AssetTargetSignatureAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTargetSignatureAllOfWithDefaults() *AssetTargetSignatureAllOf {
	this := AssetTargetSignatureAllOf{}
	var classId string = "asset.TargetSignature"
	this.ClassId = classId
	var objectType string = "asset.TargetSignature"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetTargetSignatureAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignatureAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetTargetSignatureAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetTargetSignatureAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetTargetSignatureAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetTargetSignatureAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AssetTargetSignatureAllOf) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetTargetSignatureAllOf) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AssetTargetSignatureAllOf) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AssetTargetSignatureAllOf) SetTargetId(v string) {
	o.TargetId = &v
}

func (o AssetTargetSignatureAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetTargetSignatureAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetTargetSignatureAllOf := _AssetTargetSignatureAllOf{}

	if err = json.Unmarshal(bytes, &varAssetTargetSignatureAllOf); err == nil {
		*o = AssetTargetSignatureAllOf(varAssetTargetSignatureAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetTargetSignatureAllOf struct {
	value *AssetTargetSignatureAllOf
	isSet bool
}

func (v NullableAssetTargetSignatureAllOf) Get() *AssetTargetSignatureAllOf {
	return v.value
}

func (v *NullableAssetTargetSignatureAllOf) Set(val *AssetTargetSignatureAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTargetSignatureAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTargetSignatureAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTargetSignatureAllOf(val *AssetTargetSignatureAllOf) *NullableAssetTargetSignatureAllOf {
	return &NullableAssetTargetSignatureAllOf{value: val, isSet: true}
}

func (v NullableAssetTargetSignatureAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTargetSignatureAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
