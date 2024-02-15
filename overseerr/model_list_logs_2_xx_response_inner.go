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

// checks if the ListLogs2XXResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLogs2XXResponseInner{}

// ListLogs2XXResponseInner struct for ListLogs2XXResponseInner
type ListLogs2XXResponseInner struct {
	Label *string `json:"label,omitempty"`
	Level *string `json:"level,omitempty"`
	Message *string `json:"message,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListLogs2XXResponseInner ListLogs2XXResponseInner

// NewListLogs2XXResponseInner instantiates a new ListLogs2XXResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLogs2XXResponseInner() *ListLogs2XXResponseInner {
	this := ListLogs2XXResponseInner{}
	return &this
}

// NewListLogs2XXResponseInnerWithDefaults instantiates a new ListLogs2XXResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLogs2XXResponseInnerWithDefaults() *ListLogs2XXResponseInner {
	this := ListLogs2XXResponseInner{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ListLogs2XXResponseInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs2XXResponseInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ListLogs2XXResponseInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ListLogs2XXResponseInner) SetLabel(v string) {
	o.Label = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *ListLogs2XXResponseInner) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs2XXResponseInner) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *ListLogs2XXResponseInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *ListLogs2XXResponseInner) SetLevel(v string) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ListLogs2XXResponseInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs2XXResponseInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ListLogs2XXResponseInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ListLogs2XXResponseInner) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ListLogs2XXResponseInner) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogs2XXResponseInner) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ListLogs2XXResponseInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ListLogs2XXResponseInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o ListLogs2XXResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLogs2XXResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListLogs2XXResponseInner) UnmarshalJSON(data []byte) (err error) {
	varListLogs2XXResponseInner := _ListLogs2XXResponseInner{}

	err = json.Unmarshal(data, &varListLogs2XXResponseInner)

	if err != nil {
		return err
	}

	*o = ListLogs2XXResponseInner(varListLogs2XXResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "level")
		delete(additionalProperties, "message")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListLogs2XXResponseInner struct {
	value *ListLogs2XXResponseInner
	isSet bool
}

func (v NullableListLogs2XXResponseInner) Get() *ListLogs2XXResponseInner {
	return v.value
}

func (v *NullableListLogs2XXResponseInner) Set(val *ListLogs2XXResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListLogs2XXResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListLogs2XXResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLogs2XXResponseInner(val *ListLogs2XXResponseInner) *NullableListLogs2XXResponseInner {
	return &NullableListLogs2XXResponseInner{value: val, isSet: true}
}

func (v NullableListLogs2XXResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLogs2XXResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


