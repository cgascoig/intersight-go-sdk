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

// HyperflexSnapshotInfoBrief Reference to the VmSnapshot information.
type HyperflexSnapshotInfoBrief struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                             `json:"ObjectType"`
	ReplicationStatus NullableHyperflexReplicationStatus `json:"ReplicationStatus,omitempty"`
	// Cluster site for this snapshot. * `PRIMARY` - The cluster site for this backup is primary. * `SECONDARY` - The cluster site for this backup is secondary.
	Site                      *string                          `json:"Site,omitempty"`
	SnapshotStatus            NullableHyperflexSnapshotStatus  `json:"SnapshotStatus,omitempty"`
	VmSnapshotEntityReference NullableHyperflexEntityReference `json:"VmSnapshotEntityReference,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _HyperflexSnapshotInfoBrief HyperflexSnapshotInfoBrief

// NewHyperflexSnapshotInfoBrief instantiates a new HyperflexSnapshotInfoBrief object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotInfoBrief(classId string, objectType string) *HyperflexSnapshotInfoBrief {
	this := HyperflexSnapshotInfoBrief{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSnapshotInfoBriefWithDefaults instantiates a new HyperflexSnapshotInfoBrief object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotInfoBriefWithDefaults() *HyperflexSnapshotInfoBrief {
	this := HyperflexSnapshotInfoBrief{}
	var classId string = "hyperflex.SnapshotInfoBrief"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotInfoBrief"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotInfoBrief) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotInfoBrief) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotInfoBrief) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotInfoBrief) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotInfoBrief) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotInfoBrief) SetObjectType(v string) {
	o.ObjectType = v
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotInfoBrief) GetReplicationStatus() HyperflexReplicationStatus {
	if o == nil || o.ReplicationStatus.Get() == nil {
		var ret HyperflexReplicationStatus
		return ret
	}
	return *o.ReplicationStatus.Get()
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotInfoBrief) GetReplicationStatusOk() (*HyperflexReplicationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicationStatus.Get(), o.ReplicationStatus.IsSet()
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotInfoBrief) HasReplicationStatus() bool {
	if o != nil && o.ReplicationStatus.IsSet() {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given NullableHyperflexReplicationStatus and assigns it to the ReplicationStatus field.
func (o *HyperflexSnapshotInfoBrief) SetReplicationStatus(v HyperflexReplicationStatus) {
	o.ReplicationStatus.Set(&v)
}

// SetReplicationStatusNil sets the value for ReplicationStatus to be an explicit nil
func (o *HyperflexSnapshotInfoBrief) SetReplicationStatusNil() {
	o.ReplicationStatus.Set(nil)
}

// UnsetReplicationStatus ensures that no value is present for ReplicationStatus, not even an explicit nil
func (o *HyperflexSnapshotInfoBrief) UnsetReplicationStatus() {
	o.ReplicationStatus.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *HyperflexSnapshotInfoBrief) GetSite() string {
	if o == nil || o.Site == nil {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotInfoBrief) GetSiteOk() (*string, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *HyperflexSnapshotInfoBrief) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *HyperflexSnapshotInfoBrief) SetSite(v string) {
	o.Site = &v
}

// GetSnapshotStatus returns the SnapshotStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotInfoBrief) GetSnapshotStatus() HyperflexSnapshotStatus {
	if o == nil || o.SnapshotStatus.Get() == nil {
		var ret HyperflexSnapshotStatus
		return ret
	}
	return *o.SnapshotStatus.Get()
}

// GetSnapshotStatusOk returns a tuple with the SnapshotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotInfoBrief) GetSnapshotStatusOk() (*HyperflexSnapshotStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotStatus.Get(), o.SnapshotStatus.IsSet()
}

// HasSnapshotStatus returns a boolean if a field has been set.
func (o *HyperflexSnapshotInfoBrief) HasSnapshotStatus() bool {
	if o != nil && o.SnapshotStatus.IsSet() {
		return true
	}

	return false
}

// SetSnapshotStatus gets a reference to the given NullableHyperflexSnapshotStatus and assigns it to the SnapshotStatus field.
func (o *HyperflexSnapshotInfoBrief) SetSnapshotStatus(v HyperflexSnapshotStatus) {
	o.SnapshotStatus.Set(&v)
}

// SetSnapshotStatusNil sets the value for SnapshotStatus to be an explicit nil
func (o *HyperflexSnapshotInfoBrief) SetSnapshotStatusNil() {
	o.SnapshotStatus.Set(nil)
}

// UnsetSnapshotStatus ensures that no value is present for SnapshotStatus, not even an explicit nil
func (o *HyperflexSnapshotInfoBrief) UnsetSnapshotStatus() {
	o.SnapshotStatus.Unset()
}

// GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotInfoBrief) GetVmSnapshotEntityReference() HyperflexEntityReference {
	if o == nil || o.VmSnapshotEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmSnapshotEntityReference.Get()
}

// GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotInfoBrief) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSnapshotEntityReference.Get(), o.VmSnapshotEntityReference.IsSet()
}

// HasVmSnapshotEntityReference returns a boolean if a field has been set.
func (o *HyperflexSnapshotInfoBrief) HasVmSnapshotEntityReference() bool {
	if o != nil && o.VmSnapshotEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmSnapshotEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmSnapshotEntityReference field.
func (o *HyperflexSnapshotInfoBrief) SetVmSnapshotEntityReference(v HyperflexEntityReference) {
	o.VmSnapshotEntityReference.Set(&v)
}

// SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil
func (o *HyperflexSnapshotInfoBrief) SetVmSnapshotEntityReferenceNil() {
	o.VmSnapshotEntityReference.Set(nil)
}

// UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil
func (o *HyperflexSnapshotInfoBrief) UnsetVmSnapshotEntityReference() {
	o.VmSnapshotEntityReference.Unset()
}

func (o HyperflexSnapshotInfoBrief) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ReplicationStatus.IsSet() {
		toSerialize["ReplicationStatus"] = o.ReplicationStatus.Get()
	}
	if o.Site != nil {
		toSerialize["Site"] = o.Site
	}
	if o.SnapshotStatus.IsSet() {
		toSerialize["SnapshotStatus"] = o.SnapshotStatus.Get()
	}
	if o.VmSnapshotEntityReference.IsSet() {
		toSerialize["VmSnapshotEntityReference"] = o.VmSnapshotEntityReference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSnapshotInfoBrief) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSnapshotInfoBriefWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                             `json:"ObjectType"`
		ReplicationStatus NullableHyperflexReplicationStatus `json:"ReplicationStatus,omitempty"`
		// Cluster site for this snapshot. * `PRIMARY` - The cluster site for this backup is primary. * `SECONDARY` - The cluster site for this backup is secondary.
		Site                      *string                          `json:"Site,omitempty"`
		SnapshotStatus            NullableHyperflexSnapshotStatus  `json:"SnapshotStatus,omitempty"`
		VmSnapshotEntityReference NullableHyperflexEntityReference `json:"VmSnapshotEntityReference,omitempty"`
	}

	varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct := HyperflexSnapshotInfoBriefWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSnapshotInfoBrief := _HyperflexSnapshotInfoBrief{}
		varHyperflexSnapshotInfoBrief.ClassId = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.ClassId
		varHyperflexSnapshotInfoBrief.ObjectType = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.ObjectType
		varHyperflexSnapshotInfoBrief.ReplicationStatus = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.ReplicationStatus
		varHyperflexSnapshotInfoBrief.Site = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.Site
		varHyperflexSnapshotInfoBrief.SnapshotStatus = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.SnapshotStatus
		varHyperflexSnapshotInfoBrief.VmSnapshotEntityReference = varHyperflexSnapshotInfoBriefWithoutEmbeddedStruct.VmSnapshotEntityReference
		*o = HyperflexSnapshotInfoBrief(varHyperflexSnapshotInfoBrief)
	} else {
		return err
	}

	varHyperflexSnapshotInfoBrief := _HyperflexSnapshotInfoBrief{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotInfoBrief)
	if err == nil {
		o.MoBaseComplexType = varHyperflexSnapshotInfoBrief.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ReplicationStatus")
		delete(additionalProperties, "Site")
		delete(additionalProperties, "SnapshotStatus")
		delete(additionalProperties, "VmSnapshotEntityReference")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHyperflexSnapshotInfoBrief struct {
	value *HyperflexSnapshotInfoBrief
	isSet bool
}

func (v NullableHyperflexSnapshotInfoBrief) Get() *HyperflexSnapshotInfoBrief {
	return v.value
}

func (v *NullableHyperflexSnapshotInfoBrief) Set(val *HyperflexSnapshotInfoBrief) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotInfoBrief) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotInfoBrief) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotInfoBrief(val *HyperflexSnapshotInfoBrief) *NullableHyperflexSnapshotInfoBrief {
	return &NullableHyperflexSnapshotInfoBrief{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotInfoBrief) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotInfoBrief) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
