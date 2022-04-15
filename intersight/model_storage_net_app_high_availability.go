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
	"reflect"
	"strings"
)

// StorageNetAppHighAvailability Storage failover and giveback information.
type StorageNetAppHighAvailability struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether or not storage failover is enabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// The state of the node that is giving storage back to its HA partner. * `unknown` - Default unknown giveback state. * `nothing_to_giveback` - The node has nothing to give back to its HA partner. * `not_attempted` - The node has not attempted to give back storage to its HA partner. * `in_progress` - The node is in progress of giving back storage to its HA partner. * `failed` - The node has failed to give back storage to its HA partner.
	GivebackState *string `json:"GivebackState,omitempty"`
	// The partner node name in this node's High Availability (HA) group.
	PartnerName *string `json:"PartnerName,omitempty"`
	// The partner node uuid in this node's High Availability (HA) group.
	PartnerUuid *string `json:"PartnerUuid,omitempty"`
	// The state of the node that is taking over storage from its HA partner. * `unknown` - Default unknown takeover state. * `not_possible` - It is not possible for the node to take over storage from its HA partner. * `not_attempted` - The node has not attempted to take over storage from its HA partner. * `in_takeover` - The node has taken over storage from its HA partner. * `in_progress` - The node is in progress of taking over storage from its HA partner. * `failed` - The node has failed to take over storage from its HA partner.
	TakeoverState        *string `json:"TakeoverState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppHighAvailability StorageNetAppHighAvailability

// NewStorageNetAppHighAvailability instantiates a new StorageNetAppHighAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppHighAvailability(classId string, objectType string) *StorageNetAppHighAvailability {
	this := StorageNetAppHighAvailability{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppHighAvailabilityWithDefaults instantiates a new StorageNetAppHighAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppHighAvailabilityWithDefaults() *StorageNetAppHighAvailability {
	this := StorageNetAppHighAvailability{}
	var classId string = "storage.NetAppHighAvailability"
	this.ClassId = classId
	var objectType string = "storage.NetAppHighAvailability"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppHighAvailability) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppHighAvailability) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppHighAvailability) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppHighAvailability) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StorageNetAppHighAvailability) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGivebackState returns the GivebackState field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetGivebackState() string {
	if o == nil || o.GivebackState == nil {
		var ret string
		return ret
	}
	return *o.GivebackState
}

// GetGivebackStateOk returns a tuple with the GivebackState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetGivebackStateOk() (*string, bool) {
	if o == nil || o.GivebackState == nil {
		return nil, false
	}
	return o.GivebackState, true
}

// HasGivebackState returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasGivebackState() bool {
	if o != nil && o.GivebackState != nil {
		return true
	}

	return false
}

// SetGivebackState gets a reference to the given string and assigns it to the GivebackState field.
func (o *StorageNetAppHighAvailability) SetGivebackState(v string) {
	o.GivebackState = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetPartnerName() string {
	if o == nil || o.PartnerName == nil {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetPartnerNameOk() (*string, bool) {
	if o == nil || o.PartnerName == nil {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasPartnerName() bool {
	if o != nil && o.PartnerName != nil {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *StorageNetAppHighAvailability) SetPartnerName(v string) {
	o.PartnerName = &v
}

// GetPartnerUuid returns the PartnerUuid field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetPartnerUuid() string {
	if o == nil || o.PartnerUuid == nil {
		var ret string
		return ret
	}
	return *o.PartnerUuid
}

// GetPartnerUuidOk returns a tuple with the PartnerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetPartnerUuidOk() (*string, bool) {
	if o == nil || o.PartnerUuid == nil {
		return nil, false
	}
	return o.PartnerUuid, true
}

// HasPartnerUuid returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasPartnerUuid() bool {
	if o != nil && o.PartnerUuid != nil {
		return true
	}

	return false
}

// SetPartnerUuid gets a reference to the given string and assigns it to the PartnerUuid field.
func (o *StorageNetAppHighAvailability) SetPartnerUuid(v string) {
	o.PartnerUuid = &v
}

// GetTakeoverState returns the TakeoverState field value if set, zero value otherwise.
func (o *StorageNetAppHighAvailability) GetTakeoverState() string {
	if o == nil || o.TakeoverState == nil {
		var ret string
		return ret
	}
	return *o.TakeoverState
}

// GetTakeoverStateOk returns a tuple with the TakeoverState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppHighAvailability) GetTakeoverStateOk() (*string, bool) {
	if o == nil || o.TakeoverState == nil {
		return nil, false
	}
	return o.TakeoverState, true
}

// HasTakeoverState returns a boolean if a field has been set.
func (o *StorageNetAppHighAvailability) HasTakeoverState() bool {
	if o != nil && o.TakeoverState != nil {
		return true
	}

	return false
}

// SetTakeoverState gets a reference to the given string and assigns it to the TakeoverState field.
func (o *StorageNetAppHighAvailability) SetTakeoverState(v string) {
	o.TakeoverState = &v
}

func (o StorageNetAppHighAvailability) MarshalJSON() ([]byte, error) {
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
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.GivebackState != nil {
		toSerialize["GivebackState"] = o.GivebackState
	}
	if o.PartnerName != nil {
		toSerialize["PartnerName"] = o.PartnerName
	}
	if o.PartnerUuid != nil {
		toSerialize["PartnerUuid"] = o.PartnerUuid
	}
	if o.TakeoverState != nil {
		toSerialize["TakeoverState"] = o.TakeoverState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppHighAvailability) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppHighAvailabilityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies whether or not storage failover is enabled.
		Enabled *bool `json:"Enabled,omitempty"`
		// The state of the node that is giving storage back to its HA partner. * `unknown` - Default unknown giveback state. * `nothing_to_giveback` - The node has nothing to give back to its HA partner. * `not_attempted` - The node has not attempted to give back storage to its HA partner. * `in_progress` - The node is in progress of giving back storage to its HA partner. * `failed` - The node has failed to give back storage to its HA partner.
		GivebackState *string `json:"GivebackState,omitempty"`
		// The partner node name in this node's High Availability (HA) group.
		PartnerName *string `json:"PartnerName,omitempty"`
		// The partner node uuid in this node's High Availability (HA) group.
		PartnerUuid *string `json:"PartnerUuid,omitempty"`
		// The state of the node that is taking over storage from its HA partner. * `unknown` - Default unknown takeover state. * `not_possible` - It is not possible for the node to take over storage from its HA partner. * `not_attempted` - The node has not attempted to take over storage from its HA partner. * `in_takeover` - The node has taken over storage from its HA partner. * `in_progress` - The node is in progress of taking over storage from its HA partner. * `failed` - The node has failed to take over storage from its HA partner.
		TakeoverState *string `json:"TakeoverState,omitempty"`
	}

	varStorageNetAppHighAvailabilityWithoutEmbeddedStruct := StorageNetAppHighAvailabilityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppHighAvailabilityWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppHighAvailability := _StorageNetAppHighAvailability{}
		varStorageNetAppHighAvailability.ClassId = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.ClassId
		varStorageNetAppHighAvailability.ObjectType = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.ObjectType
		varStorageNetAppHighAvailability.Enabled = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.Enabled
		varStorageNetAppHighAvailability.GivebackState = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.GivebackState
		varStorageNetAppHighAvailability.PartnerName = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.PartnerName
		varStorageNetAppHighAvailability.PartnerUuid = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.PartnerUuid
		varStorageNetAppHighAvailability.TakeoverState = varStorageNetAppHighAvailabilityWithoutEmbeddedStruct.TakeoverState
		*o = StorageNetAppHighAvailability(varStorageNetAppHighAvailability)
	} else {
		return err
	}

	varStorageNetAppHighAvailability := _StorageNetAppHighAvailability{}

	err = json.Unmarshal(bytes, &varStorageNetAppHighAvailability)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppHighAvailability.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "GivebackState")
		delete(additionalProperties, "PartnerName")
		delete(additionalProperties, "PartnerUuid")
		delete(additionalProperties, "TakeoverState")

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

type NullableStorageNetAppHighAvailability struct {
	value *StorageNetAppHighAvailability
	isSet bool
}

func (v NullableStorageNetAppHighAvailability) Get() *StorageNetAppHighAvailability {
	return v.value
}

func (v *NullableStorageNetAppHighAvailability) Set(val *StorageNetAppHighAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppHighAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppHighAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppHighAvailability(val *StorageNetAppHighAvailability) *NullableStorageNetAppHighAvailability {
	return &NullableStorageNetAppHighAvailability{value: val, isSet: true}
}

func (v NullableStorageNetAppHighAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppHighAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
