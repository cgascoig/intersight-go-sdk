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
	"reflect"
	"strings"
)

// VirtualizationBaseCluster Common attributes of a cluster of resources within a datacenter. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
type VirtualizationBaseCluster struct {
	ComputeBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType   string                   `json:"ObjectType"`
	AlarmSummary NullableCondAlarmSummary `json:"AlarmSummary,omitempty"`
	// Identifies the broad type of the underlying hypervisor. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// Identifies the version of the hypervisor running on this cluster. Keeping the hypervisor version in the cluster makes it convenient for applications to validate their deployment needs. Defaults to an empty string.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// The internally generated identity of this cluster. This entity is not manipulated by users.
	Identity          *string                               `json:"Identity,omitempty"`
	MemoryCapacity    NullableVirtualizationMemoryCapacity  `json:"MemoryCapacity,omitempty"`
	ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
	// Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster.
	TotalCores           *int64 `json:"TotalCores,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBaseCluster VirtualizationBaseCluster

// NewVirtualizationBaseCluster instantiates a new VirtualizationBaseCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBaseCluster(classId string, objectType string) *VirtualizationBaseCluster {
	this := VirtualizationBaseCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewVirtualizationBaseClusterWithDefaults instantiates a new VirtualizationBaseCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBaseClusterWithDefaults() *VirtualizationBaseCluster {
	this := VirtualizationBaseCluster{}
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBaseCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBaseCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBaseCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBaseCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseCluster) GetAlarmSummary() CondAlarmSummary {
	if o == nil || o.AlarmSummary.Get() == nil {
		var ret CondAlarmSummary
		return ret
	}
	return *o.AlarmSummary.Get()
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseCluster) GetAlarmSummaryOk() (*CondAlarmSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlarmSummary.Get(), o.AlarmSummary.IsSet()
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary.IsSet() {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given NullableCondAlarmSummary and assigns it to the AlarmSummary field.
func (o *VirtualizationBaseCluster) SetAlarmSummary(v CondAlarmSummary) {
	o.AlarmSummary.Set(&v)
}

// SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil
func (o *VirtualizationBaseCluster) SetAlarmSummaryNil() {
	o.AlarmSummary.Set(nil)
}

// UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
func (o *VirtualizationBaseCluster) UnsetAlarmSummary() {
	o.AlarmSummary.Unset()
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationBaseCluster) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetHypervisorVersion() string {
	if o == nil || o.HypervisorVersion == nil {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || o.HypervisorVersion == nil {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasHypervisorVersion() bool {
	if o != nil && o.HypervisorVersion != nil {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *VirtualizationBaseCluster) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationBaseCluster) SetIdentity(v string) {
	o.Identity = &v
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseCluster) GetMemoryCapacity() VirtualizationMemoryCapacity {
	if o == nil || o.MemoryCapacity.Get() == nil {
		var ret VirtualizationMemoryCapacity
		return ret
	}
	return *o.MemoryCapacity.Get()
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseCluster) GetMemoryCapacityOk() (*VirtualizationMemoryCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryCapacity.Get(), o.MemoryCapacity.IsSet()
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasMemoryCapacity() bool {
	if o != nil && o.MemoryCapacity.IsSet() {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given NullableVirtualizationMemoryCapacity and assigns it to the MemoryCapacity field.
func (o *VirtualizationBaseCluster) SetMemoryCapacity(v VirtualizationMemoryCapacity) {
	o.MemoryCapacity.Set(&v)
}

// SetMemoryCapacityNil sets the value for MemoryCapacity to be an explicit nil
func (o *VirtualizationBaseCluster) SetMemoryCapacityNil() {
	o.MemoryCapacity.Set(nil)
}

// UnsetMemoryCapacity ensures that no value is present for MemoryCapacity, not even an explicit nil
func (o *VirtualizationBaseCluster) UnsetMemoryCapacity() {
	o.MemoryCapacity.Unset()
}

// GetProcessorCapacity returns the ProcessorCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationBaseCluster) GetProcessorCapacity() VirtualizationComputeCapacity {
	if o == nil || o.ProcessorCapacity.Get() == nil {
		var ret VirtualizationComputeCapacity
		return ret
	}
	return *o.ProcessorCapacity.Get()
}

// GetProcessorCapacityOk returns a tuple with the ProcessorCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationBaseCluster) GetProcessorCapacityOk() (*VirtualizationComputeCapacity, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessorCapacity.Get(), o.ProcessorCapacity.IsSet()
}

// HasProcessorCapacity returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasProcessorCapacity() bool {
	if o != nil && o.ProcessorCapacity.IsSet() {
		return true
	}

	return false
}

// SetProcessorCapacity gets a reference to the given NullableVirtualizationComputeCapacity and assigns it to the ProcessorCapacity field.
func (o *VirtualizationBaseCluster) SetProcessorCapacity(v VirtualizationComputeCapacity) {
	o.ProcessorCapacity.Set(&v)
}

// SetProcessorCapacityNil sets the value for ProcessorCapacity to be an explicit nil
func (o *VirtualizationBaseCluster) SetProcessorCapacityNil() {
	o.ProcessorCapacity.Set(nil)
}

// UnsetProcessorCapacity ensures that no value is present for ProcessorCapacity, not even an explicit nil
func (o *VirtualizationBaseCluster) UnsetProcessorCapacity() {
	o.ProcessorCapacity.Unset()
}

// GetTotalCores returns the TotalCores field value if set, zero value otherwise.
func (o *VirtualizationBaseCluster) GetTotalCores() int64 {
	if o == nil || o.TotalCores == nil {
		var ret int64
		return ret
	}
	return *o.TotalCores
}

// GetTotalCoresOk returns a tuple with the TotalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBaseCluster) GetTotalCoresOk() (*int64, bool) {
	if o == nil || o.TotalCores == nil {
		return nil, false
	}
	return o.TotalCores, true
}

// HasTotalCores returns a boolean if a field has been set.
func (o *VirtualizationBaseCluster) HasTotalCores() bool {
	if o != nil && o.TotalCores != nil {
		return true
	}

	return false
}

// SetTotalCores gets a reference to the given int64 and assigns it to the TotalCores field.
func (o *VirtualizationBaseCluster) SetTotalCores(v int64) {
	o.TotalCores = &v
}

func (o VirtualizationBaseCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedComputeBaseCluster, errComputeBaseCluster := json.Marshal(o.ComputeBaseCluster)
	if errComputeBaseCluster != nil {
		return []byte{}, errComputeBaseCluster
	}
	errComputeBaseCluster = json.Unmarshal([]byte(serializedComputeBaseCluster), &toSerialize)
	if errComputeBaseCluster != nil {
		return []byte{}, errComputeBaseCluster
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlarmSummary.IsSet() {
		toSerialize["AlarmSummary"] = o.AlarmSummary.Get()
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.HypervisorVersion != nil {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.MemoryCapacity.IsSet() {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity.Get()
	}
	if o.ProcessorCapacity.IsSet() {
		toSerialize["ProcessorCapacity"] = o.ProcessorCapacity.Get()
	}
	if o.TotalCores != nil {
		toSerialize["TotalCores"] = o.TotalCores
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationBaseCluster) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationBaseClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType   string                   `json:"ObjectType"`
		AlarmSummary NullableCondAlarmSummary `json:"AlarmSummary,omitempty"`
		// Identifies the broad type of the underlying hypervisor. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// Identifies the version of the hypervisor running on this cluster. Keeping the hypervisor version in the cluster makes it convenient for applications to validate their deployment needs. Defaults to an empty string.
		HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
		// The internally generated identity of this cluster. This entity is not manipulated by users.
		Identity          *string                               `json:"Identity,omitempty"`
		MemoryCapacity    NullableVirtualizationMemoryCapacity  `json:"MemoryCapacity,omitempty"`
		ProcessorCapacity NullableVirtualizationComputeCapacity `json:"ProcessorCapacity,omitempty"`
		// Total number of CPU cores in this cluster. It is a cumulative number across all hosts in the cluster.
		TotalCores *int64 `json:"TotalCores,omitempty"`
	}

	varVirtualizationBaseClusterWithoutEmbeddedStruct := VirtualizationBaseClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseClusterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBaseCluster := _VirtualizationBaseCluster{}
		varVirtualizationBaseCluster.ClassId = varVirtualizationBaseClusterWithoutEmbeddedStruct.ClassId
		varVirtualizationBaseCluster.ObjectType = varVirtualizationBaseClusterWithoutEmbeddedStruct.ObjectType
		varVirtualizationBaseCluster.AlarmSummary = varVirtualizationBaseClusterWithoutEmbeddedStruct.AlarmSummary
		varVirtualizationBaseCluster.HypervisorType = varVirtualizationBaseClusterWithoutEmbeddedStruct.HypervisorType
		varVirtualizationBaseCluster.HypervisorVersion = varVirtualizationBaseClusterWithoutEmbeddedStruct.HypervisorVersion
		varVirtualizationBaseCluster.Identity = varVirtualizationBaseClusterWithoutEmbeddedStruct.Identity
		varVirtualizationBaseCluster.MemoryCapacity = varVirtualizationBaseClusterWithoutEmbeddedStruct.MemoryCapacity
		varVirtualizationBaseCluster.ProcessorCapacity = varVirtualizationBaseClusterWithoutEmbeddedStruct.ProcessorCapacity
		varVirtualizationBaseCluster.TotalCores = varVirtualizationBaseClusterWithoutEmbeddedStruct.TotalCores
		*o = VirtualizationBaseCluster(varVirtualizationBaseCluster)
	} else {
		return err
	}

	varVirtualizationBaseCluster := _VirtualizationBaseCluster{}

	err = json.Unmarshal(bytes, &varVirtualizationBaseCluster)
	if err == nil {
		o.ComputeBaseCluster = varVirtualizationBaseCluster.ComputeBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "MemoryCapacity")
		delete(additionalProperties, "ProcessorCapacity")
		delete(additionalProperties, "TotalCores")

		// remove fields from embedded structs
		reflectComputeBaseCluster := reflect.ValueOf(o.ComputeBaseCluster)
		for i := 0; i < reflectComputeBaseCluster.Type().NumField(); i++ {
			t := reflectComputeBaseCluster.Type().Field(i)

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

type NullableVirtualizationBaseCluster struct {
	value *VirtualizationBaseCluster
	isSet bool
}

func (v NullableVirtualizationBaseCluster) Get() *VirtualizationBaseCluster {
	return v.value
}

func (v *NullableVirtualizationBaseCluster) Set(val *VirtualizationBaseCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBaseCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBaseCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBaseCluster(val *VirtualizationBaseCluster) *NullableVirtualizationBaseCluster {
	return &NullableVirtualizationBaseCluster{value: val, isSet: true}
}

func (v NullableVirtualizationBaseCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBaseCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
