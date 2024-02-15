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

// checks if the PlexSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlexSettings{}

// PlexSettings struct for PlexSettings
type PlexSettings struct {
	Name string `json:"name"`
	MachineId string `json:"machineId"`
	Ip string `json:"ip"`
	Port float32 `json:"port"`
	UseSsl NullableBool `json:"useSsl,omitempty"`
	Libraries []PlexLibrary `json:"libraries,omitempty"`
	WebAppUrl NullableString `json:"webAppUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlexSettings PlexSettings

// NewPlexSettings instantiates a new PlexSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlexSettings(name string, machineId string, ip string, port float32) *PlexSettings {
	this := PlexSettings{}
	this.Name = name
	this.MachineId = machineId
	this.Ip = ip
	this.Port = port
	return &this
}

// NewPlexSettingsWithDefaults instantiates a new PlexSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlexSettingsWithDefaults() *PlexSettings {
	this := PlexSettings{}
	return &this
}

// GetName returns the Name field value
func (o *PlexSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlexSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlexSettings) SetName(v string) {
	o.Name = v
}

// GetMachineId returns the MachineId field value
func (o *PlexSettings) GetMachineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MachineId
}

// GetMachineIdOk returns a tuple with the MachineId field value
// and a boolean to check if the value has been set.
func (o *PlexSettings) GetMachineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MachineId, true
}

// SetMachineId sets field value
func (o *PlexSettings) SetMachineId(v string) {
	o.MachineId = v
}

// GetIp returns the Ip field value
func (o *PlexSettings) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *PlexSettings) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *PlexSettings) SetIp(v string) {
	o.Ip = v
}

// GetPort returns the Port field value
func (o *PlexSettings) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PlexSettings) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PlexSettings) SetPort(v float32) {
	o.Port = v
}

// GetUseSsl returns the UseSsl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSettings) GetUseSsl() bool {
	if o == nil || IsNil(o.UseSsl.Get()) {
		var ret bool
		return ret
	}
	return *o.UseSsl.Get()
}

// GetUseSslOk returns a tuple with the UseSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSettings) GetUseSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseSsl.Get(), o.UseSsl.IsSet()
}

// HasUseSsl returns a boolean if a field has been set.
func (o *PlexSettings) HasUseSsl() bool {
	if o != nil && o.UseSsl.IsSet() {
		return true
	}

	return false
}

// SetUseSsl gets a reference to the given NullableBool and assigns it to the UseSsl field.
func (o *PlexSettings) SetUseSsl(v bool) {
	o.UseSsl.Set(&v)
}
// SetUseSslNil sets the value for UseSsl to be an explicit nil
func (o *PlexSettings) SetUseSslNil() {
	o.UseSsl.Set(nil)
}

// UnsetUseSsl ensures that no value is present for UseSsl, not even an explicit nil
func (o *PlexSettings) UnsetUseSsl() {
	o.UseSsl.Unset()
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *PlexSettings) GetLibraries() []PlexLibrary {
	if o == nil || IsNil(o.Libraries) {
		var ret []PlexLibrary
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlexSettings) GetLibrariesOk() ([]PlexLibrary, bool) {
	if o == nil || IsNil(o.Libraries) {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *PlexSettings) HasLibraries() bool {
	if o != nil && !IsNil(o.Libraries) {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []PlexLibrary and assigns it to the Libraries field.
func (o *PlexSettings) SetLibraries(v []PlexLibrary) {
	o.Libraries = v
}

// GetWebAppUrl returns the WebAppUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlexSettings) GetWebAppUrl() string {
	if o == nil || IsNil(o.WebAppUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebAppUrl.Get()
}

// GetWebAppUrlOk returns a tuple with the WebAppUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlexSettings) GetWebAppUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebAppUrl.Get(), o.WebAppUrl.IsSet()
}

// HasWebAppUrl returns a boolean if a field has been set.
func (o *PlexSettings) HasWebAppUrl() bool {
	if o != nil && o.WebAppUrl.IsSet() {
		return true
	}

	return false
}

// SetWebAppUrl gets a reference to the given NullableString and assigns it to the WebAppUrl field.
func (o *PlexSettings) SetWebAppUrl(v string) {
	o.WebAppUrl.Set(&v)
}
// SetWebAppUrlNil sets the value for WebAppUrl to be an explicit nil
func (o *PlexSettings) SetWebAppUrlNil() {
	o.WebAppUrl.Set(nil)
}

// UnsetWebAppUrl ensures that no value is present for WebAppUrl, not even an explicit nil
func (o *PlexSettings) UnsetWebAppUrl() {
	o.WebAppUrl.Unset()
}

func (o PlexSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlexSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["machineId"] = o.MachineId
	toSerialize["ip"] = o.Ip
	toSerialize["port"] = o.Port
	if o.UseSsl.IsSet() {
		toSerialize["useSsl"] = o.UseSsl.Get()
	}
	if !IsNil(o.Libraries) {
		toSerialize["libraries"] = o.Libraries
	}
	if o.WebAppUrl.IsSet() {
		toSerialize["webAppUrl"] = o.WebAppUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlexSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"machineId",
		"ip",
		"port",
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

	varPlexSettings := _PlexSettings{}

	err = json.Unmarshal(data, &varPlexSettings)

	if err != nil {
		return err
	}

	*o = PlexSettings(varPlexSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "machineId")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "port")
		delete(additionalProperties, "useSsl")
		delete(additionalProperties, "libraries")
		delete(additionalProperties, "webAppUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlexSettings struct {
	value *PlexSettings
	isSet bool
}

func (v NullablePlexSettings) Get() *PlexSettings {
	return v.value
}

func (v *NullablePlexSettings) Set(val *PlexSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePlexSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePlexSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlexSettings(val *PlexSettings) *NullablePlexSettings {
	return &NullablePlexSettings{value: val, isSet: true}
}

func (v NullablePlexSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlexSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


