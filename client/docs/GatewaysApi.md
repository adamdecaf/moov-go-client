# \GatewaysApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGateway**](GatewaysApi.md#AddGateway) | **Post** /v1/ach/gateways | Create a new Gateway object
[**GetGateways**](GatewaysApi.md#GetGateways) | **Get** /v1/ach/gateways | Gets a list of Gatways


# **AddGateway**
> Gateway AddGateway(ctx, gateway, optional)
Create a new Gateway object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateway** | [**Gateway**](Gateway.md)|  | 
 **optional** | ***AddGatewayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddGatewayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Gateway**](Gateway.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGateways**
> Gateways GetGateways(ctx, optional)
Gets a list of Gatways

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGatewaysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Gateways**](Gateways.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

