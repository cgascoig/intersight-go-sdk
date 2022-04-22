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

// ServerProfileTemplateAllOf Definition of the list of properties defined in 'server.ProfileTemplate', excluding properties defined in parent classes.
type ServerProfileTemplateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The count of the server profiles derived from the template.
	Usage                *int64                                `json:"Usage,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServerProfileTemplateAllOf ServerProfileTemplateAllOf

// NewServerProfileTemplateAllOf instantiates a new ServerProfileTemplateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProfileTemplateAllOf(classId string, objectType string) *ServerProfileTemplateAllOf {
	this := ServerProfileTemplateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServerProfileTemplateAllOfWithDefaults instantiates a new ServerProfileTemplateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProfileTemplateAllOfWithDefaults() *ServerProfileTemplateAllOf {
	this := ServerProfileTemplateAllOf{}
	var classId string = "server.ProfileTemplate"
	this.ClassId = classId
	var objectType string = "server.ProfileTemplate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServerProfileTemplateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServerProfileTemplateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ServerProfileTemplateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServerProfileTemplateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *ServerProfileTemplateAllOf) GetUsage() int64 {
	if o == nil || o.Usage == nil {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplateAllOf) GetUsageOk() (*int64, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *ServerProfileTemplateAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *ServerProfileTemplateAllOf) SetUsage(v int64) {
	o.Usage = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ServerProfileTemplateAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProfileTemplateAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ServerProfileTemplateAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ServerProfileTemplateAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ServerProfileTemplateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServerProfileTemplateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServerProfileTemplateAllOf := _ServerProfileTemplateAllOf{}

	if err = json.Unmarshal(bytes, &varServerProfileTemplateAllOf); err == nil {
		*o = ServerProfileTemplateAllOf(varServerProfileTemplateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServerProfileTemplateAllOf struct {
	value *ServerProfileTemplateAllOf
	isSet bool
}

func (v NullableServerProfileTemplateAllOf) Get() *ServerProfileTemplateAllOf {
	return v.value
}

func (v *NullableServerProfileTemplateAllOf) Set(val *ServerProfileTemplateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProfileTemplateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProfileTemplateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProfileTemplateAllOf(val *ServerProfileTemplateAllOf) *NullableServerProfileTemplateAllOf {
	return &NullableServerProfileTemplateAllOf{value: val, isSet: true}
}

func (v NullableServerProfileTemplateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProfileTemplateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
