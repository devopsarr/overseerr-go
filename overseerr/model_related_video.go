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

// checks if the RelatedVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedVideo{}

// RelatedVideo struct for RelatedVideo
type RelatedVideo struct {
	Url *string `json:"url,omitempty"`
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *float32 `json:"size,omitempty"`
	Type *string `json:"type,omitempty"`
	Site *string `json:"site,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RelatedVideo RelatedVideo

// NewRelatedVideo instantiates a new RelatedVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedVideo() *RelatedVideo {
	this := RelatedVideo{}
	return &this
}

// NewRelatedVideoWithDefaults instantiates a new RelatedVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedVideoWithDefaults() *RelatedVideo {
	this := RelatedVideo{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RelatedVideo) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RelatedVideo) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RelatedVideo) SetUrl(v string) {
	o.Url = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RelatedVideo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RelatedVideo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RelatedVideo) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RelatedVideo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RelatedVideo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RelatedVideo) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *RelatedVideo) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *RelatedVideo) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *RelatedVideo) SetSize(v float32) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelatedVideo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelatedVideo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RelatedVideo) SetType(v string) {
	o.Type = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *RelatedVideo) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedVideo) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *RelatedVideo) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *RelatedVideo) SetSite(v string) {
	o.Site = &v
}

func (o RelatedVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RelatedVideo) UnmarshalJSON(data []byte) (err error) {
	varRelatedVideo := _RelatedVideo{}

	err = json.Unmarshal(data, &varRelatedVideo)

	if err != nil {
		return err
	}

	*o = RelatedVideo(varRelatedVideo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "key")
		delete(additionalProperties, "name")
		delete(additionalProperties, "size")
		delete(additionalProperties, "type")
		delete(additionalProperties, "site")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelatedVideo struct {
	value *RelatedVideo
	isSet bool
}

func (v NullableRelatedVideo) Get() *RelatedVideo {
	return v.value
}

func (v *NullableRelatedVideo) Set(val *RelatedVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedVideo(val *RelatedVideo) *NullableRelatedVideo {
	return &NullableRelatedVideo{value: val, isSet: true}
}

func (v NullableRelatedVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


