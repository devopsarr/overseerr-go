# TestSonarrRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** |  | 
**Port** | **float32** |  | 
**ApiKey** | **string** |  | 
**UseSsl** | **bool** |  | 
**BaseUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewTestSonarrRequest

`func NewTestSonarrRequest(hostname string, port float32, apiKey string, useSsl bool, ) *TestSonarrRequest`

NewTestSonarrRequest instantiates a new TestSonarrRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestSonarrRequestWithDefaults

`func NewTestSonarrRequestWithDefaults() *TestSonarrRequest`

NewTestSonarrRequestWithDefaults instantiates a new TestSonarrRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *TestSonarrRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TestSonarrRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TestSonarrRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *TestSonarrRequest) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TestSonarrRequest) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TestSonarrRequest) SetPort(v float32)`

SetPort sets Port field to given value.


### GetApiKey

`func (o *TestSonarrRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TestSonarrRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TestSonarrRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetUseSsl

`func (o *TestSonarrRequest) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *TestSonarrRequest) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *TestSonarrRequest) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.


### GetBaseUrl

`func (o *TestSonarrRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *TestSonarrRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *TestSonarrRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *TestSonarrRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


