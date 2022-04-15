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
)

// CloudSkuInstanceTypeAllOf Definition of the list of properties defined in 'cloud.SkuInstanceType', excluding properties defined in parent classes.
type CloudSkuInstanceTypeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates if the instance type supports 32 or 64 bit or both. * `64Bit` - Enum to indicate that the instance type suppports only 64 bit architecture. * `32Bit` - Enum to indicate that the instance type supports only 32 bit architecture. * `both` - Enum to indicate that the instance type supports both 32 and 64 bit architecture.
	ArchitectureType *string `json:"ArchitectureType,omitempty"`
	// The cpu unit for this instance type. * `VIRTUAL_CPU` - The CPU unit used for virtual machines. * `MILLI_CPU` - The CPU unit used by containers.
	CpuUnit *string `json:"CpuUnit,omitempty"`
	// Does the instanceType support CUDA architecture.
	CudaSupport *bool `json:"CudaSupport,omitempty"`
	// Total size of local storage for this instance type.
	LocalStorageSize *float32 `json:"LocalStorageSize,omitempty"`
	// The units for this storage. For e.g. MB or GB or TB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
	LocalStorageSizeUnit *string `json:"LocalStorageSizeUnit,omitempty"`
	// The RAM size for this instance type.
	MemorySize *float32 `json:"MemorySize,omitempty"`
	// Units to represent memory, for e.g. MB or GB. * `MB` - Enum to represent mega byte storage unit. * `GB` - Enum to represent Giga byte storage unit. * `TB` - Enum to represent Tera byte storage unit.
	MemorySizeUnit *string `json:"MemorySizeUnit,omitempty"`
	// Network performance of this instance type.
	NetworkPerformance *string `json:"NetworkPerformance,omitempty"`
	// The number of CPUs in this instance type.
	NumOfCpus *int64 `json:"NumOfCpus,omitempty"`
	// Maximum number of NICs supported by this instance type.
	NumOfNics            *int64 `json:"NumOfNics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudSkuInstanceTypeAllOf CloudSkuInstanceTypeAllOf

// NewCloudSkuInstanceTypeAllOf instantiates a new CloudSkuInstanceTypeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudSkuInstanceTypeAllOf(classId string, objectType string) *CloudSkuInstanceTypeAllOf {
	this := CloudSkuInstanceTypeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var architectureType string = "64Bit"
	this.ArchitectureType = &architectureType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	var localStorageSizeUnit string = "MB"
	this.LocalStorageSizeUnit = &localStorageSizeUnit
	var memorySizeUnit string = "MB"
	this.MemorySizeUnit = &memorySizeUnit
	return &this
}

// NewCloudSkuInstanceTypeAllOfWithDefaults instantiates a new CloudSkuInstanceTypeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudSkuInstanceTypeAllOfWithDefaults() *CloudSkuInstanceTypeAllOf {
	this := CloudSkuInstanceTypeAllOf{}
	var classId string = "cloud.SkuInstanceType"
	this.ClassId = classId
	var objectType string = "cloud.SkuInstanceType"
	this.ObjectType = objectType
	var architectureType string = "64Bit"
	this.ArchitectureType = &architectureType
	var cpuUnit string = "VIRTUAL_CPU"
	this.CpuUnit = &cpuUnit
	var localStorageSizeUnit string = "MB"
	this.LocalStorageSizeUnit = &localStorageSizeUnit
	var memorySizeUnit string = "MB"
	this.MemorySizeUnit = &memorySizeUnit
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudSkuInstanceTypeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudSkuInstanceTypeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudSkuInstanceTypeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudSkuInstanceTypeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArchitectureType returns the ArchitectureType field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetArchitectureType() string {
	if o == nil || o.ArchitectureType == nil {
		var ret string
		return ret
	}
	return *o.ArchitectureType
}

// GetArchitectureTypeOk returns a tuple with the ArchitectureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetArchitectureTypeOk() (*string, bool) {
	if o == nil || o.ArchitectureType == nil {
		return nil, false
	}
	return o.ArchitectureType, true
}

// HasArchitectureType returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasArchitectureType() bool {
	if o != nil && o.ArchitectureType != nil {
		return true
	}

	return false
}

// SetArchitectureType gets a reference to the given string and assigns it to the ArchitectureType field.
func (o *CloudSkuInstanceTypeAllOf) SetArchitectureType(v string) {
	o.ArchitectureType = &v
}

// GetCpuUnit returns the CpuUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetCpuUnit() string {
	if o == nil || o.CpuUnit == nil {
		var ret string
		return ret
	}
	return *o.CpuUnit
}

// GetCpuUnitOk returns a tuple with the CpuUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetCpuUnitOk() (*string, bool) {
	if o == nil || o.CpuUnit == nil {
		return nil, false
	}
	return o.CpuUnit, true
}

// HasCpuUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasCpuUnit() bool {
	if o != nil && o.CpuUnit != nil {
		return true
	}

	return false
}

// SetCpuUnit gets a reference to the given string and assigns it to the CpuUnit field.
func (o *CloudSkuInstanceTypeAllOf) SetCpuUnit(v string) {
	o.CpuUnit = &v
}

// GetCudaSupport returns the CudaSupport field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetCudaSupport() bool {
	if o == nil || o.CudaSupport == nil {
		var ret bool
		return ret
	}
	return *o.CudaSupport
}

// GetCudaSupportOk returns a tuple with the CudaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetCudaSupportOk() (*bool, bool) {
	if o == nil || o.CudaSupport == nil {
		return nil, false
	}
	return o.CudaSupport, true
}

// HasCudaSupport returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasCudaSupport() bool {
	if o != nil && o.CudaSupport != nil {
		return true
	}

	return false
}

// SetCudaSupport gets a reference to the given bool and assigns it to the CudaSupport field.
func (o *CloudSkuInstanceTypeAllOf) SetCudaSupport(v bool) {
	o.CudaSupport = &v
}

// GetLocalStorageSize returns the LocalStorageSize field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetLocalStorageSize() float32 {
	if o == nil || o.LocalStorageSize == nil {
		var ret float32
		return ret
	}
	return *o.LocalStorageSize
}

// GetLocalStorageSizeOk returns a tuple with the LocalStorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetLocalStorageSizeOk() (*float32, bool) {
	if o == nil || o.LocalStorageSize == nil {
		return nil, false
	}
	return o.LocalStorageSize, true
}

// HasLocalStorageSize returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasLocalStorageSize() bool {
	if o != nil && o.LocalStorageSize != nil {
		return true
	}

	return false
}

// SetLocalStorageSize gets a reference to the given float32 and assigns it to the LocalStorageSize field.
func (o *CloudSkuInstanceTypeAllOf) SetLocalStorageSize(v float32) {
	o.LocalStorageSize = &v
}

// GetLocalStorageSizeUnit returns the LocalStorageSizeUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetLocalStorageSizeUnit() string {
	if o == nil || o.LocalStorageSizeUnit == nil {
		var ret string
		return ret
	}
	return *o.LocalStorageSizeUnit
}

// GetLocalStorageSizeUnitOk returns a tuple with the LocalStorageSizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetLocalStorageSizeUnitOk() (*string, bool) {
	if o == nil || o.LocalStorageSizeUnit == nil {
		return nil, false
	}
	return o.LocalStorageSizeUnit, true
}

// HasLocalStorageSizeUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasLocalStorageSizeUnit() bool {
	if o != nil && o.LocalStorageSizeUnit != nil {
		return true
	}

	return false
}

// SetLocalStorageSizeUnit gets a reference to the given string and assigns it to the LocalStorageSizeUnit field.
func (o *CloudSkuInstanceTypeAllOf) SetLocalStorageSizeUnit(v string) {
	o.LocalStorageSizeUnit = &v
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetMemorySize() float32 {
	if o == nil || o.MemorySize == nil {
		var ret float32
		return ret
	}
	return *o.MemorySize
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetMemorySizeOk() (*float32, bool) {
	if o == nil || o.MemorySize == nil {
		return nil, false
	}
	return o.MemorySize, true
}

// HasMemorySize returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasMemorySize() bool {
	if o != nil && o.MemorySize != nil {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given float32 and assigns it to the MemorySize field.
func (o *CloudSkuInstanceTypeAllOf) SetMemorySize(v float32) {
	o.MemorySize = &v
}

// GetMemorySizeUnit returns the MemorySizeUnit field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetMemorySizeUnit() string {
	if o == nil || o.MemorySizeUnit == nil {
		var ret string
		return ret
	}
	return *o.MemorySizeUnit
}

// GetMemorySizeUnitOk returns a tuple with the MemorySizeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetMemorySizeUnitOk() (*string, bool) {
	if o == nil || o.MemorySizeUnit == nil {
		return nil, false
	}
	return o.MemorySizeUnit, true
}

// HasMemorySizeUnit returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasMemorySizeUnit() bool {
	if o != nil && o.MemorySizeUnit != nil {
		return true
	}

	return false
}

// SetMemorySizeUnit gets a reference to the given string and assigns it to the MemorySizeUnit field.
func (o *CloudSkuInstanceTypeAllOf) SetMemorySizeUnit(v string) {
	o.MemorySizeUnit = &v
}

// GetNetworkPerformance returns the NetworkPerformance field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetNetworkPerformance() string {
	if o == nil || o.NetworkPerformance == nil {
		var ret string
		return ret
	}
	return *o.NetworkPerformance
}

// GetNetworkPerformanceOk returns a tuple with the NetworkPerformance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetNetworkPerformanceOk() (*string, bool) {
	if o == nil || o.NetworkPerformance == nil {
		return nil, false
	}
	return o.NetworkPerformance, true
}

// HasNetworkPerformance returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasNetworkPerformance() bool {
	if o != nil && o.NetworkPerformance != nil {
		return true
	}

	return false
}

// SetNetworkPerformance gets a reference to the given string and assigns it to the NetworkPerformance field.
func (o *CloudSkuInstanceTypeAllOf) SetNetworkPerformance(v string) {
	o.NetworkPerformance = &v
}

// GetNumOfCpus returns the NumOfCpus field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetNumOfCpus() int64 {
	if o == nil || o.NumOfCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumOfCpus
}

// GetNumOfCpusOk returns a tuple with the NumOfCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetNumOfCpusOk() (*int64, bool) {
	if o == nil || o.NumOfCpus == nil {
		return nil, false
	}
	return o.NumOfCpus, true
}

// HasNumOfCpus returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasNumOfCpus() bool {
	if o != nil && o.NumOfCpus != nil {
		return true
	}

	return false
}

// SetNumOfCpus gets a reference to the given int64 and assigns it to the NumOfCpus field.
func (o *CloudSkuInstanceTypeAllOf) SetNumOfCpus(v int64) {
	o.NumOfCpus = &v
}

// GetNumOfNics returns the NumOfNics field value if set, zero value otherwise.
func (o *CloudSkuInstanceTypeAllOf) GetNumOfNics() int64 {
	if o == nil || o.NumOfNics == nil {
		var ret int64
		return ret
	}
	return *o.NumOfNics
}

// GetNumOfNicsOk returns a tuple with the NumOfNics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudSkuInstanceTypeAllOf) GetNumOfNicsOk() (*int64, bool) {
	if o == nil || o.NumOfNics == nil {
		return nil, false
	}
	return o.NumOfNics, true
}

// HasNumOfNics returns a boolean if a field has been set.
func (o *CloudSkuInstanceTypeAllOf) HasNumOfNics() bool {
	if o != nil && o.NumOfNics != nil {
		return true
	}

	return false
}

// SetNumOfNics gets a reference to the given int64 and assigns it to the NumOfNics field.
func (o *CloudSkuInstanceTypeAllOf) SetNumOfNics(v int64) {
	o.NumOfNics = &v
}

func (o CloudSkuInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArchitectureType != nil {
		toSerialize["ArchitectureType"] = o.ArchitectureType
	}
	if o.CpuUnit != nil {
		toSerialize["CpuUnit"] = o.CpuUnit
	}
	if o.CudaSupport != nil {
		toSerialize["CudaSupport"] = o.CudaSupport
	}
	if o.LocalStorageSize != nil {
		toSerialize["LocalStorageSize"] = o.LocalStorageSize
	}
	if o.LocalStorageSizeUnit != nil {
		toSerialize["LocalStorageSizeUnit"] = o.LocalStorageSizeUnit
	}
	if o.MemorySize != nil {
		toSerialize["MemorySize"] = o.MemorySize
	}
	if o.MemorySizeUnit != nil {
		toSerialize["MemorySizeUnit"] = o.MemorySizeUnit
	}
	if o.NetworkPerformance != nil {
		toSerialize["NetworkPerformance"] = o.NetworkPerformance
	}
	if o.NumOfCpus != nil {
		toSerialize["NumOfCpus"] = o.NumOfCpus
	}
	if o.NumOfNics != nil {
		toSerialize["NumOfNics"] = o.NumOfNics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudSkuInstanceTypeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudSkuInstanceTypeAllOf := _CloudSkuInstanceTypeAllOf{}

	if err = json.Unmarshal(bytes, &varCloudSkuInstanceTypeAllOf); err == nil {
		*o = CloudSkuInstanceTypeAllOf(varCloudSkuInstanceTypeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArchitectureType")
		delete(additionalProperties, "CpuUnit")
		delete(additionalProperties, "CudaSupport")
		delete(additionalProperties, "LocalStorageSize")
		delete(additionalProperties, "LocalStorageSizeUnit")
		delete(additionalProperties, "MemorySize")
		delete(additionalProperties, "MemorySizeUnit")
		delete(additionalProperties, "NetworkPerformance")
		delete(additionalProperties, "NumOfCpus")
		delete(additionalProperties, "NumOfNics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudSkuInstanceTypeAllOf struct {
	value *CloudSkuInstanceTypeAllOf
	isSet bool
}

func (v NullableCloudSkuInstanceTypeAllOf) Get() *CloudSkuInstanceTypeAllOf {
	return v.value
}

func (v *NullableCloudSkuInstanceTypeAllOf) Set(val *CloudSkuInstanceTypeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudSkuInstanceTypeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudSkuInstanceTypeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudSkuInstanceTypeAllOf(val *CloudSkuInstanceTypeAllOf) *NullableCloudSkuInstanceTypeAllOf {
	return &NullableCloudSkuInstanceTypeAllOf{value: val, isSet: true}
}

func (v NullableCloudSkuInstanceTypeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudSkuInstanceTypeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
