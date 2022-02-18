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
	"reflect"
	"strings"
)

// ApplianceSystemStatus Status of the Intersight Appliance.
type ApplianceSystemStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Operational status of the Intersight Appliance. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
	OperationalStatus *string                `json:"OperationalStatus,omitempty"`
	StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
	// An array of relationships to applianceAppStatus resources.
	AppStatuses []ApplianceAppStatusRelationship `json:"AppStatuses,omitempty"`
	// An array of relationships to applianceGroupStatus resources.
	GroupStatuses        []ApplianceGroupStatusRelationship   `json:"GroupStatuses,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	SystemInfo           *ApplianceSystemInfoRelationship     `json:"SystemInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceSystemStatus ApplianceSystemStatus

// NewApplianceSystemStatus instantiates a new ApplianceSystemStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceSystemStatus(classId string, objectType string) *ApplianceSystemStatus {
	this := ApplianceSystemStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceSystemStatusWithDefaults instantiates a new ApplianceSystemStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceSystemStatusWithDefaults() *ApplianceSystemStatus {
	this := ApplianceSystemStatus{}
	var classId string = "appliance.SystemStatus"
	this.ClassId = classId
	var objectType string = "appliance.SystemStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceSystemStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceSystemStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceSystemStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceSystemStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceSystemStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceSystemStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ApplianceSystemStatus) GetOperationalStatus() string {
	if o == nil || o.OperationalStatus == nil {
		var ret string
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemStatus) GetOperationalStatusOk() (*string, bool) {
	if o == nil || o.OperationalStatus == nil {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasOperationalStatus() bool {
	if o != nil && o.OperationalStatus != nil {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given string and assigns it to the OperationalStatus field.
func (o *ApplianceSystemStatus) SetOperationalStatus(v string) {
	o.OperationalStatus = &v
}

// GetStatusChecks returns the StatusChecks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceSystemStatus) GetStatusChecks() []ApplianceStatusCheck {
	if o == nil {
		var ret []ApplianceStatusCheck
		return ret
	}
	return o.StatusChecks
}

// GetStatusChecksOk returns a tuple with the StatusChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceSystemStatus) GetStatusChecksOk() (*[]ApplianceStatusCheck, bool) {
	if o == nil || o.StatusChecks == nil {
		return nil, false
	}
	return &o.StatusChecks, true
}

// HasStatusChecks returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasStatusChecks() bool {
	if o != nil && o.StatusChecks != nil {
		return true
	}

	return false
}

// SetStatusChecks gets a reference to the given []ApplianceStatusCheck and assigns it to the StatusChecks field.
func (o *ApplianceSystemStatus) SetStatusChecks(v []ApplianceStatusCheck) {
	o.StatusChecks = v
}

// GetAppStatuses returns the AppStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceSystemStatus) GetAppStatuses() []ApplianceAppStatusRelationship {
	if o == nil {
		var ret []ApplianceAppStatusRelationship
		return ret
	}
	return o.AppStatuses
}

// GetAppStatusesOk returns a tuple with the AppStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceSystemStatus) GetAppStatusesOk() (*[]ApplianceAppStatusRelationship, bool) {
	if o == nil || o.AppStatuses == nil {
		return nil, false
	}
	return &o.AppStatuses, true
}

// HasAppStatuses returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasAppStatuses() bool {
	if o != nil && o.AppStatuses != nil {
		return true
	}

	return false
}

// SetAppStatuses gets a reference to the given []ApplianceAppStatusRelationship and assigns it to the AppStatuses field.
func (o *ApplianceSystemStatus) SetAppStatuses(v []ApplianceAppStatusRelationship) {
	o.AppStatuses = v
}

// GetGroupStatuses returns the GroupStatuses field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceSystemStatus) GetGroupStatuses() []ApplianceGroupStatusRelationship {
	if o == nil {
		var ret []ApplianceGroupStatusRelationship
		return ret
	}
	return o.GroupStatuses
}

// GetGroupStatusesOk returns a tuple with the GroupStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceSystemStatus) GetGroupStatusesOk() (*[]ApplianceGroupStatusRelationship, bool) {
	if o == nil || o.GroupStatuses == nil {
		return nil, false
	}
	return &o.GroupStatuses, true
}

// HasGroupStatuses returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasGroupStatuses() bool {
	if o != nil && o.GroupStatuses != nil {
		return true
	}

	return false
}

// SetGroupStatuses gets a reference to the given []ApplianceGroupStatusRelationship and assigns it to the GroupStatuses field.
func (o *ApplianceSystemStatus) SetGroupStatuses(v []ApplianceGroupStatusRelationship) {
	o.GroupStatuses = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ApplianceSystemStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceSystemStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSystemInfo returns the SystemInfo field value if set, zero value otherwise.
func (o *ApplianceSystemStatus) GetSystemInfo() ApplianceSystemInfoRelationship {
	if o == nil || o.SystemInfo == nil {
		var ret ApplianceSystemInfoRelationship
		return ret
	}
	return *o.SystemInfo
}

// GetSystemInfoOk returns a tuple with the SystemInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceSystemStatus) GetSystemInfoOk() (*ApplianceSystemInfoRelationship, bool) {
	if o == nil || o.SystemInfo == nil {
		return nil, false
	}
	return o.SystemInfo, true
}

// HasSystemInfo returns a boolean if a field has been set.
func (o *ApplianceSystemStatus) HasSystemInfo() bool {
	if o != nil && o.SystemInfo != nil {
		return true
	}

	return false
}

// SetSystemInfo gets a reference to the given ApplianceSystemInfoRelationship and assigns it to the SystemInfo field.
func (o *ApplianceSystemStatus) SetSystemInfo(v ApplianceSystemInfoRelationship) {
	o.SystemInfo = &v
}

func (o ApplianceSystemStatus) MarshalJSON() ([]byte, error) {
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
	if o.OperationalStatus != nil {
		toSerialize["OperationalStatus"] = o.OperationalStatus
	}
	if o.StatusChecks != nil {
		toSerialize["StatusChecks"] = o.StatusChecks
	}
	if o.AppStatuses != nil {
		toSerialize["AppStatuses"] = o.AppStatuses
	}
	if o.GroupStatuses != nil {
		toSerialize["GroupStatuses"] = o.GroupStatuses
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SystemInfo != nil {
		toSerialize["SystemInfo"] = o.SystemInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceSystemStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceSystemStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Operational status of the Intersight Appliance. Operational status is based on the result of the status checks. If result of any check is Critical, then its value is Impaired. Otherwise, if result of any check is Warning, then its value is AttentionNeeded. If all checks are OK, then its value is Operational. * `Unknown` - Operational status of the Intersight Appliance entity is Unknown. * `Operational` - Operational status of the Intersight Appliance entity is Operational. * `Impaired` - Operational status of the Intersight Appliance entity is Impaired. * `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded.
		OperationalStatus *string                `json:"OperationalStatus,omitempty"`
		StatusChecks      []ApplianceStatusCheck `json:"StatusChecks,omitempty"`
		// An array of relationships to applianceAppStatus resources.
		AppStatuses []ApplianceAppStatusRelationship `json:"AppStatuses,omitempty"`
		// An array of relationships to applianceGroupStatus resources.
		GroupStatuses    []ApplianceGroupStatusRelationship   `json:"GroupStatuses,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		SystemInfo       *ApplianceSystemInfoRelationship     `json:"SystemInfo,omitempty"`
	}

	varApplianceSystemStatusWithoutEmbeddedStruct := ApplianceSystemStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceSystemStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceSystemStatus := _ApplianceSystemStatus{}
		varApplianceSystemStatus.ClassId = varApplianceSystemStatusWithoutEmbeddedStruct.ClassId
		varApplianceSystemStatus.ObjectType = varApplianceSystemStatusWithoutEmbeddedStruct.ObjectType
		varApplianceSystemStatus.OperationalStatus = varApplianceSystemStatusWithoutEmbeddedStruct.OperationalStatus
		varApplianceSystemStatus.StatusChecks = varApplianceSystemStatusWithoutEmbeddedStruct.StatusChecks
		varApplianceSystemStatus.AppStatuses = varApplianceSystemStatusWithoutEmbeddedStruct.AppStatuses
		varApplianceSystemStatus.GroupStatuses = varApplianceSystemStatusWithoutEmbeddedStruct.GroupStatuses
		varApplianceSystemStatus.RegisteredDevice = varApplianceSystemStatusWithoutEmbeddedStruct.RegisteredDevice
		varApplianceSystemStatus.SystemInfo = varApplianceSystemStatusWithoutEmbeddedStruct.SystemInfo
		*o = ApplianceSystemStatus(varApplianceSystemStatus)
	} else {
		return err
	}

	varApplianceSystemStatus := _ApplianceSystemStatus{}

	err = json.Unmarshal(bytes, &varApplianceSystemStatus)
	if err == nil {
		o.MoBaseMo = varApplianceSystemStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperationalStatus")
		delete(additionalProperties, "StatusChecks")
		delete(additionalProperties, "AppStatuses")
		delete(additionalProperties, "GroupStatuses")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SystemInfo")

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

type NullableApplianceSystemStatus struct {
	value *ApplianceSystemStatus
	isSet bool
}

func (v NullableApplianceSystemStatus) Get() *ApplianceSystemStatus {
	return v.value
}

func (v *NullableApplianceSystemStatus) Set(val *ApplianceSystemStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceSystemStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceSystemStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceSystemStatus(val *ApplianceSystemStatus) *NullableApplianceSystemStatus {
	return &NullableApplianceSystemStatus{value: val, isSet: true}
}

func (v NullableApplianceSystemStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceSystemStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
