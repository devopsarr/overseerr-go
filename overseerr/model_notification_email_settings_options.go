/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.33.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the NotificationEmailSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationEmailSettingsOptions{}

// NotificationEmailSettingsOptions struct for NotificationEmailSettingsOptions
type NotificationEmailSettingsOptions struct {
	EmailFrom *string `json:"emailFrom,omitempty"`
	SenderName *string `json:"senderName,omitempty"`
	SmtpHost *string `json:"smtpHost,omitempty"`
	SmtpPort *float32 `json:"smtpPort,omitempty"`
	Secure *bool `json:"secure,omitempty"`
	IgnoreTls *bool `json:"ignoreTls,omitempty"`
	RequireTls *bool `json:"requireTls,omitempty"`
	AuthUser NullableString `json:"authUser,omitempty"`
	AuthPass NullableString `json:"authPass,omitempty"`
	AllowSelfSigned *bool `json:"allowSelfSigned,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationEmailSettingsOptions NotificationEmailSettingsOptions

// NewNotificationEmailSettingsOptions instantiates a new NotificationEmailSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEmailSettingsOptions() *NotificationEmailSettingsOptions {
	this := NotificationEmailSettingsOptions{}
	return &this
}

// NewNotificationEmailSettingsOptionsWithDefaults instantiates a new NotificationEmailSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEmailSettingsOptionsWithDefaults() *NotificationEmailSettingsOptions {
	this := NotificationEmailSettingsOptions{}
	return &this
}

// GetEmailFrom returns the EmailFrom field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetEmailFrom() string {
	if o == nil || IsNil(o.EmailFrom) {
		var ret string
		return ret
	}
	return *o.EmailFrom
}

// GetEmailFromOk returns a tuple with the EmailFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetEmailFromOk() (*string, bool) {
	if o == nil || IsNil(o.EmailFrom) {
		return nil, false
	}
	return o.EmailFrom, true
}

// HasEmailFrom returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasEmailFrom() bool {
	if o != nil && !IsNil(o.EmailFrom) {
		return true
	}

	return false
}

// SetEmailFrom gets a reference to the given string and assigns it to the EmailFrom field.
func (o *NotificationEmailSettingsOptions) SetEmailFrom(v string) {
	o.EmailFrom = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *NotificationEmailSettingsOptions) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSmtpHost returns the SmtpHost field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetSmtpHost() string {
	if o == nil || IsNil(o.SmtpHost) {
		var ret string
		return ret
	}
	return *o.SmtpHost
}

// GetSmtpHostOk returns a tuple with the SmtpHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetSmtpHostOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpHost) {
		return nil, false
	}
	return o.SmtpHost, true
}

// HasSmtpHost returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasSmtpHost() bool {
	if o != nil && !IsNil(o.SmtpHost) {
		return true
	}

	return false
}

// SetSmtpHost gets a reference to the given string and assigns it to the SmtpHost field.
func (o *NotificationEmailSettingsOptions) SetSmtpHost(v string) {
	o.SmtpHost = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetSmtpPort() float32 {
	if o == nil || IsNil(o.SmtpPort) {
		var ret float32
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetSmtpPortOk() (*float32, bool) {
	if o == nil || IsNil(o.SmtpPort) {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasSmtpPort() bool {
	if o != nil && !IsNil(o.SmtpPort) {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given float32 and assigns it to the SmtpPort field.
func (o *NotificationEmailSettingsOptions) SetSmtpPort(v float32) {
	o.SmtpPort = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *NotificationEmailSettingsOptions) SetSecure(v bool) {
	o.Secure = &v
}

// GetIgnoreTls returns the IgnoreTls field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetIgnoreTls() bool {
	if o == nil || IsNil(o.IgnoreTls) {
		var ret bool
		return ret
	}
	return *o.IgnoreTls
}

// GetIgnoreTlsOk returns a tuple with the IgnoreTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetIgnoreTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreTls) {
		return nil, false
	}
	return o.IgnoreTls, true
}

// HasIgnoreTls returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasIgnoreTls() bool {
	if o != nil && !IsNil(o.IgnoreTls) {
		return true
	}

	return false
}

// SetIgnoreTls gets a reference to the given bool and assigns it to the IgnoreTls field.
func (o *NotificationEmailSettingsOptions) SetIgnoreTls(v bool) {
	o.IgnoreTls = &v
}

// GetRequireTls returns the RequireTls field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetRequireTls() bool {
	if o == nil || IsNil(o.RequireTls) {
		var ret bool
		return ret
	}
	return *o.RequireTls
}

// GetRequireTlsOk returns a tuple with the RequireTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetRequireTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireTls) {
		return nil, false
	}
	return o.RequireTls, true
}

// HasRequireTls returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasRequireTls() bool {
	if o != nil && !IsNil(o.RequireTls) {
		return true
	}

	return false
}

// SetRequireTls gets a reference to the given bool and assigns it to the RequireTls field.
func (o *NotificationEmailSettingsOptions) SetRequireTls(v bool) {
	o.RequireTls = &v
}

// GetAuthUser returns the AuthUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationEmailSettingsOptions) GetAuthUser() string {
	if o == nil || IsNil(o.AuthUser.Get()) {
		var ret string
		return ret
	}
	return *o.AuthUser.Get()
}

// GetAuthUserOk returns a tuple with the AuthUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationEmailSettingsOptions) GetAuthUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthUser.Get(), o.AuthUser.IsSet()
}

// HasAuthUser returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasAuthUser() bool {
	if o != nil && o.AuthUser.IsSet() {
		return true
	}

	return false
}

// SetAuthUser gets a reference to the given NullableString and assigns it to the AuthUser field.
func (o *NotificationEmailSettingsOptions) SetAuthUser(v string) {
	o.AuthUser.Set(&v)
}
// SetAuthUserNil sets the value for AuthUser to be an explicit nil
func (o *NotificationEmailSettingsOptions) SetAuthUserNil() {
	o.AuthUser.Set(nil)
}

// UnsetAuthUser ensures that no value is present for AuthUser, not even an explicit nil
func (o *NotificationEmailSettingsOptions) UnsetAuthUser() {
	o.AuthUser.Unset()
}

// GetAuthPass returns the AuthPass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationEmailSettingsOptions) GetAuthPass() string {
	if o == nil || IsNil(o.AuthPass.Get()) {
		var ret string
		return ret
	}
	return *o.AuthPass.Get()
}

// GetAuthPassOk returns a tuple with the AuthPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationEmailSettingsOptions) GetAuthPassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthPass.Get(), o.AuthPass.IsSet()
}

// HasAuthPass returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasAuthPass() bool {
	if o != nil && o.AuthPass.IsSet() {
		return true
	}

	return false
}

// SetAuthPass gets a reference to the given NullableString and assigns it to the AuthPass field.
func (o *NotificationEmailSettingsOptions) SetAuthPass(v string) {
	o.AuthPass.Set(&v)
}
// SetAuthPassNil sets the value for AuthPass to be an explicit nil
func (o *NotificationEmailSettingsOptions) SetAuthPassNil() {
	o.AuthPass.Set(nil)
}

// UnsetAuthPass ensures that no value is present for AuthPass, not even an explicit nil
func (o *NotificationEmailSettingsOptions) UnsetAuthPass() {
	o.AuthPass.Unset()
}

// GetAllowSelfSigned returns the AllowSelfSigned field value if set, zero value otherwise.
func (o *NotificationEmailSettingsOptions) GetAllowSelfSigned() bool {
	if o == nil || IsNil(o.AllowSelfSigned) {
		var ret bool
		return ret
	}
	return *o.AllowSelfSigned
}

// GetAllowSelfSignedOk returns a tuple with the AllowSelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettingsOptions) GetAllowSelfSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSelfSigned) {
		return nil, false
	}
	return o.AllowSelfSigned, true
}

// HasAllowSelfSigned returns a boolean if a field has been set.
func (o *NotificationEmailSettingsOptions) HasAllowSelfSigned() bool {
	if o != nil && !IsNil(o.AllowSelfSigned) {
		return true
	}

	return false
}

// SetAllowSelfSigned gets a reference to the given bool and assigns it to the AllowSelfSigned field.
func (o *NotificationEmailSettingsOptions) SetAllowSelfSigned(v bool) {
	o.AllowSelfSigned = &v
}

func (o NotificationEmailSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationEmailSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailFrom) {
		toSerialize["emailFrom"] = o.EmailFrom
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	if !IsNil(o.SmtpHost) {
		toSerialize["smtpHost"] = o.SmtpHost
	}
	if !IsNil(o.SmtpPort) {
		toSerialize["smtpPort"] = o.SmtpPort
	}
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !IsNil(o.IgnoreTls) {
		toSerialize["ignoreTls"] = o.IgnoreTls
	}
	if !IsNil(o.RequireTls) {
		toSerialize["requireTls"] = o.RequireTls
	}
	if o.AuthUser.IsSet() {
		toSerialize["authUser"] = o.AuthUser.Get()
	}
	if o.AuthPass.IsSet() {
		toSerialize["authPass"] = o.AuthPass.Get()
	}
	if !IsNil(o.AllowSelfSigned) {
		toSerialize["allowSelfSigned"] = o.AllowSelfSigned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationEmailSettingsOptions) UnmarshalJSON(data []byte) (err error) {
	varNotificationEmailSettingsOptions := _NotificationEmailSettingsOptions{}

	err = json.Unmarshal(data, &varNotificationEmailSettingsOptions)

	if err != nil {
		return err
	}

	*o = NotificationEmailSettingsOptions(varNotificationEmailSettingsOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emailFrom")
		delete(additionalProperties, "senderName")
		delete(additionalProperties, "smtpHost")
		delete(additionalProperties, "smtpPort")
		delete(additionalProperties, "secure")
		delete(additionalProperties, "ignoreTls")
		delete(additionalProperties, "requireTls")
		delete(additionalProperties, "authUser")
		delete(additionalProperties, "authPass")
		delete(additionalProperties, "allowSelfSigned")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationEmailSettingsOptions struct {
	value *NotificationEmailSettingsOptions
	isSet bool
}

func (v NullableNotificationEmailSettingsOptions) Get() *NotificationEmailSettingsOptions {
	return v.value
}

func (v *NullableNotificationEmailSettingsOptions) Set(val *NotificationEmailSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEmailSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEmailSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEmailSettingsOptions(val *NotificationEmailSettingsOptions) *NullableNotificationEmailSettingsOptions {
	return &NullableNotificationEmailSettingsOptions{value: val, isSet: true}
}

func (v NullableNotificationEmailSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEmailSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


