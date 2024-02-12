# \MediaAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMediaByStatus**](MediaAPI.md#CreateMediaByStatus) | **Post** /media/{mediaId}/{status} | Update media status
[**DeleteMedia**](MediaAPI.md#DeleteMedia) | **Delete** /media/{mediaId} | Delete media item
[**GetMedia**](MediaAPI.md#GetMedia) | **Get** /media | Get media
[**GetMediaWatchData**](MediaAPI.md#GetMediaWatchData) | **Get** /media/{mediaId}/watch_data | Get watch data



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
    resp, r, err := apiClient.MediaAPI.CreateMediaByStatus(context.Background(), mediaId, status).CreateMediaByStatusRequest(createMediaByStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.CreateMediaByStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMediaByStatus`: MediaInfo
    fmt.Fprintf(os.Stdout, "Response from `MediaAPI.CreateMediaByStatus`: %v\n", resp)
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
    resp, r, err := apiClient.MediaAPI.DeleteMedia(context.Background(), mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.DeleteMedia``: %v\n", err)
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

> GetMedia2XXResponse GetMedia(ctx).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()

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
    resp, r, err := apiClient.MediaAPI.GetMedia(context.Background()).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.GetMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMedia`: GetMedia2XXResponse
    fmt.Fprintf(os.Stdout, "Response from `MediaAPI.GetMedia`: %v\n", resp)
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

[**GetMedia2XXResponse**](GetMedia2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaWatchData

> GetMediaWatchData2XXResponse GetMediaWatchData(ctx, mediaId).Execute()

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
    resp, r, err := apiClient.MediaAPI.GetMediaWatchData(context.Background(), mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.GetMediaWatchData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMediaWatchData`: GetMediaWatchData2XXResponse
    fmt.Fprintf(os.Stdout, "Response from `MediaAPI.GetMediaWatchData`: %v\n", resp)
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

[**GetMediaWatchData2XXResponse**](GetMediaWatchData2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

