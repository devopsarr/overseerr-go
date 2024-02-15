# GetSearch2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **float32** |  | [optional] 
**TotalPages** | Pointer to **float32** |  | [optional] 
**TotalResults** | Pointer to **float32** |  | [optional] 
**Results** | Pointer to [**[]GetSearch2XXResponseResultsInner**](GetSearch2XXResponseResultsInner.md) |  | [optional] 

## Methods

### NewGetSearch2XXResponse

`func NewGetSearch2XXResponse() *GetSearch2XXResponse`

NewGetSearch2XXResponse instantiates a new GetSearch2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSearch2XXResponseWithDefaults

`func NewGetSearch2XXResponseWithDefaults() *GetSearch2XXResponse`

NewGetSearch2XXResponseWithDefaults instantiates a new GetSearch2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetSearch2XXResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSearch2XXResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSearch2XXResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetSearch2XXResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetSearch2XXResponse) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetSearch2XXResponse) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetSearch2XXResponse) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetSearch2XXResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *GetSearch2XXResponse) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *GetSearch2XXResponse) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *GetSearch2XXResponse) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *GetSearch2XXResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetResults

`func (o *GetSearch2XXResponse) GetResults() []GetSearch2XXResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSearch2XXResponse) GetResultsOk() (*[]GetSearch2XXResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSearch2XXResponse) SetResults(v []GetSearch2XXResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetSearch2XXResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


