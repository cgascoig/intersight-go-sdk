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
)

// NotificationAggregationSettingsAllOf Definition of the list of properties defined in 'notification.AggregationSettings', excluding properties defined in parent classes.
type NotificationAggregationSettingsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Show if aggregation is enabled or not. If aggregation is enabled, events will be grouped and delivered as one notification, if it is disabled - each event will be delivered as independent notification.
	Enabled *bool `json:"Enabled,omitempty"`
	// Limit of events in one group. If this limit is reached, a window will be closed regardless of the time limit.
	EventsLimit *int64 `json:"EventsLimit,omitempty"`
	// Time in seconds which means the max time after which the window will be closed.
	Size *string `json:"Size,omitempty"`
	// Time in seconds which means how long after the last event Intersight should wait for the next event to come. If there's no new event that matches with the same Subscription within this time, then the window will be closed.
	Step                 *string `json:"Step,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAggregationSettingsAllOf NotificationAggregationSettingsAllOf

// NewNotificationAggregationSettingsAllOf instantiates a new NotificationAggregationSettingsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAggregationSettingsAllOf(classId string, objectType string) *NotificationAggregationSettingsAllOf {
	this := NotificationAggregationSettingsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationAggregationSettingsAllOfWithDefaults instantiates a new NotificationAggregationSettingsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAggregationSettingsAllOfWithDefaults() *NotificationAggregationSettingsAllOf {
	this := NotificationAggregationSettingsAllOf{}
	var classId string = "notification.AggregationSettings"
	this.ClassId = classId
	var objectType string = "notification.AggregationSettings"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAggregationSettingsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAggregationSettingsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAggregationSettingsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAggregationSettingsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationAggregationSettingsAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationAggregationSettingsAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationAggregationSettingsAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventsLimit returns the EventsLimit field value if set, zero value otherwise.
func (o *NotificationAggregationSettingsAllOf) GetEventsLimit() int64 {
	if o == nil || o.EventsLimit == nil {
		var ret int64
		return ret
	}
	return *o.EventsLimit
}

// GetEventsLimitOk returns a tuple with the EventsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetEventsLimitOk() (*int64, bool) {
	if o == nil || o.EventsLimit == nil {
		return nil, false
	}
	return o.EventsLimit, true
}

// HasEventsLimit returns a boolean if a field has been set.
func (o *NotificationAggregationSettingsAllOf) HasEventsLimit() bool {
	if o != nil && o.EventsLimit != nil {
		return true
	}

	return false
}

// SetEventsLimit gets a reference to the given int64 and assigns it to the EventsLimit field.
func (o *NotificationAggregationSettingsAllOf) SetEventsLimit(v int64) {
	o.EventsLimit = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *NotificationAggregationSettingsAllOf) GetSize() string {
	if o == nil || o.Size == nil {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetSizeOk() (*string, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *NotificationAggregationSettingsAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *NotificationAggregationSettingsAllOf) SetSize(v string) {
	o.Size = &v
}

// GetStep returns the Step field value if set, zero value otherwise.
func (o *NotificationAggregationSettingsAllOf) GetStep() string {
	if o == nil || o.Step == nil {
		var ret string
		return ret
	}
	return *o.Step
}

// GetStepOk returns a tuple with the Step field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAggregationSettingsAllOf) GetStepOk() (*string, bool) {
	if o == nil || o.Step == nil {
		return nil, false
	}
	return o.Step, true
}

// HasStep returns a boolean if a field has been set.
func (o *NotificationAggregationSettingsAllOf) HasStep() bool {
	if o != nil && o.Step != nil {
		return true
	}

	return false
}

// SetStep gets a reference to the given string and assigns it to the Step field.
func (o *NotificationAggregationSettingsAllOf) SetStep(v string) {
	o.Step = &v
}

func (o NotificationAggregationSettingsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.EventsLimit != nil {
		toSerialize["EventsLimit"] = o.EventsLimit
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Step != nil {
		toSerialize["Step"] = o.Step
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAggregationSettingsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationAggregationSettingsAllOf := _NotificationAggregationSettingsAllOf{}

	if err = json.Unmarshal(bytes, &varNotificationAggregationSettingsAllOf); err == nil {
		*o = NotificationAggregationSettingsAllOf(varNotificationAggregationSettingsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "EventsLimit")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "Step")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationAggregationSettingsAllOf struct {
	value *NotificationAggregationSettingsAllOf
	isSet bool
}

func (v NullableNotificationAggregationSettingsAllOf) Get() *NotificationAggregationSettingsAllOf {
	return v.value
}

func (v *NullableNotificationAggregationSettingsAllOf) Set(val *NotificationAggregationSettingsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAggregationSettingsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAggregationSettingsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAggregationSettingsAllOf(val *NotificationAggregationSettingsAllOf) *NullableNotificationAggregationSettingsAllOf {
	return &NullableNotificationAggregationSettingsAllOf{value: val, isSet: true}
}

func (v NullableNotificationAggregationSettingsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAggregationSettingsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
