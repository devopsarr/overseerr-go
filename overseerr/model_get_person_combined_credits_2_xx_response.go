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

// checks if the GetPersonCombinedCredits2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPersonCombinedCredits2XXResponse{}

// GetPersonCombinedCredits2XXResponse struct for GetPersonCombinedCredits2XXResponse
type GetPersonCombinedCredits2XXResponse struct {
	Cast []CreditCast `json:"cast,omitempty"`
	Crew []CreditCrew `json:"crew,omitempty"`
	Id *float32 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPersonCombinedCredits2XXResponse GetPersonCombinedCredits2XXResponse

// NewGetPersonCombinedCredits2XXResponse instantiates a new GetPersonCombinedCredits2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPersonCombinedCredits2XXResponse() *GetPersonCombinedCredits2XXResponse {
	this := GetPersonCombinedCredits2XXResponse{}
	return &this
}

// NewGetPersonCombinedCredits2XXResponseWithDefaults instantiates a new GetPersonCombinedCredits2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPersonCombinedCredits2XXResponseWithDefaults() *GetPersonCombinedCredits2XXResponse {
	this := GetPersonCombinedCredits2XXResponse{}
	return &this
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *GetPersonCombinedCredits2XXResponse) GetCast() []CreditCast {
	if o == nil || IsNil(o.Cast) {
		var ret []CreditCast
		return ret
	}
	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonCombinedCredits2XXResponse) GetCastOk() ([]CreditCast, bool) {
	if o == nil || IsNil(o.Cast) {
		return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *GetPersonCombinedCredits2XXResponse) HasCast() bool {
	if o != nil && !IsNil(o.Cast) {
		return true
	}

	return false
}

// SetCast gets a reference to the given []CreditCast and assigns it to the Cast field.
func (o *GetPersonCombinedCredits2XXResponse) SetCast(v []CreditCast) {
	o.Cast = v
}

// GetCrew returns the Crew field value if set, zero value otherwise.
func (o *GetPersonCombinedCredits2XXResponse) GetCrew() []CreditCrew {
	if o == nil || IsNil(o.Crew) {
		var ret []CreditCrew
		return ret
	}
	return o.Crew
}

// GetCrewOk returns a tuple with the Crew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonCombinedCredits2XXResponse) GetCrewOk() ([]CreditCrew, bool) {
	if o == nil || IsNil(o.Crew) {
		return nil, false
	}
	return o.Crew, true
}

// HasCrew returns a boolean if a field has been set.
func (o *GetPersonCombinedCredits2XXResponse) HasCrew() bool {
	if o != nil && !IsNil(o.Crew) {
		return true
	}

	return false
}

// SetCrew gets a reference to the given []CreditCrew and assigns it to the Crew field.
func (o *GetPersonCombinedCredits2XXResponse) SetCrew(v []CreditCrew) {
	o.Crew = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPersonCombinedCredits2XXResponse) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonCombinedCredits2XXResponse) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPersonCombinedCredits2XXResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GetPersonCombinedCredits2XXResponse) SetId(v float32) {
	o.Id = &v
}

func (o GetPersonCombinedCredits2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPersonCombinedCredits2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cast) {
		toSerialize["cast"] = o.Cast
	}
	if !IsNil(o.Crew) {
		toSerialize["crew"] = o.Crew
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPersonCombinedCredits2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetPersonCombinedCredits2XXResponse := _GetPersonCombinedCredits2XXResponse{}

	err = json.Unmarshal(data, &varGetPersonCombinedCredits2XXResponse)

	if err != nil {
		return err
	}

	*o = GetPersonCombinedCredits2XXResponse(varGetPersonCombinedCredits2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cast")
		delete(additionalProperties, "crew")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPersonCombinedCredits2XXResponse struct {
	value *GetPersonCombinedCredits2XXResponse
	isSet bool
}

func (v NullableGetPersonCombinedCredits2XXResponse) Get() *GetPersonCombinedCredits2XXResponse {
	return v.value
}

func (v *NullableGetPersonCombinedCredits2XXResponse) Set(val *GetPersonCombinedCredits2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPersonCombinedCredits2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPersonCombinedCredits2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPersonCombinedCredits2XXResponse(val *GetPersonCombinedCredits2XXResponse) *NullableGetPersonCombinedCredits2XXResponse {
	return &NullableGetPersonCombinedCredits2XXResponse{value: val, isSet: true}
}

func (v NullableGetPersonCombinedCredits2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPersonCombinedCredits2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


