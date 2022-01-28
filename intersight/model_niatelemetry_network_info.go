/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5208
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryNetworkInfo Stores information related to a network on dcnm controller.
type NiatelemetryNetworkInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                   `json:"ObjectType"`
	ActiveNode NullableNiatelemetryNode `json:"ActiveNode,omitempty"`
	// Returns hostname of the network.
	Hostname *string `json:"Hostname,omitempty"`
	// Returns management IP of the network.
	ManagementtIp *string `json:"ManagementtIp,omitempty"`
	// Returns out of band IP of the network.
	OutofbandIp          *string                  `json:"OutofbandIp,omitempty"`
	StandbyNode          NullableNiatelemetryNode `json:"StandbyNode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNetworkInfo NiatelemetryNetworkInfo

// NewNiatelemetryNetworkInfo instantiates a new NiatelemetryNetworkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNetworkInfo(classId string, objectType string) *NiatelemetryNetworkInfo {
	this := NiatelemetryNetworkInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNetworkInfoWithDefaults instantiates a new NiatelemetryNetworkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNetworkInfoWithDefaults() *NiatelemetryNetworkInfo {
	this := NiatelemetryNetworkInfo{}
	var classId string = "niatelemetry.NetworkInfo"
	this.ClassId = classId
	var objectType string = "niatelemetry.NetworkInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNetworkInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNetworkInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNetworkInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNetworkInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNetworkInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNetworkInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveNode returns the ActiveNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNetworkInfo) GetActiveNode() NiatelemetryNode {
	if o == nil || o.ActiveNode.Get() == nil {
		var ret NiatelemetryNode
		return ret
	}
	return *o.ActiveNode.Get()
}

// GetActiveNodeOk returns a tuple with the ActiveNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNetworkInfo) GetActiveNodeOk() (*NiatelemetryNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveNode.Get(), o.ActiveNode.IsSet()
}

// HasActiveNode returns a boolean if a field has been set.
func (o *NiatelemetryNetworkInfo) HasActiveNode() bool {
	if o != nil && o.ActiveNode.IsSet() {
		return true
	}

	return false
}

// SetActiveNode gets a reference to the given NullableNiatelemetryNode and assigns it to the ActiveNode field.
func (o *NiatelemetryNetworkInfo) SetActiveNode(v NiatelemetryNode) {
	o.ActiveNode.Set(&v)
}

// SetActiveNodeNil sets the value for ActiveNode to be an explicit nil
func (o *NiatelemetryNetworkInfo) SetActiveNodeNil() {
	o.ActiveNode.Set(nil)
}

// UnsetActiveNode ensures that no value is present for ActiveNode, not even an explicit nil
func (o *NiatelemetryNetworkInfo) UnsetActiveNode() {
	o.ActiveNode.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NiatelemetryNetworkInfo) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNetworkInfo) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NiatelemetryNetworkInfo) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NiatelemetryNetworkInfo) SetHostname(v string) {
	o.Hostname = &v
}

// GetManagementtIp returns the ManagementtIp field value if set, zero value otherwise.
func (o *NiatelemetryNetworkInfo) GetManagementtIp() string {
	if o == nil || o.ManagementtIp == nil {
		var ret string
		return ret
	}
	return *o.ManagementtIp
}

// GetManagementtIpOk returns a tuple with the ManagementtIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNetworkInfo) GetManagementtIpOk() (*string, bool) {
	if o == nil || o.ManagementtIp == nil {
		return nil, false
	}
	return o.ManagementtIp, true
}

// HasManagementtIp returns a boolean if a field has been set.
func (o *NiatelemetryNetworkInfo) HasManagementtIp() bool {
	if o != nil && o.ManagementtIp != nil {
		return true
	}

	return false
}

// SetManagementtIp gets a reference to the given string and assigns it to the ManagementtIp field.
func (o *NiatelemetryNetworkInfo) SetManagementtIp(v string) {
	o.ManagementtIp = &v
}

// GetOutofbandIp returns the OutofbandIp field value if set, zero value otherwise.
func (o *NiatelemetryNetworkInfo) GetOutofbandIp() string {
	if o == nil || o.OutofbandIp == nil {
		var ret string
		return ret
	}
	return *o.OutofbandIp
}

// GetOutofbandIpOk returns a tuple with the OutofbandIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNetworkInfo) GetOutofbandIpOk() (*string, bool) {
	if o == nil || o.OutofbandIp == nil {
		return nil, false
	}
	return o.OutofbandIp, true
}

// HasOutofbandIp returns a boolean if a field has been set.
func (o *NiatelemetryNetworkInfo) HasOutofbandIp() bool {
	if o != nil && o.OutofbandIp != nil {
		return true
	}

	return false
}

// SetOutofbandIp gets a reference to the given string and assigns it to the OutofbandIp field.
func (o *NiatelemetryNetworkInfo) SetOutofbandIp(v string) {
	o.OutofbandIp = &v
}

// GetStandbyNode returns the StandbyNode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryNetworkInfo) GetStandbyNode() NiatelemetryNode {
	if o == nil || o.StandbyNode.Get() == nil {
		var ret NiatelemetryNode
		return ret
	}
	return *o.StandbyNode.Get()
}

// GetStandbyNodeOk returns a tuple with the StandbyNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryNetworkInfo) GetStandbyNodeOk() (*NiatelemetryNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.StandbyNode.Get(), o.StandbyNode.IsSet()
}

// HasStandbyNode returns a boolean if a field has been set.
func (o *NiatelemetryNetworkInfo) HasStandbyNode() bool {
	if o != nil && o.StandbyNode.IsSet() {
		return true
	}

	return false
}

// SetStandbyNode gets a reference to the given NullableNiatelemetryNode and assigns it to the StandbyNode field.
func (o *NiatelemetryNetworkInfo) SetStandbyNode(v NiatelemetryNode) {
	o.StandbyNode.Set(&v)
}

// SetStandbyNodeNil sets the value for StandbyNode to be an explicit nil
func (o *NiatelemetryNetworkInfo) SetStandbyNodeNil() {
	o.StandbyNode.Set(nil)
}

// UnsetStandbyNode ensures that no value is present for StandbyNode, not even an explicit nil
func (o *NiatelemetryNetworkInfo) UnsetStandbyNode() {
	o.StandbyNode.Unset()
}

func (o NiatelemetryNetworkInfo) MarshalJSON() ([]byte, error) {
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
	if o.ActiveNode.IsSet() {
		toSerialize["ActiveNode"] = o.ActiveNode.Get()
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.ManagementtIp != nil {
		toSerialize["ManagementtIp"] = o.ManagementtIp
	}
	if o.OutofbandIp != nil {
		toSerialize["OutofbandIp"] = o.OutofbandIp
	}
	if o.StandbyNode.IsSet() {
		toSerialize["StandbyNode"] = o.StandbyNode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNetworkInfo) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryNetworkInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                   `json:"ObjectType"`
		ActiveNode NullableNiatelemetryNode `json:"ActiveNode,omitempty"`
		// Returns hostname of the network.
		Hostname *string `json:"Hostname,omitempty"`
		// Returns management IP of the network.
		ManagementtIp *string `json:"ManagementtIp,omitempty"`
		// Returns out of band IP of the network.
		OutofbandIp *string                  `json:"OutofbandIp,omitempty"`
		StandbyNode NullableNiatelemetryNode `json:"StandbyNode,omitempty"`
	}

	varNiatelemetryNetworkInfoWithoutEmbeddedStruct := NiatelemetryNetworkInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryNetworkInfoWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryNetworkInfo := _NiatelemetryNetworkInfo{}
		varNiatelemetryNetworkInfo.ClassId = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.ClassId
		varNiatelemetryNetworkInfo.ObjectType = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.ObjectType
		varNiatelemetryNetworkInfo.ActiveNode = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.ActiveNode
		varNiatelemetryNetworkInfo.Hostname = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.Hostname
		varNiatelemetryNetworkInfo.ManagementtIp = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.ManagementtIp
		varNiatelemetryNetworkInfo.OutofbandIp = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.OutofbandIp
		varNiatelemetryNetworkInfo.StandbyNode = varNiatelemetryNetworkInfoWithoutEmbeddedStruct.StandbyNode
		*o = NiatelemetryNetworkInfo(varNiatelemetryNetworkInfo)
	} else {
		return err
	}

	varNiatelemetryNetworkInfo := _NiatelemetryNetworkInfo{}

	err = json.Unmarshal(bytes, &varNiatelemetryNetworkInfo)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryNetworkInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveNode")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "ManagementtIp")
		delete(additionalProperties, "OutofbandIp")
		delete(additionalProperties, "StandbyNode")

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

type NullableNiatelemetryNetworkInfo struct {
	value *NiatelemetryNetworkInfo
	isSet bool
}

func (v NullableNiatelemetryNetworkInfo) Get() *NiatelemetryNetworkInfo {
	return v.value
}

func (v *NullableNiatelemetryNetworkInfo) Set(val *NiatelemetryNetworkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNetworkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNetworkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNetworkInfo(val *NiatelemetryNetworkInfo) *NullableNiatelemetryNetworkInfo {
	return &NullableNiatelemetryNetworkInfo{value: val, isSet: true}
}

func (v NullableNiatelemetryNetworkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNetworkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}