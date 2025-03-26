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

// checks if the SonarrSeriesSeasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarrSeriesSeasonsInner{}

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
	if o == nil || IsNil(o.SeasonNumber) {
		var ret float32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesSeasonsInner) GetSeasonNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.SeasonNumber) {
		return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *SonarrSeriesSeasonsInner) HasSeasonNumber() bool {
	if o != nil && !IsNil(o.SeasonNumber) {
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
	if o == nil || IsNil(o.Monitored) {
		var ret bool
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSeriesSeasonsInner) GetMonitoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *SonarrSeriesSeasonsInner) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given bool and assigns it to the Monitored field.
func (o *SonarrSeriesSeasonsInner) SetMonitored(v bool) {
	o.Monitored = &v
}

func (o SonarrSeriesSeasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarrSeriesSeasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SonarrSeriesSeasonsInner) UnmarshalJSON(data []byte) (err error) {
	varSonarrSeriesSeasonsInner := _SonarrSeriesSeasonsInner{}

	err = json.Unmarshal(data, &varSonarrSeriesSeasonsInner)

	if err != nil {
		return err
	}

	*o = SonarrSeriesSeasonsInner(varSonarrSeriesSeasonsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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


