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

// ViewsView An abstract representation of a view. A view provides read-only access to the object model. A view can combine and transform data from multiple data sources.
type ViewsView struct {
	MoBaseMo
	AdditionalProperties map[string]interface{}
}

type _ViewsView ViewsView

// NewViewsView instantiates a new ViewsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewsView(classId string, objectType string) *ViewsView {
	this := ViewsView{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewViewsViewWithDefaults instantiates a new ViewsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewsViewWithDefaults() *ViewsView {
	this := ViewsView{}
	return &this
}

func (o ViewsView) MarshalJSON() ([]byte, error) {
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

func (o *ViewsView) UnmarshalJSON(bytes []byte) (err error) {
	type ViewsViewWithoutEmbeddedStruct struct {
	}

	varViewsViewWithoutEmbeddedStruct := ViewsViewWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varViewsViewWithoutEmbeddedStruct)
	if err == nil {
		varViewsView := _ViewsView{}
		*o = ViewsView(varViewsView)
	} else {
		return err
	}

	varViewsView := _ViewsView{}

	err = json.Unmarshal(bytes, &varViewsView)
	if err == nil {
		o.MoBaseMo = varViewsView.MoBaseMo
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

type NullableViewsView struct {
	value *ViewsView
	isSet bool
}

func (v NullableViewsView) Get() *ViewsView {
	return v.value
}

func (v *NullableViewsView) Set(val *ViewsView) {
	v.value = val
	v.isSet = true
}

func (v NullableViewsView) IsSet() bool {
	return v.isSet
}

func (v *NullableViewsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewsView(val *ViewsView) *NullableViewsView {
	return &NullableViewsView{value: val, isSet: true}
}

func (v NullableViewsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
