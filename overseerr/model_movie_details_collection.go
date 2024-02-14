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

// MovieDetailsCollection struct for MovieDetailsCollection
type MovieDetailsCollection struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	BackdropPath *string `json:"backdropPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovieDetailsCollection MovieDetailsCollection

// NewMovieDetailsCollection instantiates a new MovieDetailsCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieDetailsCollection() *MovieDetailsCollection {
	this := MovieDetailsCollection{}
	return &this
}

// NewMovieDetailsCollectionWithDefaults instantiates a new MovieDetailsCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieDetailsCollectionWithDefaults() *MovieDetailsCollection {
	this := MovieDetailsCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MovieDetailsCollection) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCollection) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MovieDetailsCollection) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *MovieDetailsCollection) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MovieDetailsCollection) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCollection) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MovieDetailsCollection) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MovieDetailsCollection) SetName(v string) {
	o.Name = &v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *MovieDetailsCollection) GetPosterPath() string {
	if o == nil || isNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCollection) GetPosterPathOk() (*string, bool) {
	if o == nil || isNil(o.PosterPath) {
    return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *MovieDetailsCollection) HasPosterPath() bool {
	if o != nil && !isNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *MovieDetailsCollection) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetBackdropPath returns the BackdropPath field value if set, zero value otherwise.
func (o *MovieDetailsCollection) GetBackdropPath() string {
	if o == nil || isNil(o.BackdropPath) {
		var ret string
		return ret
	}
	return *o.BackdropPath
}

// GetBackdropPathOk returns a tuple with the BackdropPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsCollection) GetBackdropPathOk() (*string, bool) {
	if o == nil || isNil(o.BackdropPath) {
    return nil, false
	}
	return o.BackdropPath, true
}

// HasBackdropPath returns a boolean if a field has been set.
func (o *MovieDetailsCollection) HasBackdropPath() bool {
	if o != nil && !isNil(o.BackdropPath) {
		return true
	}

	return false
}

// SetBackdropPath gets a reference to the given string and assigns it to the BackdropPath field.
func (o *MovieDetailsCollection) SetBackdropPath(v string) {
	o.BackdropPath = &v
}

func (o MovieDetailsCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !isNil(o.BackdropPath) {
		toSerialize["backdropPath"] = o.BackdropPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MovieDetailsCollection) UnmarshalJSON(bytes []byte) (err error) {
	varMovieDetailsCollection := _MovieDetailsCollection{}

	if err = json.Unmarshal(bytes, &varMovieDetailsCollection); err == nil {
		*o = MovieDetailsCollection(varMovieDetailsCollection)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "posterPath")
		delete(additionalProperties, "backdropPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovieDetailsCollection struct {
	value *MovieDetailsCollection
	isSet bool
}

func (v NullableMovieDetailsCollection) Get() *MovieDetailsCollection {
	return v.value
}

func (v *NullableMovieDetailsCollection) Set(val *MovieDetailsCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieDetailsCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieDetailsCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieDetailsCollection(val *MovieDetailsCollection) *NullableMovieDetailsCollection {
	return &NullableMovieDetailsCollection{value: val, isSet: true}
}

func (v NullableMovieDetailsCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieDetailsCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


