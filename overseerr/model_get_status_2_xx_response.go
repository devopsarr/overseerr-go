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

// GetStatus2XXResponse struct for GetStatus2XXResponse
type GetStatus2XXResponse struct {
	Version *string `json:"version,omitempty"`
	CommitTag *string `json:"commitTag,omitempty"`
	UpdateAvailable *bool `json:"updateAvailable,omitempty"`
	CommitsBehind *float32 `json:"commitsBehind,omitempty"`
	RestartRequired *bool `json:"restartRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetStatus2XXResponse GetStatus2XXResponse

// NewGetStatus2XXResponse instantiates a new GetStatus2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus2XXResponse() *GetStatus2XXResponse {
	this := GetStatus2XXResponse{}
	return &this
}

// NewGetStatus2XXResponseWithDefaults instantiates a new GetStatus2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus2XXResponseWithDefaults() *GetStatus2XXResponse {
	this := GetStatus2XXResponse{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetStatus2XXResponse) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus2XXResponse) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetStatus2XXResponse) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetStatus2XXResponse) SetVersion(v string) {
	o.Version = &v
}

// GetCommitTag returns the CommitTag field value if set, zero value otherwise.
func (o *GetStatus2XXResponse) GetCommitTag() string {
	if o == nil || isNil(o.CommitTag) {
		var ret string
		return ret
	}
	return *o.CommitTag
}

// GetCommitTagOk returns a tuple with the CommitTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus2XXResponse) GetCommitTagOk() (*string, bool) {
	if o == nil || isNil(o.CommitTag) {
    return nil, false
	}
	return o.CommitTag, true
}

// HasCommitTag returns a boolean if a field has been set.
func (o *GetStatus2XXResponse) HasCommitTag() bool {
	if o != nil && !isNil(o.CommitTag) {
		return true
	}

	return false
}

// SetCommitTag gets a reference to the given string and assigns it to the CommitTag field.
func (o *GetStatus2XXResponse) SetCommitTag(v string) {
	o.CommitTag = &v
}

// GetUpdateAvailable returns the UpdateAvailable field value if set, zero value otherwise.
func (o *GetStatus2XXResponse) GetUpdateAvailable() bool {
	if o == nil || isNil(o.UpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.UpdateAvailable
}

// GetUpdateAvailableOk returns a tuple with the UpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus2XXResponse) GetUpdateAvailableOk() (*bool, bool) {
	if o == nil || isNil(o.UpdateAvailable) {
    return nil, false
	}
	return o.UpdateAvailable, true
}

// HasUpdateAvailable returns a boolean if a field has been set.
func (o *GetStatus2XXResponse) HasUpdateAvailable() bool {
	if o != nil && !isNil(o.UpdateAvailable) {
		return true
	}

	return false
}

// SetUpdateAvailable gets a reference to the given bool and assigns it to the UpdateAvailable field.
func (o *GetStatus2XXResponse) SetUpdateAvailable(v bool) {
	o.UpdateAvailable = &v
}

// GetCommitsBehind returns the CommitsBehind field value if set, zero value otherwise.
func (o *GetStatus2XXResponse) GetCommitsBehind() float32 {
	if o == nil || isNil(o.CommitsBehind) {
		var ret float32
		return ret
	}
	return *o.CommitsBehind
}

// GetCommitsBehindOk returns a tuple with the CommitsBehind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus2XXResponse) GetCommitsBehindOk() (*float32, bool) {
	if o == nil || isNil(o.CommitsBehind) {
    return nil, false
	}
	return o.CommitsBehind, true
}

// HasCommitsBehind returns a boolean if a field has been set.
func (o *GetStatus2XXResponse) HasCommitsBehind() bool {
	if o != nil && !isNil(o.CommitsBehind) {
		return true
	}

	return false
}

// SetCommitsBehind gets a reference to the given float32 and assigns it to the CommitsBehind field.
func (o *GetStatus2XXResponse) SetCommitsBehind(v float32) {
	o.CommitsBehind = &v
}

// GetRestartRequired returns the RestartRequired field value if set, zero value otherwise.
func (o *GetStatus2XXResponse) GetRestartRequired() bool {
	if o == nil || isNil(o.RestartRequired) {
		var ret bool
		return ret
	}
	return *o.RestartRequired
}

// GetRestartRequiredOk returns a tuple with the RestartRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatus2XXResponse) GetRestartRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.RestartRequired) {
    return nil, false
	}
	return o.RestartRequired, true
}

// HasRestartRequired returns a boolean if a field has been set.
func (o *GetStatus2XXResponse) HasRestartRequired() bool {
	if o != nil && !isNil(o.RestartRequired) {
		return true
	}

	return false
}

// SetRestartRequired gets a reference to the given bool and assigns it to the RestartRequired field.
func (o *GetStatus2XXResponse) SetRestartRequired(v bool) {
	o.RestartRequired = &v
}

func (o GetStatus2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.CommitTag) {
		toSerialize["commitTag"] = o.CommitTag
	}
	if !isNil(o.UpdateAvailable) {
		toSerialize["updateAvailable"] = o.UpdateAvailable
	}
	if !isNil(o.CommitsBehind) {
		toSerialize["commitsBehind"] = o.CommitsBehind
	}
	if !isNil(o.RestartRequired) {
		toSerialize["restartRequired"] = o.RestartRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GetStatus2XXResponse) UnmarshalJSON(bytes []byte) (err error) {
	varGetStatus2XXResponse := _GetStatus2XXResponse{}

	if err = json.Unmarshal(bytes, &varGetStatus2XXResponse); err == nil {
		*o = GetStatus2XXResponse(varGetStatus2XXResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "version")
		delete(additionalProperties, "commitTag")
		delete(additionalProperties, "updateAvailable")
		delete(additionalProperties, "commitsBehind")
		delete(additionalProperties, "restartRequired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetStatus2XXResponse struct {
	value *GetStatus2XXResponse
	isSet bool
}

func (v NullableGetStatus2XXResponse) Get() *GetStatus2XXResponse {
	return v.value
}

func (v *NullableGetStatus2XXResponse) Set(val *GetStatus2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus2XXResponse(val *GetStatus2XXResponse) *NullableGetStatus2XXResponse {
	return &NullableGetStatus2XXResponse{value: val, isSet: true}
}

func (v NullableGetStatus2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

