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

// checks if the CreateUserImportFromPlexRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserImportFromPlexRequest{}

// CreateUserImportFromPlexRequest struct for CreateUserImportFromPlexRequest
type CreateUserImportFromPlexRequest struct {
	PlexIds []string `json:"plexIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserImportFromPlexRequest CreateUserImportFromPlexRequest

// NewCreateUserImportFromPlexRequest instantiates a new CreateUserImportFromPlexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserImportFromPlexRequest() *CreateUserImportFromPlexRequest {
	this := CreateUserImportFromPlexRequest{}
	return &this
}

// NewCreateUserImportFromPlexRequestWithDefaults instantiates a new CreateUserImportFromPlexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserImportFromPlexRequestWithDefaults() *CreateUserImportFromPlexRequest {
	this := CreateUserImportFromPlexRequest{}
	return &this
}

// GetPlexIds returns the PlexIds field value if set, zero value otherwise.
func (o *CreateUserImportFromPlexRequest) GetPlexIds() []string {
	if o == nil || IsNil(o.PlexIds) {
		var ret []string
		return ret
	}
	return o.PlexIds
}

// GetPlexIdsOk returns a tuple with the PlexIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserImportFromPlexRequest) GetPlexIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlexIds) {
		return nil, false
	}
	return o.PlexIds, true
}

// HasPlexIds returns a boolean if a field has been set.
func (o *CreateUserImportFromPlexRequest) HasPlexIds() bool {
	if o != nil && !IsNil(o.PlexIds) {
		return true
	}

	return false
}

// SetPlexIds gets a reference to the given []string and assigns it to the PlexIds field.
func (o *CreateUserImportFromPlexRequest) SetPlexIds(v []string) {
	o.PlexIds = v
}

func (o CreateUserImportFromPlexRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserImportFromPlexRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlexIds) {
		toSerialize["plexIds"] = o.PlexIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateUserImportFromPlexRequest) UnmarshalJSON(data []byte) (err error) {
	varCreateUserImportFromPlexRequest := _CreateUserImportFromPlexRequest{}

	err = json.Unmarshal(data, &varCreateUserImportFromPlexRequest)

	if err != nil {
		return err
	}

	*o = CreateUserImportFromPlexRequest(varCreateUserImportFromPlexRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "plexIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserImportFromPlexRequest struct {
	value *CreateUserImportFromPlexRequest
	isSet bool
}

func (v NullableCreateUserImportFromPlexRequest) Get() *CreateUserImportFromPlexRequest {
	return v.value
}

func (v *NullableCreateUserImportFromPlexRequest) Set(val *CreateUserImportFromPlexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserImportFromPlexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserImportFromPlexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserImportFromPlexRequest(val *CreateUserImportFromPlexRequest) *NullableCreateUserImportFromPlexRequest {
	return &NullableCreateUserImportFromPlexRequest{value: val, isSet: true}
}

func (v NullableCreateUserImportFromPlexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserImportFromPlexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


