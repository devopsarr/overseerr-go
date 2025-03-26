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

// checks if the GotifySettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GotifySettingsOptions{}

// GotifySettingsOptions struct for GotifySettingsOptions
type GotifySettingsOptions struct {
	Url *string `json:"url,omitempty"`
	Token *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GotifySettingsOptions GotifySettingsOptions

// NewGotifySettingsOptions instantiates a new GotifySettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGotifySettingsOptions() *GotifySettingsOptions {
	this := GotifySettingsOptions{}
	return &this
}

// NewGotifySettingsOptionsWithDefaults instantiates a new GotifySettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGotifySettingsOptionsWithDefaults() *GotifySettingsOptions {
	this := GotifySettingsOptions{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GotifySettingsOptions) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GotifySettingsOptions) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GotifySettingsOptions) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GotifySettingsOptions) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GotifySettingsOptions) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GotifySettingsOptions) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GotifySettingsOptions) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GotifySettingsOptions) SetToken(v string) {
	o.Token = &v
}

func (o GotifySettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GotifySettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GotifySettingsOptions) UnmarshalJSON(data []byte) (err error) {
	varGotifySettingsOptions := _GotifySettingsOptions{}

	err = json.Unmarshal(data, &varGotifySettingsOptions)

	if err != nil {
		return err
	}

	*o = GotifySettingsOptions(varGotifySettingsOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGotifySettingsOptions struct {
	value *GotifySettingsOptions
	isSet bool
}

func (v NullableGotifySettingsOptions) Get() *GotifySettingsOptions {
	return v.value
}

func (v *NullableGotifySettingsOptions) Set(val *GotifySettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGotifySettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGotifySettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGotifySettingsOptions(val *GotifySettingsOptions) *NullableGotifySettingsOptions {
	return &NullableGotifySettingsOptions{value: val, isSet: true}
}

func (v NullableGotifySettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGotifySettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


