# GetDiscoverMoviesGenreByGenreId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**TotalPages** | Pointer to **float32** |  | [optional] 
**TotalResults** | Pointer to **float32** |  | [optional] 
**Genre** | Pointer to [**Genre**](Genre.md) |  | [optional] 
**Results** | Pointer to [**[]MovieResult**](MovieResult.md) |  | [optional] 

## Methods

### NewGetDiscoverMoviesGenreByGenreId200Response

`func NewGetDiscoverMoviesGenreByGenreId200Response() *GetDiscoverMoviesGenreByGenreId200Response`

NewGetDiscoverMoviesGenreByGenreId200Response instantiates a new GetDiscoverMoviesGenreByGenreId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoverMoviesGenreByGenreId200ResponseWithDefaults

`func NewGetDiscoverMoviesGenreByGenreId200ResponseWithDefaults() *GetDiscoverMoviesGenreByGenreId200Response`

NewGetDiscoverMoviesGenreByGenreId200ResponseWithDefaults instantiates a new GetDiscoverMoviesGenreByGenreId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetDiscoverMoviesGenreByGenreId200Response) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetDiscoverMoviesGenreByGenreId200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetDiscoverMoviesGenreByGenreId200Response) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetDiscoverMoviesGenreByGenreId200Response) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetGenre

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetGenre() Genre`

GetGenre returns the Genre field if non-nil, zero value otherwise.

### GetGenreOk

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetGenreOk() (*Genre, bool)`

GetGenreOk returns a tuple with the Genre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenre

`func (o *GetDiscoverMoviesGenreByGenreId200Response) SetGenre(v Genre)`

SetGenre sets Genre field to given value.

### HasGenre

`func (o *GetDiscoverMoviesGenreByGenreId200Response) HasGenre() bool`

HasGenre returns a boolean if a field has been set.

### GetResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetResults() []MovieResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetDiscoverMoviesGenreByGenreId200Response) GetResultsOk() (*[]MovieResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) SetResults(v []MovieResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetDiscoverMoviesGenreByGenreId200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


