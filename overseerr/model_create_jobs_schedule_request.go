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

// CreateJobsScheduleRequest struct for CreateJobsScheduleRequest
type CreateJobsScheduleRequest struct {
	Schedule *string `json:"schedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateJobsScheduleRequest CreateJobsScheduleRequest

// NewCreateJobsScheduleRequest instantiates a new CreateJobsScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateJobsScheduleRequest() *CreateJobsScheduleRequest {
	this := CreateJobsScheduleRequest{}
	return &this
}

// NewCreateJobsScheduleRequestWithDefaults instantiates a new CreateJobsScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateJobsScheduleRequestWithDefaults() *CreateJobsScheduleRequest {
	this := CreateJobsScheduleRequest{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CreateJobsScheduleRequest) GetSchedule() string {
	if o == nil || isNil(o.Schedule) {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateJobsScheduleRequest) GetScheduleOk() (*string, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CreateJobsScheduleRequest) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *CreateJobsScheduleRequest) SetSchedule(v string) {
	o.Schedule = &v
}

func (o CreateJobsScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateJobsScheduleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateJobsScheduleRequest := _CreateJobsScheduleRequest{}

	if err = json.Unmarshal(bytes, &varCreateJobsScheduleRequest); err == nil {
		*o = CreateJobsScheduleRequest(varCreateJobsScheduleRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateJobsScheduleRequest struct {
	value *CreateJobsScheduleRequest
	isSet bool
}

func (v NullableCreateJobsScheduleRequest) Get() *CreateJobsScheduleRequest {
	return v.value
}

func (v *NullableCreateJobsScheduleRequest) Set(val *CreateJobsScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateJobsScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateJobsScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateJobsScheduleRequest(val *CreateJobsScheduleRequest) *NullableCreateJobsScheduleRequest {
	return &NullableCreateJobsScheduleRequest{value: val, isSet: true}
}

func (v NullableCreateJobsScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateJobsScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


