# GetMedia2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]MediaInfo**](MediaInfo.md) |  | [optional] 

## Methods

### NewGetMedia2XXResponse

`func NewGetMedia2XXResponse() *GetMedia2XXResponse`

NewGetMedia2XXResponse instantiates a new GetMedia2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMedia2XXResponseWithDefaults

`func NewGetMedia2XXResponseWithDefaults() *GetMedia2XXResponse`

NewGetMedia2XXResponseWithDefaults instantiates a new GetMedia2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *GetMedia2XXResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *GetMedia2XXResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *GetMedia2XXResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *GetMedia2XXResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *GetMedia2XXResponse) GetResults() []MediaInfo`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMedia2XXResponse) GetResultsOk() (*[]MediaInfo, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMedia2XXResponse) SetResults(v []MediaInfo)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMedia2XXResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


