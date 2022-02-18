/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentIoCardBase Abstract class for hif and nif ports collection that can be extended by chassis/febric extender.
type EquipmentIoCardBase struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Connectivity Status of FEX/IOM to Switch - A or B or AB.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field is to provide description for the iocard module model.
	Description *string `json:"Description,omitempty"`
	// Module Identifier for the IO module.
	ModuleId   *int64   `json:"ModuleId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of IO card or fabric extender.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the IO module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the IO module.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the Product Name for the iocard module model.
	ProductName *string `json:"ProductName,omitempty"`
	// This field identifies the Stock Keeping Unit for the IO card module.
	Sku *string `json:"Sku,omitempty"`
	// This field identifies the version of the IO card module.
	Version *string `json:"Version,omitempty"`
	// This field identifies the Vendor ID for the IO card module.
	Vid *string `json:"Vid,omitempty"`
	// An array of relationships to etherHostPort resources.
	HostPorts      []EtherHostPortRelationship       `json:"HostPorts,omitempty"`
	MgmtController *ManagementControllerRelationship `json:"MgmtController,omitempty"`
	// An array of relationships to etherNetworkPort resources.
	NetworkPorts         []EtherNetworkPortRelationship `json:"NetworkPorts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardBase EquipmentIoCardBase

// NewEquipmentIoCardBase instantiates a new EquipmentIoCardBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardBase(classId string, objectType string) *EquipmentIoCardBase {
	this := EquipmentIoCardBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardBaseWithDefaults instantiates a new EquipmentIoCardBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardBaseWithDefaults() *EquipmentIoCardBase {
	this := EquipmentIoCardBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentIoCardBase) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentIoCardBase) SetDescription(v string) {
	o.Description = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentIoCardBase) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentIoCardBase) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentIoCardBase) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentIoCardBase) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentIoCardBase) SetPid(v string) {
	o.Pid = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EquipmentIoCardBase) SetProductName(v string) {
	o.ProductName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentIoCardBase) SetSku(v string) {
	o.Sku = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentIoCardBase) SetVersion(v string) {
	o.Version = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentIoCardBase) SetVid(v string) {
	o.Vid = &v
}

// GetHostPorts returns the HostPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetHostPorts() []EtherHostPortRelationship {
	if o == nil {
		var ret []EtherHostPortRelationship
		return ret
	}
	return o.HostPorts
}

// GetHostPortsOk returns a tuple with the HostPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetHostPortsOk() (*[]EtherHostPortRelationship, bool) {
	if o == nil || o.HostPorts == nil {
		return nil, false
	}
	return &o.HostPorts, true
}

// HasHostPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasHostPorts() bool {
	if o != nil && o.HostPorts != nil {
		return true
	}

	return false
}

// SetHostPorts gets a reference to the given []EtherHostPortRelationship and assigns it to the HostPorts field.
func (o *EquipmentIoCardBase) SetHostPorts(v []EtherHostPortRelationship) {
	o.HostPorts = v
}

// GetMgmtController returns the MgmtController field value if set, zero value otherwise.
func (o *EquipmentIoCardBase) GetMgmtController() ManagementControllerRelationship {
	if o == nil || o.MgmtController == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.MgmtController
}

// GetMgmtControllerOk returns a tuple with the MgmtController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBase) GetMgmtControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.MgmtController == nil {
		return nil, false
	}
	return o.MgmtController, true
}

// HasMgmtController returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasMgmtController() bool {
	if o != nil && o.MgmtController != nil {
		return true
	}

	return false
}

// SetMgmtController gets a reference to the given ManagementControllerRelationship and assigns it to the MgmtController field.
func (o *EquipmentIoCardBase) SetMgmtController(v ManagementControllerRelationship) {
	o.MgmtController = &v
}

// GetNetworkPorts returns the NetworkPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBase) GetNetworkPorts() []EtherNetworkPortRelationship {
	if o == nil {
		var ret []EtherNetworkPortRelationship
		return ret
	}
	return o.NetworkPorts
}

// GetNetworkPortsOk returns a tuple with the NetworkPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBase) GetNetworkPortsOk() (*[]EtherNetworkPortRelationship, bool) {
	if o == nil || o.NetworkPorts == nil {
		return nil, false
	}
	return &o.NetworkPorts, true
}

// HasNetworkPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBase) HasNetworkPorts() bool {
	if o != nil && o.NetworkPorts != nil {
		return true
	}

	return false
}

// SetNetworkPorts gets a reference to the given []EtherNetworkPortRelationship and assigns it to the NetworkPorts field.
func (o *EquipmentIoCardBase) SetNetworkPorts(v []EtherNetworkPortRelationship) {
	o.NetworkPorts = v
}

func (o EquipmentIoCardBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.HostPorts != nil {
		toSerialize["HostPorts"] = o.HostPorts
	}
	if o.MgmtController != nil {
		toSerialize["MgmtController"] = o.MgmtController
	}
	if o.NetworkPorts != nil {
		toSerialize["NetworkPorts"] = o.NetworkPorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCardBase) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentIoCardBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Connectivity Status of FEX/IOM to Switch - A or B or AB.
		ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
		// This field is to provide description for the iocard module model.
		Description *string `json:"Description,omitempty"`
		// Module Identifier for the IO module.
		ModuleId   *int64   `json:"ModuleId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational state of IO card or fabric extender.
		OperState *string `json:"OperState,omitempty"`
		// Part Number identifier for the IO module.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for the IO module.
		Pid *string `json:"Pid,omitempty"`
		// This field identifies the Product Name for the iocard module model.
		ProductName *string `json:"ProductName,omitempty"`
		// This field identifies the Stock Keeping Unit for the IO card module.
		Sku *string `json:"Sku,omitempty"`
		// This field identifies the version of the IO card module.
		Version *string `json:"Version,omitempty"`
		// This field identifies the Vendor ID for the IO card module.
		Vid *string `json:"Vid,omitempty"`
		// An array of relationships to etherHostPort resources.
		HostPorts      []EtherHostPortRelationship       `json:"HostPorts,omitempty"`
		MgmtController *ManagementControllerRelationship `json:"MgmtController,omitempty"`
		// An array of relationships to etherNetworkPort resources.
		NetworkPorts []EtherNetworkPortRelationship `json:"NetworkPorts,omitempty"`
	}

	varEquipmentIoCardBaseWithoutEmbeddedStruct := EquipmentIoCardBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentIoCardBaseWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentIoCardBase := _EquipmentIoCardBase{}
		varEquipmentIoCardBase.ClassId = varEquipmentIoCardBaseWithoutEmbeddedStruct.ClassId
		varEquipmentIoCardBase.ObjectType = varEquipmentIoCardBaseWithoutEmbeddedStruct.ObjectType
		varEquipmentIoCardBase.ConnectionStatus = varEquipmentIoCardBaseWithoutEmbeddedStruct.ConnectionStatus
		varEquipmentIoCardBase.Description = varEquipmentIoCardBaseWithoutEmbeddedStruct.Description
		varEquipmentIoCardBase.ModuleId = varEquipmentIoCardBaseWithoutEmbeddedStruct.ModuleId
		varEquipmentIoCardBase.OperReason = varEquipmentIoCardBaseWithoutEmbeddedStruct.OperReason
		varEquipmentIoCardBase.OperState = varEquipmentIoCardBaseWithoutEmbeddedStruct.OperState
		varEquipmentIoCardBase.PartNumber = varEquipmentIoCardBaseWithoutEmbeddedStruct.PartNumber
		varEquipmentIoCardBase.Pid = varEquipmentIoCardBaseWithoutEmbeddedStruct.Pid
		varEquipmentIoCardBase.ProductName = varEquipmentIoCardBaseWithoutEmbeddedStruct.ProductName
		varEquipmentIoCardBase.Sku = varEquipmentIoCardBaseWithoutEmbeddedStruct.Sku
		varEquipmentIoCardBase.Version = varEquipmentIoCardBaseWithoutEmbeddedStruct.Version
		varEquipmentIoCardBase.Vid = varEquipmentIoCardBaseWithoutEmbeddedStruct.Vid
		varEquipmentIoCardBase.HostPorts = varEquipmentIoCardBaseWithoutEmbeddedStruct.HostPorts
		varEquipmentIoCardBase.MgmtController = varEquipmentIoCardBaseWithoutEmbeddedStruct.MgmtController
		varEquipmentIoCardBase.NetworkPorts = varEquipmentIoCardBaseWithoutEmbeddedStruct.NetworkPorts
		*o = EquipmentIoCardBase(varEquipmentIoCardBase)
	} else {
		return err
	}

	varEquipmentIoCardBase := _EquipmentIoCardBase{}

	err = json.Unmarshal(bytes, &varEquipmentIoCardBase)
	if err == nil {
		o.EquipmentBase = varEquipmentIoCardBase.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "HostPorts")
		delete(additionalProperties, "MgmtController")
		delete(additionalProperties, "NetworkPorts")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentIoCardBase struct {
	value *EquipmentIoCardBase
	isSet bool
}

func (v NullableEquipmentIoCardBase) Get() *EquipmentIoCardBase {
	return v.value
}

func (v *NullableEquipmentIoCardBase) Set(val *EquipmentIoCardBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardBase(val *EquipmentIoCardBase) *NullableEquipmentIoCardBase {
	return &NullableEquipmentIoCardBase{value: val, isSet: true}
}

func (v NullableEquipmentIoCardBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
