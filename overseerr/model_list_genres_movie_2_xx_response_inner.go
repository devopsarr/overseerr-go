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

// checks if the ListGenresMovie2XXResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGenresMovie2XXResponseInner{}

// ListGenresMovie2XXResponseInner struct for ListGenresMovie2XXResponseInner
type ListGenresMovie2XXResponseInner struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListGenresMovie2XXResponseInner ListGenresMovie2XXResponseInner

// NewListGenresMovie2XXResponseInner instantiates a new ListGenresMovie2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGenresMovie2XXResponseInner() *ListGenresMovie2XXResponseInner {
	this := ListGenresMovie2XXResponseInner{}
	return &this
}

// NewListGenresMovie2XXResponseInnerWithDefaults instantiates a new ListGenresMovie2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGenresMovie2XXResponseInnerWithDefaults() *ListGenresMovie2XXResponseInner {
	this := ListGenresMovie2XXResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListGenresMovie2XXResponseInner) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGenresMovie2XXResponseInner) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListGenresMovie2XXResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ListGenresMovie2XXResponseInner) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListGenresMovie2XXResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGenresMovie2XXResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListGenresMovie2XXResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListGenresMovie2XXResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListGenresMovie2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGenresMovie2XXResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListGenresMovie2XXResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListGenresMovie2XXResponseInner := _ListGenresMovie2XXResponseInner{}

	err = json.Unmarshal(data, &varListGenresMovie2XXResponseInner)

	if err != nil {
		return err
	}

	*o = ListGenresMovie2XXResponseInner(varListGenresMovie2XXResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListGenresMovie2XXResponseInner struct {
	value *ListGenresMovie2XXResponseInner
	isSet bool
}

func (v NullableListGenresMovie2XXResponseInner) Get() *ListGenresMovie2XXResponseInner {
	return v.value
}

func (v *NullableListGenresMovie2XXResponseInner) Set(val *ListGenresMovie2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListGenresMovie2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListGenresMovie2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGenresMovie2XXResponseInner(val *ListGenresMovie2XXResponseInner) *NullableListGenresMovie2XXResponseInner {
	return &NullableListGenresMovie2XXResponseInner{value: val, isSet: true}
}

func (v NullableListGenresMovie2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGenresMovie2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


