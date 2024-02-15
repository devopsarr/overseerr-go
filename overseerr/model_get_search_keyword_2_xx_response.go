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

// checks if the GetSearchKeyword2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSearchKeyword2XXResponse{}

// GetSearchKeyword2XXResponse struct for GetSearchKeyword2XXResponse
type GetSearchKeyword2XXResponse struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Results []Keyword `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSearchKeyword2XXResponse GetSearchKeyword2XXResponse

// NewGetSearchKeyword2XXResponse instantiates a new GetSearchKeyword2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSearchKeyword2XXResponse() *GetSearchKeyword2XXResponse {
	this := GetSearchKeyword2XXResponse{}
	return &this
}

// NewGetSearchKeyword2XXResponseWithDefaults instantiates a new GetSearchKeyword2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSearchKeyword2XXResponseWithDefaults() *GetSearchKeyword2XXResponse {
	this := GetSearchKeyword2XXResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetSearchKeyword2XXResponse) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearchKeyword2XXResponse) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetSearchKeyword2XXResponse) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *GetSearchKeyword2XXResponse) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetSearchKeyword2XXResponse) GetTotalPages() float32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearchKeyword2XXResponse) GetTotalPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetSearchKeyword2XXResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *GetSearchKeyword2XXResponse) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *GetSearchKeyword2XXResponse) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearchKeyword2XXResponse) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *GetSearchKeyword2XXResponse) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *GetSearchKeyword2XXResponse) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetSearchKeyword2XXResponse) GetResults() []Keyword {
	if o == nil || IsNil(o.Results) {
		var ret []Keyword
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearchKeyword2XXResponse) GetResultsOk() ([]Keyword, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetSearchKeyword2XXResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Keyword and assigns it to the Results field.
func (o *GetSearchKeyword2XXResponse) SetResults(v []Keyword) {
	o.Results = v
}

func (o GetSearchKeyword2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSearchKeyword2XXResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSearchKeyword2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSearchKeyword2XXResponse := _GetSearchKeyword2XXResponse{}

	err = json.Unmarshal(data, &varGetSearchKeyword2XXResponse)

	if err != nil {
		return err
	}

	*o = GetSearchKeyword2XXResponse(varGetSearchKeyword2XXResponse)

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

type NullableGetSearchKeyword2XXResponse struct {
	value *GetSearchKeyword2XXResponse
	isSet bool
}

func (v NullableGetSearchKeyword2XXResponse) Get() *GetSearchKeyword2XXResponse {
	return v.value
}

func (v *NullableGetSearchKeyword2XXResponse) Set(val *GetSearchKeyword2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSearchKeyword2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSearchKeyword2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSearchKeyword2XXResponse(val *GetSearchKeyword2XXResponse) *NullableGetSearchKeyword2XXResponse {
	return &NullableGetSearchKeyword2XXResponse{value: val, isSet: true}
}

func (v NullableGetSearchKeyword2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSearchKeyword2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


