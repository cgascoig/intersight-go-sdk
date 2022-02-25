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

// CloudNetworkAddress Network address details such as IP address, IP version, IP allocation type.
type CloudNetworkAddress struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP (Internet Protocol) address value.
	Ip *string `json:"Ip,omitempty"`
	// IP address allocation type (DYNAMIC | STATIC | IPAM_CALLOUT | PRE_ALLOCATE). * `Dynamic` - IP address allocation type is dynamic. * `Static` - IP address allocation type is static. * `IpamCallout` - IP address is assigned with the results of callout scripts execution. * `PreAllocate` - IP address allocation type is PreAllocate .
	IpAllocation *string `json:"IpAllocation,omitempty"`
	// Whether IP address is of type IPv4 or IPv6. * `IPv4` - Internet protocol version 4. * `IPv6` - Internet protocol version 6.
	IpVersion            *string `json:"IpVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkAddress CloudNetworkAddress

// NewCloudNetworkAddress instantiates a new CloudNetworkAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkAddress(classId string, objectType string) *CloudNetworkAddress {
	this := CloudNetworkAddress{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudNetworkAddressWithDefaults instantiates a new CloudNetworkAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkAddressWithDefaults() *CloudNetworkAddress {
	this := CloudNetworkAddress{}
	var classId string = "cloud.NetworkAddress"
	this.ClassId = classId
	var objectType string = "cloud.NetworkAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudNetworkAddress) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAddress) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudNetworkAddress) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudNetworkAddress) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkAddress) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudNetworkAddress) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CloudNetworkAddress) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAddress) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CloudNetworkAddress) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CloudNetworkAddress) SetIp(v string) {
	o.Ip = &v
}

// GetIpAllocation returns the IpAllocation field value if set, zero value otherwise.
func (o *CloudNetworkAddress) GetIpAllocation() string {
	if o == nil || o.IpAllocation == nil {
		var ret string
		return ret
	}
	return *o.IpAllocation
}

// GetIpAllocationOk returns a tuple with the IpAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAddress) GetIpAllocationOk() (*string, bool) {
	if o == nil || o.IpAllocation == nil {
		return nil, false
	}
	return o.IpAllocation, true
}

// HasIpAllocation returns a boolean if a field has been set.
func (o *CloudNetworkAddress) HasIpAllocation() bool {
	if o != nil && o.IpAllocation != nil {
		return true
	}

	return false
}

// SetIpAllocation gets a reference to the given string and assigns it to the IpAllocation field.
func (o *CloudNetworkAddress) SetIpAllocation(v string) {
	o.IpAllocation = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *CloudNetworkAddress) GetIpVersion() string {
	if o == nil || o.IpVersion == nil {
		var ret string
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkAddress) GetIpVersionOk() (*string, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *CloudNetworkAddress) HasIpVersion() bool {
	if o != nil && o.IpVersion != nil {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given string and assigns it to the IpVersion field.
func (o *CloudNetworkAddress) SetIpVersion(v string) {
	o.IpVersion = &v
}

func (o CloudNetworkAddress) MarshalJSON() ([]byte, error) {
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
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	if o.IpAllocation != nil {
		toSerialize["IpAllocation"] = o.IpAllocation
	}
	if o.IpVersion != nil {
		toSerialize["IpVersion"] = o.IpVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudNetworkAddress) UnmarshalJSON(bytes []byte) (err error) {
	type CloudNetworkAddressWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP (Internet Protocol) address value.
		Ip *string `json:"Ip,omitempty"`
		// IP address allocation type (DYNAMIC | STATIC | IPAM_CALLOUT | PRE_ALLOCATE). * `Dynamic` - IP address allocation type is dynamic. * `Static` - IP address allocation type is static. * `IpamCallout` - IP address is assigned with the results of callout scripts execution. * `PreAllocate` - IP address allocation type is PreAllocate .
		IpAllocation *string `json:"IpAllocation,omitempty"`
		// Whether IP address is of type IPv4 or IPv6. * `IPv4` - Internet protocol version 4. * `IPv6` - Internet protocol version 6.
		IpVersion *string `json:"IpVersion,omitempty"`
	}

	varCloudNetworkAddressWithoutEmbeddedStruct := CloudNetworkAddressWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudNetworkAddressWithoutEmbeddedStruct)
	if err == nil {
		varCloudNetworkAddress := _CloudNetworkAddress{}
		varCloudNetworkAddress.ClassId = varCloudNetworkAddressWithoutEmbeddedStruct.ClassId
		varCloudNetworkAddress.ObjectType = varCloudNetworkAddressWithoutEmbeddedStruct.ObjectType
		varCloudNetworkAddress.Ip = varCloudNetworkAddressWithoutEmbeddedStruct.Ip
		varCloudNetworkAddress.IpAllocation = varCloudNetworkAddressWithoutEmbeddedStruct.IpAllocation
		varCloudNetworkAddress.IpVersion = varCloudNetworkAddressWithoutEmbeddedStruct.IpVersion
		*o = CloudNetworkAddress(varCloudNetworkAddress)
	} else {
		return err
	}

	varCloudNetworkAddress := _CloudNetworkAddress{}

	err = json.Unmarshal(bytes, &varCloudNetworkAddress)
	if err == nil {
		o.MoBaseComplexType = varCloudNetworkAddress.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "IpAllocation")
		delete(additionalProperties, "IpVersion")

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

type NullableCloudNetworkAddress struct {
	value *CloudNetworkAddress
	isSet bool
}

func (v NullableCloudNetworkAddress) Get() *CloudNetworkAddress {
	return v.value
}

func (v *NullableCloudNetworkAddress) Set(val *CloudNetworkAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkAddress(val *CloudNetworkAddress) *NullableCloudNetworkAddress {
	return &NullableCloudNetworkAddress{value: val, isSet: true}
}

func (v NullableCloudNetworkAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
