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

// GetUserWatchData2XXResponse struct for GetUserWatchData2XXResponse
type GetUserWatchData2XXResponse struct {
	RecentlyWatched []*MediaInfo `json:"recentlyWatched,omitempty"`
	PlayCount *float32 `json:"playCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserWatchData2XXResponse GetUserWatchData2XXResponse

// NewGetUserWatchData2XXResponse instantiates a new GetUserWatchData2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserWatchData2XXResponse() *GetUserWatchData2XXResponse {
	this := GetUserWatchData2XXResponse{}
	return &this
}

// NewGetUserWatchData2XXResponseWithDefaults instantiates a new GetUserWatchData2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserWatchData2XXResponseWithDefaults() *GetUserWatchData2XXResponse {
	this := GetUserWatchData2XXResponse{}
	return &this
}

// GetRecentlyWatched returns the RecentlyWatched field value if set, zero value otherwise.
func (o *GetUserWatchData2XXResponse) GetRecentlyWatched() []*MediaInfo {
	if o == nil || isNil(o.RecentlyWatched) {
		var ret []*MediaInfo
		return ret
	}
	return o.RecentlyWatched
}

// GetRecentlyWatchedOk returns a tuple with the RecentlyWatched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchData2XXResponse) GetRecentlyWatchedOk() ([]*MediaInfo, bool) {
	if o == nil || isNil(o.RecentlyWatched) {
    return nil, false
	}
	return o.RecentlyWatched, true
}

// HasRecentlyWatched returns a boolean if a field has been set.
func (o *GetUserWatchData2XXResponse) HasRecentlyWatched() bool {
	if o != nil && !isNil(o.RecentlyWatched) {
		return true
	}

	return false
}

// SetRecentlyWatched gets a reference to the given []MediaInfo and assigns it to the RecentlyWatched field.
func (o *GetUserWatchData2XXResponse) SetRecentlyWatched(v []*MediaInfo) {
	o.RecentlyWatched = v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *GetUserWatchData2XXResponse) GetPlayCount() float32 {
	if o == nil || isNil(o.PlayCount) {
		var ret float32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchData2XXResponse) GetPlayCountOk() (*float32, bool) {
	if o == nil || isNil(o.PlayCount) {
    return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *GetUserWatchData2XXResponse) HasPlayCount() bool {
	if o != nil && !isNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given float32 and assigns it to the PlayCount field.
func (o *GetUserWatchData2XXResponse) SetPlayCount(v float32) {
	o.PlayCount = &v
}

func (o GetUserWatchData2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RecentlyWatched) {
		toSerialize["recentlyWatched"] = o.RecentlyWatched
	}
	if !isNil(o.PlayCount) {
		toSerialize["playCount"] = o.PlayCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetUserWatchData2XXResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserWatchData2XXResponse := _GetUserWatchData2XXResponse{}

	if err = json.Unmarshal(bytes, &varGetUserWatchData2XXResponse); err == nil {
		*o = GetUserWatchData2XXResponse(varGetUserWatchData2XXResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recentlyWatched")
		delete(additionalProperties, "playCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserWatchData2XXResponse struct {
	value *GetUserWatchData2XXResponse
	isSet bool
}

func (v NullableGetUserWatchData2XXResponse) Get() *GetUserWatchData2XXResponse {
	return v.value
}

func (v *NullableGetUserWatchData2XXResponse) Set(val *GetUserWatchData2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserWatchData2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserWatchData2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserWatchData2XXResponse(val *GetUserWatchData2XXResponse) *NullableGetUserWatchData2XXResponse {
	return &NullableGetUserWatchData2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserWatchData2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserWatchData2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

