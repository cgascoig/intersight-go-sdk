/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4870
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// CommHttpProxyPolicyAllOf Definition of the list of properties defined in 'comm.HttpProxyPolicy', excluding properties defined in parent classes.
type CommHttpProxyPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommHttpProxyPolicyAllOf CommHttpProxyPolicyAllOf

// NewCommHttpProxyPolicyAllOf instantiates a new CommHttpProxyPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommHttpProxyPolicyAllOf(classId string, objectType string) *CommHttpProxyPolicyAllOf {
	this := CommHttpProxyPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCommHttpProxyPolicyAllOfWithDefaults instantiates a new CommHttpProxyPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommHttpProxyPolicyAllOfWithDefaults() *CommHttpProxyPolicyAllOf {
	this := CommHttpProxyPolicyAllOf{}
	var classId string = "comm.HttpProxyPolicy"
	this.ClassId = classId
	var objectType string = "comm.HttpProxyPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommHttpProxyPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommHttpProxyPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommHttpProxyPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommHttpProxyPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommHttpProxyPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommHttpProxyPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommHttpProxyPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommHttpProxyPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *CommHttpProxyPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *CommHttpProxyPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

func (o CommHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommHttpProxyPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCommHttpProxyPolicyAllOf := _CommHttpProxyPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varCommHttpProxyPolicyAllOf); err == nil {
		*o = CommHttpProxyPolicyAllOf(varCommHttpProxyPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterProfiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommHttpProxyPolicyAllOf struct {
	value *CommHttpProxyPolicyAllOf
	isSet bool
}

func (v NullableCommHttpProxyPolicyAllOf) Get() *CommHttpProxyPolicyAllOf {
	return v.value
}

func (v *NullableCommHttpProxyPolicyAllOf) Set(val *CommHttpProxyPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommHttpProxyPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommHttpProxyPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommHttpProxyPolicyAllOf(val *CommHttpProxyPolicyAllOf) *NullableCommHttpProxyPolicyAllOf {
	return &NullableCommHttpProxyPolicyAllOf{value: val, isSet: true}
}

func (v NullableCommHttpProxyPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommHttpProxyPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
