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

// GetTvRatings200Response struct for GetTvRatings200Response
type GetTvRatings200Response struct {
	Title *string `json:"title,omitempty"`
	Year *float32 `json:"year,omitempty"`
	Url *string `json:"url,omitempty"`
	CriticsScore *float32 `json:"criticsScore,omitempty"`
	CriticsRating *string `json:"criticsRating,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetTvRatings200Response GetTvRatings200Response

// NewGetTvRatings200Response instantiates a new GetTvRatings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTvRatings200Response() *GetTvRatings200Response {
	this := GetTvRatings200Response{}
	return &this
}

// NewGetTvRatings200ResponseWithDefaults instantiates a new GetTvRatings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTvRatings200ResponseWithDefaults() *GetTvRatings200Response {
	this := GetTvRatings200Response{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetTvRatings200Response) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTvRatings200Response) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetTvRatings200Response) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetTvRatings200Response) SetTitle(v string) {
	o.Title = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetTvRatings200Response) GetYear() float32 {
	if o == nil || isNil(o.Year) {
		var ret float32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTvRatings200Response) GetYearOk() (*float32, bool) {
	if o == nil || isNil(o.Year) {
    return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetTvRatings200Response) HasYear() bool {
	if o != nil && !isNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given float32 and assigns it to the Year field.
func (o *GetTvRatings200Response) SetYear(v float32) {
	o.Year = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetTvRatings200Response) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTvRatings200Response) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetTvRatings200Response) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetTvRatings200Response) SetUrl(v string) {
	o.Url = &v
}

// GetCriticsScore returns the CriticsScore field value if set, zero value otherwise.
func (o *GetTvRatings200Response) GetCriticsScore() float32 {
	if o == nil || isNil(o.CriticsScore) {
		var ret float32
		return ret
	}
	return *o.CriticsScore
}

// GetCriticsScoreOk returns a tuple with the CriticsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTvRatings200Response) GetCriticsScoreOk() (*float32, bool) {
	if o == nil || isNil(o.CriticsScore) {
    return nil, false
	}
	return o.CriticsScore, true
}

// HasCriticsScore returns a boolean if a field has been set.
func (o *GetTvRatings200Response) HasCriticsScore() bool {
	if o != nil && !isNil(o.CriticsScore) {
		return true
	}

	return false
}

// SetCriticsScore gets a reference to the given float32 and assigns it to the CriticsScore field.
func (o *GetTvRatings200Response) SetCriticsScore(v float32) {
	o.CriticsScore = &v
}

// GetCriticsRating returns the CriticsRating field value if set, zero value otherwise.
func (o *GetTvRatings200Response) GetCriticsRating() string {
	if o == nil || isNil(o.CriticsRating) {
		var ret string
		return ret
	}
	return *o.CriticsRating
}

// GetCriticsRatingOk returns a tuple with the CriticsRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTvRatings200Response) GetCriticsRatingOk() (*string, bool) {
	if o == nil || isNil(o.CriticsRating) {
    return nil, false
	}
	return o.CriticsRating, true
}

// HasCriticsRating returns a boolean if a field has been set.
func (o *GetTvRatings200Response) HasCriticsRating() bool {
	if o != nil && !isNil(o.CriticsRating) {
		return true
	}

	return false
}

// SetCriticsRating gets a reference to the given string and assigns it to the CriticsRating field.
func (o *GetTvRatings200Response) SetCriticsRating(v string) {
	o.CriticsRating = &v
}

func (o GetTvRatings200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.CriticsScore) {
		toSerialize["criticsScore"] = o.CriticsScore
	}
	if !isNil(o.CriticsRating) {
		toSerialize["criticsRating"] = o.CriticsRating
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetTvRatings200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetTvRatings200Response := _GetTvRatings200Response{}

	if err = json.Unmarshal(bytes, &varGetTvRatings200Response); err == nil {
		*o = GetTvRatings200Response(varGetTvRatings200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "year")
		delete(additionalProperties, "url")
		delete(additionalProperties, "criticsScore")
		delete(additionalProperties, "criticsRating")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetTvRatings200Response struct {
	value *GetTvRatings200Response
	isSet bool
}

func (v NullableGetTvRatings200Response) Get() *GetTvRatings200Response {
	return v.value
}

func (v *NullableGetTvRatings200Response) Set(val *GetTvRatings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTvRatings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTvRatings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTvRatings200Response(val *GetTvRatings200Response) *NullableGetTvRatings200Response {
	return &NullableGetTvRatings200Response{value: val, isSet: true}
}

func (v NullableGetTvRatings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTvRatings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

