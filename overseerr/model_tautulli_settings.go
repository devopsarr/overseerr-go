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

// TautulliSettings struct for TautulliSettings
type TautulliSettings struct {
	Hostname NullableString `json:"hostname,omitempty"`
	Port NullableFloat32 `json:"port,omitempty"`
	UseSsl NullableBool `json:"useSsl,omitempty"`
	ApiKey NullableString `json:"apiKey,omitempty"`
	ExternalUrl NullableString `json:"externalUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TautulliSettings TautulliSettings

// NewTautulliSettings instantiates a new TautulliSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTautulliSettings() *TautulliSettings {
	this := TautulliSettings{}
	return &this
}

// NewTautulliSettingsWithDefaults instantiates a new TautulliSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTautulliSettingsWithDefaults() *TautulliSettings {
	this := TautulliSettings{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TautulliSettings) GetHostname() string {
	if o == nil || isNil(o.Hostname.Get()) {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TautulliSettings) GetHostnameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *TautulliSettings) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *TautulliSettings) SetHostname(v string) {
	o.Hostname.Set(&v)
}
// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *TautulliSettings) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *TautulliSettings) UnsetHostname() {
	o.Hostname.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TautulliSettings) GetPort() float32 {
	if o == nil || isNil(o.Port.Get()) {
		var ret float32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TautulliSettings) GetPortOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TautulliSettings) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableFloat32 and assigns it to the Port field.
func (o *TautulliSettings) SetPort(v float32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TautulliSettings) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TautulliSettings) UnsetPort() {
	o.Port.Unset()
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TautulliSettings) GetUseSsl() bool {
	if o == nil || isNil(o.UseSsl.Get()) {
		var ret bool
		return ret
	}
	return *o.UseSsl.Get()
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TautulliSettings) GetUseSslOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.UseSsl.Get(), o.UseSsl.IsSet()
}

// HasUseSsl returns a boolean if a field has been set.
func (o *TautulliSettings) HasUseSsl() bool {
	if o != nil && o.UseSsl.IsSet() {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given NullableBool and assigns it to the UseSsl field.
func (o *TautulliSettings) SetUseSsl(v bool) {
	o.UseSsl.Set(&v)
}
// SetUseSslNil sets the value for UseSsl to be an explicit nil
func (o *TautulliSettings) SetUseSslNil() {
	o.UseSsl.Set(nil)
}

// UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
func (o *TautulliSettings) UnsetUseSsl() {
	o.UseSsl.Unset()
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TautulliSettings) GetApiKey() string {
	if o == nil || isNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TautulliSettings) GetApiKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *TautulliSettings) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *TautulliSettings) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *TautulliSettings) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *TautulliSettings) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TautulliSettings) GetExternalUrl() string {
	if o == nil || isNil(o.ExternalUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalUrl.Get()
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TautulliSettings) GetExternalUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExternalUrl.Get(), o.ExternalUrl.IsSet()
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *TautulliSettings) HasExternalUrl() bool {
	if o != nil && o.ExternalUrl.IsSet() {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given NullableString and assigns it to the ExternalUrl field.
func (o *TautulliSettings) SetExternalUrl(v string) {
	o.ExternalUrl.Set(&v)
}
// SetExternalUrlNil sets the value for ExternalUrl to be an explicit nil
func (o *TautulliSettings) SetExternalUrlNil() {
	o.ExternalUrl.Set(nil)
}

// UnsetExternalUrl ensures that no value is present for ExternalUrl, not even an explicit nil
func (o *TautulliSettings) UnsetExternalUrl() {
	o.ExternalUrl.Unset()
}

func (o TautulliSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.UseSsl.IsSet() {
		toSerialize["useSsl"] = o.UseSsl.Get()
	}
	if o.ApiKey.IsSet() {
		toSerialize["apiKey"] = o.ApiKey.Get()
	}
	if o.ExternalUrl.IsSet() {
		toSerialize["externalUrl"] = o.ExternalUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TautulliSettings) UnmarshalJSON(bytes []byte) (err error) {
	varTautulliSettings := _TautulliSettings{}

	if err = json.Unmarshal(bytes, &varTautulliSettings); err == nil {
		*o = TautulliSettings(varTautulliSettings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "port")
		delete(additionalProperties, "useSsl")
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "externalUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTautulliSettings struct {
	value *TautulliSettings
	isSet bool
}

func (v NullableTautulliSettings) Get() *TautulliSettings {
	return v.value
}

func (v *NullableTautulliSettings) Set(val *TautulliSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTautulliSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTautulliSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTautulliSettings(val *TautulliSettings) *NullableTautulliSettings {
	return &NullableTautulliSettings{value: val, isSet: true}
}

func (v NullableTautulliSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTautulliSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


