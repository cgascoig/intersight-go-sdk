/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowSolutionDefinition Solution definition is a collection of actions and associated workflow definition that can be used to deploy a solution.
type WorkflowSolutionDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Solution definition can declare that only one instance can be allowed within the customer account.
	AllowMultipleSolutionInstances *bool `json:"AllowMultipleSolutionInstances,omitempty"`
	// The Cisco Validated Design (CVD) Identifier that this solution provides.
	CvdId *string `json:"CvdId,omitempty"`
	// The flag to indicate that solution instance will be deleted after the completion of decommission action.
	DeleteInstanceOnDecommission *bool `json:"DeleteInstanceOnDecommission,omitempty"`
	// The description for this solution which provides information on what are the pre-requisites to deploy the solution and what features are supported on the solution.
	Description *string `json:"Description,omitempty"`
	// A user friendly short name to identify the solution. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// License entitlement required to run this solution. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
	LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
	// The name for this solution definition. You can have multiple versions of the solution with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).
	Name             *string                `json:"Name,omitempty"`
	OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The version of the solution to support multiple versions.
	Version *int64 `json:"Version,omitempty"`
	// An array of relationships to workflowSolutionActionDefinition resources.
	ActionDefinitions    []WorkflowSolutionActionDefinitionRelationship `json:"ActionDefinitions,omitempty"`
	Catalog              *WorkflowCatalogRelationship                   `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowSolutionDefinition WorkflowSolutionDefinition

// NewWorkflowSolutionDefinition instantiates a new WorkflowSolutionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSolutionDefinition(classId string, objectType string) *WorkflowSolutionDefinition {
	this := WorkflowSolutionDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allowMultipleSolutionInstances bool = true
	this.AllowMultipleSolutionInstances = &allowMultipleSolutionInstances
	var deleteInstanceOnDecommission bool = false
	this.DeleteInstanceOnDecommission = &deleteInstanceOnDecommission
	var version int64 = 1
	this.Version = &version
	return &this
}

// NewWorkflowSolutionDefinitionWithDefaults instantiates a new WorkflowSolutionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSolutionDefinitionWithDefaults() *WorkflowSolutionDefinition {
	this := WorkflowSolutionDefinition{}
	var classId string = "workflow.SolutionDefinition"
	this.ClassId = classId
	var objectType string = "workflow.SolutionDefinition"
	this.ObjectType = objectType
	var allowMultipleSolutionInstances bool = true
	this.AllowMultipleSolutionInstances = &allowMultipleSolutionInstances
	var deleteInstanceOnDecommission bool = false
	this.DeleteInstanceOnDecommission = &deleteInstanceOnDecommission
	var version int64 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSolutionDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSolutionDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSolutionDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSolutionDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllowMultipleSolutionInstances returns the AllowMultipleSolutionInstances field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetAllowMultipleSolutionInstances() bool {
	if o == nil || o.AllowMultipleSolutionInstances == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultipleSolutionInstances
}

// GetAllowMultipleSolutionInstancesOk returns a tuple with the AllowMultipleSolutionInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetAllowMultipleSolutionInstancesOk() (*bool, bool) {
	if o == nil || o.AllowMultipleSolutionInstances == nil {
		return nil, false
	}
	return o.AllowMultipleSolutionInstances, true
}

// HasAllowMultipleSolutionInstances returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasAllowMultipleSolutionInstances() bool {
	if o != nil && o.AllowMultipleSolutionInstances != nil {
		return true
	}

	return false
}

// SetAllowMultipleSolutionInstances gets a reference to the given bool and assigns it to the AllowMultipleSolutionInstances field.
func (o *WorkflowSolutionDefinition) SetAllowMultipleSolutionInstances(v bool) {
	o.AllowMultipleSolutionInstances = &v
}

// GetCvdId returns the CvdId field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetCvdId() string {
	if o == nil || o.CvdId == nil {
		var ret string
		return ret
	}
	return *o.CvdId
}

// GetCvdIdOk returns a tuple with the CvdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetCvdIdOk() (*string, bool) {
	if o == nil || o.CvdId == nil {
		return nil, false
	}
	return o.CvdId, true
}

// HasCvdId returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasCvdId() bool {
	if o != nil && o.CvdId != nil {
		return true
	}

	return false
}

// SetCvdId gets a reference to the given string and assigns it to the CvdId field.
func (o *WorkflowSolutionDefinition) SetCvdId(v string) {
	o.CvdId = &v
}

// GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetDeleteInstanceOnDecommission() bool {
	if o == nil || o.DeleteInstanceOnDecommission == nil {
		var ret bool
		return ret
	}
	return *o.DeleteInstanceOnDecommission
}

// GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetDeleteInstanceOnDecommissionOk() (*bool, bool) {
	if o == nil || o.DeleteInstanceOnDecommission == nil {
		return nil, false
	}
	return o.DeleteInstanceOnDecommission, true
}

// HasDeleteInstanceOnDecommission returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasDeleteInstanceOnDecommission() bool {
	if o != nil && o.DeleteInstanceOnDecommission != nil {
		return true
	}

	return false
}

// SetDeleteInstanceOnDecommission gets a reference to the given bool and assigns it to the DeleteInstanceOnDecommission field.
func (o *WorkflowSolutionDefinition) SetDeleteInstanceOnDecommission(v bool) {
	o.DeleteInstanceOnDecommission = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSolutionDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowSolutionDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetLicenseEntitlement() string {
	if o == nil || o.LicenseEntitlement == nil {
		var ret string
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetLicenseEntitlementOk() (*string, bool) {
	if o == nil || o.LicenseEntitlement == nil {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasLicenseEntitlement() bool {
	if o != nil && o.LicenseEntitlement != nil {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given string and assigns it to the LicenseEntitlement field.
func (o *WorkflowSolutionDefinition) SetLicenseEntitlement(v string) {
	o.LicenseEntitlement = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSolutionDefinition) SetName(v string) {
	o.Name = &v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionDefinition) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionDefinition) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return &o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowSolutionDefinition) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowSolutionDefinition) SetVersion(v int64) {
	o.Version = &v
}

// GetActionDefinitions returns the ActionDefinitions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionDefinition) GetActionDefinitions() []WorkflowSolutionActionDefinitionRelationship {
	if o == nil {
		var ret []WorkflowSolutionActionDefinitionRelationship
		return ret
	}
	return o.ActionDefinitions
}

// GetActionDefinitionsOk returns a tuple with the ActionDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionDefinition) GetActionDefinitionsOk() (*[]WorkflowSolutionActionDefinitionRelationship, bool) {
	if o == nil || o.ActionDefinitions == nil {
		return nil, false
	}
	return &o.ActionDefinitions, true
}

// HasActionDefinitions returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasActionDefinitions() bool {
	if o != nil && o.ActionDefinitions != nil {
		return true
	}

	return false
}

// SetActionDefinitions gets a reference to the given []WorkflowSolutionActionDefinitionRelationship and assigns it to the ActionDefinitions field.
func (o *WorkflowSolutionDefinition) SetActionDefinitions(v []WorkflowSolutionActionDefinitionRelationship) {
	o.ActionDefinitions = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowSolutionDefinition) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowSolutionDefinition) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowSolutionDefinition) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowSolutionDefinition) MarshalJSON() ([]byte, error) {
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
	if o.AllowMultipleSolutionInstances != nil {
		toSerialize["AllowMultipleSolutionInstances"] = o.AllowMultipleSolutionInstances
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

func (o *WorkflowSolutionDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowSolutionDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Solution definition can declare that only one instance can be allowed within the customer account.
		AllowMultipleSolutionInstances *bool `json:"AllowMultipleSolutionInstances,omitempty"`
		// The Cisco Validated Design (CVD) Identifier that this solution provides.
		CvdId *string `json:"CvdId,omitempty"`
		// The flag to indicate that solution instance will be deleted after the completion of decommission action.
		DeleteInstanceOnDecommission *bool `json:"DeleteInstanceOnDecommission,omitempty"`
		// The description for this solution which provides information on what are the pre-requisites to deploy the solution and what features are supported on the solution.
		Description *string `json:"Description,omitempty"`
		// A user friendly short name to identify the solution. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
		Label *string `json:"Label,omitempty"`
		// License entitlement required to run this solution. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
		LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
		// The name for this solution definition. You can have multiple versions of the solution with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_).
		Name             *string                `json:"Name,omitempty"`
		OutputDefinition []WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
		// The version of the solution to support multiple versions.
		Version *int64 `json:"Version,omitempty"`
		// An array of relationships to workflowSolutionActionDefinition resources.
		ActionDefinitions []WorkflowSolutionActionDefinitionRelationship `json:"ActionDefinitions,omitempty"`
		Catalog           *WorkflowCatalogRelationship                   `json:"Catalog,omitempty"`
	}

	varWorkflowSolutionDefinitionWithoutEmbeddedStruct := WorkflowSolutionDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSolutionDefinition := _WorkflowSolutionDefinition{}
		varWorkflowSolutionDefinition.ClassId = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.ClassId
		varWorkflowSolutionDefinition.ObjectType = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.ObjectType
		varWorkflowSolutionDefinition.AllowMultipleSolutionInstances = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.AllowMultipleSolutionInstances
		varWorkflowSolutionDefinition.CvdId = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.CvdId
		varWorkflowSolutionDefinition.DeleteInstanceOnDecommission = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.DeleteInstanceOnDecommission
		varWorkflowSolutionDefinition.Description = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.Description
		varWorkflowSolutionDefinition.Label = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.Label
		varWorkflowSolutionDefinition.LicenseEntitlement = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.LicenseEntitlement
		varWorkflowSolutionDefinition.Name = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.Name
		varWorkflowSolutionDefinition.OutputDefinition = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.OutputDefinition
		varWorkflowSolutionDefinition.Version = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.Version
		varWorkflowSolutionDefinition.ActionDefinitions = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.ActionDefinitions
		varWorkflowSolutionDefinition.Catalog = varWorkflowSolutionDefinitionWithoutEmbeddedStruct.Catalog
		*o = WorkflowSolutionDefinition(varWorkflowSolutionDefinition)
	} else {
		return err
	}

	varWorkflowSolutionDefinition := _WorkflowSolutionDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionDefinition)
	if err == nil {
		o.MoBaseMo = varWorkflowSolutionDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowMultipleSolutionInstances")
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

type NullableWorkflowSolutionDefinition struct {
	value *WorkflowSolutionDefinition
	isSet bool
}

func (v NullableWorkflowSolutionDefinition) Get() *WorkflowSolutionDefinition {
	return v.value
}

func (v *NullableWorkflowSolutionDefinition) Set(val *WorkflowSolutionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSolutionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSolutionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSolutionDefinition(val *WorkflowSolutionDefinition) *NullableWorkflowSolutionDefinition {
	return &NullableWorkflowSolutionDefinition{value: val, isSet: true}
}

func (v NullableWorkflowSolutionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSolutionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
