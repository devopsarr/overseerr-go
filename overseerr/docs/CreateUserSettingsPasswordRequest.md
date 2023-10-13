# CreateUserSettingsPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **NullableString** |  | [optional] 
**NewPassword** | **string** |  | 

## Methods

### NewCreateUserSettingsPasswordRequest

`func NewCreateUserSettingsPasswordRequest(newPassword string, ) *CreateUserSettingsPasswordRequest`

NewCreateUserSettingsPasswordRequest instantiates a new CreateUserSettingsPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSettingsPasswordRequestWithDefaults

`func NewCreateUserSettingsPasswordRequestWithDefaults() *CreateUserSettingsPasswordRequest`

NewCreateUserSettingsPasswordRequestWithDefaults instantiates a new CreateUserSettingsPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *CreateUserSettingsPasswordRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *CreateUserSettingsPasswordRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *CreateUserSettingsPasswordRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *CreateUserSettingsPasswordRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *CreateUserSettingsPasswordRequest) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *CreateUserSettingsPasswordRequest) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetNewPassword

`func (o *CreateUserSettingsPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *CreateUserSettingsPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *CreateUserSettingsPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


