/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationIpAddressInfo Ipaddress details of the VM (host) launched.
type VirtualizationIpAddressInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of the device on network which forwards local traffic to other networks.
	GatewayIp *string `json:"GatewayIp,omitempty"`
	// An IP address is a 32-bit number. It uniquely identifies a host in given network.
	IpAddress *string `json:"IpAddress,omitempty"`
	// A 32 bit number which helps to identify the host and rest of the network.
	SubnetMask           *string `json:"SubnetMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIpAddressInfo VirtualizationIpAddressInfo

// NewVirtualizationIpAddressInfo instantiates a new VirtualizationIpAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIpAddressInfo(classId string, objectType string) *VirtualizationIpAddressInfo {
	this := VirtualizationIpAddressInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIpAddressInfoWithDefaults instantiates a new VirtualizationIpAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIpAddressInfoWithDefaults() *VirtualizationIpAddressInfo {
	this := VirtualizationIpAddressInfo{}
	var classId string = "virtualization.IpAddressInfo"
	this.ClassId = classId
	var objectType string = "virtualization.IpAddressInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIpAddressInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIpAddressInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIpAddressInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIpAddressInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIpAddressInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIpAddressInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *VirtualizationIpAddressInfo) GetGatewayIp() string {
	if o == nil || o.GatewayIp == nil {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIpAddressInfo) GetGatewayIpOk() (*string, bool) {
	if o == nil || o.GatewayIp == nil {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *VirtualizationIpAddressInfo) HasGatewayIp() bool {
	if o != nil && o.GatewayIp != nil {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *VirtualizationIpAddressInfo) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VirtualizationIpAddressInfo) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIpAddressInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationIpAddressInfo) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *VirtualizationIpAddressInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *VirtualizationIpAddressInfo) GetSubnetMask() string {
	if o == nil || o.SubnetMask == nil {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIpAddressInfo) GetSubnetMaskOk() (*string, bool) {
	if o == nil || o.SubnetMask == nil {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *VirtualizationIpAddressInfo) HasSubnetMask() bool {
	if o != nil && o.SubnetMask != nil {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *VirtualizationIpAddressInfo) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

func (o VirtualizationIpAddressInfo) MarshalJSON() ([]byte, error) {
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
	if o.GatewayIp != nil {
		toSerialize["GatewayIp"] = o.GatewayIp
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.SubnetMask != nil {
		toSerialize["SubnetMask"] = o.SubnetMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIpAddressInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIpAddressInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of the device on network which forwards local traffic to other networks.
		GatewayIp *string `json:"GatewayIp,omitempty"`
		// An IP address is a 32-bit number. It uniquely identifies a host in given network.
		IpAddress *string `json:"IpAddress,omitempty"`
		// A 32 bit number which helps to identify the host and rest of the network.
		SubnetMask *string `json:"SubnetMask,omitempty"`
	}

	varVirtualizationIpAddressInfoWithoutEmbeddedStruct := VirtualizationIpAddressInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIpAddressInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIpAddressInfo := _VirtualizationIpAddressInfo{}
		varVirtualizationIpAddressInfo.ClassId = varVirtualizationIpAddressInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationIpAddressInfo.ObjectType = varVirtualizationIpAddressInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationIpAddressInfo.GatewayIp = varVirtualizationIpAddressInfoWithoutEmbeddedStruct.GatewayIp
		varVirtualizationIpAddressInfo.IpAddress = varVirtualizationIpAddressInfoWithoutEmbeddedStruct.IpAddress
		varVirtualizationIpAddressInfo.SubnetMask = varVirtualizationIpAddressInfoWithoutEmbeddedStruct.SubnetMask
		*o = VirtualizationIpAddressInfo(varVirtualizationIpAddressInfo)
	} else {
		return err
	}

	varVirtualizationIpAddressInfo := _VirtualizationIpAddressInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationIpAddressInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationIpAddressInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GatewayIp")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "SubnetMask")

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

type NullableVirtualizationIpAddressInfo struct {
	value *VirtualizationIpAddressInfo
	isSet bool
}

func (v NullableVirtualizationIpAddressInfo) Get() *VirtualizationIpAddressInfo {
	return v.value
}

func (v *NullableVirtualizationIpAddressInfo) Set(val *VirtualizationIpAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIpAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIpAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIpAddressInfo(val *VirtualizationIpAddressInfo) *NullableVirtualizationIpAddressInfo {
	return &NullableVirtualizationIpAddressInfo{value: val, isSet: true}
}

func (v NullableVirtualizationIpAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIpAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
