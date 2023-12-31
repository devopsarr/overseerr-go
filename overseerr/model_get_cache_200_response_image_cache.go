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

// GetCache200ResponseImageCache struct for GetCache200ResponseImageCache
type GetCache200ResponseImageCache struct {
	Tmdb *GetCache200ResponseImageCacheTmdb `json:"tmdb,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCache200ResponseImageCache GetCache200ResponseImageCache

// NewGetCache200ResponseImageCache instantiates a new GetCache200ResponseImageCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCache200ResponseImageCache() *GetCache200ResponseImageCache {
	this := GetCache200ResponseImageCache{}
	return &this
}

// NewGetCache200ResponseImageCacheWithDefaults instantiates a new GetCache200ResponseImageCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCache200ResponseImageCacheWithDefaults() *GetCache200ResponseImageCache {
	this := GetCache200ResponseImageCache{}
	return &this
}

// GetTmdb returns the Tmdb field value if set, zero value otherwise.
func (o *GetCache200ResponseImageCache) GetTmdb() GetCache200ResponseImageCacheTmdb {
	if o == nil || isNil(o.Tmdb) {
		var ret GetCache200ResponseImageCacheTmdb
		return ret
	}
	return *o.Tmdb
}

// GetTmdbOk returns a tuple with the Tmdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache200ResponseImageCache) GetTmdbOk() (*GetCache200ResponseImageCacheTmdb, bool) {
	if o == nil || isNil(o.Tmdb) {
    return nil, false
	}
	return o.Tmdb, true
}

// HasTmdb returns a boolean if a field has been set.
func (o *GetCache200ResponseImageCache) HasTmdb() bool {
	if o != nil && !isNil(o.Tmdb) {
		return true
	}

	return false
}

// SetTmdb gets a reference to the given GetCache200ResponseImageCacheTmdb and assigns it to the Tmdb field.
func (o *GetCache200ResponseImageCache) SetTmdb(v GetCache200ResponseImageCacheTmdb) {
	o.Tmdb = &v
}

func (o GetCache200ResponseImageCache) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tmdb) {
		toSerialize["tmdb"] = o.Tmdb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetCache200ResponseImageCache) UnmarshalJSON(bytes []byte) (err error) {
	varGetCache200ResponseImageCache := _GetCache200ResponseImageCache{}

	if err = json.Unmarshal(bytes, &varGetCache200ResponseImageCache); err == nil {
		*o = GetCache200ResponseImageCache(varGetCache200ResponseImageCache)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tmdb")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCache200ResponseImageCache struct {
	value *GetCache200ResponseImageCache
	isSet bool
}

func (v NullableGetCache200ResponseImageCache) Get() *GetCache200ResponseImageCache {
	return v.value
}

func (v *NullableGetCache200ResponseImageCache) Set(val *GetCache200ResponseImageCache) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCache200ResponseImageCache) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCache200ResponseImageCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCache200ResponseImageCache(val *GetCache200ResponseImageCache) *NullableGetCache200ResponseImageCache {
	return &NullableGetCache200ResponseImageCache{value: val, isSet: true}
}

func (v NullableGetCache200ResponseImageCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCache200ResponseImageCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


