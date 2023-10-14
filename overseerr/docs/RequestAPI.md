# \RequestAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequest**](RequestAPI.md#CreateRequest) | **Post** /request | Create new request
[**CreateRequestByStatus**](RequestAPI.md#CreateRequestByStatus) | **Post** /request/{requestId}/{status} | Update a request&#39;s status
[**CreateRequestRetry**](RequestAPI.md#CreateRequestRetry) | **Post** /request/{requestId}/retry | Retry failed request
[**DeleteRequest**](RequestAPI.md#DeleteRequest) | **Delete** /request/{requestId} | Delete request
[**GetRequest**](RequestAPI.md#GetRequest) | **Get** /request | Get all requests
[**GetRequestByRequestId**](RequestAPI.md#GetRequestByRequestId) | **Get** /request/{requestId} | Get MediaRequest
[**GetRequestCount**](RequestAPI.md#GetRequestCount) | **Get** /request/count | Gets request counts
[**UpdateRequest**](RequestAPI.md#UpdateRequest) | **Put** /request/{requestId} | Update MediaRequest



## CreateRequest

> MediaRequest CreateRequest(ctx).CreateRequestRequest(createRequestRequest).Execute()

Create new request



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
    createRequestRequest := *overseerrClient.NewCreateRequestRequest("movie", float32(123)) // CreateRequestRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.CreateRequest(context.Background()).CreateRequestRequest(createRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.CreateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequest`: MediaRequest
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRequestRequest** | [**CreateRequestRequest**](CreateRequestRequest.md) |  | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequestByStatus

> MediaRequest CreateRequestByStatus(ctx, requestId, status).Execute()

Update a request's status



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
    requestId := "1" // string | Request ID
    status := "status_example" // string | New status

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.CreateRequestByStatus(context.Background(), requestId, status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.CreateRequestByStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestByStatus`: MediaRequest
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.CreateRequestByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequestRetry

> MediaRequest CreateRequestRetry(ctx, requestId).Execute()

Retry failed request



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
    requestId := "1" // string | Request ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.CreateRequestRetry(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.CreateRequestRetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestRetry`: MediaRequest
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.CreateRequestRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequest

> DeleteRequest(ctx, requestId).Execute()

Delete request



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
    requestId := "1" // string | Request ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.DeleteRequest(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.DeleteRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestRequest struct via the builder pattern


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


## GetRequest

> GetUserRequests200Response GetRequest(ctx).Take(take).Skip(skip).Filter(filter).Sort(sort).RequestedBy(requestedBy).Execute()

Get all requests



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
    requestedBy := float32(1) // float32 |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.GetRequest(context.Background()).Take(take).Skip(skip).Filter(filter).Sort(sort).RequestedBy(requestedBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.GetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequest`: GetUserRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.GetRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **filter** | **string** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]
 **requestedBy** | **float32** |  | 

### Return type

[**GetUserRequests200Response**](GetUserRequests200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestByRequestId

> MediaRequest GetRequestByRequestId(ctx, requestId).Execute()

Get MediaRequest



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
    requestId := "1" // string | Request ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.GetRequestByRequestId(context.Background(), requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.GetRequestByRequestId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestByRequestId`: MediaRequest
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.GetRequestByRequestId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestByRequestIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestCount

> GetRequestCount200Response GetRequestCount(ctx).Execute()

Gets request counts



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

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.GetRequestCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.GetRequestCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestCount`: GetRequestCount200Response
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.GetRequestCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestCountRequest struct via the builder pattern


### Return type

[**GetRequestCount200Response**](GetRequestCount200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRequest

> MediaRequest UpdateRequest(ctx, requestId).UpdateRequestRequest(updateRequestRequest).Execute()

Update MediaRequest



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
    requestId := "1" // string | Request ID
    updateRequestRequest := *overseerrClient.NewUpdateRequestRequest("MediaType_example") // UpdateRequestRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestAPI.UpdateRequest(context.Background(), requestId).UpdateRequestRequest(updateRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.UpdateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequest`: MediaRequest
    fmt.Fprintf(os.Stdout, "Response from `RequestAPI.UpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRequestRequest** | [**UpdateRequestRequest**](UpdateRequestRequest.md) |  | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

