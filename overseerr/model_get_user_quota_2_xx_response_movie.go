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

// GetUserQuota2XXResponseMovie struct for GetUserQuota2XXResponseMovie
type GetUserQuota2XXResponseMovie struct {
	Days *float32 `json:"days,omitempty"`
	Limit *float32 `json:"limit,omitempty"`
	Used *float32 `json:"used,omitempty"`
	Remaining *float32 `json:"remaining,omitempty"`
	Restricted *bool `json:"restricted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserQuota2XXResponseMovie GetUserQuota2XXResponseMovie

// NewGetUserQuota2XXResponseMovie instantiates a new GetUserQuota2XXResponseMovie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserQuota2XXResponseMovie() *GetUserQuota2XXResponseMovie {
	this := GetUserQuota2XXResponseMovie{}
	return &this
}

// NewGetUserQuota2XXResponseMovieWithDefaults instantiates a new GetUserQuota2XXResponseMovie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserQuota2XXResponseMovieWithDefaults() *GetUserQuota2XXResponseMovie {
	this := GetUserQuota2XXResponseMovie{}
	return &this
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponseMovie) GetDays() float32 {
	if o == nil || isNil(o.Days) {
		var ret float32
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponseMovie) GetDaysOk() (*float32, bool) {
	if o == nil || isNil(o.Days) {
    return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponseMovie) HasDays() bool {
	if o != nil && !isNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given float32 and assigns it to the Days field.
func (o *GetUserQuota2XXResponseMovie) SetDays(v float32) {
	o.Days = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponseMovie) GetLimit() float32 {
	if o == nil || isNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponseMovie) GetLimitOk() (*float32, bool) {
	if o == nil || isNil(o.Limit) {
    return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponseMovie) HasLimit() bool {
	if o != nil && !isNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *GetUserQuota2XXResponseMovie) SetLimit(v float32) {
	o.Limit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponseMovie) GetUsed() float32 {
	if o == nil || isNil(o.Used) {
		var ret float32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponseMovie) GetUsedOk() (*float32, bool) {
	if o == nil || isNil(o.Used) {
    return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponseMovie) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given float32 and assigns it to the Used field.
func (o *GetUserQuota2XXResponseMovie) SetUsed(v float32) {
	o.Used = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponseMovie) GetRemaining() float32 {
	if o == nil || isNil(o.Remaining) {
		var ret float32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponseMovie) GetRemainingOk() (*float32, bool) {
	if o == nil || isNil(o.Remaining) {
    return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponseMovie) HasRemaining() bool {
	if o != nil && !isNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given float32 and assigns it to the Remaining field.
func (o *GetUserQuota2XXResponseMovie) SetRemaining(v float32) {
	o.Remaining = &v
}

// GetRestricted returns the Restricted field value if set, zero value otherwise.
func (o *GetUserQuota2XXResponseMovie) GetRestricted() bool {
	if o == nil || isNil(o.Restricted) {
		var ret bool
		return ret
	}
	return *o.Restricted
}

// GetRestrictedOk returns a tuple with the Restricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserQuota2XXResponseMovie) GetRestrictedOk() (*bool, bool) {
	if o == nil || isNil(o.Restricted) {
    return nil, false
	}
	return o.Restricted, true
}

// HasRestricted returns a boolean if a field has been set.
func (o *GetUserQuota2XXResponseMovie) HasRestricted() bool {
	if o != nil && !isNil(o.Restricted) {
		return true
	}

	return false
}

// SetRestricted gets a reference to the given bool and assigns it to the Restricted field.
func (o *GetUserQuota2XXResponseMovie) SetRestricted(v bool) {
	o.Restricted = &v
}

func (o GetUserQuota2XXResponseMovie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !isNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	if !isNil(o.Restricted) {
		toSerialize["restricted"] = o.Restricted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetUserQuota2XXResponseMovie) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserQuota2XXResponseMovie := _GetUserQuota2XXResponseMovie{}

	if err = json.Unmarshal(bytes, &varGetUserQuota2XXResponseMovie); err == nil {
		*o = GetUserQuota2XXResponseMovie(varGetUserQuota2XXResponseMovie)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "days")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "used")
		delete(additionalProperties, "remaining")
		delete(additionalProperties, "restricted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserQuota2XXResponseMovie struct {
	value *GetUserQuota2XXResponseMovie
	isSet bool
}

func (v NullableGetUserQuota2XXResponseMovie) Get() *GetUserQuota2XXResponseMovie {
	return v.value
}

func (v *NullableGetUserQuota2XXResponseMovie) Set(val *GetUserQuota2XXResponseMovie) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserQuota2XXResponseMovie) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserQuota2XXResponseMovie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserQuota2XXResponseMovie(val *GetUserQuota2XXResponseMovie) *NullableGetUserQuota2XXResponseMovie {
	return &NullableGetUserQuota2XXResponseMovie{value: val, isSet: true}
}

func (v NullableGetUserQuota2XXResponseMovie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserQuota2XXResponseMovie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

