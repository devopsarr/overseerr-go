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

// GetDiscoverTvNetworkByNetworkId200Response struct for GetDiscoverTvNetworkByNetworkId200Response
type GetDiscoverTvNetworkByNetworkId200Response struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Network *Network `json:"network,omitempty"`
	Results []*TvResult `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDiscoverTvNetworkByNetworkId200Response GetDiscoverTvNetworkByNetworkId200Response

// NewGetDiscoverTvNetworkByNetworkId200Response instantiates a new GetDiscoverTvNetworkByNetworkId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoverTvNetworkByNetworkId200Response() *GetDiscoverTvNetworkByNetworkId200Response {
	this := GetDiscoverTvNetworkByNetworkId200Response{}
	return &this
}

// NewGetDiscoverTvNetworkByNetworkId200ResponseWithDefaults instantiates a new GetDiscoverTvNetworkByNetworkId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoverTvNetworkByNetworkId200ResponseWithDefaults() *GetDiscoverTvNetworkByNetworkId200Response {
	this := GetDiscoverTvNetworkByNetworkId200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetPage() float32 {
	if o == nil || isNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetPageOk() (*float32, bool) {
	if o == nil || isNil(o.Page) {
    return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) HasPage() bool {
	if o != nil && !isNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *GetDiscoverTvNetworkByNetworkId200Response) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetTotalPages() float32 {
	if o == nil || isNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetTotalPagesOk() (*float32, bool) {
	if o == nil || isNil(o.TotalPages) {
    return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) HasTotalPages() bool {
	if o != nil && !isNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *GetDiscoverTvNetworkByNetworkId200Response) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetTotalResults() float32 {
	if o == nil || isNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetTotalResultsOk() (*float32, bool) {
	if o == nil || isNil(o.TotalResults) {
    return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) HasTotalResults() bool {
	if o != nil && !isNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *GetDiscoverTvNetworkByNetworkId200Response) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetNetwork() Network {
	if o == nil || isNil(o.Network) {
		var ret Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetNetworkOk() (*Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given Network and assigns it to the Network field.
func (o *GetDiscoverTvNetworkByNetworkId200Response) SetNetwork(v Network) {
	o.Network = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetResults() []*TvResult {
	if o == nil || isNil(o.Results) {
		var ret []*TvResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) GetResultsOk() ([]*TvResult, bool) {
	if o == nil || isNil(o.Results) {
    return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetDiscoverTvNetworkByNetworkId200Response) HasResults() bool {
	if o != nil && !isNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []TvResult and assigns it to the Results field.
func (o *GetDiscoverTvNetworkByNetworkId200Response) SetResults(v []*TvResult) {
	o.Results = v
}

func (o GetDiscoverTvNetworkByNetworkId200Response) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetDiscoverTvNetworkByNetworkId200Response) UnmarshalJSON(bytes []byte) (err error) {
	varGetDiscoverTvNetworkByNetworkId200Response := _GetDiscoverTvNetworkByNetworkId200Response{}

	if err = json.Unmarshal(bytes, &varGetDiscoverTvNetworkByNetworkId200Response); err == nil {
		*o = GetDiscoverTvNetworkByNetworkId200Response(varGetDiscoverTvNetworkByNetworkId200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "totalPages")
		delete(additionalProperties, "totalResults")
		delete(additionalProperties, "network")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDiscoverTvNetworkByNetworkId200Response struct {
	value *GetDiscoverTvNetworkByNetworkId200Response
	isSet bool
}

func (v NullableGetDiscoverTvNetworkByNetworkId200Response) Get() *GetDiscoverTvNetworkByNetworkId200Response {
	return v.value
}

func (v *NullableGetDiscoverTvNetworkByNetworkId200Response) Set(val *GetDiscoverTvNetworkByNetworkId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoverTvNetworkByNetworkId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoverTvNetworkByNetworkId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoverTvNetworkByNetworkId200Response(val *GetDiscoverTvNetworkByNetworkId200Response) *NullableGetDiscoverTvNetworkByNetworkId200Response {
	return &NullableGetDiscoverTvNetworkByNetworkId200Response{value: val, isSet: true}
}

func (v NullableGetDiscoverTvNetworkByNetworkId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoverTvNetworkByNetworkId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


