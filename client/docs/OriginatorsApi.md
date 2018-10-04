# \OriginatorsApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOriginator**](OriginatorsApi.md#AddOriginator) | **Post** /v1/ach/originators | Create a new Originator object
[**DeleteOriginator**](OriginatorsApi.md#DeleteOriginator) | **Delete** /v1/ach/originators/{originatorId} | Permanently deletes an Originator and associated Customers, Depositories, and Transfers. It cannot be undone. Also immediately cancels any active Transfers for the Originator.
[**GetOriginatorByID**](OriginatorsApi.md#GetOriginatorByID) | **Get** /v1/ach/originators/{originatorId} | Retrieves the details of an existing Originator. You need only supply the unique Originator identifier that was returned upon customer creation.
[**GetOriginators**](OriginatorsApi.md#GetOriginators) | **Get** /v1/ach/originators | Gets a list of Originators
[**UpdateOriginator**](OriginatorsApi.md#UpdateOriginator) | **Patch** /v1/ach/originators/{originatorId} | Updates the specified Originator by setting the values of the parameters passed. Any parameters not provided will be left unchanged.


# **AddOriginator**
> Originator AddOriginator(ctx, originator, optional)
Create a new Originator object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **originator** | [**Originator**](Originator.md)| A JSON object containing a new Originator | 
 **optional** | ***AddOriginatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddOriginatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Originator**](Originator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOriginator**
> DeleteOriginator(ctx, originatorId, optional)
Permanently deletes an Originator and associated Customers, Depositories, and Transfers. It cannot be undone. Also immediately cancels any active Transfers for the Originator.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **originatorId** | **string**| Originator ID | 
 **optional** | ***DeleteOriginatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOriginatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOriginatorByID**
> Originator GetOriginatorByID(ctx, originatorId, optional)
Retrieves the details of an existing Originator. You need only supply the unique Originator identifier that was returned upon customer creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **originatorId** | **string**| Originator ID | 
 **optional** | ***GetOriginatorByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOriginatorByIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **expand** | **optional.String**| Return nested objects rather than ID&#39;s in the response body. | 

### Return type

[**Originator**](Originator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOriginators**
> Originators GetOriginators(ctx, optional)
Gets a list of Originators

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOriginatorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOriginatorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **expand** | **optional.String**| Return nested objects rather than ID&#39;s in the response body. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Originators**](Originators.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOriginator**
> Originator UpdateOriginator(ctx, originatorId, originator, optional)
Updates the specified Originator by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **originatorId** | **string**| Originator ID | 
  **originator** | [**Originator**](Originator.md)| A JSON object containing a new Originator | 
 **optional** | ***UpdateOriginatorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateOriginatorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Originator**](Originator.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

