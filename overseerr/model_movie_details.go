/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// checks if the MovieDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieDetails{}

// MovieDetails struct for MovieDetails
type MovieDetails struct {
	Id *float32 `json:"id,omitempty"`
	ImdbId *string `json:"imdbId,omitempty"`
	Adult *bool `json:"adult,omitempty"`
	BackdropPath *string `json:"backdropPath,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	Budget *float32 `json:"budget,omitempty"`
	Genres []Genre `json:"genres,omitempty"`
	Homepage *string `json:"homepage,omitempty"`
	RelatedVideos []RelatedVideo `json:"relatedVideos,omitempty"`
	OriginalLanguage *string `json:"originalLanguage,omitempty"`
	OriginalTitle *string `json:"originalTitle,omitempty"`
	Overview *string `json:"overview,omitempty"`
	Popularity *float32 `json:"popularity,omitempty"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies,omitempty"`
	ProductionCountries []MovieDetailsProductionCountriesInner `json:"productionCountries,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
	Releases *MovieDetailsReleases `json:"releases,omitempty"`
	Revenue NullableFloat32 `json:"revenue,omitempty"`
	Runtime *float32 `json:"runtime,omitempty"`
	SpokenLanguages []SpokenLanguage `json:"spokenLanguages,omitempty"`
	Status *string `json:"status,omitempty"`
	Tagline *string `json:"tagline,omitempty"`
	Title *string `json:"title,omitempty"`
	Video *bool `json:"video,omitempty"`
	VoteAverage *float32 `json:"voteAverage,omitempty"`
	VoteCount *float32 `json:"voteCount,omitempty"`
	Credits *MovieDetailsCredits `json:"credits,omitempty"`
	Collection *MovieDetailsCollection `json:"collection,omitempty"`
	ExternalIds *ExternalIds `json:"externalIds,omitempty"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
	WatchProviders [][]WatchProvidersInner `json:"watchProviders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovieDetails MovieDetails

// NewMovieDetails instantiates a new MovieDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieDetails() *MovieDetails {
	this := MovieDetails{}
	return &this
}

// NewMovieDetailsWithDefaults instantiates a new MovieDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieDetailsWithDefaults() *MovieDetails {
	this := MovieDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MovieDetails) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MovieDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *MovieDetails) SetId(v float32) {
	o.Id = &v
}

// GetImdbId returns the ImdbId field value if set, zero value otherwise.
func (o *MovieDetails) GetImdbId() string {
	if o == nil || IsNil(o.ImdbId) {
		var ret string
		return ret
	}
	return *o.ImdbId
}

// GetImdbIdOk returns a tuple with the ImdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetImdbIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImdbId) {
		return nil, false
	}
	return o.ImdbId, true
}

// HasImdbId returns a boolean if a field has been set.
func (o *MovieDetails) HasImdbId() bool {
	if o != nil && !IsNil(o.ImdbId) {
		return true
	}

	return false
}

// SetImdbId gets a reference to the given string and assigns it to the ImdbId field.
func (o *MovieDetails) SetImdbId(v string) {
	o.ImdbId = &v
}

// GetAdult returns the Adult field value if set, zero value otherwise.
func (o *MovieDetails) GetAdult() bool {
	if o == nil || IsNil(o.Adult) {
		var ret bool
		return ret
	}
	return *o.Adult
}

// GetAdultOk returns a tuple with the Adult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetAdultOk() (*bool, bool) {
	if o == nil || IsNil(o.Adult) {
		return nil, false
	}
	return o.Adult, true
}

// HasAdult returns a boolean if a field has been set.
func (o *MovieDetails) HasAdult() bool {
	if o != nil && !IsNil(o.Adult) {
		return true
	}

	return false
}

// SetAdult gets a reference to the given bool and assigns it to the Adult field.
func (o *MovieDetails) SetAdult(v bool) {
	o.Adult = &v
}

// GetBackdropPath returns the BackdropPath field value if set, zero value otherwise.
func (o *MovieDetails) GetBackdropPath() string {
	if o == nil || IsNil(o.BackdropPath) {
		var ret string
		return ret
	}
	return *o.BackdropPath
}

// GetBackdropPathOk returns a tuple with the BackdropPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetBackdropPathOk() (*string, bool) {
	if o == nil || IsNil(o.BackdropPath) {
		return nil, false
	}
	return o.BackdropPath, true
}

// HasBackdropPath returns a boolean if a field has been set.
func (o *MovieDetails) HasBackdropPath() bool {
	if o != nil && !IsNil(o.BackdropPath) {
		return true
	}

	return false
}

// SetBackdropPath gets a reference to the given string and assigns it to the BackdropPath field.
func (o *MovieDetails) SetBackdropPath(v string) {
	o.BackdropPath = &v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *MovieDetails) GetPosterPath() string {
	if o == nil || IsNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetPosterPathOk() (*string, bool) {
	if o == nil || IsNil(o.PosterPath) {
		return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *MovieDetails) HasPosterPath() bool {
	if o != nil && !IsNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *MovieDetails) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *MovieDetails) GetBudget() float32 {
	if o == nil || IsNil(o.Budget) {
		var ret float32
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *MovieDetails) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given float32 and assigns it to the Budget field.
func (o *MovieDetails) SetBudget(v float32) {
	o.Budget = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *MovieDetails) GetGenres() []Genre {
	if o == nil || IsNil(o.Genres) {
		var ret []Genre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetGenresOk() ([]Genre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *MovieDetails) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []Genre and assigns it to the Genres field.
func (o *MovieDetails) SetGenres(v []Genre) {
	o.Genres = v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *MovieDetails) GetHomepage() string {
	if o == nil || IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetHomepageOk() (*string, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *MovieDetails) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *MovieDetails) SetHomepage(v string) {
	o.Homepage = &v
}

// GetRelatedVideos returns the RelatedVideos field value if set, zero value otherwise.
func (o *MovieDetails) GetRelatedVideos() []RelatedVideo {
	if o == nil || IsNil(o.RelatedVideos) {
		var ret []RelatedVideo
		return ret
	}
	return o.RelatedVideos
}

// GetRelatedVideosOk returns a tuple with the RelatedVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetRelatedVideosOk() ([]RelatedVideo, bool) {
	if o == nil || IsNil(o.RelatedVideos) {
		return nil, false
	}
	return o.RelatedVideos, true
}

// HasRelatedVideos returns a boolean if a field has been set.
func (o *MovieDetails) HasRelatedVideos() bool {
	if o != nil && !IsNil(o.RelatedVideos) {
		return true
	}

	return false
}

// SetRelatedVideos gets a reference to the given []RelatedVideo and assigns it to the RelatedVideos field.
func (o *MovieDetails) SetRelatedVideos(v []RelatedVideo) {
	o.RelatedVideos = v
}

// GetOriginalLanguage returns the OriginalLanguage field value if set, zero value otherwise.
func (o *MovieDetails) GetOriginalLanguage() string {
	if o == nil || IsNil(o.OriginalLanguage) {
		var ret string
		return ret
	}
	return *o.OriginalLanguage
}

// GetOriginalLanguageOk returns a tuple with the OriginalLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetOriginalLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalLanguage) {
		return nil, false
	}
	return o.OriginalLanguage, true
}

// HasOriginalLanguage returns a boolean if a field has been set.
func (o *MovieDetails) HasOriginalLanguage() bool {
	if o != nil && !IsNil(o.OriginalLanguage) {
		return true
	}

	return false
}

// SetOriginalLanguage gets a reference to the given string and assigns it to the OriginalLanguage field.
func (o *MovieDetails) SetOriginalLanguage(v string) {
	o.OriginalLanguage = &v
}

// GetOriginalTitle returns the OriginalTitle field value if set, zero value otherwise.
func (o *MovieDetails) GetOriginalTitle() string {
	if o == nil || IsNil(o.OriginalTitle) {
		var ret string
		return ret
	}
	return *o.OriginalTitle
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetOriginalTitleOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalTitle) {
		return nil, false
	}
	return o.OriginalTitle, true
}

// HasOriginalTitle returns a boolean if a field has been set.
func (o *MovieDetails) HasOriginalTitle() bool {
	if o != nil && !IsNil(o.OriginalTitle) {
		return true
	}

	return false
}

// SetOriginalTitle gets a reference to the given string and assigns it to the OriginalTitle field.
func (o *MovieDetails) SetOriginalTitle(v string) {
	o.OriginalTitle = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *MovieDetails) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *MovieDetails) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *MovieDetails) SetOverview(v string) {
	o.Overview = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *MovieDetails) GetPopularity() float32 {
	if o == nil || IsNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetPopularityOk() (*float32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *MovieDetails) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *MovieDetails) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetProductionCompanies returns the ProductionCompanies field value if set, zero value otherwise.
func (o *MovieDetails) GetProductionCompanies() []ProductionCompany {
	if o == nil || IsNil(o.ProductionCompanies) {
		var ret []ProductionCompany
		return ret
	}
	return o.ProductionCompanies
}

// GetProductionCompaniesOk returns a tuple with the ProductionCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetProductionCompaniesOk() ([]ProductionCompany, bool) {
	if o == nil || IsNil(o.ProductionCompanies) {
		return nil, false
	}
	return o.ProductionCompanies, true
}

// HasProductionCompanies returns a boolean if a field has been set.
func (o *MovieDetails) HasProductionCompanies() bool {
	if o != nil && !IsNil(o.ProductionCompanies) {
		return true
	}

	return false
}

// SetProductionCompanies gets a reference to the given []ProductionCompany and assigns it to the ProductionCompanies field.
func (o *MovieDetails) SetProductionCompanies(v []ProductionCompany) {
	o.ProductionCompanies = v
}

// GetProductionCountries returns the ProductionCountries field value if set, zero value otherwise.
func (o *MovieDetails) GetProductionCountries() []MovieDetailsProductionCountriesInner {
	if o == nil || IsNil(o.ProductionCountries) {
		var ret []MovieDetailsProductionCountriesInner
		return ret
	}
	return o.ProductionCountries
}

// GetProductionCountriesOk returns a tuple with the ProductionCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetProductionCountriesOk() ([]MovieDetailsProductionCountriesInner, bool) {
	if o == nil || IsNil(o.ProductionCountries) {
		return nil, false
	}
	return o.ProductionCountries, true
}

// HasProductionCountries returns a boolean if a field has been set.
func (o *MovieDetails) HasProductionCountries() bool {
	if o != nil && !IsNil(o.ProductionCountries) {
		return true
	}

	return false
}

// SetProductionCountries gets a reference to the given []MovieDetailsProductionCountriesInner and assigns it to the ProductionCountries field.
func (o *MovieDetails) SetProductionCountries(v []MovieDetailsProductionCountriesInner) {
	o.ProductionCountries = v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *MovieDetails) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *MovieDetails) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *MovieDetails) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *MovieDetails) GetReleases() MovieDetailsReleases {
	if o == nil || IsNil(o.Releases) {
		var ret MovieDetailsReleases
		return ret
	}
	return *o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetReleasesOk() (*MovieDetailsReleases, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *MovieDetails) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given MovieDetailsReleases and assigns it to the Releases field.
func (o *MovieDetails) SetReleases(v MovieDetailsReleases) {
	o.Releases = &v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MovieDetails) GetRevenue() float32 {
	if o == nil || IsNil(o.Revenue.Get()) {
		var ret float32
		return ret
	}
	return *o.Revenue.Get()
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MovieDetails) GetRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revenue.Get(), o.Revenue.IsSet()
}

// HasRevenue returns a boolean if a field has been set.
func (o *MovieDetails) HasRevenue() bool {
	if o != nil && o.Revenue.IsSet() {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given NullableFloat32 and assigns it to the Revenue field.
func (o *MovieDetails) SetRevenue(v float32) {
	o.Revenue.Set(&v)
}
// SetRevenueNil sets the value for Revenue to be an explicit nil
func (o *MovieDetails) SetRevenueNil() {
	o.Revenue.Set(nil)
}

// UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
func (o *MovieDetails) UnsetRevenue() {
	o.Revenue.Unset()
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *MovieDetails) GetRuntime() float32 {
	if o == nil || IsNil(o.Runtime) {
		var ret float32
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetRuntimeOk() (*float32, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *MovieDetails) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given float32 and assigns it to the Runtime field.
func (o *MovieDetails) SetRuntime(v float32) {
	o.Runtime = &v
}

// GetSpokenLanguages returns the SpokenLanguages field value if set, zero value otherwise.
func (o *MovieDetails) GetSpokenLanguages() []SpokenLanguage {
	if o == nil || IsNil(o.SpokenLanguages) {
		var ret []SpokenLanguage
		return ret
	}
	return o.SpokenLanguages
}

// GetSpokenLanguagesOk returns a tuple with the SpokenLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetSpokenLanguagesOk() ([]SpokenLanguage, bool) {
	if o == nil || IsNil(o.SpokenLanguages) {
		return nil, false
	}
	return o.SpokenLanguages, true
}

// HasSpokenLanguages returns a boolean if a field has been set.
func (o *MovieDetails) HasSpokenLanguages() bool {
	if o != nil && !IsNil(o.SpokenLanguages) {
		return true
	}

	return false
}

// SetSpokenLanguages gets a reference to the given []SpokenLanguage and assigns it to the SpokenLanguages field.
func (o *MovieDetails) SetSpokenLanguages(v []SpokenLanguage) {
	o.SpokenLanguages = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MovieDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MovieDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MovieDetails) SetStatus(v string) {
	o.Status = &v
}

// GetTagline returns the Tagline field value if set, zero value otherwise.
func (o *MovieDetails) GetTagline() string {
	if o == nil || IsNil(o.Tagline) {
		var ret string
		return ret
	}
	return *o.Tagline
}

// GetTaglineOk returns a tuple with the Tagline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetTaglineOk() (*string, bool) {
	if o == nil || IsNil(o.Tagline) {
		return nil, false
	}
	return o.Tagline, true
}

// HasTagline returns a boolean if a field has been set.
func (o *MovieDetails) HasTagline() bool {
	if o != nil && !IsNil(o.Tagline) {
		return true
	}

	return false
}

// SetTagline gets a reference to the given string and assigns it to the Tagline field.
func (o *MovieDetails) SetTagline(v string) {
	o.Tagline = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MovieDetails) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MovieDetails) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MovieDetails) SetTitle(v string) {
	o.Title = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *MovieDetails) GetVideo() bool {
	if o == nil || IsNil(o.Video) {
		var ret bool
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.Video) {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *MovieDetails) HasVideo() bool {
	if o != nil && !IsNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given bool and assigns it to the Video field.
func (o *MovieDetails) SetVideo(v bool) {
	o.Video = &v
}

// GetVoteAverage returns the VoteAverage field value if set, zero value otherwise.
func (o *MovieDetails) GetVoteAverage() float32 {
	if o == nil || IsNil(o.VoteAverage) {
		var ret float32
		return ret
	}
	return *o.VoteAverage
}

// GetVoteAverageOk returns a tuple with the VoteAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetVoteAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteAverage) {
		return nil, false
	}
	return o.VoteAverage, true
}

// HasVoteAverage returns a boolean if a field has been set.
func (o *MovieDetails) HasVoteAverage() bool {
	if o != nil && !IsNil(o.VoteAverage) {
		return true
	}

	return false
}

// SetVoteAverage gets a reference to the given float32 and assigns it to the VoteAverage field.
func (o *MovieDetails) SetVoteAverage(v float32) {
	o.VoteAverage = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *MovieDetails) GetVoteCount() float32 {
	if o == nil || IsNil(o.VoteCount) {
		var ret float32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetVoteCountOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteCount) {
		return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *MovieDetails) HasVoteCount() bool {
	if o != nil && !IsNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given float32 and assigns it to the VoteCount field.
func (o *MovieDetails) SetVoteCount(v float32) {
	o.VoteCount = &v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *MovieDetails) GetCredits() MovieDetailsCredits {
	if o == nil || IsNil(o.Credits) {
		var ret MovieDetailsCredits
		return ret
	}
	return *o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetCreditsOk() (*MovieDetailsCredits, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *MovieDetails) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given MovieDetailsCredits and assigns it to the Credits field.
func (o *MovieDetails) SetCredits(v MovieDetailsCredits) {
	o.Credits = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *MovieDetails) GetCollection() MovieDetailsCollection {
	if o == nil || IsNil(o.Collection) {
		var ret MovieDetailsCollection
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetCollectionOk() (*MovieDetailsCollection, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *MovieDetails) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given MovieDetailsCollection and assigns it to the Collection field.
func (o *MovieDetails) SetCollection(v MovieDetailsCollection) {
	o.Collection = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *MovieDetails) GetExternalIds() ExternalIds {
	if o == nil || IsNil(o.ExternalIds) {
		var ret ExternalIds
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetExternalIdsOk() (*ExternalIds, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *MovieDetails) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given ExternalIds and assigns it to the ExternalIds field.
func (o *MovieDetails) SetExternalIds(v ExternalIds) {
	o.ExternalIds = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *MovieDetails) GetMediaInfo() MediaInfo {
	if o == nil || IsNil(o.MediaInfo) {
		var ret MediaInfo
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetMediaInfoOk() (*MediaInfo, bool) {
	if o == nil || IsNil(o.MediaInfo) {
		return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *MovieDetails) HasMediaInfo() bool {
	if o != nil && !IsNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfo and assigns it to the MediaInfo field.
func (o *MovieDetails) SetMediaInfo(v MediaInfo) {
	o.MediaInfo = &v
}

// GetWatchProviders returns the WatchProviders field value if set, zero value otherwise.
func (o *MovieDetails) GetWatchProviders() [][]WatchProvidersInner {
	if o == nil || IsNil(o.WatchProviders) {
		var ret [][]WatchProvidersInner
		return ret
	}
	return o.WatchProviders
}

// GetWatchProvidersOk returns a tuple with the WatchProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetails) GetWatchProvidersOk() ([][]WatchProvidersInner, bool) {
	if o == nil || IsNil(o.WatchProviders) {
		return nil, false
	}
	return o.WatchProviders, true
}

// HasWatchProviders returns a boolean if a field has been set.
func (o *MovieDetails) HasWatchProviders() bool {
	if o != nil && !IsNil(o.WatchProviders) {
		return true
	}

	return false
}

// SetWatchProviders gets a reference to the given [][]WatchProvidersInner and assigns it to the WatchProviders field.
func (o *MovieDetails) SetWatchProviders(v [][]WatchProvidersInner) {
	o.WatchProviders = v
}

func (o MovieDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImdbId) {
		toSerialize["imdbId"] = o.ImdbId
	}
	if !IsNil(o.Adult) {
		toSerialize["adult"] = o.Adult
	}
	if !IsNil(o.BackdropPath) {
		toSerialize["backdropPath"] = o.BackdropPath
	}
	if !IsNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !IsNil(o.RelatedVideos) {
		toSerialize["relatedVideos"] = o.RelatedVideos
	}
	if !IsNil(o.OriginalLanguage) {
		toSerialize["originalLanguage"] = o.OriginalLanguage
	}
	if !IsNil(o.OriginalTitle) {
		toSerialize["originalTitle"] = o.OriginalTitle
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.ProductionCompanies) {
		toSerialize["productionCompanies"] = o.ProductionCompanies
	}
	if !IsNil(o.ProductionCountries) {
		toSerialize["productionCountries"] = o.ProductionCountries
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}
	if o.Revenue.IsSet() {
		toSerialize["revenue"] = o.Revenue.Get()
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	if !IsNil(o.SpokenLanguages) {
		toSerialize["spokenLanguages"] = o.SpokenLanguages
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tagline) {
		toSerialize["tagline"] = o.Tagline
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !IsNil(o.VoteAverage) {
		toSerialize["voteAverage"] = o.VoteAverage
	}
	if !IsNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !IsNil(o.WatchProviders) {
		toSerialize["watchProviders"] = o.WatchProviders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MovieDetails) UnmarshalJSON(data []byte) (err error) {
	varMovieDetails := _MovieDetails{}

	err = json.Unmarshal(data, &varMovieDetails)

	if err != nil {
		return err
	}

	*o = MovieDetails(varMovieDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "imdbId")
		delete(additionalProperties, "adult")
		delete(additionalProperties, "backdropPath")
		delete(additionalProperties, "posterPath")
		delete(additionalProperties, "budget")
		delete(additionalProperties, "genres")
		delete(additionalProperties, "homepage")
		delete(additionalProperties, "relatedVideos")
		delete(additionalProperties, "originalLanguage")
		delete(additionalProperties, "originalTitle")
		delete(additionalProperties, "overview")
		delete(additionalProperties, "popularity")
		delete(additionalProperties, "productionCompanies")
		delete(additionalProperties, "productionCountries")
		delete(additionalProperties, "releaseDate")
		delete(additionalProperties, "releases")
		delete(additionalProperties, "revenue")
		delete(additionalProperties, "runtime")
		delete(additionalProperties, "spokenLanguages")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tagline")
		delete(additionalProperties, "title")
		delete(additionalProperties, "video")
		delete(additionalProperties, "voteAverage")
		delete(additionalProperties, "voteCount")
		delete(additionalProperties, "credits")
		delete(additionalProperties, "collection")
		delete(additionalProperties, "externalIds")
		delete(additionalProperties, "mediaInfo")
		delete(additionalProperties, "watchProviders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovieDetails struct {
	value *MovieDetails
	isSet bool
}

func (v NullableMovieDetails) Get() *MovieDetails {
	return v.value
}

func (v *NullableMovieDetails) Set(val *MovieDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieDetails(val *MovieDetails) *NullableMovieDetails {
	return &NullableMovieDetails{value: val, isSet: true}
}

func (v NullableMovieDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


