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
	"reflect"
	"strings"
)

// KubernetesAddonConfiguration Configuration settings for an Addon.
type KubernetesAddonConfiguration struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Addon install strategy to determine whether an addon is installed if not present. * `None` - Unspecified install strategy. * `NoAction` - No install action performed. * `InstallOnly` - Only install in green field. No action in case of failure or removal. * `Always` - Attempt install if chart is not already installed.
	InstallStrategy *string              `json:"InstallStrategy,omitempty"`
	OverrideSets    []KubernetesKeyValue `json:"OverrideSets,omitempty"`
	// Properties that can be overridden for an addon.
	Overrides *string `json:"Overrides,omitempty"`
	// Name for the helm release.
	ReleaseName *string `json:"ReleaseName,omitempty"`
	// Namespace for the helm release.
	ReleaseNamespace *string `json:"ReleaseNamespace,omitempty"`
	// Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade. * `None` - Unspecified upgrade strategy. * `NoAction` - This choice enables No upgrades to be performed. * `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure. * `AlwaysReinstall` - Always remove older release and reinstall.
	UpgradeStrategy      *string `json:"UpgradeStrategy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesAddonConfiguration KubernetesAddonConfiguration

// NewKubernetesAddonConfiguration instantiates a new KubernetesAddonConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAddonConfiguration(classId string, objectType string) *KubernetesAddonConfiguration {
	this := KubernetesAddonConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	var installStrategy string = "None"
	this.InstallStrategy = &installStrategy
	var upgradeStrategy string = "None"
	this.UpgradeStrategy = &upgradeStrategy
	return &this
}

// NewKubernetesAddonConfigurationWithDefaults instantiates a new KubernetesAddonConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAddonConfigurationWithDefaults() *KubernetesAddonConfiguration {
	this := KubernetesAddonConfiguration{}
	var classId string = "kubernetes.AddonConfiguration"
	this.ClassId = classId
	var objectType string = "kubernetes.AddonConfiguration"
	this.ObjectType = objectType
	var installStrategy string = "None"
	this.InstallStrategy = &installStrategy
	var upgradeStrategy string = "None"
	this.UpgradeStrategy = &upgradeStrategy
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesAddonConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesAddonConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesAddonConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesAddonConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInstallStrategy returns the InstallStrategy field value if set, zero value otherwise.
func (o *KubernetesAddonConfiguration) GetInstallStrategy() string {
	if o == nil || o.InstallStrategy == nil {
		var ret string
		return ret
	}
	return *o.InstallStrategy
}

// GetInstallStrategyOk returns a tuple with the InstallStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetInstallStrategyOk() (*string, bool) {
	if o == nil || o.InstallStrategy == nil {
		return nil, false
	}
	return o.InstallStrategy, true
}

// HasInstallStrategy returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasInstallStrategy() bool {
	if o != nil && o.InstallStrategy != nil {
		return true
	}

	return false
}

// SetInstallStrategy gets a reference to the given string and assigns it to the InstallStrategy field.
func (o *KubernetesAddonConfiguration) SetInstallStrategy(v string) {
	o.InstallStrategy = &v
}

// GetOverrideSets returns the OverrideSets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAddonConfiguration) GetOverrideSets() []KubernetesKeyValue {
	if o == nil {
		var ret []KubernetesKeyValue
		return ret
	}
	return o.OverrideSets
}

// GetOverrideSetsOk returns a tuple with the OverrideSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAddonConfiguration) GetOverrideSetsOk() (*[]KubernetesKeyValue, bool) {
	if o == nil || o.OverrideSets == nil {
		return nil, false
	}
	return &o.OverrideSets, true
}

// HasOverrideSets returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasOverrideSets() bool {
	if o != nil && o.OverrideSets != nil {
		return true
	}

	return false
}

// SetOverrideSets gets a reference to the given []KubernetesKeyValue and assigns it to the OverrideSets field.
func (o *KubernetesAddonConfiguration) SetOverrideSets(v []KubernetesKeyValue) {
	o.OverrideSets = v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *KubernetesAddonConfiguration) GetOverrides() string {
	if o == nil || o.Overrides == nil {
		var ret string
		return ret
	}
	return *o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetOverridesOk() (*string, bool) {
	if o == nil || o.Overrides == nil {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasOverrides() bool {
	if o != nil && o.Overrides != nil {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given string and assigns it to the Overrides field.
func (o *KubernetesAddonConfiguration) SetOverrides(v string) {
	o.Overrides = &v
}

// GetReleaseName returns the ReleaseName field value if set, zero value otherwise.
func (o *KubernetesAddonConfiguration) GetReleaseName() string {
	if o == nil || o.ReleaseName == nil {
		var ret string
		return ret
	}
	return *o.ReleaseName
}

// GetReleaseNameOk returns a tuple with the ReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetReleaseNameOk() (*string, bool) {
	if o == nil || o.ReleaseName == nil {
		return nil, false
	}
	return o.ReleaseName, true
}

// HasReleaseName returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasReleaseName() bool {
	if o != nil && o.ReleaseName != nil {
		return true
	}

	return false
}

// SetReleaseName gets a reference to the given string and assigns it to the ReleaseName field.
func (o *KubernetesAddonConfiguration) SetReleaseName(v string) {
	o.ReleaseName = &v
}

// GetReleaseNamespace returns the ReleaseNamespace field value if set, zero value otherwise.
func (o *KubernetesAddonConfiguration) GetReleaseNamespace() string {
	if o == nil || o.ReleaseNamespace == nil {
		var ret string
		return ret
	}
	return *o.ReleaseNamespace
}

// GetReleaseNamespaceOk returns a tuple with the ReleaseNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetReleaseNamespaceOk() (*string, bool) {
	if o == nil || o.ReleaseNamespace == nil {
		return nil, false
	}
	return o.ReleaseNamespace, true
}

// HasReleaseNamespace returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasReleaseNamespace() bool {
	if o != nil && o.ReleaseNamespace != nil {
		return true
	}

	return false
}

// SetReleaseNamespace gets a reference to the given string and assigns it to the ReleaseNamespace field.
func (o *KubernetesAddonConfiguration) SetReleaseNamespace(v string) {
	o.ReleaseNamespace = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *KubernetesAddonConfiguration) GetUpgradeStrategy() string {
	if o == nil || o.UpgradeStrategy == nil {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAddonConfiguration) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || o.UpgradeStrategy == nil {
		return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *KubernetesAddonConfiguration) HasUpgradeStrategy() bool {
	if o != nil && o.UpgradeStrategy != nil {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *KubernetesAddonConfiguration) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

func (o KubernetesAddonConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InstallStrategy != nil {
		toSerialize["InstallStrategy"] = o.InstallStrategy
	}
	if o.OverrideSets != nil {
		toSerialize["OverrideSets"] = o.OverrideSets
	}
	if o.Overrides != nil {
		toSerialize["Overrides"] = o.Overrides
	}
	if o.ReleaseName != nil {
		toSerialize["ReleaseName"] = o.ReleaseName
	}
	if o.ReleaseNamespace != nil {
		toSerialize["ReleaseNamespace"] = o.ReleaseNamespace
	}
	if o.UpgradeStrategy != nil {
		toSerialize["UpgradeStrategy"] = o.UpgradeStrategy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesAddonConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesAddonConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Addon install strategy to determine whether an addon is installed if not present. * `None` - Unspecified install strategy. * `NoAction` - No install action performed. * `InstallOnly` - Only install in green field. No action in case of failure or removal. * `Always` - Attempt install if chart is not already installed.
		InstallStrategy *string              `json:"InstallStrategy,omitempty"`
		OverrideSets    []KubernetesKeyValue `json:"OverrideSets,omitempty"`
		// Properties that can be overridden for an addon.
		Overrides *string `json:"Overrides,omitempty"`
		// Name for the helm release.
		ReleaseName *string `json:"ReleaseName,omitempty"`
		// Namespace for the helm release.
		ReleaseNamespace *string `json:"ReleaseNamespace,omitempty"`
		// Addon upgrade strategy to determine whether an addon configuration is overwritten on upgrade. * `None` - Unspecified upgrade strategy. * `NoAction` - This choice enables No upgrades to be performed. * `UpgradeOnly` - Attempt upgrade if chart or overrides options change, no action on upgrade failure. * `ReinstallOnFailure` - Attempt upgrade first. Remove and install on upgrade failure. * `AlwaysReinstall` - Always remove older release and reinstall.
		UpgradeStrategy *string `json:"UpgradeStrategy,omitempty"`
	}

	varKubernetesAddonConfigurationWithoutEmbeddedStruct := KubernetesAddonConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesAddonConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAddonConfiguration := _KubernetesAddonConfiguration{}
		varKubernetesAddonConfiguration.ClassId = varKubernetesAddonConfigurationWithoutEmbeddedStruct.ClassId
		varKubernetesAddonConfiguration.ObjectType = varKubernetesAddonConfigurationWithoutEmbeddedStruct.ObjectType
		varKubernetesAddonConfiguration.InstallStrategy = varKubernetesAddonConfigurationWithoutEmbeddedStruct.InstallStrategy
		varKubernetesAddonConfiguration.OverrideSets = varKubernetesAddonConfigurationWithoutEmbeddedStruct.OverrideSets
		varKubernetesAddonConfiguration.Overrides = varKubernetesAddonConfigurationWithoutEmbeddedStruct.Overrides
		varKubernetesAddonConfiguration.ReleaseName = varKubernetesAddonConfigurationWithoutEmbeddedStruct.ReleaseName
		varKubernetesAddonConfiguration.ReleaseNamespace = varKubernetesAddonConfigurationWithoutEmbeddedStruct.ReleaseNamespace
		varKubernetesAddonConfiguration.UpgradeStrategy = varKubernetesAddonConfigurationWithoutEmbeddedStruct.UpgradeStrategy
		*o = KubernetesAddonConfiguration(varKubernetesAddonConfiguration)
	} else {
		return err
	}

	varKubernetesAddonConfiguration := _KubernetesAddonConfiguration{}

	err = json.Unmarshal(bytes, &varKubernetesAddonConfiguration)
	if err == nil {
		o.MoBaseComplexType = varKubernetesAddonConfiguration.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InstallStrategy")
		delete(additionalProperties, "OverrideSets")
		delete(additionalProperties, "Overrides")
		delete(additionalProperties, "ReleaseName")
		delete(additionalProperties, "ReleaseNamespace")
		delete(additionalProperties, "UpgradeStrategy")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableKubernetesAddonConfiguration struct {
	value *KubernetesAddonConfiguration
	isSet bool
}

func (v NullableKubernetesAddonConfiguration) Get() *KubernetesAddonConfiguration {
	return v.value
}

func (v *NullableKubernetesAddonConfiguration) Set(val *KubernetesAddonConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAddonConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAddonConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAddonConfiguration(val *KubernetesAddonConfiguration) *NullableKubernetesAddonConfiguration {
	return &NullableKubernetesAddonConfiguration{value: val, isSet: true}
}

func (v NullableKubernetesAddonConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAddonConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
