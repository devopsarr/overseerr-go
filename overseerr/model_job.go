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

// checks if the Job type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Job{}

// Job struct for Job
type Job struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Interval *string `json:"interval,omitempty"`
	Name *string `json:"name,omitempty"`
	NextExecutionTime *string `json:"nextExecutionTime,omitempty"`
	Running *bool `json:"running,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Job Job

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob() *Job {
	this := Job{}
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Job) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Job) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Job) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Job) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Job) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Job) SetType(v string) {
	o.Type = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *Job) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *Job) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *Job) SetInterval(v string) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Job) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Job) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Job) SetName(v string) {
	o.Name = &v
}

// GetNextExecutionTime returns the NextExecutionTime field value if set, zero value otherwise.
func (o *Job) GetNextExecutionTime() string {
	if o == nil || IsNil(o.NextExecutionTime) {
		var ret string
		return ret
	}
	return *o.NextExecutionTime
}

// GetNextExecutionTimeOk returns a tuple with the NextExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetNextExecutionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextExecutionTime) {
		return nil, false
	}
	return o.NextExecutionTime, true
}

// HasNextExecutionTime returns a boolean if a field has been set.
func (o *Job) HasNextExecutionTime() bool {
	if o != nil && !IsNil(o.NextExecutionTime) {
		return true
	}

	return false
}

// SetNextExecutionTime gets a reference to the given string and assigns it to the NextExecutionTime field.
func (o *Job) SetNextExecutionTime(v string) {
	o.NextExecutionTime = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *Job) GetRunning() bool {
	if o == nil || IsNil(o.Running) {
		var ret bool
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.Running) {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *Job) HasRunning() bool {
	if o != nil && !IsNil(o.Running) {
		return true
	}

	return false
}

// SetRunning gets a reference to the given bool and assigns it to the Running field.
func (o *Job) SetRunning(v bool) {
	o.Running = &v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Job) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NextExecutionTime) {
		toSerialize["nextExecutionTime"] = o.NextExecutionTime
	}
	if !IsNil(o.Running) {
		toSerialize["running"] = o.Running
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Job) UnmarshalJSON(data []byte) (err error) {
	varJob := _Job{}

	err = json.Unmarshal(data, &varJob)

	if err != nil {
		return err
	}

	*o = Job(varJob)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nextExecutionTime")
		delete(additionalProperties, "running")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


