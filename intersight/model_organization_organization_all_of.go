/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// OrganizationOrganizationAllOf Definition of the list of properties defined in 'organization.Organization', excluding properties defined in parent classes.
type OrganizationOrganizationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The informative description about the usage of this organization.
	Description *string `json:"Description,omitempty"`
	// The name of the organization. There can be multiple organizations under an account.
	Name    *string                 `json:"Name,omitempty"`
	Account *IamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to resourceGroup resources.
	ResourceGroups       []ResourceGroupRelationship `json:"ResourceGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationOrganizationAllOf OrganizationOrganizationAllOf

// NewOrganizationOrganizationAllOf instantiates a new OrganizationOrganizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationOrganizationAllOf(classId string, objectType string) *OrganizationOrganizationAllOf {
	this := OrganizationOrganizationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOrganizationOrganizationAllOfWithDefaults instantiates a new OrganizationOrganizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationOrganizationAllOfWithDefaults() *OrganizationOrganizationAllOf {
	this := OrganizationOrganizationAllOf{}
	var classId string = "organization.Organization"
	this.ClassId = classId
	var objectType string = "organization.Organization"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OrganizationOrganizationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OrganizationOrganizationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OrganizationOrganizationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OrganizationOrganizationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OrganizationOrganizationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OrganizationOrganizationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationOrganizationAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationOrganizationAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationOrganizationAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationOrganizationAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationOrganizationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationOrganizationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationOrganizationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationOrganizationAllOf) SetName(v string) {
	o.Name = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *OrganizationOrganizationAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationOrganizationAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *OrganizationOrganizationAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *OrganizationOrganizationAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetResourceGroups returns the ResourceGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationOrganizationAllOf) GetResourceGroups() []ResourceGroupRelationship {
	if o == nil {
		var ret []ResourceGroupRelationship
		return ret
	}
	return o.ResourceGroups
}

// GetResourceGroupsOk returns a tuple with the ResourceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationOrganizationAllOf) GetResourceGroupsOk() (*[]ResourceGroupRelationship, bool) {
	if o == nil || o.ResourceGroups == nil {
		return nil, false
	}
	return &o.ResourceGroups, true
}

// HasResourceGroups returns a boolean if a field has been set.
func (o *OrganizationOrganizationAllOf) HasResourceGroups() bool {
	if o != nil && o.ResourceGroups != nil {
		return true
	}

	return false
}

// SetResourceGroups gets a reference to the given []ResourceGroupRelationship and assigns it to the ResourceGroups field.
func (o *OrganizationOrganizationAllOf) SetResourceGroups(v []ResourceGroupRelationship) {
	o.ResourceGroups = v
}

func (o OrganizationOrganizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.ResourceGroups != nil {
		toSerialize["ResourceGroups"] = o.ResourceGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrganizationOrganizationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOrganizationOrganizationAllOf := _OrganizationOrganizationAllOf{}

	if err = json.Unmarshal(bytes, &varOrganizationOrganizationAllOf); err == nil {
		*o = OrganizationOrganizationAllOf(varOrganizationOrganizationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "ResourceGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationOrganizationAllOf struct {
	value *OrganizationOrganizationAllOf
	isSet bool
}

func (v NullableOrganizationOrganizationAllOf) Get() *OrganizationOrganizationAllOf {
	return v.value
}

func (v *NullableOrganizationOrganizationAllOf) Set(val *OrganizationOrganizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationOrganizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationOrganizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationOrganizationAllOf(val *OrganizationOrganizationAllOf) *NullableOrganizationOrganizationAllOf {
	return &NullableOrganizationOrganizationAllOf{value: val, isSet: true}
}

func (v NullableOrganizationOrganizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationOrganizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
