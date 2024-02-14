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

// Cast struct for Cast
type Cast struct {
	Id *float32 `json:"id,omitempty"`
	CastId *float32 `json:"castId,omitempty"`
	Character *string `json:"character,omitempty"`
	CreditId *string `json:"creditId,omitempty"`
	Gender *float32 `json:"gender,omitempty"`
	Name *string `json:"name,omitempty"`
	Order *float32 `json:"order,omitempty"`
	ProfilePath NullableString `json:"profilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Cast Cast

// NewCast instantiates a new Cast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCast() *Cast {
	this := Cast{}
	return &this
}

// NewCastWithDefaults instantiates a new Cast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithDefaults() *Cast {
	this := Cast{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cast) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cast) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Cast) SetId(v float32) {
	o.Id = &v
}

// GetCastId returns the CastId field value if set, zero value otherwise.
func (o *Cast) GetCastId() float32 {
	if o == nil || isNil(o.CastId) {
		var ret float32
		return ret
	}
	return *o.CastId
}

// GetCastIdOk returns a tuple with the CastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetCastIdOk() (*float32, bool) {
	if o == nil || isNil(o.CastId) {
    return nil, false
	}
	return o.CastId, true
}

// HasCastId returns a boolean if a field has been set.
func (o *Cast) HasCastId() bool {
	if o != nil && !isNil(o.CastId) {
		return true
	}

	return false
}

// SetCastId gets a reference to the given float32 and assigns it to the CastId field.
func (o *Cast) SetCastId(v float32) {
	o.CastId = &v
}

// GetCharacter returns the Character field value if set, zero value otherwise.
func (o *Cast) GetCharacter() string {
	if o == nil || isNil(o.Character) {
		var ret string
		return ret
	}
	return *o.Character
}

// GetCharacterOk returns a tuple with the Character field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetCharacterOk() (*string, bool) {
	if o == nil || isNil(o.Character) {
    return nil, false
	}
	return o.Character, true
}

// HasCharacter returns a boolean if a field has been set.
func (o *Cast) HasCharacter() bool {
	if o != nil && !isNil(o.Character) {
		return true
	}

	return false
}

// SetCharacter gets a reference to the given string and assigns it to the Character field.
func (o *Cast) SetCharacter(v string) {
	o.Character = &v
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *Cast) GetCreditId() string {
	if o == nil || isNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetCreditIdOk() (*string, bool) {
	if o == nil || isNil(o.CreditId) {
    return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *Cast) HasCreditId() bool {
	if o != nil && !isNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *Cast) SetCreditId(v string) {
	o.CreditId = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Cast) GetGender() float32 {
	if o == nil || isNil(o.Gender) {
		var ret float32
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetGenderOk() (*float32, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Cast) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given float32 and assigns it to the Gender field.
func (o *Cast) SetGender(v float32) {
	o.Gender = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Cast) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Cast) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Cast) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Cast) GetOrder() float32 {
	if o == nil || isNil(o.Order) {
		var ret float32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetOrderOk() (*float32, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Cast) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given float32 and assigns it to the Order field.
func (o *Cast) SetOrder(v float32) {
	o.Order = &v
}

// GetProfilePath returns the ProfilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cast) GetProfilePath() string {
	if o == nil || isNil(o.ProfilePath.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilePath.Get()
}

// GetProfilePathOk returns a tuple with the ProfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetProfilePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProfilePath.Get(), o.ProfilePath.IsSet()
}

// HasProfilePath returns a boolean if a field has been set.
func (o *Cast) HasProfilePath() bool {
	if o != nil && o.ProfilePath.IsSet() {
		return true
	}

	return false
}

// SetProfilePath gets a reference to the given NullableString and assigns it to the ProfilePath field.
func (o *Cast) SetProfilePath(v string) {
	o.ProfilePath.Set(&v)
}
// SetProfilePathNil sets the value for ProfilePath to be an explicit nil
func (o *Cast) SetProfilePathNil() {
	o.ProfilePath.Set(nil)
}

// UnsetProfilePath ensures that no value is present for ProfilePath, not even an explicit nil
func (o *Cast) UnsetProfilePath() {
	o.ProfilePath.Unset()
}

func (o Cast) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.CastId) {
		toSerialize["castId"] = o.CastId
	}
	if !isNil(o.Character) {
		toSerialize["character"] = o.Character
	}
	if !isNil(o.CreditId) {
		toSerialize["creditId"] = o.CreditId
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if o.ProfilePath.IsSet() {
		toSerialize["profilePath"] = o.ProfilePath.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Cast) UnmarshalJSON(bytes []byte) (err error) {
	varCast := _Cast{}

	if err = json.Unmarshal(bytes, &varCast); err == nil {
		*o = Cast(varCast)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "castId")
		delete(additionalProperties, "character")
		delete(additionalProperties, "creditId")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "name")
		delete(additionalProperties, "order")
		delete(additionalProperties, "profilePath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCast struct {
	value *Cast
	isSet bool
}

func (v NullableCast) Get() *Cast {
	return v.value
}

func (v *NullableCast) Set(val *Cast) {
	v.value = val
	v.isSet = true
}

func (v NullableCast) IsSet() bool {
	return v.isSet
}

func (v *NullableCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCast(val *Cast) *NullableCast {
	return &NullableCast{value: val, isSet: true}
}

func (v NullableCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


