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

// GetUserSettingsPassword200Response struct for GetUserSettingsPassword200Response
type GetUserSettingsPassword200Response struct {
	HasPassword *bool `json:"hasPassword,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserSettingsPassword200Response GetUserSettingsPassword200Response

// NewGetUserSettingsPassword200Response instantiates a new GetUserSettingsPassword200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserSettingsPassword200Response() *GetUserSettingsPassword200Response {
	this := GetUserSettingsPassword200Response{}
	return &this
}

// NewGetUserSettingsPassword200ResponseWithDefaults instantiates a new GetUserSettingsPassword200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserSettingsPassword200ResponseWithDefaults() *GetUserSettingsPassword200Response {
	this := GetUserSettingsPassword200Response{}
	return &this
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *GetUserSettingsPassword200Response) GetHasPassword() bool {
	if o == nil || isNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSettingsPassword200Response) GetHasPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.HasPassword) {
    return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GetUserSettingsPassword200Response) HasHasPassword() bool {
	if o != nil && !isNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *GetUserSettingsPassword200Response) SetHasPassword(v bool) {
	o.HasPassword = &v
}

func (o GetUserSettingsPassword200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetUserSettingsPassword200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserSettingsPassword200Response := _GetUserSettingsPassword200Response{}

	if err = json.Unmarshal(bytes, &varGetUserSettingsPassword200Response); err == nil {
		*o = GetUserSettingsPassword200Response(varGetUserSettingsPassword200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "hasPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserSettingsPassword200Response struct {
	value *GetUserSettingsPassword200Response
	isSet bool
}

func (v NullableGetUserSettingsPassword200Response) Get() *GetUserSettingsPassword200Response {
	return v.value
}

func (v *NullableGetUserSettingsPassword200Response) Set(val *GetUserSettingsPassword200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserSettingsPassword200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserSettingsPassword200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserSettingsPassword200Response(val *GetUserSettingsPassword200Response) *NullableGetUserSettingsPassword200Response {
	return &NullableGetUserSettingsPassword200Response{value: val, isSet: true}
}

func (v NullableGetUserSettingsPassword200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserSettingsPassword200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

