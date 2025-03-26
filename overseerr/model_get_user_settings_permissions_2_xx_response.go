/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the GetUserSettingsPermissions2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserSettingsPermissions2XXResponse{}

// GetUserSettingsPermissions2XXResponse struct for GetUserSettingsPermissions2XXResponse
type GetUserSettingsPermissions2XXResponse struct {
	Permissions *float32 `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserSettingsPermissions2XXResponse GetUserSettingsPermissions2XXResponse

// NewGetUserSettingsPermissions2XXResponse instantiates a new GetUserSettingsPermissions2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserSettingsPermissions2XXResponse() *GetUserSettingsPermissions2XXResponse {
	this := GetUserSettingsPermissions2XXResponse{}
	return &this
}

// NewGetUserSettingsPermissions2XXResponseWithDefaults instantiates a new GetUserSettingsPermissions2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserSettingsPermissions2XXResponseWithDefaults() *GetUserSettingsPermissions2XXResponse {
	this := GetUserSettingsPermissions2XXResponse{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *GetUserSettingsPermissions2XXResponse) GetPermissions() float32 {
	if o == nil || IsNil(o.Permissions) {
		var ret float32
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSettingsPermissions2XXResponse) GetPermissionsOk() (*float32, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *GetUserSettingsPermissions2XXResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given float32 and assigns it to the Permissions field.
func (o *GetUserSettingsPermissions2XXResponse) SetPermissions(v float32) {
	o.Permissions = &v
}

func (o GetUserSettingsPermissions2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserSettingsPermissions2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserSettingsPermissions2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUserSettingsPermissions2XXResponse := _GetUserSettingsPermissions2XXResponse{}

	err = json.Unmarshal(data, &varGetUserSettingsPermissions2XXResponse)

	if err != nil {
		return err
	}

	*o = GetUserSettingsPermissions2XXResponse(varGetUserSettingsPermissions2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserSettingsPermissions2XXResponse struct {
	value *GetUserSettingsPermissions2XXResponse
	isSet bool
}

func (v NullableGetUserSettingsPermissions2XXResponse) Get() *GetUserSettingsPermissions2XXResponse {
	return v.value
}

func (v *NullableGetUserSettingsPermissions2XXResponse) Set(val *GetUserSettingsPermissions2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserSettingsPermissions2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserSettingsPermissions2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserSettingsPermissions2XXResponse(val *GetUserSettingsPermissions2XXResponse) *NullableGetUserSettingsPermissions2XXResponse {
	return &NullableGetUserSettingsPermissions2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserSettingsPermissions2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserSettingsPermissions2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


