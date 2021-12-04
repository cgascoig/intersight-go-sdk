/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricVsanList This resource list is returned as a response to a HTTP GET request that does not include a specific resource identifier.
type FabricVsanList struct {
	MoBaseResponse
	// The total number of 'fabric.Vsan' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
	Count *int32 `json:"Count,omitempty"`
	// The array of 'fabric.Vsan' resources matching the request.
	Results              []FabricVsan `json:"Results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVsanList FabricVsanList

// NewFabricVsanList instantiates a new FabricVsanList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVsanList(objectType string) *FabricVsanList {
	this := FabricVsanList{}
	this.ObjectType = objectType
	return &this
}

// NewFabricVsanListWithDefaults instantiates a new FabricVsanList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVsanListWithDefaults() *FabricVsanList {
	this := FabricVsanList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FabricVsanList) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanList) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FabricVsanList) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FabricVsanList) SetCount(v int32) {
	o.Count = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricVsanList) GetResults() []FabricVsan {
	if o == nil {
		var ret []FabricVsan
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricVsanList) GetResultsOk() (*[]FabricVsan, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *FabricVsanList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []FabricVsan and assigns it to the Results field.
func (o *FabricVsanList) SetResults(v []FabricVsan) {
	o.Results = v
}

func (o FabricVsanList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseResponse, errMoBaseResponse := json.Marshal(o.MoBaseResponse)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	errMoBaseResponse = json.Unmarshal([]byte(serializedMoBaseResponse), &toSerialize)
	if errMoBaseResponse != nil {
		return []byte{}, errMoBaseResponse
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVsanList) UnmarshalJSON(bytes []byte) (err error) {
	type FabricVsanListWithoutEmbeddedStruct struct {
		// The total number of 'fabric.Vsan' resources matching the request, accross all pages. The 'Count' attribute is included when the HTTP GET request includes the '$inlinecount' parameter.
		Count *int32 `json:"Count,omitempty"`
		// The array of 'fabric.Vsan' resources matching the request.
		Results []FabricVsan `json:"Results,omitempty"`
	}

	varFabricVsanListWithoutEmbeddedStruct := FabricVsanListWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricVsanListWithoutEmbeddedStruct)
	if err == nil {
		varFabricVsanList := _FabricVsanList{}
		varFabricVsanList.Count = varFabricVsanListWithoutEmbeddedStruct.Count
		varFabricVsanList.Results = varFabricVsanListWithoutEmbeddedStruct.Results
		*o = FabricVsanList(varFabricVsanList)
	} else {
		return err
	}

	varFabricVsanList := _FabricVsanList{}

	err = json.Unmarshal(bytes, &varFabricVsanList)
	if err == nil {
		o.MoBaseResponse = varFabricVsanList.MoBaseResponse
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Count")
		delete(additionalProperties, "Results")

		// remove fields from embedded structs
		reflectMoBaseResponse := reflect.ValueOf(o.MoBaseResponse)
		for i := 0; i < reflectMoBaseResponse.Type().NumField(); i++ {
			t := reflectMoBaseResponse.Type().Field(i)

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

type NullableFabricVsanList struct {
	value *FabricVsanList
	isSet bool
}

func (v NullableFabricVsanList) Get() *FabricVsanList {
	return v.value
}

func (v *NullableFabricVsanList) Set(val *FabricVsanList) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVsanList) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVsanList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVsanList(val *FabricVsanList) *NullableFabricVsanList {
	return &NullableFabricVsanList{value: val, isSet: true}
}

func (v NullableFabricVsanList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVsanList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
