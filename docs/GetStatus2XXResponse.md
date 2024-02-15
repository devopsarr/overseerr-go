# GetStatus2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**CommitTag** | Pointer to **string** |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] 
**CommitsBehind** | Pointer to **float32** |  | [optional] 
**RestartRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetStatus2XXResponse

`func NewGetStatus2XXResponse() *GetStatus2XXResponse`

NewGetStatus2XXResponse instantiates a new GetStatus2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus2XXResponseWithDefaults

`func NewGetStatus2XXResponseWithDefaults() *GetStatus2XXResponse`

NewGetStatus2XXResponseWithDefaults instantiates a new GetStatus2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetStatus2XXResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetStatus2XXResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetStatus2XXResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetStatus2XXResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommitTag

`func (o *GetStatus2XXResponse) GetCommitTag() string`

GetCommitTag returns the CommitTag field if non-nil, zero value otherwise.

### GetCommitTagOk

`func (o *GetStatus2XXResponse) GetCommitTagOk() (*string, bool)`

GetCommitTagOk returns a tuple with the CommitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTag

`func (o *GetStatus2XXResponse) SetCommitTag(v string)`

SetCommitTag sets CommitTag field to given value.

### HasCommitTag

`func (o *GetStatus2XXResponse) HasCommitTag() bool`

HasCommitTag returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *GetStatus2XXResponse) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *GetStatus2XXResponse) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *GetStatus2XXResponse) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *GetStatus2XXResponse) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetCommitsBehind

`func (o *GetStatus2XXResponse) GetCommitsBehind() float32`

GetCommitsBehind returns the CommitsBehind field if non-nil, zero value otherwise.

### GetCommitsBehindOk

`func (o *GetStatus2XXResponse) GetCommitsBehindOk() (*float32, bool)`

GetCommitsBehindOk returns a tuple with the CommitsBehind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsBehind

`func (o *GetStatus2XXResponse) SetCommitsBehind(v float32)`

SetCommitsBehind sets CommitsBehind field to given value.

### HasCommitsBehind

`func (o *GetStatus2XXResponse) HasCommitsBehind() bool`

HasCommitsBehind returns a boolean if a field has been set.

### GetRestartRequired

`func (o *GetStatus2XXResponse) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *GetStatus2XXResponse) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *GetStatus2XXResponse) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *GetStatus2XXResponse) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


