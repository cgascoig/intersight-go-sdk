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
)

// NiatelemetrySshVersionTwoAllOf Definition of the list of properties defined in 'niatelemetry.SshVersionTwo', excluding properties defined in parent classes.
type NiatelemetrySshVersionTwoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of SSH V2 in APIC.
	AdminSt *string `json:"AdminSt,omitempty"`
	// Dn of SSH V2 attribute in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Password auth for SSH V2 in APIC.
	PasswordAuth *string `json:"PasswordAuth,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// SSH Ciphers for SSH V2 in APIC.
	SshCiphers *string `json:"SshCiphers,omitempty"`
	// SSH MACS for SSH V2 in APIC.
	SshMacs              *string                              `json:"SshMacs,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetrySshVersionTwoAllOf NiatelemetrySshVersionTwoAllOf

// NewNiatelemetrySshVersionTwoAllOf instantiates a new NiatelemetrySshVersionTwoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetrySshVersionTwoAllOf(classId string, objectType string) *NiatelemetrySshVersionTwoAllOf {
	this := NiatelemetrySshVersionTwoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetrySshVersionTwoAllOfWithDefaults instantiates a new NiatelemetrySshVersionTwoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetrySshVersionTwoAllOfWithDefaults() *NiatelemetrySshVersionTwoAllOf {
	this := NiatelemetrySshVersionTwoAllOf{}
	var classId string = "niatelemetry.SshVersionTwo"
	this.ClassId = classId
	var objectType string = "niatelemetry.SshVersionTwo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetrySshVersionTwoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetrySshVersionTwoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetrySshVersionTwoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetrySshVersionTwoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSt returns the AdminSt field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetAdminSt() string {
	if o == nil || o.AdminSt == nil {
		var ret string
		return ret
	}
	return *o.AdminSt
}

// GetAdminStOk returns a tuple with the AdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetAdminStOk() (*string, bool) {
	if o == nil || o.AdminSt == nil {
		return nil, false
	}
	return o.AdminSt, true
}

// HasAdminSt returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasAdminSt() bool {
	if o != nil && o.AdminSt != nil {
		return true
	}

	return false
}

// SetAdminSt gets a reference to the given string and assigns it to the AdminSt field.
func (o *NiatelemetrySshVersionTwoAllOf) SetAdminSt(v string) {
	o.AdminSt = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetrySshVersionTwoAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetPasswordAuth returns the PasswordAuth field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetPasswordAuth() string {
	if o == nil || o.PasswordAuth == nil {
		var ret string
		return ret
	}
	return *o.PasswordAuth
}

// GetPasswordAuthOk returns a tuple with the PasswordAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetPasswordAuthOk() (*string, bool) {
	if o == nil || o.PasswordAuth == nil {
		return nil, false
	}
	return o.PasswordAuth, true
}

// HasPasswordAuth returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasPasswordAuth() bool {
	if o != nil && o.PasswordAuth != nil {
		return true
	}

	return false
}

// SetPasswordAuth gets a reference to the given string and assigns it to the PasswordAuth field.
func (o *NiatelemetrySshVersionTwoAllOf) SetPasswordAuth(v string) {
	o.PasswordAuth = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetrySshVersionTwoAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetrySshVersionTwoAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetrySshVersionTwoAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSshCiphers returns the SshCiphers field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetSshCiphers() string {
	if o == nil || o.SshCiphers == nil {
		var ret string
		return ret
	}
	return *o.SshCiphers
}

// GetSshCiphersOk returns a tuple with the SshCiphers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetSshCiphersOk() (*string, bool) {
	if o == nil || o.SshCiphers == nil {
		return nil, false
	}
	return o.SshCiphers, true
}

// HasSshCiphers returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasSshCiphers() bool {
	if o != nil && o.SshCiphers != nil {
		return true
	}

	return false
}

// SetSshCiphers gets a reference to the given string and assigns it to the SshCiphers field.
func (o *NiatelemetrySshVersionTwoAllOf) SetSshCiphers(v string) {
	o.SshCiphers = &v
}

// GetSshMacs returns the SshMacs field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetSshMacs() string {
	if o == nil || o.SshMacs == nil {
		var ret string
		return ret
	}
	return *o.SshMacs
}

// GetSshMacsOk returns a tuple with the SshMacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetSshMacsOk() (*string, bool) {
	if o == nil || o.SshMacs == nil {
		return nil, false
	}
	return o.SshMacs, true
}

// HasSshMacs returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasSshMacs() bool {
	if o != nil && o.SshMacs != nil {
		return true
	}

	return false
}

// SetSshMacs gets a reference to the given string and assigns it to the SshMacs field.
func (o *NiatelemetrySshVersionTwoAllOf) SetSshMacs(v string) {
	o.SshMacs = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetrySshVersionTwoAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetrySshVersionTwoAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetrySshVersionTwoAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetrySshVersionTwoAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetrySshVersionTwoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSt != nil {
		toSerialize["AdminSt"] = o.AdminSt
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.PasswordAuth != nil {
		toSerialize["PasswordAuth"] = o.PasswordAuth
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
	if o.SshCiphers != nil {
		toSerialize["SshCiphers"] = o.SshCiphers
	}
	if o.SshMacs != nil {
		toSerialize["SshMacs"] = o.SshMacs
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetrySshVersionTwoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetrySshVersionTwoAllOf := _NiatelemetrySshVersionTwoAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetrySshVersionTwoAllOf); err == nil {
		*o = NiatelemetrySshVersionTwoAllOf(varNiatelemetrySshVersionTwoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSt")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "PasswordAuth")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SshCiphers")
		delete(additionalProperties, "SshMacs")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetrySshVersionTwoAllOf struct {
	value *NiatelemetrySshVersionTwoAllOf
	isSet bool
}

func (v NullableNiatelemetrySshVersionTwoAllOf) Get() *NiatelemetrySshVersionTwoAllOf {
	return v.value
}

func (v *NullableNiatelemetrySshVersionTwoAllOf) Set(val *NiatelemetrySshVersionTwoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetrySshVersionTwoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetrySshVersionTwoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetrySshVersionTwoAllOf(val *NiatelemetrySshVersionTwoAllOf) *NullableNiatelemetrySshVersionTwoAllOf {
	return &NullableNiatelemetrySshVersionTwoAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetrySshVersionTwoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetrySshVersionTwoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
