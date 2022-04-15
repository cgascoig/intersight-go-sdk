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

// NotificationAccountSubscription AccountSubscription is a concrete type that extends abstract Subscription type and intent to be used only for account level subscriptions by Account Administrator.
type NotificationAccountSubscription struct {
	NotificationSubscription
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the subscription.
	Name *string `json:"Name,omitempty"`
	// The chosen subscription type imposes it is own validation rules. When 'email' type is chosen, actions array can contain only one entry and it is entry should be of can be only notification.SendEmail; conditions can contain only notification.AlarmMoCondition and condition types should be unique. When the 'webhook' type is chosen, the actions array can contain only one entry and it is entry should be of can be only notification.TriggerWebhook; conditions can contain up to a limited amount of entries and all of them should be of type notification.MoCondition. * `email` - Email type requires usage of notification.SendEmail complex types for actionsand notification.AlarmMoCondition complex types for conditions. * `webhook` - Webhook type requires usage of notification.TriggerWebhook complex types for actionsand notification.MoCondition complex types for conditions.
	Type *string `json:"Type,omitempty"`
	// Used to verify the actions of the Subscription MO. For a 'webhook' type Ping event is sent to verify that the webhook server is accessible. For an 'email' type there will be a verification email sent. * `none` - No actions will be verified. Default value. * `all` - All actions will be re-verified. The previous state of the verification will be preserved.
	Verify               *string                 `json:"Verify,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationAccountSubscription NotificationAccountSubscription

// NewNotificationAccountSubscription instantiates a new NotificationAccountSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationAccountSubscription(classId string, objectType string) *NotificationAccountSubscription {
	this := NotificationAccountSubscription{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "email"
	this.Type = &type_
	var verify string = "none"
	this.Verify = &verify
	return &this
}

// NewNotificationAccountSubscriptionWithDefaults instantiates a new NotificationAccountSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationAccountSubscriptionWithDefaults() *NotificationAccountSubscription {
	this := NotificationAccountSubscription{}
	var classId string = "notification.AccountSubscription"
	this.ClassId = classId
	var objectType string = "notification.AccountSubscription"
	this.ObjectType = objectType
	var type_ string = "email"
	this.Type = &type_
	var verify string = "none"
	this.Verify = &verify
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationAccountSubscription) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationAccountSubscription) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationAccountSubscription) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationAccountSubscription) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationAccountSubscription) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationAccountSubscription) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationAccountSubscription) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationAccountSubscription) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationAccountSubscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationAccountSubscription) SetType(v string) {
	o.Type = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *NotificationAccountSubscription) GetVerify() string {
	if o == nil || o.Verify == nil {
		var ret string
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetVerifyOk() (*string, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *NotificationAccountSubscription) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given string and assigns it to the Verify field.
func (o *NotificationAccountSubscription) SetVerify(v string) {
	o.Verify = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NotificationAccountSubscription) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationAccountSubscription) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *NotificationAccountSubscription) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *NotificationAccountSubscription) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o NotificationAccountSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationSubscription, errNotificationSubscription := json.Marshal(o.NotificationSubscription)
	if errNotificationSubscription != nil {
		return []byte{}, errNotificationSubscription
	}
	errNotificationSubscription = json.Unmarshal([]byte(serializedNotificationSubscription), &toSerialize)
	if errNotificationSubscription != nil {
		return []byte{}, errNotificationSubscription
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Verify != nil {
		toSerialize["Verify"] = o.Verify
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationAccountSubscription) UnmarshalJSON(bytes []byte) (err error) {
	type NotificationAccountSubscriptionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the subscription.
		Name *string `json:"Name,omitempty"`
		// The chosen subscription type imposes it is own validation rules. When 'email' type is chosen, actions array can contain only one entry and it is entry should be of can be only notification.SendEmail; conditions can contain only notification.AlarmMoCondition and condition types should be unique. When the 'webhook' type is chosen, the actions array can contain only one entry and it is entry should be of can be only notification.TriggerWebhook; conditions can contain up to a limited amount of entries and all of them should be of type notification.MoCondition. * `email` - Email type requires usage of notification.SendEmail complex types for actionsand notification.AlarmMoCondition complex types for conditions. * `webhook` - Webhook type requires usage of notification.TriggerWebhook complex types for actionsand notification.MoCondition complex types for conditions.
		Type *string `json:"Type,omitempty"`
		// Used to verify the actions of the Subscription MO. For a 'webhook' type Ping event is sent to verify that the webhook server is accessible. For an 'email' type there will be a verification email sent. * `none` - No actions will be verified. Default value. * `all` - All actions will be re-verified. The previous state of the verification will be preserved.
		Verify  *string                 `json:"Verify,omitempty"`
		Account *IamAccountRelationship `json:"Account,omitempty"`
	}

	varNotificationAccountSubscriptionWithoutEmbeddedStruct := NotificationAccountSubscriptionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNotificationAccountSubscriptionWithoutEmbeddedStruct)
	if err == nil {
		varNotificationAccountSubscription := _NotificationAccountSubscription{}
		varNotificationAccountSubscription.ClassId = varNotificationAccountSubscriptionWithoutEmbeddedStruct.ClassId
		varNotificationAccountSubscription.ObjectType = varNotificationAccountSubscriptionWithoutEmbeddedStruct.ObjectType
		varNotificationAccountSubscription.Name = varNotificationAccountSubscriptionWithoutEmbeddedStruct.Name
		varNotificationAccountSubscription.Type = varNotificationAccountSubscriptionWithoutEmbeddedStruct.Type
		varNotificationAccountSubscription.Verify = varNotificationAccountSubscriptionWithoutEmbeddedStruct.Verify
		varNotificationAccountSubscription.Account = varNotificationAccountSubscriptionWithoutEmbeddedStruct.Account
		*o = NotificationAccountSubscription(varNotificationAccountSubscription)
	} else {
		return err
	}

	varNotificationAccountSubscription := _NotificationAccountSubscription{}

	err = json.Unmarshal(bytes, &varNotificationAccountSubscription)
	if err == nil {
		o.NotificationSubscription = varNotificationAccountSubscription.NotificationSubscription
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Verify")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectNotificationSubscription := reflect.ValueOf(o.NotificationSubscription)
		for i := 0; i < reflectNotificationSubscription.Type().NumField(); i++ {
			t := reflectNotificationSubscription.Type().Field(i)

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

type NullableNotificationAccountSubscription struct {
	value *NotificationAccountSubscription
	isSet bool
}

func (v NullableNotificationAccountSubscription) Get() *NotificationAccountSubscription {
	return v.value
}

func (v *NullableNotificationAccountSubscription) Set(val *NotificationAccountSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationAccountSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationAccountSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationAccountSubscription(val *NotificationAccountSubscription) *NullableNotificationAccountSubscription {
	return &NullableNotificationAccountSubscription{value: val, isSet: true}
}

func (v NullableNotificationAccountSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationAccountSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
