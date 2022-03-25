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
	"reflect"
	"strings"
)

// CrdCustomResource Custom kubernetes resource launcher service.
type CrdCustomResource struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of custom resource or 'kind' in K8s.
	DcLauncher *string `json:"DcLauncher,omitempty"`
	// Docker image URL for public cloud DC.
	Image *string `json:"Image,omitempty"`
	// A string property called name which is used as part of a uniqueness constraint. See 'identity' specification in this MO definition.
	Name *string `json:"Name,omitempty"`
	// Namespace to launch the deployment associated with the custom resource.
	Namespace *string `json:"Namespace,omitempty"`
	// Port used for public cloud DC.
	Port       *int64                            `json:"Port,omitempty"`
	Properties []CrdCustomResourceConfigProperty `json:"Properties,omitempty"`
	// Target ID for public cloud DC.
	TargetId *string `json:"TargetId,omitempty"`
	// Target Moid for public cloud DC.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// Target type for public cloud DC.
	TargetType           *string                 `json:"TargetType,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrdCustomResource CrdCustomResource

// NewCrdCustomResource instantiates a new CrdCustomResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrdCustomResource(classId string, objectType string) *CrdCustomResource {
	this := CrdCustomResource{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCrdCustomResourceWithDefaults instantiates a new CrdCustomResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrdCustomResourceWithDefaults() *CrdCustomResource {
	this := CrdCustomResource{}
	var classId string = "crd.CustomResource"
	this.ClassId = classId
	var objectType string = "crd.CustomResource"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CrdCustomResource) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CrdCustomResource) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CrdCustomResource) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CrdCustomResource) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDcLauncher returns the DcLauncher field value if set, zero value otherwise.
func (o *CrdCustomResource) GetDcLauncher() string {
	if o == nil || o.DcLauncher == nil {
		var ret string
		return ret
	}
	return *o.DcLauncher
}

// GetDcLauncherOk returns a tuple with the DcLauncher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetDcLauncherOk() (*string, bool) {
	if o == nil || o.DcLauncher == nil {
		return nil, false
	}
	return o.DcLauncher, true
}

// HasDcLauncher returns a boolean if a field has been set.
func (o *CrdCustomResource) HasDcLauncher() bool {
	if o != nil && o.DcLauncher != nil {
		return true
	}

	return false
}

// SetDcLauncher gets a reference to the given string and assigns it to the DcLauncher field.
func (o *CrdCustomResource) SetDcLauncher(v string) {
	o.DcLauncher = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *CrdCustomResource) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *CrdCustomResource) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *CrdCustomResource) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CrdCustomResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CrdCustomResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CrdCustomResource) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CrdCustomResource) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CrdCustomResource) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CrdCustomResource) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *CrdCustomResource) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *CrdCustomResource) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *CrdCustomResource) SetPort(v int64) {
	o.Port = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrdCustomResource) GetProperties() []CrdCustomResourceConfigProperty {
	if o == nil {
		var ret []CrdCustomResourceConfigProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrdCustomResource) GetPropertiesOk() (*[]CrdCustomResourceConfigProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CrdCustomResource) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []CrdCustomResourceConfigProperty and assigns it to the Properties field.
func (o *CrdCustomResource) SetProperties(v []CrdCustomResourceConfigProperty) {
	o.Properties = v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *CrdCustomResource) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *CrdCustomResource) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *CrdCustomResource) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *CrdCustomResource) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *CrdCustomResource) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *CrdCustomResource) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *CrdCustomResource) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *CrdCustomResource) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *CrdCustomResource) SetTargetType(v string) {
	o.TargetType = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CrdCustomResource) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrdCustomResource) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CrdCustomResource) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *CrdCustomResource) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o CrdCustomResource) MarshalJSON() ([]byte, error) {
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
	if o.DcLauncher != nil {
		toSerialize["DcLauncher"] = o.DcLauncher
	}
	if o.Image != nil {
		toSerialize["Image"] = o.Image
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.TargetId != nil {
		toSerialize["TargetId"] = o.TargetId
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CrdCustomResource) UnmarshalJSON(bytes []byte) (err error) {
	type CrdCustomResourceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of custom resource or 'kind' in K8s.
		DcLauncher *string `json:"DcLauncher,omitempty"`
		// Docker image URL for public cloud DC.
		Image *string `json:"Image,omitempty"`
		// A string property called name which is used as part of a uniqueness constraint. See 'identity' specification in this MO definition.
		Name *string `json:"Name,omitempty"`
		// Namespace to launch the deployment associated with the custom resource.
		Namespace *string `json:"Namespace,omitempty"`
		// Port used for public cloud DC.
		Port       *int64                            `json:"Port,omitempty"`
		Properties []CrdCustomResourceConfigProperty `json:"Properties,omitempty"`
		// Target ID for public cloud DC.
		TargetId *string `json:"TargetId,omitempty"`
		// Target Moid for public cloud DC.
		TargetMoid *string `json:"TargetMoid,omitempty"`
		// Target type for public cloud DC.
		TargetType *string                 `json:"TargetType,omitempty"`
		Account    *IamAccountRelationship `json:"Account,omitempty"`
	}

	varCrdCustomResourceWithoutEmbeddedStruct := CrdCustomResourceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCrdCustomResourceWithoutEmbeddedStruct)
	if err == nil {
		varCrdCustomResource := _CrdCustomResource{}
		varCrdCustomResource.ClassId = varCrdCustomResourceWithoutEmbeddedStruct.ClassId
		varCrdCustomResource.ObjectType = varCrdCustomResourceWithoutEmbeddedStruct.ObjectType
		varCrdCustomResource.DcLauncher = varCrdCustomResourceWithoutEmbeddedStruct.DcLauncher
		varCrdCustomResource.Image = varCrdCustomResourceWithoutEmbeddedStruct.Image
		varCrdCustomResource.Name = varCrdCustomResourceWithoutEmbeddedStruct.Name
		varCrdCustomResource.Namespace = varCrdCustomResourceWithoutEmbeddedStruct.Namespace
		varCrdCustomResource.Port = varCrdCustomResourceWithoutEmbeddedStruct.Port
		varCrdCustomResource.Properties = varCrdCustomResourceWithoutEmbeddedStruct.Properties
		varCrdCustomResource.TargetId = varCrdCustomResourceWithoutEmbeddedStruct.TargetId
		varCrdCustomResource.TargetMoid = varCrdCustomResourceWithoutEmbeddedStruct.TargetMoid
		varCrdCustomResource.TargetType = varCrdCustomResourceWithoutEmbeddedStruct.TargetType
		varCrdCustomResource.Account = varCrdCustomResourceWithoutEmbeddedStruct.Account
		*o = CrdCustomResource(varCrdCustomResource)
	} else {
		return err
	}

	varCrdCustomResource := _CrdCustomResource{}

	err = json.Unmarshal(bytes, &varCrdCustomResource)
	if err == nil {
		o.MoBaseMo = varCrdCustomResource.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DcLauncher")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "TargetId")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Account")

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

type NullableCrdCustomResource struct {
	value *CrdCustomResource
	isSet bool
}

func (v NullableCrdCustomResource) Get() *CrdCustomResource {
	return v.value
}

func (v *NullableCrdCustomResource) Set(val *CrdCustomResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCrdCustomResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCrdCustomResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrdCustomResource(val *CrdCustomResource) *NullableCrdCustomResource {
	return &NullableCrdCustomResource{value: val, isSet: true}
}

func (v NullableCrdCustomResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrdCustomResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
