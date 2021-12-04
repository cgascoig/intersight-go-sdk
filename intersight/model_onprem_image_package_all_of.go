/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// OnpremImagePackageAllOf Definition of the list of properties defined in 'onprem.ImagePackage', excluding properties defined in parent classes.
type OnpremImagePackageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Optional file path of the image package.
	FilePath *string `json:"FilePath,omitempty"`
	// Image file's fingerprint. Fingerprint is calculated using SHA256 algorithm.
	FileSha *string `json:"FileSha,omitempty"`
	// Image file size in bytes.
	FileSize *int64 `json:"FileSize,omitempty"`
	// Image file's last modified date and time.
	FileTime *time.Time `json:"FileTime,omitempty"`
	// Filename of the image package.
	Filename *string `json:"Filename,omitempty"`
	// Name of the software image package.
	Name *string `json:"Name,omitempty"`
	// Image package type (e.g. service, system etc.).
	PackageType *string `json:"PackageType,omitempty"`
	// Image package version string.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnpremImagePackageAllOf OnpremImagePackageAllOf

// NewOnpremImagePackageAllOf instantiates a new OnpremImagePackageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremImagePackageAllOf(classId string, objectType string) *OnpremImagePackageAllOf {
	this := OnpremImagePackageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOnpremImagePackageAllOfWithDefaults instantiates a new OnpremImagePackageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremImagePackageAllOfWithDefaults() *OnpremImagePackageAllOf {
	this := OnpremImagePackageAllOf{}
	var classId string = "onprem.ImagePackage"
	this.ClassId = classId
	var objectType string = "onprem.ImagePackage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OnpremImagePackageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OnpremImagePackageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OnpremImagePackageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OnpremImagePackageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *OnpremImagePackageAllOf) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileSha returns the FileSha field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetFileSha() string {
	if o == nil || o.FileSha == nil {
		var ret string
		return ret
	}
	return *o.FileSha
}

// GetFileShaOk returns a tuple with the FileSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetFileShaOk() (*string, bool) {
	if o == nil || o.FileSha == nil {
		return nil, false
	}
	return o.FileSha, true
}

// HasFileSha returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasFileSha() bool {
	if o != nil && o.FileSha != nil {
		return true
	}

	return false
}

// SetFileSha gets a reference to the given string and assigns it to the FileSha field.
func (o *OnpremImagePackageAllOf) SetFileSha(v string) {
	o.FileSha = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetFileSize() int64 {
	if o == nil || o.FileSize == nil {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetFileSizeOk() (*int64, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *OnpremImagePackageAllOf) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetFileTime returns the FileTime field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetFileTime() time.Time {
	if o == nil || o.FileTime == nil {
		var ret time.Time
		return ret
	}
	return *o.FileTime
}

// GetFileTimeOk returns a tuple with the FileTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetFileTimeOk() (*time.Time, bool) {
	if o == nil || o.FileTime == nil {
		return nil, false
	}
	return o.FileTime, true
}

// HasFileTime returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasFileTime() bool {
	if o != nil && o.FileTime != nil {
		return true
	}

	return false
}

// SetFileTime gets a reference to the given time.Time and assigns it to the FileTime field.
func (o *OnpremImagePackageAllOf) SetFileTime(v time.Time) {
	o.FileTime = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *OnpremImagePackageAllOf) SetFilename(v string) {
	o.Filename = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OnpremImagePackageAllOf) SetName(v string) {
	o.Name = &v
}

// GetPackageType returns the PackageType field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetPackageType() string {
	if o == nil || o.PackageType == nil {
		var ret string
		return ret
	}
	return *o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetPackageTypeOk() (*string, bool) {
	if o == nil || o.PackageType == nil {
		return nil, false
	}
	return o.PackageType, true
}

// HasPackageType returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasPackageType() bool {
	if o != nil && o.PackageType != nil {
		return true
	}

	return false
}

// SetPackageType gets a reference to the given string and assigns it to the PackageType field.
func (o *OnpremImagePackageAllOf) SetPackageType(v string) {
	o.PackageType = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OnpremImagePackageAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremImagePackageAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OnpremImagePackageAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OnpremImagePackageAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o OnpremImagePackageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FileSha != nil {
		toSerialize["FileSha"] = o.FileSha
	}
	if o.FileSize != nil {
		toSerialize["FileSize"] = o.FileSize
	}
	if o.FileTime != nil {
		toSerialize["FileTime"] = o.FileTime
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PackageType != nil {
		toSerialize["PackageType"] = o.PackageType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OnpremImagePackageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOnpremImagePackageAllOf := _OnpremImagePackageAllOf{}

	if err = json.Unmarshal(bytes, &varOnpremImagePackageAllOf); err == nil {
		*o = OnpremImagePackageAllOf(varOnpremImagePackageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FileSha")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "FileTime")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PackageType")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnpremImagePackageAllOf struct {
	value *OnpremImagePackageAllOf
	isSet bool
}

func (v NullableOnpremImagePackageAllOf) Get() *OnpremImagePackageAllOf {
	return v.value
}

func (v *NullableOnpremImagePackageAllOf) Set(val *OnpremImagePackageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremImagePackageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremImagePackageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremImagePackageAllOf(val *OnpremImagePackageAllOf) *NullableOnpremImagePackageAllOf {
	return &NullableOnpremImagePackageAllOf{value: val, isSet: true}
}

func (v NullableOnpremImagePackageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremImagePackageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
