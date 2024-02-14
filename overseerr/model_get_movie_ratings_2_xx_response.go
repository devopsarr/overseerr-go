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

// GetMovieRatings2XXResponse struct for GetMovieRatings2XXResponse
type GetMovieRatings2XXResponse struct {
	Title *string `json:"title,omitempty"`
	Year *float32 `json:"year,omitempty"`
	Url *string `json:"url,omitempty"`
	CriticsScore *float32 `json:"criticsScore,omitempty"`
	CriticsRating *string `json:"criticsRating,omitempty"`
	AudienceScore *float32 `json:"audienceScore,omitempty"`
	AudienceRating *string `json:"audienceRating,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMovieRatings2XXResponse GetMovieRatings2XXResponse

// NewGetMovieRatings2XXResponse instantiates a new GetMovieRatings2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMovieRatings2XXResponse() *GetMovieRatings2XXResponse {
	this := GetMovieRatings2XXResponse{}
	return &this
}

// NewGetMovieRatings2XXResponseWithDefaults instantiates a new GetMovieRatings2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMovieRatings2XXResponseWithDefaults() *GetMovieRatings2XXResponse {
	this := GetMovieRatings2XXResponse{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetMovieRatings2XXResponse) SetTitle(v string) {
	o.Title = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetYear() float32 {
	if o == nil || isNil(o.Year) {
		var ret float32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetYearOk() (*float32, bool) {
	if o == nil || isNil(o.Year) {
    return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasYear() bool {
	if o != nil && !isNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given float32 and assigns it to the Year field.
func (o *GetMovieRatings2XXResponse) SetYear(v float32) {
	o.Year = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetMovieRatings2XXResponse) SetUrl(v string) {
	o.Url = &v
}

// GetCriticsScore returns the CriticsScore field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetCriticsScore() float32 {
	if o == nil || isNil(o.CriticsScore) {
		var ret float32
		return ret
	}
	return *o.CriticsScore
}

// GetCriticsScoreOk returns a tuple with the CriticsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetCriticsScoreOk() (*float32, bool) {
	if o == nil || isNil(o.CriticsScore) {
    return nil, false
	}
	return o.CriticsScore, true
}

// HasCriticsScore returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasCriticsScore() bool {
	if o != nil && !isNil(o.CriticsScore) {
		return true
	}

	return false
}

// SetCriticsScore gets a reference to the given float32 and assigns it to the CriticsScore field.
func (o *GetMovieRatings2XXResponse) SetCriticsScore(v float32) {
	o.CriticsScore = &v
}

// GetCriticsRating returns the CriticsRating field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetCriticsRating() string {
	if o == nil || isNil(o.CriticsRating) {
		var ret string
		return ret
	}
	return *o.CriticsRating
}

// GetCriticsRatingOk returns a tuple with the CriticsRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetCriticsRatingOk() (*string, bool) {
	if o == nil || isNil(o.CriticsRating) {
    return nil, false
	}
	return o.CriticsRating, true
}

// HasCriticsRating returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasCriticsRating() bool {
	if o != nil && !isNil(o.CriticsRating) {
		return true
	}

	return false
}

// SetCriticsRating gets a reference to the given string and assigns it to the CriticsRating field.
func (o *GetMovieRatings2XXResponse) SetCriticsRating(v string) {
	o.CriticsRating = &v
}

// GetAudienceScore returns the AudienceScore field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetAudienceScore() float32 {
	if o == nil || isNil(o.AudienceScore) {
		var ret float32
		return ret
	}
	return *o.AudienceScore
}

// GetAudienceScoreOk returns a tuple with the AudienceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetAudienceScoreOk() (*float32, bool) {
	if o == nil || isNil(o.AudienceScore) {
    return nil, false
	}
	return o.AudienceScore, true
}

// HasAudienceScore returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasAudienceScore() bool {
	if o != nil && !isNil(o.AudienceScore) {
		return true
	}

	return false
}

// SetAudienceScore gets a reference to the given float32 and assigns it to the AudienceScore field.
func (o *GetMovieRatings2XXResponse) SetAudienceScore(v float32) {
	o.AudienceScore = &v
}

// GetAudienceRating returns the AudienceRating field value if set, zero value otherwise.
func (o *GetMovieRatings2XXResponse) GetAudienceRating() string {
	if o == nil || isNil(o.AudienceRating) {
		var ret string
		return ret
	}
	return *o.AudienceRating
}

// GetAudienceRatingOk returns a tuple with the AudienceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovieRatings2XXResponse) GetAudienceRatingOk() (*string, bool) {
	if o == nil || isNil(o.AudienceRating) {
    return nil, false
	}
	return o.AudienceRating, true
}

// HasAudienceRating returns a boolean if a field has been set.
func (o *GetMovieRatings2XXResponse) HasAudienceRating() bool {
	if o != nil && !isNil(o.AudienceRating) {
		return true
	}

	return false
}

// SetAudienceRating gets a reference to the given string and assigns it to the AudienceRating field.
func (o *GetMovieRatings2XXResponse) SetAudienceRating(v string) {
	o.AudienceRating = &v
}

func (o GetMovieRatings2XXResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AudienceScore) {
		toSerialize["audienceScore"] = o.AudienceScore
	}
	if !isNil(o.AudienceRating) {
		toSerialize["audienceRating"] = o.AudienceRating
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetMovieRatings2XXResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetMovieRatings2XXResponse := _GetMovieRatings2XXResponse{}

	if err = json.Unmarshal(bytes, &varGetMovieRatings2XXResponse); err == nil {
		*o = GetMovieRatings2XXResponse(varGetMovieRatings2XXResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "year")
		delete(additionalProperties, "url")
		delete(additionalProperties, "criticsScore")
		delete(additionalProperties, "criticsRating")
		delete(additionalProperties, "audienceScore")
		delete(additionalProperties, "audienceRating")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMovieRatings2XXResponse struct {
	value *GetMovieRatings2XXResponse
	isSet bool
}

func (v NullableGetMovieRatings2XXResponse) Get() *GetMovieRatings2XXResponse {
	return v.value
}

func (v *NullableGetMovieRatings2XXResponse) Set(val *GetMovieRatings2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMovieRatings2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMovieRatings2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMovieRatings2XXResponse(val *GetMovieRatings2XXResponse) *NullableGetMovieRatings2XXResponse {
	return &NullableGetMovieRatings2XXResponse{value: val, isSet: true}
}

func (v NullableGetMovieRatings2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMovieRatings2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


