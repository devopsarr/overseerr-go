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

// checks if the ListDiscoverGenresliderMovie2XXResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDiscoverGenresliderMovie2XXResponseInner{}

// ListDiscoverGenresliderMovie2XXResponseInner struct for ListDiscoverGenresliderMovie2XXResponseInner
type ListDiscoverGenresliderMovie2XXResponseInner struct {
	Id *float32 `json:"id,omitempty"`
	Backdrops []string `json:"backdrops,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDiscoverGenresliderMovie2XXResponseInner ListDiscoverGenresliderMovie2XXResponseInner

// NewListDiscoverGenresliderMovie2XXResponseInner instantiates a new ListDiscoverGenresliderMovie2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDiscoverGenresliderMovie2XXResponseInner() *ListDiscoverGenresliderMovie2XXResponseInner {
	this := ListDiscoverGenresliderMovie2XXResponseInner{}
	return &this
}

// NewListDiscoverGenresliderMovie2XXResponseInnerWithDefaults instantiates a new ListDiscoverGenresliderMovie2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDiscoverGenresliderMovie2XXResponseInnerWithDefaults() *ListDiscoverGenresliderMovie2XXResponseInner {
	this := ListDiscoverGenresliderMovie2XXResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) SetId(v float32) {
	o.Id = &v
}

// GetBackdrops returns the Backdrops field value if set, zero value otherwise.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetBackdrops() []string {
	if o == nil || IsNil(o.Backdrops) {
		var ret []string
		return ret
	}
	return o.Backdrops
}

// GetBackdropsOk returns a tuple with the Backdrops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetBackdropsOk() ([]string, bool) {
	if o == nil || IsNil(o.Backdrops) {
		return nil, false
	}
	return o.Backdrops, true
}

// HasBackdrops returns a boolean if a field has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) HasBackdrops() bool {
	if o != nil && !IsNil(o.Backdrops) {
		return true
	}

	return false
}

// SetBackdrops gets a reference to the given []string and assigns it to the Backdrops field.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) SetBackdrops(v []string) {
	o.Backdrops = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListDiscoverGenresliderMovie2XXResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListDiscoverGenresliderMovie2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDiscoverGenresliderMovie2XXResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Backdrops) {
		toSerialize["backdrops"] = o.Backdrops
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDiscoverGenresliderMovie2XXResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListDiscoverGenresliderMovie2XXResponseInner := _ListDiscoverGenresliderMovie2XXResponseInner{}

	err = json.Unmarshal(data, &varListDiscoverGenresliderMovie2XXResponseInner)

	if err != nil {
		return err
	}

	*o = ListDiscoverGenresliderMovie2XXResponseInner(varListDiscoverGenresliderMovie2XXResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "backdrops")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDiscoverGenresliderMovie2XXResponseInner struct {
	value *ListDiscoverGenresliderMovie2XXResponseInner
	isSet bool
}

func (v NullableListDiscoverGenresliderMovie2XXResponseInner) Get() *ListDiscoverGenresliderMovie2XXResponseInner {
	return v.value
}

func (v *NullableListDiscoverGenresliderMovie2XXResponseInner) Set(val *ListDiscoverGenresliderMovie2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListDiscoverGenresliderMovie2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListDiscoverGenresliderMovie2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDiscoverGenresliderMovie2XXResponseInner(val *ListDiscoverGenresliderMovie2XXResponseInner) *NullableListDiscoverGenresliderMovie2XXResponseInner {
	return &NullableListDiscoverGenresliderMovie2XXResponseInner{value: val, isSet: true}
}

func (v NullableListDiscoverGenresliderMovie2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDiscoverGenresliderMovie2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


