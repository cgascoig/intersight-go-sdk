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

// ResourceGroupMember A resolved member of a resource group.
type ResourceGroupMember struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                     `json:"ObjectType"`
	Group                *ResourceGroupRelationship `json:"Group,omitempty"`
	Resource             *MoBaseMoRelationship      `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceGroupMember ResourceGroupMember

// NewResourceGroupMember instantiates a new ResourceGroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceGroupMember(classId string, objectType string) *ResourceGroupMember {
	this := ResourceGroupMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceGroupMemberWithDefaults instantiates a new ResourceGroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceGroupMemberWithDefaults() *ResourceGroupMember {
	this := ResourceGroupMember{}
	var classId string = "resource.GroupMember"
	this.ClassId = classId
	var objectType string = "resource.GroupMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceGroupMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceGroupMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceGroupMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceGroupMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceGroupMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceGroupMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ResourceGroupMember) GetGroup() ResourceGroupRelationship {
	if o == nil || o.Group == nil {
		var ret ResourceGroupRelationship
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroupMember) GetGroupOk() (*ResourceGroupRelationship, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ResourceGroupMember) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given ResourceGroupRelationship and assigns it to the Group field.
func (o *ResourceGroupMember) SetGroup(v ResourceGroupRelationship) {
	o.Group = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceGroupMember) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGroupMember) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceGroupMember) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourceGroupMember) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o ResourceGroupMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Group != nil {
		toSerialize["Group"] = o.Group
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceGroupMember) UnmarshalJSON(bytes []byte) (err error) {
	type ResourceGroupMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                     `json:"ObjectType"`
		Group      *ResourceGroupRelationship `json:"Group,omitempty"`
		Resource   *MoBaseMoRelationship      `json:"Resource,omitempty"`
	}

	varResourceGroupMemberWithoutEmbeddedStruct := ResourceGroupMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varResourceGroupMemberWithoutEmbeddedStruct)
	if err == nil {
		varResourceGroupMember := _ResourceGroupMember{}
		varResourceGroupMember.ClassId = varResourceGroupMemberWithoutEmbeddedStruct.ClassId
		varResourceGroupMember.ObjectType = varResourceGroupMemberWithoutEmbeddedStruct.ObjectType
		varResourceGroupMember.Group = varResourceGroupMemberWithoutEmbeddedStruct.Group
		varResourceGroupMember.Resource = varResourceGroupMemberWithoutEmbeddedStruct.Resource
		*o = ResourceGroupMember(varResourceGroupMember)
	} else {
		return err
	}

	varResourceGroupMember := _ResourceGroupMember{}

	err = json.Unmarshal(bytes, &varResourceGroupMember)
	if err == nil {
		o.MoBaseMo = varResourceGroupMember.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Group")
		delete(additionalProperties, "Resource")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableResourceGroupMember struct {
	value *ResourceGroupMember
	isSet bool
}

func (v NullableResourceGroupMember) Get() *ResourceGroupMember {
	return v.value
}

func (v *NullableResourceGroupMember) Set(val *ResourceGroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceGroupMember(val *ResourceGroupMember) *NullableResourceGroupMember {
	return &NullableResourceGroupMember{value: val, isSet: true}
}

func (v NullableResourceGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
