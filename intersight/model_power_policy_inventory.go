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

// PowerPolicyInventory Power Management policy models a configuration that can be applied to Chassis or Server to manage Power Related Features.
type PowerPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis.
	AllocatedBudget *int64 `json:"AllocatedBudget,omitempty"`
	// Sets the Dynamic Power Rebalancing of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	DynamicRebalancing *string `json:"DynamicRebalancing,omitempty"`
	// Sets the Extended Power Capacity of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	ExtendedPowerCapacity *string `json:"ExtendedPowerCapacity,omitempty"`
	// Sets the Power Priority of the System. This field is only supported for Cisco UCS X series servers. * `Low` - Set the Power Priority to Low. * `Medium` - Set the Power Priority to Medium. * `High` - Set the Power Priority to High.
	PowerPriority *string `json:"PowerPriority,omitempty"`
	// Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	PowerProfiling *string `json:"PowerProfiling,omitempty"`
	// Sets the Power Restore State of the Server. This field is only supported for Cisco UCS X series servers. * `AlwaysOff` - Set the Power Restore Mode to Off. * `AlwaysOn` - Set the Power Restore Mode to On. * `LastState` - Set the Power Restore Mode to LastState.
	PowerRestoreState *string `json:"PowerRestoreState,omitempty"`
	// Sets the Power Save mode of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	PowerSaveMode *string `json:"PowerSaveMode,omitempty"`
	// Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis. * `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis.
	RedundancyMode       *string               `json:"RedundancyMode,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerPolicyInventory PowerPolicyInventory

// NewPowerPolicyInventory instantiates a new PowerPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerPolicyInventory(classId string, objectType string) *PowerPolicyInventory {
	this := PowerPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPowerPolicyInventoryWithDefaults instantiates a new PowerPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerPolicyInventoryWithDefaults() *PowerPolicyInventory {
	this := PowerPolicyInventory{}
	var classId string = "power.PolicyInventory"
	this.ClassId = classId
	var objectType string = "power.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PowerPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PowerPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PowerPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PowerPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocatedBudget returns the AllocatedBudget field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetAllocatedBudget() int64 {
	if o == nil || o.AllocatedBudget == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedBudget
}

// GetAllocatedBudgetOk returns a tuple with the AllocatedBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetAllocatedBudgetOk() (*int64, bool) {
	if o == nil || o.AllocatedBudget == nil {
		return nil, false
	}
	return o.AllocatedBudget, true
}

// HasAllocatedBudget returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasAllocatedBudget() bool {
	if o != nil && o.AllocatedBudget != nil {
		return true
	}

	return false
}

// SetAllocatedBudget gets a reference to the given int64 and assigns it to the AllocatedBudget field.
func (o *PowerPolicyInventory) SetAllocatedBudget(v int64) {
	o.AllocatedBudget = &v
}

// GetDynamicRebalancing returns the DynamicRebalancing field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetDynamicRebalancing() string {
	if o == nil || o.DynamicRebalancing == nil {
		var ret string
		return ret
	}
	return *o.DynamicRebalancing
}

// GetDynamicRebalancingOk returns a tuple with the DynamicRebalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetDynamicRebalancingOk() (*string, bool) {
	if o == nil || o.DynamicRebalancing == nil {
		return nil, false
	}
	return o.DynamicRebalancing, true
}

// HasDynamicRebalancing returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasDynamicRebalancing() bool {
	if o != nil && o.DynamicRebalancing != nil {
		return true
	}

	return false
}

// SetDynamicRebalancing gets a reference to the given string and assigns it to the DynamicRebalancing field.
func (o *PowerPolicyInventory) SetDynamicRebalancing(v string) {
	o.DynamicRebalancing = &v
}

// GetExtendedPowerCapacity returns the ExtendedPowerCapacity field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetExtendedPowerCapacity() string {
	if o == nil || o.ExtendedPowerCapacity == nil {
		var ret string
		return ret
	}
	return *o.ExtendedPowerCapacity
}

// GetExtendedPowerCapacityOk returns a tuple with the ExtendedPowerCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetExtendedPowerCapacityOk() (*string, bool) {
	if o == nil || o.ExtendedPowerCapacity == nil {
		return nil, false
	}
	return o.ExtendedPowerCapacity, true
}

// HasExtendedPowerCapacity returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasExtendedPowerCapacity() bool {
	if o != nil && o.ExtendedPowerCapacity != nil {
		return true
	}

	return false
}

// SetExtendedPowerCapacity gets a reference to the given string and assigns it to the ExtendedPowerCapacity field.
func (o *PowerPolicyInventory) SetExtendedPowerCapacity(v string) {
	o.ExtendedPowerCapacity = &v
}

// GetPowerPriority returns the PowerPriority field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetPowerPriority() string {
	if o == nil || o.PowerPriority == nil {
		var ret string
		return ret
	}
	return *o.PowerPriority
}

// GetPowerPriorityOk returns a tuple with the PowerPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetPowerPriorityOk() (*string, bool) {
	if o == nil || o.PowerPriority == nil {
		return nil, false
	}
	return o.PowerPriority, true
}

// HasPowerPriority returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasPowerPriority() bool {
	if o != nil && o.PowerPriority != nil {
		return true
	}

	return false
}

// SetPowerPriority gets a reference to the given string and assigns it to the PowerPriority field.
func (o *PowerPolicyInventory) SetPowerPriority(v string) {
	o.PowerPriority = &v
}

// GetPowerProfiling returns the PowerProfiling field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetPowerProfiling() string {
	if o == nil || o.PowerProfiling == nil {
		var ret string
		return ret
	}
	return *o.PowerProfiling
}

// GetPowerProfilingOk returns a tuple with the PowerProfiling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetPowerProfilingOk() (*string, bool) {
	if o == nil || o.PowerProfiling == nil {
		return nil, false
	}
	return o.PowerProfiling, true
}

// HasPowerProfiling returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasPowerProfiling() bool {
	if o != nil && o.PowerProfiling != nil {
		return true
	}

	return false
}

// SetPowerProfiling gets a reference to the given string and assigns it to the PowerProfiling field.
func (o *PowerPolicyInventory) SetPowerProfiling(v string) {
	o.PowerProfiling = &v
}

// GetPowerRestoreState returns the PowerRestoreState field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetPowerRestoreState() string {
	if o == nil || o.PowerRestoreState == nil {
		var ret string
		return ret
	}
	return *o.PowerRestoreState
}

// GetPowerRestoreStateOk returns a tuple with the PowerRestoreState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetPowerRestoreStateOk() (*string, bool) {
	if o == nil || o.PowerRestoreState == nil {
		return nil, false
	}
	return o.PowerRestoreState, true
}

// HasPowerRestoreState returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasPowerRestoreState() bool {
	if o != nil && o.PowerRestoreState != nil {
		return true
	}

	return false
}

// SetPowerRestoreState gets a reference to the given string and assigns it to the PowerRestoreState field.
func (o *PowerPolicyInventory) SetPowerRestoreState(v string) {
	o.PowerRestoreState = &v
}

// GetPowerSaveMode returns the PowerSaveMode field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetPowerSaveMode() string {
	if o == nil || o.PowerSaveMode == nil {
		var ret string
		return ret
	}
	return *o.PowerSaveMode
}

// GetPowerSaveModeOk returns a tuple with the PowerSaveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetPowerSaveModeOk() (*string, bool) {
	if o == nil || o.PowerSaveMode == nil {
		return nil, false
	}
	return o.PowerSaveMode, true
}

// HasPowerSaveMode returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasPowerSaveMode() bool {
	if o != nil && o.PowerSaveMode != nil {
		return true
	}

	return false
}

// SetPowerSaveMode gets a reference to the given string and assigns it to the PowerSaveMode field.
func (o *PowerPolicyInventory) SetPowerSaveMode(v string) {
	o.PowerSaveMode = &v
}

// GetRedundancyMode returns the RedundancyMode field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetRedundancyMode() string {
	if o == nil || o.RedundancyMode == nil {
		var ret string
		return ret
	}
	return *o.RedundancyMode
}

// GetRedundancyModeOk returns a tuple with the RedundancyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetRedundancyModeOk() (*string, bool) {
	if o == nil || o.RedundancyMode == nil {
		return nil, false
	}
	return o.RedundancyMode, true
}

// HasRedundancyMode returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasRedundancyMode() bool {
	if o != nil && o.RedundancyMode != nil {
		return true
	}

	return false
}

// SetRedundancyMode gets a reference to the given string and assigns it to the RedundancyMode field.
func (o *PowerPolicyInventory) SetRedundancyMode(v string) {
	o.RedundancyMode = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *PowerPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *PowerPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *PowerPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o PowerPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocatedBudget != nil {
		toSerialize["AllocatedBudget"] = o.AllocatedBudget
	}
	if o.DynamicRebalancing != nil {
		toSerialize["DynamicRebalancing"] = o.DynamicRebalancing
	}
	if o.ExtendedPowerCapacity != nil {
		toSerialize["ExtendedPowerCapacity"] = o.ExtendedPowerCapacity
	}
	if o.PowerPriority != nil {
		toSerialize["PowerPriority"] = o.PowerPriority
	}
	if o.PowerProfiling != nil {
		toSerialize["PowerProfiling"] = o.PowerProfiling
	}
	if o.PowerRestoreState != nil {
		toSerialize["PowerRestoreState"] = o.PowerRestoreState
	}
	if o.PowerSaveMode != nil {
		toSerialize["PowerSaveMode"] = o.PowerSaveMode
	}
	if o.RedundancyMode != nil {
		toSerialize["RedundancyMode"] = o.RedundancyMode
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PowerPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type PowerPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Sets the Allocated Power Budget of the System (in Watts). This field is only supported for Cisco UCS X series Chassis.
		AllocatedBudget *int64 `json:"AllocatedBudget,omitempty"`
		// Sets the Dynamic Power Rebalancing of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
		DynamicRebalancing *string `json:"DynamicRebalancing,omitempty"`
		// Sets the Extended Power Capacity of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
		ExtendedPowerCapacity *string `json:"ExtendedPowerCapacity,omitempty"`
		// Sets the Power Priority of the System. This field is only supported for Cisco UCS X series servers. * `Low` - Set the Power Priority to Low. * `Medium` - Set the Power Priority to Medium. * `High` - Set the Power Priority to High.
		PowerPriority *string `json:"PowerPriority,omitempty"`
		// Sets the Power Profiling of the Server. This field is only supported for Cisco UCS X series servers. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
		PowerProfiling *string `json:"PowerProfiling,omitempty"`
		// Sets the Power Restore State of the Server. This field is only supported for Cisco UCS X series servers. * `AlwaysOff` - Set the Power Restore Mode to Off. * `AlwaysOn` - Set the Power Restore Mode to On. * `LastState` - Set the Power Restore Mode to LastState.
		PowerRestoreState *string `json:"PowerRestoreState,omitempty"`
		// Sets the Power Save mode of the System. This option is only supported for Cisco UCS X series Chassis. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
		PowerSaveMode *string `json:"PowerSaveMode,omitempty"`
		// Sets the Power Redundancy of the System. N+2 mode is only supported for Cisco UCS X series Chassis. * `Grid` - Grid Mode requires two power sources. If one source fails, the surviving PSUs connected to the other source provides power to the chassis. * `NotRedundant` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements. No Redundant PSUs are maintained. * `N+1` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus one additional PSU for redundancy. * `N+2` - Power Manager turns on the minimum number of PSUs required to support chassis power requirements plus two additional PSU for redundancy. This Mode is only supported for UCS X series Chassis.
		RedundancyMode *string               `json:"RedundancyMode,omitempty"`
		TargetMo       *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varPowerPolicyInventoryWithoutEmbeddedStruct := PowerPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPowerPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varPowerPolicyInventory := _PowerPolicyInventory{}
		varPowerPolicyInventory.ClassId = varPowerPolicyInventoryWithoutEmbeddedStruct.ClassId
		varPowerPolicyInventory.ObjectType = varPowerPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varPowerPolicyInventory.AllocatedBudget = varPowerPolicyInventoryWithoutEmbeddedStruct.AllocatedBudget
		varPowerPolicyInventory.DynamicRebalancing = varPowerPolicyInventoryWithoutEmbeddedStruct.DynamicRebalancing
		varPowerPolicyInventory.ExtendedPowerCapacity = varPowerPolicyInventoryWithoutEmbeddedStruct.ExtendedPowerCapacity
		varPowerPolicyInventory.PowerPriority = varPowerPolicyInventoryWithoutEmbeddedStruct.PowerPriority
		varPowerPolicyInventory.PowerProfiling = varPowerPolicyInventoryWithoutEmbeddedStruct.PowerProfiling
		varPowerPolicyInventory.PowerRestoreState = varPowerPolicyInventoryWithoutEmbeddedStruct.PowerRestoreState
		varPowerPolicyInventory.PowerSaveMode = varPowerPolicyInventoryWithoutEmbeddedStruct.PowerSaveMode
		varPowerPolicyInventory.RedundancyMode = varPowerPolicyInventoryWithoutEmbeddedStruct.RedundancyMode
		varPowerPolicyInventory.TargetMo = varPowerPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = PowerPolicyInventory(varPowerPolicyInventory)
	} else {
		return err
	}

	varPowerPolicyInventory := _PowerPolicyInventory{}

	err = json.Unmarshal(bytes, &varPowerPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varPowerPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocatedBudget")
		delete(additionalProperties, "DynamicRebalancing")
		delete(additionalProperties, "ExtendedPowerCapacity")
		delete(additionalProperties, "PowerPriority")
		delete(additionalProperties, "PowerProfiling")
		delete(additionalProperties, "PowerRestoreState")
		delete(additionalProperties, "PowerSaveMode")
		delete(additionalProperties, "RedundancyMode")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullablePowerPolicyInventory struct {
	value *PowerPolicyInventory
	isSet bool
}

func (v NullablePowerPolicyInventory) Get() *PowerPolicyInventory {
	return v.value
}

func (v *NullablePowerPolicyInventory) Set(val *PowerPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerPolicyInventory(val *PowerPolicyInventory) *NullablePowerPolicyInventory {
	return &NullablePowerPolicyInventory{value: val, isSet: true}
}

func (v NullablePowerPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
