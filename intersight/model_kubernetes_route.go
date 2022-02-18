/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KubernetesRoute Routes include a destination network and a gateway.
type KubernetesRoute struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The destination subnet, if set to 0.0.0.0/0 then the Route is considered a default route.
	To *string `json:"To,omitempty"`
	// Via is the gateway for traffic destined for the subnet in the To property.
	Via                  *string `json:"Via,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesRoute KubernetesRoute

// NewKubernetesRoute instantiates a new KubernetesRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesRoute(classId string, objectType string) *KubernetesRoute {
	this := KubernetesRoute{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesRouteWithDefaults instantiates a new KubernetesRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesRouteWithDefaults() *KubernetesRoute {
	this := KubernetesRoute{}
	var classId string = "kubernetes.Route"
	this.ClassId = classId
	var objectType string = "kubernetes.Route"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesRoute) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesRoute) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesRoute) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesRoute) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesRoute) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesRoute) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *KubernetesRoute) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesRoute) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *KubernetesRoute) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *KubernetesRoute) SetTo(v string) {
	o.To = &v
}

// GetVia returns the Via field value if set, zero value otherwise.
func (o *KubernetesRoute) GetVia() string {
	if o == nil || o.Via == nil {
		var ret string
		return ret
	}
	return *o.Via
}

// GetViaOk returns a tuple with the Via field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesRoute) GetViaOk() (*string, bool) {
	if o == nil || o.Via == nil {
		return nil, false
	}
	return o.Via, true
}

// HasVia returns a boolean if a field has been set.
func (o *KubernetesRoute) HasVia() bool {
	if o != nil && o.Via != nil {
		return true
	}

	return false
}

// SetVia gets a reference to the given string and assigns it to the Via field.
func (o *KubernetesRoute) SetVia(v string) {
	o.Via = &v
}

func (o KubernetesRoute) MarshalJSON() ([]byte, error) {
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
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	if o.Via != nil {
		toSerialize["Via"] = o.Via
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesRoute) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesRouteWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The destination subnet, if set to 0.0.0.0/0 then the Route is considered a default route.
		To *string `json:"To,omitempty"`
		// Via is the gateway for traffic destined for the subnet in the To property.
		Via *string `json:"Via,omitempty"`
	}

	varKubernetesRouteWithoutEmbeddedStruct := KubernetesRouteWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesRouteWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesRoute := _KubernetesRoute{}
		varKubernetesRoute.ClassId = varKubernetesRouteWithoutEmbeddedStruct.ClassId
		varKubernetesRoute.ObjectType = varKubernetesRouteWithoutEmbeddedStruct.ObjectType
		varKubernetesRoute.To = varKubernetesRouteWithoutEmbeddedStruct.To
		varKubernetesRoute.Via = varKubernetesRouteWithoutEmbeddedStruct.Via
		*o = KubernetesRoute(varKubernetesRoute)
	} else {
		return err
	}

	varKubernetesRoute := _KubernetesRoute{}

	err = json.Unmarshal(bytes, &varKubernetesRoute)
	if err == nil {
		o.MoBaseComplexType = varKubernetesRoute.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "To")
		delete(additionalProperties, "Via")

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

type NullableKubernetesRoute struct {
	value *KubernetesRoute
	isSet bool
}

func (v NullableKubernetesRoute) Get() *KubernetesRoute {
	return v.value
}

func (v *NullableKubernetesRoute) Set(val *KubernetesRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesRoute(val *KubernetesRoute) *NullableKubernetesRoute {
	return &NullableKubernetesRoute{value: val, isSet: true}
}

func (v NullableKubernetesRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
