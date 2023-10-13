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

// TestRadarr200Response struct for TestRadarr200Response
type TestRadarr200Response struct {
	Profiles []*ServiceProfile `json:"profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestRadarr200Response TestRadarr200Response

// NewTestRadarr200Response instantiates a new TestRadarr200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRadarr200Response() *TestRadarr200Response {
	this := TestRadarr200Response{}
	return &this
}

// NewTestRadarr200ResponseWithDefaults instantiates a new TestRadarr200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRadarr200ResponseWithDefaults() *TestRadarr200Response {
	this := TestRadarr200Response{}
	return &this
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *TestRadarr200Response) GetProfiles() []*ServiceProfile {
	if o == nil || isNil(o.Profiles) {
		var ret []*ServiceProfile
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRadarr200Response) GetProfilesOk() ([]*ServiceProfile, bool) {
	if o == nil || isNil(o.Profiles) {
    return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *TestRadarr200Response) HasProfiles() bool {
	if o != nil && !isNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []ServiceProfile and assigns it to the Profiles field.
func (o *TestRadarr200Response) SetProfiles(v []*ServiceProfile) {
	o.Profiles = v
}

func (o TestRadarr200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestRadarr200Response) UnmarshalJSON(bytes []byte) (err error) {
	varTestRadarr200Response := _TestRadarr200Response{}

	if err = json.Unmarshal(bytes, &varTestRadarr200Response); err == nil {
		*o = TestRadarr200Response(varTestRadarr200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestRadarr200Response struct {
	value *TestRadarr200Response
	isSet bool
}

func (v NullableTestRadarr200Response) Get() *TestRadarr200Response {
	return v.value
}

func (v *NullableTestRadarr200Response) Set(val *TestRadarr200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRadarr200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRadarr200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRadarr200Response(val *TestRadarr200Response) *NullableTestRadarr200Response {
	return &NullableTestRadarr200Response{value: val, isSet: true}
}

func (v NullableTestRadarr200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRadarr200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


