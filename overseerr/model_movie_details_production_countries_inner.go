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

// checks if the MovieDetailsProductionCountriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieDetailsProductionCountriesInner{}

// MovieDetailsProductionCountriesInner struct for MovieDetailsProductionCountriesInner
type MovieDetailsProductionCountriesInner struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovieDetailsProductionCountriesInner MovieDetailsProductionCountriesInner

// NewMovieDetailsProductionCountriesInner instantiates a new MovieDetailsProductionCountriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieDetailsProductionCountriesInner() *MovieDetailsProductionCountriesInner {
	this := MovieDetailsProductionCountriesInner{}
	return &this
}

// NewMovieDetailsProductionCountriesInnerWithDefaults instantiates a new MovieDetailsProductionCountriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieDetailsProductionCountriesInnerWithDefaults() *MovieDetailsProductionCountriesInner {
	this := MovieDetailsProductionCountriesInner{}
	return &this
}

// GetIso31661 returns the Iso31661 field value if set, zero value otherwise.
func (o *MovieDetailsProductionCountriesInner) GetIso31661() string {
	if o == nil || IsNil(o.Iso31661) {
		var ret string
		return ret
	}
	return *o.Iso31661
}

// GetIso31661Ok returns a tuple with the Iso31661 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsProductionCountriesInner) GetIso31661Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso31661) {
		return nil, false
	}
	return o.Iso31661, true
}

// HasIso31661 returns a boolean if a field has been set.
func (o *MovieDetailsProductionCountriesInner) HasIso31661() bool {
	if o != nil && !IsNil(o.Iso31661) {
		return true
	}

	return false
}

// SetIso31661 gets a reference to the given string and assigns it to the Iso31661 field.
func (o *MovieDetailsProductionCountriesInner) SetIso31661(v string) {
	o.Iso31661 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MovieDetailsProductionCountriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsProductionCountriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MovieDetailsProductionCountriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MovieDetailsProductionCountriesInner) SetName(v string) {
	o.Name = &v
}

func (o MovieDetailsProductionCountriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieDetailsProductionCountriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iso31661) {
		toSerialize["iso_3166_1"] = o.Iso31661
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MovieDetailsProductionCountriesInner) UnmarshalJSON(data []byte) (err error) {
	varMovieDetailsProductionCountriesInner := _MovieDetailsProductionCountriesInner{}

	err = json.Unmarshal(data, &varMovieDetailsProductionCountriesInner)

	if err != nil {
		return err
	}

	*o = MovieDetailsProductionCountriesInner(varMovieDetailsProductionCountriesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_3166_1")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovieDetailsProductionCountriesInner struct {
	value *MovieDetailsProductionCountriesInner
	isSet bool
}

func (v NullableMovieDetailsProductionCountriesInner) Get() *MovieDetailsProductionCountriesInner {
	return v.value
}

func (v *NullableMovieDetailsProductionCountriesInner) Set(val *MovieDetailsProductionCountriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieDetailsProductionCountriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieDetailsProductionCountriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieDetailsProductionCountriesInner(val *MovieDetailsProductionCountriesInner) *NullableMovieDetailsProductionCountriesInner {
	return &NullableMovieDetailsProductionCountriesInner{value: val, isSet: true}
}

func (v NullableMovieDetailsProductionCountriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieDetailsProductionCountriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


