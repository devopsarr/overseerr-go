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

// GetSearch200Response struct for GetSearch200Response
type GetSearch200Response struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Results []*GetSearch200ResponseResultsInner `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSearch200Response GetSearch200Response

// NewGetSearch200Response instantiates a new GetSearch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSearch200Response() *GetSearch200Response {
	this := GetSearch200Response{}
	return &this
}

// NewGetSearch200ResponseWithDefaults instantiates a new GetSearch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSearch200ResponseWithDefaults() *GetSearch200Response {
	this := GetSearch200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetSearch200Response) GetPage() float32 {
	if o == nil || isNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearch200Response) GetPageOk() (*float32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetSearch200Response) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *GetSearch200Response) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetSearch200Response) GetTotalPages() float32 {
	if o == nil || isNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearch200Response) GetTotalPagesOk() (*float32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetSearch200Response) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *GetSearch200Response) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *GetSearch200Response) GetTotalResults() float32 {
	if o == nil || isNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearch200Response) GetTotalResultsOk() (*float32, bool) {
	if o == nil || isNil(o.TotalResults) {
    return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *GetSearch200Response) HasTotalResults() bool {
	if o != nil && !isNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *GetSearch200Response) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetSearch200Response) GetResults() []*GetSearch200ResponseResultsInner {
	if o == nil || isNil(o.Results) {
		var ret []*GetSearch200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSearch200Response) GetResultsOk() ([]*GetSearch200ResponseResultsInner, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetSearch200Response) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GetSearch200ResponseResultsInner and assigns it to the Results field.
func (o *GetSearch200Response) SetResults(v []*GetSearch200ResponseResultsInner) {
	o.Results = v
}

func (o GetSearch200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !isNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !isNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetSearch200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetSearch200Response := _GetSearch200Response{}

	if err = json.Unmarshal(bytes, &varGetSearch200Response); err == nil {
		*o = GetSearch200Response(varGetSearch200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "totalPages")
		delete(additionalProperties, "totalResults")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSearch200Response struct {
	value *GetSearch200Response
	isSet bool
}

func (v NullableGetSearch200Response) Get() *GetSearch200Response {
	return v.value
}

func (v *NullableGetSearch200Response) Set(val *GetSearch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSearch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSearch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSearch200Response(val *GetSearch200Response) *NullableGetSearch200Response {
	return &NullableGetSearch200Response{value: val, isSet: true}
}

func (v NullableGetSearch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSearch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


