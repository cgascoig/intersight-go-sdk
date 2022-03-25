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

// CloudBasePlacement A base placement definition extended by all cloud specific placement objects.
type CloudBasePlacement struct {
	VirtualizationBasePlacement
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType  string                   `json:"ObjectType"`
	BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
	// The internally generated identity of this placement. This entity is not manipulated by users. It aids in uniquely identifying the placement object.
	Identity             *string                       `json:"Identity,omitempty"`
	RegionInfo           NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
	ZoneInfo             NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBasePlacement CloudBasePlacement

// NewCloudBasePlacement instantiates a new CloudBasePlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBasePlacement(classId string, objectType string) *CloudBasePlacement {
	this := CloudBasePlacement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudBasePlacementWithDefaults instantiates a new CloudBasePlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBasePlacementWithDefaults() *CloudBasePlacement {
	this := CloudBasePlacement{}
	var classId string = "cloud.AwsVpc"
	this.ClassId = classId
	var objectType string = "cloud.AwsVpc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBasePlacement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBasePlacement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBasePlacement) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBasePlacement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBasePlacement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBasePlacement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingUnit returns the BillingUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBasePlacement) GetBillingUnit() CloudBillingUnit {
	if o == nil || o.BillingUnit.Get() == nil {
		var ret CloudBillingUnit
		return ret
	}
	return *o.BillingUnit.Get()
}

// GetBillingUnitOk returns a tuple with the BillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBasePlacement) GetBillingUnitOk() (*CloudBillingUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnit.Get(), o.BillingUnit.IsSet()
}

// HasBillingUnit returns a boolean if a field has been set.
func (o *CloudBasePlacement) HasBillingUnit() bool {
	if o != nil && o.BillingUnit.IsSet() {
		return true
	}

	return false
}

// SetBillingUnit gets a reference to the given NullableCloudBillingUnit and assigns it to the BillingUnit field.
func (o *CloudBasePlacement) SetBillingUnit(v CloudBillingUnit) {
	o.BillingUnit.Set(&v)
}

// SetBillingUnitNil sets the value for BillingUnit to be an explicit nil
func (o *CloudBasePlacement) SetBillingUnitNil() {
	o.BillingUnit.Set(nil)
}

// UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
func (o *CloudBasePlacement) UnsetBillingUnit() {
	o.BillingUnit.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CloudBasePlacement) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBasePlacement) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CloudBasePlacement) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *CloudBasePlacement) SetIdentity(v string) {
	o.Identity = &v
}

// GetRegionInfo returns the RegionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBasePlacement) GetRegionInfo() CloudCloudRegion {
	if o == nil || o.RegionInfo.Get() == nil {
		var ret CloudCloudRegion
		return ret
	}
	return *o.RegionInfo.Get()
}

// GetRegionInfoOk returns a tuple with the RegionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBasePlacement) GetRegionInfoOk() (*CloudCloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegionInfo.Get(), o.RegionInfo.IsSet()
}

// HasRegionInfo returns a boolean if a field has been set.
func (o *CloudBasePlacement) HasRegionInfo() bool {
	if o != nil && o.RegionInfo.IsSet() {
		return true
	}

	return false
}

// SetRegionInfo gets a reference to the given NullableCloudCloudRegion and assigns it to the RegionInfo field.
func (o *CloudBasePlacement) SetRegionInfo(v CloudCloudRegion) {
	o.RegionInfo.Set(&v)
}

// SetRegionInfoNil sets the value for RegionInfo to be an explicit nil
func (o *CloudBasePlacement) SetRegionInfoNil() {
	o.RegionInfo.Set(nil)
}

// UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
func (o *CloudBasePlacement) UnsetRegionInfo() {
	o.RegionInfo.Unset()
}

// GetZoneInfo returns the ZoneInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBasePlacement) GetZoneInfo() CloudAvailabilityZone {
	if o == nil || o.ZoneInfo.Get() == nil {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.ZoneInfo.Get()
}

// GetZoneInfoOk returns a tuple with the ZoneInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBasePlacement) GetZoneInfoOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneInfo.Get(), o.ZoneInfo.IsSet()
}

// HasZoneInfo returns a boolean if a field has been set.
func (o *CloudBasePlacement) HasZoneInfo() bool {
	if o != nil && o.ZoneInfo.IsSet() {
		return true
	}

	return false
}

// SetZoneInfo gets a reference to the given NullableCloudAvailabilityZone and assigns it to the ZoneInfo field.
func (o *CloudBasePlacement) SetZoneInfo(v CloudAvailabilityZone) {
	o.ZoneInfo.Set(&v)
}

// SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil
func (o *CloudBasePlacement) SetZoneInfoNil() {
	o.ZoneInfo.Set(nil)
}

// UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil
func (o *CloudBasePlacement) UnsetZoneInfo() {
	o.ZoneInfo.Unset()
}

func (o CloudBasePlacement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBasePlacement, errVirtualizationBasePlacement := json.Marshal(o.VirtualizationBasePlacement)
	if errVirtualizationBasePlacement != nil {
		return []byte{}, errVirtualizationBasePlacement
	}
	errVirtualizationBasePlacement = json.Unmarshal([]byte(serializedVirtualizationBasePlacement), &toSerialize)
	if errVirtualizationBasePlacement != nil {
		return []byte{}, errVirtualizationBasePlacement
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingUnit.IsSet() {
		toSerialize["BillingUnit"] = o.BillingUnit.Get()
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.RegionInfo.IsSet() {
		toSerialize["RegionInfo"] = o.RegionInfo.Get()
	}
	if o.ZoneInfo.IsSet() {
		toSerialize["ZoneInfo"] = o.ZoneInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBasePlacement) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBasePlacementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType  string                   `json:"ObjectType"`
		BillingUnit NullableCloudBillingUnit `json:"BillingUnit,omitempty"`
		// The internally generated identity of this placement. This entity is not manipulated by users. It aids in uniquely identifying the placement object.
		Identity   *string                       `json:"Identity,omitempty"`
		RegionInfo NullableCloudCloudRegion      `json:"RegionInfo,omitempty"`
		ZoneInfo   NullableCloudAvailabilityZone `json:"ZoneInfo,omitempty"`
	}

	varCloudBasePlacementWithoutEmbeddedStruct := CloudBasePlacementWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBasePlacementWithoutEmbeddedStruct)
	if err == nil {
		varCloudBasePlacement := _CloudBasePlacement{}
		varCloudBasePlacement.ClassId = varCloudBasePlacementWithoutEmbeddedStruct.ClassId
		varCloudBasePlacement.ObjectType = varCloudBasePlacementWithoutEmbeddedStruct.ObjectType
		varCloudBasePlacement.BillingUnit = varCloudBasePlacementWithoutEmbeddedStruct.BillingUnit
		varCloudBasePlacement.Identity = varCloudBasePlacementWithoutEmbeddedStruct.Identity
		varCloudBasePlacement.RegionInfo = varCloudBasePlacementWithoutEmbeddedStruct.RegionInfo
		varCloudBasePlacement.ZoneInfo = varCloudBasePlacementWithoutEmbeddedStruct.ZoneInfo
		*o = CloudBasePlacement(varCloudBasePlacement)
	} else {
		return err
	}

	varCloudBasePlacement := _CloudBasePlacement{}

	err = json.Unmarshal(bytes, &varCloudBasePlacement)
	if err == nil {
		o.VirtualizationBasePlacement = varCloudBasePlacement.VirtualizationBasePlacement
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnit")
		delete(additionalProperties, "Identity")
		delete(additionalProperties, "RegionInfo")
		delete(additionalProperties, "ZoneInfo")

		// remove fields from embedded structs
		reflectVirtualizationBasePlacement := reflect.ValueOf(o.VirtualizationBasePlacement)
		for i := 0; i < reflectVirtualizationBasePlacement.Type().NumField(); i++ {
			t := reflectVirtualizationBasePlacement.Type().Field(i)

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

type NullableCloudBasePlacement struct {
	value *CloudBasePlacement
	isSet bool
}

func (v NullableCloudBasePlacement) Get() *CloudBasePlacement {
	return v.value
}

func (v *NullableCloudBasePlacement) Set(val *CloudBasePlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBasePlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBasePlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBasePlacement(val *CloudBasePlacement) *NullableCloudBasePlacement {
	return &NullableCloudBasePlacement{value: val, isSet: true}
}

func (v NullableCloudBasePlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBasePlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
