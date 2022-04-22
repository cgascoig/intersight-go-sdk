/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryPodSnmpPolicies Object to capture Pod SNMP Policy details.
type NiatelemetryPodSnmpPolicies struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin State of the Snmp Pol in APIC.
	AdminSt *string `json:"AdminSt,omitempty"`
	// Dn of the Snmp Pol in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// List of Dn of the SNMP Client grp in APIC.
	SnmpClientGrp *string `json:"SnmpClientGrp,omitempty"`
	// List of Dn of the SNMP Community in APIC.
	SnmpCommunity *string `json:"SnmpCommunity,omitempty"`
	// List of Dn of the SNMP Trap Fwd Server in APIC.
	SnmpTrapFwdServer *string `json:"SnmpTrapFwdServer,omitempty"`
	// List of Dn of the SNMP user in APIC.
	SnmpUser             *string                              `json:"SnmpUser,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodSnmpPolicies NiatelemetryPodSnmpPolicies

// NewNiatelemetryPodSnmpPolicies instantiates a new NiatelemetryPodSnmpPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodSnmpPolicies(classId string, objectType string) *NiatelemetryPodSnmpPolicies {
	this := NiatelemetryPodSnmpPolicies{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodSnmpPoliciesWithDefaults instantiates a new NiatelemetryPodSnmpPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodSnmpPoliciesWithDefaults() *NiatelemetryPodSnmpPolicies {
	this := NiatelemetryPodSnmpPolicies{}
	var classId string = "niatelemetry.PodSnmpPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodSnmpPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodSnmpPolicies) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodSnmpPolicies) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodSnmpPolicies) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodSnmpPolicies) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSt returns the AdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetAdminSt() string {
	if o == nil || o.AdminSt == nil {
		var ret string
		return ret
	}
	return *o.AdminSt
}

// GetAdminStOk returns a tuple with the AdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetAdminStOk() (*string, bool) {
	if o == nil || o.AdminSt == nil {
		return nil, false
	}
	return o.AdminSt, true
}

// HasAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasAdminSt() bool {
	if o != nil && o.AdminSt != nil {
		return true
	}

	return false
}

// SetAdminSt gets a reference to the given string and assigns it to the AdminSt field.
func (o *NiatelemetryPodSnmpPolicies) SetAdminSt(v string) {
	o.AdminSt = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodSnmpPolicies) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodSnmpPolicies) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodSnmpPolicies) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodSnmpPolicies) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSnmpClientGrp returns the SnmpClientGrp field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrp() string {
	if o == nil || o.SnmpClientGrp == nil {
		var ret string
		return ret
	}
	return *o.SnmpClientGrp
}

// GetSnmpClientGrpOk returns a tuple with the SnmpClientGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpClientGrpOk() (*string, bool) {
	if o == nil || o.SnmpClientGrp == nil {
		return nil, false
	}
	return o.SnmpClientGrp, true
}

// HasSnmpClientGrp returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpClientGrp() bool {
	if o != nil && o.SnmpClientGrp != nil {
		return true
	}

	return false
}

// SetSnmpClientGrp gets a reference to the given string and assigns it to the SnmpClientGrp field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpClientGrp(v string) {
	o.SnmpClientGrp = &v
}

// GetSnmpCommunity returns the SnmpCommunity field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunity() string {
	if o == nil || o.SnmpCommunity == nil {
		var ret string
		return ret
	}
	return *o.SnmpCommunity
}

// GetSnmpCommunityOk returns a tuple with the SnmpCommunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpCommunityOk() (*string, bool) {
	if o == nil || o.SnmpCommunity == nil {
		return nil, false
	}
	return o.SnmpCommunity, true
}

// HasSnmpCommunity returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpCommunity() bool {
	if o != nil && o.SnmpCommunity != nil {
		return true
	}

	return false
}

// SetSnmpCommunity gets a reference to the given string and assigns it to the SnmpCommunity field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpCommunity(v string) {
	o.SnmpCommunity = &v
}

// GetSnmpTrapFwdServer returns the SnmpTrapFwdServer field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServer() string {
	if o == nil || o.SnmpTrapFwdServer == nil {
		var ret string
		return ret
	}
	return *o.SnmpTrapFwdServer
}

// GetSnmpTrapFwdServerOk returns a tuple with the SnmpTrapFwdServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpTrapFwdServerOk() (*string, bool) {
	if o == nil || o.SnmpTrapFwdServer == nil {
		return nil, false
	}
	return o.SnmpTrapFwdServer, true
}

// HasSnmpTrapFwdServer returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpTrapFwdServer() bool {
	if o != nil && o.SnmpTrapFwdServer != nil {
		return true
	}

	return false
}

// SetSnmpTrapFwdServer gets a reference to the given string and assigns it to the SnmpTrapFwdServer field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpTrapFwdServer(v string) {
	o.SnmpTrapFwdServer = &v
}

// GetSnmpUser returns the SnmpUser field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpUser() string {
	if o == nil || o.SnmpUser == nil {
		var ret string
		return ret
	}
	return *o.SnmpUser
}

// GetSnmpUserOk returns a tuple with the SnmpUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetSnmpUserOk() (*string, bool) {
	if o == nil || o.SnmpUser == nil {
		return nil, false
	}
	return o.SnmpUser, true
}

// HasSnmpUser returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasSnmpUser() bool {
	if o != nil && o.SnmpUser != nil {
		return true
	}

	return false
}

// SetSnmpUser gets a reference to the given string and assigns it to the SnmpUser field.
func (o *NiatelemetryPodSnmpPolicies) SetSnmpUser(v string) {
	o.SnmpUser = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodSnmpPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodSnmpPolicies) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodSnmpPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryPodSnmpPolicies) MarshalJSON() ([]byte, error) {
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
	if o.AdminSt != nil {
		toSerialize["AdminSt"] = o.AdminSt
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
	if o.SnmpClientGrp != nil {
		toSerialize["SnmpClientGrp"] = o.SnmpClientGrp
	}
	if o.SnmpCommunity != nil {
		toSerialize["SnmpCommunity"] = o.SnmpCommunity
	}
	if o.SnmpTrapFwdServer != nil {
		toSerialize["SnmpTrapFwdServer"] = o.SnmpTrapFwdServer
	}
	if o.SnmpUser != nil {
		toSerialize["SnmpUser"] = o.SnmpUser
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryPodSnmpPolicies) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin State of the Snmp Pol in APIC.
		AdminSt *string `json:"AdminSt,omitempty"`
		// Dn of the Snmp Pol in APIC.
		PolDn *string `json:"PolDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName *string `json:"SiteName,omitempty"`
		// List of Dn of the SNMP Client grp in APIC.
		SnmpClientGrp *string `json:"SnmpClientGrp,omitempty"`
		// List of Dn of the SNMP Community in APIC.
		SnmpCommunity *string `json:"SnmpCommunity,omitempty"`
		// List of Dn of the SNMP Trap Fwd Server in APIC.
		SnmpTrapFwdServer *string `json:"SnmpTrapFwdServer,omitempty"`
		// List of Dn of the SNMP user in APIC.
		SnmpUser         *string                              `json:"SnmpUser,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct := NiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryPodSnmpPolicies := _NiatelemetryPodSnmpPolicies{}
		varNiatelemetryPodSnmpPolicies.ClassId = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.ClassId
		varNiatelemetryPodSnmpPolicies.ObjectType = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.ObjectType
		varNiatelemetryPodSnmpPolicies.AdminSt = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.AdminSt
		varNiatelemetryPodSnmpPolicies.PolDn = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.PolDn
		varNiatelemetryPodSnmpPolicies.RecordType = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RecordType
		varNiatelemetryPodSnmpPolicies.RecordVersion = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryPodSnmpPolicies.SiteName = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SiteName
		varNiatelemetryPodSnmpPolicies.SnmpClientGrp = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpClientGrp
		varNiatelemetryPodSnmpPolicies.SnmpCommunity = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpCommunity
		varNiatelemetryPodSnmpPolicies.SnmpTrapFwdServer = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpTrapFwdServer
		varNiatelemetryPodSnmpPolicies.SnmpUser = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.SnmpUser
		varNiatelemetryPodSnmpPolicies.RegisteredDevice = varNiatelemetryPodSnmpPoliciesWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryPodSnmpPolicies(varNiatelemetryPodSnmpPolicies)
	} else {
		return err
	}

	varNiatelemetryPodSnmpPolicies := _NiatelemetryPodSnmpPolicies{}

	err = json.Unmarshal(bytes, &varNiatelemetryPodSnmpPolicies)
	if err == nil {
		o.MoBaseMo = varNiatelemetryPodSnmpPolicies.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSt")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SnmpClientGrp")
		delete(additionalProperties, "SnmpCommunity")
		delete(additionalProperties, "SnmpTrapFwdServer")
		delete(additionalProperties, "SnmpUser")
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

type NullableNiatelemetryPodSnmpPolicies struct {
	value *NiatelemetryPodSnmpPolicies
	isSet bool
}

func (v NullableNiatelemetryPodSnmpPolicies) Get() *NiatelemetryPodSnmpPolicies {
	return v.value
}

func (v *NullableNiatelemetryPodSnmpPolicies) Set(val *NiatelemetryPodSnmpPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodSnmpPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodSnmpPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodSnmpPolicies(val *NiatelemetryPodSnmpPolicies) *NullableNiatelemetryPodSnmpPolicies {
	return &NullableNiatelemetryPodSnmpPolicies{value: val, isSet: true}
}

func (v NullableNiatelemetryPodSnmpPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodSnmpPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
