/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// NiatelemetryHttpsAclEpgContractMapAllOf Definition of the list of properties defined in 'niatelemetry.HttpsAclEpgContractMap', excluding properties defined in parent classes.
type NiatelemetryHttpsAclEpgContractMapAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of HTTPS ACL contract for APIC.
	ContractName *string `json:"ContractName,omitempty"`
	// Dn of the HTTPS ACL contract provider for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of EPGs of the HTTPS ACL contract for APIC.
	EpgName *string `json:"EpgName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// TDn of the HTTPS ACL contract provider for APIC.
	TargetDn             *string                              `json:"TargetDn,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryHttpsAclEpgContractMapAllOf NiatelemetryHttpsAclEpgContractMapAllOf

// NewNiatelemetryHttpsAclEpgContractMapAllOf instantiates a new NiatelemetryHttpsAclEpgContractMapAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryHttpsAclEpgContractMapAllOf(classId string, objectType string) *NiatelemetryHttpsAclEpgContractMapAllOf {
	this := NiatelemetryHttpsAclEpgContractMapAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryHttpsAclEpgContractMapAllOfWithDefaults instantiates a new NiatelemetryHttpsAclEpgContractMapAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryHttpsAclEpgContractMapAllOfWithDefaults() *NiatelemetryHttpsAclEpgContractMapAllOf {
	this := NiatelemetryHttpsAclEpgContractMapAllOf{}
	var classId string = "niatelemetry.HttpsAclEpgContractMap"
	this.ClassId = classId
	var objectType string = "niatelemetry.HttpsAclEpgContractMap"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetContractName() string {
	if o == nil || o.ContractName == nil {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetContractNameOk() (*string, bool) {
	if o == nil || o.ContractName == nil {
		return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasContractName() bool {
	if o != nil && o.ContractName != nil {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetContractName(v string) {
	o.ContractName = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetEpgName returns the EpgName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetEpgName() string {
	if o == nil || o.EpgName == nil {
		var ret string
		return ret
	}
	return *o.EpgName
}

// GetEpgNameOk returns a tuple with the EpgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetEpgNameOk() (*string, bool) {
	if o == nil || o.EpgName == nil {
		return nil, false
	}
	return o.EpgName, true
}

// HasEpgName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasEpgName() bool {
	if o != nil && o.EpgName != nil {
		return true
	}

	return false
}

// SetEpgName gets a reference to the given string and assigns it to the EpgName field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetEpgName(v string) {
	o.EpgName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetTargetDn returns the TargetDn field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetTargetDn() string {
	if o == nil || o.TargetDn == nil {
		var ret string
		return ret
	}
	return *o.TargetDn
}

// GetTargetDnOk returns a tuple with the TargetDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetTargetDnOk() (*string, bool) {
	if o == nil || o.TargetDn == nil {
		return nil, false
	}
	return o.TargetDn, true
}

// HasTargetDn returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasTargetDn() bool {
	if o != nil && o.TargetDn != nil {
		return true
	}

	return false
}

// SetTargetDn gets a reference to the given string and assigns it to the TargetDn field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetTargetDn(v string) {
	o.TargetDn = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryHttpsAclEpgContractMapAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryHttpsAclEpgContractMapAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ContractName != nil {
		toSerialize["ContractName"] = o.ContractName
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.EpgName != nil {
		toSerialize["EpgName"] = o.EpgName
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
	if o.TargetDn != nil {
		toSerialize["TargetDn"] = o.TargetDn
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryHttpsAclEpgContractMapAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryHttpsAclEpgContractMapAllOf := _NiatelemetryHttpsAclEpgContractMapAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryHttpsAclEpgContractMapAllOf); err == nil {
		*o = NiatelemetryHttpsAclEpgContractMapAllOf(varNiatelemetryHttpsAclEpgContractMapAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ContractName")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "EpgName")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "TargetDn")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryHttpsAclEpgContractMapAllOf struct {
	value *NiatelemetryHttpsAclEpgContractMapAllOf
	isSet bool
}

func (v NullableNiatelemetryHttpsAclEpgContractMapAllOf) Get() *NiatelemetryHttpsAclEpgContractMapAllOf {
	return v.value
}

func (v *NullableNiatelemetryHttpsAclEpgContractMapAllOf) Set(val *NiatelemetryHttpsAclEpgContractMapAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryHttpsAclEpgContractMapAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryHttpsAclEpgContractMapAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryHttpsAclEpgContractMapAllOf(val *NiatelemetryHttpsAclEpgContractMapAllOf) *NullableNiatelemetryHttpsAclEpgContractMapAllOf {
	return &NullableNiatelemetryHttpsAclEpgContractMapAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryHttpsAclEpgContractMapAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryHttpsAclEpgContractMapAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
