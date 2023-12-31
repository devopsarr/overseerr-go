/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
	"fmt"
)

// CreateUserRegisterPushSubscriptionRequest struct for CreateUserRegisterPushSubscriptionRequest
type CreateUserRegisterPushSubscriptionRequest struct {
	Endpoint string `json:"endpoint"`
	Auth string `json:"auth"`
	P256dh string `json:"p256dh"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserRegisterPushSubscriptionRequest CreateUserRegisterPushSubscriptionRequest

// NewCreateUserRegisterPushSubscriptionRequest instantiates a new CreateUserRegisterPushSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRegisterPushSubscriptionRequest(endpoint string, auth string, p256dh string) *CreateUserRegisterPushSubscriptionRequest {
	this := CreateUserRegisterPushSubscriptionRequest{}
	this.Endpoint = endpoint
	this.Auth = auth
	this.P256dh = p256dh
	return &this
}

// NewCreateUserRegisterPushSubscriptionRequestWithDefaults instantiates a new CreateUserRegisterPushSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRegisterPushSubscriptionRequestWithDefaults() *CreateUserRegisterPushSubscriptionRequest {
	this := CreateUserRegisterPushSubscriptionRequest{}
	return &this
}

// GetEndpoint returns the Endpoint field value
func (o *CreateUserRegisterPushSubscriptionRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CreateUserRegisterPushSubscriptionRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CreateUserRegisterPushSubscriptionRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetAuth returns the Auth field value
func (o *CreateUserRegisterPushSubscriptionRequest) GetAuth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CreateUserRegisterPushSubscriptionRequest) GetAuthOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value
func (o *CreateUserRegisterPushSubscriptionRequest) SetAuth(v string) {
	o.Auth = v
}

// GetP256dh returns the P256dh field value
func (o *CreateUserRegisterPushSubscriptionRequest) GetP256dh() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.P256dh
}

// GetP256dhOk returns a tuple with the P256dh field value
// and a boolean to check if the value has been set.
func (o *CreateUserRegisterPushSubscriptionRequest) GetP256dhOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.P256dh, true
}

// SetP256dh sets field value
func (o *CreateUserRegisterPushSubscriptionRequest) SetP256dh(v string) {
	o.P256dh = v
}

func (o CreateUserRegisterPushSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["endpoint"] = o.Endpoint
	}
	if true {
		toSerialize["auth"] = o.Auth
	}
	if true {
		toSerialize["p256dh"] = o.P256dh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUserRegisterPushSubscriptionRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUserRegisterPushSubscriptionRequest := _CreateUserRegisterPushSubscriptionRequest{}

	if err = json.Unmarshal(bytes, &varCreateUserRegisterPushSubscriptionRequest); err == nil {
		*o = CreateUserRegisterPushSubscriptionRequest(varCreateUserRegisterPushSubscriptionRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "auth")
		delete(additionalProperties, "p256dh")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserRegisterPushSubscriptionRequest struct {
	value *CreateUserRegisterPushSubscriptionRequest
	isSet bool
}

func (v NullableCreateUserRegisterPushSubscriptionRequest) Get() *CreateUserRegisterPushSubscriptionRequest {
	return v.value
}

func (v *NullableCreateUserRegisterPushSubscriptionRequest) Set(val *CreateUserRegisterPushSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRegisterPushSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRegisterPushSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRegisterPushSubscriptionRequest(val *CreateUserRegisterPushSubscriptionRequest) *NullableCreateUserRegisterPushSubscriptionRequest {
	return &NullableCreateUserRegisterPushSubscriptionRequest{value: val, isSet: true}
}

func (v NullableCreateUserRegisterPushSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRegisterPushSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


