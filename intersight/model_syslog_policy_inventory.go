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
	"reflect"
	"strings"
)

// SyslogPolicyInventory The syslog policy configure the syslog server to receive CIMC log entries.
type SyslogPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	LocalClients         []SyslogLocalClientBase  `json:"LocalClients,omitempty"`
	RemoteClients        []SyslogRemoteClientBase `json:"RemoteClients,omitempty"`
	TargetMo             *MoBaseMoRelationship    `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SyslogPolicyInventory SyslogPolicyInventory

// NewSyslogPolicyInventory instantiates a new SyslogPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyslogPolicyInventory(classId string, objectType string) *SyslogPolicyInventory {
	this := SyslogPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSyslogPolicyInventoryWithDefaults instantiates a new SyslogPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyslogPolicyInventoryWithDefaults() *SyslogPolicyInventory {
	this := SyslogPolicyInventory{}
	var classId string = "syslog.PolicyInventory"
	this.ClassId = classId
	var objectType string = "syslog.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SyslogPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SyslogPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SyslogPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SyslogPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SyslogPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocalClients returns the LocalClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicyInventory) GetLocalClients() []SyslogLocalClientBase {
	if o == nil {
		var ret []SyslogLocalClientBase
		return ret
	}
	return o.LocalClients
}

// GetLocalClientsOk returns a tuple with the LocalClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicyInventory) GetLocalClientsOk() (*[]SyslogLocalClientBase, bool) {
	if o == nil || o.LocalClients == nil {
		return nil, false
	}
	return &o.LocalClients, true
}

// HasLocalClients returns a boolean if a field has been set.
func (o *SyslogPolicyInventory) HasLocalClients() bool {
	if o != nil && o.LocalClients != nil {
		return true
	}

	return false
}

// SetLocalClients gets a reference to the given []SyslogLocalClientBase and assigns it to the LocalClients field.
func (o *SyslogPolicyInventory) SetLocalClients(v []SyslogLocalClientBase) {
	o.LocalClients = v
}

// GetRemoteClients returns the RemoteClients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyslogPolicyInventory) GetRemoteClients() []SyslogRemoteClientBase {
	if o == nil {
		var ret []SyslogRemoteClientBase
		return ret
	}
	return o.RemoteClients
}

// GetRemoteClientsOk returns a tuple with the RemoteClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyslogPolicyInventory) GetRemoteClientsOk() (*[]SyslogRemoteClientBase, bool) {
	if o == nil || o.RemoteClients == nil {
		return nil, false
	}
	return &o.RemoteClients, true
}

// HasRemoteClients returns a boolean if a field has been set.
func (o *SyslogPolicyInventory) HasRemoteClients() bool {
	if o != nil && o.RemoteClients != nil {
		return true
	}

	return false
}

// SetRemoteClients gets a reference to the given []SyslogRemoteClientBase and assigns it to the RemoteClients field.
func (o *SyslogPolicyInventory) SetRemoteClients(v []SyslogRemoteClientBase) {
	o.RemoteClients = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *SyslogPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyslogPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *SyslogPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *SyslogPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o SyslogPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocalClients != nil {
		toSerialize["LocalClients"] = o.LocalClients
	}
	if o.RemoteClients != nil {
		toSerialize["RemoteClients"] = o.RemoteClients
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SyslogPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type SyslogPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                   `json:"ObjectType"`
		LocalClients  []SyslogLocalClientBase  `json:"LocalClients,omitempty"`
		RemoteClients []SyslogRemoteClientBase `json:"RemoteClients,omitempty"`
		TargetMo      *MoBaseMoRelationship    `json:"TargetMo,omitempty"`
	}

	varSyslogPolicyInventoryWithoutEmbeddedStruct := SyslogPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSyslogPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varSyslogPolicyInventory := _SyslogPolicyInventory{}
		varSyslogPolicyInventory.ClassId = varSyslogPolicyInventoryWithoutEmbeddedStruct.ClassId
		varSyslogPolicyInventory.ObjectType = varSyslogPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varSyslogPolicyInventory.LocalClients = varSyslogPolicyInventoryWithoutEmbeddedStruct.LocalClients
		varSyslogPolicyInventory.RemoteClients = varSyslogPolicyInventoryWithoutEmbeddedStruct.RemoteClients
		varSyslogPolicyInventory.TargetMo = varSyslogPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = SyslogPolicyInventory(varSyslogPolicyInventory)
	} else {
		return err
	}

	varSyslogPolicyInventory := _SyslogPolicyInventory{}

	err = json.Unmarshal(bytes, &varSyslogPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varSyslogPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocalClients")
		delete(additionalProperties, "RemoteClients")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableSyslogPolicyInventory struct {
	value *SyslogPolicyInventory
	isSet bool
}

func (v NullableSyslogPolicyInventory) Get() *SyslogPolicyInventory {
	return v.value
}

func (v *NullableSyslogPolicyInventory) Set(val *SyslogPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableSyslogPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableSyslogPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyslogPolicyInventory(val *SyslogPolicyInventory) *NullableSyslogPolicyInventory {
	return &NullableSyslogPolicyInventory{value: val, isSet: true}
}

func (v NullableSyslogPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyslogPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
