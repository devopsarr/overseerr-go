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

// checks if the UpdateIssueCommentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIssueCommentRequest{}

// UpdateIssueCommentRequest struct for UpdateIssueCommentRequest
type UpdateIssueCommentRequest struct {
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIssueCommentRequest UpdateIssueCommentRequest

// NewUpdateIssueCommentRequest instantiates a new UpdateIssueCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIssueCommentRequest() *UpdateIssueCommentRequest {
	this := UpdateIssueCommentRequest{}
	return &this
}

// NewUpdateIssueCommentRequestWithDefaults instantiates a new UpdateIssueCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIssueCommentRequestWithDefaults() *UpdateIssueCommentRequest {
	this := UpdateIssueCommentRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UpdateIssueCommentRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueCommentRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UpdateIssueCommentRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UpdateIssueCommentRequest) SetMessage(v string) {
	o.Message = &v
}

func (o UpdateIssueCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIssueCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIssueCommentRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateIssueCommentRequest := _UpdateIssueCommentRequest{}

	err = json.Unmarshal(data, &varUpdateIssueCommentRequest)

	if err != nil {
		return err
	}

	*o = UpdateIssueCommentRequest(varUpdateIssueCommentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIssueCommentRequest struct {
	value *UpdateIssueCommentRequest
	isSet bool
}

func (v NullableUpdateIssueCommentRequest) Get() *UpdateIssueCommentRequest {
	return v.value
}

func (v *NullableUpdateIssueCommentRequest) Set(val *UpdateIssueCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIssueCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIssueCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIssueCommentRequest(val *UpdateIssueCommentRequest) *NullableUpdateIssueCommentRequest {
	return &NullableUpdateIssueCommentRequest{value: val, isSet: true}
}

func (v NullableUpdateIssueCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIssueCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


