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

// checks if the NotificationEmailSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationEmailSettings{}

// NotificationEmailSettings struct for NotificationEmailSettings
type NotificationEmailSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Types *float32 `json:"types,omitempty"`
	Options *NotificationEmailSettingsOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationEmailSettings NotificationEmailSettings

// NewNotificationEmailSettings instantiates a new NotificationEmailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationEmailSettings() *NotificationEmailSettings {
	this := NotificationEmailSettings{}
	return &this
}

// NewNotificationEmailSettingsWithDefaults instantiates a new NotificationEmailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationEmailSettingsWithDefaults() *NotificationEmailSettings {
	this := NotificationEmailSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NotificationEmailSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NotificationEmailSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NotificationEmailSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *NotificationEmailSettings) GetTypes() float32 {
	if o == nil || IsNil(o.Types) {
		var ret float32
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettings) GetTypesOk() (*float32, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *NotificationEmailSettings) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given float32 and assigns it to the Types field.
func (o *NotificationEmailSettings) SetTypes(v float32) {
	o.Types = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *NotificationEmailSettings) GetOptions() NotificationEmailSettingsOptions {
	if o == nil || IsNil(o.Options) {
		var ret NotificationEmailSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationEmailSettings) GetOptionsOk() (*NotificationEmailSettingsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *NotificationEmailSettings) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NotificationEmailSettingsOptions and assigns it to the Options field.
func (o *NotificationEmailSettings) SetOptions(v NotificationEmailSettingsOptions) {
	o.Options = &v
}

func (o NotificationEmailSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationEmailSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationEmailSettings) UnmarshalJSON(data []byte) (err error) {
	varNotificationEmailSettings := _NotificationEmailSettings{}

	err = json.Unmarshal(data, &varNotificationEmailSettings)

	if err != nil {
		return err
	}

	*o = NotificationEmailSettings(varNotificationEmailSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "types")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationEmailSettings struct {
	value *NotificationEmailSettings
	isSet bool
}

func (v NullableNotificationEmailSettings) Get() *NotificationEmailSettings {
	return v.value
}

func (v *NullableNotificationEmailSettings) Set(val *NotificationEmailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEmailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEmailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEmailSettings(val *NotificationEmailSettings) *NullableNotificationEmailSettings {
	return &NullableNotificationEmailSettings{value: val, isSet: true}
}

func (v NullableNotificationEmailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEmailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


