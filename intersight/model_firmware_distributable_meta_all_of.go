/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareDistributableMetaAllOf Definition of the list of properties defined in 'firmware.DistributableMeta', excluding properties defined in parent classes.
type FirmwareDistributableMetaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The S3 bucket name where the images are present, if source is external cloud store.
	BucketName *string `json:"BucketName,omitempty"`
	// The type of distributable image, example huu, scu, driver, os. * `Distributable` - Stores firmware host utility images and fabric images. * `DriverDistributable` - Stores driver distributable images. * `ServerConfigurationUtilityDistributable` - Stores server configuration utility images. * `OperatingSystemFile` - Stores operating system iso images. * `HyperflexDistributable` - It stores HyperFlex images.
	FileType *string `json:"FileType,omitempty"`
	// The version from which user can download images from amazon store, if source is external cloud store.
	FromVersion *string `json:"FromVersion,omitempty"`
	// The mdfid of the image provided by cisco.com.
	Mdfid *string `json:"Mdfid,omitempty"`
	// The software type id provided by cisco.com.
	SoftwareTypeId *string `json:"SoftwareTypeId,omitempty"`
	// The image can be downloaded from cisco.com or external cloud store. * `Cisco` - External repository hosted on cisco.com. * `IntersightCloud` - Repository hosted by the Intersight Cloud. * `LocalMachine` - The file is available on the local client machine. Used as an upload source type. * `NetworkShare` - External repository in the customer datacenter. This will typically be a file server.
	Source          *string  `json:"Source,omitempty"`
	SupportedModels []string `json:"SupportedModels,omitempty"`
	// The version till which user can download images from amazon store, if source is external cloud store.
	ToVersion            *string `json:"ToVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDistributableMetaAllOf FirmwareDistributableMetaAllOf

// NewFirmwareDistributableMetaAllOf instantiates a new FirmwareDistributableMetaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDistributableMetaAllOf(classId string, objectType string) *FirmwareDistributableMetaAllOf {
	this := FirmwareDistributableMetaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// NewFirmwareDistributableMetaAllOfWithDefaults instantiates a new FirmwareDistributableMetaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDistributableMetaAllOfWithDefaults() *FirmwareDistributableMetaAllOf {
	this := FirmwareDistributableMetaAllOf{}
	var classId string = "firmware.DistributableMeta"
	this.ClassId = classId
	var objectType string = "firmware.DistributableMeta"
	this.ObjectType = objectType
	var fileType string = "Distributable"
	this.FileType = &fileType
	var source string = "Cisco"
	this.Source = &source
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareDistributableMetaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareDistributableMetaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareDistributableMetaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareDistributableMetaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *FirmwareDistributableMetaAllOf) SetBucketName(v string) {
	o.BucketName = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *FirmwareDistributableMetaAllOf) SetFileType(v string) {
	o.FileType = &v
}

// GetFromVersion returns the FromVersion field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetFromVersion() string {
	if o == nil || o.FromVersion == nil {
		var ret string
		return ret
	}
	return *o.FromVersion
}

// GetFromVersionOk returns a tuple with the FromVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetFromVersionOk() (*string, bool) {
	if o == nil || o.FromVersion == nil {
		return nil, false
	}
	return o.FromVersion, true
}

// HasFromVersion returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasFromVersion() bool {
	if o != nil && o.FromVersion != nil {
		return true
	}

	return false
}

// SetFromVersion gets a reference to the given string and assigns it to the FromVersion field.
func (o *FirmwareDistributableMetaAllOf) SetFromVersion(v string) {
	o.FromVersion = &v
}

// GetMdfid returns the Mdfid field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetMdfid() string {
	if o == nil || o.Mdfid == nil {
		var ret string
		return ret
	}
	return *o.Mdfid
}

// GetMdfidOk returns a tuple with the Mdfid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetMdfidOk() (*string, bool) {
	if o == nil || o.Mdfid == nil {
		return nil, false
	}
	return o.Mdfid, true
}

// HasMdfid returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasMdfid() bool {
	if o != nil && o.Mdfid != nil {
		return true
	}

	return false
}

// SetMdfid gets a reference to the given string and assigns it to the Mdfid field.
func (o *FirmwareDistributableMetaAllOf) SetMdfid(v string) {
	o.Mdfid = &v
}

// GetSoftwareTypeId returns the SoftwareTypeId field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetSoftwareTypeId() string {
	if o == nil || o.SoftwareTypeId == nil {
		var ret string
		return ret
	}
	return *o.SoftwareTypeId
}

// GetSoftwareTypeIdOk returns a tuple with the SoftwareTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetSoftwareTypeIdOk() (*string, bool) {
	if o == nil || o.SoftwareTypeId == nil {
		return nil, false
	}
	return o.SoftwareTypeId, true
}

// HasSoftwareTypeId returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasSoftwareTypeId() bool {
	if o != nil && o.SoftwareTypeId != nil {
		return true
	}

	return false
}

// SetSoftwareTypeId gets a reference to the given string and assigns it to the SoftwareTypeId field.
func (o *FirmwareDistributableMetaAllOf) SetSoftwareTypeId(v string) {
	o.SoftwareTypeId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *FirmwareDistributableMetaAllOf) SetSource(v string) {
	o.Source = &v
}

// GetSupportedModels returns the SupportedModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareDistributableMetaAllOf) GetSupportedModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedModels
}

// GetSupportedModelsOk returns a tuple with the SupportedModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareDistributableMetaAllOf) GetSupportedModelsOk() (*[]string, bool) {
	if o == nil || o.SupportedModels == nil {
		return nil, false
	}
	return &o.SupportedModels, true
}

// HasSupportedModels returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasSupportedModels() bool {
	if o != nil && o.SupportedModels != nil {
		return true
	}

	return false
}

// SetSupportedModels gets a reference to the given []string and assigns it to the SupportedModels field.
func (o *FirmwareDistributableMetaAllOf) SetSupportedModels(v []string) {
	o.SupportedModels = v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *FirmwareDistributableMetaAllOf) GetToVersion() string {
	if o == nil || o.ToVersion == nil {
		var ret string
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDistributableMetaAllOf) GetToVersionOk() (*string, bool) {
	if o == nil || o.ToVersion == nil {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *FirmwareDistributableMetaAllOf) HasToVersion() bool {
	if o != nil && o.ToVersion != nil {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given string and assigns it to the ToVersion field.
func (o *FirmwareDistributableMetaAllOf) SetToVersion(v string) {
	o.ToVersion = &v
}

func (o FirmwareDistributableMetaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BucketName != nil {
		toSerialize["BucketName"] = o.BucketName
	}
	if o.FileType != nil {
		toSerialize["FileType"] = o.FileType
	}
	if o.FromVersion != nil {
		toSerialize["FromVersion"] = o.FromVersion
	}
	if o.Mdfid != nil {
		toSerialize["Mdfid"] = o.Mdfid
	}
	if o.SoftwareTypeId != nil {
		toSerialize["SoftwareTypeId"] = o.SoftwareTypeId
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.SupportedModels != nil {
		toSerialize["SupportedModels"] = o.SupportedModels
	}
	if o.ToVersion != nil {
		toSerialize["ToVersion"] = o.ToVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDistributableMetaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareDistributableMetaAllOf := _FirmwareDistributableMetaAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareDistributableMetaAllOf); err == nil {
		*o = FirmwareDistributableMetaAllOf(varFirmwareDistributableMetaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "FromVersion")
		delete(additionalProperties, "Mdfid")
		delete(additionalProperties, "SoftwareTypeId")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "SupportedModels")
		delete(additionalProperties, "ToVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareDistributableMetaAllOf struct {
	value *FirmwareDistributableMetaAllOf
	isSet bool
}

func (v NullableFirmwareDistributableMetaAllOf) Get() *FirmwareDistributableMetaAllOf {
	return v.value
}

func (v *NullableFirmwareDistributableMetaAllOf) Set(val *FirmwareDistributableMetaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDistributableMetaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDistributableMetaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDistributableMetaAllOf(val *FirmwareDistributableMetaAllOf) *NullableFirmwareDistributableMetaAllOf {
	return &NullableFirmwareDistributableMetaAllOf{value: val, isSet: true}
}

func (v NullableFirmwareDistributableMetaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDistributableMetaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
