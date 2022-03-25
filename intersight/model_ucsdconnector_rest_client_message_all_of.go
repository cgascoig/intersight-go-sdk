/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5808
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// UcsdconnectorRestClientMessageAllOf Definition of the list of properties defined in 'ucsdconnector.RestClientMessage', excluding properties defined in parent classes.
type UcsdconnectorRestClientMessageAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Payload which is sent along with the request. Most applicable to POST methods.
	Body *string `json:"Body,omitempty"`
	// Headers to be passed with the HTTP rest request.
	Header interface{} `json:"Header,omitempty"`
	// REST Method, should be set to one of [HTTP.MethodGet, HTTP.MethodPost].
	Method *string `json:"Method,omitempty"`
	// REST URL endpoint to which the HTTP request is sent.
	RestUrl              *string `json:"RestUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UcsdconnectorRestClientMessageAllOf UcsdconnectorRestClientMessageAllOf

// NewUcsdconnectorRestClientMessageAllOf instantiates a new UcsdconnectorRestClientMessageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdconnectorRestClientMessageAllOf(classId string, objectType string) *UcsdconnectorRestClientMessageAllOf {
	this := UcsdconnectorRestClientMessageAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUcsdconnectorRestClientMessageAllOfWithDefaults instantiates a new UcsdconnectorRestClientMessageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdconnectorRestClientMessageAllOfWithDefaults() *UcsdconnectorRestClientMessageAllOf {
	this := UcsdconnectorRestClientMessageAllOf{}
	var classId string = "ucsdconnector.RestClientMessage"
	this.ClassId = classId
	var objectType string = "ucsdconnector.RestClientMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UcsdconnectorRestClientMessageAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UcsdconnectorRestClientMessageAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UcsdconnectorRestClientMessageAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *UcsdconnectorRestClientMessageAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UcsdconnectorRestClientMessageAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UcsdconnectorRestClientMessageAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UcsdconnectorRestClientMessageAllOf) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdconnectorRestClientMessageAllOf) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UcsdconnectorRestClientMessageAllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UcsdconnectorRestClientMessageAllOf) SetBody(v string) {
	o.Body = &v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UcsdconnectorRestClientMessageAllOf) GetHeader() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UcsdconnectorRestClientMessageAllOf) GetHeaderOk() (*interface{}, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return &o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *UcsdconnectorRestClientMessageAllOf) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given interface{} and assigns it to the Header field.
func (o *UcsdconnectorRestClientMessageAllOf) SetHeader(v interface{}) {
	o.Header = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *UcsdconnectorRestClientMessageAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdconnectorRestClientMessageAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *UcsdconnectorRestClientMessageAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *UcsdconnectorRestClientMessageAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetRestUrl returns the RestUrl field value if set, zero value otherwise.
func (o *UcsdconnectorRestClientMessageAllOf) GetRestUrl() string {
	if o == nil || o.RestUrl == nil {
		var ret string
		return ret
	}
	return *o.RestUrl
}

// GetRestUrlOk returns a tuple with the RestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdconnectorRestClientMessageAllOf) GetRestUrlOk() (*string, bool) {
	if o == nil || o.RestUrl == nil {
		return nil, false
	}
	return o.RestUrl, true
}

// HasRestUrl returns a boolean if a field has been set.
func (o *UcsdconnectorRestClientMessageAllOf) HasRestUrl() bool {
	if o != nil && o.RestUrl != nil {
		return true
	}

	return false
}

// SetRestUrl gets a reference to the given string and assigns it to the RestUrl field.
func (o *UcsdconnectorRestClientMessageAllOf) SetRestUrl(v string) {
	o.RestUrl = &v
}

func (o UcsdconnectorRestClientMessageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.Header != nil {
		toSerialize["Header"] = o.Header
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.RestUrl != nil {
		toSerialize["RestUrl"] = o.RestUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UcsdconnectorRestClientMessageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varUcsdconnectorRestClientMessageAllOf := _UcsdconnectorRestClientMessageAllOf{}

	if err = json.Unmarshal(bytes, &varUcsdconnectorRestClientMessageAllOf); err == nil {
		*o = UcsdconnectorRestClientMessageAllOf(varUcsdconnectorRestClientMessageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "Header")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "RestUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUcsdconnectorRestClientMessageAllOf struct {
	value *UcsdconnectorRestClientMessageAllOf
	isSet bool
}

func (v NullableUcsdconnectorRestClientMessageAllOf) Get() *UcsdconnectorRestClientMessageAllOf {
	return v.value
}

func (v *NullableUcsdconnectorRestClientMessageAllOf) Set(val *UcsdconnectorRestClientMessageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdconnectorRestClientMessageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdconnectorRestClientMessageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdconnectorRestClientMessageAllOf(val *UcsdconnectorRestClientMessageAllOf) *NullableUcsdconnectorRestClientMessageAllOf {
	return &NullableUcsdconnectorRestClientMessageAllOf{value: val, isSet: true}
}

func (v NullableUcsdconnectorRestClientMessageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdconnectorRestClientMessageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
