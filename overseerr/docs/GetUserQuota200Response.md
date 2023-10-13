# GetUserQuota200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Movie** | Pointer to [**GetUserQuota200ResponseMovie**](GetUserQuota200ResponseMovie.md) |  | [optional] 
**Tv** | Pointer to [**GetUserQuota200ResponseMovie**](GetUserQuota200ResponseMovie.md) |  | [optional] 

## Methods

### NewGetUserQuota200Response

`func NewGetUserQuota200Response() *GetUserQuota200Response`

NewGetUserQuota200Response instantiates a new GetUserQuota200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserQuota200ResponseWithDefaults

`func NewGetUserQuota200ResponseWithDefaults() *GetUserQuota200Response`

NewGetUserQuota200ResponseWithDefaults instantiates a new GetUserQuota200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovie

`func (o *GetUserQuota200Response) GetMovie() GetUserQuota200ResponseMovie`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *GetUserQuota200Response) GetMovieOk() (*GetUserQuota200ResponseMovie, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *GetUserQuota200Response) SetMovie(v GetUserQuota200ResponseMovie)`

SetMovie sets Movie field to given value.

### HasMovie

`func (o *GetUserQuota200Response) HasMovie() bool`

HasMovie returns a boolean if a field has been set.

### GetTv

`func (o *GetUserQuota200Response) GetTv() GetUserQuota200ResponseMovie`

GetTv returns the Tv field if non-nil, zero value otherwise.

### GetTvOk

`func (o *GetUserQuota200Response) GetTvOk() (*GetUserQuota200ResponseMovie, bool)`

GetTvOk returns a tuple with the Tv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTv

`func (o *GetUserQuota200Response) SetTv(v GetUserQuota200ResponseMovie)`

SetTv sets Tv field to given value.

### HasTv

`func (o *GetUserQuota200Response) HasTv() bool`

HasTv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


