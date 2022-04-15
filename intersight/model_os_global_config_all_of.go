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

// OsGlobalConfigAllOf Definition of the list of properties defined in 'os.GlobalConfig', excluding properties defined in parent classes.
type OsGlobalConfigAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Configuration file.
	ConfigurationFileName *string `json:"ConfigurationFileName,omitempty"`
	// Configuration source for the OS Installation.
	ConfigurationSource *string `json:"ConfigurationSource,omitempty"`
	// The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.
	InstallMethod *string `json:"InstallMethod,omitempty"`
	// The Prefill install Target Name.
	InstallTargetType         *string                             `json:"InstallTargetType,omitempty"`
	OperatingSystemParameters NullableOsOperatingSystemParameters `json:"OperatingSystemParameters,omitempty"`
	// The Operating System Image name.
	OsImageName *string `json:"OsImageName,omitempty"`
	// The name of the Server Configuration Utilities Image.
	ScuImageName *string `json:"ScuImageName,omitempty"`
	// The Windows OS edition, this property required only for Windows server.
	WindowsEdition       *string `json:"WindowsEdition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsGlobalConfigAllOf OsGlobalConfigAllOf

// NewOsGlobalConfigAllOf instantiates a new OsGlobalConfigAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsGlobalConfigAllOf(classId string, objectType string) *OsGlobalConfigAllOf {
	this := OsGlobalConfigAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsGlobalConfigAllOfWithDefaults instantiates a new OsGlobalConfigAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsGlobalConfigAllOfWithDefaults() *OsGlobalConfigAllOf {
	this := OsGlobalConfigAllOf{}
	var classId string = "os.GlobalConfig"
	this.ClassId = classId
	var objectType string = "os.GlobalConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsGlobalConfigAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsGlobalConfigAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsGlobalConfigAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsGlobalConfigAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigurationFileName returns the ConfigurationFileName field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetConfigurationFileName() string {
	if o == nil || o.ConfigurationFileName == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationFileName
}

// GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetConfigurationFileNameOk() (*string, bool) {
	if o == nil || o.ConfigurationFileName == nil {
		return nil, false
	}
	return o.ConfigurationFileName, true
}

// HasConfigurationFileName returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasConfigurationFileName() bool {
	if o != nil && o.ConfigurationFileName != nil {
		return true
	}

	return false
}

// SetConfigurationFileName gets a reference to the given string and assigns it to the ConfigurationFileName field.
func (o *OsGlobalConfigAllOf) SetConfigurationFileName(v string) {
	o.ConfigurationFileName = &v
}

// GetConfigurationSource returns the ConfigurationSource field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetConfigurationSource() string {
	if o == nil || o.ConfigurationSource == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationSource
}

// GetConfigurationSourceOk returns a tuple with the ConfigurationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetConfigurationSourceOk() (*string, bool) {
	if o == nil || o.ConfigurationSource == nil {
		return nil, false
	}
	return o.ConfigurationSource, true
}

// HasConfigurationSource returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasConfigurationSource() bool {
	if o != nil && o.ConfigurationSource != nil {
		return true
	}

	return false
}

// SetConfigurationSource gets a reference to the given string and assigns it to the ConfigurationSource field.
func (o *OsGlobalConfigAllOf) SetConfigurationSource(v string) {
	o.ConfigurationSource = &v
}

// GetInstallMethod returns the InstallMethod field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetInstallMethod() string {
	if o == nil || o.InstallMethod == nil {
		var ret string
		return ret
	}
	return *o.InstallMethod
}

// GetInstallMethodOk returns a tuple with the InstallMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetInstallMethodOk() (*string, bool) {
	if o == nil || o.InstallMethod == nil {
		return nil, false
	}
	return o.InstallMethod, true
}

// HasInstallMethod returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasInstallMethod() bool {
	if o != nil && o.InstallMethod != nil {
		return true
	}

	return false
}

// SetInstallMethod gets a reference to the given string and assigns it to the InstallMethod field.
func (o *OsGlobalConfigAllOf) SetInstallMethod(v string) {
	o.InstallMethod = &v
}

// GetInstallTargetType returns the InstallTargetType field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetInstallTargetType() string {
	if o == nil || o.InstallTargetType == nil {
		var ret string
		return ret
	}
	return *o.InstallTargetType
}

// GetInstallTargetTypeOk returns a tuple with the InstallTargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetInstallTargetTypeOk() (*string, bool) {
	if o == nil || o.InstallTargetType == nil {
		return nil, false
	}
	return o.InstallTargetType, true
}

// HasInstallTargetType returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasInstallTargetType() bool {
	if o != nil && o.InstallTargetType != nil {
		return true
	}

	return false
}

// SetInstallTargetType gets a reference to the given string and assigns it to the InstallTargetType field.
func (o *OsGlobalConfigAllOf) SetInstallTargetType(v string) {
	o.InstallTargetType = &v
}

// GetOperatingSystemParameters returns the OperatingSystemParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsGlobalConfigAllOf) GetOperatingSystemParameters() OsOperatingSystemParameters {
	if o == nil || o.OperatingSystemParameters.Get() == nil {
		var ret OsOperatingSystemParameters
		return ret
	}
	return *o.OperatingSystemParameters.Get()
}

// GetOperatingSystemParametersOk returns a tuple with the OperatingSystemParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsGlobalConfigAllOf) GetOperatingSystemParametersOk() (*OsOperatingSystemParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystemParameters.Get(), o.OperatingSystemParameters.IsSet()
}

// HasOperatingSystemParameters returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasOperatingSystemParameters() bool {
	if o != nil && o.OperatingSystemParameters.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystemParameters gets a reference to the given NullableOsOperatingSystemParameters and assigns it to the OperatingSystemParameters field.
func (o *OsGlobalConfigAllOf) SetOperatingSystemParameters(v OsOperatingSystemParameters) {
	o.OperatingSystemParameters.Set(&v)
}

// SetOperatingSystemParametersNil sets the value for OperatingSystemParameters to be an explicit nil
func (o *OsGlobalConfigAllOf) SetOperatingSystemParametersNil() {
	o.OperatingSystemParameters.Set(nil)
}

// UnsetOperatingSystemParameters ensures that no value is present for OperatingSystemParameters, not even an explicit nil
func (o *OsGlobalConfigAllOf) UnsetOperatingSystemParameters() {
	o.OperatingSystemParameters.Unset()
}

// GetOsImageName returns the OsImageName field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetOsImageName() string {
	if o == nil || o.OsImageName == nil {
		var ret string
		return ret
	}
	return *o.OsImageName
}

// GetOsImageNameOk returns a tuple with the OsImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetOsImageNameOk() (*string, bool) {
	if o == nil || o.OsImageName == nil {
		return nil, false
	}
	return o.OsImageName, true
}

// HasOsImageName returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasOsImageName() bool {
	if o != nil && o.OsImageName != nil {
		return true
	}

	return false
}

// SetOsImageName gets a reference to the given string and assigns it to the OsImageName field.
func (o *OsGlobalConfigAllOf) SetOsImageName(v string) {
	o.OsImageName = &v
}

// GetScuImageName returns the ScuImageName field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetScuImageName() string {
	if o == nil || o.ScuImageName == nil {
		var ret string
		return ret
	}
	return *o.ScuImageName
}

// GetScuImageNameOk returns a tuple with the ScuImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetScuImageNameOk() (*string, bool) {
	if o == nil || o.ScuImageName == nil {
		return nil, false
	}
	return o.ScuImageName, true
}

// HasScuImageName returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasScuImageName() bool {
	if o != nil && o.ScuImageName != nil {
		return true
	}

	return false
}

// SetScuImageName gets a reference to the given string and assigns it to the ScuImageName field.
func (o *OsGlobalConfigAllOf) SetScuImageName(v string) {
	o.ScuImageName = &v
}

// GetWindowsEdition returns the WindowsEdition field value if set, zero value otherwise.
func (o *OsGlobalConfigAllOf) GetWindowsEdition() string {
	if o == nil || o.WindowsEdition == nil {
		var ret string
		return ret
	}
	return *o.WindowsEdition
}

// GetWindowsEditionOk returns a tuple with the WindowsEdition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsGlobalConfigAllOf) GetWindowsEditionOk() (*string, bool) {
	if o == nil || o.WindowsEdition == nil {
		return nil, false
	}
	return o.WindowsEdition, true
}

// HasWindowsEdition returns a boolean if a field has been set.
func (o *OsGlobalConfigAllOf) HasWindowsEdition() bool {
	if o != nil && o.WindowsEdition != nil {
		return true
	}

	return false
}

// SetWindowsEdition gets a reference to the given string and assigns it to the WindowsEdition field.
func (o *OsGlobalConfigAllOf) SetWindowsEdition(v string) {
	o.WindowsEdition = &v
}

func (o OsGlobalConfigAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigurationFileName != nil {
		toSerialize["ConfigurationFileName"] = o.ConfigurationFileName
	}
	if o.ConfigurationSource != nil {
		toSerialize["ConfigurationSource"] = o.ConfigurationSource
	}
	if o.InstallMethod != nil {
		toSerialize["InstallMethod"] = o.InstallMethod
	}
	if o.InstallTargetType != nil {
		toSerialize["InstallTargetType"] = o.InstallTargetType
	}
	if o.OperatingSystemParameters.IsSet() {
		toSerialize["OperatingSystemParameters"] = o.OperatingSystemParameters.Get()
	}
	if o.OsImageName != nil {
		toSerialize["OsImageName"] = o.OsImageName
	}
	if o.ScuImageName != nil {
		toSerialize["ScuImageName"] = o.ScuImageName
	}
	if o.WindowsEdition != nil {
		toSerialize["WindowsEdition"] = o.WindowsEdition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsGlobalConfigAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOsGlobalConfigAllOf := _OsGlobalConfigAllOf{}

	if err = json.Unmarshal(bytes, &varOsGlobalConfigAllOf); err == nil {
		*o = OsGlobalConfigAllOf(varOsGlobalConfigAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigurationFileName")
		delete(additionalProperties, "ConfigurationSource")
		delete(additionalProperties, "InstallMethod")
		delete(additionalProperties, "InstallTargetType")
		delete(additionalProperties, "OperatingSystemParameters")
		delete(additionalProperties, "OsImageName")
		delete(additionalProperties, "ScuImageName")
		delete(additionalProperties, "WindowsEdition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOsGlobalConfigAllOf struct {
	value *OsGlobalConfigAllOf
	isSet bool
}

func (v NullableOsGlobalConfigAllOf) Get() *OsGlobalConfigAllOf {
	return v.value
}

func (v *NullableOsGlobalConfigAllOf) Set(val *OsGlobalConfigAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOsGlobalConfigAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOsGlobalConfigAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsGlobalConfigAllOf(val *OsGlobalConfigAllOf) *NullableOsGlobalConfigAllOf {
	return &NullableOsGlobalConfigAllOf{value: val, isSet: true}
}

func (v NullableOsGlobalConfigAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsGlobalConfigAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
