/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricEstimateImpact Before submitting switch profile deploy operation, the estimate impact helps to know the list of components be impacted and require switch reboot.
type FabricEstimateImpact struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                           `json:"ObjectType"`
	Profile              *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricEstimateImpact FabricEstimateImpact

// NewFabricEstimateImpact instantiates a new FabricEstimateImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricEstimateImpact(classId string, objectType string) *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricEstimateImpactWithDefaults instantiates a new FabricEstimateImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricEstimateImpactWithDefaults() *FabricEstimateImpact {
	this := FabricEstimateImpact{}
	var classId string = "fabric.EstimateImpact"
	this.ClassId = classId
	var objectType string = "fabric.EstimateImpact"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricEstimateImpact) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricEstimateImpact) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricEstimateImpact) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricEstimateImpact) SetObjectType(v string) {
	o.ObjectType = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *FabricEstimateImpact) GetProfile() FabricSwitchProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret FabricSwitchProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricEstimateImpact) GetProfileOk() (*FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *FabricEstimateImpact) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given FabricSwitchProfileRelationship and assigns it to the Profile field.
func (o *FabricEstimateImpact) SetProfile(v FabricSwitchProfileRelationship) {
	o.Profile = &v
}

func (o FabricEstimateImpact) MarshalJSON() ([]byte, error) {
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
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricEstimateImpact) UnmarshalJSON(bytes []byte) (err error) {
	type FabricEstimateImpactWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                           `json:"ObjectType"`
		Profile    *FabricSwitchProfileRelationship `json:"Profile,omitempty"`
	}

	varFabricEstimateImpactWithoutEmbeddedStruct := FabricEstimateImpactWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricEstimateImpactWithoutEmbeddedStruct)
	if err == nil {
		varFabricEstimateImpact := _FabricEstimateImpact{}
		varFabricEstimateImpact.ClassId = varFabricEstimateImpactWithoutEmbeddedStruct.ClassId
		varFabricEstimateImpact.ObjectType = varFabricEstimateImpactWithoutEmbeddedStruct.ObjectType
		varFabricEstimateImpact.Profile = varFabricEstimateImpactWithoutEmbeddedStruct.Profile
		*o = FabricEstimateImpact(varFabricEstimateImpact)
	} else {
		return err
	}

	varFabricEstimateImpact := _FabricEstimateImpact{}

	err = json.Unmarshal(bytes, &varFabricEstimateImpact)
	if err == nil {
		o.MoBaseMo = varFabricEstimateImpact.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Profile")

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

type NullableFabricEstimateImpact struct {
	value *FabricEstimateImpact
	isSet bool
}

func (v NullableFabricEstimateImpact) Get() *FabricEstimateImpact {
	return v.value
}

func (v *NullableFabricEstimateImpact) Set(val *FabricEstimateImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricEstimateImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricEstimateImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricEstimateImpact(val *FabricEstimateImpact) *NullableFabricEstimateImpact {
	return &NullableFabricEstimateImpact{value: val, isSet: true}
}

func (v NullableFabricEstimateImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricEstimateImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
