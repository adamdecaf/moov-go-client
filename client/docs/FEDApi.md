# \FEDApi

All URIs are relative to *http://localhost:9999*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchFEDACH**](FEDApi.md#SearchFEDACH) | **Get** /v1/fed/ach/search | Search FEDACH names and metadata



## SearchFEDACH

> AchDictionary SearchFEDACH(ctx, optional)

Search FEDACH names and metadata

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchFEDACHOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchFEDACHOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Optional User ID used to perform this search | 
 **name** | **optional.String**| FEDACH Financial Institution Name | 
 **routingNumber** | **optional.String**| FEDACH Routing Number for a Financial Institution | 
 **state** | **optional.String**| FEDACH Financial Institution State | 
 **city** | **optional.String**| FEDACH Financial Institution City | 
 **postalCode** | **optional.String**| FEDACH Financial Institution Postal Code | 
 **limit** | **optional.Int32**| Maximum results returned by a search | 

### Return type

[**AchDictionary**](ACHDictionary.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

