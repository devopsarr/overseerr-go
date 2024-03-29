# \PersonAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPersonByPersonId**](PersonAPI.md#GetPersonByPersonId) | **Get** /person/{personId} | Get person details
[**GetPersonCombinedCredits**](PersonAPI.md#GetPersonCombinedCredits) | **Get** /person/{personId}/combined_credits | Get combined credits



## GetPersonByPersonId

> PersonDetails GetPersonByPersonId(ctx, personId).Language(language).Execute()

Get person details



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
	personId := float32(287) // float32 | 
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonAPI.GetPersonByPersonId(context.Background(), personId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GetPersonByPersonId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonByPersonId`: PersonDetails
	fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GetPersonByPersonId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonByPersonIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonCombinedCredits

> GetPersonCombinedCredits2XXResponse GetPersonCombinedCredits(ctx, personId).Language(language).Execute()

Get combined credits



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
	personId := float32(287) // float32 | 
	language := "en" // string |  (optional)

	configuration := overseerrClient.NewConfiguration()
	apiClient := overseerrClient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonAPI.GetPersonCombinedCredits(context.Background(), personId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.GetPersonCombinedCredits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonCombinedCredits`: GetPersonCombinedCredits2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `PersonAPI.GetPersonCombinedCredits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonCombinedCreditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**GetPersonCombinedCredits2XXResponse**](GetPersonCombinedCredits2XXResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

