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

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	Email *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	Permissions *float32 `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserRequest CreateUserRequest

// NewCreateUserRequest instantiates a new CreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequest() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestWithDefaults() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateUserRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateUserRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CreateUserRequest) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
    return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateUserRequest) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CreateUserRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *CreateUserRequest) GetPermissions() float32 {
	if o == nil || isNil(o.Permissions) {
		var ret float32
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetPermissionsOk() (*float32, bool) {
	if o == nil || isNil(o.Permissions) {
    return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *CreateUserRequest) HasPermissions() bool {
	if o != nil && !isNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given float32 and assigns it to the Permissions field.
func (o *CreateUserRequest) SetPermissions(v float32) {
	o.Permissions = &v
}

func (o CreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUserRequest := _CreateUserRequest{}

	if err = json.Unmarshal(bytes, &varCreateUserRequest); err == nil {
		*o = CreateUserRequest(varCreateUserRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "username")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserRequest struct {
	value *CreateUserRequest
	isSet bool
}

func (v NullableCreateUserRequest) Get() *CreateUserRequest {
	return v.value
}

func (v *NullableCreateUserRequest) Set(val *CreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequest(val *CreateUserRequest) *NullableCreateUserRequest {
	return &NullableCreateUserRequest{value: val, isSet: true}
}

func (v NullableCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


