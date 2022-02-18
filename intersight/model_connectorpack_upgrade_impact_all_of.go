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
)

// ConnectorpackUpgradeImpactAllOf Definition of the list of properties defined in 'connectorpack.UpgradeImpact', excluding properties defined in parent classes.
type ConnectorpackUpgradeImpactAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                             `json:"ObjectType"`
	ConnectorPack []ConnectorpackConnectorPackUpdate `json:"ConnectorPack,omitempty"`
	// States whether the UCS Director is eligible for an upgrade. Set to true if connector packs are available for upgrade, else set to false.
	IsEligibleForUpgrade *bool `json:"IsEligibleForUpgrade,omitempty"`
	// States whether all the requisite updates have been downloaded to the target UCS Director. Set to true if all connector packs required to upgrade UCS Director to the next iteration have been downloaded, else set to false.
	IsUpdateDownloaded   *bool                     `json:"IsUpdateDownloaded,omitempty"`
	UcsdInfo             *IaasUcsdInfoRelationship `json:"UcsdInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorpackUpgradeImpactAllOf ConnectorpackUpgradeImpactAllOf

// NewConnectorpackUpgradeImpactAllOf instantiates a new ConnectorpackUpgradeImpactAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackUpgradeImpactAllOf(classId string, objectType string) *ConnectorpackUpgradeImpactAllOf {
	this := ConnectorpackUpgradeImpactAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorpackUpgradeImpactAllOfWithDefaults instantiates a new ConnectorpackUpgradeImpactAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackUpgradeImpactAllOfWithDefaults() *ConnectorpackUpgradeImpactAllOf {
	this := ConnectorpackUpgradeImpactAllOf{}
	var classId string = "connectorpack.UpgradeImpact"
	this.ClassId = classId
	var objectType string = "connectorpack.UpgradeImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorpackUpgradeImpactAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpactAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorpackUpgradeImpactAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorpackUpgradeImpactAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpactAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorpackUpgradeImpactAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectorPack returns the ConnectorPack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectorpackUpgradeImpactAllOf) GetConnectorPack() []ConnectorpackConnectorPackUpdate {
	if o == nil {
		var ret []ConnectorpackConnectorPackUpdate
		return ret
	}
	return o.ConnectorPack
}

// GetConnectorPackOk returns a tuple with the ConnectorPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectorpackUpgradeImpactAllOf) GetConnectorPackOk() (*[]ConnectorpackConnectorPackUpdate, bool) {
	if o == nil || o.ConnectorPack == nil {
		return nil, false
	}
	return &o.ConnectorPack, true
}

// HasConnectorPack returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpactAllOf) HasConnectorPack() bool {
	if o != nil && o.ConnectorPack != nil {
		return true
	}

	return false
}

// SetConnectorPack gets a reference to the given []ConnectorpackConnectorPackUpdate and assigns it to the ConnectorPack field.
func (o *ConnectorpackUpgradeImpactAllOf) SetConnectorPack(v []ConnectorpackConnectorPackUpdate) {
	o.ConnectorPack = v
}

// GetIsEligibleForUpgrade returns the IsEligibleForUpgrade field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpactAllOf) GetIsEligibleForUpgrade() bool {
	if o == nil || o.IsEligibleForUpgrade == nil {
		var ret bool
		return ret
	}
	return *o.IsEligibleForUpgrade
}

// GetIsEligibleForUpgradeOk returns a tuple with the IsEligibleForUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpactAllOf) GetIsEligibleForUpgradeOk() (*bool, bool) {
	if o == nil || o.IsEligibleForUpgrade == nil {
		return nil, false
	}
	return o.IsEligibleForUpgrade, true
}

// HasIsEligibleForUpgrade returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpactAllOf) HasIsEligibleForUpgrade() bool {
	if o != nil && o.IsEligibleForUpgrade != nil {
		return true
	}

	return false
}

// SetIsEligibleForUpgrade gets a reference to the given bool and assigns it to the IsEligibleForUpgrade field.
func (o *ConnectorpackUpgradeImpactAllOf) SetIsEligibleForUpgrade(v bool) {
	o.IsEligibleForUpgrade = &v
}

// GetIsUpdateDownloaded returns the IsUpdateDownloaded field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpactAllOf) GetIsUpdateDownloaded() bool {
	if o == nil || o.IsUpdateDownloaded == nil {
		var ret bool
		return ret
	}
	return *o.IsUpdateDownloaded
}

// GetIsUpdateDownloadedOk returns a tuple with the IsUpdateDownloaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpactAllOf) GetIsUpdateDownloadedOk() (*bool, bool) {
	if o == nil || o.IsUpdateDownloaded == nil {
		return nil, false
	}
	return o.IsUpdateDownloaded, true
}

// HasIsUpdateDownloaded returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpactAllOf) HasIsUpdateDownloaded() bool {
	if o != nil && o.IsUpdateDownloaded != nil {
		return true
	}

	return false
}

// SetIsUpdateDownloaded gets a reference to the given bool and assigns it to the IsUpdateDownloaded field.
func (o *ConnectorpackUpgradeImpactAllOf) SetIsUpdateDownloaded(v bool) {
	o.IsUpdateDownloaded = &v
}

// GetUcsdInfo returns the UcsdInfo field value if set, zero value otherwise.
func (o *ConnectorpackUpgradeImpactAllOf) GetUcsdInfo() IaasUcsdInfoRelationship {
	if o == nil || o.UcsdInfo == nil {
		var ret IaasUcsdInfoRelationship
		return ret
	}
	return *o.UcsdInfo
}

// GetUcsdInfoOk returns a tuple with the UcsdInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackUpgradeImpactAllOf) GetUcsdInfoOk() (*IaasUcsdInfoRelationship, bool) {
	if o == nil || o.UcsdInfo == nil {
		return nil, false
	}
	return o.UcsdInfo, true
}

// HasUcsdInfo returns a boolean if a field has been set.
func (o *ConnectorpackUpgradeImpactAllOf) HasUcsdInfo() bool {
	if o != nil && o.UcsdInfo != nil {
		return true
	}

	return false
}

// SetUcsdInfo gets a reference to the given IaasUcsdInfoRelationship and assigns it to the UcsdInfo field.
func (o *ConnectorpackUpgradeImpactAllOf) SetUcsdInfo(v IaasUcsdInfoRelationship) {
	o.UcsdInfo = &v
}

func (o ConnectorpackUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectorPack != nil {
		toSerialize["ConnectorPack"] = o.ConnectorPack
	}
	if o.IsEligibleForUpgrade != nil {
		toSerialize["IsEligibleForUpgrade"] = o.IsEligibleForUpgrade
	}
	if o.IsUpdateDownloaded != nil {
		toSerialize["IsUpdateDownloaded"] = o.IsUpdateDownloaded
	}
	if o.UcsdInfo != nil {
		toSerialize["UcsdInfo"] = o.UcsdInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorpackUpgradeImpactAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorpackUpgradeImpactAllOf := _ConnectorpackUpgradeImpactAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorpackUpgradeImpactAllOf); err == nil {
		*o = ConnectorpackUpgradeImpactAllOf(varConnectorpackUpgradeImpactAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectorPack")
		delete(additionalProperties, "IsEligibleForUpgrade")
		delete(additionalProperties, "IsUpdateDownloaded")
		delete(additionalProperties, "UcsdInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorpackUpgradeImpactAllOf struct {
	value *ConnectorpackUpgradeImpactAllOf
	isSet bool
}

func (v NullableConnectorpackUpgradeImpactAllOf) Get() *ConnectorpackUpgradeImpactAllOf {
	return v.value
}

func (v *NullableConnectorpackUpgradeImpactAllOf) Set(val *ConnectorpackUpgradeImpactAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackUpgradeImpactAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackUpgradeImpactAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackUpgradeImpactAllOf(val *ConnectorpackUpgradeImpactAllOf) *NullableConnectorpackUpgradeImpactAllOf {
	return &NullableConnectorpackUpgradeImpactAllOf{value: val, isSet: true}
}

func (v NullableConnectorpackUpgradeImpactAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackUpgradeImpactAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
