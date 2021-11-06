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
)

// VnicRoceSettingsAllOf Definition of the list of properties defined in 'vnic.RoceSettings', excluding properties defined in parent classes.
type VnicRoceSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Class of Service for RoCE on this virtual interface. * `5` - RDMA CoS Service Level 5. * `1` - RDMA CoS Service Level 1. * `2` - RDMA CoS Service Level 2. * `4` - RDMA CoS Service Level 4. * `6` - RDMA CoS Service Level 6.
	ClassOfService *int32 `json:"ClassOfService,omitempty"`
	// If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface.
	Enabled *bool `json:"Enabled,omitempty"`
	// The number of memory regions per adapter. Recommended value = integer power of 2.
	MemoryRegions *int64 `json:"MemoryRegions,omitempty"`
	// The number of queue pairs per adapter. Recommended value = integer power of 2.
	QueuePairs *int64 `json:"QueuePairs,omitempty"`
	// The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance.
	ResourceGroups *int64 `json:"ResourceGroups,omitempty"`
	// Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters. * `1` - RDMA over Converged Ethernet Protocol Version 1. * `2` - RDMA over Converged Ethernet Protocol Version 2.
	Version              *int32 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicRoceSettingsAllOf VnicRoceSettingsAllOf

// NewVnicRoceSettingsAllOf instantiates a new VnicRoceSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicRoceSettingsAllOf(classId string, objectType string) *VnicRoceSettingsAllOf {
	this := VnicRoceSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var classOfService int32 = 5
	this.ClassOfService = &classOfService
	var memoryRegions int64 = 131072
	this.MemoryRegions = &memoryRegions
	var queuePairs int64 = 256
	this.QueuePairs = &queuePairs
	var resourceGroups int64 = 4
	this.ResourceGroups = &resourceGroups
	var version int32 = 1
	this.Version = &version
	return &this
}

// NewVnicRoceSettingsAllOfWithDefaults instantiates a new VnicRoceSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicRoceSettingsAllOfWithDefaults() *VnicRoceSettingsAllOf {
	this := VnicRoceSettingsAllOf{}
	var classId string = "vnic.RoceSettings"
	this.ClassId = classId
	var objectType string = "vnic.RoceSettings"
	this.ObjectType = objectType
	var classOfService int32 = 5
	this.ClassOfService = &classOfService
	var memoryRegions int64 = 131072
	this.MemoryRegions = &memoryRegions
	var queuePairs int64 = 256
	this.QueuePairs = &queuePairs
	var resourceGroups int64 = 4
	this.ResourceGroups = &resourceGroups
	var version int32 = 1
	this.Version = &version
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicRoceSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicRoceSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicRoceSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicRoceSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClassOfService returns the ClassOfService field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetClassOfService() int32 {
	if o == nil || o.ClassOfService == nil {
		var ret int32
		return ret
	}
	return *o.ClassOfService
}

// GetClassOfServiceOk returns a tuple with the ClassOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetClassOfServiceOk() (*int32, bool) {
	if o == nil || o.ClassOfService == nil {
		return nil, false
	}
	return o.ClassOfService, true
}

// HasClassOfService returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasClassOfService() bool {
	if o != nil && o.ClassOfService != nil {
		return true
	}

	return false
}

// SetClassOfService gets a reference to the given int32 and assigns it to the ClassOfService field.
func (o *VnicRoceSettingsAllOf) SetClassOfService(v int32) {
	o.ClassOfService = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VnicRoceSettingsAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMemoryRegions returns the MemoryRegions field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetMemoryRegions() int64 {
	if o == nil || o.MemoryRegions == nil {
		var ret int64
		return ret
	}
	return *o.MemoryRegions
}

// GetMemoryRegionsOk returns a tuple with the MemoryRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetMemoryRegionsOk() (*int64, bool) {
	if o == nil || o.MemoryRegions == nil {
		return nil, false
	}
	return o.MemoryRegions, true
}

// HasMemoryRegions returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasMemoryRegions() bool {
	if o != nil && o.MemoryRegions != nil {
		return true
	}

	return false
}

// SetMemoryRegions gets a reference to the given int64 and assigns it to the MemoryRegions field.
func (o *VnicRoceSettingsAllOf) SetMemoryRegions(v int64) {
	o.MemoryRegions = &v
}

// GetQueuePairs returns the QueuePairs field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetQueuePairs() int64 {
	if o == nil || o.QueuePairs == nil {
		var ret int64
		return ret
	}
	return *o.QueuePairs
}

// GetQueuePairsOk returns a tuple with the QueuePairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetQueuePairsOk() (*int64, bool) {
	if o == nil || o.QueuePairs == nil {
		return nil, false
	}
	return o.QueuePairs, true
}

// HasQueuePairs returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasQueuePairs() bool {
	if o != nil && o.QueuePairs != nil {
		return true
	}

	return false
}

// SetQueuePairs gets a reference to the given int64 and assigns it to the QueuePairs field.
func (o *VnicRoceSettingsAllOf) SetQueuePairs(v int64) {
	o.QueuePairs = &v
}

// GetResourceGroups returns the ResourceGroups field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetResourceGroups() int64 {
	if o == nil || o.ResourceGroups == nil {
		var ret int64
		return ret
	}
	return *o.ResourceGroups
}

// GetResourceGroupsOk returns a tuple with the ResourceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetResourceGroupsOk() (*int64, bool) {
	if o == nil || o.ResourceGroups == nil {
		return nil, false
	}
	return o.ResourceGroups, true
}

// HasResourceGroups returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasResourceGroups() bool {
	if o != nil && o.ResourceGroups != nil {
		return true
	}

	return false
}

// SetResourceGroups gets a reference to the given int64 and assigns it to the ResourceGroups field.
func (o *VnicRoceSettingsAllOf) SetResourceGroups(v int64) {
	o.ResourceGroups = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VnicRoceSettingsAllOf) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicRoceSettingsAllOf) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VnicRoceSettingsAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *VnicRoceSettingsAllOf) SetVersion(v int32) {
	o.Version = &v
}

func (o VnicRoceSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClassOfService != nil {
		toSerialize["ClassOfService"] = o.ClassOfService
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.MemoryRegions != nil {
		toSerialize["MemoryRegions"] = o.MemoryRegions
	}
	if o.QueuePairs != nil {
		toSerialize["QueuePairs"] = o.QueuePairs
	}
	if o.ResourceGroups != nil {
		toSerialize["ResourceGroups"] = o.ResourceGroups
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicRoceSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicRoceSettingsAllOf := _VnicRoceSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varVnicRoceSettingsAllOf); err == nil {
		*o = VnicRoceSettingsAllOf(varVnicRoceSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClassOfService")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "MemoryRegions")
		delete(additionalProperties, "QueuePairs")
		delete(additionalProperties, "ResourceGroups")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicRoceSettingsAllOf struct {
	value *VnicRoceSettingsAllOf
	isSet bool
}

func (v NullableVnicRoceSettingsAllOf) Get() *VnicRoceSettingsAllOf {
	return v.value
}

func (v *NullableVnicRoceSettingsAllOf) Set(val *VnicRoceSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicRoceSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicRoceSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicRoceSettingsAllOf(val *VnicRoceSettingsAllOf) *NullableVnicRoceSettingsAllOf {
	return &NullableVnicRoceSettingsAllOf{value: val, isSet: true}
}

func (v NullableVnicRoceSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicRoceSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
