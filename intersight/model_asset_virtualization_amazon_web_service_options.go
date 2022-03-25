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

// AssetVirtualizationAmazonWebServiceOptions Captures Virtualization Service configuration specific to AWS.
type AssetVirtualizationAmazonWebServiceOptions struct {
	AssetVirtualizationCloudOptions
	AdditionalProperties map[string]interface{}
}

type _AssetVirtualizationAmazonWebServiceOptions AssetVirtualizationAmazonWebServiceOptions

// NewAssetVirtualizationAmazonWebServiceOptions instantiates a new AssetVirtualizationAmazonWebServiceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVirtualizationAmazonWebServiceOptions(classId string, objectType string) *AssetVirtualizationAmazonWebServiceOptions {
	this := AssetVirtualizationAmazonWebServiceOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVirtualizationAmazonWebServiceOptionsWithDefaults instantiates a new AssetVirtualizationAmazonWebServiceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVirtualizationAmazonWebServiceOptionsWithDefaults() *AssetVirtualizationAmazonWebServiceOptions {
	this := AssetVirtualizationAmazonWebServiceOptions{}
	return &this
}

func (o AssetVirtualizationAmazonWebServiceOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetVirtualizationCloudOptions, errAssetVirtualizationCloudOptions := json.Marshal(o.AssetVirtualizationCloudOptions)
	if errAssetVirtualizationCloudOptions != nil {
		return []byte{}, errAssetVirtualizationCloudOptions
	}
	errAssetVirtualizationCloudOptions = json.Unmarshal([]byte(serializedAssetVirtualizationCloudOptions), &toSerialize)
	if errAssetVirtualizationCloudOptions != nil {
		return []byte{}, errAssetVirtualizationCloudOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetVirtualizationAmazonWebServiceOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct struct {
	}

	varAssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct := AssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetVirtualizationAmazonWebServiceOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetVirtualizationAmazonWebServiceOptions := _AssetVirtualizationAmazonWebServiceOptions{}
		*o = AssetVirtualizationAmazonWebServiceOptions(varAssetVirtualizationAmazonWebServiceOptions)
	} else {
		return err
	}

	varAssetVirtualizationAmazonWebServiceOptions := _AssetVirtualizationAmazonWebServiceOptions{}

	err = json.Unmarshal(bytes, &varAssetVirtualizationAmazonWebServiceOptions)
	if err == nil {
		o.AssetVirtualizationCloudOptions = varAssetVirtualizationAmazonWebServiceOptions.AssetVirtualizationCloudOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAssetVirtualizationCloudOptions := reflect.ValueOf(o.AssetVirtualizationCloudOptions)
		for i := 0; i < reflectAssetVirtualizationCloudOptions.Type().NumField(); i++ {
			t := reflectAssetVirtualizationCloudOptions.Type().Field(i)

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

type NullableAssetVirtualizationAmazonWebServiceOptions struct {
	value *AssetVirtualizationAmazonWebServiceOptions
	isSet bool
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) Get() *AssetVirtualizationAmazonWebServiceOptions {
	return v.value
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) Set(val *AssetVirtualizationAmazonWebServiceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVirtualizationAmazonWebServiceOptions(val *AssetVirtualizationAmazonWebServiceOptions) *NullableAssetVirtualizationAmazonWebServiceOptions {
	return &NullableAssetVirtualizationAmazonWebServiceOptions{value: val, isSet: true}
}

func (v NullableAssetVirtualizationAmazonWebServiceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVirtualizationAmazonWebServiceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
