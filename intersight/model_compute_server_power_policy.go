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
	"reflect"
	"strings"
)

// ComputeServerPowerPolicy Policy to determine the required power task during server profile deploy/undeploy.
type ComputeServerPowerPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured power state of server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
	PowerState *string `json:"PowerState,omitempty"`
	// The name of the server it is associated with.
	ServerName   *string                               `json:"ServerName,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship      `json:"RegisteredDevice,omitempty"`
	Server               *ComputePhysicalRelationship              `json:"Server,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerPowerPolicy ComputeServerPowerPolicy

// NewComputeServerPowerPolicy instantiates a new ComputeServerPowerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerPowerPolicy(classId string, objectType string) *ComputeServerPowerPolicy {
	this := ComputeServerPowerPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var powerState string = "Policy"
	this.PowerState = &powerState
	return &this
}

// NewComputeServerPowerPolicyWithDefaults instantiates a new ComputeServerPowerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerPowerPolicyWithDefaults() *ComputeServerPowerPolicy {
	this := ComputeServerPowerPolicy{}
	var classId string = "compute.ServerPowerPolicy"
	this.ClassId = classId
	var objectType string = "compute.ServerPowerPolicy"
	this.ObjectType = objectType
	var powerState string = "Policy"
	this.PowerState = &powerState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerPowerPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerPowerPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerPowerPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerPowerPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicy) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *ComputeServerPowerPolicy) SetPowerState(v string) {
	o.PowerState = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicy) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ComputeServerPowerPolicy) SetServerName(v string) {
	o.ServerName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ComputeServerPowerPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerPowerPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerPowerPolicy) GetProfilesOk() (*[]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *ComputeServerPowerPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeServerPowerPolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ComputeServerPowerPolicy) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerPowerPolicy) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ComputeServerPowerPolicy) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *ComputeServerPowerPolicy) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o ComputeServerPowerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ServerName != nil {
		toSerialize["ServerName"] = o.ServerName
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerPowerPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeServerPowerPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// User configured power state of server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
		PowerState *string `json:"PowerState,omitempty"`
		// The name of the server it is associated with.
		ServerName   *string                               `json:"ServerName,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles         []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship      `json:"RegisteredDevice,omitempty"`
		Server           *ComputePhysicalRelationship              `json:"Server,omitempty"`
	}

	varComputeServerPowerPolicyWithoutEmbeddedStruct := ComputeServerPowerPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeServerPowerPolicyWithoutEmbeddedStruct)
	if err == nil {
		varComputeServerPowerPolicy := _ComputeServerPowerPolicy{}
		varComputeServerPowerPolicy.ClassId = varComputeServerPowerPolicyWithoutEmbeddedStruct.ClassId
		varComputeServerPowerPolicy.ObjectType = varComputeServerPowerPolicyWithoutEmbeddedStruct.ObjectType
		varComputeServerPowerPolicy.PowerState = varComputeServerPowerPolicyWithoutEmbeddedStruct.PowerState
		varComputeServerPowerPolicy.ServerName = varComputeServerPowerPolicyWithoutEmbeddedStruct.ServerName
		varComputeServerPowerPolicy.Organization = varComputeServerPowerPolicyWithoutEmbeddedStruct.Organization
		varComputeServerPowerPolicy.Profiles = varComputeServerPowerPolicyWithoutEmbeddedStruct.Profiles
		varComputeServerPowerPolicy.RegisteredDevice = varComputeServerPowerPolicyWithoutEmbeddedStruct.RegisteredDevice
		varComputeServerPowerPolicy.Server = varComputeServerPowerPolicyWithoutEmbeddedStruct.Server
		*o = ComputeServerPowerPolicy(varComputeServerPowerPolicy)
	} else {
		return err
	}

	varComputeServerPowerPolicy := _ComputeServerPowerPolicy{}

	err = json.Unmarshal(bytes, &varComputeServerPowerPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varComputeServerPowerPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "ServerName")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Server")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableComputeServerPowerPolicy struct {
	value *ComputeServerPowerPolicy
	isSet bool
}

func (v NullableComputeServerPowerPolicy) Get() *ComputeServerPowerPolicy {
	return v.value
}

func (v *NullableComputeServerPowerPolicy) Set(val *ComputeServerPowerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerPowerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerPowerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerPowerPolicy(val *ComputeServerPowerPolicy) *NullableComputeServerPowerPolicy {
	return &NullableComputeServerPowerPolicy{value: val, isSet: true}
}

func (v NullableComputeServerPowerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerPowerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
