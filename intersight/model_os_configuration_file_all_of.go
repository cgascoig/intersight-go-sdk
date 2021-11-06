/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// OsConfigurationFileAllOf Definition of the list of properties defined in 'os.ConfigurationFile', excluding properties defined in parent classes.
type OsConfigurationFileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the OS ConfigurationFile.
	Description *string `json:"Description,omitempty"`
	// The content of the entire configuration file is stored as value. The content can either be a static file content or a template content. The template is expected to conform to the golang template syntax. The values from os.Answers properties will be used to populate this template.
	FileContent *string `json:"FileContent,omitempty"`
	// The internal flag is set to true when configuration file is uploaded from OS Install wizard. Internal Configuration files will not be displayed in Answer Management Page.
	Internal *bool `json:"Internal,omitempty"`
	// The name of the OS ConfigurationFile that uniquely identifies the configuration file.
	Name         *string         `json:"Name,omitempty"`
	Placeholders []OsPlaceHolder `json:"Placeholders,omitempty"`
	// An internal property that is used to distinguish between the pre-canned OS configuration file entries and user provided entries.
	Supported *bool                  `json:"Supported,omitempty"`
	Catalog   *OsCatalogRelationship `json:"Catalog,omitempty"`
	// An array of relationships to hclOperatingSystem resources.
	Distributions        []HclOperatingSystemRelationship `json:"Distributions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsConfigurationFileAllOf OsConfigurationFileAllOf

// NewOsConfigurationFileAllOf instantiates a new OsConfigurationFileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsConfigurationFileAllOf(classId string, objectType string) *OsConfigurationFileAllOf {
	this := OsConfigurationFileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var internal bool = false
	this.Internal = &internal
	return &this
}

// NewOsConfigurationFileAllOfWithDefaults instantiates a new OsConfigurationFileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsConfigurationFileAllOfWithDefaults() *OsConfigurationFileAllOf {
	this := OsConfigurationFileAllOf{}
	var classId string = "os.ConfigurationFile"
	this.ClassId = classId
	var objectType string = "os.ConfigurationFile"
	this.ObjectType = objectType
	var internal bool = false
	this.Internal = &internal
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsConfigurationFileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsConfigurationFileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsConfigurationFileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsConfigurationFileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OsConfigurationFileAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetFileContent returns the FileContent field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetFileContent() string {
	if o == nil || o.FileContent == nil {
		var ret string
		return ret
	}
	return *o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetFileContentOk() (*string, bool) {
	if o == nil || o.FileContent == nil {
		return nil, false
	}
	return o.FileContent, true
}

// HasFileContent returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasFileContent() bool {
	if o != nil && o.FileContent != nil {
		return true
	}

	return false
}

// SetFileContent gets a reference to the given string and assigns it to the FileContent field.
func (o *OsConfigurationFileAllOf) SetFileContent(v string) {
	o.FileContent = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *OsConfigurationFileAllOf) SetInternal(v bool) {
	o.Internal = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsConfigurationFileAllOf) SetName(v string) {
	o.Name = &v
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsConfigurationFileAllOf) GetPlaceholders() []OsPlaceHolder {
	if o == nil {
		var ret []OsPlaceHolder
		return ret
	}
	return o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsConfigurationFileAllOf) GetPlaceholdersOk() (*[]OsPlaceHolder, bool) {
	if o == nil || o.Placeholders == nil {
		return nil, false
	}
	return &o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasPlaceholders() bool {
	if o != nil && o.Placeholders != nil {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given []OsPlaceHolder and assigns it to the Placeholders field.
func (o *OsConfigurationFileAllOf) SetPlaceholders(v []OsPlaceHolder) {
	o.Placeholders = v
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetSupportedOk() (*bool, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasSupported() bool {
	if o != nil && o.Supported != nil {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *OsConfigurationFileAllOf) SetSupported(v bool) {
	o.Supported = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *OsConfigurationFileAllOf) GetCatalog() OsCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret OsCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsConfigurationFileAllOf) GetCatalogOk() (*OsCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given OsCatalogRelationship and assigns it to the Catalog field.
func (o *OsConfigurationFileAllOf) SetCatalog(v OsCatalogRelationship) {
	o.Catalog = &v
}

// GetDistributions returns the Distributions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsConfigurationFileAllOf) GetDistributions() []HclOperatingSystemRelationship {
	if o == nil {
		var ret []HclOperatingSystemRelationship
		return ret
	}
	return o.Distributions
}

// GetDistributionsOk returns a tuple with the Distributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsConfigurationFileAllOf) GetDistributionsOk() (*[]HclOperatingSystemRelationship, bool) {
	if o == nil || o.Distributions == nil {
		return nil, false
	}
	return &o.Distributions, true
}

// HasDistributions returns a boolean if a field has been set.
func (o *OsConfigurationFileAllOf) HasDistributions() bool {
	if o != nil && o.Distributions != nil {
		return true
	}

	return false
}

// SetDistributions gets a reference to the given []HclOperatingSystemRelationship and assigns it to the Distributions field.
func (o *OsConfigurationFileAllOf) SetDistributions(v []HclOperatingSystemRelationship) {
	o.Distributions = v
}

func (o OsConfigurationFileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.FileContent != nil {
		toSerialize["FileContent"] = o.FileContent
	}
	if o.Internal != nil {
		toSerialize["Internal"] = o.Internal
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Placeholders != nil {
		toSerialize["Placeholders"] = o.Placeholders
	}
	if o.Supported != nil {
		toSerialize["Supported"] = o.Supported
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Distributions != nil {
		toSerialize["Distributions"] = o.Distributions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsConfigurationFileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsConfigurationFileAllOf := _OsConfigurationFileAllOf{}

	if err = json.Unmarshal(bytes, &varOsConfigurationFileAllOf); err == nil {
		*o = OsConfigurationFileAllOf(varOsConfigurationFileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "FileContent")
		delete(additionalProperties, "Internal")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Placeholders")
		delete(additionalProperties, "Supported")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Distributions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsConfigurationFileAllOf struct {
	value *OsConfigurationFileAllOf
	isSet bool
}

func (v NullableOsConfigurationFileAllOf) Get() *OsConfigurationFileAllOf {
	return v.value
}

func (v *NullableOsConfigurationFileAllOf) Set(val *OsConfigurationFileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsConfigurationFileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsConfigurationFileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsConfigurationFileAllOf(val *OsConfigurationFileAllOf) *NullableOsConfigurationFileAllOf {
	return &NullableOsConfigurationFileAllOf{value: val, isSet: true}
}

func (v NullableOsConfigurationFileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsConfigurationFileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
