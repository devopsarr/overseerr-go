/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: v1.33.2
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


// TvAPIService TvAPI service
type TvAPIService service

type ApiGetTvByTvIdRequest struct {
	ctx context.Context
	ApiService *TvAPIService
	tvId float32
	language *string
}

func (r ApiGetTvByTvIdRequest) Language(language string) ApiGetTvByTvIdRequest {
	r.language = &language
	return r
}

func (r ApiGetTvByTvIdRequest) Execute() (*TvDetails, *http.Response, error) {
	return r.ApiService.GetTvByTvIdExecute(r)
}

/*
GetTvByTvId Get TV details

Returns full TV details in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tvId
 @return ApiGetTvByTvIdRequest
*/
func (a *TvAPIService) GetTvByTvId(ctx context.Context, tvId float32) ApiGetTvByTvIdRequest {
	return ApiGetTvByTvIdRequest{
		ApiService: a,
		ctx: ctx,
		tvId: tvId,
	}
}

// Execute executes the request
//  @return TvDetails
func (a *TvAPIService) GetTvByTvIdExecute(r ApiGetTvByTvIdRequest) (*TvDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TvDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvAPIService.GetTvByTvId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tv/{tvId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tvId"+"}", url.PathEscape(parameterValueToString(r.tvId, "tvId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
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

type ApiGetTvRatingsRequest struct {
	ctx context.Context
	ApiService *TvAPIService
	tvId float32
}

func (r ApiGetTvRatingsRequest) Execute() (*GetTvRatings2XXResponse, *http.Response, error) {
	return r.ApiService.GetTvRatingsExecute(r)
}

/*
GetTvRatings Get TV ratings

Returns ratings based on provided tvId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tvId
 @return ApiGetTvRatingsRequest
*/
func (a *TvAPIService) GetTvRatings(ctx context.Context, tvId float32) ApiGetTvRatingsRequest {
	return ApiGetTvRatingsRequest{
		ApiService: a,
		ctx: ctx,
		tvId: tvId,
	}
}

// Execute executes the request
//  @return GetTvRatings2XXResponse
func (a *TvAPIService) GetTvRatingsExecute(r ApiGetTvRatingsRequest) (*GetTvRatings2XXResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTvRatings2XXResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvAPIService.GetTvRatings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tv/{tvId}/ratings"
	localVarPath = strings.Replace(localVarPath, "{"+"tvId"+"}", url.PathEscape(parameterValueToString(r.tvId, "tvId")), -1)

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

type ApiGetTvRecommendationsRequest struct {
	ctx context.Context
	ApiService *TvAPIService
	tvId float32
	page *float32
	language *string
}

func (r ApiGetTvRecommendationsRequest) Page(page float32) ApiGetTvRecommendationsRequest {
	r.page = &page
	return r
}

func (r ApiGetTvRecommendationsRequest) Language(language string) ApiGetTvRecommendationsRequest {
	r.language = &language
	return r
}

func (r ApiGetTvRecommendationsRequest) Execute() (*GetDiscoverTv2XXResponse, *http.Response, error) {
	return r.ApiService.GetTvRecommendationsExecute(r)
}

/*
GetTvRecommendations Get recommended TV series

Returns list of recommended TV series based on the provided tvId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tvId
 @return ApiGetTvRecommendationsRequest
*/
func (a *TvAPIService) GetTvRecommendations(ctx context.Context, tvId float32) ApiGetTvRecommendationsRequest {
	return ApiGetTvRecommendationsRequest{
		ApiService: a,
		ctx: ctx,
		tvId: tvId,
	}
}

// Execute executes the request
//  @return GetDiscoverTv2XXResponse
func (a *TvAPIService) GetTvRecommendationsExecute(r ApiGetTvRecommendationsRequest) (*GetDiscoverTv2XXResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoverTv2XXResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvAPIService.GetTvRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tv/{tvId}/recommendations"
	localVarPath = strings.Replace(localVarPath, "{"+"tvId"+"}", url.PathEscape(parameterValueToString(r.tvId, "tvId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue float32 = 1
		r.page = &defaultValue
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
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

type ApiGetTvSeasonBySeasonIdRequest struct {
	ctx context.Context
	ApiService *TvAPIService
	tvId float32
	seasonId float32
	language *string
}

func (r ApiGetTvSeasonBySeasonIdRequest) Language(language string) ApiGetTvSeasonBySeasonIdRequest {
	r.language = &language
	return r
}

func (r ApiGetTvSeasonBySeasonIdRequest) Execute() (*Season, *http.Response, error) {
	return r.ApiService.GetTvSeasonBySeasonIdExecute(r)
}

/*
GetTvSeasonBySeasonId Get season details and episode list

Returns season details with a list of episodes in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tvId
 @param seasonId
 @return ApiGetTvSeasonBySeasonIdRequest
*/
func (a *TvAPIService) GetTvSeasonBySeasonId(ctx context.Context, tvId float32, seasonId float32) ApiGetTvSeasonBySeasonIdRequest {
	return ApiGetTvSeasonBySeasonIdRequest{
		ApiService: a,
		ctx: ctx,
		tvId: tvId,
		seasonId: seasonId,
	}
}

// Execute executes the request
//  @return Season
func (a *TvAPIService) GetTvSeasonBySeasonIdExecute(r ApiGetTvSeasonBySeasonIdRequest) (*Season, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Season
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvAPIService.GetTvSeasonBySeasonId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tv/{tvId}/season/{seasonId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tvId"+"}", url.PathEscape(parameterValueToString(r.tvId, "tvId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"seasonId"+"}", url.PathEscape(parameterValueToString(r.seasonId, "seasonId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
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

type ApiGetTvSimilarRequest struct {
	ctx context.Context
	ApiService *TvAPIService
	tvId float32
	page *float32
	language *string
}

func (r ApiGetTvSimilarRequest) Page(page float32) ApiGetTvSimilarRequest {
	r.page = &page
	return r
}

func (r ApiGetTvSimilarRequest) Language(language string) ApiGetTvSimilarRequest {
	r.language = &language
	return r
}

func (r ApiGetTvSimilarRequest) Execute() (*GetDiscoverTv2XXResponse, *http.Response, error) {
	return r.ApiService.GetTvSimilarExecute(r)
}

/*
GetTvSimilar Get similar TV series

Returns list of similar TV series based on the provided tvId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tvId
 @return ApiGetTvSimilarRequest
*/
func (a *TvAPIService) GetTvSimilar(ctx context.Context, tvId float32) ApiGetTvSimilarRequest {
	return ApiGetTvSimilarRequest{
		ApiService: a,
		ctx: ctx,
		tvId: tvId,
	}
}

// Execute executes the request
//  @return GetDiscoverTv2XXResponse
func (a *TvAPIService) GetTvSimilarExecute(r ApiGetTvSimilarRequest) (*GetDiscoverTv2XXResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoverTv2XXResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvAPIService.GetTvSimilar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tv/{tvId}/similar"
	localVarPath = strings.Replace(localVarPath, "{"+"tvId"+"}", url.PathEscape(parameterValueToString(r.tvId, "tvId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue float32 = 1
		r.page = &defaultValue
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
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
