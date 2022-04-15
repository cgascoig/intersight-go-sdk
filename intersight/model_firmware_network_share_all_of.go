/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6207
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// FirmwareNetworkShareAllOf Definition of the list of properties defined in 'firmware.NetworkShare', excluding properties defined in parent classes.
type FirmwareNetworkShareAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                     `json:"ObjectType"`
	CifsServer NullableFirmwareCifsServer `json:"CifsServer,omitempty"`
	HttpServer NullableFirmwareHttpServer `json:"HttpServer,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image. * `nfs` - File server protocol used is NFS. * `cifs` - File server protocol used is CIFS. * `www` - File server protocol used is WWW.
	MapType   *string                   `json:"MapType,omitempty"`
	NfsServer NullableFirmwareNfsServer `json:"NfsServer,omitempty"`
	// Password as configured on the file server.
	Password *string `json:"Password,omitempty"`
	// Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade. * `nw_upgrade_full` - Network upgrade option for full upgrade. * `nw_upgrade_mount_only` - Network upgrade mount only.
	Upgradeoption *string `json:"Upgradeoption,omitempty"`
	// Username as configured on the file server.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareNetworkShareAllOf FirmwareNetworkShareAllOf

// NewFirmwareNetworkShareAllOf instantiates a new FirmwareNetworkShareAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareNetworkShareAllOf(classId string, objectType string) *FirmwareNetworkShareAllOf {
	this := FirmwareNetworkShareAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var mapType string = "nfs"
	this.MapType = &mapType
	var upgradeoption string = "nw_upgrade_full"
	this.Upgradeoption = &upgradeoption
	return &this
}

// NewFirmwareNetworkShareAllOfWithDefaults instantiates a new FirmwareNetworkShareAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareNetworkShareAllOfWithDefaults() *FirmwareNetworkShareAllOf {
	this := FirmwareNetworkShareAllOf{}
	var classId string = "firmware.NetworkShare"
	this.ClassId = classId
	var objectType string = "firmware.NetworkShare"
	this.ObjectType = objectType
	var mapType string = "nfs"
	this.MapType = &mapType
	var upgradeoption string = "nw_upgrade_full"
	this.Upgradeoption = &upgradeoption
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareNetworkShareAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareNetworkShareAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareNetworkShareAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareNetworkShareAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCifsServer returns the CifsServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShareAllOf) GetCifsServer() FirmwareCifsServer {
	if o == nil || o.CifsServer.Get() == nil {
		var ret FirmwareCifsServer
		return ret
	}
	return *o.CifsServer.Get()
}

// GetCifsServerOk returns a tuple with the CifsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShareAllOf) GetCifsServerOk() (*FirmwareCifsServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.CifsServer.Get(), o.CifsServer.IsSet()
}

// HasCifsServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasCifsServer() bool {
	if o != nil && o.CifsServer.IsSet() {
		return true
	}

	return false
}

// SetCifsServer gets a reference to the given NullableFirmwareCifsServer and assigns it to the CifsServer field.
func (o *FirmwareNetworkShareAllOf) SetCifsServer(v FirmwareCifsServer) {
	o.CifsServer.Set(&v)
}

// SetCifsServerNil sets the value for CifsServer to be an explicit nil
func (o *FirmwareNetworkShareAllOf) SetCifsServerNil() {
	o.CifsServer.Set(nil)
}

// UnsetCifsServer ensures that no value is present for CifsServer, not even an explicit nil
func (o *FirmwareNetworkShareAllOf) UnsetCifsServer() {
	o.CifsServer.Unset()
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShareAllOf) GetHttpServer() FirmwareHttpServer {
	if o == nil || o.HttpServer.Get() == nil {
		var ret FirmwareHttpServer
		return ret
	}
	return *o.HttpServer.Get()
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShareAllOf) GetHttpServerOk() (*FirmwareHttpServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpServer.Get(), o.HttpServer.IsSet()
}

// HasHttpServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasHttpServer() bool {
	if o != nil && o.HttpServer.IsSet() {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given NullableFirmwareHttpServer and assigns it to the HttpServer field.
func (o *FirmwareNetworkShareAllOf) SetHttpServer(v FirmwareHttpServer) {
	o.HttpServer.Set(&v)
}

// SetHttpServerNil sets the value for HttpServer to be an explicit nil
func (o *FirmwareNetworkShareAllOf) SetHttpServerNil() {
	o.HttpServer.Set(nil)
}

// UnsetHttpServer ensures that no value is present for HttpServer, not even an explicit nil
func (o *FirmwareNetworkShareAllOf) UnsetHttpServer() {
	o.HttpServer.Unset()
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *FirmwareNetworkShareAllOf) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *FirmwareNetworkShareAllOf) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetMapType returns the MapType field value if set, zero value otherwise.
func (o *FirmwareNetworkShareAllOf) GetMapType() string {
	if o == nil || o.MapType == nil {
		var ret string
		return ret
	}
	return *o.MapType
}

// GetMapTypeOk returns a tuple with the MapType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetMapTypeOk() (*string, bool) {
	if o == nil || o.MapType == nil {
		return nil, false
	}
	return o.MapType, true
}

// HasMapType returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasMapType() bool {
	if o != nil && o.MapType != nil {
		return true
	}

	return false
}

// SetMapType gets a reference to the given string and assigns it to the MapType field.
func (o *FirmwareNetworkShareAllOf) SetMapType(v string) {
	o.MapType = &v
}

// GetNfsServer returns the NfsServer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareNetworkShareAllOf) GetNfsServer() FirmwareNfsServer {
	if o == nil || o.NfsServer.Get() == nil {
		var ret FirmwareNfsServer
		return ret
	}
	return *o.NfsServer.Get()
}

// GetNfsServerOk returns a tuple with the NfsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareNetworkShareAllOf) GetNfsServerOk() (*FirmwareNfsServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.NfsServer.Get(), o.NfsServer.IsSet()
}

// HasNfsServer returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasNfsServer() bool {
	if o != nil && o.NfsServer.IsSet() {
		return true
	}

	return false
}

// SetNfsServer gets a reference to the given NullableFirmwareNfsServer and assigns it to the NfsServer field.
func (o *FirmwareNetworkShareAllOf) SetNfsServer(v FirmwareNfsServer) {
	o.NfsServer.Set(&v)
}

// SetNfsServerNil sets the value for NfsServer to be an explicit nil
func (o *FirmwareNetworkShareAllOf) SetNfsServerNil() {
	o.NfsServer.Set(nil)
}

// UnsetNfsServer ensures that no value is present for NfsServer, not even an explicit nil
func (o *FirmwareNetworkShareAllOf) UnsetNfsServer() {
	o.NfsServer.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FirmwareNetworkShareAllOf) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FirmwareNetworkShareAllOf) SetPassword(v string) {
	o.Password = &v
}

// GetUpgradeoption returns the Upgradeoption field value if set, zero value otherwise.
func (o *FirmwareNetworkShareAllOf) GetUpgradeoption() string {
	if o == nil || o.Upgradeoption == nil {
		var ret string
		return ret
	}
	return *o.Upgradeoption
}

// GetUpgradeoptionOk returns a tuple with the Upgradeoption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetUpgradeoptionOk() (*string, bool) {
	if o == nil || o.Upgradeoption == nil {
		return nil, false
	}
	return o.Upgradeoption, true
}

// HasUpgradeoption returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasUpgradeoption() bool {
	if o != nil && o.Upgradeoption != nil {
		return true
	}

	return false
}

// SetUpgradeoption gets a reference to the given string and assigns it to the Upgradeoption field.
func (o *FirmwareNetworkShareAllOf) SetUpgradeoption(v string) {
	o.Upgradeoption = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *FirmwareNetworkShareAllOf) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNetworkShareAllOf) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *FirmwareNetworkShareAllOf) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *FirmwareNetworkShareAllOf) SetUsername(v string) {
	o.Username = &v
}

func (o FirmwareNetworkShareAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CifsServer.IsSet() {
		toSerialize["CifsServer"] = o.CifsServer.Get()
	}
	if o.HttpServer.IsSet() {
		toSerialize["HttpServer"] = o.HttpServer.Get()
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.MapType != nil {
		toSerialize["MapType"] = o.MapType
	}
	if o.NfsServer.IsSet() {
		toSerialize["NfsServer"] = o.NfsServer.Get()
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Upgradeoption != nil {
		toSerialize["Upgradeoption"] = o.Upgradeoption
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareNetworkShareAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareNetworkShareAllOf := _FirmwareNetworkShareAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareNetworkShareAllOf); err == nil {
		*o = FirmwareNetworkShareAllOf(varFirmwareNetworkShareAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CifsServer")
		delete(additionalProperties, "HttpServer")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "MapType")
		delete(additionalProperties, "NfsServer")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Upgradeoption")
		delete(additionalProperties, "Username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareNetworkShareAllOf struct {
	value *FirmwareNetworkShareAllOf
	isSet bool
}

func (v NullableFirmwareNetworkShareAllOf) Get() *FirmwareNetworkShareAllOf {
	return v.value
}

func (v *NullableFirmwareNetworkShareAllOf) Set(val *FirmwareNetworkShareAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareNetworkShareAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareNetworkShareAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareNetworkShareAllOf(val *FirmwareNetworkShareAllOf) *NullableFirmwareNetworkShareAllOf {
	return &NullableFirmwareNetworkShareAllOf{value: val, isSet: true}
}

func (v NullableFirmwareNetworkShareAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareNetworkShareAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
