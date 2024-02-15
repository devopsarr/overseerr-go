# GetServiceRadarrByRadarrId2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**RadarrSettings**](RadarrSettings.md) |  | [optional] 
**Profiles** | Pointer to [**ServiceProfile**](ServiceProfile.md) |  | [optional] 

## Methods

### NewGetServiceRadarrByRadarrId2XXResponse

`func NewGetServiceRadarrByRadarrId2XXResponse() *GetServiceRadarrByRadarrId2XXResponse`

NewGetServiceRadarrByRadarrId2XXResponse instantiates a new GetServiceRadarrByRadarrId2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceRadarrByRadarrId2XXResponseWithDefaults

`func NewGetServiceRadarrByRadarrId2XXResponseWithDefaults() *GetServiceRadarrByRadarrId2XXResponse`

NewGetServiceRadarrByRadarrId2XXResponseWithDefaults instantiates a new GetServiceRadarrByRadarrId2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *GetServiceRadarrByRadarrId2XXResponse) GetServer() RadarrSettings`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetServiceRadarrByRadarrId2XXResponse) GetServerOk() (*RadarrSettings, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetServiceRadarrByRadarrId2XXResponse) SetServer(v RadarrSettings)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetServiceRadarrByRadarrId2XXResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetProfiles

`func (o *GetServiceRadarrByRadarrId2XXResponse) GetProfiles() ServiceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetServiceRadarrByRadarrId2XXResponse) GetProfilesOk() (*ServiceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetServiceRadarrByRadarrId2XXResponse) SetProfiles(v ServiceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetServiceRadarrByRadarrId2XXResponse) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


