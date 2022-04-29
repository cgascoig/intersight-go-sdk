/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-6484
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"time"
)

// OprsDeploymentAllOf Definition of the list of properties defined in 'oprs.Deployment', excluding properties defined in parent classes.
type OprsDeploymentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Available number of replicas.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
	// The expected number of replicas.
	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty"`
	// The type of event which was triggered.
	Event  *string      `json:"Event,omitempty"`
	Labels []OprsKvpair `json:"Labels,omitempty"`
	// Agent name for which the event is triggered.
	Name *string `json:"Name,omitempty"`
	// Name space in which the agents are running.
	Namespace *string `json:"Namespace,omitempty"`
	// Status which shows if the resource is healthy or not. * `` - An Unknown status indicates that the resource status is not known. * `Healthy` - A healthy status indicates that the resource is healthy and running as per spec. * `Unhealthy` - An unhealthy status indicates that the resource is down.
	Status *string `json:"Status,omitempty"`
	// The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events.
	TimeStamp *time.Time `json:"TimeStamp,omitempty"`
	// Number of replicas Unavailable.
	UnavailableReplicas  *int64                               `json:"UnavailableReplicas,omitempty"`
	Assist               *AssetDeviceRegistrationRelationship `json:"Assist,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OprsDeploymentAllOf OprsDeploymentAllOf

// NewOprsDeploymentAllOf instantiates a new OprsDeploymentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOprsDeploymentAllOf(classId string, objectType string) *OprsDeploymentAllOf {
	this := OprsDeploymentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// NewOprsDeploymentAllOfWithDefaults instantiates a new OprsDeploymentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOprsDeploymentAllOfWithDefaults() *OprsDeploymentAllOf {
	this := OprsDeploymentAllOf{}
	var classId string = "oprs.Deployment"
	this.ClassId = classId
	var objectType string = "oprs.Deployment"
	this.ObjectType = objectType
	var status string = ""
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *OprsDeploymentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OprsDeploymentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OprsDeploymentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OprsDeploymentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvailableReplicas returns the AvailableReplicas field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetAvailableReplicas() int64 {
	if o == nil || o.AvailableReplicas == nil {
		var ret int64
		return ret
	}
	return *o.AvailableReplicas
}

// GetAvailableReplicasOk returns a tuple with the AvailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetAvailableReplicasOk() (*int64, bool) {
	if o == nil || o.AvailableReplicas == nil {
		return nil, false
	}
	return o.AvailableReplicas, true
}

// HasAvailableReplicas returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasAvailableReplicas() bool {
	if o != nil && o.AvailableReplicas != nil {
		return true
	}

	return false
}

// SetAvailableReplicas gets a reference to the given int64 and assigns it to the AvailableReplicas field.
func (o *OprsDeploymentAllOf) SetAvailableReplicas(v int64) {
	o.AvailableReplicas = &v
}

// GetDesiredReplicas returns the DesiredReplicas field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetDesiredReplicas() int64 {
	if o == nil || o.DesiredReplicas == nil {
		var ret int64
		return ret
	}
	return *o.DesiredReplicas
}

// GetDesiredReplicasOk returns a tuple with the DesiredReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetDesiredReplicasOk() (*int64, bool) {
	if o == nil || o.DesiredReplicas == nil {
		return nil, false
	}
	return o.DesiredReplicas, true
}

// HasDesiredReplicas returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasDesiredReplicas() bool {
	if o != nil && o.DesiredReplicas != nil {
		return true
	}

	return false
}

// SetDesiredReplicas gets a reference to the given int64 and assigns it to the DesiredReplicas field.
func (o *OprsDeploymentAllOf) SetDesiredReplicas(v int64) {
	o.DesiredReplicas = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *OprsDeploymentAllOf) SetEvent(v string) {
	o.Event = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OprsDeploymentAllOf) GetLabels() []OprsKvpair {
	if o == nil {
		var ret []OprsKvpair
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OprsDeploymentAllOf) GetLabelsOk() (*[]OprsKvpair, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []OprsKvpair and assigns it to the Labels field.
func (o *OprsDeploymentAllOf) SetLabels(v []OprsKvpair) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OprsDeploymentAllOf) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *OprsDeploymentAllOf) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OprsDeploymentAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetTimeStamp() time.Time {
	if o == nil || o.TimeStamp == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || o.TimeStamp == nil {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasTimeStamp() bool {
	if o != nil && o.TimeStamp != nil {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *OprsDeploymentAllOf) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetUnavailableReplicas returns the UnavailableReplicas field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetUnavailableReplicas() int64 {
	if o == nil || o.UnavailableReplicas == nil {
		var ret int64
		return ret
	}
	return *o.UnavailableReplicas
}

// GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetUnavailableReplicasOk() (*int64, bool) {
	if o == nil || o.UnavailableReplicas == nil {
		return nil, false
	}
	return o.UnavailableReplicas, true
}

// HasUnavailableReplicas returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasUnavailableReplicas() bool {
	if o != nil && o.UnavailableReplicas != nil {
		return true
	}

	return false
}

// SetUnavailableReplicas gets a reference to the given int64 and assigns it to the UnavailableReplicas field.
func (o *OprsDeploymentAllOf) SetUnavailableReplicas(v int64) {
	o.UnavailableReplicas = &v
}

// GetAssist returns the Assist field value if set, zero value otherwise.
func (o *OprsDeploymentAllOf) GetAssist() AssetDeviceRegistrationRelationship {
	if o == nil || o.Assist == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Assist
}

// GetAssistOk returns a tuple with the Assist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OprsDeploymentAllOf) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Assist == nil {
		return nil, false
	}
	return o.Assist, true
}

// HasAssist returns a boolean if a field has been set.
func (o *OprsDeploymentAllOf) HasAssist() bool {
	if o != nil && o.Assist != nil {
		return true
	}

	return false
}

// SetAssist gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Assist field.
func (o *OprsDeploymentAllOf) SetAssist(v AssetDeviceRegistrationRelationship) {
	o.Assist = &v
}

func (o OprsDeploymentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvailableReplicas != nil {
		toSerialize["AvailableReplicas"] = o.AvailableReplicas
	}
	if o.DesiredReplicas != nil {
		toSerialize["DesiredReplicas"] = o.DesiredReplicas
	}
	if o.Event != nil {
		toSerialize["Event"] = o.Event
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.TimeStamp != nil {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if o.UnavailableReplicas != nil {
		toSerialize["UnavailableReplicas"] = o.UnavailableReplicas
	}
	if o.Assist != nil {
		toSerialize["Assist"] = o.Assist
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OprsDeploymentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varOprsDeploymentAllOf := _OprsDeploymentAllOf{}

	if err = json.Unmarshal(bytes, &varOprsDeploymentAllOf); err == nil {
		*o = OprsDeploymentAllOf(varOprsDeploymentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailableReplicas")
		delete(additionalProperties, "DesiredReplicas")
		delete(additionalProperties, "Event")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TimeStamp")
		delete(additionalProperties, "UnavailableReplicas")
		delete(additionalProperties, "Assist")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOprsDeploymentAllOf struct {
	value *OprsDeploymentAllOf
	isSet bool
}

func (v NullableOprsDeploymentAllOf) Get() *OprsDeploymentAllOf {
	return v.value
}

func (v *NullableOprsDeploymentAllOf) Set(val *OprsDeploymentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOprsDeploymentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOprsDeploymentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOprsDeploymentAllOf(val *OprsDeploymentAllOf) *NullableOprsDeploymentAllOf {
	return &NullableOprsDeploymentAllOf{value: val, isSet: true}
}

func (v NullableOprsDeploymentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOprsDeploymentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
