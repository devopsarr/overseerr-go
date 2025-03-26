/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the DiscordSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscordSettingsOptions{}

// DiscordSettingsOptions struct for DiscordSettingsOptions
type DiscordSettingsOptions struct {
	BotUsername *string `json:"botUsername,omitempty"`
	BotAvatarUrl *string `json:"botAvatarUrl,omitempty"`
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	EnableMentions *bool `json:"enableMentions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscordSettingsOptions DiscordSettingsOptions

// NewDiscordSettingsOptions instantiates a new DiscordSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscordSettingsOptions() *DiscordSettingsOptions {
	this := DiscordSettingsOptions{}
	return &this
}

// NewDiscordSettingsOptionsWithDefaults instantiates a new DiscordSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscordSettingsOptionsWithDefaults() *DiscordSettingsOptions {
	this := DiscordSettingsOptions{}
	return &this
}

// GetBotUsername returns the BotUsername field value if set, zero value otherwise.
func (o *DiscordSettingsOptions) GetBotUsername() string {
	if o == nil || IsNil(o.BotUsername) {
		var ret string
		return ret
	}
	return *o.BotUsername
}

// GetBotUsernameOk returns a tuple with the BotUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordSettingsOptions) GetBotUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BotUsername) {
		return nil, false
	}
	return o.BotUsername, true
}

// HasBotUsername returns a boolean if a field has been set.
func (o *DiscordSettingsOptions) HasBotUsername() bool {
	if o != nil && !IsNil(o.BotUsername) {
		return true
	}

	return false
}

// SetBotUsername gets a reference to the given string and assigns it to the BotUsername field.
func (o *DiscordSettingsOptions) SetBotUsername(v string) {
	o.BotUsername = &v
}

// GetBotAvatarUrl returns the BotAvatarUrl field value if set, zero value otherwise.
func (o *DiscordSettingsOptions) GetBotAvatarUrl() string {
	if o == nil || IsNil(o.BotAvatarUrl) {
		var ret string
		return ret
	}
	return *o.BotAvatarUrl
}

// GetBotAvatarUrlOk returns a tuple with the BotAvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordSettingsOptions) GetBotAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BotAvatarUrl) {
		return nil, false
	}
	return o.BotAvatarUrl, true
}

// HasBotAvatarUrl returns a boolean if a field has been set.
func (o *DiscordSettingsOptions) HasBotAvatarUrl() bool {
	if o != nil && !IsNil(o.BotAvatarUrl) {
		return true
	}

	return false
}

// SetBotAvatarUrl gets a reference to the given string and assigns it to the BotAvatarUrl field.
func (o *DiscordSettingsOptions) SetBotAvatarUrl(v string) {
	o.BotAvatarUrl = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *DiscordSettingsOptions) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordSettingsOptions) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *DiscordSettingsOptions) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *DiscordSettingsOptions) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetEnableMentions returns the EnableMentions field value if set, zero value otherwise.
func (o *DiscordSettingsOptions) GetEnableMentions() bool {
	if o == nil || IsNil(o.EnableMentions) {
		var ret bool
		return ret
	}
	return *o.EnableMentions
}

// GetEnableMentionsOk returns a tuple with the EnableMentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscordSettingsOptions) GetEnableMentionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMentions) {
		return nil, false
	}
	return o.EnableMentions, true
}

// HasEnableMentions returns a boolean if a field has been set.
func (o *DiscordSettingsOptions) HasEnableMentions() bool {
	if o != nil && !IsNil(o.EnableMentions) {
		return true
	}

	return false
}

// SetEnableMentions gets a reference to the given bool and assigns it to the EnableMentions field.
func (o *DiscordSettingsOptions) SetEnableMentions(v bool) {
	o.EnableMentions = &v
}

func (o DiscordSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscordSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BotUsername) {
		toSerialize["botUsername"] = o.BotUsername
	}
	if !IsNil(o.BotAvatarUrl) {
		toSerialize["botAvatarUrl"] = o.BotAvatarUrl
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	if !IsNil(o.EnableMentions) {
		toSerialize["enableMentions"] = o.EnableMentions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscordSettingsOptions) UnmarshalJSON(data []byte) (err error) {
	varDiscordSettingsOptions := _DiscordSettingsOptions{}

	err = json.Unmarshal(data, &varDiscordSettingsOptions)

	if err != nil {
		return err
	}

	*o = DiscordSettingsOptions(varDiscordSettingsOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "botUsername")
		delete(additionalProperties, "botAvatarUrl")
		delete(additionalProperties, "webhookUrl")
		delete(additionalProperties, "enableMentions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscordSettingsOptions struct {
	value *DiscordSettingsOptions
	isSet bool
}

func (v NullableDiscordSettingsOptions) Get() *DiscordSettingsOptions {
	return v.value
}

func (v *NullableDiscordSettingsOptions) Set(val *DiscordSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscordSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscordSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscordSettingsOptions(val *DiscordSettingsOptions) *NullableDiscordSettingsOptions {
	return &NullableDiscordSettingsOptions{value: val, isSet: true}
}

func (v NullableDiscordSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscordSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


