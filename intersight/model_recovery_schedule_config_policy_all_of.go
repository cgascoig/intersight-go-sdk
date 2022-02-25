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

// RecoveryScheduleConfigPolicyAllOf Definition of the list of properties defined in 'recovery.ScheduleConfigPolicy', excluding properties defined in parent classes.
type RecoveryScheduleConfigPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                         `json:"ObjectType"`
	Schedule   NullableRecoveryBackupSchedule `json:"Schedule,omitempty"`
	// An array of relationships to recoveryBackupProfile resources.
	BackupProfiles       []RecoveryBackupProfileRelationship   `json:"BackupProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryScheduleConfigPolicyAllOf RecoveryScheduleConfigPolicyAllOf

// NewRecoveryScheduleConfigPolicyAllOf instantiates a new RecoveryScheduleConfigPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryScheduleConfigPolicyAllOf(classId string, objectType string) *RecoveryScheduleConfigPolicyAllOf {
	this := RecoveryScheduleConfigPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryScheduleConfigPolicyAllOfWithDefaults instantiates a new RecoveryScheduleConfigPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryScheduleConfigPolicyAllOfWithDefaults() *RecoveryScheduleConfigPolicyAllOf {
	this := RecoveryScheduleConfigPolicyAllOf{}
	var classId string = "recovery.ScheduleConfigPolicy"
	this.ClassId = classId
	var objectType string = "recovery.ScheduleConfigPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryScheduleConfigPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryScheduleConfigPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryScheduleConfigPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryScheduleConfigPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryScheduleConfigPolicyAllOf) GetSchedule() RecoveryBackupSchedule {
	if o == nil || o.Schedule.Get() == nil {
		var ret RecoveryBackupSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryScheduleConfigPolicyAllOf) GetScheduleOk() (*RecoveryBackupSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableRecoveryBackupSchedule and assigns it to the Schedule field.
func (o *RecoveryScheduleConfigPolicyAllOf) SetSchedule(v RecoveryBackupSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *RecoveryScheduleConfigPolicyAllOf) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *RecoveryScheduleConfigPolicyAllOf) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetBackupProfiles returns the BackupProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryScheduleConfigPolicyAllOf) GetBackupProfiles() []RecoveryBackupProfileRelationship {
	if o == nil {
		var ret []RecoveryBackupProfileRelationship
		return ret
	}
	return o.BackupProfiles
}

// GetBackupProfilesOk returns a tuple with the BackupProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryScheduleConfigPolicyAllOf) GetBackupProfilesOk() (*[]RecoveryBackupProfileRelationship, bool) {
	if o == nil || o.BackupProfiles == nil {
		return nil, false
	}
	return &o.BackupProfiles, true
}

// HasBackupProfiles returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) HasBackupProfiles() bool {
	if o != nil && o.BackupProfiles != nil {
		return true
	}

	return false
}

// SetBackupProfiles gets a reference to the given []RecoveryBackupProfileRelationship and assigns it to the BackupProfiles field.
func (o *RecoveryScheduleConfigPolicyAllOf) SetBackupProfiles(v []RecoveryBackupProfileRelationship) {
	o.BackupProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryScheduleConfigPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryScheduleConfigPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryScheduleConfigPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o RecoveryScheduleConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.BackupProfiles != nil {
		toSerialize["BackupProfiles"] = o.BackupProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryScheduleConfigPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryScheduleConfigPolicyAllOf := _RecoveryScheduleConfigPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varRecoveryScheduleConfigPolicyAllOf); err == nil {
		*o = RecoveryScheduleConfigPolicyAllOf(varRecoveryScheduleConfigPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "BackupProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryScheduleConfigPolicyAllOf struct {
	value *RecoveryScheduleConfigPolicyAllOf
	isSet bool
}

func (v NullableRecoveryScheduleConfigPolicyAllOf) Get() *RecoveryScheduleConfigPolicyAllOf {
	return v.value
}

func (v *NullableRecoveryScheduleConfigPolicyAllOf) Set(val *RecoveryScheduleConfigPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryScheduleConfigPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryScheduleConfigPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryScheduleConfigPolicyAllOf(val *RecoveryScheduleConfigPolicyAllOf) *NullableRecoveryScheduleConfigPolicyAllOf {
	return &NullableRecoveryScheduleConfigPolicyAllOf{value: val, isSet: true}
}

func (v NullableRecoveryScheduleConfigPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryScheduleConfigPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
