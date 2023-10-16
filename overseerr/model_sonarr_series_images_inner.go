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

// SonarrSeriesImagesInner struct for SonarrSeriesImagesInner
type SonarrSeriesImagesInner struct {
	CoverType *string `json:"coverType,omitempty"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SonarrSeriesImagesInner SonarrSeriesImagesInner

// NewSonarrSeriesImagesInner instantiates a new SonarrSeriesImagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarrSeriesImagesInner() *SonarrSeriesImagesInner {
	this := SonarrSeriesImagesInner{}
	return &this
}

// NewSonarrSeriesImagesInnerWithDefaults instantiates a new SonarrSeriesImagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarrSeriesImagesInnerWithDefaults() *SonarrSeriesImagesInner {
	this := SonarrSeriesImagesInner{}
	return &this
}

// GetCoverType returns the CoverType field value if set, zero value otherwise.
func (o *SonarrSeriesImagesInner) GetCoverType() string {
	if o == nil || isNil(o.CoverType) {
		var ret string
		return ret
	}
	return *o.CoverType
}

// GetCoverTypeOk returns a tuple with the CoverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesImagesInner) GetCoverTypeOk() (*string, bool) {
	if o == nil || isNil(o.CoverType) {
    return nil, false
	}
	return o.CoverType, true
}

// HasCoverType returns a boolean if a field has been set.
func (o *SonarrSeriesImagesInner) HasCoverType() bool {
	if o != nil && !isNil(o.CoverType) {
		return true
	}

	return false
}

// SetCoverType gets a reference to the given string and assigns it to the CoverType field.
func (o *SonarrSeriesImagesInner) SetCoverType(v string) {
	o.CoverType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SonarrSeriesImagesInner) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesImagesInner) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SonarrSeriesImagesInner) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SonarrSeriesImagesInner) SetUrl(v string) {
	o.Url = &v
}

func (o SonarrSeriesImagesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CoverType) {
		toSerialize["coverType"] = o.CoverType
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SonarrSeriesImagesInner) UnmarshalJSON(bytes []byte) (err error) {
	varSonarrSeriesImagesInner := _SonarrSeriesImagesInner{}

	if err = json.Unmarshal(bytes, &varSonarrSeriesImagesInner); err == nil {
		*o = SonarrSeriesImagesInner(varSonarrSeriesImagesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "coverType")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSonarrSeriesImagesInner struct {
	value *SonarrSeriesImagesInner
	isSet bool
}

func (v NullableSonarrSeriesImagesInner) Get() *SonarrSeriesImagesInner {
	return v.value
}

func (v *NullableSonarrSeriesImagesInner) Set(val *SonarrSeriesImagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarrSeriesImagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarrSeriesImagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarrSeriesImagesInner(val *SonarrSeriesImagesInner) *NullableSonarrSeriesImagesInner {
	return &NullableSonarrSeriesImagesInner{value: val, isSet: true}
}

func (v NullableSonarrSeriesImagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarrSeriesImagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

