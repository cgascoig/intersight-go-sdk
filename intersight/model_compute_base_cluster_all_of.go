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
)

// ComputeBaseClusterAllOf Definition of the list of properties defined in 'compute.BaseCluster', excluding properties defined in parent classes.
type ComputeBaseClusterAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// An array of relationships to storageBaseCluster resources.
	StorageClusters      []StorageBaseClusterRelationship `json:"StorageClusters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBaseClusterAllOf ComputeBaseClusterAllOf

// NewComputeBaseClusterAllOf instantiates a new ComputeBaseClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBaseClusterAllOf(classId string, objectType string) *ComputeBaseClusterAllOf {
	this := ComputeBaseClusterAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewComputeBaseClusterAllOfWithDefaults instantiates a new ComputeBaseClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBaseClusterAllOfWithDefaults() *ComputeBaseClusterAllOf {
	this := ComputeBaseClusterAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeBaseClusterAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeBaseClusterAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeBaseClusterAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeBaseClusterAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeBaseClusterAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeBaseClusterAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetStorageClusters returns the StorageClusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBaseClusterAllOf) GetStorageClusters() []StorageBaseClusterRelationship {
	if o == nil {
		var ret []StorageBaseClusterRelationship
		return ret
	}
	return o.StorageClusters
}

// GetStorageClustersOk returns a tuple with the StorageClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBaseClusterAllOf) GetStorageClustersOk() (*[]StorageBaseClusterRelationship, bool) {
	if o == nil || o.StorageClusters == nil {
		return nil, false
	}
	return &o.StorageClusters, true
}

// HasStorageClusters returns a boolean if a field has been set.
func (o *ComputeBaseClusterAllOf) HasStorageClusters() bool {
	if o != nil && o.StorageClusters != nil {
		return true
	}

	return false
}

// SetStorageClusters gets a reference to the given []StorageBaseClusterRelationship and assigns it to the StorageClusters field.
func (o *ComputeBaseClusterAllOf) SetStorageClusters(v []StorageBaseClusterRelationship) {
	o.StorageClusters = v
}

func (o ComputeBaseClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.StorageClusters != nil {
		toSerialize["StorageClusters"] = o.StorageClusters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBaseClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeBaseClusterAllOf := _ComputeBaseClusterAllOf{}

	if err = json.Unmarshal(bytes, &varComputeBaseClusterAllOf); err == nil {
		*o = ComputeBaseClusterAllOf(varComputeBaseClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "StorageClusters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeBaseClusterAllOf struct {
	value *ComputeBaseClusterAllOf
	isSet bool
}

func (v NullableComputeBaseClusterAllOf) Get() *ComputeBaseClusterAllOf {
	return v.value
}

func (v *NullableComputeBaseClusterAllOf) Set(val *ComputeBaseClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBaseClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBaseClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBaseClusterAllOf(val *ComputeBaseClusterAllOf) *NullableComputeBaseClusterAllOf {
	return &NullableComputeBaseClusterAllOf{value: val, isSet: true}
}

func (v NullableComputeBaseClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBaseClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
