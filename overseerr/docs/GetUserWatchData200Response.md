# GetUserWatchData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentlyWatched** | Pointer to [**[]MediaInfo**](MediaInfo.md) |  | [optional] 
**PlayCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetUserWatchData200Response

`func NewGetUserWatchData200Response() *GetUserWatchData200Response`

NewGetUserWatchData200Response instantiates a new GetUserWatchData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserWatchData200ResponseWithDefaults

`func NewGetUserWatchData200ResponseWithDefaults() *GetUserWatchData200Response`

NewGetUserWatchData200ResponseWithDefaults instantiates a new GetUserWatchData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecentlyWatched

`func (o *GetUserWatchData200Response) GetRecentlyWatched() []MediaInfo`

GetRecentlyWatched returns the RecentlyWatched field if non-nil, zero value otherwise.

### GetRecentlyWatchedOk

`func (o *GetUserWatchData200Response) GetRecentlyWatchedOk() (*[]MediaInfo, bool)`

GetRecentlyWatchedOk returns a tuple with the RecentlyWatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyWatched

`func (o *GetUserWatchData200Response) SetRecentlyWatched(v []MediaInfo)`

SetRecentlyWatched sets RecentlyWatched field to given value.

### HasRecentlyWatched

`func (o *GetUserWatchData200Response) HasRecentlyWatched() bool`

HasRecentlyWatched returns a boolean if a field has been set.

### GetPlayCount

`func (o *GetUserWatchData200Response) GetPlayCount() float32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *GetUserWatchData200Response) GetPlayCountOk() (*float32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *GetUserWatchData200Response) SetPlayCount(v float32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *GetUserWatchData200Response) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


