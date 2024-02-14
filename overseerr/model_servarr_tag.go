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

// ServarrTag struct for ServarrTag
type ServarrTag struct {
	Id *float32 `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServarrTag ServarrTag

// NewServarrTag instantiates a new ServarrTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServarrTag() *ServarrTag {
	this := ServarrTag{}
	return &this
}

// NewServarrTagWithDefaults instantiates a new ServarrTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServarrTagWithDefaults() *ServarrTag {
	this := ServarrTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServarrTag) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServarrTag) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServarrTag) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ServarrTag) SetId(v float32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ServarrTag) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServarrTag) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ServarrTag) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ServarrTag) SetLabel(v string) {
	o.Label = &v
}

func (o ServarrTag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServarrTag) UnmarshalJSON(bytes []byte) (err error) {
	varServarrTag := _ServarrTag{}

	if err = json.Unmarshal(bytes, &varServarrTag); err == nil {
		*o = ServarrTag(varServarrTag)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServarrTag struct {
	value *ServarrTag
	isSet bool
}

func (v NullableServarrTag) Get() *ServarrTag {
	return v.value
}

func (v *NullableServarrTag) Set(val *ServarrTag) {
	v.value = val
	v.isSet = true
}

func (v NullableServarrTag) IsSet() bool {
	return v.isSet
}

func (v *NullableServarrTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServarrTag(val *ServarrTag) *NullableServarrTag {
	return &NullableServarrTag{value: val, isSet: true}
}

func (v NullableServarrTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServarrTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


