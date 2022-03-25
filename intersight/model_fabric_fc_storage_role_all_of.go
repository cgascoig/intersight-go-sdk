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
)

// FabricFcStorageRoleAllOf Definition of the list of properties defined in 'fabric.FcStorageRole', excluding properties defined in parent classes.
type FabricFcStorageRoleAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `8Gbps` - Admin configurable speed 8Gbps. * `16Gbps` - Admin configurable speed 16Gbps. * `32Gbps` - Admin configurable speed 32Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Virtual San Identifier associated to the FC port.
	VsanId               *int64 `json:"VsanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcStorageRoleAllOf FabricFcStorageRoleAllOf

// NewFabricFcStorageRoleAllOf instantiates a new FabricFcStorageRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcStorageRoleAllOf(classId string, objectType string) *FabricFcStorageRoleAllOf {
	this := FabricFcStorageRoleAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	return &this
}

// NewFabricFcStorageRoleAllOfWithDefaults instantiates a new FabricFcStorageRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcStorageRoleAllOfWithDefaults() *FabricFcStorageRoleAllOf {
	this := FabricFcStorageRoleAllOf{}
	var classId string = "fabric.FcStorageRole"
	this.ClassId = classId
	var objectType string = "fabric.FcStorageRole"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcStorageRoleAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRoleAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcStorageRoleAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcStorageRoleAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRoleAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcStorageRoleAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcStorageRoleAllOf) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRoleAllOf) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcStorageRoleAllOf) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcStorageRoleAllOf) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricFcStorageRoleAllOf) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcStorageRoleAllOf) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricFcStorageRoleAllOf) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricFcStorageRoleAllOf) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o FabricFcStorageRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcStorageRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFabricFcStorageRoleAllOf := _FabricFcStorageRoleAllOf{}

	if err = json.Unmarshal(bytes, &varFabricFcStorageRoleAllOf); err == nil {
		*o = FabricFcStorageRoleAllOf(varFabricFcStorageRoleAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "VsanId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricFcStorageRoleAllOf struct {
	value *FabricFcStorageRoleAllOf
	isSet bool
}

func (v NullableFabricFcStorageRoleAllOf) Get() *FabricFcStorageRoleAllOf {
	return v.value
}

func (v *NullableFabricFcStorageRoleAllOf) Set(val *FabricFcStorageRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcStorageRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcStorageRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcStorageRoleAllOf(val *FabricFcStorageRoleAllOf) *NullableFabricFcStorageRoleAllOf {
	return &NullableFabricFcStorageRoleAllOf{value: val, isSet: true}
}

func (v NullableFabricFcStorageRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcStorageRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
