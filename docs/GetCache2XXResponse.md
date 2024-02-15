# GetCache2XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageCache** | Pointer to [**GetCache2XXResponseImageCache**](GetCache2XXResponseImageCache.md) |  | [optional] 
**ApiCaches** | Pointer to [**[]GetCache2XXResponseApiCachesInner**](GetCache2XXResponseApiCachesInner.md) |  | [optional] 

## Methods

### NewGetCache2XXResponse

`func NewGetCache2XXResponse() *GetCache2XXResponse`

NewGetCache2XXResponse instantiates a new GetCache2XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCache2XXResponseWithDefaults

`func NewGetCache2XXResponseWithDefaults() *GetCache2XXResponse`

NewGetCache2XXResponseWithDefaults instantiates a new GetCache2XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageCache

`func (o *GetCache2XXResponse) GetImageCache() GetCache2XXResponseImageCache`

GetImageCache returns the ImageCache field if non-nil, zero value otherwise.

### GetImageCacheOk

`func (o *GetCache2XXResponse) GetImageCacheOk() (*GetCache2XXResponseImageCache, bool)`

GetImageCacheOk returns a tuple with the ImageCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCache

`func (o *GetCache2XXResponse) SetImageCache(v GetCache2XXResponseImageCache)`

SetImageCache sets ImageCache field to given value.

### HasImageCache

`func (o *GetCache2XXResponse) HasImageCache() bool`

HasImageCache returns a boolean if a field has been set.

### GetApiCaches

`func (o *GetCache2XXResponse) GetApiCaches() []GetCache2XXResponseApiCachesInner`

GetApiCaches returns the ApiCaches field if non-nil, zero value otherwise.

### GetApiCachesOk

`func (o *GetCache2XXResponse) GetApiCachesOk() (*[]GetCache2XXResponseApiCachesInner, bool)`

GetApiCachesOk returns a tuple with the ApiCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCaches

`func (o *GetCache2XXResponse) SetApiCaches(v []GetCache2XXResponseApiCachesInner)`

SetApiCaches sets ApiCaches field to given value.

### HasApiCaches

`func (o *GetCache2XXResponse) HasApiCaches() bool`

HasApiCaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


