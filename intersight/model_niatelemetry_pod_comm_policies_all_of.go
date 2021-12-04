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
)

// NiatelemetryPodCommPoliciesAllOf Definition of the list of properties defined in 'niatelemetry.PodCommPolicies', excluding properties defined in parent classes.
type NiatelemetryPodCommPoliciesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Comm Http Admin State of the Comm Pol in APIC.
	CommHttpAdminSt *string `json:"CommHttpAdminSt,omitempty"`
	// Comm Https Admin State of the Comm Pol in APIC.
	CommHttpsAdminSt *string `json:"CommHttpsAdminSt,omitempty"`
	// Comm Ssh Admin State of the Comm Pol in APIC.
	CommSshAdminSt *string `json:"CommSshAdminSt,omitempty"`
	// Comm Telnet Admin State of the Comm Pol in APIC.
	CommTelnetAdminSt *string `json:"CommTelnetAdminSt,omitempty"`
	// Dn of the Comm Pol in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodCommPoliciesAllOf NiatelemetryPodCommPoliciesAllOf

// NewNiatelemetryPodCommPoliciesAllOf instantiates a new NiatelemetryPodCommPoliciesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodCommPoliciesAllOf(classId string, objectType string) *NiatelemetryPodCommPoliciesAllOf {
	this := NiatelemetryPodCommPoliciesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodCommPoliciesAllOfWithDefaults instantiates a new NiatelemetryPodCommPoliciesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodCommPoliciesAllOfWithDefaults() *NiatelemetryPodCommPoliciesAllOf {
	this := NiatelemetryPodCommPoliciesAllOf{}
	var classId string = "niatelemetry.PodCommPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodCommPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodCommPoliciesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodCommPoliciesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodCommPoliciesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodCommPoliciesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommHttpAdminSt returns the CommHttpAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpAdminSt() string {
	if o == nil || o.CommHttpAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommHttpAdminSt
}

// GetCommHttpAdminStOk returns a tuple with the CommHttpAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpAdminStOk() (*string, bool) {
	if o == nil || o.CommHttpAdminSt == nil {
		return nil, false
	}
	return o.CommHttpAdminSt, true
}

// HasCommHttpAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasCommHttpAdminSt() bool {
	if o != nil && o.CommHttpAdminSt != nil {
		return true
	}

	return false
}

// SetCommHttpAdminSt gets a reference to the given string and assigns it to the CommHttpAdminSt field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetCommHttpAdminSt(v string) {
	o.CommHttpAdminSt = &v
}

// GetCommHttpsAdminSt returns the CommHttpsAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpsAdminSt() string {
	if o == nil || o.CommHttpsAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommHttpsAdminSt
}

// GetCommHttpsAdminStOk returns a tuple with the CommHttpsAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommHttpsAdminStOk() (*string, bool) {
	if o == nil || o.CommHttpsAdminSt == nil {
		return nil, false
	}
	return o.CommHttpsAdminSt, true
}

// HasCommHttpsAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasCommHttpsAdminSt() bool {
	if o != nil && o.CommHttpsAdminSt != nil {
		return true
	}

	return false
}

// SetCommHttpsAdminSt gets a reference to the given string and assigns it to the CommHttpsAdminSt field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetCommHttpsAdminSt(v string) {
	o.CommHttpsAdminSt = &v
}

// GetCommSshAdminSt returns the CommSshAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommSshAdminSt() string {
	if o == nil || o.CommSshAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommSshAdminSt
}

// GetCommSshAdminStOk returns a tuple with the CommSshAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommSshAdminStOk() (*string, bool) {
	if o == nil || o.CommSshAdminSt == nil {
		return nil, false
	}
	return o.CommSshAdminSt, true
}

// HasCommSshAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasCommSshAdminSt() bool {
	if o != nil && o.CommSshAdminSt != nil {
		return true
	}

	return false
}

// SetCommSshAdminSt gets a reference to the given string and assigns it to the CommSshAdminSt field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetCommSshAdminSt(v string) {
	o.CommSshAdminSt = &v
}

// GetCommTelnetAdminSt returns the CommTelnetAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommTelnetAdminSt() string {
	if o == nil || o.CommTelnetAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommTelnetAdminSt
}

// GetCommTelnetAdminStOk returns a tuple with the CommTelnetAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetCommTelnetAdminStOk() (*string, bool) {
	if o == nil || o.CommTelnetAdminSt == nil {
		return nil, false
	}
	return o.CommTelnetAdminSt, true
}

// HasCommTelnetAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasCommTelnetAdminSt() bool {
	if o != nil && o.CommTelnetAdminSt != nil {
		return true
	}

	return false
}

// SetCommTelnetAdminSt gets a reference to the given string and assigns it to the CommTelnetAdminSt field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetCommTelnetAdminSt(v string) {
	o.CommTelnetAdminSt = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPoliciesAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodCommPoliciesAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryPodCommPoliciesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CommHttpAdminSt != nil {
		toSerialize["CommHttpAdminSt"] = o.CommHttpAdminSt
	}
	if o.CommHttpsAdminSt != nil {
		toSerialize["CommHttpsAdminSt"] = o.CommHttpsAdminSt
	}
	if o.CommSshAdminSt != nil {
		toSerialize["CommSshAdminSt"] = o.CommSshAdminSt
	}
	if o.CommTelnetAdminSt != nil {
		toSerialize["CommTelnetAdminSt"] = o.CommTelnetAdminSt
	}
	if o.PolDn != nil {
		toSerialize["PolDn"] = o.PolDn
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

func (o *NiatelemetryPodCommPoliciesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryPodCommPoliciesAllOf := _NiatelemetryPodCommPoliciesAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryPodCommPoliciesAllOf); err == nil {
		*o = NiatelemetryPodCommPoliciesAllOf(varNiatelemetryPodCommPoliciesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommHttpAdminSt")
		delete(additionalProperties, "CommHttpsAdminSt")
		delete(additionalProperties, "CommSshAdminSt")
		delete(additionalProperties, "CommTelnetAdminSt")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryPodCommPoliciesAllOf struct {
	value *NiatelemetryPodCommPoliciesAllOf
	isSet bool
}

func (v NullableNiatelemetryPodCommPoliciesAllOf) Get() *NiatelemetryPodCommPoliciesAllOf {
	return v.value
}

func (v *NullableNiatelemetryPodCommPoliciesAllOf) Set(val *NiatelemetryPodCommPoliciesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodCommPoliciesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodCommPoliciesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodCommPoliciesAllOf(val *NiatelemetryPodCommPoliciesAllOf) *NullableNiatelemetryPodCommPoliciesAllOf {
	return &NullableNiatelemetryPodCommPoliciesAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryPodCommPoliciesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodCommPoliciesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
