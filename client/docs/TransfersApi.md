# \TransfersApi

All URIs are relative to *http://localhost:9999*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTransfer**](TransfersApi.md#AddTransfer) | **Post** /v1/transfers | Create Transfer
[**DeleteTransferByID**](TransfersApi.md#DeleteTransferByID) | **Delete** /v1/transfers/{transferID} | Delete Transfer
[**GetTransferByID**](TransfersApi.md#GetTransferByID) | **Get** /v1/transfers/{transferID} | Get Transfer
[**GetTransfers**](TransfersApi.md#GetTransfers) | **Get** /v1/transfers | List Transfers



## AddTransfer

> Transfer AddTransfer(ctx, xUserID, createTransfer, optional)

Create Transfer

Create a new transfer between a Source and a Destination. Transfers can only be modified in the pending status. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xUserID** | **string**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
**createTransfer** | [**CreateTransfer**](CreateTransfer.md)|  | 
 **optional** | ***AddTransferOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddTransferOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 

### Return type

[**Transfer**](Transfer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransferByID

> DeleteTransferByID(ctx, transferID, xUserID, optional)

Delete Transfer

Remove a transfer for the specified userID. Its status will be updated as transfer is processed. It is only possible to delete (recall) a Transfer before it has been released from the financial institution. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferID** | **string**| transferID to delete | 
**xUserID** | **string**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
 **optional** | ***DeleteTransferByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTransferByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferByID

> Transfer GetTransferByID(ctx, transferID, xUserID, optional)

Get Transfer

Get a Transfer object for the supplied userID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferID** | **string**| transferID to retrieve | 
**xUserID** | **string**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
 **optional** | ***GetTransferByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransferByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 

### Return type

[**Transfer**](Transfer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfers

> []Transfer GetTransfers(ctx, xUserID, optional)

List Transfers

List all Transfers created for the given userID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xUserID** | **string**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
 **optional** | ***GetTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTransfersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| The number of items to skip before starting to collect the result set | [default to 0]
 **limit** | **optional.Int32**| The number of items to return | [default to 25]
 **status** | [**optional.Interface of TransferStatus**](.md)| Return only Transfers in this TransferStatus | 
 **startDate** | **optional.Time**| Return Transfers that are scheduled for this date or later in ISO-8601 format YYYY-MM-DD. Can optionally be used with endDate to specify a date range.  | 
 **endDate** | **optional.Time**| Return Transfers that are scheduled for this date or earlier in ISO-8601 format YYYY-MM-DD. Can optionally be used with startDate to specify a date range.  | 
 **organizationIDs** | **optional.String**| Comma separated list of organizationID values to return Transfer objects for. | 
 **customerIDs** | **optional.String**| Comma separated list of customerID values to return Transfer objects for. | 
 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 

### Return type

[**[]Transfer**](Transfer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

