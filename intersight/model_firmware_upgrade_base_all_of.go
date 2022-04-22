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

// FirmwareUpgradeBaseAllOf Definition of the list of properties defined in 'firmware.UpgradeBase', excluding properties defined in parent classes.
type FirmwareUpgradeBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType     string                               `json:"ObjectType"`
	DirectDownload NullableFirmwareDirectDownload       `json:"DirectDownload,omitempty"`
	FileServer     NullableSoftwarerepositoryFileServer `json:"FileServer,omitempty"`
	NetworkShare   NullableFirmwareNetworkShare         `json:"NetworkShare,omitempty"`
	// User has the option to skip the estimate impact calculation.
	SkipEstimateImpact *bool `json:"SkipEstimateImpact,omitempty"`
	// Status of the upgrade operation. * `NONE` - Upgrade status is not populated. * `IN_PROGRESS` - The upgrade is in progress. * `SUCCESSFUL` - The upgrade successfully completed. * `FAILED` - The upgrade shows failed status. * `TERMINATED` - The upgrade has been terminated.
	Status *string `json:"Status,omitempty"`
	// Desired upgrade mode to choose either direct download based upgrade or network share upgrade. * `direct_upgrade` - Upgrade mode is direct download. * `network_upgrade` - Upgrade mode is network upgrade.
	UpgradeType          *string                                  `json:"UpgradeType,omitempty"`
	Distributable        *FirmwareDistributableRelationship       `json:"Distributable,omitempty"`
	Release              *SoftwarerepositoryReleaseRelationship   `json:"Release,omitempty"`
	UpgradeImpact        *FirmwareUpgradeImpactStatusRelationship `json:"UpgradeImpact,omitempty"`
	UpgradeStatus        *FirmwareUpgradeStatusRelationship       `json:"UpgradeStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeBaseAllOf FirmwareUpgradeBaseAllOf

// NewFirmwareUpgradeBaseAllOf instantiates a new FirmwareUpgradeBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeBaseAllOf(classId string, objectType string) *FirmwareUpgradeBaseAllOf {
	this := FirmwareUpgradeBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "NONE"
	this.Status = &status
	var upgradeType string = "direct_upgrade"
	this.UpgradeType = &upgradeType
	return &this
}

// NewFirmwareUpgradeBaseAllOfWithDefaults instantiates a new FirmwareUpgradeBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeBaseAllOfWithDefaults() *FirmwareUpgradeBaseAllOf {
	this := FirmwareUpgradeBaseAllOf{}
	var status string = "NONE"
	this.Status = &status
	var upgradeType string = "direct_upgrade"
	this.UpgradeType = &upgradeType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareUpgradeBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareUpgradeBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareUpgradeBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareUpgradeBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDirectDownload returns the DirectDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeBaseAllOf) GetDirectDownload() FirmwareDirectDownload {
	if o == nil || o.DirectDownload.Get() == nil {
		var ret FirmwareDirectDownload
		return ret
	}
	return *o.DirectDownload.Get()
}

// GetDirectDownloadOk returns a tuple with the DirectDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeBaseAllOf) GetDirectDownloadOk() (*FirmwareDirectDownload, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectDownload.Get(), o.DirectDownload.IsSet()
}

// HasDirectDownload returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasDirectDownload() bool {
	if o != nil && o.DirectDownload.IsSet() {
		return true
	}

	return false
}

// SetDirectDownload gets a reference to the given NullableFirmwareDirectDownload and assigns it to the DirectDownload field.
func (o *FirmwareUpgradeBaseAllOf) SetDirectDownload(v FirmwareDirectDownload) {
	o.DirectDownload.Set(&v)
}

// SetDirectDownloadNil sets the value for DirectDownload to be an explicit nil
func (o *FirmwareUpgradeBaseAllOf) SetDirectDownloadNil() {
	o.DirectDownload.Set(nil)
}

// UnsetDirectDownload ensures that no value is present for DirectDownload, not even an explicit nil
func (o *FirmwareUpgradeBaseAllOf) UnsetDirectDownload() {
	o.DirectDownload.Unset()
}

// GetFileServer returns the FileServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeBaseAllOf) GetFileServer() SoftwarerepositoryFileServer {
	if o == nil || o.FileServer.Get() == nil {
		var ret SoftwarerepositoryFileServer
		return ret
	}
	return *o.FileServer.Get()
}

// GetFileServerOk returns a tuple with the FileServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeBaseAllOf) GetFileServerOk() (*SoftwarerepositoryFileServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileServer.Get(), o.FileServer.IsSet()
}

// HasFileServer returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasFileServer() bool {
	if o != nil && o.FileServer.IsSet() {
		return true
	}

	return false
}

// SetFileServer gets a reference to the given NullableSoftwarerepositoryFileServer and assigns it to the FileServer field.
func (o *FirmwareUpgradeBaseAllOf) SetFileServer(v SoftwarerepositoryFileServer) {
	o.FileServer.Set(&v)
}

// SetFileServerNil sets the value for FileServer to be an explicit nil
func (o *FirmwareUpgradeBaseAllOf) SetFileServerNil() {
	o.FileServer.Set(nil)
}

// UnsetFileServer ensures that no value is present for FileServer, not even an explicit nil
func (o *FirmwareUpgradeBaseAllOf) UnsetFileServer() {
	o.FileServer.Unset()
}

// GetNetworkShare returns the NetworkShare field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeBaseAllOf) GetNetworkShare() FirmwareNetworkShare {
	if o == nil || o.NetworkShare.Get() == nil {
		var ret FirmwareNetworkShare
		return ret
	}
	return *o.NetworkShare.Get()
}

// GetNetworkShareOk returns a tuple with the NetworkShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeBaseAllOf) GetNetworkShareOk() (*FirmwareNetworkShare, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkShare.Get(), o.NetworkShare.IsSet()
}

// HasNetworkShare returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasNetworkShare() bool {
	if o != nil && o.NetworkShare.IsSet() {
		return true
	}

	return false
}

// SetNetworkShare gets a reference to the given NullableFirmwareNetworkShare and assigns it to the NetworkShare field.
func (o *FirmwareUpgradeBaseAllOf) SetNetworkShare(v FirmwareNetworkShare) {
	o.NetworkShare.Set(&v)
}

// SetNetworkShareNil sets the value for NetworkShare to be an explicit nil
func (o *FirmwareUpgradeBaseAllOf) SetNetworkShareNil() {
	o.NetworkShare.Set(nil)
}

// UnsetNetworkShare ensures that no value is present for NetworkShare, not even an explicit nil
func (o *FirmwareUpgradeBaseAllOf) UnsetNetworkShare() {
	o.NetworkShare.Unset()
}

// GetSkipEstimateImpact returns the SkipEstimateImpact field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetSkipEstimateImpact() bool {
	if o == nil || o.SkipEstimateImpact == nil {
		var ret bool
		return ret
	}
	return *o.SkipEstimateImpact
}

// GetSkipEstimateImpactOk returns a tuple with the SkipEstimateImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetSkipEstimateImpactOk() (*bool, bool) {
	if o == nil || o.SkipEstimateImpact == nil {
		return nil, false
	}
	return o.SkipEstimateImpact, true
}

// HasSkipEstimateImpact returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasSkipEstimateImpact() bool {
	if o != nil && o.SkipEstimateImpact != nil {
		return true
	}

	return false
}

// SetSkipEstimateImpact gets a reference to the given bool and assigns it to the SkipEstimateImpact field.
func (o *FirmwareUpgradeBaseAllOf) SetSkipEstimateImpact(v bool) {
	o.SkipEstimateImpact = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FirmwareUpgradeBaseAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetUpgradeType returns the UpgradeType field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeType() string {
	if o == nil || o.UpgradeType == nil {
		var ret string
		return ret
	}
	return *o.UpgradeType
}

// GetUpgradeTypeOk returns a tuple with the UpgradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeTypeOk() (*string, bool) {
	if o == nil || o.UpgradeType == nil {
		return nil, false
	}
	return o.UpgradeType, true
}

// HasUpgradeType returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasUpgradeType() bool {
	if o != nil && o.UpgradeType != nil {
		return true
	}

	return false
}

// SetUpgradeType gets a reference to the given string and assigns it to the UpgradeType field.
func (o *FirmwareUpgradeBaseAllOf) SetUpgradeType(v string) {
	o.UpgradeType = &v
}

// GetDistributable returns the Distributable field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetDistributable() FirmwareDistributableRelationship {
	if o == nil || o.Distributable == nil {
		var ret FirmwareDistributableRelationship
		return ret
	}
	return *o.Distributable
}

// GetDistributableOk returns a tuple with the Distributable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetDistributableOk() (*FirmwareDistributableRelationship, bool) {
	if o == nil || o.Distributable == nil {
		return nil, false
	}
	return o.Distributable, true
}

// HasDistributable returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasDistributable() bool {
	if o != nil && o.Distributable != nil {
		return true
	}

	return false
}

// SetDistributable gets a reference to the given FirmwareDistributableRelationship and assigns it to the Distributable field.
func (o *FirmwareUpgradeBaseAllOf) SetDistributable(v FirmwareDistributableRelationship) {
	o.Distributable = &v
}

// GetRelease returns the Release field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetRelease() SoftwarerepositoryReleaseRelationship {
	if o == nil || o.Release == nil {
		var ret SoftwarerepositoryReleaseRelationship
		return ret
	}
	return *o.Release
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool) {
	if o == nil || o.Release == nil {
		return nil, false
	}
	return o.Release, true
}

// HasRelease returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasRelease() bool {
	if o != nil && o.Release != nil {
		return true
	}

	return false
}

// SetRelease gets a reference to the given SoftwarerepositoryReleaseRelationship and assigns it to the Release field.
func (o *FirmwareUpgradeBaseAllOf) SetRelease(v SoftwarerepositoryReleaseRelationship) {
	o.Release = &v
}

// GetUpgradeImpact returns the UpgradeImpact field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeImpact() FirmwareUpgradeImpactStatusRelationship {
	if o == nil || o.UpgradeImpact == nil {
		var ret FirmwareUpgradeImpactStatusRelationship
		return ret
	}
	return *o.UpgradeImpact
}

// GetUpgradeImpactOk returns a tuple with the UpgradeImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeImpactOk() (*FirmwareUpgradeImpactStatusRelationship, bool) {
	if o == nil || o.UpgradeImpact == nil {
		return nil, false
	}
	return o.UpgradeImpact, true
}

// HasUpgradeImpact returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasUpgradeImpact() bool {
	if o != nil && o.UpgradeImpact != nil {
		return true
	}

	return false
}

// SetUpgradeImpact gets a reference to the given FirmwareUpgradeImpactStatusRelationship and assigns it to the UpgradeImpact field.
func (o *FirmwareUpgradeBaseAllOf) SetUpgradeImpact(v FirmwareUpgradeImpactStatusRelationship) {
	o.UpgradeImpact = &v
}

// GetUpgradeStatus returns the UpgradeStatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeStatus() FirmwareUpgradeStatusRelationship {
	if o == nil || o.UpgradeStatus == nil {
		var ret FirmwareUpgradeStatusRelationship
		return ret
	}
	return *o.UpgradeStatus
}

// GetUpgradeStatusOk returns a tuple with the UpgradeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeBaseAllOf) GetUpgradeStatusOk() (*FirmwareUpgradeStatusRelationship, bool) {
	if o == nil || o.UpgradeStatus == nil {
		return nil, false
	}
	return o.UpgradeStatus, true
}

// HasUpgradeStatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeBaseAllOf) HasUpgradeStatus() bool {
	if o != nil && o.UpgradeStatus != nil {
		return true
	}

	return false
}

// SetUpgradeStatus gets a reference to the given FirmwareUpgradeStatusRelationship and assigns it to the UpgradeStatus field.
func (o *FirmwareUpgradeBaseAllOf) SetUpgradeStatus(v FirmwareUpgradeStatusRelationship) {
	o.UpgradeStatus = &v
}

func (o FirmwareUpgradeBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DirectDownload.IsSet() {
		toSerialize["DirectDownload"] = o.DirectDownload.Get()
	}
	if o.FileServer.IsSet() {
		toSerialize["FileServer"] = o.FileServer.Get()
	}
	if o.NetworkShare.IsSet() {
		toSerialize["NetworkShare"] = o.NetworkShare.Get()
	}
	if o.SkipEstimateImpact != nil {
		toSerialize["SkipEstimateImpact"] = o.SkipEstimateImpact
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UpgradeType != nil {
		toSerialize["UpgradeType"] = o.UpgradeType
	}
	if o.Distributable != nil {
		toSerialize["Distributable"] = o.Distributable
	}
	if o.Release != nil {
		toSerialize["Release"] = o.Release
	}
	if o.UpgradeImpact != nil {
		toSerialize["UpgradeImpact"] = o.UpgradeImpact
	}
	if o.UpgradeStatus != nil {
		toSerialize["UpgradeStatus"] = o.UpgradeStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareUpgradeBaseAllOf := _FirmwareUpgradeBaseAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareUpgradeBaseAllOf); err == nil {
		*o = FirmwareUpgradeBaseAllOf(varFirmwareUpgradeBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DirectDownload")
		delete(additionalProperties, "FileServer")
		delete(additionalProperties, "NetworkShare")
		delete(additionalProperties, "SkipEstimateImpact")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UpgradeType")
		delete(additionalProperties, "Distributable")
		delete(additionalProperties, "Release")
		delete(additionalProperties, "UpgradeImpact")
		delete(additionalProperties, "UpgradeStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareUpgradeBaseAllOf struct {
	value *FirmwareUpgradeBaseAllOf
	isSet bool
}

func (v NullableFirmwareUpgradeBaseAllOf) Get() *FirmwareUpgradeBaseAllOf {
	return v.value
}

func (v *NullableFirmwareUpgradeBaseAllOf) Set(val *FirmwareUpgradeBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeBaseAllOf(val *FirmwareUpgradeBaseAllOf) *NullableFirmwareUpgradeBaseAllOf {
	return &NullableFirmwareUpgradeBaseAllOf{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
