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

// Crew struct for Crew
type Crew struct {
	Id *float32 `json:"id,omitempty"`
	CreditId *string `json:"creditId,omitempty"`
	Gender *float32 `json:"gender,omitempty"`
	Name *string `json:"name,omitempty"`
	Job *string `json:"job,omitempty"`
	Department *string `json:"department,omitempty"`
	ProfilePath NullableString `json:"profilePath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Crew Crew

// NewCrew instantiates a new Crew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrew() *Crew {
	this := Crew{}
	return &this
}

// NewCrewWithDefaults instantiates a new Crew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrewWithDefaults() *Crew {
	this := Crew{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Crew) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Crew) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Crew) SetId(v float32) {
	o.Id = &v
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *Crew) GetCreditId() string {
	if o == nil || isNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetCreditIdOk() (*string, bool) {
	if o == nil || isNil(o.CreditId) {
    return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *Crew) HasCreditId() bool {
	if o != nil && !isNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *Crew) SetCreditId(v string) {
	o.CreditId = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Crew) GetGender() float32 {
	if o == nil || isNil(o.Gender) {
		var ret float32
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetGenderOk() (*float32, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Crew) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given float32 and assigns it to the Gender field.
func (o *Crew) SetGender(v float32) {
	o.Gender = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Crew) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Crew) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Crew) SetName(v string) {
	o.Name = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *Crew) GetJob() string {
	if o == nil || isNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetJobOk() (*string, bool) {
	if o == nil || isNil(o.Job) {
    return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *Crew) HasJob() bool {
	if o != nil && !isNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *Crew) SetJob(v string) {
	o.Job = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Crew) GetDepartment() string {
	if o == nil || isNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Crew) GetDepartmentOk() (*string, bool) {
	if o == nil || isNil(o.Department) {
    return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Crew) HasDepartment() bool {
	if o != nil && !isNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Crew) SetDepartment(v string) {
	o.Department = &v
}

// GetProfilePath returns the ProfilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Crew) GetProfilePath() string {
	if o == nil || isNil(o.ProfilePath.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilePath.Get()
}

// GetProfilePathOk returns a tuple with the ProfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Crew) GetProfilePathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProfilePath.Get(), o.ProfilePath.IsSet()
}

// HasProfilePath returns a boolean if a field has been set.
func (o *Crew) HasProfilePath() bool {
	if o != nil && o.ProfilePath.IsSet() {
		return true
	}

	return false
}

// SetProfilePath gets a reference to the given NullableString and assigns it to the ProfilePath field.
func (o *Crew) SetProfilePath(v string) {
	o.ProfilePath.Set(&v)
}
// SetProfilePathNil sets the value for ProfilePath to be an explicit nil
func (o *Crew) SetProfilePathNil() {
	o.ProfilePath.Set(nil)
}

// UnsetProfilePath ensures that no value is present for ProfilePath, not even an explicit nil
func (o *Crew) UnsetProfilePath() {
	o.ProfilePath.Unset()
}

func (o Crew) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !isNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !isNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if o.ProfilePath.IsSet() {
		toSerialize["profilePath"] = o.ProfilePath.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Crew) UnmarshalJSON(bytes []byte) (err error) {
	varCrew := _Crew{}

	if err = json.Unmarshal(bytes, &varCrew); err == nil {
		*o = Crew(varCrew)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "creditId")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "name")
		delete(additionalProperties, "job")
		delete(additionalProperties, "department")
		delete(additionalProperties, "profilePath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrew struct {
	value *Crew
	isSet bool
}

func (v NullableCrew) Get() *Crew {
	return v.value
}

func (v *NullableCrew) Set(val *Crew) {
	v.value = val
	v.isSet = true
}

func (v NullableCrew) IsSet() bool {
	return v.isSet
}

func (v *NullableCrew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrew(val *Crew) *NullableCrew {
	return &NullableCrew{value: val, isSet: true}
}

func (v NullableCrew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

