/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// HyperflexEntityReference Reference to another object possibly managed by another service.
type HyperflexEntityReference struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Configuration number for this reference.
	Confignum *int64 `json:"Confignum,omitempty"`
	// Uuid of the entity for this reference.
	Id *string `json:"Id,omitempty"`
	// Type of entity id for this reference. * `VCMOID` - The entity reference ID type is VC MOID. * `VMBIOSUUID` - The entity reference ID type is VM Bios UUID. * `VMDSPATH` - The entity reference ID type is VM Datastore Path. * `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID. * `VMNAME` - The entity reference ID type is VM Name.
	Idtype *string `json:"Idtype,omitempty"`
	// Name of the entity for this entity reference.
	Name *string `json:"Name,omitempty"`
	// Type of the entity for this entity reference. * `DISK` - This entity type is a disk. * `PNODE` - This entity type is a P node. * `NODE` - This entity type is a node. * `CLUSTER` - This entity type is a cluster. * `DATASTORE` - This entity is a datastore. * `VIRTNODE` - This entity is a HyperFlex virtual node. * `VIRTCLUSTER` - This entity type is a virtual cluster. * `VIRTDATASTORE` - This entity type is a virtual data store. * `VIRTMACHINE` - This entity type is a virtual machine. * `PDISK` - This entity type is a P disk. * `PDATASTORE` - This entity type is a P Datastore. * `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot. * `FOLDER` - This entity type is a folder. * `RESOURCEPOOL` - This entity type is a resource pool. * `FILE` - This entity type is a file. * `VIRTDATACENTER` - This entity type is a virtual data center. * `REPLICATION_APPLIANCE` - This entity type is a replication appliance. * `REPLICATION_JOB` - This entity type is a replication job. * `IP_POOL` - This entity type is an IP Pool. * `REPLICATION_INFO` - This entity type is a replication information. * `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot. * `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot. * `DP_VM_CONFIG` - This entity type is a DP VM Configuration. * `DP_VM` - This entity type is a DP VM. * `DP_VMGROUP` - This entity type is a DP VM Group. * `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point. * `CLUSTER_PAIR` - This entity is a cluster pair. * `HX_TASK` - This entity type is a HyperFlex task. * `ZONE` - This entity type is a zone.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexEntityReference HyperflexEntityReference

// NewHyperflexEntityReference instantiates a new HyperflexEntityReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexEntityReference(classId string, objectType string) *HyperflexEntityReference {
	this := HyperflexEntityReference{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexEntityReferenceWithDefaults instantiates a new HyperflexEntityReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexEntityReferenceWithDefaults() *HyperflexEntityReference {
	this := HyperflexEntityReference{}
	var classId string = "hyperflex.EntityReference"
	this.ClassId = classId
	var objectType string = "hyperflex.EntityReference"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexEntityReference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexEntityReference) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexEntityReference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexEntityReference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfignum returns the Confignum field value if set, zero value otherwise.
func (o *HyperflexEntityReference) GetConfignum() int64 {
	if o == nil || o.Confignum == nil {
		var ret int64
		return ret
	}
	return *o.Confignum
}

// GetConfignumOk returns a tuple with the Confignum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetConfignumOk() (*int64, bool) {
	if o == nil || o.Confignum == nil {
		return nil, false
	}
	return o.Confignum, true
}

// HasConfignum returns a boolean if a field has been set.
func (o *HyperflexEntityReference) HasConfignum() bool {
	if o != nil && o.Confignum != nil {
		return true
	}

	return false
}

// SetConfignum gets a reference to the given int64 and assigns it to the Confignum field.
func (o *HyperflexEntityReference) SetConfignum(v int64) {
	o.Confignum = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HyperflexEntityReference) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HyperflexEntityReference) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HyperflexEntityReference) SetId(v string) {
	o.Id = &v
}

// GetIdtype returns the Idtype field value if set, zero value otherwise.
func (o *HyperflexEntityReference) GetIdtype() string {
	if o == nil || o.Idtype == nil {
		var ret string
		return ret
	}
	return *o.Idtype
}

// GetIdtypeOk returns a tuple with the Idtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetIdtypeOk() (*string, bool) {
	if o == nil || o.Idtype == nil {
		return nil, false
	}
	return o.Idtype, true
}

// HasIdtype returns a boolean if a field has been set.
func (o *HyperflexEntityReference) HasIdtype() bool {
	if o != nil && o.Idtype != nil {
		return true
	}

	return false
}

// SetIdtype gets a reference to the given string and assigns it to the Idtype field.
func (o *HyperflexEntityReference) SetIdtype(v string) {
	o.Idtype = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexEntityReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexEntityReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexEntityReference) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HyperflexEntityReference) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexEntityReference) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HyperflexEntityReference) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HyperflexEntityReference) SetType(v string) {
	o.Type = &v
}

func (o HyperflexEntityReference) MarshalJSON() ([]byte, error) {
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
	if o.Confignum != nil {
		toSerialize["Confignum"] = o.Confignum
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Idtype != nil {
		toSerialize["Idtype"] = o.Idtype
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexEntityReference) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexEntityReferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Configuration number for this reference.
		Confignum *int64 `json:"Confignum,omitempty"`
		// Uuid of the entity for this reference.
		Id *string `json:"Id,omitempty"`
		// Type of entity id for this reference. * `VCMOID` - The entity reference ID type is VC MOID. * `VMBIOSUUID` - The entity reference ID type is VM Bios UUID. * `VMDSPATH` - The entity reference ID type is VM Datastore Path. * `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID. * `VMNAME` - The entity reference ID type is VM Name.
		Idtype *string `json:"Idtype,omitempty"`
		// Name of the entity for this entity reference.
		Name *string `json:"Name,omitempty"`
		// Type of the entity for this entity reference. * `DISK` - This entity type is a disk. * `PNODE` - This entity type is a P node. * `NODE` - This entity type is a node. * `CLUSTER` - This entity type is a cluster. * `DATASTORE` - This entity is a datastore. * `VIRTNODE` - This entity is a HyperFlex virtual node. * `VIRTCLUSTER` - This entity type is a virtual cluster. * `VIRTDATASTORE` - This entity type is a virtual data store. * `VIRTMACHINE` - This entity type is a virtual machine. * `PDISK` - This entity type is a P disk. * `PDATASTORE` - This entity type is a P Datastore. * `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot. * `FOLDER` - This entity type is a folder. * `RESOURCEPOOL` - This entity type is a resource pool. * `FILE` - This entity type is a file. * `VIRTDATACENTER` - This entity type is a virtual data center. * `REPLICATION_APPLIANCE` - This entity type is a replication appliance. * `REPLICATION_JOB` - This entity type is a replication job. * `IP_POOL` - This entity type is an IP Pool. * `REPLICATION_INFO` - This entity type is a replication information. * `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot. * `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot. * `DP_VM_CONFIG` - This entity type is a DP VM Configuration. * `DP_VM` - This entity type is a DP VM. * `DP_VMGROUP` - This entity type is a DP VM Group. * `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point. * `CLUSTER_PAIR` - This entity is a cluster pair. * `HX_TASK` - This entity type is a HyperFlex task. * `ZONE` - This entity type is a zone.
		Type *string `json:"Type,omitempty"`
	}

	varHyperflexEntityReferenceWithoutEmbeddedStruct := HyperflexEntityReferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexEntityReferenceWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexEntityReference := _HyperflexEntityReference{}
		varHyperflexEntityReference.ClassId = varHyperflexEntityReferenceWithoutEmbeddedStruct.ClassId
		varHyperflexEntityReference.ObjectType = varHyperflexEntityReferenceWithoutEmbeddedStruct.ObjectType
		varHyperflexEntityReference.Confignum = varHyperflexEntityReferenceWithoutEmbeddedStruct.Confignum
		varHyperflexEntityReference.Id = varHyperflexEntityReferenceWithoutEmbeddedStruct.Id
		varHyperflexEntityReference.Idtype = varHyperflexEntityReferenceWithoutEmbeddedStruct.Idtype
		varHyperflexEntityReference.Name = varHyperflexEntityReferenceWithoutEmbeddedStruct.Name
		varHyperflexEntityReference.Type = varHyperflexEntityReferenceWithoutEmbeddedStruct.Type
		*o = HyperflexEntityReference(varHyperflexEntityReference)
	} else {
		return err
	}

	varHyperflexEntityReference := _HyperflexEntityReference{}

	err = json.Unmarshal(bytes, &varHyperflexEntityReference)
	if err == nil {
		o.MoBaseComplexType = varHyperflexEntityReference.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Confignum")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Idtype")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")

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

type NullableHyperflexEntityReference struct {
	value *HyperflexEntityReference
	isSet bool
}

func (v NullableHyperflexEntityReference) Get() *HyperflexEntityReference {
	return v.value
}

func (v *NullableHyperflexEntityReference) Set(val *HyperflexEntityReference) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexEntityReference) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexEntityReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexEntityReference(val *HyperflexEntityReference) *NullableHyperflexEntityReference {
	return &NullableHyperflexEntityReference{value: val, isSet: true}
}

func (v NullableHyperflexEntityReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexEntityReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
