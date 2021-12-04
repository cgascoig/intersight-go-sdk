/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// PolicyConfigChange Defines the configuration changes at the summary level including configuration changes and disruptions.
type PolicyConfigChange struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string   `json:"ObjectType"`
	Changes              []string `json:"Changes,omitempty"`
	Disruptions          []string `json:"Disruptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigChange PolicyConfigChange

// NewPolicyConfigChange instantiates a new PolicyConfigChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigChange(classId string, objectType string) *PolicyConfigChange {
	this := PolicyConfigChange{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigChangeWithDefaults instantiates a new PolicyConfigChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigChangeWithDefaults() *PolicyConfigChange {
	this := PolicyConfigChange{}
	var classId string = "policy.ConfigChange"
	this.ClassId = classId
	var objectType string = "policy.ConfigChange"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigChange) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChange) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigChange) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigChange) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigChange) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigChange) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChange) GetChanges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChange) GetChangesOk() (*[]string, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return &o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *PolicyConfigChange) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []string and assigns it to the Changes field.
func (o *PolicyConfigChange) SetChanges(v []string) {
	o.Changes = v
}

// GetDisruptions returns the Disruptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigChange) GetDisruptions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disruptions
}

// GetDisruptionsOk returns a tuple with the Disruptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigChange) GetDisruptionsOk() (*[]string, bool) {
	if o == nil || o.Disruptions == nil {
		return nil, false
	}
	return &o.Disruptions, true
}

// HasDisruptions returns a boolean if a field has been set.
func (o *PolicyConfigChange) HasDisruptions() bool {
	if o != nil && o.Disruptions != nil {
		return true
	}

	return false
}

// SetDisruptions gets a reference to the given []string and assigns it to the Disruptions field.
func (o *PolicyConfigChange) SetDisruptions(v []string) {
	o.Disruptions = v
}

func (o PolicyConfigChange) MarshalJSON() ([]byte, error) {
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
	if o.Changes != nil {
		toSerialize["Changes"] = o.Changes
	}
	if o.Disruptions != nil {
		toSerialize["Disruptions"] = o.Disruptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyConfigChange) UnmarshalJSON(bytes []byte) (err error) {
	type PolicyConfigChangeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string   `json:"ObjectType"`
		Changes     []string `json:"Changes,omitempty"`
		Disruptions []string `json:"Disruptions,omitempty"`
	}

	varPolicyConfigChangeWithoutEmbeddedStruct := PolicyConfigChangeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPolicyConfigChangeWithoutEmbeddedStruct)
	if err == nil {
		varPolicyConfigChange := _PolicyConfigChange{}
		varPolicyConfigChange.ClassId = varPolicyConfigChangeWithoutEmbeddedStruct.ClassId
		varPolicyConfigChange.ObjectType = varPolicyConfigChangeWithoutEmbeddedStruct.ObjectType
		varPolicyConfigChange.Changes = varPolicyConfigChangeWithoutEmbeddedStruct.Changes
		varPolicyConfigChange.Disruptions = varPolicyConfigChangeWithoutEmbeddedStruct.Disruptions
		*o = PolicyConfigChange(varPolicyConfigChange)
	} else {
		return err
	}

	varPolicyConfigChange := _PolicyConfigChange{}

	err = json.Unmarshal(bytes, &varPolicyConfigChange)
	if err == nil {
		o.MoBaseComplexType = varPolicyConfigChange.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Changes")
		delete(additionalProperties, "Disruptions")

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

type NullablePolicyConfigChange struct {
	value *PolicyConfigChange
	isSet bool
}

func (v NullablePolicyConfigChange) Get() *PolicyConfigChange {
	return v.value
}

func (v *NullablePolicyConfigChange) Set(val *PolicyConfigChange) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigChange) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigChange(val *PolicyConfigChange) *NullablePolicyConfigChange {
	return &NullablePolicyConfigChange{value: val, isSet: true}
}

func (v NullablePolicyConfigChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
