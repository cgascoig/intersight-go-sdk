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
	"reflect"
	"strings"
)

// NiatelemetryApicSnmpClientGrpDetails Object to capture the SNMP client grp details in APIC.
type NiatelemetryApicSnmpClientGrpDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the SNMP community in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Name of SNMP client grp in APIC.
	Name *string `json:"Name,omitempty"`
	// Dn of the parent SNMP Policy in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// List of address of restricted clients for particular client grp.
	RestrictedClients *string `json:"RestrictedClients,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicSnmpClientGrpDetails NiatelemetryApicSnmpClientGrpDetails

// NewNiatelemetryApicSnmpClientGrpDetails instantiates a new NiatelemetryApicSnmpClientGrpDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicSnmpClientGrpDetails(classId string, objectType string) *NiatelemetryApicSnmpClientGrpDetails {
	this := NiatelemetryApicSnmpClientGrpDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicSnmpClientGrpDetailsWithDefaults instantiates a new NiatelemetryApicSnmpClientGrpDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicSnmpClientGrpDetailsWithDefaults() *NiatelemetryApicSnmpClientGrpDetails {
	this := NiatelemetryApicSnmpClientGrpDetails{}
	var classId string = "niatelemetry.ApicSnmpClientGrpDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicSnmpClientGrpDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicSnmpClientGrpDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicSnmpClientGrpDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicSnmpClientGrpDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicSnmpClientGrpDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetDn(v string) {
	o.Dn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetName(v string) {
	o.Name = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetRestrictedClients returns the RestrictedClients field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRestrictedClients() string {
	if o == nil || o.RestrictedClients == nil {
		var ret string
		return ret
	}
	return *o.RestrictedClients
}

// GetRestrictedClientsOk returns a tuple with the RestrictedClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRestrictedClientsOk() (*string, bool) {
	if o == nil || o.RestrictedClients == nil {
		return nil, false
	}
	return o.RestrictedClients, true
}

// HasRestrictedClients returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasRestrictedClients() bool {
	if o != nil && o.RestrictedClients != nil {
		return true
	}

	return false
}

// SetRestrictedClients gets a reference to the given string and assigns it to the RestrictedClients field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetRestrictedClients(v string) {
	o.RestrictedClients = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicSnmpClientGrpDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicSnmpClientGrpDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicSnmpClientGrpDetails) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
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
	if o.RestrictedClients != nil {
		toSerialize["RestrictedClients"] = o.RestrictedClients
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

func (o *NiatelemetryApicSnmpClientGrpDetails) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the SNMP community in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Name of SNMP client grp in APIC.
		Name *string `json:"Name,omitempty"`
		// Dn of the parent SNMP Policy in APIC.
		PolDn *string `json:"PolDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// List of address of restricted clients for particular client grp.
		RestrictedClients *string `json:"RestrictedClients,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct := NiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicSnmpClientGrpDetails := _NiatelemetryApicSnmpClientGrpDetails{}
		varNiatelemetryApicSnmpClientGrpDetails.ClassId = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicSnmpClientGrpDetails.ObjectType = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicSnmpClientGrpDetails.Dn = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicSnmpClientGrpDetails.Name = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.Name
		varNiatelemetryApicSnmpClientGrpDetails.PolDn = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.PolDn
		varNiatelemetryApicSnmpClientGrpDetails.RecordType = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicSnmpClientGrpDetails.RecordVersion = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicSnmpClientGrpDetails.RestrictedClients = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.RestrictedClients
		varNiatelemetryApicSnmpClientGrpDetails.SiteName = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicSnmpClientGrpDetails.RegisteredDevice = varNiatelemetryApicSnmpClientGrpDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicSnmpClientGrpDetails(varNiatelemetryApicSnmpClientGrpDetails)
	} else {
		return err
	}

	varNiatelemetryApicSnmpClientGrpDetails := _NiatelemetryApicSnmpClientGrpDetails{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicSnmpClientGrpDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicSnmpClientGrpDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "RestrictedClients")
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

type NullableNiatelemetryApicSnmpClientGrpDetails struct {
	value *NiatelemetryApicSnmpClientGrpDetails
	isSet bool
}

func (v NullableNiatelemetryApicSnmpClientGrpDetails) Get() *NiatelemetryApicSnmpClientGrpDetails {
	return v.value
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetails) Set(val *NiatelemetryApicSnmpClientGrpDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicSnmpClientGrpDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicSnmpClientGrpDetails(val *NiatelemetryApicSnmpClientGrpDetails) *NullableNiatelemetryApicSnmpClientGrpDetails {
	return &NullableNiatelemetryApicSnmpClientGrpDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryApicSnmpClientGrpDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicSnmpClientGrpDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
