# \MediaApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMediaByStatus**](MediaApi.md#CreateMediaByStatus) | **Post** /media/{mediaId}/{status} | Update media status
[**DeleteMedia**](MediaApi.md#DeleteMedia) | **Delete** /media/{mediaId} | Delete media item
[**GetMedia**](MediaApi.md#GetMedia) | **Get** /media | Get media
[**GetMediaWatchData**](MediaApi.md#GetMediaWatchData) | **Get** /media/{mediaId}/watch_data | Get watch data



## CreateMediaByStatus

> MediaInfo CreateMediaByStatus(ctx, mediaId, status).CreateMediaByStatusRequest(createMediaByStatusRequest).Execute()

Update media status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    overseerrClient "./openapi"
)

func main() {
    mediaId := "1" // string | Media ID
    status := "available" // string | New status
    createMediaByStatusRequest := *overseerrClient.NewCreateMediaByStatusRequest() // CreateMediaByStatusRequest |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.CreateMediaByStatus(context.Background(), mediaId, status).CreateMediaByStatusRequest(createMediaByStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.CreateMediaByStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMediaByStatus`: MediaInfo
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.CreateMediaByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMediaByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMediaByStatusRequest** | [**CreateMediaByStatusRequest**](CreateMediaByStatusRequest.md) |  | 

### Return type

[**MediaInfo**](MediaInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMedia

> DeleteMedia(ctx, mediaId).Execute()

Delete media item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    overseerrClient "./openapi"
)

func main() {
    mediaId := "1" // string | Media ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.DeleteMedia(context.Background(), mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.DeleteMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMedia

> GetMedia200Response GetMedia(ctx).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()

Get media



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    overseerrClient "./openapi"
)

func main() {
    take := float32(20) // float32 |  (optional)
    skip := float32(0) // float32 |  (optional)
    filter := "filter_example" // string |  (optional)
    sort := "sort_example" // string |  (optional) (default to "added")

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.GetMedia(context.Background()).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.GetMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMedia`: GetMedia200Response
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.GetMedia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **filter** | **string** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]

### Return type

[**GetMedia200Response**](GetMedia200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaWatchData

> GetMediaWatchData200Response GetMediaWatchData(ctx, mediaId).Execute()

Get watch data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    overseerrClient "./openapi"
)

func main() {
    mediaId := "1" // string | Media ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.GetMediaWatchData(context.Background(), mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.GetMediaWatchData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaWatchData`: GetMediaWatchData200Response
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.GetMediaWatchData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaWatchDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMediaWatchData200Response**](GetMediaWatchData200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

