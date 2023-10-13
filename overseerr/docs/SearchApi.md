# \SearchApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiscoverKeywordMovies**](SearchApi.md#GetDiscoverKeywordMovies) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
[**GetDiscoverMovies**](SearchApi.md#GetDiscoverMovies) | **Get** /discover/movies | Discover movies
[**GetDiscoverMoviesGenreByGenreId**](SearchApi.md#GetDiscoverMoviesGenreByGenreId) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
[**GetDiscoverMoviesLanguageByLanguage**](SearchApi.md#GetDiscoverMoviesLanguageByLanguage) | **Get** /discover/movies/language/{language} | Discover movies by original language
[**GetDiscoverMoviesStudioByStudioId**](SearchApi.md#GetDiscoverMoviesStudioByStudioId) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
[**GetDiscoverMoviesUpcoming**](SearchApi.md#GetDiscoverMoviesUpcoming) | **Get** /discover/movies/upcoming | Upcoming movies
[**GetDiscoverTrending**](SearchApi.md#GetDiscoverTrending) | **Get** /discover/trending | Trending movies and TV
[**GetDiscoverTv**](SearchApi.md#GetDiscoverTv) | **Get** /discover/tv | Discover TV shows
[**GetDiscoverTvGenreByGenreId**](SearchApi.md#GetDiscoverTvGenreByGenreId) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
[**GetDiscoverTvLanguageByLanguage**](SearchApi.md#GetDiscoverTvLanguageByLanguage) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
[**GetDiscoverTvNetworkByNetworkId**](SearchApi.md#GetDiscoverTvNetworkByNetworkId) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
[**GetDiscoverTvUpcoming**](SearchApi.md#GetDiscoverTvUpcoming) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
[**GetDiscoverWatchlist**](SearchApi.md#GetDiscoverWatchlist) | **Get** /discover/watchlist | Get the Plex watchlist.
[**GetSearch**](SearchApi.md#GetSearch) | **Get** /search | Search for movies, TV shows, or people
[**GetSearchCompany**](SearchApi.md#GetSearchCompany) | **Get** /search/company | Search for companies
[**GetSearchKeyword**](SearchApi.md#GetSearchKeyword) | **Get** /search/keyword | Search for keywords
[**ListDiscoverGenresliderMovie**](SearchApi.md#ListDiscoverGenresliderMovie) | **Get** /discover/genreslider/movie | Get genre slider data for movies
[**ListDiscoverGenresliderTv**](SearchApi.md#ListDiscoverGenresliderTv) | **Get** /discover/genreslider/tv | Get genre slider data for TV series



## GetDiscoverKeywordMovies

> GetDiscoverMovies200Response GetDiscoverKeywordMovies(ctx, keywordId).Page(page).Language(language).Execute()

Get movies from keyword



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
    keywordId := float32(207317) // float32 | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverKeywordMovies(context.Background(), keywordId).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverKeywordMovies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverKeywordMovies`: GetDiscoverMovies200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverKeywordMovies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywordId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverKeywordMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMovies200Response**](GetDiscoverMovies200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverMovies

> GetDiscoverMovies200Response GetDiscoverMovies(ctx).Page(page).Language(language).Genre(genre).Studio(studio).Keywords(keywords).SortBy(sortBy).PrimaryReleaseDateGte(primaryReleaseDateGte).PrimaryReleaseDateLte(primaryReleaseDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Execute()

Discover movies



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
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)
    genre := "18" // string |  (optional)
    studio := float32(1) // float32 |  (optional)
    keywords := "1,2" // string |  (optional)
    sortBy := "popularity.desc" // string |  (optional)
    primaryReleaseDateGte := "2022-01-01T00:00:00Z" // string |  (optional)
    primaryReleaseDateLte := "2023-01-01T00:00:00Z" // string |  (optional)
    withRuntimeGte := float32(60) // float32 |  (optional)
    withRuntimeLte := float32(120) // float32 |  (optional)
    voteAverageGte := float32(7) // float32 |  (optional)
    voteAverageLte := float32(10) // float32 |  (optional)
    voteCountGte := float32(7) // float32 |  (optional)
    voteCountLte := float32(10) // float32 |  (optional)
    watchRegion := "US" // string |  (optional)
    watchProviders := "8|9" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverMovies(context.Background()).Page(page).Language(language).Genre(genre).Studio(studio).Keywords(keywords).SortBy(sortBy).PrimaryReleaseDateGte(primaryReleaseDateGte).PrimaryReleaseDateLte(primaryReleaseDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverMovies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverMovies`: GetDiscoverMovies200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverMovies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 
 **genre** | **string** |  | 
 **studio** | **float32** |  | 
 **keywords** | **string** |  | 
 **sortBy** | **string** |  | 
 **primaryReleaseDateGte** | **string** |  | 
 **primaryReleaseDateLte** | **string** |  | 
 **withRuntimeGte** | **float32** |  | 
 **withRuntimeLte** | **float32** |  | 
 **voteAverageGte** | **float32** |  | 
 **voteAverageLte** | **float32** |  | 
 **voteCountGte** | **float32** |  | 
 **voteCountLte** | **float32** |  | 
 **watchRegion** | **string** |  | 
 **watchProviders** | **string** |  | 

### Return type

[**GetDiscoverMovies200Response**](GetDiscoverMovies200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverMoviesGenreByGenreId

> GetDiscoverMoviesGenreByGenreId200Response GetDiscoverMoviesGenreByGenreId(ctx, genreId).Page(page).Language(language).Execute()

Discover movies by genre



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
    genreId := "1" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverMoviesGenreByGenreId(context.Background(), genreId).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverMoviesGenreByGenreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverMoviesGenreByGenreId`: GetDiscoverMoviesGenreByGenreId200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverMoviesGenreByGenreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverMoviesGenreByGenreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMoviesGenreByGenreId200Response**](GetDiscoverMoviesGenreByGenreId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverMoviesLanguageByLanguage

> GetDiscoverMoviesLanguageByLanguage200Response GetDiscoverMoviesLanguageByLanguage(ctx, language).Page(page).Language2(language2).Execute()

Discover movies by original language



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
    language := "en" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language2 := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverMoviesLanguageByLanguage(context.Background(), language).Page(page).Language2(language2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverMoviesLanguageByLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverMoviesLanguageByLanguage`: GetDiscoverMoviesLanguageByLanguage200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverMoviesLanguageByLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverMoviesLanguageByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language2** | **string** |  | 

### Return type

[**GetDiscoverMoviesLanguageByLanguage200Response**](GetDiscoverMoviesLanguageByLanguage200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverMoviesStudioByStudioId

> GetDiscoverMoviesStudioByStudioId200Response GetDiscoverMoviesStudioByStudioId(ctx, studioId).Page(page).Language(language).Execute()

Discover movies by studio



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
    studioId := "1" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverMoviesStudioByStudioId(context.Background(), studioId).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverMoviesStudioByStudioId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverMoviesStudioByStudioId`: GetDiscoverMoviesStudioByStudioId200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverMoviesStudioByStudioId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studioId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverMoviesStudioByStudioIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMoviesStudioByStudioId200Response**](GetDiscoverMoviesStudioByStudioId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverMoviesUpcoming

> GetDiscoverMovies200Response GetDiscoverMoviesUpcoming(ctx).Page(page).Language(language).Execute()

Upcoming movies



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
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverMoviesUpcoming(context.Background()).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverMoviesUpcoming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverMoviesUpcoming`: GetDiscoverMovies200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverMoviesUpcoming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverMoviesUpcomingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverMovies200Response**](GetDiscoverMovies200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTrending

> GetSearch200Response GetDiscoverTrending(ctx).Page(page).Language(language).Execute()

Trending movies and TV



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
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTrending(context.Background()).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTrending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTrending`: GetSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTrending`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTrendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetSearch200Response**](GetSearch200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTv

> GetDiscoverTv200Response GetDiscoverTv(ctx).Page(page).Language(language).Genre(genre).Network(network).Keywords(keywords).SortBy(sortBy).FirstAirDateGte(firstAirDateGte).FirstAirDateLte(firstAirDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Execute()

Discover TV shows



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
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)
    genre := "18" // string |  (optional)
    network := float32(1) // float32 |  (optional)
    keywords := "1,2" // string |  (optional)
    sortBy := "popularity.desc" // string |  (optional)
    firstAirDateGte := "2022-01-01T00:00:00Z" // string |  (optional)
    firstAirDateLte := "2023-01-01T00:00:00Z" // string |  (optional)
    withRuntimeGte := float32(60) // float32 |  (optional)
    withRuntimeLte := float32(120) // float32 |  (optional)
    voteAverageGte := float32(7) // float32 |  (optional)
    voteAverageLte := float32(10) // float32 |  (optional)
    voteCountGte := float32(7) // float32 |  (optional)
    voteCountLte := float32(10) // float32 |  (optional)
    watchRegion := "US" // string |  (optional)
    watchProviders := "8|9" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTv(context.Background()).Page(page).Language(language).Genre(genre).Network(network).Keywords(keywords).SortBy(sortBy).FirstAirDateGte(firstAirDateGte).FirstAirDateLte(firstAirDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTv`: GetDiscoverTv200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 
 **genre** | **string** |  | 
 **network** | **float32** |  | 
 **keywords** | **string** |  | 
 **sortBy** | **string** |  | 
 **firstAirDateGte** | **string** |  | 
 **firstAirDateLte** | **string** |  | 
 **withRuntimeGte** | **float32** |  | 
 **withRuntimeLte** | **float32** |  | 
 **voteAverageGte** | **float32** |  | 
 **voteAverageLte** | **float32** |  | 
 **voteCountGte** | **float32** |  | 
 **voteCountLte** | **float32** |  | 
 **watchRegion** | **string** |  | 
 **watchProviders** | **string** |  | 

### Return type

[**GetDiscoverTv200Response**](GetDiscoverTv200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTvGenreByGenreId

> GetDiscoverTvGenreByGenreId200Response GetDiscoverTvGenreByGenreId(ctx, genreId).Page(page).Language(language).Execute()

Discover TV shows by genre



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
    genreId := "1" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTvGenreByGenreId(context.Background(), genreId).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTvGenreByGenreId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTvGenreByGenreId`: GetDiscoverTvGenreByGenreId200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTvGenreByGenreId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTvGenreByGenreIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverTvGenreByGenreId200Response**](GetDiscoverTvGenreByGenreId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTvLanguageByLanguage

> GetDiscoverTvLanguageByLanguage200Response GetDiscoverTvLanguageByLanguage(ctx, language).Page(page).Language2(language2).Execute()

Discover TV shows by original language



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
    language := "en" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language2 := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTvLanguageByLanguage(context.Background(), language).Page(page).Language2(language2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTvLanguageByLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTvLanguageByLanguage`: GetDiscoverTvLanguageByLanguage200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTvLanguageByLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTvLanguageByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language2** | **string** |  | 

### Return type

[**GetDiscoverTvLanguageByLanguage200Response**](GetDiscoverTvLanguageByLanguage200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTvNetworkByNetworkId

> GetDiscoverTvNetworkByNetworkId200Response GetDiscoverTvNetworkByNetworkId(ctx, networkId).Page(page).Language(language).Execute()

Discover TV shows by network



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
    networkId := "1" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTvNetworkByNetworkId(context.Background(), networkId).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTvNetworkByNetworkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTvNetworkByNetworkId`: GetDiscoverTvNetworkByNetworkId200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTvNetworkByNetworkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTvNetworkByNetworkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverTvNetworkByNetworkId200Response**](GetDiscoverTvNetworkByNetworkId200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverTvUpcoming

> GetDiscoverTv200Response GetDiscoverTvUpcoming(ctx).Page(page).Language(language).Execute()

Discover Upcoming TV shows



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
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverTvUpcoming(context.Background()).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverTvUpcoming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverTvUpcoming`: GetDiscoverTv200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverTvUpcoming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverTvUpcomingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetDiscoverTv200Response**](GetDiscoverTv200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverWatchlist

> GetUserWatchlist200Response GetDiscoverWatchlist(ctx).Page(page).Execute()

Get the Plex watchlist.

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
    page := float32(1) // float32 |  (optional) (default to 1)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetDiscoverWatchlist(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetDiscoverWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscoverWatchlist`: GetUserWatchlist200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetDiscoverWatchlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]

### Return type

[**GetUserWatchlist200Response**](GetUserWatchlist200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearch

> GetSearch200Response GetSearch(ctx).Query(query).Page(page).Language(language).Execute()

Search for movies, TV shows, or people



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
    query := "Mulan" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)
    language := "en" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearch(context.Background()).Query(query).Page(page).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearch`: GetSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**GetSearch200Response**](GetSearch200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchCompany

> GetSearchCompany200Response GetSearchCompany(ctx).Query(query).Page(page).Execute()

Search for companies



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
    query := "Disney" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearchCompany(context.Background()).Query(query).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchCompany`: GetSearchCompany200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]

### Return type

[**GetSearchCompany200Response**](GetSearchCompany200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchKeyword

> GetSearchKeyword200Response GetSearchKeyword(ctx).Query(query).Page(page).Execute()

Search for keywords



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
    query := "christmas" // string | 
    page := float32(1) // float32 |  (optional) (default to 1)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearchKeyword(context.Background()).Query(query).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearchKeyword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearchKeyword`: GetSearchKeyword200Response
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearchKeyword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]

### Return type

[**GetSearchKeyword200Response**](GetSearchKeyword200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscoverGenresliderMovie

> []ListDiscoverGenresliderMovie200ResponseInner ListDiscoverGenresliderMovie(ctx).Language(language).Execute()

Get genre slider data for movies



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
    resp, r, err := apiClient.SearchApi.ListDiscoverGenresliderMovie(context.Background()).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ListDiscoverGenresliderMovie``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscoverGenresliderMovie`: []ListDiscoverGenresliderMovie200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ListDiscoverGenresliderMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDiscoverGenresliderMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListDiscoverGenresliderMovie200ResponseInner**](ListDiscoverGenresliderMovie200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscoverGenresliderTv

> []ListDiscoverGenresliderMovie200ResponseInner ListDiscoverGenresliderTv(ctx).Language(language).Execute()

Get genre slider data for TV series



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
    resp, r, err := apiClient.SearchApi.ListDiscoverGenresliderTv(context.Background()).Language(language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.ListDiscoverGenresliderTv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscoverGenresliderTv`: []ListDiscoverGenresliderMovie200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.ListDiscoverGenresliderTv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDiscoverGenresliderTvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]ListDiscoverGenresliderMovie200ResponseInner**](ListDiscoverGenresliderMovie200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

