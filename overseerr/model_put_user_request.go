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

// checks if the PutUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutUserRequest{}

// PutUserRequest struct for PutUserRequest
type PutUserRequest struct {
	Ids []int32 `json:"ids,omitempty"`
	Permissions *int32 `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PutUserRequest PutUserRequest

// NewPutUserRequest instantiates a new PutUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutUserRequest() *PutUserRequest {
	this := PutUserRequest{}
	return &this
}

// NewPutUserRequestWithDefaults instantiates a new PutUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutUserRequestWithDefaults() *PutUserRequest {
	this := PutUserRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *PutUserRequest) GetIds() []int32 {
	if o == nil || IsNil(o.Ids) {
		var ret []int32
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutUserRequest) GetIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *PutUserRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int32 and assigns it to the Ids field.
func (o *PutUserRequest) SetIds(v []int32) {
	o.Ids = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PutUserRequest) GetPermissions() int32 {
	if o == nil || IsNil(o.Permissions) {
		var ret int32
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutUserRequest) GetPermissionsOk() (*int32, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PutUserRequest) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given int32 and assigns it to the Permissions field.
func (o *PutUserRequest) SetPermissions(v int32) {
	o.Permissions = &v
}

func (o PutUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PutUserRequest) UnmarshalJSON(data []byte) (err error) {
	varPutUserRequest := _PutUserRequest{}

	err = json.Unmarshal(data, &varPutUserRequest)

	if err != nil {
		return err
	}

	*o = PutUserRequest(varPutUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePutUserRequest struct {
	value *PutUserRequest
	isSet bool
}

func (v NullablePutUserRequest) Get() *PutUserRequest {
	return v.value
}

func (v *NullablePutUserRequest) Set(val *PutUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutUserRequest(val *PutUserRequest) *NullablePutUserRequest {
	return &NullablePutUserRequest{value: val, isSet: true}
}

func (v NullablePutUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


