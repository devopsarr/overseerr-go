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

// GetUserRequests200Response struct for GetUserRequests200Response
type GetUserRequests200Response struct {
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
	Results []*MediaRequest `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserRequests200Response GetUserRequests200Response

// NewGetUserRequests200Response instantiates a new GetUserRequests200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserRequests200Response() *GetUserRequests200Response {
	this := GetUserRequests200Response{}
	return &this
}

// NewGetUserRequests200ResponseWithDefaults instantiates a new GetUserRequests200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserRequests200ResponseWithDefaults() *GetUserRequests200Response {
	this := GetUserRequests200Response{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *GetUserRequests200Response) GetPageInfo() PageInfo {
	if o == nil || isNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequests200Response) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || isNil(o.PageInfo) {
    return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *GetUserRequests200Response) HasPageInfo() bool {
	if o != nil && !isNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *GetUserRequests200Response) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetUserRequests200Response) GetResults() []*MediaRequest {
	if o == nil || isNil(o.Results) {
		var ret []*MediaRequest
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserRequests200Response) GetResultsOk() ([]*MediaRequest, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetUserRequests200Response) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MediaRequest and assigns it to the Results field.
func (o *GetUserRequests200Response) SetResults(v []*MediaRequest) {
	o.Results = v
}

func (o GetUserRequests200Response) MarshalJSON() ([]byte, error) {
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

func (o *GetUserRequests200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetUserRequests200Response := _GetUserRequests200Response{}

	if err = json.Unmarshal(bytes, &varGetUserRequests200Response); err == nil {
		*o = GetUserRequests200Response(varGetUserRequests200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pageInfo")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserRequests200Response struct {
	value *GetUserRequests200Response
	isSet bool
}

func (v NullableGetUserRequests200Response) Get() *GetUserRequests200Response {
	return v.value
}

func (v *NullableGetUserRequests200Response) Set(val *GetUserRequests200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserRequests200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserRequests200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserRequests200Response(val *GetUserRequests200Response) *NullableGetUserRequests200Response {
	return &NullableGetUserRequests200Response{value: val, isSet: true}
}

func (v NullableGetUserRequests200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserRequests200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


