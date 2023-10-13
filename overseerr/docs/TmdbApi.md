# \TmdbApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkByNetworkId**](TmdbApi.md#GetNetworkByNetworkId) | **Get** /network/{networkId} | Get TV network details
[**GetStudioByStudioId**](TmdbApi.md#GetStudioByStudioId) | **Get** /studio/{studioId} | Get movie studio details
[**ListBackdrops**](TmdbApi.md#ListBackdrops) | **Get** /backdrops | Get backdrops of trending items
[**ListGenresMovie**](TmdbApi.md#ListGenresMovie) | **Get** /genres/movie | Get list of official TMDB movie genres
[**ListGenresTv**](TmdbApi.md#ListGenresTv) | **Get** /genres/tv | Get list of official TMDB movie genres
[**ListLanguages**](TmdbApi.md#ListLanguages) | **Get** /languages | Languages supported by TMDB
[**ListRegions**](TmdbApi.md#ListRegions) | **Get** /regions | Regions supported by TMDB



## GetNetworkByNetworkId

> ProductionCompany GetNetworkByNetworkId(ctx, networkId).Execute()

Get TV network details



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
    networkId := float32(1) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TmdbApi.GetNetworkByNetworkId(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.GetNetworkByNetworkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkByNetworkId`: ProductionCompany
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.GetNetworkByNetworkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkByNetworkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioByStudioId

> ProductionCompany GetStudioByStudioId(ctx, studioId).Execute()

Get movie studio details



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
    studioId := float32(2) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TmdbApi.GetStudioByStudioId(context.Background(), studioId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.GetStudioByStudioId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStudioByStudioId`: ProductionCompany
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.GetStudioByStudioId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studioId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioByStudioIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackdrops

> []string ListBackdrops(ctx).Execute()

Get backdrops of trending items



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
    resp, r, err := apiClient.TmdbApi.ListBackdrops(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.ListBackdrops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackdrops`: []string
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.ListBackdrops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBackdropsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGenresMovie

> []ListGenresMovie200ResponseInner ListGenresMovie(ctx).Language(language).Execute()

Get list of official TMDB movie genres



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
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TmdbApi.ListGenresMovie(context.Background()).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.ListGenresMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGenresMovie`: []ListGenresMovie200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.ListGenresMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenresMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListGenresMovie200ResponseInner**](ListGenresMovie200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGenresTv

> []ListGenresTv200ResponseInner ListGenresTv(ctx).Language(language).Execute()

Get list of official TMDB movie genres



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
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.TmdbApi.ListGenresTv(context.Background()).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.ListGenresTv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGenresTv`: []ListGenresTv200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.ListGenresTv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenresTvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListGenresTv200ResponseInner**](ListGenresTv200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguages

> []ListLanguages200ResponseInner ListLanguages(ctx).Execute()

Languages supported by TMDB



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
    resp, r, err := apiClient.TmdbApi.ListLanguages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.ListLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLanguages`: []ListLanguages200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.ListLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLanguagesRequest struct via the builder pattern


### Return type

[**[]ListLanguages200ResponseInner**](ListLanguages200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegions

> []ListRegions200ResponseInner ListRegions(ctx).Execute()

Regions supported by TMDB



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
    resp, r, err := apiClient.TmdbApi.ListRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TmdbApi.ListRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegions`: []ListRegions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `TmdbApi.ListRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRegionsRequest struct via the builder pattern


### Return type

[**[]ListRegions200ResponseInner**](ListRegions200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

