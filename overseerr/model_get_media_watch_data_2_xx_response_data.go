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

// checks if the GetMediaWatchData2XXResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMediaWatchData2XXResponseData{}

// GetMediaWatchData2XXResponseData struct for GetMediaWatchData2XXResponseData
type GetMediaWatchData2XXResponseData struct {
	PlayCount7Days *float32 `json:"playCount7Days,omitempty"`
	PlayCount30Days *float32 `json:"playCount30Days,omitempty"`
	PlayCount *float32 `json:"playCount,omitempty"`
	Users []User `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMediaWatchData2XXResponseData GetMediaWatchData2XXResponseData

// NewGetMediaWatchData2XXResponseData instantiates a new GetMediaWatchData2XXResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMediaWatchData2XXResponseData() *GetMediaWatchData2XXResponseData {
	this := GetMediaWatchData2XXResponseData{}
	return &this
}

// NewGetMediaWatchData2XXResponseDataWithDefaults instantiates a new GetMediaWatchData2XXResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMediaWatchData2XXResponseDataWithDefaults() *GetMediaWatchData2XXResponseData {
	this := GetMediaWatchData2XXResponseData{}
	return &this
}

// GetPlayCount7Days returns the PlayCount7Days field value if set, zero value otherwise.
func (o *GetMediaWatchData2XXResponseData) GetPlayCount7Days() float32 {
	if o == nil || IsNil(o.PlayCount7Days) {
		var ret float32
		return ret
	}
	return *o.PlayCount7Days
}

// GetPlayCount7DaysOk returns a tuple with the PlayCount7Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMediaWatchData2XXResponseData) GetPlayCount7DaysOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount7Days) {
		return nil, false
	}
	return o.PlayCount7Days, true
}

// HasPlayCount7Days returns a boolean if a field has been set.
func (o *GetMediaWatchData2XXResponseData) HasPlayCount7Days() bool {
	if o != nil && !IsNil(o.PlayCount7Days) {
		return true
	}

	return false
}

// SetPlayCount7Days gets a reference to the given float32 and assigns it to the PlayCount7Days field.
func (o *GetMediaWatchData2XXResponseData) SetPlayCount7Days(v float32) {
	o.PlayCount7Days = &v
}

// GetPlayCount30Days returns the PlayCount30Days field value if set, zero value otherwise.
func (o *GetMediaWatchData2XXResponseData) GetPlayCount30Days() float32 {
	if o == nil || IsNil(o.PlayCount30Days) {
		var ret float32
		return ret
	}
	return *o.PlayCount30Days
}

// GetPlayCount30DaysOk returns a tuple with the PlayCount30Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMediaWatchData2XXResponseData) GetPlayCount30DaysOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount30Days) {
		return nil, false
	}
	return o.PlayCount30Days, true
}

// HasPlayCount30Days returns a boolean if a field has been set.
func (o *GetMediaWatchData2XXResponseData) HasPlayCount30Days() bool {
	if o != nil && !IsNil(o.PlayCount30Days) {
		return true
	}

	return false
}

// SetPlayCount30Days gets a reference to the given float32 and assigns it to the PlayCount30Days field.
func (o *GetMediaWatchData2XXResponseData) SetPlayCount30Days(v float32) {
	o.PlayCount30Days = &v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *GetMediaWatchData2XXResponseData) GetPlayCount() float32 {
	if o == nil || IsNil(o.PlayCount) {
		var ret float32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMediaWatchData2XXResponseData) GetPlayCountOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount) {
		return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *GetMediaWatchData2XXResponseData) HasPlayCount() bool {
	if o != nil && !IsNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given float32 and assigns it to the PlayCount field.
func (o *GetMediaWatchData2XXResponseData) SetPlayCount(v float32) {
	o.PlayCount = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetMediaWatchData2XXResponseData) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMediaWatchData2XXResponseData) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetMediaWatchData2XXResponseData) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *GetMediaWatchData2XXResponseData) SetUsers(v []User) {
	o.Users = v
}

func (o GetMediaWatchData2XXResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMediaWatchData2XXResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlayCount7Days) {
		toSerialize["playCount7Days"] = o.PlayCount7Days
	}
	if !IsNil(o.PlayCount30Days) {
		toSerialize["playCount30Days"] = o.PlayCount30Days
	}
	if !IsNil(o.PlayCount) {
		toSerialize["playCount"] = o.PlayCount
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetMediaWatchData2XXResponseData) UnmarshalJSON(data []byte) (err error) {
	varGetMediaWatchData2XXResponseData := _GetMediaWatchData2XXResponseData{}

	err = json.Unmarshal(data, &varGetMediaWatchData2XXResponseData)

	if err != nil {
		return err
	}

	*o = GetMediaWatchData2XXResponseData(varGetMediaWatchData2XXResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "playCount7Days")
		delete(additionalProperties, "playCount30Days")
		delete(additionalProperties, "playCount")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMediaWatchData2XXResponseData struct {
	value *GetMediaWatchData2XXResponseData
	isSet bool
}

func (v NullableGetMediaWatchData2XXResponseData) Get() *GetMediaWatchData2XXResponseData {
	return v.value
}

func (v *NullableGetMediaWatchData2XXResponseData) Set(val *GetMediaWatchData2XXResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMediaWatchData2XXResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMediaWatchData2XXResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMediaWatchData2XXResponseData(val *GetMediaWatchData2XXResponseData) *NullableGetMediaWatchData2XXResponseData {
	return &NullableGetMediaWatchData2XXResponseData{value: val, isSet: true}
}

func (v NullableGetMediaWatchData2XXResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMediaWatchData2XXResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


