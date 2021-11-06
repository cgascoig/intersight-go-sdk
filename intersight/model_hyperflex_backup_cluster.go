/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexBackupCluster BackupCluster object associated with Hyperflex cluster describing the backup related information.
type HyperflexBackupCluster struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Defines the backup source cluster and its references.
	BackupDataStore *string `json:"BackupDataStore,omitempty"`
	// UUID for the cluster to allow lookups across unclaim/reclaim.
	SrcClusterUuid       *string                       `json:"SrcClusterUuid,omitempty"`
	SrcCluster           *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
	TgtCluster           *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexBackupCluster HyperflexBackupCluster

// NewHyperflexBackupCluster instantiates a new HyperflexBackupCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexBackupCluster(classId string, objectType string) *HyperflexBackupCluster {
	this := HyperflexBackupCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexBackupClusterWithDefaults instantiates a new HyperflexBackupCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexBackupClusterWithDefaults() *HyperflexBackupCluster {
	this := HyperflexBackupCluster{}
	var classId string = "hyperflex.BackupCluster"
	this.ClassId = classId
	var objectType string = "hyperflex.BackupCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexBackupCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexBackupCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexBackupCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexBackupCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBackupDataStore returns the BackupDataStore field value if set, zero value otherwise.
func (o *HyperflexBackupCluster) GetBackupDataStore() string {
	if o == nil || o.BackupDataStore == nil {
		var ret string
		return ret
	}
	return *o.BackupDataStore
}

// GetBackupDataStoreOk returns a tuple with the BackupDataStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetBackupDataStoreOk() (*string, bool) {
	if o == nil || o.BackupDataStore == nil {
		return nil, false
	}
	return o.BackupDataStore, true
}

// HasBackupDataStore returns a boolean if a field has been set.
func (o *HyperflexBackupCluster) HasBackupDataStore() bool {
	if o != nil && o.BackupDataStore != nil {
		return true
	}

	return false
}

// SetBackupDataStore gets a reference to the given string and assigns it to the BackupDataStore field.
func (o *HyperflexBackupCluster) SetBackupDataStore(v string) {
	o.BackupDataStore = &v
}

// GetSrcClusterUuid returns the SrcClusterUuid field value if set, zero value otherwise.
func (o *HyperflexBackupCluster) GetSrcClusterUuid() string {
	if o == nil || o.SrcClusterUuid == nil {
		var ret string
		return ret
	}
	return *o.SrcClusterUuid
}

// GetSrcClusterUuidOk returns a tuple with the SrcClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetSrcClusterUuidOk() (*string, bool) {
	if o == nil || o.SrcClusterUuid == nil {
		return nil, false
	}
	return o.SrcClusterUuid, true
}

// HasSrcClusterUuid returns a boolean if a field has been set.
func (o *HyperflexBackupCluster) HasSrcClusterUuid() bool {
	if o != nil && o.SrcClusterUuid != nil {
		return true
	}

	return false
}

// SetSrcClusterUuid gets a reference to the given string and assigns it to the SrcClusterUuid field.
func (o *HyperflexBackupCluster) SetSrcClusterUuid(v string) {
	o.SrcClusterUuid = &v
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexBackupCluster) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexBackupCluster) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexBackupCluster) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexBackupCluster) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexBackupCluster) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexBackupCluster) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexBackupCluster) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

func (o HyperflexBackupCluster) MarshalJSON() ([]byte, error) {
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
	if o.BackupDataStore != nil {
		toSerialize["BackupDataStore"] = o.BackupDataStore
	}
	if o.SrcClusterUuid != nil {
		toSerialize["SrcClusterUuid"] = o.SrcClusterUuid
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexBackupCluster) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexBackupClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Defines the backup source cluster and its references.
		BackupDataStore *string `json:"BackupDataStore,omitempty"`
		// UUID for the cluster to allow lookups across unclaim/reclaim.
		SrcClusterUuid *string                       `json:"SrcClusterUuid,omitempty"`
		SrcCluster     *HyperflexClusterRelationship `json:"SrcCluster,omitempty"`
		TgtCluster     *HyperflexClusterRelationship `json:"TgtCluster,omitempty"`
	}

	varHyperflexBackupClusterWithoutEmbeddedStruct := HyperflexBackupClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexBackupClusterWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexBackupCluster := _HyperflexBackupCluster{}
		varHyperflexBackupCluster.ClassId = varHyperflexBackupClusterWithoutEmbeddedStruct.ClassId
		varHyperflexBackupCluster.ObjectType = varHyperflexBackupClusterWithoutEmbeddedStruct.ObjectType
		varHyperflexBackupCluster.BackupDataStore = varHyperflexBackupClusterWithoutEmbeddedStruct.BackupDataStore
		varHyperflexBackupCluster.SrcClusterUuid = varHyperflexBackupClusterWithoutEmbeddedStruct.SrcClusterUuid
		varHyperflexBackupCluster.SrcCluster = varHyperflexBackupClusterWithoutEmbeddedStruct.SrcCluster
		varHyperflexBackupCluster.TgtCluster = varHyperflexBackupClusterWithoutEmbeddedStruct.TgtCluster
		*o = HyperflexBackupCluster(varHyperflexBackupCluster)
	} else {
		return err
	}

	varHyperflexBackupCluster := _HyperflexBackupCluster{}

	err = json.Unmarshal(bytes, &varHyperflexBackupCluster)
	if err == nil {
		o.MoBaseMo = varHyperflexBackupCluster.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupDataStore")
		delete(additionalProperties, "SrcClusterUuid")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")

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

type NullableHyperflexBackupCluster struct {
	value *HyperflexBackupCluster
	isSet bool
}

func (v NullableHyperflexBackupCluster) Get() *HyperflexBackupCluster {
	return v.value
}

func (v *NullableHyperflexBackupCluster) Set(val *HyperflexBackupCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexBackupCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexBackupCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexBackupCluster(val *HyperflexBackupCluster) *NullableHyperflexBackupCluster {
	return &NullableHyperflexBackupCluster{value: val, isSet: true}
}

func (v NullableHyperflexBackupCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexBackupCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
