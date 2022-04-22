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
)

// NiatelemetryApicFexDetailsAllOf Definition of the list of properties defined in 'niatelemetry.ApicFexDetails', excluding properties defined in parent classes.
type NiatelemetryApicFexDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn for the FEX in the inventory.
	Dn *string `json:"Dn,omitempty"`
	// Model number of the FEX in APIC.
	ModelNumber *string `json:"ModelNumber,omitempty"`
	// Node id for the FEX in the inventory.
	NodeId *int64 `json:"NodeId,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Serial number of the FEX in APIC.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName *string `json:"SiteName,omitempty"`
	// Vendor name of the FEX in APIC.
	VendorName *string `json:"VendorName,omitempty"`
	// VID for the FEX in the inventory.
	Vid                  *string                              `json:"Vid,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicFexDetailsAllOf NiatelemetryApicFexDetailsAllOf

// NewNiatelemetryApicFexDetailsAllOf instantiates a new NiatelemetryApicFexDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicFexDetailsAllOf(classId string, objectType string) *NiatelemetryApicFexDetailsAllOf {
	this := NiatelemetryApicFexDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicFexDetailsAllOfWithDefaults instantiates a new NiatelemetryApicFexDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicFexDetailsAllOfWithDefaults() *NiatelemetryApicFexDetailsAllOf {
	this := NiatelemetryApicFexDetailsAllOf{}
	var classId string = "niatelemetry.ApicFexDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicFexDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicFexDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicFexDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicFexDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicFexDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicFexDetailsAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasModelNumber() bool {
	if o != nil && o.ModelNumber != nil {
		return true
	}

	return false
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *NiatelemetryApicFexDetailsAllOf) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetNodeId() int64 {
	if o == nil || o.NodeId == nil {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetNodeIdOk() (*int64, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *NiatelemetryApicFexDetailsAllOf) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicFexDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicFexDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *NiatelemetryApicFexDetailsAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicFexDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetVendorName() string {
	if o == nil || o.VendorName == nil {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetVendorNameOk() (*string, bool) {
	if o == nil || o.VendorName == nil {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasVendorName() bool {
	if o != nil && o.VendorName != nil {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *NiatelemetryApicFexDetailsAllOf) SetVendorName(v string) {
	o.VendorName = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *NiatelemetryApicFexDetailsAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicFexDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicFexDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicFexDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicFexDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicFexDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.ModelNumber != nil {
		toSerialize["ModelNumber"] = o.ModelNumber
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.VendorName != nil {
		toSerialize["VendorName"] = o.VendorName
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicFexDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicFexDetailsAllOf := _NiatelemetryApicFexDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicFexDetailsAllOf); err == nil {
		*o = NiatelemetryApicFexDetailsAllOf(varNiatelemetryApicFexDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "ModelNumber")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "VendorName")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicFexDetailsAllOf struct {
	value *NiatelemetryApicFexDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryApicFexDetailsAllOf) Get() *NiatelemetryApicFexDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicFexDetailsAllOf) Set(val *NiatelemetryApicFexDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicFexDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicFexDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicFexDetailsAllOf(val *NiatelemetryApicFexDetailsAllOf) *NullableNiatelemetryApicFexDetailsAllOf {
	return &NullableNiatelemetryApicFexDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicFexDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicFexDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
