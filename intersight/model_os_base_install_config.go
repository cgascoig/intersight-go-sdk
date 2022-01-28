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

// OsBaseInstallConfig BaseInstallConfig models the configuration required to install OS.
type OsBaseInstallConfig struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType           string            `json:"ObjectType"`
	AdditionalParameters []OsPlaceHolder   `json:"AdditionalParameters,omitempty"`
	Answers              NullableOsAnswers `json:"Answers,omitempty"`
	// User provided description about the OS install configuration.
	Description *string `json:"Description,omitempty"`
	// The failure message of the API.
	ErrorMsg *string `json:"ErrorMsg,omitempty"`
	// The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now. * `vMedia` - OS image is mounted as vMedia in target server for OS installation.
	InstallMethod *string                 `json:"InstallMethod,omitempty"`
	InstallTarget NullableOsInstallTarget `json:"InstallTarget,omitempty"`
	// Denotes API operating status as pending, in_progress, completed_ok, completed_error based on the execution status. * `Pending` - The initial value of the OperStatus. * `InProgress` - The OperStatus value will be InProgress during execution. * `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk. * `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError. * `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.
	OperState                 *string                             `json:"OperState,omitempty"`
	OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _OsBaseInstallConfig OsBaseInstallConfig

// NewOsBaseInstallConfig instantiates a new OsBaseInstallConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsBaseInstallConfig(classId string, objectType string) *OsBaseInstallConfig {
	this := OsBaseInstallConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var installMethod string = "vMedia"
	this.InstallMethod = &installMethod
	return &this
}

// NewOsBaseInstallConfigWithDefaults instantiates a new OsBaseInstallConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsBaseInstallConfigWithDefaults() *OsBaseInstallConfig {
	this := OsBaseInstallConfig{}
	var classId string = "os.Install"
	this.ClassId = classId
	var objectType string = "os.Install"
	this.ObjectType = objectType
	var installMethod string = "vMedia"
	this.InstallMethod = &installMethod
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsBaseInstallConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsBaseInstallConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsBaseInstallConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsBaseInstallConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalParameters returns the AdditionalParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBaseInstallConfig) GetAdditionalParameters() []OsPlaceHolder {
	if o == nil {
		var ret []OsPlaceHolder
		return ret
	}
	return o.AdditionalParameters
}

// GetAdditionalParametersOk returns a tuple with the AdditionalParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBaseInstallConfig) GetAdditionalParametersOk() (*[]OsPlaceHolder, bool) {
	if o == nil || o.AdditionalParameters == nil {
		return nil, false
	}
	return &o.AdditionalParameters, true
}

// HasAdditionalParameters returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasAdditionalParameters() bool {
	if o != nil && o.AdditionalParameters != nil {
		return true
	}

	return false
}

// SetAdditionalParameters gets a reference to the given []OsPlaceHolder and assigns it to the AdditionalParameters field.
func (o *OsBaseInstallConfig) SetAdditionalParameters(v []OsPlaceHolder) {
	o.AdditionalParameters = v
}

// GetAnswers returns the Answers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBaseInstallConfig) GetAnswers() OsAnswers {
	if o == nil || o.Answers.Get() == nil {
		var ret OsAnswers
		return ret
	}
	return *o.Answers.Get()
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBaseInstallConfig) GetAnswersOk() (*OsAnswers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Answers.Get(), o.Answers.IsSet()
}

// HasAnswers returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasAnswers() bool {
	if o != nil && o.Answers.IsSet() {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given NullableOsAnswers and assigns it to the Answers field.
func (o *OsBaseInstallConfig) SetAnswers(v OsAnswers) {
	o.Answers.Set(&v)
}

// SetAnswersNil sets the value for Answers to be an explicit nil
func (o *OsBaseInstallConfig) SetAnswersNil() {
	o.Answers.Set(nil)
}

// UnsetAnswers ensures that no value is present for Answers, not even an explicit nil
func (o *OsBaseInstallConfig) UnsetAnswers() {
	o.Answers.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OsBaseInstallConfig) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OsBaseInstallConfig) SetDescription(v string) {
	o.Description = &v
}

// GetErrorMsg returns the ErrorMsg field value if set, zero value otherwise.
func (o *OsBaseInstallConfig) GetErrorMsg() string {
	if o == nil || o.ErrorMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrorMsg
}

// GetErrorMsgOk returns a tuple with the ErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetErrorMsgOk() (*string, bool) {
	if o == nil || o.ErrorMsg == nil {
		return nil, false
	}
	return o.ErrorMsg, true
}

// HasErrorMsg returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasErrorMsg() bool {
	if o != nil && o.ErrorMsg != nil {
		return true
	}

	return false
}

// SetErrorMsg gets a reference to the given string and assigns it to the ErrorMsg field.
func (o *OsBaseInstallConfig) SetErrorMsg(v string) {
	o.ErrorMsg = &v
}

// GetInstallMethod returns the InstallMethod field value if set, zero value otherwise.
func (o *OsBaseInstallConfig) GetInstallMethod() string {
	if o == nil || o.InstallMethod == nil {
		var ret string
		return ret
	}
	return *o.InstallMethod
}

// GetInstallMethodOk returns a tuple with the InstallMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetInstallMethodOk() (*string, bool) {
	if o == nil || o.InstallMethod == nil {
		return nil, false
	}
	return o.InstallMethod, true
}

// HasInstallMethod returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasInstallMethod() bool {
	if o != nil && o.InstallMethod != nil {
		return true
	}

	return false
}

// SetInstallMethod gets a reference to the given string and assigns it to the InstallMethod field.
func (o *OsBaseInstallConfig) SetInstallMethod(v string) {
	o.InstallMethod = &v
}

// GetInstallTarget returns the InstallTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBaseInstallConfig) GetInstallTarget() OsInstallTarget {
	if o == nil || o.InstallTarget.Get() == nil {
		var ret OsInstallTarget
		return ret
	}
	return *o.InstallTarget.Get()
}

// GetInstallTargetOk returns a tuple with the InstallTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBaseInstallConfig) GetInstallTargetOk() (*OsInstallTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallTarget.Get(), o.InstallTarget.IsSet()
}

// HasInstallTarget returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasInstallTarget() bool {
	if o != nil && o.InstallTarget.IsSet() {
		return true
	}

	return false
}

// SetInstallTarget gets a reference to the given NullableOsInstallTarget and assigns it to the InstallTarget field.
func (o *OsBaseInstallConfig) SetInstallTarget(v OsInstallTarget) {
	o.InstallTarget.Set(&v)
}

// SetInstallTargetNil sets the value for InstallTarget to be an explicit nil
func (o *OsBaseInstallConfig) SetInstallTargetNil() {
	o.InstallTarget.Set(nil)
}

// UnsetInstallTarget ensures that no value is present for InstallTarget, not even an explicit nil
func (o *OsBaseInstallConfig) UnsetInstallTarget() {
	o.InstallTarget.Unset()
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *OsBaseInstallConfig) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsBaseInstallConfig) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *OsBaseInstallConfig) SetOperState(v string) {
	o.OperState = &v
}

// GetOperatingSystemParameters returns the OperatingSystemParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsBaseInstallConfig) GetOperatingSystemParameters() OsOperatingSystemParameters {
	if o == nil || o.OperatingSystemParameters.Get() == nil {
		var ret OsOperatingSystemParameters
		return ret
	}
	return *o.OperatingSystemParameters.Get()
}

// GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsBaseInstallConfig) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystemParameters.Get(), o.OperatingSystemParameters.IsSet()
}

// HasOperatingSystemParameters returns a boolean if a field has been set.
func (o *OsBaseInstallConfig) HasOperatingSystemParameters() bool {
	if o != nil && o.OperatingSystemParameters.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystemParameters gets a reference to the given NullableOsOperatingSystemParameters and assigns it to the OperatingSystemParameters field.
func (o *OsBaseInstallConfig) SetOperatingSystemParameters(v OsOperatingSystemParameters) {
	o.OperatingSystemParameters.Set(&v)
}

// SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil
func (o *OsBaseInstallConfig) SetOperatingSystemParametersNil() {
	o.OperatingSystemParameters.Set(nil)
}

// UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
func (o *OsBaseInstallConfig) UnsetOperatingSystemParameters() {
	o.OperatingSystemParameters.Unset()
}

func (o OsBaseInstallConfig) MarshalJSON() ([]byte, error) {
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
	if o.AdditionalParameters != nil {
		toSerialize["AdditionalParameters"] = o.AdditionalParameters
	}
	if o.Answers.IsSet() {
		toSerialize["Answers"] = o.Answers.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ErrorMsg != nil {
		toSerialize["ErrorMsg"] = o.ErrorMsg
	}
	if o.InstallMethod != nil {
		toSerialize["InstallMethod"] = o.InstallMethod
	}
	if o.InstallTarget.IsSet() {
		toSerialize["InstallTarget"] = o.InstallTarget.Get()
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperatingSystemParameters.IsSet() {
		toSerialize["OperatingSystemParameters"] = o.OperatingSystemParameters.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsBaseInstallConfig) UnmarshalJSON(bytes []byte) (err error) {
	type OsBaseInstallConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType           string            `json:"ObjectType"`
		AdditionalParameters []OsPlaceHolder   `json:"AdditionalParameters,omitempty"`
		Answers              NullableOsAnswers `json:"Answers,omitempty"`
		// User provided description about the OS install configuration.
		Description *string `json:"Description,omitempty"`
		// The failure message of the API.
		ErrorMsg *string `json:"ErrorMsg,omitempty"`
		// The install method to be used for OS installation - vMedia, iPXE.  Only vMedia is supported as of now. * `vMedia` - OS image is mounted as vMedia in target server for OS installation.
		InstallMethod *string                 `json:"InstallMethod,omitempty"`
		InstallTarget NullableOsInstallTarget `json:"InstallTarget,omitempty"`
		// Denotes API operating status as pending, in_progress, completed_ok, completed_error based on the execution status. * `Pending` - The initial value of the OperStatus. * `InProgress` - The OperStatus value will be InProgress during execution. * `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk. * `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError. * `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning.
		OperState                 *string                             `json:"OperState,omitempty"`
		OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	}

	varOsBaseInstallConfigWithoutEmbeddedStruct := OsBaseInstallConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsBaseInstallConfigWithoutEmbeddedStruct)
	if err == nil {
		varOsBaseInstallConfig := _OsBaseInstallConfig{}
		varOsBaseInstallConfig.ClassId = varOsBaseInstallConfigWithoutEmbeddedStruct.ClassId
		varOsBaseInstallConfig.ObjectType = varOsBaseInstallConfigWithoutEmbeddedStruct.ObjectType
		varOsBaseInstallConfig.AdditionalParameters = varOsBaseInstallConfigWithoutEmbeddedStruct.AdditionalParameters
		varOsBaseInstallConfig.Answers = varOsBaseInstallConfigWithoutEmbeddedStruct.Answers
		varOsBaseInstallConfig.Description = varOsBaseInstallConfigWithoutEmbeddedStruct.Description
		varOsBaseInstallConfig.ErrorMsg = varOsBaseInstallConfigWithoutEmbeddedStruct.ErrorMsg
		varOsBaseInstallConfig.InstallMethod = varOsBaseInstallConfigWithoutEmbeddedStruct.InstallMethod
		varOsBaseInstallConfig.InstallTarget = varOsBaseInstallConfigWithoutEmbeddedStruct.InstallTarget
		varOsBaseInstallConfig.OperState = varOsBaseInstallConfigWithoutEmbeddedStruct.OperState
		varOsBaseInstallConfig.OperatingSystemParameters = varOsBaseInstallConfigWithoutEmbeddedStruct.OperatingSystemParameters
		*o = OsBaseInstallConfig(varOsBaseInstallConfig)
	} else {
		return err
	}

	varOsBaseInstallConfig := _OsBaseInstallConfig{}

	err = json.Unmarshal(bytes, &varOsBaseInstallConfig)
	if err == nil {
		o.MoBaseMo = varOsBaseInstallConfig.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalParameters")
		delete(additionalProperties, "Answers")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ErrorMsg")
		delete(additionalProperties, "InstallMethod")
		delete(additionalProperties, "InstallTarget")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperatingSystemParameters")

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

type NullableOsBaseInstallConfig struct {
	value *OsBaseInstallConfig
	isSet bool
}

func (v NullableOsBaseInstallConfig) Get() *OsBaseInstallConfig {
	return v.value
}

func (v *NullableOsBaseInstallConfig) Set(val *OsBaseInstallConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOsBaseInstallConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOsBaseInstallConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsBaseInstallConfig(val *OsBaseInstallConfig) *NullableOsBaseInstallConfig {
	return &NullableOsBaseInstallConfig{value: val, isSet: true}
}

func (v NullableOsBaseInstallConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsBaseInstallConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
