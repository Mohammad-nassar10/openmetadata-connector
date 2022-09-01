/*
OpenMetadata Apis

# Overview  OpenMetadata supports REST APIs for getting metadata in and out of metadata store. The API resources are grouped under following categories: - **Data assets** - includes resources for data entities, such as `databases`, `tables`, and `topics`. Resources for data assets created from data, such as `dashboards`, `reports`, `metrics`, and `ML Features` are part of this collection. `pipelines`, `notebooks`, etc. that are used for creating data assets are also available as resources as of this collection. - **Teams and Users** - includes `users`, `teams`, a special type of user called `bots` that performs many automated tasks such as ingestion. - **Services** - are services that OpenMetadata integrates with. Currently `databaseService` is the only service under this collection that represents data sources. In the future, services related to Dashboards, Reports, ETL pipelines will be added under this collection. - **Glossary** - OpenMetadata supports hierarchical tags that can be used to build business vocabulary to describe and classify data available under `tags` resource.

API version: 0.11.0
Contact: openmetadata-dev@googlegroups.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// LineageApiService LineageApi service
type LineageApiService service

type ApiAddLineageEdgeRequest struct {
	ctx        context.Context
	ApiService *LineageApiService
	addLineage *AddLineage
}

func (r ApiAddLineageEdgeRequest) AddLineage(addLineage AddLineage) ApiAddLineageEdgeRequest {
	r.addLineage = &addLineage
	return r
}

func (r ApiAddLineageEdgeRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddLineageEdgeExecute(r)
}

/*
AddLineageEdge Add a lineage edge

Add a lineage edge with from entity as upstream node and to entity as downstream node.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddLineageEdgeRequest
*/
func (a *LineageApiService) AddLineageEdge(ctx context.Context) ApiAddLineageEdgeRequest {
	return ApiAddLineageEdgeRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *LineageApiService) AddLineageEdgeExecute(r ApiAddLineageEdgeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineageApiService.AddLineageEdge")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lineage"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addLineage
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteLineageEdgeRequest struct {
	ctx        context.Context
	ApiService *LineageApiService
	fromEntity string
	fromId     string
	toEntity   string
	toId       string
}

func (r ApiDeleteLineageEdgeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLineageEdgeExecute(r)
}

/*
DeleteLineageEdge Delete a lineage edge

Delete a lineage edge with from entity as upstream node and to entity as downstream node.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param fromEntity Entity type of upstream entity of the edge
 @param fromId Entity id
 @param toEntity Entity type for downstream entity of the edge
 @param toId Entity id
 @return ApiDeleteLineageEdgeRequest
*/
func (a *LineageApiService) DeleteLineageEdge(ctx context.Context, fromEntity string, fromId string, toEntity string, toId string) ApiDeleteLineageEdgeRequest {
	return ApiDeleteLineageEdgeRequest{
		ApiService: a,
		ctx:        ctx,
		fromEntity: fromEntity,
		fromId:     fromId,
		toEntity:   toEntity,
		toId:       toId,
	}
}

// Execute executes the request
func (a *LineageApiService) DeleteLineageEdgeExecute(r ApiDeleteLineageEdgeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineageApiService.DeleteLineageEdge")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lineage/{fromEntity}/{fromId}/{toEntity}/{toId}"
	localVarPath = strings.Replace(localVarPath, "{"+"fromEntity"+"}", url.PathEscape(parameterToString(r.fromEntity, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fromId"+"}", url.PathEscape(parameterToString(r.fromId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toEntity"+"}", url.PathEscape(parameterToString(r.toEntity, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"toId"+"}", url.PathEscape(parameterToString(r.toId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetLineageRequest struct {
	ctx             context.Context
	ApiService      *LineageApiService
	entity          string
	id              string
	upstreamDepth   *int32
	downstreamDepth *int32
}

// Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3)
func (r ApiGetLineageRequest) UpstreamDepth(upstreamDepth int32) ApiGetLineageRequest {
	r.upstreamDepth = &upstreamDepth
	return r
}

// Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3)
func (r ApiGetLineageRequest) DownstreamDepth(downstreamDepth int32) ApiGetLineageRequest {
	r.downstreamDepth = &downstreamDepth
	return r
}

func (r ApiGetLineageRequest) Execute() (*EntityLineage, *http.Response, error) {
	return r.ApiService.GetLineageExecute(r)
}

/*
GetLineage Get lineage

Get lineage details for an entity identified by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entity Entity type for which lineage is requested
 @param id Entity id
 @return ApiGetLineageRequest
*/
func (a *LineageApiService) GetLineage(ctx context.Context, entity string, id string) ApiGetLineageRequest {
	return ApiGetLineageRequest{
		ApiService: a,
		ctx:        ctx,
		entity:     entity,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityLineage
func (a *LineageApiService) GetLineageExecute(r ApiGetLineageRequest) (*EntityLineage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityLineage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineageApiService.GetLineage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lineage/{entity}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity"+"}", url.PathEscape(parameterToString(r.entity, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.upstreamDepth != nil {
		localVarQueryParams.Add("upstreamDepth", parameterToString(*r.upstreamDepth, ""))
	}
	if r.downstreamDepth != nil {
		localVarQueryParams.Add("downstreamDepth", parameterToString(*r.downstreamDepth, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetLineageByFQNRequest struct {
	ctx             context.Context
	ApiService      *LineageApiService
	entity          string
	fqn             string
	upstreamDepth   *int32
	downstreamDepth *int32
}

// Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3)
func (r ApiGetLineageByFQNRequest) UpstreamDepth(upstreamDepth int32) ApiGetLineageByFQNRequest {
	r.upstreamDepth = &upstreamDepth
	return r
}

// Upstream depth of lineage (default&#x3D;1, min&#x3D;0, max&#x3D;3)
func (r ApiGetLineageByFQNRequest) DownstreamDepth(downstreamDepth int32) ApiGetLineageByFQNRequest {
	r.downstreamDepth = &downstreamDepth
	return r
}

func (r ApiGetLineageByFQNRequest) Execute() (*EntityLineage, *http.Response, error) {
	return r.ApiService.GetLineageByFQNExecute(r)
}

/*
GetLineageByFQN Get lineage by name

Get lineage details for an entity identified by fully qualified name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entity Entity type for which lineage is requested
 @param fqn Fully qualified name of the entity that uniquely identifies an entity
 @return ApiGetLineageByFQNRequest
*/
func (a *LineageApiService) GetLineageByFQN(ctx context.Context, entity string, fqn string) ApiGetLineageByFQNRequest {
	return ApiGetLineageByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		entity:     entity,
		fqn:        fqn,
	}
}

// Execute executes the request
//  @return EntityLineage
func (a *LineageApiService) GetLineageByFQNExecute(r ApiGetLineageByFQNRequest) (*EntityLineage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityLineage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LineageApiService.GetLineageByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lineage/{entity}/name/{fqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity"+"}", url.PathEscape(parameterToString(r.entity, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fqn"+"}", url.PathEscape(parameterToString(r.fqn, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.upstreamDepth != nil {
		localVarQueryParams.Add("upstreamDepth", parameterToString(*r.upstreamDepth, ""))
	}
	if r.downstreamDepth != nil {
		localVarQueryParams.Add("downstreamDepth", parameterToString(*r.downstreamDepth, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
