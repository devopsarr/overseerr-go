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


// MediaApiService MediaApi service
type MediaApiService service
type ApiCreateMediaByStatusRequest struct {
	ctx context.Context
	ApiService *MediaApiService
	mediaId string
	status string
	createMediaByStatusRequest *CreateMediaByStatusRequest
}

func (r ApiCreateMediaByStatusRequest) CreateMediaByStatusRequest(createMediaByStatusRequest CreateMediaByStatusRequest) ApiCreateMediaByStatusRequest {
	r.createMediaByStatusRequest = &createMediaByStatusRequest
	return r
}

func (r ApiCreateMediaByStatusRequest) Execute() (*MediaInfo, *http.Response, error) {
	return r.ApiService.CreateMediaByStatusExecute(r)
}

/*
CreateMediaByStatus Update media status

Updates a media item's status and returns the media in JSON format

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @param status New status
 @return ApiCreateMediaByStatusRequest
*/
func (a *MediaApiService) CreateMediaByStatus(ctx context.Context, mediaId string, status string) ApiCreateMediaByStatusRequest {
	return ApiCreateMediaByStatusRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
		status: status,
	}
}

// Execute executes the request
//  @return MediaInfo
func (a *MediaApiService) CreateMediaByStatusExecute(r ApiCreateMediaByStatusRequest) (*MediaInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MediaInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaApiService.CreateMediaByStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}/{status}"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterToString(r.mediaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterToString(r.status, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createMediaByStatusRequest
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
type ApiDeleteMediaRequest struct {
	ctx context.Context
	ApiService *MediaApiService
	mediaId string
}

func (r ApiDeleteMediaRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMediaExecute(r)
}

/*
DeleteMedia Delete media item

Removes a media item. The `MANAGE_REQUESTS` permission is required to perform this action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @return ApiDeleteMediaRequest
*/
func (a *MediaApiService) DeleteMedia(ctx context.Context, mediaId string) ApiDeleteMediaRequest {
	return ApiDeleteMediaRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
	}
}

// Execute executes the request
func (a *MediaApiService) DeleteMediaExecute(r ApiDeleteMediaRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaApiService.DeleteMedia")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterToString(r.mediaId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
type ApiGetMediaRequest struct {
	ctx context.Context
	ApiService *MediaApiService
	take *float32
	skip *float32
	filter *string
	sort *string
}

func (r ApiGetMediaRequest) Take(take float32) ApiGetMediaRequest {
	r.take = &take
	return r
}

func (r ApiGetMediaRequest) Skip(skip float32) ApiGetMediaRequest {
	r.skip = &skip
	return r
}

func (r ApiGetMediaRequest) Filter(filter string) ApiGetMediaRequest {
	r.filter = &filter
	return r
}

func (r ApiGetMediaRequest) Sort(sort string) ApiGetMediaRequest {
	r.sort = &sort
	return r
}

func (r ApiGetMediaRequest) Execute() (*GetMedia200Response, *http.Response, error) {
	return r.ApiService.GetMediaExecute(r)
}

/*
GetMedia Get media

Returns all media (can be filtered and limited) in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMediaRequest
*/
func (a *MediaApiService) GetMedia(ctx context.Context) ApiGetMediaRequest {
	return ApiGetMediaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMedia200Response
func (a *MediaApiService) GetMediaExecute(r ApiGetMediaRequest) (*GetMedia200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMedia200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaApiService.GetMedia")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.take != nil {
		localVarQueryParams.Add("take", parameterToString(*r.take, ""))
	}
	if r.skip != nil {
		localVarQueryParams.Add("skip", parameterToString(*r.skip, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
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
type ApiGetMediaWatchDataRequest struct {
	ctx context.Context
	ApiService *MediaApiService
	mediaId string
}

func (r ApiGetMediaWatchDataRequest) Execute() (*GetMediaWatchData200Response, *http.Response, error) {
	return r.ApiService.GetMediaWatchDataExecute(r)
}

/*
GetMediaWatchData Get watch data

Returns play count, play duration, and users who have watched the media.

Requires the `ADMIN` permission.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @return ApiGetMediaWatchDataRequest
*/
func (a *MediaApiService) GetMediaWatchData(ctx context.Context, mediaId string) ApiGetMediaWatchDataRequest {
	return ApiGetMediaWatchDataRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
	}
}

// Execute executes the request
//  @return GetMediaWatchData200Response
func (a *MediaApiService) GetMediaWatchDataExecute(r ApiGetMediaWatchDataRequest) (*GetMediaWatchData200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMediaWatchData200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaApiService.GetMediaWatchData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}/watch_data"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterToString(r.mediaId, "")), -1)

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
