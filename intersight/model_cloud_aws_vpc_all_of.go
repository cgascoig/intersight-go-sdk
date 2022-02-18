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
)

// CloudAwsVpcAllOf Definition of the list of properties defined in 'cloud.AwsVpc', excluding properties defined in parent classes.
type CloudAwsVpcAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If true, instances in the vpc get public DNS hostnames.
	DnsHostName *bool `json:"DnsHostName,omitempty"`
	// Indicates whether the Dns resolution is supported.
	DnsResolution *bool    `json:"DnsResolution,omitempty"`
	Ipv4Cidr      []string `json:"Ipv4Cidr,omitempty"`
	Ipv6Cidr      []string `json:"Ipv6Cidr,omitempty"`
	// If true, indicates that this is default VPC.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// The state of the VPC (pending | available).
	State *string `json:"State,omitempty"`
	// The allowed tenancy of instances launched into the VPC.
	Tenancy              *string                          `json:"Tenancy,omitempty"`
	VpcTags              []CloudCloudTag                  `json:"VpcTags,omitempty"`
	AwsBillingUnit       *CloudAwsBillingUnitRelationship `json:"AwsBillingUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsVpcAllOf CloudAwsVpcAllOf

// NewCloudAwsVpcAllOf instantiates a new CloudAwsVpcAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsVpcAllOf(classId string, objectType string) *CloudAwsVpcAllOf {
	this := CloudAwsVpcAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsVpcAllOfWithDefaults instantiates a new CloudAwsVpcAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsVpcAllOfWithDefaults() *CloudAwsVpcAllOf {
	this := CloudAwsVpcAllOf{}
	var classId string = "cloud.AwsVpc"
	this.ClassId = classId
	var objectType string = "cloud.AwsVpc"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsVpcAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsVpcAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsVpcAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsVpcAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDnsHostName returns the DnsHostName field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetDnsHostName() bool {
	if o == nil || o.DnsHostName == nil {
		var ret bool
		return ret
	}
	return *o.DnsHostName
}

// GetDnsHostNameOk returns a tuple with the DnsHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetDnsHostNameOk() (*bool, bool) {
	if o == nil || o.DnsHostName == nil {
		return nil, false
	}
	return o.DnsHostName, true
}

// HasDnsHostName returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasDnsHostName() bool {
	if o != nil && o.DnsHostName != nil {
		return true
	}

	return false
}

// SetDnsHostName gets a reference to the given bool and assigns it to the DnsHostName field.
func (o *CloudAwsVpcAllOf) SetDnsHostName(v bool) {
	o.DnsHostName = &v
}

// GetDnsResolution returns the DnsResolution field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetDnsResolution() bool {
	if o == nil || o.DnsResolution == nil {
		var ret bool
		return ret
	}
	return *o.DnsResolution
}

// GetDnsResolutionOk returns a tuple with the DnsResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetDnsResolutionOk() (*bool, bool) {
	if o == nil || o.DnsResolution == nil {
		return nil, false
	}
	return o.DnsResolution, true
}

// HasDnsResolution returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasDnsResolution() bool {
	if o != nil && o.DnsResolution != nil {
		return true
	}

	return false
}

// SetDnsResolution gets a reference to the given bool and assigns it to the DnsResolution field.
func (o *CloudAwsVpcAllOf) SetDnsResolution(v bool) {
	o.DnsResolution = &v
}

// GetIpv4Cidr returns the Ipv4Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpcAllOf) GetIpv4Cidr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ipv4Cidr
}

// GetIpv4CidrOk returns a tuple with the Ipv4Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpcAllOf) GetIpv4CidrOk() (*[]string, bool) {
	if o == nil || o.Ipv4Cidr == nil {
		return nil, false
	}
	return &o.Ipv4Cidr, true
}

// HasIpv4Cidr returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasIpv4Cidr() bool {
	if o != nil && o.Ipv4Cidr != nil {
		return true
	}

	return false
}

// SetIpv4Cidr gets a reference to the given []string and assigns it to the Ipv4Cidr field.
func (o *CloudAwsVpcAllOf) SetIpv4Cidr(v []string) {
	o.Ipv4Cidr = v
}

// GetIpv6Cidr returns the Ipv6Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpcAllOf) GetIpv6Cidr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ipv6Cidr
}

// GetIpv6CidrOk returns a tuple with the Ipv6Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpcAllOf) GetIpv6CidrOk() (*[]string, bool) {
	if o == nil || o.Ipv6Cidr == nil {
		return nil, false
	}
	return &o.Ipv6Cidr, true
}

// HasIpv6Cidr returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasIpv6Cidr() bool {
	if o != nil && o.Ipv6Cidr != nil {
		return true
	}

	return false
}

// SetIpv6Cidr gets a reference to the given []string and assigns it to the Ipv6Cidr field.
func (o *CloudAwsVpcAllOf) SetIpv6Cidr(v []string) {
	o.Ipv6Cidr = v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudAwsVpcAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudAwsVpcAllOf) SetState(v string) {
	o.State = &v
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *CloudAwsVpcAllOf) SetTenancy(v string) {
	o.Tenancy = &v
}

// GetVpcTags returns the VpcTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsVpcAllOf) GetVpcTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.VpcTags
}

// GetVpcTagsOk returns a tuple with the VpcTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsVpcAllOf) GetVpcTagsOk() (*[]CloudCloudTag, bool) {
	if o == nil || o.VpcTags == nil {
		return nil, false
	}
	return &o.VpcTags, true
}

// HasVpcTags returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasVpcTags() bool {
	if o != nil && o.VpcTags != nil {
		return true
	}

	return false
}

// SetVpcTags gets a reference to the given []CloudCloudTag and assigns it to the VpcTags field.
func (o *CloudAwsVpcAllOf) SetVpcTags(v []CloudCloudTag) {
	o.VpcTags = v
}

// GetAwsBillingUnit returns the AwsBillingUnit field value if set, zero value otherwise.
func (o *CloudAwsVpcAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship {
	if o == nil || o.AwsBillingUnit == nil {
		var ret CloudAwsBillingUnitRelationship
		return ret
	}
	return *o.AwsBillingUnit
}

// GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsVpcAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool) {
	if o == nil || o.AwsBillingUnit == nil {
		return nil, false
	}
	return o.AwsBillingUnit, true
}

// HasAwsBillingUnit returns a boolean if a field has been set.
func (o *CloudAwsVpcAllOf) HasAwsBillingUnit() bool {
	if o != nil && o.AwsBillingUnit != nil {
		return true
	}

	return false
}

// SetAwsBillingUnit gets a reference to the given CloudAwsBillingUnitRelationship and assigns it to the AwsBillingUnit field.
func (o *CloudAwsVpcAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship) {
	o.AwsBillingUnit = &v
}

func (o CloudAwsVpcAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DnsHostName != nil {
		toSerialize["DnsHostName"] = o.DnsHostName
	}
	if o.DnsResolution != nil {
		toSerialize["DnsResolution"] = o.DnsResolution
	}
	if o.Ipv4Cidr != nil {
		toSerialize["Ipv4Cidr"] = o.Ipv4Cidr
	}
	if o.Ipv6Cidr != nil {
		toSerialize["Ipv6Cidr"] = o.Ipv6Cidr
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	if o.VpcTags != nil {
		toSerialize["VpcTags"] = o.VpcTags
	}
	if o.AwsBillingUnit != nil {
		toSerialize["AwsBillingUnit"] = o.AwsBillingUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsVpcAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsVpcAllOf := _CloudAwsVpcAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsVpcAllOf); err == nil {
		*o = CloudAwsVpcAllOf(varCloudAwsVpcAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DnsHostName")
		delete(additionalProperties, "DnsResolution")
		delete(additionalProperties, "Ipv4Cidr")
		delete(additionalProperties, "Ipv6Cidr")
		delete(additionalProperties, "IsDefault")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Tenancy")
		delete(additionalProperties, "VpcTags")
		delete(additionalProperties, "AwsBillingUnit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsVpcAllOf struct {
	value *CloudAwsVpcAllOf
	isSet bool
}

func (v NullableCloudAwsVpcAllOf) Get() *CloudAwsVpcAllOf {
	return v.value
}

func (v *NullableCloudAwsVpcAllOf) Set(val *CloudAwsVpcAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsVpcAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsVpcAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsVpcAllOf(val *CloudAwsVpcAllOf) *NullableCloudAwsVpcAllOf {
	return &NullableCloudAwsVpcAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsVpcAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsVpcAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
