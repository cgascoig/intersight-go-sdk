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
	"reflect"
	"strings"
)

// CloudTfcWorkspaceVariables The variables associated with a workspace.
type CloudTfcWorkspaceVariables struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Whether this is a Terraform environment variable. Valid values are \"terraform\" or \"env\".
	Category *string `json:"Category,omitempty"`
	// The description of the variable.
	Description *string `json:"Description,omitempty"`
	// Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables.
	Hcl *bool `json:"Hcl,omitempty"`
	// The unique identifier for this variable.
	Identity *string `json:"Identity,omitempty"`
	// The name of the variables.
	Key *string `json:"Key,omitempty"`
	// Whether the value is sensitive. If true then variable is written once and not visible thereafter.
	Sensitive *bool `json:"Sensitive,omitempty"`
	// The value of the variables.
	Value                *string `json:"Value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudTfcWorkspaceVariables CloudTfcWorkspaceVariables

// NewCloudTfcWorkspaceVariables instantiates a new CloudTfcWorkspaceVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTfcWorkspaceVariables(classId string, objectType string) *CloudTfcWorkspaceVariables {
	this := CloudTfcWorkspaceVariables{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudTfcWorkspaceVariablesWithDefaults instantiates a new CloudTfcWorkspaceVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTfcWorkspaceVariablesWithDefaults() *CloudTfcWorkspaceVariables {
	this := CloudTfcWorkspaceVariables{}
	var classId string = "cloud.TfcWorkspaceVariables"
	this.ClassId = classId
	var objectType string = "cloud.TfcWorkspaceVariables"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudTfcWorkspaceVariables) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudTfcWorkspaceVariables) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudTfcWorkspaceVariables) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudTfcWorkspaceVariables) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CloudTfcWorkspaceVariables) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudTfcWorkspaceVariables) SetDescription(v string) {
	o.Description = &v
}

// GetHcl returns the Hcl field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetHcl() bool {
	if o == nil || o.Hcl == nil {
		var ret bool
		return ret
	}
	return *o.Hcl
}

// GetHclOk returns a tuple with the Hcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetHclOk() (*bool, bool) {
	if o == nil || o.Hcl == nil {
		return nil, false
	}
	return o.Hcl, true
}

// HasHcl returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasHcl() bool {
	if o != nil && o.Hcl != nil {
		return true
	}

	return false
}

// SetHcl gets a reference to the given bool and assigns it to the Hcl field.
func (o *CloudTfcWorkspaceVariables) SetHcl(v bool) {
	o.Hcl = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudTfcWorkspaceVariables) SetIdentity(v string) {
	o.Identity = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CloudTfcWorkspaceVariables) SetKey(v string) {
	o.Key = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasSensitive() bool {
	if o != nil && o.Sensitive != nil {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *CloudTfcWorkspaceVariables) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CloudTfcWorkspaceVariables) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTfcWorkspaceVariables) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CloudTfcWorkspaceVariables) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CloudTfcWorkspaceVariables) SetValue(v string) {
	o.Value = &v
}

func (o CloudTfcWorkspaceVariables) MarshalJSON() ([]byte, error) {
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
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Hcl != nil {
		toSerialize["Hcl"] = o.Hcl
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Sensitive != nil {
		toSerialize["Sensitive"] = o.Sensitive
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudTfcWorkspaceVariables) UnmarshalJSON(bytes []byte) (err error) {
	type CloudTfcWorkspaceVariablesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Whether this is a Terraform environment variable. Valid values are \"terraform\" or \"env\".
		Category *string `json:"Category,omitempty"`
		// The description of the variable.
		Description *string `json:"Description,omitempty"`
		// Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables.
		Hcl *bool `json:"Hcl,omitempty"`
		// The unique identifier for this variable.
		Identity *string `json:"Identity,omitempty"`
		// The name of the variables.
		Key *string `json:"Key,omitempty"`
		// Whether the value is sensitive. If true then variable is written once and not visible thereafter.
		Sensitive *bool `json:"Sensitive,omitempty"`
		// The value of the variables.
		Value *string `json:"Value,omitempty"`
	}

	varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct := CloudTfcWorkspaceVariablesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct)
	if err == nil {
		varCloudTfcWorkspaceVariables := _CloudTfcWorkspaceVariables{}
		varCloudTfcWorkspaceVariables.ClassId = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.ClassId
		varCloudTfcWorkspaceVariables.ObjectType = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.ObjectType
		varCloudTfcWorkspaceVariables.Category = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Category
		varCloudTfcWorkspaceVariables.Description = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Description
		varCloudTfcWorkspaceVariables.Hcl = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Hcl
		varCloudTfcWorkspaceVariables.Identity = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Identity
		varCloudTfcWorkspaceVariables.Key = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Key
		varCloudTfcWorkspaceVariables.Sensitive = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Sensitive
		varCloudTfcWorkspaceVariables.Value = varCloudTfcWorkspaceVariablesWithoutEmbeddedStruct.Value
		*o = CloudTfcWorkspaceVariables(varCloudTfcWorkspaceVariables)
	} else {
		return err
	}

	varCloudTfcWorkspaceVariables := _CloudTfcWorkspaceVariables{}

	err = json.Unmarshal(bytes, &varCloudTfcWorkspaceVariables)
	if err == nil {
		o.MoBaseComplexType = varCloudTfcWorkspaceVariables.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Hcl")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Sensitive")
		delete(additionalProperties, "Value")

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

type NullableCloudTfcWorkspaceVariables struct {
	value *CloudTfcWorkspaceVariables
	isSet bool
}

func (v NullableCloudTfcWorkspaceVariables) Get() *CloudTfcWorkspaceVariables {
	return v.value
}

func (v *NullableCloudTfcWorkspaceVariables) Set(val *CloudTfcWorkspaceVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTfcWorkspaceVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTfcWorkspaceVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTfcWorkspaceVariables(val *CloudTfcWorkspaceVariables) *NullableCloudTfcWorkspaceVariables {
	return &NullableCloudTfcWorkspaceVariables{value: val, isSet: true}
}

func (v NullableCloudTfcWorkspaceVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTfcWorkspaceVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
