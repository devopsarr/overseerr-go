# \UsersApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthResetPassword**](UsersApi.md#CreateAuthResetPassword) | **Post** /auth/reset-password | Send a reset password email
[**CreateAuthResetPasswordByGuid**](UsersApi.md#CreateAuthResetPasswordByGuid) | **Post** /auth/reset-password/{guid} | Reset the password for a user
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /user | Create new user
[**CreateUserImportFromPlex**](UsersApi.md#CreateUserImportFromPlex) | **Post** /user/import-from-plex | Import all users from Plex
[**CreateUserRegisterPushSubscription**](UsersApi.md#CreateUserRegisterPushSubscription) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
[**CreateUserSettingsMain**](UsersApi.md#CreateUserSettingsMain) | **Post** /user/{userId}/settings/main | Update general settings for a user
[**CreateUserSettingsNotifications**](UsersApi.md#CreateUserSettingsNotifications) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
[**CreateUserSettingsPassword**](UsersApi.md#CreateUserSettingsPassword) | **Post** /user/{userId}/settings/password | Update password for a user
[**CreateUserSettingsPermissions**](UsersApi.md#CreateUserSettingsPermissions) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /user/{userId} | Delete user by ID
[**GetAuthMe**](UsersApi.md#GetAuthMe) | **Get** /auth/me | Get logged-in user
[**GetUser**](UsersApi.md#GetUser) | **Get** /user | Get all users
[**GetUserByUserId**](UsersApi.md#GetUserByUserId) | **Get** /user/{userId} | Get user by ID
[**GetUserQuota**](UsersApi.md#GetUserQuota) | **Get** /user/{userId}/quota | Get quotas for a specific user
[**GetUserRequests**](UsersApi.md#GetUserRequests) | **Get** /user/{userId}/requests | Get requests for a specific user
[**GetUserSettingsMain**](UsersApi.md#GetUserSettingsMain) | **Get** /user/{userId}/settings/main | Get general settings for a user
[**GetUserSettingsNotifications**](UsersApi.md#GetUserSettingsNotifications) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
[**GetUserSettingsPassword**](UsersApi.md#GetUserSettingsPassword) | **Get** /user/{userId}/settings/password | Get password page informatiom
[**GetUserSettingsPermissions**](UsersApi.md#GetUserSettingsPermissions) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
[**GetUserWatchData**](UsersApi.md#GetUserWatchData) | **Get** /user/{userId}/watch_data | Get watch data
[**GetUserWatchlist**](UsersApi.md#GetUserWatchlist) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user
[**ListPlexUsers**](UsersApi.md#ListPlexUsers) | **Get** /settings/plex/users | Get Plex users
[**PutUser**](UsersApi.md#PutUser) | **Put** /user | Update batch of users
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /user/{userId} | Update a user by user ID



## CreateAuthResetPassword

> CreateAuthLogout200Response CreateAuthResetPassword(ctx).CreateAuthResetPasswordRequest(createAuthResetPasswordRequest).Execute()

Send a reset password email



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
    createAuthResetPasswordRequest := *overseerrClient.NewCreateAuthResetPasswordRequest("Email_example") // CreateAuthResetPasswordRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateAuthResetPassword(context.Background()).CreateAuthResetPasswordRequest(createAuthResetPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateAuthResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthResetPassword`: CreateAuthLogout200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateAuthResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAuthResetPasswordRequest** | [**CreateAuthResetPasswordRequest**](CreateAuthResetPasswordRequest.md) |  | 

### Return type

[**CreateAuthLogout200Response**](CreateAuthLogout200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthResetPasswordByGuid

> CreateAuthLogout200Response CreateAuthResetPasswordByGuid(ctx, guid).CreateAuthResetPasswordByGuidRequest(createAuthResetPasswordByGuidRequest).Execute()

Reset the password for a user



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
    guid := "9afef5a7-ec89-4d5f-9397-261e96970b50" // string | 
    createAuthResetPasswordByGuidRequest := *overseerrClient.NewCreateAuthResetPasswordByGuidRequest("Password_example") // CreateAuthResetPasswordByGuidRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateAuthResetPasswordByGuid(context.Background(), guid).CreateAuthResetPasswordByGuidRequest(createAuthResetPasswordByGuidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateAuthResetPasswordByGuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthResetPasswordByGuid`: CreateAuthLogout200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateAuthResetPasswordByGuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthResetPasswordByGuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAuthResetPasswordByGuidRequest** | [**CreateAuthResetPasswordByGuidRequest**](CreateAuthResetPasswordByGuidRequest.md) |  | 

### Return type

[**CreateAuthLogout200Response**](CreateAuthLogout200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Create new user



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
    createUserRequest := *overseerrClient.NewCreateUserRequest() // CreateUserRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserImportFromPlex

> []User CreateUserImportFromPlex(ctx).CreateUserImportFromPlexRequest(createUserImportFromPlexRequest).Execute()

Import all users from Plex



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
    createUserImportFromPlexRequest := *overseerrClient.NewCreateUserImportFromPlexRequest() // CreateUserImportFromPlexRequest |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserImportFromPlex(context.Background()).CreateUserImportFromPlexRequest(createUserImportFromPlexRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserImportFromPlex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserImportFromPlex`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUserImportFromPlex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserImportFromPlexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserImportFromPlexRequest** | [**CreateUserImportFromPlexRequest**](CreateUserImportFromPlexRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserRegisterPushSubscription

> CreateUserRegisterPushSubscription(ctx).CreateUserRegisterPushSubscriptionRequest(createUserRegisterPushSubscriptionRequest).Execute()

Register a web push /user/registerPushSubscription



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
    createUserRegisterPushSubscriptionRequest := *overseerrClient.NewCreateUserRegisterPushSubscriptionRequest("Endpoint_example", "Auth_example", "P256dh_example") // CreateUserRegisterPushSubscriptionRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserRegisterPushSubscription(context.Background()).CreateUserRegisterPushSubscriptionRequest(createUserRegisterPushSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserRegisterPushSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRegisterPushSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRegisterPushSubscriptionRequest** | [**CreateUserRegisterPushSubscriptionRequest**](CreateUserRegisterPushSubscriptionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSettingsMain

> GetUserSettingsMain200Response CreateUserSettingsMain(ctx, userId).CreateUserSettingsMainRequest(createUserSettingsMainRequest).Execute()

Update general settings for a user



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
    userId := float32(8.14) // float32 | 
    createUserSettingsMainRequest := *overseerrClient.NewCreateUserSettingsMainRequest() // CreateUserSettingsMainRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserSettingsMain(context.Background(), userId).CreateUserSettingsMainRequest(createUserSettingsMainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserSettingsMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserSettingsMain`: GetUserSettingsMain200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUserSettingsMain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSettingsMainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserSettingsMainRequest** | [**CreateUserSettingsMainRequest**](CreateUserSettingsMainRequest.md) |  | 

### Return type

[**GetUserSettingsMain200Response**](GetUserSettingsMain200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSettingsNotifications

> UserSettingsNotifications CreateUserSettingsNotifications(ctx, userId).UserSettingsNotifications(userSettingsNotifications).Execute()

Update notification settings for a user



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
    userId := float32(8.14) // float32 | 
    userSettingsNotifications := *overseerrClient.NewUserSettingsNotifications() // UserSettingsNotifications | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserSettingsNotifications(context.Background(), userId).UserSettingsNotifications(userSettingsNotifications).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserSettingsNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserSettingsNotifications`: UserSettingsNotifications
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUserSettingsNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSettingsNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userSettingsNotifications** | [**UserSettingsNotifications**](UserSettingsNotifications.md) |  | 

### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSettingsPassword

> CreateUserSettingsPassword(ctx, userId).CreateUserSettingsPasswordRequest(createUserSettingsPasswordRequest).Execute()

Update password for a user



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
    userId := float32(8.14) // float32 | 
    createUserSettingsPasswordRequest := *overseerrClient.NewCreateUserSettingsPasswordRequest("NewPassword_example") // CreateUserSettingsPasswordRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserSettingsPassword(context.Background(), userId).CreateUserSettingsPasswordRequest(createUserSettingsPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserSettingsPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSettingsPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserSettingsPasswordRequest** | [**CreateUserSettingsPasswordRequest**](CreateUserSettingsPasswordRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSettingsPermissions

> GetUserSettingsPermissions200Response CreateUserSettingsPermissions(ctx, userId).CreateUserSettingsPermissionsRequest(createUserSettingsPermissionsRequest).Execute()

Update permission settings for a user



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
    userId := float32(8.14) // float32 | 
    createUserSettingsPermissionsRequest := *overseerrClient.NewCreateUserSettingsPermissionsRequest(float32(123)) // CreateUserSettingsPermissionsRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.CreateUserSettingsPermissions(context.Background(), userId).CreateUserSettingsPermissionsRequest(createUserSettingsPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.CreateUserSettingsPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserSettingsPermissions`: GetUserSettingsPermissions200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.CreateUserSettingsPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSettingsPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserSettingsPermissionsRequest** | [**CreateUserSettingsPermissionsRequest**](CreateUserSettingsPermissionsRequest.md) |  | 

### Return type

[**GetUserSettingsPermissions200Response**](GetUserSettingsPermissions200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> User DeleteUser(ctx, userId).Execute()

Delete user by ID



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.DeleteUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthMe

> User GetAuthMe(ctx).Execute()

Get logged-in user



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
    resp, r, err := apiClient.UsersApi.GetAuthMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetAuthMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthMe`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetAuthMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthMeRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUser200Response GetUser(ctx).Take(take).Skip(skip).Sort(sort).Execute()

Get all users



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
    take := float32(20) // float32 |  (optional)
    skip := float32(0) // float32 |  (optional)
    sort := "sort_example" // string |  (optional) (default to "created")

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUser(context.Background()).Take(take).Skip(skip).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: GetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **sort** | **string** |  | [default to &quot;created&quot;]

### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByUserId

> User GetUserByUserId(ctx, userId).Execute()

Get user by ID



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserByUserId(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserByUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserByUserId`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserByUserId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserQuota

> GetUserQuota200Response GetUserQuota(ctx, userId).Execute()

Get quotas for a specific user



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserQuota(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserQuota`: GetUserQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserQuota200Response**](GetUserQuota200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRequests

> GetUserRequests200Response GetUserRequests(ctx, userId).Take(take).Skip(skip).Execute()

Get requests for a specific user



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
    userId := float32(8.14) // float32 | 
    take := float32(20) // float32 |  (optional)
    skip := float32(0) // float32 |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserRequests(context.Background(), userId).Take(take).Skip(skip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRequests`: GetUserRequests200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **take** | **float32** |  | 
 **skip** | **float32** |  | 

### Return type

[**GetUserRequests200Response**](GetUserRequests200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsMain

> GetUserSettingsMain200Response GetUserSettingsMain(ctx, userId).Execute()

Get general settings for a user



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserSettingsMain(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserSettingsMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsMain`: GetUserSettingsMain200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserSettingsMain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsMainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserSettingsMain200Response**](GetUserSettingsMain200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsNotifications

> UserSettingsNotifications GetUserSettingsNotifications(ctx, userId).Execute()

Get notification settings for a user



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserSettingsNotifications(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserSettingsNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsNotifications`: UserSettingsNotifications
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserSettingsNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsPassword

> GetUserSettingsPassword200Response GetUserSettingsPassword(ctx, userId).Execute()

Get password page informatiom



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserSettingsPassword(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserSettingsPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsPassword`: GetUserSettingsPassword200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserSettingsPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserSettingsPassword200Response**](GetUserSettingsPassword200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsPermissions

> GetUserSettingsPermissions200Response GetUserSettingsPermissions(ctx, userId).Execute()

Get permission settings for a user



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserSettingsPermissions(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserSettingsPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsPermissions`: GetUserSettingsPermissions200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserSettingsPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserSettingsPermissions200Response**](GetUserSettingsPermissions200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserWatchData

> GetUserWatchData200Response GetUserWatchData(ctx, userId).Execute()

Get watch data



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
    userId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserWatchData(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserWatchData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWatchData`: GetUserWatchData200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserWatchData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWatchDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserWatchData200Response**](GetUserWatchData200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserWatchlist

> GetUserWatchlist200Response GetUserWatchlist(ctx, userId).Page(page).Execute()

Get the Plex watchlist for a specific user



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
    userId := float32(8.14) // float32 | 
    page := float32(1) // float32 |  (optional) (default to 1)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUserWatchlist(context.Background(), userId).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserWatchlist`: GetUserWatchlist200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserWatchlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserWatchlistRequest struct via the builder pattern


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


## ListPlexUsers

> []ListPlexUsers200ResponseInner ListPlexUsers(ctx).Execute()

Get Plex users



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
    resp, r, err := apiClient.UsersApi.ListPlexUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListPlexUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexUsers`: []ListPlexUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListPlexUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPlexUsersRequest struct via the builder pattern


### Return type

[**[]ListPlexUsers200ResponseInner**](ListPlexUsers200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUser

> []User PutUser(ctx).PutUserRequest(putUserRequest).Execute()

Update batch of users



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
    putUserRequest := *overseerrClient.NewPutUserRequest() // PutUserRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.PutUser(context.Background()).PutUserRequest(putUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.PutUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUser`: []User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.PutUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putUserRequest** | [**PutUserRequest**](PutUserRequest.md) |  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).User(user).Execute()

Update a user by user ID



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
    userId := float32(8.14) // float32 | 
    user := *overseerrClient.NewUser(int32(1), "hey@itsme.com", "2020-09-02T05:02:23.000Z", "2020-09-02T05:02:23.000Z") // User | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.UpdateUser(context.Background(), userId).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

