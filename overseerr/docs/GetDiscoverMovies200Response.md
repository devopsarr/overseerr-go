# GetDiscoverMovies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**TotalPages** | Pointer to **float32** |  | [optional] 
**TotalResults** | Pointer to **float32** |  | [optional] 
**Results** | Pointer to [**[]MovieResult**](MovieResult.md) |  | [optional] 

## Methods

### NewGetDiscoverMovies200Response

`func NewGetDiscoverMovies200Response() *GetDiscoverMovies200Response`

NewGetDiscoverMovies200Response instantiates a new GetDiscoverMovies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoverMovies200ResponseWithDefaults

`func NewGetDiscoverMovies200ResponseWithDefaults() *GetDiscoverMovies200Response`

NewGetDiscoverMovies200ResponseWithDefaults instantiates a new GetDiscoverMovies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetDiscoverMovies200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetDiscoverMovies200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetDiscoverMovies200Response) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetDiscoverMovies200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetDiscoverMovies200Response) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetDiscoverMovies200Response) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetDiscoverMovies200Response) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetDiscoverMovies200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *GetDiscoverMovies200Response) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *GetDiscoverMovies200Response) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *GetDiscoverMovies200Response) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *GetDiscoverMovies200Response) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResults

`func (o *GetDiscoverMovies200Response) GetResults() []MovieResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetDiscoverMovies200Response) GetResultsOk() (*[]MovieResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetDiscoverMovies200Response) SetResults(v []MovieResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetDiscoverMovies200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


