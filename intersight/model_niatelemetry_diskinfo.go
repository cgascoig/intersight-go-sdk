/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-5313
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryDiskinfo Object that carries all the fields needed for Disk usage. This determines the usage of disk capacity of a device.
type NiatelemetryDiskinfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes.
	Free *int64 `json:"Free,omitempty"`
	// Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried.
	Name *string `json:"Name,omitempty"`
	// The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition.
	Total *int64 `json:"Total,omitempty"`
	// The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes.
	Used                 *int64 `json:"Used,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryDiskinfo NiatelemetryDiskinfo

// NewNiatelemetryDiskinfo instantiates a new NiatelemetryDiskinfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryDiskinfo(classId string, objectType string) *NiatelemetryDiskinfo {
	this := NiatelemetryDiskinfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryDiskinfoWithDefaults instantiates a new NiatelemetryDiskinfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryDiskinfoWithDefaults() *NiatelemetryDiskinfo {
	this := NiatelemetryDiskinfo{}
	var classId string = "niatelemetry.Diskinfo"
	this.ClassId = classId
	var objectType string = "niatelemetry.Diskinfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryDiskinfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryDiskinfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryDiskinfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryDiskinfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *NiatelemetryDiskinfo) GetFree() int64 {
	if o == nil || o.Free == nil {
		var ret int64
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetFreeOk() (*int64, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *NiatelemetryDiskinfo) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given int64 and assigns it to the Free field.
func (o *NiatelemetryDiskinfo) SetFree(v int64) {
	o.Free = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiatelemetryDiskinfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiatelemetryDiskinfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiatelemetryDiskinfo) SetName(v string) {
	o.Name = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *NiatelemetryDiskinfo) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *NiatelemetryDiskinfo) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *NiatelemetryDiskinfo) SetTotal(v int64) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *NiatelemetryDiskinfo) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryDiskinfo) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *NiatelemetryDiskinfo) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *NiatelemetryDiskinfo) SetUsed(v int64) {
	o.Used = &v
}

func (o NiatelemetryDiskinfo) MarshalJSON() ([]byte, error) {
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
	if o.Free != nil {
		toSerialize["Free"] = o.Free
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Total != nil {
		toSerialize["Total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["Used"] = o.Used
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryDiskinfo) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryDiskinfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The free disk capacity, currently the type of this field is set to integer. This determines how much memory is free in Bytes.
		Free *int64 `json:"Free,omitempty"`
		// Disk Name used to identified the disk usage record. This determines the name of the disk partition that is inventoried.
		Name *string `json:"Name,omitempty"`
		// The total disk capacity, it should be the sum of free and used, currently the type of this field is set to integer. This determines the total memory for this partition.
		Total *int64 `json:"Total,omitempty"`
		// The used disk capacity, currently the type of this field is set to integer. This determines how much memory is used in Bytes.
		Used *int64 `json:"Used,omitempty"`
	}

	varNiatelemetryDiskinfoWithoutEmbeddedStruct := NiatelemetryDiskinfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryDiskinfoWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryDiskinfo := _NiatelemetryDiskinfo{}
		varNiatelemetryDiskinfo.ClassId = varNiatelemetryDiskinfoWithoutEmbeddedStruct.ClassId
		varNiatelemetryDiskinfo.ObjectType = varNiatelemetryDiskinfoWithoutEmbeddedStruct.ObjectType
		varNiatelemetryDiskinfo.Free = varNiatelemetryDiskinfoWithoutEmbeddedStruct.Free
		varNiatelemetryDiskinfo.Name = varNiatelemetryDiskinfoWithoutEmbeddedStruct.Name
		varNiatelemetryDiskinfo.Total = varNiatelemetryDiskinfoWithoutEmbeddedStruct.Total
		varNiatelemetryDiskinfo.Used = varNiatelemetryDiskinfoWithoutEmbeddedStruct.Used
		*o = NiatelemetryDiskinfo(varNiatelemetryDiskinfo)
	} else {
		return err
	}

	varNiatelemetryDiskinfo := _NiatelemetryDiskinfo{}

	err = json.Unmarshal(bytes, &varNiatelemetryDiskinfo)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryDiskinfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Free")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Total")
		delete(additionalProperties, "Used")

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

type NullableNiatelemetryDiskinfo struct {
	value *NiatelemetryDiskinfo
	isSet bool
}

func (v NullableNiatelemetryDiskinfo) Get() *NiatelemetryDiskinfo {
	return v.value
}

func (v *NullableNiatelemetryDiskinfo) Set(val *NiatelemetryDiskinfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryDiskinfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryDiskinfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryDiskinfo(val *NiatelemetryDiskinfo) *NullableNiatelemetryDiskinfo {
	return &NullableNiatelemetryDiskinfo{value: val, isSet: true}
}

func (v NullableNiatelemetryDiskinfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryDiskinfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
