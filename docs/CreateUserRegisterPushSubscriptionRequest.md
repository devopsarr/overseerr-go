# CreateUserRegisterPushSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** |  | 
**Auth** | **string** |  | 
**P256dh** | **string** |  | 
**UserAgent** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserRegisterPushSubscriptionRequest

`func NewCreateUserRegisterPushSubscriptionRequest(endpoint string, auth string, p256dh string, ) *CreateUserRegisterPushSubscriptionRequest`

NewCreateUserRegisterPushSubscriptionRequest instantiates a new CreateUserRegisterPushSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRegisterPushSubscriptionRequestWithDefaults

`func NewCreateUserRegisterPushSubscriptionRequestWithDefaults() *CreateUserRegisterPushSubscriptionRequest`

NewCreateUserRegisterPushSubscriptionRequestWithDefaults instantiates a new CreateUserRegisterPushSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CreateUserRegisterPushSubscriptionRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateUserRegisterPushSubscriptionRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateUserRegisterPushSubscriptionRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAuth

`func (o *CreateUserRegisterPushSubscriptionRequest) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CreateUserRegisterPushSubscriptionRequest) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CreateUserRegisterPushSubscriptionRequest) SetAuth(v string)`

SetAuth sets Auth field to given value.


### GetP256dh

`func (o *CreateUserRegisterPushSubscriptionRequest) GetP256dh() string`

GetP256dh returns the P256dh field if non-nil, zero value otherwise.

### GetP256dhOk

`func (o *CreateUserRegisterPushSubscriptionRequest) GetP256dhOk() (*string, bool)`

GetP256dhOk returns a tuple with the P256dh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP256dh

`func (o *CreateUserRegisterPushSubscriptionRequest) SetP256dh(v string)`

SetP256dh sets P256dh field to given value.


### GetUserAgent

`func (o *CreateUserRegisterPushSubscriptionRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CreateUserRegisterPushSubscriptionRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CreateUserRegisterPushSubscriptionRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CreateUserRegisterPushSubscriptionRequest) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


