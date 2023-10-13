/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"encoding/json"
)

// CreditCrew struct for CreditCrew
type CreditCrew struct {
	Id *float32 `json:"id,omitempty"`
	OriginalLanguage *string `json:"originalLanguage,omitempty"`
	EpisodeCount *float32 `json:"episodeCount,omitempty"`
	Overview *string `json:"overview,omitempty"`
	OriginCountry []*string `json:"originCountry,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	VoteCount *float32 `json:"voteCount,omitempty"`
	Name *string `json:"name,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Popularity *float32 `json:"popularity,omitempty"`
	CreditId *string `json:"creditId,omitempty"`
	BackdropPath *string `json:"backdropPath,omitempty"`
	FirstAirDate *string `json:"firstAirDate,omitempty"`
	VoteAverage *float32 `json:"voteAverage,omitempty"`
	GenreIds []*float32 `json:"genreIds,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	OriginalTitle *string `json:"originalTitle,omitempty"`
	Video *bool `json:"video,omitempty"`
	Title *string `json:"title,omitempty"`
	Adult *bool `json:"adult,omitempty"`
	ReleaseDate *string `json:"releaseDate,omitempty"`
	Department *string `json:"department,omitempty"`
	Job *string `json:"job,omitempty"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditCrew CreditCrew

// NewCreditCrew instantiates a new CreditCrew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditCrew() *CreditCrew {
	this := CreditCrew{}
	return &this
}

// NewCreditCrewWithDefaults instantiates a new CreditCrew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditCrewWithDefaults() *CreditCrew {
	this := CreditCrew{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreditCrew) GetId() float32 {
	if o == nil || isNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetIdOk() (*float32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreditCrew) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *CreditCrew) SetId(v float32) {
	o.Id = &v
}

// GetOriginalLanguage returns the OriginalLanguage field value if set, zero value otherwise.
func (o *CreditCrew) GetOriginalLanguage() string {
	if o == nil || isNil(o.OriginalLanguage) {
		var ret string
		return ret
	}
	return *o.OriginalLanguage
}

// GetOriginalLanguageOk returns a tuple with the OriginalLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetOriginalLanguageOk() (*string, bool) {
	if o == nil || isNil(o.OriginalLanguage) {
    return nil, false
	}
	return o.OriginalLanguage, true
}

// HasOriginalLanguage returns a boolean if a field has been set.
func (o *CreditCrew) HasOriginalLanguage() bool {
	if o != nil && !isNil(o.OriginalLanguage) {
		return true
	}

	return false
}

// SetOriginalLanguage gets a reference to the given string and assigns it to the OriginalLanguage field.
func (o *CreditCrew) SetOriginalLanguage(v string) {
	o.OriginalLanguage = &v
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *CreditCrew) GetEpisodeCount() float32 {
	if o == nil || isNil(o.EpisodeCount) {
		var ret float32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetEpisodeCountOk() (*float32, bool) {
	if o == nil || isNil(o.EpisodeCount) {
    return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *CreditCrew) HasEpisodeCount() bool {
	if o != nil && !isNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given float32 and assigns it to the EpisodeCount field.
func (o *CreditCrew) SetEpisodeCount(v float32) {
	o.EpisodeCount = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *CreditCrew) GetOverview() string {
	if o == nil || isNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetOverviewOk() (*string, bool) {
	if o == nil || isNil(o.Overview) {
    return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *CreditCrew) HasOverview() bool {
	if o != nil && !isNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *CreditCrew) SetOverview(v string) {
	o.Overview = &v
}

// GetOriginCountry returns the OriginCountry field value if set, zero value otherwise.
func (o *CreditCrew) GetOriginCountry() []*string {
	if o == nil || isNil(o.OriginCountry) {
		var ret []*string
		return ret
	}
	return o.OriginCountry
}

// GetOriginCountryOk returns a tuple with the OriginCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetOriginCountryOk() ([]*string, bool) {
	if o == nil || isNil(o.OriginCountry) {
    return nil, false
	}
	return o.OriginCountry, true
}

// HasOriginCountry returns a boolean if a field has been set.
func (o *CreditCrew) HasOriginCountry() bool {
	if o != nil && !isNil(o.OriginCountry) {
		return true
	}

	return false
}

// SetOriginCountry gets a reference to the given []string and assigns it to the OriginCountry field.
func (o *CreditCrew) SetOriginCountry(v []*string) {
	o.OriginCountry = v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *CreditCrew) GetOriginalName() string {
	if o == nil || isNil(o.OriginalName) {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetOriginalNameOk() (*string, bool) {
	if o == nil || isNil(o.OriginalName) {
    return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *CreditCrew) HasOriginalName() bool {
	if o != nil && !isNil(o.OriginalName) {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *CreditCrew) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *CreditCrew) GetVoteCount() float32 {
	if o == nil || isNil(o.VoteCount) {
		var ret float32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetVoteCountOk() (*float32, bool) {
	if o == nil || isNil(o.VoteCount) {
    return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *CreditCrew) HasVoteCount() bool {
	if o != nil && !isNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given float32 and assigns it to the VoteCount field.
func (o *CreditCrew) SetVoteCount(v float32) {
	o.VoteCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreditCrew) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreditCrew) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreditCrew) SetName(v string) {
	o.Name = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *CreditCrew) GetMediaType() string {
	if o == nil || isNil(o.MediaType) {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetMediaTypeOk() (*string, bool) {
	if o == nil || isNil(o.MediaType) {
    return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *CreditCrew) HasMediaType() bool {
	if o != nil && !isNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *CreditCrew) SetMediaType(v string) {
	o.MediaType = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *CreditCrew) GetPopularity() float32 {
	if o == nil || isNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetPopularityOk() (*float32, bool) {
	if o == nil || isNil(o.Popularity) {
    return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *CreditCrew) HasPopularity() bool {
	if o != nil && !isNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *CreditCrew) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *CreditCrew) GetCreditId() string {
	if o == nil || isNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetCreditIdOk() (*string, bool) {
	if o == nil || isNil(o.CreditId) {
    return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *CreditCrew) HasCreditId() bool {
	if o != nil && !isNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *CreditCrew) SetCreditId(v string) {
	o.CreditId = &v
}

// GetBackdropPath returns the BackdropPath field value if set, zero value otherwise.
func (o *CreditCrew) GetBackdropPath() string {
	if o == nil || isNil(o.BackdropPath) {
		var ret string
		return ret
	}
	return *o.BackdropPath
}

// GetBackdropPathOk returns a tuple with the BackdropPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetBackdropPathOk() (*string, bool) {
	if o == nil || isNil(o.BackdropPath) {
    return nil, false
	}
	return o.BackdropPath, true
}

// HasBackdropPath returns a boolean if a field has been set.
func (o *CreditCrew) HasBackdropPath() bool {
	if o != nil && !isNil(o.BackdropPath) {
		return true
	}

	return false
}

// SetBackdropPath gets a reference to the given string and assigns it to the BackdropPath field.
func (o *CreditCrew) SetBackdropPath(v string) {
	o.BackdropPath = &v
}

// GetFirstAirDate returns the FirstAirDate field value if set, zero value otherwise.
func (o *CreditCrew) GetFirstAirDate() string {
	if o == nil || isNil(o.FirstAirDate) {
		var ret string
		return ret
	}
	return *o.FirstAirDate
}

// GetFirstAirDateOk returns a tuple with the FirstAirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetFirstAirDateOk() (*string, bool) {
	if o == nil || isNil(o.FirstAirDate) {
    return nil, false
	}
	return o.FirstAirDate, true
}

// HasFirstAirDate returns a boolean if a field has been set.
func (o *CreditCrew) HasFirstAirDate() bool {
	if o != nil && !isNil(o.FirstAirDate) {
		return true
	}

	return false
}

// SetFirstAirDate gets a reference to the given string and assigns it to the FirstAirDate field.
func (o *CreditCrew) SetFirstAirDate(v string) {
	o.FirstAirDate = &v
}

// GetVoteAverage returns the VoteAverage field value if set, zero value otherwise.
func (o *CreditCrew) GetVoteAverage() float32 {
	if o == nil || isNil(o.VoteAverage) {
		var ret float32
		return ret
	}
	return *o.VoteAverage
}

// GetVoteAverageOk returns a tuple with the VoteAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetVoteAverageOk() (*float32, bool) {
	if o == nil || isNil(o.VoteAverage) {
    return nil, false
	}
	return o.VoteAverage, true
}

// HasVoteAverage returns a boolean if a field has been set.
func (o *CreditCrew) HasVoteAverage() bool {
	if o != nil && !isNil(o.VoteAverage) {
		return true
	}

	return false
}

// SetVoteAverage gets a reference to the given float32 and assigns it to the VoteAverage field.
func (o *CreditCrew) SetVoteAverage(v float32) {
	o.VoteAverage = &v
}

// GetGenreIds returns the GenreIds field value if set, zero value otherwise.
func (o *CreditCrew) GetGenreIds() []*float32 {
	if o == nil || isNil(o.GenreIds) {
		var ret []*float32
		return ret
	}
	return o.GenreIds
}

// GetGenreIdsOk returns a tuple with the GenreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetGenreIdsOk() ([]*float32, bool) {
	if o == nil || isNil(o.GenreIds) {
    return nil, false
	}
	return o.GenreIds, true
}

// HasGenreIds returns a boolean if a field has been set.
func (o *CreditCrew) HasGenreIds() bool {
	if o != nil && !isNil(o.GenreIds) {
		return true
	}

	return false
}

// SetGenreIds gets a reference to the given []float32 and assigns it to the GenreIds field.
func (o *CreditCrew) SetGenreIds(v []*float32) {
	o.GenreIds = v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *CreditCrew) GetPosterPath() string {
	if o == nil || isNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetPosterPathOk() (*string, bool) {
	if o == nil || isNil(o.PosterPath) {
    return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *CreditCrew) HasPosterPath() bool {
	if o != nil && !isNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *CreditCrew) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetOriginalTitle returns the OriginalTitle field value if set, zero value otherwise.
func (o *CreditCrew) GetOriginalTitle() string {
	if o == nil || isNil(o.OriginalTitle) {
		var ret string
		return ret
	}
	return *o.OriginalTitle
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetOriginalTitleOk() (*string, bool) {
	if o == nil || isNil(o.OriginalTitle) {
    return nil, false
	}
	return o.OriginalTitle, true
}

// HasOriginalTitle returns a boolean if a field has been set.
func (o *CreditCrew) HasOriginalTitle() bool {
	if o != nil && !isNil(o.OriginalTitle) {
		return true
	}

	return false
}

// SetOriginalTitle gets a reference to the given string and assigns it to the OriginalTitle field.
func (o *CreditCrew) SetOriginalTitle(v string) {
	o.OriginalTitle = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *CreditCrew) GetVideo() bool {
	if o == nil || isNil(o.Video) {
		var ret bool
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetVideoOk() (*bool, bool) {
	if o == nil || isNil(o.Video) {
    return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *CreditCrew) HasVideo() bool {
	if o != nil && !isNil(o.Video) {
		return true
	}

	return false
}

// SetVideo gets a reference to the given bool and assigns it to the Video field.
func (o *CreditCrew) SetVideo(v bool) {
	o.Video = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreditCrew) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreditCrew) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreditCrew) SetTitle(v string) {
	o.Title = &v
}

// GetAdult returns the Adult field value if set, zero value otherwise.
func (o *CreditCrew) GetAdult() bool {
	if o == nil || isNil(o.Adult) {
		var ret bool
		return ret
	}
	return *o.Adult
}

// GetAdultOk returns a tuple with the Adult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetAdultOk() (*bool, bool) {
	if o == nil || isNil(o.Adult) {
    return nil, false
	}
	return o.Adult, true
}

// HasAdult returns a boolean if a field has been set.
func (o *CreditCrew) HasAdult() bool {
	if o != nil && !isNil(o.Adult) {
		return true
	}

	return false
}

// SetAdult gets a reference to the given bool and assigns it to the Adult field.
func (o *CreditCrew) SetAdult(v bool) {
	o.Adult = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *CreditCrew) GetReleaseDate() string {
	if o == nil || isNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetReleaseDateOk() (*string, bool) {
	if o == nil || isNil(o.ReleaseDate) {
    return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *CreditCrew) HasReleaseDate() bool {
	if o != nil && !isNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *CreditCrew) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *CreditCrew) GetDepartment() string {
	if o == nil || isNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetDepartmentOk() (*string, bool) {
	if o == nil || isNil(o.Department) {
    return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *CreditCrew) HasDepartment() bool {
	if o != nil && !isNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *CreditCrew) SetDepartment(v string) {
	o.Department = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreditCrew) GetJob() string {
	if o == nil || isNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetJobOk() (*string, bool) {
	if o == nil || isNil(o.Job) {
    return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreditCrew) HasJob() bool {
	if o != nil && !isNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *CreditCrew) SetJob(v string) {
	o.Job = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *CreditCrew) GetMediaInfo() MediaInfo {
	if o == nil || isNil(o.MediaInfo) {
		var ret MediaInfo
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditCrew) GetMediaInfoOk() (*MediaInfo, bool) {
	if o == nil || isNil(o.MediaInfo) {
    return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *CreditCrew) HasMediaInfo() bool {
	if o != nil && !isNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfo and assigns it to the MediaInfo field.
func (o *CreditCrew) SetMediaInfo(v MediaInfo) {
	o.MediaInfo = &v
}

func (o CreditCrew) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OriginalLanguage) {
		toSerialize["originalLanguage"] = o.OriginalLanguage
	}
	if !isNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	if !isNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !isNil(o.OriginCountry) {
		toSerialize["originCountry"] = o.OriginCountry
	}
	if !isNil(o.OriginalName) {
		toSerialize["originalName"] = o.OriginalName
	}
	if !isNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !isNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !isNil(o.CreditId) {
		toSerialize["creditId"] = o.CreditId
	}
	if !isNil(o.BackdropPath) {
		toSerialize["backdropPath"] = o.BackdropPath
	}
	if !isNil(o.FirstAirDate) {
		toSerialize["firstAirDate"] = o.FirstAirDate
	}
	if !isNil(o.VoteAverage) {
		toSerialize["voteAverage"] = o.VoteAverage
	}
	if !isNil(o.GenreIds) {
		toSerialize["genreIds"] = o.GenreIds
	}
	if !isNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !isNil(o.OriginalTitle) {
		toSerialize["originalTitle"] = o.OriginalTitle
	}
	if !isNil(o.Video) {
		toSerialize["video"] = o.Video
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Adult) {
		toSerialize["adult"] = o.Adult
	}
	if !isNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !isNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !isNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !isNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditCrew) UnmarshalJSON(bytes []byte) (err error) {
	varCreditCrew := _CreditCrew{}

	if err = json.Unmarshal(bytes, &varCreditCrew); err == nil {
		*o = CreditCrew(varCreditCrew)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "originalLanguage")
		delete(additionalProperties, "episodeCount")
		delete(additionalProperties, "overview")
		delete(additionalProperties, "originCountry")
		delete(additionalProperties, "originalName")
		delete(additionalProperties, "voteCount")
		delete(additionalProperties, "name")
		delete(additionalProperties, "mediaType")
		delete(additionalProperties, "popularity")
		delete(additionalProperties, "creditId")
		delete(additionalProperties, "backdropPath")
		delete(additionalProperties, "firstAirDate")
		delete(additionalProperties, "voteAverage")
		delete(additionalProperties, "genreIds")
		delete(additionalProperties, "posterPath")
		delete(additionalProperties, "originalTitle")
		delete(additionalProperties, "video")
		delete(additionalProperties, "title")
		delete(additionalProperties, "adult")
		delete(additionalProperties, "releaseDate")
		delete(additionalProperties, "department")
		delete(additionalProperties, "job")
		delete(additionalProperties, "mediaInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditCrew struct {
	value *CreditCrew
	isSet bool
}

func (v NullableCreditCrew) Get() *CreditCrew {
	return v.value
}

func (v *NullableCreditCrew) Set(val *CreditCrew) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditCrew) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditCrew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditCrew(val *CreditCrew) *NullableCreditCrew {
	return &NullableCreditCrew{value: val, isSet: true}
}

func (v NullableCreditCrew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditCrew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


