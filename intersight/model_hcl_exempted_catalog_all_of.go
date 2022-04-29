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

// HclExemptedCatalogAllOf Definition of the list of properties defined in 'hcl.ExemptedCatalog', excluding properties defined in parent classes.
type HclExemptedCatalogAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Reason for the exemption.
	Comments *string `json:"Comments,omitempty"`
	// A unique descriptive name of the exemption.
	Name *string `json:"Name,omitempty"`
	// Vendor of the Operating System.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating system.
	OsVersion *string `json:"OsVersion,omitempty"`
	// Name of the processor supported for the server.
	ProcessorName *string  `json:"ProcessorName,omitempty"`
	ProductModels []string `json:"ProductModels,omitempty"`
	// Type of the product/adapter say GPU for graphic cards. * `` - Default type of the Product. * `Adapter` - Represents network adapter cards. * `StorageController` - Represents storage controllers. * `GPU` - Represents graphics cards.
	ProductType *string `json:"ProductType,omitempty"`
	// Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
	ServerPid *string `json:"ServerPid,omitempty"`
	// Version of the UCS software.
	UcsVersion *string `json:"UcsVersion,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclExemptedCatalogAllOf HclExemptedCatalogAllOf

// NewHclExemptedCatalogAllOf instantiates a new HclExemptedCatalogAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclExemptedCatalogAllOf(classId string, objectType string) *HclExemptedCatalogAllOf {
	this := HclExemptedCatalogAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var productType string = ""
	this.ProductType = &productType
	return &this
}

// NewHclExemptedCatalogAllOfWithDefaults instantiates a new HclExemptedCatalogAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclExemptedCatalogAllOfWithDefaults() *HclExemptedCatalogAllOf {
	this := HclExemptedCatalogAllOf{}
	var classId string = "hcl.ExemptedCatalog"
	this.ClassId = classId
	var objectType string = "hcl.ExemptedCatalog"
	this.ObjectType = objectType
	var productType string = ""
	this.ProductType = &productType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclExemptedCatalogAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclExemptedCatalogAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclExemptedCatalogAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclExemptedCatalogAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetComments() string {
	if o == nil || o.Comments == nil {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetCommentsOk() (*string, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasComments() bool {
	if o != nil && o.Comments != nil {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *HclExemptedCatalogAllOf) SetComments(v string) {
	o.Comments = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HclExemptedCatalogAllOf) SetName(v string) {
	o.Name = &v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclExemptedCatalogAllOf) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclExemptedCatalogAllOf) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProcessorName returns the ProcessorName field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetProcessorName() string {
	if o == nil || o.ProcessorName == nil {
		var ret string
		return ret
	}
	return *o.ProcessorName
}

// GetProcessorNameOk returns a tuple with the ProcessorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetProcessorNameOk() (*string, bool) {
	if o == nil || o.ProcessorName == nil {
		return nil, false
	}
	return o.ProcessorName, true
}

// HasProcessorName returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasProcessorName() bool {
	if o != nil && o.ProcessorName != nil {
		return true
	}

	return false
}

// SetProcessorName gets a reference to the given string and assigns it to the ProcessorName field.
func (o *HclExemptedCatalogAllOf) SetProcessorName(v string) {
	o.ProcessorName = &v
}

// GetProductModels returns the ProductModels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclExemptedCatalogAllOf) GetProductModels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductModels
}

// GetProductModelsOk returns a tuple with the ProductModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclExemptedCatalogAllOf) GetProductModelsOk() (*[]string, bool) {
	if o == nil || o.ProductModels == nil {
		return nil, false
	}
	return &o.ProductModels, true
}

// HasProductModels returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasProductModels() bool {
	if o != nil && o.ProductModels != nil {
		return true
	}

	return false
}

// SetProductModels gets a reference to the given []string and assigns it to the ProductModels field.
func (o *HclExemptedCatalogAllOf) SetProductModels(v []string) {
	o.ProductModels = v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetProductType() string {
	if o == nil || o.ProductType == nil {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetProductTypeOk() (*string, bool) {
	if o == nil || o.ProductType == nil {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasProductType() bool {
	if o != nil && o.ProductType != nil {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *HclExemptedCatalogAllOf) SetProductType(v string) {
	o.ProductType = &v
}

// GetServerPid returns the ServerPid field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetServerPid() string {
	if o == nil || o.ServerPid == nil {
		var ret string
		return ret
	}
	return *o.ServerPid
}

// GetServerPidOk returns a tuple with the ServerPid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetServerPidOk() (*string, bool) {
	if o == nil || o.ServerPid == nil {
		return nil, false
	}
	return o.ServerPid, true
}

// HasServerPid returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasServerPid() bool {
	if o != nil && o.ServerPid != nil {
		return true
	}

	return false
}

// SetServerPid gets a reference to the given string and assigns it to the ServerPid field.
func (o *HclExemptedCatalogAllOf) SetServerPid(v string) {
	o.ServerPid = &v
}

// GetUcsVersion returns the UcsVersion field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetUcsVersion() string {
	if o == nil || o.UcsVersion == nil {
		var ret string
		return ret
	}
	return *o.UcsVersion
}

// GetUcsVersionOk returns a tuple with the UcsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetUcsVersionOk() (*string, bool) {
	if o == nil || o.UcsVersion == nil {
		return nil, false
	}
	return o.UcsVersion, true
}

// HasUcsVersion returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasUcsVersion() bool {
	if o != nil && o.UcsVersion != nil {
		return true
	}

	return false
}

// SetUcsVersion gets a reference to the given string and assigns it to the UcsVersion field.
func (o *HclExemptedCatalogAllOf) SetUcsVersion(v string) {
	o.UcsVersion = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *HclExemptedCatalogAllOf) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclExemptedCatalogAllOf) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *HclExemptedCatalogAllOf) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *HclExemptedCatalogAllOf) SetVersionType(v string) {
	o.VersionType = &v
}

func (o HclExemptedCatalogAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Comments != nil {
		toSerialize["Comments"] = o.Comments
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.ProcessorName != nil {
		toSerialize["ProcessorName"] = o.ProcessorName
	}
	if o.ProductModels != nil {
		toSerialize["ProductModels"] = o.ProductModels
	}
	if o.ProductType != nil {
		toSerialize["ProductType"] = o.ProductType
	}
	if o.ServerPid != nil {
		toSerialize["ServerPid"] = o.ServerPid
	}
	if o.UcsVersion != nil {
		toSerialize["UcsVersion"] = o.UcsVersion
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclExemptedCatalogAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclExemptedCatalogAllOf := _HclExemptedCatalogAllOf{}

	if err = json.Unmarshal(bytes, &varHclExemptedCatalogAllOf); err == nil {
		*o = HclExemptedCatalogAllOf(varHclExemptedCatalogAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "ProcessorName")
		delete(additionalProperties, "ProductModels")
		delete(additionalProperties, "ProductType")
		delete(additionalProperties, "ServerPid")
		delete(additionalProperties, "UcsVersion")
		delete(additionalProperties, "VersionType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclExemptedCatalogAllOf struct {
	value *HclExemptedCatalogAllOf
	isSet bool
}

func (v NullableHclExemptedCatalogAllOf) Get() *HclExemptedCatalogAllOf {
	return v.value
}

func (v *NullableHclExemptedCatalogAllOf) Set(val *HclExemptedCatalogAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclExemptedCatalogAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclExemptedCatalogAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclExemptedCatalogAllOf(val *HclExemptedCatalogAllOf) *NullableHclExemptedCatalogAllOf {
	return &NullableHclExemptedCatalogAllOf{value: val, isSet: true}
}

func (v NullableHclExemptedCatalogAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclExemptedCatalogAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
