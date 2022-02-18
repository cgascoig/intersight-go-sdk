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

// ConvergedinfraHealthCheckExecution Health check execution result for a health check definition on a converged infra pod.
type ConvergedinfraHealthCheckExecution struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error details of a health check execution failure.
	Error *string `json:"Error,omitempty"`
	// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the pod.
	Result *string `json:"Result,omitempty"`
	// Status of the health check execution. * `Unknown` - Indicates that the health heck execution status is unknown. This mostly happens in case where health check could not be performed due to connectivity issues. * `Succeeded` - Indicates that the health check execution has succeeded. * `Failed` - Indicates that the health check execution has failed. * `Timedout` - Indicates that the health check execution timed out before completion.
	Status *string `json:"Status,omitempty"`
	// A brief summary of health check results.
	Summary               *string                                          `json:"Summary,omitempty"`
	HealthCheckDefinition *ConvergedinfraHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ConvergedinfraHealthCheckExecution ConvergedinfraHealthCheckExecution

// NewConvergedinfraHealthCheckExecution instantiates a new ConvergedinfraHealthCheckExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraHealthCheckExecution(classId string, objectType string) *ConvergedinfraHealthCheckExecution {
	this := ConvergedinfraHealthCheckExecution{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraHealthCheckExecutionWithDefaults instantiates a new ConvergedinfraHealthCheckExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraHealthCheckExecutionWithDefaults() *ConvergedinfraHealthCheckExecution {
	this := ConvergedinfraHealthCheckExecution{}
	var classId string = "convergedinfra.HealthCheckExecution"
	this.ClassId = classId
	var objectType string = "convergedinfra.HealthCheckExecution"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraHealthCheckExecution) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraHealthCheckExecution) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraHealthCheckExecution) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraHealthCheckExecution) SetObjectType(v string) {
	o.ObjectType = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecution) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecution) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ConvergedinfraHealthCheckExecution) SetError(v string) {
	o.Error = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecution) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecution) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ConvergedinfraHealthCheckExecution) SetResult(v string) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecution) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecution) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConvergedinfraHealthCheckExecution) SetStatus(v string) {
	o.Status = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecution) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecution) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ConvergedinfraHealthCheckExecution) SetSummary(v string) {
	o.Summary = &v
}

// GetHealthCheckDefinition returns the HealthCheckDefinition field value if set, zero value otherwise.
func (o *ConvergedinfraHealthCheckExecution) GetHealthCheckDefinition() ConvergedinfraHealthCheckDefinitionRelationship {
	if o == nil || o.HealthCheckDefinition == nil {
		var ret ConvergedinfraHealthCheckDefinitionRelationship
		return ret
	}
	return *o.HealthCheckDefinition
}

// GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraHealthCheckExecution) GetHealthCheckDefinitionOk() (*ConvergedinfraHealthCheckDefinitionRelationship, bool) {
	if o == nil || o.HealthCheckDefinition == nil {
		return nil, false
	}
	return o.HealthCheckDefinition, true
}

// HasHealthCheckDefinition returns a boolean if a field has been set.
func (o *ConvergedinfraHealthCheckExecution) HasHealthCheckDefinition() bool {
	if o != nil && o.HealthCheckDefinition != nil {
		return true
	}

	return false
}

// SetHealthCheckDefinition gets a reference to the given ConvergedinfraHealthCheckDefinitionRelationship and assigns it to the HealthCheckDefinition field.
func (o *ConvergedinfraHealthCheckExecution) SetHealthCheckDefinition(v ConvergedinfraHealthCheckDefinitionRelationship) {
	o.HealthCheckDefinition = &v
}

func (o ConvergedinfraHealthCheckExecution) MarshalJSON() ([]byte, error) {
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
	if o.Error != nil {
		toSerialize["Error"] = o.Error
	}
	if o.Result != nil {
		toSerialize["Result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Summary != nil {
		toSerialize["Summary"] = o.Summary
	}
	if o.HealthCheckDefinition != nil {
		toSerialize["HealthCheckDefinition"] = o.HealthCheckDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraHealthCheckExecution) UnmarshalJSON(bytes []byte) (err error) {
	type ConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Error details of a health check execution failure.
		Error *string `json:"Error,omitempty"`
		// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the pod.
		Result *string `json:"Result,omitempty"`
		// Status of the health check execution. * `Unknown` - Indicates that the health heck execution status is unknown. This mostly happens in case where health check could not be performed due to connectivity issues. * `Succeeded` - Indicates that the health check execution has succeeded. * `Failed` - Indicates that the health check execution has failed. * `Timedout` - Indicates that the health check execution timed out before completion.
		Status *string `json:"Status,omitempty"`
		// A brief summary of health check results.
		Summary               *string                                          `json:"Summary,omitempty"`
		HealthCheckDefinition *ConvergedinfraHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	}

	varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct := ConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraHealthCheckExecution := _ConvergedinfraHealthCheckExecution{}
		varConvergedinfraHealthCheckExecution.ClassId = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.ClassId
		varConvergedinfraHealthCheckExecution.ObjectType = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.ObjectType
		varConvergedinfraHealthCheckExecution.Error = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.Error
		varConvergedinfraHealthCheckExecution.Result = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.Result
		varConvergedinfraHealthCheckExecution.Status = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.Status
		varConvergedinfraHealthCheckExecution.Summary = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.Summary
		varConvergedinfraHealthCheckExecution.HealthCheckDefinition = varConvergedinfraHealthCheckExecutionWithoutEmbeddedStruct.HealthCheckDefinition
		*o = ConvergedinfraHealthCheckExecution(varConvergedinfraHealthCheckExecution)
	} else {
		return err
	}

	varConvergedinfraHealthCheckExecution := _ConvergedinfraHealthCheckExecution{}

	err = json.Unmarshal(bytes, &varConvergedinfraHealthCheckExecution)
	if err == nil {
		o.MoBaseMo = varConvergedinfraHealthCheckExecution.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "HealthCheckDefinition")

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

type NullableConvergedinfraHealthCheckExecution struct {
	value *ConvergedinfraHealthCheckExecution
	isSet bool
}

func (v NullableConvergedinfraHealthCheckExecution) Get() *ConvergedinfraHealthCheckExecution {
	return v.value
}

func (v *NullableConvergedinfraHealthCheckExecution) Set(val *ConvergedinfraHealthCheckExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraHealthCheckExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraHealthCheckExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraHealthCheckExecution(val *ConvergedinfraHealthCheckExecution) *NullableConvergedinfraHealthCheckExecution {
	return &NullableConvergedinfraHealthCheckExecution{value: val, isSet: true}
}

func (v NullableConvergedinfraHealthCheckExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraHealthCheckExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
