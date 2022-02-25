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

// SoftwarerepositoryLocalMachineAllOf Definition of the list of properties defined in 'softwarerepository.LocalMachine', excluding properties defined in parent classes.
type SoftwarerepositoryLocalMachineAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When the import action in the file MO is updated with 'GeneratePreSignedDownloadUrl', Intersight returns a pre-signed URL in this property as part of the patch response. The user is expected to subsequently download the file using this URL.
	DownloadUrl *string `json:"DownloadUrl,omitempty"`
	// The chunk size (in bytes) for each part of the file to be uploaded.
	PartSize *int64 `json:"PartSize,omitempty"`
	// When the import action in file MO is updated with 'GeneratePreSignedUploadUrl', Intersight shall return a upload Id in this property as part of the PATCH response.
	UploadId *string `json:"UploadId,omitempty"`
	// When a file MO is created with 'LocalMachine' as the source, Intersight returns a pre-signed URL in this property as part of the POST response. The user is expected to subsequently upload the file content using this URL. Once the upload is completed, the user is expected to patch the uploader object's transfer state to success.
	UploadUrl            *string  `json:"UploadUrl,omitempty"`
	UploadUrls           []string `json:"UploadUrls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryLocalMachineAllOf SoftwarerepositoryLocalMachineAllOf

// NewSoftwarerepositoryLocalMachineAllOf instantiates a new SoftwarerepositoryLocalMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryLocalMachineAllOf(classId string, objectType string) *SoftwarerepositoryLocalMachineAllOf {
	this := SoftwarerepositoryLocalMachineAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryLocalMachineAllOfWithDefaults instantiates a new SoftwarerepositoryLocalMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryLocalMachineAllOfWithDefaults() *SoftwarerepositoryLocalMachineAllOf {
	this := SoftwarerepositoryLocalMachineAllOf{}
	var classId string = "softwarerepository.LocalMachine"
	this.ClassId = classId
	var objectType string = "softwarerepository.LocalMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryLocalMachineAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryLocalMachineAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryLocalMachineAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryLocalMachineAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachineAllOf) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetDownloadUrlOk() (*string, bool) {
	if o == nil || o.DownloadUrl == nil {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl != nil {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *SoftwarerepositoryLocalMachineAllOf) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetPartSize returns the PartSize field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachineAllOf) GetPartSize() int64 {
	if o == nil || o.PartSize == nil {
		var ret int64
		return ret
	}
	return *o.PartSize
}

// GetPartSizeOk returns a tuple with the PartSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetPartSizeOk() (*int64, bool) {
	if o == nil || o.PartSize == nil {
		return nil, false
	}
	return o.PartSize, true
}

// HasPartSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) HasPartSize() bool {
	if o != nil && o.PartSize != nil {
		return true
	}

	return false
}

// SetPartSize gets a reference to the given int64 and assigns it to the PartSize field.
func (o *SoftwarerepositoryLocalMachineAllOf) SetPartSize(v int64) {
	o.PartSize = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadId() bool {
	if o != nil && o.UploadId != nil {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadId(v string) {
	o.UploadId = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrl() string {
	if o == nil || o.UploadUrl == nil {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrlOk() (*string, bool) {
	if o == nil || o.UploadUrl == nil {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadUrl() bool {
	if o != nil && o.UploadUrl != nil {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadUrls returns the UploadUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UploadUrls
}

// GetUploadUrlsOk returns a tuple with the UploadUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwarerepositoryLocalMachineAllOf) GetUploadUrlsOk() (*[]string, bool) {
	if o == nil || o.UploadUrls == nil {
		return nil, false
	}
	return &o.UploadUrls, true
}

// HasUploadUrls returns a boolean if a field has been set.
func (o *SoftwarerepositoryLocalMachineAllOf) HasUploadUrls() bool {
	if o != nil && o.UploadUrls != nil {
		return true
	}

	return false
}

// SetUploadUrls gets a reference to the given []string and assigns it to the UploadUrls field.
func (o *SoftwarerepositoryLocalMachineAllOf) SetUploadUrls(v []string) {
	o.UploadUrls = v
}

func (o SoftwarerepositoryLocalMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DownloadUrl != nil {
		toSerialize["DownloadUrl"] = o.DownloadUrl
	}
	if o.PartSize != nil {
		toSerialize["PartSize"] = o.PartSize
	}
	if o.UploadId != nil {
		toSerialize["UploadId"] = o.UploadId
	}
	if o.UploadUrl != nil {
		toSerialize["UploadUrl"] = o.UploadUrl
	}
	if o.UploadUrls != nil {
		toSerialize["UploadUrls"] = o.UploadUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryLocalMachineAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwarerepositoryLocalMachineAllOf := _SoftwarerepositoryLocalMachineAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwarerepositoryLocalMachineAllOf); err == nil {
		*o = SoftwarerepositoryLocalMachineAllOf(varSoftwarerepositoryLocalMachineAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DownloadUrl")
		delete(additionalProperties, "PartSize")
		delete(additionalProperties, "UploadId")
		delete(additionalProperties, "UploadUrl")
		delete(additionalProperties, "UploadUrls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwarerepositoryLocalMachineAllOf struct {
	value *SoftwarerepositoryLocalMachineAllOf
	isSet bool
}

func (v NullableSoftwarerepositoryLocalMachineAllOf) Get() *SoftwarerepositoryLocalMachineAllOf {
	return v.value
}

func (v *NullableSoftwarerepositoryLocalMachineAllOf) Set(val *SoftwarerepositoryLocalMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryLocalMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryLocalMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryLocalMachineAllOf(val *SoftwarerepositoryLocalMachineAllOf) *NullableSoftwarerepositoryLocalMachineAllOf {
	return &NullableSoftwarerepositoryLocalMachineAllOf{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryLocalMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryLocalMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
