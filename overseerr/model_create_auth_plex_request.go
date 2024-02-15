/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.33.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateAuthPlexRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAuthPlexRequest{}

// CreateAuthPlexRequest struct for CreateAuthPlexRequest
type CreateAuthPlexRequest struct {
	AuthToken string `json:"authToken"`
	AdditionalProperties map[string]interface{}
}

type _CreateAuthPlexRequest CreateAuthPlexRequest

// NewCreateAuthPlexRequest instantiates a new CreateAuthPlexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAuthPlexRequest(authToken string) *CreateAuthPlexRequest {
	this := CreateAuthPlexRequest{}
	this.AuthToken = authToken
	return &this
}

// NewCreateAuthPlexRequestWithDefaults instantiates a new CreateAuthPlexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAuthPlexRequestWithDefaults() *CreateAuthPlexRequest {
	this := CreateAuthPlexRequest{}
	return &this
}

// GetAuthToken returns the AuthToken field value
func (o *CreateAuthPlexRequest) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *CreateAuthPlexRequest) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *CreateAuthPlexRequest) SetAuthToken(v string) {
	o.AuthToken = v
}

func (o CreateAuthPlexRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAuthPlexRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authToken"] = o.AuthToken

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAuthPlexRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authToken",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateAuthPlexRequest := _CreateAuthPlexRequest{}

	err = json.Unmarshal(data, &varCreateAuthPlexRequest)

	if err != nil {
		return err
	}

	*o = CreateAuthPlexRequest(varCreateAuthPlexRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAuthPlexRequest struct {
	value *CreateAuthPlexRequest
	isSet bool
}

func (v NullableCreateAuthPlexRequest) Get() *CreateAuthPlexRequest {
	return v.value
}

func (v *NullableCreateAuthPlexRequest) Set(val *CreateAuthPlexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAuthPlexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAuthPlexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAuthPlexRequest(val *CreateAuthPlexRequest) *NullableCreateAuthPlexRequest {
	return &NullableCreateAuthPlexRequest{value: val, isSet: true}
}

func (v NullableCreateAuthPlexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAuthPlexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


