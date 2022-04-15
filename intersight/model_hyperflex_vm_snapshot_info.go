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
	"reflect"
	"strings"
)

// HyperflexVmSnapshotInfo Virtual Machine Snapshot information like replication status, snapshot point and status.
type HyperflexVmSnapshotInfo struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType       string                                   `json:"ObjectType"`
	ClusterIdSnapMap []HyperflexMapClusterIdToStSnapshotPoint `json:"ClusterIdSnapMap,omitempty"`
	// Combined status code from replication and snapshot status to display in the Intersight UI. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
	DisplayStatus *string                     `json:"DisplayStatus,omitempty"`
	Error         NullableHyperflexErrorStack `json:"Error,omitempty"`
	// The name of the Virtual Machine and the time stamp of the snapshot.
	Label *string `json:"Label,omitempty"`
	// Quiesce Mode for snapshot taken on Hyperflex cluster. * `NONE` - The snapshot quiesce mode is none. * `CRASH` - The snapshot quiesce mode is crash. * `VMTOOLS` - The snapshot quiesce mode is VMTOOLS. * `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.
	Mode           *string                          `json:"Mode,omitempty"`
	ParentSnapshot NullableHyperflexEntityReference `json:"ParentSnapshot,omitempty"`
	// Replication status of the least successful target cluster. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
	ReplicationStatus *string `json:"ReplicationStatus,omitempty"`
	// Error message from snapshot creation or replcation if any exist.
	SnapshotErrorMsg *string `json:"SnapshotErrorMsg,omitempty"`
	// Snapshot status of the source cluster. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none. * `INIT` - This snapshot status code is initializing.
	SnapshotStatus *string `json:"SnapshotStatus,omitempty"`
	// Timestamp the snapshot was created on the source cluster.
	SourceTimestamp *int64 `json:"SourceTimestamp,omitempty"`
	// Name of the cluster this snapshot resides on.
	SrcClusterName *string `json:"SrcClusterName,omitempty"`
	// Timestamp the snapshot was finished replicating on the target cluster.
	TargetCompletionTimestamp *int64 `json:"TargetCompletionTimestamp,omitempty"`
	// Name of the cluster this snapshot is replicated to.
	TgtClusterName            *string                            `json:"TgtClusterName,omitempty"`
	VmEntityReference         NullableHyperflexEntityReference   `json:"VmEntityReference,omitempty"`
	VmSnapshotEntityReference NullableHyperflexEntityReference   `json:"VmSnapshotEntityReference,omitempty"`
	SrcCluster                *HyperflexClusterRelationship      `json:"SrcCluster,omitempty"`
	TgtCluster                *HyperflexClusterRelationship      `json:"TgtCluster,omitempty"`
	VmBackupInfo              *HyperflexVmBackupInfoRelationship `json:"VmBackupInfo,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _HyperflexVmSnapshotInfo HyperflexVmSnapshotInfo

// NewHyperflexVmSnapshotInfo instantiates a new HyperflexVmSnapshotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVmSnapshotInfo(classId string, objectType string) *HyperflexVmSnapshotInfo {
	this := HyperflexVmSnapshotInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVmSnapshotInfoWithDefaults instantiates a new HyperflexVmSnapshotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVmSnapshotInfoWithDefaults() *HyperflexVmSnapshotInfo {
	this := HyperflexVmSnapshotInfo{}
	var classId string = "hyperflex.VmSnapshotInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.VmSnapshotInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVmSnapshotInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVmSnapshotInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVmSnapshotInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVmSnapshotInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterIdSnapMap returns the ClusterIdSnapMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfo) GetClusterIdSnapMap() []HyperflexMapClusterIdToStSnapshotPoint {
	if o == nil {
		var ret []HyperflexMapClusterIdToStSnapshotPoint
		return ret
	}
	return o.ClusterIdSnapMap
}

// GetClusterIdSnapMapOk returns a tuple with the ClusterIdSnapMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfo) GetClusterIdSnapMapOk() (*[]HyperflexMapClusterIdToStSnapshotPoint, bool) {
	if o == nil || o.ClusterIdSnapMap == nil {
		return nil, false
	}
	return &o.ClusterIdSnapMap, true
}

// HasClusterIdSnapMap returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasClusterIdSnapMap() bool {
	if o != nil && o.ClusterIdSnapMap != nil {
		return true
	}

	return false
}

// SetClusterIdSnapMap gets a reference to the given []HyperflexMapClusterIdToStSnapshotPoint and assigns it to the ClusterIdSnapMap field.
func (o *HyperflexVmSnapshotInfo) SetClusterIdSnapMap(v []HyperflexMapClusterIdToStSnapshotPoint) {
	o.ClusterIdSnapMap = v
}

// GetDisplayStatus returns the DisplayStatus field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetDisplayStatus() string {
	if o == nil || o.DisplayStatus == nil {
		var ret string
		return ret
	}
	return *o.DisplayStatus
}

// GetDisplayStatusOk returns a tuple with the DisplayStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetDisplayStatusOk() (*string, bool) {
	if o == nil || o.DisplayStatus == nil {
		return nil, false
	}
	return o.DisplayStatus, true
}

// HasDisplayStatus returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasDisplayStatus() bool {
	if o != nil && o.DisplayStatus != nil {
		return true
	}

	return false
}

// SetDisplayStatus gets a reference to the given string and assigns it to the DisplayStatus field.
func (o *HyperflexVmSnapshotInfo) SetDisplayStatus(v string) {
	o.DisplayStatus = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfo) GetError() HyperflexErrorStack {
	if o == nil || o.Error.Get() == nil {
		var ret HyperflexErrorStack
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfo) GetErrorOk() (*HyperflexErrorStack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableHyperflexErrorStack and assigns it to the Error field.
func (o *HyperflexVmSnapshotInfo) SetError(v HyperflexErrorStack) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *HyperflexVmSnapshotInfo) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HyperflexVmSnapshotInfo) UnsetError() {
	o.Error.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *HyperflexVmSnapshotInfo) SetLabel(v string) {
	o.Label = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *HyperflexVmSnapshotInfo) SetMode(v string) {
	o.Mode = &v
}

// GetParentSnapshot returns the ParentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfo) GetParentSnapshot() HyperflexEntityReference {
	if o == nil || o.ParentSnapshot.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.ParentSnapshot.Get()
}

// GetParentSnapshotOk returns a tuple with the ParentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfo) GetParentSnapshotOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentSnapshot.Get(), o.ParentSnapshot.IsSet()
}

// HasParentSnapshot returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasParentSnapshot() bool {
	if o != nil && o.ParentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetParentSnapshot gets a reference to the given NullableHyperflexEntityReference and assigns it to the ParentSnapshot field.
func (o *HyperflexVmSnapshotInfo) SetParentSnapshot(v HyperflexEntityReference) {
	o.ParentSnapshot.Set(&v)
}

// SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil
func (o *HyperflexVmSnapshotInfo) SetParentSnapshotNil() {
	o.ParentSnapshot.Set(nil)
}

// UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
func (o *HyperflexVmSnapshotInfo) UnsetParentSnapshot() {
	o.ParentSnapshot.Unset()
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetReplicationStatus() string {
	if o == nil || o.ReplicationStatus == nil {
		var ret string
		return ret
	}
	return *o.ReplicationStatus
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetReplicationStatusOk() (*string, bool) {
	if o == nil || o.ReplicationStatus == nil {
		return nil, false
	}
	return o.ReplicationStatus, true
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasReplicationStatus() bool {
	if o != nil && o.ReplicationStatus != nil {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given string and assigns it to the ReplicationStatus field.
func (o *HyperflexVmSnapshotInfo) SetReplicationStatus(v string) {
	o.ReplicationStatus = &v
}

// GetSnapshotErrorMsg returns the SnapshotErrorMsg field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetSnapshotErrorMsg() string {
	if o == nil || o.SnapshotErrorMsg == nil {
		var ret string
		return ret
	}
	return *o.SnapshotErrorMsg
}

// GetSnapshotErrorMsgOk returns a tuple with the SnapshotErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetSnapshotErrorMsgOk() (*string, bool) {
	if o == nil || o.SnapshotErrorMsg == nil {
		return nil, false
	}
	return o.SnapshotErrorMsg, true
}

// HasSnapshotErrorMsg returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasSnapshotErrorMsg() bool {
	if o != nil && o.SnapshotErrorMsg != nil {
		return true
	}

	return false
}

// SetSnapshotErrorMsg gets a reference to the given string and assigns it to the SnapshotErrorMsg field.
func (o *HyperflexVmSnapshotInfo) SetSnapshotErrorMsg(v string) {
	o.SnapshotErrorMsg = &v
}

// GetSnapshotStatus returns the SnapshotStatus field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetSnapshotStatus() string {
	if o == nil || o.SnapshotStatus == nil {
		var ret string
		return ret
	}
	return *o.SnapshotStatus
}

// GetSnapshotStatusOk returns a tuple with the SnapshotStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetSnapshotStatusOk() (*string, bool) {
	if o == nil || o.SnapshotStatus == nil {
		return nil, false
	}
	return o.SnapshotStatus, true
}

// HasSnapshotStatus returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasSnapshotStatus() bool {
	if o != nil && o.SnapshotStatus != nil {
		return true
	}

	return false
}

// SetSnapshotStatus gets a reference to the given string and assigns it to the SnapshotStatus field.
func (o *HyperflexVmSnapshotInfo) SetSnapshotStatus(v string) {
	o.SnapshotStatus = &v
}

// GetSourceTimestamp returns the SourceTimestamp field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetSourceTimestamp() int64 {
	if o == nil || o.SourceTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.SourceTimestamp
}

// GetSourceTimestampOk returns a tuple with the SourceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetSourceTimestampOk() (*int64, bool) {
	if o == nil || o.SourceTimestamp == nil {
		return nil, false
	}
	return o.SourceTimestamp, true
}

// HasSourceTimestamp returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasSourceTimestamp() bool {
	if o != nil && o.SourceTimestamp != nil {
		return true
	}

	return false
}

// SetSourceTimestamp gets a reference to the given int64 and assigns it to the SourceTimestamp field.
func (o *HyperflexVmSnapshotInfo) SetSourceTimestamp(v int64) {
	o.SourceTimestamp = &v
}

// GetSrcClusterName returns the SrcClusterName field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetSrcClusterName() string {
	if o == nil || o.SrcClusterName == nil {
		var ret string
		return ret
	}
	return *o.SrcClusterName
}

// GetSrcClusterNameOk returns a tuple with the SrcClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetSrcClusterNameOk() (*string, bool) {
	if o == nil || o.SrcClusterName == nil {
		return nil, false
	}
	return o.SrcClusterName, true
}

// HasSrcClusterName returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasSrcClusterName() bool {
	if o != nil && o.SrcClusterName != nil {
		return true
	}

	return false
}

// SetSrcClusterName gets a reference to the given string and assigns it to the SrcClusterName field.
func (o *HyperflexVmSnapshotInfo) SetSrcClusterName(v string) {
	o.SrcClusterName = &v
}

// GetTargetCompletionTimestamp returns the TargetCompletionTimestamp field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetTargetCompletionTimestamp() int64 {
	if o == nil || o.TargetCompletionTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.TargetCompletionTimestamp
}

// GetTargetCompletionTimestampOk returns a tuple with the TargetCompletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetTargetCompletionTimestampOk() (*int64, bool) {
	if o == nil || o.TargetCompletionTimestamp == nil {
		return nil, false
	}
	return o.TargetCompletionTimestamp, true
}

// HasTargetCompletionTimestamp returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasTargetCompletionTimestamp() bool {
	if o != nil && o.TargetCompletionTimestamp != nil {
		return true
	}

	return false
}

// SetTargetCompletionTimestamp gets a reference to the given int64 and assigns it to the TargetCompletionTimestamp field.
func (o *HyperflexVmSnapshotInfo) SetTargetCompletionTimestamp(v int64) {
	o.TargetCompletionTimestamp = &v
}

// GetTgtClusterName returns the TgtClusterName field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetTgtClusterName() string {
	if o == nil || o.TgtClusterName == nil {
		var ret string
		return ret
	}
	return *o.TgtClusterName
}

// GetTgtClusterNameOk returns a tuple with the TgtClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetTgtClusterNameOk() (*string, bool) {
	if o == nil || o.TgtClusterName == nil {
		return nil, false
	}
	return o.TgtClusterName, true
}

// HasTgtClusterName returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasTgtClusterName() bool {
	if o != nil && o.TgtClusterName != nil {
		return true
	}

	return false
}

// SetTgtClusterName gets a reference to the given string and assigns it to the TgtClusterName field.
func (o *HyperflexVmSnapshotInfo) SetTgtClusterName(v string) {
	o.TgtClusterName = &v
}

// GetVmEntityReference returns the VmEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfo) GetVmEntityReference() HyperflexEntityReference {
	if o == nil || o.VmEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmEntityReference.Get()
}

// GetVmEntityReferenceOk returns a tuple with the VmEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfo) GetVmEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmEntityReference.Get(), o.VmEntityReference.IsSet()
}

// HasVmEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasVmEntityReference() bool {
	if o != nil && o.VmEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmEntityReference field.
func (o *HyperflexVmSnapshotInfo) SetVmEntityReference(v HyperflexEntityReference) {
	o.VmEntityReference.Set(&v)
}

// SetVmEntityReferenceNil sets the value for VmEntityReference to be an explicit nil
func (o *HyperflexVmSnapshotInfo) SetVmEntityReferenceNil() {
	o.VmEntityReference.Set(nil)
}

// UnsetVmEntityReference ensures that no value is present for VmEntityReference, not even an explicit nil
func (o *HyperflexVmSnapshotInfo) UnsetVmEntityReference() {
	o.VmEntityReference.Unset()
}

// GetVmSnapshotEntityReference returns the VmSnapshotEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVmSnapshotInfo) GetVmSnapshotEntityReference() HyperflexEntityReference {
	if o == nil || o.VmSnapshotEntityReference.Get() == nil {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.VmSnapshotEntityReference.Get()
}

// GetVmSnapshotEntityReferenceOk returns a tuple with the VmSnapshotEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVmSnapshotInfo) GetVmSnapshotEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmSnapshotEntityReference.Get(), o.VmSnapshotEntityReference.IsSet()
}

// HasVmSnapshotEntityReference returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasVmSnapshotEntityReference() bool {
	if o != nil && o.VmSnapshotEntityReference.IsSet() {
		return true
	}

	return false
}

// SetVmSnapshotEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the VmSnapshotEntityReference field.
func (o *HyperflexVmSnapshotInfo) SetVmSnapshotEntityReference(v HyperflexEntityReference) {
	o.VmSnapshotEntityReference.Set(&v)
}

// SetVmSnapshotEntityReferenceNil sets the value for VmSnapshotEntityReference to be an explicit nil
func (o *HyperflexVmSnapshotInfo) SetVmSnapshotEntityReferenceNil() {
	o.VmSnapshotEntityReference.Set(nil)
}

// UnsetVmSnapshotEntityReference ensures that no value is present for VmSnapshotEntityReference, not even an explicit nil
func (o *HyperflexVmSnapshotInfo) UnsetVmSnapshotEntityReference() {
	o.VmSnapshotEntityReference.Unset()
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || o.SrcCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.SrcCluster == nil {
		return nil, false
	}
	return o.SrcCluster, true
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasSrcCluster() bool {
	if o != nil && o.SrcCluster != nil {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexVmSnapshotInfo) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster = &v
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || o.TgtCluster == nil {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil || o.TgtCluster == nil {
		return nil, false
	}
	return o.TgtCluster, true
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasTgtCluster() bool {
	if o != nil && o.TgtCluster != nil {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given HyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexVmSnapshotInfo) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster = &v
}

// GetVmBackupInfo returns the VmBackupInfo field value if set, zero value otherwise.
func (o *HyperflexVmSnapshotInfo) GetVmBackupInfo() HyperflexVmBackupInfoRelationship {
	if o == nil || o.VmBackupInfo == nil {
		var ret HyperflexVmBackupInfoRelationship
		return ret
	}
	return *o.VmBackupInfo
}

// GetVmBackupInfoOk returns a tuple with the VmBackupInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVmSnapshotInfo) GetVmBackupInfoOk() (*HyperflexVmBackupInfoRelationship, bool) {
	if o == nil || o.VmBackupInfo == nil {
		return nil, false
	}
	return o.VmBackupInfo, true
}

// HasVmBackupInfo returns a boolean if a field has been set.
func (o *HyperflexVmSnapshotInfo) HasVmBackupInfo() bool {
	if o != nil && o.VmBackupInfo != nil {
		return true
	}

	return false
}

// SetVmBackupInfo gets a reference to the given HyperflexVmBackupInfoRelationship and assigns it to the VmBackupInfo field.
func (o *HyperflexVmSnapshotInfo) SetVmBackupInfo(v HyperflexVmBackupInfoRelationship) {
	o.VmBackupInfo = &v
}

func (o HyperflexVmSnapshotInfo) MarshalJSON() ([]byte, error) {
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
	if o.ClusterIdSnapMap != nil {
		toSerialize["ClusterIdSnapMap"] = o.ClusterIdSnapMap
	}
	if o.DisplayStatus != nil {
		toSerialize["DisplayStatus"] = o.DisplayStatus
	}
	if o.Error.IsSet() {
		toSerialize["Error"] = o.Error.Get()
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.ParentSnapshot.IsSet() {
		toSerialize["ParentSnapshot"] = o.ParentSnapshot.Get()
	}
	if o.ReplicationStatus != nil {
		toSerialize["ReplicationStatus"] = o.ReplicationStatus
	}
	if o.SnapshotErrorMsg != nil {
		toSerialize["SnapshotErrorMsg"] = o.SnapshotErrorMsg
	}
	if o.SnapshotStatus != nil {
		toSerialize["SnapshotStatus"] = o.SnapshotStatus
	}
	if o.SourceTimestamp != nil {
		toSerialize["SourceTimestamp"] = o.SourceTimestamp
	}
	if o.SrcClusterName != nil {
		toSerialize["SrcClusterName"] = o.SrcClusterName
	}
	if o.TargetCompletionTimestamp != nil {
		toSerialize["TargetCompletionTimestamp"] = o.TargetCompletionTimestamp
	}
	if o.TgtClusterName != nil {
		toSerialize["TgtClusterName"] = o.TgtClusterName
	}
	if o.VmEntityReference.IsSet() {
		toSerialize["VmEntityReference"] = o.VmEntityReference.Get()
	}
	if o.VmSnapshotEntityReference.IsSet() {
		toSerialize["VmSnapshotEntityReference"] = o.VmSnapshotEntityReference.Get()
	}
	if o.SrcCluster != nil {
		toSerialize["SrcCluster"] = o.SrcCluster
	}
	if o.TgtCluster != nil {
		toSerialize["TgtCluster"] = o.TgtCluster
	}
	if o.VmBackupInfo != nil {
		toSerialize["VmBackupInfo"] = o.VmBackupInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVmSnapshotInfo) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexVmSnapshotInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                                   `json:"ObjectType"`
		ClusterIdSnapMap []HyperflexMapClusterIdToStSnapshotPoint `json:"ClusterIdSnapMap,omitempty"`
		// Combined status code from replication and snapshot status to display in the Intersight UI. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
		DisplayStatus *string                     `json:"DisplayStatus,omitempty"`
		Error         NullableHyperflexErrorStack `json:"Error,omitempty"`
		// The name of the Virtual Machine and the time stamp of the snapshot.
		Label *string `json:"Label,omitempty"`
		// Quiesce Mode for snapshot taken on Hyperflex cluster. * `NONE` - The snapshot quiesce mode is none. * `CRASH` - The snapshot quiesce mode is crash. * `VMTOOLS` - The snapshot quiesce mode is VMTOOLS. * `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.
		Mode           *string                          `json:"Mode,omitempty"`
		ParentSnapshot NullableHyperflexEntityReference `json:"ParentSnapshot,omitempty"`
		// Replication status of the least successful target cluster. * `NONE` - Snapshot replication state is none. * `SUCCESS` - Snapshot completed successfully. * `FAILED` - Snapshot failed replication status code. * `PAUSED` - Snapshot replication paused status code. * `IN_USE` - Snapshot replica in use status code. * `STARTING` - Snapshot replication starting. * `REPLICATING` - Snapshot replication in progress.
		ReplicationStatus *string `json:"ReplicationStatus,omitempty"`
		// Error message from snapshot creation or replcation if any exist.
		SnapshotErrorMsg *string `json:"SnapshotErrorMsg,omitempty"`
		// Snapshot status of the source cluster. * `SUCCESS` - This snapshot status code is success. * `FAILED` - This snapshot status code is failed. * `IN_PROGRESS` - This snapshot status code is in progress. * `DELETING` - This snapshot status code is deleting. * `DELETED` - This snapshot status code is deleted. * `NONE` - This snapshot status code is none. * `INIT` - This snapshot status code is initializing.
		SnapshotStatus *string `json:"SnapshotStatus,omitempty"`
		// Timestamp the snapshot was created on the source cluster.
		SourceTimestamp *int64 `json:"SourceTimestamp,omitempty"`
		// Name of the cluster this snapshot resides on.
		SrcClusterName *string `json:"SrcClusterName,omitempty"`
		// Timestamp the snapshot was finished replicating on the target cluster.
		TargetCompletionTimestamp *int64 `json:"TargetCompletionTimestamp,omitempty"`
		// Name of the cluster this snapshot is replicated to.
		TgtClusterName            *string                            `json:"TgtClusterName,omitempty"`
		VmEntityReference         NullableHyperflexEntityReference   `json:"VmEntityReference,omitempty"`
		VmSnapshotEntityReference NullableHyperflexEntityReference   `json:"VmSnapshotEntityReference,omitempty"`
		SrcCluster                *HyperflexClusterRelationship      `json:"SrcCluster,omitempty"`
		TgtCluster                *HyperflexClusterRelationship      `json:"TgtCluster,omitempty"`
		VmBackupInfo              *HyperflexVmBackupInfoRelationship `json:"VmBackupInfo,omitempty"`
	}

	varHyperflexVmSnapshotInfoWithoutEmbeddedStruct := HyperflexVmSnapshotInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexVmSnapshotInfoWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexVmSnapshotInfo := _HyperflexVmSnapshotInfo{}
		varHyperflexVmSnapshotInfo.ClassId = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.ClassId
		varHyperflexVmSnapshotInfo.ObjectType = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.ObjectType
		varHyperflexVmSnapshotInfo.ClusterIdSnapMap = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.ClusterIdSnapMap
		varHyperflexVmSnapshotInfo.DisplayStatus = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.DisplayStatus
		varHyperflexVmSnapshotInfo.Error = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.Error
		varHyperflexVmSnapshotInfo.Label = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.Label
		varHyperflexVmSnapshotInfo.Mode = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.Mode
		varHyperflexVmSnapshotInfo.ParentSnapshot = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.ParentSnapshot
		varHyperflexVmSnapshotInfo.ReplicationStatus = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.ReplicationStatus
		varHyperflexVmSnapshotInfo.SnapshotErrorMsg = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.SnapshotErrorMsg
		varHyperflexVmSnapshotInfo.SnapshotStatus = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.SnapshotStatus
		varHyperflexVmSnapshotInfo.SourceTimestamp = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.SourceTimestamp
		varHyperflexVmSnapshotInfo.SrcClusterName = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.SrcClusterName
		varHyperflexVmSnapshotInfo.TargetCompletionTimestamp = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.TargetCompletionTimestamp
		varHyperflexVmSnapshotInfo.TgtClusterName = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.TgtClusterName
		varHyperflexVmSnapshotInfo.VmEntityReference = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.VmEntityReference
		varHyperflexVmSnapshotInfo.VmSnapshotEntityReference = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.VmSnapshotEntityReference
		varHyperflexVmSnapshotInfo.SrcCluster = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.SrcCluster
		varHyperflexVmSnapshotInfo.TgtCluster = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.TgtCluster
		varHyperflexVmSnapshotInfo.VmBackupInfo = varHyperflexVmSnapshotInfoWithoutEmbeddedStruct.VmBackupInfo
		*o = HyperflexVmSnapshotInfo(varHyperflexVmSnapshotInfo)
	} else {
		return err
	}

	varHyperflexVmSnapshotInfo := _HyperflexVmSnapshotInfo{}

	err = json.Unmarshal(bytes, &varHyperflexVmSnapshotInfo)
	if err == nil {
		o.MoBaseMo = varHyperflexVmSnapshotInfo.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterIdSnapMap")
		delete(additionalProperties, "DisplayStatus")
		delete(additionalProperties, "Error")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "ParentSnapshot")
		delete(additionalProperties, "ReplicationStatus")
		delete(additionalProperties, "SnapshotErrorMsg")
		delete(additionalProperties, "SnapshotStatus")
		delete(additionalProperties, "SourceTimestamp")
		delete(additionalProperties, "SrcClusterName")
		delete(additionalProperties, "TargetCompletionTimestamp")
		delete(additionalProperties, "TgtClusterName")
		delete(additionalProperties, "VmEntityReference")
		delete(additionalProperties, "VmSnapshotEntityReference")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")
		delete(additionalProperties, "VmBackupInfo")

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

type NullableHyperflexVmSnapshotInfo struct {
	value *HyperflexVmSnapshotInfo
	isSet bool
}

func (v NullableHyperflexVmSnapshotInfo) Get() *HyperflexVmSnapshotInfo {
	return v.value
}

func (v *NullableHyperflexVmSnapshotInfo) Set(val *HyperflexVmSnapshotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVmSnapshotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVmSnapshotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVmSnapshotInfo(val *HyperflexVmSnapshotInfo) *NullableHyperflexVmSnapshotInfo {
	return &NullableHyperflexVmSnapshotInfo{value: val, isSet: true}
}

func (v NullableHyperflexVmSnapshotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVmSnapshotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
