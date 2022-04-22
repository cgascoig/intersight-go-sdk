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

// IamSystem System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
type IamSystem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to iamEndPointPrivilege resources.
	EndPointPrivileges []IamEndPointPrivilegeRelationship `json:"EndPointPrivileges,omitempty"`
	// An array of relationships to iamEndPointRole resources.
	EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
	Idp           *IamIdpRelationship           `json:"Idp,omitempty"`
	// An array of relationships to iamPrivilegeSet resources.
	PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
	// An array of relationships to iamPrivilege resources.
	Privileges []IamPrivilegeRelationship `json:"Privileges,omitempty"`
	// An array of relationships to iamRole resources.
	Roles                []IamRoleRelationship           `json:"Roles,omitempty"`
	ServiceProvider      *IamServiceProviderRelationship `json:"ServiceProvider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSystem IamSystem

// NewIamSystem instantiates a new IamSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSystem(classId string, objectType string) *IamSystem {
	this := IamSystem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSystemWithDefaults instantiates a new IamSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSystemWithDefaults() *IamSystem {
	this := IamSystem{}
	var classId string = "iam.System"
	this.ClassId = classId
	var objectType string = "iam.System"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSystem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSystem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSystem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSystem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSystem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSystem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEndPointPrivileges returns the EndPointPrivileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSystem) GetEndPointPrivileges() []IamEndPointPrivilegeRelationship {
	if o == nil {
		var ret []IamEndPointPrivilegeRelationship
		return ret
	}
	return o.EndPointPrivileges
}

// GetEndPointPrivilegesOk returns a tuple with the EndPointPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSystem) GetEndPointPrivilegesOk() (*[]IamEndPointPrivilegeRelationship, bool) {
	if o == nil || o.EndPointPrivileges == nil {
		return nil, false
	}
	return &o.EndPointPrivileges, true
}

// HasEndPointPrivileges returns a boolean if a field has been set.
func (o *IamSystem) HasEndPointPrivileges() bool {
	if o != nil && o.EndPointPrivileges != nil {
		return true
	}

	return false
}

// SetEndPointPrivileges gets a reference to the given []IamEndPointPrivilegeRelationship and assigns it to the EndPointPrivileges field.
func (o *IamSystem) SetEndPointPrivileges(v []IamEndPointPrivilegeRelationship) {
	o.EndPointPrivileges = v
}

// GetEndPointRoles returns the EndPointRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSystem) GetEndPointRoles() []IamEndPointRoleRelationship {
	if o == nil {
		var ret []IamEndPointRoleRelationship
		return ret
	}
	return o.EndPointRoles
}

// GetEndPointRolesOk returns a tuple with the EndPointRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSystem) GetEndPointRolesOk() (*[]IamEndPointRoleRelationship, bool) {
	if o == nil || o.EndPointRoles == nil {
		return nil, false
	}
	return &o.EndPointRoles, true
}

// HasEndPointRoles returns a boolean if a field has been set.
func (o *IamSystem) HasEndPointRoles() bool {
	if o != nil && o.EndPointRoles != nil {
		return true
	}

	return false
}

// SetEndPointRoles gets a reference to the given []IamEndPointRoleRelationship and assigns it to the EndPointRoles field.
func (o *IamSystem) SetEndPointRoles(v []IamEndPointRoleRelationship) {
	o.EndPointRoles = v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamSystem) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSystem) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamSystem) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamSystem) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetPrivilegeSets returns the PrivilegeSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSystem) GetPrivilegeSets() []IamPrivilegeSetRelationship {
	if o == nil {
		var ret []IamPrivilegeSetRelationship
		return ret
	}
	return o.PrivilegeSets
}

// GetPrivilegeSetsOk returns a tuple with the PrivilegeSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSystem) GetPrivilegeSetsOk() (*[]IamPrivilegeSetRelationship, bool) {
	if o == nil || o.PrivilegeSets == nil {
		return nil, false
	}
	return &o.PrivilegeSets, true
}

// HasPrivilegeSets returns a boolean if a field has been set.
func (o *IamSystem) HasPrivilegeSets() bool {
	if o != nil && o.PrivilegeSets != nil {
		return true
	}

	return false
}

// SetPrivilegeSets gets a reference to the given []IamPrivilegeSetRelationship and assigns it to the PrivilegeSets field.
func (o *IamSystem) SetPrivilegeSets(v []IamPrivilegeSetRelationship) {
	o.PrivilegeSets = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSystem) GetPrivileges() []IamPrivilegeRelationship {
	if o == nil {
		var ret []IamPrivilegeRelationship
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSystem) GetPrivilegesOk() (*[]IamPrivilegeRelationship, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *IamSystem) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given []IamPrivilegeRelationship and assigns it to the Privileges field.
func (o *IamSystem) SetPrivileges(v []IamPrivilegeRelationship) {
	o.Privileges = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSystem) GetRoles() []IamRoleRelationship {
	if o == nil {
		var ret []IamRoleRelationship
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSystem) GetRolesOk() (*[]IamRoleRelationship, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *IamSystem) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IamRoleRelationship and assigns it to the Roles field.
func (o *IamSystem) SetRoles(v []IamRoleRelationship) {
	o.Roles = v
}

// GetServiceProvider returns the ServiceProvider field value if set, zero value otherwise.
func (o *IamSystem) GetServiceProvider() IamServiceProviderRelationship {
	if o == nil || o.ServiceProvider == nil {
		var ret IamServiceProviderRelationship
		return ret
	}
	return *o.ServiceProvider
}

// GetServiceProviderOk returns a tuple with the ServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSystem) GetServiceProviderOk() (*IamServiceProviderRelationship, bool) {
	if o == nil || o.ServiceProvider == nil {
		return nil, false
	}
	return o.ServiceProvider, true
}

// HasServiceProvider returns a boolean if a field has been set.
func (o *IamSystem) HasServiceProvider() bool {
	if o != nil && o.ServiceProvider != nil {
		return true
	}

	return false
}

// SetServiceProvider gets a reference to the given IamServiceProviderRelationship and assigns it to the ServiceProvider field.
func (o *IamSystem) SetServiceProvider(v IamServiceProviderRelationship) {
	o.ServiceProvider = &v
}

func (o IamSystem) MarshalJSON() ([]byte, error) {
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
	if o.EndPointPrivileges != nil {
		toSerialize["EndPointPrivileges"] = o.EndPointPrivileges
	}
	if o.EndPointRoles != nil {
		toSerialize["EndPointRoles"] = o.EndPointRoles
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.PrivilegeSets != nil {
		toSerialize["PrivilegeSets"] = o.PrivilegeSets
	}
	if o.Privileges != nil {
		toSerialize["Privileges"] = o.Privileges
	}
	if o.Roles != nil {
		toSerialize["Roles"] = o.Roles
	}
	if o.ServiceProvider != nil {
		toSerialize["ServiceProvider"] = o.ServiceProvider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSystem) UnmarshalJSON(bytes []byte) (err error) {
	type IamSystemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// An array of relationships to iamEndPointPrivilege resources.
		EndPointPrivileges []IamEndPointPrivilegeRelationship `json:"EndPointPrivileges,omitempty"`
		// An array of relationships to iamEndPointRole resources.
		EndPointRoles []IamEndPointRoleRelationship `json:"EndPointRoles,omitempty"`
		Idp           *IamIdpRelationship           `json:"Idp,omitempty"`
		// An array of relationships to iamPrivilegeSet resources.
		PrivilegeSets []IamPrivilegeSetRelationship `json:"PrivilegeSets,omitempty"`
		// An array of relationships to iamPrivilege resources.
		Privileges []IamPrivilegeRelationship `json:"Privileges,omitempty"`
		// An array of relationships to iamRole resources.
		Roles           []IamRoleRelationship           `json:"Roles,omitempty"`
		ServiceProvider *IamServiceProviderRelationship `json:"ServiceProvider,omitempty"`
	}

	varIamSystemWithoutEmbeddedStruct := IamSystemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamSystemWithoutEmbeddedStruct)
	if err == nil {
		varIamSystem := _IamSystem{}
		varIamSystem.ClassId = varIamSystemWithoutEmbeddedStruct.ClassId
		varIamSystem.ObjectType = varIamSystemWithoutEmbeddedStruct.ObjectType
		varIamSystem.EndPointPrivileges = varIamSystemWithoutEmbeddedStruct.EndPointPrivileges
		varIamSystem.EndPointRoles = varIamSystemWithoutEmbeddedStruct.EndPointRoles
		varIamSystem.Idp = varIamSystemWithoutEmbeddedStruct.Idp
		varIamSystem.PrivilegeSets = varIamSystemWithoutEmbeddedStruct.PrivilegeSets
		varIamSystem.Privileges = varIamSystemWithoutEmbeddedStruct.Privileges
		varIamSystem.Roles = varIamSystemWithoutEmbeddedStruct.Roles
		varIamSystem.ServiceProvider = varIamSystemWithoutEmbeddedStruct.ServiceProvider
		*o = IamSystem(varIamSystem)
	} else {
		return err
	}

	varIamSystem := _IamSystem{}

	err = json.Unmarshal(bytes, &varIamSystem)
	if err == nil {
		o.MoBaseMo = varIamSystem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EndPointPrivileges")
		delete(additionalProperties, "EndPointRoles")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "PrivilegeSets")
		delete(additionalProperties, "Privileges")
		delete(additionalProperties, "Roles")
		delete(additionalProperties, "ServiceProvider")

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

type NullableIamSystem struct {
	value *IamSystem
	isSet bool
}

func (v NullableIamSystem) Get() *IamSystem {
	return v.value
}

func (v *NullableIamSystem) Set(val *IamSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSystem(val *IamSystem) *NullableIamSystem {
	return &NullableIamSystem{value: val, isSet: true}
}

func (v NullableIamSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
