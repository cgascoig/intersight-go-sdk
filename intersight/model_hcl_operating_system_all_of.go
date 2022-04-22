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

// HclOperatingSystemAllOf Definition of the list of properties defined in 'hcl.OperatingSystem', excluding properties defined in parent classes.
type HclOperatingSystemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version of the Operating System.
	Version              *string                               `json:"Version,omitempty"`
	Vendor               *HclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclOperatingSystemAllOf HclOperatingSystemAllOf

// NewHclOperatingSystemAllOf instantiates a new HclOperatingSystemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclOperatingSystemAllOf(classId string, objectType string) *HclOperatingSystemAllOf {
	this := HclOperatingSystemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclOperatingSystemAllOfWithDefaults instantiates a new HclOperatingSystemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclOperatingSystemAllOfWithDefaults() *HclOperatingSystemAllOf {
	this := HclOperatingSystemAllOf{}
	var classId string = "hcl.OperatingSystem"
	this.ClassId = classId
	var objectType string = "hcl.OperatingSystem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclOperatingSystemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclOperatingSystemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclOperatingSystemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclOperatingSystemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclOperatingSystemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclOperatingSystemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HclOperatingSystemAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclOperatingSystemAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HclOperatingSystemAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HclOperatingSystemAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *HclOperatingSystemAllOf) GetVendor() HclOperatingSystemVendorRelationship {
	if o == nil || o.Vendor == nil {
		var ret HclOperatingSystemVendorRelationship
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclOperatingSystemAllOf) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *HclOperatingSystemAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given HclOperatingSystemVendorRelationship and assigns it to the Vendor field.
func (o *HclOperatingSystemAllOf) SetVendor(v HclOperatingSystemVendorRelationship) {
	o.Vendor = &v
}

func (o HclOperatingSystemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclOperatingSystemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclOperatingSystemAllOf := _HclOperatingSystemAllOf{}

	if err = json.Unmarshal(bytes, &varHclOperatingSystemAllOf); err == nil {
		*o = HclOperatingSystemAllOf(varHclOperatingSystemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Vendor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclOperatingSystemAllOf struct {
	value *HclOperatingSystemAllOf
	isSet bool
}

func (v NullableHclOperatingSystemAllOf) Get() *HclOperatingSystemAllOf {
	return v.value
}

func (v *NullableHclOperatingSystemAllOf) Set(val *HclOperatingSystemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclOperatingSystemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclOperatingSystemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclOperatingSystemAllOf(val *HclOperatingSystemAllOf) *NullableHclOperatingSystemAllOf {
	return &NullableHclOperatingSystemAllOf{value: val, isSet: true}
}

func (v NullableHclOperatingSystemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclOperatingSystemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
