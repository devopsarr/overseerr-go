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

// GetUserSettingsMain2XXResponse struct for GetUserSettingsMain2XXResponse
type GetUserSettingsMain2XXResponse struct {
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserSettingsMain2XXResponse GetUserSettingsMain2XXResponse

// NewGetUserSettingsMain2XXResponse instantiates a new GetUserSettingsMain2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserSettingsMain2XXResponse() *GetUserSettingsMain2XXResponse {
	this := GetUserSettingsMain2XXResponse{}
	return &this
}

// NewGetUserSettingsMain2XXResponseWithDefaults instantiates a new GetUserSettingsMain2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserSettingsMain2XXResponseWithDefaults() *GetUserSettingsMain2XXResponse {
	this := GetUserSettingsMain2XXResponse{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetUserSettingsMain2XXResponse) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSettingsMain2XXResponse) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetUserSettingsMain2XXResponse) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetUserSettingsMain2XXResponse) SetUsername(v string) {
	o.Username = &v
}

func (o GetUserSettingsMain2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetUserSettingsMain2XXResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserSettingsMain2XXResponse := _GetUserSettingsMain2XXResponse{}

	if err = json.Unmarshal(bytes, &varGetUserSettingsMain2XXResponse); err == nil {
		*o = GetUserSettingsMain2XXResponse(varGetUserSettingsMain2XXResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserSettingsMain2XXResponse struct {
	value *GetUserSettingsMain2XXResponse
	isSet bool
}

func (v NullableGetUserSettingsMain2XXResponse) Get() *GetUserSettingsMain2XXResponse {
	return v.value
}

func (v *NullableGetUserSettingsMain2XXResponse) Set(val *GetUserSettingsMain2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserSettingsMain2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserSettingsMain2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserSettingsMain2XXResponse(val *GetUserSettingsMain2XXResponse) *NullableGetUserSettingsMain2XXResponse {
	return &NullableGetUserSettingsMain2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserSettingsMain2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserSettingsMain2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

