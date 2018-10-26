# \CustomersApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomers**](CustomersApi.md#AddCustomers) | **Post** /v1/ach/customers | Create a new Customer object
[**DeleteCustomer**](CustomersApi.md#DeleteCustomer) | **Delete** /v1/ach/customers/{customerId} | Permanently deletes a customer and associated depositories and transfers. It cannot be undone. Immediately cancels any active Transfers for the customer.
[**GetCustomerByID**](CustomersApi.md#GetCustomerByID) | **Get** /v1/ach/customers/{customerId} | Get a Customer by ID
[**GetCustomers**](CustomersApi.md#GetCustomers) | **Get** /v1/ach/customers | Gets a list of Customers
[**GetDepositoriesByCustomerID**](CustomersApi.md#GetDepositoriesByCustomerID) | **Get** /v1/ach/customers/{customerId}/depositories | Get a list of Depository accounts for a Customer
[**GetDepositoriesByID**](CustomersApi.md#GetDepositoriesByID) | **Get** /v1/ach/customers/{customerId}/depositories/{depositoryId} | Get a Depository accounts for a Customer based on it&#39;s ID
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Patch** /v1/ach/customers/{customerId} | Updates the specified Customer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.


# **AddCustomers**
> Customer AddCustomers(ctx, uNKNOWNBASETYPE, optional)
Create a new Customer object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| A JSON object containing a new Customer | 
 **optional** | ***AddCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomer**
> DeleteCustomer(ctx, customerId, optional)
Permanently deletes a customer and associated depositories and transfers. It cannot be undone. Immediately cancels any active Transfers for the customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***DeleteCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **optional.String**| OAuth2 Bearer token | 
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

# **GetCustomerByID**
> Customer GetCustomerByID(ctx, customerId, optional)
Get a Customer by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***GetCustomerByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomerByIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomers**
> Customers GetCustomers(ctx, optional)
Gets a list of Customers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 
 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Customers**](Customers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDepositoriesByCustomerID**
> Depositories GetDepositoriesByCustomerID(ctx, customerId, optional)
Get a list of Depository accounts for a Customer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
 **optional** | ***GetDepositoriesByCustomerIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDepositoriesByCustomerIDOpts struct

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

# **GetDepositoriesByID**
> Depository GetDepositoriesByID(ctx, customerId, depositoryId, optional)
Get a Depository accounts for a Customer based on it's ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **depositoryId** | **string**| Depository ID | 
 **optional** | ***GetDepositoriesByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetDepositoriesByIDOpts struct

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

# **UpdateCustomer**
> Customer UpdateCustomer(ctx, customerId, uNKNOWNBASETYPE, optional)
Updates the specified Customer by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| Customer ID | 
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| A JSON object containing a new Customer | 
 **optional** | ***UpdateCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateCustomerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cookie** | **optional.String**| moov_auth Cookie | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Customer**](Customer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

