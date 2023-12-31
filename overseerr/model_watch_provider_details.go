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

// WatchProviderDetails struct for WatchProviderDetails
type WatchProviderDetails struct {
	DisplayPriority *float32 `json:"displayPriority,omitempty"`
	LogoPath *string `json:"logoPath,omitempty"`
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WatchProviderDetails WatchProviderDetails

// NewWatchProviderDetails instantiates a new WatchProviderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchProviderDetails() *WatchProviderDetails {
	this := WatchProviderDetails{}
	return &this
}

// NewWatchProviderDetailsWithDefaults instantiates a new WatchProviderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchProviderDetailsWithDefaults() *WatchProviderDetails {
	this := WatchProviderDetails{}
	return &this
}

// GetDisplayPriority returns the DisplayPriority field value if set, zero value otherwise.
func (o *WatchProviderDetails) GetDisplayPriority() float32 {
	if o == nil || isNil(o.DisplayPriority) {
		var ret float32
		return ret
	}
	return *o.DisplayPriority
}

// GetDisplayPriorityOk returns a tuple with the DisplayPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProviderDetails) GetDisplayPriorityOk() (*float32, bool) {
	if o == nil || isNil(o.DisplayPriority) {
    return nil, false
	}
	return o.DisplayPriority, true
}

// HasDisplayPriority returns a boolean if a field has been set.
func (o *WatchProviderDetails) HasDisplayPriority() bool {
	if o != nil && !isNil(o.DisplayPriority) {
		return true
	}

	return false
}

// SetDisplayPriority gets a reference to the given float32 and assigns it to the DisplayPriority field.
func (o *WatchProviderDetails) SetDisplayPriority(v float32) {
	o.DisplayPriority = &v
}

// GetLogoPath returns the LogoPath field value if set, zero value otherwise.
func (o *WatchProviderDetails) GetLogoPath() string {
	if o == nil || isNil(o.LogoPath) {
		var ret string
		return ret
	}
	return *o.LogoPath
}

// GetLogoPathOk returns a tuple with the LogoPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProviderDetails) GetLogoPathOk() (*string, bool) {
	if o == nil || isNil(o.LogoPath) {
    return nil, false
	}
	return o.LogoPath, true
}

// HasLogoPath returns a boolean if a field has been set.
func (o *WatchProviderDetails) HasLogoPath() bool {
	if o != nil && !isNil(o.LogoPath) {
		return true
	}

	return false
}

// SetLogoPath gets a reference to the given string and assigns it to the LogoPath field.
func (o *WatchProviderDetails) SetLogoPath(v string) {
	o.LogoPath = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WatchProviderDetails) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProviderDetails) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WatchProviderDetails) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *WatchProviderDetails) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WatchProviderDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchProviderDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WatchProviderDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WatchProviderDetails) SetName(v string) {
	o.Name = &v
}

func (o WatchProviderDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DisplayPriority) {
		toSerialize["displayPriority"] = o.DisplayPriority
	}
	if !isNil(o.LogoPath) {
		toSerialize["logoPath"] = o.LogoPath
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchProviderDetails) UnmarshalJSON(bytes []byte) (err error) {
	varWatchProviderDetails := _WatchProviderDetails{}

	if err = json.Unmarshal(bytes, &varWatchProviderDetails); err == nil {
		*o = WatchProviderDetails(varWatchProviderDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "displayPriority")
		delete(additionalProperties, "logoPath")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchProviderDetails struct {
	value *WatchProviderDetails
	isSet bool
}

func (v NullableWatchProviderDetails) Get() *WatchProviderDetails {
	return v.value
}

func (v *NullableWatchProviderDetails) Set(val *WatchProviderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchProviderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchProviderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchProviderDetails(val *WatchProviderDetails) *NullableWatchProviderDetails {
	return &NullableWatchProviderDetails{value: val, isSet: true}
}

func (v NullableWatchProviderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchProviderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


