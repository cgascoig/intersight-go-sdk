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
	"reflect"
	"strings"
)

// SoftwarerepositoryAuthorization User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.
type SoftwarerepositoryAuthorization struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Indicates whether the value of the 'userId' property has been set.
	IsUserIdSet *bool `json:"IsUserIdSet,omitempty"`
	// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	Password *string `json:"Password,omitempty"`
	// The external repository for which this authorization has been provided. The only supported repository today is cisco.com. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
	RepositoryType *string `json:"RepositoryType,omitempty"`
	// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
	UserId               *string                 `json:"UserId,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryAuthorization SoftwarerepositoryAuthorization

// NewSoftwarerepositoryAuthorization instantiates a new SoftwarerepositoryAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryAuthorization(classId string, objectType string) *SoftwarerepositoryAuthorization {
	this := SoftwarerepositoryAuthorization{}
	this.ClassId = classId
	this.ObjectType = objectType
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// NewSoftwarerepositoryAuthorizationWithDefaults instantiates a new SoftwarerepositoryAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryAuthorizationWithDefaults() *SoftwarerepositoryAuthorization {
	this := SoftwarerepositoryAuthorization{}
	var classId string = "softwarerepository.Authorization"
	this.ClassId = classId
	var objectType string = "softwarerepository.Authorization"
	this.ObjectType = objectType
	var repositoryType string = "Cisco"
	this.RepositoryType = &repositoryType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryAuthorization) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryAuthorization) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryAuthorization) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryAuthorization) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryAuthorization) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsUserIdSet returns the IsUserIdSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetIsUserIdSet() bool {
	if o == nil || o.IsUserIdSet == nil {
		var ret bool
		return ret
	}
	return *o.IsUserIdSet
}

// GetIsUserIdSetOk returns a tuple with the IsUserIdSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetIsUserIdSetOk() (*bool, bool) {
	if o == nil || o.IsUserIdSet == nil {
		return nil, false
	}
	return o.IsUserIdSet, true
}

// HasIsUserIdSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasIsUserIdSet() bool {
	if o != nil && o.IsUserIdSet != nil {
		return true
	}

	return false
}

// SetIsUserIdSet gets a reference to the given bool and assigns it to the IsUserIdSet field.
func (o *SoftwarerepositoryAuthorization) SetIsUserIdSet(v bool) {
	o.IsUserIdSet = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryAuthorization) SetPassword(v string) {
	o.Password = &v
}

// GetRepositoryType returns the RepositoryType field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetRepositoryType() string {
	if o == nil || o.RepositoryType == nil {
		var ret string
		return ret
	}
	return *o.RepositoryType
}

// GetRepositoryTypeOk returns a tuple with the RepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetRepositoryTypeOk() (*string, bool) {
	if o == nil || o.RepositoryType == nil {
		return nil, false
	}
	return o.RepositoryType, true
}

// HasRepositoryType returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasRepositoryType() bool {
	if o != nil && o.RepositoryType != nil {
		return true
	}

	return false
}

// SetRepositoryType gets a reference to the given string and assigns it to the RepositoryType field.
func (o *SoftwarerepositoryAuthorization) SetRepositoryType(v string) {
	o.RepositoryType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SoftwarerepositoryAuthorization) SetUserId(v string) {
	o.UserId = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *SoftwarerepositoryAuthorization) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryAuthorization) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwarerepositoryAuthorization) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *SoftwarerepositoryAuthorization) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o SoftwarerepositoryAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.IsUserIdSet != nil {
		toSerialize["IsUserIdSet"] = o.IsUserIdSet
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.RepositoryType != nil {
		toSerialize["RepositoryType"] = o.RepositoryType
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryAuthorization) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryAuthorizationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Indicates whether the value of the 'userId' property has been set.
		IsUserIdSet *bool `json:"IsUserIdSet,omitempty"`
		// The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
		Password *string `json:"Password,omitempty"`
		// The external repository for which this authorization has been provided. The only supported repository today is cisco.com. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
		RepositoryType *string `json:"RepositoryType,omitempty"`
		// The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf.
		UserId  *string                 `json:"UserId,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
	}

	varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct := SoftwarerepositoryAuthorizationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryAuthorization := _SoftwarerepositoryAuthorization{}
		varSoftwarerepositoryAuthorization.ClassId = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryAuthorization.ObjectType = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryAuthorization.IsPasswordSet = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.IsPasswordSet
		varSoftwarerepositoryAuthorization.IsUserIdSet = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.IsUserIdSet
		varSoftwarerepositoryAuthorization.Password = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.Password
		varSoftwarerepositoryAuthorization.RepositoryType = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.RepositoryType
		varSoftwarerepositoryAuthorization.UserId = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.UserId
		varSoftwarerepositoryAuthorization.Account = varSoftwarerepositoryAuthorizationWithoutEmbeddedStruct.Account
		*o = SoftwarerepositoryAuthorization(varSoftwarerepositoryAuthorization)
	} else {
		return err
	}

	varSoftwarerepositoryAuthorization := _SoftwarerepositoryAuthorization{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryAuthorization)
	if err == nil {
		o.MoBaseMo = varSoftwarerepositoryAuthorization.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "IsUserIdSet")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "RepositoryType")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableSoftwarerepositoryAuthorization struct {
	value *SoftwarerepositoryAuthorization
	isSet bool
}

func (v NullableSoftwarerepositoryAuthorization) Get() *SoftwarerepositoryAuthorization {
	return v.value
}

func (v *NullableSoftwarerepositoryAuthorization) Set(val *SoftwarerepositoryAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryAuthorization(val *SoftwarerepositoryAuthorization) *NullableSoftwarerepositoryAuthorization {
	return &NullableSoftwarerepositoryAuthorization{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
