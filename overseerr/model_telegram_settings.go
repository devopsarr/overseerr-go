/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// TelegramSettings struct for TelegramSettings
type TelegramSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Types *float32 `json:"types,omitempty"`
	Options *TelegramSettingsOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelegramSettings TelegramSettings

// NewTelegramSettings instantiates a new TelegramSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelegramSettings() *TelegramSettings {
	this := TelegramSettings{}
	return &this
}

// NewTelegramSettingsWithDefaults instantiates a new TelegramSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelegramSettingsWithDefaults() *TelegramSettings {
	this := TelegramSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *TelegramSettings) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *TelegramSettings) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *TelegramSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *TelegramSettings) GetTypes() float32 {
	if o == nil || isNil(o.Types) {
		var ret float32
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettings) GetTypesOk() (*float32, bool) {
	if o == nil || isNil(o.Types) {
    return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *TelegramSettings) HasTypes() bool {
	if o != nil && !isNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given float32 and assigns it to the Types field.
func (o *TelegramSettings) SetTypes(v float32) {
	o.Types = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TelegramSettings) GetOptions() TelegramSettingsOptions {
	if o == nil || isNil(o.Options) {
		var ret TelegramSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettings) GetOptionsOk() (*TelegramSettingsOptions, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TelegramSettings) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TelegramSettingsOptions and assigns it to the Options field.
func (o *TelegramSettings) SetOptions(v TelegramSettingsOptions) {
	o.Options = &v
}

func (o TelegramSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelegramSettings) UnmarshalJSON(bytes []byte) (err error) {
	varTelegramSettings := _TelegramSettings{}

	if err = json.Unmarshal(bytes, &varTelegramSettings); err == nil {
		*o = TelegramSettings(varTelegramSettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "types")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelegramSettings struct {
	value *TelegramSettings
	isSet bool
}

func (v NullableTelegramSettings) Get() *TelegramSettings {
	return v.value
}

func (v *NullableTelegramSettings) Set(val *TelegramSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTelegramSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTelegramSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelegramSettings(val *TelegramSettings) *NullableTelegramSettings {
	return &NullableTelegramSettings{value: val, isSet: true}
}

func (v NullableTelegramSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelegramSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


