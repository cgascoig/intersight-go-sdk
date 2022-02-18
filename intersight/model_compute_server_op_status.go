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

// ComputeServerOpStatus The running workflow details for each server operation that can be performed on the servers.
type ComputeServerOpStatus struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - The state denotes that the settings are applied successfully in the target server. Applying - The state denotes that the settings are being applied in the target server. Failed - The state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// The workflow type being started. The workflow name to distinguish workflow by type.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeServerOpStatus ComputeServerOpStatus

// NewComputeServerOpStatus instantiates a new ComputeServerOpStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerOpStatus(classId string, objectType string) *ComputeServerOpStatus {
	this := ComputeServerOpStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// NewComputeServerOpStatusWithDefaults instantiates a new ComputeServerOpStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerOpStatusWithDefaults() *ComputeServerOpStatus {
	this := ComputeServerOpStatus{}
	var classId string = "compute.ServerOpStatus"
	this.ClassId = classId
	var objectType string = "compute.ServerOpStatus"
	this.ObjectType = objectType
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerOpStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerOpStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerOpStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerOpStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *ComputeServerOpStatus) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatus) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *ComputeServerOpStatus) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *ComputeServerOpStatus) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *ComputeServerOpStatus) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerOpStatus) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *ComputeServerOpStatus) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *ComputeServerOpStatus) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o ComputeServerOpStatus) MarshalJSON() ([]byte, error) {
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
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerOpStatus) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeServerOpStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - The state denotes that the settings are applied successfully in the target server. Applying - The state denotes that the settings are being applied in the target server. Failed - The state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
		ConfigState *string `json:"ConfigState,omitempty"`
		// The workflow type being started. The workflow name to distinguish workflow by type.
		WorkflowType *string `json:"WorkflowType,omitempty"`
	}

	varComputeServerOpStatusWithoutEmbeddedStruct := ComputeServerOpStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeServerOpStatusWithoutEmbeddedStruct)
	if err == nil {
		varComputeServerOpStatus := _ComputeServerOpStatus{}
		varComputeServerOpStatus.ClassId = varComputeServerOpStatusWithoutEmbeddedStruct.ClassId
		varComputeServerOpStatus.ObjectType = varComputeServerOpStatusWithoutEmbeddedStruct.ObjectType
		varComputeServerOpStatus.ConfigState = varComputeServerOpStatusWithoutEmbeddedStruct.ConfigState
		varComputeServerOpStatus.WorkflowType = varComputeServerOpStatusWithoutEmbeddedStruct.WorkflowType
		*o = ComputeServerOpStatus(varComputeServerOpStatus)
	} else {
		return err
	}

	varComputeServerOpStatus := _ComputeServerOpStatus{}

	err = json.Unmarshal(bytes, &varComputeServerOpStatus)
	if err == nil {
		o.MoBaseComplexType = varComputeServerOpStatus.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "WorkflowType")

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

type NullableComputeServerOpStatus struct {
	value *ComputeServerOpStatus
	isSet bool
}

func (v NullableComputeServerOpStatus) Get() *ComputeServerOpStatus {
	return v.value
}

func (v *NullableComputeServerOpStatus) Set(val *ComputeServerOpStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerOpStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerOpStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerOpStatus(val *ComputeServerOpStatus) *NullableComputeServerOpStatus {
	return &NullableComputeServerOpStatus{value: val, isSet: true}
}

func (v NullableComputeServerOpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerOpStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
