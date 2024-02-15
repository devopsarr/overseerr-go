# \MoviesAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovieByMovieId**](MoviesAPI.md#GetMovieByMovieId) | **Get** /movie/{movieId} | Get movie details
[**GetMovieRatings**](MoviesAPI.md#GetMovieRatings) | **Get** /movie/{movieId}/ratings | Get movie ratings
[**GetMovieRecommendations**](MoviesAPI.md#GetMovieRecommendations) | **Get** /movie/{movieId}/recommendations | Get recommended movies
[**GetMovieSimilar**](MoviesAPI.md#GetMovieSimilar) | **Get** /movie/{movieId}/similar | Get similar movies



## GetMovieByMovieId

> MovieDetails GetMovieByMovieId(ctx, movieId).Language(language).Execute()

Get movie details



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
	movieId := float32(337401) // float32 | 
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.GetMovieByMovieId(context.Background(), movieId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.GetMovieByMovieId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieByMovieId`: MovieDetails
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.GetMovieByMovieId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieByMovieIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**MovieDetails**](MovieDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieRatings

> GetMovieRatings2XXResponse GetMovieRatings(ctx, movieId).Execute()

Get movie ratings



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
	movieId := float32(337401) // float32 | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.GetMovieRatings(context.Background(), movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.GetMovieRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieRatings`: GetMovieRatings2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.GetMovieRatings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieRatingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMovieRatings2XXResponse**](GetMovieRatings2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieRecommendations

> GetDiscoverMovies2XXResponse GetMovieRecommendations(ctx, movieId).Page(page).Language(language).Execute()

Get recommended movies



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
	movieId := float32(337401) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.GetMovieRecommendations(context.Background(), movieId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.GetMovieRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieRecommendations`: GetDiscoverMovies2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.GetMovieRecommendations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMovies2XXResponse**](GetDiscoverMovies2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieSimilar

> GetDiscoverMovies2XXResponse GetMovieSimilar(ctx, movieId).Page(page).Language(language).Execute()

Get similar movies



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
	movieId := float32(337401) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.GetMovieSimilar(context.Background(), movieId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.GetMovieSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieSimilar`: GetDiscoverMovies2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.GetMovieSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMovies2XXResponse**](GetDiscoverMovies2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

