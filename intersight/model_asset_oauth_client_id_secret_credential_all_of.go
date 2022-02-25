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

// AssetOauthClientIdSecretCredentialAllOf Definition of the list of properties defined in 'asset.OauthClientIdSecretCredential', excluding properties defined in parent classes.
type AssetOauthClientIdSecretCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The client ID used to authenticate with a managed target.
	ClientId *string `json:"ClientId,omitempty"`
	// The client secret used to authenticate with a managed target.
	ClientSecret *string `json:"ClientSecret,omitempty"`
	// Indicates whether the value of the 'clientSecret' property has been set.
	IsClientSecretSet    *bool `json:"IsClientSecretSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOauthClientIdSecretCredentialAllOf AssetOauthClientIdSecretCredentialAllOf

// NewAssetOauthClientIdSecretCredentialAllOf instantiates a new AssetOauthClientIdSecretCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOauthClientIdSecretCredentialAllOf(classId string, objectType string) *AssetOauthClientIdSecretCredentialAllOf {
	this := AssetOauthClientIdSecretCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOauthClientIdSecretCredentialAllOfWithDefaults instantiates a new AssetOauthClientIdSecretCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOauthClientIdSecretCredentialAllOfWithDefaults() *AssetOauthClientIdSecretCredentialAllOf {
	this := AssetOauthClientIdSecretCredentialAllOf{}
	var classId string = "asset.OauthClientIdSecretCredential"
	this.ClassId = classId
	var objectType string = "asset.OauthClientIdSecretCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOauthClientIdSecretCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOauthClientIdSecretCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOauthClientIdSecretCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetOauthClientIdSecretCredentialAllOf) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AssetOauthClientIdSecretCredentialAllOf) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetIsClientSecretSet returns the IsClientSecretSet field value if set, zero value otherwise.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetIsClientSecretSet() bool {
	if o == nil || o.IsClientSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.IsClientSecretSet
}

// GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) GetIsClientSecretSetOk() (*bool, bool) {
	if o == nil || o.IsClientSecretSet == nil {
		return nil, false
	}
	return o.IsClientSecretSet, true
}

// HasIsClientSecretSet returns a boolean if a field has been set.
func (o *AssetOauthClientIdSecretCredentialAllOf) HasIsClientSecretSet() bool {
	if o != nil && o.IsClientSecretSet != nil {
		return true
	}

	return false
}

// SetIsClientSecretSet gets a reference to the given bool and assigns it to the IsClientSecretSet field.
func (o *AssetOauthClientIdSecretCredentialAllOf) SetIsClientSecretSet(v bool) {
	o.IsClientSecretSet = &v
}

func (o AssetOauthClientIdSecretCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["ClientSecret"] = o.ClientSecret
	}
	if o.IsClientSecretSet != nil {
		toSerialize["IsClientSecretSet"] = o.IsClientSecretSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOauthClientIdSecretCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetOauthClientIdSecretCredentialAllOf := _AssetOauthClientIdSecretCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varAssetOauthClientIdSecretCredentialAllOf); err == nil {
		*o = AssetOauthClientIdSecretCredentialAllOf(varAssetOauthClientIdSecretCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "ClientSecret")
		delete(additionalProperties, "IsClientSecretSet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOauthClientIdSecretCredentialAllOf struct {
	value *AssetOauthClientIdSecretCredentialAllOf
	isSet bool
}

func (v NullableAssetOauthClientIdSecretCredentialAllOf) Get() *AssetOauthClientIdSecretCredentialAllOf {
	return v.value
}

func (v *NullableAssetOauthClientIdSecretCredentialAllOf) Set(val *AssetOauthClientIdSecretCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOauthClientIdSecretCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOauthClientIdSecretCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOauthClientIdSecretCredentialAllOf(val *AssetOauthClientIdSecretCredentialAllOf) *NullableAssetOauthClientIdSecretCredentialAllOf {
	return &NullableAssetOauthClientIdSecretCredentialAllOf{value: val, isSet: true}
}

func (v NullableAssetOauthClientIdSecretCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOauthClientIdSecretCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
