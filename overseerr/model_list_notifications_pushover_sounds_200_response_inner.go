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

// ListNotificationsPushoverSounds200ResponseInner struct for ListNotificationsPushoverSounds200ResponseInner
type ListNotificationsPushoverSounds200ResponseInner struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListNotificationsPushoverSounds200ResponseInner ListNotificationsPushoverSounds200ResponseInner

// NewListNotificationsPushoverSounds200ResponseInner instantiates a new ListNotificationsPushoverSounds200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNotificationsPushoverSounds200ResponseInner() *ListNotificationsPushoverSounds200ResponseInner {
	this := ListNotificationsPushoverSounds200ResponseInner{}
	return &this
}

// NewListNotificationsPushoverSounds200ResponseInnerWithDefaults instantiates a new ListNotificationsPushoverSounds200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNotificationsPushoverSounds200ResponseInnerWithDefaults() *ListNotificationsPushoverSounds200ResponseInner {
	this := ListNotificationsPushoverSounds200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListNotificationsPushoverSounds200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNotificationsPushoverSounds200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListNotificationsPushoverSounds200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListNotificationsPushoverSounds200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListNotificationsPushoverSounds200ResponseInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNotificationsPushoverSounds200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListNotificationsPushoverSounds200ResponseInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListNotificationsPushoverSounds200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

func (o ListNotificationsPushoverSounds200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListNotificationsPushoverSounds200ResponseInner) UnmarshalJSON(bytes []byte) (err error) {
	varListNotificationsPushoverSounds200ResponseInner := _ListNotificationsPushoverSounds200ResponseInner{}

	if err = json.Unmarshal(bytes, &varListNotificationsPushoverSounds200ResponseInner); err == nil {
		*o = ListNotificationsPushoverSounds200ResponseInner(varListNotificationsPushoverSounds200ResponseInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListNotificationsPushoverSounds200ResponseInner struct {
	value *ListNotificationsPushoverSounds200ResponseInner
	isSet bool
}

func (v NullableListNotificationsPushoverSounds200ResponseInner) Get() *ListNotificationsPushoverSounds200ResponseInner {
	return v.value
}

func (v *NullableListNotificationsPushoverSounds200ResponseInner) Set(val *ListNotificationsPushoverSounds200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListNotificationsPushoverSounds200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListNotificationsPushoverSounds200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNotificationsPushoverSounds200ResponseInner(val *ListNotificationsPushoverSounds200ResponseInner) *NullableListNotificationsPushoverSounds200ResponseInner {
	return &NullableListNotificationsPushoverSounds200ResponseInner{value: val, isSet: true}
}

func (v NullableListNotificationsPushoverSounds200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNotificationsPushoverSounds200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


