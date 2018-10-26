# \DepositoriesApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDepository**](DepositoriesApi.md#AddDepository) | **Post** /v1/ach/depositories | Create a new depository account for a Customer ID or Originator ID defined in the Parent parameter
[**CreateMicroDeposits**](DepositoriesApi.md#CreateMicroDeposits) | **Post** /v1/ach/depositories/{depositoryId}/micro-deposits | Initiates micro deposits to be sent to the Depository institution for account validation
[**DeleteDepository**](DepositoriesApi.md#DeleteDepository) | **Delete** /v1/ach/depositories/{depositoryId} | Permanently deletes a depository and associated transfers. It cannot be undone. Immediately cancels any active Transfers for the depository.
[**GetDepositories**](DepositoriesApi.md#GetDepositories) | **Get** /v1/ach/depositories | A list of all Depository objects for the authentication context.
[**GetDepositoryByID**](DepositoriesApi.md#GetDepositoryByID) | **Get** /v1/ach/depositories/{depositoryId} | Get a Depository object for the supplied ID
[**UpdateDepository**](DepositoriesApi.md#UpdateDepository) | **Patch** /v1/ach/depositories/{depositoryId} | Updates the specified Depository by setting the values of the parameters passed. Any parameters not provided will be left unchanged.


# **AddDepository**
> Depository AddDepository(ctx, optional)
Create a new depository account for a Customer ID or Originator ID defined in the Parent parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AddDepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddDepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **depository** | [**optional.Interface of Depository**](Depository.md)|  | 

### Return type

[**Depository**](Depository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMicroDeposits**
> CreateMicroDeposits(ctx, depositoryId, optional)
Initiates micro deposits to be sent to the Depository institution for account validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **depositoryId** | **string**| Depository ID | 
 **optional** | ***CreateMicroDepositsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateMicroDepositsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amounts** | [**optional.Interface of []string**](string.md)| amounts that have been deposited | 
 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDepository**
> DeleteDepository(ctx, depositoryId, optional)
Permanently deletes a depository and associated transfers. It cannot be undone. Immediately cancels any active Transfers for the depository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **depositoryId** | **string**| Depository ID | 
 **optional** | ***DeleteDepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteDepositoryOpts struct

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

# **GetDepositories**
> Depositories GetDepositories(ctx, optional)
A list of all Depository objects for the authentication context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetDepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDepositoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Depositories**](Depositories.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDepositoryByID**
> Depository GetDepositoryByID(ctx, depositoryId, optional)
Get a Depository object for the supplied ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **depositoryId** | **string**| Depository ID | 
 **optional** | ***GetDepositoryByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDepositoryByIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Depository**](Depository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDepository**
> Depository UpdateDepository(ctx, depositoryId, optional)
Updates the specified Depository by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **depositoryId** | **string**| Depository ID | 
 **optional** | ***UpdateDepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateDepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **depository** | [**optional.Interface of Depository**](Depository.md)|  | 

### Return type

[**Depository**](Depository.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

