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
)

// MetaRelationshipDefinition Definitions for the relationship in a meta.
type MetaRelationshipDefinition struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// API access definition for this relationship. * `NoAccess` - The property is not accessible from the API. * `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * `ReadWrite` - The property has read/write access. * `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.
	ApiAccess *string `json:"ApiAccess,omitempty"`
	// Specifies whether the relationship is a collection.
	Collection *bool `json:"Collection,omitempty"`
	// When turned off, the peer MO is not exported when the local MO is exported.
	Export *bool `json:"Export,omitempty"`
	// When turned on, the local MO is exported when the peer is exported.
	ExportWithPeer *bool `json:"ExportWithPeer,omitempty"`
	// The name of the relationship.
	Name *string `json:"Name,omitempty"`
	// Name of relationship in peer managed object.
	PeerRelName *string `json:"PeerRelName,omitempty"`
	// When turned on, peer MO corresponding to the reference provided in relation is updated with a reference to the current MO.
	PeerSync *bool `json:"PeerSync,omitempty"`
	// Fully qualified type of the peer managed object.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetaRelationshipDefinition MetaRelationshipDefinition

// NewMetaRelationshipDefinition instantiates a new MetaRelationshipDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaRelationshipDefinition(classId string, objectType string) *MetaRelationshipDefinition {
	this := MetaRelationshipDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetaRelationshipDefinitionWithDefaults instantiates a new MetaRelationshipDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaRelationshipDefinitionWithDefaults() *MetaRelationshipDefinition {
	this := MetaRelationshipDefinition{}
	var classId string = "meta.RelationshipDefinition"
	this.ClassId = classId
	var objectType string = "meta.RelationshipDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetaRelationshipDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetaRelationshipDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MetaRelationshipDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetaRelationshipDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiAccess returns the ApiAccess field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetApiAccess() string {
	if o == nil || o.ApiAccess == nil {
		var ret string
		return ret
	}
	return *o.ApiAccess
}

// GetApiAccessOk returns a tuple with the ApiAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetApiAccessOk() (*string, bool) {
	if o == nil || o.ApiAccess == nil {
		return nil, false
	}
	return o.ApiAccess, true
}

// HasApiAccess returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasApiAccess() bool {
	if o != nil && o.ApiAccess != nil {
		return true
	}

	return false
}

// SetApiAccess gets a reference to the given string and assigns it to the ApiAccess field.
func (o *MetaRelationshipDefinition) SetApiAccess(v string) {
	o.ApiAccess = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetCollection() bool {
	if o == nil || o.Collection == nil {
		var ret bool
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetCollectionOk() (*bool, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given bool and assigns it to the Collection field.
func (o *MetaRelationshipDefinition) SetCollection(v bool) {
	o.Collection = &v
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetExport() bool {
	if o == nil || o.Export == nil {
		var ret bool
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetExportOk() (*bool, bool) {
	if o == nil || o.Export == nil {
		return nil, false
	}
	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasExport() bool {
	if o != nil && o.Export != nil {
		return true
	}

	return false
}

// SetExport gets a reference to the given bool and assigns it to the Export field.
func (o *MetaRelationshipDefinition) SetExport(v bool) {
	o.Export = &v
}

// GetExportWithPeer returns the ExportWithPeer field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetExportWithPeer() bool {
	if o == nil || o.ExportWithPeer == nil {
		var ret bool
		return ret
	}
	return *o.ExportWithPeer
}

// GetExportWithPeerOk returns a tuple with the ExportWithPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetExportWithPeerOk() (*bool, bool) {
	if o == nil || o.ExportWithPeer == nil {
		return nil, false
	}
	return o.ExportWithPeer, true
}

// HasExportWithPeer returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasExportWithPeer() bool {
	if o != nil && o.ExportWithPeer != nil {
		return true
	}

	return false
}

// SetExportWithPeer gets a reference to the given bool and assigns it to the ExportWithPeer field.
func (o *MetaRelationshipDefinition) SetExportWithPeer(v bool) {
	o.ExportWithPeer = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MetaRelationshipDefinition) SetName(v string) {
	o.Name = &v
}

// GetPeerRelName returns the PeerRelName field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetPeerRelName() string {
	if o == nil || o.PeerRelName == nil {
		var ret string
		return ret
	}
	return *o.PeerRelName
}

// GetPeerRelNameOk returns a tuple with the PeerRelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetPeerRelNameOk() (*string, bool) {
	if o == nil || o.PeerRelName == nil {
		return nil, false
	}
	return o.PeerRelName, true
}

// HasPeerRelName returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasPeerRelName() bool {
	if o != nil && o.PeerRelName != nil {
		return true
	}

	return false
}

// SetPeerRelName gets a reference to the given string and assigns it to the PeerRelName field.
func (o *MetaRelationshipDefinition) SetPeerRelName(v string) {
	o.PeerRelName = &v
}

// GetPeerSync returns the PeerSync field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetPeerSync() bool {
	if o == nil || o.PeerSync == nil {
		var ret bool
		return ret
	}
	return *o.PeerSync
}

// GetPeerSyncOk returns a tuple with the PeerSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetPeerSyncOk() (*bool, bool) {
	if o == nil || o.PeerSync == nil {
		return nil, false
	}
	return o.PeerSync, true
}

// HasPeerSync returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasPeerSync() bool {
	if o != nil && o.PeerSync != nil {
		return true
	}

	return false
}

// SetPeerSync gets a reference to the given bool and assigns it to the PeerSync field.
func (o *MetaRelationshipDefinition) SetPeerSync(v bool) {
	o.PeerSync = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetaRelationshipDefinition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaRelationshipDefinition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetaRelationshipDefinition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetaRelationshipDefinition) SetType(v string) {
	o.Type = &v
}

func (o MetaRelationshipDefinition) MarshalJSON() ([]byte, error) {
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
	if o.ApiAccess != nil {
		toSerialize["ApiAccess"] = o.ApiAccess
	}
	if o.Collection != nil {
		toSerialize["Collection"] = o.Collection
	}
	if o.Export != nil {
		toSerialize["Export"] = o.Export
	}
	if o.ExportWithPeer != nil {
		toSerialize["ExportWithPeer"] = o.ExportWithPeer
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PeerRelName != nil {
		toSerialize["PeerRelName"] = o.PeerRelName
	}
	if o.PeerSync != nil {
		toSerialize["PeerSync"] = o.PeerSync
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MetaRelationshipDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type MetaRelationshipDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// API access definition for this relationship. * `NoAccess` - The property is not accessible from the API. * `ReadOnly` - The value of the property is read-only.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa ReadOnly property. * `CreateOnly` - The value of the property can be set when the REST resource is created. It cannot be changed after object creation.An HTTP 4xx status code is returned when the user sends a POST/PUT/PATCH request that containsa CreateOnly property.CreateOnly properties are returned in the response body of HTTP GET requests. * `ReadWrite` - The property has read/write access. * `WriteOnly` - The value of the property can be set but it is never returned in the response body of supported HTTP methods.This settings is used for sensitive properties such as passwords. * `ReadOnCreate` - The value of the property is returned in the HTTP POST response body when the REST resource is created.The property is not writeable and cannot be queried through a GET request after the resource has been created.
		ApiAccess *string `json:"ApiAccess,omitempty"`
		// Specifies whether the relationship is a collection.
		Collection *bool `json:"Collection,omitempty"`
		// When turned off, the peer MO is not exported when the local MO is exported.
		Export *bool `json:"Export,omitempty"`
		// When turned on, the local MO is exported when the peer is exported.
		ExportWithPeer *bool `json:"ExportWithPeer,omitempty"`
		// The name of the relationship.
		Name *string `json:"Name,omitempty"`
		// Name of relationship in peer managed object.
		PeerRelName *string `json:"PeerRelName,omitempty"`
		// When turned on, peer MO corresponding to the reference provided in relation is updated with a reference to the current MO.
		PeerSync *bool `json:"PeerSync,omitempty"`
		// Fully qualified type of the peer managed object.
		Type *string `json:"Type,omitempty"`
	}

	varMetaRelationshipDefinitionWithoutEmbeddedStruct := MetaRelationshipDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMetaRelationshipDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varMetaRelationshipDefinition := _MetaRelationshipDefinition{}
		varMetaRelationshipDefinition.ClassId = varMetaRelationshipDefinitionWithoutEmbeddedStruct.ClassId
		varMetaRelationshipDefinition.ObjectType = varMetaRelationshipDefinitionWithoutEmbeddedStruct.ObjectType
		varMetaRelationshipDefinition.ApiAccess = varMetaRelationshipDefinitionWithoutEmbeddedStruct.ApiAccess
		varMetaRelationshipDefinition.Collection = varMetaRelationshipDefinitionWithoutEmbeddedStruct.Collection
		varMetaRelationshipDefinition.Export = varMetaRelationshipDefinitionWithoutEmbeddedStruct.Export
		varMetaRelationshipDefinition.ExportWithPeer = varMetaRelationshipDefinitionWithoutEmbeddedStruct.ExportWithPeer
		varMetaRelationshipDefinition.Name = varMetaRelationshipDefinitionWithoutEmbeddedStruct.Name
		varMetaRelationshipDefinition.PeerRelName = varMetaRelationshipDefinitionWithoutEmbeddedStruct.PeerRelName
		varMetaRelationshipDefinition.PeerSync = varMetaRelationshipDefinitionWithoutEmbeddedStruct.PeerSync
		varMetaRelationshipDefinition.Type = varMetaRelationshipDefinitionWithoutEmbeddedStruct.Type
		*o = MetaRelationshipDefinition(varMetaRelationshipDefinition)
	} else {
		return err
	}

	varMetaRelationshipDefinition := _MetaRelationshipDefinition{}

	err = json.Unmarshal(bytes, &varMetaRelationshipDefinition)
	if err == nil {
		o.MoBaseComplexType = varMetaRelationshipDefinition.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiAccess")
		delete(additionalProperties, "Collection")
		delete(additionalProperties, "Export")
		delete(additionalProperties, "ExportWithPeer")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PeerRelName")
		delete(additionalProperties, "PeerSync")
		delete(additionalProperties, "Type")

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

type NullableMetaRelationshipDefinition struct {
	value *MetaRelationshipDefinition
	isSet bool
}

func (v NullableMetaRelationshipDefinition) Get() *MetaRelationshipDefinition {
	return v.value
}

func (v *NullableMetaRelationshipDefinition) Set(val *MetaRelationshipDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaRelationshipDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaRelationshipDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaRelationshipDefinition(val *MetaRelationshipDefinition) *NullableMetaRelationshipDefinition {
	return &NullableMetaRelationshipDefinition{value: val, isSet: true}
}

func (v NullableMetaRelationshipDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaRelationshipDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
