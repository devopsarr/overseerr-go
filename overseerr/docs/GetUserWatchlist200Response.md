# GetUserWatchlist200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**TotalPages** | Pointer to **float32** |  | [optional] 
**TotalResults** | Pointer to **float32** |  | [optional] 
**Results** | Pointer to [**[]GetUserWatchlist200ResponseResultsInner**](GetUserWatchlist200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewGetUserWatchlist200Response

`func NewGetUserWatchlist200Response() *GetUserWatchlist200Response`

NewGetUserWatchlist200Response instantiates a new GetUserWatchlist200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserWatchlist200ResponseWithDefaults

`func NewGetUserWatchlist200ResponseWithDefaults() *GetUserWatchlist200Response`

NewGetUserWatchlist200ResponseWithDefaults instantiates a new GetUserWatchlist200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetUserWatchlist200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetUserWatchlist200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetUserWatchlist200Response) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetUserWatchlist200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetUserWatchlist200Response) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetUserWatchlist200Response) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetUserWatchlist200Response) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetUserWatchlist200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *GetUserWatchlist200Response) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *GetUserWatchlist200Response) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *GetUserWatchlist200Response) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *GetUserWatchlist200Response) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResults

`func (o *GetUserWatchlist200Response) GetResults() []GetUserWatchlist200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetUserWatchlist200Response) GetResultsOk() (*[]GetUserWatchlist200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetUserWatchlist200Response) SetResults(v []GetUserWatchlist200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetUserWatchlist200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


