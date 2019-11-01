# \WireFilesApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWireFile**](WireFilesApi.md#CreateWireFile) | **Post** /v1/wire/files/create | Create a new File object
[**DeleteWireFileByID**](WireFilesApi.md#DeleteWireFileByID) | **Delete** /v1/wire/files/{fileID} | Permanently deletes a File and associated FEDWireMessage. It cannot be undone.
[**GetWireFileByID**](WireFilesApi.md#GetWireFileByID) | **Get** /v1/wire/files/{fileID} | Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.
[**GetWireFileContents**](WireFilesApi.md#GetWireFileContents) | **Get** /v1/wire/files/{fileID}/contents | Assembles the existing file witha FEDWireMessage, Returns plaintext file.
[**GetWireFiles**](WireFilesApi.md#GetWireFiles) | **Get** /v1/wire/files | Gets a list of Files
[**UpdateWireFileByID**](WireFilesApi.md#UpdateWireFileByID) | **Post** /v1/wire/files/{fileID} | Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
[**ValidateWireFile**](WireFilesApi.md#ValidateWireFile) | **Get** /v1/wire/files/{fileID}/validate | Validates the existing file. You need only supply the unique File identifier that was returned upon creation.



## CreateWireFile

> File2 CreateWireFile(ctx, createFile2, optional)

Create a new File object

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createFile2** | [**CreateFile2**](CreateFile2.md)| Content of the WIRE file (in json or raw text) | 
 **optional** | ***CreateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File2**](File_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWireFileByID

> DeleteWireFileByID(ctx, fileID, optional)

Permanently deletes a File and associated FEDWireMessage. It cannot be undone.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***DeleteWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFileByID

> File2 GetWireFileByID(ctx, fileID, optional)

Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***GetWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File2**](File_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFileContents

> string GetWireFileContents(ctx, fileID, optional)

Assembles the existing file witha FEDWireMessage, Returns plaintext file.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***GetWireFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFileContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireFiles

> []File2 GetWireFiles(ctx, optional)

Gets a list of Files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWireFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWireFilesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]File2**](File_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWireFileByID

> File2 UpdateWireFileByID(ctx, fileID, createFile2, optional)

Updates the specified FEDWire Message by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
**createFile2** | [**CreateFile2**](CreateFile2.md)|  | 
 **optional** | ***UpdateWireFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWireFileByIDOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**File2**](File_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateWireFile

> File2 ValidateWireFile(ctx, fileID, optional)

Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileID** | **string**| File ID | 
 **optional** | ***ValidateWireFileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateWireFileOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File2**](File_2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

