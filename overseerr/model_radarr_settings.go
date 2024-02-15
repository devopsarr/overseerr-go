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

// checks if the RadarrSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadarrSettings{}

// RadarrSettings struct for RadarrSettings
type RadarrSettings struct {
	Id *float32 `json:"id,omitempty"`
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	Port float32 `json:"port"`
	ApiKey string `json:"apiKey"`
	UseSsl bool `json:"useSsl"`
	BaseUrl *string `json:"baseUrl,omitempty"`
	ActiveProfileId float32 `json:"activeProfileId"`
	ActiveProfileName string `json:"activeProfileName"`
	ActiveDirectory string `json:"activeDirectory"`
	Is4k bool `json:"is4k"`
	MinimumAvailability string `json:"minimumAvailability"`
	IsDefault bool `json:"isDefault"`
	ExternalUrl *string `json:"externalUrl,omitempty"`
	SyncEnabled *bool `json:"syncEnabled,omitempty"`
	PreventSearch *bool `json:"preventSearch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RadarrSettings RadarrSettings

// NewRadarrSettings instantiates a new RadarrSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadarrSettings(name string, hostname string, port float32, apiKey string, useSsl bool, activeProfileId float32, activeProfileName string, activeDirectory string, is4k bool, minimumAvailability string, isDefault bool) *RadarrSettings {
	this := RadarrSettings{}
	this.Name = name
	this.Hostname = hostname
	this.Port = port
	this.ApiKey = apiKey
	this.UseSsl = useSsl
	this.ActiveProfileId = activeProfileId
	this.ActiveProfileName = activeProfileName
	this.ActiveDirectory = activeDirectory
	this.Is4k = is4k
	this.MinimumAvailability = minimumAvailability
	this.IsDefault = isDefault
	return &this
}

// NewRadarrSettingsWithDefaults instantiates a new RadarrSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadarrSettingsWithDefaults() *RadarrSettings {
	this := RadarrSettings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RadarrSettings) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RadarrSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *RadarrSettings) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RadarrSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RadarrSettings) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *RadarrSettings) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *RadarrSettings) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *RadarrSettings) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *RadarrSettings) SetPort(v float32) {
	o.Port = v
}

// GetApiKey returns the ApiKey field value
func (o *RadarrSettings) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *RadarrSettings) SetApiKey(v string) {
	o.ApiKey = v
}

// GetUseSsl returns the UseSsl field value
func (o *RadarrSettings) GetUseSsl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetUseSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSsl, true
}

// SetUseSsl sets field value
func (o *RadarrSettings) SetUseSsl(v bool) {
	o.UseSsl = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *RadarrSettings) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *RadarrSettings) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *RadarrSettings) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetActiveProfileId returns the ActiveProfileId field value
func (o *RadarrSettings) GetActiveProfileId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ActiveProfileId
}

// GetActiveProfileIdOk returns a tuple with the ActiveProfileId field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetActiveProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveProfileId, true
}

// SetActiveProfileId sets field value
func (o *RadarrSettings) SetActiveProfileId(v float32) {
	o.ActiveProfileId = v
}

// GetActiveProfileName returns the ActiveProfileName field value
func (o *RadarrSettings) GetActiveProfileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveProfileName
}

// GetActiveProfileNameOk returns a tuple with the ActiveProfileName field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetActiveProfileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveProfileName, true
}

// SetActiveProfileName sets field value
func (o *RadarrSettings) SetActiveProfileName(v string) {
	o.ActiveProfileName = v
}

// GetActiveDirectory returns the ActiveDirectory field value
func (o *RadarrSettings) GetActiveDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveDirectory
}

// GetActiveDirectoryOk returns a tuple with the ActiveDirectory field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetActiveDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveDirectory, true
}

// SetActiveDirectory sets field value
func (o *RadarrSettings) SetActiveDirectory(v string) {
	o.ActiveDirectory = v
}

// GetIs4k returns the Is4k field value
func (o *RadarrSettings) GetIs4k() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Is4k
}

// GetIs4kOk returns a tuple with the Is4k field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetIs4kOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Is4k, true
}

// SetIs4k sets field value
func (o *RadarrSettings) SetIs4k(v bool) {
	o.Is4k = v
}

// GetMinimumAvailability returns the MinimumAvailability field value
func (o *RadarrSettings) GetMinimumAvailability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumAvailability
}

// GetMinimumAvailabilityOk returns a tuple with the MinimumAvailability field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetMinimumAvailabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumAvailability, true
}

// SetMinimumAvailability sets field value
func (o *RadarrSettings) SetMinimumAvailability(v string) {
	o.MinimumAvailability = v
}

// GetIsDefault returns the IsDefault field value
func (o *RadarrSettings) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *RadarrSettings) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *RadarrSettings) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *RadarrSettings) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *RadarrSettings) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *RadarrSettings) GetSyncEnabled() bool {
	if o == nil || IsNil(o.SyncEnabled) {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncEnabled) {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *RadarrSettings) HasSyncEnabled() bool {
	if o != nil && !IsNil(o.SyncEnabled) {
		return true
	}

	return false
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *RadarrSettings) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// GetPreventSearch returns the PreventSearch field value if set, zero value otherwise.
func (o *RadarrSettings) GetPreventSearch() bool {
	if o == nil || IsNil(o.PreventSearch) {
		var ret bool
		return ret
	}
	return *o.PreventSearch
}

// GetPreventSearchOk returns a tuple with the PreventSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RadarrSettings) GetPreventSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventSearch) {
		return nil, false
	}
	return o.PreventSearch, true
}

// HasPreventSearch returns a boolean if a field has been set.
func (o *RadarrSettings) HasPreventSearch() bool {
	if o != nil && !IsNil(o.PreventSearch) {
		return true
	}

	return false
}

// SetPreventSearch gets a reference to the given bool and assigns it to the PreventSearch field.
func (o *RadarrSettings) SetPreventSearch(v bool) {
	o.PreventSearch = &v
}

func (o RadarrSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadarrSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["useSsl"] = o.UseSsl
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	toSerialize["activeProfileId"] = o.ActiveProfileId
	toSerialize["activeProfileName"] = o.ActiveProfileName
	toSerialize["activeDirectory"] = o.ActiveDirectory
	toSerialize["is4k"] = o.Is4k
	toSerialize["minimumAvailability"] = o.MinimumAvailability
	toSerialize["isDefault"] = o.IsDefault
	if !IsNil(o.ExternalUrl) {
		toSerialize["externalUrl"] = o.ExternalUrl
	}
	if !IsNil(o.SyncEnabled) {
		toSerialize["syncEnabled"] = o.SyncEnabled
	}
	if !IsNil(o.PreventSearch) {
		toSerialize["preventSearch"] = o.PreventSearch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RadarrSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"hostname",
		"port",
		"apiKey",
		"useSsl",
		"activeProfileId",
		"activeProfileName",
		"activeDirectory",
		"is4k",
		"minimumAvailability",
		"isDefault",
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

	varRadarrSettings := _RadarrSettings{}

	err = json.Unmarshal(data, &varRadarrSettings)

	if err != nil {
		return err
	}

	*o = RadarrSettings(varRadarrSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "port")
		delete(additionalProperties, "apiKey")
		delete(additionalProperties, "useSsl")
		delete(additionalProperties, "baseUrl")
		delete(additionalProperties, "activeProfileId")
		delete(additionalProperties, "activeProfileName")
		delete(additionalProperties, "activeDirectory")
		delete(additionalProperties, "is4k")
		delete(additionalProperties, "minimumAvailability")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "externalUrl")
		delete(additionalProperties, "syncEnabled")
		delete(additionalProperties, "preventSearch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRadarrSettings struct {
	value *RadarrSettings
	isSet bool
}

func (v NullableRadarrSettings) Get() *RadarrSettings {
	return v.value
}

func (v *NullableRadarrSettings) Set(val *RadarrSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRadarrSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRadarrSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadarrSettings(val *RadarrSettings) *NullableRadarrSettings {
	return &NullableRadarrSettings{value: val, isSet: true}
}

func (v NullableRadarrSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadarrSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


