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

// LicenseIwoCustomerOp Customer operation object to refresh the registration or re-authenticate, pre-created.
type LicenseIwoCustomerOp struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The workload optimizer license administrative state. Set this property to 'true' to activate the workload optimizer license entitlements.
	ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
	// Active workload optimizer license tier set by user. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
	ActiveLicenseType *string `json:"ActiveLicenseType,omitempty"`
	// Enable trial for Intersight licensing.
	EnableTrial *bool `json:"EnableTrial,omitempty"`
	// The default Trial or Grace period customer is entitled to.
	EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
	// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
	ExtraEvaluation      *int64                                 `json:"ExtraEvaluation,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseIwoCustomerOp LicenseIwoCustomerOp

// NewLicenseIwoCustomerOp instantiates a new LicenseIwoCustomerOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseIwoCustomerOp(classId string, objectType string) *LicenseIwoCustomerOp {
	this := LicenseIwoCustomerOp{}
	this.ClassId = classId
	this.ObjectType = objectType
	var activeLicenseType string = "Base"
	this.ActiveLicenseType = &activeLicenseType
	return &this
}

// NewLicenseIwoCustomerOpWithDefaults instantiates a new LicenseIwoCustomerOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseIwoCustomerOpWithDefaults() *LicenseIwoCustomerOp {
	this := LicenseIwoCustomerOp{}
	var classId string = "license.IwoCustomerOp"
	this.ClassId = classId
	var objectType string = "license.IwoCustomerOp"
	this.ObjectType = objectType
	var activeLicenseType string = "Base"
	this.ActiveLicenseType = &activeLicenseType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseIwoCustomerOp) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseIwoCustomerOp) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseIwoCustomerOp) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseIwoCustomerOp) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveAdmin returns the ActiveAdmin field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetActiveAdmin() bool {
	if o == nil || o.ActiveAdmin == nil {
		var ret bool
		return ret
	}
	return *o.ActiveAdmin
}

// GetActiveAdminOk returns a tuple with the ActiveAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetActiveAdminOk() (*bool, bool) {
	if o == nil || o.ActiveAdmin == nil {
		return nil, false
	}
	return o.ActiveAdmin, true
}

// HasActiveAdmin returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasActiveAdmin() bool {
	if o != nil && o.ActiveAdmin != nil {
		return true
	}

	return false
}

// SetActiveAdmin gets a reference to the given bool and assigns it to the ActiveAdmin field.
func (o *LicenseIwoCustomerOp) SetActiveAdmin(v bool) {
	o.ActiveAdmin = &v
}

// GetActiveLicenseType returns the ActiveLicenseType field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetActiveLicenseType() string {
	if o == nil || o.ActiveLicenseType == nil {
		var ret string
		return ret
	}
	return *o.ActiveLicenseType
}

// GetActiveLicenseTypeOk returns a tuple with the ActiveLicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetActiveLicenseTypeOk() (*string, bool) {
	if o == nil || o.ActiveLicenseType == nil {
		return nil, false
	}
	return o.ActiveLicenseType, true
}

// HasActiveLicenseType returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasActiveLicenseType() bool {
	if o != nil && o.ActiveLicenseType != nil {
		return true
	}

	return false
}

// SetActiveLicenseType gets a reference to the given string and assigns it to the ActiveLicenseType field.
func (o *LicenseIwoCustomerOp) SetActiveLicenseType(v string) {
	o.ActiveLicenseType = &v
}

// GetEnableTrial returns the EnableTrial field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetEnableTrial() bool {
	if o == nil || o.EnableTrial == nil {
		var ret bool
		return ret
	}
	return *o.EnableTrial
}

// GetEnableTrialOk returns a tuple with the EnableTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetEnableTrialOk() (*bool, bool) {
	if o == nil || o.EnableTrial == nil {
		return nil, false
	}
	return o.EnableTrial, true
}

// HasEnableTrial returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasEnableTrial() bool {
	if o != nil && o.EnableTrial != nil {
		return true
	}

	return false
}

// SetEnableTrial gets a reference to the given bool and assigns it to the EnableTrial field.
func (o *LicenseIwoCustomerOp) SetEnableTrial(v bool) {
	o.EnableTrial = &v
}

// GetEvaluationPeriod returns the EvaluationPeriod field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetEvaluationPeriod() int64 {
	if o == nil || o.EvaluationPeriod == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationPeriod
}

// GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetEvaluationPeriodOk() (*int64, bool) {
	if o == nil || o.EvaluationPeriod == nil {
		return nil, false
	}
	return o.EvaluationPeriod, true
}

// HasEvaluationPeriod returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasEvaluationPeriod() bool {
	if o != nil && o.EvaluationPeriod != nil {
		return true
	}

	return false
}

// SetEvaluationPeriod gets a reference to the given int64 and assigns it to the EvaluationPeriod field.
func (o *LicenseIwoCustomerOp) SetEvaluationPeriod(v int64) {
	o.EvaluationPeriod = &v
}

// GetExtraEvaluation returns the ExtraEvaluation field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetExtraEvaluation() int64 {
	if o == nil || o.ExtraEvaluation == nil {
		var ret int64
		return ret
	}
	return *o.ExtraEvaluation
}

// GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetExtraEvaluationOk() (*int64, bool) {
	if o == nil || o.ExtraEvaluation == nil {
		return nil, false
	}
	return o.ExtraEvaluation, true
}

// HasExtraEvaluation returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasExtraEvaluation() bool {
	if o != nil && o.ExtraEvaluation != nil {
		return true
	}

	return false
}

// SetExtraEvaluation gets a reference to the given int64 and assigns it to the ExtraEvaluation field.
func (o *LicenseIwoCustomerOp) SetExtraEvaluation(v int64) {
	o.ExtraEvaluation = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseIwoCustomerOp) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIwoCustomerOp) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseIwoCustomerOp) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseIwoCustomerOp) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseIwoCustomerOp) MarshalJSON() ([]byte, error) {
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
	if o.ActiveAdmin != nil {
		toSerialize["ActiveAdmin"] = o.ActiveAdmin
	}
	if o.ActiveLicenseType != nil {
		toSerialize["ActiveLicenseType"] = o.ActiveLicenseType
	}
	if o.EnableTrial != nil {
		toSerialize["EnableTrial"] = o.EnableTrial
	}
	if o.EvaluationPeriod != nil {
		toSerialize["EvaluationPeriod"] = o.EvaluationPeriod
	}
	if o.ExtraEvaluation != nil {
		toSerialize["ExtraEvaluation"] = o.ExtraEvaluation
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseIwoCustomerOp) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseIwoCustomerOpWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The workload optimizer license administrative state. Set this property to 'true' to activate the workload optimizer license entitlements.
		ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
		// Active workload optimizer license tier set by user. * `Base` - Base as a License type. It is default license type. * `Essential` - Essential as a License type. * `Standard` - Standard as a License type. * `Advantage` - Advantage as a License type. * `Premier` - Premier as a License type. * `IWO-Essential` - IWO-Essential as a License type. * `IWO-Advantage` - IWO-Advantage as a License type. * `IWO-Premier` - IWO-Premier as a License type. * `IKS-Advantage` - IKS-Advantage as a License type.
		ActiveLicenseType *string `json:"ActiveLicenseType,omitempty"`
		// Enable trial for Intersight licensing.
		EnableTrial *bool `json:"EnableTrial,omitempty"`
		// The default Trial or Grace period customer is entitled to.
		EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
		// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
		ExtraEvaluation    *int64                                 `json:"ExtraEvaluation,omitempty"`
		AccountLicenseData *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseIwoCustomerOpWithoutEmbeddedStruct := LicenseIwoCustomerOpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseIwoCustomerOpWithoutEmbeddedStruct)
	if err == nil {
		varLicenseIwoCustomerOp := _LicenseIwoCustomerOp{}
		varLicenseIwoCustomerOp.ClassId = varLicenseIwoCustomerOpWithoutEmbeddedStruct.ClassId
		varLicenseIwoCustomerOp.ObjectType = varLicenseIwoCustomerOpWithoutEmbeddedStruct.ObjectType
		varLicenseIwoCustomerOp.ActiveAdmin = varLicenseIwoCustomerOpWithoutEmbeddedStruct.ActiveAdmin
		varLicenseIwoCustomerOp.ActiveLicenseType = varLicenseIwoCustomerOpWithoutEmbeddedStruct.ActiveLicenseType
		varLicenseIwoCustomerOp.EnableTrial = varLicenseIwoCustomerOpWithoutEmbeddedStruct.EnableTrial
		varLicenseIwoCustomerOp.EvaluationPeriod = varLicenseIwoCustomerOpWithoutEmbeddedStruct.EvaluationPeriod
		varLicenseIwoCustomerOp.ExtraEvaluation = varLicenseIwoCustomerOpWithoutEmbeddedStruct.ExtraEvaluation
		varLicenseIwoCustomerOp.AccountLicenseData = varLicenseIwoCustomerOpWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseIwoCustomerOp(varLicenseIwoCustomerOp)
	} else {
		return err
	}

	varLicenseIwoCustomerOp := _LicenseIwoCustomerOp{}

	err = json.Unmarshal(bytes, &varLicenseIwoCustomerOp)
	if err == nil {
		o.MoBaseMo = varLicenseIwoCustomerOp.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveAdmin")
		delete(additionalProperties, "ActiveLicenseType")
		delete(additionalProperties, "EnableTrial")
		delete(additionalProperties, "EvaluationPeriod")
		delete(additionalProperties, "ExtraEvaluation")
		delete(additionalProperties, "AccountLicenseData")

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

type NullableLicenseIwoCustomerOp struct {
	value *LicenseIwoCustomerOp
	isSet bool
}

func (v NullableLicenseIwoCustomerOp) Get() *LicenseIwoCustomerOp {
	return v.value
}

func (v *NullableLicenseIwoCustomerOp) Set(val *LicenseIwoCustomerOp) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIwoCustomerOp) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIwoCustomerOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIwoCustomerOp(val *LicenseIwoCustomerOp) *NullableLicenseIwoCustomerOp {
	return &NullableLicenseIwoCustomerOp{value: val, isSet: true}
}

func (v NullableLicenseIwoCustomerOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIwoCustomerOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
