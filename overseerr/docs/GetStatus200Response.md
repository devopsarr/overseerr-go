# GetStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**CommitTag** | Pointer to **string** |  | [optional] 
**UpdateAvailable** | Pointer to **bool** |  | [optional] 
**CommitsBehind** | Pointer to **float32** |  | [optional] 
**RestartRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetStatus200Response

`func NewGetStatus200Response() *GetStatus200Response`

NewGetStatus200Response instantiates a new GetStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus200ResponseWithDefaults

`func NewGetStatus200ResponseWithDefaults() *GetStatus200Response`

NewGetStatus200ResponseWithDefaults instantiates a new GetStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetStatus200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetStatus200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetStatus200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetStatus200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommitTag

`func (o *GetStatus200Response) GetCommitTag() string`

GetCommitTag returns the CommitTag field if non-nil, zero value otherwise.

### GetCommitTagOk

`func (o *GetStatus200Response) GetCommitTagOk() (*string, bool)`

GetCommitTagOk returns a tuple with the CommitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTag

`func (o *GetStatus200Response) SetCommitTag(v string)`

SetCommitTag sets CommitTag field to given value.

### HasCommitTag

`func (o *GetStatus200Response) HasCommitTag() bool`

HasCommitTag returns a boolean if a field has been set.

### GetUpdateAvailable

`func (o *GetStatus200Response) GetUpdateAvailable() bool`

GetUpdateAvailable returns the UpdateAvailable field if non-nil, zero value otherwise.

### GetUpdateAvailableOk

`func (o *GetStatus200Response) GetUpdateAvailableOk() (*bool, bool)`

GetUpdateAvailableOk returns a tuple with the UpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAvailable

`func (o *GetStatus200Response) SetUpdateAvailable(v bool)`

SetUpdateAvailable sets UpdateAvailable field to given value.

### HasUpdateAvailable

`func (o *GetStatus200Response) HasUpdateAvailable() bool`

HasUpdateAvailable returns a boolean if a field has been set.

### GetCommitsBehind

`func (o *GetStatus200Response) GetCommitsBehind() float32`

GetCommitsBehind returns the CommitsBehind field if non-nil, zero value otherwise.

### GetCommitsBehindOk

`func (o *GetStatus200Response) GetCommitsBehindOk() (*float32, bool)`

GetCommitsBehindOk returns a tuple with the CommitsBehind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitsBehind

`func (o *GetStatus200Response) SetCommitsBehind(v float32)`

SetCommitsBehind sets CommitsBehind field to given value.

### HasCommitsBehind

`func (o *GetStatus200Response) HasCommitsBehind() bool`

HasCommitsBehind returns a boolean if a field has been set.

### GetRestartRequired

`func (o *GetStatus200Response) GetRestartRequired() bool`

GetRestartRequired returns the RestartRequired field if non-nil, zero value otherwise.

### GetRestartRequiredOk

`func (o *GetStatus200Response) GetRestartRequiredOk() (*bool, bool)`

GetRestartRequiredOk returns a tuple with the RestartRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartRequired

`func (o *GetStatus200Response) SetRestartRequired(v bool)`

SetRestartRequired sets RestartRequired field to given value.

### HasRestartRequired

`func (o *GetStatus200Response) HasRestartRequired() bool`

HasRestartRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


