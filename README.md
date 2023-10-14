# Go API client for overseerr

This is the documentation for the Overseerr API backend.

Two primary authentication methods are supported:

- **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie.
- **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

[comment]: # (x-release-please-start-version)
- Package version: 0.0.0

[comment]: # (x-release-please-end)
- API version: 1.0.0

- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import overseerr "github.com/devopsarr/overseerr-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), overseerr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), overseerr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), overseerr.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), overseerr.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:5055/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthAPI* | [**CreateAuthLocal**](overseerr/docs/AuthAPI.md#createauthlocal) | **Post** /auth/local | Sign in using a local account
*AuthAPI* | [**CreateAuthLogout**](overseerr/docs/AuthAPI.md#createauthlogout) | **Post** /auth/logout | Sign out and clear session cookie
*AuthAPI* | [**CreateAuthPlex**](overseerr/docs/AuthAPI.md#createauthplex) | **Post** /auth/plex | Sign in using a Plex token
*AuthAPI* | [**GetAuthMe**](overseerr/docs/AuthAPI.md#getauthme) | **Get** /auth/me | Get logged-in user
*CollectionAPI* | [**GetCollectionByCollectionId**](overseerr/docs/CollectionAPI.md#getcollectionbycollectionid) | **Get** /collection/{collectionId} | Get collection details
*IssueAPI* | [**CreateIssue**](overseerr/docs/IssueAPI.md#createissue) | **Post** /issue | Create new issue
*IssueAPI* | [**CreateIssueByStatus**](overseerr/docs/IssueAPI.md#createissuebystatus) | **Post** /issue/{issueId}/{status} | Update an issue&#39;s status
*IssueAPI* | [**CreateIssueComment**](overseerr/docs/IssueAPI.md#createissuecomment) | **Post** /issue/{issueId}/comment | Create a comment
*IssueAPI* | [**DeleteIssue**](overseerr/docs/IssueAPI.md#deleteissue) | **Delete** /issue/{issueId} | Delete issue
*IssueAPI* | [**DeleteIssueComment**](overseerr/docs/IssueAPI.md#deleteissuecomment) | **Delete** /issueComment/{commentId} | Delete issue comment
*IssueAPI* | [**GetIssue**](overseerr/docs/IssueAPI.md#getissue) | **Get** /issue | Get all issues
*IssueAPI* | [**GetIssueByIssueId**](overseerr/docs/IssueAPI.md#getissuebyissueid) | **Get** /issue/{issueId} | Get issue
*IssueAPI* | [**GetIssueCommentByCommentId**](overseerr/docs/IssueAPI.md#getissuecommentbycommentid) | **Get** /issueComment/{commentId} | Get issue comment
*IssueAPI* | [**GetIssueCount**](overseerr/docs/IssueAPI.md#getissuecount) | **Get** /issue/count | Gets issue counts
*IssueAPI* | [**UpdateIssueComment**](overseerr/docs/IssueAPI.md#updateissuecomment) | **Put** /issueComment/{commentId} | Update issue comment
*MediaAPI* | [**CreateMediaByStatus**](overseerr/docs/MediaAPI.md#createmediabystatus) | **Post** /media/{mediaId}/{status} | Update media status
*MediaAPI* | [**DeleteMedia**](overseerr/docs/MediaAPI.md#deletemedia) | **Delete** /media/{mediaId} | Delete media item
*MediaAPI* | [**GetMedia**](overseerr/docs/MediaAPI.md#getmedia) | **Get** /media | Get media
*MediaAPI* | [**GetMediaWatchData**](overseerr/docs/MediaAPI.md#getmediawatchdata) | **Get** /media/{mediaId}/watch_data | Get watch data
*MoviesAPI* | [**GetMovieByMovieId**](overseerr/docs/MoviesAPI.md#getmoviebymovieid) | **Get** /movie/{movieId} | Get movie details
*MoviesAPI* | [**GetMovieRatings**](overseerr/docs/MoviesAPI.md#getmovieratings) | **Get** /movie/{movieId}/ratings | Get movie ratings
*MoviesAPI* | [**GetMovieRatingscombined**](overseerr/docs/MoviesAPI.md#getmovieratingscombined) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
*MoviesAPI* | [**GetMovieRecommendations**](overseerr/docs/MoviesAPI.md#getmovierecommendations) | **Get** /movie/{movieId}/recommendations | Get recommended movies
*MoviesAPI* | [**GetMovieSimilar**](overseerr/docs/MoviesAPI.md#getmoviesimilar) | **Get** /movie/{movieId}/similar | Get similar movies
*OtherAPI* | [**GetKeywordByKeywordId**](overseerr/docs/OtherAPI.md#getkeywordbykeywordid) | **Get** /keyword/{keywordId} | Get keyword
*OtherAPI* | [**ListWatchprovidersMovies**](overseerr/docs/OtherAPI.md#listwatchprovidersmovies) | **Get** /watchproviders/movies | Get watch provider movies
*OtherAPI* | [**ListWatchprovidersRegions**](overseerr/docs/OtherAPI.md#listwatchprovidersregions) | **Get** /watchproviders/regions | Get watch provider regions
*OtherAPI* | [**ListWatchprovidersTv**](overseerr/docs/OtherAPI.md#listwatchproviderstv) | **Get** /watchproviders/tv | Get watch provider series
*PersonAPI* | [**GetPersonByPersonId**](overseerr/docs/PersonAPI.md#getpersonbypersonid) | **Get** /person/{personId} | Get person details
*PersonAPI* | [**GetPersonCombinedCredits**](overseerr/docs/PersonAPI.md#getpersoncombinedcredits) | **Get** /person/{personId}/combined_credits | Get combined credits
*PublicAPI* | [**GetStatus**](overseerr/docs/PublicAPI.md#getstatus) | **Get** /status | Get Overseerr status
*PublicAPI* | [**GetStatusAppdata**](overseerr/docs/PublicAPI.md#getstatusappdata) | **Get** /status/appdata | Get application data volume status
*RequestAPI* | [**CreateRequest**](overseerr/docs/RequestAPI.md#createrequest) | **Post** /request | Create new request
*RequestAPI* | [**CreateRequestByStatus**](overseerr/docs/RequestAPI.md#createrequestbystatus) | **Post** /request/{requestId}/{status} | Update a request&#39;s status
*RequestAPI* | [**CreateRequestRetry**](overseerr/docs/RequestAPI.md#createrequestretry) | **Post** /request/{requestId}/retry | Retry failed request
*RequestAPI* | [**DeleteRequest**](overseerr/docs/RequestAPI.md#deleterequest) | **Delete** /request/{requestId} | Delete request
*RequestAPI* | [**GetRequest**](overseerr/docs/RequestAPI.md#getrequest) | **Get** /request | Get all requests
*RequestAPI* | [**GetRequestByRequestId**](overseerr/docs/RequestAPI.md#getrequestbyrequestid) | **Get** /request/{requestId} | Get MediaRequest
*RequestAPI* | [**GetRequestCount**](overseerr/docs/RequestAPI.md#getrequestcount) | **Get** /request/count | Gets request counts
*RequestAPI* | [**UpdateRequest**](overseerr/docs/RequestAPI.md#updaterequest) | **Put** /request/{requestId} | Update MediaRequest
*SearchAPI* | [**GetDiscoverKeywordMovies**](overseerr/docs/SearchAPI.md#getdiscoverkeywordmovies) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
*SearchAPI* | [**GetDiscoverMovies**](overseerr/docs/SearchAPI.md#getdiscovermovies) | **Get** /discover/movies | Discover movies
*SearchAPI* | [**GetDiscoverMoviesGenreByGenreId**](overseerr/docs/SearchAPI.md#getdiscovermoviesgenrebygenreid) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
*SearchAPI* | [**GetDiscoverMoviesLanguageByLanguage**](overseerr/docs/SearchAPI.md#getdiscovermovieslanguagebylanguage) | **Get** /discover/movies/language/{language} | Discover movies by original language
*SearchAPI* | [**GetDiscoverMoviesStudioByStudioId**](overseerr/docs/SearchAPI.md#getdiscovermoviesstudiobystudioid) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
*SearchAPI* | [**GetDiscoverMoviesUpcoming**](overseerr/docs/SearchAPI.md#getdiscovermoviesupcoming) | **Get** /discover/movies/upcoming | Upcoming movies
*SearchAPI* | [**GetDiscoverTrending**](overseerr/docs/SearchAPI.md#getdiscovertrending) | **Get** /discover/trending | Trending movies and TV
*SearchAPI* | [**GetDiscoverTv**](overseerr/docs/SearchAPI.md#getdiscovertv) | **Get** /discover/tv | Discover TV shows
*SearchAPI* | [**GetDiscoverTvGenreByGenreId**](overseerr/docs/SearchAPI.md#getdiscovertvgenrebygenreid) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
*SearchAPI* | [**GetDiscoverTvLanguageByLanguage**](overseerr/docs/SearchAPI.md#getdiscovertvlanguagebylanguage) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
*SearchAPI* | [**GetDiscoverTvNetworkByNetworkId**](overseerr/docs/SearchAPI.md#getdiscovertvnetworkbynetworkid) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
*SearchAPI* | [**GetDiscoverTvUpcoming**](overseerr/docs/SearchAPI.md#getdiscovertvupcoming) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
*SearchAPI* | [**GetDiscoverWatchlist**](overseerr/docs/SearchAPI.md#getdiscoverwatchlist) | **Get** /discover/watchlist | Get the Plex watchlist.
*SearchAPI* | [**GetSearch**](overseerr/docs/SearchAPI.md#getsearch) | **Get** /search | Search for movies, TV shows, or people
*SearchAPI* | [**GetSearchCompany**](overseerr/docs/SearchAPI.md#getsearchcompany) | **Get** /search/company | Search for companies
*SearchAPI* | [**GetSearchKeyword**](overseerr/docs/SearchAPI.md#getsearchkeyword) | **Get** /search/keyword | Search for keywords
*SearchAPI* | [**ListDiscoverGenresliderMovie**](overseerr/docs/SearchAPI.md#listdiscovergenreslidermovie) | **Get** /discover/genreslider/movie | Get genre slider data for movies
*SearchAPI* | [**ListDiscoverGenresliderTv**](overseerr/docs/SearchAPI.md#listdiscovergenreslidertv) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
*ServiceAPI* | [**GetServiceRadarrByRadarrId**](overseerr/docs/ServiceAPI.md#getserviceradarrbyradarrid) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
*ServiceAPI* | [**GetServiceSonarrBySonarrId**](overseerr/docs/ServiceAPI.md#getservicesonarrbysonarrid) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
*ServiceAPI* | [**ListServiceRadarr**](overseerr/docs/ServiceAPI.md#listserviceradarr) | **Get** /service/radarr | Get non-sensitive Radarr server list
*ServiceAPI* | [**ListServiceSonarr**](overseerr/docs/ServiceAPI.md#listservicesonarr) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
*ServiceAPI* | [**ListServiceSonarrLookupByTmdbId**](overseerr/docs/ServiceAPI.md#listservicesonarrlookupbytmdbid) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr
*SettingsAPI* | [**CreateCacheFlush**](overseerr/docs/SettingsAPI.md#createcacheflush) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
*SettingsAPI* | [**CreateDiscover**](overseerr/docs/SettingsAPI.md#creatediscover) | **Post** /settings/discover | Batch update all sliders.
*SettingsAPI* | [**CreateDiscoverAdd**](overseerr/docs/SettingsAPI.md#creatediscoveradd) | **Post** /settings/discover/add | Add a new slider
*SettingsAPI* | [**CreateInitialize**](overseerr/docs/SettingsAPI.md#createinitialize) | **Post** /settings/initialize | Initialize application
*SettingsAPI* | [**CreateJobsCancel**](overseerr/docs/SettingsAPI.md#createjobscancel) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
*SettingsAPI* | [**CreateJobsRun**](overseerr/docs/SettingsAPI.md#createjobsrun) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
*SettingsAPI* | [**CreateJobsSchedule**](overseerr/docs/SettingsAPI.md#createjobsschedule) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
*SettingsAPI* | [**CreateMain**](overseerr/docs/SettingsAPI.md#createmain) | **Post** /settings/main | Update main settings
*SettingsAPI* | [**CreateMainRegenerate**](overseerr/docs/SettingsAPI.md#createmainregenerate) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
*SettingsAPI* | [**CreateNotificationsDiscord**](overseerr/docs/SettingsAPI.md#createnotificationsdiscord) | **Post** /settings/notifications/discord | Update Discord notification settings
*SettingsAPI* | [**CreateNotificationsEmail**](overseerr/docs/SettingsAPI.md#createnotificationsemail) | **Post** /settings/notifications/email | Update email notification settings
*SettingsAPI* | [**CreateNotificationsGotify**](overseerr/docs/SettingsAPI.md#createnotificationsgotify) | **Post** /settings/notifications/gotify | Update Gotify notification settings
*SettingsAPI* | [**CreateNotificationsLunasea**](overseerr/docs/SettingsAPI.md#createnotificationslunasea) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
*SettingsAPI* | [**CreateNotificationsPushbullet**](overseerr/docs/SettingsAPI.md#createnotificationspushbullet) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
*SettingsAPI* | [**CreateNotificationsPushover**](overseerr/docs/SettingsAPI.md#createnotificationspushover) | **Post** /settings/notifications/pushover | Update Pushover notification settings
*SettingsAPI* | [**CreateNotificationsSlack**](overseerr/docs/SettingsAPI.md#createnotificationsslack) | **Post** /settings/notifications/slack | Update Slack notification settings
*SettingsAPI* | [**CreateNotificationsTelegram**](overseerr/docs/SettingsAPI.md#createnotificationstelegram) | **Post** /settings/notifications/telegram | Update Telegram notification settings
*SettingsAPI* | [**CreateNotificationsWebhook**](overseerr/docs/SettingsAPI.md#createnotificationswebhook) | **Post** /settings/notifications/webhook | Update webhook notification settings
*SettingsAPI* | [**CreateNotificationsWebpush**](overseerr/docs/SettingsAPI.md#createnotificationswebpush) | **Post** /settings/notifications/webpush | Update Web Push notification settings
*SettingsAPI* | [**CreatePlex**](overseerr/docs/SettingsAPI.md#createplex) | **Post** /settings/plex | Update Plex settings
*SettingsAPI* | [**CreatePlexSync**](overseerr/docs/SettingsAPI.md#createplexsync) | **Post** /settings/plex/sync | Start full Plex library scan
*SettingsAPI* | [**CreateRadarr**](overseerr/docs/SettingsAPI.md#createradarr) | **Post** /settings/radarr | Create Radarr instance
*SettingsAPI* | [**CreateSonarr**](overseerr/docs/SettingsAPI.md#createsonarr) | **Post** /settings/sonarr | Create Sonarr instance
*SettingsAPI* | [**CreateTautulli**](overseerr/docs/SettingsAPI.md#createtautulli) | **Post** /settings/tautulli | Update Tautulli settings
*SettingsAPI* | [**DeleteDiscover**](overseerr/docs/SettingsAPI.md#deletediscover) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
*SettingsAPI* | [**DeleteRadarr**](overseerr/docs/SettingsAPI.md#deleteradarr) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
*SettingsAPI* | [**DeleteSonarr**](overseerr/docs/SettingsAPI.md#deletesonarr) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
*SettingsAPI* | [**GetAbout**](overseerr/docs/SettingsAPI.md#getabout) | **Get** /settings/about | Get server stats
*SettingsAPI* | [**GetCache**](overseerr/docs/SettingsAPI.md#getcache) | **Get** /settings/cache | Get a list of active caches
*SettingsAPI* | [**GetDiscoverReset**](overseerr/docs/SettingsAPI.md#getdiscoverreset) | **Get** /settings/discover/reset | Reset all discover sliders
*SettingsAPI* | [**GetMain**](overseerr/docs/SettingsAPI.md#getmain) | **Get** /settings/main | Get main settings
*SettingsAPI* | [**GetNotificationsDiscord**](overseerr/docs/SettingsAPI.md#getnotificationsdiscord) | **Get** /settings/notifications/discord | Get Discord notification settings
*SettingsAPI* | [**GetNotificationsEmail**](overseerr/docs/SettingsAPI.md#getnotificationsemail) | **Get** /settings/notifications/email | Get email notification settings
*SettingsAPI* | [**GetNotificationsGotify**](overseerr/docs/SettingsAPI.md#getnotificationsgotify) | **Get** /settings/notifications/gotify | Get Gotify notification settings
*SettingsAPI* | [**GetNotificationsLunasea**](overseerr/docs/SettingsAPI.md#getnotificationslunasea) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
*SettingsAPI* | [**GetNotificationsPushbullet**](overseerr/docs/SettingsAPI.md#getnotificationspushbullet) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
*SettingsAPI* | [**GetNotificationsPushover**](overseerr/docs/SettingsAPI.md#getnotificationspushover) | **Get** /settings/notifications/pushover | Get Pushover notification settings
*SettingsAPI* | [**GetNotificationsSlack**](overseerr/docs/SettingsAPI.md#getnotificationsslack) | **Get** /settings/notifications/slack | Get Slack notification settings
*SettingsAPI* | [**GetNotificationsTelegram**](overseerr/docs/SettingsAPI.md#getnotificationstelegram) | **Get** /settings/notifications/telegram | Get Telegram notification settings
*SettingsAPI* | [**GetNotificationsWebhook**](overseerr/docs/SettingsAPI.md#getnotificationswebhook) | **Get** /settings/notifications/webhook | Get webhook notification settings
*SettingsAPI* | [**GetNotificationsWebpush**](overseerr/docs/SettingsAPI.md#getnotificationswebpush) | **Get** /settings/notifications/webpush | Get Web Push notification settings
*SettingsAPI* | [**GetPlex**](overseerr/docs/SettingsAPI.md#getplex) | **Get** /settings/plex | Get Plex settings
*SettingsAPI* | [**GetPlexSync**](overseerr/docs/SettingsAPI.md#getplexsync) | **Get** /settings/plex/sync | Get status of full Plex library scan
*SettingsAPI* | [**GetPublic**](overseerr/docs/SettingsAPI.md#getpublic) | **Get** /settings/public | Get public settings
*SettingsAPI* | [**GetTautulli**](overseerr/docs/SettingsAPI.md#gettautulli) | **Get** /settings/tautulli | Get Tautulli settings
*SettingsAPI* | [**ListDiscover**](overseerr/docs/SettingsAPI.md#listdiscover) | **Get** /settings/discover | Get all discover sliders
*SettingsAPI* | [**ListJobs**](overseerr/docs/SettingsAPI.md#listjobs) | **Get** /settings/jobs | Get scheduled jobs
*SettingsAPI* | [**ListLogs**](overseerr/docs/SettingsAPI.md#listlogs) | **Get** /settings/logs | Returns logs
*SettingsAPI* | [**ListPlexDevicesServers**](overseerr/docs/SettingsAPI.md#listplexdevicesservers) | **Get** /settings/plex/devices/servers | Gets the user&#39;s available Plex servers
*SettingsAPI* | [**ListPlexLibrary**](overseerr/docs/SettingsAPI.md#listplexlibrary) | **Get** /settings/plex/library | Get Plex libraries
*SettingsAPI* | [**ListPlexUsers**](overseerr/docs/SettingsAPI.md#listplexusers) | **Get** /settings/plex/users | Get Plex users
*SettingsAPI* | [**ListRadarr**](overseerr/docs/SettingsAPI.md#listradarr) | **Get** /settings/radarr | Get Radarr settings
*SettingsAPI* | [**ListRadarrProfiles**](overseerr/docs/SettingsAPI.md#listradarrprofiles) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
*SettingsAPI* | [**ListSonarr**](overseerr/docs/SettingsAPI.md#listsonarr) | **Get** /settings/sonarr | Get Sonarr settings
*SettingsAPI* | [**TestNotificationsDiscord**](overseerr/docs/SettingsAPI.md#testnotificationsdiscord) | **Post** /settings/notifications/discord/test | Test Discord settings
*SettingsAPI* | [**TestNotificationsEmail**](overseerr/docs/SettingsAPI.md#testnotificationsemail) | **Post** /settings/notifications/email/test | Test email settings
*SettingsAPI* | [**TestNotificationsGotify**](overseerr/docs/SettingsAPI.md#testnotificationsgotify) | **Post** /settings/notifications/gotify/test | Test Gotify settings
*SettingsAPI* | [**TestNotificationsLunasea**](overseerr/docs/SettingsAPI.md#testnotificationslunasea) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
*SettingsAPI* | [**TestNotificationsPushbullet**](overseerr/docs/SettingsAPI.md#testnotificationspushbullet) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
*SettingsAPI* | [**TestNotificationsPushover**](overseerr/docs/SettingsAPI.md#testnotificationspushover) | **Post** /settings/notifications/pushover/test | Test Pushover settings
*SettingsAPI* | [**TestNotificationsSlack**](overseerr/docs/SettingsAPI.md#testnotificationsslack) | **Post** /settings/notifications/slack/test | Test Slack settings
*SettingsAPI* | [**TestNotificationsTelegram**](overseerr/docs/SettingsAPI.md#testnotificationstelegram) | **Post** /settings/notifications/telegram/test | Test Telegram settings
*SettingsAPI* | [**TestNotificationsWebhook**](overseerr/docs/SettingsAPI.md#testnotificationswebhook) | **Post** /settings/notifications/webhook/test | Test webhook settings
*SettingsAPI* | [**TestNotificationsWebpush**](overseerr/docs/SettingsAPI.md#testnotificationswebpush) | **Post** /settings/notifications/webpush/test | Test Web Push settings
*SettingsAPI* | [**TestRadarr**](overseerr/docs/SettingsAPI.md#testradarr) | **Post** /settings/radarr/test | Test Radarr configuration
*SettingsAPI* | [**TestSonarr**](overseerr/docs/SettingsAPI.md#testsonarr) | **Post** /settings/sonarr/test | Test Sonarr configuration
*SettingsAPI* | [**UpdateDiscover**](overseerr/docs/SettingsAPI.md#updatediscover) | **Put** /settings/discover/{sliderId} | Update a single slider
*SettingsAPI* | [**UpdateRadarr**](overseerr/docs/SettingsAPI.md#updateradarr) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
*SettingsAPI* | [**UpdateSonarr**](overseerr/docs/SettingsAPI.md#updatesonarr) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance
*TmdbAPI* | [**GetNetworkByNetworkId**](overseerr/docs/TmdbAPI.md#getnetworkbynetworkid) | **Get** /network/{networkId} | Get TV network details
*TmdbAPI* | [**GetStudioByStudioId**](overseerr/docs/TmdbAPI.md#getstudiobystudioid) | **Get** /studio/{studioId} | Get movie studio details
*TmdbAPI* | [**ListBackdrops**](overseerr/docs/TmdbAPI.md#listbackdrops) | **Get** /backdrops | Get backdrops of trending items
*TmdbAPI* | [**ListGenresMovie**](overseerr/docs/TmdbAPI.md#listgenresmovie) | **Get** /genres/movie | Get list of official TMDB movie genres
*TmdbAPI* | [**ListGenresTv**](overseerr/docs/TmdbAPI.md#listgenrestv) | **Get** /genres/tv | Get list of official TMDB movie genres
*TmdbAPI* | [**ListLanguages**](overseerr/docs/TmdbAPI.md#listlanguages) | **Get** /languages | Languages supported by TMDB
*TmdbAPI* | [**ListRegions**](overseerr/docs/TmdbAPI.md#listregions) | **Get** /regions | Regions supported by TMDB
*TvAPI* | [**GetTvByTvId**](overseerr/docs/TvAPI.md#gettvbytvid) | **Get** /tv/{tvId} | Get TV details
*TvAPI* | [**GetTvRatings**](overseerr/docs/TvAPI.md#gettvratings) | **Get** /tv/{tvId}/ratings | Get TV ratings
*TvAPI* | [**GetTvRecommendations**](overseerr/docs/TvAPI.md#gettvrecommendations) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
*TvAPI* | [**GetTvSeasonBySeasonId**](overseerr/docs/TvAPI.md#gettvseasonbyseasonid) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
*TvAPI* | [**GetTvSimilar**](overseerr/docs/TvAPI.md#gettvsimilar) | **Get** /tv/{tvId}/similar | Get similar TV series
*UsersAPI* | [**CreateAuthResetPassword**](overseerr/docs/UsersAPI.md#createauthresetpassword) | **Post** /auth/reset-password | Send a reset password email
*UsersAPI* | [**CreateAuthResetPasswordByGuid**](overseerr/docs/UsersAPI.md#createauthresetpasswordbyguid) | **Post** /auth/reset-password/{guid} | Reset the password for a user
*UsersAPI* | [**CreateUser**](overseerr/docs/UsersAPI.md#createuser) | **Post** /user | Create new user
*UsersAPI* | [**CreateUserImportFromPlex**](overseerr/docs/UsersAPI.md#createuserimportfromplex) | **Post** /user/import-from-plex | Import all users from Plex
*UsersAPI* | [**CreateUserRegisterPushSubscription**](overseerr/docs/UsersAPI.md#createuserregisterpushsubscription) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
*UsersAPI* | [**CreateUserSettingsMain**](overseerr/docs/UsersAPI.md#createusersettingsmain) | **Post** /user/{userId}/settings/main | Update general settings for a user
*UsersAPI* | [**CreateUserSettingsNotifications**](overseerr/docs/UsersAPI.md#createusersettingsnotifications) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
*UsersAPI* | [**CreateUserSettingsPassword**](overseerr/docs/UsersAPI.md#createusersettingspassword) | **Post** /user/{userId}/settings/password | Update password for a user
*UsersAPI* | [**CreateUserSettingsPermissions**](overseerr/docs/UsersAPI.md#createusersettingspermissions) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
*UsersAPI* | [**DeleteUser**](overseerr/docs/UsersAPI.md#deleteuser) | **Delete** /user/{userId} | Delete user by ID
*UsersAPI* | [**GetAuthMe**](overseerr/docs/UsersAPI.md#getauthme) | **Get** /auth/me | Get logged-in user
*UsersAPI* | [**GetUser**](overseerr/docs/UsersAPI.md#getuser) | **Get** /user | Get all users
*UsersAPI* | [**GetUserByUserId**](overseerr/docs/UsersAPI.md#getuserbyuserid) | **Get** /user/{userId} | Get user by ID
*UsersAPI* | [**GetUserQuota**](overseerr/docs/UsersAPI.md#getuserquota) | **Get** /user/{userId}/quota | Get quotas for a specific user
*UsersAPI* | [**GetUserRequests**](overseerr/docs/UsersAPI.md#getuserrequests) | **Get** /user/{userId}/requests | Get requests for a specific user
*UsersAPI* | [**GetUserSettingsMain**](overseerr/docs/UsersAPI.md#getusersettingsmain) | **Get** /user/{userId}/settings/main | Get general settings for a user
*UsersAPI* | [**GetUserSettingsNotifications**](overseerr/docs/UsersAPI.md#getusersettingsnotifications) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
*UsersAPI* | [**GetUserSettingsPassword**](overseerr/docs/UsersAPI.md#getusersettingspassword) | **Get** /user/{userId}/settings/password | Get password page informatiom
*UsersAPI* | [**GetUserSettingsPermissions**](overseerr/docs/UsersAPI.md#getusersettingspermissions) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
*UsersAPI* | [**GetUserWatchData**](overseerr/docs/UsersAPI.md#getuserwatchdata) | **Get** /user/{userId}/watch_data | Get watch data
*UsersAPI* | [**GetUserWatchlist**](overseerr/docs/UsersAPI.md#getuserwatchlist) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user
*UsersAPI* | [**ListPlexUsers**](overseerr/docs/UsersAPI.md#listplexusers) | **Get** /settings/plex/users | Get Plex users
*UsersAPI* | [**PutUser**](overseerr/docs/UsersAPI.md#putuser) | **Put** /user | Update batch of users
*UsersAPI* | [**UpdateUser**](overseerr/docs/UsersAPI.md#updateuser) | **Put** /user/{userId} | Update a user by user ID


## Documentation For Models

 - [Cast](docs/Cast.md)
 - [Collection](docs/Collection.md)
 - [Company](docs/Company.md)
 - [CreateAuthLocalRequest](docs/CreateAuthLocalRequest.md)
 - [CreateAuthLogout200Response](docs/CreateAuthLogout200Response.md)
 - [CreateAuthPlexRequest](docs/CreateAuthPlexRequest.md)
 - [CreateAuthResetPasswordByGuidRequest](docs/CreateAuthResetPasswordByGuidRequest.md)
 - [CreateAuthResetPasswordRequest](docs/CreateAuthResetPasswordRequest.md)
 - [CreateDiscoverAddRequest](docs/CreateDiscoverAddRequest.md)
 - [CreateIssueCommentRequest](docs/CreateIssueCommentRequest.md)
 - [CreateIssueRequest](docs/CreateIssueRequest.md)
 - [CreateJobsScheduleRequest](docs/CreateJobsScheduleRequest.md)
 - [CreateMediaByStatusRequest](docs/CreateMediaByStatusRequest.md)
 - [CreatePlexSyncRequest](docs/CreatePlexSyncRequest.md)
 - [CreateRequestRequest](docs/CreateRequestRequest.md)
 - [CreateRequestRequestSeasons](docs/CreateRequestRequestSeasons.md)
 - [CreateUserImportFromPlexRequest](docs/CreateUserImportFromPlexRequest.md)
 - [CreateUserRegisterPushSubscriptionRequest](docs/CreateUserRegisterPushSubscriptionRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserSettingsMainRequest](docs/CreateUserSettingsMainRequest.md)
 - [CreateUserSettingsPasswordRequest](docs/CreateUserSettingsPasswordRequest.md)
 - [CreateUserSettingsPermissionsRequest](docs/CreateUserSettingsPermissionsRequest.md)
 - [CreditCast](docs/CreditCast.md)
 - [CreditCrew](docs/CreditCrew.md)
 - [Crew](docs/Crew.md)
 - [DiscordSettings](docs/DiscordSettings.md)
 - [DiscordSettingsOptions](docs/DiscordSettingsOptions.md)
 - [DiscoverSlider](docs/DiscoverSlider.md)
 - [Episode](docs/Episode.md)
 - [ExternalIds](docs/ExternalIds.md)
 - [Genre](docs/Genre.md)
 - [GetAbout200Response](docs/GetAbout200Response.md)
 - [GetCache200Response](docs/GetCache200Response.md)
 - [GetCache200ResponseApiCachesInner](docs/GetCache200ResponseApiCachesInner.md)
 - [GetCache200ResponseApiCachesInnerStats](docs/GetCache200ResponseApiCachesInnerStats.md)
 - [GetCache200ResponseImageCache](docs/GetCache200ResponseImageCache.md)
 - [GetCache200ResponseImageCacheTmdb](docs/GetCache200ResponseImageCacheTmdb.md)
 - [GetDiscoverMovies200Response](docs/GetDiscoverMovies200Response.md)
 - [GetDiscoverMoviesGenreByGenreId200Response](docs/GetDiscoverMoviesGenreByGenreId200Response.md)
 - [GetDiscoverMoviesLanguageByLanguage200Response](docs/GetDiscoverMoviesLanguageByLanguage200Response.md)
 - [GetDiscoverMoviesStudioByStudioId200Response](docs/GetDiscoverMoviesStudioByStudioId200Response.md)
 - [GetDiscoverTv200Response](docs/GetDiscoverTv200Response.md)
 - [GetDiscoverTvGenreByGenreId200Response](docs/GetDiscoverTvGenreByGenreId200Response.md)
 - [GetDiscoverTvLanguageByLanguage200Response](docs/GetDiscoverTvLanguageByLanguage200Response.md)
 - [GetDiscoverTvNetworkByNetworkId200Response](docs/GetDiscoverTvNetworkByNetworkId200Response.md)
 - [GetIssue200Response](docs/GetIssue200Response.md)
 - [GetIssueCount200Response](docs/GetIssueCount200Response.md)
 - [GetMedia200Response](docs/GetMedia200Response.md)
 - [GetMediaWatchData200Response](docs/GetMediaWatchData200Response.md)
 - [GetMediaWatchData200ResponseData](docs/GetMediaWatchData200ResponseData.md)
 - [GetMovieRatings200Response](docs/GetMovieRatings200Response.md)
 - [GetMovieRatingscombined200Response](docs/GetMovieRatingscombined200Response.md)
 - [GetMovieRatingscombined200ResponseImdb](docs/GetMovieRatingscombined200ResponseImdb.md)
 - [GetPersonCombinedCredits200Response](docs/GetPersonCombinedCredits200Response.md)
 - [GetPlexSync200Response](docs/GetPlexSync200Response.md)
 - [GetRequestCount200Response](docs/GetRequestCount200Response.md)
 - [GetSearch200Response](docs/GetSearch200Response.md)
 - [GetSearch200ResponseResultsInner](docs/GetSearch200ResponseResultsInner.md)
 - [GetSearchCompany200Response](docs/GetSearchCompany200Response.md)
 - [GetSearchKeyword200Response](docs/GetSearchKeyword200Response.md)
 - [GetServiceRadarrByRadarrId200Response](docs/GetServiceRadarrByRadarrId200Response.md)
 - [GetServiceSonarrBySonarrId200Response](docs/GetServiceSonarrBySonarrId200Response.md)
 - [GetStatus200Response](docs/GetStatus200Response.md)
 - [GetStatusAppdata200Response](docs/GetStatusAppdata200Response.md)
 - [GetTvRatings200Response](docs/GetTvRatings200Response.md)
 - [GetUser200Response](docs/GetUser200Response.md)
 - [GetUserQuota200Response](docs/GetUserQuota200Response.md)
 - [GetUserQuota200ResponseMovie](docs/GetUserQuota200ResponseMovie.md)
 - [GetUserRequests200Response](docs/GetUserRequests200Response.md)
 - [GetUserSettingsMain200Response](docs/GetUserSettingsMain200Response.md)
 - [GetUserSettingsPassword200Response](docs/GetUserSettingsPassword200Response.md)
 - [GetUserSettingsPermissions200Response](docs/GetUserSettingsPermissions200Response.md)
 - [GetUserWatchData200Response](docs/GetUserWatchData200Response.md)
 - [GetUserWatchlist200Response](docs/GetUserWatchlist200Response.md)
 - [GetUserWatchlist200ResponseResultsInner](docs/GetUserWatchlist200ResponseResultsInner.md)
 - [GotifySettings](docs/GotifySettings.md)
 - [GotifySettingsOptions](docs/GotifySettingsOptions.md)
 - [Issue](docs/Issue.md)
 - [IssueComment](docs/IssueComment.md)
 - [Job](docs/Job.md)
 - [Keyword](docs/Keyword.md)
 - [ListDiscoverGenresliderMovie200ResponseInner](docs/ListDiscoverGenresliderMovie200ResponseInner.md)
 - [ListGenresMovie200ResponseInner](docs/ListGenresMovie200ResponseInner.md)
 - [ListGenresTv200ResponseInner](docs/ListGenresTv200ResponseInner.md)
 - [ListLanguages200ResponseInner](docs/ListLanguages200ResponseInner.md)
 - [ListLogs200ResponseInner](docs/ListLogs200ResponseInner.md)
 - [ListPlexUsers200ResponseInner](docs/ListPlexUsers200ResponseInner.md)
 - [ListRegions200ResponseInner](docs/ListRegions200ResponseInner.md)
 - [LunaSeaSettings](docs/LunaSeaSettings.md)
 - [LunaSeaSettingsOptions](docs/LunaSeaSettingsOptions.md)
 - [MainSettings](docs/MainSettings.md)
 - [MediaInfo](docs/MediaInfo.md)
 - [MediaRequest](docs/MediaRequest.md)
 - [MediaRequestModifiedBy](docs/MediaRequestModifiedBy.md)
 - [MovieDetails](docs/MovieDetails.md)
 - [MovieDetailsCollection](docs/MovieDetailsCollection.md)
 - [MovieDetailsCredits](docs/MovieDetailsCredits.md)
 - [MovieDetailsProductionCountriesInner](docs/MovieDetailsProductionCountriesInner.md)
 - [MovieDetailsReleases](docs/MovieDetailsReleases.md)
 - [MovieDetailsReleasesResultsInner](docs/MovieDetailsReleasesResultsInner.md)
 - [MovieDetailsReleasesResultsInnerReleaseDatesInner](docs/MovieDetailsReleasesResultsInnerReleaseDatesInner.md)
 - [MovieResult](docs/MovieResult.md)
 - [Network](docs/Network.md)
 - [NotificationAgentTypes](docs/NotificationAgentTypes.md)
 - [NotificationEmailSettings](docs/NotificationEmailSettings.md)
 - [NotificationEmailSettingsOptions](docs/NotificationEmailSettingsOptions.md)
 - [PageInfo](docs/PageInfo.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonResult](docs/PersonResult.md)
 - [PersonResultKnownForInner](docs/PersonResultKnownForInner.md)
 - [PlexConnection](docs/PlexConnection.md)
 - [PlexDevice](docs/PlexDevice.md)
 - [PlexLibrary](docs/PlexLibrary.md)
 - [PlexSettings](docs/PlexSettings.md)
 - [ProductionCompany](docs/ProductionCompany.md)
 - [PublicSettings](docs/PublicSettings.md)
 - [PushbulletSettings](docs/PushbulletSettings.md)
 - [PushbulletSettingsOptions](docs/PushbulletSettingsOptions.md)
 - [PushoverSettings](docs/PushoverSettings.md)
 - [PushoverSettingsOptions](docs/PushoverSettingsOptions.md)
 - [PutUserRequest](docs/PutUserRequest.md)
 - [RadarrSettings](docs/RadarrSettings.md)
 - [RelatedVideo](docs/RelatedVideo.md)
 - [Season](docs/Season.md)
 - [ServarrTag](docs/ServarrTag.md)
 - [ServiceProfile](docs/ServiceProfile.md)
 - [SlackSettings](docs/SlackSettings.md)
 - [SlackSettingsOptions](docs/SlackSettingsOptions.md)
 - [SonarrSeries](docs/SonarrSeries.md)
 - [SonarrSeriesAddOptionsInner](docs/SonarrSeriesAddOptionsInner.md)
 - [SonarrSeriesImagesInner](docs/SonarrSeriesImagesInner.md)
 - [SonarrSeriesRatingsInner](docs/SonarrSeriesRatingsInner.md)
 - [SonarrSeriesSeasonsInner](docs/SonarrSeriesSeasonsInner.md)
 - [SonarrSettings](docs/SonarrSettings.md)
 - [SpokenLanguage](docs/SpokenLanguage.md)
 - [TautulliSettings](docs/TautulliSettings.md)
 - [TelegramSettings](docs/TelegramSettings.md)
 - [TelegramSettingsOptions](docs/TelegramSettingsOptions.md)
 - [TestRadarr200Response](docs/TestRadarr200Response.md)
 - [TestRadarrRequest](docs/TestRadarrRequest.md)
 - [TestSonarrRequest](docs/TestSonarrRequest.md)
 - [TvDetails](docs/TvDetails.md)
 - [TvDetailsContentRatings](docs/TvDetailsContentRatings.md)
 - [TvDetailsContentRatingsResultsInner](docs/TvDetailsContentRatingsResultsInner.md)
 - [TvDetailsCreatedByInner](docs/TvDetailsCreatedByInner.md)
 - [TvResult](docs/TvResult.md)
 - [UpdateDiscoverRequest](docs/UpdateDiscoverRequest.md)
 - [UpdateIssueCommentRequest](docs/UpdateIssueCommentRequest.md)
 - [UpdateRequestRequest](docs/UpdateRequestRequest.md)
 - [User](docs/User.md)
 - [UserSettings](docs/UserSettings.md)
 - [UserSettingsNotifications](docs/UserSettingsNotifications.md)
 - [WatchProviderDetails](docs/WatchProviderDetails.md)
 - [WatchProviderRegion](docs/WatchProviderRegion.md)
 - [WatchProvidersInner](docs/WatchProvidersInner.md)
 - [WebPushSettings](docs/WebPushSettings.md)
 - [WebhookSettings](docs/WebhookSettings.md)
 - [WebhookSettingsOptions](docs/WebhookSettingsOptions.md)


## Documentation For Authorization



### cookieAuth

- **Type**: API key
- **API key parameter name**: connect.sid
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: connect.sid and passed in as the auth context for each request.


### apiKey

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



