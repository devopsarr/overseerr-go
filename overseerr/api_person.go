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


// PersonApiService PersonApi service
type PersonApiService service
type ApiGetPersonByPersonIdRequest struct {
	ctx context.Context
	ApiService *PersonApiService
	personId float32
	language *string
}

func (r ApiGetPersonByPersonIdRequest) Language(language string) ApiGetPersonByPersonIdRequest {
	r.language = &language
	return r
}

func (r ApiGetPersonByPersonIdRequest) Execute() (*PersonDetails, *http.Response, error) {
	return r.ApiService.GetPersonByPersonIdExecute(r)
}

/*
GetPersonByPersonId Get person details

Returns person details based on provided personId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId
 @return ApiGetPersonByPersonIdRequest
*/
func (a *PersonApiService) GetPersonByPersonId(ctx context.Context, personId float32) ApiGetPersonByPersonIdRequest {
	return ApiGetPersonByPersonIdRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return PersonDetails
func (a *PersonApiService) GetPersonByPersonIdExecute(r ApiGetPersonByPersonIdRequest) (*PersonDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PersonDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonApiService.GetPersonByPersonId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/person/{personId}"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterToString(r.personId, "")), -1)

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
type ApiGetPersonCombinedCreditsRequest struct {
	ctx context.Context
	ApiService *PersonApiService
	personId float32
	language *string
}

func (r ApiGetPersonCombinedCreditsRequest) Language(language string) ApiGetPersonCombinedCreditsRequest {
	r.language = &language
	return r
}

func (r ApiGetPersonCombinedCreditsRequest) Execute() (*GetPersonCombinedCredits200Response, *http.Response, error) {
	return r.ApiService.GetPersonCombinedCreditsExecute(r)
}

/*
GetPersonCombinedCredits Get combined credits

Returns the person's combined credits based on the provided personId in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param personId
 @return ApiGetPersonCombinedCreditsRequest
*/
func (a *PersonApiService) GetPersonCombinedCredits(ctx context.Context, personId float32) ApiGetPersonCombinedCreditsRequest {
	return ApiGetPersonCombinedCreditsRequest{
		ApiService: a,
		ctx: ctx,
		personId: personId,
	}
}

// Execute executes the request
//  @return GetPersonCombinedCredits200Response
func (a *PersonApiService) GetPersonCombinedCreditsExecute(r ApiGetPersonCombinedCreditsRequest) (*GetPersonCombinedCredits200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPersonCombinedCredits200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PersonApiService.GetPersonCombinedCredits")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/person/{personId}/combined_credits"
	localVarPath = strings.Replace(localVarPath, "{"+"personId"+"}", url.PathEscape(parameterToString(r.personId, "")), -1)

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
