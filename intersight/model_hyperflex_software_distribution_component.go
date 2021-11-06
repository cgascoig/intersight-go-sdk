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
	"reflect"
	"strings"
)

// HyperflexSoftwareDistributionComponent A HyperFlex Software Distribution Component.
type HyperflexSoftwareDistributionComponent struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The bucket name where the files are present, if source is external cloud store.
	BucketName *string `json:"BucketName,omitempty"`
	// The HyperFlex Software Distribution Component Identifier.
	ComponentId *string `json:"ComponentId,omitempty"`
	// The HyperFlex Software Distribution Component Name.
	ComponentName *string `json:"ComponentName,omitempty"`
	// File location on the cloud storage.
	FilePath        *string  `json:"FilePath,omitempty"`
	FilesToDownload []string `json:"FilesToDownload,omitempty"`
	// The HyperFlex Software Distribution Component Version.
	Version                     *string                                           `json:"Version,omitempty"`
	SoftwareDistributionVersion *HyperflexSoftwareDistributionVersionRelationship `json:"SoftwareDistributionVersion,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _HyperflexSoftwareDistributionComponent HyperflexSoftwareDistributionComponent

// NewHyperflexSoftwareDistributionComponent instantiates a new HyperflexSoftwareDistributionComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareDistributionComponent(classId string, objectType string) *HyperflexSoftwareDistributionComponent {
	this := HyperflexSoftwareDistributionComponent{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareDistributionComponentWithDefaults instantiates a new HyperflexSoftwareDistributionComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareDistributionComponentWithDefaults() *HyperflexSoftwareDistributionComponent {
	this := HyperflexSoftwareDistributionComponent{}
	var classId string = "hyperflex.SoftwareDistributionComponent"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareDistributionComponent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareDistributionComponent) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareDistributionComponent) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareDistributionComponent) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareDistributionComponent) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *HyperflexSoftwareDistributionComponent) SetBucketName(v string) {
	o.BucketName = &v
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetComponentId() string {
	if o == nil || o.ComponentId == nil {
		var ret string
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetComponentIdOk() (*string, bool) {
	if o == nil || o.ComponentId == nil {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasComponentId() bool {
	if o != nil && o.ComponentId != nil {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given string and assigns it to the ComponentId field.
func (o *HyperflexSoftwareDistributionComponent) SetComponentId(v string) {
	o.ComponentId = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasComponentName() bool {
	if o != nil && o.ComponentName != nil {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *HyperflexSoftwareDistributionComponent) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *HyperflexSoftwareDistributionComponent) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFilesToDownload returns the FilesToDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownload() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilesToDownload
}

// GetFilesToDownloadOk returns a tuple with the FilesToDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionComponent) GetFilesToDownloadOk() (*[]string, bool) {
	if o == nil || o.FilesToDownload == nil {
		return nil, false
	}
	return &o.FilesToDownload, true
}

// HasFilesToDownload returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasFilesToDownload() bool {
	if o != nil && o.FilesToDownload != nil {
		return true
	}

	return false
}

// SetFilesToDownload gets a reference to the given []string and assigns it to the FilesToDownload field.
func (o *HyperflexSoftwareDistributionComponent) SetFilesToDownload(v []string) {
	o.FilesToDownload = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexSoftwareDistributionComponent) SetVersion(v string) {
	o.Version = &v
}

// GetSoftwareDistributionVersion returns the SoftwareDistributionVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersion() HyperflexSoftwareDistributionVersionRelationship {
	if o == nil || o.SoftwareDistributionVersion == nil {
		var ret HyperflexSoftwareDistributionVersionRelationship
		return ret
	}
	return *o.SoftwareDistributionVersion
}

// GetSoftwareDistributionVersionOk returns a tuple with the SoftwareDistributionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponent) GetSoftwareDistributionVersionOk() (*HyperflexSoftwareDistributionVersionRelationship, bool) {
	if o == nil || o.SoftwareDistributionVersion == nil {
		return nil, false
	}
	return o.SoftwareDistributionVersion, true
}

// HasSoftwareDistributionVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponent) HasSoftwareDistributionVersion() bool {
	if o != nil && o.SoftwareDistributionVersion != nil {
		return true
	}

	return false
}

// SetSoftwareDistributionVersion gets a reference to the given HyperflexSoftwareDistributionVersionRelationship and assigns it to the SoftwareDistributionVersion field.
func (o *HyperflexSoftwareDistributionComponent) SetSoftwareDistributionVersion(v HyperflexSoftwareDistributionVersionRelationship) {
	o.SoftwareDistributionVersion = &v
}

func (o HyperflexSoftwareDistributionComponent) MarshalJSON() ([]byte, error) {
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
	if o.BucketName != nil {
		toSerialize["BucketName"] = o.BucketName
	}
	if o.ComponentId != nil {
		toSerialize["ComponentId"] = o.ComponentId
	}
	if o.ComponentName != nil {
		toSerialize["ComponentName"] = o.ComponentName
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FilesToDownload != nil {
		toSerialize["FilesToDownload"] = o.FilesToDownload
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.SoftwareDistributionVersion != nil {
		toSerialize["SoftwareDistributionVersion"] = o.SoftwareDistributionVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSoftwareDistributionComponent) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSoftwareDistributionComponentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The bucket name where the files are present, if source is external cloud store.
		BucketName *string `json:"BucketName,omitempty"`
		// The HyperFlex Software Distribution Component Identifier.
		ComponentId *string `json:"ComponentId,omitempty"`
		// The HyperFlex Software Distribution Component Name.
		ComponentName *string `json:"ComponentName,omitempty"`
		// File location on the cloud storage.
		FilePath        *string  `json:"FilePath,omitempty"`
		FilesToDownload []string `json:"FilesToDownload,omitempty"`
		// The HyperFlex Software Distribution Component Version.
		Version                     *string                                           `json:"Version,omitempty"`
		SoftwareDistributionVersion *HyperflexSoftwareDistributionVersionRelationship `json:"SoftwareDistributionVersion,omitempty"`
	}

	varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct := HyperflexSoftwareDistributionComponentWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSoftwareDistributionComponent := _HyperflexSoftwareDistributionComponent{}
		varHyperflexSoftwareDistributionComponent.ClassId = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ClassId
		varHyperflexSoftwareDistributionComponent.ObjectType = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ObjectType
		varHyperflexSoftwareDistributionComponent.BucketName = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.BucketName
		varHyperflexSoftwareDistributionComponent.ComponentId = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ComponentId
		varHyperflexSoftwareDistributionComponent.ComponentName = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.ComponentName
		varHyperflexSoftwareDistributionComponent.FilePath = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.FilePath
		varHyperflexSoftwareDistributionComponent.FilesToDownload = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.FilesToDownload
		varHyperflexSoftwareDistributionComponent.Version = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.Version
		varHyperflexSoftwareDistributionComponent.SoftwareDistributionVersion = varHyperflexSoftwareDistributionComponentWithoutEmbeddedStruct.SoftwareDistributionVersion
		*o = HyperflexSoftwareDistributionComponent(varHyperflexSoftwareDistributionComponent)
	} else {
		return err
	}

	varHyperflexSoftwareDistributionComponent := _HyperflexSoftwareDistributionComponent{}

	err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionComponent)
	if err == nil {
		o.MoBaseMo = varHyperflexSoftwareDistributionComponent.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "ComponentId")
		delete(additionalProperties, "ComponentName")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FilesToDownload")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "SoftwareDistributionVersion")

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

type NullableHyperflexSoftwareDistributionComponent struct {
	value *HyperflexSoftwareDistributionComponent
	isSet bool
}

func (v NullableHyperflexSoftwareDistributionComponent) Get() *HyperflexSoftwareDistributionComponent {
	return v.value
}

func (v *NullableHyperflexSoftwareDistributionComponent) Set(val *HyperflexSoftwareDistributionComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareDistributionComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareDistributionComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareDistributionComponent(val *HyperflexSoftwareDistributionComponent) *NullableHyperflexSoftwareDistributionComponent {
	return &NullableHyperflexSoftwareDistributionComponent{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareDistributionComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareDistributionComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
