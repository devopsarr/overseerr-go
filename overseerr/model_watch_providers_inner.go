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

// WatchProvidersInner struct for WatchProvidersInner
type WatchProvidersInner struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	Link *string `json:"link,omitempty"`
	Buy []*WatchProviderDetails `json:"buy,omitempty"`
	Flatrate []*WatchProviderDetails `json:"flatrate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WatchProvidersInner WatchProvidersInner

// NewWatchProvidersInner instantiates a new WatchProvidersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchProvidersInner() *WatchProvidersInner {
	this := WatchProvidersInner{}
	return &this
}

// NewWatchProvidersInnerWithDefaults instantiates a new WatchProvidersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchProvidersInnerWithDefaults() *WatchProvidersInner {
	this := WatchProvidersInner{}
	return &this
}

// GetIso31661 returns the Iso31661 field value if set, zero value otherwise.
func (o *WatchProvidersInner) GetIso31661() string {
	if o == nil || isNil(o.Iso31661) {
		var ret string
		return ret
	}
	return *o.Iso31661
}

// GetIso31661Ok returns a tuple with the Iso31661 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProvidersInner) GetIso31661Ok() (*string, bool) {
	if o == nil || isNil(o.Iso31661) {
    return nil, false
	}
	return o.Iso31661, true
}

// HasIso31661 returns a boolean if a field has been set.
func (o *WatchProvidersInner) HasIso31661() bool {
	if o != nil && !isNil(o.Iso31661) {
		return true
	}

	return false
}

// SetIso31661 gets a reference to the given string and assigns it to the Iso31661 field.
func (o *WatchProvidersInner) SetIso31661(v string) {
	o.Iso31661 = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *WatchProvidersInner) GetLink() string {
	if o == nil || isNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProvidersInner) GetLinkOk() (*string, bool) {
	if o == nil || isNil(o.Link) {
    return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *WatchProvidersInner) HasLink() bool {
	if o != nil && !isNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *WatchProvidersInner) SetLink(v string) {
	o.Link = &v
}

// GetBuy returns the Buy field value if set, zero value otherwise.
func (o *WatchProvidersInner) GetBuy() []*WatchProviderDetails {
	if o == nil || isNil(o.Buy) {
		var ret []*WatchProviderDetails
		return ret
	}
	return o.Buy
}

// GetBuyOk returns a tuple with the Buy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProvidersInner) GetBuyOk() ([]*WatchProviderDetails, bool) {
	if o == nil || isNil(o.Buy) {
    return nil, false
	}
	return o.Buy, true
}

// HasBuy returns a boolean if a field has been set.
func (o *WatchProvidersInner) HasBuy() bool {
	if o != nil && !isNil(o.Buy) {
		return true
	}

	return false
}

// SetBuy gets a reference to the given []WatchProviderDetails and assigns it to the Buy field.
func (o *WatchProvidersInner) SetBuy(v []*WatchProviderDetails) {
	o.Buy = v
}

// GetFlatrate returns the Flatrate field value if set, zero value otherwise.
func (o *WatchProvidersInner) GetFlatrate() []*WatchProviderDetails {
	if o == nil || isNil(o.Flatrate) {
		var ret []*WatchProviderDetails
		return ret
	}
	return o.Flatrate
}

// GetFlatrateOk returns a tuple with the Flatrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProvidersInner) GetFlatrateOk() ([]*WatchProviderDetails, bool) {
	if o == nil || isNil(o.Flatrate) {
    return nil, false
	}
	return o.Flatrate, true
}

// HasFlatrate returns a boolean if a field has been set.
func (o *WatchProvidersInner) HasFlatrate() bool {
	if o != nil && !isNil(o.Flatrate) {
		return true
	}

	return false
}

// SetFlatrate gets a reference to the given []WatchProviderDetails and assigns it to the Flatrate field.
func (o *WatchProvidersInner) SetFlatrate(v []*WatchProviderDetails) {
	o.Flatrate = v
}

func (o WatchProvidersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Iso31661) {
		toSerialize["iso_3166_1"] = o.Iso31661
	}
	if !isNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !isNil(o.Buy) {
		toSerialize["buy"] = o.Buy
	}
	if !isNil(o.Flatrate) {
		toSerialize["flatrate"] = o.Flatrate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchProvidersInner) UnmarshalJSON(bytes []byte) (err error) {
	varWatchProvidersInner := _WatchProvidersInner{}

	if err = json.Unmarshal(bytes, &varWatchProvidersInner); err == nil {
		*o = WatchProvidersInner(varWatchProvidersInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iso_3166_1")
		delete(additionalProperties, "link")
		delete(additionalProperties, "buy")
		delete(additionalProperties, "flatrate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchProvidersInner struct {
	value *WatchProvidersInner
	isSet bool
}

func (v NullableWatchProvidersInner) Get() *WatchProvidersInner {
	return v.value
}

func (v *NullableWatchProvidersInner) Set(val *WatchProvidersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchProvidersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchProvidersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchProvidersInner(val *WatchProvidersInner) *NullableWatchProvidersInner {
	return &NullableWatchProvidersInner{value: val, isSet: true}
}

func (v NullableWatchProvidersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchProvidersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


