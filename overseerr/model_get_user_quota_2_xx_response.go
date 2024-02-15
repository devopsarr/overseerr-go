/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.33.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the GetUserQuota2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserQuota2XXResponse{}

// GetUserQuota2XXResponse struct for GetUserQuota2XXResponse
type GetUserQuota2XXResponse struct {
	Movie *GetUserQuota2XXResponseMovie `json:"movie,omitempty"`
	Tv *GetUserQuota2XXResponseMovie `json:"tv,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserQuota2XXResponse GetUserQuota2XXResponse

// NewGetUserQuota2XXResponse instantiates a new GetUserQuota2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserQuota2XXResponse() *GetUserQuota2XXResponse {
	this := GetUserQuota2XXResponse{}
	return &this
}

// NewGetUserQuota2XXResponseWithDefaults instantiates a new GetUserQuota2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserQuota2XXResponseWithDefaults() *GetUserQuota2XXResponse {
	this := GetUserQuota2XXResponse{}
	return &this
}

// GetMovie returns the Movie field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponse) GetMovie() GetUserQuota2XXResponseMovie {
	if o == nil || IsNil(o.Movie) {
		var ret GetUserQuota2XXResponseMovie
		return ret
	}
	return *o.Movie
}

// GetMovieOk returns a tuple with the Movie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponse) GetMovieOk() (*GetUserQuota2XXResponseMovie, bool) {
	if o == nil || IsNil(o.Movie) {
		return nil, false
	}
	return o.Movie, true
}

// HasMovie returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponse) HasMovie() bool {
	if o != nil && !IsNil(o.Movie) {
		return true
	}

	return false
}

// SetMovie gets a reference to the given GetUserQuota2XXResponseMovie and assigns it to the Movie field.
func (o *GetUserQuota2XXResponse) SetMovie(v GetUserQuota2XXResponseMovie) {
	o.Movie = &v
}

// GetTv returns the Tv field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponse) GetTv() GetUserQuota2XXResponseMovie {
	if o == nil || IsNil(o.Tv) {
		var ret GetUserQuota2XXResponseMovie
		return ret
	}
	return *o.Tv
}

// GetTvOk returns a tuple with the Tv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponse) GetTvOk() (*GetUserQuota2XXResponseMovie, bool) {
	if o == nil || IsNil(o.Tv) {
		return nil, false
	}
	return o.Tv, true
}

// HasTv returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponse) HasTv() bool {
	if o != nil && !IsNil(o.Tv) {
		return true
	}

	return false
}

// SetTv gets a reference to the given GetUserQuota2XXResponseMovie and assigns it to the Tv field.
func (o *GetUserQuota2XXResponse) SetTv(v GetUserQuota2XXResponseMovie) {
	o.Tv = &v
}

func (o GetUserQuota2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserQuota2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Movie) {
		toSerialize["movie"] = o.Movie
	}
	if !IsNil(o.Tv) {
		toSerialize["tv"] = o.Tv
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserQuota2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUserQuota2XXResponse := _GetUserQuota2XXResponse{}

	err = json.Unmarshal(data, &varGetUserQuota2XXResponse)

	if err != nil {
		return err
	}

	*o = GetUserQuota2XXResponse(varGetUserQuota2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "movie")
		delete(additionalProperties, "tv")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserQuota2XXResponse struct {
	value *GetUserQuota2XXResponse
	isSet bool
}

func (v NullableGetUserQuota2XXResponse) Get() *GetUserQuota2XXResponse {
	return v.value
}

func (v *NullableGetUserQuota2XXResponse) Set(val *GetUserQuota2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserQuota2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserQuota2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserQuota2XXResponse(val *GetUserQuota2XXResponse) *NullableGetUserQuota2XXResponse {
	return &NullableGetUserQuota2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserQuota2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserQuota2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


