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

// checks if the WebhookSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSettingsOptions{}

// WebhookSettingsOptions struct for WebhookSettingsOptions
type WebhookSettingsOptions struct {
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	AuthHeader *string `json:"authHeader,omitempty"`
	JsonPayload *string `json:"jsonPayload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookSettingsOptions WebhookSettingsOptions

// NewWebhookSettingsOptions instantiates a new WebhookSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSettingsOptions() *WebhookSettingsOptions {
	this := WebhookSettingsOptions{}
	return &this
}

// NewWebhookSettingsOptionsWithDefaults instantiates a new WebhookSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSettingsOptionsWithDefaults() *WebhookSettingsOptions {
	this := WebhookSettingsOptions{}
	return &this
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *WebhookSettingsOptions) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettingsOptions) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *WebhookSettingsOptions) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *WebhookSettingsOptions) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetAuthHeader returns the AuthHeader field value if set, zero value otherwise.
func (o *WebhookSettingsOptions) GetAuthHeader() string {
	if o == nil || IsNil(o.AuthHeader) {
		var ret string
		return ret
	}
	return *o.AuthHeader
}

// GetAuthHeaderOk returns a tuple with the AuthHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettingsOptions) GetAuthHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.AuthHeader) {
		return nil, false
	}
	return o.AuthHeader, true
}

// HasAuthHeader returns a boolean if a field has been set.
func (o *WebhookSettingsOptions) HasAuthHeader() bool {
	if o != nil && !IsNil(o.AuthHeader) {
		return true
	}

	return false
}

// SetAuthHeader gets a reference to the given string and assigns it to the AuthHeader field.
func (o *WebhookSettingsOptions) SetAuthHeader(v string) {
	o.AuthHeader = &v
}

// GetJsonPayload returns the JsonPayload field value if set, zero value otherwise.
func (o *WebhookSettingsOptions) GetJsonPayload() string {
	if o == nil || IsNil(o.JsonPayload) {
		var ret string
		return ret
	}
	return *o.JsonPayload
}

// GetJsonPayloadOk returns a tuple with the JsonPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSettingsOptions) GetJsonPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.JsonPayload) {
		return nil, false
	}
	return o.JsonPayload, true
}

// HasJsonPayload returns a boolean if a field has been set.
func (o *WebhookSettingsOptions) HasJsonPayload() bool {
	if o != nil && !IsNil(o.JsonPayload) {
		return true
	}

	return false
}

// SetJsonPayload gets a reference to the given string and assigns it to the JsonPayload field.
func (o *WebhookSettingsOptions) SetJsonPayload(v string) {
	o.JsonPayload = &v
}

func (o WebhookSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	if !IsNil(o.AuthHeader) {
		toSerialize["authHeader"] = o.AuthHeader
	}
	if !IsNil(o.JsonPayload) {
		toSerialize["jsonPayload"] = o.JsonPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookSettingsOptions) UnmarshalJSON(data []byte) (err error) {
	varWebhookSettingsOptions := _WebhookSettingsOptions{}

	err = json.Unmarshal(data, &varWebhookSettingsOptions)

	if err != nil {
		return err
	}

	*o = WebhookSettingsOptions(varWebhookSettingsOptions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "webhookUrl")
		delete(additionalProperties, "authHeader")
		delete(additionalProperties, "jsonPayload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookSettingsOptions struct {
	value *WebhookSettingsOptions
	isSet bool
}

func (v NullableWebhookSettingsOptions) Get() *WebhookSettingsOptions {
	return v.value
}

func (v *NullableWebhookSettingsOptions) Set(val *WebhookSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSettingsOptions(val *WebhookSettingsOptions) *NullableWebhookSettingsOptions {
	return &NullableWebhookSettingsOptions{value: val, isSet: true}
}

func (v NullableWebhookSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


