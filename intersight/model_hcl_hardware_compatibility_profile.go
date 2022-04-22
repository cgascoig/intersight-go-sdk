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

// HclHardwareCompatibilityProfile Profile giving server hardware details, OS details and UCS Version details.
type HclHardwareCompatibilityProfile struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Url for the ISO with the drivers supported for the server.
	DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
	// Error code indicating the compatibility status. * `Success` - The input combination is valid. * `Unknown` - Unknown API request to the service. * `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database. * `InvalidUcsVersion` - UCS Version is not in expected format. * `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * `OSUnknown` - OS vendor or version is not known as per the HCL database. * `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported. * `ProductUnknown` - Product is not known as per the HCL database. * `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version. * `DriverNameNotSupported` - Driver protocol or name is not supported for the given product. * `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination. * `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination. * `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * `InternalError` - API requests to the service have either failed or timed out. * `MarshallingError` - Error in JSON marshalling. * `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// Identifier of the hardware compatibility profile.
	Id *string `json:"Id,omitempty"`
	// Vendor of the Operating System running on the server.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating System running on the server.
	OsVersion *string `json:"OsVersion,omitempty"`
	// Personality indicating the personality with the which the sever is used in a Hyperflex environment.
	Personality *string `json:"Personality,omitempty"`
	// Model of the processor present in the server.
	ProcessorModel *string      `json:"ProcessorModel,omitempty"`
	Products       []HclProduct `json:"Products,omitempty"`
	// Model of the server as returned by UCSM/CIMC XML API.
	ServerModel *string `json:"ServerModel,omitempty"`
	// Revision of the server model.
	ServerRevision *string `json:"ServerRevision,omitempty"`
	// Version of the UCS software.
	UcsVersion *string `json:"UcsVersion,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * `UCSM` - The server is managed by UCS Manager. * `IMC` - The server is standalone managed by CIMC.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclHardwareCompatibilityProfile HclHardwareCompatibilityProfile

// NewHclHardwareCompatibilityProfile instantiates a new HclHardwareCompatibilityProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHardwareCompatibilityProfile(classId string, objectType string) *HclHardwareCompatibilityProfile {
	this := HclHardwareCompatibilityProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// NewHclHardwareCompatibilityProfileWithDefaults instantiates a new HclHardwareCompatibilityProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHardwareCompatibilityProfileWithDefaults() *HclHardwareCompatibilityProfile {
	this := HclHardwareCompatibilityProfile{}
	var classId string = "hcl.HardwareCompatibilityProfile"
	this.ClassId = classId
	var objectType string = "hcl.HardwareCompatibilityProfile"
	this.ObjectType = objectType
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclHardwareCompatibilityProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclHardwareCompatibilityProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HclHardwareCompatibilityProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclHardwareCompatibilityProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriverIsoUrl returns the DriverIsoUrl field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrl() string {
	if o == nil || o.DriverIsoUrl == nil {
		var ret string
		return ret
	}
	return *o.DriverIsoUrl
}

// GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetDriverIsoUrlOk() (*string, bool) {
	if o == nil || o.DriverIsoUrl == nil {
		return nil, false
	}
	return o.DriverIsoUrl, true
}

// HasDriverIsoUrl returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasDriverIsoUrl() bool {
	if o != nil && o.DriverIsoUrl != nil {
		return true
	}

	return false
}

// SetDriverIsoUrl gets a reference to the given string and assigns it to the DriverIsoUrl field.
func (o *HclHardwareCompatibilityProfile) SetDriverIsoUrl(v string) {
	o.DriverIsoUrl = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclHardwareCompatibilityProfile) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclHardwareCompatibilityProfile) SetId(v string) {
	o.Id = &v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclHardwareCompatibilityProfile) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclHardwareCompatibilityProfile) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPersonality returns the Personality field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetPersonality() string {
	if o == nil || o.Personality == nil {
		var ret string
		return ret
	}
	return *o.Personality
}

// GetPersonalityOk returns a tuple with the Personality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetPersonalityOk() (*string, bool) {
	if o == nil || o.Personality == nil {
		return nil, false
	}
	return o.Personality, true
}

// HasPersonality returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasPersonality() bool {
	if o != nil && o.Personality != nil {
		return true
	}

	return false
}

// SetPersonality gets a reference to the given string and assigns it to the Personality field.
func (o *HclHardwareCompatibilityProfile) SetPersonality(v string) {
	o.Personality = &v
}

// GetProcessorModel returns the ProcessorModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetProcessorModel() string {
	if o == nil || o.ProcessorModel == nil {
		var ret string
		return ret
	}
	return *o.ProcessorModel
}

// GetProcessorModelOk returns a tuple with the ProcessorModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetProcessorModelOk() (*string, bool) {
	if o == nil || o.ProcessorModel == nil {
		return nil, false
	}
	return o.ProcessorModel, true
}

// HasProcessorModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasProcessorModel() bool {
	if o != nil && o.ProcessorModel != nil {
		return true
	}

	return false
}

// SetProcessorModel gets a reference to the given string and assigns it to the ProcessorModel field.
func (o *HclHardwareCompatibilityProfile) SetProcessorModel(v string) {
	o.ProcessorModel = &v
}

// GetProducts returns the Products field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclHardwareCompatibilityProfile) GetProducts() []HclProduct {
	if o == nil {
		var ret []HclProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclHardwareCompatibilityProfile) GetProductsOk() (*[]HclProduct, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return &o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []HclProduct and assigns it to the Products field.
func (o *HclHardwareCompatibilityProfile) SetProducts(v []HclProduct) {
	o.Products = v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HclHardwareCompatibilityProfile) SetServerModel(v string) {
	o.ServerModel = &v
}

// GetServerRevision returns the ServerRevision field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetServerRevision() string {
	if o == nil || o.ServerRevision == nil {
		var ret string
		return ret
	}
	return *o.ServerRevision
}

// GetServerRevisionOk returns a tuple with the ServerRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetServerRevisionOk() (*string, bool) {
	if o == nil || o.ServerRevision == nil {
		return nil, false
	}
	return o.ServerRevision, true
}

// HasServerRevision returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasServerRevision() bool {
	if o != nil && o.ServerRevision != nil {
		return true
	}

	return false
}

// SetServerRevision gets a reference to the given string and assigns it to the ServerRevision field.
func (o *HclHardwareCompatibilityProfile) SetServerRevision(v string) {
	o.ServerRevision = &v
}

// GetUcsVersion returns the UcsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetUcsVersion() string {
	if o == nil || o.UcsVersion == nil {
		var ret string
		return ret
	}
	return *o.UcsVersion
}

// GetUcsVersionOk returns a tuple with the UcsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetUcsVersionOk() (*string, bool) {
	if o == nil || o.UcsVersion == nil {
		return nil, false
	}
	return o.UcsVersion, true
}

// HasUcsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasUcsVersion() bool {
	if o != nil && o.UcsVersion != nil {
		return true
	}

	return false
}

// SetUcsVersion gets a reference to the given string and assigns it to the UcsVersion field.
func (o *HclHardwareCompatibilityProfile) SetUcsVersion(v string) {
	o.UcsVersion = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfile) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfile) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfile) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *HclHardwareCompatibilityProfile) SetVersionType(v string) {
	o.VersionType = &v
}

func (o HclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
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
	if o.DriverIsoUrl != nil {
		toSerialize["DriverIsoUrl"] = o.DriverIsoUrl
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.Personality != nil {
		toSerialize["Personality"] = o.Personality
	}
	if o.ProcessorModel != nil {
		toSerialize["ProcessorModel"] = o.ProcessorModel
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	if o.ServerRevision != nil {
		toSerialize["ServerRevision"] = o.ServerRevision
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

func (o *HclHardwareCompatibilityProfile) UnmarshalJSON(bytes []byte) (err error) {
	type HclHardwareCompatibilityProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Url for the ISO with the drivers supported for the server.
		DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
		// Error code indicating the compatibility status. * `Success` - The input combination is valid. * `Unknown` - Unknown API request to the service. * `UnknownServer` - An invalid server model is given API requests or the server model is not present in the HCL database. * `InvalidUcsVersion` - UCS Version is not in expected format. * `ProcessorNotSupported` - Processor is not supported with the given Server or the Processor does not exist in the HCL database. * `OSNotSupported` - OS version is not supported with the given server, processor combination or OS information is not present in the HCL database. * `OSUnknown` - OS vendor or version is not known as per the HCL database. * `UCSVersionNotSupported` - UCS Version is not supported with the given server, processor and OS combination or the UCS version is not present in the HCL database. * `UcsVersionServerOSCombinationNotSupported` - Combination of UCS version, server (model and processor) and os version is not supported. * `ProductUnknown` - Product is not known as per the HCL database. * `ProductNotSupported` - Product is not supported in the given UCS version, server (model and processor) and operating system version. * `DriverNameNotSupported` - Driver protocol or name is not supported for the given product. * `FirmwareVersionNotSupported` - Firmware version not supported for the component and the server, operating system combination. * `DriverVersionNotSupported` - Driver version not supported for the component and the server, operating system combination. * `FirmwareVersionDriverVersionCombinationNotSupported` - Both Firmware and Driver versions are not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverVersionNotSupported` - Firmware and Driver version combination not supported for the component and the server, operating system combination. * `FirmwareVersionAndDriverNameNotSupported` - Firmware Version and Driver name or not supported with the component and the server, operating system combination. * `InternalError` - API requests to the service have either failed or timed out. * `MarshallingError` - Error in JSON marshalling. * `Exempted` - An exempted error code means that the product is part of the exempted Catalog and should be ignored for HCL validation purposes.
		ErrorCode *string `json:"ErrorCode,omitempty"`
		// Identifier of the hardware compatibility profile.
		Id *string `json:"Id,omitempty"`
		// Vendor of the Operating System running on the server.
		OsVendor *string `json:"OsVendor,omitempty"`
		// Version of the Operating System running on the server.
		OsVersion *string `json:"OsVersion,omitempty"`
		// Personality indicating the personality with the which the sever is used in a Hyperflex environment.
		Personality *string `json:"Personality,omitempty"`
		// Model of the processor present in the server.
		ProcessorModel *string      `json:"ProcessorModel,omitempty"`
		Products       []HclProduct `json:"Products,omitempty"`
		// Model of the server as returned by UCSM/CIMC XML API.
		ServerModel *string `json:"ServerModel,omitempty"`
		// Revision of the server model.
		ServerRevision *string `json:"ServerRevision,omitempty"`
		// Version of the UCS software.
		UcsVersion *string `json:"UcsVersion,omitempty"`
		// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release. * `UCSM` - The server is managed by UCS Manager. * `IMC` - The server is standalone managed by CIMC.
		VersionType *string `json:"VersionType,omitempty"`
	}

	varHclHardwareCompatibilityProfileWithoutEmbeddedStruct := HclHardwareCompatibilityProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHclHardwareCompatibilityProfileWithoutEmbeddedStruct)
	if err == nil {
		varHclHardwareCompatibilityProfile := _HclHardwareCompatibilityProfile{}
		varHclHardwareCompatibilityProfile.ClassId = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ClassId
		varHclHardwareCompatibilityProfile.ObjectType = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ObjectType
		varHclHardwareCompatibilityProfile.DriverIsoUrl = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.DriverIsoUrl
		varHclHardwareCompatibilityProfile.ErrorCode = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ErrorCode
		varHclHardwareCompatibilityProfile.Id = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.Id
		varHclHardwareCompatibilityProfile.OsVendor = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.OsVendor
		varHclHardwareCompatibilityProfile.OsVersion = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.OsVersion
		varHclHardwareCompatibilityProfile.Personality = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.Personality
		varHclHardwareCompatibilityProfile.ProcessorModel = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ProcessorModel
		varHclHardwareCompatibilityProfile.Products = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.Products
		varHclHardwareCompatibilityProfile.ServerModel = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ServerModel
		varHclHardwareCompatibilityProfile.ServerRevision = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.ServerRevision
		varHclHardwareCompatibilityProfile.UcsVersion = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.UcsVersion
		varHclHardwareCompatibilityProfile.VersionType = varHclHardwareCompatibilityProfileWithoutEmbeddedStruct.VersionType
		*o = HclHardwareCompatibilityProfile(varHclHardwareCompatibilityProfile)
	} else {
		return err
	}

	varHclHardwareCompatibilityProfile := _HclHardwareCompatibilityProfile{}

	err = json.Unmarshal(bytes, &varHclHardwareCompatibilityProfile)
	if err == nil {
		o.MoBaseComplexType = varHclHardwareCompatibilityProfile.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriverIsoUrl")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "Personality")
		delete(additionalProperties, "ProcessorModel")
		delete(additionalProperties, "Products")
		delete(additionalProperties, "ServerModel")
		delete(additionalProperties, "ServerRevision")
		delete(additionalProperties, "UcsVersion")
		delete(additionalProperties, "VersionType")

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

type NullableHclHardwareCompatibilityProfile struct {
	value *HclHardwareCompatibilityProfile
	isSet bool
}

func (v NullableHclHardwareCompatibilityProfile) Get() *HclHardwareCompatibilityProfile {
	return v.value
}

func (v *NullableHclHardwareCompatibilityProfile) Set(val *HclHardwareCompatibilityProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHardwareCompatibilityProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHardwareCompatibilityProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHardwareCompatibilityProfile(val *HclHardwareCompatibilityProfile) *NullableHclHardwareCompatibilityProfile {
	return &NullableHclHardwareCompatibilityProfile{value: val, isSet: true}
}

func (v NullableHclHardwareCompatibilityProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHardwareCompatibilityProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
