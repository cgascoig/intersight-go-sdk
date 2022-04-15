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
	"time"
)

// TamMilestoneAllOf Definition of the list of properties defined in 'tam.Milestone', excluding properties defined in parent classes.
type TamMilestoneAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Date when the specified end-of-life milestone advisory is reached.
	Date *time.Time `json:"Date,omitempty"`
	// A description of the milestone defined by an end-of-life advisory.
	Description *string `json:"Description,omitempty"`
	// Number of days (exclusive) relative to the milestone date when the milestone is considered to be not in effect. A nagative number indicates number of days ahead of the milestone date. The default is 2147483647 (0x7FFFFFFF) which means the milestone date range's upper bound is omitted.
	EndOffset *int32 `json:"EndOffset,omitempty"`
	// Extra hint to the type of label to be used in display in addition to severity and effective date. How to use it is at UI's descretion. * `upcoming` - This end-of-life (EOL) milestone is upcoming. The label may be changed to more urgent type such as 'imminent' as time progress approaching effective date. * `imminent` - This end-of-life (EOL) milestone is imminent. There will be no label change before the effective date. * `past` - This end-of-life (EOL) milestone has past the effective date.
	LabelHint *string `json:"LabelHint,omitempty"`
	// Milestone type as defined in Cisco end-of-life (EOL) policy (https://www.cisco.com/c/en/us/products/eos-eol-policy.html) when the specified end-of-life milestone advisory is reached. * `unknown` - The type of end-of-life milestone is not defined. * `endOfSoftwareMaintenanceDate` - The last date that Cisco Engineering may release any final software maintenance releases or bug fixes. After this date, Cisco Engineering may no longer develop, repair, maintain, or test the product software and only critical security updates will be provided on this release train.  * `lastDateOfSupport` - The last date to receive service and support for the software. After this date, all support services for the software are unavailable, and the software becomes obsolete.
	MilestoneType *string `json:"MilestoneType,omitempty"`
	// A milestone defined by an end-of-life advisory.
	Name *string `json:"Name,omitempty"`
	// Number of days (inclusive) relative to the milestone date when the milestone is considered to be in effect. A nagative number indicates number of days ahead of the milestone date. The default is 0 which means the milestone take effect exactly on the same date as the specified milestone date. A negative value of -2147483648 (0x80000000) indicates the milestone date range's lower bound is omitted.
	StartOffset          *int32 `json:"StartOffset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamMilestoneAllOf TamMilestoneAllOf

// NewTamMilestoneAllOf instantiates a new TamMilestoneAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamMilestoneAllOf(classId string, objectType string) *TamMilestoneAllOf {
	this := TamMilestoneAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var endOffset int32 = 2147483647
	this.EndOffset = &endOffset
	var labelHint string = "upcoming"
	this.LabelHint = &labelHint
	var milestoneType string = "unknown"
	this.MilestoneType = &milestoneType
	var startOffset int32 = 0
	this.StartOffset = &startOffset
	return &this
}

// NewTamMilestoneAllOfWithDefaults instantiates a new TamMilestoneAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamMilestoneAllOfWithDefaults() *TamMilestoneAllOf {
	this := TamMilestoneAllOf{}
	var classId string = "tam.Milestone"
	this.ClassId = classId
	var objectType string = "tam.Milestone"
	this.ObjectType = objectType
	var endOffset int32 = 2147483647
	this.EndOffset = &endOffset
	var labelHint string = "upcoming"
	this.LabelHint = &labelHint
	var milestoneType string = "unknown"
	this.MilestoneType = &milestoneType
	var startOffset int32 = 0
	this.StartOffset = &startOffset
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamMilestoneAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamMilestoneAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamMilestoneAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamMilestoneAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *TamMilestoneAllOf) SetDate(v time.Time) {
	o.Date = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TamMilestoneAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetEndOffset returns the EndOffset field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetEndOffset() int32 {
	if o == nil || o.EndOffset == nil {
		var ret int32
		return ret
	}
	return *o.EndOffset
}

// GetEndOffsetOk returns a tuple with the EndOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetEndOffsetOk() (*int32, bool) {
	if o == nil || o.EndOffset == nil {
		return nil, false
	}
	return o.EndOffset, true
}

// HasEndOffset returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasEndOffset() bool {
	if o != nil && o.EndOffset != nil {
		return true
	}

	return false
}

// SetEndOffset gets a reference to the given int32 and assigns it to the EndOffset field.
func (o *TamMilestoneAllOf) SetEndOffset(v int32) {
	o.EndOffset = &v
}

// GetLabelHint returns the LabelHint field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetLabelHint() string {
	if o == nil || o.LabelHint == nil {
		var ret string
		return ret
	}
	return *o.LabelHint
}

// GetLabelHintOk returns a tuple with the LabelHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetLabelHintOk() (*string, bool) {
	if o == nil || o.LabelHint == nil {
		return nil, false
	}
	return o.LabelHint, true
}

// HasLabelHint returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasLabelHint() bool {
	if o != nil && o.LabelHint != nil {
		return true
	}

	return false
}

// SetLabelHint gets a reference to the given string and assigns it to the LabelHint field.
func (o *TamMilestoneAllOf) SetLabelHint(v string) {
	o.LabelHint = &v
}

// GetMilestoneType returns the MilestoneType field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetMilestoneType() string {
	if o == nil || o.MilestoneType == nil {
		var ret string
		return ret
	}
	return *o.MilestoneType
}

// GetMilestoneTypeOk returns a tuple with the MilestoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetMilestoneTypeOk() (*string, bool) {
	if o == nil || o.MilestoneType == nil {
		return nil, false
	}
	return o.MilestoneType, true
}

// HasMilestoneType returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasMilestoneType() bool {
	if o != nil && o.MilestoneType != nil {
		return true
	}

	return false
}

// SetMilestoneType gets a reference to the given string and assigns it to the MilestoneType field.
func (o *TamMilestoneAllOf) SetMilestoneType(v string) {
	o.MilestoneType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TamMilestoneAllOf) SetName(v string) {
	o.Name = &v
}

// GetStartOffset returns the StartOffset field value if set, zero value otherwise.
func (o *TamMilestoneAllOf) GetStartOffset() int32 {
	if o == nil || o.StartOffset == nil {
		var ret int32
		return ret
	}
	return *o.StartOffset
}

// GetStartOffsetOk returns a tuple with the StartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamMilestoneAllOf) GetStartOffsetOk() (*int32, bool) {
	if o == nil || o.StartOffset == nil {
		return nil, false
	}
	return o.StartOffset, true
}

// HasStartOffset returns a boolean if a field has been set.
func (o *TamMilestoneAllOf) HasStartOffset() bool {
	if o != nil && o.StartOffset != nil {
		return true
	}

	return false
}

// SetStartOffset gets a reference to the given int32 and assigns it to the StartOffset field.
func (o *TamMilestoneAllOf) SetStartOffset(v int32) {
	o.StartOffset = &v
}

func (o TamMilestoneAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Date != nil {
		toSerialize["Date"] = o.Date
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.EndOffset != nil {
		toSerialize["EndOffset"] = o.EndOffset
	}
	if o.LabelHint != nil {
		toSerialize["LabelHint"] = o.LabelHint
	}
	if o.MilestoneType != nil {
		toSerialize["MilestoneType"] = o.MilestoneType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartOffset != nil {
		toSerialize["StartOffset"] = o.StartOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamMilestoneAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTamMilestoneAllOf := _TamMilestoneAllOf{}

	if err = json.Unmarshal(bytes, &varTamMilestoneAllOf); err == nil {
		*o = TamMilestoneAllOf(varTamMilestoneAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Date")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "EndOffset")
		delete(additionalProperties, "LabelHint")
		delete(additionalProperties, "MilestoneType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StartOffset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTamMilestoneAllOf struct {
	value *TamMilestoneAllOf
	isSet bool
}

func (v NullableTamMilestoneAllOf) Get() *TamMilestoneAllOf {
	return v.value
}

func (v *NullableTamMilestoneAllOf) Set(val *TamMilestoneAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTamMilestoneAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTamMilestoneAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamMilestoneAllOf(val *TamMilestoneAllOf) *NullableTamMilestoneAllOf {
	return &NullableTamMilestoneAllOf{value: val, isSet: true}
}

func (v NullableTamMilestoneAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamMilestoneAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
