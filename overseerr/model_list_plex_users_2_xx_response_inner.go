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

// checks if the ListPlexUsers2XXResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPlexUsers2XXResponseInner{}

// ListPlexUsers2XXResponseInner struct for ListPlexUsers2XXResponseInner
type ListPlexUsers2XXResponseInner struct {
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	Thumb *string `json:"thumb,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListPlexUsers2XXResponseInner ListPlexUsers2XXResponseInner

// NewListPlexUsers2XXResponseInner instantiates a new ListPlexUsers2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPlexUsers2XXResponseInner() *ListPlexUsers2XXResponseInner {
	this := ListPlexUsers2XXResponseInner{}
	return &this
}

// NewListPlexUsers2XXResponseInnerWithDefaults instantiates a new ListPlexUsers2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPlexUsers2XXResponseInnerWithDefaults() *ListPlexUsers2XXResponseInner {
	this := ListPlexUsers2XXResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListPlexUsers2XXResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlexUsers2XXResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListPlexUsers2XXResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListPlexUsers2XXResponseInner) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ListPlexUsers2XXResponseInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlexUsers2XXResponseInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ListPlexUsers2XXResponseInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ListPlexUsers2XXResponseInner) SetTitle(v string) {
	o.Title = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListPlexUsers2XXResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlexUsers2XXResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ListPlexUsers2XXResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListPlexUsers2XXResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ListPlexUsers2XXResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlexUsers2XXResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ListPlexUsers2XXResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ListPlexUsers2XXResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetThumb returns the Thumb field value if set, zero value otherwise.
func (o *ListPlexUsers2XXResponseInner) GetThumb() string {
	if o == nil || IsNil(o.Thumb) {
		var ret string
		return ret
	}
	return *o.Thumb
}

// GetThumbOk returns a tuple with the Thumb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPlexUsers2XXResponseInner) GetThumbOk() (*string, bool) {
	if o == nil || IsNil(o.Thumb) {
		return nil, false
	}
	return o.Thumb, true
}

// HasThumb returns a boolean if a field has been set.
func (o *ListPlexUsers2XXResponseInner) HasThumb() bool {
	if o != nil && !IsNil(o.Thumb) {
		return true
	}

	return false
}

// SetThumb gets a reference to the given string and assigns it to the Thumb field.
func (o *ListPlexUsers2XXResponseInner) SetThumb(v string) {
	o.Thumb = &v
}

func (o ListPlexUsers2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPlexUsers2XXResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Thumb) {
		toSerialize["thumb"] = o.Thumb
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListPlexUsers2XXResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListPlexUsers2XXResponseInner := _ListPlexUsers2XXResponseInner{}

	err = json.Unmarshal(data, &varListPlexUsers2XXResponseInner)

	if err != nil {
		return err
	}

	*o = ListPlexUsers2XXResponseInner(varListPlexUsers2XXResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "username")
		delete(additionalProperties, "email")
		delete(additionalProperties, "thumb")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListPlexUsers2XXResponseInner struct {
	value *ListPlexUsers2XXResponseInner
	isSet bool
}

func (v NullableListPlexUsers2XXResponseInner) Get() *ListPlexUsers2XXResponseInner {
	return v.value
}

func (v *NullableListPlexUsers2XXResponseInner) Set(val *ListPlexUsers2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListPlexUsers2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListPlexUsers2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPlexUsers2XXResponseInner(val *ListPlexUsers2XXResponseInner) *NullableListPlexUsers2XXResponseInner {
	return &NullableListPlexUsers2XXResponseInner{value: val, isSet: true}
}

func (v NullableListPlexUsers2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPlexUsers2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


