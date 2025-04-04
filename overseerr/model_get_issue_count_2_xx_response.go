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

// checks if the GetIssueCount2XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssueCount2XXResponse{}

// GetIssueCount2XXResponse struct for GetIssueCount2XXResponse
type GetIssueCount2XXResponse struct {
	Total *float32 `json:"total,omitempty"`
	Video *float32 `json:"video,omitempty"`
	Audio *float32 `json:"audio,omitempty"`
	Subtitles *float32 `json:"subtitles,omitempty"`
	Others *float32 `json:"others,omitempty"`
	Open *float32 `json:"open,omitempty"`
	Closed *float32 `json:"closed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetIssueCount2XXResponse GetIssueCount2XXResponse

// NewGetIssueCount2XXResponse instantiates a new GetIssueCount2XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssueCount2XXResponse() *GetIssueCount2XXResponse {
	this := GetIssueCount2XXResponse{}
	return &this
}

// NewGetIssueCount2XXResponseWithDefaults instantiates a new GetIssueCount2XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssueCount2XXResponseWithDefaults() *GetIssueCount2XXResponse {
	this := GetIssueCount2XXResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetIssueCount2XXResponse) SetTotal(v float32) {
	o.Total = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetVideo() float32 {
	if o == nil || IsNil(o.Video) {
		var ret float32
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetVideoOk() (*float32, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given float32 and assigns it to the Video field.
func (o *GetIssueCount2XXResponse) SetVideo(v float32) {
	o.Video = &v
}

// GetAudio returns the Audio field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetAudio() float32 {
	if o == nil || IsNil(o.Audio) {
		var ret float32
		return ret
	}
	return *o.Audio
}

// GetAudioOk returns a tuple with the Audio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetAudioOk() (*float32, bool) {
	if o == nil || IsNil(o.Audio) {
		return nil, false
	}
	return o.Audio, true
}

// HasAudio returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasAudio() bool {
	if o != nil && !IsNil(o.Audio) {
		return true
	}

	return false
}

// SetAudio gets a reference to the given float32 and assigns it to the Audio field.
func (o *GetIssueCount2XXResponse) SetAudio(v float32) {
	o.Audio = &v
}

// GetSubtitles returns the Subtitles field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetSubtitles() float32 {
	if o == nil || IsNil(o.Subtitles) {
		var ret float32
		return ret
	}
	return *o.Subtitles
}

// GetSubtitlesOk returns a tuple with the Subtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetSubtitlesOk() (*float32, bool) {
	if o == nil || IsNil(o.Subtitles) {
		return nil, false
	}
	return o.Subtitles, true
}

// HasSubtitles returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasSubtitles() bool {
	if o != nil && !IsNil(o.Subtitles) {
		return true
	}

	return false
}

// SetSubtitles gets a reference to the given float32 and assigns it to the Subtitles field.
func (o *GetIssueCount2XXResponse) SetSubtitles(v float32) {
	o.Subtitles = &v
}

// GetOthers returns the Others field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetOthers() float32 {
	if o == nil || IsNil(o.Others) {
		var ret float32
		return ret
	}
	return *o.Others
}

// GetOthersOk returns a tuple with the Others field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetOthersOk() (*float32, bool) {
	if o == nil || IsNil(o.Others) {
		return nil, false
	}
	return o.Others, true
}

// HasOthers returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasOthers() bool {
	if o != nil && !IsNil(o.Others) {
		return true
	}

	return false
}

// SetOthers gets a reference to the given float32 and assigns it to the Others field.
func (o *GetIssueCount2XXResponse) SetOthers(v float32) {
	o.Others = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetOpen() float32 {
	if o == nil || IsNil(o.Open) {
		var ret float32
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetOpenOk() (*float32, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given float32 and assigns it to the Open field.
func (o *GetIssueCount2XXResponse) SetOpen(v float32) {
	o.Open = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *GetIssueCount2XXResponse) GetClosed() float32 {
	if o == nil || IsNil(o.Closed) {
		var ret float32
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCount2XXResponse) GetClosedOk() (*float32, bool) {
	if o == nil || IsNil(o.Closed) {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *GetIssueCount2XXResponse) HasClosed() bool {
	if o != nil && !IsNil(o.Closed) {
		return true
	}

	return false
}

// SetClosed gets a reference to the given float32 and assigns it to the Closed field.
func (o *GetIssueCount2XXResponse) SetClosed(v float32) {
	o.Closed = &v
}

func (o GetIssueCount2XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssueCount2XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.Audio) {
		toSerialize["audio"] = o.Audio
	}
	if !IsNil(o.Subtitles) {
		toSerialize["subtitles"] = o.Subtitles
	}
	if !IsNil(o.Others) {
		toSerialize["others"] = o.Others
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Closed) {
		toSerialize["closed"] = o.Closed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetIssueCount2XXResponse) UnmarshalJSON(data []byte) (err error) {
	varGetIssueCount2XXResponse := _GetIssueCount2XXResponse{}

	err = json.Unmarshal(data, &varGetIssueCount2XXResponse)

	if err != nil {
		return err
	}

	*o = GetIssueCount2XXResponse(varGetIssueCount2XXResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "video")
		delete(additionalProperties, "audio")
		delete(additionalProperties, "subtitles")
		delete(additionalProperties, "others")
		delete(additionalProperties, "open")
		delete(additionalProperties, "closed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetIssueCount2XXResponse struct {
	value *GetIssueCount2XXResponse
	isSet bool
}

func (v NullableGetIssueCount2XXResponse) Get() *GetIssueCount2XXResponse {
	return v.value
}

func (v *NullableGetIssueCount2XXResponse) Set(val *GetIssueCount2XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssueCount2XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssueCount2XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssueCount2XXResponse(val *GetIssueCount2XXResponse) *NullableGetIssueCount2XXResponse {
	return &NullableGetIssueCount2XXResponse{value: val, isSet: true}
}

func (v NullableGetIssueCount2XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssueCount2XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


