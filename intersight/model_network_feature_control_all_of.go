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

// NetworkFeatureControlAllOf Definition of the list of properties defined in 'network.FeatureControl', excluding properties defined in parent classes.
type NetworkFeatureControlAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The admin state of the feature.
	AdminState *string `json:"AdminState,omitempty"`
	// The number of instances of the feature.
	Instance *int64 `json:"Instance,omitempty"`
	// The name to identify the feature.
	Name *string `json:"Name,omitempty"`
	// The operational state of the feature.
	OperationalState     *string                              `json:"OperationalState,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkFeatureControlAllOf NetworkFeatureControlAllOf

// NewNetworkFeatureControlAllOf instantiates a new NetworkFeatureControlAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeatureControlAllOf(classId string, objectType string) *NetworkFeatureControlAllOf {
	this := NetworkFeatureControlAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkFeatureControlAllOfWithDefaults instantiates a new NetworkFeatureControlAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeatureControlAllOfWithDefaults() *NetworkFeatureControlAllOf {
	this := NetworkFeatureControlAllOf{}
	var classId string = "network.FeatureControl"
	this.ClassId = classId
	var objectType string = "network.FeatureControl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkFeatureControlAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkFeatureControlAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkFeatureControlAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkFeatureControlAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NetworkFeatureControlAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetInstance() int64 {
	if o == nil || o.Instance == nil {
		var ret int64
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetInstanceOk() (*int64, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given int64 and assigns it to the Instance field.
func (o *NetworkFeatureControlAllOf) SetInstance(v int64) {
	o.Instance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkFeatureControlAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetOperationalState() string {
	if o == nil || o.OperationalState == nil {
		var ret string
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetOperationalStateOk() (*string, bool) {
	if o == nil || o.OperationalState == nil {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasOperationalState() bool {
	if o != nil && o.OperationalState != nil {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given string and assigns it to the OperationalState field.
func (o *NetworkFeatureControlAllOf) SetOperationalState(v string) {
	o.OperationalState = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkFeatureControlAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkFeatureControlAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControlAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkFeatureControlAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkFeatureControlAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkFeatureControlAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.Instance != nil {
		toSerialize["Instance"] = o.Instance
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperationalState != nil {
		toSerialize["OperationalState"] = o.OperationalState
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkFeatureControlAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkFeatureControlAllOf := _NetworkFeatureControlAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkFeatureControlAllOf); err == nil {
		*o = NetworkFeatureControlAllOf(varNetworkFeatureControlAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "Instance")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperationalState")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkFeatureControlAllOf struct {
	value *NetworkFeatureControlAllOf
	isSet bool
}

func (v NullableNetworkFeatureControlAllOf) Get() *NetworkFeatureControlAllOf {
	return v.value
}

func (v *NullableNetworkFeatureControlAllOf) Set(val *NetworkFeatureControlAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeatureControlAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeatureControlAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeatureControlAllOf(val *NetworkFeatureControlAllOf) *NullableNetworkFeatureControlAllOf {
	return &NullableNetworkFeatureControlAllOf{value: val, isSet: true}
}

func (v NullableNetworkFeatureControlAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeatureControlAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
