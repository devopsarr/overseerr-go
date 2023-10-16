# \SettingsAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCacheFlush**](SettingsAPI.md#CreateCacheFlush) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
[**CreateDiscover**](SettingsAPI.md#CreateDiscover) | **Post** /settings/discover | Batch update all sliders.
[**CreateDiscoverAdd**](SettingsAPI.md#CreateDiscoverAdd) | **Post** /settings/discover/add | Add a new slider
[**CreateInitialize**](SettingsAPI.md#CreateInitialize) | **Post** /settings/initialize | Initialize application
[**CreateJobsCancel**](SettingsAPI.md#CreateJobsCancel) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
[**CreateJobsRun**](SettingsAPI.md#CreateJobsRun) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
[**CreateJobsSchedule**](SettingsAPI.md#CreateJobsSchedule) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
[**CreateMain**](SettingsAPI.md#CreateMain) | **Post** /settings/main | Update main settings
[**CreateMainRegenerate**](SettingsAPI.md#CreateMainRegenerate) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
[**CreateNotificationsDiscord**](SettingsAPI.md#CreateNotificationsDiscord) | **Post** /settings/notifications/discord | Update Discord notification settings
[**CreateNotificationsEmail**](SettingsAPI.md#CreateNotificationsEmail) | **Post** /settings/notifications/email | Update email notification settings
[**CreateNotificationsGotify**](SettingsAPI.md#CreateNotificationsGotify) | **Post** /settings/notifications/gotify | Update Gotify notification settings
[**CreateNotificationsLunasea**](SettingsAPI.md#CreateNotificationsLunasea) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
[**CreateNotificationsPushbullet**](SettingsAPI.md#CreateNotificationsPushbullet) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
[**CreateNotificationsPushover**](SettingsAPI.md#CreateNotificationsPushover) | **Post** /settings/notifications/pushover | Update Pushover notification settings
[**CreateNotificationsSlack**](SettingsAPI.md#CreateNotificationsSlack) | **Post** /settings/notifications/slack | Update Slack notification settings
[**CreateNotificationsTelegram**](SettingsAPI.md#CreateNotificationsTelegram) | **Post** /settings/notifications/telegram | Update Telegram notification settings
[**CreateNotificationsWebhook**](SettingsAPI.md#CreateNotificationsWebhook) | **Post** /settings/notifications/webhook | Update webhook notification settings
[**CreateNotificationsWebpush**](SettingsAPI.md#CreateNotificationsWebpush) | **Post** /settings/notifications/webpush | Update Web Push notification settings
[**CreatePlex**](SettingsAPI.md#CreatePlex) | **Post** /settings/plex | Update Plex settings
[**CreatePlexSync**](SettingsAPI.md#CreatePlexSync) | **Post** /settings/plex/sync | Start full Plex library scan
[**CreateRadarr**](SettingsAPI.md#CreateRadarr) | **Post** /settings/radarr | Create Radarr instance
[**CreateSonarr**](SettingsAPI.md#CreateSonarr) | **Post** /settings/sonarr | Create Sonarr instance
[**CreateTautulli**](SettingsAPI.md#CreateTautulli) | **Post** /settings/tautulli | Update Tautulli settings
[**DeleteDiscover**](SettingsAPI.md#DeleteDiscover) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
[**DeleteRadarr**](SettingsAPI.md#DeleteRadarr) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
[**DeleteSonarr**](SettingsAPI.md#DeleteSonarr) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
[**GetAbout**](SettingsAPI.md#GetAbout) | **Get** /settings/about | Get server stats
[**GetCache**](SettingsAPI.md#GetCache) | **Get** /settings/cache | Get a list of active caches
[**GetDiscoverReset**](SettingsAPI.md#GetDiscoverReset) | **Get** /settings/discover/reset | Reset all discover sliders
[**GetMain**](SettingsAPI.md#GetMain) | **Get** /settings/main | Get main settings
[**GetNotificationsDiscord**](SettingsAPI.md#GetNotificationsDiscord) | **Get** /settings/notifications/discord | Get Discord notification settings
[**GetNotificationsEmail**](SettingsAPI.md#GetNotificationsEmail) | **Get** /settings/notifications/email | Get email notification settings
[**GetNotificationsGotify**](SettingsAPI.md#GetNotificationsGotify) | **Get** /settings/notifications/gotify | Get Gotify notification settings
[**GetNotificationsLunasea**](SettingsAPI.md#GetNotificationsLunasea) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
[**GetNotificationsPushbullet**](SettingsAPI.md#GetNotificationsPushbullet) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
[**GetNotificationsPushover**](SettingsAPI.md#GetNotificationsPushover) | **Get** /settings/notifications/pushover | Get Pushover notification settings
[**GetNotificationsSlack**](SettingsAPI.md#GetNotificationsSlack) | **Get** /settings/notifications/slack | Get Slack notification settings
[**GetNotificationsTelegram**](SettingsAPI.md#GetNotificationsTelegram) | **Get** /settings/notifications/telegram | Get Telegram notification settings
[**GetNotificationsWebhook**](SettingsAPI.md#GetNotificationsWebhook) | **Get** /settings/notifications/webhook | Get webhook notification settings
[**GetNotificationsWebpush**](SettingsAPI.md#GetNotificationsWebpush) | **Get** /settings/notifications/webpush | Get Web Push notification settings
[**GetPlex**](SettingsAPI.md#GetPlex) | **Get** /settings/plex | Get Plex settings
[**GetPlexSync**](SettingsAPI.md#GetPlexSync) | **Get** /settings/plex/sync | Get status of full Plex library scan
[**GetPublic**](SettingsAPI.md#GetPublic) | **Get** /settings/public | Get public settings
[**GetTautulli**](SettingsAPI.md#GetTautulli) | **Get** /settings/tautulli | Get Tautulli settings
[**ListDiscover**](SettingsAPI.md#ListDiscover) | **Get** /settings/discover | Get all discover sliders
[**ListJobs**](SettingsAPI.md#ListJobs) | **Get** /settings/jobs | Get scheduled jobs
[**ListLogs**](SettingsAPI.md#ListLogs) | **Get** /settings/logs | Returns logs
[**ListNotificationsPushoverSounds**](SettingsAPI.md#ListNotificationsPushoverSounds) | **Get** /settings/notifications/pushover/sounds | Get Pushover sounds
[**ListPlexDevicesServers**](SettingsAPI.md#ListPlexDevicesServers) | **Get** /settings/plex/devices/servers | Gets the user&#39;s available Plex servers
[**ListPlexLibrary**](SettingsAPI.md#ListPlexLibrary) | **Get** /settings/plex/library | Get Plex libraries
[**ListPlexUsers**](SettingsAPI.md#ListPlexUsers) | **Get** /settings/plex/users | Get Plex users
[**ListRadarr**](SettingsAPI.md#ListRadarr) | **Get** /settings/radarr | Get Radarr settings
[**ListRadarrProfiles**](SettingsAPI.md#ListRadarrProfiles) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
[**ListSonarr**](SettingsAPI.md#ListSonarr) | **Get** /settings/sonarr | Get Sonarr settings
[**TestNotificationsDiscord**](SettingsAPI.md#TestNotificationsDiscord) | **Post** /settings/notifications/discord/test | Test Discord settings
[**TestNotificationsEmail**](SettingsAPI.md#TestNotificationsEmail) | **Post** /settings/notifications/email/test | Test email settings
[**TestNotificationsGotify**](SettingsAPI.md#TestNotificationsGotify) | **Post** /settings/notifications/gotify/test | Test Gotify settings
[**TestNotificationsLunasea**](SettingsAPI.md#TestNotificationsLunasea) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
[**TestNotificationsPushbullet**](SettingsAPI.md#TestNotificationsPushbullet) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
[**TestNotificationsPushover**](SettingsAPI.md#TestNotificationsPushover) | **Post** /settings/notifications/pushover/test | Test Pushover settings
[**TestNotificationsSlack**](SettingsAPI.md#TestNotificationsSlack) | **Post** /settings/notifications/slack/test | Test Slack settings
[**TestNotificationsTelegram**](SettingsAPI.md#TestNotificationsTelegram) | **Post** /settings/notifications/telegram/test | Test Telegram settings
[**TestNotificationsWebhook**](SettingsAPI.md#TestNotificationsWebhook) | **Post** /settings/notifications/webhook/test | Test webhook settings
[**TestNotificationsWebpush**](SettingsAPI.md#TestNotificationsWebpush) | **Post** /settings/notifications/webpush/test | Test Web Push settings
[**TestRadarr**](SettingsAPI.md#TestRadarr) | **Post** /settings/radarr/test | Test Radarr configuration
[**TestSonarr**](SettingsAPI.md#TestSonarr) | **Post** /settings/sonarr/test | Test Sonarr configuration
[**UpdateDiscover**](SettingsAPI.md#UpdateDiscover) | **Put** /settings/discover/{sliderId} | Update a single slider
[**UpdateRadarr**](SettingsAPI.md#UpdateRadarr) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
[**UpdateSonarr**](SettingsAPI.md#UpdateSonarr) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance



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
    resp, r, err := apiClient.SettingsAPI.CreateCacheFlush(context.Background(), cacheId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateCacheFlush``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.CreateDiscover(context.Background()).DiscoverSlider(discoverSlider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscover`: []DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateDiscover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateDiscoverAdd(context.Background()).CreateDiscoverAddRequest(createDiscoverAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateDiscoverAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscoverAdd`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateDiscoverAdd`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateInitialize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateInitialize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInitialize`: PublicSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateInitialize`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateJobsCancel(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateJobsCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsCancel`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateJobsCancel`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateJobsRun(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateJobsRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsRun`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateJobsRun`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateJobsSchedule(context.Background(), jobId).CreateJobsScheduleRequest(createJobsScheduleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateJobsSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateJobsSchedule`: Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateJobsSchedule`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateMain(context.Background()).MainSettings(mainSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMain`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateMain`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateMainRegenerate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateMainRegenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMainRegenerate`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateMainRegenerate`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsDiscord(context.Background()).DiscordSettings(discordSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsDiscord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsDiscord`: DiscordSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsDiscord`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsEmail(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsEmail`: NotificationEmailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsEmail`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsGotify(context.Background()).GotifySettings(gotifySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsGotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsGotify`: GotifySettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsGotify`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsLunasea(context.Background()).LunaSeaSettings(lunaSeaSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsLunasea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsLunasea`: LunaSeaSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsLunasea`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsPushbullet(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsPushbullet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsPushbullet`: PushbulletSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsPushbullet`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsPushover(context.Background()).PushoverSettings(pushoverSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsPushover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsPushover`: PushoverSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsPushover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsSlack(context.Background()).SlackSettings(slackSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsSlack`: SlackSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsSlack`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsTelegram(context.Background()).TelegramSettings(telegramSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsTelegram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsTelegram`: TelegramSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsTelegram`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsWebhook(context.Background()).WebhookSettings(webhookSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsWebhook`: WebhookSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsWebhook`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateNotificationsWebpush(context.Background()).WebPushSettings(webPushSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateNotificationsWebpush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsWebpush`: WebPushSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateNotificationsWebpush`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreatePlex(context.Background()).PlexSettings(plexSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreatePlex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlex`: PlexSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreatePlex`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreatePlexSync(context.Background()).CreatePlexSyncRequest(createPlexSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreatePlexSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlexSync`: GetPlexSync200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreatePlexSync`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateRadarr(context.Background()).RadarrSettings(radarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateSonarr(context.Background()).SonarrSettings(sonarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateSonarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.CreateTautulli(context.Background()).TautulliSettings(tautulliSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CreateTautulli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTautulli`: TautulliSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CreateTautulli`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.DeleteDiscover(context.Background(), sliderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDiscover`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.DeleteDiscover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.DeleteRadarr(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.DeleteRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.DeleteSonarr(context.Background(), sonarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.DeleteSonarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetAbout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetAbout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAbout`: GetAbout200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetAbout`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetCache(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCache`: GetCache200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetCache`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetDiscoverReset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetDiscoverReset``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.GetMain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetMain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMain`: MainSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetMain`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsDiscord(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsDiscord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsDiscord`: DiscordSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsDiscord`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsEmail(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsEmail`: NotificationEmailSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsEmail`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsGotify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsGotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsGotify`: GotifySettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsGotify`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsLunasea(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsLunasea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsLunasea`: LunaSeaSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsLunasea`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsPushbullet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsPushbullet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsPushbullet`: PushbulletSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsPushbullet`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsPushover(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsPushover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsPushover`: PushoverSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsPushover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsSlack(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsSlack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsSlack`: SlackSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsSlack`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsTelegram(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsTelegram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsTelegram`: TelegramSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsTelegram`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsWebhook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsWebhook`: WebhookSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsWebhook`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetNotificationsWebpush(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetNotificationsWebpush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationsWebpush`: WebPushSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetNotificationsWebpush`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetPlex(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetPlex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlex`: PlexSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetPlex`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetPlexSync(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetPlexSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlexSync`: GetPlexSync200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetPlexSync`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetPublic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublic`: PublicSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetPublic`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.GetTautulli(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetTautulli``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTautulli`: TautulliSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetTautulli`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListDiscover(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscover`: []DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListDiscover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobs`: []Job
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListJobs`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListLogs(context.Background()).Take(take).Skip(skip).Filter(filter).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLogs`: []ListLogs200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListLogs`: %v\n", resp)
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


## ListNotificationsPushoverSounds

> []ListNotificationsPushoverSounds200ResponseInner ListNotificationsPushoverSounds(ctx).Token(token).Execute()

Get Pushover sounds



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
    token := "token_example" // string | 

    configuration := overseerrClient.NewConfiguration()
    apiClient := overseerrClient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsAPI.ListNotificationsPushoverSounds(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListNotificationsPushoverSounds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationsPushoverSounds`: []ListNotificationsPushoverSounds200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListNotificationsPushoverSounds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsPushoverSoundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** |  | 

### Return type

[**[]ListNotificationsPushoverSounds200ResponseInner**](ListNotificationsPushoverSounds200ResponseInner.md)

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
    resp, r, err := apiClient.SettingsAPI.ListPlexDevicesServers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListPlexDevicesServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexDevicesServers`: []PlexDevice
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListPlexDevicesServers`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListPlexLibrary(context.Background()).Sync(sync).Enable(enable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListPlexLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexLibrary`: []PlexLibrary
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListPlexLibrary`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListPlexUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListPlexUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPlexUsers`: []ListPlexUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListPlexUsers`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListRadarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRadarr`: []RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListRadarrProfiles(context.Background(), radarrId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListRadarrProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRadarrProfiles`: []ServiceProfile
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListRadarrProfiles`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.ListSonarr(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.ListSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSonarr`: []SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.ListSonarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsDiscord(context.Background()).DiscordSettings(discordSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsDiscord``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsEmail(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsEmail``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsGotify(context.Background()).GotifySettings(gotifySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsGotify``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsLunasea(context.Background()).LunaSeaSettings(lunaSeaSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsLunasea``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsPushbullet(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsPushbullet``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsPushover(context.Background()).PushoverSettings(pushoverSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsPushover``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsSlack(context.Background()).SlackSettings(slackSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsSlack``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsTelegram(context.Background()).TelegramSettings(telegramSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsTelegram``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsWebhook(context.Background()).WebhookSettings(webhookSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsWebhook``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestNotificationsWebpush(context.Background()).WebPushSettings(webPushSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestNotificationsWebpush``: %v\n", err)
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
    resp, r, err := apiClient.SettingsAPI.TestRadarr(context.Background()).TestRadarrRequest(testRadarrRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestRadarr`: TestRadarr200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.TestRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.TestSonarr(context.Background()).TestSonarrRequest(testSonarrRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.TestSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSonarr`: TestRadarr200Response
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.TestSonarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.UpdateDiscover(context.Background(), sliderId).UpdateDiscoverRequest(updateDiscoverRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateDiscover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDiscover`: DiscoverSlider
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateDiscover`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.UpdateRadarr(context.Background(), radarrId).RadarrSettings(radarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateRadarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRadarr`: RadarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateRadarr`: %v\n", resp)
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
    resp, r, err := apiClient.SettingsAPI.UpdateSonarr(context.Background(), sonarrId).SonarrSettings(sonarrSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.UpdateSonarr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSonarr`: SonarrSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.UpdateSonarr`: %v\n", resp)
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

