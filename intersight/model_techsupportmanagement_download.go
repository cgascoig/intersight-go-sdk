/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5517
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// TechsupportmanagementDownload Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
type TechsupportmanagementDownload struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                              `json:"ObjectType"`
	TechSupportStatus    *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TechsupportmanagementDownload TechsupportmanagementDownload

// NewTechsupportmanagementDownload instantiates a new TechsupportmanagementDownload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTechsupportmanagementDownload(classId string, objectType string) *TechsupportmanagementDownload {
	this := TechsupportmanagementDownload{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTechsupportmanagementDownloadWithDefaults instantiates a new TechsupportmanagementDownload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTechsupportmanagementDownloadWithDefaults() *TechsupportmanagementDownload {
	this := TechsupportmanagementDownload{}
	var classId string = "techsupportmanagement.Download"
	this.ClassId = classId
	var objectType string = "techsupportmanagement.Download"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TechsupportmanagementDownload) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownload) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TechsupportmanagementDownload) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TechsupportmanagementDownload) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownload) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TechsupportmanagementDownload) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTechSupportStatus returns the TechSupportStatus field value if set, zero value otherwise.
func (o *TechsupportmanagementDownload) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship {
	if o == nil || o.TechSupportStatus == nil {
		var ret TechsupportmanagementTechSupportStatusRelationship
		return ret
	}
	return *o.TechSupportStatus
}

// GetTechSupportStatusOk returns a tuple with the TechSupportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TechsupportmanagementDownload) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool) {
	if o == nil || o.TechSupportStatus == nil {
		return nil, false
	}
	return o.TechSupportStatus, true
}

// HasTechSupportStatus returns a boolean if a field has been set.
func (o *TechsupportmanagementDownload) HasTechSupportStatus() bool {
	if o != nil && o.TechSupportStatus != nil {
		return true
	}

	return false
}

// SetTechSupportStatus gets a reference to the given TechsupportmanagementTechSupportStatusRelationship and assigns it to the TechSupportStatus field.
func (o *TechsupportmanagementDownload) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship) {
	o.TechSupportStatus = &v
}

func (o TechsupportmanagementDownload) MarshalJSON() ([]byte, error) {
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
	if o.TechSupportStatus != nil {
		toSerialize["TechSupportStatus"] = o.TechSupportStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TechsupportmanagementDownload) UnmarshalJSON(bytes []byte) (err error) {
	type TechsupportmanagementDownloadWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                                              `json:"ObjectType"`
		TechSupportStatus *TechsupportmanagementTechSupportStatusRelationship `json:"TechSupportStatus,omitempty"`
	}

	varTechsupportmanagementDownloadWithoutEmbeddedStruct := TechsupportmanagementDownloadWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementDownloadWithoutEmbeddedStruct)
	if err == nil {
		varTechsupportmanagementDownload := _TechsupportmanagementDownload{}
		varTechsupportmanagementDownload.ClassId = varTechsupportmanagementDownloadWithoutEmbeddedStruct.ClassId
		varTechsupportmanagementDownload.ObjectType = varTechsupportmanagementDownloadWithoutEmbeddedStruct.ObjectType
		varTechsupportmanagementDownload.TechSupportStatus = varTechsupportmanagementDownloadWithoutEmbeddedStruct.TechSupportStatus
		*o = TechsupportmanagementDownload(varTechsupportmanagementDownload)
	} else {
		return err
	}

	varTechsupportmanagementDownload := _TechsupportmanagementDownload{}

	err = json.Unmarshal(bytes, &varTechsupportmanagementDownload)
	if err == nil {
		o.MoBaseMo = varTechsupportmanagementDownload.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TechSupportStatus")

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

type NullableTechsupportmanagementDownload struct {
	value *TechsupportmanagementDownload
	isSet bool
}

func (v NullableTechsupportmanagementDownload) Get() *TechsupportmanagementDownload {
	return v.value
}

func (v *NullableTechsupportmanagementDownload) Set(val *TechsupportmanagementDownload) {
	v.value = val
	v.isSet = true
}

func (v NullableTechsupportmanagementDownload) IsSet() bool {
	return v.isSet
}

func (v *NullableTechsupportmanagementDownload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTechsupportmanagementDownload(val *TechsupportmanagementDownload) *NullableTechsupportmanagementDownload {
	return &NullableTechsupportmanagementDownload{value: val, isSet: true}
}

func (v NullableTechsupportmanagementDownload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTechsupportmanagementDownload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
