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
)

// WorkflowSubWorkflowTaskAllOf Definition of the list of properties defined in 'workflow.SubWorkflowTask', excluding properties defined in parent classes.
type WorkflowSubWorkflowTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the catalog moid that this task belongs.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used.
	Version *int64 `json:"Version,omitempty"`
	// The resolved referenced workflow definition managed object.
	WorkflowDefinitionId *string `json:"WorkflowDefinitionId,omitempty"`
	// The qualified name of workflow that should be executed as a task.
	WorkflowDefinitionName *string `json:"WorkflowDefinitionName,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _WorkflowSubWorkflowTaskAllOf WorkflowSubWorkflowTaskAllOf

// NewWorkflowSubWorkflowTaskAllOf instantiates a new WorkflowSubWorkflowTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSubWorkflowTaskAllOf(classId string, objectType string) *WorkflowSubWorkflowTaskAllOf {
	this := WorkflowSubWorkflowTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowSubWorkflowTaskAllOfWithDefaults instantiates a new WorkflowSubWorkflowTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSubWorkflowTaskAllOfWithDefaults() *WorkflowSubWorkflowTaskAllOf {
	this := WorkflowSubWorkflowTaskAllOf{}
	var classId string = "workflow.SubWorkflowTask"
	this.ClassId = classId
	var objectType string = "workflow.SubWorkflowTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSubWorkflowTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSubWorkflowTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSubWorkflowTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSubWorkflowTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowSubWorkflowTaskAllOf) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowSubWorkflowTaskAllOf) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowSubWorkflowTaskAllOf) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowSubWorkflowTaskAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowSubWorkflowTaskAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowSubWorkflowTaskAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetWorkflowDefinitionId returns the WorkflowDefinitionId field value if set, zero value otherwise.
func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionId() string {
	if o == nil || o.WorkflowDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowDefinitionId
}

// GetWorkflowDefinitionIdOk returns a tuple with the WorkflowDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionIdOk() (*string, bool) {
	if o == nil || o.WorkflowDefinitionId == nil {
		return nil, false
	}
	return o.WorkflowDefinitionId, true
}

// HasWorkflowDefinitionId returns a boolean if a field has been set.
func (o *WorkflowSubWorkflowTaskAllOf) HasWorkflowDefinitionId() bool {
	if o != nil && o.WorkflowDefinitionId != nil {
		return true
	}

	return false
}

// SetWorkflowDefinitionId gets a reference to the given string and assigns it to the WorkflowDefinitionId field.
func (o *WorkflowSubWorkflowTaskAllOf) SetWorkflowDefinitionId(v string) {
	o.WorkflowDefinitionId = &v
}

// GetWorkflowDefinitionName returns the WorkflowDefinitionName field value if set, zero value otherwise.
func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionName() string {
	if o == nil || o.WorkflowDefinitionName == nil {
		var ret string
		return ret
	}
	return *o.WorkflowDefinitionName
}

// GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSubWorkflowTaskAllOf) GetWorkflowDefinitionNameOk() (*string, bool) {
	if o == nil || o.WorkflowDefinitionName == nil {
		return nil, false
	}
	return o.WorkflowDefinitionName, true
}

// HasWorkflowDefinitionName returns a boolean if a field has been set.
func (o *WorkflowSubWorkflowTaskAllOf) HasWorkflowDefinitionName() bool {
	if o != nil && o.WorkflowDefinitionName != nil {
		return true
	}

	return false
}

// SetWorkflowDefinitionName gets a reference to the given string and assigns it to the WorkflowDefinitionName field.
func (o *WorkflowSubWorkflowTaskAllOf) SetWorkflowDefinitionName(v string) {
	o.WorkflowDefinitionName = &v
}

func (o WorkflowSubWorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WorkflowDefinitionId != nil {
		toSerialize["WorkflowDefinitionId"] = o.WorkflowDefinitionId
	}
	if o.WorkflowDefinitionName != nil {
		toSerialize["WorkflowDefinitionName"] = o.WorkflowDefinitionName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSubWorkflowTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSubWorkflowTaskAllOf := _WorkflowSubWorkflowTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSubWorkflowTaskAllOf); err == nil {
		*o = WorkflowSubWorkflowTaskAllOf(varWorkflowSubWorkflowTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WorkflowDefinitionId")
		delete(additionalProperties, "WorkflowDefinitionName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSubWorkflowTaskAllOf struct {
	value *WorkflowSubWorkflowTaskAllOf
	isSet bool
}

func (v NullableWorkflowSubWorkflowTaskAllOf) Get() *WorkflowSubWorkflowTaskAllOf {
	return v.value
}

func (v *NullableWorkflowSubWorkflowTaskAllOf) Set(val *WorkflowSubWorkflowTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSubWorkflowTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSubWorkflowTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSubWorkflowTaskAllOf(val *WorkflowSubWorkflowTaskAllOf) *NullableWorkflowSubWorkflowTaskAllOf {
	return &NullableWorkflowSubWorkflowTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSubWorkflowTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSubWorkflowTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
