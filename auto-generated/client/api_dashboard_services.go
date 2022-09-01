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

// DashboardServicesApiService DashboardServicesApi service
type DashboardServicesApiService service

type ApiCreateDashboardServiceRequest struct {
	ctx                    context.Context
	ApiService             *DashboardServicesApiService
	createDashboardService *CreateDashboardService
}

func (r ApiCreateDashboardServiceRequest) CreateDashboardService(createDashboardService CreateDashboardService) ApiCreateDashboardServiceRequest {
	r.createDashboardService = &createDashboardService
	return r
}

func (r ApiCreateDashboardServiceRequest) Execute() (*DashboardService, *http.Response, error) {
	return r.ApiService.CreateDashboardServiceExecute(r)
}

/*
CreateDashboardService Create a dashboard service

Create a new dashboard service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDashboardServiceRequest
*/
func (a *DashboardServicesApiService) CreateDashboardService(ctx context.Context) ApiCreateDashboardServiceRequest {
	return ApiCreateDashboardServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DashboardService
func (a *DashboardServicesApiService) CreateDashboardServiceExecute(r ApiCreateDashboardServiceRequest) (*DashboardService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.CreateDashboardService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices"

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
	localVarPostBody = r.createDashboardService
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

type ApiCreateOrUpdateDashboardServiceRequest struct {
	ctx                    context.Context
	ApiService             *DashboardServicesApiService
	createDashboardService *CreateDashboardService
}

func (r ApiCreateOrUpdateDashboardServiceRequest) CreateDashboardService(createDashboardService CreateDashboardService) ApiCreateOrUpdateDashboardServiceRequest {
	r.createDashboardService = &createDashboardService
	return r
}

func (r ApiCreateOrUpdateDashboardServiceRequest) Execute() (*DashboardService, *http.Response, error) {
	return r.ApiService.CreateOrUpdateDashboardServiceExecute(r)
}

/*
CreateOrUpdateDashboardService Update a Dashboard service

Update an existing dashboard service identified by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrUpdateDashboardServiceRequest
*/
func (a *DashboardServicesApiService) CreateOrUpdateDashboardService(ctx context.Context) ApiCreateOrUpdateDashboardServiceRequest {
	return ApiCreateOrUpdateDashboardServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DashboardService
func (a *DashboardServicesApiService) CreateOrUpdateDashboardServiceExecute(r ApiCreateOrUpdateDashboardServiceRequest) (*DashboardService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.CreateOrUpdateDashboardService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices"

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
	localVarPostBody = r.createDashboardService
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

type ApiDeleteDashboardServiceRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	id         string
	recursive  *bool
	hardDelete *bool
}

// Recursively delete this entity and it&#39;s children. (Default &#x60;false&#x60;)
func (r ApiDeleteDashboardServiceRequest) Recursive(recursive bool) ApiDeleteDashboardServiceRequest {
	r.recursive = &recursive
	return r
}

// Hard delete the entity. (Default &#x3D; &#x60;false&#x60;)
func (r ApiDeleteDashboardServiceRequest) HardDelete(hardDelete bool) ApiDeleteDashboardServiceRequest {
	r.hardDelete = &hardDelete
	return r
}

func (r ApiDeleteDashboardServiceRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDashboardServiceExecute(r)
}

/*
DeleteDashboardService Delete a Dashboard service

Delete a Dashboard services. If dashboard (and charts) belong to the service, it can't be deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the dashboard service
 @return ApiDeleteDashboardServiceRequest
*/
func (a *DashboardServicesApiService) DeleteDashboardService(ctx context.Context, id string) ApiDeleteDashboardServiceRequest {
	return ApiDeleteDashboardServiceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DashboardServicesApiService) DeleteDashboardServiceExecute(r ApiDeleteDashboardServiceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.DeleteDashboardService")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recursive != nil {
		localVarQueryParams.Add("recursive", parameterToString(*r.recursive, ""))
	}
	if r.hardDelete != nil {
		localVarQueryParams.Add("hardDelete", parameterToString(*r.hardDelete, ""))
	}
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

type ApiGetDashboardServiceByFQNRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	name       string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDashboardServiceByFQNRequest) Fields(fields string) ApiGetDashboardServiceByFQNRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDashboardServiceByFQNRequest) Include(include string) ApiGetDashboardServiceByFQNRequest {
	r.include = &include
	return r
}

func (r ApiGetDashboardServiceByFQNRequest) Execute() (*DashboardService, *http.Response, error) {
	return r.ApiService.GetDashboardServiceByFQNExecute(r)
}

/*
GetDashboardServiceByFQN Get dashboard service by name

Get a dashboard service by the service `name`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name
 @return ApiGetDashboardServiceByFQNRequest
*/
func (a *DashboardServicesApiService) GetDashboardServiceByFQN(ctx context.Context, name string) ApiGetDashboardServiceByFQNRequest {
	return ApiGetDashboardServiceByFQNRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//  @return DashboardService
func (a *DashboardServicesApiService) GetDashboardServiceByFQNExecute(r ApiGetDashboardServiceByFQNRequest) (*DashboardService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.GetDashboardServiceByFQN")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices/name/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
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

type ApiGetDashboardServiceByIDRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	id         string
	fields     *string
	include    *string
}

// Fields requested in the returned resource
func (r ApiGetDashboardServiceByIDRequest) Fields(fields string) ApiGetDashboardServiceByIDRequest {
	r.fields = &fields
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiGetDashboardServiceByIDRequest) Include(include string) ApiGetDashboardServiceByIDRequest {
	r.include = &include
	return r
}

func (r ApiGetDashboardServiceByIDRequest) Execute() (*DashboardService, *http.Response, error) {
	return r.ApiService.GetDashboardServiceByIDExecute(r)
}

/*
GetDashboardServiceByID Get a dashboard service

Get a dashboard service by `id`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiGetDashboardServiceByIDRequest
*/
func (a *DashboardServicesApiService) GetDashboardServiceByID(ctx context.Context, id string) ApiGetDashboardServiceByIDRequest {
	return ApiGetDashboardServiceByIDRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return DashboardService
func (a *DashboardServicesApiService) GetDashboardServiceByIDExecute(r ApiGetDashboardServiceByIDRequest) (*DashboardService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.GetDashboardServiceByID")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
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

type ApiGetSpecificDashboardServiceVersionRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	id         string
	version    string
}

func (r ApiGetSpecificDashboardServiceVersionRequest) Execute() (*DashboardService, *http.Response, error) {
	return r.ApiService.GetSpecificDashboardServiceVersionExecute(r)
}

/*
GetSpecificDashboardServiceVersion Get a version of the dashboard service

Get a version of the dashboard service by given `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id dashboard service Id
 @param version dashboard service version number in the form `major`.`minor`
 @return ApiGetSpecificDashboardServiceVersionRequest
*/
func (a *DashboardServicesApiService) GetSpecificDashboardServiceVersion(ctx context.Context, id string, version string) ApiGetSpecificDashboardServiceVersionRequest {
	return ApiGetSpecificDashboardServiceVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		version:    version,
	}
}

// Execute executes the request
//  @return DashboardService
func (a *DashboardServicesApiService) GetSpecificDashboardServiceVersionExecute(r ApiGetSpecificDashboardServiceVersionRequest) (*DashboardService, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardService
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.GetSpecificDashboardServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices/{id}/versions/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

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

type ApiListAllDashboardServiceVersionRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	id         string
}

func (r ApiListAllDashboardServiceVersionRequest) Execute() (*EntityHistory, *http.Response, error) {
	return r.ApiService.ListAllDashboardServiceVersionExecute(r)
}

/*
ListAllDashboardServiceVersion List dashboard service versions

Get a list of all the versions of a dashboard service identified by `id`

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id dashboard service Id
 @return ApiListAllDashboardServiceVersionRequest
*/
func (a *DashboardServicesApiService) ListAllDashboardServiceVersion(ctx context.Context, id string) ApiListAllDashboardServiceVersionRequest {
	return ApiListAllDashboardServiceVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return EntityHistory
func (a *DashboardServicesApiService) ListAllDashboardServiceVersionExecute(r ApiListAllDashboardServiceVersionRequest) (*EntityHistory, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.ListAllDashboardServiceVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices/{id}/versions"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiListDashboardsServiceRequest struct {
	ctx        context.Context
	ApiService *DashboardServicesApiService
	name       *string
	fields     *string
	limit      *int32
	before     *string
	after      *string
	include    *string
}

func (r ApiListDashboardsServiceRequest) Name(name string) ApiListDashboardsServiceRequest {
	r.name = &name
	return r
}

// Fields requested in the returned resource
func (r ApiListDashboardsServiceRequest) Fields(fields string) ApiListDashboardsServiceRequest {
	r.fields = &fields
	return r
}

func (r ApiListDashboardsServiceRequest) Limit(limit int32) ApiListDashboardsServiceRequest {
	r.limit = &limit
	return r
}

// Returns list of dashboard services before this cursor
func (r ApiListDashboardsServiceRequest) Before(before string) ApiListDashboardsServiceRequest {
	r.before = &before
	return r
}

// Returns list of dashboard services after this cursor
func (r ApiListDashboardsServiceRequest) After(after string) ApiListDashboardsServiceRequest {
	r.after = &after
	return r
}

// Include all, deleted, or non-deleted entities.
func (r ApiListDashboardsServiceRequest) Include(include string) ApiListDashboardsServiceRequest {
	r.include = &include
	return r
}

func (r ApiListDashboardsServiceRequest) Execute() (*DashboardServiceList, *http.Response, error) {
	return r.ApiService.ListDashboardsServiceExecute(r)
}

/*
ListDashboardsService List dashboard services

Get a list of dashboard services.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDashboardsServiceRequest
*/
func (a *DashboardServicesApiService) ListDashboardsService(ctx context.Context) ApiListDashboardsServiceRequest {
	return ApiListDashboardsServiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return DashboardServiceList
func (a *DashboardServicesApiService) ListDashboardsServiceExecute(r ApiListDashboardsServiceRequest) (*DashboardServiceList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DashboardServiceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DashboardServicesApiService.ListDashboardsService")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/services/dashboardServices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.before != nil {
		localVarQueryParams.Add("before", parameterToString(*r.before, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
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
