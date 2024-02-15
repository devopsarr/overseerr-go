# \OtherAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKeywordByKeywordId**](OtherAPI.md#GetKeywordByKeywordId) | **Get** /keyword/{keywordId} | Get keyword
[**ListWatchprovidersMovies**](OtherAPI.md#ListWatchprovidersMovies) | **Get** /watchproviders/movies | Get watch provider movies
[**ListWatchprovidersRegions**](OtherAPI.md#ListWatchprovidersRegions) | **Get** /watchproviders/regions | Get watch provider regions
[**ListWatchprovidersTv**](OtherAPI.md#ListWatchprovidersTv) | **Get** /watchproviders/tv | Get watch provider series



## GetKeywordByKeywordId

> Keyword GetKeywordByKeywordId(ctx, keywordId).Execute()

Get keyword



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
	keywordId := float32(1) // float32 | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.GetKeywordByKeywordId(context.Background(), keywordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.GetKeywordByKeywordId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeywordByKeywordId`: Keyword
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.GetKeywordByKeywordId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywordId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeywordByKeywordIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keyword**](Keyword.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchprovidersMovies

> []WatchProviderDetails ListWatchprovidersMovies(ctx).WatchRegion(watchRegion).Execute()

Get watch provider movies



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
	watchRegion := "US" // string | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.ListWatchprovidersMovies(context.Background()).WatchRegion(watchRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.ListWatchprovidersMovies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchprovidersMovies`: []WatchProviderDetails
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.ListWatchprovidersMovies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWatchprovidersMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchprovidersRegions

> []WatchProviderRegion ListWatchprovidersRegions(ctx).Execute()

Get watch provider regions



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
	resp, r, err := apiClient.OtherAPI.ListWatchprovidersRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.ListWatchprovidersRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchprovidersRegions`: []WatchProviderRegion
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.ListWatchprovidersRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchprovidersRegionsRequest struct via the builder pattern


### Return type

[**[]WatchProviderRegion**](WatchProviderRegion.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchprovidersTv

> []WatchProviderDetails ListWatchprovidersTv(ctx).WatchRegion(watchRegion).Execute()

Get watch provider series



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
	watchRegion := "US" // string | 

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.ListWatchprovidersTv(context.Background()).WatchRegion(watchRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.ListWatchprovidersTv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchprovidersTv`: []WatchProviderDetails
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.ListWatchprovidersTv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWatchprovidersTvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

