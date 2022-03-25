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

// RecoveryRestoreAllOf Definition of the list of properties defined in 'recovery.Restore', excluding properties defined in parent classes.
type RecoveryRestoreAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                  `json:"ObjectType"`
	ConfigParams         NullableRecoveryConfigParams            `json:"ConfigParams,omitempty"`
	BackupInfo           *RecoveryAbstractBackupInfoRelationship `json:"BackupInfo,omitempty"`
	Device               *AssetDeviceRegistrationRelationship    `json:"Device,omitempty"`
	Organization         *OrganizationOrganizationRelationship   `json:"Organization,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship       `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryRestoreAllOf RecoveryRestoreAllOf

// NewRecoveryRestoreAllOf instantiates a new RecoveryRestoreAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryRestoreAllOf(classId string, objectType string) *RecoveryRestoreAllOf {
	this := RecoveryRestoreAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryRestoreAllOfWithDefaults instantiates a new RecoveryRestoreAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryRestoreAllOfWithDefaults() *RecoveryRestoreAllOf {
	this := RecoveryRestoreAllOf{}
	var classId string = "recovery.Restore"
	this.ClassId = classId
	var objectType string = "recovery.Restore"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryRestoreAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryRestoreAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryRestoreAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryRestoreAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryRestoreAllOf) GetConfigParams() RecoveryConfigParams {
	if o == nil || o.ConfigParams.Get() == nil {
		var ret RecoveryConfigParams
		return ret
	}
	return *o.ConfigParams.Get()
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryRestoreAllOf) GetConfigParamsOk() (*RecoveryConfigParams, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigParams.Get(), o.ConfigParams.IsSet()
}

// HasConfigParams returns a boolean if a field has been set.
func (o *RecoveryRestoreAllOf) HasConfigParams() bool {
	if o != nil && o.ConfigParams.IsSet() {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given NullableRecoveryConfigParams and assigns it to the ConfigParams field.
func (o *RecoveryRestoreAllOf) SetConfigParams(v RecoveryConfigParams) {
	o.ConfigParams.Set(&v)
}

// SetConfigParamsNil sets the value for ConfigParams to be an explicit nil
func (o *RecoveryRestoreAllOf) SetConfigParamsNil() {
	o.ConfigParams.Set(nil)
}

// UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
func (o *RecoveryRestoreAllOf) UnsetConfigParams() {
	o.ConfigParams.Unset()
}

// GetBackupInfo returns the BackupInfo field value if set, zero value otherwise.
func (o *RecoveryRestoreAllOf) GetBackupInfo() RecoveryAbstractBackupInfoRelationship {
	if o == nil || o.BackupInfo == nil {
		var ret RecoveryAbstractBackupInfoRelationship
		return ret
	}
	return *o.BackupInfo
}

// GetBackupInfoOk returns a tuple with the BackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetBackupInfoOk() (*RecoveryAbstractBackupInfoRelationship, bool) {
	if o == nil || o.BackupInfo == nil {
		return nil, false
	}
	return o.BackupInfo, true
}

// HasBackupInfo returns a boolean if a field has been set.
func (o *RecoveryRestoreAllOf) HasBackupInfo() bool {
	if o != nil && o.BackupInfo != nil {
		return true
	}

	return false
}

// SetBackupInfo gets a reference to the given RecoveryAbstractBackupInfoRelationship and assigns it to the BackupInfo field.
func (o *RecoveryRestoreAllOf) SetBackupInfo(v RecoveryAbstractBackupInfoRelationship) {
	o.BackupInfo = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *RecoveryRestoreAllOf) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.Device == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *RecoveryRestoreAllOf) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *RecoveryRestoreAllOf) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *RecoveryRestoreAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *RecoveryRestoreAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *RecoveryRestoreAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *RecoveryRestoreAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryRestoreAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *RecoveryRestoreAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *RecoveryRestoreAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o RecoveryRestoreAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigParams.IsSet() {
		toSerialize["ConfigParams"] = o.ConfigParams.Get()
	}
	if o.BackupInfo != nil {
		toSerialize["BackupInfo"] = o.BackupInfo
	}
	if o.Device != nil {
		toSerialize["Device"] = o.Device
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecoveryRestoreAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecoveryRestoreAllOf := _RecoveryRestoreAllOf{}

	if err = json.Unmarshal(bytes, &varRecoveryRestoreAllOf); err == nil {
		*o = RecoveryRestoreAllOf(varRecoveryRestoreAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigParams")
		delete(additionalProperties, "BackupInfo")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecoveryRestoreAllOf struct {
	value *RecoveryRestoreAllOf
	isSet bool
}

func (v NullableRecoveryRestoreAllOf) Get() *RecoveryRestoreAllOf {
	return v.value
}

func (v *NullableRecoveryRestoreAllOf) Set(val *RecoveryRestoreAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryRestoreAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryRestoreAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryRestoreAllOf(val *RecoveryRestoreAllOf) *NullableRecoveryRestoreAllOf {
	return &NullableRecoveryRestoreAllOf{value: val, isSet: true}
}

func (v NullableRecoveryRestoreAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryRestoreAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
