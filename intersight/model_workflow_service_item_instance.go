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

// WorkflowServiceItemInstance Service item instance is one instance of a service item based on a service item definition.
type WorkflowServiceItemInstance struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The description for this service item instance.
	Description *string `json:"Description,omitempty"`
	// Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted. * `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * `Okay` - The last action on the service item instance completed and the service item instance is in Okay state. * `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned.
	LastStatus *string `json:"LastStatus,omitempty"`
	// A name of the service item instance. Name of the service item instance must be unique within a type of Service item definition.
	Name *string `json:"Name,omitempty"`
	// Status of the service item instance which controls the actions that can be performed on this instance. * `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * `Okay` - The last action on the service item instance completed and the service item instance is in Okay state. * `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned.
	Status *string `json:"Status,omitempty"`
	// The user identifier which indicates the user that started this workflow.
	UserId                *string                                    `json:"UserId,omitempty"`
	Organization          *OrganizationOrganizationRelationship      `json:"Organization,omitempty"`
	ServiceItemDefinition *WorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemInstance WorkflowServiceItemInstance

// NewWorkflowServiceItemInstance instantiates a new WorkflowServiceItemInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemInstance(classId string, objectType string) *WorkflowServiceItemInstance {
	this := WorkflowServiceItemInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemInstanceWithDefaults instantiates a new WorkflowServiceItemInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemInstanceWithDefaults() *WorkflowServiceItemInstance {
	this := WorkflowServiceItemInstance{}
	var classId string = "workflow.ServiceItemInstance"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemInstance"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemInstance) SetDescription(v string) {
	o.Description = &v
}

// GetLastStatus returns the LastStatus field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetLastStatus() string {
	if o == nil || o.LastStatus == nil {
		var ret string
		return ret
	}
	return *o.LastStatus
}

// GetLastStatusOk returns a tuple with the LastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetLastStatusOk() (*string, bool) {
	if o == nil || o.LastStatus == nil {
		return nil, false
	}
	return o.LastStatus, true
}

// HasLastStatus returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasLastStatus() bool {
	if o != nil && o.LastStatus != nil {
		return true
	}

	return false
}

// SetLastStatus gets a reference to the given string and assigns it to the LastStatus field.
func (o *WorkflowServiceItemInstance) SetLastStatus(v string) {
	o.LastStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemInstance) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowServiceItemInstance) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *WorkflowServiceItemInstance) SetUserId(v string) {
	o.UserId = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *WorkflowServiceItemInstance) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemInstance) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || o.ServiceItemDefinition == nil {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemInstance) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemDefinition == nil {
		return nil, false
	}
	return o.ServiceItemDefinition, true
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemInstance) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given WorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemInstance) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition = &v
}

func (o WorkflowServiceItemInstance) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.LastStatus != nil {
		toSerialize["LastStatus"] = o.LastStatus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UserId != nil {
		toSerialize["UserId"] = o.UserId
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.ServiceItemDefinition != nil {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemInstance) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowServiceItemInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The description for this service item instance.
		Description *string `json:"Description,omitempty"`
		// Last status of the service item instance which will be reverted when an ongoing service item action instance is aborted. * `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * `Okay` - The last action on the service item instance completed and the service item instance is in Okay state. * `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned.
		LastStatus *string `json:"LastStatus,omitempty"`
		// A name of the service item instance. Name of the service item instance must be unique within a type of Service item definition.
		Name *string `json:"Name,omitempty"`
		// Status of the service item instance which controls the actions that can be performed on this instance. * `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state. * `Okay` - The last action on the service item instance completed and the service item instance is in Okay state. * `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned.
		Status *string `json:"Status,omitempty"`
		// The user identifier which indicates the user that started this workflow.
		UserId                *string                                    `json:"UserId,omitempty"`
		Organization          *OrganizationOrganizationRelationship      `json:"Organization,omitempty"`
		ServiceItemDefinition *WorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	}

	varWorkflowServiceItemInstanceWithoutEmbeddedStruct := WorkflowServiceItemInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemInstanceWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemInstance := _WorkflowServiceItemInstance{}
		varWorkflowServiceItemInstance.ClassId = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemInstance.ObjectType = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemInstance.Description = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.Description
		varWorkflowServiceItemInstance.LastStatus = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.LastStatus
		varWorkflowServiceItemInstance.Name = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.Name
		varWorkflowServiceItemInstance.Status = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.Status
		varWorkflowServiceItemInstance.UserId = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.UserId
		varWorkflowServiceItemInstance.Organization = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.Organization
		varWorkflowServiceItemInstance.ServiceItemDefinition = varWorkflowServiceItemInstanceWithoutEmbeddedStruct.ServiceItemDefinition
		*o = WorkflowServiceItemInstance(varWorkflowServiceItemInstance)
	} else {
		return err
	}

	varWorkflowServiceItemInstance := _WorkflowServiceItemInstance{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemInstance)
	if err == nil {
		o.MoBaseMo = varWorkflowServiceItemInstance.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "LastStatus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UserId")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "ServiceItemDefinition")

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

type NullableWorkflowServiceItemInstance struct {
	value *WorkflowServiceItemInstance
	isSet bool
}

func (v NullableWorkflowServiceItemInstance) Get() *WorkflowServiceItemInstance {
	return v.value
}

func (v *NullableWorkflowServiceItemInstance) Set(val *WorkflowServiceItemInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemInstance(val *WorkflowServiceItemInstance) *NullableWorkflowServiceItemInstance {
	return &NullableWorkflowServiceItemInstance{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}