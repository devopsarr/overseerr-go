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

// GetUserWatchlist200ResponseResultsInner struct for GetUserWatchlist200ResponseResultsInner
type GetUserWatchlist200ResponseResultsInner struct {
	TmdbId *float32 `json:"tmdbId,omitempty"`
	RatingKey *string `json:"ratingKey,omitempty"`
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserWatchlist200ResponseResultsInner GetUserWatchlist200ResponseResultsInner

// NewGetUserWatchlist200ResponseResultsInner instantiates a new GetUserWatchlist200ResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserWatchlist200ResponseResultsInner() *GetUserWatchlist200ResponseResultsInner {
	this := GetUserWatchlist200ResponseResultsInner{}
	return &this
}

// NewGetUserWatchlist200ResponseResultsInnerWithDefaults instantiates a new GetUserWatchlist200ResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserWatchlist200ResponseResultsInnerWithDefaults() *GetUserWatchlist200ResponseResultsInner {
	this := GetUserWatchlist200ResponseResultsInner{}
	return &this
}

// GetTmdbId returns the TmdbId field value if set, zero value otherwise.
func (o *GetUserWatchlist200ResponseResultsInner) GetTmdbId() float32 {
	if o == nil || isNil(o.TmdbId) {
		var ret float32
		return ret
	}
	return *o.TmdbId
}

// GetTmdbIdOk returns a tuple with the TmdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist200ResponseResultsInner) GetTmdbIdOk() (*float32, bool) {
	if o == nil || isNil(o.TmdbId) {
    return nil, false
	}
	return o.TmdbId, true
}

// HasTmdbId returns a boolean if a field has been set.
func (o *GetUserWatchlist200ResponseResultsInner) HasTmdbId() bool {
	if o != nil && !isNil(o.TmdbId) {
		return true
	}

	return false
}

// SetTmdbId gets a reference to the given float32 and assigns it to the TmdbId field.
func (o *GetUserWatchlist200ResponseResultsInner) SetTmdbId(v float32) {
	o.TmdbId = &v
}

// GetRatingKey returns the RatingKey field value if set, zero value otherwise.
func (o *GetUserWatchlist200ResponseResultsInner) GetRatingKey() string {
	if o == nil || isNil(o.RatingKey) {
		var ret string
		return ret
	}
	return *o.RatingKey
}

// GetRatingKeyOk returns a tuple with the RatingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist200ResponseResultsInner) GetRatingKeyOk() (*string, bool) {
	if o == nil || isNil(o.RatingKey) {
    return nil, false
	}
	return o.RatingKey, true
}

// HasRatingKey returns a boolean if a field has been set.
func (o *GetUserWatchlist200ResponseResultsInner) HasRatingKey() bool {
	if o != nil && !isNil(o.RatingKey) {
		return true
	}

	return false
}

// SetRatingKey gets a reference to the given string and assigns it to the RatingKey field.
func (o *GetUserWatchlist200ResponseResultsInner) SetRatingKey(v string) {
	o.RatingKey = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetUserWatchlist200ResponseResultsInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist200ResponseResultsInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetUserWatchlist200ResponseResultsInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetUserWatchlist200ResponseResultsInner) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetUserWatchlist200ResponseResultsInner) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist200ResponseResultsInner) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetUserWatchlist200ResponseResultsInner) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetUserWatchlist200ResponseResultsInner) SetTitle(v string) {
	o.Title = &v
}

func (o GetUserWatchlist200ResponseResultsInner) MarshalJSON() ([]byte, error) {
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

func (o *GetUserWatchlist200ResponseResultsInner) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserWatchlist200ResponseResultsInner := _GetUserWatchlist200ResponseResultsInner{}

	if err = json.Unmarshal(bytes, &varGetUserWatchlist200ResponseResultsInner); err == nil {
		*o = GetUserWatchlist200ResponseResultsInner(varGetUserWatchlist200ResponseResultsInner)
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

type NullableGetUserWatchlist200ResponseResultsInner struct {
	value *GetUserWatchlist200ResponseResultsInner
	isSet bool
}

func (v NullableGetUserWatchlist200ResponseResultsInner) Get() *GetUserWatchlist200ResponseResultsInner {
	return v.value
}

func (v *NullableGetUserWatchlist200ResponseResultsInner) Set(val *GetUserWatchlist200ResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserWatchlist200ResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserWatchlist200ResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserWatchlist200ResponseResultsInner(val *GetUserWatchlist200ResponseResultsInner) *NullableGetUserWatchlist200ResponseResultsInner {
	return &NullableGetUserWatchlist200ResponseResultsInner{value: val, isSet: true}
}

func (v NullableGetUserWatchlist200ResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserWatchlist200ResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


