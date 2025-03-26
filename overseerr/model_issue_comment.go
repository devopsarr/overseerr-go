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

// checks if the IssueComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueComment{}

// IssueComment struct for IssueComment
type IssueComment struct {
	Id *float32 `json:"id,omitempty"`
	User *User `json:"user,omitempty"`
	Message *string `json:"message,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssueComment IssueComment

// NewIssueComment instantiates a new IssueComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueComment() *IssueComment {
	this := IssueComment{}
	return &this
}

// NewIssueCommentWithDefaults instantiates a new IssueComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueCommentWithDefaults() *IssueComment {
	this := IssueComment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssueComment) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssueComment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *IssueComment) SetId(v float32) {
	o.Id = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IssueComment) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IssueComment) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *IssueComment) SetUser(v User) {
	o.User = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IssueComment) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IssueComment) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IssueComment) SetMessage(v string) {
	o.Message = &v
}

func (o IssueComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IssueComment) UnmarshalJSON(data []byte) (err error) {
	varIssueComment := _IssueComment{}

	err = json.Unmarshal(data, &varIssueComment)

	if err != nil {
		return err
	}

	*o = IssueComment(varIssueComment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user")
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIssueComment struct {
	value *IssueComment
	isSet bool
}

func (v NullableIssueComment) Get() *IssueComment {
	return v.value
}

func (v *NullableIssueComment) Set(val *IssueComment) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueComment) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueComment(val *IssueComment) *NullableIssueComment {
	return &NullableIssueComment{value: val, isSet: true}
}

func (v NullableIssueComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


