# \FilesApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBatchToFile**](FilesApi.md#AddBatchToFile) | **Post** /v1/ach/files/{file_id}/batches | Add Batch to File
[**AddFile**](FilesApi.md#AddFile) | **Post** /v1/ach/files/create | Create a new File object
[**DeleteACHFile**](FilesApi.md#DeleteACHFile) | **Delete** /v1/ach/files/{file_id} | Permanently deletes a File and associated Batches. It cannot be undone.
[**DeleteFileBatch**](FilesApi.md#DeleteFileBatch) | **Delete** /v1/ach/files/{file_id}/batches/{batch_id} | Delete a Batch from a File
[**GetFileBatch**](FilesApi.md#GetFileBatch) | **Get** /v1/ach/files/{file_id}/batches/{batch_id} | Get a specific Batch on a FIle
[**GetFileBatches**](FilesApi.md#GetFileBatches) | **Get** /v1/ach/files/{file_id}/batches | Get the batches on a File.
[**GetFileByID**](FilesApi.md#GetFileByID) | **Get** /v1/ach/files/{file_id} | Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.
[**GetFileContents**](FilesApi.md#GetFileContents) | **Get** /v1/ach/files/{file_id}/contents | Assembles the existing file (batches and controls) records, computes sequence numbers and totals. Returns plaintext file.
[**GetFiles**](FilesApi.md#GetFiles) | **Get** /v1/ach/files | Gets a list of Files
[**UpdateFile**](FilesApi.md#UpdateFile) | **Post** /v1/ach/files/{file_id} | Updates the specified File Header by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
[**ValidateFile**](FilesApi.md#ValidateFile) | **Get** /v1/ach/files/{file_id}/validate | Validates the existing file. You need only supply the unique File identifier that was returned upon creation.


# **AddBatchToFile**
> AddBatchToFile(ctx, fileId, optional)
Add Batch to File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***AddBatchToFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddBatchToFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFile**
> File AddFile(ctx, file, optional)
Create a new File object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | [**File**](File.md)| Content of the ACH file (in json or raw text) | 
 **optional** | ***AddFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteACHFile**
> DeleteACHFile(ctx, fileId, optional)
Permanently deletes a File and associated Batches. It cannot be undone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***DeleteACHFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteACHFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFileBatch**
> DeleteFileBatch(ctx, fileId, batchId, optional)
Delete a Batch from a File

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
  **batchId** | **string**| Batch ID | 
 **optional** | ***DeleteFileBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteFileBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileBatch**
> Batch GetFileBatch(ctx, fileId, batchId, optional)
Get a specific Batch on a FIle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
  **batchId** | **string**| Batch ID | 
 **optional** | ***GetFileBatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileBatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

[**Batch**](Batch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileBatches**
> Batches GetFileBatches(ctx, fileId, optional)
Get the batches on a File.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***GetFileBatchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileBatchesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

[**Batches**](Batches.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileByID**
> File GetFileByID(ctx, fileId, optional)
Retrieves the details of an existing File. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***GetFileByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileByIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileContents**
> string GetFileContents(ctx, fileId, optional)
Assembles the existing file (batches and controls) records, computes sequence numbers and totals. Returns plaintext file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***GetFileContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFileContentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFiles**
> Files GetFiles(ctx, optional)
Gets a list of Files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

[**Files**](Files.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFile**
> File UpdateFile(ctx, fileId, optional)
Updates the specified File Header by setting the values of the parameters passed. Any parameters not provided will be left unchanged.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***UpdateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **file** | [**optional.Interface of File**](File.md)|  | 

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFile**
> File ValidateFile(ctx, fileId, optional)
Validates the existing file. You need only supply the unique File identifier that was returned upon creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **string**| File ID | 
 **optional** | ***ValidateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

[**File**](File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

