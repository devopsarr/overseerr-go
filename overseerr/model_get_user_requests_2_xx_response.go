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

// checks if the GetUserRequests2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserRequests2XXResponse{}

// GetUserRequests2XXResponse struct for GetUserRequests2XXResponse
type GetUserRequests2XXResponse struct {
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
	Results []MediaRequest `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserRequests2XXResponse GetUserRequests2XXResponse

// NewGetUserRequests2XXResponse instantiates a new GetUserRequests2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserRequests2XXResponse() *GetUserRequests2XXResponse {
	this := GetUserRequests2XXResponse{}
	return &this
}

// NewGetUserRequests2XXResponseWithDefaults instantiates a new GetUserRequests2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserRequests2XXResponseWithDefaults() *GetUserRequests2XXResponse {
	this := GetUserRequests2XXResponse{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *GetUserRequests2XXResponse) GetPageInfo() PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequests2XXResponse) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *GetUserRequests2XXResponse) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *GetUserRequests2XXResponse) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetUserRequests2XXResponse) GetResults() []MediaRequest {
	if o == nil || IsNil(o.Results) {
		var ret []MediaRequest
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequests2XXResponse) GetResultsOk() ([]MediaRequest, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetUserRequests2XXResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MediaRequest and assigns it to the Results field.
func (o *GetUserRequests2XXResponse) SetResults(v []MediaRequest) {
	o.Results = v
}

func (o GetUserRequests2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserRequests2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserRequests2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUserRequests2XXResponse := _GetUserRequests2XXResponse{}

	err = json.Unmarshal(data, &varGetUserRequests2XXResponse)

	if err != nil {
		return err
	}

	*o = GetUserRequests2XXResponse(varGetUserRequests2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pageInfo")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserRequests2XXResponse struct {
	value *GetUserRequests2XXResponse
	isSet bool
}

func (v NullableGetUserRequests2XXResponse) Get() *GetUserRequests2XXResponse {
	return v.value
}

func (v *NullableGetUserRequests2XXResponse) Set(val *GetUserRequests2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRequests2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRequests2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRequests2XXResponse(val *GetUserRequests2XXResponse) *NullableGetUserRequests2XXResponse {
	return &NullableGetUserRequests2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserRequests2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRequests2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


