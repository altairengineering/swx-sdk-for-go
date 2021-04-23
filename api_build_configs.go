/*
 * Accounts & Users Service - Public API
 *
 * IN PROGRESS->This is the guide to use the different endpoints to manage the clusters.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// BuildConfigsApiService BuildConfigsApi service
type BuildConfigsApiService service

type ApiBuildConfCreateRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	modelsBuildConfigRequest *ModelsBuildConfigRequest
}

func (r ApiBuildConfCreateRequest) ModelsBuildConfigRequest(modelsBuildConfigRequest ModelsBuildConfigRequest) ApiBuildConfCreateRequest {
	r.modelsBuildConfigRequest = &modelsBuildConfigRequest
	return r
}

func (r ApiBuildConfCreateRequest) Execute() (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	return r.ApiService.BuildConfCreateExecute(r)
}

/*
 * BuildConfCreate Create Build Configuration
 * Create Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @return ApiBuildConfCreateRequest
 */
func (a *BuildConfigsApiService) BuildConfCreate(ctx _context.Context, space string) ApiBuildConfCreateRequest {
	return ApiBuildConfCreateRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
	}
}

/*
 * Execute executes the request
 * @return ModelsBuildConfigResponse
 */
func (a *BuildConfigsApiService) BuildConfCreateExecute(r ApiBuildConfCreateRequest) (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsBuildConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.modelsBuildConfigRequest == nil {
		return localVarReturnValue, nil, reportError("modelsBuildConfigRequest is required and must be specified")
	}

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
	localVarPostBody = r.modelsBuildConfigRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildConfDeleteRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	buildConfigID string
}


func (r ApiBuildConfDeleteRequest) Execute() (ModelsResourcesDeleteResponse, *_nethttp.Response, error) {
	return r.ApiService.BuildConfDeleteExecute(r)
}

/*
 * BuildConfDelete Delete Build Configuration
 * Delete Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @param buildConfigID Build Configuration ID
 * @return ApiBuildConfDeleteRequest
 */
func (a *BuildConfigsApiService) BuildConfDelete(ctx _context.Context, space string, buildConfigID string) ApiBuildConfDeleteRequest {
	return ApiBuildConfDeleteRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		buildConfigID: buildConfigID,
	}
}

/*
 * Execute executes the request
 * @return ModelsResourcesDeleteResponse
 */
func (a *BuildConfigsApiService) BuildConfDeleteExecute(r ApiBuildConfDeleteRequest) (ModelsResourcesDeleteResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsResourcesDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfDelete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/{build-configID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build-configID"+"}", _neturl.PathEscape(parameterToString(r.buildConfigID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildConfGetRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	buildConfigID string
}


func (r ApiBuildConfGetRequest) Execute() (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	return r.ApiService.BuildConfGetExecute(r)
}

/*
 * BuildConfGet Get Build Configuration
 * Get Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @param buildConfigID Build Configuration ID
 * @return ApiBuildConfGetRequest
 */
func (a *BuildConfigsApiService) BuildConfGet(ctx _context.Context, space string, buildConfigID string) ApiBuildConfGetRequest {
	return ApiBuildConfGetRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		buildConfigID: buildConfigID,
	}
}

/*
 * Execute executes the request
 * @return ModelsBuildConfigResponse
 */
func (a *BuildConfigsApiService) BuildConfGetExecute(r ApiBuildConfGetRequest) (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsBuildConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/{build-configID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build-configID"+"}", _neturl.PathEscape(parameterToString(r.buildConfigID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildConfListRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
}


func (r ApiBuildConfListRequest) Execute() (DataPagingBuildConfigs, *_nethttp.Response, error) {
	return r.ApiService.BuildConfListExecute(r)
}

/*
 * BuildConfList List Build Configuration
 * List Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @return ApiBuildConfListRequest
 */
func (a *BuildConfigsApiService) BuildConfList(ctx _context.Context, space string) ApiBuildConfListRequest {
	return ApiBuildConfListRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
	}
}

/*
 * Execute executes the request
 * @return DataPagingBuildConfigs
 */
func (a *BuildConfigsApiService) BuildConfListExecute(r ApiBuildConfListRequest) (DataPagingBuildConfigs, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DataPagingBuildConfigs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildConfUpdateRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	buildConfigID string
	modelsBuildConfigRequest *ModelsBuildConfigRequest
}

func (r ApiBuildConfUpdateRequest) ModelsBuildConfigRequest(modelsBuildConfigRequest ModelsBuildConfigRequest) ApiBuildConfUpdateRequest {
	r.modelsBuildConfigRequest = &modelsBuildConfigRequest
	return r
}

func (r ApiBuildConfUpdateRequest) Execute() (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	return r.ApiService.BuildConfUpdateExecute(r)
}

/*
 * BuildConfUpdate Update Build Configuration
 * Update Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @param buildConfigID Build Configuration ID
 * @return ApiBuildConfUpdateRequest
 */
func (a *BuildConfigsApiService) BuildConfUpdate(ctx _context.Context, space string, buildConfigID string) ApiBuildConfUpdateRequest {
	return ApiBuildConfUpdateRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		buildConfigID: buildConfigID,
	}
}

/*
 * Execute executes the request
 * @return ModelsBuildConfigResponse
 */
func (a *BuildConfigsApiService) BuildConfUpdateExecute(r ApiBuildConfUpdateRequest) (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsBuildConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/{build-configID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build-configID"+"}", _neturl.PathEscape(parameterToString(r.buildConfigID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.modelsBuildConfigRequest == nil {
		return localVarReturnValue, nil, reportError("modelsBuildConfigRequest is required and must be specified")
	}

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
	localVarPostBody = r.modelsBuildConfigRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildConfUpdateParcialyRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	buildConfigID string
	modelsBuildConfigRequest *ModelsBuildConfigRequest
}

func (r ApiBuildConfUpdateParcialyRequest) ModelsBuildConfigRequest(modelsBuildConfigRequest ModelsBuildConfigRequest) ApiBuildConfUpdateParcialyRequest {
	r.modelsBuildConfigRequest = &modelsBuildConfigRequest
	return r
}

func (r ApiBuildConfUpdateParcialyRequest) Execute() (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	return r.ApiService.BuildConfUpdateParcialyExecute(r)
}

/*
 * BuildConfUpdateParcialy Update Build Configuration
 * Update Build Configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @param buildConfigID Build Configuration ID
 * @return ApiBuildConfUpdateParcialyRequest
 */
func (a *BuildConfigsApiService) BuildConfUpdateParcialy(ctx _context.Context, space string, buildConfigID string) ApiBuildConfUpdateParcialyRequest {
	return ApiBuildConfUpdateParcialyRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		buildConfigID: buildConfigID,
	}
}

/*
 * Execute executes the request
 * @return ModelsBuildConfigResponse
 */
func (a *BuildConfigsApiService) BuildConfUpdateParcialyExecute(r ApiBuildConfUpdateParcialyRequest) (ModelsBuildConfigResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsBuildConfigResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildConfUpdateParcialy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/{build-configID}/"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build-configID"+"}", _neturl.PathEscape(parameterToString(r.buildConfigID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.modelsBuildConfigRequest == nil {
		return localVarReturnValue, nil, reportError("modelsBuildConfigRequest is required and must be specified")
	}

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
	localVarPostBody = r.modelsBuildConfigRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBuildGenerateFileRequest struct {
	ctx _context.Context
	ApiService *BuildConfigsApiService
	space string
	buildConfigID string
}


func (r ApiBuildGenerateFileRequest) Execute() (ModelsBuildConfigGenerateFile, *_nethttp.Response, error) {
	return r.ApiService.BuildGenerateFileExecute(r)
}

/*
 * BuildGenerateFile Create a temporary endpoint with the Build Configuration file
 * Get generate file for a specific build configuration
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space Space ID
 * @param buildConfigID Build Configuration ID
 * @return ApiBuildGenerateFileRequest
 */
func (a *BuildConfigsApiService) BuildGenerateFile(ctx _context.Context, space string, buildConfigID string) ApiBuildGenerateFileRequest {
	return ApiBuildGenerateFileRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		buildConfigID: buildConfigID,
	}
}

/*
 * Execute executes the request
 * @return ModelsBuildConfigGenerateFile
 */
func (a *BuildConfigsApiService) BuildGenerateFileExecute(r ApiBuildGenerateFileRequest) (ModelsBuildConfigGenerateFile, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelsBuildConfigGenerateFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BuildConfigsApiService.BuildGenerateFile")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/build-configs/{build-configID}/generate-file"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build-configID"+"}", _neturl.PathEscape(parameterToString(r.buildConfigID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
