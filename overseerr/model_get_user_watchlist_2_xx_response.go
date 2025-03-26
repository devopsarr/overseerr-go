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

// checks if the GetUserWatchlist2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserWatchlist2XXResponse{}

// GetUserWatchlist2XXResponse struct for GetUserWatchlist2XXResponse
type GetUserWatchlist2XXResponse struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Results []GetUserWatchlist2XXResponseResultsInner `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetUserWatchlist2XXResponse GetUserWatchlist2XXResponse

// NewGetUserWatchlist2XXResponse instantiates a new GetUserWatchlist2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserWatchlist2XXResponse() *GetUserWatchlist2XXResponse {
	this := GetUserWatchlist2XXResponse{}
	return &this
}

// NewGetUserWatchlist2XXResponseWithDefaults instantiates a new GetUserWatchlist2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserWatchlist2XXResponseWithDefaults() *GetUserWatchlist2XXResponse {
	this := GetUserWatchlist2XXResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponse) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponse) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *GetUserWatchlist2XXResponse) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponse) GetTotalPages() float32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponse) GetTotalPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *GetUserWatchlist2XXResponse) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponse) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponse) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *GetUserWatchlist2XXResponse) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetUserWatchlist2XXResponse) GetResults() []GetUserWatchlist2XXResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []GetUserWatchlist2XXResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserWatchlist2XXResponse) GetResultsOk() ([]GetUserWatchlist2XXResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetUserWatchlist2XXResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GetUserWatchlist2XXResponseResultsInner and assigns it to the Results field.
func (o *GetUserWatchlist2XXResponse) SetResults(v []GetUserWatchlist2XXResponseResultsInner) {
	o.Results = v
}

func (o GetUserWatchlist2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserWatchlist2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetUserWatchlist2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetUserWatchlist2XXResponse := _GetUserWatchlist2XXResponse{}

	err = json.Unmarshal(data, &varGetUserWatchlist2XXResponse)

	if err != nil {
		return err
	}

	*o = GetUserWatchlist2XXResponse(varGetUserWatchlist2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "totalPages")
		delete(additionalProperties, "totalResults")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetUserWatchlist2XXResponse struct {
	value *GetUserWatchlist2XXResponse
	isSet bool
}

func (v NullableGetUserWatchlist2XXResponse) Get() *GetUserWatchlist2XXResponse {
	return v.value
}

func (v *NullableGetUserWatchlist2XXResponse) Set(val *GetUserWatchlist2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserWatchlist2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserWatchlist2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserWatchlist2XXResponse(val *GetUserWatchlist2XXResponse) *NullableGetUserWatchlist2XXResponse {
	return &NullableGetUserWatchlist2XXResponse{value: val, isSet: true}
}

func (v NullableGetUserWatchlist2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserWatchlist2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


