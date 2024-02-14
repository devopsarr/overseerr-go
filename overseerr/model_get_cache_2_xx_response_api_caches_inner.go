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

// GetCache2XXResponseApiCachesInner struct for GetCache2XXResponseApiCachesInner
type GetCache2XXResponseApiCachesInner struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Stats *GetCache2XXResponseApiCachesInnerStats `json:"stats,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetCache2XXResponseApiCachesInner GetCache2XXResponseApiCachesInner

// NewGetCache2XXResponseApiCachesInner instantiates a new GetCache2XXResponseApiCachesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCache2XXResponseApiCachesInner() *GetCache2XXResponseApiCachesInner {
	this := GetCache2XXResponseApiCachesInner{}
	return &this
}

// NewGetCache2XXResponseApiCachesInnerWithDefaults instantiates a new GetCache2XXResponseApiCachesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCache2XXResponseApiCachesInnerWithDefaults() *GetCache2XXResponseApiCachesInner {
	this := GetCache2XXResponseApiCachesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetCache2XXResponseApiCachesInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache2XXResponseApiCachesInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetCache2XXResponseApiCachesInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetCache2XXResponseApiCachesInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCache2XXResponseApiCachesInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache2XXResponseApiCachesInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCache2XXResponseApiCachesInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCache2XXResponseApiCachesInner) SetName(v string) {
	o.Name = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *GetCache2XXResponseApiCachesInner) GetStats() GetCache2XXResponseApiCachesInnerStats {
	if o == nil || isNil(o.Stats) {
		var ret GetCache2XXResponseApiCachesInnerStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCache2XXResponseApiCachesInner) GetStatsOk() (*GetCache2XXResponseApiCachesInnerStats, bool) {
	if o == nil || isNil(o.Stats) {
    return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *GetCache2XXResponseApiCachesInner) HasStats() bool {
	if o != nil && !isNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given GetCache2XXResponseApiCachesInnerStats and assigns it to the Stats field.
func (o *GetCache2XXResponseApiCachesInner) SetStats(v GetCache2XXResponseApiCachesInnerStats) {
	o.Stats = &v
}

func (o GetCache2XXResponseApiCachesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetCache2XXResponseApiCachesInner) UnmarshalJSON(bytes []byte) (err error) {
	varGetCache2XXResponseApiCachesInner := _GetCache2XXResponseApiCachesInner{}

	if err = json.Unmarshal(bytes, &varGetCache2XXResponseApiCachesInner); err == nil {
		*o = GetCache2XXResponseApiCachesInner(varGetCache2XXResponseApiCachesInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCache2XXResponseApiCachesInner struct {
	value *GetCache2XXResponseApiCachesInner
	isSet bool
}

func (v NullableGetCache2XXResponseApiCachesInner) Get() *GetCache2XXResponseApiCachesInner {
	return v.value
}

func (v *NullableGetCache2XXResponseApiCachesInner) Set(val *GetCache2XXResponseApiCachesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCache2XXResponseApiCachesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCache2XXResponseApiCachesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCache2XXResponseApiCachesInner(val *GetCache2XXResponseApiCachesInner) *NullableGetCache2XXResponseApiCachesInner {
	return &NullableGetCache2XXResponseApiCachesInner{value: val, isSet: true}
}

func (v NullableGetCache2XXResponseApiCachesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCache2XXResponseApiCachesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


