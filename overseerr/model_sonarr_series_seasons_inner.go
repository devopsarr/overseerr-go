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

// SonarrSeriesSeasonsInner struct for SonarrSeriesSeasonsInner
type SonarrSeriesSeasonsInner struct {
	SeasonNumber *float32 `json:"seasonNumber,omitempty"`
	Monitored *bool `json:"monitored,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SonarrSeriesSeasonsInner SonarrSeriesSeasonsInner

// NewSonarrSeriesSeasonsInner instantiates a new SonarrSeriesSeasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarrSeriesSeasonsInner() *SonarrSeriesSeasonsInner {
	this := SonarrSeriesSeasonsInner{}
	return &this
}

// NewSonarrSeriesSeasonsInnerWithDefaults instantiates a new SonarrSeriesSeasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarrSeriesSeasonsInnerWithDefaults() *SonarrSeriesSeasonsInner {
	this := SonarrSeriesSeasonsInner{}
	return &this
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *SonarrSeriesSeasonsInner) GetSeasonNumber() float32 {
	if o == nil || isNil(o.SeasonNumber) {
		var ret float32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesSeasonsInner) GetSeasonNumberOk() (*float32, bool) {
	if o == nil || isNil(o.SeasonNumber) {
    return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *SonarrSeriesSeasonsInner) HasSeasonNumber() bool {
	if o != nil && !isNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given float32 and assigns it to the SeasonNumber field.
func (o *SonarrSeriesSeasonsInner) SetSeasonNumber(v float32) {
	o.SeasonNumber = &v
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *SonarrSeriesSeasonsInner) GetMonitored() bool {
	if o == nil || isNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesSeasonsInner) GetMonitoredOk() (*bool, bool) {
	if o == nil || isNil(o.Monitored) {
    return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *SonarrSeriesSeasonsInner) HasMonitored() bool {
	if o != nil && !isNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *SonarrSeriesSeasonsInner) SetMonitored(v bool) {
	o.Monitored = &v
}

func (o SonarrSeriesSeasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !isNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SonarrSeriesSeasonsInner) UnmarshalJSON(bytes []byte) (err error) {
	varSonarrSeriesSeasonsInner := _SonarrSeriesSeasonsInner{}

	if err = json.Unmarshal(bytes, &varSonarrSeriesSeasonsInner); err == nil {
		*o = SonarrSeriesSeasonsInner(varSonarrSeriesSeasonsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "seasonNumber")
		delete(additionalProperties, "monitored")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSonarrSeriesSeasonsInner struct {
	value *SonarrSeriesSeasonsInner
	isSet bool
}

func (v NullableSonarrSeriesSeasonsInner) Get() *SonarrSeriesSeasonsInner {
	return v.value
}

func (v *NullableSonarrSeriesSeasonsInner) Set(val *SonarrSeriesSeasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarrSeriesSeasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarrSeriesSeasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarrSeriesSeasonsInner(val *SonarrSeriesSeasonsInner) *NullableSonarrSeriesSeasonsInner {
	return &NullableSonarrSeriesSeasonsInner{value: val, isSet: true}
}

func (v NullableSonarrSeriesSeasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarrSeriesSeasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


