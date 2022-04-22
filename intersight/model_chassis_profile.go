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

// ChassisProfile A profile specifying configuration settings for a chassis.
type ChassisProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                     `json:"ObjectType"`
	ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
	// The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight.
	TargetPlatform    *string                       `json:"TargetPlatform,omitempty"`
	AssignedChassis   *EquipmentChassisRelationship `json:"AssignedChassis,omitempty"`
	AssociatedChassis *EquipmentChassisRelationship `json:"AssociatedChassis,omitempty"`
	// An array of relationships to chassisConfigChangeDetail resources.
	ConfigChangeDetails []ChassisConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
	ConfigResult        *ChassisConfigResultRelationship        `json:"ConfigResult,omitempty"`
	// An array of relationships to chassisIomProfile resources.
	IomProfiles  []ChassisIomProfileRelationship       `json:"IomProfiles,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to workflowWorkflowInfo resources.
	RunningWorkflows     []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisProfile ChassisProfile

// NewChassisProfile instantiates a new ChassisProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisProfile(classId string, objectType string) *ChassisProfile {
	this := ChassisProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var targetPlatform string = "FIAttached"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewChassisProfileWithDefaults instantiates a new ChassisProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisProfileWithDefaults() *ChassisProfile {
	this := ChassisProfile{}
	var classId string = "chassis.Profile"
	this.ClassId = classId
	var objectType string = "chassis.Profile"
	this.ObjectType = objectType
	var targetPlatform string = "FIAttached"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigChanges returns the ConfigChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetConfigChanges() PolicyConfigChange {
	if o == nil || o.ConfigChanges.Get() == nil {
		var ret PolicyConfigChange
		return ret
	}
	return *o.ConfigChanges.Get()
}

// GetConfigChangesOk returns a tuple with the ConfigChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetConfigChangesOk() (*PolicyConfigChange, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigChanges.Get(), o.ConfigChanges.IsSet()
}

// HasConfigChanges returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigChanges() bool {
	if o != nil && o.ConfigChanges.IsSet() {
		return true
	}

	return false
}

// SetConfigChanges gets a reference to the given NullablePolicyConfigChange and assigns it to the ConfigChanges field.
func (o *ChassisProfile) SetConfigChanges(v PolicyConfigChange) {
	o.ConfigChanges.Set(&v)
}

// SetConfigChangesNil sets the value for ConfigChanges to be an explicit nil
func (o *ChassisProfile) SetConfigChangesNil() {
	o.ConfigChanges.Set(nil)
}

// UnsetConfigChanges ensures that no value is present for ConfigChanges, not even an explicit nil
func (o *ChassisProfile) UnsetConfigChanges() {
	o.ConfigChanges.Unset()
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
func (o *ChassisProfile) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *ChassisProfile) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
func (o *ChassisProfile) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetAssignedChassis returns the AssignedChassis field value if set, zero value otherwise.
func (o *ChassisProfile) GetAssignedChassis() EquipmentChassisRelationship {
	if o == nil || o.AssignedChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssignedChassis
}

// GetAssignedChassisOk returns a tuple with the AssignedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetAssignedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.AssignedChassis == nil {
		return nil, false
	}
	return o.AssignedChassis, true
}

// HasAssignedChassis returns a boolean if a field has been set.
func (o *ChassisProfile) HasAssignedChassis() bool {
	if o != nil && o.AssignedChassis != nil {
		return true
	}

	return false
}

// SetAssignedChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the AssignedChassis field.
func (o *ChassisProfile) SetAssignedChassis(v EquipmentChassisRelationship) {
	o.AssignedChassis = &v
}

// GetAssociatedChassis returns the AssociatedChassis field value if set, zero value otherwise.
func (o *ChassisProfile) GetAssociatedChassis() EquipmentChassisRelationship {
	if o == nil || o.AssociatedChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.AssociatedChassis
}

// GetAssociatedChassisOk returns a tuple with the AssociatedChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetAssociatedChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.AssociatedChassis == nil {
		return nil, false
	}
	return o.AssociatedChassis, true
}

// HasAssociatedChassis returns a boolean if a field has been set.
func (o *ChassisProfile) HasAssociatedChassis() bool {
	if o != nil && o.AssociatedChassis != nil {
		return true
	}

	return false
}

// SetAssociatedChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the AssociatedChassis field.
func (o *ChassisProfile) SetAssociatedChassis(v EquipmentChassisRelationship) {
	o.AssociatedChassis = &v
}

// GetConfigChangeDetails returns the ConfigChangeDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetConfigChangeDetails() []ChassisConfigChangeDetailRelationship {
	if o == nil {
		var ret []ChassisConfigChangeDetailRelationship
		return ret
	}
	return o.ConfigChangeDetails
}

// GetConfigChangeDetailsOk returns a tuple with the ConfigChangeDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetConfigChangeDetailsOk() (*[]ChassisConfigChangeDetailRelationship, bool) {
	if o == nil || o.ConfigChangeDetails == nil {
		return nil, false
	}
	return &o.ConfigChangeDetails, true
}

// HasConfigChangeDetails returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigChangeDetails() bool {
	if o != nil && o.ConfigChangeDetails != nil {
		return true
	}

	return false
}

// SetConfigChangeDetails gets a reference to the given []ChassisConfigChangeDetailRelationship and assigns it to the ConfigChangeDetails field.
func (o *ChassisProfile) SetConfigChangeDetails(v []ChassisConfigChangeDetailRelationship) {
	o.ConfigChangeDetails = v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *ChassisProfile) GetConfigResult() ChassisConfigResultRelationship {
	if o == nil || o.ConfigResult == nil {
		var ret ChassisConfigResultRelationship
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetConfigResultOk() (*ChassisConfigResultRelationship, bool) {
	if o == nil || o.ConfigResult == nil {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *ChassisProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult != nil {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given ChassisConfigResultRelationship and assigns it to the ConfigResult field.
func (o *ChassisProfile) SetConfigResult(v ChassisConfigResultRelationship) {
	o.ConfigResult = &v
}

// GetIomProfiles returns the IomProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetIomProfiles() []ChassisIomProfileRelationship {
	if o == nil {
		var ret []ChassisIomProfileRelationship
		return ret
	}
	return o.IomProfiles
}

// GetIomProfilesOk returns a tuple with the IomProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetIomProfilesOk() (*[]ChassisIomProfileRelationship, bool) {
	if o == nil || o.IomProfiles == nil {
		return nil, false
	}
	return &o.IomProfiles, true
}

// HasIomProfiles returns a boolean if a field has been set.
func (o *ChassisProfile) HasIomProfiles() bool {
	if o != nil && o.IomProfiles != nil {
		return true
	}

	return false
}

// SetIomProfiles gets a reference to the given []ChassisIomProfileRelationship and assigns it to the IomProfiles field.
func (o *ChassisProfile) SetIomProfiles(v []ChassisIomProfileRelationship) {
	o.IomProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ChassisProfile) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisProfile) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisProfile) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisProfile) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetRunningWorkflows returns the RunningWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisProfile) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship {
	if o == nil {
		var ret []WorkflowWorkflowInfoRelationship
		return ret
	}
	return o.RunningWorkflows
}

// GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisProfile) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflows == nil {
		return nil, false
	}
	return &o.RunningWorkflows, true
}

// HasRunningWorkflows returns a boolean if a field has been set.
func (o *ChassisProfile) HasRunningWorkflows() bool {
	if o != nil && o.RunningWorkflows != nil {
		return true
	}

	return false
}

// SetRunningWorkflows gets a reference to the given []WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflows field.
func (o *ChassisProfile) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflows = v
}

func (o ChassisProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return []byte{}, errPolicyAbstractConfigProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigChanges.IsSet() {
		toSerialize["ConfigChanges"] = o.ConfigChanges.Get()
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.AssignedChassis != nil {
		toSerialize["AssignedChassis"] = o.AssignedChassis
	}
	if o.AssociatedChassis != nil {
		toSerialize["AssociatedChassis"] = o.AssociatedChassis
	}
	if o.ConfigChangeDetails != nil {
		toSerialize["ConfigChangeDetails"] = o.ConfigChangeDetails
	}
	if o.ConfigResult != nil {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if o.IomProfiles != nil {
		toSerialize["IomProfiles"] = o.IomProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.RunningWorkflows != nil {
		toSerialize["RunningWorkflows"] = o.RunningWorkflows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisProfile) UnmarshalJSON(bytes []byte) (err error) {
	type ChassisProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                     `json:"ObjectType"`
		ConfigChanges NullablePolicyConfigChange `json:"ConfigChanges,omitempty"`
		// The platform for which the chassis profile is applicable. It can either be a chassis that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `FIAttached` - Chassis which are connected to a Fabric Interconnect that is managed by Intersight.
		TargetPlatform    *string                       `json:"TargetPlatform,omitempty"`
		AssignedChassis   *EquipmentChassisRelationship `json:"AssignedChassis,omitempty"`
		AssociatedChassis *EquipmentChassisRelationship `json:"AssociatedChassis,omitempty"`
		// An array of relationships to chassisConfigChangeDetail resources.
		ConfigChangeDetails []ChassisConfigChangeDetailRelationship `json:"ConfigChangeDetails,omitempty"`
		ConfigResult        *ChassisConfigResultRelationship        `json:"ConfigResult,omitempty"`
		// An array of relationships to chassisIomProfile resources.
		IomProfiles  []ChassisIomProfileRelationship       `json:"IomProfiles,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to workflowWorkflowInfo resources.
		RunningWorkflows []WorkflowWorkflowInfoRelationship `json:"RunningWorkflows,omitempty"`
	}

	varChassisProfileWithoutEmbeddedStruct := ChassisProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varChassisProfileWithoutEmbeddedStruct)
	if err == nil {
		varChassisProfile := _ChassisProfile{}
		varChassisProfile.ClassId = varChassisProfileWithoutEmbeddedStruct.ClassId
		varChassisProfile.ObjectType = varChassisProfileWithoutEmbeddedStruct.ObjectType
		varChassisProfile.ConfigChanges = varChassisProfileWithoutEmbeddedStruct.ConfigChanges
		varChassisProfile.TargetPlatform = varChassisProfileWithoutEmbeddedStruct.TargetPlatform
		varChassisProfile.AssignedChassis = varChassisProfileWithoutEmbeddedStruct.AssignedChassis
		varChassisProfile.AssociatedChassis = varChassisProfileWithoutEmbeddedStruct.AssociatedChassis
		varChassisProfile.ConfigChangeDetails = varChassisProfileWithoutEmbeddedStruct.ConfigChangeDetails
		varChassisProfile.ConfigResult = varChassisProfileWithoutEmbeddedStruct.ConfigResult
		varChassisProfile.IomProfiles = varChassisProfileWithoutEmbeddedStruct.IomProfiles
		varChassisProfile.Organization = varChassisProfileWithoutEmbeddedStruct.Organization
		varChassisProfile.RunningWorkflows = varChassisProfileWithoutEmbeddedStruct.RunningWorkflows
		*o = ChassisProfile(varChassisProfile)
	} else {
		return err
	}

	varChassisProfile := _ChassisProfile{}

	err = json.Unmarshal(bytes, &varChassisProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varChassisProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigChanges")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "AssignedChassis")
		delete(additionalProperties, "AssociatedChassis")
		delete(additionalProperties, "ConfigChangeDetails")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "IomProfiles")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "RunningWorkflows")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableChassisProfile struct {
	value *ChassisProfile
	isSet bool
}

func (v NullableChassisProfile) Get() *ChassisProfile {
	return v.value
}

func (v *NullableChassisProfile) Set(val *ChassisProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisProfile(val *ChassisProfile) *NullableChassisProfile {
	return &NullableChassisProfile{value: val, isSet: true}
}

func (v NullableChassisProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
