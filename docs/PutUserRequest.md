# PutUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** |  | [optional] 
**Permissions** | Pointer to **int32** |  | [optional] 

## Methods

### NewPutUserRequest

`func NewPutUserRequest() *PutUserRequest`

NewPutUserRequest instantiates a new PutUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutUserRequestWithDefaults

`func NewPutUserRequestWithDefaults() *PutUserRequest`

NewPutUserRequestWithDefaults instantiates a new PutUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *PutUserRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *PutUserRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *PutUserRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *PutUserRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetPermissions

`func (o *PutUserRequest) GetPermissions() int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PutUserRequest) GetPermissionsOk() (*int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PutUserRequest) SetPermissions(v int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PutUserRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


