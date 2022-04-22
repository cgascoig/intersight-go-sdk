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
)

// HyperflexAppCatalogAllOf Definition of the list of properties defined in 'hyperflex.AppCatalog', excluding properties defined in parent classes.
type HyperflexAppCatalogAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The catalog version used in HyperFlex cluster configuration service.
	Version              *string                                    `json:"Version,omitempty"`
	FeatureLimitExternal *HyperflexFeatureLimitExternalRelationship `json:"FeatureLimitExternal,omitempty"`
	FeatureLimitInternal *HyperflexFeatureLimitInternalRelationship `json:"FeatureLimitInternal,omitempty"`
	// An array of relationships to hyperflexHxdpVersion resources.
	HxdpVersions []HyperflexHxdpVersionRelationship `json:"HxdpVersions,omitempty"`
	// An array of relationships to hyperflexCapabilityInfo resources.
	HyperflexCapabilityInfos []HyperflexCapabilityInfoRelationship `json:"HyperflexCapabilityInfos,omitempty"`
	// An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources.
	HyperflexSoftwareCompatibilityInfos []HclHyperflexSoftwareCompatibilityInfoRelationship `json:"HyperflexSoftwareCompatibilityInfos,omitempty"`
	ServerFirmwareVersion               *HyperflexServerFirmwareVersionRelationship         `json:"ServerFirmwareVersion,omitempty"`
	ServerModel                         *HyperflexServerModelRelationship                   `json:"ServerModel,omitempty"`
	// An array of relationships to hyperflexSoftwareDistributionEntry resources.
	SoftwareDistributions []HyperflexSoftwareDistributionEntryRelationship `json:"SoftwareDistributions,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _HyperflexAppCatalogAllOf HyperflexAppCatalogAllOf

// NewHyperflexAppCatalogAllOf instantiates a new HyperflexAppCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAppCatalogAllOf(classId string, objectType string) *HyperflexAppCatalogAllOf {
	this := HyperflexAppCatalogAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexAppCatalogAllOfWithDefaults instantiates a new HyperflexAppCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAppCatalogAllOfWithDefaults() *HyperflexAppCatalogAllOf {
	this := HyperflexAppCatalogAllOf{}
	var classId string = "hyperflex.AppCatalog"
	this.ClassId = classId
	var objectType string = "hyperflex.AppCatalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAppCatalogAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAppCatalogAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAppCatalogAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAppCatalogAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexAppCatalogAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetFeatureLimitExternal returns the FeatureLimitExternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship {
	if o == nil || o.FeatureLimitExternal == nil {
		var ret HyperflexFeatureLimitExternalRelationship
		return ret
	}
	return *o.FeatureLimitExternal
}

// GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool) {
	if o == nil || o.FeatureLimitExternal == nil {
		return nil, false
	}
	return o.FeatureLimitExternal, true
}

// HasFeatureLimitExternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasFeatureLimitExternal() bool {
	if o != nil && o.FeatureLimitExternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitExternal gets a reference to the given HyperflexFeatureLimitExternalRelationship and assigns it to the FeatureLimitExternal field.
func (o *HyperflexAppCatalogAllOf) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship) {
	o.FeatureLimitExternal = &v
}

// GetFeatureLimitInternal returns the FeatureLimitInternal field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship {
	if o == nil || o.FeatureLimitInternal == nil {
		var ret HyperflexFeatureLimitInternalRelationship
		return ret
	}
	return *o.FeatureLimitInternal
}

// GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool) {
	if o == nil || o.FeatureLimitInternal == nil {
		return nil, false
	}
	return o.FeatureLimitInternal, true
}

// HasFeatureLimitInternal returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasFeatureLimitInternal() bool {
	if o != nil && o.FeatureLimitInternal != nil {
		return true
	}

	return false
}

// SetFeatureLimitInternal gets a reference to the given HyperflexFeatureLimitInternalRelationship and assigns it to the FeatureLimitInternal field.
func (o *HyperflexAppCatalogAllOf) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship) {
	o.FeatureLimitInternal = &v
}

// GetHxdpVersions returns the HxdpVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalogAllOf) GetHxdpVersions() []HyperflexHxdpVersionRelationship {
	if o == nil {
		var ret []HyperflexHxdpVersionRelationship
		return ret
	}
	return o.HxdpVersions
}

// GetHxdpVersionsOk returns a tuple with the HxdpVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalogAllOf) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool) {
	if o == nil || o.HxdpVersions == nil {
		return nil, false
	}
	return &o.HxdpVersions, true
}

// HasHxdpVersions returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHxdpVersions() bool {
	if o != nil && o.HxdpVersions != nil {
		return true
	}

	return false
}

// SetHxdpVersions gets a reference to the given []HyperflexHxdpVersionRelationship and assigns it to the HxdpVersions field.
func (o *HyperflexAppCatalogAllOf) SetHxdpVersions(v []HyperflexHxdpVersionRelationship) {
	o.HxdpVersions = v
}

// GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship {
	if o == nil {
		var ret []HyperflexCapabilityInfoRelationship
		return ret
	}
	return o.HyperflexCapabilityInfos
}

// GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalogAllOf) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool) {
	if o == nil || o.HyperflexCapabilityInfos == nil {
		return nil, false
	}
	return &o.HyperflexCapabilityInfos, true
}

// HasHyperflexCapabilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHyperflexCapabilityInfos() bool {
	if o != nil && o.HyperflexCapabilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexCapabilityInfos gets a reference to the given []HyperflexCapabilityInfoRelationship and assigns it to the HyperflexCapabilityInfos field.
func (o *HyperflexAppCatalogAllOf) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship) {
	o.HyperflexCapabilityInfos = v
}

// GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship {
	if o == nil {
		var ret []HclHyperflexSoftwareCompatibilityInfoRelationship
		return ret
	}
	return o.HyperflexSoftwareCompatibilityInfos
}

// GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalogAllOf) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool) {
	if o == nil || o.HyperflexSoftwareCompatibilityInfos == nil {
		return nil, false
	}
	return &o.HyperflexSoftwareCompatibilityInfos, true
}

// HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasHyperflexSoftwareCompatibilityInfos() bool {
	if o != nil && o.HyperflexSoftwareCompatibilityInfos != nil {
		return true
	}

	return false
}

// SetHyperflexSoftwareCompatibilityInfos gets a reference to the given []HclHyperflexSoftwareCompatibilityInfoRelationship and assigns it to the HyperflexSoftwareCompatibilityInfos field.
func (o *HyperflexAppCatalogAllOf) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship) {
	o.HyperflexSoftwareCompatibilityInfos = v
}

// GetServerFirmwareVersion returns the ServerFirmwareVersion field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship {
	if o == nil || o.ServerFirmwareVersion == nil {
		var ret HyperflexServerFirmwareVersionRelationship
		return ret
	}
	return *o.ServerFirmwareVersion
}

// GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool) {
	if o == nil || o.ServerFirmwareVersion == nil {
		return nil, false
	}
	return o.ServerFirmwareVersion, true
}

// HasServerFirmwareVersion returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasServerFirmwareVersion() bool {
	if o != nil && o.ServerFirmwareVersion != nil {
		return true
	}

	return false
}

// SetServerFirmwareVersion gets a reference to the given HyperflexServerFirmwareVersionRelationship and assigns it to the ServerFirmwareVersion field.
func (o *HyperflexAppCatalogAllOf) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship) {
	o.ServerFirmwareVersion = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HyperflexAppCatalogAllOf) GetServerModel() HyperflexServerModelRelationship {
	if o == nil || o.ServerModel == nil {
		var ret HyperflexServerModelRelationship
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAppCatalogAllOf) GetServerModelOk() (*HyperflexServerModelRelationship, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given HyperflexServerModelRelationship and assigns it to the ServerModel field.
func (o *HyperflexAppCatalogAllOf) SetServerModel(v HyperflexServerModelRelationship) {
	o.ServerModel = &v
}

// GetSoftwareDistributions returns the SoftwareDistributions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAppCatalogAllOf) GetSoftwareDistributions() []HyperflexSoftwareDistributionEntryRelationship {
	if o == nil {
		var ret []HyperflexSoftwareDistributionEntryRelationship
		return ret
	}
	return o.SoftwareDistributions
}

// GetSoftwareDistributionsOk returns a tuple with the SoftwareDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAppCatalogAllOf) GetSoftwareDistributionsOk() (*[]HyperflexSoftwareDistributionEntryRelationship, bool) {
	if o == nil || o.SoftwareDistributions == nil {
		return nil, false
	}
	return &o.SoftwareDistributions, true
}

// HasSoftwareDistributions returns a boolean if a field has been set.
func (o *HyperflexAppCatalogAllOf) HasSoftwareDistributions() bool {
	if o != nil && o.SoftwareDistributions != nil {
		return true
	}

	return false
}

// SetSoftwareDistributions gets a reference to the given []HyperflexSoftwareDistributionEntryRelationship and assigns it to the SoftwareDistributions field.
func (o *HyperflexAppCatalogAllOf) SetSoftwareDistributions(v []HyperflexSoftwareDistributionEntryRelationship) {
	o.SoftwareDistributions = v
}

func (o HyperflexAppCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.FeatureLimitExternal != nil {
		toSerialize["FeatureLimitExternal"] = o.FeatureLimitExternal
	}
	if o.FeatureLimitInternal != nil {
		toSerialize["FeatureLimitInternal"] = o.FeatureLimitInternal
	}
	if o.HxdpVersions != nil {
		toSerialize["HxdpVersions"] = o.HxdpVersions
	}
	if o.HyperflexCapabilityInfos != nil {
		toSerialize["HyperflexCapabilityInfos"] = o.HyperflexCapabilityInfos
	}
	if o.HyperflexSoftwareCompatibilityInfos != nil {
		toSerialize["HyperflexSoftwareCompatibilityInfos"] = o.HyperflexSoftwareCompatibilityInfos
	}
	if o.ServerFirmwareVersion != nil {
		toSerialize["ServerFirmwareVersion"] = o.ServerFirmwareVersion
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	if o.SoftwareDistributions != nil {
		toSerialize["SoftwareDistributions"] = o.SoftwareDistributions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAppCatalogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAppCatalogAllOf := _HyperflexAppCatalogAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAppCatalogAllOf); err == nil {
		*o = HyperflexAppCatalogAllOf(varHyperflexAppCatalogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "FeatureLimitExternal")
		delete(additionalProperties, "FeatureLimitInternal")
		delete(additionalProperties, "HxdpVersions")
		delete(additionalProperties, "HyperflexCapabilityInfos")
		delete(additionalProperties, "HyperflexSoftwareCompatibilityInfos")
		delete(additionalProperties, "ServerFirmwareVersion")
		delete(additionalProperties, "ServerModel")
		delete(additionalProperties, "SoftwareDistributions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAppCatalogAllOf struct {
	value *HyperflexAppCatalogAllOf
	isSet bool
}

func (v NullableHyperflexAppCatalogAllOf) Get() *HyperflexAppCatalogAllOf {
	return v.value
}

func (v *NullableHyperflexAppCatalogAllOf) Set(val *HyperflexAppCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAppCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAppCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAppCatalogAllOf(val *HyperflexAppCatalogAllOf) *NullableHyperflexAppCatalogAllOf {
	return &NullableHyperflexAppCatalogAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAppCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAppCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
