# \PublicAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatus**](PublicAPI.md#GetStatus) | **Get** /status | Get Overseerr status
[**GetStatusAppdata**](PublicAPI.md#GetStatusAppdata) | **Get** /status/appdata | Get application data volume status



## GetStatus

> GetStatus2XXResponse GetStatus(ctx).Execute()

Get Overseerr status



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
	resp, r, err := apiClient.PublicAPI.GetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: GetStatus2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**GetStatus2XXResponse**](GetStatus2XXResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusAppdata

> GetStatusAppdata2XXResponse GetStatusAppdata(ctx).Execute()

Get application data volume status



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
	resp, r, err := apiClient.PublicAPI.GetStatusAppdata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.GetStatusAppdata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusAppdata`: GetStatusAppdata2XXResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.GetStatusAppdata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusAppdataRequest struct via the builder pattern


### Return type

[**GetStatusAppdata2XXResponse**](GetStatusAppdata2XXResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

