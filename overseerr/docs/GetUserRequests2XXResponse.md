# GetUserRequests2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageInfo** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]MediaRequest**](MediaRequest.md) |  | [optional] 

## Methods

### NewGetUserRequests2XXResponse

`func NewGetUserRequests2XXResponse() *GetUserRequests2XXResponse`

NewGetUserRequests2XXResponse instantiates a new GetUserRequests2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRequests2XXResponseWithDefaults

`func NewGetUserRequests2XXResponseWithDefaults() *GetUserRequests2XXResponse`

NewGetUserRequests2XXResponseWithDefaults instantiates a new GetUserRequests2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageInfo

`func (o *GetUserRequests2XXResponse) GetPageInfo() PageInfo`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *GetUserRequests2XXResponse) GetPageInfoOk() (*PageInfo, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *GetUserRequests2XXResponse) SetPageInfo(v PageInfo)`

SetPageInfo sets PageInfo field to given value.

### HasPageInfo

`func (o *GetUserRequests2XXResponse) HasPageInfo() bool`

HasPageInfo returns a boolean if a field has been set.

### GetResults

`func (o *GetUserRequests2XXResponse) GetResults() []MediaRequest`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetUserRequests2XXResponse) GetResultsOk() (*[]MediaRequest, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetUserRequests2XXResponse) SetResults(v []MediaRequest)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetUserRequests2XXResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


