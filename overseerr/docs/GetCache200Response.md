# GetCache200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageCache** | Pointer to [**GetCache200ResponseImageCache**](GetCache200ResponseImageCache.md) |  | [optional] 
**ApiCaches** | Pointer to [**[]GetCache200ResponseApiCachesInner**](GetCache200ResponseApiCachesInner.md) |  | [optional] 

## Methods

### NewGetCache200Response

`func NewGetCache200Response() *GetCache200Response`

NewGetCache200Response instantiates a new GetCache200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCache200ResponseWithDefaults

`func NewGetCache200ResponseWithDefaults() *GetCache200Response`

NewGetCache200ResponseWithDefaults instantiates a new GetCache200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageCache

`func (o *GetCache200Response) GetImageCache() GetCache200ResponseImageCache`

GetImageCache returns the ImageCache field if non-nil, zero value otherwise.

### GetImageCacheOk

`func (o *GetCache200Response) GetImageCacheOk() (*GetCache200ResponseImageCache, bool)`

GetImageCacheOk returns a tuple with the ImageCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCache

`func (o *GetCache200Response) SetImageCache(v GetCache200ResponseImageCache)`

SetImageCache sets ImageCache field to given value.

### HasImageCache

`func (o *GetCache200Response) HasImageCache() bool`

HasImageCache returns a boolean if a field has been set.

### GetApiCaches

`func (o *GetCache200Response) GetApiCaches() []GetCache200ResponseApiCachesInner`

GetApiCaches returns the ApiCaches field if non-nil, zero value otherwise.

### GetApiCachesOk

`func (o *GetCache200Response) GetApiCachesOk() (*[]GetCache200ResponseApiCachesInner, bool)`

GetApiCachesOk returns a tuple with the ApiCaches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCaches

`func (o *GetCache200Response) SetApiCaches(v []GetCache200ResponseApiCachesInner)`

SetApiCaches sets ApiCaches field to given value.

### HasApiCaches

`func (o *GetCache200Response) HasApiCaches() bool`

HasApiCaches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


