# \TmdbAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkByNetworkId**](TmdbAPI.md#GetNetworkByNetworkId) | **Get** /network/{networkId} | Get TV network details
[**GetStudioByStudioId**](TmdbAPI.md#GetStudioByStudioId) | **Get** /studio/{studioId} | Get movie studio details
[**ListBackdrops**](TmdbAPI.md#ListBackdrops) | **Get** /backdrops | Get backdrops of trending items
[**ListGenresMovie**](TmdbAPI.md#ListGenresMovie) | **Get** /genres/movie | Get list of official TMDB movie genres
[**ListGenresTv**](TmdbAPI.md#ListGenresTv) | **Get** /genres/tv | Get list of official TMDB movie genres
[**ListLanguages**](TmdbAPI.md#ListLanguages) | **Get** /languages | Languages supported by TMDB
[**ListRegions**](TmdbAPI.md#ListRegions) | **Get** /regions | Regions supported by TMDB



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
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {
	networkId := float32(1) // float32 | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.GetNetworkByNetworkId(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.GetNetworkByNetworkId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkByNetworkId`: ProductionCompany
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.GetNetworkByNetworkId`: %v\n", resp)
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
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {
	studioId := float32(2) // float32 | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.GetStudioByStudioId(context.Background(), studioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.GetStudioByStudioId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioByStudioId`: ProductionCompany
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.GetStudioByStudioId`: %v\n", resp)
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
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.ListBackdrops(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.ListBackdrops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackdrops`: []string
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.ListBackdrops`: %v\n", resp)
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

> []ListGenresMovie2XXResponseInner ListGenresMovie(ctx).Language(language).Execute()

Get list of official TMDB movie genres



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.ListGenresMovie(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.ListGenresMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenresMovie`: []ListGenresMovie2XXResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.ListGenresMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenresMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListGenresMovie2XXResponseInner**](ListGenresMovie2XXResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGenresTv

> []ListGenresTv2XXResponseInner ListGenresTv(ctx).Language(language).Execute()

Get list of official TMDB movie genres



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.ListGenresTv(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.ListGenresTv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGenresTv`: []ListGenresTv2XXResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.ListGenresTv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGenresTvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListGenresTv2XXResponseInner**](ListGenresTv2XXResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLanguages

> []ListLanguages2XXResponseInner ListLanguages(ctx).Execute()

Languages supported by TMDB



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.ListLanguages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.ListLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLanguages`: []ListLanguages2XXResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.ListLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLanguagesRequest struct via the builder pattern


### Return type

[**[]ListLanguages2XXResponseInner**](ListLanguages2XXResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegions

> []ListRegions2XXResponseInner ListRegions(ctx).Execute()

Regions supported by TMDB



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	overseerrClient "github.com/devopsarr/overseerr-go/overseerr"
)

func main() {

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.ListRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.ListRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRegions`: []ListRegions2XXResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.ListRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRegionsRequest struct via the builder pattern


### Return type

[**[]ListRegions2XXResponseInner**](ListRegions2XXResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

