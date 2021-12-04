/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentIoCardBaseAllOf Definition of the list of properties defined in 'equipment.IoCardBase', excluding properties defined in parent classes.
type EquipmentIoCardBaseAllOf struct {
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

type _EquipmentIoCardBaseAllOf EquipmentIoCardBaseAllOf

// NewEquipmentIoCardBaseAllOf instantiates a new EquipmentIoCardBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardBaseAllOf(classId string, objectType string) *EquipmentIoCardBaseAllOf {
	this := EquipmentIoCardBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardBaseAllOfWithDefaults instantiates a new EquipmentIoCardBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardBaseAllOfWithDefaults() *EquipmentIoCardBaseAllOf {
	this := EquipmentIoCardBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentIoCardBaseAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentIoCardBaseAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentIoCardBaseAllOf) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBaseAllOf) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBaseAllOf) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentIoCardBaseAllOf) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentIoCardBaseAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentIoCardBaseAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentIoCardBaseAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EquipmentIoCardBaseAllOf) SetProductName(v string) {
	o.ProductName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentIoCardBaseAllOf) SetSku(v string) {
	o.Sku = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EquipmentIoCardBaseAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentIoCardBaseAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetHostPorts returns the HostPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBaseAllOf) GetHostPorts() []EtherHostPortRelationship {
	if o == nil {
		var ret []EtherHostPortRelationship
		return ret
	}
	return o.HostPorts
}

// GetHostPortsOk returns a tuple with the HostPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBaseAllOf) GetHostPortsOk() (*[]EtherHostPortRelationship, bool) {
	if o == nil || o.HostPorts == nil {
		return nil, false
	}
	return &o.HostPorts, true
}

// HasHostPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasHostPorts() bool {
	if o != nil && o.HostPorts != nil {
		return true
	}

	return false
}

// SetHostPorts gets a reference to the given []EtherHostPortRelationship and assigns it to the HostPorts field.
func (o *EquipmentIoCardBaseAllOf) SetHostPorts(v []EtherHostPortRelationship) {
	o.HostPorts = v
}

// GetMgmtController returns the MgmtController field value if set, zero value otherwise.
func (o *EquipmentIoCardBaseAllOf) GetMgmtController() ManagementControllerRelationship {
	if o == nil || o.MgmtController == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.MgmtController
}

// GetMgmtControllerOk returns a tuple with the MgmtController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardBaseAllOf) GetMgmtControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.MgmtController == nil {
		return nil, false
	}
	return o.MgmtController, true
}

// HasMgmtController returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasMgmtController() bool {
	if o != nil && o.MgmtController != nil {
		return true
	}

	return false
}

// SetMgmtController gets a reference to the given ManagementControllerRelationship and assigns it to the MgmtController field.
func (o *EquipmentIoCardBaseAllOf) SetMgmtController(v ManagementControllerRelationship) {
	o.MgmtController = &v
}

// GetNetworkPorts returns the NetworkPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentIoCardBaseAllOf) GetNetworkPorts() []EtherNetworkPortRelationship {
	if o == nil {
		var ret []EtherNetworkPortRelationship
		return ret
	}
	return o.NetworkPorts
}

// GetNetworkPortsOk returns a tuple with the NetworkPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentIoCardBaseAllOf) GetNetworkPortsOk() (*[]EtherNetworkPortRelationship, bool) {
	if o == nil || o.NetworkPorts == nil {
		return nil, false
	}
	return &o.NetworkPorts, true
}

// HasNetworkPorts returns a boolean if a field has been set.
func (o *EquipmentIoCardBaseAllOf) HasNetworkPorts() bool {
	if o != nil && o.NetworkPorts != nil {
		return true
	}

	return false
}

// SetNetworkPorts gets a reference to the given []EtherNetworkPortRelationship and assigns it to the NetworkPorts field.
func (o *EquipmentIoCardBaseAllOf) SetNetworkPorts(v []EtherNetworkPortRelationship) {
	o.NetworkPorts = v
}

func (o EquipmentIoCardBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *EquipmentIoCardBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoCardBaseAllOf := _EquipmentIoCardBaseAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoCardBaseAllOf); err == nil {
		*o = EquipmentIoCardBaseAllOf(varEquipmentIoCardBaseAllOf)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoCardBaseAllOf struct {
	value *EquipmentIoCardBaseAllOf
	isSet bool
}

func (v NullableEquipmentIoCardBaseAllOf) Get() *EquipmentIoCardBaseAllOf {
	return v.value
}

func (v *NullableEquipmentIoCardBaseAllOf) Set(val *EquipmentIoCardBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardBaseAllOf(val *EquipmentIoCardBaseAllOf) *NullableEquipmentIoCardBaseAllOf {
	return &NullableEquipmentIoCardBaseAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoCardBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
