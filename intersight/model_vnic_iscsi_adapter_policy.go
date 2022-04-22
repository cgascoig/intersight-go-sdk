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

// VnicIscsiAdapterPolicy Set of iSCSI properties that govern the host-side behavior of the adapter.
type VnicIscsiAdapterPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable.
	ConnectionTimeOut *int64 `json:"ConnectionTimeOut,omitempty"`
	// The number of seconds to wait before the initiator assumes that the DHCP server is unavailable.
	DhcpTimeout *int64 `json:"DhcpTimeout,omitempty"`
	// The number of times to retry the connection in case of a failure during iSCSI LUN discovery.
	LunBusyRetryCount    *int64                                `json:"LunBusyRetryCount,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicIscsiAdapterPolicy VnicIscsiAdapterPolicy

// NewVnicIscsiAdapterPolicy instantiates a new VnicIscsiAdapterPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiAdapterPolicy(classId string, objectType string) *VnicIscsiAdapterPolicy {
	this := VnicIscsiAdapterPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiAdapterPolicyWithDefaults instantiates a new VnicIscsiAdapterPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiAdapterPolicyWithDefaults() *VnicIscsiAdapterPolicy {
	this := VnicIscsiAdapterPolicy{}
	var classId string = "vnic.IscsiAdapterPolicy"
	this.ClassId = classId
	var objectType string = "vnic.IscsiAdapterPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiAdapterPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiAdapterPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiAdapterPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiAdapterPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConnectionTimeOut returns the ConnectionTimeOut field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicy) GetConnectionTimeOut() int64 {
	if o == nil || o.ConnectionTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.ConnectionTimeOut
}

// GetConnectionTimeOutOk returns a tuple with the ConnectionTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetConnectionTimeOutOk() (*int64, bool) {
	if o == nil || o.ConnectionTimeOut == nil {
		return nil, false
	}
	return o.ConnectionTimeOut, true
}

// HasConnectionTimeOut returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicy) HasConnectionTimeOut() bool {
	if o != nil && o.ConnectionTimeOut != nil {
		return true
	}

	return false
}

// SetConnectionTimeOut gets a reference to the given int64 and assigns it to the ConnectionTimeOut field.
func (o *VnicIscsiAdapterPolicy) SetConnectionTimeOut(v int64) {
	o.ConnectionTimeOut = &v
}

// GetDhcpTimeout returns the DhcpTimeout field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicy) GetDhcpTimeout() int64 {
	if o == nil || o.DhcpTimeout == nil {
		var ret int64
		return ret
	}
	return *o.DhcpTimeout
}

// GetDhcpTimeoutOk returns a tuple with the DhcpTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetDhcpTimeoutOk() (*int64, bool) {
	if o == nil || o.DhcpTimeout == nil {
		return nil, false
	}
	return o.DhcpTimeout, true
}

// HasDhcpTimeout returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicy) HasDhcpTimeout() bool {
	if o != nil && o.DhcpTimeout != nil {
		return true
	}

	return false
}

// SetDhcpTimeout gets a reference to the given int64 and assigns it to the DhcpTimeout field.
func (o *VnicIscsiAdapterPolicy) SetDhcpTimeout(v int64) {
	o.DhcpTimeout = &v
}

// GetLunBusyRetryCount returns the LunBusyRetryCount field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicy) GetLunBusyRetryCount() int64 {
	if o == nil || o.LunBusyRetryCount == nil {
		var ret int64
		return ret
	}
	return *o.LunBusyRetryCount
}

// GetLunBusyRetryCountOk returns a tuple with the LunBusyRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetLunBusyRetryCountOk() (*int64, bool) {
	if o == nil || o.LunBusyRetryCount == nil {
		return nil, false
	}
	return o.LunBusyRetryCount, true
}

// HasLunBusyRetryCount returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicy) HasLunBusyRetryCount() bool {
	if o != nil && o.LunBusyRetryCount != nil {
		return true
	}

	return false
}

// SetLunBusyRetryCount gets a reference to the given int64 and assigns it to the LunBusyRetryCount field.
func (o *VnicIscsiAdapterPolicy) SetLunBusyRetryCount(v int64) {
	o.LunBusyRetryCount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicIscsiAdapterPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiAdapterPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicIscsiAdapterPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicIscsiAdapterPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicIscsiAdapterPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConnectionTimeOut != nil {
		toSerialize["ConnectionTimeOut"] = o.ConnectionTimeOut
	}
	if o.DhcpTimeout != nil {
		toSerialize["DhcpTimeout"] = o.DhcpTimeout
	}
	if o.LunBusyRetryCount != nil {
		toSerialize["LunBusyRetryCount"] = o.LunBusyRetryCount
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicIscsiAdapterPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicIscsiAdapterPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable.
		ConnectionTimeOut *int64 `json:"ConnectionTimeOut,omitempty"`
		// The number of seconds to wait before the initiator assumes that the DHCP server is unavailable.
		DhcpTimeout *int64 `json:"DhcpTimeout,omitempty"`
		// The number of times to retry the connection in case of a failure during iSCSI LUN discovery.
		LunBusyRetryCount *int64                                `json:"LunBusyRetryCount,omitempty"`
		Organization      *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicIscsiAdapterPolicyWithoutEmbeddedStruct := VnicIscsiAdapterPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicIscsiAdapterPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicIscsiAdapterPolicy := _VnicIscsiAdapterPolicy{}
		varVnicIscsiAdapterPolicy.ClassId = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.ClassId
		varVnicIscsiAdapterPolicy.ObjectType = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.ObjectType
		varVnicIscsiAdapterPolicy.ConnectionTimeOut = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.ConnectionTimeOut
		varVnicIscsiAdapterPolicy.DhcpTimeout = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.DhcpTimeout
		varVnicIscsiAdapterPolicy.LunBusyRetryCount = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.LunBusyRetryCount
		varVnicIscsiAdapterPolicy.Organization = varVnicIscsiAdapterPolicyWithoutEmbeddedStruct.Organization
		*o = VnicIscsiAdapterPolicy(varVnicIscsiAdapterPolicy)
	} else {
		return err
	}

	varVnicIscsiAdapterPolicy := _VnicIscsiAdapterPolicy{}

	err = json.Unmarshal(bytes, &varVnicIscsiAdapterPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicIscsiAdapterPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConnectionTimeOut")
		delete(additionalProperties, "DhcpTimeout")
		delete(additionalProperties, "LunBusyRetryCount")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicIscsiAdapterPolicy struct {
	value *VnicIscsiAdapterPolicy
	isSet bool
}

func (v NullableVnicIscsiAdapterPolicy) Get() *VnicIscsiAdapterPolicy {
	return v.value
}

func (v *NullableVnicIscsiAdapterPolicy) Set(val *VnicIscsiAdapterPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiAdapterPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiAdapterPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiAdapterPolicy(val *VnicIscsiAdapterPolicy) *NullableVnicIscsiAdapterPolicy {
	return &NullableVnicIscsiAdapterPolicy{value: val, isSet: true}
}

func (v NullableVnicIscsiAdapterPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiAdapterPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
