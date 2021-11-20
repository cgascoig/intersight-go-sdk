/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4903
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowBuildTaskMetaOwnerAllOf Definition of the list of properties defined in 'workflow.BuildTaskMetaOwner', excluding properties defined in parent classes.
type WorkflowBuildTaskMetaOwnerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The microservice owner responsible for the tasks.
	Service              *string  `json:"Service,omitempty"`
	WorkflowTypes        []string `json:"WorkflowTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBuildTaskMetaOwnerAllOf WorkflowBuildTaskMetaOwnerAllOf

// NewWorkflowBuildTaskMetaOwnerAllOf instantiates a new WorkflowBuildTaskMetaOwnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBuildTaskMetaOwnerAllOf(classId string, objectType string) *WorkflowBuildTaskMetaOwnerAllOf {
	this := WorkflowBuildTaskMetaOwnerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBuildTaskMetaOwnerAllOfWithDefaults instantiates a new WorkflowBuildTaskMetaOwnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBuildTaskMetaOwnerAllOfWithDefaults() *WorkflowBuildTaskMetaOwnerAllOf {
	this := WorkflowBuildTaskMetaOwnerAllOf{}
	var classId string = "workflow.BuildTaskMetaOwner"
	this.ClassId = classId
	var objectType string = "workflow.BuildTaskMetaOwner"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBuildTaskMetaOwnerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBuildTaskMetaOwnerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaOwnerAllOf) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *WorkflowBuildTaskMetaOwnerAllOf) SetService(v string) {
	o.Service = &v
}

// GetWorkflowTypes returns the WorkflowTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetWorkflowTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WorkflowTypes
}

// GetWorkflowTypesOk returns a tuple with the WorkflowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBuildTaskMetaOwnerAllOf) GetWorkflowTypesOk() (*[]string, bool) {
	if o == nil || o.WorkflowTypes == nil {
		return nil, false
	}
	return &o.WorkflowTypes, true
}

// HasWorkflowTypes returns a boolean if a field has been set.
func (o *WorkflowBuildTaskMetaOwnerAllOf) HasWorkflowTypes() bool {
	if o != nil && o.WorkflowTypes != nil {
		return true
	}

	return false
}

// SetWorkflowTypes gets a reference to the given []string and assigns it to the WorkflowTypes field.
func (o *WorkflowBuildTaskMetaOwnerAllOf) SetWorkflowTypes(v []string) {
	o.WorkflowTypes = v
}

func (o WorkflowBuildTaskMetaOwnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Service != nil {
		toSerialize["Service"] = o.Service
	}
	if o.WorkflowTypes != nil {
		toSerialize["WorkflowTypes"] = o.WorkflowTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBuildTaskMetaOwnerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowBuildTaskMetaOwnerAllOf := _WorkflowBuildTaskMetaOwnerAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowBuildTaskMetaOwnerAllOf); err == nil {
		*o = WorkflowBuildTaskMetaOwnerAllOf(varWorkflowBuildTaskMetaOwnerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Service")
		delete(additionalProperties, "WorkflowTypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowBuildTaskMetaOwnerAllOf struct {
	value *WorkflowBuildTaskMetaOwnerAllOf
	isSet bool
}

func (v NullableWorkflowBuildTaskMetaOwnerAllOf) Get() *WorkflowBuildTaskMetaOwnerAllOf {
	return v.value
}

func (v *NullableWorkflowBuildTaskMetaOwnerAllOf) Set(val *WorkflowBuildTaskMetaOwnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBuildTaskMetaOwnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBuildTaskMetaOwnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBuildTaskMetaOwnerAllOf(val *WorkflowBuildTaskMetaOwnerAllOf) *NullableWorkflowBuildTaskMetaOwnerAllOf {
	return &NullableWorkflowBuildTaskMetaOwnerAllOf{value: val, isSet: true}
}

func (v NullableWorkflowBuildTaskMetaOwnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBuildTaskMetaOwnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
