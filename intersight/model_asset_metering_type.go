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

// AssetMeteringType Type for saving Metering details.
type AssetMeteringType struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Metric type used to calculate metering for the device sent from the IB Contract. example  Node, vMemory, vCPU. * `None` - A default value to catch cases where metric type is not correctly detected. * `Node` - The metering of the device is on the basis of Power state. * `Storage` - The metering of the device is on the basis of used Storage. * `vMemory` - The metering of the device is on the basis of VM Memory. * `vCPU` - The metering of the device is on the basis of vCPU. * `vStorage` - The metering of the device is on the basis of used virtual Storage. * `Switch` - The metering of the device is on the basis of Switch.
	Name *string `json:"Name,omitempty"`
	// Metric unit used to calculate metering for the device sent from the IB Contract. example  Node, GiB, Cores. * `None` - A default value to catch cases where metric unit is not correctly detected. * `Node` - It is applicable for Node Metric type. * `GiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `TiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `Cores` - It is applicable for vCPU Metric type. * `Switch` - It is applicable for Switch Metric type. * `Port` - It is applicable for Switch Metric type.
	Unit                 *string `json:"Unit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetMeteringType AssetMeteringType

// NewAssetMeteringType instantiates a new AssetMeteringType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetMeteringType(classId string, objectType string) *AssetMeteringType {
	this := AssetMeteringType{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetMeteringTypeWithDefaults instantiates a new AssetMeteringType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetMeteringTypeWithDefaults() *AssetMeteringType {
	this := AssetMeteringType{}
	var classId string = "asset.MeteringType"
	this.ClassId = classId
	var objectType string = "asset.MeteringType"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetMeteringType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetMeteringType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetMeteringType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetMeteringType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetMeteringType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetMeteringType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetMeteringType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMeteringType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetMeteringType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetMeteringType) SetName(v string) {
	o.Name = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AssetMeteringType) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMeteringType) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AssetMeteringType) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AssetMeteringType) SetUnit(v string) {
	o.Unit = &v
}

func (o AssetMeteringType) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetMeteringType) UnmarshalJSON(bytes []byte) (err error) {
	type AssetMeteringTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Metric type used to calculate metering for the device sent from the IB Contract. example  Node, vMemory, vCPU. * `None` - A default value to catch cases where metric type is not correctly detected. * `Node` - The metering of the device is on the basis of Power state. * `Storage` - The metering of the device is on the basis of used Storage. * `vMemory` - The metering of the device is on the basis of VM Memory. * `vCPU` - The metering of the device is on the basis of vCPU. * `vStorage` - The metering of the device is on the basis of used virtual Storage. * `Switch` - The metering of the device is on the basis of Switch.
		Name *string `json:"Name,omitempty"`
		// Metric unit used to calculate metering for the device sent from the IB Contract. example  Node, GiB, Cores. * `None` - A default value to catch cases where metric unit is not correctly detected. * `Node` - It is applicable for Node Metric type. * `GiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `TiB` - It is applicable for VMemory, vStorage and Storage Metric types. * `Cores` - It is applicable for vCPU Metric type. * `Switch` - It is applicable for Switch Metric type. * `Port` - It is applicable for Switch Metric type.
		Unit *string `json:"Unit,omitempty"`
	}

	varAssetMeteringTypeWithoutEmbeddedStruct := AssetMeteringTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetMeteringTypeWithoutEmbeddedStruct)
	if err == nil {
		varAssetMeteringType := _AssetMeteringType{}
		varAssetMeteringType.ClassId = varAssetMeteringTypeWithoutEmbeddedStruct.ClassId
		varAssetMeteringType.ObjectType = varAssetMeteringTypeWithoutEmbeddedStruct.ObjectType
		varAssetMeteringType.Name = varAssetMeteringTypeWithoutEmbeddedStruct.Name
		varAssetMeteringType.Unit = varAssetMeteringTypeWithoutEmbeddedStruct.Unit
		*o = AssetMeteringType(varAssetMeteringType)
	} else {
		return err
	}

	varAssetMeteringType := _AssetMeteringType{}

	err = json.Unmarshal(bytes, &varAssetMeteringType)
	if err == nil {
		o.MoBaseComplexType = varAssetMeteringType.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Unit")

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

type NullableAssetMeteringType struct {
	value *AssetMeteringType
	isSet bool
}

func (v NullableAssetMeteringType) Get() *AssetMeteringType {
	return v.value
}

func (v *NullableAssetMeteringType) Set(val *AssetMeteringType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetMeteringType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetMeteringType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetMeteringType(val *AssetMeteringType) *NullableAssetMeteringType {
	return &NullableAssetMeteringType{value: val, isSet: true}
}

func (v NullableAssetMeteringType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetMeteringType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
