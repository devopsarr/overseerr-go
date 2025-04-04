# Go API client for overseerr

This is the documentation for the Overseerr API backend.

Two primary authentication methods are supported:

- **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie.
- **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.34.0
- Package version: 1.0.1 <!--- x-release-please-version -->
- Generator version: 7.12.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import overseerr "github.com/devopsarr/overseerr-go/overseerr"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `overseerr.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), overseerr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `overseerr.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), overseerr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `overseerr.ContextOperationServerIndices` and `overseerr.ContextOperationServerVariables` context maps.

```go
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
*AuthAPI* | [**CreateAuthLocal**](docs/AuthAPI.md#createauthlocal) | **Post** /auth/local | Sign in using a local account
*AuthAPI* | [**CreateAuthLogout**](docs/AuthAPI.md#createauthlogout) | **Post** /auth/logout | Sign out and clear session cookie
*AuthAPI* | [**CreateAuthPlex**](docs/AuthAPI.md#createauthplex) | **Post** /auth/plex | Sign in using a Plex token
*AuthAPI* | [**GetAuthMe**](docs/AuthAPI.md#getauthme) | **Get** /auth/me | Get logged-in user
*CollectionAPI* | [**GetCollectionByCollectionId**](docs/CollectionAPI.md#getcollectionbycollectionid) | **Get** /collection/{collectionId} | Get collection details
*IssueAPI* | [**CreateIssue**](docs/IssueAPI.md#createissue) | **Post** /issue | Create new issue
*IssueAPI* | [**CreateIssueByStatus**](docs/IssueAPI.md#createissuebystatus) | **Post** /issue/{issueId}/{status} | Update an issue&#39;s status
*IssueAPI* | [**CreateIssueComment**](docs/IssueAPI.md#createissuecomment) | **Post** /issue/{issueId}/comment | Create a comment
*IssueAPI* | [**DeleteIssue**](docs/IssueAPI.md#deleteissue) | **Delete** /issue/{issueId} | Delete issue
*IssueAPI* | [**DeleteIssueComment**](docs/IssueAPI.md#deleteissuecomment) | **Delete** /issueComment/{commentId} | Delete issue comment
*IssueAPI* | [**GetIssue**](docs/IssueAPI.md#getissue) | **Get** /issue | Get all issues
*IssueAPI* | [**GetIssueByIssueId**](docs/IssueAPI.md#getissuebyissueid) | **Get** /issue/{issueId} | Get issue
*IssueAPI* | [**GetIssueCommentByCommentId**](docs/IssueAPI.md#getissuecommentbycommentid) | **Get** /issueComment/{commentId} | Get issue comment
*IssueAPI* | [**GetIssueCount**](docs/IssueAPI.md#getissuecount) | **Get** /issue/count | Gets issue counts
*IssueAPI* | [**UpdateIssueComment**](docs/IssueAPI.md#updateissuecomment) | **Put** /issueComment/{commentId} | Update issue comment
*MediaAPI* | [**CreateMediaByStatus**](docs/MediaAPI.md#createmediabystatus) | **Post** /media/{mediaId}/{status} | Update media status
*MediaAPI* | [**DeleteMedia**](docs/MediaAPI.md#deletemedia) | **Delete** /media/{mediaId} | Delete media item
*MediaAPI* | [**GetMedia**](docs/MediaAPI.md#getmedia) | **Get** /media | Get media
*MediaAPI* | [**GetMediaWatchData**](docs/MediaAPI.md#getmediawatchdata) | **Get** /media/{mediaId}/watch_data | Get watch data
*MoviesAPI* | [**GetMovieByMovieId**](docs/MoviesAPI.md#getmoviebymovieid) | **Get** /movie/{movieId} | Get movie details
*MoviesAPI* | [**GetMovieRatings**](docs/MoviesAPI.md#getmovieratings) | **Get** /movie/{movieId}/ratings | Get movie ratings
*MoviesAPI* | [**GetMovieRatingscombined**](docs/MoviesAPI.md#getmovieratingscombined) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
*MoviesAPI* | [**GetMovieRecommendations**](docs/MoviesAPI.md#getmovierecommendations) | **Get** /movie/{movieId}/recommendations | Get recommended movies
*MoviesAPI* | [**GetMovieSimilar**](docs/MoviesAPI.md#getmoviesimilar) | **Get** /movie/{movieId}/similar | Get similar movies
*OtherAPI* | [**GetKeywordByKeywordId**](docs/OtherAPI.md#getkeywordbykeywordid) | **Get** /keyword/{keywordId} | Get keyword
*OtherAPI* | [**ListWatchprovidersMovies**](docs/OtherAPI.md#listwatchprovidersmovies) | **Get** /watchproviders/movies | Get watch provider movies
*OtherAPI* | [**ListWatchprovidersRegions**](docs/OtherAPI.md#listwatchprovidersregions) | **Get** /watchproviders/regions | Get watch provider regions
*OtherAPI* | [**ListWatchprovidersTv**](docs/OtherAPI.md#listwatchproviderstv) | **Get** /watchproviders/tv | Get watch provider series
*PersonAPI* | [**GetPersonByPersonId**](docs/PersonAPI.md#getpersonbypersonid) | **Get** /person/{personId} | Get person details
*PersonAPI* | [**GetPersonCombinedCredits**](docs/PersonAPI.md#getpersoncombinedcredits) | **Get** /person/{personId}/combined_credits | Get combined credits
*PublicAPI* | [**GetStatus**](docs/PublicAPI.md#getstatus) | **Get** /status | Get Overseerr status
*PublicAPI* | [**GetStatusAppdata**](docs/PublicAPI.md#getstatusappdata) | **Get** /status/appdata | Get application data volume status
*RequestAPI* | [**CreateRequest**](docs/RequestAPI.md#createrequest) | **Post** /request | Create new request
*RequestAPI* | [**CreateRequestByStatus**](docs/RequestAPI.md#createrequestbystatus) | **Post** /request/{requestId}/{status} | Update a request&#39;s status
*RequestAPI* | [**CreateRequestRetry**](docs/RequestAPI.md#createrequestretry) | **Post** /request/{requestId}/retry | Retry failed request
*RequestAPI* | [**DeleteRequest**](docs/RequestAPI.md#deleterequest) | **Delete** /request/{requestId} | Delete request
*RequestAPI* | [**GetRequest**](docs/RequestAPI.md#getrequest) | **Get** /request | Get all requests
*RequestAPI* | [**GetRequestByRequestId**](docs/RequestAPI.md#getrequestbyrequestid) | **Get** /request/{requestId} | Get MediaRequest
*RequestAPI* | [**GetRequestCount**](docs/RequestAPI.md#getrequestcount) | **Get** /request/count | Gets request counts
*RequestAPI* | [**UpdateRequest**](docs/RequestAPI.md#updaterequest) | **Put** /request/{requestId} | Update MediaRequest
*SearchAPI* | [**GetDiscoverKeywordMovies**](docs/SearchAPI.md#getdiscoverkeywordmovies) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
*SearchAPI* | [**GetDiscoverMovies**](docs/SearchAPI.md#getdiscovermovies) | **Get** /discover/movies | Discover movies
*SearchAPI* | [**GetDiscoverMoviesGenreByGenreId**](docs/SearchAPI.md#getdiscovermoviesgenrebygenreid) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
*SearchAPI* | [**GetDiscoverMoviesLanguageByLanguage**](docs/SearchAPI.md#getdiscovermovieslanguagebylanguage) | **Get** /discover/movies/language/{language} | Discover movies by original language
*SearchAPI* | [**GetDiscoverMoviesStudioByStudioId**](docs/SearchAPI.md#getdiscovermoviesstudiobystudioid) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
*SearchAPI* | [**GetDiscoverMoviesUpcoming**](docs/SearchAPI.md#getdiscovermoviesupcoming) | **Get** /discover/movies/upcoming | Upcoming movies
*SearchAPI* | [**GetDiscoverTrending**](docs/SearchAPI.md#getdiscovertrending) | **Get** /discover/trending | Trending movies and TV
*SearchAPI* | [**GetDiscoverTv**](docs/SearchAPI.md#getdiscovertv) | **Get** /discover/tv | Discover TV shows
*SearchAPI* | [**GetDiscoverTvGenreByGenreId**](docs/SearchAPI.md#getdiscovertvgenrebygenreid) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
*SearchAPI* | [**GetDiscoverTvLanguageByLanguage**](docs/SearchAPI.md#getdiscovertvlanguagebylanguage) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
*SearchAPI* | [**GetDiscoverTvNetworkByNetworkId**](docs/SearchAPI.md#getdiscovertvnetworkbynetworkid) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
*SearchAPI* | [**GetDiscoverTvUpcoming**](docs/SearchAPI.md#getdiscovertvupcoming) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
*SearchAPI* | [**GetDiscoverWatchlist**](docs/SearchAPI.md#getdiscoverwatchlist) | **Get** /discover/watchlist | Get the Plex watchlist.
*SearchAPI* | [**GetSearch**](docs/SearchAPI.md#getsearch) | **Get** /search | Search for movies, TV shows, or people
*SearchAPI* | [**GetSearchCompany**](docs/SearchAPI.md#getsearchcompany) | **Get** /search/company | Search for companies
*SearchAPI* | [**GetSearchKeyword**](docs/SearchAPI.md#getsearchkeyword) | **Get** /search/keyword | Search for keywords
*SearchAPI* | [**ListDiscoverGenresliderMovie**](docs/SearchAPI.md#listdiscovergenreslidermovie) | **Get** /discover/genreslider/movie | Get genre slider data for movies
*SearchAPI* | [**ListDiscoverGenresliderTv**](docs/SearchAPI.md#listdiscovergenreslidertv) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
*ServiceAPI* | [**GetServiceRadarrByRadarrId**](docs/ServiceAPI.md#getserviceradarrbyradarrid) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
*ServiceAPI* | [**GetServiceSonarrBySonarrId**](docs/ServiceAPI.md#getservicesonarrbysonarrid) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
*ServiceAPI* | [**ListServiceRadarr**](docs/ServiceAPI.md#listserviceradarr) | **Get** /service/radarr | Get non-sensitive Radarr server list
*ServiceAPI* | [**ListServiceSonarr**](docs/ServiceAPI.md#listservicesonarr) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
*ServiceAPI* | [**ListServiceSonarrLookupByTmdbId**](docs/ServiceAPI.md#listservicesonarrlookupbytmdbid) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr
*SettingsAPI* | [**CreateCacheFlush**](docs/SettingsAPI.md#createcacheflush) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
*SettingsAPI* | [**CreateDiscover**](docs/SettingsAPI.md#creatediscover) | **Post** /settings/discover | Batch update all sliders.
*SettingsAPI* | [**CreateDiscoverAdd**](docs/SettingsAPI.md#creatediscoveradd) | **Post** /settings/discover/add | Add a new slider
*SettingsAPI* | [**CreateInitialize**](docs/SettingsAPI.md#createinitialize) | **Post** /settings/initialize | Initialize application
*SettingsAPI* | [**CreateJobsCancel**](docs/SettingsAPI.md#createjobscancel) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
*SettingsAPI* | [**CreateJobsRun**](docs/SettingsAPI.md#createjobsrun) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
*SettingsAPI* | [**CreateJobsSchedule**](docs/SettingsAPI.md#createjobsschedule) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
*SettingsAPI* | [**CreateMain**](docs/SettingsAPI.md#createmain) | **Post** /settings/main | Update main settings
*SettingsAPI* | [**CreateMainRegenerate**](docs/SettingsAPI.md#createmainregenerate) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
*SettingsAPI* | [**CreateNotificationsDiscord**](docs/SettingsAPI.md#createnotificationsdiscord) | **Post** /settings/notifications/discord | Update Discord notification settings
*SettingsAPI* | [**CreateNotificationsEmail**](docs/SettingsAPI.md#createnotificationsemail) | **Post** /settings/notifications/email | Update email notification settings
*SettingsAPI* | [**CreateNotificationsGotify**](docs/SettingsAPI.md#createnotificationsgotify) | **Post** /settings/notifications/gotify | Update Gotify notification settings
*SettingsAPI* | [**CreateNotificationsLunasea**](docs/SettingsAPI.md#createnotificationslunasea) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
*SettingsAPI* | [**CreateNotificationsPushbullet**](docs/SettingsAPI.md#createnotificationspushbullet) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
*SettingsAPI* | [**CreateNotificationsPushover**](docs/SettingsAPI.md#createnotificationspushover) | **Post** /settings/notifications/pushover | Update Pushover notification settings
*SettingsAPI* | [**CreateNotificationsSlack**](docs/SettingsAPI.md#createnotificationsslack) | **Post** /settings/notifications/slack | Update Slack notification settings
*SettingsAPI* | [**CreateNotificationsTelegram**](docs/SettingsAPI.md#createnotificationstelegram) | **Post** /settings/notifications/telegram | Update Telegram notification settings
*SettingsAPI* | [**CreateNotificationsWebhook**](docs/SettingsAPI.md#createnotificationswebhook) | **Post** /settings/notifications/webhook | Update webhook notification settings
*SettingsAPI* | [**CreateNotificationsWebpush**](docs/SettingsAPI.md#createnotificationswebpush) | **Post** /settings/notifications/webpush | Update Web Push notification settings
*SettingsAPI* | [**CreatePlex**](docs/SettingsAPI.md#createplex) | **Post** /settings/plex | Update Plex settings
*SettingsAPI* | [**CreatePlexSync**](docs/SettingsAPI.md#createplexsync) | **Post** /settings/plex/sync | Start full Plex library scan
*SettingsAPI* | [**CreateRadarr**](docs/SettingsAPI.md#createradarr) | **Post** /settings/radarr | Create Radarr instance
*SettingsAPI* | [**CreateSonarr**](docs/SettingsAPI.md#createsonarr) | **Post** /settings/sonarr | Create Sonarr instance
*SettingsAPI* | [**CreateTautulli**](docs/SettingsAPI.md#createtautulli) | **Post** /settings/tautulli | Update Tautulli settings
*SettingsAPI* | [**DeleteDiscover**](docs/SettingsAPI.md#deletediscover) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
*SettingsAPI* | [**DeleteRadarr**](docs/SettingsAPI.md#deleteradarr) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
*SettingsAPI* | [**DeleteSonarr**](docs/SettingsAPI.md#deletesonarr) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
*SettingsAPI* | [**GetAbout**](docs/SettingsAPI.md#getabout) | **Get** /settings/about | Get server stats
*SettingsAPI* | [**GetCache**](docs/SettingsAPI.md#getcache) | **Get** /settings/cache | Get a list of active caches
*SettingsAPI* | [**GetDiscoverReset**](docs/SettingsAPI.md#getdiscoverreset) | **Get** /settings/discover/reset | Reset all discover sliders
*SettingsAPI* | [**GetMain**](docs/SettingsAPI.md#getmain) | **Get** /settings/main | Get main settings
*SettingsAPI* | [**GetNotificationsDiscord**](docs/SettingsAPI.md#getnotificationsdiscord) | **Get** /settings/notifications/discord | Get Discord notification settings
*SettingsAPI* | [**GetNotificationsEmail**](docs/SettingsAPI.md#getnotificationsemail) | **Get** /settings/notifications/email | Get email notification settings
*SettingsAPI* | [**GetNotificationsGotify**](docs/SettingsAPI.md#getnotificationsgotify) | **Get** /settings/notifications/gotify | Get Gotify notification settings
*SettingsAPI* | [**GetNotificationsLunasea**](docs/SettingsAPI.md#getnotificationslunasea) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
*SettingsAPI* | [**GetNotificationsPushbullet**](docs/SettingsAPI.md#getnotificationspushbullet) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
*SettingsAPI* | [**GetNotificationsPushover**](docs/SettingsAPI.md#getnotificationspushover) | **Get** /settings/notifications/pushover | Get Pushover notification settings
*SettingsAPI* | [**GetNotificationsSlack**](docs/SettingsAPI.md#getnotificationsslack) | **Get** /settings/notifications/slack | Get Slack notification settings
*SettingsAPI* | [**GetNotificationsTelegram**](docs/SettingsAPI.md#getnotificationstelegram) | **Get** /settings/notifications/telegram | Get Telegram notification settings
*SettingsAPI* | [**GetNotificationsWebhook**](docs/SettingsAPI.md#getnotificationswebhook) | **Get** /settings/notifications/webhook | Get webhook notification settings
*SettingsAPI* | [**GetNotificationsWebpush**](docs/SettingsAPI.md#getnotificationswebpush) | **Get** /settings/notifications/webpush | Get Web Push notification settings
*SettingsAPI* | [**GetPlex**](docs/SettingsAPI.md#getplex) | **Get** /settings/plex | Get Plex settings
*SettingsAPI* | [**GetPlexSync**](docs/SettingsAPI.md#getplexsync) | **Get** /settings/plex/sync | Get status of full Plex library scan
*SettingsAPI* | [**GetPublic**](docs/SettingsAPI.md#getpublic) | **Get** /settings/public | Get public settings
*SettingsAPI* | [**GetTautulli**](docs/SettingsAPI.md#gettautulli) | **Get** /settings/tautulli | Get Tautulli settings
*SettingsAPI* | [**ListDiscover**](docs/SettingsAPI.md#listdiscover) | **Get** /settings/discover | Get all discover sliders
*SettingsAPI* | [**ListJobs**](docs/SettingsAPI.md#listjobs) | **Get** /settings/jobs | Get scheduled jobs
*SettingsAPI* | [**ListLogs**](docs/SettingsAPI.md#listlogs) | **Get** /settings/logs | Returns logs
*SettingsAPI* | [**ListNotificationsPushoverSounds**](docs/SettingsAPI.md#listnotificationspushoversounds) | **Get** /settings/notifications/pushover/sounds | Get Pushover sounds
*SettingsAPI* | [**ListPlexDevicesServers**](docs/SettingsAPI.md#listplexdevicesservers) | **Get** /settings/plex/devices/servers | Gets the user&#39;s available Plex servers
*SettingsAPI* | [**ListPlexLibrary**](docs/SettingsAPI.md#listplexlibrary) | **Get** /settings/plex/library | Get Plex libraries
*SettingsAPI* | [**ListPlexUsers**](docs/SettingsAPI.md#listplexusers) | **Get** /settings/plex/users | Get Plex users
*SettingsAPI* | [**ListRadarr**](docs/SettingsAPI.md#listradarr) | **Get** /settings/radarr | Get Radarr settings
*SettingsAPI* | [**ListRadarrProfiles**](docs/SettingsAPI.md#listradarrprofiles) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
*SettingsAPI* | [**ListSonarr**](docs/SettingsAPI.md#listsonarr) | **Get** /settings/sonarr | Get Sonarr settings
*SettingsAPI* | [**TestNotificationsDiscord**](docs/SettingsAPI.md#testnotificationsdiscord) | **Post** /settings/notifications/discord/test | Test Discord settings
*SettingsAPI* | [**TestNotificationsEmail**](docs/SettingsAPI.md#testnotificationsemail) | **Post** /settings/notifications/email/test | Test email settings
*SettingsAPI* | [**TestNotificationsGotify**](docs/SettingsAPI.md#testnotificationsgotify) | **Post** /settings/notifications/gotify/test | Test Gotify settings
*SettingsAPI* | [**TestNotificationsLunasea**](docs/SettingsAPI.md#testnotificationslunasea) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
*SettingsAPI* | [**TestNotificationsPushbullet**](docs/SettingsAPI.md#testnotificationspushbullet) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
*SettingsAPI* | [**TestNotificationsPushover**](docs/SettingsAPI.md#testnotificationspushover) | **Post** /settings/notifications/pushover/test | Test Pushover settings
*SettingsAPI* | [**TestNotificationsSlack**](docs/SettingsAPI.md#testnotificationsslack) | **Post** /settings/notifications/slack/test | Test Slack settings
*SettingsAPI* | [**TestNotificationsTelegram**](docs/SettingsAPI.md#testnotificationstelegram) | **Post** /settings/notifications/telegram/test | Test Telegram settings
*SettingsAPI* | [**TestNotificationsWebhook**](docs/SettingsAPI.md#testnotificationswebhook) | **Post** /settings/notifications/webhook/test | Test webhook settings
*SettingsAPI* | [**TestNotificationsWebpush**](docs/SettingsAPI.md#testnotificationswebpush) | **Post** /settings/notifications/webpush/test | Test Web Push settings
*SettingsAPI* | [**TestRadarr**](docs/SettingsAPI.md#testradarr) | **Post** /settings/radarr/test | Test Radarr configuration
*SettingsAPI* | [**TestSonarr**](docs/SettingsAPI.md#testsonarr) | **Post** /settings/sonarr/test | Test Sonarr configuration
*SettingsAPI* | [**UpdateDiscover**](docs/SettingsAPI.md#updatediscover) | **Put** /settings/discover/{sliderId} | Update a single slider
*SettingsAPI* | [**UpdateRadarr**](docs/SettingsAPI.md#updateradarr) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
*SettingsAPI* | [**UpdateSonarr**](docs/SettingsAPI.md#updatesonarr) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance
*TmdbAPI* | [**GetNetworkByNetworkId**](docs/TmdbAPI.md#getnetworkbynetworkid) | **Get** /network/{networkId} | Get TV network details
*TmdbAPI* | [**GetStudioByStudioId**](docs/TmdbAPI.md#getstudiobystudioid) | **Get** /studio/{studioId} | Get movie studio details
*TmdbAPI* | [**ListBackdrops**](docs/TmdbAPI.md#listbackdrops) | **Get** /backdrops | Get backdrops of trending items
*TmdbAPI* | [**ListGenresMovie**](docs/TmdbAPI.md#listgenresmovie) | **Get** /genres/movie | Get list of official TMDB movie genres
*TmdbAPI* | [**ListGenresTv**](docs/TmdbAPI.md#listgenrestv) | **Get** /genres/tv | Get list of official TMDB movie genres
*TmdbAPI* | [**ListLanguages**](docs/TmdbAPI.md#listlanguages) | **Get** /languages | Languages supported by TMDB
*TmdbAPI* | [**ListRegions**](docs/TmdbAPI.md#listregions) | **Get** /regions | Regions supported by TMDB
*TvAPI* | [**GetTvByTvId**](docs/TvAPI.md#gettvbytvid) | **Get** /tv/{tvId} | Get TV details
*TvAPI* | [**GetTvRatings**](docs/TvAPI.md#gettvratings) | **Get** /tv/{tvId}/ratings | Get TV ratings
*TvAPI* | [**GetTvRecommendations**](docs/TvAPI.md#gettvrecommendations) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
*TvAPI* | [**GetTvSeasonBySeasonId**](docs/TvAPI.md#gettvseasonbyseasonid) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
*TvAPI* | [**GetTvSimilar**](docs/TvAPI.md#gettvsimilar) | **Get** /tv/{tvId}/similar | Get similar TV series
*UsersAPI* | [**CreateAuthResetPassword**](docs/UsersAPI.md#createauthresetpassword) | **Post** /auth/reset-password | Send a reset password email
*UsersAPI* | [**CreateAuthResetPasswordByGuid**](docs/UsersAPI.md#createauthresetpasswordbyguid) | **Post** /auth/reset-password/{guid} | Reset the password for a user
*UsersAPI* | [**CreateUser**](docs/UsersAPI.md#createuser) | **Post** /user | Create new user
*UsersAPI* | [**CreateUserImportFromPlex**](docs/UsersAPI.md#createuserimportfromplex) | **Post** /user/import-from-plex | Import all users from Plex
*UsersAPI* | [**CreateUserRegisterPushSubscription**](docs/UsersAPI.md#createuserregisterpushsubscription) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
*UsersAPI* | [**CreateUserSettingsMain**](docs/UsersAPI.md#createusersettingsmain) | **Post** /user/{userId}/settings/main | Update general settings for a user
*UsersAPI* | [**CreateUserSettingsNotifications**](docs/UsersAPI.md#createusersettingsnotifications) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
*UsersAPI* | [**CreateUserSettingsPassword**](docs/UsersAPI.md#createusersettingspassword) | **Post** /user/{userId}/settings/password | Update password for a user
*UsersAPI* | [**CreateUserSettingsPermissions**](docs/UsersAPI.md#createusersettingspermissions) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
*UsersAPI* | [**DeleteUser**](docs/UsersAPI.md#deleteuser) | **Delete** /user/{userId} | Delete user by ID
*UsersAPI* | [**DeleteUserPushSubscription**](docs/UsersAPI.md#deleteuserpushsubscription) | **Delete** /user/{userId}/pushSubscription/{key} | Delete user push subscription by key
*UsersAPI* | [**GetUser**](docs/UsersAPI.md#getuser) | **Get** /user | Get all users
*UsersAPI* | [**GetUserByUserId**](docs/UsersAPI.md#getuserbyuserid) | **Get** /user/{userId} | Get user by ID
*UsersAPI* | [**GetUserPushSubscriptionByKey**](docs/UsersAPI.md#getuserpushsubscriptionbykey) | **Get** /user/{userId}/pushSubscription/{key} | Get web push notification settings for a user
*UsersAPI* | [**GetUserPushSubscriptions**](docs/UsersAPI.md#getuserpushsubscriptions) | **Get** /user/{userId}/pushSubscriptions | Get all web push notification settings for a user
*UsersAPI* | [**GetUserQuota**](docs/UsersAPI.md#getuserquota) | **Get** /user/{userId}/quota | Get quotas for a specific user
*UsersAPI* | [**GetUserRequests**](docs/UsersAPI.md#getuserrequests) | **Get** /user/{userId}/requests | Get requests for a specific user
*UsersAPI* | [**GetUserSettingsMain**](docs/UsersAPI.md#getusersettingsmain) | **Get** /user/{userId}/settings/main | Get general settings for a user
*UsersAPI* | [**GetUserSettingsNotifications**](docs/UsersAPI.md#getusersettingsnotifications) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
*UsersAPI* | [**GetUserSettingsPassword**](docs/UsersAPI.md#getusersettingspassword) | **Get** /user/{userId}/settings/password | Get password page informatiom
*UsersAPI* | [**GetUserSettingsPermissions**](docs/UsersAPI.md#getusersettingspermissions) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
*UsersAPI* | [**GetUserWatchData**](docs/UsersAPI.md#getuserwatchdata) | **Get** /user/{userId}/watch_data | Get watch data
*UsersAPI* | [**GetUserWatchlist**](docs/UsersAPI.md#getuserwatchlist) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user
*UsersAPI* | [**PutUser**](docs/UsersAPI.md#putuser) | **Put** /user | Update batch of users
*UsersAPI* | [**UpdateUser**](docs/UsersAPI.md#updateuser) | **Put** /user/{userId} | Update a user by user ID


## Documentation For Models

 - [Cast](docs/Cast.md)
 - [Collection](docs/Collection.md)
 - [Company](docs/Company.md)
 - [CreateAuthLocalRequest](docs/CreateAuthLocalRequest.md)
 - [CreateAuthLogout2XXResponse](docs/CreateAuthLogout2XXResponse.md)
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
 - [GetAbout2XXResponse](docs/GetAbout2XXResponse.md)
 - [GetCache2XXResponse](docs/GetCache2XXResponse.md)
 - [GetCache2XXResponseApiCachesInner](docs/GetCache2XXResponseApiCachesInner.md)
 - [GetCache2XXResponseApiCachesInnerStats](docs/GetCache2XXResponseApiCachesInnerStats.md)
 - [GetCache2XXResponseImageCache](docs/GetCache2XXResponseImageCache.md)
 - [GetCache2XXResponseImageCacheTmdb](docs/GetCache2XXResponseImageCacheTmdb.md)
 - [GetDiscoverMovies2XXResponse](docs/GetDiscoverMovies2XXResponse.md)
 - [GetDiscoverMoviesGenreByGenreId2XXResponse](docs/GetDiscoverMoviesGenreByGenreId2XXResponse.md)
 - [GetDiscoverMoviesLanguageByLanguage2XXResponse](docs/GetDiscoverMoviesLanguageByLanguage2XXResponse.md)
 - [GetDiscoverMoviesStudioByStudioId2XXResponse](docs/GetDiscoverMoviesStudioByStudioId2XXResponse.md)
 - [GetDiscoverTv2XXResponse](docs/GetDiscoverTv2XXResponse.md)
 - [GetDiscoverTvGenreByGenreId2XXResponse](docs/GetDiscoverTvGenreByGenreId2XXResponse.md)
 - [GetDiscoverTvLanguageByLanguage2XXResponse](docs/GetDiscoverTvLanguageByLanguage2XXResponse.md)
 - [GetDiscoverTvNetworkByNetworkId2XXResponse](docs/GetDiscoverTvNetworkByNetworkId2XXResponse.md)
 - [GetIssue2XXResponse](docs/GetIssue2XXResponse.md)
 - [GetIssueCount2XXResponse](docs/GetIssueCount2XXResponse.md)
 - [GetMedia2XXResponse](docs/GetMedia2XXResponse.md)
 - [GetMediaWatchData2XXResponse](docs/GetMediaWatchData2XXResponse.md)
 - [GetMediaWatchData2XXResponseData](docs/GetMediaWatchData2XXResponseData.md)
 - [GetMovieRatings2XXResponse](docs/GetMovieRatings2XXResponse.md)
 - [GetMovieRatingscombined2XXResponse](docs/GetMovieRatingscombined2XXResponse.md)
 - [GetMovieRatingscombined2XXResponseImdb](docs/GetMovieRatingscombined2XXResponseImdb.md)
 - [GetPersonCombinedCredits2XXResponse](docs/GetPersonCombinedCredits2XXResponse.md)
 - [GetPlexSync2XXResponse](docs/GetPlexSync2XXResponse.md)
 - [GetRequestCount2XXResponse](docs/GetRequestCount2XXResponse.md)
 - [GetSearch2XXResponse](docs/GetSearch2XXResponse.md)
 - [GetSearch2XXResponseResultsInner](docs/GetSearch2XXResponseResultsInner.md)
 - [GetSearchCompany2XXResponse](docs/GetSearchCompany2XXResponse.md)
 - [GetSearchKeyword2XXResponse](docs/GetSearchKeyword2XXResponse.md)
 - [GetServiceRadarrByRadarrId2XXResponse](docs/GetServiceRadarrByRadarrId2XXResponse.md)
 - [GetServiceSonarrBySonarrId2XXResponse](docs/GetServiceSonarrBySonarrId2XXResponse.md)
 - [GetStatus2XXResponse](docs/GetStatus2XXResponse.md)
 - [GetStatusAppdata2XXResponse](docs/GetStatusAppdata2XXResponse.md)
 - [GetTvRatings2XXResponse](docs/GetTvRatings2XXResponse.md)
 - [GetUser2XXResponse](docs/GetUser2XXResponse.md)
 - [GetUserPushSubscriptions2XXResponse](docs/GetUserPushSubscriptions2XXResponse.md)
 - [GetUserQuota2XXResponse](docs/GetUserQuota2XXResponse.md)
 - [GetUserQuota2XXResponseMovie](docs/GetUserQuota2XXResponseMovie.md)
 - [GetUserRequests2XXResponse](docs/GetUserRequests2XXResponse.md)
 - [GetUserSettingsMain2XXResponse](docs/GetUserSettingsMain2XXResponse.md)
 - [GetUserSettingsPassword2XXResponse](docs/GetUserSettingsPassword2XXResponse.md)
 - [GetUserSettingsPermissions2XXResponse](docs/GetUserSettingsPermissions2XXResponse.md)
 - [GetUserWatchData2XXResponse](docs/GetUserWatchData2XXResponse.md)
 - [GetUserWatchlist2XXResponse](docs/GetUserWatchlist2XXResponse.md)
 - [GetUserWatchlist2XXResponseResultsInner](docs/GetUserWatchlist2XXResponseResultsInner.md)
 - [GotifySettings](docs/GotifySettings.md)
 - [GotifySettingsOptions](docs/GotifySettingsOptions.md)
 - [Issue](docs/Issue.md)
 - [IssueComment](docs/IssueComment.md)
 - [Job](docs/Job.md)
 - [Keyword](docs/Keyword.md)
 - [ListDiscoverGenresliderMovie2XXResponseInner](docs/ListDiscoverGenresliderMovie2XXResponseInner.md)
 - [ListGenresMovie2XXResponseInner](docs/ListGenresMovie2XXResponseInner.md)
 - [ListGenresTv2XXResponseInner](docs/ListGenresTv2XXResponseInner.md)
 - [ListLanguages2XXResponseInner](docs/ListLanguages2XXResponseInner.md)
 - [ListLogs2XXResponseInner](docs/ListLogs2XXResponseInner.md)
 - [ListNotificationsPushoverSounds2XXResponseInner](docs/ListNotificationsPushoverSounds2XXResponseInner.md)
 - [ListPlexUsers2XXResponseInner](docs/ListPlexUsers2XXResponseInner.md)
 - [ListRegions2XXResponseInner](docs/ListRegions2XXResponseInner.md)
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
 - [TestRadarr2XXResponse](docs/TestRadarr2XXResponse.md)
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


Authentication schemes defined for the API:
### cookieAuth

- **Type**: API key
- **API key parameter name**: connect.sid
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: cookieAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		overseerr.ContextAPIKeys,
		map[string]overseerr.APIKey{
			"cookieAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### apiKey

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		overseerr.ContextAPIKeys,
		map[string]overseerr.APIKey{
			"apiKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


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



