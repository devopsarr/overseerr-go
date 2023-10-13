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

// CreateAuthResetPasswordRequest struct for CreateAuthResetPasswordRequest
type CreateAuthResetPasswordRequest struct {
	Email string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthResetPasswordRequest CreateAuthResetPasswordRequest

// NewCreateAuthResetPasswordRequest instantiates a new CreateAuthResetPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthResetPasswordRequest(email string) *CreateAuthResetPasswordRequest {
	this := CreateAuthResetPasswordRequest{}
	this.Email = email
	return &this
}

// NewCreateAuthResetPasswordRequestWithDefaults instantiates a new CreateAuthResetPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthResetPasswordRequestWithDefaults() *CreateAuthResetPasswordRequest {
	this := CreateAuthResetPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateAuthResetPasswordRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateAuthResetPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateAuthResetPasswordRequest) SetEmail(v string) {
	o.Email = v
}

func (o CreateAuthResetPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateAuthResetPasswordRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateAuthResetPasswordRequest := _CreateAuthResetPasswordRequest{}

	if err = json.Unmarshal(bytes, &varCreateAuthResetPasswordRequest); err == nil {
		*o = CreateAuthResetPasswordRequest(varCreateAuthResetPasswordRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthResetPasswordRequest struct {
	value *CreateAuthResetPasswordRequest
	isSet bool
}

func (v NullableCreateAuthResetPasswordRequest) Get() *CreateAuthResetPasswordRequest {
	return v.value
}

func (v *NullableCreateAuthResetPasswordRequest) Set(val *CreateAuthResetPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthResetPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthResetPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthResetPasswordRequest(val *CreateAuthResetPasswordRequest) *NullableCreateAuthResetPasswordRequest {
	return &NullableCreateAuthResetPasswordRequest{value: val, isSet: true}
}

func (v NullableCreateAuthResetPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthResetPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


