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

// EquipmentDeviceSummary Aggregation of properties pertaining to different inventory MOs.
type EquipmentDeviceSummary struct {
	ViewsView
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The distinguished name for the Network Element.
	Dn *string `json:"Dn,omitempty"`
	// The model information of the Network Element.
	Model *string `json:"Model,omitempty"`
	// The serial number for the Network Element.
	Serial *string `json:"Serial,omitempty"`
	// The source object type of this view MO.
	SourceObjectType     *string                              `json:"SourceObjectType,omitempty"`
	ComputeRackUnit      *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentDeviceSummary EquipmentDeviceSummary

// NewEquipmentDeviceSummary instantiates a new EquipmentDeviceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentDeviceSummary(classId string, objectType string) *EquipmentDeviceSummary {
	this := EquipmentDeviceSummary{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentDeviceSummaryWithDefaults instantiates a new EquipmentDeviceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentDeviceSummaryWithDefaults() *EquipmentDeviceSummary {
	this := EquipmentDeviceSummary{}
	var classId string = "equipment.DeviceSummary"
	this.ClassId = classId
	var objectType string = "equipment.DeviceSummary"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentDeviceSummary) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentDeviceSummary) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentDeviceSummary) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentDeviceSummary) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *EquipmentDeviceSummary) SetDn(v string) {
	o.Dn = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentDeviceSummary) SetModel(v string) {
	o.Model = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentDeviceSummary) SetSerial(v string) {
	o.Serial = &v
}

// GetSourceObjectType returns the SourceObjectType field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetSourceObjectType() string {
	if o == nil || o.SourceObjectType == nil {
		var ret string
		return ret
	}
	return *o.SourceObjectType
}

// GetSourceObjectTypeOk returns a tuple with the SourceObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetSourceObjectTypeOk() (*string, bool) {
	if o == nil || o.SourceObjectType == nil {
		return nil, false
	}
	return o.SourceObjectType, true
}

// HasSourceObjectType returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasSourceObjectType() bool {
	if o != nil && o.SourceObjectType != nil {
		return true
	}

	return false
}

// SetSourceObjectType gets a reference to the given string and assigns it to the SourceObjectType field.
func (o *EquipmentDeviceSummary) SetSourceObjectType(v string) {
	o.SourceObjectType = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentDeviceSummary) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentDeviceSummary) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentDeviceSummary) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentDeviceSummary) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeviceSummary) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentDeviceSummary) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentDeviceSummary) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentDeviceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedViewsView, errViewsView := json.Marshal(o.ViewsView)
	if errViewsView != nil {
		return []byte{}, errViewsView
	}
	errViewsView = json.Unmarshal([]byte(serializedViewsView), &toSerialize)
	if errViewsView != nil {
		return []byte{}, errViewsView
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SourceObjectType != nil {
		toSerialize["SourceObjectType"] = o.SourceObjectType
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentDeviceSummary) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentDeviceSummaryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The distinguished name for the Network Element.
		Dn *string `json:"Dn,omitempty"`
		// The model information of the Network Element.
		Model *string `json:"Model,omitempty"`
		// The serial number for the Network Element.
		Serial *string `json:"Serial,omitempty"`
		// The source object type of this view MO.
		SourceObjectType    *string                              `json:"SourceObjectType,omitempty"`
		ComputeRackUnit     *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		EquipmentChassis    *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentDeviceSummaryWithoutEmbeddedStruct := EquipmentDeviceSummaryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentDeviceSummaryWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentDeviceSummary := _EquipmentDeviceSummary{}
		varEquipmentDeviceSummary.ClassId = varEquipmentDeviceSummaryWithoutEmbeddedStruct.ClassId
		varEquipmentDeviceSummary.ObjectType = varEquipmentDeviceSummaryWithoutEmbeddedStruct.ObjectType
		varEquipmentDeviceSummary.Dn = varEquipmentDeviceSummaryWithoutEmbeddedStruct.Dn
		varEquipmentDeviceSummary.Model = varEquipmentDeviceSummaryWithoutEmbeddedStruct.Model
		varEquipmentDeviceSummary.Serial = varEquipmentDeviceSummaryWithoutEmbeddedStruct.Serial
		varEquipmentDeviceSummary.SourceObjectType = varEquipmentDeviceSummaryWithoutEmbeddedStruct.SourceObjectType
		varEquipmentDeviceSummary.ComputeRackUnit = varEquipmentDeviceSummaryWithoutEmbeddedStruct.ComputeRackUnit
		varEquipmentDeviceSummary.EquipmentChassis = varEquipmentDeviceSummaryWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentDeviceSummary.InventoryDeviceInfo = varEquipmentDeviceSummaryWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentDeviceSummary.RegisteredDevice = varEquipmentDeviceSummaryWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentDeviceSummary(varEquipmentDeviceSummary)
	} else {
		return err
	}

	varEquipmentDeviceSummary := _EquipmentDeviceSummary{}

	err = json.Unmarshal(bytes, &varEquipmentDeviceSummary)
	if err == nil {
		o.ViewsView = varEquipmentDeviceSummary.ViewsView
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "SourceObjectType")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectViewsView := reflect.ValueOf(o.ViewsView)
		for i := 0; i < reflectViewsView.Type().NumField(); i++ {
			t := reflectViewsView.Type().Field(i)

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

type NullableEquipmentDeviceSummary struct {
	value *EquipmentDeviceSummary
	isSet bool
}

func (v NullableEquipmentDeviceSummary) Get() *EquipmentDeviceSummary {
	return v.value
}

func (v *NullableEquipmentDeviceSummary) Set(val *EquipmentDeviceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentDeviceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentDeviceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentDeviceSummary(val *EquipmentDeviceSummary) *NullableEquipmentDeviceSummary {
	return &NullableEquipmentDeviceSummary{value: val, isSet: true}
}

func (v NullableEquipmentDeviceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentDeviceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
