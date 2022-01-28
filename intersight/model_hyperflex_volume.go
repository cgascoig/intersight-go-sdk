/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// HyperflexVolume A HyperFlex Volume entity.
type HyperflexVolume struct {
	StorageBaseVolume
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Provisioned Capacity of the volume in bytes.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Client (tenant) ID to which the volume belongs.
	ClientId *string `json:"ClientId,omitempty"`
	// The name of the kubernetes cluster to which the volume is associated.
	KubernetesClusterName *string `json:"KubernetesClusterName,omitempty"`
	// Last modified time as UTC of the volume.
	LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
	// UUID of LUN associated with the volume.
	LunUuid *string `json:"LunUuid,omitempty"`
	// Serial number of the volume.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The unique identifier for this volume.
	Uuid *string `json:"Uuid,omitempty"`
	// Access Mode of the volume. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. * `` - Unknown disk access mode.
	VolumeAccessMode *string `json:"VolumeAccessMode,omitempty"`
	// Volume creation time in UTC.
	VolumeCreateTime *time.Time `json:"VolumeCreateTime,omitempty"`
	// The mode of the HyperFlex volume. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
	VolumeMode *string `json:"VolumeMode,omitempty"`
	// The type of the HyperFlex volume.
	VolumeType           *string                                   `json:"VolumeType,omitempty"`
	Cluster              *HyperflexClusterRelationship             `json:"Cluster,omitempty"`
	IweVirtualDisk       *VirtualizationIweVirtualDiskRelationship `json:"IweVirtualDisk,omitempty"`
	StorageContainer     *HyperflexStorageContainerRelationship    `json:"StorageContainer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVolume HyperflexVolume

// NewHyperflexVolume instantiates a new HyperflexVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVolume(classId string, objectType string) *HyperflexVolume {
	this := HyperflexVolume{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVolumeWithDefaults instantiates a new HyperflexVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVolumeWithDefaults() *HyperflexVolume {
	this := HyperflexVolume{}
	var classId string = "hyperflex.Volume"
	this.ClassId = classId
	var objectType string = "hyperflex.Volume"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVolume) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVolume) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVolume) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVolume) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *HyperflexVolume) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *HyperflexVolume) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *HyperflexVolume) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *HyperflexVolume) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *HyperflexVolume) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *HyperflexVolume) SetClientId(v string) {
	o.ClientId = &v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value if set, zero value otherwise.
func (o *HyperflexVolume) GetKubernetesClusterName() string {
	if o == nil || o.KubernetesClusterName == nil {
		var ret string
		return ret
	}
	return *o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil || o.KubernetesClusterName == nil {
		return nil, false
	}
	return o.KubernetesClusterName, true
}

// HasKubernetesClusterName returns a boolean if a field has been set.
func (o *HyperflexVolume) HasKubernetesClusterName() bool {
	if o != nil && o.KubernetesClusterName != nil {
		return true
	}

	return false
}

// SetKubernetesClusterName gets a reference to the given string and assigns it to the KubernetesClusterName field.
func (o *HyperflexVolume) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *HyperflexVolume) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *HyperflexVolume) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *HyperflexVolume) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetLunUuid returns the LunUuid field value if set, zero value otherwise.
func (o *HyperflexVolume) GetLunUuid() string {
	if o == nil || o.LunUuid == nil {
		var ret string
		return ret
	}
	return *o.LunUuid
}

// GetLunUuidOk returns a tuple with the LunUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetLunUuidOk() (*string, bool) {
	if o == nil || o.LunUuid == nil {
		return nil, false
	}
	return o.LunUuid, true
}

// HasLunUuid returns a boolean if a field has been set.
func (o *HyperflexVolume) HasLunUuid() bool {
	if o != nil && o.LunUuid != nil {
		return true
	}

	return false
}

// SetLunUuid gets a reference to the given string and assigns it to the LunUuid field.
func (o *HyperflexVolume) SetLunUuid(v string) {
	o.LunUuid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HyperflexVolume) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HyperflexVolume) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HyperflexVolume) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *HyperflexVolume) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *HyperflexVolume) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *HyperflexVolume) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeAccessMode returns the VolumeAccessMode field value if set, zero value otherwise.
func (o *HyperflexVolume) GetVolumeAccessMode() string {
	if o == nil || o.VolumeAccessMode == nil {
		var ret string
		return ret
	}
	return *o.VolumeAccessMode
}

// GetVolumeAccessModeOk returns a tuple with the VolumeAccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetVolumeAccessModeOk() (*string, bool) {
	if o == nil || o.VolumeAccessMode == nil {
		return nil, false
	}
	return o.VolumeAccessMode, true
}

// HasVolumeAccessMode returns a boolean if a field has been set.
func (o *HyperflexVolume) HasVolumeAccessMode() bool {
	if o != nil && o.VolumeAccessMode != nil {
		return true
	}

	return false
}

// SetVolumeAccessMode gets a reference to the given string and assigns it to the VolumeAccessMode field.
func (o *HyperflexVolume) SetVolumeAccessMode(v string) {
	o.VolumeAccessMode = &v
}

// GetVolumeCreateTime returns the VolumeCreateTime field value if set, zero value otherwise.
func (o *HyperflexVolume) GetVolumeCreateTime() time.Time {
	if o == nil || o.VolumeCreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.VolumeCreateTime
}

// GetVolumeCreateTimeOk returns a tuple with the VolumeCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetVolumeCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.VolumeCreateTime == nil {
		return nil, false
	}
	return o.VolumeCreateTime, true
}

// HasVolumeCreateTime returns a boolean if a field has been set.
func (o *HyperflexVolume) HasVolumeCreateTime() bool {
	if o != nil && o.VolumeCreateTime != nil {
		return true
	}

	return false
}

// SetVolumeCreateTime gets a reference to the given time.Time and assigns it to the VolumeCreateTime field.
func (o *HyperflexVolume) SetVolumeCreateTime(v time.Time) {
	o.VolumeCreateTime = &v
}

// GetVolumeMode returns the VolumeMode field value if set, zero value otherwise.
func (o *HyperflexVolume) GetVolumeMode() string {
	if o == nil || o.VolumeMode == nil {
		var ret string
		return ret
	}
	return *o.VolumeMode
}

// GetVolumeModeOk returns a tuple with the VolumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetVolumeModeOk() (*string, bool) {
	if o == nil || o.VolumeMode == nil {
		return nil, false
	}
	return o.VolumeMode, true
}

// HasVolumeMode returns a boolean if a field has been set.
func (o *HyperflexVolume) HasVolumeMode() bool {
	if o != nil && o.VolumeMode != nil {
		return true
	}

	return false
}

// SetVolumeMode gets a reference to the given string and assigns it to the VolumeMode field.
func (o *HyperflexVolume) SetVolumeMode(v string) {
	o.VolumeMode = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *HyperflexVolume) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *HyperflexVolume) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *HyperflexVolume) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexVolume) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexVolume) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexVolume) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetIweVirtualDisk returns the IweVirtualDisk field value if set, zero value otherwise.
func (o *HyperflexVolume) GetIweVirtualDisk() VirtualizationIweVirtualDiskRelationship {
	if o == nil || o.IweVirtualDisk == nil {
		var ret VirtualizationIweVirtualDiskRelationship
		return ret
	}
	return *o.IweVirtualDisk
}

// GetIweVirtualDiskOk returns a tuple with the IweVirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetIweVirtualDiskOk() (*VirtualizationIweVirtualDiskRelationship, bool) {
	if o == nil || o.IweVirtualDisk == nil {
		return nil, false
	}
	return o.IweVirtualDisk, true
}

// HasIweVirtualDisk returns a boolean if a field has been set.
func (o *HyperflexVolume) HasIweVirtualDisk() bool {
	if o != nil && o.IweVirtualDisk != nil {
		return true
	}

	return false
}

// SetIweVirtualDisk gets a reference to the given VirtualizationIweVirtualDiskRelationship and assigns it to the IweVirtualDisk field.
func (o *HyperflexVolume) SetIweVirtualDisk(v VirtualizationIweVirtualDiskRelationship) {
	o.IweVirtualDisk = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *HyperflexVolume) GetStorageContainer() HyperflexStorageContainerRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret HyperflexStorageContainerRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVolume) GetStorageContainerOk() (*HyperflexStorageContainerRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *HyperflexVolume) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given HyperflexStorageContainerRelationship and assigns it to the StorageContainer field.
func (o *HyperflexVolume) SetStorageContainer(v HyperflexStorageContainerRelationship) {
	o.StorageContainer = &v
}

func (o HyperflexVolume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseVolume, errStorageBaseVolume := json.Marshal(o.StorageBaseVolume)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	errStorageBaseVolume = json.Unmarshal([]byte(serializedStorageBaseVolume), &toSerialize)
	if errStorageBaseVolume != nil {
		return []byte{}, errStorageBaseVolume
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.ClientId != nil {
		toSerialize["ClientId"] = o.ClientId
	}
	if o.KubernetesClusterName != nil {
		toSerialize["KubernetesClusterName"] = o.KubernetesClusterName
	}
	if o.LastModifiedTime != nil {
		toSerialize["LastModifiedTime"] = o.LastModifiedTime
	}
	if o.LunUuid != nil {
		toSerialize["LunUuid"] = o.LunUuid
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VolumeAccessMode != nil {
		toSerialize["VolumeAccessMode"] = o.VolumeAccessMode
	}
	if o.VolumeCreateTime != nil {
		toSerialize["VolumeCreateTime"] = o.VolumeCreateTime
	}
	if o.VolumeMode != nil {
		toSerialize["VolumeMode"] = o.VolumeMode
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.IweVirtualDisk != nil {
		toSerialize["IweVirtualDisk"] = o.IweVirtualDisk
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVolume) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexVolumeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Provisioned Capacity of the volume in bytes.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Client (tenant) ID to which the volume belongs.
		ClientId *string `json:"ClientId,omitempty"`
		// The name of the kubernetes cluster to which the volume is associated.
		KubernetesClusterName *string `json:"KubernetesClusterName,omitempty"`
		// Last modified time as UTC of the volume.
		LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
		// UUID of LUN associated with the volume.
		LunUuid *string `json:"LunUuid,omitempty"`
		// Serial number of the volume.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The unique identifier for this volume.
		Uuid *string `json:"Uuid,omitempty"`
		// Access Mode of the volume. * `ReadWriteOnce` - Read write permisisons to a Virtual disk by a single virtual machine. * `ReadWriteMany` - Read write permisisons to a Virtual disk by multiple virtual machines. * `ReadOnlyMany` - Read only permisisons to a Virtual disk by multiple virtual machines. * `` - Unknown disk access mode.
		VolumeAccessMode *string `json:"VolumeAccessMode,omitempty"`
		// Volume creation time in UTC.
		VolumeCreateTime *time.Time `json:"VolumeCreateTime,omitempty"`
		// The mode of the HyperFlex volume. * `Block` - It is a Block virtual disk. * `Filesystem` - It is a File system virtual disk. * `` - Disk mode is either unknown or not supported.
		VolumeMode *string `json:"VolumeMode,omitempty"`
		// The type of the HyperFlex volume.
		VolumeType       *string                                   `json:"VolumeType,omitempty"`
		Cluster          *HyperflexClusterRelationship             `json:"Cluster,omitempty"`
		IweVirtualDisk   *VirtualizationIweVirtualDiskRelationship `json:"IweVirtualDisk,omitempty"`
		StorageContainer *HyperflexStorageContainerRelationship    `json:"StorageContainer,omitempty"`
	}

	varHyperflexVolumeWithoutEmbeddedStruct := HyperflexVolumeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexVolumeWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexVolume := _HyperflexVolume{}
		varHyperflexVolume.ClassId = varHyperflexVolumeWithoutEmbeddedStruct.ClassId
		varHyperflexVolume.ObjectType = varHyperflexVolumeWithoutEmbeddedStruct.ObjectType
		varHyperflexVolume.Capacity = varHyperflexVolumeWithoutEmbeddedStruct.Capacity
		varHyperflexVolume.ClientId = varHyperflexVolumeWithoutEmbeddedStruct.ClientId
		varHyperflexVolume.KubernetesClusterName = varHyperflexVolumeWithoutEmbeddedStruct.KubernetesClusterName
		varHyperflexVolume.LastModifiedTime = varHyperflexVolumeWithoutEmbeddedStruct.LastModifiedTime
		varHyperflexVolume.LunUuid = varHyperflexVolumeWithoutEmbeddedStruct.LunUuid
		varHyperflexVolume.SerialNumber = varHyperflexVolumeWithoutEmbeddedStruct.SerialNumber
		varHyperflexVolume.Uuid = varHyperflexVolumeWithoutEmbeddedStruct.Uuid
		varHyperflexVolume.VolumeAccessMode = varHyperflexVolumeWithoutEmbeddedStruct.VolumeAccessMode
		varHyperflexVolume.VolumeCreateTime = varHyperflexVolumeWithoutEmbeddedStruct.VolumeCreateTime
		varHyperflexVolume.VolumeMode = varHyperflexVolumeWithoutEmbeddedStruct.VolumeMode
		varHyperflexVolume.VolumeType = varHyperflexVolumeWithoutEmbeddedStruct.VolumeType
		varHyperflexVolume.Cluster = varHyperflexVolumeWithoutEmbeddedStruct.Cluster
		varHyperflexVolume.IweVirtualDisk = varHyperflexVolumeWithoutEmbeddedStruct.IweVirtualDisk
		varHyperflexVolume.StorageContainer = varHyperflexVolumeWithoutEmbeddedStruct.StorageContainer
		*o = HyperflexVolume(varHyperflexVolume)
	} else {
		return err
	}

	varHyperflexVolume := _HyperflexVolume{}

	err = json.Unmarshal(bytes, &varHyperflexVolume)
	if err == nil {
		o.StorageBaseVolume = varHyperflexVolume.StorageBaseVolume
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "KubernetesClusterName")
		delete(additionalProperties, "LastModifiedTime")
		delete(additionalProperties, "LunUuid")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeAccessMode")
		delete(additionalProperties, "VolumeCreateTime")
		delete(additionalProperties, "VolumeMode")
		delete(additionalProperties, "VolumeType")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "IweVirtualDisk")
		delete(additionalProperties, "StorageContainer")

		// remove fields from embedded structs
		reflectStorageBaseVolume := reflect.ValueOf(o.StorageBaseVolume)
		for i := 0; i < reflectStorageBaseVolume.Type().NumField(); i++ {
			t := reflectStorageBaseVolume.Type().Field(i)

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

type NullableHyperflexVolume struct {
	value *HyperflexVolume
	isSet bool
}

func (v NullableHyperflexVolume) Get() *HyperflexVolume {
	return v.value
}

func (v *NullableHyperflexVolume) Set(val *HyperflexVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVolume(val *HyperflexVolume) *NullableHyperflexVolume {
	return &NullableHyperflexVolume{value: val, isSet: true}
}

func (v NullableHyperflexVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
