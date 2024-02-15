# \TvAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTvByTvId**](TvAPI.md#GetTvByTvId) | **Get** /tv/{tvId} | Get TV details
[**GetTvRatings**](TvAPI.md#GetTvRatings) | **Get** /tv/{tvId}/ratings | Get TV ratings
[**GetTvRecommendations**](TvAPI.md#GetTvRecommendations) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
[**GetTvSeasonBySeasonId**](TvAPI.md#GetTvSeasonBySeasonId) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
[**GetTvSimilar**](TvAPI.md#GetTvSimilar) | **Get** /tv/{tvId}/similar | Get similar TV series



## GetTvByTvId

> TvDetails GetTvByTvId(ctx, tvId).Language(language).Execute()

Get TV details



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
	tvId := float32(76479) // float32 | 
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.GetTvByTvId(context.Background(), tvId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.GetTvByTvId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTvByTvId`: TvDetails
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.GetTvByTvId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tvId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTvByTvIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**TvDetails**](TvDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTvRatings

> GetTvRatings2XXResponse GetTvRatings(ctx, tvId).Execute()

Get TV ratings



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
	tvId := float32(76479) // float32 | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.GetTvRatings(context.Background(), tvId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.GetTvRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTvRatings`: GetTvRatings2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.GetTvRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tvId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTvRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTvRatings2XXResponse**](GetTvRatings2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTvRecommendations

> GetDiscoverTv2XXResponse GetTvRecommendations(ctx, tvId).Page(page).Language(language).Execute()

Get recommended TV series



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
	tvId := float32(76479) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.GetTvRecommendations(context.Background(), tvId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.GetTvRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTvRecommendations`: GetDiscoverTv2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.GetTvRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tvId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTvRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverTv2XXResponse**](GetDiscoverTv2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTvSeasonBySeasonId

> Season GetTvSeasonBySeasonId(ctx, tvId, seasonId).Language(language).Execute()

Get season details and episode list



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
	tvId := float32(76479) // float32 | 
	seasonId := float32(1) // float32 | 
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.GetTvSeasonBySeasonId(context.Background(), tvId, seasonId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.GetTvSeasonBySeasonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTvSeasonBySeasonId`: Season
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.GetTvSeasonBySeasonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tvId** | **float32** |  | 
**seasonId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTvSeasonBySeasonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **string** |  | 

### Return type

[**Season**](Season.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTvSimilar

> GetDiscoverTv2XXResponse GetTvSimilar(ctx, tvId).Page(page).Language(language).Execute()

Get similar TV series



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
	tvId := float32(76479) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.GetTvSimilar(context.Background(), tvId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.GetTvSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTvSimilar`: GetDiscoverTv2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.GetTvSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tvId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTvSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverTv2XXResponse**](GetDiscoverTv2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

