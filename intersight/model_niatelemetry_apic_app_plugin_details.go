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

// NiatelemetryApicAppPluginDetails Object to capture APIC App plugin details.
type NiatelemetryApicAppPluginDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// App Id of the plugin in APIC.
	AppId *string `json:"AppId,omitempty"`
	// Permissions to the app plugin in APIC.
	Permission *string `json:"Permission,omitempty"`
	// Permission Level of the app plugin in APIC.
	PermissionLevel *string `json:"PermissionLevel,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicAppPluginDetails NiatelemetryApicAppPluginDetails

// NewNiatelemetryApicAppPluginDetails instantiates a new NiatelemetryApicAppPluginDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicAppPluginDetails(classId string, objectType string) *NiatelemetryApicAppPluginDetails {
	this := NiatelemetryApicAppPluginDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicAppPluginDetailsWithDefaults instantiates a new NiatelemetryApicAppPluginDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicAppPluginDetailsWithDefaults() *NiatelemetryApicAppPluginDetails {
	this := NiatelemetryApicAppPluginDetails{}
	var classId string = "niatelemetry.ApicAppPluginDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicAppPluginDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicAppPluginDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicAppPluginDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicAppPluginDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicAppPluginDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *NiatelemetryApicAppPluginDetails) SetAppId(v string) {
	o.AppId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *NiatelemetryApicAppPluginDetails) SetPermission(v string) {
	o.Permission = &v
}

// GetPermissionLevel returns the PermissionLevel field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetPermissionLevel() string {
	if o == nil || o.PermissionLevel == nil {
		var ret string
		return ret
	}
	return *o.PermissionLevel
}

// GetPermissionLevelOk returns a tuple with the PermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetPermissionLevelOk() (*string, bool) {
	if o == nil || o.PermissionLevel == nil {
		return nil, false
	}
	return o.PermissionLevel, true
}

// HasPermissionLevel returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasPermissionLevel() bool {
	if o != nil && o.PermissionLevel != nil {
		return true
	}

	return false
}

// SetPermissionLevel gets a reference to the given string and assigns it to the PermissionLevel field.
func (o *NiatelemetryApicAppPluginDetails) SetPermissionLevel(v string) {
	o.PermissionLevel = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicAppPluginDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicAppPluginDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicAppPluginDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicAppPluginDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicAppPluginDetails) MarshalJSON() ([]byte, error) {
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
	if o.AppId != nil {
		toSerialize["AppId"] = o.AppId
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.PermissionLevel != nil {
		toSerialize["PermissionLevel"] = o.PermissionLevel
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicAppPluginDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// App Id of the plugin in APIC.
		AppId *string `json:"AppId,omitempty"`
		// Permissions to the app plugin in APIC.
		Permission *string `json:"Permission,omitempty"`
		// Permission Level of the app plugin in APIC.
		PermissionLevel *string `json:"PermissionLevel,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct := NiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicAppPluginDetails := _NiatelemetryApicAppPluginDetails{}
		varNiatelemetryApicAppPluginDetails.ClassId = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicAppPluginDetails.ObjectType = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicAppPluginDetails.AppId = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.AppId
		varNiatelemetryApicAppPluginDetails.Permission = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.Permission
		varNiatelemetryApicAppPluginDetails.PermissionLevel = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.PermissionLevel
		varNiatelemetryApicAppPluginDetails.RecordType = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicAppPluginDetails.RecordVersion = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicAppPluginDetails.SiteName = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicAppPluginDetails.RegisteredDevice = varNiatelemetryApicAppPluginDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicAppPluginDetails(varNiatelemetryApicAppPluginDetails)
	} else {
		return err
	}

	varNiatelemetryApicAppPluginDetails := _NiatelemetryApicAppPluginDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicAppPluginDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicAppPluginDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppId")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "PermissionLevel")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryApicAppPluginDetails struct {
	value *NiatelemetryApicAppPluginDetails
	isSet bool
}

func (v NullableNiatelemetryApicAppPluginDetails) Get() *NiatelemetryApicAppPluginDetails {
	return v.value
}

func (v *NullableNiatelemetryApicAppPluginDetails) Set(val *NiatelemetryApicAppPluginDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicAppPluginDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicAppPluginDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicAppPluginDetails(val *NiatelemetryApicAppPluginDetails) *NullableNiatelemetryApicAppPluginDetails {
	return &NullableNiatelemetryApicAppPluginDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicAppPluginDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicAppPluginDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
