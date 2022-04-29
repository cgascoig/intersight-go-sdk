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
)

// StorageNetAppNodeAllOf Definition of the list of properties defined in 'storage.NetAppNode', excluding properties defined in parent classes.
type StorageNetAppNodeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType            string                                  `json:"ObjectType"`
	AvgPerformanceMetrics *StorageNetAppPerformanceMetricsAverage `json:"AvgPerformanceMetrics,omitempty"`
	// Storage node option for cdpd state. * `unknown` - The cdpd option is unknown on the node. * `on` - The cdpd option is enabled on the node. * `off` - The cdpd option is disabled on the node.
	CdpdEnabled *string `json:"CdpdEnabled,omitempty"`
	// The health of the NetApp Node.
	Health           *bool                                 `json:"Health,omitempty"`
	HighAvailability NullableStorageNetAppHighAvailability `json:"HighAvailability,omitempty"`
	// Unique identifier of NetApp Node across data center.
	Key *string `json:"Key,omitempty"`
	// The system id of the NetApp Node.
	Systemid *string `json:"Systemid,omitempty"`
	// Universally unique identifier of NetApp Node.
	Uuid  *string                           `json:"Uuid,omitempty"`
	Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	// An array of relationships to storageNetAppNodeEvent resources.
	Events               []StorageNetAppNodeEventRelationship `json:"Events,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNodeAllOf StorageNetAppNodeAllOf

// NewStorageNetAppNodeAllOf instantiates a new StorageNetAppNodeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNodeAllOf(classId string, objectType string) *StorageNetAppNodeAllOf {
	this := StorageNetAppNodeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNodeAllOfWithDefaults instantiates a new StorageNetAppNodeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNodeAllOfWithDefaults() *StorageNetAppNodeAllOf {
	this := StorageNetAppNodeAllOf{}
	var classId string = "storage.NetAppNode"
	this.ClassId = classId
	var objectType string = "storage.NetAppNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNodeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNodeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNodeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNodeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvgPerformanceMetrics returns the AvgPerformanceMetrics field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetAvgPerformanceMetrics() StorageNetAppPerformanceMetricsAverage {
	if o == nil || o.AvgPerformanceMetrics == nil {
		var ret StorageNetAppPerformanceMetricsAverage
		return ret
	}
	return *o.AvgPerformanceMetrics
}

// GetAvgPerformanceMetricsOk returns a tuple with the AvgPerformanceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetAvgPerformanceMetricsOk() (*StorageNetAppPerformanceMetricsAverage, bool) {
	if o == nil || o.AvgPerformanceMetrics == nil {
		return nil, false
	}
	return o.AvgPerformanceMetrics, true
}

// HasAvgPerformanceMetrics returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasAvgPerformanceMetrics() bool {
	if o != nil && o.AvgPerformanceMetrics != nil {
		return true
	}

	return false
}

// SetAvgPerformanceMetrics gets a reference to the given StorageNetAppPerformanceMetricsAverage and assigns it to the AvgPerformanceMetrics field.
func (o *StorageNetAppNodeAllOf) SetAvgPerformanceMetrics(v StorageNetAppPerformanceMetricsAverage) {
	o.AvgPerformanceMetrics = &v
}

// GetCdpdEnabled returns the CdpdEnabled field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetCdpdEnabled() string {
	if o == nil || o.CdpdEnabled == nil {
		var ret string
		return ret
	}
	return *o.CdpdEnabled
}

// GetCdpdEnabledOk returns a tuple with the CdpdEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetCdpdEnabledOk() (*string, bool) {
	if o == nil || o.CdpdEnabled == nil {
		return nil, false
	}
	return o.CdpdEnabled, true
}

// HasCdpdEnabled returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasCdpdEnabled() bool {
	if o != nil && o.CdpdEnabled != nil {
		return true
	}

	return false
}

// SetCdpdEnabled gets a reference to the given string and assigns it to the CdpdEnabled field.
func (o *StorageNetAppNodeAllOf) SetCdpdEnabled(v string) {
	o.CdpdEnabled = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetHealth() bool {
	if o == nil || o.Health == nil {
		var ret bool
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetHealthOk() (*bool, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given bool and assigns it to the Health field.
func (o *StorageNetAppNodeAllOf) SetHealth(v bool) {
	o.Health = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNodeAllOf) GetHighAvailability() StorageNetAppHighAvailability {
	if o == nil || o.HighAvailability.Get() == nil {
		var ret StorageNetAppHighAvailability
		return ret
	}
	return *o.HighAvailability.Get()
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNodeAllOf) GetHighAvailabilityOk() (*StorageNetAppHighAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighAvailability.Get(), o.HighAvailability.IsSet()
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasHighAvailability() bool {
	if o != nil && o.HighAvailability.IsSet() {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given NullableStorageNetAppHighAvailability and assigns it to the HighAvailability field.
func (o *StorageNetAppNodeAllOf) SetHighAvailability(v StorageNetAppHighAvailability) {
	o.HighAvailability.Set(&v)
}

// SetHighAvailabilityNil sets the value for HighAvailability to be an explicit nil
func (o *StorageNetAppNodeAllOf) SetHighAvailabilityNil() {
	o.HighAvailability.Set(nil)
}

// UnsetHighAvailability ensures that no value is present for HighAvailability, not even an explicit nil
func (o *StorageNetAppNodeAllOf) UnsetHighAvailability() {
	o.HighAvailability.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *StorageNetAppNodeAllOf) SetKey(v string) {
	o.Key = &v
}

// GetSystemid returns the Systemid field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetSystemid() string {
	if o == nil || o.Systemid == nil {
		var ret string
		return ret
	}
	return *o.Systemid
}

// GetSystemidOk returns a tuple with the Systemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetSystemidOk() (*string, bool) {
	if o == nil || o.Systemid == nil {
		return nil, false
	}
	return o.Systemid, true
}

// HasSystemid returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasSystemid() bool {
	if o != nil && o.Systemid != nil {
		return true
	}

	return false
}

// SetSystemid gets a reference to the given string and assigns it to the Systemid field.
func (o *StorageNetAppNodeAllOf) SetSystemid(v string) {
	o.Systemid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppNodeAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppNodeAllOf) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNodeAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNodeAllOf) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppNodeAllOf) GetEvents() []StorageNetAppNodeEventRelationship {
	if o == nil {
		var ret []StorageNetAppNodeEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppNodeAllOf) GetEventsOk() (*[]StorageNetAppNodeEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppNodeAllOf) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppNodeEventRelationship and assigns it to the Events field.
func (o *StorageNetAppNodeAllOf) SetEvents(v []StorageNetAppNodeEventRelationship) {
	o.Events = v
}

func (o StorageNetAppNodeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvgPerformanceMetrics != nil {
		toSerialize["AvgPerformanceMetrics"] = o.AvgPerformanceMetrics
	}
	if o.CdpdEnabled != nil {
		toSerialize["CdpdEnabled"] = o.CdpdEnabled
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.HighAvailability.IsSet() {
		toSerialize["HighAvailability"] = o.HighAvailability.Get()
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.Systemid != nil {
		toSerialize["Systemid"] = o.Systemid
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNodeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNodeAllOf := _StorageNetAppNodeAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNodeAllOf); err == nil {
		*o = StorageNetAppNodeAllOf(varStorageNetAppNodeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvgPerformanceMetrics")
		delete(additionalProperties, "CdpdEnabled")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "HighAvailability")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "Systemid")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNodeAllOf struct {
	value *StorageNetAppNodeAllOf
	isSet bool
}

func (v NullableStorageNetAppNodeAllOf) Get() *StorageNetAppNodeAllOf {
	return v.value
}

func (v *NullableStorageNetAppNodeAllOf) Set(val *StorageNetAppNodeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNodeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNodeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNodeAllOf(val *StorageNetAppNodeAllOf) *NullableStorageNetAppNodeAllOf {
	return &NullableStorageNetAppNodeAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNodeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNodeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
