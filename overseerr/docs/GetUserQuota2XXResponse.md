# GetUserQuota2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Movie** | Pointer to [**GetUserQuota2XXResponseMovie**](GetUserQuota2XXResponseMovie.md) |  | [optional] 
**Tv** | Pointer to [**GetUserQuota2XXResponseMovie**](GetUserQuota2XXResponseMovie.md) |  | [optional] 

## Methods

### NewGetUserQuota2XXResponse

`func NewGetUserQuota2XXResponse() *GetUserQuota2XXResponse`

NewGetUserQuota2XXResponse instantiates a new GetUserQuota2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserQuota2XXResponseWithDefaults

`func NewGetUserQuota2XXResponseWithDefaults() *GetUserQuota2XXResponse`

NewGetUserQuota2XXResponseWithDefaults instantiates a new GetUserQuota2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovie

`func (o *GetUserQuota2XXResponse) GetMovie() GetUserQuota2XXResponseMovie`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *GetUserQuota2XXResponse) GetMovieOk() (*GetUserQuota2XXResponseMovie, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *GetUserQuota2XXResponse) SetMovie(v GetUserQuota2XXResponseMovie)`

SetMovie sets Movie field to given value.

### HasMovie

`func (o *GetUserQuota2XXResponse) HasMovie() bool`

HasMovie returns a boolean if a field has been set.

### GetTv

`func (o *GetUserQuota2XXResponse) GetTv() GetUserQuota2XXResponseMovie`

GetTv returns the Tv field if non-nil, zero value otherwise.

### GetTvOk

`func (o *GetUserQuota2XXResponse) GetTvOk() (*GetUserQuota2XXResponseMovie, bool)`

GetTvOk returns a tuple with the Tv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTv

`func (o *GetUserQuota2XXResponse) SetTv(v GetUserQuota2XXResponseMovie)`

SetTv sets Tv field to given value.

### HasTv

`func (o *GetUserQuota2XXResponse) HasTv() bool`

HasTv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


