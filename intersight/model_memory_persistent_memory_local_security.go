/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// MemoryPersistentMemoryLocalSecurity Persistent Memory Local Security specification to be applied to the associated servers through this policy.
type MemoryPersistentMemoryLocalSecurity struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Persistent Memory Security state.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'securePassphrase' property has been set.
	IsSecurePassphraseSet *bool `json:"IsSecurePassphraseSet,omitempty"`
	// Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters =, \\u0021, &, \\#, $, %, +, ^, @, _, *, -.
	SecurePassphrase     *string `json:"SecurePassphrase,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPersistentMemoryLocalSecurity MemoryPersistentMemoryLocalSecurity

// NewMemoryPersistentMemoryLocalSecurity instantiates a new MemoryPersistentMemoryLocalSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryLocalSecurity(classId string, objectType string) *MemoryPersistentMemoryLocalSecurity {
	this := MemoryPersistentMemoryLocalSecurity{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewMemoryPersistentMemoryLocalSecurityWithDefaults instantiates a new MemoryPersistentMemoryLocalSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryLocalSecurityWithDefaults() *MemoryPersistentMemoryLocalSecurity {
	this := MemoryPersistentMemoryLocalSecurity{}
	var classId string = "memory.PersistentMemoryLocalSecurity"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryLocalSecurity"
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryLocalSecurity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLocalSecurity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryLocalSecurity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryLocalSecurity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLocalSecurity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryLocalSecurity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLocalSecurity) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLocalSecurity) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLocalSecurity) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MemoryPersistentMemoryLocalSecurity) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLocalSecurity) GetIsSecurePassphraseSet() bool {
	if o == nil || o.IsSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsSecurePassphraseSet
}

// GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLocalSecurity) GetIsSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsSecurePassphraseSet, true
}

// HasIsSecurePassphraseSet returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLocalSecurity) HasIsSecurePassphraseSet() bool {
	if o != nil && o.IsSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsSecurePassphraseSet gets a reference to the given bool and assigns it to the IsSecurePassphraseSet field.
func (o *MemoryPersistentMemoryLocalSecurity) SetIsSecurePassphraseSet(v bool) {
	o.IsSecurePassphraseSet = &v
}

// GetSecurePassphrase returns the SecurePassphrase field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryLocalSecurity) GetSecurePassphrase() string {
	if o == nil || o.SecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.SecurePassphrase
}

// GetSecurePassphraseOk returns a tuple with the SecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryLocalSecurity) GetSecurePassphraseOk() (*string, bool) {
	if o == nil || o.SecurePassphrase == nil {
		return nil, false
	}
	return o.SecurePassphrase, true
}

// HasSecurePassphrase returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryLocalSecurity) HasSecurePassphrase() bool {
	if o != nil && o.SecurePassphrase != nil {
		return true
	}

	return false
}

// SetSecurePassphrase gets a reference to the given string and assigns it to the SecurePassphrase field.
func (o *MemoryPersistentMemoryLocalSecurity) SetSecurePassphrase(v string) {
	o.SecurePassphrase = &v
}

func (o MemoryPersistentMemoryLocalSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IsSecurePassphraseSet != nil {
		toSerialize["IsSecurePassphraseSet"] = o.IsSecurePassphraseSet
	}
	if o.SecurePassphrase != nil {
		toSerialize["SecurePassphrase"] = o.SecurePassphrase
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryPersistentMemoryLocalSecurity) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Persistent Memory Security state.
		Enabled *bool `json:"Enabled,omitempty"`
		// Indicates whether the value of the 'securePassphrase' property has been set.
		IsSecurePassphraseSet *bool `json:"IsSecurePassphraseSet,omitempty"`
		// Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters =, \\u0021, &, \\#, $, %, +, ^, @, _, *, -.
		SecurePassphrase *string `json:"SecurePassphrase,omitempty"`
	}

	varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct := MemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryLocalSecurity := _MemoryPersistentMemoryLocalSecurity{}
		varMemoryPersistentMemoryLocalSecurity.ClassId = varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryLocalSecurity.ObjectType = varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryLocalSecurity.Enabled = varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct.Enabled
		varMemoryPersistentMemoryLocalSecurity.IsSecurePassphraseSet = varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct.IsSecurePassphraseSet
		varMemoryPersistentMemoryLocalSecurity.SecurePassphrase = varMemoryPersistentMemoryLocalSecurityWithoutEmbeddedStruct.SecurePassphrase
		*o = MemoryPersistentMemoryLocalSecurity(varMemoryPersistentMemoryLocalSecurity)
	} else {
		return err
	}

	varMemoryPersistentMemoryLocalSecurity := _MemoryPersistentMemoryLocalSecurity{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryLocalSecurity)
	if err == nil {
		o.MoBaseComplexType = varMemoryPersistentMemoryLocalSecurity.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsSecurePassphraseSet")
		delete(additionalProperties, "SecurePassphrase")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableMemoryPersistentMemoryLocalSecurity struct {
	value *MemoryPersistentMemoryLocalSecurity
	isSet bool
}

func (v NullableMemoryPersistentMemoryLocalSecurity) Get() *MemoryPersistentMemoryLocalSecurity {
	return v.value
}

func (v *NullableMemoryPersistentMemoryLocalSecurity) Set(val *MemoryPersistentMemoryLocalSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryLocalSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryLocalSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryLocalSecurity(val *MemoryPersistentMemoryLocalSecurity) *NullableMemoryPersistentMemoryLocalSecurity {
	return &NullableMemoryPersistentMemoryLocalSecurity{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryLocalSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryLocalSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
