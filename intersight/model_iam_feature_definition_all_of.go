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

// IamFeatureDefinitionAllOf Definition of the list of properties defined in 'iam.FeatureDefinition', excluding properties defined in parent classes.
type IamFeatureDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The beta feature that will be enabled for specific account. * `IWO` - Intersight Workflow Optimizer. * `Hitachi` - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework. * `KubernetesExtension` - Extension to the IKS and Adopted Clusters. * `NetAppIO` - Support to claim NetApp Storage arrays as IO targets. * `IvsPublicCloud` - Enables virtualization service for public clouds. * `TerraformCloud` - Enables an ability to create Terraform Cloud. * `IWE` - Enables an ability to use Intersight Workload Engine. * `WashingtonEFT` - Support for EFT customers to use Washington firmware images for upgrades. * `Solutions` - Support for managing solutions.
	Feature              *string `json:"Feature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamFeatureDefinitionAllOf IamFeatureDefinitionAllOf

// NewIamFeatureDefinitionAllOf instantiates a new IamFeatureDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamFeatureDefinitionAllOf(classId string, objectType string) *IamFeatureDefinitionAllOf {
	this := IamFeatureDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// NewIamFeatureDefinitionAllOfWithDefaults instantiates a new IamFeatureDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamFeatureDefinitionAllOfWithDefaults() *IamFeatureDefinitionAllOf {
	this := IamFeatureDefinitionAllOf{}
	var classId string = "iam.FeatureDefinition"
	this.ClassId = classId
	var objectType string = "iam.FeatureDefinition"
	this.ObjectType = objectType
	var feature string = "IWO"
	this.Feature = &feature
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamFeatureDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamFeatureDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamFeatureDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamFeatureDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *IamFeatureDefinitionAllOf) GetFeature() string {
	if o == nil || o.Feature == nil {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamFeatureDefinitionAllOf) GetFeatureOk() (*string, bool) {
	if o == nil || o.Feature == nil {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *IamFeatureDefinitionAllOf) HasFeature() bool {
	if o != nil && o.Feature != nil {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *IamFeatureDefinitionAllOf) SetFeature(v string) {
	o.Feature = &v
}

func (o IamFeatureDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Feature != nil {
		toSerialize["Feature"] = o.Feature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamFeatureDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamFeatureDefinitionAllOf := _IamFeatureDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varIamFeatureDefinitionAllOf); err == nil {
		*o = IamFeatureDefinitionAllOf(varIamFeatureDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Feature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamFeatureDefinitionAllOf struct {
	value *IamFeatureDefinitionAllOf
	isSet bool
}

func (v NullableIamFeatureDefinitionAllOf) Get() *IamFeatureDefinitionAllOf {
	return v.value
}

func (v *NullableIamFeatureDefinitionAllOf) Set(val *IamFeatureDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamFeatureDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamFeatureDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamFeatureDefinitionAllOf(val *IamFeatureDefinitionAllOf) *NullableIamFeatureDefinitionAllOf {
	return &NullableIamFeatureDefinitionAllOf{value: val, isSet: true}
}

func (v NullableIamFeatureDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamFeatureDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
