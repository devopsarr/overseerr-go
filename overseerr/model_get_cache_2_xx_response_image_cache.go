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

// checks if the GetCache2XXResponseImageCache type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCache2XXResponseImageCache{}

// GetCache2XXResponseImageCache struct for GetCache2XXResponseImageCache
type GetCache2XXResponseImageCache struct {
	Tmdb *GetCache2XXResponseImageCacheTmdb `json:"tmdb,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCache2XXResponseImageCache GetCache2XXResponseImageCache

// NewGetCache2XXResponseImageCache instantiates a new GetCache2XXResponseImageCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCache2XXResponseImageCache() *GetCache2XXResponseImageCache {
	this := GetCache2XXResponseImageCache{}
	return &this
}

// NewGetCache2XXResponseImageCacheWithDefaults instantiates a new GetCache2XXResponseImageCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCache2XXResponseImageCacheWithDefaults() *GetCache2XXResponseImageCache {
	this := GetCache2XXResponseImageCache{}
	return &this
}

// GetTmdb returns the Tmdb field value if set, zero value otherwise.
func (o *GetCache2XXResponseImageCache) GetTmdb() GetCache2XXResponseImageCacheTmdb {
	if o == nil || IsNil(o.Tmdb) {
		var ret GetCache2XXResponseImageCacheTmdb
		return ret
	}
	return *o.Tmdb
}

// GetTmdbOk returns a tuple with the Tmdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache2XXResponseImageCache) GetTmdbOk() (*GetCache2XXResponseImageCacheTmdb, bool) {
	if o == nil || IsNil(o.Tmdb) {
		return nil, false
	}
	return o.Tmdb, true
}

// HasTmdb returns a boolean if a field has been set.
func (o *GetCache2XXResponseImageCache) HasTmdb() bool {
	if o != nil && !IsNil(o.Tmdb) {
		return true
	}

	return false
}

// SetTmdb gets a reference to the given GetCache2XXResponseImageCacheTmdb and assigns it to the Tmdb field.
func (o *GetCache2XXResponseImageCache) SetTmdb(v GetCache2XXResponseImageCacheTmdb) {
	o.Tmdb = &v
}

func (o GetCache2XXResponseImageCache) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCache2XXResponseImageCache) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tmdb) {
		toSerialize["tmdb"] = o.Tmdb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCache2XXResponseImageCache) UnmarshalJSON(data []byte) (err error) {
	varGetCache2XXResponseImageCache := _GetCache2XXResponseImageCache{}

	err = json.Unmarshal(data, &varGetCache2XXResponseImageCache)

	if err != nil {
		return err
	}

	*o = GetCache2XXResponseImageCache(varGetCache2XXResponseImageCache)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tmdb")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCache2XXResponseImageCache struct {
	value *GetCache2XXResponseImageCache
	isSet bool
}

func (v NullableGetCache2XXResponseImageCache) Get() *GetCache2XXResponseImageCache {
	return v.value
}

func (v *NullableGetCache2XXResponseImageCache) Set(val *GetCache2XXResponseImageCache) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCache2XXResponseImageCache) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCache2XXResponseImageCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCache2XXResponseImageCache(val *GetCache2XXResponseImageCache) *NullableGetCache2XXResponseImageCache {
	return &NullableGetCache2XXResponseImageCache{value: val, isSet: true}
}

func (v NullableGetCache2XXResponseImageCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCache2XXResponseImageCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


