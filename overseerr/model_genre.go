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

// Genre struct for Genre
type Genre struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Genre Genre

// NewGenre instantiates a new Genre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenre() *Genre {
	this := Genre{}
	return &this
}

// NewGenreWithDefaults instantiates a new Genre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenreWithDefaults() *Genre {
	this := Genre{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Genre) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Genre) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Genre) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Genre) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Genre) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Genre) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Genre) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Genre) SetName(v string) {
	o.Name = &v
}

func (o Genre) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Genre) UnmarshalJSON(bytes []byte) (err error) {
	varGenre := _Genre{}

	if err = json.Unmarshal(bytes, &varGenre); err == nil {
		*o = Genre(varGenre)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenre struct {
	value *Genre
	isSet bool
}

func (v NullableGenre) Get() *Genre {
	return v.value
}

func (v *NullableGenre) Set(val *Genre) {
	v.value = val
	v.isSet = true
}

func (v NullableGenre) IsSet() bool {
	return v.isSet
}

func (v *NullableGenre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenre(val *Genre) *NullableGenre {
	return &NullableGenre{value: val, isSet: true}
}

func (v NullableGenre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


