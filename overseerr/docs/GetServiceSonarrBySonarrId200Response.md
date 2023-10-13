# GetServiceSonarrBySonarrId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**SonarrSettings**](SonarrSettings.md) |  | [optional] 
**Profiles** | Pointer to [**ServiceProfile**](ServiceProfile.md) |  | [optional] 

## Methods

### NewGetServiceSonarrBySonarrId200Response

`func NewGetServiceSonarrBySonarrId200Response() *GetServiceSonarrBySonarrId200Response`

NewGetServiceSonarrBySonarrId200Response instantiates a new GetServiceSonarrBySonarrId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceSonarrBySonarrId200ResponseWithDefaults

`func NewGetServiceSonarrBySonarrId200ResponseWithDefaults() *GetServiceSonarrBySonarrId200Response`

NewGetServiceSonarrBySonarrId200ResponseWithDefaults instantiates a new GetServiceSonarrBySonarrId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *GetServiceSonarrBySonarrId200Response) GetServer() SonarrSettings`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetServiceSonarrBySonarrId200Response) GetServerOk() (*SonarrSettings, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetServiceSonarrBySonarrId200Response) SetServer(v SonarrSettings)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetServiceSonarrBySonarrId200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetProfiles

`func (o *GetServiceSonarrBySonarrId200Response) GetProfiles() ServiceProfile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GetServiceSonarrBySonarrId200Response) GetProfilesOk() (*ServiceProfile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GetServiceSonarrBySonarrId200Response) SetProfiles(v ServiceProfile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GetServiceSonarrBySonarrId200Response) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


