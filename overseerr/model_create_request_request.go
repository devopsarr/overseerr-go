/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateRequestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRequestRequest{}

// CreateRequestRequest struct for CreateRequestRequest
type CreateRequestRequest struct {
	MediaType string `json:"mediaType"`
	MediaId float32 `json:"mediaId"`
	TvdbId *float32 `json:"tvdbId,omitempty"`
	Seasons *CreateRequestRequestSeasons `json:"seasons,omitempty"`
	Is4k *bool `json:"is4k,omitempty"`
	ServerId *float32 `json:"serverId,omitempty"`
	ProfileId *float32 `json:"profileId,omitempty"`
	RootFolder *string `json:"rootFolder,omitempty"`
	LanguageProfileId *float32 `json:"languageProfileId,omitempty"`
	UserId NullableFloat32 `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRequestRequest CreateRequestRequest

// NewCreateRequestRequest instantiates a new CreateRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRequestRequest(mediaType string, mediaId float32) *CreateRequestRequest {
	this := CreateRequestRequest{}
	this.MediaType = mediaType
	this.MediaId = mediaId
	return &this
}

// NewCreateRequestRequestWithDefaults instantiates a new CreateRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRequestRequestWithDefaults() *CreateRequestRequest {
	this := CreateRequestRequest{}
	return &this
}

// GetMediaType returns the MediaType field value
func (o *CreateRequestRequest) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *CreateRequestRequest) SetMediaType(v string) {
	o.MediaType = v
}

// GetMediaId returns the MediaId field value
func (o *CreateRequestRequest) GetMediaId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetMediaIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaId, true
}

// SetMediaId sets field value
func (o *CreateRequestRequest) SetMediaId(v float32) {
	o.MediaId = v
}

// GetTvdbId returns the TvdbId field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetTvdbId() float32 {
	if o == nil || IsNil(o.TvdbId) {
		var ret float32
		return ret
	}
	return *o.TvdbId
}

// GetTvdbIdOk returns a tuple with the TvdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetTvdbIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TvdbId) {
		return nil, false
	}
	return o.TvdbId, true
}

// HasTvdbId returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasTvdbId() bool {
	if o != nil && !IsNil(o.TvdbId) {
		return true
	}

	return false
}

// SetTvdbId gets a reference to the given float32 and assigns it to the TvdbId field.
func (o *CreateRequestRequest) SetTvdbId(v float32) {
	o.TvdbId = &v
}

// GetSeasons returns the Seasons field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetSeasons() CreateRequestRequestSeasons {
	if o == nil || IsNil(o.Seasons) {
		var ret CreateRequestRequestSeasons
		return ret
	}
	return *o.Seasons
}

// GetSeasonsOk returns a tuple with the Seasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetSeasonsOk() (*CreateRequestRequestSeasons, bool) {
	if o == nil || IsNil(o.Seasons) {
		return nil, false
	}
	return o.Seasons, true
}

// HasSeasons returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasSeasons() bool {
	if o != nil && !IsNil(o.Seasons) {
		return true
	}

	return false
}

// SetSeasons gets a reference to the given CreateRequestRequestSeasons and assigns it to the Seasons field.
func (o *CreateRequestRequest) SetSeasons(v CreateRequestRequestSeasons) {
	o.Seasons = &v
}

// GetIs4k returns the Is4k field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetIs4k() bool {
	if o == nil || IsNil(o.Is4k) {
		var ret bool
		return ret
	}
	return *o.Is4k
}

// GetIs4kOk returns a tuple with the Is4k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetIs4kOk() (*bool, bool) {
	if o == nil || IsNil(o.Is4k) {
		return nil, false
	}
	return o.Is4k, true
}

// HasIs4k returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasIs4k() bool {
	if o != nil && !IsNil(o.Is4k) {
		return true
	}

	return false
}

// SetIs4k gets a reference to the given bool and assigns it to the Is4k field.
func (o *CreateRequestRequest) SetIs4k(v bool) {
	o.Is4k = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetServerId() float32 {
	if o == nil || IsNil(o.ServerId) {
		var ret float32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetServerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given float32 and assigns it to the ServerId field.
func (o *CreateRequestRequest) SetServerId(v float32) {
	o.ServerId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetProfileId() float32 {
	if o == nil || IsNil(o.ProfileId) {
		var ret float32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given float32 and assigns it to the ProfileId field.
func (o *CreateRequestRequest) SetProfileId(v float32) {
	o.ProfileId = &v
}

// GetRootFolder returns the RootFolder field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetRootFolder() string {
	if o == nil || IsNil(o.RootFolder) {
		var ret string
		return ret
	}
	return *o.RootFolder
}

// GetRootFolderOk returns a tuple with the RootFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetRootFolderOk() (*string, bool) {
	if o == nil || IsNil(o.RootFolder) {
		return nil, false
	}
	return o.RootFolder, true
}

// HasRootFolder returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasRootFolder() bool {
	if o != nil && !IsNil(o.RootFolder) {
		return true
	}

	return false
}

// SetRootFolder gets a reference to the given string and assigns it to the RootFolder field.
func (o *CreateRequestRequest) SetRootFolder(v string) {
	o.RootFolder = &v
}

// GetLanguageProfileId returns the LanguageProfileId field value if set, zero value otherwise.
func (o *CreateRequestRequest) GetLanguageProfileId() float32 {
	if o == nil || IsNil(o.LanguageProfileId) {
		var ret float32
		return ret
	}
	return *o.LanguageProfileId
}

// GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRequestRequest) GetLanguageProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LanguageProfileId) {
		return nil, false
	}
	return o.LanguageProfileId, true
}

// HasLanguageProfileId returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasLanguageProfileId() bool {
	if o != nil && !IsNil(o.LanguageProfileId) {
		return true
	}

	return false
}

// SetLanguageProfileId gets a reference to the given float32 and assigns it to the LanguageProfileId field.
func (o *CreateRequestRequest) SetLanguageProfileId(v float32) {
	o.LanguageProfileId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRequestRequest) GetUserId() float32 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret float32
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRequestRequest) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CreateRequestRequest) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableFloat32 and assigns it to the UserId field.
func (o *CreateRequestRequest) SetUserId(v float32) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CreateRequestRequest) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CreateRequestRequest) UnsetUserId() {
	o.UserId.Unset()
}

func (o CreateRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mediaType"] = o.MediaType
	toSerialize["mediaId"] = o.MediaId
	if !IsNil(o.TvdbId) {
		toSerialize["tvdbId"] = o.TvdbId
	}
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	if !IsNil(o.Is4k) {
		toSerialize["is4k"] = o.Is4k
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.RootFolder) {
		toSerialize["rootFolder"] = o.RootFolder
	}
	if !IsNil(o.LanguageProfileId) {
		toSerialize["languageProfileId"] = o.LanguageProfileId
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRequestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mediaType",
		"mediaId",
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

	varCreateRequestRequest := _CreateRequestRequest{}

	err = json.Unmarshal(data, &varCreateRequestRequest)

	if err != nil {
		return err
	}

	*o = CreateRequestRequest(varCreateRequestRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mediaType")
		delete(additionalProperties, "mediaId")
		delete(additionalProperties, "tvdbId")
		delete(additionalProperties, "seasons")
		delete(additionalProperties, "is4k")
		delete(additionalProperties, "serverId")
		delete(additionalProperties, "profileId")
		delete(additionalProperties, "rootFolder")
		delete(additionalProperties, "languageProfileId")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRequestRequest struct {
	value *CreateRequestRequest
	isSet bool
}

func (v NullableCreateRequestRequest) Get() *CreateRequestRequest {
	return v.value
}

func (v *NullableCreateRequestRequest) Set(val *CreateRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRequestRequest(val *CreateRequestRequest) *NullableCreateRequestRequest {
	return &NullableCreateRequestRequest{value: val, isSet: true}
}

func (v NullableCreateRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


