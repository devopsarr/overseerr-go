# \ServiceApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceRadarrByRadarrId**](ServiceApi.md#GetServiceRadarrByRadarrId) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
[**GetServiceSonarrBySonarrId**](ServiceApi.md#GetServiceSonarrBySonarrId) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
[**ListServiceRadarr**](ServiceApi.md#ListServiceRadarr) | **Get** /service/radarr | Get non-sensitive Radarr server list
[**ListServiceSonarr**](ServiceApi.md#ListServiceSonarr) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
[**ListServiceSonarrLookupByTmdbId**](ServiceApi.md#ListServiceSonarrLookupByTmdbId) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr



## GetServiceRadarrByRadarrId

> GetServiceRadarrByRadarrId200Response GetServiceRadarrByRadarrId(ctx, radarrId).Execute()

Get Radarr server quality profiles and root folders



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
    radarrId := float32(0) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.GetServiceRadarrByRadarrId(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceRadarrByRadarrId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceRadarrByRadarrId`: GetServiceRadarrByRadarrId200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceRadarrByRadarrId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radarrId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRadarrByRadarrIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceRadarrByRadarrId200Response**](GetServiceRadarrByRadarrId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSonarrBySonarrId

> GetServiceSonarrBySonarrId200Response GetServiceSonarrBySonarrId(ctx, sonarrId).Execute()

Get Sonarr server quality profiles and root folders



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
    sonarrId := float32(0) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.GetServiceSonarrBySonarrId(context.Background(), sonarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.GetServiceSonarrBySonarrId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSonarrBySonarrId`: GetServiceSonarrBySonarrId200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.GetServiceSonarrBySonarrId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sonarrId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceSonarrBySonarrIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetServiceSonarrBySonarrId200Response**](GetServiceSonarrBySonarrId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceRadarr

> []RadarrSettings ListServiceRadarr(ctx).Execute()

Get non-sensitive Radarr server list



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
    resp, r, err := apiClient.ServiceApi.ListServiceRadarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.ListServiceRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceRadarr`: []RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.ListServiceRadarr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceRadarrRequest struct via the builder pattern


### Return type

[**[]RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceSonarr

> []SonarrSettings ListServiceSonarr(ctx).Execute()

Get non-sensitive Sonarr server list



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
    resp, r, err := apiClient.ServiceApi.ListServiceSonarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.ListServiceSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceSonarr`: []SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.ListServiceSonarr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceSonarrRequest struct via the builder pattern


### Return type

[**[]SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceSonarrLookupByTmdbId

> []SonarrSeries ListServiceSonarrLookupByTmdbId(ctx, tmdbId).Execute()

Get series from Sonarr



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
    tmdbId := float32(0) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceApi.ListServiceSonarrLookupByTmdbId(context.Background(), tmdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceApi.ListServiceSonarrLookupByTmdbId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceSonarrLookupByTmdbId`: []SonarrSeries
    fmt.Fprintf(os.Stdout, "Response from `ServiceApi.ListServiceSonarrLookupByTmdbId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmdbId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceSonarrLookupByTmdbIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SonarrSeries**](SonarrSeries.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

