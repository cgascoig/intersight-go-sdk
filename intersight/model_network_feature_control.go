/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NetworkFeatureControl List of features available on a switch along with the index and admin state. These features will allow the user to perform certain set of actions on the switch and get a view of the status of the sub-set feature names.
type NetworkFeatureControl struct {
	InventoryBase
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

type _NetworkFeatureControl NetworkFeatureControl

// NewNetworkFeatureControl instantiates a new NetworkFeatureControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFeatureControl(classId string, objectType string) *NetworkFeatureControl {
	this := NetworkFeatureControl{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkFeatureControlWithDefaults instantiates a new NetworkFeatureControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFeatureControlWithDefaults() *NetworkFeatureControl {
	this := NetworkFeatureControl{}
	var classId string = "network.FeatureControl"
	this.ClassId = classId
	var objectType string = "network.FeatureControl"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkFeatureControl) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkFeatureControl) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkFeatureControl) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkFeatureControl) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NetworkFeatureControl) SetAdminState(v string) {
	o.AdminState = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetInstance() int64 {
	if o == nil || o.Instance == nil {
		var ret int64
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetInstanceOk() (*int64, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given int64 and assigns it to the Instance field.
func (o *NetworkFeatureControl) SetInstance(v int64) {
	o.Instance = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkFeatureControl) SetName(v string) {
	o.Name = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetOperationalState() string {
	if o == nil || o.OperationalState == nil {
		var ret string
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetOperationalStateOk() (*string, bool) {
	if o == nil || o.OperationalState == nil {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasOperationalState() bool {
	if o != nil && o.OperationalState != nil {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given string and assigns it to the OperationalState field.
func (o *NetworkFeatureControl) SetOperationalState(v string) {
	o.OperationalState = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkFeatureControl) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkFeatureControl) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFeatureControl) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkFeatureControl) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkFeatureControl) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkFeatureControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
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

func (o *NetworkFeatureControl) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkFeatureControlWithoutEmbeddedStruct struct {
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
		OperationalState *string                              `json:"OperationalState,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkFeatureControlWithoutEmbeddedStruct := NetworkFeatureControlWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkFeatureControlWithoutEmbeddedStruct)
	if err == nil {
		varNetworkFeatureControl := _NetworkFeatureControl{}
		varNetworkFeatureControl.ClassId = varNetworkFeatureControlWithoutEmbeddedStruct.ClassId
		varNetworkFeatureControl.ObjectType = varNetworkFeatureControlWithoutEmbeddedStruct.ObjectType
		varNetworkFeatureControl.AdminState = varNetworkFeatureControlWithoutEmbeddedStruct.AdminState
		varNetworkFeatureControl.Instance = varNetworkFeatureControlWithoutEmbeddedStruct.Instance
		varNetworkFeatureControl.Name = varNetworkFeatureControlWithoutEmbeddedStruct.Name
		varNetworkFeatureControl.OperationalState = varNetworkFeatureControlWithoutEmbeddedStruct.OperationalState
		varNetworkFeatureControl.NetworkElement = varNetworkFeatureControlWithoutEmbeddedStruct.NetworkElement
		varNetworkFeatureControl.RegisteredDevice = varNetworkFeatureControlWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkFeatureControl(varNetworkFeatureControl)
	} else {
		return err
	}

	varNetworkFeatureControl := _NetworkFeatureControl{}

	err = json.Unmarshal(bytes, &varNetworkFeatureControl)
	if err == nil {
		o.InventoryBase = varNetworkFeatureControl.InventoryBase
	} else {
		return err
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

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkFeatureControl struct {
	value *NetworkFeatureControl
	isSet bool
}

func (v NullableNetworkFeatureControl) Get() *NetworkFeatureControl {
	return v.value
}

func (v *NullableNetworkFeatureControl) Set(val *NetworkFeatureControl) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFeatureControl) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFeatureControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFeatureControl(val *NetworkFeatureControl) *NullableNetworkFeatureControl {
	return &NullableNetworkFeatureControl{value: val, isSet: true}
}

func (v NullableNetworkFeatureControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFeatureControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
