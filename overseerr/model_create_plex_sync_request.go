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

// checks if the CreatePlexSyncRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePlexSyncRequest{}

// CreatePlexSyncRequest struct for CreatePlexSyncRequest
type CreatePlexSyncRequest struct {
	Cancel *bool `json:"cancel,omitempty"`
	Start *bool `json:"start,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePlexSyncRequest CreatePlexSyncRequest

// NewCreatePlexSyncRequest instantiates a new CreatePlexSyncRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlexSyncRequest() *CreatePlexSyncRequest {
	this := CreatePlexSyncRequest{}
	return &this
}

// NewCreatePlexSyncRequestWithDefaults instantiates a new CreatePlexSyncRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlexSyncRequestWithDefaults() *CreatePlexSyncRequest {
	this := CreatePlexSyncRequest{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *CreatePlexSyncRequest) GetCancel() bool {
	if o == nil || IsNil(o.Cancel) {
		var ret bool
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlexSyncRequest) GetCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *CreatePlexSyncRequest) HasCancel() bool {
	if o != nil && !IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given bool and assigns it to the Cancel field.
func (o *CreatePlexSyncRequest) SetCancel(v bool) {
	o.Cancel = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CreatePlexSyncRequest) GetStart() bool {
	if o == nil || IsNil(o.Start) {
		var ret bool
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlexSyncRequest) GetStartOk() (*bool, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CreatePlexSyncRequest) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given bool and assigns it to the Start field.
func (o *CreatePlexSyncRequest) SetStart(v bool) {
	o.Start = &v
}

func (o CreatePlexSyncRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePlexSyncRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePlexSyncRequest) UnmarshalJSON(data []byte) (err error) {
	varCreatePlexSyncRequest := _CreatePlexSyncRequest{}

	err = json.Unmarshal(data, &varCreatePlexSyncRequest)

	if err != nil {
		return err
	}

	*o = CreatePlexSyncRequest(varCreatePlexSyncRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cancel")
		delete(additionalProperties, "start")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePlexSyncRequest struct {
	value *CreatePlexSyncRequest
	isSet bool
}

func (v NullableCreatePlexSyncRequest) Get() *CreatePlexSyncRequest {
	return v.value
}

func (v *NullableCreatePlexSyncRequest) Set(val *CreatePlexSyncRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlexSyncRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlexSyncRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlexSyncRequest(val *CreatePlexSyncRequest) *NullableCreatePlexSyncRequest {
	return &NullableCreatePlexSyncRequest{value: val, isSet: true}
}

func (v NullableCreatePlexSyncRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePlexSyncRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


