/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the ListRegions2XXResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRegions2XXResponseInner{}

// ListRegions2XXResponseInner struct for ListRegions2XXResponseInner
type ListRegions2XXResponseInner struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	EnglishName *string `json:"english_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRegions2XXResponseInner ListRegions2XXResponseInner

// NewListRegions2XXResponseInner instantiates a new ListRegions2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRegions2XXResponseInner() *ListRegions2XXResponseInner {
	this := ListRegions2XXResponseInner{}
	return &this
}

// NewListRegions2XXResponseInnerWithDefaults instantiates a new ListRegions2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRegions2XXResponseInnerWithDefaults() *ListRegions2XXResponseInner {
	this := ListRegions2XXResponseInner{}
	return &this
}

// GetIso31661 returns the Iso31661 field value if set, zero value otherwise.
func (o *ListRegions2XXResponseInner) GetIso31661() string {
	if o == nil || IsNil(o.Iso31661) {
		var ret string
		return ret
	}
	return *o.Iso31661
}

// GetIso31661Ok returns a tuple with the Iso31661 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegions2XXResponseInner) GetIso31661Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso31661) {
		return nil, false
	}
	return o.Iso31661, true
}

// HasIso31661 returns a boolean if a field has been set.
func (o *ListRegions2XXResponseInner) HasIso31661() bool {
	if o != nil && !IsNil(o.Iso31661) {
		return true
	}

	return false
}

// SetIso31661 gets a reference to the given string and assigns it to the Iso31661 field.
func (o *ListRegions2XXResponseInner) SetIso31661(v string) {
	o.Iso31661 = &v
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise.
func (o *ListRegions2XXResponseInner) GetEnglishName() string {
	if o == nil || IsNil(o.EnglishName) {
		var ret string
		return ret
	}
	return *o.EnglishName
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRegions2XXResponseInner) GetEnglishNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnglishName) {
		return nil, false
	}
	return o.EnglishName, true
}

// HasEnglishName returns a boolean if a field has been set.
func (o *ListRegions2XXResponseInner) HasEnglishName() bool {
	if o != nil && !IsNil(o.EnglishName) {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given string and assigns it to the EnglishName field.
func (o *ListRegions2XXResponseInner) SetEnglishName(v string) {
	o.EnglishName = &v
}

func (o ListRegions2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRegions2XXResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iso31661) {
		toSerialize["iso_3166_1"] = o.Iso31661
	}
	if !IsNil(o.EnglishName) {
		toSerialize["english_name"] = o.EnglishName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRegions2XXResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListRegions2XXResponseInner := _ListRegions2XXResponseInner{}

	err = json.Unmarshal(data, &varListRegions2XXResponseInner)

	if err != nil {
		return err
	}

	*o = ListRegions2XXResponseInner(varListRegions2XXResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_3166_1")
		delete(additionalProperties, "english_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListRegions2XXResponseInner struct {
	value *ListRegions2XXResponseInner
	isSet bool
}

func (v NullableListRegions2XXResponseInner) Get() *ListRegions2XXResponseInner {
	return v.value
}

func (v *NullableListRegions2XXResponseInner) Set(val *ListRegions2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListRegions2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListRegions2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRegions2XXResponseInner(val *ListRegions2XXResponseInner) *NullableListRegions2XXResponseInner {
	return &NullableListRegions2XXResponseInner{value: val, isSet: true}
}

func (v NullableListRegions2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRegions2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


