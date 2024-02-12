# \ServiceAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceRadarrByRadarrId**](ServiceAPI.md#GetServiceRadarrByRadarrId) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
[**GetServiceSonarrBySonarrId**](ServiceAPI.md#GetServiceSonarrBySonarrId) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
[**ListServiceRadarr**](ServiceAPI.md#ListServiceRadarr) | **Get** /service/radarr | Get non-sensitive Radarr server list
[**ListServiceSonarr**](ServiceAPI.md#ListServiceSonarr) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
[**ListServiceSonarrLookupByTmdbId**](ServiceAPI.md#ListServiceSonarrLookupByTmdbId) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr



## GetServiceRadarrByRadarrId

> GetServiceRadarrByRadarrId2XXResponse GetServiceRadarrByRadarrId(ctx, radarrId).Execute()

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
    resp, r, err := apiClient.ServiceAPI.GetServiceRadarrByRadarrId(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetServiceRadarrByRadarrId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceRadarrByRadarrId`: GetServiceRadarrByRadarrId2XXResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetServiceRadarrByRadarrId`: %v\n", resp)
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

[**GetServiceRadarrByRadarrId2XXResponse**](GetServiceRadarrByRadarrId2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceSonarrBySonarrId

> GetServiceSonarrBySonarrId2XXResponse GetServiceSonarrBySonarrId(ctx, sonarrId).Execute()

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
    resp, r, err := apiClient.ServiceAPI.GetServiceSonarrBySonarrId(context.Background(), sonarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.GetServiceSonarrBySonarrId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceSonarrBySonarrId`: GetServiceSonarrBySonarrId2XXResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.GetServiceSonarrBySonarrId`: %v\n", resp)
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

[**GetServiceSonarrBySonarrId2XXResponse**](GetServiceSonarrBySonarrId2XXResponse.md)

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
    resp, r, err := apiClient.ServiceAPI.ListServiceRadarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ListServiceRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceRadarr`: []RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ListServiceRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.ServiceAPI.ListServiceSonarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ListServiceSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceSonarr`: []SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ListServiceSonarr`: %v\n", resp)
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
    resp, r, err := apiClient.ServiceAPI.ListServiceSonarrLookupByTmdbId(context.Background(), tmdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ListServiceSonarrLookupByTmdbId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceSonarrLookupByTmdbId`: []SonarrSeries
    fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ListServiceSonarrLookupByTmdbId`: %v\n", resp)
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

