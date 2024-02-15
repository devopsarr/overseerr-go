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

// checks if the LunaSeaSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LunaSeaSettings{}

// LunaSeaSettings struct for LunaSeaSettings
type LunaSeaSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
	Types *float32 `json:"types,omitempty"`
	Options *LunaSeaSettingsOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LunaSeaSettings LunaSeaSettings

// NewLunaSeaSettings instantiates a new LunaSeaSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLunaSeaSettings() *LunaSeaSettings {
	this := LunaSeaSettings{}
	return &this
}

// NewLunaSeaSettingsWithDefaults instantiates a new LunaSeaSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLunaSeaSettingsWithDefaults() *LunaSeaSettings {
	this := LunaSeaSettings{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LunaSeaSettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LunaSeaSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LunaSeaSettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LunaSeaSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *LunaSeaSettings) GetTypes() float32 {
	if o == nil || IsNil(o.Types) {
		var ret float32
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LunaSeaSettings) GetTypesOk() (*float32, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *LunaSeaSettings) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given float32 and assigns it to the Types field.
func (o *LunaSeaSettings) SetTypes(v float32) {
	o.Types = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *LunaSeaSettings) GetOptions() LunaSeaSettingsOptions {
	if o == nil || IsNil(o.Options) {
		var ret LunaSeaSettingsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LunaSeaSettings) GetOptionsOk() (*LunaSeaSettingsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *LunaSeaSettings) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given LunaSeaSettingsOptions and assigns it to the Options field.
func (o *LunaSeaSettings) SetOptions(v LunaSeaSettingsOptions) {
	o.Options = &v
}

func (o LunaSeaSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LunaSeaSettings) ToMap() (map[string]interface{}, error) {
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

func (o *LunaSeaSettings) UnmarshalJSON(data []byte) (err error) {
	varLunaSeaSettings := _LunaSeaSettings{}

	err = json.Unmarshal(data, &varLunaSeaSettings)

	if err != nil {
		return err
	}

	*o = LunaSeaSettings(varLunaSeaSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "types")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLunaSeaSettings struct {
	value *LunaSeaSettings
	isSet bool
}

func (v NullableLunaSeaSettings) Get() *LunaSeaSettings {
	return v.value
}

func (v *NullableLunaSeaSettings) Set(val *LunaSeaSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLunaSeaSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLunaSeaSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLunaSeaSettings(val *LunaSeaSettings) *NullableLunaSeaSettings {
	return &NullableLunaSeaSettings{value: val, isSet: true}
}

func (v NullableLunaSeaSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLunaSeaSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


