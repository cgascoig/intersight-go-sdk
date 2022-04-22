/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6341
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

// HyperflexClusterHealthCheckExecutionSnapshot Health check execution snapshot of the HyperFlex cluster.
type HyperflexClusterHealthCheckExecutionSnapshot struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Timestamp of the last health check execution on the HyperFlex cluster.
	Timestamp            *time.Time                           `json:"Timestamp,omitempty"`
	HxCluster            *HyperflexClusterRelationship        `json:"HxCluster,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship    `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexClusterHealthCheckExecutionSnapshot HyperflexClusterHealthCheckExecutionSnapshot

// NewHyperflexClusterHealthCheckExecutionSnapshot instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexClusterHealthCheckExecutionSnapshot(classId string, objectType string) *HyperflexClusterHealthCheckExecutionSnapshot {
	this := HyperflexClusterHealthCheckExecutionSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults instantiates a new HyperflexClusterHealthCheckExecutionSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexClusterHealthCheckExecutionSnapshotWithDefaults() *HyperflexClusterHealthCheckExecutionSnapshot {
	this := HyperflexClusterHealthCheckExecutionSnapshot{}
	var classId string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ClassId = classId
	var objectType string = "hyperflex.ClusterHealthCheckExecutionSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetHxCluster returns the HxCluster field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxCluster() HyperflexClusterRelationship {
	if o == nil || o.HxCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.HxCluster
}

// GetHxClusterOk returns a tuple with the HxCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetHxClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.HxCluster == nil {
		return nil, false
	}
	return o.HxCluster, true
}

// HasHxCluster returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasHxCluster() bool {
	if o != nil && o.HxCluster != nil {
		return true
	}

	return false
}

// SetHxCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the HxCluster field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetHxCluster(v HyperflexClusterRelationship) {
	o.HxCluster = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *HyperflexClusterHealthCheckExecutionSnapshot) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o HyperflexClusterHealthCheckExecutionSnapshot) MarshalJSON() ([]byte, error) {
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
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.HxCluster != nil {
		toSerialize["HxCluster"] = o.HxCluster
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexClusterHealthCheckExecutionSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Timestamp of the last health check execution on the HyperFlex cluster.
		Timestamp        *time.Time                           `json:"Timestamp,omitempty"`
		HxCluster        *HyperflexClusterRelationship        `json:"HxCluster,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		Workflow         *WorkflowWorkflowInfoRelationship    `json:"Workflow,omitempty"`
	}

	varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct := HyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexClusterHealthCheckExecutionSnapshot := _HyperflexClusterHealthCheckExecutionSnapshot{}
		varHyperflexClusterHealthCheckExecutionSnapshot.ClassId = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.ClassId
		varHyperflexClusterHealthCheckExecutionSnapshot.ObjectType = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.ObjectType
		varHyperflexClusterHealthCheckExecutionSnapshot.Timestamp = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.Timestamp
		varHyperflexClusterHealthCheckExecutionSnapshot.HxCluster = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.HxCluster
		varHyperflexClusterHealthCheckExecutionSnapshot.RegisteredDevice = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.RegisteredDevice
		varHyperflexClusterHealthCheckExecutionSnapshot.Workflow = varHyperflexClusterHealthCheckExecutionSnapshotWithoutEmbeddedStruct.Workflow
		*o = HyperflexClusterHealthCheckExecutionSnapshot(varHyperflexClusterHealthCheckExecutionSnapshot)
	} else {
		return err
	}

	varHyperflexClusterHealthCheckExecutionSnapshot := _HyperflexClusterHealthCheckExecutionSnapshot{}

	err = json.Unmarshal(bytes, &varHyperflexClusterHealthCheckExecutionSnapshot)
	if err == nil {
		o.MoBaseMo = varHyperflexClusterHealthCheckExecutionSnapshot.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "HxCluster")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Workflow")

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

type NullableHyperflexClusterHealthCheckExecutionSnapshot struct {
	value *HyperflexClusterHealthCheckExecutionSnapshot
	isSet bool
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) Get() *HyperflexClusterHealthCheckExecutionSnapshot {
	return v.value
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) Set(val *HyperflexClusterHealthCheckExecutionSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexClusterHealthCheckExecutionSnapshot(val *HyperflexClusterHealthCheckExecutionSnapshot) *NullableHyperflexClusterHealthCheckExecutionSnapshot {
	return &NullableHyperflexClusterHealthCheckExecutionSnapshot{value: val, isSet: true}
}

func (v NullableHyperflexClusterHealthCheckExecutionSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexClusterHealthCheckExecutionSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
