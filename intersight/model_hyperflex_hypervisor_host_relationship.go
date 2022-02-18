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

// HyperflexHypervisorHostRelationship - A relationship to the 'hyperflex.HypervisorHost' resource, or the expanded 'hyperflex.HypervisorHost' resource, or the 'null' value.
type HyperflexHypervisorHostRelationship struct {
	HyperflexHypervisorHost *HyperflexHypervisorHost
	MoMoRef                 *MoMoRef
}

// HyperflexHypervisorHostAsHyperflexHypervisorHostRelationship is a convenience function that returns HyperflexHypervisorHost wrapped in HyperflexHypervisorHostRelationship
func HyperflexHypervisorHostAsHyperflexHypervisorHostRelationship(v *HyperflexHypervisorHost) HyperflexHypervisorHostRelationship {
	return HyperflexHypervisorHostRelationship{HyperflexHypervisorHost: v}
}

// MoMoRefAsHyperflexHypervisorHostRelationship is a convenience function that returns MoMoRef wrapped in HyperflexHypervisorHostRelationship
func MoMoRefAsHyperflexHypervisorHostRelationship(v *MoMoRef) HyperflexHypervisorHostRelationship {
	return HyperflexHypervisorHostRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *HyperflexHypervisorHostRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'hyperflex.HypervisorHost'
	if jsonDict["ClassId"] == "hyperflex.HypervisorHost" {
		// try to unmarshal JSON data into HyperflexHypervisorHost
		err = json.Unmarshal(data, &dst.HyperflexHypervisorHost)
		if err == nil {
			return nil // data stored in dst.HyperflexHypervisorHost, return on the first match
		} else {
			dst.HyperflexHypervisorHost = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHypervisorHostRelationship as HyperflexHypervisorHost: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal HyperflexHypervisorHostRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src HyperflexHypervisorHostRelationship) MarshalJSON() ([]byte, error) {
	if src.HyperflexHypervisorHost != nil {
		return json.Marshal(&src.HyperflexHypervisorHost)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *HyperflexHypervisorHostRelationship) GetActualInstance() interface{} {
	if obj.HyperflexHypervisorHost != nil {
		return obj.HyperflexHypervisorHost
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableHyperflexHypervisorHostRelationship struct {
	value *HyperflexHypervisorHostRelationship
	isSet bool
}

func (v NullableHyperflexHypervisorHostRelationship) Get() *HyperflexHypervisorHostRelationship {
	return v.value
}

func (v *NullableHyperflexHypervisorHostRelationship) Set(val *HyperflexHypervisorHostRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHypervisorHostRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHypervisorHostRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHypervisorHostRelationship(val *HyperflexHypervisorHostRelationship) *NullableHyperflexHypervisorHostRelationship {
	return &NullableHyperflexHypervisorHostRelationship{value: val, isSet: true}
}

func (v NullableHyperflexHypervisorHostRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHypervisorHostRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
