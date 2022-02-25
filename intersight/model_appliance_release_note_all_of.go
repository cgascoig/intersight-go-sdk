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
)

// ApplianceReleaseNoteAllOf Definition of the list of properties defined in 'appliance.ReleaseNote', excluding properties defined in parent classes.
type ApplianceReleaseNoteAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string              `json:"ObjectType"`
	Notes      []OnpremUpgradeNote `json:"Notes,omitempty"`
	// Version number of the pending upgrade.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceReleaseNoteAllOf ApplianceReleaseNoteAllOf

// NewApplianceReleaseNoteAllOf instantiates a new ApplianceReleaseNoteAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceReleaseNoteAllOf(classId string, objectType string) *ApplianceReleaseNoteAllOf {
	this := ApplianceReleaseNoteAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceReleaseNoteAllOfWithDefaults instantiates a new ApplianceReleaseNoteAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceReleaseNoteAllOfWithDefaults() *ApplianceReleaseNoteAllOf {
	this := ApplianceReleaseNoteAllOf{}
	var classId string = "appliance.ReleaseNote"
	this.ClassId = classId
	var objectType string = "appliance.ReleaseNote"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceReleaseNoteAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceReleaseNoteAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceReleaseNoteAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceReleaseNoteAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceReleaseNoteAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceReleaseNoteAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceReleaseNoteAllOf) GetNotes() []OnpremUpgradeNote {
	if o == nil {
		var ret []OnpremUpgradeNote
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceReleaseNoteAllOf) GetNotesOk() (*[]OnpremUpgradeNote, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return &o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApplianceReleaseNoteAllOf) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []OnpremUpgradeNote and assigns it to the Notes field.
func (o *ApplianceReleaseNoteAllOf) SetNotes(v []OnpremUpgradeNote) {
	o.Notes = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceReleaseNoteAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceReleaseNoteAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceReleaseNoteAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceReleaseNoteAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o ApplianceReleaseNoteAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Notes != nil {
		toSerialize["Notes"] = o.Notes
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceReleaseNoteAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceReleaseNoteAllOf := _ApplianceReleaseNoteAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceReleaseNoteAllOf); err == nil {
		*o = ApplianceReleaseNoteAllOf(varApplianceReleaseNoteAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Notes")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceReleaseNoteAllOf struct {
	value *ApplianceReleaseNoteAllOf
	isSet bool
}

func (v NullableApplianceReleaseNoteAllOf) Get() *ApplianceReleaseNoteAllOf {
	return v.value
}

func (v *NullableApplianceReleaseNoteAllOf) Set(val *ApplianceReleaseNoteAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceReleaseNoteAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceReleaseNoteAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceReleaseNoteAllOf(val *ApplianceReleaseNoteAllOf) *NullableApplianceReleaseNoteAllOf {
	return &NullableApplianceReleaseNoteAllOf{value: val, isSet: true}
}

func (v NullableApplianceReleaseNoteAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceReleaseNoteAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
