# CreateRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaType** | **string** |  | 
**MediaId** | **float32** |  | 
**TvdbId** | Pointer to **float32** |  | [optional] 
**Seasons** | Pointer to [**CreateRequestRequestSeasons**](CreateRequestRequestSeasons.md) |  | [optional] 
**Is4k** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **float32** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**RootFolder** | Pointer to **string** |  | [optional] 
**LanguageProfileId** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewCreateRequestRequest

`func NewCreateRequestRequest(mediaType string, mediaId float32, ) *CreateRequestRequest`

NewCreateRequestRequest instantiates a new CreateRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestRequestWithDefaults

`func NewCreateRequestRequestWithDefaults() *CreateRequestRequest`

NewCreateRequestRequestWithDefaults instantiates a new CreateRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaType

`func (o *CreateRequestRequest) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreateRequestRequest) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreateRequestRequest) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMediaId

`func (o *CreateRequestRequest) GetMediaId() float32`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *CreateRequestRequest) GetMediaIdOk() (*float32, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *CreateRequestRequest) SetMediaId(v float32)`

SetMediaId sets MediaId field to given value.


### GetTvdbId

`func (o *CreateRequestRequest) GetTvdbId() float32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *CreateRequestRequest) GetTvdbIdOk() (*float32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *CreateRequestRequest) SetTvdbId(v float32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *CreateRequestRequest) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### GetSeasons

`func (o *CreateRequestRequest) GetSeasons() CreateRequestRequestSeasons`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *CreateRequestRequest) GetSeasonsOk() (*CreateRequestRequestSeasons, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *CreateRequestRequest) SetSeasons(v CreateRequestRequestSeasons)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *CreateRequestRequest) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetIs4k

`func (o *CreateRequestRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *CreateRequestRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *CreateRequestRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *CreateRequestRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.

### GetServerId

`func (o *CreateRequestRequest) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *CreateRequestRequest) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *CreateRequestRequest) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *CreateRequestRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetProfileId

`func (o *CreateRequestRequest) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateRequestRequest) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateRequestRequest) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *CreateRequestRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRootFolder

`func (o *CreateRequestRequest) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *CreateRequestRequest) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *CreateRequestRequest) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *CreateRequestRequest) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.

### GetLanguageProfileId

`func (o *CreateRequestRequest) GetLanguageProfileId() float32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *CreateRequestRequest) GetLanguageProfileIdOk() (*float32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *CreateRequestRequest) SetLanguageProfileId(v float32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *CreateRequestRequest) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateRequestRequest) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateRequestRequest) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateRequestRequest) SetUserId(v float32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateRequestRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateRequestRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateRequestRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


