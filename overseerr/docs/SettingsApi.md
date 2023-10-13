# \SettingsApi

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCacheFlush**](SettingsApi.md#CreateCacheFlush) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
[**CreateDiscover**](SettingsApi.md#CreateDiscover) | **Post** /settings/discover | Batch update all sliders.
[**CreateDiscoverAdd**](SettingsApi.md#CreateDiscoverAdd) | **Post** /settings/discover/add | Add a new slider
[**CreateInitialize**](SettingsApi.md#CreateInitialize) | **Post** /settings/initialize | Initialize application
[**CreateJobsCancel**](SettingsApi.md#CreateJobsCancel) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
[**CreateJobsRun**](SettingsApi.md#CreateJobsRun) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
[**CreateJobsSchedule**](SettingsApi.md#CreateJobsSchedule) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
[**CreateMain**](SettingsApi.md#CreateMain) | **Post** /settings/main | Update main settings
[**CreateMainRegenerate**](SettingsApi.md#CreateMainRegenerate) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
[**CreateNotificationsDiscord**](SettingsApi.md#CreateNotificationsDiscord) | **Post** /settings/notifications/discord | Update Discord notification settings
[**CreateNotificationsEmail**](SettingsApi.md#CreateNotificationsEmail) | **Post** /settings/notifications/email | Update email notification settings
[**CreateNotificationsGotify**](SettingsApi.md#CreateNotificationsGotify) | **Post** /settings/notifications/gotify | Update Gotify notification settings
[**CreateNotificationsLunasea**](SettingsApi.md#CreateNotificationsLunasea) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
[**CreateNotificationsPushbullet**](SettingsApi.md#CreateNotificationsPushbullet) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
[**CreateNotificationsPushover**](SettingsApi.md#CreateNotificationsPushover) | **Post** /settings/notifications/pushover | Update Pushover notification settings
[**CreateNotificationsSlack**](SettingsApi.md#CreateNotificationsSlack) | **Post** /settings/notifications/slack | Update Slack notification settings
[**CreateNotificationsTelegram**](SettingsApi.md#CreateNotificationsTelegram) | **Post** /settings/notifications/telegram | Update Telegram notification settings
[**CreateNotificationsWebhook**](SettingsApi.md#CreateNotificationsWebhook) | **Post** /settings/notifications/webhook | Update webhook notification settings
[**CreateNotificationsWebpush**](SettingsApi.md#CreateNotificationsWebpush) | **Post** /settings/notifications/webpush | Update Web Push notification settings
[**CreatePlex**](SettingsApi.md#CreatePlex) | **Post** /settings/plex | Update Plex settings
[**CreatePlexSync**](SettingsApi.md#CreatePlexSync) | **Post** /settings/plex/sync | Start full Plex library scan
[**CreateRadarr**](SettingsApi.md#CreateRadarr) | **Post** /settings/radarr | Create Radarr instance
[**CreateSonarr**](SettingsApi.md#CreateSonarr) | **Post** /settings/sonarr | Create Sonarr instance
[**CreateTautulli**](SettingsApi.md#CreateTautulli) | **Post** /settings/tautulli | Update Tautulli settings
[**DeleteDiscover**](SettingsApi.md#DeleteDiscover) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
[**DeleteRadarr**](SettingsApi.md#DeleteRadarr) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
[**DeleteSonarr**](SettingsApi.md#DeleteSonarr) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
[**GetAbout**](SettingsApi.md#GetAbout) | **Get** /settings/about | Get server stats
[**GetCache**](SettingsApi.md#GetCache) | **Get** /settings/cache | Get a list of active caches
[**GetDiscoverReset**](SettingsApi.md#GetDiscoverReset) | **Get** /settings/discover/reset | Reset all discover sliders
[**GetMain**](SettingsApi.md#GetMain) | **Get** /settings/main | Get main settings
[**GetNotificationsDiscord**](SettingsApi.md#GetNotificationsDiscord) | **Get** /settings/notifications/discord | Get Discord notification settings
[**GetNotificationsEmail**](SettingsApi.md#GetNotificationsEmail) | **Get** /settings/notifications/email | Get email notification settings
[**GetNotificationsGotify**](SettingsApi.md#GetNotificationsGotify) | **Get** /settings/notifications/gotify | Get Gotify notification settings
[**GetNotificationsLunasea**](SettingsApi.md#GetNotificationsLunasea) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
[**GetNotificationsPushbullet**](SettingsApi.md#GetNotificationsPushbullet) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
[**GetNotificationsPushover**](SettingsApi.md#GetNotificationsPushover) | **Get** /settings/notifications/pushover | Get Pushover notification settings
[**GetNotificationsSlack**](SettingsApi.md#GetNotificationsSlack) | **Get** /settings/notifications/slack | Get Slack notification settings
[**GetNotificationsTelegram**](SettingsApi.md#GetNotificationsTelegram) | **Get** /settings/notifications/telegram | Get Telegram notification settings
[**GetNotificationsWebhook**](SettingsApi.md#GetNotificationsWebhook) | **Get** /settings/notifications/webhook | Get webhook notification settings
[**GetNotificationsWebpush**](SettingsApi.md#GetNotificationsWebpush) | **Get** /settings/notifications/webpush | Get Web Push notification settings
[**GetPlex**](SettingsApi.md#GetPlex) | **Get** /settings/plex | Get Plex settings
[**GetPlexSync**](SettingsApi.md#GetPlexSync) | **Get** /settings/plex/sync | Get status of full Plex library scan
[**GetPublic**](SettingsApi.md#GetPublic) | **Get** /settings/public | Get public settings
[**GetTautulli**](SettingsApi.md#GetTautulli) | **Get** /settings/tautulli | Get Tautulli settings
[**ListDiscover**](SettingsApi.md#ListDiscover) | **Get** /settings/discover | Get all discover sliders
[**ListJobs**](SettingsApi.md#ListJobs) | **Get** /settings/jobs | Get scheduled jobs
[**ListLogs**](SettingsApi.md#ListLogs) | **Get** /settings/logs | Returns logs
[**ListPlexDevicesServers**](SettingsApi.md#ListPlexDevicesServers) | **Get** /settings/plex/devices/servers | Gets the user&#39;s available Plex servers
[**ListPlexLibrary**](SettingsApi.md#ListPlexLibrary) | **Get** /settings/plex/library | Get Plex libraries
[**ListPlexUsers**](SettingsApi.md#ListPlexUsers) | **Get** /settings/plex/users | Get Plex users
[**ListRadarr**](SettingsApi.md#ListRadarr) | **Get** /settings/radarr | Get Radarr settings
[**ListRadarrProfiles**](SettingsApi.md#ListRadarrProfiles) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
[**ListSonarr**](SettingsApi.md#ListSonarr) | **Get** /settings/sonarr | Get Sonarr settings
[**TestNotificationsDiscord**](SettingsApi.md#TestNotificationsDiscord) | **Post** /settings/notifications/discord/test | Test Discord settings
[**TestNotificationsEmail**](SettingsApi.md#TestNotificationsEmail) | **Post** /settings/notifications/email/test | Test email settings
[**TestNotificationsGotify**](SettingsApi.md#TestNotificationsGotify) | **Post** /settings/notifications/gotify/test | Test Gotify settings
[**TestNotificationsLunasea**](SettingsApi.md#TestNotificationsLunasea) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
[**TestNotificationsPushbullet**](SettingsApi.md#TestNotificationsPushbullet) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
[**TestNotificationsPushover**](SettingsApi.md#TestNotificationsPushover) | **Post** /settings/notifications/pushover/test | Test Pushover settings
[**TestNotificationsSlack**](SettingsApi.md#TestNotificationsSlack) | **Post** /settings/notifications/slack/test | Test Slack settings
[**TestNotificationsTelegram**](SettingsApi.md#TestNotificationsTelegram) | **Post** /settings/notifications/telegram/test | Test Telegram settings
[**TestNotificationsWebhook**](SettingsApi.md#TestNotificationsWebhook) | **Post** /settings/notifications/webhook/test | Test webhook settings
[**TestNotificationsWebpush**](SettingsApi.md#TestNotificationsWebpush) | **Post** /settings/notifications/webpush/test | Test Web Push settings
[**TestRadarr**](SettingsApi.md#TestRadarr) | **Post** /settings/radarr/test | Test Radarr configuration
[**TestSonarr**](SettingsApi.md#TestSonarr) | **Post** /settings/sonarr/test | Test Sonarr configuration
[**UpdateDiscover**](SettingsApi.md#UpdateDiscover) | **Put** /settings/discover/{sliderId} | Update a single slider
[**UpdateRadarr**](SettingsApi.md#UpdateRadarr) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
[**UpdateSonarr**](SettingsApi.md#UpdateSonarr) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance



## CreateCacheFlush

> CreateCacheFlush(ctx, cacheId).Execute()

Flush a specific cache



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
    cacheId := "cacheId_example" // string | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateCacheFlush(context.Background(), cacheId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateCacheFlush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cacheId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCacheFlushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDiscover

> []DiscoverSlider CreateDiscover(ctx).DiscoverSlider(discoverSlider).Execute()

Batch update all sliders.



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
    discoverSlider := []overseerrClient.DiscoverSlider{*overseerrClient.NewDiscoverSlider(float32(1), "Title_example", false, "1234")} // []DiscoverSlider | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateDiscover(context.Background()).DiscoverSlider(discoverSlider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscover`: []DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateDiscover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discoverSlider** | [**[]DiscoverSlider**](DiscoverSlider.md) |  | 

### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDiscoverAdd

> DiscoverSlider CreateDiscoverAdd(ctx).CreateDiscoverAddRequest(createDiscoverAddRequest).Execute()

Add a new slider



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
    createDiscoverAddRequest := *overseerrClient.NewCreateDiscoverAddRequest() // CreateDiscoverAddRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateDiscoverAdd(context.Background()).CreateDiscoverAddRequest(createDiscoverAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateDiscoverAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscoverAdd`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateDiscoverAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscoverAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDiscoverAddRequest** | [**CreateDiscoverAddRequest**](CreateDiscoverAddRequest.md) |  | 

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInitialize

> PublicSettings CreateInitialize(ctx).Execute()

Initialize application



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
    resp, r, err := apiClient.SettingsApi.CreateInitialize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateInitialize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInitialize`: PublicSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateInitialize`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInitializeRequest struct via the builder pattern


### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobsCancel

> Job CreateJobsCancel(ctx, jobId).Execute()

Cancel a specific job



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
    jobId := "jobId_example" // string | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateJobsCancel(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateJobsCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsCancel`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateJobsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobsRun

> Job CreateJobsRun(ctx, jobId).Execute()

Invoke a specific job



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
    jobId := "jobId_example" // string | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateJobsRun(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateJobsRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsRun`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateJobsRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobsRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateJobsSchedule

> Job CreateJobsSchedule(ctx, jobId).CreateJobsScheduleRequest(createJobsScheduleRequest).Execute()

Modify job schedule



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
    jobId := "jobId_example" // string | 
    createJobsScheduleRequest := *overseerrClient.NewCreateJobsScheduleRequest() // CreateJobsScheduleRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateJobsSchedule(context.Background(), jobId).CreateJobsScheduleRequest(createJobsScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateJobsSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsSchedule`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateJobsSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateJobsScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createJobsScheduleRequest** | [**CreateJobsScheduleRequest**](CreateJobsScheduleRequest.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMain

> MainSettings CreateMain(ctx).MainSettings(mainSettings).Execute()

Update main settings



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
    mainSettings := *overseerrClient.NewMainSettings() // MainSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateMain(context.Background()).MainSettings(mainSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMain`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateMain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mainSettings** | [**MainSettings**](MainSettings.md) |  | 

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMainRegenerate

> MainSettings CreateMainRegenerate(ctx).Execute()

Get main settings with newly-generated API key



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
    resp, r, err := apiClient.SettingsApi.CreateMainRegenerate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateMainRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMainRegenerate`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateMainRegenerate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMainRegenerateRequest struct via the builder pattern


### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsDiscord

> DiscordSettings CreateNotificationsDiscord(ctx).DiscordSettings(discordSettings).Execute()

Update Discord notification settings



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
    discordSettings := *overseerrClient.NewDiscordSettings() // DiscordSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsDiscord(context.Background()).DiscordSettings(discordSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsDiscord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsDiscord`: DiscordSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsDiscord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsDiscordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discordSettings** | [**DiscordSettings**](DiscordSettings.md) |  | 

### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsEmail

> NotificationEmailSettings CreateNotificationsEmail(ctx).NotificationEmailSettings(notificationEmailSettings).Execute()

Update email notification settings



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
    notificationEmailSettings := *overseerrClient.NewNotificationEmailSettings() // NotificationEmailSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsEmail(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsEmail`: NotificationEmailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationEmailSettings** | [**NotificationEmailSettings**](NotificationEmailSettings.md) |  | 

### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsGotify

> GotifySettings CreateNotificationsGotify(ctx).GotifySettings(gotifySettings).Execute()

Update Gotify notification settings



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
    gotifySettings := *overseerrClient.NewGotifySettings() // GotifySettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsGotify(context.Background()).GotifySettings(gotifySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsGotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsGotify`: GotifySettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsGotify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsGotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gotifySettings** | [**GotifySettings**](GotifySettings.md) |  | 

### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsLunasea

> LunaSeaSettings CreateNotificationsLunasea(ctx).LunaSeaSettings(lunaSeaSettings).Execute()

Update LunaSea notification settings



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
    lunaSeaSettings := *overseerrClient.NewLunaSeaSettings() // LunaSeaSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsLunasea(context.Background()).LunaSeaSettings(lunaSeaSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsLunasea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsLunasea`: LunaSeaSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsLunasea`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsLunaseaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lunaSeaSettings** | [**LunaSeaSettings**](LunaSeaSettings.md) |  | 

### Return type

[**LunaSeaSettings**](LunaSeaSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsPushbullet

> PushbulletSettings CreateNotificationsPushbullet(ctx).PushbulletSettings(pushbulletSettings).Execute()

Update Pushbullet notification settings



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
    pushbulletSettings := *overseerrClient.NewPushbulletSettings() // PushbulletSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsPushbullet(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsPushbullet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsPushbullet`: PushbulletSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsPushbullet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsPushbulletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushbulletSettings** | [**PushbulletSettings**](PushbulletSettings.md) |  | 

### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsPushover

> PushoverSettings CreateNotificationsPushover(ctx).PushoverSettings(pushoverSettings).Execute()

Update Pushover notification settings



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
    pushoverSettings := *overseerrClient.NewPushoverSettings() // PushoverSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsPushover(context.Background()).PushoverSettings(pushoverSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsPushover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsPushover`: PushoverSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsPushover`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsPushoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushoverSettings** | [**PushoverSettings**](PushoverSettings.md) |  | 

### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsSlack

> SlackSettings CreateNotificationsSlack(ctx).SlackSettings(slackSettings).Execute()

Update Slack notification settings



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
    slackSettings := *overseerrClient.NewSlackSettings() // SlackSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsSlack(context.Background()).SlackSettings(slackSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsSlack`: SlackSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsSlack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsSlackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slackSettings** | [**SlackSettings**](SlackSettings.md) |  | 

### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsTelegram

> TelegramSettings CreateNotificationsTelegram(ctx).TelegramSettings(telegramSettings).Execute()

Update Telegram notification settings



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
    telegramSettings := *overseerrClient.NewTelegramSettings() // TelegramSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsTelegram(context.Background()).TelegramSettings(telegramSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsTelegram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsTelegram`: TelegramSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsTelegram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsTelegramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telegramSettings** | [**TelegramSettings**](TelegramSettings.md) |  | 

### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsWebhook

> WebhookSettings CreateNotificationsWebhook(ctx).WebhookSettings(webhookSettings).Execute()

Update webhook notification settings



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
    webhookSettings := *overseerrClient.NewWebhookSettings() // WebhookSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsWebhook(context.Background()).WebhookSettings(webhookSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsWebhook`: WebhookSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSettings** | [**WebhookSettings**](WebhookSettings.md) |  | 

### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotificationsWebpush

> WebPushSettings CreateNotificationsWebpush(ctx).WebPushSettings(webPushSettings).Execute()

Update Web Push notification settings



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
    webPushSettings := *overseerrClient.NewWebPushSettings() // WebPushSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateNotificationsWebpush(context.Background()).WebPushSettings(webPushSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateNotificationsWebpush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsWebpush`: WebPushSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateNotificationsWebpush`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsWebpushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webPushSettings** | [**WebPushSettings**](WebPushSettings.md) |  | 

### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlex

> PlexSettings CreatePlex(ctx).PlexSettings(plexSettings).Execute()

Update Plex settings



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
    plexSettings := *overseerrClient.NewPlexSettings("Main Server", "1234123412341234", "127.0.0.1", float32(32400)) // PlexSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreatePlex(context.Background()).PlexSettings(plexSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreatePlex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlex`: PlexSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreatePlex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plexSettings** | [**PlexSettings**](PlexSettings.md) |  | 

### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlexSync

> GetPlexSync200Response CreatePlexSync(ctx).CreatePlexSyncRequest(createPlexSyncRequest).Execute()

Start full Plex library scan



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
    createPlexSyncRequest := *overseerrClient.NewCreatePlexSyncRequest() // CreatePlexSyncRequest |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreatePlexSync(context.Background()).CreatePlexSyncRequest(createPlexSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreatePlexSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlexSync`: GetPlexSync200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreatePlexSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlexSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlexSyncRequest** | [**CreatePlexSyncRequest**](CreatePlexSyncRequest.md) |  | 

### Return type

[**GetPlexSync200Response**](GetPlexSync200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRadarr

> RadarrSettings CreateRadarr(ctx).RadarrSettings(radarrSettings).Execute()

Create Radarr instance



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
    radarrSettings := *overseerrClient.NewRadarrSettings("Radarr Main", "127.0.0.1", float32(7878), "exampleapikey", false, float32(1), "720p/1080p", "/movies", false, "In Cinema", false) // RadarrSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateRadarr(context.Background()).RadarrSettings(radarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateRadarr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRadarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **radarrSettings** | [**RadarrSettings**](RadarrSettings.md) |  | 

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSonarr

> SonarrSettings CreateSonarr(ctx).SonarrSettings(sonarrSettings).Execute()

Create Sonarr instance



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
    sonarrSettings := *overseerrClient.NewSonarrSettings("Sonarr Main", "127.0.0.1", float32(8989), "exampleapikey", false, float32(1), "720p/1080p", "/tv/", false, false, false) // SonarrSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateSonarr(context.Background()).SonarrSettings(sonarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateSonarr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSonarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sonarrSettings** | [**SonarrSettings**](SonarrSettings.md) |  | 

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTautulli

> TautulliSettings CreateTautulli(ctx).TautulliSettings(tautulliSettings).Execute()

Update Tautulli settings



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
    tautulliSettings := *overseerrClient.NewTautulliSettings() // TautulliSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.CreateTautulli(context.Background()).TautulliSettings(tautulliSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.CreateTautulli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTautulli`: TautulliSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.CreateTautulli`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTautulliRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tautulliSettings** | [**TautulliSettings**](TautulliSettings.md) |  | 

### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDiscover

> DiscoverSlider DeleteDiscover(ctx, sliderId).Execute()

Delete slider by ID



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
    sliderId := float32(8.14) // float32 | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeleteDiscover(context.Background(), sliderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDiscover`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.DeleteDiscover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliderId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRadarr

> RadarrSettings DeleteRadarr(ctx, radarrId).Execute()

Delete Radarr instance



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
    radarrId := int32(56) // int32 | Radarr instance ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeleteRadarr(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.DeleteRadarr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radarrId** | **int32** | Radarr instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRadarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSonarr

> SonarrSettings DeleteSonarr(ctx, sonarrId).Execute()

Delete Sonarr instance



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
    sonarrId := int32(56) // int32 | Sonarr instance ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.DeleteSonarr(context.Background(), sonarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.DeleteSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.DeleteSonarr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sonarrId** | **int32** | Sonarr instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSonarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAbout

> GetAbout200Response GetAbout(ctx).Execute()

Get server stats



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
    resp, r, err := apiClient.SettingsApi.GetAbout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetAbout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAbout`: GetAbout200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetAbout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutRequest struct via the builder pattern


### Return type

[**GetAbout200Response**](GetAbout200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCache

> GetCache200Response GetCache(ctx).Execute()

Get a list of active caches



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
    resp, r, err := apiClient.SettingsApi.GetCache(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCache`: GetCache200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetCache`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCacheRequest struct via the builder pattern


### Return type

[**GetCache200Response**](GetCache200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDiscoverReset

> GetDiscoverReset(ctx).Execute()

Reset all discover sliders



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
    resp, r, err := apiClient.SettingsApi.GetDiscoverReset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetDiscoverReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscoverResetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMain

> MainSettings GetMain(ctx).Execute()

Get main settings



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
    resp, r, err := apiClient.SettingsApi.GetMain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMain`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetMain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMainRequest struct via the builder pattern


### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsDiscord

> DiscordSettings GetNotificationsDiscord(ctx).Execute()

Get Discord notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsDiscord(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsDiscord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsDiscord`: DiscordSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsDiscord`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsDiscordRequest struct via the builder pattern


### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsEmail

> NotificationEmailSettings GetNotificationsEmail(ctx).Execute()

Get email notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsEmail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsEmail`: NotificationEmailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsEmail`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsEmailRequest struct via the builder pattern


### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsGotify

> GotifySettings GetNotificationsGotify(ctx).Execute()

Get Gotify notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsGotify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsGotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsGotify`: GotifySettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsGotify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsGotifyRequest struct via the builder pattern


### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsLunasea

> LunaSeaSettings GetNotificationsLunasea(ctx).Execute()

Get LunaSea notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsLunasea(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsLunasea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsLunasea`: LunaSeaSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsLunasea`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsLunaseaRequest struct via the builder pattern


### Return type

[**LunaSeaSettings**](LunaSeaSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsPushbullet

> PushbulletSettings GetNotificationsPushbullet(ctx).Execute()

Get Pushbullet notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsPushbullet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsPushbullet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsPushbullet`: PushbulletSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsPushbullet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsPushbulletRequest struct via the builder pattern


### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsPushover

> PushoverSettings GetNotificationsPushover(ctx).Execute()

Get Pushover notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsPushover(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsPushover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsPushover`: PushoverSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsPushover`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsPushoverRequest struct via the builder pattern


### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsSlack

> SlackSettings GetNotificationsSlack(ctx).Execute()

Get Slack notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsSlack(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsSlack`: SlackSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsSlack`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsSlackRequest struct via the builder pattern


### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsTelegram

> TelegramSettings GetNotificationsTelegram(ctx).Execute()

Get Telegram notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsTelegram(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsTelegram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsTelegram`: TelegramSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsTelegram`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsTelegramRequest struct via the builder pattern


### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsWebhook

> WebhookSettings GetNotificationsWebhook(ctx).Execute()

Get webhook notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsWebhook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsWebhook`: WebhookSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsWebhookRequest struct via the builder pattern


### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationsWebpush

> WebPushSettings GetNotificationsWebpush(ctx).Execute()

Get Web Push notification settings



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
    resp, r, err := apiClient.SettingsApi.GetNotificationsWebpush(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetNotificationsWebpush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsWebpush`: WebPushSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetNotificationsWebpush`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsWebpushRequest struct via the builder pattern


### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlex

> PlexSettings GetPlex(ctx).Execute()

Get Plex settings



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
    resp, r, err := apiClient.SettingsApi.GetPlex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetPlex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlex`: PlexSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetPlex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlexRequest struct via the builder pattern


### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlexSync

> GetPlexSync200Response GetPlexSync(ctx).Execute()

Get status of full Plex library scan



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
    resp, r, err := apiClient.SettingsApi.GetPlexSync(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetPlexSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlexSync`: GetPlexSync200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetPlexSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlexSyncRequest struct via the builder pattern


### Return type

[**GetPlexSync200Response**](GetPlexSync200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublic

> PublicSettings GetPublic(ctx).Execute()

Get public settings



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
    resp, r, err := apiClient.SettingsApi.GetPublic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublic`: PublicSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicRequest struct via the builder pattern


### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTautulli

> TautulliSettings GetTautulli(ctx).Execute()

Get Tautulli settings



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
    resp, r, err := apiClient.SettingsApi.GetTautulli(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.GetTautulli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTautulli`: TautulliSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.GetTautulli`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTautulliRequest struct via the builder pattern


### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDiscover

> []DiscoverSlider ListDiscover(ctx).Execute()

Get all discover sliders



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
    resp, r, err := apiClient.SettingsApi.ListDiscover(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscover`: []DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListDiscover`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscoverRequest struct via the builder pattern


### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobs

> []Job ListJobs(ctx).Execute()

Get scheduled jobs



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
    resp, r, err := apiClient.SettingsApi.ListJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: []Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJobsRequest struct via the builder pattern


### Return type

[**[]Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLogs

> []ListLogs200ResponseInner ListLogs(ctx).Take(take).Skip(skip).Filter(filter).Search(search).Execute()

Returns logs



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
    take := float32(25) // float32 |  (optional)
    skip := float32(0) // float32 |  (optional)
    filter := "filter_example" // string |  (optional) (default to "debug")
    search := "plex" // string |  (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ListLogs(context.Background()).Take(take).Skip(skip).Filter(filter).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: []ListLogs200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **filter** | **string** |  | [default to &quot;debug&quot;]
 **search** | **string** |  | 

### Return type

[**[]ListLogs200ResponseInner**](ListLogs200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlexDevicesServers

> []PlexDevice ListPlexDevicesServers(ctx).Execute()

Gets the user's available Plex servers



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
    resp, r, err := apiClient.SettingsApi.ListPlexDevicesServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListPlexDevicesServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexDevicesServers`: []PlexDevice
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListPlexDevicesServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPlexDevicesServersRequest struct via the builder pattern


### Return type

[**[]PlexDevice**](PlexDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPlexLibrary

> []PlexLibrary ListPlexLibrary(ctx).Sync(sync).Enable(enable).Execute()

Get Plex libraries



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
    sync := "sync_example" // string | Syncs the current libraries with the current Plex server (optional)
    enable := "enable_example" // string | Comma separated list of libraries to enable. Any libraries not passed will be disabled! (optional)

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ListPlexLibrary(context.Background()).Sync(sync).Enable(enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListPlexLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexLibrary`: []PlexLibrary
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListPlexLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPlexLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sync** | **string** | Syncs the current libraries with the current Plex server | 
 **enable** | **string** | Comma separated list of libraries to enable. Any libraries not passed will be disabled! | 

### Return type

[**[]PlexLibrary**](PlexLibrary.md)

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
    resp, r, err := apiClient.SettingsApi.ListPlexUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListPlexUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexUsers`: []ListPlexUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListPlexUsers`: %v\n", resp)
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


## ListRadarr

> []RadarrSettings ListRadarr(ctx).Execute()

Get Radarr settings



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
    resp, r, err := apiClient.SettingsApi.ListRadarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRadarr`: []RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListRadarr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRadarrRequest struct via the builder pattern


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


## ListRadarrProfiles

> []ServiceProfile ListRadarrProfiles(ctx, radarrId).Execute()

Get available Radarr profiles



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
    radarrId := int32(56) // int32 | Radarr instance ID

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.ListRadarrProfiles(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListRadarrProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRadarrProfiles`: []ServiceProfile
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListRadarrProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radarrId** | **int32** | Radarr instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRadarrProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ServiceProfile**](ServiceProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSonarr

> []SonarrSettings ListSonarr(ctx).Execute()

Get Sonarr settings



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
    resp, r, err := apiClient.SettingsApi.ListSonarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.ListSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSonarr`: []SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.ListSonarr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSonarrRequest struct via the builder pattern


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


## TestNotificationsDiscord

> TestNotificationsDiscord(ctx).DiscordSettings(discordSettings).Execute()

Test Discord settings



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
    discordSettings := *overseerrClient.NewDiscordSettings() // DiscordSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsDiscord(context.Background()).DiscordSettings(discordSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsDiscord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsDiscordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discordSettings** | [**DiscordSettings**](DiscordSettings.md) |  | 

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


## TestNotificationsEmail

> TestNotificationsEmail(ctx).NotificationEmailSettings(notificationEmailSettings).Execute()

Test email settings



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
    notificationEmailSettings := *overseerrClient.NewNotificationEmailSettings() // NotificationEmailSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsEmail(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationEmailSettings** | [**NotificationEmailSettings**](NotificationEmailSettings.md) |  | 

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


## TestNotificationsGotify

> TestNotificationsGotify(ctx).GotifySettings(gotifySettings).Execute()

Test Gotify settings



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
    gotifySettings := *overseerrClient.NewGotifySettings() // GotifySettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsGotify(context.Background()).GotifySettings(gotifySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsGotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsGotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gotifySettings** | [**GotifySettings**](GotifySettings.md) |  | 

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


## TestNotificationsLunasea

> TestNotificationsLunasea(ctx).LunaSeaSettings(lunaSeaSettings).Execute()

Test LunaSea settings



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
    lunaSeaSettings := *overseerrClient.NewLunaSeaSettings() // LunaSeaSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsLunasea(context.Background()).LunaSeaSettings(lunaSeaSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsLunasea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsLunaseaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lunaSeaSettings** | [**LunaSeaSettings**](LunaSeaSettings.md) |  | 

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


## TestNotificationsPushbullet

> TestNotificationsPushbullet(ctx).PushbulletSettings(pushbulletSettings).Execute()

Test Pushbullet settings



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
    pushbulletSettings := *overseerrClient.NewPushbulletSettings() // PushbulletSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsPushbullet(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsPushbullet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsPushbulletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushbulletSettings** | [**PushbulletSettings**](PushbulletSettings.md) |  | 

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


## TestNotificationsPushover

> TestNotificationsPushover(ctx).PushoverSettings(pushoverSettings).Execute()

Test Pushover settings



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
    pushoverSettings := *overseerrClient.NewPushoverSettings() // PushoverSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsPushover(context.Background()).PushoverSettings(pushoverSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsPushover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsPushoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushoverSettings** | [**PushoverSettings**](PushoverSettings.md) |  | 

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


## TestNotificationsSlack

> TestNotificationsSlack(ctx).SlackSettings(slackSettings).Execute()

Test Slack settings



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
    slackSettings := *overseerrClient.NewSlackSettings() // SlackSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsSlack(context.Background()).SlackSettings(slackSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsSlackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slackSettings** | [**SlackSettings**](SlackSettings.md) |  | 

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


## TestNotificationsTelegram

> TestNotificationsTelegram(ctx).TelegramSettings(telegramSettings).Execute()

Test Telegram settings



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
    telegramSettings := *overseerrClient.NewTelegramSettings() // TelegramSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsTelegram(context.Background()).TelegramSettings(telegramSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsTelegram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsTelegramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telegramSettings** | [**TelegramSettings**](TelegramSettings.md) |  | 

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


## TestNotificationsWebhook

> TestNotificationsWebhook(ctx).WebhookSettings(webhookSettings).Execute()

Test webhook settings



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
    webhookSettings := *overseerrClient.NewWebhookSettings() // WebhookSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsWebhook(context.Background()).WebhookSettings(webhookSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookSettings** | [**WebhookSettings**](WebhookSettings.md) |  | 

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


## TestNotificationsWebpush

> TestNotificationsWebpush(ctx).WebPushSettings(webPushSettings).Execute()

Test Web Push settings



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
    webPushSettings := *overseerrClient.NewWebPushSettings() // WebPushSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestNotificationsWebpush(context.Background()).WebPushSettings(webPushSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestNotificationsWebpush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestNotificationsWebpushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webPushSettings** | [**WebPushSettings**](WebPushSettings.md) |  | 

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


## TestRadarr

> TestRadarr200Response TestRadarr(ctx).TestRadarrRequest(testRadarrRequest).Execute()

Test Radarr configuration



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
    testRadarrRequest := *overseerrClient.NewTestRadarrRequest("127.0.0.1", float32(7878), "yourapikey", false) // TestRadarrRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestRadarr(context.Background()).TestRadarrRequest(testRadarrRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestRadarr`: TestRadarr200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.TestRadarr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestRadarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testRadarrRequest** | [**TestRadarrRequest**](TestRadarrRequest.md) |  | 

### Return type

[**TestRadarr200Response**](TestRadarr200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSonarr

> TestRadarr200Response TestSonarr(ctx).TestSonarrRequest(testSonarrRequest).Execute()

Test Sonarr configuration



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
    testSonarrRequest := *overseerrClient.NewTestSonarrRequest("127.0.0.1", float32(8989), "yourapikey", false) // TestSonarrRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.TestSonarr(context.Background()).TestSonarrRequest(testSonarrRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.TestSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSonarr`: TestRadarr200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.TestSonarr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSonarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testSonarrRequest** | [**TestSonarrRequest**](TestSonarrRequest.md) |  | 

### Return type

[**TestRadarr200Response**](TestRadarr200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDiscover

> DiscoverSlider UpdateDiscover(ctx, sliderId).UpdateDiscoverRequest(updateDiscoverRequest).Execute()

Update a single slider



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
    sliderId := float32(8.14) // float32 | 
    updateDiscoverRequest := *overseerrClient.NewUpdateDiscoverRequest() // UpdateDiscoverRequest | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.UpdateDiscover(context.Background(), sliderId).UpdateDiscoverRequest(updateDiscoverRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.UpdateDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDiscover`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.UpdateDiscover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliderId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDiscoverRequest** | [**UpdateDiscoverRequest**](UpdateDiscoverRequest.md) |  | 

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRadarr

> RadarrSettings UpdateRadarr(ctx, radarrId).RadarrSettings(radarrSettings).Execute()

Update Radarr instance



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
    radarrId := int32(56) // int32 | Radarr instance ID
    radarrSettings := *overseerrClient.NewRadarrSettings("Radarr Main", "127.0.0.1", float32(7878), "exampleapikey", false, float32(1), "720p/1080p", "/movies", false, "In Cinema", false) // RadarrSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.UpdateRadarr(context.Background(), radarrId).RadarrSettings(radarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.UpdateRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.UpdateRadarr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**radarrId** | **int32** | Radarr instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRadarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **radarrSettings** | [**RadarrSettings**](RadarrSettings.md) |  | 

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSonarr

> SonarrSettings UpdateSonarr(ctx, sonarrId).SonarrSettings(sonarrSettings).Execute()

Update Sonarr instance



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
    sonarrId := int32(56) // int32 | Sonarr instance ID
    sonarrSettings := *overseerrClient.NewSonarrSettings("Sonarr Main", "127.0.0.1", float32(8989), "exampleapikey", false, float32(1), "720p/1080p", "/tv/", false, false, false) // SonarrSettings | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsApi.UpdateSonarr(context.Background(), sonarrId).SonarrSettings(sonarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsApi.UpdateSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsApi.UpdateSonarr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sonarrId** | **int32** | Sonarr instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSonarrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sonarrSettings** | [**SonarrSettings**](SonarrSettings.md) |  | 

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

