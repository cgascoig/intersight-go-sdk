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

// IamTrustPointAllOf Definition of the list of properties defined in 'iam.TrustPoint', excluding properties defined in parent classes.
type IamTrustPointAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string            `json:"ObjectType"`
	Certificates []X509Certificate `json:"Certificates,omitempty"`
	// The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format.
	Chain                *string                 `json:"Chain,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamTrustPointAllOf IamTrustPointAllOf

// NewIamTrustPointAllOf instantiates a new IamTrustPointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamTrustPointAllOf(classId string, objectType string) *IamTrustPointAllOf {
	this := IamTrustPointAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamTrustPointAllOfWithDefaults instantiates a new IamTrustPointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamTrustPointAllOfWithDefaults() *IamTrustPointAllOf {
	this := IamTrustPointAllOf{}
	var classId string = "iam.TrustPoint"
	this.ClassId = classId
	var objectType string = "iam.TrustPoint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamTrustPointAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamTrustPointAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamTrustPointAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamTrustPointAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamTrustPointAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamTrustPointAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamTrustPointAllOf) GetCertificates() []X509Certificate {
	if o == nil {
		var ret []X509Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamTrustPointAllOf) GetCertificatesOk() (*[]X509Certificate, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return &o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *IamTrustPointAllOf) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []X509Certificate and assigns it to the Certificates field.
func (o *IamTrustPointAllOf) SetCertificates(v []X509Certificate) {
	o.Certificates = v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *IamTrustPointAllOf) GetChain() string {
	if o == nil || o.Chain == nil {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamTrustPointAllOf) GetChainOk() (*string, bool) {
	if o == nil || o.Chain == nil {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *IamTrustPointAllOf) HasChain() bool {
	if o != nil && o.Chain != nil {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *IamTrustPointAllOf) SetChain(v string) {
	o.Chain = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamTrustPointAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamTrustPointAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamTrustPointAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamTrustPointAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamTrustPointAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Certificates != nil {
		toSerialize["Certificates"] = o.Certificates
	}
	if o.Chain != nil {
		toSerialize["Chain"] = o.Chain
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamTrustPointAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamTrustPointAllOf := _IamTrustPointAllOf{}

	if err = json.Unmarshal(bytes, &varIamTrustPointAllOf); err == nil {
		*o = IamTrustPointAllOf(varIamTrustPointAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificates")
		delete(additionalProperties, "Chain")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamTrustPointAllOf struct {
	value *IamTrustPointAllOf
	isSet bool
}

func (v NullableIamTrustPointAllOf) Get() *IamTrustPointAllOf {
	return v.value
}

func (v *NullableIamTrustPointAllOf) Set(val *IamTrustPointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamTrustPointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamTrustPointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamTrustPointAllOf(val *IamTrustPointAllOf) *NullableIamTrustPointAllOf {
	return &NullableIamTrustPointAllOf{value: val, isSet: true}
}

func (v NullableIamTrustPointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamTrustPointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
