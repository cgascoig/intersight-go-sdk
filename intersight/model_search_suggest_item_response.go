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
	"fmt"
)

// SearchSuggestItemResponse - The response body of a HTTP GET request for the 'search.SuggesthItem' resource.
type SearchSuggestItemResponse struct {
	MoDocumentCount       *MoDocumentCount
	SearchSuggestItemList *SearchSuggestItemList
}

// MoDocumentCountAsSearchSuggestItemResponse is a convenience function that returns MoDocumentCount wrapped in SearchSuggestItemResponse
func MoDocumentCountAsSearchSuggestItemResponse(v *MoDocumentCount) SearchSuggestItemResponse {
	return SearchSuggestItemResponse{MoDocumentCount: v}
}

// SearchSuggestItemListAsSearchSuggestItemResponse is a convenience function that returns SearchSuggestItemList wrapped in SearchSuggestItemResponse
func SearchSuggestItemListAsSearchSuggestItemResponse(v *SearchSuggestItemList) SearchSuggestItemResponse {
	return SearchSuggestItemResponse{SearchSuggestItemList: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SearchSuggestItemResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal SearchSuggestItemResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'search.SuggestItem.List'
	if jsonDict["ObjectType"] == "search.SuggestItem.List" {
		// try to unmarshal JSON data into SearchSuggestItemList
		err = json.Unmarshal(data, &dst.SearchSuggestItemList)
		if err == nil {
			return nil // data stored in dst.SearchSuggestItemList, return on the first match
		} else {
			dst.SearchSuggestItemList = nil
			return fmt.Errorf("Failed to unmarshal SearchSuggestItemResponse as SearchSuggestItemList: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SearchSuggestItemResponse) MarshalJSON() ([]byte, error) {
	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.SearchSuggestItemList != nil {
		return json.Marshal(&src.SearchSuggestItemList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SearchSuggestItemResponse) GetActualInstance() interface{} {
	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.SearchSuggestItemList != nil {
		return obj.SearchSuggestItemList
	}

	// all schemas are nil
	return nil
}

type NullableSearchSuggestItemResponse struct {
	value *SearchSuggestItemResponse
	isSet bool
}

func (v NullableSearchSuggestItemResponse) Get() *SearchSuggestItemResponse {
	return v.value
}

func (v *NullableSearchSuggestItemResponse) Set(val *SearchSuggestItemResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSuggestItemResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSuggestItemResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSuggestItemResponse(val *SearchSuggestItemResponse) *NullableSearchSuggestItemResponse {
	return &NullableSearchSuggestItemResponse{value: val, isSet: true}
}

func (v NullableSearchSuggestItemResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSuggestItemResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
