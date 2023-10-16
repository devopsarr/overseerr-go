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

// ListGenresMovie200ResponseInner struct for ListGenresMovie200ResponseInner
type ListGenresMovie200ResponseInner struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListGenresMovie200ResponseInner ListGenresMovie200ResponseInner

// NewListGenresMovie200ResponseInner instantiates a new ListGenresMovie200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGenresMovie200ResponseInner() *ListGenresMovie200ResponseInner {
	this := ListGenresMovie200ResponseInner{}
	return &this
}

// NewListGenresMovie200ResponseInnerWithDefaults instantiates a new ListGenresMovie200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGenresMovie200ResponseInnerWithDefaults() *ListGenresMovie200ResponseInner {
	this := ListGenresMovie200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListGenresMovie200ResponseInner) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGenresMovie200ResponseInner) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListGenresMovie200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ListGenresMovie200ResponseInner) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListGenresMovie200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGenresMovie200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListGenresMovie200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListGenresMovie200ResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListGenresMovie200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListGenresMovie200ResponseInner) UnmarshalJSON(bytes []byte) (err error) {
	varListGenresMovie200ResponseInner := _ListGenresMovie200ResponseInner{}

	if err = json.Unmarshal(bytes, &varListGenresMovie200ResponseInner); err == nil {
		*o = ListGenresMovie200ResponseInner(varListGenresMovie200ResponseInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListGenresMovie200ResponseInner struct {
	value *ListGenresMovie200ResponseInner
	isSet bool
}

func (v NullableListGenresMovie200ResponseInner) Get() *ListGenresMovie200ResponseInner {
	return v.value
}

func (v *NullableListGenresMovie200ResponseInner) Set(val *ListGenresMovie200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListGenresMovie200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListGenresMovie200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGenresMovie200ResponseInner(val *ListGenresMovie200ResponseInner) *NullableListGenresMovie200ResponseInner {
	return &NullableListGenresMovie200ResponseInner{value: val, isSet: true}
}

func (v NullableListGenresMovie200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGenresMovie200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

