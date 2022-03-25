/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// AssetProductInformation Type for saving the product information.
type AssetProductInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                          `json:"ObjectType"`
	BillTo     NullableAssetAddressInformation `json:"BillTo,omitempty"`
	// Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".
	Description *string `json:"Description,omitempty"`
	// Family that the product belongs to. Example \"UCSB\".
	Family *string `json:"Family,omitempty"`
	// Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".
	Group *string `json:"Group,omitempty"`
	// Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".
	Number *string                         `json:"Number,omitempty"`
	ShipTo NullableAssetAddressInformation `json:"ShipTo,omitempty"`
	// Sub type of the product being specified. example \"UCS 6200 SER\".
	SubType              *string `json:"SubType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetProductInformation AssetProductInformation

// NewAssetProductInformation instantiates a new AssetProductInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetProductInformation(classId string, objectType string) *AssetProductInformation {
	this := AssetProductInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetProductInformationWithDefaults instantiates a new AssetProductInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetProductInformationWithDefaults() *AssetProductInformation {
	this := AssetProductInformation{}
	var classId string = "asset.ProductInformation"
	this.ClassId = classId
	var objectType string = "asset.ProductInformation"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetProductInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetProductInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetProductInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetProductInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillTo returns the BillTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetProductInformation) GetBillTo() AssetAddressInformation {
	if o == nil || o.BillTo.Get() == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.BillTo.Get()
}

// GetBillToOk returns a tuple with the BillTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetProductInformation) GetBillToOk() (*AssetAddressInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillTo.Get(), o.BillTo.IsSet()
}

// HasBillTo returns a boolean if a field has been set.
func (o *AssetProductInformation) HasBillTo() bool {
	if o != nil && o.BillTo.IsSet() {
		return true
	}

	return false
}

// SetBillTo gets a reference to the given NullableAssetAddressInformation and assigns it to the BillTo field.
func (o *AssetProductInformation) SetBillTo(v AssetAddressInformation) {
	o.BillTo.Set(&v)
}

// SetBillToNil sets the value for BillTo to be an explicit nil
func (o *AssetProductInformation) SetBillToNil() {
	o.BillTo.Set(nil)
}

// UnsetBillTo ensures that no value is present for BillTo, not even an explicit nil
func (o *AssetProductInformation) UnsetBillTo() {
	o.BillTo.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetProductInformation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetProductInformation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetProductInformation) SetDescription(v string) {
	o.Description = &v
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AssetProductInformation) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AssetProductInformation) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AssetProductInformation) SetFamily(v string) {
	o.Family = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AssetProductInformation) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AssetProductInformation) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *AssetProductInformation) SetGroup(v string) {
	o.Group = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *AssetProductInformation) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *AssetProductInformation) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *AssetProductInformation) SetNumber(v string) {
	o.Number = &v
}

// GetShipTo returns the ShipTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetProductInformation) GetShipTo() AssetAddressInformation {
	if o == nil || o.ShipTo.Get() == nil {
		var ret AssetAddressInformation
		return ret
	}
	return *o.ShipTo.Get()
}

// GetShipToOk returns a tuple with the ShipTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetProductInformation) GetShipToOk() (*AssetAddressInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipTo.Get(), o.ShipTo.IsSet()
}

// HasShipTo returns a boolean if a field has been set.
func (o *AssetProductInformation) HasShipTo() bool {
	if o != nil && o.ShipTo.IsSet() {
		return true
	}

	return false
}

// SetShipTo gets a reference to the given NullableAssetAddressInformation and assigns it to the ShipTo field.
func (o *AssetProductInformation) SetShipTo(v AssetAddressInformation) {
	o.ShipTo.Set(&v)
}

// SetShipToNil sets the value for ShipTo to be an explicit nil
func (o *AssetProductInformation) SetShipToNil() {
	o.ShipTo.Set(nil)
}

// UnsetShipTo ensures that no value is present for ShipTo, not even an explicit nil
func (o *AssetProductInformation) UnsetShipTo() {
	o.ShipTo.Unset()
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *AssetProductInformation) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProductInformation) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *AssetProductInformation) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *AssetProductInformation) SetSubType(v string) {
	o.SubType = &v
}

func (o AssetProductInformation) MarshalJSON() ([]byte, error) {
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
	if o.BillTo.IsSet() {
		toSerialize["BillTo"] = o.BillTo.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Family != nil {
		toSerialize["Family"] = o.Family
	}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.Number != nil {
		toSerialize["Number"] = o.Number
	}
	if o.ShipTo.IsSet() {
		toSerialize["ShipTo"] = o.ShipTo.Get()
	}
	if o.SubType != nil {
		toSerialize["SubType"] = o.SubType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetProductInformation) UnmarshalJSON(bytes []byte) (err error) {
	type AssetProductInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                          `json:"ObjectType"`
		BillTo     NullableAssetAddressInformation `json:"BillTo,omitempty"`
		// Short description of the Cisco product that helps identify the product easily. example \"DISTI:UCS 6248UP 1RU Fabric Int/No PSU/32 UP/ 12p LIC\".
		Description *string `json:"Description,omitempty"`
		// Family that the product belongs to. Example \"UCSB\".
		Family *string `json:"Family,omitempty"`
		// Group that the product belongs to. It is one higher level categorization above family. example \"Switch\".
		Group *string `json:"Group,omitempty"`
		// Product number that identifies the product. example PID. example \"UCS-FI-6248UP-CH2\".
		Number *string                         `json:"Number,omitempty"`
		ShipTo NullableAssetAddressInformation `json:"ShipTo,omitempty"`
		// Sub type of the product being specified. example \"UCS 6200 SER\".
		SubType *string `json:"SubType,omitempty"`
	}

	varAssetProductInformationWithoutEmbeddedStruct := AssetProductInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetProductInformationWithoutEmbeddedStruct)
	if err == nil {
		varAssetProductInformation := _AssetProductInformation{}
		varAssetProductInformation.ClassId = varAssetProductInformationWithoutEmbeddedStruct.ClassId
		varAssetProductInformation.ObjectType = varAssetProductInformationWithoutEmbeddedStruct.ObjectType
		varAssetProductInformation.BillTo = varAssetProductInformationWithoutEmbeddedStruct.BillTo
		varAssetProductInformation.Description = varAssetProductInformationWithoutEmbeddedStruct.Description
		varAssetProductInformation.Family = varAssetProductInformationWithoutEmbeddedStruct.Family
		varAssetProductInformation.Group = varAssetProductInformationWithoutEmbeddedStruct.Group
		varAssetProductInformation.Number = varAssetProductInformationWithoutEmbeddedStruct.Number
		varAssetProductInformation.ShipTo = varAssetProductInformationWithoutEmbeddedStruct.ShipTo
		varAssetProductInformation.SubType = varAssetProductInformationWithoutEmbeddedStruct.SubType
		*o = AssetProductInformation(varAssetProductInformation)
	} else {
		return err
	}

	varAssetProductInformation := _AssetProductInformation{}

	err = json.Unmarshal(bytes, &varAssetProductInformation)
	if err == nil {
		o.MoBaseComplexType = varAssetProductInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillTo")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Family")
		delete(additionalProperties, "Group")
		delete(additionalProperties, "Number")
		delete(additionalProperties, "ShipTo")
		delete(additionalProperties, "SubType")

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

type NullableAssetProductInformation struct {
	value *AssetProductInformation
	isSet bool
}

func (v NullableAssetProductInformation) Get() *AssetProductInformation {
	return v.value
}

func (v *NullableAssetProductInformation) Set(val *AssetProductInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetProductInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetProductInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetProductInformation(val *AssetProductInformation) *NullableAssetProductInformation {
	return &NullableAssetProductInformation{value: val, isSet: true}
}

func (v NullableAssetProductInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetProductInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
