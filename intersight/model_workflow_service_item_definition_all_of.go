/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// WorkflowServiceItemDefinitionAllOf Definition of the list of properties defined in 'workflow.ServiceItemDefinition', excluding properties defined in parent classes.
type WorkflowServiceItemDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Service item definition can declare that only one instance can be allowed within the customer account.
	AllowMultipleServiceItemInstances *bool `json:"AllowMultipleServiceItemInstances,omitempty"`
	// The Cisco Validated Design (CVD) Identifier that this service item provides.
	CvdId *string `json:"CvdId,omitempty"`
	// The flag to indicate that service item instance will be deleted after the completion of decommission action.
	DeleteInstanceOnDecommission *bool `json:"DeleteInstanceOnDecommission,omitempty"`
	// The description for this service item which provides information on what are the pre-requisites to deploy the service item and what features are supported on the service item.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the service item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// License entitlement required to run this service item. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
	LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
	// The name for this service item definition. You can have multiple versions of the service item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).
	Name             *string                `json:"Name,omitempty"`
	OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The version of the service item to support multiple versions.
	Version *int64 `json:"Version,omitempty"`
	// An array of relationships to workflowServiceItemActionDefinition resources.
	ActionDefinitions    []WorkflowServiceItemActionDefinitionRelationship `json:"ActionDefinitions,omitempty"`
	Catalog              *WorkflowCatalogRelationship                      `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowServiceItemDefinitionAllOf WorkflowServiceItemDefinitionAllOf

// NewWorkflowServiceItemDefinitionAllOf instantiates a new WorkflowServiceItemDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemDefinitionAllOf(classId string, objectType string) *WorkflowServiceItemDefinitionAllOf {
	this := WorkflowServiceItemDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allowMultipleServiceItemInstances bool = true
	this.AllowMultipleServiceItemInstances = &allowMultipleServiceItemInstances
	var deleteInstanceOnDecommission bool = false
	this.DeleteInstanceOnDecommission = &deleteInstanceOnDecommission
	var version int64 = 1
	this.Version = &version
	return &this
}

// NewWorkflowServiceItemDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemDefinitionAllOfWithDefaults() *WorkflowServiceItemDefinitionAllOf {
	this := WorkflowServiceItemDefinitionAllOf{}
	var classId string = "workflow.ServiceItemDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemDefinition"
	this.ObjectType = objectType
	var allowMultipleServiceItemInstances bool = true
	this.AllowMultipleServiceItemInstances = &allowMultipleServiceItemInstances
	var deleteInstanceOnDecommission bool = false
	this.DeleteInstanceOnDecommission = &deleteInstanceOnDecommission
	var version int64 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllowMultipleServiceItemInstances returns the AllowMultipleServiceItemInstances field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetAllowMultipleServiceItemInstances() bool {
	if o == nil || o.AllowMultipleServiceItemInstances == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultipleServiceItemInstances
}

// GetAllowMultipleServiceItemInstancesOk returns a tuple with the AllowMultipleServiceItemInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetAllowMultipleServiceItemInstancesOk() (*bool, bool) {
	if o == nil || o.AllowMultipleServiceItemInstances == nil {
		return nil, false
	}
	return o.AllowMultipleServiceItemInstances, true
}

// HasAllowMultipleServiceItemInstances returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasAllowMultipleServiceItemInstances() bool {
	if o != nil && o.AllowMultipleServiceItemInstances != nil {
		return true
	}

	return false
}

// SetAllowMultipleServiceItemInstances gets a reference to the given bool and assigns it to the AllowMultipleServiceItemInstances field.
func (o *WorkflowServiceItemDefinitionAllOf) SetAllowMultipleServiceItemInstances(v bool) {
	o.AllowMultipleServiceItemInstances = &v
}

// GetCvdId returns the CvdId field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetCvdId() string {
	if o == nil || o.CvdId == nil {
		var ret string
		return ret
	}
	return *o.CvdId
}

// GetCvdIdOk returns a tuple with the CvdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetCvdIdOk() (*string, bool) {
	if o == nil || o.CvdId == nil {
		return nil, false
	}
	return o.CvdId, true
}

// HasCvdId returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasCvdId() bool {
	if o != nil && o.CvdId != nil {
		return true
	}

	return false
}

// SetCvdId gets a reference to the given string and assigns it to the CvdId field.
func (o *WorkflowServiceItemDefinitionAllOf) SetCvdId(v string) {
	o.CvdId = &v
}

// GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetDeleteInstanceOnDecommission() bool {
	if o == nil || o.DeleteInstanceOnDecommission == nil {
		var ret bool
		return ret
	}
	return *o.DeleteInstanceOnDecommission
}

// GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetDeleteInstanceOnDecommissionOk() (*bool, bool) {
	if o == nil || o.DeleteInstanceOnDecommission == nil {
		return nil, false
	}
	return o.DeleteInstanceOnDecommission, true
}

// HasDeleteInstanceOnDecommission returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasDeleteInstanceOnDecommission() bool {
	if o != nil && o.DeleteInstanceOnDecommission != nil {
		return true
	}

	return false
}

// SetDeleteInstanceOnDecommission gets a reference to the given bool and assigns it to the DeleteInstanceOnDecommission field.
func (o *WorkflowServiceItemDefinitionAllOf) SetDeleteInstanceOnDecommission(v bool) {
	o.DeleteInstanceOnDecommission = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowServiceItemDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetLicenseEntitlement() string {
	if o == nil || o.LicenseEntitlement == nil {
		var ret string
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool) {
	if o == nil || o.LicenseEntitlement == nil {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasLicenseEntitlement() bool {
	if o != nil && o.LicenseEntitlement != nil {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given string and assigns it to the LicenseEntitlement field.
func (o *WorkflowServiceItemDefinitionAllOf) SetLicenseEntitlement(v string) {
	o.LicenseEntitlement = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemDefinitionAllOf) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemDefinitionAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return &o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowServiceItemDefinitionAllOf) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowServiceItemDefinitionAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetActionDefinitions returns the ActionDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemDefinitionAllOf) GetActionDefinitions() []WorkflowServiceItemActionDefinitionRelationship {
	if o == nil {
		var ret []WorkflowServiceItemActionDefinitionRelationship
		return ret
	}
	return o.ActionDefinitions
}

// GetActionDefinitionsOk returns a tuple with the ActionDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemDefinitionAllOf) GetActionDefinitionsOk() (*[]WorkflowServiceItemActionDefinitionRelationship, bool) {
	if o == nil || o.ActionDefinitions == nil {
		return nil, false
	}
	return &o.ActionDefinitions, true
}

// HasActionDefinitions returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasActionDefinitions() bool {
	if o != nil && o.ActionDefinitions != nil {
		return true
	}

	return false
}

// SetActionDefinitions gets a reference to the given []WorkflowServiceItemActionDefinitionRelationship and assigns it to the ActionDefinitions field.
func (o *WorkflowServiceItemDefinitionAllOf) SetActionDefinitions(v []WorkflowServiceItemActionDefinitionRelationship) {
	o.ActionDefinitions = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowServiceItemDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowServiceItemDefinitionAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowServiceItemDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowServiceItemDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllowMultipleServiceItemInstances != nil {
		toSerialize["AllowMultipleServiceItemInstances"] = o.AllowMultipleServiceItemInstances
	}
	if o.CvdId != nil {
		toSerialize["CvdId"] = o.CvdId
	}
	if o.DeleteInstanceOnDecommission != nil {
		toSerialize["DeleteInstanceOnDecommission"] = o.DeleteInstanceOnDecommission
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.LicenseEntitlement != nil {
		toSerialize["LicenseEntitlement"] = o.LicenseEntitlement
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.ActionDefinitions != nil {
		toSerialize["ActionDefinitions"] = o.ActionDefinitions
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemDefinitionAllOf := _WorkflowServiceItemDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemDefinitionAllOf); err == nil {
		*o = WorkflowServiceItemDefinitionAllOf(varWorkflowServiceItemDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowMultipleServiceItemInstances")
		delete(additionalProperties, "CvdId")
		delete(additionalProperties, "DeleteInstanceOnDecommission")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "LicenseEntitlement")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "ActionDefinitions")
		delete(additionalProperties, "Catalog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemDefinitionAllOf struct {
	value *WorkflowServiceItemDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemDefinitionAllOf) Get() *WorkflowServiceItemDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemDefinitionAllOf) Set(val *WorkflowServiceItemDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemDefinitionAllOf(val *WorkflowServiceItemDefinitionAllOf) *NullableWorkflowServiceItemDefinitionAllOf {
	return &NullableWorkflowServiceItemDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
