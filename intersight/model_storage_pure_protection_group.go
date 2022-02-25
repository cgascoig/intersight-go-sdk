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
	"reflect"
	"strings"
)

// StoragePureProtectionGroup Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
type StoragePureProtectionGroup struct {
	StorageBaseProtectionGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Overall size of all snapshots in the protection group, represented in bytes.
	Size *int64 `json:"Size,omitempty"`
	// Name of PureStorage array name on which the protection group is created.
	Source  *string                       `json:"Source,omitempty"`
	Targets []string                      `json:"Targets,omitempty"`
	Array   *StoragePureArrayRelationship `json:"Array,omitempty"`
	// An array of relationships to storagePureHostGroup resources.
	HostGroups []StoragePureHostGroupRelationship `json:"HostGroups,omitempty"`
	// An array of relationships to storagePureHost resources.
	Hosts            []StoragePureHostRelationship        `json:"Hosts,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storagePureVolume resources.
	Volumes              []StoragePureVolumeRelationship `json:"Volumes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureProtectionGroup StoragePureProtectionGroup

// NewStoragePureProtectionGroup instantiates a new StoragePureProtectionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureProtectionGroup(classId string, objectType string) *StoragePureProtectionGroup {
	this := StoragePureProtectionGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureProtectionGroupWithDefaults instantiates a new StoragePureProtectionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureProtectionGroupWithDefaults() *StoragePureProtectionGroup {
	this := StoragePureProtectionGroup{}
	var classId string = "storage.PureProtectionGroup"
	this.ClassId = classId
	var objectType string = "storage.PureProtectionGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureProtectionGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureProtectionGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureProtectionGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureProtectionGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StoragePureProtectionGroup) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StoragePureProtectionGroup) SetSize(v int64) {
	o.Size = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StoragePureProtectionGroup) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StoragePureProtectionGroup) SetSource(v string) {
	o.Source = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroup) GetTargets() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroup) GetTargetsOk() (*[]string, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return &o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *StoragePureProtectionGroup) SetTargets(v []string) {
	o.Targets = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureProtectionGroup) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureProtectionGroup) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHostGroups returns the HostGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroup) GetHostGroups() []StoragePureHostGroupRelationship {
	if o == nil {
		var ret []StoragePureHostGroupRelationship
		return ret
	}
	return o.HostGroups
}

// GetHostGroupsOk returns a tuple with the HostGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroup) GetHostGroupsOk() (*[]StoragePureHostGroupRelationship, bool) {
	if o == nil || o.HostGroups == nil {
		return nil, false
	}
	return &o.HostGroups, true
}

// HasHostGroups returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasHostGroups() bool {
	if o != nil && o.HostGroups != nil {
		return true
	}

	return false
}

// SetHostGroups gets a reference to the given []StoragePureHostGroupRelationship and assigns it to the HostGroups field.
func (o *StoragePureProtectionGroup) SetHostGroups(v []StoragePureHostGroupRelationship) {
	o.HostGroups = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroup) GetHosts() []StoragePureHostRelationship {
	if o == nil {
		var ret []StoragePureHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroup) GetHostsOk() (*[]StoragePureHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []StoragePureHostRelationship and assigns it to the Hosts field.
func (o *StoragePureProtectionGroup) SetHosts(v []StoragePureHostRelationship) {
	o.Hosts = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureProtectionGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureProtectionGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureProtectionGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureProtectionGroup) GetVolumes() []StoragePureVolumeRelationship {
	if o == nil {
		var ret []StoragePureVolumeRelationship
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureProtectionGroup) GetVolumesOk() (*[]StoragePureVolumeRelationship, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *StoragePureProtectionGroup) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []StoragePureVolumeRelationship and assigns it to the Volumes field.
func (o *StoragePureProtectionGroup) SetVolumes(v []StoragePureVolumeRelationship) {
	o.Volumes = v
}

func (o StoragePureProtectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseProtectionGroup, errStorageBaseProtectionGroup := json.Marshal(o.StorageBaseProtectionGroup)
	if errStorageBaseProtectionGroup != nil {
		return []byte{}, errStorageBaseProtectionGroup
	}
	errStorageBaseProtectionGroup = json.Unmarshal([]byte(serializedStorageBaseProtectionGroup), &toSerialize)
	if errStorageBaseProtectionGroup != nil {
		return []byte{}, errStorageBaseProtectionGroup
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["Source"] = o.Source
	}
	if o.Targets != nil {
		toSerialize["Targets"] = o.Targets
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.HostGroups != nil {
		toSerialize["HostGroups"] = o.HostGroups
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureProtectionGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureProtectionGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Overall size of all snapshots in the protection group, represented in bytes.
		Size *int64 `json:"Size,omitempty"`
		// Name of PureStorage array name on which the protection group is created.
		Source  *string                       `json:"Source,omitempty"`
		Targets []string                      `json:"Targets,omitempty"`
		Array   *StoragePureArrayRelationship `json:"Array,omitempty"`
		// An array of relationships to storagePureHostGroup resources.
		HostGroups []StoragePureHostGroupRelationship `json:"HostGroups,omitempty"`
		// An array of relationships to storagePureHost resources.
		Hosts            []StoragePureHostRelationship        `json:"Hosts,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to storagePureVolume resources.
		Volumes []StoragePureVolumeRelationship `json:"Volumes,omitempty"`
	}

	varStoragePureProtectionGroupWithoutEmbeddedStruct := StoragePureProtectionGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureProtectionGroupWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureProtectionGroup := _StoragePureProtectionGroup{}
		varStoragePureProtectionGroup.ClassId = varStoragePureProtectionGroupWithoutEmbeddedStruct.ClassId
		varStoragePureProtectionGroup.ObjectType = varStoragePureProtectionGroupWithoutEmbeddedStruct.ObjectType
		varStoragePureProtectionGroup.Size = varStoragePureProtectionGroupWithoutEmbeddedStruct.Size
		varStoragePureProtectionGroup.Source = varStoragePureProtectionGroupWithoutEmbeddedStruct.Source
		varStoragePureProtectionGroup.Targets = varStoragePureProtectionGroupWithoutEmbeddedStruct.Targets
		varStoragePureProtectionGroup.Array = varStoragePureProtectionGroupWithoutEmbeddedStruct.Array
		varStoragePureProtectionGroup.HostGroups = varStoragePureProtectionGroupWithoutEmbeddedStruct.HostGroups
		varStoragePureProtectionGroup.Hosts = varStoragePureProtectionGroupWithoutEmbeddedStruct.Hosts
		varStoragePureProtectionGroup.RegisteredDevice = varStoragePureProtectionGroupWithoutEmbeddedStruct.RegisteredDevice
		varStoragePureProtectionGroup.Volumes = varStoragePureProtectionGroupWithoutEmbeddedStruct.Volumes
		*o = StoragePureProtectionGroup(varStoragePureProtectionGroup)
	} else {
		return err
	}

	varStoragePureProtectionGroup := _StoragePureProtectionGroup{}

	err = json.Unmarshal(bytes, &varStoragePureProtectionGroup)
	if err == nil {
		o.StorageBaseProtectionGroup = varStoragePureProtectionGroup.StorageBaseProtectionGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Source")
		delete(additionalProperties, "Targets")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "HostGroups")
		delete(additionalProperties, "Hosts")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Volumes")

		// remove fields from embedded structs
		reflectStorageBaseProtectionGroup := reflect.ValueOf(o.StorageBaseProtectionGroup)
		for i := 0; i < reflectStorageBaseProtectionGroup.Type().NumField(); i++ {
			t := reflectStorageBaseProtectionGroup.Type().Field(i)

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

type NullableStoragePureProtectionGroup struct {
	value *StoragePureProtectionGroup
	isSet bool
}

func (v NullableStoragePureProtectionGroup) Get() *StoragePureProtectionGroup {
	return v.value
}

func (v *NullableStoragePureProtectionGroup) Set(val *StoragePureProtectionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureProtectionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureProtectionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureProtectionGroup(val *StoragePureProtectionGroup) *NullableStoragePureProtectionGroup {
	return &NullableStoragePureProtectionGroup{value: val, isSet: true}
}

func (v NullableStoragePureProtectionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureProtectionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
