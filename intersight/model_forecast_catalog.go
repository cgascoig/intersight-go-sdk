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

// ForecastCatalog A catalog for managing forecast settings.
type ForecastCatalog struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The time at which the regression model needs to run for all the metrics specified in catalog.
	SchedTime *string `json:"SchedTime,omitempty"`
	// The catalog version used in forecast configuration service.
	Version *string `json:"Version,omitempty"`
	// An array of relationships to forecastDefinition resources.
	Definition           []ForecastDefinitionRelationship `json:"Definition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForecastCatalog ForecastCatalog

// NewForecastCatalog instantiates a new ForecastCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForecastCatalog(classId string, objectType string) *ForecastCatalog {
	this := ForecastCatalog{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewForecastCatalogWithDefaults instantiates a new ForecastCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForecastCatalogWithDefaults() *ForecastCatalog {
	this := ForecastCatalog{}
	var classId string = "forecast.Catalog"
	this.ClassId = classId
	var objectType string = "forecast.Catalog"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ForecastCatalog) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ForecastCatalog) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ForecastCatalog) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ForecastCatalog) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ForecastCatalog) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ForecastCatalog) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSchedTime returns the SchedTime field value if set, zero value otherwise.
func (o *ForecastCatalog) GetSchedTime() string {
	if o == nil || o.SchedTime == nil {
		var ret string
		return ret
	}
	return *o.SchedTime
}

// GetSchedTimeOk returns a tuple with the SchedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastCatalog) GetSchedTimeOk() (*string, bool) {
	if o == nil || o.SchedTime == nil {
		return nil, false
	}
	return o.SchedTime, true
}

// HasSchedTime returns a boolean if a field has been set.
func (o *ForecastCatalog) HasSchedTime() bool {
	if o != nil && o.SchedTime != nil {
		return true
	}

	return false
}

// SetSchedTime gets a reference to the given string and assigns it to the SchedTime field.
func (o *ForecastCatalog) SetSchedTime(v string) {
	o.SchedTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ForecastCatalog) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForecastCatalog) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ForecastCatalog) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ForecastCatalog) SetVersion(v string) {
	o.Version = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForecastCatalog) GetDefinition() []ForecastDefinitionRelationship {
	if o == nil {
		var ret []ForecastDefinitionRelationship
		return ret
	}
	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForecastCatalog) GetDefinitionOk() (*[]ForecastDefinitionRelationship, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return &o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *ForecastCatalog) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given []ForecastDefinitionRelationship and assigns it to the Definition field.
func (o *ForecastCatalog) SetDefinition(v []ForecastDefinitionRelationship) {
	o.Definition = v
}

func (o ForecastCatalog) MarshalJSON() ([]byte, error) {
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
	if o.SchedTime != nil {
		toSerialize["SchedTime"] = o.SchedTime
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Definition != nil {
		toSerialize["Definition"] = o.Definition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ForecastCatalog) UnmarshalJSON(bytes []byte) (err error) {
	type ForecastCatalogWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The time at which the regression model needs to run for all the metrics specified in catalog.
		SchedTime *string `json:"SchedTime,omitempty"`
		// The catalog version used in forecast configuration service.
		Version *string `json:"Version,omitempty"`
		// An array of relationships to forecastDefinition resources.
		Definition []ForecastDefinitionRelationship `json:"Definition,omitempty"`
	}

	varForecastCatalogWithoutEmbeddedStruct := ForecastCatalogWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varForecastCatalogWithoutEmbeddedStruct)
	if err == nil {
		varForecastCatalog := _ForecastCatalog{}
		varForecastCatalog.ClassId = varForecastCatalogWithoutEmbeddedStruct.ClassId
		varForecastCatalog.ObjectType = varForecastCatalogWithoutEmbeddedStruct.ObjectType
		varForecastCatalog.SchedTime = varForecastCatalogWithoutEmbeddedStruct.SchedTime
		varForecastCatalog.Version = varForecastCatalogWithoutEmbeddedStruct.Version
		varForecastCatalog.Definition = varForecastCatalogWithoutEmbeddedStruct.Definition
		*o = ForecastCatalog(varForecastCatalog)
	} else {
		return err
	}

	varForecastCatalog := _ForecastCatalog{}

	err = json.Unmarshal(bytes, &varForecastCatalog)
	if err == nil {
		o.MoBaseMo = varForecastCatalog.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SchedTime")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Definition")

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

type NullableForecastCatalog struct {
	value *ForecastCatalog
	isSet bool
}

func (v NullableForecastCatalog) Get() *ForecastCatalog {
	return v.value
}

func (v *NullableForecastCatalog) Set(val *ForecastCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableForecastCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableForecastCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForecastCatalog(val *ForecastCatalog) *NullableForecastCatalog {
	return &NullableForecastCatalog{value: val, isSet: true}
}

func (v NullableForecastCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForecastCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
