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

// IamPermissionReferenceAllOf Definition of the list of properties defined in 'iam.PermissionReference', excluding properties defined in parent classes.
type IamPermissionReferenceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MOID of the permission which user has access to.
	PermissionIdentifier *string `json:"PermissionIdentifier,omitempty"`
	// Name of the permission which user has access to.
	PermissionName       *string `json:"PermissionName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamPermissionReferenceAllOf IamPermissionReferenceAllOf

// NewIamPermissionReferenceAllOf instantiates a new IamPermissionReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermissionReferenceAllOf(classId string, objectType string) *IamPermissionReferenceAllOf {
	this := IamPermissionReferenceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamPermissionReferenceAllOfWithDefaults instantiates a new IamPermissionReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionReferenceAllOfWithDefaults() *IamPermissionReferenceAllOf {
	this := IamPermissionReferenceAllOf{}
	var classId string = "iam.PermissionReference"
	this.ClassId = classId
	var objectType string = "iam.PermissionReference"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamPermissionReferenceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamPermissionReferenceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamPermissionReferenceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamPermissionReferenceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamPermissionReferenceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamPermissionReferenceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermissionIdentifier returns the PermissionIdentifier field value if set, zero value otherwise.
func (o *IamPermissionReferenceAllOf) GetPermissionIdentifier() string {
	if o == nil || o.PermissionIdentifier == nil {
		var ret string
		return ret
	}
	return *o.PermissionIdentifier
}

// GetPermissionIdentifierOk returns a tuple with the PermissionIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionReferenceAllOf) GetPermissionIdentifierOk() (*string, bool) {
	if o == nil || o.PermissionIdentifier == nil {
		return nil, false
	}
	return o.PermissionIdentifier, true
}

// HasPermissionIdentifier returns a boolean if a field has been set.
func (o *IamPermissionReferenceAllOf) HasPermissionIdentifier() bool {
	if o != nil && o.PermissionIdentifier != nil {
		return true
	}

	return false
}

// SetPermissionIdentifier gets a reference to the given string and assigns it to the PermissionIdentifier field.
func (o *IamPermissionReferenceAllOf) SetPermissionIdentifier(v string) {
	o.PermissionIdentifier = &v
}

// GetPermissionName returns the PermissionName field value if set, zero value otherwise.
func (o *IamPermissionReferenceAllOf) GetPermissionName() string {
	if o == nil || o.PermissionName == nil {
		var ret string
		return ret
	}
	return *o.PermissionName
}

// GetPermissionNameOk returns a tuple with the PermissionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermissionReferenceAllOf) GetPermissionNameOk() (*string, bool) {
	if o == nil || o.PermissionName == nil {
		return nil, false
	}
	return o.PermissionName, true
}

// HasPermissionName returns a boolean if a field has been set.
func (o *IamPermissionReferenceAllOf) HasPermissionName() bool {
	if o != nil && o.PermissionName != nil {
		return true
	}

	return false
}

// SetPermissionName gets a reference to the given string and assigns it to the PermissionName field.
func (o *IamPermissionReferenceAllOf) SetPermissionName(v string) {
	o.PermissionName = &v
}

func (o IamPermissionReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PermissionIdentifier != nil {
		toSerialize["PermissionIdentifier"] = o.PermissionIdentifier
	}
	if o.PermissionName != nil {
		toSerialize["PermissionName"] = o.PermissionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamPermissionReferenceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamPermissionReferenceAllOf := _IamPermissionReferenceAllOf{}

	if err = json.Unmarshal(bytes, &varIamPermissionReferenceAllOf); err == nil {
		*o = IamPermissionReferenceAllOf(varIamPermissionReferenceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PermissionIdentifier")
		delete(additionalProperties, "PermissionName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamPermissionReferenceAllOf struct {
	value *IamPermissionReferenceAllOf
	isSet bool
}

func (v NullableIamPermissionReferenceAllOf) Get() *IamPermissionReferenceAllOf {
	return v.value
}

func (v *NullableIamPermissionReferenceAllOf) Set(val *IamPermissionReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamPermissionReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamPermissionReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamPermissionReferenceAllOf(val *IamPermissionReferenceAllOf) *NullableIamPermissionReferenceAllOf {
	return &NullableIamPermissionReferenceAllOf{value: val, isSet: true}
}

func (v NullableIamPermissionReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamPermissionReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
