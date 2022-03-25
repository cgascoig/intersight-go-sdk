/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VnicFcErrorRecoverySettings Fibre Channel Error Recovery Settings.
type VnicFcErrorRecoverySettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enables Fibre Channel Error recovery.
	Enabled *bool `json:"Enabled,omitempty"`
	// The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.
	IoRetryCount *int64 `json:"IoRetryCount,omitempty"`
	// The number of seconds the adapter waits before aborting the pending command and resending the same IO request.
	IoRetryTimeout *int64 `json:"IoRetryTimeout,omitempty"`
	// The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.
	LinkDownTimeout *int64 `json:"LinkDownTimeout,omitempty"`
	// The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.
	PortDownTimeout      *int64 `json:"PortDownTimeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcErrorRecoverySettings VnicFcErrorRecoverySettings

// NewVnicFcErrorRecoverySettings instantiates a new VnicFcErrorRecoverySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcErrorRecoverySettings(classId string, objectType string) *VnicFcErrorRecoverySettings {
	this := VnicFcErrorRecoverySettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ioRetryCount int64 = 8
	this.IoRetryCount = &ioRetryCount
	var ioRetryTimeout int64 = 5
	this.IoRetryTimeout = &ioRetryTimeout
	var linkDownTimeout int64 = 30000
	this.LinkDownTimeout = &linkDownTimeout
	var portDownTimeout int64 = 10000
	this.PortDownTimeout = &portDownTimeout
	return &this
}

// NewVnicFcErrorRecoverySettingsWithDefaults instantiates a new VnicFcErrorRecoverySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcErrorRecoverySettingsWithDefaults() *VnicFcErrorRecoverySettings {
	this := VnicFcErrorRecoverySettings{}
	var classId string = "vnic.FcErrorRecoverySettings"
	this.ClassId = classId
	var objectType string = "vnic.FcErrorRecoverySettings"
	this.ObjectType = objectType
	var ioRetryCount int64 = 8
	this.IoRetryCount = &ioRetryCount
	var ioRetryTimeout int64 = 5
	this.IoRetryTimeout = &ioRetryTimeout
	var linkDownTimeout int64 = 30000
	this.LinkDownTimeout = &linkDownTimeout
	var portDownTimeout int64 = 10000
	this.PortDownTimeout = &portDownTimeout
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcErrorRecoverySettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcErrorRecoverySettings) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcErrorRecoverySettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcErrorRecoverySettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicFcErrorRecoverySettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicFcErrorRecoverySettings) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicFcErrorRecoverySettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIoRetryCount returns the IoRetryCount field value if set, zero value otherwise.
func (o *VnicFcErrorRecoverySettings) GetIoRetryCount() int64 {
	if o == nil || o.IoRetryCount == nil {
		var ret int64
		return ret
	}
	return *o.IoRetryCount
}

// GetIoRetryCountOk returns a tuple with the IoRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetIoRetryCountOk() (*int64, bool) {
	if o == nil || o.IoRetryCount == nil {
		return nil, false
	}
	return o.IoRetryCount, true
}

// HasIoRetryCount returns a boolean if a field has been set.
func (o *VnicFcErrorRecoverySettings) HasIoRetryCount() bool {
	if o != nil && o.IoRetryCount != nil {
		return true
	}

	return false
}

// SetIoRetryCount gets a reference to the given int64 and assigns it to the IoRetryCount field.
func (o *VnicFcErrorRecoverySettings) SetIoRetryCount(v int64) {
	o.IoRetryCount = &v
}

// GetIoRetryTimeout returns the IoRetryTimeout field value if set, zero value otherwise.
func (o *VnicFcErrorRecoverySettings) GetIoRetryTimeout() int64 {
	if o == nil || o.IoRetryTimeout == nil {
		var ret int64
		return ret
	}
	return *o.IoRetryTimeout
}

// GetIoRetryTimeoutOk returns a tuple with the IoRetryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetIoRetryTimeoutOk() (*int64, bool) {
	if o == nil || o.IoRetryTimeout == nil {
		return nil, false
	}
	return o.IoRetryTimeout, true
}

// HasIoRetryTimeout returns a boolean if a field has been set.
func (o *VnicFcErrorRecoverySettings) HasIoRetryTimeout() bool {
	if o != nil && o.IoRetryTimeout != nil {
		return true
	}

	return false
}

// SetIoRetryTimeout gets a reference to the given int64 and assigns it to the IoRetryTimeout field.
func (o *VnicFcErrorRecoverySettings) SetIoRetryTimeout(v int64) {
	o.IoRetryTimeout = &v
}

// GetLinkDownTimeout returns the LinkDownTimeout field value if set, zero value otherwise.
func (o *VnicFcErrorRecoverySettings) GetLinkDownTimeout() int64 {
	if o == nil || o.LinkDownTimeout == nil {
		var ret int64
		return ret
	}
	return *o.LinkDownTimeout
}

// GetLinkDownTimeoutOk returns a tuple with the LinkDownTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetLinkDownTimeoutOk() (*int64, bool) {
	if o == nil || o.LinkDownTimeout == nil {
		return nil, false
	}
	return o.LinkDownTimeout, true
}

// HasLinkDownTimeout returns a boolean if a field has been set.
func (o *VnicFcErrorRecoverySettings) HasLinkDownTimeout() bool {
	if o != nil && o.LinkDownTimeout != nil {
		return true
	}

	return false
}

// SetLinkDownTimeout gets a reference to the given int64 and assigns it to the LinkDownTimeout field.
func (o *VnicFcErrorRecoverySettings) SetLinkDownTimeout(v int64) {
	o.LinkDownTimeout = &v
}

// GetPortDownTimeout returns the PortDownTimeout field value if set, zero value otherwise.
func (o *VnicFcErrorRecoverySettings) GetPortDownTimeout() int64 {
	if o == nil || o.PortDownTimeout == nil {
		var ret int64
		return ret
	}
	return *o.PortDownTimeout
}

// GetPortDownTimeoutOk returns a tuple with the PortDownTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcErrorRecoverySettings) GetPortDownTimeoutOk() (*int64, bool) {
	if o == nil || o.PortDownTimeout == nil {
		return nil, false
	}
	return o.PortDownTimeout, true
}

// HasPortDownTimeout returns a boolean if a field has been set.
func (o *VnicFcErrorRecoverySettings) HasPortDownTimeout() bool {
	if o != nil && o.PortDownTimeout != nil {
		return true
	}

	return false
}

// SetPortDownTimeout gets a reference to the given int64 and assigns it to the PortDownTimeout field.
func (o *VnicFcErrorRecoverySettings) SetPortDownTimeout(v int64) {
	o.PortDownTimeout = &v
}

func (o VnicFcErrorRecoverySettings) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.IoRetryCount != nil {
		toSerialize["IoRetryCount"] = o.IoRetryCount
	}
	if o.IoRetryTimeout != nil {
		toSerialize["IoRetryTimeout"] = o.IoRetryTimeout
	}
	if o.LinkDownTimeout != nil {
		toSerialize["LinkDownTimeout"] = o.LinkDownTimeout
	}
	if o.PortDownTimeout != nil {
		toSerialize["PortDownTimeout"] = o.PortDownTimeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcErrorRecoverySettings) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcErrorRecoverySettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enables Fibre Channel Error recovery.
		Enabled *bool `json:"Enabled,omitempty"`
		// The number of times an I/O request to a port is retried because the port is busy before the system decides the port is unavailable.
		IoRetryCount *int64 `json:"IoRetryCount,omitempty"`
		// The number of seconds the adapter waits before aborting the pending command and resending the same IO request.
		IoRetryTimeout *int64 `json:"IoRetryTimeout,omitempty"`
		// The number of milliseconds the port should actually be down before it is marked down and fabric connectivity is lost.
		LinkDownTimeout *int64 `json:"LinkDownTimeout,omitempty"`
		// The number of milliseconds a remote Fibre Channel port should be offline before informing the SCSI upper layer that the port is unavailable. For a server with a VIC adapter running ESXi, the recommended value is 10000. For a server with a port used to boot a Windows OS from the SAN, the recommended value is 5000 milliseconds.
		PortDownTimeout *int64 `json:"PortDownTimeout,omitempty"`
	}

	varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct := VnicFcErrorRecoverySettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcErrorRecoverySettings := _VnicFcErrorRecoverySettings{}
		varVnicFcErrorRecoverySettings.ClassId = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.ClassId
		varVnicFcErrorRecoverySettings.ObjectType = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.ObjectType
		varVnicFcErrorRecoverySettings.Enabled = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.Enabled
		varVnicFcErrorRecoverySettings.IoRetryCount = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.IoRetryCount
		varVnicFcErrorRecoverySettings.IoRetryTimeout = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.IoRetryTimeout
		varVnicFcErrorRecoverySettings.LinkDownTimeout = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.LinkDownTimeout
		varVnicFcErrorRecoverySettings.PortDownTimeout = varVnicFcErrorRecoverySettingsWithoutEmbeddedStruct.PortDownTimeout
		*o = VnicFcErrorRecoverySettings(varVnicFcErrorRecoverySettings)
	} else {
		return err
	}

	varVnicFcErrorRecoverySettings := _VnicFcErrorRecoverySettings{}

	err = json.Unmarshal(bytes, &varVnicFcErrorRecoverySettings)
	if err == nil {
		o.MoBaseComplexType = varVnicFcErrorRecoverySettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IoRetryCount")
		delete(additionalProperties, "IoRetryTimeout")
		delete(additionalProperties, "LinkDownTimeout")
		delete(additionalProperties, "PortDownTimeout")

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

type NullableVnicFcErrorRecoverySettings struct {
	value *VnicFcErrorRecoverySettings
	isSet bool
}

func (v NullableVnicFcErrorRecoverySettings) Get() *VnicFcErrorRecoverySettings {
	return v.value
}

func (v *NullableVnicFcErrorRecoverySettings) Set(val *VnicFcErrorRecoverySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcErrorRecoverySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcErrorRecoverySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcErrorRecoverySettings(val *VnicFcErrorRecoverySettings) *NullableVnicFcErrorRecoverySettings {
	return &NullableVnicFcErrorRecoverySettings{value: val, isSet: true}
}

func (v NullableVnicFcErrorRecoverySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcErrorRecoverySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
