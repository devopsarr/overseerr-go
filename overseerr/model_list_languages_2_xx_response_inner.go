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

// ListLanguages2XXResponseInner struct for ListLanguages2XXResponseInner
type ListLanguages2XXResponseInner struct {
	Iso6391 *string `json:"iso_639_1,omitempty"`
	EnglishName *string `json:"english_name,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListLanguages2XXResponseInner ListLanguages2XXResponseInner

// NewListLanguages2XXResponseInner instantiates a new ListLanguages2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLanguages2XXResponseInner() *ListLanguages2XXResponseInner {
	this := ListLanguages2XXResponseInner{}
	return &this
}

// NewListLanguages2XXResponseInnerWithDefaults instantiates a new ListLanguages2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLanguages2XXResponseInnerWithDefaults() *ListLanguages2XXResponseInner {
	this := ListLanguages2XXResponseInner{}
	return &this
}

// GetIso6391 returns the Iso6391 field value if set, zero value otherwise.
func (o *ListLanguages2XXResponseInner) GetIso6391() string {
	if o == nil || isNil(o.Iso6391) {
		var ret string
		return ret
	}
	return *o.Iso6391
}

// GetIso6391Ok returns a tuple with the Iso6391 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLanguages2XXResponseInner) GetIso6391Ok() (*string, bool) {
	if o == nil || isNil(o.Iso6391) {
    return nil, false
	}
	return o.Iso6391, true
}

// HasIso6391 returns a boolean if a field has been set.
func (o *ListLanguages2XXResponseInner) HasIso6391() bool {
	if o != nil && !isNil(o.Iso6391) {
		return true
	}

	return false
}

// SetIso6391 gets a reference to the given string and assigns it to the Iso6391 field.
func (o *ListLanguages2XXResponseInner) SetIso6391(v string) {
	o.Iso6391 = &v
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise.
func (o *ListLanguages2XXResponseInner) GetEnglishName() string {
	if o == nil || isNil(o.EnglishName) {
		var ret string
		return ret
	}
	return *o.EnglishName
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLanguages2XXResponseInner) GetEnglishNameOk() (*string, bool) {
	if o == nil || isNil(o.EnglishName) {
    return nil, false
	}
	return o.EnglishName, true
}

// HasEnglishName returns a boolean if a field has been set.
func (o *ListLanguages2XXResponseInner) HasEnglishName() bool {
	if o != nil && !isNil(o.EnglishName) {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given string and assigns it to the EnglishName field.
func (o *ListLanguages2XXResponseInner) SetEnglishName(v string) {
	o.EnglishName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListLanguages2XXResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLanguages2XXResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListLanguages2XXResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListLanguages2XXResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListLanguages2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iso6391) {
		toSerialize["iso_639_1"] = o.Iso6391
	}
	if !isNil(o.EnglishName) {
		toSerialize["english_name"] = o.EnglishName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListLanguages2XXResponseInner) UnmarshalJSON(bytes []byte) (err error) {
	varListLanguages2XXResponseInner := _ListLanguages2XXResponseInner{}

	if err = json.Unmarshal(bytes, &varListLanguages2XXResponseInner); err == nil {
		*o = ListLanguages2XXResponseInner(varListLanguages2XXResponseInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_639_1")
		delete(additionalProperties, "english_name")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListLanguages2XXResponseInner struct {
	value *ListLanguages2XXResponseInner
	isSet bool
}

func (v NullableListLanguages2XXResponseInner) Get() *ListLanguages2XXResponseInner {
	return v.value
}

func (v *NullableListLanguages2XXResponseInner) Set(val *ListLanguages2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListLanguages2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListLanguages2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLanguages2XXResponseInner(val *ListLanguages2XXResponseInner) *NullableListLanguages2XXResponseInner {
	return &NullableListLanguages2XXResponseInner{value: val, isSet: true}
}

func (v NullableListLanguages2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLanguages2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

