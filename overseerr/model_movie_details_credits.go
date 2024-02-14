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

// MovieDetailsCredits struct for MovieDetailsCredits
type MovieDetailsCredits struct {
	Cast []*Cast `json:"cast,omitempty"`
	Crew []*Crew `json:"crew,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovieDetailsCredits MovieDetailsCredits

// NewMovieDetailsCredits instantiates a new MovieDetailsCredits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieDetailsCredits() *MovieDetailsCredits {
	this := MovieDetailsCredits{}
	return &this
}

// NewMovieDetailsCreditsWithDefaults instantiates a new MovieDetailsCredits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieDetailsCreditsWithDefaults() *MovieDetailsCredits {
	this := MovieDetailsCredits{}
	return &this
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *MovieDetailsCredits) GetCast() []*Cast {
	if o == nil || isNil(o.Cast) {
		var ret []*Cast
		return ret
	}
	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCredits) GetCastOk() ([]*Cast, bool) {
	if o == nil || isNil(o.Cast) {
    return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *MovieDetailsCredits) HasCast() bool {
	if o != nil && !isNil(o.Cast) {
		return true
	}

	return false
}

// SetCast gets a reference to the given []Cast and assigns it to the Cast field.
func (o *MovieDetailsCredits) SetCast(v []*Cast) {
	o.Cast = v
}

// GetCrew returns the Crew field value if set, zero value otherwise.
func (o *MovieDetailsCredits) GetCrew() []*Crew {
	if o == nil || isNil(o.Crew) {
		var ret []*Crew
		return ret
	}
	return o.Crew
}

// GetCrewOk returns a tuple with the Crew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCredits) GetCrewOk() ([]*Crew, bool) {
	if o == nil || isNil(o.Crew) {
    return nil, false
	}
	return o.Crew, true
}

// HasCrew returns a boolean if a field has been set.
func (o *MovieDetailsCredits) HasCrew() bool {
	if o != nil && !isNil(o.Crew) {
		return true
	}

	return false
}

// SetCrew gets a reference to the given []Crew and assigns it to the Crew field.
func (o *MovieDetailsCredits) SetCrew(v []*Crew) {
	o.Crew = v
}

func (o MovieDetailsCredits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cast) {
		toSerialize["cast"] = o.Cast
	}
	if !isNil(o.Crew) {
		toSerialize["crew"] = o.Crew
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MovieDetailsCredits) UnmarshalJSON(bytes []byte) (err error) {
	varMovieDetailsCredits := _MovieDetailsCredits{}

	if err = json.Unmarshal(bytes, &varMovieDetailsCredits); err == nil {
		*o = MovieDetailsCredits(varMovieDetailsCredits)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cast")
		delete(additionalProperties, "crew")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovieDetailsCredits struct {
	value *MovieDetailsCredits
	isSet bool
}

func (v NullableMovieDetailsCredits) Get() *MovieDetailsCredits {
	return v.value
}

func (v *NullableMovieDetailsCredits) Set(val *MovieDetailsCredits) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieDetailsCredits) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieDetailsCredits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieDetailsCredits(val *MovieDetailsCredits) *NullableMovieDetailsCredits {
	return &NullableMovieDetailsCredits{value: val, isSet: true}
}

func (v NullableMovieDetailsCredits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieDetailsCredits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


