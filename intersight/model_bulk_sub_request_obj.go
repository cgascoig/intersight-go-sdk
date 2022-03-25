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

// BulkSubRequestObj The sub request object is created for every subrequest in the incoming request.
type BulkSubRequestObj struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string    `json:"ObjectType"`
	Body       *MoBaseMo `json:"Body,omitempty"`
	// The body of the sub-request in string format.
	BodyString *string `json:"BodyString,omitempty"`
	// The time at which processing of this request completed.
	ExecutionCompletionTime *string `json:"ExecutionCompletionTime,omitempty"`
	// The time at which processing of this request started.
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty"`
	// This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence.
	IsObjectPresent *bool                 `json:"IsObjectPresent,omitempty"`
	Result          NullableBulkApiResult `json:"Result,omitempty"`
	// Skip the already present objects. The value from the Request.
	SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
	// The status of the request. * `Pending` - Indicates that the request is yet to be processed. * `ObjPresenceCheckInProgress` - Indicates that the checking for object presence is in progress. * `ObjPresenceCheckInComplete` - Indicates that the request is being processed. * `ObjPresenceCheckFailed` - Indicates that the checking for object presence failed. * `Processing` - Indicates that the request is being processed. * `TimedOut` - Indicates that the request processing timed out. * `Completed` - Indicates that the request processing is complete. * `Skipped` - Indicates that the request was skipped.
	Status *string `json:"Status,omitempty"`
	// This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence.
	SystemDefinedObjectDetected *bool `json:"SystemDefinedObjectDetected,omitempty"`
	// Used with PATCH & DELETE actions. The moid of an existing object instance.
	TargetMoid *string `json:"TargetMoid,omitempty"`
	// The URI on which this bulk action is to be performed.
	Uri *string `json:"Uri,omitempty"`
	// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
	Verb                 *string                  `json:"Verb,omitempty"`
	Request              *BulkRequestRelationship `json:"Request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkSubRequestObj BulkSubRequestObj

// NewBulkSubRequestObj instantiates a new BulkSubRequestObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSubRequestObj(classId string, objectType string) *BulkSubRequestObj {
	this := BulkSubRequestObj{}
	this.ClassId = classId
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// NewBulkSubRequestObjWithDefaults instantiates a new BulkSubRequestObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSubRequestObjWithDefaults() *BulkSubRequestObj {
	this := BulkSubRequestObj{}
	var classId string = "bulk.SubRequestObj"
	this.ClassId = classId
	var objectType string = "bulk.SubRequestObj"
	this.ObjectType = objectType
	var verb string = "POST"
	this.Verb = &verb
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkSubRequestObj) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkSubRequestObj) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BulkSubRequestObj) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkSubRequestObj) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetBody() MoBaseMo {
	if o == nil || o.Body == nil {
		var ret MoBaseMo
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetBodyOk() (*MoBaseMo, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given MoBaseMo and assigns it to the Body field.
func (o *BulkSubRequestObj) SetBody(v MoBaseMo) {
	o.Body = &v
}

// GetBodyString returns the BodyString field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetBodyString() string {
	if o == nil || o.BodyString == nil {
		var ret string
		return ret
	}
	return *o.BodyString
}

// GetBodyStringOk returns a tuple with the BodyString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetBodyStringOk() (*string, bool) {
	if o == nil || o.BodyString == nil {
		return nil, false
	}
	return o.BodyString, true
}

// HasBodyString returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasBodyString() bool {
	if o != nil && o.BodyString != nil {
		return true
	}

	return false
}

// SetBodyString gets a reference to the given string and assigns it to the BodyString field.
func (o *BulkSubRequestObj) SetBodyString(v string) {
	o.BodyString = &v
}

// GetExecutionCompletionTime returns the ExecutionCompletionTime field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetExecutionCompletionTime() string {
	if o == nil || o.ExecutionCompletionTime == nil {
		var ret string
		return ret
	}
	return *o.ExecutionCompletionTime
}

// GetExecutionCompletionTimeOk returns a tuple with the ExecutionCompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetExecutionCompletionTimeOk() (*string, bool) {
	if o == nil || o.ExecutionCompletionTime == nil {
		return nil, false
	}
	return o.ExecutionCompletionTime, true
}

// HasExecutionCompletionTime returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasExecutionCompletionTime() bool {
	if o != nil && o.ExecutionCompletionTime != nil {
		return true
	}

	return false
}

// SetExecutionCompletionTime gets a reference to the given string and assigns it to the ExecutionCompletionTime field.
func (o *BulkSubRequestObj) SetExecutionCompletionTime(v string) {
	o.ExecutionCompletionTime = &v
}

// GetExecutionStartTime returns the ExecutionStartTime field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetExecutionStartTime() string {
	if o == nil || o.ExecutionStartTime == nil {
		var ret string
		return ret
	}
	return *o.ExecutionStartTime
}

// GetExecutionStartTimeOk returns a tuple with the ExecutionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetExecutionStartTimeOk() (*string, bool) {
	if o == nil || o.ExecutionStartTime == nil {
		return nil, false
	}
	return o.ExecutionStartTime, true
}

// HasExecutionStartTime returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasExecutionStartTime() bool {
	if o != nil && o.ExecutionStartTime != nil {
		return true
	}

	return false
}

// SetExecutionStartTime gets a reference to the given string and assigns it to the ExecutionStartTime field.
func (o *BulkSubRequestObj) SetExecutionStartTime(v string) {
	o.ExecutionStartTime = &v
}

// GetIsObjectPresent returns the IsObjectPresent field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetIsObjectPresent() bool {
	if o == nil || o.IsObjectPresent == nil {
		var ret bool
		return ret
	}
	return *o.IsObjectPresent
}

// GetIsObjectPresentOk returns a tuple with the IsObjectPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetIsObjectPresentOk() (*bool, bool) {
	if o == nil || o.IsObjectPresent == nil {
		return nil, false
	}
	return o.IsObjectPresent, true
}

// HasIsObjectPresent returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasIsObjectPresent() bool {
	if o != nil && o.IsObjectPresent != nil {
		return true
	}

	return false
}

// SetIsObjectPresent gets a reference to the given bool and assigns it to the IsObjectPresent field.
func (o *BulkSubRequestObj) SetIsObjectPresent(v bool) {
	o.IsObjectPresent = &v
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkSubRequestObj) GetResult() BulkApiResult {
	if o == nil || o.Result.Get() == nil {
		var ret BulkApiResult
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkSubRequestObj) GetResultOk() (*BulkApiResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableBulkApiResult and assigns it to the Result field.
func (o *BulkSubRequestObj) SetResult(v BulkApiResult) {
	o.Result.Set(&v)
}

// SetResultNil sets the value for Result to be an explicit nil
func (o *BulkSubRequestObj) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *BulkSubRequestObj) UnsetResult() {
	o.Result.Unset()
}

// GetSkipDuplicates returns the SkipDuplicates field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetSkipDuplicates() bool {
	if o == nil || o.SkipDuplicates == nil {
		var ret bool
		return ret
	}
	return *o.SkipDuplicates
}

// GetSkipDuplicatesOk returns a tuple with the SkipDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetSkipDuplicatesOk() (*bool, bool) {
	if o == nil || o.SkipDuplicates == nil {
		return nil, false
	}
	return o.SkipDuplicates, true
}

// HasSkipDuplicates returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasSkipDuplicates() bool {
	if o != nil && o.SkipDuplicates != nil {
		return true
	}

	return false
}

// SetSkipDuplicates gets a reference to the given bool and assigns it to the SkipDuplicates field.
func (o *BulkSubRequestObj) SetSkipDuplicates(v bool) {
	o.SkipDuplicates = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkSubRequestObj) SetStatus(v string) {
	o.Status = &v
}

// GetSystemDefinedObjectDetected returns the SystemDefinedObjectDetected field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetSystemDefinedObjectDetected() bool {
	if o == nil || o.SystemDefinedObjectDetected == nil {
		var ret bool
		return ret
	}
	return *o.SystemDefinedObjectDetected
}

// GetSystemDefinedObjectDetectedOk returns a tuple with the SystemDefinedObjectDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetSystemDefinedObjectDetectedOk() (*bool, bool) {
	if o == nil || o.SystemDefinedObjectDetected == nil {
		return nil, false
	}
	return o.SystemDefinedObjectDetected, true
}

// HasSystemDefinedObjectDetected returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasSystemDefinedObjectDetected() bool {
	if o != nil && o.SystemDefinedObjectDetected != nil {
		return true
	}

	return false
}

// SetSystemDefinedObjectDetected gets a reference to the given bool and assigns it to the SystemDefinedObjectDetected field.
func (o *BulkSubRequestObj) SetSystemDefinedObjectDetected(v bool) {
	o.SystemDefinedObjectDetected = &v
}

// GetTargetMoid returns the TargetMoid field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetTargetMoid() string {
	if o == nil || o.TargetMoid == nil {
		var ret string
		return ret
	}
	return *o.TargetMoid
}

// GetTargetMoidOk returns a tuple with the TargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetTargetMoidOk() (*string, bool) {
	if o == nil || o.TargetMoid == nil {
		return nil, false
	}
	return o.TargetMoid, true
}

// HasTargetMoid returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasTargetMoid() bool {
	if o != nil && o.TargetMoid != nil {
		return true
	}

	return false
}

// SetTargetMoid gets a reference to the given string and assigns it to the TargetMoid field.
func (o *BulkSubRequestObj) SetTargetMoid(v string) {
	o.TargetMoid = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkSubRequestObj) SetUri(v string) {
	o.Uri = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetVerb() string {
	if o == nil || o.Verb == nil {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetVerbOk() (*string, bool) {
	if o == nil || o.Verb == nil {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasVerb() bool {
	if o != nil && o.Verb != nil {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *BulkSubRequestObj) SetVerb(v string) {
	o.Verb = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *BulkSubRequestObj) GetRequest() BulkRequestRelationship {
	if o == nil || o.Request == nil {
		var ret BulkRequestRelationship
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSubRequestObj) GetRequestOk() (*BulkRequestRelationship, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *BulkSubRequestObj) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given BulkRequestRelationship and assigns it to the Request field.
func (o *BulkSubRequestObj) SetRequest(v BulkRequestRelationship) {
	o.Request = &v
}

func (o BulkSubRequestObj) MarshalJSON() ([]byte, error) {
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
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.BodyString != nil {
		toSerialize["BodyString"] = o.BodyString
	}
	if o.ExecutionCompletionTime != nil {
		toSerialize["ExecutionCompletionTime"] = o.ExecutionCompletionTime
	}
	if o.ExecutionStartTime != nil {
		toSerialize["ExecutionStartTime"] = o.ExecutionStartTime
	}
	if o.IsObjectPresent != nil {
		toSerialize["IsObjectPresent"] = o.IsObjectPresent
	}
	if o.Result.IsSet() {
		toSerialize["Result"] = o.Result.Get()
	}
	if o.SkipDuplicates != nil {
		toSerialize["SkipDuplicates"] = o.SkipDuplicates
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SystemDefinedObjectDetected != nil {
		toSerialize["SystemDefinedObjectDetected"] = o.SystemDefinedObjectDetected
	}
	if o.TargetMoid != nil {
		toSerialize["TargetMoid"] = o.TargetMoid
	}
	if o.Uri != nil {
		toSerialize["Uri"] = o.Uri
	}
	if o.Verb != nil {
		toSerialize["Verb"] = o.Verb
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkSubRequestObj) UnmarshalJSON(bytes []byte) (err error) {
	type BulkSubRequestObjWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string    `json:"ObjectType"`
		Body       *MoBaseMo `json:"Body,omitempty"`
		// The body of the sub-request in string format.
		BodyString *string `json:"BodyString,omitempty"`
		// The time at which processing of this request completed.
		ExecutionCompletionTime *string `json:"ExecutionCompletionTime,omitempty"`
		// The time at which processing of this request started.
		ExecutionStartTime *string `json:"ExecutionStartTime,omitempty"`
		// This flag indicates if an already existing object was found or not after execution of the action CheckObjectPresence.
		IsObjectPresent *bool                 `json:"IsObjectPresent,omitempty"`
		Result          NullableBulkApiResult `json:"Result,omitempty"`
		// Skip the already present objects. The value from the Request.
		SkipDuplicates *bool `json:"SkipDuplicates,omitempty"`
		// The status of the request. * `Pending` - Indicates that the request is yet to be processed. * `ObjPresenceCheckInProgress` - Indicates that the checking for object presence is in progress. * `ObjPresenceCheckInComplete` - Indicates that the request is being processed. * `ObjPresenceCheckFailed` - Indicates that the checking for object presence failed. * `Processing` - Indicates that the request is being processed. * `TimedOut` - Indicates that the request processing timed out. * `Completed` - Indicates that the request processing is complete. * `Skipped` - Indicates that the request was skipped.
		Status *string `json:"Status,omitempty"`
		// This flag indicates if the a system defined object was detected after execution of the action CheckObjectPresence.
		SystemDefinedObjectDetected *bool `json:"SystemDefinedObjectDetected,omitempty"`
		// Used with PATCH & DELETE actions. The moid of an existing object instance.
		TargetMoid *string `json:"TargetMoid,omitempty"`
		// The URI on which this bulk action is to be performed.
		Uri *string `json:"Uri,omitempty"`
		// The type of operation to be performed. One of - Post (Create), Patch (Update) or Delete (Remove). * `POST` - Used to create a REST resource. * `PATCH` - Used to update a REST resource. * `DELETE` - Used to delete a REST resource.
		Verb    *string                  `json:"Verb,omitempty"`
		Request *BulkRequestRelationship `json:"Request,omitempty"`
	}

	varBulkSubRequestObjWithoutEmbeddedStruct := BulkSubRequestObjWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBulkSubRequestObjWithoutEmbeddedStruct)
	if err == nil {
		varBulkSubRequestObj := _BulkSubRequestObj{}
		varBulkSubRequestObj.ClassId = varBulkSubRequestObjWithoutEmbeddedStruct.ClassId
		varBulkSubRequestObj.ObjectType = varBulkSubRequestObjWithoutEmbeddedStruct.ObjectType
		varBulkSubRequestObj.Body = varBulkSubRequestObjWithoutEmbeddedStruct.Body
		varBulkSubRequestObj.BodyString = varBulkSubRequestObjWithoutEmbeddedStruct.BodyString
		varBulkSubRequestObj.ExecutionCompletionTime = varBulkSubRequestObjWithoutEmbeddedStruct.ExecutionCompletionTime
		varBulkSubRequestObj.ExecutionStartTime = varBulkSubRequestObjWithoutEmbeddedStruct.ExecutionStartTime
		varBulkSubRequestObj.IsObjectPresent = varBulkSubRequestObjWithoutEmbeddedStruct.IsObjectPresent
		varBulkSubRequestObj.Result = varBulkSubRequestObjWithoutEmbeddedStruct.Result
		varBulkSubRequestObj.SkipDuplicates = varBulkSubRequestObjWithoutEmbeddedStruct.SkipDuplicates
		varBulkSubRequestObj.Status = varBulkSubRequestObjWithoutEmbeddedStruct.Status
		varBulkSubRequestObj.SystemDefinedObjectDetected = varBulkSubRequestObjWithoutEmbeddedStruct.SystemDefinedObjectDetected
		varBulkSubRequestObj.TargetMoid = varBulkSubRequestObjWithoutEmbeddedStruct.TargetMoid
		varBulkSubRequestObj.Uri = varBulkSubRequestObjWithoutEmbeddedStruct.Uri
		varBulkSubRequestObj.Verb = varBulkSubRequestObjWithoutEmbeddedStruct.Verb
		varBulkSubRequestObj.Request = varBulkSubRequestObjWithoutEmbeddedStruct.Request
		*o = BulkSubRequestObj(varBulkSubRequestObj)
	} else {
		return err
	}

	varBulkSubRequestObj := _BulkSubRequestObj{}

	err = json.Unmarshal(bytes, &varBulkSubRequestObj)
	if err == nil {
		o.MoBaseMo = varBulkSubRequestObj.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "BodyString")
		delete(additionalProperties, "ExecutionCompletionTime")
		delete(additionalProperties, "ExecutionStartTime")
		delete(additionalProperties, "IsObjectPresent")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "SkipDuplicates")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SystemDefinedObjectDetected")
		delete(additionalProperties, "TargetMoid")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "Verb")
		delete(additionalProperties, "Request")

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

type NullableBulkSubRequestObj struct {
	value *BulkSubRequestObj
	isSet bool
}

func (v NullableBulkSubRequestObj) Get() *BulkSubRequestObj {
	return v.value
}

func (v *NullableBulkSubRequestObj) Set(val *BulkSubRequestObj) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSubRequestObj) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSubRequestObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSubRequestObj(val *BulkSubRequestObj) *NullableBulkSubRequestObj {
	return &NullableBulkSubRequestObj{value: val, isSet: true}
}

func (v NullableBulkSubRequestObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSubRequestObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
