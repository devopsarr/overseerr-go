# GetPlexSync2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to **bool** |  | [optional] 
**Progress** | Pointer to **float32** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**CurrentLibrary** | Pointer to [**PlexLibrary**](PlexLibrary.md) |  | [optional] 
**Libraries** | Pointer to [**[]PlexLibrary**](PlexLibrary.md) |  | [optional] 

## Methods

### NewGetPlexSync2XXResponse

`func NewGetPlexSync2XXResponse() *GetPlexSync2XXResponse`

NewGetPlexSync2XXResponse instantiates a new GetPlexSync2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlexSync2XXResponseWithDefaults

`func NewGetPlexSync2XXResponseWithDefaults() *GetPlexSync2XXResponse`

NewGetPlexSync2XXResponseWithDefaults instantiates a new GetPlexSync2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *GetPlexSync2XXResponse) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *GetPlexSync2XXResponse) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *GetPlexSync2XXResponse) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *GetPlexSync2XXResponse) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetProgress

`func (o *GetPlexSync2XXResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *GetPlexSync2XXResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *GetPlexSync2XXResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *GetPlexSync2XXResponse) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetTotal

`func (o *GetPlexSync2XXResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPlexSync2XXResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPlexSync2XXResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetPlexSync2XXResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetCurrentLibrary

`func (o *GetPlexSync2XXResponse) GetCurrentLibrary() PlexLibrary`

GetCurrentLibrary returns the CurrentLibrary field if non-nil, zero value otherwise.

### GetCurrentLibraryOk

`func (o *GetPlexSync2XXResponse) GetCurrentLibraryOk() (*PlexLibrary, bool)`

GetCurrentLibraryOk returns a tuple with the CurrentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLibrary

`func (o *GetPlexSync2XXResponse) SetCurrentLibrary(v PlexLibrary)`

SetCurrentLibrary sets CurrentLibrary field to given value.

### HasCurrentLibrary

`func (o *GetPlexSync2XXResponse) HasCurrentLibrary() bool`

HasCurrentLibrary returns a boolean if a field has been set.

### GetLibraries

`func (o *GetPlexSync2XXResponse) GetLibraries() []PlexLibrary`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *GetPlexSync2XXResponse) GetLibrariesOk() (*[]PlexLibrary, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *GetPlexSync2XXResponse) SetLibraries(v []PlexLibrary)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *GetPlexSync2XXResponse) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


