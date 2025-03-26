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

// checks if the UserSettingsNotifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSettingsNotifications{}

// UserSettingsNotifications struct for UserSettingsNotifications
type UserSettingsNotifications struct {
	NotificationTypes *NotificationAgentTypes `json:"notificationTypes,omitempty"`
	EmailEnabled *bool `json:"emailEnabled,omitempty"`
	PgpKey NullableString `json:"pgpKey,omitempty"`
	DiscordEnabled *bool `json:"discordEnabled,omitempty"`
	DiscordEnabledTypes NullableFloat32 `json:"discordEnabledTypes,omitempty"`
	DiscordId NullableString `json:"discordId,omitempty"`
	PushbulletAccessToken NullableString `json:"pushbulletAccessToken,omitempty"`
	PushoverApplicationToken NullableString `json:"pushoverApplicationToken,omitempty"`
	PushoverUserKey NullableString `json:"pushoverUserKey,omitempty"`
	PushoverSound NullableString `json:"pushoverSound,omitempty"`
	TelegramEnabled *bool `json:"telegramEnabled,omitempty"`
	TelegramBotUsername NullableString `json:"telegramBotUsername,omitempty"`
	TelegramChatId NullableString `json:"telegramChatId,omitempty"`
	TelegramSendSilently NullableBool `json:"telegramSendSilently,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSettingsNotifications UserSettingsNotifications

// NewUserSettingsNotifications instantiates a new UserSettingsNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSettingsNotifications() *UserSettingsNotifications {
	this := UserSettingsNotifications{}
	return &this
}

// NewUserSettingsNotificationsWithDefaults instantiates a new UserSettingsNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingsNotificationsWithDefaults() *UserSettingsNotifications {
	this := UserSettingsNotifications{}
	return &this
}

// GetNotificationTypes returns the NotificationTypes field value if set, zero value otherwise.
func (o *UserSettingsNotifications) GetNotificationTypes() NotificationAgentTypes {
	if o == nil || IsNil(o.NotificationTypes) {
		var ret NotificationAgentTypes
		return ret
	}
	return *o.NotificationTypes
}

// GetNotificationTypesOk returns a tuple with the NotificationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsNotifications) GetNotificationTypesOk() (*NotificationAgentTypes, bool) {
	if o == nil || IsNil(o.NotificationTypes) {
		return nil, false
	}
	return o.NotificationTypes, true
}

// HasNotificationTypes returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasNotificationTypes() bool {
	if o != nil && !IsNil(o.NotificationTypes) {
		return true
	}

	return false
}

// SetNotificationTypes gets a reference to the given NotificationAgentTypes and assigns it to the NotificationTypes field.
func (o *UserSettingsNotifications) SetNotificationTypes(v NotificationAgentTypes) {
	o.NotificationTypes = &v
}

// GetEmailEnabled returns the EmailEnabled field value if set, zero value otherwise.
func (o *UserSettingsNotifications) GetEmailEnabled() bool {
	if o == nil || IsNil(o.EmailEnabled) {
		var ret bool
		return ret
	}
	return *o.EmailEnabled
}

// GetEmailEnabledOk returns a tuple with the EmailEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsNotifications) GetEmailEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailEnabled) {
		return nil, false
	}
	return o.EmailEnabled, true
}

// HasEmailEnabled returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasEmailEnabled() bool {
	if o != nil && !IsNil(o.EmailEnabled) {
		return true
	}

	return false
}

// SetEmailEnabled gets a reference to the given bool and assigns it to the EmailEnabled field.
func (o *UserSettingsNotifications) SetEmailEnabled(v bool) {
	o.EmailEnabled = &v
}

// GetPgpKey returns the PgpKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetPgpKey() string {
	if o == nil || IsNil(o.PgpKey.Get()) {
		var ret string
		return ret
	}
	return *o.PgpKey.Get()
}

// GetPgpKeyOk returns a tuple with the PgpKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetPgpKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PgpKey.Get(), o.PgpKey.IsSet()
}

// HasPgpKey returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasPgpKey() bool {
	if o != nil && o.PgpKey.IsSet() {
		return true
	}

	return false
}

// SetPgpKey gets a reference to the given NullableString and assigns it to the PgpKey field.
func (o *UserSettingsNotifications) SetPgpKey(v string) {
	o.PgpKey.Set(&v)
}
// SetPgpKeyNil sets the value for PgpKey to be an explicit nil
func (o *UserSettingsNotifications) SetPgpKeyNil() {
	o.PgpKey.Set(nil)
}

// UnsetPgpKey ensures that no value is present for PgpKey, not even an explicit nil
func (o *UserSettingsNotifications) UnsetPgpKey() {
	o.PgpKey.Unset()
}

// GetDiscordEnabled returns the DiscordEnabled field value if set, zero value otherwise.
func (o *UserSettingsNotifications) GetDiscordEnabled() bool {
	if o == nil || IsNil(o.DiscordEnabled) {
		var ret bool
		return ret
	}
	return *o.DiscordEnabled
}

// GetDiscordEnabledOk returns a tuple with the DiscordEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsNotifications) GetDiscordEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscordEnabled) {
		return nil, false
	}
	return o.DiscordEnabled, true
}

// HasDiscordEnabled returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasDiscordEnabled() bool {
	if o != nil && !IsNil(o.DiscordEnabled) {
		return true
	}

	return false
}

// SetDiscordEnabled gets a reference to the given bool and assigns it to the DiscordEnabled field.
func (o *UserSettingsNotifications) SetDiscordEnabled(v bool) {
	o.DiscordEnabled = &v
}

// GetDiscordEnabledTypes returns the DiscordEnabledTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetDiscordEnabledTypes() float32 {
	if o == nil || IsNil(o.DiscordEnabledTypes.Get()) {
		var ret float32
		return ret
	}
	return *o.DiscordEnabledTypes.Get()
}

// GetDiscordEnabledTypesOk returns a tuple with the DiscordEnabledTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetDiscordEnabledTypesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscordEnabledTypes.Get(), o.DiscordEnabledTypes.IsSet()
}

// HasDiscordEnabledTypes returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasDiscordEnabledTypes() bool {
	if o != nil && o.DiscordEnabledTypes.IsSet() {
		return true
	}

	return false
}

// SetDiscordEnabledTypes gets a reference to the given NullableFloat32 and assigns it to the DiscordEnabledTypes field.
func (o *UserSettingsNotifications) SetDiscordEnabledTypes(v float32) {
	o.DiscordEnabledTypes.Set(&v)
}
// SetDiscordEnabledTypesNil sets the value for DiscordEnabledTypes to be an explicit nil
func (o *UserSettingsNotifications) SetDiscordEnabledTypesNil() {
	o.DiscordEnabledTypes.Set(nil)
}

// UnsetDiscordEnabledTypes ensures that no value is present for DiscordEnabledTypes, not even an explicit nil
func (o *UserSettingsNotifications) UnsetDiscordEnabledTypes() {
	o.DiscordEnabledTypes.Unset()
}

// GetDiscordId returns the DiscordId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetDiscordId() string {
	if o == nil || IsNil(o.DiscordId.Get()) {
		var ret string
		return ret
	}
	return *o.DiscordId.Get()
}

// GetDiscordIdOk returns a tuple with the DiscordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetDiscordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscordId.Get(), o.DiscordId.IsSet()
}

// HasDiscordId returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasDiscordId() bool {
	if o != nil && o.DiscordId.IsSet() {
		return true
	}

	return false
}

// SetDiscordId gets a reference to the given NullableString and assigns it to the DiscordId field.
func (o *UserSettingsNotifications) SetDiscordId(v string) {
	o.DiscordId.Set(&v)
}
// SetDiscordIdNil sets the value for DiscordId to be an explicit nil
func (o *UserSettingsNotifications) SetDiscordIdNil() {
	o.DiscordId.Set(nil)
}

// UnsetDiscordId ensures that no value is present for DiscordId, not even an explicit nil
func (o *UserSettingsNotifications) UnsetDiscordId() {
	o.DiscordId.Unset()
}

// GetPushbulletAccessToken returns the PushbulletAccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetPushbulletAccessToken() string {
	if o == nil || IsNil(o.PushbulletAccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.PushbulletAccessToken.Get()
}

// GetPushbulletAccessTokenOk returns a tuple with the PushbulletAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetPushbulletAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushbulletAccessToken.Get(), o.PushbulletAccessToken.IsSet()
}

// HasPushbulletAccessToken returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasPushbulletAccessToken() bool {
	if o != nil && o.PushbulletAccessToken.IsSet() {
		return true
	}

	return false
}

// SetPushbulletAccessToken gets a reference to the given NullableString and assigns it to the PushbulletAccessToken field.
func (o *UserSettingsNotifications) SetPushbulletAccessToken(v string) {
	o.PushbulletAccessToken.Set(&v)
}
// SetPushbulletAccessTokenNil sets the value for PushbulletAccessToken to be an explicit nil
func (o *UserSettingsNotifications) SetPushbulletAccessTokenNil() {
	o.PushbulletAccessToken.Set(nil)
}

// UnsetPushbulletAccessToken ensures that no value is present for PushbulletAccessToken, not even an explicit nil
func (o *UserSettingsNotifications) UnsetPushbulletAccessToken() {
	o.PushbulletAccessToken.Unset()
}

// GetPushoverApplicationToken returns the PushoverApplicationToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetPushoverApplicationToken() string {
	if o == nil || IsNil(o.PushoverApplicationToken.Get()) {
		var ret string
		return ret
	}
	return *o.PushoverApplicationToken.Get()
}

// GetPushoverApplicationTokenOk returns a tuple with the PushoverApplicationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetPushoverApplicationTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushoverApplicationToken.Get(), o.PushoverApplicationToken.IsSet()
}

// HasPushoverApplicationToken returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasPushoverApplicationToken() bool {
	if o != nil && o.PushoverApplicationToken.IsSet() {
		return true
	}

	return false
}

// SetPushoverApplicationToken gets a reference to the given NullableString and assigns it to the PushoverApplicationToken field.
func (o *UserSettingsNotifications) SetPushoverApplicationToken(v string) {
	o.PushoverApplicationToken.Set(&v)
}
// SetPushoverApplicationTokenNil sets the value for PushoverApplicationToken to be an explicit nil
func (o *UserSettingsNotifications) SetPushoverApplicationTokenNil() {
	o.PushoverApplicationToken.Set(nil)
}

// UnsetPushoverApplicationToken ensures that no value is present for PushoverApplicationToken, not even an explicit nil
func (o *UserSettingsNotifications) UnsetPushoverApplicationToken() {
	o.PushoverApplicationToken.Unset()
}

// GetPushoverUserKey returns the PushoverUserKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetPushoverUserKey() string {
	if o == nil || IsNil(o.PushoverUserKey.Get()) {
		var ret string
		return ret
	}
	return *o.PushoverUserKey.Get()
}

// GetPushoverUserKeyOk returns a tuple with the PushoverUserKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetPushoverUserKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushoverUserKey.Get(), o.PushoverUserKey.IsSet()
}

// HasPushoverUserKey returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasPushoverUserKey() bool {
	if o != nil && o.PushoverUserKey.IsSet() {
		return true
	}

	return false
}

// SetPushoverUserKey gets a reference to the given NullableString and assigns it to the PushoverUserKey field.
func (o *UserSettingsNotifications) SetPushoverUserKey(v string) {
	o.PushoverUserKey.Set(&v)
}
// SetPushoverUserKeyNil sets the value for PushoverUserKey to be an explicit nil
func (o *UserSettingsNotifications) SetPushoverUserKeyNil() {
	o.PushoverUserKey.Set(nil)
}

// UnsetPushoverUserKey ensures that no value is present for PushoverUserKey, not even an explicit nil
func (o *UserSettingsNotifications) UnsetPushoverUserKey() {
	o.PushoverUserKey.Unset()
}

// GetPushoverSound returns the PushoverSound field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetPushoverSound() string {
	if o == nil || IsNil(o.PushoverSound.Get()) {
		var ret string
		return ret
	}
	return *o.PushoverSound.Get()
}

// GetPushoverSoundOk returns a tuple with the PushoverSound field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetPushoverSoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushoverSound.Get(), o.PushoverSound.IsSet()
}

// HasPushoverSound returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasPushoverSound() bool {
	if o != nil && o.PushoverSound.IsSet() {
		return true
	}

	return false
}

// SetPushoverSound gets a reference to the given NullableString and assigns it to the PushoverSound field.
func (o *UserSettingsNotifications) SetPushoverSound(v string) {
	o.PushoverSound.Set(&v)
}
// SetPushoverSoundNil sets the value for PushoverSound to be an explicit nil
func (o *UserSettingsNotifications) SetPushoverSoundNil() {
	o.PushoverSound.Set(nil)
}

// UnsetPushoverSound ensures that no value is present for PushoverSound, not even an explicit nil
func (o *UserSettingsNotifications) UnsetPushoverSound() {
	o.PushoverSound.Unset()
}

// GetTelegramEnabled returns the TelegramEnabled field value if set, zero value otherwise.
func (o *UserSettingsNotifications) GetTelegramEnabled() bool {
	if o == nil || IsNil(o.TelegramEnabled) {
		var ret bool
		return ret
	}
	return *o.TelegramEnabled
}

// GetTelegramEnabledOk returns a tuple with the TelegramEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSettingsNotifications) GetTelegramEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TelegramEnabled) {
		return nil, false
	}
	return o.TelegramEnabled, true
}

// HasTelegramEnabled returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasTelegramEnabled() bool {
	if o != nil && !IsNil(o.TelegramEnabled) {
		return true
	}

	return false
}

// SetTelegramEnabled gets a reference to the given bool and assigns it to the TelegramEnabled field.
func (o *UserSettingsNotifications) SetTelegramEnabled(v bool) {
	o.TelegramEnabled = &v
}

// GetTelegramBotUsername returns the TelegramBotUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetTelegramBotUsername() string {
	if o == nil || IsNil(o.TelegramBotUsername.Get()) {
		var ret string
		return ret
	}
	return *o.TelegramBotUsername.Get()
}

// GetTelegramBotUsernameOk returns a tuple with the TelegramBotUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetTelegramBotUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TelegramBotUsername.Get(), o.TelegramBotUsername.IsSet()
}

// HasTelegramBotUsername returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasTelegramBotUsername() bool {
	if o != nil && o.TelegramBotUsername.IsSet() {
		return true
	}

	return false
}

// SetTelegramBotUsername gets a reference to the given NullableString and assigns it to the TelegramBotUsername field.
func (o *UserSettingsNotifications) SetTelegramBotUsername(v string) {
	o.TelegramBotUsername.Set(&v)
}
// SetTelegramBotUsernameNil sets the value for TelegramBotUsername to be an explicit nil
func (o *UserSettingsNotifications) SetTelegramBotUsernameNil() {
	o.TelegramBotUsername.Set(nil)
}

// UnsetTelegramBotUsername ensures that no value is present for TelegramBotUsername, not even an explicit nil
func (o *UserSettingsNotifications) UnsetTelegramBotUsername() {
	o.TelegramBotUsername.Unset()
}

// GetTelegramChatId returns the TelegramChatId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetTelegramChatId() string {
	if o == nil || IsNil(o.TelegramChatId.Get()) {
		var ret string
		return ret
	}
	return *o.TelegramChatId.Get()
}

// GetTelegramChatIdOk returns a tuple with the TelegramChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetTelegramChatIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TelegramChatId.Get(), o.TelegramChatId.IsSet()
}

// HasTelegramChatId returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasTelegramChatId() bool {
	if o != nil && o.TelegramChatId.IsSet() {
		return true
	}

	return false
}

// SetTelegramChatId gets a reference to the given NullableString and assigns it to the TelegramChatId field.
func (o *UserSettingsNotifications) SetTelegramChatId(v string) {
	o.TelegramChatId.Set(&v)
}
// SetTelegramChatIdNil sets the value for TelegramChatId to be an explicit nil
func (o *UserSettingsNotifications) SetTelegramChatIdNil() {
	o.TelegramChatId.Set(nil)
}

// UnsetTelegramChatId ensures that no value is present for TelegramChatId, not even an explicit nil
func (o *UserSettingsNotifications) UnsetTelegramChatId() {
	o.TelegramChatId.Unset()
}

// GetTelegramSendSilently returns the TelegramSendSilently field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSettingsNotifications) GetTelegramSendSilently() bool {
	if o == nil || IsNil(o.TelegramSendSilently.Get()) {
		var ret bool
		return ret
	}
	return *o.TelegramSendSilently.Get()
}

// GetTelegramSendSilentlyOk returns a tuple with the TelegramSendSilently field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSettingsNotifications) GetTelegramSendSilentlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TelegramSendSilently.Get(), o.TelegramSendSilently.IsSet()
}

// HasTelegramSendSilently returns a boolean if a field has been set.
func (o *UserSettingsNotifications) HasTelegramSendSilently() bool {
	if o != nil && o.TelegramSendSilently.IsSet() {
		return true
	}

	return false
}

// SetTelegramSendSilently gets a reference to the given NullableBool and assigns it to the TelegramSendSilently field.
func (o *UserSettingsNotifications) SetTelegramSendSilently(v bool) {
	o.TelegramSendSilently.Set(&v)
}
// SetTelegramSendSilentlyNil sets the value for TelegramSendSilently to be an explicit nil
func (o *UserSettingsNotifications) SetTelegramSendSilentlyNil() {
	o.TelegramSendSilently.Set(nil)
}

// UnsetTelegramSendSilently ensures that no value is present for TelegramSendSilently, not even an explicit nil
func (o *UserSettingsNotifications) UnsetTelegramSendSilently() {
	o.TelegramSendSilently.Unset()
}

func (o UserSettingsNotifications) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSettingsNotifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationTypes) {
		toSerialize["notificationTypes"] = o.NotificationTypes
	}
	if !IsNil(o.EmailEnabled) {
		toSerialize["emailEnabled"] = o.EmailEnabled
	}
	if o.PgpKey.IsSet() {
		toSerialize["pgpKey"] = o.PgpKey.Get()
	}
	if !IsNil(o.DiscordEnabled) {
		toSerialize["discordEnabled"] = o.DiscordEnabled
	}
	if o.DiscordEnabledTypes.IsSet() {
		toSerialize["discordEnabledTypes"] = o.DiscordEnabledTypes.Get()
	}
	if o.DiscordId.IsSet() {
		toSerialize["discordId"] = o.DiscordId.Get()
	}
	if o.PushbulletAccessToken.IsSet() {
		toSerialize["pushbulletAccessToken"] = o.PushbulletAccessToken.Get()
	}
	if o.PushoverApplicationToken.IsSet() {
		toSerialize["pushoverApplicationToken"] = o.PushoverApplicationToken.Get()
	}
	if o.PushoverUserKey.IsSet() {
		toSerialize["pushoverUserKey"] = o.PushoverUserKey.Get()
	}
	if o.PushoverSound.IsSet() {
		toSerialize["pushoverSound"] = o.PushoverSound.Get()
	}
	if !IsNil(o.TelegramEnabled) {
		toSerialize["telegramEnabled"] = o.TelegramEnabled
	}
	if o.TelegramBotUsername.IsSet() {
		toSerialize["telegramBotUsername"] = o.TelegramBotUsername.Get()
	}
	if o.TelegramChatId.IsSet() {
		toSerialize["telegramChatId"] = o.TelegramChatId.Get()
	}
	if o.TelegramSendSilently.IsSet() {
		toSerialize["telegramSendSilently"] = o.TelegramSendSilently.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSettingsNotifications) UnmarshalJSON(data []byte) (err error) {
	varUserSettingsNotifications := _UserSettingsNotifications{}

	err = json.Unmarshal(data, &varUserSettingsNotifications)

	if err != nil {
		return err
	}

	*o = UserSettingsNotifications(varUserSettingsNotifications)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notificationTypes")
		delete(additionalProperties, "emailEnabled")
		delete(additionalProperties, "pgpKey")
		delete(additionalProperties, "discordEnabled")
		delete(additionalProperties, "discordEnabledTypes")
		delete(additionalProperties, "discordId")
		delete(additionalProperties, "pushbulletAccessToken")
		delete(additionalProperties, "pushoverApplicationToken")
		delete(additionalProperties, "pushoverUserKey")
		delete(additionalProperties, "pushoverSound")
		delete(additionalProperties, "telegramEnabled")
		delete(additionalProperties, "telegramBotUsername")
		delete(additionalProperties, "telegramChatId")
		delete(additionalProperties, "telegramSendSilently")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSettingsNotifications struct {
	value *UserSettingsNotifications
	isSet bool
}

func (v NullableUserSettingsNotifications) Get() *UserSettingsNotifications {
	return v.value
}

func (v *NullableUserSettingsNotifications) Set(val *UserSettingsNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSettingsNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSettingsNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSettingsNotifications(val *UserSettingsNotifications) *NullableUserSettingsNotifications {
	return &NullableUserSettingsNotifications{value: val, isSet: true}
}

func (v NullableUserSettingsNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSettingsNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


