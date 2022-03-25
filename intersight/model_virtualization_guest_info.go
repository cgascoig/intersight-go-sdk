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

// VirtualizationGuestInfo Captures the common details of the guest personality that runs in a VM.
type VirtualizationGuestInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).
	Hostname *string `json:"Hostname,omitempty"`
	// Primary IP address of the guest os.
	IpAddress *string `json:"IpAddress,omitempty"`
	// The name of the guest running on this VM. This may not be the same as the hostname.
	Name *string `json:"Name,omitempty"`
	// The name of the guest OS running on this VM (Cent OS 4/5/6/7).
	OperatingSystem      *string `json:"OperatingSystem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationGuestInfo VirtualizationGuestInfo

// NewVirtualizationGuestInfo instantiates a new VirtualizationGuestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationGuestInfo(classId string, objectType string) *VirtualizationGuestInfo {
	this := VirtualizationGuestInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationGuestInfoWithDefaults instantiates a new VirtualizationGuestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationGuestInfoWithDefaults() *VirtualizationGuestInfo {
	this := VirtualizationGuestInfo{}
	var classId string = "virtualization.GuestInfo"
	this.ClassId = classId
	var objectType string = "virtualization.GuestInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationGuestInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationGuestInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationGuestInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationGuestInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *VirtualizationGuestInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *VirtualizationGuestInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *VirtualizationGuestInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VirtualizationGuestInfo) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationGuestInfo) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *VirtualizationGuestInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationGuestInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationGuestInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationGuestInfo) SetName(v string) {
	o.Name = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *VirtualizationGuestInfo) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationGuestInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *VirtualizationGuestInfo) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *VirtualizationGuestInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

func (o VirtualizationGuestInfo) MarshalJSON() ([]byte, error) {
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
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperatingSystem != nil {
		toSerialize["OperatingSystem"] = o.OperatingSystem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationGuestInfo) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationGuestInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name provided to the host OS (example, ubuntu6410, test-gateway, etc.).
		Hostname *string `json:"Hostname,omitempty"`
		// Primary IP address of the guest os.
		IpAddress *string `json:"IpAddress,omitempty"`
		// The name of the guest running on this VM. This may not be the same as the hostname.
		Name *string `json:"Name,omitempty"`
		// The name of the guest OS running on this VM (Cent OS 4/5/6/7).
		OperatingSystem *string `json:"OperatingSystem,omitempty"`
	}

	varVirtualizationGuestInfoWithoutEmbeddedStruct := VirtualizationGuestInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationGuestInfoWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationGuestInfo := _VirtualizationGuestInfo{}
		varVirtualizationGuestInfo.ClassId = varVirtualizationGuestInfoWithoutEmbeddedStruct.ClassId
		varVirtualizationGuestInfo.ObjectType = varVirtualizationGuestInfoWithoutEmbeddedStruct.ObjectType
		varVirtualizationGuestInfo.Hostname = varVirtualizationGuestInfoWithoutEmbeddedStruct.Hostname
		varVirtualizationGuestInfo.IpAddress = varVirtualizationGuestInfoWithoutEmbeddedStruct.IpAddress
		varVirtualizationGuestInfo.Name = varVirtualizationGuestInfoWithoutEmbeddedStruct.Name
		varVirtualizationGuestInfo.OperatingSystem = varVirtualizationGuestInfoWithoutEmbeddedStruct.OperatingSystem
		*o = VirtualizationGuestInfo(varVirtualizationGuestInfo)
	} else {
		return err
	}

	varVirtualizationGuestInfo := _VirtualizationGuestInfo{}

	err = json.Unmarshal(bytes, &varVirtualizationGuestInfo)
	if err == nil {
		o.MoBaseComplexType = varVirtualizationGuestInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperatingSystem")

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

type NullableVirtualizationGuestInfo struct {
	value *VirtualizationGuestInfo
	isSet bool
}

func (v NullableVirtualizationGuestInfo) Get() *VirtualizationGuestInfo {
	return v.value
}

func (v *NullableVirtualizationGuestInfo) Set(val *VirtualizationGuestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationGuestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationGuestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationGuestInfo(val *VirtualizationGuestInfo) *NullableVirtualizationGuestInfo {
	return &NullableVirtualizationGuestInfo{value: val, isSet: true}
}

func (v NullableVirtualizationGuestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationGuestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
