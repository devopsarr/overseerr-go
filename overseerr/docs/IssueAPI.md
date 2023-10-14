# \IssueAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssue**](IssueAPI.md#CreateIssue) | **Post** /issue | Create new issue
[**CreateIssueByStatus**](IssueAPI.md#CreateIssueByStatus) | **Post** /issue/{issueId}/{status} | Update an issue&#39;s status
[**CreateIssueComment**](IssueAPI.md#CreateIssueComment) | **Post** /issue/{issueId}/comment | Create a comment
[**DeleteIssue**](IssueAPI.md#DeleteIssue) | **Delete** /issue/{issueId} | Delete issue
[**DeleteIssueComment**](IssueAPI.md#DeleteIssueComment) | **Delete** /issueComment/{commentId} | Delete issue comment
[**GetIssue**](IssueAPI.md#GetIssue) | **Get** /issue | Get all issues
[**GetIssueByIssueId**](IssueAPI.md#GetIssueByIssueId) | **Get** /issue/{issueId} | Get issue
[**GetIssueCommentByCommentId**](IssueAPI.md#GetIssueCommentByCommentId) | **Get** /issueComment/{commentId} | Get issue comment
[**GetIssueCount**](IssueAPI.md#GetIssueCount) | **Get** /issue/count | Gets issue counts
[**UpdateIssueComment**](IssueAPI.md#UpdateIssueComment) | **Put** /issueComment/{commentId} | Update issue comment



## CreateIssue

> Issue CreateIssue(ctx).CreateIssueRequest(createIssueRequest).Execute()

Create new issue



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
    createIssueRequest := *overseerrClient.NewCreateIssueRequest() // CreateIssueRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.CreateIssue(context.Background()).CreateIssueRequest(createIssueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.CreateIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIssue`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.CreateIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIssueRequest** | [**CreateIssueRequest**](CreateIssueRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueByStatus

> Issue CreateIssueByStatus(ctx, issueId, status).Execute()

Update an issue's status



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
    issueId := "1" // string | Issue ID
    status := "status_example" // string | New status

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.CreateIssueByStatus(context.Background(), issueId, status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.CreateIssueByStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIssueByStatus`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.CreateIssueByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **string** | Issue ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueComment

> Issue CreateIssueComment(ctx, issueId).CreateIssueCommentRequest(createIssueCommentRequest).Execute()

Create a comment



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
    issueId := float32(1) // float32 | 
    createIssueCommentRequest := *overseerrClient.NewCreateIssueCommentRequest("Message_example") // CreateIssueCommentRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.CreateIssueComment(context.Background(), issueId).CreateIssueCommentRequest(createIssueCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.CreateIssueComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIssueComment`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.CreateIssueComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIssueCommentRequest** | [**CreateIssueCommentRequest**](CreateIssueCommentRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssue

> DeleteIssue(ctx, issueId).Execute()

Delete issue



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
    issueId := "1" // string | Issue ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.DeleteIssue(context.Background(), issueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.DeleteIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **string** | Issue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueRequest struct via the builder pattern


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


## DeleteIssueComment

> DeleteIssueComment(ctx, commentId).Execute()

Delete issue comment



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
    commentId := "1" // string | Issue Comment ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.DeleteIssueComment(context.Background(), commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.DeleteIssueComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | Issue Comment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueCommentRequest struct via the builder pattern


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


## GetIssue

> GetIssue200Response GetIssue(ctx).Take(take).Skip(skip).Sort(sort).Filter(filter).RequestedBy(requestedBy).Execute()

Get all issues



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
    sort := "sort_example" // string |  (optional) (default to "added")
    filter := "filter_example" // string |  (optional) (default to "open")
    requestedBy := float32(1) // float32 |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.GetIssue(context.Background()).Take(take).Skip(skip).Sort(sort).Filter(filter).RequestedBy(requestedBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssue`: GetIssue200Response
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.GetIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]
 **filter** | **string** |  | [default to &quot;open&quot;]
 **requestedBy** | **float32** |  | 

### Return type

[**GetIssue200Response**](GetIssue200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueByIssueId

> Issue GetIssueByIssueId(ctx, issueId).Execute()

Get issue



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
    issueId := float32(1) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.GetIssueByIssueId(context.Background(), issueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssueByIssueId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssueByIssueId`: Issue
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.GetIssueByIssueId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueByIssueIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueCommentByCommentId

> IssueComment GetIssueCommentByCommentId(ctx, commentId).Execute()

Get issue comment



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
    commentId := "1" // string | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.GetIssueCommentByCommentId(context.Background(), commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssueCommentByCommentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssueCommentByCommentId`: IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.GetIssueCommentByCommentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueCommentByCommentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueCount

> GetIssueCount200Response GetIssueCount(ctx).Execute()

Gets issue counts



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
    resp, r, err := apiClient.IssueAPI.GetIssueCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssueCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIssueCount`: GetIssueCount200Response
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.GetIssueCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueCountRequest struct via the builder pattern


### Return type

[**GetIssueCount200Response**](GetIssueCount200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssueComment

> IssueComment UpdateIssueComment(ctx, commentId).UpdateIssueCommentRequest(updateIssueCommentRequest).Execute()

Update issue comment



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
    commentId := "1" // string | 
    updateIssueCommentRequest := *overseerrClient.NewUpdateIssueCommentRequest() // UpdateIssueCommentRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssueAPI.UpdateIssueComment(context.Background(), commentId).UpdateIssueCommentRequest(updateIssueCommentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.UpdateIssueComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIssueComment`: IssueComment
    fmt.Fprintf(os.Stdout, "Response from `IssueAPI.UpdateIssueComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIssueCommentRequest** | [**UpdateIssueCommentRequest**](UpdateIssueCommentRequest.md) |  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

