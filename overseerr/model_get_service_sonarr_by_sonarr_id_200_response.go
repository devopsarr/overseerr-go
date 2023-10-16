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

// GetServiceSonarrBySonarrId200Response struct for GetServiceSonarrBySonarrId200Response
type GetServiceSonarrBySonarrId200Response struct {
	Server *SonarrSettings `json:"server,omitempty"`
	Profiles *ServiceProfile `json:"profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetServiceSonarrBySonarrId200Response GetServiceSonarrBySonarrId200Response

// NewGetServiceSonarrBySonarrId200Response instantiates a new GetServiceSonarrBySonarrId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceSonarrBySonarrId200Response() *GetServiceSonarrBySonarrId200Response {
	this := GetServiceSonarrBySonarrId200Response{}
	return &this
}

// NewGetServiceSonarrBySonarrId200ResponseWithDefaults instantiates a new GetServiceSonarrBySonarrId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceSonarrBySonarrId200ResponseWithDefaults() *GetServiceSonarrBySonarrId200Response {
	this := GetServiceSonarrBySonarrId200Response{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *GetServiceSonarrBySonarrId200Response) GetServer() SonarrSettings {
	if o == nil || isNil(o.Server) {
		var ret SonarrSettings
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceSonarrBySonarrId200Response) GetServerOk() (*SonarrSettings, bool) {
	if o == nil || isNil(o.Server) {
    return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *GetServiceSonarrBySonarrId200Response) HasServer() bool {
	if o != nil && !isNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given SonarrSettings and assigns it to the Server field.
func (o *GetServiceSonarrBySonarrId200Response) SetServer(v SonarrSettings) {
	o.Server = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *GetServiceSonarrBySonarrId200Response) GetProfiles() ServiceProfile {
	if o == nil || isNil(o.Profiles) {
		var ret ServiceProfile
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceSonarrBySonarrId200Response) GetProfilesOk() (*ServiceProfile, bool) {
	if o == nil || isNil(o.Profiles) {
    return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *GetServiceSonarrBySonarrId200Response) HasProfiles() bool {
	if o != nil && !isNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given ServiceProfile and assigns it to the Profiles field.
func (o *GetServiceSonarrBySonarrId200Response) SetProfiles(v ServiceProfile) {
	o.Profiles = &v
}

func (o GetServiceSonarrBySonarrId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !isNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetServiceSonarrBySonarrId200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetServiceSonarrBySonarrId200Response := _GetServiceSonarrBySonarrId200Response{}

	if err = json.Unmarshal(bytes, &varGetServiceSonarrBySonarrId200Response); err == nil {
		*o = GetServiceSonarrBySonarrId200Response(varGetServiceSonarrBySonarrId200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "server")
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetServiceSonarrBySonarrId200Response struct {
	value *GetServiceSonarrBySonarrId200Response
	isSet bool
}

func (v NullableGetServiceSonarrBySonarrId200Response) Get() *GetServiceSonarrBySonarrId200Response {
	return v.value
}

func (v *NullableGetServiceSonarrBySonarrId200Response) Set(val *GetServiceSonarrBySonarrId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceSonarrBySonarrId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceSonarrBySonarrId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceSonarrBySonarrId200Response(val *GetServiceSonarrBySonarrId200Response) *NullableGetServiceSonarrBySonarrId200Response {
	return &NullableGetServiceSonarrBySonarrId200Response{value: val, isSet: true}
}

func (v NullableGetServiceSonarrBySonarrId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceSonarrBySonarrId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

