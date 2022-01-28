/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationHost Depicts operations to control the life cycle of a Hypervisor Host.
type VirtualizationHost struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Action to be performed on a host (Create, PowerState, Migrate, Clone etc). * `None` - A place holder for the default value. * `EnterMaintenanceMode` - Put a host into maintenance mode. * `ExitMaintenanceMode` - Put a host into active mode. * `PowerOffStorageController` - Power off HyperFlex storage controller node running on selected hypervisor host. * `PowerOnStorageController` - Power on HyperFlex storage controller node running on selected hypervisor host.
	Action *string `json:"Action,omitempty"`
	// Flag to indicate whether the configuration is created from inventory object.
	Discovered *bool `json:"Discovered,omitempty"`
	// If true, move powered-off and suspended virtual machines to other hosts in the cluster.
	Evacuate   *bool                                       `json:"Evacuate,omitempty"`
	HostConfig NullableVirtualizationBaseHostConfiguration `json:"HostConfig,omitempty"`
	// Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
	HypervisorType *string `json:"HypervisorType,omitempty"`
	// Unique identifier assigned to the hypervisor host.
	Identity *string `json:"Identity,omitempty"`
	// Expected state of host (enter maintenance, exit maintenance). * `None` - A place holder for the default value. * `Enter` - Power action is performed on the virtual machine. * `Exit` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.
	MaintenanceState *string `json:"MaintenanceState,omitempty"`
	// Commercial model information about this hardware.
	Model *string `json:"Model,omitempty"`
	// Name of the hypervisor host. It must be unique within the target endpoint.
	Name *string `json:"Name,omitempty"`
	// Serial number of this host (internally generated).
	Serial *string `json:"Serial,omitempty"`
	// Commercial vendor details of this hardware.
	Vendor               *string                              `json:"Vendor,omitempty"`
	Inventory            *VirtualizationBaseHostRelationship  `json:"Inventory,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	WorkflowInfo         *WorkflowWorkflowInfoRelationship    `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationHost VirtualizationHost

// NewVirtualizationHost instantiates a new VirtualizationHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationHost(classId string, objectType string) *VirtualizationHost {
	this := VirtualizationHost{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var maintenanceState string = "None"
	this.MaintenanceState = &maintenanceState
	return &this
}

// NewVirtualizationHostWithDefaults instantiates a new VirtualizationHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationHostWithDefaults() *VirtualizationHost {
	this := VirtualizationHost{}
	var classId string = "virtualization.Host"
	this.ClassId = classId
	var objectType string = "virtualization.Host"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	var maintenanceState string = "None"
	this.MaintenanceState = &maintenanceState
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationHost) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationHost) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationHost) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationHost) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *VirtualizationHost) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *VirtualizationHost) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *VirtualizationHost) SetAction(v string) {
	o.Action = &v
}

// GetDiscovered returns the Discovered field value if set, zero value otherwise.
func (o *VirtualizationHost) GetDiscovered() bool {
	if o == nil || o.Discovered == nil {
		var ret bool
		return ret
	}
	return *o.Discovered
}

// GetDiscoveredOk returns a tuple with the Discovered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetDiscoveredOk() (*bool, bool) {
	if o == nil || o.Discovered == nil {
		return nil, false
	}
	return o.Discovered, true
}

// HasDiscovered returns a boolean if a field has been set.
func (o *VirtualizationHost) HasDiscovered() bool {
	if o != nil && o.Discovered != nil {
		return true
	}

	return false
}

// SetDiscovered gets a reference to the given bool and assigns it to the Discovered field.
func (o *VirtualizationHost) SetDiscovered(v bool) {
	o.Discovered = &v
}

// GetEvacuate returns the Evacuate field value if set, zero value otherwise.
func (o *VirtualizationHost) GetEvacuate() bool {
	if o == nil || o.Evacuate == nil {
		var ret bool
		return ret
	}
	return *o.Evacuate
}

// GetEvacuateOk returns a tuple with the Evacuate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetEvacuateOk() (*bool, bool) {
	if o == nil || o.Evacuate == nil {
		return nil, false
	}
	return o.Evacuate, true
}

// HasEvacuate returns a boolean if a field has been set.
func (o *VirtualizationHost) HasEvacuate() bool {
	if o != nil && o.Evacuate != nil {
		return true
	}

	return false
}

// SetEvacuate gets a reference to the given bool and assigns it to the Evacuate field.
func (o *VirtualizationHost) SetEvacuate(v bool) {
	o.Evacuate = &v
}

// GetHostConfig returns the HostConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationHost) GetHostConfig() VirtualizationBaseHostConfiguration {
	if o == nil || o.HostConfig.Get() == nil {
		var ret VirtualizationBaseHostConfiguration
		return ret
	}
	return *o.HostConfig.Get()
}

// GetHostConfigOk returns a tuple with the HostConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationHost) GetHostConfigOk() (*VirtualizationBaseHostConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostConfig.Get(), o.HostConfig.IsSet()
}

// HasHostConfig returns a boolean if a field has been set.
func (o *VirtualizationHost) HasHostConfig() bool {
	if o != nil && o.HostConfig.IsSet() {
		return true
	}

	return false
}

// SetHostConfig gets a reference to the given NullableVirtualizationBaseHostConfiguration and assigns it to the HostConfig field.
func (o *VirtualizationHost) SetHostConfig(v VirtualizationBaseHostConfiguration) {
	o.HostConfig.Set(&v)
}

// SetHostConfigNil sets the value for HostConfig to be an explicit nil
func (o *VirtualizationHost) SetHostConfigNil() {
	o.HostConfig.Set(nil)
}

// UnsetHostConfig ensures that no value is present for HostConfig, not even an explicit nil
func (o *VirtualizationHost) UnsetHostConfig() {
	o.HostConfig.Unset()
}

// GetHypervisorType returns the HypervisorType field value if set, zero value otherwise.
func (o *VirtualizationHost) GetHypervisorType() string {
	if o == nil || o.HypervisorType == nil {
		var ret string
		return ret
	}
	return *o.HypervisorType
}

// GetHypervisorTypeOk returns a tuple with the HypervisorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetHypervisorTypeOk() (*string, bool) {
	if o == nil || o.HypervisorType == nil {
		return nil, false
	}
	return o.HypervisorType, true
}

// HasHypervisorType returns a boolean if a field has been set.
func (o *VirtualizationHost) HasHypervisorType() bool {
	if o != nil && o.HypervisorType != nil {
		return true
	}

	return false
}

// SetHypervisorType gets a reference to the given string and assigns it to the HypervisorType field.
func (o *VirtualizationHost) SetHypervisorType(v string) {
	o.HypervisorType = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *VirtualizationHost) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *VirtualizationHost) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *VirtualizationHost) SetIdentity(v string) {
	o.Identity = &v
}

// GetMaintenanceState returns the MaintenanceState field value if set, zero value otherwise.
func (o *VirtualizationHost) GetMaintenanceState() string {
	if o == nil || o.MaintenanceState == nil {
		var ret string
		return ret
	}
	return *o.MaintenanceState
}

// GetMaintenanceStateOk returns a tuple with the MaintenanceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetMaintenanceStateOk() (*string, bool) {
	if o == nil || o.MaintenanceState == nil {
		return nil, false
	}
	return o.MaintenanceState, true
}

// HasMaintenanceState returns a boolean if a field has been set.
func (o *VirtualizationHost) HasMaintenanceState() bool {
	if o != nil && o.MaintenanceState != nil {
		return true
	}

	return false
}

// SetMaintenanceState gets a reference to the given string and assigns it to the MaintenanceState field.
func (o *VirtualizationHost) SetMaintenanceState(v string) {
	o.MaintenanceState = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *VirtualizationHost) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *VirtualizationHost) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *VirtualizationHost) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationHost) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationHost) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationHost) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *VirtualizationHost) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *VirtualizationHost) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *VirtualizationHost) SetSerial(v string) {
	o.Serial = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *VirtualizationHost) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *VirtualizationHost) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *VirtualizationHost) SetVendor(v string) {
	o.Vendor = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *VirtualizationHost) GetInventory() VirtualizationBaseHostRelationship {
	if o == nil || o.Inventory == nil {
		var ret VirtualizationBaseHostRelationship
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetInventoryOk() (*VirtualizationBaseHostRelationship, bool) {
	if o == nil || o.Inventory == nil {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *VirtualizationHost) HasInventory() bool {
	if o != nil && o.Inventory != nil {
		return true
	}

	return false
}

// SetInventory gets a reference to the given VirtualizationBaseHostRelationship and assigns it to the Inventory field.
func (o *VirtualizationHost) SetInventory(v VirtualizationBaseHostRelationship) {
	o.Inventory = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationHost) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationHost) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationHost) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise.
func (o *VirtualizationHost) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.WorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationHost) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.WorkflowInfo == nil {
		return nil, false
	}
	return o.WorkflowInfo, true
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *VirtualizationHost) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo != nil {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *VirtualizationHost) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo = &v
}

func (o VirtualizationHost) MarshalJSON() ([]byte, error) {
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
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.Discovered != nil {
		toSerialize["Discovered"] = o.Discovered
	}
	if o.Evacuate != nil {
		toSerialize["Evacuate"] = o.Evacuate
	}
	if o.HostConfig.IsSet() {
		toSerialize["HostConfig"] = o.HostConfig.Get()
	}
	if o.HypervisorType != nil {
		toSerialize["HypervisorType"] = o.HypervisorType
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.MaintenanceState != nil {
		toSerialize["MaintenanceState"] = o.MaintenanceState
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Inventory != nil {
		toSerialize["Inventory"] = o.Inventory
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.WorkflowInfo != nil {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationHost) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationHostWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Action to be performed on a host (Create, PowerState, Migrate, Clone etc). * `None` - A place holder for the default value. * `EnterMaintenanceMode` - Put a host into maintenance mode. * `ExitMaintenanceMode` - Put a host into active mode. * `PowerOffStorageController` - Power off HyperFlex storage controller node running on selected hypervisor host. * `PowerOnStorageController` - Power on HyperFlex storage controller node running on selected hypervisor host.
		Action *string `json:"Action,omitempty"`
		// Flag to indicate whether the configuration is created from inventory object.
		Discovered *bool `json:"Discovered,omitempty"`
		// If true, move powered-off and suspended virtual machines to other hosts in the cluster.
		Evacuate   *bool                                       `json:"Evacuate,omitempty"`
		HostConfig NullableVirtualizationBaseHostConfiguration `json:"HostConfig,omitempty"`
		// Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with. * `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version. * `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform. * `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine. * `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V. * `Unknown` - The hypervisor running on the HyperFlex cluster is not known.
		HypervisorType *string `json:"HypervisorType,omitempty"`
		// Unique identifier assigned to the hypervisor host.
		Identity *string `json:"Identity,omitempty"`
		// Expected state of host (enter maintenance, exit maintenance). * `None` - A place holder for the default value. * `Enter` - Power action is performed on the virtual machine. * `Exit` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.
		MaintenanceState *string `json:"MaintenanceState,omitempty"`
		// Commercial model information about this hardware.
		Model *string `json:"Model,omitempty"`
		// Name of the hypervisor host. It must be unique within the target endpoint.
		Name *string `json:"Name,omitempty"`
		// Serial number of this host (internally generated).
		Serial *string `json:"Serial,omitempty"`
		// Commercial vendor details of this hardware.
		Vendor           *string                              `json:"Vendor,omitempty"`
		Inventory        *VirtualizationBaseHostRelationship  `json:"Inventory,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		WorkflowInfo     *WorkflowWorkflowInfoRelationship    `json:"WorkflowInfo,omitempty"`
	}

	varVirtualizationHostWithoutEmbeddedStruct := VirtualizationHostWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationHostWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationHost := _VirtualizationHost{}
		varVirtualizationHost.ClassId = varVirtualizationHostWithoutEmbeddedStruct.ClassId
		varVirtualizationHost.ObjectType = varVirtualizationHostWithoutEmbeddedStruct.ObjectType
		varVirtualizationHost.Action = varVirtualizationHostWithoutEmbeddedStruct.Action
		varVirtualizationHost.Discovered = varVirtualizationHostWithoutEmbeddedStruct.Discovered
		varVirtualizationHost.Evacuate = varVirtualizationHostWithoutEmbeddedStruct.Evacuate
		varVirtualizationHost.HostConfig = varVirtualizationHostWithoutEmbeddedStruct.HostConfig
		varVirtualizationHost.HypervisorType = varVirtualizationHostWithoutEmbeddedStruct.HypervisorType
		varVirtualizationHost.Identity = varVirtualizationHostWithoutEmbeddedStruct.Identity
		varVirtualizationHost.MaintenanceState = varVirtualizationHostWithoutEmbeddedStruct.MaintenanceState
		varVirtualizationHost.Model = varVirtualizationHostWithoutEmbeddedStruct.Model
		varVirtualizationHost.Name = varVirtualizationHostWithoutEmbeddedStruct.Name
		varVirtualizationHost.Serial = varVirtualizationHostWithoutEmbeddedStruct.Serial
		varVirtualizationHost.Vendor = varVirtualizationHostWithoutEmbeddedStruct.Vendor
		varVirtualizationHost.Inventory = varVirtualizationHostWithoutEmbeddedStruct.Inventory
		varVirtualizationHost.RegisteredDevice = varVirtualizationHostWithoutEmbeddedStruct.RegisteredDevice
		varVirtualizationHost.WorkflowInfo = varVirtualizationHostWithoutEmbeddedStruct.WorkflowInfo
		*o = VirtualizationHost(varVirtualizationHost)
	} else {
		return err
	}

	varVirtualizationHost := _VirtualizationHost{}

	err = json.Unmarshal(bytes, &varVirtualizationHost)
	if err == nil {
		o.MoBaseMo = varVirtualizationHost.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "Discovered")
		delete(additionalProperties, "Evacuate")
		delete(additionalProperties, "HostConfig")
		delete(additionalProperties, "HypervisorType")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "MaintenanceState")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Inventory")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableVirtualizationHost struct {
	value *VirtualizationHost
	isSet bool
}

func (v NullableVirtualizationHost) Get() *VirtualizationHost {
	return v.value
}

func (v *NullableVirtualizationHost) Set(val *VirtualizationHost) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationHost) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationHost(val *VirtualizationHost) *NullableVirtualizationHost {
	return &NullableVirtualizationHost{value: val, isSet: true}
}

func (v NullableVirtualizationHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
