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

// CreateAuthResetPasswordByGuidRequest struct for CreateAuthResetPasswordByGuidRequest
type CreateAuthResetPasswordByGuidRequest struct {
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthResetPasswordByGuidRequest CreateAuthResetPasswordByGuidRequest

// NewCreateAuthResetPasswordByGuidRequest instantiates a new CreateAuthResetPasswordByGuidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthResetPasswordByGuidRequest(password string) *CreateAuthResetPasswordByGuidRequest {
	this := CreateAuthResetPasswordByGuidRequest{}
	this.Password = password
	return &this
}

// NewCreateAuthResetPasswordByGuidRequestWithDefaults instantiates a new CreateAuthResetPasswordByGuidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthResetPasswordByGuidRequestWithDefaults() *CreateAuthResetPasswordByGuidRequest {
	this := CreateAuthResetPasswordByGuidRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *CreateAuthResetPasswordByGuidRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateAuthResetPasswordByGuidRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateAuthResetPasswordByGuidRequest) SetPassword(v string) {
	o.Password = v
}

func (o CreateAuthResetPasswordByGuidRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateAuthResetPasswordByGuidRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateAuthResetPasswordByGuidRequest := _CreateAuthResetPasswordByGuidRequest{}

	if err = json.Unmarshal(bytes, &varCreateAuthResetPasswordByGuidRequest); err == nil {
		*o = CreateAuthResetPasswordByGuidRequest(varCreateAuthResetPasswordByGuidRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthResetPasswordByGuidRequest struct {
	value *CreateAuthResetPasswordByGuidRequest
	isSet bool
}

func (v NullableCreateAuthResetPasswordByGuidRequest) Get() *CreateAuthResetPasswordByGuidRequest {
	return v.value
}

func (v *NullableCreateAuthResetPasswordByGuidRequest) Set(val *CreateAuthResetPasswordByGuidRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthResetPasswordByGuidRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthResetPasswordByGuidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthResetPasswordByGuidRequest(val *CreateAuthResetPasswordByGuidRequest) *NullableCreateAuthResetPasswordByGuidRequest {
	return &NullableCreateAuthResetPasswordByGuidRequest{value: val, isSet: true}
}

func (v NullableCreateAuthResetPasswordByGuidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthResetPasswordByGuidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

