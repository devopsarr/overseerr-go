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

// checks if the PushbulletSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushbulletSettingsOptions{}

// PushbulletSettingsOptions struct for PushbulletSettingsOptions
type PushbulletSettingsOptions struct {
	AccessToken *string `json:"accessToken,omitempty"`
	ChannelTag NullableString `json:"channelTag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushbulletSettingsOptions PushbulletSettingsOptions

// NewPushbulletSettingsOptions instantiates a new PushbulletSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushbulletSettingsOptions() *PushbulletSettingsOptions {
	this := PushbulletSettingsOptions{}
	return &this
}

// NewPushbulletSettingsOptionsWithDefaults instantiates a new PushbulletSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushbulletSettingsOptionsWithDefaults() *PushbulletSettingsOptions {
	this := PushbulletSettingsOptions{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *PushbulletSettingsOptions) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushbulletSettingsOptions) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *PushbulletSettingsOptions) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *PushbulletSettingsOptions) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetChannelTag returns the ChannelTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PushbulletSettingsOptions) GetChannelTag() string {
	if o == nil || IsNil(o.ChannelTag.Get()) {
		var ret string
		return ret
	}
	return *o.ChannelTag.Get()
}

// GetChannelTagOk returns a tuple with the ChannelTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PushbulletSettingsOptions) GetChannelTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelTag.Get(), o.ChannelTag.IsSet()
}

// HasChannelTag returns a boolean if a field has been set.
func (o *PushbulletSettingsOptions) HasChannelTag() bool {
	if o != nil && o.ChannelTag.IsSet() {
		return true
	}

	return false
}

// SetChannelTag gets a reference to the given NullableString and assigns it to the ChannelTag field.
func (o *PushbulletSettingsOptions) SetChannelTag(v string) {
	o.ChannelTag.Set(&v)
}
// SetChannelTagNil sets the value for ChannelTag to be an explicit nil
func (o *PushbulletSettingsOptions) SetChannelTagNil() {
	o.ChannelTag.Set(nil)
}

// UnsetChannelTag ensures that no value is present for ChannelTag, not even an explicit nil
func (o *PushbulletSettingsOptions) UnsetChannelTag() {
	o.ChannelTag.Unset()
}

func (o PushbulletSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushbulletSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if o.ChannelTag.IsSet() {
		toSerialize["channelTag"] = o.ChannelTag.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PushbulletSettingsOptions) UnmarshalJSON(data []byte) (err error) {
	varPushbulletSettingsOptions := _PushbulletSettingsOptions{}

	err = json.Unmarshal(data, &varPushbulletSettingsOptions)

	if err != nil {
		return err
	}

	*o = PushbulletSettingsOptions(varPushbulletSettingsOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessToken")
		delete(additionalProperties, "channelTag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePushbulletSettingsOptions struct {
	value *PushbulletSettingsOptions
	isSet bool
}

func (v NullablePushbulletSettingsOptions) Get() *PushbulletSettingsOptions {
	return v.value
}

func (v *NullablePushbulletSettingsOptions) Set(val *PushbulletSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePushbulletSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePushbulletSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushbulletSettingsOptions(val *PushbulletSettingsOptions) *NullablePushbulletSettingsOptions {
	return &NullablePushbulletSettingsOptions{value: val, isSet: true}
}

func (v NullablePushbulletSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushbulletSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


