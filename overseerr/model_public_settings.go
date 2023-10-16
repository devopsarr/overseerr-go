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

// PublicSettings struct for PublicSettings
type PublicSettings struct {
	Initialized *bool `json:"initialized,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicSettings PublicSettings

// NewPublicSettings instantiates a new PublicSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSettings() *PublicSettings {
	this := PublicSettings{}
	return &this
}

// NewPublicSettingsWithDefaults instantiates a new PublicSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSettingsWithDefaults() *PublicSettings {
	this := PublicSettings{}
	return &this
}

// GetInitialized returns the Initialized field value if set, zero value otherwise.
func (o *PublicSettings) GetInitialized() bool {
	if o == nil || isNil(o.Initialized) {
		var ret bool
		return ret
	}
	return *o.Initialized
}

// GetInitializedOk returns a tuple with the Initialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSettings) GetInitializedOk() (*bool, bool) {
	if o == nil || isNil(o.Initialized) {
    return nil, false
	}
	return o.Initialized, true
}

// HasInitialized returns a boolean if a field has been set.
func (o *PublicSettings) HasInitialized() bool {
	if o != nil && !isNil(o.Initialized) {
		return true
	}

	return false
}

// SetInitialized gets a reference to the given bool and assigns it to the Initialized field.
func (o *PublicSettings) SetInitialized(v bool) {
	o.Initialized = &v
}

func (o PublicSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Initialized) {
		toSerialize["initialized"] = o.Initialized
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PublicSettings) UnmarshalJSON(bytes []byte) (err error) {
	varPublicSettings := _PublicSettings{}

	if err = json.Unmarshal(bytes, &varPublicSettings); err == nil {
		*o = PublicSettings(varPublicSettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "initialized")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicSettings struct {
	value *PublicSettings
	isSet bool
}

func (v NullablePublicSettings) Get() *PublicSettings {
	return v.value
}

func (v *NullablePublicSettings) Set(val *PublicSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSettings(val *PublicSettings) *NullablePublicSettings {
	return &NullablePublicSettings{value: val, isSet: true}
}

func (v NullablePublicSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

