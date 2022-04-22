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

// HclDriverImageAllOf Definition of the list of properties defined in 'hcl.DriverImage', excluding properties defined in parent classes.
type HclDriverImageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// URL of the driver ISO images.
	DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * `UCSM` - The server is managed by UCS Manager. * `IMC` - The server is standalone managed by CIMC.
	ManagementType *string `json:"ManagementType,omitempty"`
	// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
	ServerPid            *string `json:"ServerPid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclDriverImageAllOf HclDriverImageAllOf

// NewHclDriverImageAllOf instantiates a new HclDriverImageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclDriverImageAllOf(classId string, objectType string) *HclDriverImageAllOf {
	this := HclDriverImageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var managementType string = "UCSM"
	this.ManagementType = &managementType
	return &this
}

// NewHclDriverImageAllOfWithDefaults instantiates a new HclDriverImageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclDriverImageAllOfWithDefaults() *HclDriverImageAllOf {
	this := HclDriverImageAllOf{}
	var classId string = "hcl.DriverImage"
	this.ClassId = classId
	var objectType string = "hcl.DriverImage"
	this.ObjectType = objectType
	var managementType string = "UCSM"
	this.ManagementType = &managementType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclDriverImageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclDriverImageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclDriverImageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclDriverImageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclDriverImageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclDriverImageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverIsoUrl returns the DriverIsoUrl field value if set, zero value otherwise.
func (o *HclDriverImageAllOf) GetDriverIsoUrl() string {
	if o == nil || o.DriverIsoUrl == nil {
		var ret string
		return ret
	}
	return *o.DriverIsoUrl
}

// GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImageAllOf) GetDriverIsoUrlOk() (*string, bool) {
	if o == nil || o.DriverIsoUrl == nil {
		return nil, false
	}
	return o.DriverIsoUrl, true
}

// HasDriverIsoUrl returns a boolean if a field has been set.
func (o *HclDriverImageAllOf) HasDriverIsoUrl() bool {
	if o != nil && o.DriverIsoUrl != nil {
		return true
	}

	return false
}

// SetDriverIsoUrl gets a reference to the given string and assigns it to the DriverIsoUrl field.
func (o *HclDriverImageAllOf) SetDriverIsoUrl(v string) {
	o.DriverIsoUrl = &v
}

// GetManagementType returns the ManagementType field value if set, zero value otherwise.
func (o *HclDriverImageAllOf) GetManagementType() string {
	if o == nil || o.ManagementType == nil {
		var ret string
		return ret
	}
	return *o.ManagementType
}

// GetManagementTypeOk returns a tuple with the ManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImageAllOf) GetManagementTypeOk() (*string, bool) {
	if o == nil || o.ManagementType == nil {
		return nil, false
	}
	return o.ManagementType, true
}

// HasManagementType returns a boolean if a field has been set.
func (o *HclDriverImageAllOf) HasManagementType() bool {
	if o != nil && o.ManagementType != nil {
		return true
	}

	return false
}

// SetManagementType gets a reference to the given string and assigns it to the ManagementType field.
func (o *HclDriverImageAllOf) SetManagementType(v string) {
	o.ManagementType = &v
}

// GetServerPid returns the ServerPid field value if set, zero value otherwise.
func (o *HclDriverImageAllOf) GetServerPid() string {
	if o == nil || o.ServerPid == nil {
		var ret string
		return ret
	}
	return *o.ServerPid
}

// GetServerPidOk returns a tuple with the ServerPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclDriverImageAllOf) GetServerPidOk() (*string, bool) {
	if o == nil || o.ServerPid == nil {
		return nil, false
	}
	return o.ServerPid, true
}

// HasServerPid returns a boolean if a field has been set.
func (o *HclDriverImageAllOf) HasServerPid() bool {
	if o != nil && o.ServerPid != nil {
		return true
	}

	return false
}

// SetServerPid gets a reference to the given string and assigns it to the ServerPid field.
func (o *HclDriverImageAllOf) SetServerPid(v string) {
	o.ServerPid = &v
}

func (o HclDriverImageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriverIsoUrl != nil {
		toSerialize["DriverIsoUrl"] = o.DriverIsoUrl
	}
	if o.ManagementType != nil {
		toSerialize["ManagementType"] = o.ManagementType
	}
	if o.ServerPid != nil {
		toSerialize["ServerPid"] = o.ServerPid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclDriverImageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclDriverImageAllOf := _HclDriverImageAllOf{}

	if err = json.Unmarshal(bytes, &varHclDriverImageAllOf); err == nil {
		*o = HclDriverImageAllOf(varHclDriverImageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverIsoUrl")
		delete(additionalProperties, "ManagementType")
		delete(additionalProperties, "ServerPid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclDriverImageAllOf struct {
	value *HclDriverImageAllOf
	isSet bool
}

func (v NullableHclDriverImageAllOf) Get() *HclDriverImageAllOf {
	return v.value
}

func (v *NullableHclDriverImageAllOf) Set(val *HclDriverImageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclDriverImageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclDriverImageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclDriverImageAllOf(val *HclDriverImageAllOf) *NullableHclDriverImageAllOf {
	return &NullableHclDriverImageAllOf{value: val, isSet: true}
}

func (v NullableHclDriverImageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclDriverImageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
