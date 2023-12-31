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

// GetCache200ResponseImageCacheTmdb struct for GetCache200ResponseImageCacheTmdb
type GetCache200ResponseImageCacheTmdb struct {
	Size *float32 `json:"size,omitempty"`
	ImageCount *float32 `json:"imageCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCache200ResponseImageCacheTmdb GetCache200ResponseImageCacheTmdb

// NewGetCache200ResponseImageCacheTmdb instantiates a new GetCache200ResponseImageCacheTmdb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCache200ResponseImageCacheTmdb() *GetCache200ResponseImageCacheTmdb {
	this := GetCache200ResponseImageCacheTmdb{}
	return &this
}

// NewGetCache200ResponseImageCacheTmdbWithDefaults instantiates a new GetCache200ResponseImageCacheTmdb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCache200ResponseImageCacheTmdbWithDefaults() *GetCache200ResponseImageCacheTmdb {
	this := GetCache200ResponseImageCacheTmdb{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetCache200ResponseImageCacheTmdb) GetSize() float32 {
	if o == nil || isNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache200ResponseImageCacheTmdb) GetSizeOk() (*float32, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetCache200ResponseImageCacheTmdb) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *GetCache200ResponseImageCacheTmdb) SetSize(v float32) {
	o.Size = &v
}

// GetImageCount returns the ImageCount field value if set, zero value otherwise.
func (o *GetCache200ResponseImageCacheTmdb) GetImageCount() float32 {
	if o == nil || isNil(o.ImageCount) {
		var ret float32
		return ret
	}
	return *o.ImageCount
}

// GetImageCountOk returns a tuple with the ImageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache200ResponseImageCacheTmdb) GetImageCountOk() (*float32, bool) {
	if o == nil || isNil(o.ImageCount) {
    return nil, false
	}
	return o.ImageCount, true
}

// HasImageCount returns a boolean if a field has been set.
func (o *GetCache200ResponseImageCacheTmdb) HasImageCount() bool {
	if o != nil && !isNil(o.ImageCount) {
		return true
	}

	return false
}

// SetImageCount gets a reference to the given float32 and assigns it to the ImageCount field.
func (o *GetCache200ResponseImageCacheTmdb) SetImageCount(v float32) {
	o.ImageCount = &v
}

func (o GetCache200ResponseImageCacheTmdb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.ImageCount) {
		toSerialize["imageCount"] = o.ImageCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetCache200ResponseImageCacheTmdb) UnmarshalJSON(bytes []byte) (err error) {
	varGetCache200ResponseImageCacheTmdb := _GetCache200ResponseImageCacheTmdb{}

	if err = json.Unmarshal(bytes, &varGetCache200ResponseImageCacheTmdb); err == nil {
		*o = GetCache200ResponseImageCacheTmdb(varGetCache200ResponseImageCacheTmdb)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "size")
		delete(additionalProperties, "imageCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCache200ResponseImageCacheTmdb struct {
	value *GetCache200ResponseImageCacheTmdb
	isSet bool
}

func (v NullableGetCache200ResponseImageCacheTmdb) Get() *GetCache200ResponseImageCacheTmdb {
	return v.value
}

func (v *NullableGetCache200ResponseImageCacheTmdb) Set(val *GetCache200ResponseImageCacheTmdb) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCache200ResponseImageCacheTmdb) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCache200ResponseImageCacheTmdb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCache200ResponseImageCacheTmdb(val *GetCache200ResponseImageCacheTmdb) *NullableGetCache200ResponseImageCacheTmdb {
	return &NullableGetCache200ResponseImageCacheTmdb{value: val, isSet: true}
}

func (v NullableGetCache200ResponseImageCacheTmdb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCache200ResponseImageCacheTmdb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


