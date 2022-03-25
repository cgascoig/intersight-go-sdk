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

// VirtualizationCiscoHypervisorManager A hypervisor manager to manage Cisco Intersight Workload Engine compute clusters and is available per user account.
type VirtualizationCiscoHypervisorManager struct {
	VirtualizationBaseHypervisorManager
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                  `json:"ObjectType"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCiscoHypervisorManager VirtualizationCiscoHypervisorManager

// NewVirtualizationCiscoHypervisorManager instantiates a new VirtualizationCiscoHypervisorManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCiscoHypervisorManager(classId string, objectType string) *VirtualizationCiscoHypervisorManager {
	this := VirtualizationCiscoHypervisorManager{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCiscoHypervisorManagerWithDefaults instantiates a new VirtualizationCiscoHypervisorManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCiscoHypervisorManagerWithDefaults() *VirtualizationCiscoHypervisorManager {
	this := VirtualizationCiscoHypervisorManager{}
	var classId string = "virtualization.CiscoHypervisorManager"
	this.ClassId = classId
	var objectType string = "virtualization.CiscoHypervisorManager"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCiscoHypervisorManager) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCiscoHypervisorManager) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCiscoHypervisorManager) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCiscoHypervisorManager) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCiscoHypervisorManager) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCiscoHypervisorManager) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *VirtualizationCiscoHypervisorManager) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCiscoHypervisorManager) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *VirtualizationCiscoHypervisorManager) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *VirtualizationCiscoHypervisorManager) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o VirtualizationCiscoHypervisorManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHypervisorManager, errVirtualizationBaseHypervisorManager := json.Marshal(o.VirtualizationBaseHypervisorManager)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}
	errVirtualizationBaseHypervisorManager = json.Unmarshal([]byte(serializedVirtualizationBaseHypervisorManager), &toSerialize)
	if errVirtualizationBaseHypervisorManager != nil {
		return []byte{}, errVirtualizationBaseHypervisorManager
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCiscoHypervisorManager) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varVirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct := VirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationCiscoHypervisorManager := _VirtualizationCiscoHypervisorManager{}
		varVirtualizationCiscoHypervisorManager.ClassId = varVirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct.ClassId
		varVirtualizationCiscoHypervisorManager.ObjectType = varVirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct.ObjectType
		varVirtualizationCiscoHypervisorManager.Account = varVirtualizationCiscoHypervisorManagerWithoutEmbeddedStruct.Account
		*o = VirtualizationCiscoHypervisorManager(varVirtualizationCiscoHypervisorManager)
	} else {
		return err
	}

	varVirtualizationCiscoHypervisorManager := _VirtualizationCiscoHypervisorManager{}

	err = json.Unmarshal(bytes, &varVirtualizationCiscoHypervisorManager)
	if err == nil {
		o.VirtualizationBaseHypervisorManager = varVirtualizationCiscoHypervisorManager.VirtualizationBaseHypervisorManager
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectVirtualizationBaseHypervisorManager := reflect.ValueOf(o.VirtualizationBaseHypervisorManager)
		for i := 0; i < reflectVirtualizationBaseHypervisorManager.Type().NumField(); i++ {
			t := reflectVirtualizationBaseHypervisorManager.Type().Field(i)

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

type NullableVirtualizationCiscoHypervisorManager struct {
	value *VirtualizationCiscoHypervisorManager
	isSet bool
}

func (v NullableVirtualizationCiscoHypervisorManager) Get() *VirtualizationCiscoHypervisorManager {
	return v.value
}

func (v *NullableVirtualizationCiscoHypervisorManager) Set(val *VirtualizationCiscoHypervisorManager) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCiscoHypervisorManager) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCiscoHypervisorManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCiscoHypervisorManager(val *VirtualizationCiscoHypervisorManager) *NullableVirtualizationCiscoHypervisorManager {
	return &NullableVirtualizationCiscoHypervisorManager{value: val, isSet: true}
}

func (v NullableVirtualizationCiscoHypervisorManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCiscoHypervisorManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
