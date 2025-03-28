/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateUserSettingsPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserSettingsPasswordRequest{}

// CreateUserSettingsPasswordRequest struct for CreateUserSettingsPasswordRequest
type CreateUserSettingsPasswordRequest struct {
	CurrentPassword NullableString `json:"currentPassword,omitempty"`
	NewPassword string `json:"newPassword"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserSettingsPasswordRequest CreateUserSettingsPasswordRequest

// NewCreateUserSettingsPasswordRequest instantiates a new CreateUserSettingsPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserSettingsPasswordRequest(newPassword string) *CreateUserSettingsPasswordRequest {
	this := CreateUserSettingsPasswordRequest{}
	this.NewPassword = newPassword
	return &this
}

// NewCreateUserSettingsPasswordRequestWithDefaults instantiates a new CreateUserSettingsPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserSettingsPasswordRequestWithDefaults() *CreateUserSettingsPasswordRequest {
	this := CreateUserSettingsPasswordRequest{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserSettingsPasswordRequest) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentPassword.Get()
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserSettingsPasswordRequest) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPassword.Get(), o.CurrentPassword.IsSet()
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *CreateUserSettingsPasswordRequest) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword.IsSet() {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given NullableString and assigns it to the CurrentPassword field.
func (o *CreateUserSettingsPasswordRequest) SetCurrentPassword(v string) {
	o.CurrentPassword.Set(&v)
}
// SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil
func (o *CreateUserSettingsPasswordRequest) SetCurrentPasswordNil() {
	o.CurrentPassword.Set(nil)
}

// UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
func (o *CreateUserSettingsPasswordRequest) UnsetCurrentPassword() {
	o.CurrentPassword.Unset()
}

// GetNewPassword returns the NewPassword field value
func (o *CreateUserSettingsPasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *CreateUserSettingsPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *CreateUserSettingsPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
}

func (o CreateUserSettingsPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserSettingsPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPassword.IsSet() {
		toSerialize["currentPassword"] = o.CurrentPassword.Get()
	}
	toSerialize["newPassword"] = o.NewPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateUserSettingsPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPassword",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateUserSettingsPasswordRequest := _CreateUserSettingsPasswordRequest{}

	err = json.Unmarshal(data, &varCreateUserSettingsPasswordRequest)

	if err != nil {
		return err
	}

	*o = CreateUserSettingsPasswordRequest(varCreateUserSettingsPasswordRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currentPassword")
		delete(additionalProperties, "newPassword")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserSettingsPasswordRequest struct {
	value *CreateUserSettingsPasswordRequest
	isSet bool
}

func (v NullableCreateUserSettingsPasswordRequest) Get() *CreateUserSettingsPasswordRequest {
	return v.value
}

func (v *NullableCreateUserSettingsPasswordRequest) Set(val *CreateUserSettingsPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserSettingsPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserSettingsPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserSettingsPasswordRequest(val *CreateUserSettingsPasswordRequest) *NullableCreateUserSettingsPasswordRequest {
	return &NullableCreateUserSettingsPasswordRequest{value: val, isSet: true}
}

func (v NullableCreateUserSettingsPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserSettingsPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


