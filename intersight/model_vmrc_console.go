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

// VmrcConsole API to launch VMRC console to a VMware virtual machine.
type VmrcConsole struct {
	TunnelingTunnel
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                          `json:"ObjectType"`
	Vcenter              *VirtualizationVmwareVcenterRelationship        `json:"Vcenter,omitempty"`
	VirtualMachine       *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VmrcConsole VmrcConsole

// NewVmrcConsole instantiates a new VmrcConsole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVmrcConsole(classId string, objectType string) *VmrcConsole {
	this := VmrcConsole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewVmrcConsoleWithDefaults instantiates a new VmrcConsole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVmrcConsoleWithDefaults() *VmrcConsole {
	this := VmrcConsole{}
	var classId string = "vmrc.Console"
	this.ClassId = classId
	var objectType string = "vmrc.Console"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VmrcConsole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VmrcConsole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VmrcConsole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VmrcConsole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VmrcConsole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VmrcConsole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVcenter returns the Vcenter field value if set, zero value otherwise.
func (o *VmrcConsole) GetVcenter() VirtualizationVmwareVcenterRelationship {
	if o == nil || o.Vcenter == nil {
		var ret VirtualizationVmwareVcenterRelationship
		return ret
	}
	return *o.Vcenter
}

// GetVcenterOk returns a tuple with the Vcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmrcConsole) GetVcenterOk() (*VirtualizationVmwareVcenterRelationship, bool) {
	if o == nil || o.Vcenter == nil {
		return nil, false
	}
	return o.Vcenter, true
}

// HasVcenter returns a boolean if a field has been set.
func (o *VmrcConsole) HasVcenter() bool {
	if o != nil && o.Vcenter != nil {
		return true
	}

	return false
}

// SetVcenter gets a reference to the given VirtualizationVmwareVcenterRelationship and assigns it to the Vcenter field.
func (o *VmrcConsole) SetVcenter(v VirtualizationVmwareVcenterRelationship) {
	o.Vcenter = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VmrcConsole) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVmwareVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VmrcConsole) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VmrcConsole) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVmwareVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VmrcConsole) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VmrcConsole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTunnelingTunnel, errTunnelingTunnel := json.Marshal(o.TunnelingTunnel)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	errTunnelingTunnel = json.Unmarshal([]byte(serializedTunnelingTunnel), &toSerialize)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Vcenter != nil {
		toSerialize["Vcenter"] = o.Vcenter
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VmrcConsole) UnmarshalJSON(bytes []byte) (err error) {
	type VmrcConsoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType     string                                          `json:"ObjectType"`
		Vcenter        *VirtualizationVmwareVcenterRelationship        `json:"Vcenter,omitempty"`
		VirtualMachine *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varVmrcConsoleWithoutEmbeddedStruct := VmrcConsoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVmrcConsoleWithoutEmbeddedStruct)
	if err == nil {
		varVmrcConsole := _VmrcConsole{}
		varVmrcConsole.ClassId = varVmrcConsoleWithoutEmbeddedStruct.ClassId
		varVmrcConsole.ObjectType = varVmrcConsoleWithoutEmbeddedStruct.ObjectType
		varVmrcConsole.Vcenter = varVmrcConsoleWithoutEmbeddedStruct.Vcenter
		varVmrcConsole.VirtualMachine = varVmrcConsoleWithoutEmbeddedStruct.VirtualMachine
		*o = VmrcConsole(varVmrcConsole)
	} else {
		return err
	}

	varVmrcConsole := _VmrcConsole{}

	err = json.Unmarshal(bytes, &varVmrcConsole)
	if err == nil {
		o.TunnelingTunnel = varVmrcConsole.TunnelingTunnel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Vcenter")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectTunnelingTunnel := reflect.ValueOf(o.TunnelingTunnel)
		for i := 0; i < reflectTunnelingTunnel.Type().NumField(); i++ {
			t := reflectTunnelingTunnel.Type().Field(i)

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

type NullableVmrcConsole struct {
	value *VmrcConsole
	isSet bool
}

func (v NullableVmrcConsole) Get() *VmrcConsole {
	return v.value
}

func (v *NullableVmrcConsole) Set(val *VmrcConsole) {
	v.value = val
	v.isSet = true
}

func (v NullableVmrcConsole) IsSet() bool {
	return v.isSet
}

func (v *NullableVmrcConsole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVmrcConsole(val *VmrcConsole) *NullableVmrcConsole {
	return &NullableVmrcConsole{value: val, isSet: true}
}

func (v NullableVmrcConsole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVmrcConsole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
