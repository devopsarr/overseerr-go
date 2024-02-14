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

// PushoverSettings struct for PushoverSettings
type PushoverSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Types *float32 `json:"types,omitempty"`
	Options *PushoverSettingsOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushoverSettings PushoverSettings

// NewPushoverSettings instantiates a new PushoverSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushoverSettings() *PushoverSettings {
	this := PushoverSettings{}
	return &this
}

// NewPushoverSettingsWithDefaults instantiates a new PushoverSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushoverSettingsWithDefaults() *PushoverSettings {
	this := PushoverSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PushoverSettings) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushoverSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PushoverSettings) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PushoverSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *PushoverSettings) GetTypes() float32 {
	if o == nil || isNil(o.Types) {
		var ret float32
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushoverSettings) GetTypesOk() (*float32, bool) {
	if o == nil || isNil(o.Types) {
    return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *PushoverSettings) HasTypes() bool {
	if o != nil && !isNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given float32 and assigns it to the Types field.
func (o *PushoverSettings) SetTypes(v float32) {
	o.Types = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PushoverSettings) GetOptions() PushoverSettingsOptions {
	if o == nil || isNil(o.Options) {
		var ret PushoverSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushoverSettings) GetOptionsOk() (*PushoverSettingsOptions, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PushoverSettings) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given PushoverSettingsOptions and assigns it to the Options field.
func (o *PushoverSettings) SetOptions(v PushoverSettingsOptions) {
	o.Options = &v
}

func (o PushoverSettings) MarshalJSON() ([]byte, error) {
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

func (o *PushoverSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPushoverSettings := _PushoverSettings{}

	if err = json.Unmarshal(bytes, &varPushoverSettings); err == nil {
		*o = PushoverSettings(varPushoverSettings)
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

type NullablePushoverSettings struct {
	value *PushoverSettings
	isSet bool
}

func (v NullablePushoverSettings) Get() *PushoverSettings {
	return v.value
}

func (v *NullablePushoverSettings) Set(val *PushoverSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePushoverSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePushoverSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushoverSettings(val *PushoverSettings) *NullablePushoverSettings {
	return &NullablePushoverSettings{value: val, isSet: true}
}

func (v NullablePushoverSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushoverSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


