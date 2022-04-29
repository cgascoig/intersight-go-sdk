/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SoftwarerepositoryCategoryMapperAllOf Definition of the list of properties defined in 'softwarerepository.CategoryMapper', excluding properties defined in parent classes.
type SoftwarerepositoryCategoryMapperAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The category of the model series.
	Category *string `json:"Category,omitempty"`
	// The type of distributable image, example huu, scu, driver, os. * `Distributable` - Stores firmware host utility images and fabric images. * `DriverDistributable` - Stores driver distributable images. * `ServerConfigurationUtilityDistributable` - Stores server configuration utility images. * `OperatingSystemFile` - Stores operating system iso images. * `HyperflexDistributable` - It stores HyperFlex images.
	FileType *string `json:"FileType,omitempty"`
	// The type of image based on the endpoint it can upgrade. For example, ucs-c420m5-huu-3.2.1a.iso can upgrade standalone servers, so the image type is Standalone Server.
	ImageType *string `json:"ImageType,omitempty"`
	// Defines whether NFS firmware upgrade is supported with this image type.
	IsNfsUpgradeSupported *bool `json:"IsNfsUpgradeSupported,omitempty"`
	// Cisco software repository image category identifier.
	MdfId *string `json:"MdfId,omitempty"`
	// The regex that all images of this category follow.
	RegexPattern *string `json:"RegexPattern,omitempty"`
	// The image can be downloaded from cisco.com or external cloud store. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
	Source          *string  `json:"Source,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The software type id provided by cisco.com.
	SwId     *string  `json:"SwId,omitempty"`
	TagTypes []string `json:"TagTypes,omitempty"`
	// The version from which user can download images from amazon store, if source is external cloud store.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryCategoryMapperAllOf SoftwarerepositoryCategoryMapperAllOf

// NewSoftwarerepositoryCategoryMapperAllOf instantiates a new SoftwarerepositoryCategoryMapperAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryCategoryMapperAllOf(classId string, objectType string) *SoftwarerepositoryCategoryMapperAllOf {
	this := SoftwarerepositoryCategoryMapperAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// NewSoftwarerepositoryCategoryMapperAllOfWithDefaults instantiates a new SoftwarerepositoryCategoryMapperAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryCategoryMapperAllOfWithDefaults() *SoftwarerepositoryCategoryMapperAllOf {
	this := SoftwarerepositoryCategoryMapperAllOf{}
	var classId string = "softwarerepository.CategoryMapper"
	this.ClassId = classId
	var objectType string = "softwarerepository.CategoryMapper"
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryCategoryMapperAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryCategoryMapperAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryCategoryMapperAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryCategoryMapperAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetCategory(v string) {
	o.Category = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetFileType(v string) {
	o.FileType = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetImageType() string {
	if o == nil || o.ImageType == nil {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetImageTypeOk() (*string, bool) {
	if o == nil || o.ImageType == nil {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasImageType() bool {
	if o != nil && o.ImageType != nil {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetImageType(v string) {
	o.ImageType = &v
}

// GetIsNfsUpgradeSupported returns the IsNfsUpgradeSupported field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetIsNfsUpgradeSupported() bool {
	if o == nil || o.IsNfsUpgradeSupported == nil {
		var ret bool
		return ret
	}
	return *o.IsNfsUpgradeSupported
}

// GetIsNfsUpgradeSupportedOk returns a tuple with the IsNfsUpgradeSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetIsNfsUpgradeSupportedOk() (*bool, bool) {
	if o == nil || o.IsNfsUpgradeSupported == nil {
		return nil, false
	}
	return o.IsNfsUpgradeSupported, true
}

// HasIsNfsUpgradeSupported returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasIsNfsUpgradeSupported() bool {
	if o != nil && o.IsNfsUpgradeSupported != nil {
		return true
	}

	return false
}

// SetIsNfsUpgradeSupported gets a reference to the given bool and assigns it to the IsNfsUpgradeSupported field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetIsNfsUpgradeSupported(v bool) {
	o.IsNfsUpgradeSupported = &v
}

// GetMdfId returns the MdfId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfId() string {
	if o == nil || o.MdfId == nil {
		var ret string
		return ret
	}
	return *o.MdfId
}

// GetMdfIdOk returns a tuple with the MdfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetMdfIdOk() (*string, bool) {
	if o == nil || o.MdfId == nil {
		return nil, false
	}
	return o.MdfId, true
}

// HasMdfId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasMdfId() bool {
	if o != nil && o.MdfId != nil {
		return true
	}

	return false
}

// SetMdfId gets a reference to the given string and assigns it to the MdfId field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetMdfId(v string) {
	o.MdfId = &v
}

// GetRegexPattern returns the RegexPattern field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPattern() string {
	if o == nil || o.RegexPattern == nil {
		var ret string
		return ret
	}
	return *o.RegexPattern
}

// GetRegexPatternOk returns a tuple with the RegexPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetRegexPatternOk() (*string, bool) {
	if o == nil || o.RegexPattern == nil {
		return nil, false
	}
	return o.RegexPattern, true
}

// HasRegexPattern returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasRegexPattern() bool {
	if o != nil && o.RegexPattern != nil {
		return true
	}

	return false
}

// SetRegexPattern gets a reference to the given string and assigns it to the RegexPattern field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetRegexPattern(v string) {
	o.RegexPattern = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSource(v string) {
	o.Source = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetSwId returns the SwId field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwId() string {
	if o == nil || o.SwId == nil {
		var ret string
		return ret
	}
	return *o.SwId
}

// GetSwIdOk returns a tuple with the SwId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetSwIdOk() (*string, bool) {
	if o == nil || o.SwId == nil {
		return nil, false
	}
	return o.SwId, true
}

// HasSwId returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasSwId() bool {
	if o != nil && o.SwId != nil {
		return true
	}

	return false
}

// SetSwId gets a reference to the given string and assigns it to the SwId field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetSwId(v string) {
	o.SwId = &v
}

// GetTagTypes returns the TagTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagTypes
}

// GetTagTypesOk returns a tuple with the TagTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryCategoryMapperAllOf) GetTagTypesOk() (*[]string, bool) {
	if o == nil || o.TagTypes == nil {
		return nil, false
	}
	return &o.TagTypes, true
}

// HasTagTypes returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasTagTypes() bool {
	if o != nil && o.TagTypes != nil {
		return true
	}

	return false
}

// SetTagTypes gets a reference to the given []string and assigns it to the TagTypes field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetTagTypes(v []string) {
	o.TagTypes = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwarerepositoryCategoryMapperAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwarerepositoryCategoryMapperAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o SoftwarerepositoryCategoryMapperAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}
	if o.ImageType != nil {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.IsNfsUpgradeSupported != nil {
		toSerialize["IsNfsUpgradeSupported"] = o.IsNfsUpgradeSupported
	}
	if o.MdfId != nil {
		toSerialize["MdfId"] = o.MdfId
	}
	if o.RegexPattern != nil {
		toSerialize["RegexPattern"] = o.RegexPattern
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.SwId != nil {
		toSerialize["SwId"] = o.SwId
	}
	if o.TagTypes != nil {
		toSerialize["TagTypes"] = o.TagTypes
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryCategoryMapperAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryCategoryMapperAllOf := _SoftwarerepositoryCategoryMapperAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryCategoryMapperAllOf); err == nil {
		*o = SoftwarerepositoryCategoryMapperAllOf(varSoftwarerepositoryCategoryMapperAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "ImageType")
		delete(additionalProperties, "IsNfsUpgradeSupported")
		delete(additionalProperties, "MdfId")
		delete(additionalProperties, "RegexPattern")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "SwId")
		delete(additionalProperties, "TagTypes")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryCategoryMapperAllOf struct {
	value *SoftwarerepositoryCategoryMapperAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) Get() *SoftwarerepositoryCategoryMapperAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) Set(val *SoftwarerepositoryCategoryMapperAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryCategoryMapperAllOf(val *SoftwarerepositoryCategoryMapperAllOf) *NullableSoftwarerepositoryCategoryMapperAllOf {
	return &NullableSoftwarerepositoryCategoryMapperAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryCategoryMapperAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryCategoryMapperAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
