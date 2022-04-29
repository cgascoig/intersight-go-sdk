/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// VirtualizationCloudVmConfigurationAllOf Definition of the list of properties defined in 'virtualization.CloudVmConfiguration', excluding properties defined in parent classes.
type VirtualizationCloudVmConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Availability zone for the VM.
	AvailabilityZoneId *string                                           `json:"AvailabilityZoneId,omitempty"`
	Compute            NullableVirtualizationCloudVmComputeConfiguration `json:"Compute,omitempty"`
	// Virtual machine image used by this VM.
	ImageId *string `json:"ImageId,omitempty"`
	// Keypair for accessing the VM.
	KeyPairName *string                                           `json:"KeyPairName,omitempty"`
	Network     NullableVirtualizationCloudVmNetworkConfiguration `json:"Network,omitempty"`
	// Region where the VM instance is created.
	RegionId       *string                                           `json:"RegionId,omitempty"`
	SecurityGroups []string                                          `json:"SecurityGroups,omitempty"`
	Storage        NullableVirtualizationCloudVmStorageConfiguration `json:"Storage,omitempty"`
	// Unique Identifier of the cloud VM.
	VmId                 *string `json:"VmId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationCloudVmConfigurationAllOf VirtualizationCloudVmConfigurationAllOf

// NewVirtualizationCloudVmConfigurationAllOf instantiates a new VirtualizationCloudVmConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationCloudVmConfigurationAllOf(classId string, objectType string) *VirtualizationCloudVmConfigurationAllOf {
	this := VirtualizationCloudVmConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationCloudVmConfigurationAllOfWithDefaults instantiates a new VirtualizationCloudVmConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationCloudVmConfigurationAllOfWithDefaults() *VirtualizationCloudVmConfigurationAllOf {
	this := VirtualizationCloudVmConfigurationAllOf{}
	var classId string = "virtualization.AwsVmConfiguration"
	this.ClassId = classId
	var objectType string = "virtualization.AwsVmConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationCloudVmConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationCloudVmConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationCloudVmConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationCloudVmConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvailabilityZoneId returns the AvailabilityZoneId field value if set, zero value otherwise.
func (o *VirtualizationCloudVmConfigurationAllOf) GetAvailabilityZoneId() string {
	if o == nil || o.AvailabilityZoneId == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneId
}

// GetAvailabilityZoneIdOk returns a tuple with the AvailabilityZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetAvailabilityZoneIdOk() (*string, bool) {
	if o == nil || o.AvailabilityZoneId == nil {
		return nil, false
	}
	return o.AvailabilityZoneId, true
}

// HasAvailabilityZoneId returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasAvailabilityZoneId() bool {
	if o != nil && o.AvailabilityZoneId != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneId gets a reference to the given string and assigns it to the AvailabilityZoneId field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetAvailabilityZoneId(v string) {
	o.AvailabilityZoneId = &v
}

// GetCompute returns the Compute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationCloudVmConfigurationAllOf) GetCompute() VirtualizationCloudVmComputeConfiguration {
	if o == nil || o.Compute.Get() == nil {
		var ret VirtualizationCloudVmComputeConfiguration
		return ret
	}
	return *o.Compute.Get()
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationCloudVmConfigurationAllOf) GetComputeOk() (*VirtualizationCloudVmComputeConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Compute.Get(), o.Compute.IsSet()
}

// HasCompute returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasCompute() bool {
	if o != nil && o.Compute.IsSet() {
		return true
	}

	return false
}

// SetCompute gets a reference to the given NullableVirtualizationCloudVmComputeConfiguration and assigns it to the Compute field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetCompute(v VirtualizationCloudVmComputeConfiguration) {
	o.Compute.Set(&v)
}

// SetComputeNil sets the value for Compute to be an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) SetComputeNil() {
	o.Compute.Set(nil)
}

// UnsetCompute ensures that no value is present for Compute, not even an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) UnsetCompute() {
	o.Compute.Unset()
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *VirtualizationCloudVmConfigurationAllOf) GetImageId() string {
	if o == nil || o.ImageId == nil {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetImageIdOk() (*string, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetImageId(v string) {
	o.ImageId = &v
}

// GetKeyPairName returns the KeyPairName field value if set, zero value otherwise.
func (o *VirtualizationCloudVmConfigurationAllOf) GetKeyPairName() string {
	if o == nil || o.KeyPairName == nil {
		var ret string
		return ret
	}
	return *o.KeyPairName
}

// GetKeyPairNameOk returns a tuple with the KeyPairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetKeyPairNameOk() (*string, bool) {
	if o == nil || o.KeyPairName == nil {
		return nil, false
	}
	return o.KeyPairName, true
}

// HasKeyPairName returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasKeyPairName() bool {
	if o != nil && o.KeyPairName != nil {
		return true
	}

	return false
}

// SetKeyPairName gets a reference to the given string and assigns it to the KeyPairName field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetKeyPairName(v string) {
	o.KeyPairName = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationCloudVmConfigurationAllOf) GetNetwork() VirtualizationCloudVmNetworkConfiguration {
	if o == nil || o.Network.Get() == nil {
		var ret VirtualizationCloudVmNetworkConfiguration
		return ret
	}
	return *o.Network.Get()
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationCloudVmConfigurationAllOf) GetNetworkOk() (*VirtualizationCloudVmNetworkConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Network.Get(), o.Network.IsSet()
}

// HasNetwork returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasNetwork() bool {
	if o != nil && o.Network.IsSet() {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given NullableVirtualizationCloudVmNetworkConfiguration and assigns it to the Network field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetNetwork(v VirtualizationCloudVmNetworkConfiguration) {
	o.Network.Set(&v)
}

// SetNetworkNil sets the value for Network to be an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) SetNetworkNil() {
	o.Network.Set(nil)
}

// UnsetNetwork ensures that no value is present for Network, not even an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) UnsetNetwork() {
	o.Network.Unset()
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *VirtualizationCloudVmConfigurationAllOf) GetRegionId() string {
	if o == nil || o.RegionId == nil {
		var ret string
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetRegionIdOk() (*string, bool) {
	if o == nil || o.RegionId == nil {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasRegionId() bool {
	if o != nil && o.RegionId != nil {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given string and assigns it to the RegionId field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetRegionId(v string) {
	o.RegionId = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationCloudVmConfigurationAllOf) GetSecurityGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationCloudVmConfigurationAllOf) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return &o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetSecurityGroups(v []string) {
	o.SecurityGroups = v
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationCloudVmConfigurationAllOf) GetStorage() VirtualizationCloudVmStorageConfiguration {
	if o == nil || o.Storage.Get() == nil {
		var ret VirtualizationCloudVmStorageConfiguration
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationCloudVmConfigurationAllOf) GetStorageOk() (*VirtualizationCloudVmStorageConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableVirtualizationCloudVmStorageConfiguration and assigns it to the Storage field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetStorage(v VirtualizationCloudVmStorageConfiguration) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *VirtualizationCloudVmConfigurationAllOf) UnsetStorage() {
	o.Storage.Unset()
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *VirtualizationCloudVmConfigurationAllOf) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *VirtualizationCloudVmConfigurationAllOf) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *VirtualizationCloudVmConfigurationAllOf) SetVmId(v string) {
	o.VmId = &v
}

func (o VirtualizationCloudVmConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvailabilityZoneId != nil {
		toSerialize["AvailabilityZoneId"] = o.AvailabilityZoneId
	}
	if o.Compute.IsSet() {
		toSerialize["Compute"] = o.Compute.Get()
	}
	if o.ImageId != nil {
		toSerialize["ImageId"] = o.ImageId
	}
	if o.KeyPairName != nil {
		toSerialize["KeyPairName"] = o.KeyPairName
	}
	if o.Network.IsSet() {
		toSerialize["Network"] = o.Network.Get()
	}
	if o.RegionId != nil {
		toSerialize["RegionId"] = o.RegionId
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.Storage.IsSet() {
		toSerialize["Storage"] = o.Storage.Get()
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationCloudVmConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationCloudVmConfigurationAllOf := _VirtualizationCloudVmConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationCloudVmConfigurationAllOf); err == nil {
		*o = VirtualizationCloudVmConfigurationAllOf(varVirtualizationCloudVmConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailabilityZoneId")
		delete(additionalProperties, "Compute")
		delete(additionalProperties, "ImageId")
		delete(additionalProperties, "KeyPairName")
		delete(additionalProperties, "Network")
		delete(additionalProperties, "RegionId")
		delete(additionalProperties, "SecurityGroups")
		delete(additionalProperties, "Storage")
		delete(additionalProperties, "VmId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationCloudVmConfigurationAllOf struct {
	value *VirtualizationCloudVmConfigurationAllOf
	isSet bool
}

func (v NullableVirtualizationCloudVmConfigurationAllOf) Get() *VirtualizationCloudVmConfigurationAllOf {
	return v.value
}

func (v *NullableVirtualizationCloudVmConfigurationAllOf) Set(val *VirtualizationCloudVmConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationCloudVmConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationCloudVmConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationCloudVmConfigurationAllOf(val *VirtualizationCloudVmConfigurationAllOf) *NullableVirtualizationCloudVmConfigurationAllOf {
	return &NullableVirtualizationCloudVmConfigurationAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationCloudVmConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationCloudVmConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
