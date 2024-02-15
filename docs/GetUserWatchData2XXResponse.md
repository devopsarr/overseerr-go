# GetUserWatchData2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentlyWatched** | Pointer to [**[]MediaInfo**](MediaInfo.md) |  | [optional] 
**PlayCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetUserWatchData2XXResponse

`func NewGetUserWatchData2XXResponse() *GetUserWatchData2XXResponse`

NewGetUserWatchData2XXResponse instantiates a new GetUserWatchData2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserWatchData2XXResponseWithDefaults

`func NewGetUserWatchData2XXResponseWithDefaults() *GetUserWatchData2XXResponse`

NewGetUserWatchData2XXResponseWithDefaults instantiates a new GetUserWatchData2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecentlyWatched

`func (o *GetUserWatchData2XXResponse) GetRecentlyWatched() []MediaInfo`

GetRecentlyWatched returns the RecentlyWatched field if non-nil, zero value otherwise.

### GetRecentlyWatchedOk

`func (o *GetUserWatchData2XXResponse) GetRecentlyWatchedOk() (*[]MediaInfo, bool)`

GetRecentlyWatchedOk returns a tuple with the RecentlyWatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyWatched

`func (o *GetUserWatchData2XXResponse) SetRecentlyWatched(v []MediaInfo)`

SetRecentlyWatched sets RecentlyWatched field to given value.

### HasRecentlyWatched

`func (o *GetUserWatchData2XXResponse) HasRecentlyWatched() bool`

HasRecentlyWatched returns a boolean if a field has been set.

### GetPlayCount

`func (o *GetUserWatchData2XXResponse) GetPlayCount() float32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *GetUserWatchData2XXResponse) GetPlayCountOk() (*float32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *GetUserWatchData2XXResponse) SetPlayCount(v float32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *GetUserWatchData2XXResponse) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


