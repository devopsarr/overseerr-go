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

// GetMedia2XXResponse struct for GetMedia2XXResponse
type GetMedia2XXResponse struct {
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
	Results []*MediaInfo `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetMedia2XXResponse GetMedia2XXResponse

// NewGetMedia2XXResponse instantiates a new GetMedia2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMedia2XXResponse() *GetMedia2XXResponse {
	this := GetMedia2XXResponse{}
	return &this
}

// NewGetMedia2XXResponseWithDefaults instantiates a new GetMedia2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMedia2XXResponseWithDefaults() *GetMedia2XXResponse {
	this := GetMedia2XXResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *GetMedia2XXResponse) GetPageInfo() PageInfo {
	if o == nil || isNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMedia2XXResponse) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || isNil(o.PageInfo) {
    return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *GetMedia2XXResponse) HasPageInfo() bool {
	if o != nil && !isNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *GetMedia2XXResponse) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetMedia2XXResponse) GetResults() []*MediaInfo {
	if o == nil || isNil(o.Results) {
		var ret []*MediaInfo
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMedia2XXResponse) GetResultsOk() ([]*MediaInfo, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetMedia2XXResponse) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MediaInfo and assigns it to the Results field.
func (o *GetMedia2XXResponse) SetResults(v []*MediaInfo) {
	o.Results = v
}

func (o GetMedia2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetMedia2XXResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetMedia2XXResponse := _GetMedia2XXResponse{}

	if err = json.Unmarshal(bytes, &varGetMedia2XXResponse); err == nil {
		*o = GetMedia2XXResponse(varGetMedia2XXResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pageInfo")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetMedia2XXResponse struct {
	value *GetMedia2XXResponse
	isSet bool
}

func (v NullableGetMedia2XXResponse) Get() *GetMedia2XXResponse {
	return v.value
}

func (v *NullableGetMedia2XXResponse) Set(val *GetMedia2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMedia2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMedia2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMedia2XXResponse(val *GetMedia2XXResponse) *NullableGetMedia2XXResponse {
	return &NullableGetMedia2XXResponse{value: val, isSet: true}
}

func (v NullableGetMedia2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMedia2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

