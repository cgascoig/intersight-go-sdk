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

// VnicFcQosPolicyInventory A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
type VnicFcQosPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The burst traffic, in bytes, allowed on the vHBA.
	Burst *int64 `json:"Burst,omitempty"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
	MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
	// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority *string `json:"Priority,omitempty"`
	// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit            *int64                `json:"RateLimit,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcQosPolicyInventory VnicFcQosPolicyInventory

// NewVnicFcQosPolicyInventory instantiates a new VnicFcQosPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcQosPolicyInventory(classId string, objectType string) *VnicFcQosPolicyInventory {
	this := VnicFcQosPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicFcQosPolicyInventoryWithDefaults instantiates a new VnicFcQosPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcQosPolicyInventoryWithDefaults() *VnicFcQosPolicyInventory {
	this := VnicFcQosPolicyInventory{}
	var classId string = "vnic.FcQosPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.FcQosPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcQosPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcQosPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcQosPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcQosPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBurst returns the Burst field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetBurst() int64 {
	if o == nil || o.Burst == nil {
		var ret int64
		return ret
	}
	return *o.Burst
}

// GetBurstOk returns a tuple with the Burst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetBurstOk() (*int64, bool) {
	if o == nil || o.Burst == nil {
		return nil, false
	}
	return o.Burst, true
}

// HasBurst returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasBurst() bool {
	if o != nil && o.Burst != nil {
		return true
	}

	return false
}

// SetBurst gets a reference to the given int64 and assigns it to the Burst field.
func (o *VnicFcQosPolicyInventory) SetBurst(v int64) {
	o.Burst = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicFcQosPolicyInventory) SetCos(v int64) {
	o.Cos = &v
}

// GetMaxDataFieldSize returns the MaxDataFieldSize field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetMaxDataFieldSize() int64 {
	if o == nil || o.MaxDataFieldSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxDataFieldSize
}

// GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetMaxDataFieldSizeOk() (*int64, bool) {
	if o == nil || o.MaxDataFieldSize == nil {
		return nil, false
	}
	return o.MaxDataFieldSize, true
}

// HasMaxDataFieldSize returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasMaxDataFieldSize() bool {
	if o != nil && o.MaxDataFieldSize != nil {
		return true
	}

	return false
}

// SetMaxDataFieldSize gets a reference to the given int64 and assigns it to the MaxDataFieldSize field.
func (o *VnicFcQosPolicyInventory) SetMaxDataFieldSize(v int64) {
	o.MaxDataFieldSize = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicFcQosPolicyInventory) SetPriority(v string) {
	o.Priority = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicFcQosPolicyInventory) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicFcQosPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicFcQosPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicFcQosPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicFcQosPolicyInventory) MarshalJSON() ([]byte, error) {
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
	if o.Burst != nil {
		toSerialize["Burst"] = o.Burst
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.MaxDataFieldSize != nil {
		toSerialize["MaxDataFieldSize"] = o.MaxDataFieldSize
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcQosPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcQosPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The burst traffic, in bytes, allowed on the vHBA.
		Burst *int64 `json:"Burst,omitempty"`
		// Class of Service to be associated to the traffic on the virtual interface.
		Cos *int64 `json:"Cos,omitempty"`
		// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
		MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
		// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority *string `json:"Priority,omitempty"`
		// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
		RateLimit *int64                `json:"RateLimit,omitempty"`
		TargetMo  *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicFcQosPolicyInventoryWithoutEmbeddedStruct := VnicFcQosPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcQosPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcQosPolicyInventory := _VnicFcQosPolicyInventory{}
		varVnicFcQosPolicyInventory.ClassId = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicFcQosPolicyInventory.ObjectType = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicFcQosPolicyInventory.Burst = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.Burst
		varVnicFcQosPolicyInventory.Cos = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.Cos
		varVnicFcQosPolicyInventory.MaxDataFieldSize = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.MaxDataFieldSize
		varVnicFcQosPolicyInventory.Priority = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.Priority
		varVnicFcQosPolicyInventory.RateLimit = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.RateLimit
		varVnicFcQosPolicyInventory.TargetMo = varVnicFcQosPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicFcQosPolicyInventory(varVnicFcQosPolicyInventory)
	} else {
		return err
	}

	varVnicFcQosPolicyInventory := _VnicFcQosPolicyInventory{}

	err = json.Unmarshal(bytes, &varVnicFcQosPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicFcQosPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Burst")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "MaxDataFieldSize")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "RateLimit")
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

type NullableVnicFcQosPolicyInventory struct {
	value *VnicFcQosPolicyInventory
	isSet bool
}

func (v NullableVnicFcQosPolicyInventory) Get() *VnicFcQosPolicyInventory {
	return v.value
}

func (v *NullableVnicFcQosPolicyInventory) Set(val *VnicFcQosPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcQosPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcQosPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcQosPolicyInventory(val *VnicFcQosPolicyInventory) *NullableVnicFcQosPolicyInventory {
	return &NullableVnicFcQosPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicFcQosPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcQosPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}