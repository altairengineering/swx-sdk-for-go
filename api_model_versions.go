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

// ModelVersionsApiService ModelVersionsApi service
type ModelVersionsApiService service

type ApiAddVersionRequest struct {
	ctx _context.Context
	ApiService *ModelVersionsApiService
	space string
	collectionName string
	modelName string
	modelVersionRequest *ModelVersionRequest
}

func (r ApiAddVersionRequest) ModelVersionRequest(modelVersionRequest ModelVersionRequest) ApiAddVersionRequest {
	r.modelVersionRequest = &modelVersionRequest
	return r
}

func (r ApiAddVersionRequest) Execute() (ModelVersionResponse, *_nethttp.Response, error) {
	return r.ApiService.AddVersionExecute(r)
}

/*
 * AddVersion Create version
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param collectionName
 * @param modelName
 * @return ApiAddVersionRequest
 */
func (a *ModelVersionsApiService) AddVersion(ctx _context.Context, space string, collectionName string, modelName string) ApiAddVersionRequest {
	return ApiAddVersionRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		collectionName: collectionName,
		modelName: modelName,
	}
}

/*
 * Execute executes the request
 * @return ModelVersionResponse
 */
func (a *ModelVersionsApiService) AddVersionExecute(r ApiAddVersionRequest) (ModelVersionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModelVersionsApiService.AddVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/collections/{collection-name}/models/{model-name}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection-name"+"}", _neturl.PathEscape(parameterToString(r.collectionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model-name"+"}", _neturl.PathEscape(parameterToString(r.modelName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.modelVersionRequest == nil {
		return localVarReturnValue, nil, reportError("modelVersionRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.modelVersionRequest
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
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiDeleteVersionRequest struct {
	ctx _context.Context
	ApiService *ModelVersionsApiService
	space string
	collectionName string
	thingId string
	modelName string
	versionName string
}


func (r ApiDeleteVersionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteVersionExecute(r)
}

/*
 * DeleteVersion Delete version
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param collectionName
 * @param thingId
 * @param modelName
 * @param versionName
 * @return ApiDeleteVersionRequest
 */
func (a *ModelVersionsApiService) DeleteVersion(ctx _context.Context, space string, collectionName string, thingId string, modelName string, versionName string) ApiDeleteVersionRequest {
	return ApiDeleteVersionRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		collectionName: collectionName,
		thingId: thingId,
		modelName: modelName,
		versionName: versionName,
	}
}

/*
 * Execute executes the request
 */
func (a *ModelVersionsApiService) DeleteVersionExecute(r ApiDeleteVersionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModelVersionsApiService.DeleteVersion")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/collections/{collection-name}/models/{model-name}/versions/{version-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection-name"+"}", _neturl.PathEscape(parameterToString(r.collectionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thing-id"+"}", _neturl.PathEscape(parameterToString(r.thingId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model-name"+"}", _neturl.PathEscape(parameterToString(r.modelName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version-name"+"}", _neturl.PathEscape(parameterToString(r.versionName, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListVersionRequest struct {
	ctx _context.Context
	ApiService *ModelVersionsApiService
	space string
	collectionName string
	modelName string
}


func (r ApiListVersionRequest) Execute() (ModelVersionListResponse, *_nethttp.Response, error) {
	return r.ApiService.ListVersionExecute(r)
}

/*
 * ListVersion List version
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param collectionName
 * @param modelName
 * @return ApiListVersionRequest
 */
func (a *ModelVersionsApiService) ListVersion(ctx _context.Context, space string, collectionName string, modelName string) ApiListVersionRequest {
	return ApiListVersionRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		collectionName: collectionName,
		modelName: modelName,
	}
}

/*
 * Execute executes the request
 * @return ModelVersionListResponse
 */
func (a *ModelVersionsApiService) ListVersionExecute(r ApiListVersionRequest) (ModelVersionListResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelVersionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModelVersionsApiService.ListVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/collections/{collection-name}/models/{model-name}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection-name"+"}", _neturl.PathEscape(parameterToString(r.collectionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model-name"+"}", _neturl.PathEscape(parameterToString(r.modelName, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

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
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiShowVersionRequest struct {
	ctx _context.Context
	ApiService *ModelVersionsApiService
	space string
	collectionName string
	thingId string
	modelName string
	versionName string
}


func (r ApiShowVersionRequest) Execute() (ModelVersionResponse, *_nethttp.Response, error) {
	return r.ApiService.ShowVersionExecute(r)
}

/*
 * ShowVersion Show version
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param collectionName
 * @param thingId
 * @param modelName
 * @param versionName
 * @return ApiShowVersionRequest
 */
func (a *ModelVersionsApiService) ShowVersion(ctx _context.Context, space string, collectionName string, thingId string, modelName string, versionName string) ApiShowVersionRequest {
	return ApiShowVersionRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		collectionName: collectionName,
		thingId: thingId,
		modelName: modelName,
		versionName: versionName,
	}
}

/*
 * Execute executes the request
 * @return ModelVersionResponse
 */
func (a *ModelVersionsApiService) ShowVersionExecute(r ApiShowVersionRequest) (ModelVersionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModelVersionsApiService.ShowVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/collections/{collection-name}/models/{model-name}/versions/{version-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection-name"+"}", _neturl.PathEscape(parameterToString(r.collectionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thing-id"+"}", _neturl.PathEscape(parameterToString(r.thingId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model-name"+"}", _neturl.PathEscape(parameterToString(r.modelName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version-name"+"}", _neturl.PathEscape(parameterToString(r.versionName, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

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
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiUpdateVersionRequest struct {
	ctx _context.Context
	ApiService *ModelVersionsApiService
	space string
	collectionName string
	thingId string
	modelName string
	versionName string
	modelVersionResponse *ModelVersionResponse
}

func (r ApiUpdateVersionRequest) ModelVersionResponse(modelVersionResponse ModelVersionResponse) ApiUpdateVersionRequest {
	r.modelVersionResponse = &modelVersionResponse
	return r
}

func (r ApiUpdateVersionRequest) Execute() (ModelVersionResponse, *_nethttp.Response, error) {
	return r.ApiService.UpdateVersionExecute(r)
}

/*
 * UpdateVersion Update version
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param space
 * @param collectionName
 * @param thingId
 * @param modelName
 * @param versionName
 * @return ApiUpdateVersionRequest
 */
func (a *ModelVersionsApiService) UpdateVersion(ctx _context.Context, space string, collectionName string, thingId string, modelName string, versionName string) ApiUpdateVersionRequest {
	return ApiUpdateVersionRequest{
		ApiService: a,
		ctx: ctx,
		space: space,
		collectionName: collectionName,
		thingId: thingId,
		modelName: modelName,
		versionName: versionName,
	}
}

/*
 * Execute executes the request
 * @return ModelVersionResponse
 */
func (a *ModelVersionsApiService) UpdateVersionExecute(r ApiUpdateVersionRequest) (ModelVersionResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ModelVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ModelVersionsApiService.UpdateVersion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/spaces/{space}/collections/{collection-name}/models/{model-name}/versions/{version-name}"
	localVarPath = strings.Replace(localVarPath, "{"+"space"+"}", _neturl.PathEscape(parameterToString(r.space, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection-name"+"}", _neturl.PathEscape(parameterToString(r.collectionName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"thing-id"+"}", _neturl.PathEscape(parameterToString(r.thingId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model-name"+"}", _neturl.PathEscape(parameterToString(r.modelName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version-name"+"}", _neturl.PathEscape(parameterToString(r.versionName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.modelVersionResponse == nil {
		return localVarReturnValue, nil, reportError("modelVersionResponse is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.modelVersionResponse
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
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BaseError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
