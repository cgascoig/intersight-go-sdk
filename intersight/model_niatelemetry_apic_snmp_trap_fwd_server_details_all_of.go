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
)

// NiatelemetryApicSnmpTrapFwdServerDetailsAllOf Definition of the list of properties defined in 'niatelemetry.ApicSnmpTrapFwdServerDetails', excluding properties defined in parent classes.
type NiatelemetryApicSnmpTrapFwdServerDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Address of SNMP Trap Fwd Server in APIC.
	Address *string `json:"Address,omitempty"`
	// Dn of the SNMP Trap Fwd Server details in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Dn of the parent SNMP Policy in APIC.
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

type _NiatelemetryApicSnmpTrapFwdServerDetailsAllOf NiatelemetryApicSnmpTrapFwdServerDetailsAllOf

// NewNiatelemetryApicSnmpTrapFwdServerDetailsAllOf instantiates a new NiatelemetryApicSnmpTrapFwdServerDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpTrapFwdServerDetailsAllOf(classId string, objectType string) *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf {
	this := NiatelemetryApicSnmpTrapFwdServerDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpTrapFwdServerDetailsAllOfWithDefaults instantiates a new NiatelemetryApicSnmpTrapFwdServerDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpTrapFwdServerDetailsAllOfWithDefaults() *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf {
	this := NiatelemetryApicSnmpTrapFwdServerDetailsAllOf{}
	var classId string = "niatelemetry.ApicSnmpTrapFwdServerDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpTrapFwdServerDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetAddress(v string) {
	o.Address = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
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

func (o *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicSnmpTrapFwdServerDetailsAllOf := _NiatelemetryApicSnmpTrapFwdServerDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpTrapFwdServerDetailsAllOf); err == nil {
		*o = NiatelemetryApicSnmpTrapFwdServerDetailsAllOf(varNiatelemetryApicSnmpTrapFwdServerDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf struct {
	value *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) Get() *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) Set(val *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf(val *NiatelemetryApicSnmpTrapFwdServerDetailsAllOf) *NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf {
	return &NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpTrapFwdServerDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
