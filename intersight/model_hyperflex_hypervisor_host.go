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

// HyperflexHypervisorHost A host appliance in the HyperFlex Cluster.
type HyperflexHypervisorHost struct {
	VirtualizationBaseHost
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Memory configured for controller virtual machine.
	ConfiguredMemory *int64 `json:"ConfiguredMemory,omitempty"`
	// The display version of HyperFlex software running on the controller VM.
	ControllerVmDisplayVersion *string `json:"ControllerVmDisplayVersion,omitempty"`
	// Platform storage software product display version running on controller VM.
	ControllerVmHxdpDisplayVersion *string `json:"ControllerVmHxdpDisplayVersion,omitempty"`
	// Platform storage software product version running on controller VM.
	ControllerVmHxdpVersion *string `json:"ControllerVmHxdpVersion,omitempty"`
	// The UUID of the controller VM which belongs to this host.
	ControllerVmUuid *string `json:"ControllerVmUuid,omitempty"`
	// The version of HyperFlex software running on the controller VM.
	ControllerVmVersion *string                                `json:"ControllerVmVersion,omitempty"`
	DataIp              NullableNetworkHyperFlexNetworkAddress `json:"DataIp,omitempty"`
	// The status of the HyperFlex host. * `UNKNOWN` - Current status of the HyperFlex host is unknown. * `ONLINE` - The HyperFlex host is online. * `OFFLINE` - The HyperFlex host is offline. * `INMAINTENANCE` - The HyperFlex host is in maintenance mode. * `DEGRADED` - Current status of the HyperFlex virtual machine is degraded.
	HostStatus *string `json:"HostStatus,omitempty"`
	// The hypervisor type of the host.
	Hypervisor *string                                `json:"Hypervisor,omitempty"`
	Ip         NullableNetworkHyperFlexNetworkAddress `json:"Ip,omitempty"`
	// Flag indicating whether the HyperFlex host is in lockdown mode.
	Lockdown *bool                                  `json:"Lockdown,omitempty"`
	MgmtIp   NullableNetworkHyperFlexNetworkAddress `json:"MgmtIp,omitempty"`
	// The operation system version of the controller VM.
	OsVersion *string `json:"OsVersion,omitempty"`
	// The role of the HyperFlex host. * `UNKNOWN` - The role of the HyperFlex host is unknown. * `STORAGE` - The HyperFlex host's role is storage. * `COMPUTE` - The HyperFlex host's role is compute.
	Role *string `json:"Role,omitempty"`
	// The controller virtual machine template version.
	TemplateVersion *string `json:"TemplateVersion,omitempty"`
	// Configured number of virtual CPUs for Controller virtual machine.
	VirtualCpus          *int32                        `json:"VirtualCpus,omitempty"`
	Cluster              *HyperflexClusterRelationship `json:"Cluster,omitempty"`
	Node                 *HyperflexNodeRelationship    `json:"Node,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHypervisorHost HyperflexHypervisorHost

// NewHyperflexHypervisorHost instantiates a new HyperflexHypervisorHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHypervisorHost(classId string, objectType string) *HyperflexHypervisorHost {
	this := HyperflexHypervisorHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewHyperflexHypervisorHostWithDefaults instantiates a new HyperflexHypervisorHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHypervisorHostWithDefaults() *HyperflexHypervisorHost {
	this := HyperflexHypervisorHost{}
	var classId string = "hyperflex.HypervisorHost"
	this.ClassId = classId
	var objectType string = "hyperflex.HypervisorHost"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHypervisorHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHypervisorHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHypervisorHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHypervisorHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfiguredMemory returns the ConfiguredMemory field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetConfiguredMemory() int64 {
	if o == nil || o.ConfiguredMemory == nil {
		var ret int64
		return ret
	}
	return *o.ConfiguredMemory
}

// GetConfiguredMemoryOk returns a tuple with the ConfiguredMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetConfiguredMemoryOk() (*int64, bool) {
	if o == nil || o.ConfiguredMemory == nil {
		return nil, false
	}
	return o.ConfiguredMemory, true
}

// HasConfiguredMemory returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasConfiguredMemory() bool {
	if o != nil && o.ConfiguredMemory != nil {
		return true
	}

	return false
}

// SetConfiguredMemory gets a reference to the given int64 and assigns it to the ConfiguredMemory field.
func (o *HyperflexHypervisorHost) SetConfiguredMemory(v int64) {
	o.ConfiguredMemory = &v
}

// GetControllerVmDisplayVersion returns the ControllerVmDisplayVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetControllerVmDisplayVersion() string {
	if o == nil || o.ControllerVmDisplayVersion == nil {
		var ret string
		return ret
	}
	return *o.ControllerVmDisplayVersion
}

// GetControllerVmDisplayVersionOk returns a tuple with the ControllerVmDisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetControllerVmDisplayVersionOk() (*string, bool) {
	if o == nil || o.ControllerVmDisplayVersion == nil {
		return nil, false
	}
	return o.ControllerVmDisplayVersion, true
}

// HasControllerVmDisplayVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasControllerVmDisplayVersion() bool {
	if o != nil && o.ControllerVmDisplayVersion != nil {
		return true
	}

	return false
}

// SetControllerVmDisplayVersion gets a reference to the given string and assigns it to the ControllerVmDisplayVersion field.
func (o *HyperflexHypervisorHost) SetControllerVmDisplayVersion(v string) {
	o.ControllerVmDisplayVersion = &v
}

// GetControllerVmHxdpDisplayVersion returns the ControllerVmHxdpDisplayVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetControllerVmHxdpDisplayVersion() string {
	if o == nil || o.ControllerVmHxdpDisplayVersion == nil {
		var ret string
		return ret
	}
	return *o.ControllerVmHxdpDisplayVersion
}

// GetControllerVmHxdpDisplayVersionOk returns a tuple with the ControllerVmHxdpDisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetControllerVmHxdpDisplayVersionOk() (*string, bool) {
	if o == nil || o.ControllerVmHxdpDisplayVersion == nil {
		return nil, false
	}
	return o.ControllerVmHxdpDisplayVersion, true
}

// HasControllerVmHxdpDisplayVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasControllerVmHxdpDisplayVersion() bool {
	if o != nil && o.ControllerVmHxdpDisplayVersion != nil {
		return true
	}

	return false
}

// SetControllerVmHxdpDisplayVersion gets a reference to the given string and assigns it to the ControllerVmHxdpDisplayVersion field.
func (o *HyperflexHypervisorHost) SetControllerVmHxdpDisplayVersion(v string) {
	o.ControllerVmHxdpDisplayVersion = &v
}

// GetControllerVmHxdpVersion returns the ControllerVmHxdpVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetControllerVmHxdpVersion() string {
	if o == nil || o.ControllerVmHxdpVersion == nil {
		var ret string
		return ret
	}
	return *o.ControllerVmHxdpVersion
}

// GetControllerVmHxdpVersionOk returns a tuple with the ControllerVmHxdpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetControllerVmHxdpVersionOk() (*string, bool) {
	if o == nil || o.ControllerVmHxdpVersion == nil {
		return nil, false
	}
	return o.ControllerVmHxdpVersion, true
}

// HasControllerVmHxdpVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasControllerVmHxdpVersion() bool {
	if o != nil && o.ControllerVmHxdpVersion != nil {
		return true
	}

	return false
}

// SetControllerVmHxdpVersion gets a reference to the given string and assigns it to the ControllerVmHxdpVersion field.
func (o *HyperflexHypervisorHost) SetControllerVmHxdpVersion(v string) {
	o.ControllerVmHxdpVersion = &v
}

// GetControllerVmUuid returns the ControllerVmUuid field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetControllerVmUuid() string {
	if o == nil || o.ControllerVmUuid == nil {
		var ret string
		return ret
	}
	return *o.ControllerVmUuid
}

// GetControllerVmUuidOk returns a tuple with the ControllerVmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetControllerVmUuidOk() (*string, bool) {
	if o == nil || o.ControllerVmUuid == nil {
		return nil, false
	}
	return o.ControllerVmUuid, true
}

// HasControllerVmUuid returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasControllerVmUuid() bool {
	if o != nil && o.ControllerVmUuid != nil {
		return true
	}

	return false
}

// SetControllerVmUuid gets a reference to the given string and assigns it to the ControllerVmUuid field.
func (o *HyperflexHypervisorHost) SetControllerVmUuid(v string) {
	o.ControllerVmUuid = &v
}

// GetControllerVmVersion returns the ControllerVmVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetControllerVmVersion() string {
	if o == nil || o.ControllerVmVersion == nil {
		var ret string
		return ret
	}
	return *o.ControllerVmVersion
}

// GetControllerVmVersionOk returns a tuple with the ControllerVmVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetControllerVmVersionOk() (*string, bool) {
	if o == nil || o.ControllerVmVersion == nil {
		return nil, false
	}
	return o.ControllerVmVersion, true
}

// HasControllerVmVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasControllerVmVersion() bool {
	if o != nil && o.ControllerVmVersion != nil {
		return true
	}

	return false
}

// SetControllerVmVersion gets a reference to the given string and assigns it to the ControllerVmVersion field.
func (o *HyperflexHypervisorHost) SetControllerVmVersion(v string) {
	o.ControllerVmVersion = &v
}

// GetDataIp returns the DataIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHypervisorHost) GetDataIp() NetworkHyperFlexNetworkAddress {
	if o == nil || o.DataIp.Get() == nil {
		var ret NetworkHyperFlexNetworkAddress
		return ret
	}
	return *o.DataIp.Get()
}

// GetDataIpOk returns a tuple with the DataIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHypervisorHost) GetDataIpOk() (*NetworkHyperFlexNetworkAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataIp.Get(), o.DataIp.IsSet()
}

// HasDataIp returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasDataIp() bool {
	if o != nil && o.DataIp.IsSet() {
		return true
	}

	return false
}

// SetDataIp gets a reference to the given NullableNetworkHyperFlexNetworkAddress and assigns it to the DataIp field.
func (o *HyperflexHypervisorHost) SetDataIp(v NetworkHyperFlexNetworkAddress) {
	o.DataIp.Set(&v)
}

// SetDataIpNil sets the value for DataIp to be an explicit nil
func (o *HyperflexHypervisorHost) SetDataIpNil() {
	o.DataIp.Set(nil)
}

// UnsetDataIp ensures that no value is present for DataIp, not even an explicit nil
func (o *HyperflexHypervisorHost) UnsetDataIp() {
	o.DataIp.Unset()
}

// GetHostStatus returns the HostStatus field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetHostStatus() string {
	if o == nil || o.HostStatus == nil {
		var ret string
		return ret
	}
	return *o.HostStatus
}

// GetHostStatusOk returns a tuple with the HostStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetHostStatusOk() (*string, bool) {
	if o == nil || o.HostStatus == nil {
		return nil, false
	}
	return o.HostStatus, true
}

// HasHostStatus returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasHostStatus() bool {
	if o != nil && o.HostStatus != nil {
		return true
	}

	return false
}

// SetHostStatus gets a reference to the given string and assigns it to the HostStatus field.
func (o *HyperflexHypervisorHost) SetHostStatus(v string) {
	o.HostStatus = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetHypervisor() string {
	if o == nil || o.Hypervisor == nil {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetHypervisorOk() (*string, bool) {
	if o == nil || o.Hypervisor == nil {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasHypervisor() bool {
	if o != nil && o.Hypervisor != nil {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *HyperflexHypervisorHost) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHypervisorHost) GetIp() NetworkHyperFlexNetworkAddress {
	if o == nil || o.Ip.Get() == nil {
		var ret NetworkHyperFlexNetworkAddress
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHypervisorHost) GetIpOk() (*NetworkHyperFlexNetworkAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableNetworkHyperFlexNetworkAddress and assigns it to the Ip field.
func (o *HyperflexHypervisorHost) SetIp(v NetworkHyperFlexNetworkAddress) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *HyperflexHypervisorHost) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HyperflexHypervisorHost) UnsetIp() {
	o.Ip.Unset()
}

// GetLockdown returns the Lockdown field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetLockdown() bool {
	if o == nil || o.Lockdown == nil {
		var ret bool
		return ret
	}
	return *o.Lockdown
}

// GetLockdownOk returns a tuple with the Lockdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetLockdownOk() (*bool, bool) {
	if o == nil || o.Lockdown == nil {
		return nil, false
	}
	return o.Lockdown, true
}

// HasLockdown returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasLockdown() bool {
	if o != nil && o.Lockdown != nil {
		return true
	}

	return false
}

// SetLockdown gets a reference to the given bool and assigns it to the Lockdown field.
func (o *HyperflexHypervisorHost) SetLockdown(v bool) {
	o.Lockdown = &v
}

// GetMgmtIp returns the MgmtIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHypervisorHost) GetMgmtIp() NetworkHyperFlexNetworkAddress {
	if o == nil || o.MgmtIp.Get() == nil {
		var ret NetworkHyperFlexNetworkAddress
		return ret
	}
	return *o.MgmtIp.Get()
}

// GetMgmtIpOk returns a tuple with the MgmtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHypervisorHost) GetMgmtIpOk() (*NetworkHyperFlexNetworkAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.MgmtIp.Get(), o.MgmtIp.IsSet()
}

// HasMgmtIp returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasMgmtIp() bool {
	if o != nil && o.MgmtIp.IsSet() {
		return true
	}

	return false
}

// SetMgmtIp gets a reference to the given NullableNetworkHyperFlexNetworkAddress and assigns it to the MgmtIp field.
func (o *HyperflexHypervisorHost) SetMgmtIp(v NetworkHyperFlexNetworkAddress) {
	o.MgmtIp.Set(&v)
}

// SetMgmtIpNil sets the value for MgmtIp to be an explicit nil
func (o *HyperflexHypervisorHost) SetMgmtIpNil() {
	o.MgmtIp.Set(nil)
}

// UnsetMgmtIp ensures that no value is present for MgmtIp, not even an explicit nil
func (o *HyperflexHypervisorHost) UnsetMgmtIp() {
	o.MgmtIp.Unset()
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HyperflexHypervisorHost) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *HyperflexHypervisorHost) SetRole(v string) {
	o.Role = &v
}

// GetTemplateVersion returns the TemplateVersion field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetTemplateVersion() string {
	if o == nil || o.TemplateVersion == nil {
		var ret string
		return ret
	}
	return *o.TemplateVersion
}

// GetTemplateVersionOk returns a tuple with the TemplateVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetTemplateVersionOk() (*string, bool) {
	if o == nil || o.TemplateVersion == nil {
		return nil, false
	}
	return o.TemplateVersion, true
}

// HasTemplateVersion returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasTemplateVersion() bool {
	if o != nil && o.TemplateVersion != nil {
		return true
	}

	return false
}

// SetTemplateVersion gets a reference to the given string and assigns it to the TemplateVersion field.
func (o *HyperflexHypervisorHost) SetTemplateVersion(v string) {
	o.TemplateVersion = &v
}

// GetVirtualCpus returns the VirtualCpus field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetVirtualCpus() int32 {
	if o == nil || o.VirtualCpus == nil {
		var ret int32
		return ret
	}
	return *o.VirtualCpus
}

// GetVirtualCpusOk returns a tuple with the VirtualCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetVirtualCpusOk() (*int32, bool) {
	if o == nil || o.VirtualCpus == nil {
		return nil, false
	}
	return o.VirtualCpus, true
}

// HasVirtualCpus returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasVirtualCpus() bool {
	if o != nil && o.VirtualCpus != nil {
		return true
	}

	return false
}

// SetVirtualCpus gets a reference to the given int32 and assigns it to the VirtualCpus field.
func (o *HyperflexHypervisorHost) SetVirtualCpus(v int32) {
	o.VirtualCpus = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetCluster() HyperflexClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexHypervisorHost) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *HyperflexHypervisorHost) GetNode() HyperflexNodeRelationship {
	if o == nil || o.Node == nil {
		var ret HyperflexNodeRelationship
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHypervisorHost) GetNodeOk() (*HyperflexNodeRelationship, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *HyperflexHypervisorHost) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given HyperflexNodeRelationship and assigns it to the Node field.
func (o *HyperflexHypervisorHost) SetNode(v HyperflexNodeRelationship) {
	o.Node = &v
}

func (o HyperflexHypervisorHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseHost, errVirtualizationBaseHost := json.Marshal(o.VirtualizationBaseHost)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	errVirtualizationBaseHost = json.Unmarshal([]byte(serializedVirtualizationBaseHost), &toSerialize)
	if errVirtualizationBaseHost != nil {
		return []byte{}, errVirtualizationBaseHost
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfiguredMemory != nil {
		toSerialize["ConfiguredMemory"] = o.ConfiguredMemory
	}
	if o.ControllerVmDisplayVersion != nil {
		toSerialize["ControllerVmDisplayVersion"] = o.ControllerVmDisplayVersion
	}
	if o.ControllerVmHxdpDisplayVersion != nil {
		toSerialize["ControllerVmHxdpDisplayVersion"] = o.ControllerVmHxdpDisplayVersion
	}
	if o.ControllerVmHxdpVersion != nil {
		toSerialize["ControllerVmHxdpVersion"] = o.ControllerVmHxdpVersion
	}
	if o.ControllerVmUuid != nil {
		toSerialize["ControllerVmUuid"] = o.ControllerVmUuid
	}
	if o.ControllerVmVersion != nil {
		toSerialize["ControllerVmVersion"] = o.ControllerVmVersion
	}
	if o.DataIp.IsSet() {
		toSerialize["DataIp"] = o.DataIp.Get()
	}
	if o.HostStatus != nil {
		toSerialize["HostStatus"] = o.HostStatus
	}
	if o.Hypervisor != nil {
		toSerialize["Hypervisor"] = o.Hypervisor
	}
	if o.Ip.IsSet() {
		toSerialize["Ip"] = o.Ip.Get()
	}
	if o.Lockdown != nil {
		toSerialize["Lockdown"] = o.Lockdown
	}
	if o.MgmtIp.IsSet() {
		toSerialize["MgmtIp"] = o.MgmtIp.Get()
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.TemplateVersion != nil {
		toSerialize["TemplateVersion"] = o.TemplateVersion
	}
	if o.VirtualCpus != nil {
		toSerialize["VirtualCpus"] = o.VirtualCpus
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Node != nil {
		toSerialize["Node"] = o.Node
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHypervisorHost) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexHypervisorHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Memory configured for controller virtual machine.
		ConfiguredMemory *int64 `json:"ConfiguredMemory,omitempty"`
		// The display version of HyperFlex software running on the controller VM.
		ControllerVmDisplayVersion *string `json:"ControllerVmDisplayVersion,omitempty"`
		// Platform storage software product display version running on controller VM.
		ControllerVmHxdpDisplayVersion *string `json:"ControllerVmHxdpDisplayVersion,omitempty"`
		// Platform storage software product version running on controller VM.
		ControllerVmHxdpVersion *string `json:"ControllerVmHxdpVersion,omitempty"`
		// The UUID of the controller VM which belongs to this host.
		ControllerVmUuid *string `json:"ControllerVmUuid,omitempty"`
		// The version of HyperFlex software running on the controller VM.
		ControllerVmVersion *string                                `json:"ControllerVmVersion,omitempty"`
		DataIp              NullableNetworkHyperFlexNetworkAddress `json:"DataIp,omitempty"`
		// The status of the HyperFlex host. * `UNKNOWN` - Current status of the HyperFlex host is unknown. * `ONLINE` - The HyperFlex host is online. * `OFFLINE` - The HyperFlex host is offline. * `INMAINTENANCE` - The HyperFlex host is in maintenance mode. * `DEGRADED` - Current status of the HyperFlex virtual machine is degraded.
		HostStatus *string `json:"HostStatus,omitempty"`
		// The hypervisor type of the host.
		Hypervisor *string                                `json:"Hypervisor,omitempty"`
		Ip         NullableNetworkHyperFlexNetworkAddress `json:"Ip,omitempty"`
		// Flag indicating whether the HyperFlex host is in lockdown mode.
		Lockdown *bool                                  `json:"Lockdown,omitempty"`
		MgmtIp   NullableNetworkHyperFlexNetworkAddress `json:"MgmtIp,omitempty"`
		// The operation system version of the controller VM.
		OsVersion *string `json:"OsVersion,omitempty"`
		// The role of the HyperFlex host. * `UNKNOWN` - The role of the HyperFlex host is unknown. * `STORAGE` - The HyperFlex host's role is storage. * `COMPUTE` - The HyperFlex host's role is compute.
		Role *string `json:"Role,omitempty"`
		// The controller virtual machine template version.
		TemplateVersion *string `json:"TemplateVersion,omitempty"`
		// Configured number of virtual CPUs for Controller virtual machine.
		VirtualCpus *int32                        `json:"VirtualCpus,omitempty"`
		Cluster     *HyperflexClusterRelationship `json:"Cluster,omitempty"`
		Node        *HyperflexNodeRelationship    `json:"Node,omitempty"`
	}

	varHyperflexHypervisorHostWithoutEmbeddedStruct := HyperflexHypervisorHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexHypervisorHostWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHypervisorHost := _HyperflexHypervisorHost{}
		varHyperflexHypervisorHost.ClassId = varHyperflexHypervisorHostWithoutEmbeddedStruct.ClassId
		varHyperflexHypervisorHost.ObjectType = varHyperflexHypervisorHostWithoutEmbeddedStruct.ObjectType
		varHyperflexHypervisorHost.ConfiguredMemory = varHyperflexHypervisorHostWithoutEmbeddedStruct.ConfiguredMemory
		varHyperflexHypervisorHost.ControllerVmDisplayVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.ControllerVmDisplayVersion
		varHyperflexHypervisorHost.ControllerVmHxdpDisplayVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.ControllerVmHxdpDisplayVersion
		varHyperflexHypervisorHost.ControllerVmHxdpVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.ControllerVmHxdpVersion
		varHyperflexHypervisorHost.ControllerVmUuid = varHyperflexHypervisorHostWithoutEmbeddedStruct.ControllerVmUuid
		varHyperflexHypervisorHost.ControllerVmVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.ControllerVmVersion
		varHyperflexHypervisorHost.DataIp = varHyperflexHypervisorHostWithoutEmbeddedStruct.DataIp
		varHyperflexHypervisorHost.HostStatus = varHyperflexHypervisorHostWithoutEmbeddedStruct.HostStatus
		varHyperflexHypervisorHost.Hypervisor = varHyperflexHypervisorHostWithoutEmbeddedStruct.Hypervisor
		varHyperflexHypervisorHost.Ip = varHyperflexHypervisorHostWithoutEmbeddedStruct.Ip
		varHyperflexHypervisorHost.Lockdown = varHyperflexHypervisorHostWithoutEmbeddedStruct.Lockdown
		varHyperflexHypervisorHost.MgmtIp = varHyperflexHypervisorHostWithoutEmbeddedStruct.MgmtIp
		varHyperflexHypervisorHost.OsVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.OsVersion
		varHyperflexHypervisorHost.Role = varHyperflexHypervisorHostWithoutEmbeddedStruct.Role
		varHyperflexHypervisorHost.TemplateVersion = varHyperflexHypervisorHostWithoutEmbeddedStruct.TemplateVersion
		varHyperflexHypervisorHost.VirtualCpus = varHyperflexHypervisorHostWithoutEmbeddedStruct.VirtualCpus
		varHyperflexHypervisorHost.Cluster = varHyperflexHypervisorHostWithoutEmbeddedStruct.Cluster
		varHyperflexHypervisorHost.Node = varHyperflexHypervisorHostWithoutEmbeddedStruct.Node
		*o = HyperflexHypervisorHost(varHyperflexHypervisorHost)
	} else {
		return err
	}

	varHyperflexHypervisorHost := _HyperflexHypervisorHost{}

	err = json.Unmarshal(bytes, &varHyperflexHypervisorHost)
	if err == nil {
		o.VirtualizationBaseHost = varHyperflexHypervisorHost.VirtualizationBaseHost
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfiguredMemory")
		delete(additionalProperties, "ControllerVmDisplayVersion")
		delete(additionalProperties, "ControllerVmHxdpDisplayVersion")
		delete(additionalProperties, "ControllerVmHxdpVersion")
		delete(additionalProperties, "ControllerVmUuid")
		delete(additionalProperties, "ControllerVmVersion")
		delete(additionalProperties, "DataIp")
		delete(additionalProperties, "HostStatus")
		delete(additionalProperties, "Hypervisor")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "Lockdown")
		delete(additionalProperties, "MgmtIp")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "TemplateVersion")
		delete(additionalProperties, "VirtualCpus")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Node")

		// remove fields from embedded structs
		reflectVirtualizationBaseHost := reflect.ValueOf(o.VirtualizationBaseHost)
		for i := 0; i < reflectVirtualizationBaseHost.Type().NumField(); i++ {
			t := reflectVirtualizationBaseHost.Type().Field(i)

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

type NullableHyperflexHypervisorHost struct {
	value *HyperflexHypervisorHost
	isSet bool
}

func (v NullableHyperflexHypervisorHost) Get() *HyperflexHypervisorHost {
	return v.value
}

func (v *NullableHyperflexHypervisorHost) Set(val *HyperflexHypervisorHost) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHypervisorHost) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHypervisorHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHypervisorHost(val *HyperflexHypervisorHost) *NullableHyperflexHypervisorHost {
	return &NullableHyperflexHypervisorHost{value: val, isSet: true}
}

func (v NullableHyperflexHypervisorHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHypervisorHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
