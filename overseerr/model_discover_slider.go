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

// checks if the DiscoverSlider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoverSlider{}

// DiscoverSlider struct for DiscoverSlider
type DiscoverSlider struct {
	Id *float32 `json:"id,omitempty"`
	Type float32 `json:"type"`
	Title NullableString `json:"title"`
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	Enabled bool `json:"enabled"`
	Data NullableString `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _DiscoverSlider DiscoverSlider

// NewDiscoverSlider instantiates a new DiscoverSlider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoverSlider(type_ float32, title NullableString, enabled bool, data NullableString) *DiscoverSlider {
	this := DiscoverSlider{}
	this.Type = type_
	this.Title = title
	this.Enabled = enabled
	this.Data = data
	return &this
}

// NewDiscoverSliderWithDefaults instantiates a new DiscoverSlider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoverSliderWithDefaults() *DiscoverSlider {
	this := DiscoverSlider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiscoverSlider) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverSlider) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiscoverSlider) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *DiscoverSlider) SetId(v float32) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *DiscoverSlider) GetType() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiscoverSlider) GetTypeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiscoverSlider) SetType(v float32) {
	o.Type = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DiscoverSlider) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoverSlider) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *DiscoverSlider) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetIsBuiltIn returns the IsBuiltIn field value if set, zero value otherwise.
func (o *DiscoverSlider) GetIsBuiltIn() bool {
	if o == nil || IsNil(o.IsBuiltIn) {
		var ret bool
		return ret
	}
	return *o.IsBuiltIn
}

// GetIsBuiltInOk returns a tuple with the IsBuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverSlider) GetIsBuiltInOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBuiltIn) {
		return nil, false
	}
	return o.IsBuiltIn, true
}

// HasIsBuiltIn returns a boolean if a field has been set.
func (o *DiscoverSlider) HasIsBuiltIn() bool {
	if o != nil && !IsNil(o.IsBuiltIn) {
		return true
	}

	return false
}

// SetIsBuiltIn gets a reference to the given bool and assigns it to the IsBuiltIn field.
func (o *DiscoverSlider) SetIsBuiltIn(v bool) {
	o.IsBuiltIn = &v
}

// GetEnabled returns the Enabled field value
func (o *DiscoverSlider) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DiscoverSlider) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DiscoverSlider) SetEnabled(v bool) {
	o.Enabled = v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DiscoverSlider) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}

	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoverSlider) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// SetData sets field value
func (o *DiscoverSlider) SetData(v string) {
	o.Data.Set(&v)
}

func (o DiscoverSlider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoverSlider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title.Get()
	if !IsNil(o.IsBuiltIn) {
		toSerialize["isBuiltIn"] = o.IsBuiltIn
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["data"] = o.Data.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoverSlider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"title",
		"enabled",
		"data",
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

	varDiscoverSlider := _DiscoverSlider{}

	err = json.Unmarshal(data, &varDiscoverSlider)

	if err != nil {
		return err
	}

	*o = DiscoverSlider(varDiscoverSlider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "isBuiltIn")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoverSlider struct {
	value *DiscoverSlider
	isSet bool
}

func (v NullableDiscoverSlider) Get() *DiscoverSlider {
	return v.value
}

func (v *NullableDiscoverSlider) Set(val *DiscoverSlider) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoverSlider) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoverSlider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoverSlider(val *DiscoverSlider) *NullableDiscoverSlider {
	return &NullableDiscoverSlider{value: val, isSet: true}
}

func (v NullableDiscoverSlider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoverSlider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


