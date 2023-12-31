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

// CreateAuthLogout200Response struct for CreateAuthLogout200Response
type CreateAuthLogout200Response struct {
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthLogout200Response CreateAuthLogout200Response

// NewCreateAuthLogout200Response instantiates a new CreateAuthLogout200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthLogout200Response() *CreateAuthLogout200Response {
	this := CreateAuthLogout200Response{}
	return &this
}

// NewCreateAuthLogout200ResponseWithDefaults instantiates a new CreateAuthLogout200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthLogout200ResponseWithDefaults() *CreateAuthLogout200Response {
	this := CreateAuthLogout200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateAuthLogout200Response) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAuthLogout200Response) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateAuthLogout200Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateAuthLogout200Response) SetStatus(v string) {
	o.Status = &v
}

func (o CreateAuthLogout200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateAuthLogout200Response) UnmarshalJSON(bytes []byte) (err error) {
	varCreateAuthLogout200Response := _CreateAuthLogout200Response{}

	if err = json.Unmarshal(bytes, &varCreateAuthLogout200Response); err == nil {
		*o = CreateAuthLogout200Response(varCreateAuthLogout200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthLogout200Response struct {
	value *CreateAuthLogout200Response
	isSet bool
}

func (v NullableCreateAuthLogout200Response) Get() *CreateAuthLogout200Response {
	return v.value
}

func (v *NullableCreateAuthLogout200Response) Set(val *CreateAuthLogout200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthLogout200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthLogout200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthLogout200Response(val *CreateAuthLogout200Response) *NullableCreateAuthLogout200Response {
	return &NullableCreateAuthLogout200Response{value: val, isSet: true}
}

func (v NullableCreateAuthLogout200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthLogout200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


