/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IamLdapGroupAllOf Definition of the list of properties defined in 'iam.LdapGroup', excluding properties defined in parent classes.
type IamLdapGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// LDAP server domain the Group resides in.
	Domain *string `json:"Domain,omitempty"`
	// LDAP Group name in the LDAP server database.
	Name *string `json:"Name,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRole         []IamEndPointRoleRelationship `json:"EndPointRole,omitempty"`
	LdapPolicy           *IamLdapPolicyRelationship    `json:"LdapPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamLdapGroupAllOf IamLdapGroupAllOf

// NewIamLdapGroupAllOf instantiates a new IamLdapGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamLdapGroupAllOf(classId string, objectType string) *IamLdapGroupAllOf {
	this := IamLdapGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamLdapGroupAllOfWithDefaults instantiates a new IamLdapGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamLdapGroupAllOfWithDefaults() *IamLdapGroupAllOf {
	this := IamLdapGroupAllOf{}
	var classId string = "iam.LdapGroup"
	this.ClassId = classId
	var objectType string = "iam.LdapGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamLdapGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamLdapGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamLdapGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamLdapGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamLdapGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamLdapGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *IamLdapGroupAllOf) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapGroupAllOf) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *IamLdapGroupAllOf) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *IamLdapGroupAllOf) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamLdapGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamLdapGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamLdapGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetEndPointRole returns the EndPointRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamLdapGroupAllOf) GetEndPointRole() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRole
}

// GetEndPointRoleOk returns a tuple with the EndPointRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamLdapGroupAllOf) GetEndPointRoleOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRole == nil {
		return nil, false
	}
	return &o.EndPointRole, true
}

// HasEndPointRole returns a boolean if a field has been set.
func (o *IamLdapGroupAllOf) HasEndPointRole() bool {
	if o != nil && o.EndPointRole != nil {
		return true
	}

	return false
}

// SetEndPointRole gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRole field.
func (o *IamLdapGroupAllOf) SetEndPointRole(v []IamEndPointRoleRelationship) {
	o.EndPointRole = v
}

// GetLdapPolicy returns the LdapPolicy field value if set, zero value otherwise.
func (o *IamLdapGroupAllOf) GetLdapPolicy() IamLdapPolicyRelationship {
	if o == nil || o.LdapPolicy == nil {
		var ret IamLdapPolicyRelationship
		return ret
	}
	return *o.LdapPolicy
}

// GetLdapPolicyOk returns a tuple with the LdapPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamLdapGroupAllOf) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool) {
	if o == nil || o.LdapPolicy == nil {
		return nil, false
	}
	return o.LdapPolicy, true
}

// HasLdapPolicy returns a boolean if a field has been set.
func (o *IamLdapGroupAllOf) HasLdapPolicy() bool {
	if o != nil && o.LdapPolicy != nil {
		return true
	}

	return false
}

// SetLdapPolicy gets a reference to the given IamLdapPolicyRelationship and assigns it to the LdapPolicy field.
func (o *IamLdapGroupAllOf) SetLdapPolicy(v IamLdapPolicyRelationship) {
	o.LdapPolicy = &v
}

func (o IamLdapGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Domain != nil {
		toSerialize["Domain"] = o.Domain
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.EndPointRole != nil {
		toSerialize["EndPointRole"] = o.EndPointRole
	}
	if o.LdapPolicy != nil {
		toSerialize["LdapPolicy"] = o.LdapPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamLdapGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamLdapGroupAllOf := _IamLdapGroupAllOf{}

	if err = json.Unmarshal(bytes, &varIamLdapGroupAllOf); err == nil {
		*o = IamLdapGroupAllOf(varIamLdapGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "EndPointRole")
		delete(additionalProperties, "LdapPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamLdapGroupAllOf struct {
	value *IamLdapGroupAllOf
	isSet bool
}

func (v NullableIamLdapGroupAllOf) Get() *IamLdapGroupAllOf {
	return v.value
}

func (v *NullableIamLdapGroupAllOf) Set(val *IamLdapGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamLdapGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamLdapGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamLdapGroupAllOf(val *IamLdapGroupAllOf) *NullableIamLdapGroupAllOf {
	return &NullableIamLdapGroupAllOf{value: val, isSet: true}
}

func (v NullableIamLdapGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamLdapGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
