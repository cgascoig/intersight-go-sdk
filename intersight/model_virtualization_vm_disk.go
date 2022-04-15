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

// VirtualizationVmDisk Virtual machine disk information.
type VirtualizationVmDisk struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Boot order for this disk.
	BootOrder *int64 `json:"BootOrder,omitempty"`
	// Virtual machine bridge name of interface. * `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration. * `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives. * `scsi` - SCSI (Small Computer System Interface) bus used..
	Bus *string `json:"Bus,omitempty"`
	// Disk name associated with virtual machine.
	Name *string `json:"Name,omitempty"`
	// Type of the Disk (hdd, cdrom). * `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image. * `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image.
	Type        *string                           `json:"Type,omitempty"`
	VirtualDisk NullableVirtualizationVdiskConfig `json:"VirtualDisk,omitempty"`
	// Virtual disk reference name.
	VirtualDiskReference *string `json:"VirtualDiskReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmDisk VirtualizationVmDisk

// NewVirtualizationVmDisk instantiates a new VirtualizationVmDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmDisk(classId string, objectType string) *VirtualizationVmDisk {
	this := VirtualizationVmDisk{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmDiskWithDefaults instantiates a new VirtualizationVmDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmDiskWithDefaults() *VirtualizationVmDisk {
	this := VirtualizationVmDisk{}
	var classId string = "virtualization.VmDisk"
	this.ClassId = classId
	var objectType string = "virtualization.VmDisk"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmDisk) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmDisk) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmDisk) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmDisk) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootOrder returns the BootOrder field value if set, zero value otherwise.
func (o *VirtualizationVmDisk) GetBootOrder() int64 {
	if o == nil || o.BootOrder == nil {
		var ret int64
		return ret
	}
	return *o.BootOrder
}

// GetBootOrderOk returns a tuple with the BootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetBootOrderOk() (*int64, bool) {
	if o == nil || o.BootOrder == nil {
		return nil, false
	}
	return o.BootOrder, true
}

// HasBootOrder returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasBootOrder() bool {
	if o != nil && o.BootOrder != nil {
		return true
	}

	return false
}

// SetBootOrder gets a reference to the given int64 and assigns it to the BootOrder field.
func (o *VirtualizationVmDisk) SetBootOrder(v int64) {
	o.BootOrder = &v
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *VirtualizationVmDisk) GetBus() string {
	if o == nil || o.Bus == nil {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetBusOk() (*string, bool) {
	if o == nil || o.Bus == nil {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasBus() bool {
	if o != nil && o.Bus != nil {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *VirtualizationVmDisk) SetBus(v string) {
	o.Bus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVmDisk) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVmDisk) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualizationVmDisk) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualizationVmDisk) SetType(v string) {
	o.Type = &v
}

// GetVirtualDisk returns the VirtualDisk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmDisk) GetVirtualDisk() VirtualizationVdiskConfig {
	if o == nil || o.VirtualDisk.Get() == nil {
		var ret VirtualizationVdiskConfig
		return ret
	}
	return *o.VirtualDisk.Get()
}

// GetVirtualDiskOk returns a tuple with the VirtualDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmDisk) GetVirtualDiskOk() (*VirtualizationVdiskConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDisk.Get(), o.VirtualDisk.IsSet()
}

// HasVirtualDisk returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasVirtualDisk() bool {
	if o != nil && o.VirtualDisk.IsSet() {
		return true
	}

	return false
}

// SetVirtualDisk gets a reference to the given NullableVirtualizationVdiskConfig and assigns it to the VirtualDisk field.
func (o *VirtualizationVmDisk) SetVirtualDisk(v VirtualizationVdiskConfig) {
	o.VirtualDisk.Set(&v)
}

// SetVirtualDiskNil sets the value for VirtualDisk to be an explicit nil
func (o *VirtualizationVmDisk) SetVirtualDiskNil() {
	o.VirtualDisk.Set(nil)
}

// UnsetVirtualDisk ensures that no value is present for VirtualDisk, not even an explicit nil
func (o *VirtualizationVmDisk) UnsetVirtualDisk() {
	o.VirtualDisk.Unset()
}

// GetVirtualDiskReference returns the VirtualDiskReference field value if set, zero value otherwise.
func (o *VirtualizationVmDisk) GetVirtualDiskReference() string {
	if o == nil || o.VirtualDiskReference == nil {
		var ret string
		return ret
	}
	return *o.VirtualDiskReference
}

// GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmDisk) GetVirtualDiskReferenceOk() (*string, bool) {
	if o == nil || o.VirtualDiskReference == nil {
		return nil, false
	}
	return o.VirtualDiskReference, true
}

// HasVirtualDiskReference returns a boolean if a field has been set.
func (o *VirtualizationVmDisk) HasVirtualDiskReference() bool {
	if o != nil && o.VirtualDiskReference != nil {
		return true
	}

	return false
}

// SetVirtualDiskReference gets a reference to the given string and assigns it to the VirtualDiskReference field.
func (o *VirtualizationVmDisk) SetVirtualDiskReference(v string) {
	o.VirtualDiskReference = &v
}

func (o VirtualizationVmDisk) MarshalJSON() ([]byte, error) {
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
	if o.BootOrder != nil {
		toSerialize["BootOrder"] = o.BootOrder
	}
	if o.Bus != nil {
		toSerialize["Bus"] = o.Bus
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VirtualDisk.IsSet() {
		toSerialize["VirtualDisk"] = o.VirtualDisk.Get()
	}
	if o.VirtualDiskReference != nil {
		toSerialize["VirtualDiskReference"] = o.VirtualDiskReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmDisk) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmDiskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Boot order for this disk.
		BootOrder *int64 `json:"BootOrder,omitempty"`
		// Virtual machine bridge name of interface. * `virtio` - Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration. * `sata` - Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives. * `scsi` - SCSI (Small Computer System Interface) bus used..
		Bus *string `json:"Bus,omitempty"`
		// Disk name associated with virtual machine.
		Name *string `json:"Name,omitempty"`
		// Type of the Disk (hdd, cdrom). * `hdd` - Allows the virtual machine to mount disk from hard disk drive (hdd) image. * `cdrom` - Allows the virtual machine to mount disk from compact disk (cd) image.
		Type        *string                           `json:"Type,omitempty"`
		VirtualDisk NullableVirtualizationVdiskConfig `json:"VirtualDisk,omitempty"`
		// Virtual disk reference name.
		VirtualDiskReference *string `json:"VirtualDiskReference,omitempty"`
	}

	varVirtualizationVmDiskWithoutEmbeddedStruct := VirtualizationVmDiskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmDiskWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmDisk := _VirtualizationVmDisk{}
		varVirtualizationVmDisk.ClassId = varVirtualizationVmDiskWithoutEmbeddedStruct.ClassId
		varVirtualizationVmDisk.ObjectType = varVirtualizationVmDiskWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmDisk.BootOrder = varVirtualizationVmDiskWithoutEmbeddedStruct.BootOrder
		varVirtualizationVmDisk.Bus = varVirtualizationVmDiskWithoutEmbeddedStruct.Bus
		varVirtualizationVmDisk.Name = varVirtualizationVmDiskWithoutEmbeddedStruct.Name
		varVirtualizationVmDisk.Type = varVirtualizationVmDiskWithoutEmbeddedStruct.Type
		varVirtualizationVmDisk.VirtualDisk = varVirtualizationVmDiskWithoutEmbeddedStruct.VirtualDisk
		varVirtualizationVmDisk.VirtualDiskReference = varVirtualizationVmDiskWithoutEmbeddedStruct.VirtualDiskReference
		*o = VirtualizationVmDisk(varVirtualizationVmDisk)
	} else {
		return err
	}

	varVirtualizationVmDisk := _VirtualizationVmDisk{}

	err = json.Unmarshal(bytes, &varVirtualizationVmDisk)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationVmDisk.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BootOrder")
		delete(additionalProperties, "Bus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VirtualDisk")
		delete(additionalProperties, "VirtualDiskReference")

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

type NullableVirtualizationVmDisk struct {
	value *VirtualizationVmDisk
	isSet bool
}

func (v NullableVirtualizationVmDisk) Get() *VirtualizationVmDisk {
	return v.value
}

func (v *NullableVirtualizationVmDisk) Set(val *VirtualizationVmDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmDisk(val *VirtualizationVmDisk) *NullableVirtualizationVmDisk {
	return &NullableVirtualizationVmDisk{value: val, isSet: true}
}

func (v NullableVirtualizationVmDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
