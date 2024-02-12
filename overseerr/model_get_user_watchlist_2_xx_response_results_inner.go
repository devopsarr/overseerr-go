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

// GetUserWatchlist2XXResponseResultsInner struct for GetUserWatchlist2XXResponseResultsInner
type GetUserWatchlist2XXResponseResultsInner struct {
	TmdbId *float32 `json:"tmdbId,omitempty"`
	RatingKey *string `json:"ratingKey,omitempty"`
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserWatchlist2XXResponseResultsInner GetUserWatchlist2XXResponseResultsInner

// NewGetUserWatchlist2XXResponseResultsInner instantiates a new GetUserWatchlist2XXResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserWatchlist2XXResponseResultsInner() *GetUserWatchlist2XXResponseResultsInner {
	this := GetUserWatchlist2XXResponseResultsInner{}
	return &this
}

// NewGetUserWatchlist2XXResponseResultsInnerWithDefaults instantiates a new GetUserWatchlist2XXResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserWatchlist2XXResponseResultsInnerWithDefaults() *GetUserWatchlist2XXResponseResultsInner {
	this := GetUserWatchlist2XXResponseResultsInner{}
	return &this
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponseResultsInner) GetTmdbId() float32 {
	if o == nil || isNil(o.TmdbId) {
		var ret float32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) GetTmdbIdOk() (*float32, bool) {
	if o == nil || isNil(o.TmdbId) {
    return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) HasTmdbId() bool {
	if o != nil && !isNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given float32 and assigns it to the TmdbId field.
func (o *GetUserWatchlist2XXResponseResultsInner) SetTmdbId(v float32) {
	o.TmdbId = &v
}

// GetRatingKey returns the RatingKey field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponseResultsInner) GetRatingKey() string {
	if o == nil || isNil(o.RatingKey) {
		var ret string
		return ret
	}
	return *o.RatingKey
}

// GetRatingKeyOk returns a tuple with the RatingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) GetRatingKeyOk() (*string, bool) {
	if o == nil || isNil(o.RatingKey) {
    return nil, false
	}
	return o.RatingKey, true
}

// HasRatingKey returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) HasRatingKey() bool {
	if o != nil && !isNil(o.RatingKey) {
		return true
	}

	return false
}

// SetRatingKey gets a reference to the given string and assigns it to the RatingKey field.
func (o *GetUserWatchlist2XXResponseResultsInner) SetRatingKey(v string) {
	o.RatingKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponseResultsInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetUserWatchlist2XXResponseResultsInner) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponseResultsInner) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponseResultsInner) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetUserWatchlist2XXResponseResultsInner) SetTitle(v string) {
	o.Title = &v
}

func (o GetUserWatchlist2XXResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TmdbId) {
		toSerialize["tmdbId"] = o.TmdbId
	}
	if !isNil(o.RatingKey) {
		toSerialize["ratingKey"] = o.RatingKey
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetUserWatchlist2XXResponseResultsInner) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserWatchlist2XXResponseResultsInner := _GetUserWatchlist2XXResponseResultsInner{}

	if err = json.Unmarshal(bytes, &varGetUserWatchlist2XXResponseResultsInner); err == nil {
		*o = GetUserWatchlist2XXResponseResultsInner(varGetUserWatchlist2XXResponseResultsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tmdbId")
		delete(additionalProperties, "ratingKey")
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserWatchlist2XXResponseResultsInner struct {
	value *GetUserWatchlist2XXResponseResultsInner
	isSet bool
}

func (v NullableGetUserWatchlist2XXResponseResultsInner) Get() *GetUserWatchlist2XXResponseResultsInner {
	return v.value
}

func (v *NullableGetUserWatchlist2XXResponseResultsInner) Set(val *GetUserWatchlist2XXResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserWatchlist2XXResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserWatchlist2XXResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserWatchlist2XXResponseResultsInner(val *GetUserWatchlist2XXResponseResultsInner) *NullableGetUserWatchlist2XXResponseResultsInner {
	return &NullableGetUserWatchlist2XXResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetUserWatchlist2XXResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserWatchlist2XXResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

