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

// StorageNetAppBaseIpInterface NetApp IP interface is a logical interface.
type StorageNetAppBaseIpInterface struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the current node of IP interface.
	CurrentNode *string `json:"CurrentNode,omitempty"`
	// Name of the current port of IP interface.
	CurrentPort *string `json:"CurrentPort,omitempty"`
	// IP interface is enabled or not.
	Enabled *string `json:"Enabled,omitempty"`
	// Name of home node of IP interface.
	HomeNode *string `json:"HomeNode,omitempty"`
	// Name of home port of IP interface.
	HomePort *string `json:"HomePort,omitempty"`
	// The IP address of interface.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
	IpFamily *string `json:"IpFamily,omitempty"`
	// The name of the IPspace of the IP interface.
	Ipspace *string `json:"Ipspace,omitempty"`
	// Reports whether the IP interface is home or has failed over to its HA peer.
	IsHome *bool `json:"IsHome,omitempty"`
	// The name of the IP interface.
	Name *string `json:"Name,omitempty"`
	// The netmask of the interface.
	Netmask *string `json:"Netmask,omitempty"`
	// Service policy name of IP interface.
	ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
	// Service policy UUID of IP interface.
	ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty"`
	Services          []string `json:"Services,omitempty"`
	// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
	State *string `json:"State,omitempty"`
	// Uuid of NetApp IP Interface.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppBaseIpInterface StorageNetAppBaseIpInterface

// NewStorageNetAppBaseIpInterface instantiates a new StorageNetAppBaseIpInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppBaseIpInterface(classId string, objectType string) *StorageNetAppBaseIpInterface {
	this := StorageNetAppBaseIpInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppBaseIpInterfaceWithDefaults instantiates a new StorageNetAppBaseIpInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppBaseIpInterfaceWithDefaults() *StorageNetAppBaseIpInterface {
	this := StorageNetAppBaseIpInterface{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppBaseIpInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppBaseIpInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppBaseIpInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppBaseIpInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentNode returns the CurrentNode field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetCurrentNode() string {
	if o == nil || o.CurrentNode == nil {
		var ret string
		return ret
	}
	return *o.CurrentNode
}

// GetCurrentNodeOk returns a tuple with the CurrentNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetCurrentNodeOk() (*string, bool) {
	if o == nil || o.CurrentNode == nil {
		return nil, false
	}
	return o.CurrentNode, true
}

// HasCurrentNode returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasCurrentNode() bool {
	if o != nil && o.CurrentNode != nil {
		return true
	}

	return false
}

// SetCurrentNode gets a reference to the given string and assigns it to the CurrentNode field.
func (o *StorageNetAppBaseIpInterface) SetCurrentNode(v string) {
	o.CurrentNode = &v
}

// GetCurrentPort returns the CurrentPort field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetCurrentPort() string {
	if o == nil || o.CurrentPort == nil {
		var ret string
		return ret
	}
	return *o.CurrentPort
}

// GetCurrentPortOk returns a tuple with the CurrentPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetCurrentPortOk() (*string, bool) {
	if o == nil || o.CurrentPort == nil {
		return nil, false
	}
	return o.CurrentPort, true
}

// HasCurrentPort returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasCurrentPort() bool {
	if o != nil && o.CurrentPort != nil {
		return true
	}

	return false
}

// SetCurrentPort gets a reference to the given string and assigns it to the CurrentPort field.
func (o *StorageNetAppBaseIpInterface) SetCurrentPort(v string) {
	o.CurrentPort = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppBaseIpInterface) SetEnabled(v string) {
	o.Enabled = &v
}

// GetHomeNode returns the HomeNode field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetHomeNode() string {
	if o == nil || o.HomeNode == nil {
		var ret string
		return ret
	}
	return *o.HomeNode
}

// GetHomeNodeOk returns a tuple with the HomeNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetHomeNodeOk() (*string, bool) {
	if o == nil || o.HomeNode == nil {
		return nil, false
	}
	return o.HomeNode, true
}

// HasHomeNode returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasHomeNode() bool {
	if o != nil && o.HomeNode != nil {
		return true
	}

	return false
}

// SetHomeNode gets a reference to the given string and assigns it to the HomeNode field.
func (o *StorageNetAppBaseIpInterface) SetHomeNode(v string) {
	o.HomeNode = &v
}

// GetHomePort returns the HomePort field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetHomePort() string {
	if o == nil || o.HomePort == nil {
		var ret string
		return ret
	}
	return *o.HomePort
}

// GetHomePortOk returns a tuple with the HomePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetHomePortOk() (*string, bool) {
	if o == nil || o.HomePort == nil {
		return nil, false
	}
	return o.HomePort, true
}

// HasHomePort returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasHomePort() bool {
	if o != nil && o.HomePort != nil {
		return true
	}

	return false
}

// SetHomePort gets a reference to the given string and assigns it to the HomePort field.
func (o *StorageNetAppBaseIpInterface) SetHomePort(v string) {
	o.HomePort = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *StorageNetAppBaseIpInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetIpFamily() string {
	if o == nil || o.IpFamily == nil {
		var ret string
		return ret
	}
	return *o.IpFamily
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetIpFamilyOk() (*string, bool) {
	if o == nil || o.IpFamily == nil {
		return nil, false
	}
	return o.IpFamily, true
}

// HasIpFamily returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasIpFamily() bool {
	if o != nil && o.IpFamily != nil {
		return true
	}

	return false
}

// SetIpFamily gets a reference to the given string and assigns it to the IpFamily field.
func (o *StorageNetAppBaseIpInterface) SetIpFamily(v string) {
	o.IpFamily = &v
}

// GetIpspace returns the Ipspace field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetIpspace() string {
	if o == nil || o.Ipspace == nil {
		var ret string
		return ret
	}
	return *o.Ipspace
}

// GetIpspaceOk returns a tuple with the Ipspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetIpspaceOk() (*string, bool) {
	if o == nil || o.Ipspace == nil {
		return nil, false
	}
	return o.Ipspace, true
}

// HasIpspace returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasIpspace() bool {
	if o != nil && o.Ipspace != nil {
		return true
	}

	return false
}

// SetIpspace gets a reference to the given string and assigns it to the Ipspace field.
func (o *StorageNetAppBaseIpInterface) SetIpspace(v string) {
	o.Ipspace = &v
}

// GetIsHome returns the IsHome field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetIsHome() bool {
	if o == nil || o.IsHome == nil {
		var ret bool
		return ret
	}
	return *o.IsHome
}

// GetIsHomeOk returns a tuple with the IsHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetIsHomeOk() (*bool, bool) {
	if o == nil || o.IsHome == nil {
		return nil, false
	}
	return o.IsHome, true
}

// HasIsHome returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasIsHome() bool {
	if o != nil && o.IsHome != nil {
		return true
	}

	return false
}

// SetIsHome gets a reference to the given bool and assigns it to the IsHome field.
func (o *StorageNetAppBaseIpInterface) SetIsHome(v bool) {
	o.IsHome = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppBaseIpInterface) SetName(v string) {
	o.Name = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *StorageNetAppBaseIpInterface) SetNetmask(v string) {
	o.Netmask = &v
}

// GetServicePolicyName returns the ServicePolicyName field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetServicePolicyName() string {
	if o == nil || o.ServicePolicyName == nil {
		var ret string
		return ret
	}
	return *o.ServicePolicyName
}

// GetServicePolicyNameOk returns a tuple with the ServicePolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetServicePolicyNameOk() (*string, bool) {
	if o == nil || o.ServicePolicyName == nil {
		return nil, false
	}
	return o.ServicePolicyName, true
}

// HasServicePolicyName returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasServicePolicyName() bool {
	if o != nil && o.ServicePolicyName != nil {
		return true
	}

	return false
}

// SetServicePolicyName gets a reference to the given string and assigns it to the ServicePolicyName field.
func (o *StorageNetAppBaseIpInterface) SetServicePolicyName(v string) {
	o.ServicePolicyName = &v
}

// GetServicePolicyUuid returns the ServicePolicyUuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetServicePolicyUuid() string {
	if o == nil || o.ServicePolicyUuid == nil {
		var ret string
		return ret
	}
	return *o.ServicePolicyUuid
}

// GetServicePolicyUuidOk returns a tuple with the ServicePolicyUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetServicePolicyUuidOk() (*string, bool) {
	if o == nil || o.ServicePolicyUuid == nil {
		return nil, false
	}
	return o.ServicePolicyUuid, true
}

// HasServicePolicyUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasServicePolicyUuid() bool {
	if o != nil && o.ServicePolicyUuid != nil {
		return true
	}

	return false
}

// SetServicePolicyUuid gets a reference to the given string and assigns it to the ServicePolicyUuid field.
func (o *StorageNetAppBaseIpInterface) SetServicePolicyUuid(v string) {
	o.ServicePolicyUuid = &v
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppBaseIpInterface) GetServices() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppBaseIpInterface) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *StorageNetAppBaseIpInterface) SetServices(v []string) {
	o.Services = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppBaseIpInterface) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppBaseIpInterface) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppBaseIpInterface) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppBaseIpInterface) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppBaseIpInterface) SetUuid(v string) {
	o.Uuid = &v
}

func (o StorageNetAppBaseIpInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentNode != nil {
		toSerialize["CurrentNode"] = o.CurrentNode
	}
	if o.CurrentPort != nil {
		toSerialize["CurrentPort"] = o.CurrentPort
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.HomeNode != nil {
		toSerialize["HomeNode"] = o.HomeNode
	}
	if o.HomePort != nil {
		toSerialize["HomePort"] = o.HomePort
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.IpFamily != nil {
		toSerialize["IpFamily"] = o.IpFamily
	}
	if o.Ipspace != nil {
		toSerialize["Ipspace"] = o.Ipspace
	}
	if o.IsHome != nil {
		toSerialize["IsHome"] = o.IsHome
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}
	if o.ServicePolicyName != nil {
		toSerialize["ServicePolicyName"] = o.ServicePolicyName
	}
	if o.ServicePolicyUuid != nil {
		toSerialize["ServicePolicyUuid"] = o.ServicePolicyUuid
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppBaseIpInterface) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppBaseIpInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the current node of IP interface.
		CurrentNode *string `json:"CurrentNode,omitempty"`
		// Name of the current port of IP interface.
		CurrentPort *string `json:"CurrentPort,omitempty"`
		// IP interface is enabled or not.
		Enabled *string `json:"Enabled,omitempty"`
		// Name of home node of IP interface.
		HomeNode *string `json:"HomeNode,omitempty"`
		// Name of home port of IP interface.
		HomePort *string `json:"HomePort,omitempty"`
		// The IP address of interface.
		IpAddress *string `json:"IpAddress,omitempty"`
		// IP address family of interface. * `IPv4` - IP address family type is IPv4. * `IPv6` - IP address family type is IP6.
		IpFamily *string `json:"IpFamily,omitempty"`
		// The name of the IPspace of the IP interface.
		Ipspace *string `json:"Ipspace,omitempty"`
		// Reports whether the IP interface is home or has failed over to its HA peer.
		IsHome *bool `json:"IsHome,omitempty"`
		// The name of the IP interface.
		Name *string `json:"Name,omitempty"`
		// The netmask of the interface.
		Netmask *string `json:"Netmask,omitempty"`
		// Service policy name of IP interface.
		ServicePolicyName *string `json:"ServicePolicyName,omitempty"`
		// Service policy UUID of IP interface.
		ServicePolicyUuid *string  `json:"ServicePolicyUuid,omitempty"`
		Services          []string `json:"Services,omitempty"`
		// The state of the IP interface. * `down` - An inactive port is listed as Down. * `up` - An active port is listed as Up. * `present` - An active port is listed as present.
		State *string `json:"State,omitempty"`
		// Uuid of NetApp IP Interface.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct := StorageNetAppBaseIpInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppBaseIpInterface := _StorageNetAppBaseIpInterface{}
		varStorageNetAppBaseIpInterface.ClassId = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppBaseIpInterface.ObjectType = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppBaseIpInterface.CurrentNode = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.CurrentNode
		varStorageNetAppBaseIpInterface.CurrentPort = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.CurrentPort
		varStorageNetAppBaseIpInterface.Enabled = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Enabled
		varStorageNetAppBaseIpInterface.HomeNode = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.HomeNode
		varStorageNetAppBaseIpInterface.HomePort = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.HomePort
		varStorageNetAppBaseIpInterface.IpAddress = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.IpAddress
		varStorageNetAppBaseIpInterface.IpFamily = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.IpFamily
		varStorageNetAppBaseIpInterface.Ipspace = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Ipspace
		varStorageNetAppBaseIpInterface.IsHome = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.IsHome
		varStorageNetAppBaseIpInterface.Name = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Name
		varStorageNetAppBaseIpInterface.Netmask = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Netmask
		varStorageNetAppBaseIpInterface.ServicePolicyName = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.ServicePolicyName
		varStorageNetAppBaseIpInterface.ServicePolicyUuid = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.ServicePolicyUuid
		varStorageNetAppBaseIpInterface.Services = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Services
		varStorageNetAppBaseIpInterface.State = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.State
		varStorageNetAppBaseIpInterface.Uuid = varStorageNetAppBaseIpInterfaceWithoutEmbeddedStruct.Uuid
		*o = StorageNetAppBaseIpInterface(varStorageNetAppBaseIpInterface)
	} else {
		return err
	}

	varStorageNetAppBaseIpInterface := _StorageNetAppBaseIpInterface{}

	err = json.Unmarshal(bytes, &varStorageNetAppBaseIpInterface)
	if err == nil {
		o.MoBaseMo = varStorageNetAppBaseIpInterface.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentNode")
		delete(additionalProperties, "CurrentPort")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "HomeNode")
		delete(additionalProperties, "HomePort")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "IpFamily")
		delete(additionalProperties, "Ipspace")
		delete(additionalProperties, "IsHome")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Netmask")
		delete(additionalProperties, "ServicePolicyName")
		delete(additionalProperties, "ServicePolicyUuid")
		delete(additionalProperties, "Services")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Uuid")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableStorageNetAppBaseIpInterface struct {
	value *StorageNetAppBaseIpInterface
	isSet bool
}

func (v NullableStorageNetAppBaseIpInterface) Get() *StorageNetAppBaseIpInterface {
	return v.value
}

func (v *NullableStorageNetAppBaseIpInterface) Set(val *StorageNetAppBaseIpInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppBaseIpInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppBaseIpInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppBaseIpInterface(val *StorageNetAppBaseIpInterface) *NullableStorageNetAppBaseIpInterface {
	return &NullableStorageNetAppBaseIpInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppBaseIpInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppBaseIpInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
