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
	"reflect"
	"strings"
)

// ViewHealthStatus Expose a Rest endpoint to return a High level, aggregated status of Intersight components for a given Intersight account user. Meant to inform the user if there's an issue with Intersight components that needs her attention. At this point, Aggregated status is reported for 'Licensing', 'Advisories' and 'Alarms' components.
type ViewHealthStatus struct {
	MoBaseMo
	AdditionalProperties map[string]interface{}
}

type _ViewHealthStatus ViewHealthStatus

// NewViewHealthStatus instantiates a new ViewHealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewHealthStatus(classId string, objectType string) *ViewHealthStatus {
	this := ViewHealthStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewViewHealthStatusWithDefaults instantiates a new ViewHealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewHealthStatusWithDefaults() *ViewHealthStatus {
	this := ViewHealthStatus{}
	return &this
}

func (o ViewHealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ViewHealthStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ViewHealthStatusWithoutEmbeddedStruct struct {
	}

	varViewHealthStatusWithoutEmbeddedStruct := ViewHealthStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varViewHealthStatusWithoutEmbeddedStruct)
	if err == nil {
		varViewHealthStatus := _ViewHealthStatus{}
		*o = ViewHealthStatus(varViewHealthStatus)
	} else {
		return err
	}

	varViewHealthStatus := _ViewHealthStatus{}

	err = json.Unmarshal(bytes, &varViewHealthStatus)
	if err == nil {
		o.MoBaseMo = varViewHealthStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

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

type NullableViewHealthStatus struct {
	value *ViewHealthStatus
	isSet bool
}

func (v NullableViewHealthStatus) Get() *ViewHealthStatus {
	return v.value
}

func (v *NullableViewHealthStatus) Set(val *ViewHealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableViewHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableViewHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewHealthStatus(val *ViewHealthStatus) *NullableViewHealthStatus {
	return &NullableViewHealthStatus{value: val, isSet: true}
}

func (v NullableViewHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
