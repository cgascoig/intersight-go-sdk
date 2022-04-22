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
)

// ConnectorpackConnectorPackUpdateAllOf Definition of the list of properties defined in 'connectorpack.ConnectorPackUpdate', excluding properties defined in parent classes.
type ConnectorpackConnectorPackUpdateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version of connector pack currently running in UCS Director.
	CurrentVersion *string `json:"CurrentVersion,omitempty"`
	// Name of the connector pack.
	Name *string `json:"Name,omitempty"`
	// Version of connector pack to be installed in the next upgrade cycle.
	NewVersion           *string `json:"NewVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorpackConnectorPackUpdateAllOf ConnectorpackConnectorPackUpdateAllOf

// NewConnectorpackConnectorPackUpdateAllOf instantiates a new ConnectorpackConnectorPackUpdateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorpackConnectorPackUpdateAllOf(classId string, objectType string) *ConnectorpackConnectorPackUpdateAllOf {
	this := ConnectorpackConnectorPackUpdateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorpackConnectorPackUpdateAllOfWithDefaults instantiates a new ConnectorpackConnectorPackUpdateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorpackConnectorPackUpdateAllOfWithDefaults() *ConnectorpackConnectorPackUpdateAllOf {
	this := ConnectorpackConnectorPackUpdateAllOf{}
	var classId string = "connectorpack.ConnectorPackUpdate"
	this.ClassId = classId
	var objectType string = "connectorpack.ConnectorPackUpdate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorpackConnectorPackUpdateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorpackConnectorPackUpdateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorpackConnectorPackUpdateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorpackConnectorPackUpdateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersion() string {
	if o == nil || o.CurrentVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetCurrentVersionOk() (*string, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasCurrentVersion() bool {
	if o != nil && o.CurrentVersion != nil {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetName(v string) {
	o.Name = &v
}

// GetNewVersion returns the NewVersion field value if set, zero value otherwise.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersion() string {
	if o == nil || o.NewVersion == nil {
		var ret string
		return ret
	}
	return *o.NewVersion
}

// GetNewVersionOk returns a tuple with the NewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) GetNewVersionOk() (*string, bool) {
	if o == nil || o.NewVersion == nil {
		return nil, false
	}
	return o.NewVersion, true
}

// HasNewVersion returns a boolean if a field has been set.
func (o *ConnectorpackConnectorPackUpdateAllOf) HasNewVersion() bool {
	if o != nil && o.NewVersion != nil {
		return true
	}

	return false
}

// SetNewVersion gets a reference to the given string and assigns it to the NewVersion field.
func (o *ConnectorpackConnectorPackUpdateAllOf) SetNewVersion(v string) {
	o.NewVersion = &v
}

func (o ConnectorpackConnectorPackUpdateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CurrentVersion != nil {
		toSerialize["CurrentVersion"] = o.CurrentVersion
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NewVersion != nil {
		toSerialize["NewVersion"] = o.NewVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorpackConnectorPackUpdateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConnectorpackConnectorPackUpdateAllOf := _ConnectorpackConnectorPackUpdateAllOf{}

	if err = json.Unmarshal(bytes, &varConnectorpackConnectorPackUpdateAllOf); err == nil {
		*o = ConnectorpackConnectorPackUpdateAllOf(varConnectorpackConnectorPackUpdateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentVersion")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NewVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectorpackConnectorPackUpdateAllOf struct {
	value *ConnectorpackConnectorPackUpdateAllOf
	isSet bool
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) Get() *ConnectorpackConnectorPackUpdateAllOf {
	return v.value
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) Set(val *ConnectorpackConnectorPackUpdateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorpackConnectorPackUpdateAllOf(val *ConnectorpackConnectorPackUpdateAllOf) *NullableConnectorpackConnectorPackUpdateAllOf {
	return &NullableConnectorpackConnectorPackUpdateAllOf{value: val, isSet: true}
}

func (v NullableConnectorpackConnectorPackUpdateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorpackConnectorPackUpdateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
