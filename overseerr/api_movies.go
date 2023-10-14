/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MoviesAPIService MoviesAPI service
type MoviesAPIService service
type ApiGetMovieByMovieIdRequest struct {
	ctx context.Context
	ApiService *MoviesAPIService
	movieId float32
	language *string
}

func (r ApiGetMovieByMovieIdRequest) Language(language string) ApiGetMovieByMovieIdRequest {
	r.language = &language
	return r
}

func (r ApiGetMovieByMovieIdRequest) Execute() (*MovieDetails, *http.Response, error) {
	return r.ApiService.GetMovieByMovieIdExecute(r)
}

/*
GetMovieByMovieId Get movie details

Returns full movie details in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param movieId
 @return ApiGetMovieByMovieIdRequest
*/
func (a *MoviesAPIService) GetMovieByMovieId(ctx context.Context, movieId float32) ApiGetMovieByMovieIdRequest {
	return ApiGetMovieByMovieIdRequest{
		ApiService: a,
		ctx: ctx,
		movieId: movieId,
	}
}

// Execute executes the request
//  @return MovieDetails
func (a *MoviesAPIService) GetMovieByMovieIdExecute(r ApiGetMovieByMovieIdRequest) (*MovieDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MovieDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieByMovieId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterToString(r.movieId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.language != nil {
		localVarQueryParams.Add("language", parameterToString(*r.language, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiGetMovieRatingsRequest struct {
	ctx context.Context
	ApiService *MoviesAPIService
	movieId float32
}

func (r ApiGetMovieRatingsRequest) Execute() (*GetMovieRatings200Response, *http.Response, error) {
	return r.ApiService.GetMovieRatingsExecute(r)
}

/*
GetMovieRatings Get movie ratings

Returns ratings based on the provided movieId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param movieId
 @return ApiGetMovieRatingsRequest
*/
func (a *MoviesAPIService) GetMovieRatings(ctx context.Context, movieId float32) ApiGetMovieRatingsRequest {
	return ApiGetMovieRatingsRequest{
		ApiService: a,
		ctx: ctx,
		movieId: movieId,
	}
}

// Execute executes the request
//  @return GetMovieRatings200Response
func (a *MoviesAPIService) GetMovieRatingsExecute(r ApiGetMovieRatingsRequest) (*GetMovieRatings200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMovieRatings200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieRatings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/ratings"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterToString(r.movieId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiGetMovieRatingscombinedRequest struct {
	ctx context.Context
	ApiService *MoviesAPIService
	movieId float32
}

func (r ApiGetMovieRatingscombinedRequest) Execute() (*GetMovieRatingscombined200Response, *http.Response, error) {
	return r.ApiService.GetMovieRatingscombinedExecute(r)
}

/*
GetMovieRatingscombined Get RT and IMDB movie ratings combined

Returns ratings from RottenTomatoes and IMDB based on the provided movieId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param movieId
 @return ApiGetMovieRatingscombinedRequest
*/
func (a *MoviesAPIService) GetMovieRatingscombined(ctx context.Context, movieId float32) ApiGetMovieRatingscombinedRequest {
	return ApiGetMovieRatingscombinedRequest{
		ApiService: a,
		ctx: ctx,
		movieId: movieId,
	}
}

// Execute executes the request
//  @return GetMovieRatingscombined200Response
func (a *MoviesAPIService) GetMovieRatingscombinedExecute(r ApiGetMovieRatingscombinedRequest) (*GetMovieRatingscombined200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMovieRatingscombined200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieRatingscombined")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/ratingscombined"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterToString(r.movieId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiGetMovieRecommendationsRequest struct {
	ctx context.Context
	ApiService *MoviesAPIService
	movieId float32
	page *float32
	language *string
}

func (r ApiGetMovieRecommendationsRequest) Page(page float32) ApiGetMovieRecommendationsRequest {
	r.page = &page
	return r
}

func (r ApiGetMovieRecommendationsRequest) Language(language string) ApiGetMovieRecommendationsRequest {
	r.language = &language
	return r
}

func (r ApiGetMovieRecommendationsRequest) Execute() (*GetDiscoverMovies200Response, *http.Response, error) {
	return r.ApiService.GetMovieRecommendationsExecute(r)
}

/*
GetMovieRecommendations Get recommended movies

Returns list of recommended movies based on provided movie ID in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param movieId
 @return ApiGetMovieRecommendationsRequest
*/
func (a *MoviesAPIService) GetMovieRecommendations(ctx context.Context, movieId float32) ApiGetMovieRecommendationsRequest {
	return ApiGetMovieRecommendationsRequest{
		ApiService: a,
		ctx: ctx,
		movieId: movieId,
	}
}

// Execute executes the request
//  @return GetDiscoverMovies200Response
func (a *MoviesAPIService) GetMovieRecommendationsExecute(r ApiGetMovieRecommendationsRequest) (*GetDiscoverMovies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoverMovies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/recommendations"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterToString(r.movieId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.language != nil {
		localVarQueryParams.Add("language", parameterToString(*r.language, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type ApiGetMovieSimilarRequest struct {
	ctx context.Context
	ApiService *MoviesAPIService
	movieId float32
	page *float32
	language *string
}

func (r ApiGetMovieSimilarRequest) Page(page float32) ApiGetMovieSimilarRequest {
	r.page = &page
	return r
}

func (r ApiGetMovieSimilarRequest) Language(language string) ApiGetMovieSimilarRequest {
	r.language = &language
	return r
}

func (r ApiGetMovieSimilarRequest) Execute() (*GetDiscoverMovies200Response, *http.Response, error) {
	return r.ApiService.GetMovieSimilarExecute(r)
}

/*
GetMovieSimilar Get similar movies

Returns list of similar movies based on the provided movieId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param movieId
 @return ApiGetMovieSimilarRequest
*/
func (a *MoviesAPIService) GetMovieSimilar(ctx context.Context, movieId float32) ApiGetMovieSimilarRequest {
	return ApiGetMovieSimilarRequest{
		ApiService: a,
		ctx: ctx,
		movieId: movieId,
	}
}

// Execute executes the request
//  @return GetDiscoverMovies200Response
func (a *MoviesAPIService) GetMovieSimilarExecute(r ApiGetMovieSimilarRequest) (*GetDiscoverMovies200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoverMovies200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieSimilar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/similar"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterToString(r.movieId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.language != nil {
		localVarQueryParams.Add("language", parameterToString(*r.language, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
