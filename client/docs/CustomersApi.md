# \CustomersApi

All URIs are relative to *http://localhost:9999*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomerAddress**](CustomersApi.md#AddCustomerAddress) | **Post** /v1/customers/{customerID}/address | Add customer address
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /v1/customers | Create customer
[**CreateCustomerAccount**](CustomersApi.md#CreateCustomerAccount) | **Post** /v1/customers/{customerID}/accounts | Create Customer Account
[**DecryptAccountNumber**](CustomersApi.md#DecryptAccountNumber) | **Post** /v1/customers/{customerID}/accounts/{accountID}/decrypt | Decrypt Account Number
[**DeleteCustomerAccount**](CustomersApi.md#DeleteCustomerAccount) | **Delete** /v1/customers/{customerID}/accounts | Delete Customer Account
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /v1/customers/{customerID} | Retrieve customer
[**GetCustomerAccounts**](CustomersApi.md#GetCustomerAccounts) | **Get** /v1/customers/{customerID}/accounts | Get Customer Accounts
[**GetCustomerDisclaimers**](CustomersApi.md#GetCustomerDisclaimers) | **Get** /v1/customers/{customerID}/disclaimers | Get customer disclaimers
[**GetCustomerDocumentContents**](CustomersApi.md#GetCustomerDocumentContents) | **Get** /v1/customers/{customerID}/documents/{documentID} | Get customer document
[**GetCustomerDocuments**](CustomersApi.md#GetCustomerDocuments) | **Get** /v1/customers/{customerID}/documents | Get customer documents
[**GetLatestOFACSearch**](CustomersApi.md#GetLatestOFACSearch) | **Get** /v1/customers/{customerID}/ofac | Get latest OFAC search
[**RefreshOFACSearch**](CustomersApi.md#RefreshOFACSearch) | **Put** /v1/customers/{customerID}/refresh/ofac | Refresh customer OFAC search
[**ReplaceCustomerMetadata**](CustomersApi.md#ReplaceCustomerMetadata) | **Put** /v1/customers/{customerID}/metadata | Update customer metadata
[**UpdateCustomerStatus**](CustomersApi.md#UpdateCustomerStatus) | **Put** /v1/customers/{customerID}/status | Update customer status
[**UploadCustomerDocument**](CustomersApi.md#UploadCustomerDocument) | **Post** /v1/customers/{customerID}/documents | Upload document
[**ValidateAccount**](CustomersApi.md#ValidateAccount) | **Post** /v1/customers/{customerID}/accounts/{accountID}/validate | Validate Account



## AddCustomerAddress

> Customer AddCustomerAddress(ctx, customerID, createCustomerAddress, optional)

Add customer address

Add an Address onto an existing Customer record

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to add the address onto | 
**createCustomerAddress** | [**CreateCustomerAddress**](CreateCustomerAddress.md)|  | 
 **optional** | ***AddCustomerAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCustomerAddressOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Customer**](Customer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomer

> Customer CreateCustomer(ctx, createCustomer, optional)

Create customer

Create a Customer object from the given details of a human or business

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createCustomer** | [**CreateCustomer**](CreateCustomer.md)|  | 
 **optional** | ***CreateCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCustomerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Customer**](Customer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomerAccount

> Account CreateCustomerAccount(ctx, customerID, optional)

Create Customer Account

Create an account for the given customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to add an Account onto | 
 **optional** | ***CreateCustomerAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCustomerAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
 **createAccount** | [**optional.Interface of CreateAccount**](CreateAccount.md)|  | 

### Return type

[**Account**](Account.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptAccountNumber

> TransitAccountNumber DecryptAccountNumber(ctx, customerID, accountID, optional)

Decrypt Account Number

Return the account number encrypted with a shared secret for application requests. This encryption key is different from the key used for persistence. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer the accountID belongs to | 
**accountID** | **string**| accountID of the Account to validate | 
 **optional** | ***DecryptAccountNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DecryptAccountNumberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**TransitAccountNumber**](TransitAccountNumber.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerAccount

> DeleteCustomerAccount(ctx, customerID, optional)

Delete Customer Account

Remove an account from the given Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to remove an Account | 
 **optional** | ***DeleteCustomerAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCustomerAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

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


## GetCustomer

> Customer GetCustomer(ctx, customerID, optional)

Retrieve customer

Get the Customer object and metadata for the customerID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID that identifies this Customer | 
 **optional** | ***GetCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Customer**](Customer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAccounts

> []Account GetCustomerAccounts(ctx, customerID, optional)

Get Customer Accounts

Retrieve all accounts for the given customer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to get Accounts for | 
 **optional** | ***GetCustomerAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**[]Account**](Account.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerDisclaimers

> []Disclaimer GetCustomerDisclaimers(ctx, customerID, optional)

Get customer disclaimers

Get active disclaimers for the given customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to get disclaimers | 
 **optional** | ***GetCustomerDisclaimersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerDisclaimersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**[]Disclaimer**](Disclaimer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerDocumentContents

> *os.File GetCustomerDocumentContents(ctx, customerID, documentID, optional)

Get customer document

Retrieve the referenced document

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to get a Document | 
**documentID** | **string**| documentID to identify a Document | 
 **optional** | ***GetCustomerDocumentContentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerDocumentContentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, image/_*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerDocuments

> []Document GetCustomerDocuments(ctx, customerID, optional)

Get customer documents

Get documents for a customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to get all Documents | 
 **optional** | ***GetCustomerDocumentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerDocumentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**[]Document**](Document.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestOFACSearch

> OfacSearch GetLatestOFACSearch(ctx, customerID, optional)

Get latest OFAC search

Get the latest OFAC search for a customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to get latest OFAC search | 
 **optional** | ***GetLatestOFACSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestOFACSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**OfacSearch**](OFACSearch.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshOFACSearch

> OfacSearch RefreshOFACSearch(ctx, customerID, optional)

Refresh customer OFAC search

Refresh OFAC search for a given Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to refresh OFAC search | 
 **optional** | ***RefreshOFACSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RefreshOFACSearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**OfacSearch**](OFACSearch.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCustomerMetadata

> Customer ReplaceCustomerMetadata(ctx, customerID, customerMetadata, optional)

Update customer metadata

Replace the metadata object for a customer. Metadata is a map of unique keys associated to values to act as foreign key relationships or arbitrary data associated to a Customer.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to add the metadata onto | 
**customerMetadata** | [**CustomerMetadata**](CustomerMetadata.md)|  | 
 **optional** | ***ReplaceCustomerMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplaceCustomerMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Customer**](Customer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerStatus

> Customer UpdateCustomerStatus(ctx, customerID, updateCustomerStatus, optional)

Update customer status

Update the status for a customer, which can only be updated by authenticated users with permissions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to update the CustomerStatus | 
**updateCustomerStatus** | [**UpdateCustomerStatus**](UpdateCustomerStatus.md)|  | 
 **optional** | ***UpdateCustomerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomerStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Customer**](Customer.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCustomerDocument

> Document UploadCustomerDocument(ctx, customerID, type_, file, optional)

Upload document

Upload a document for the given customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer to add a document | 
**type_** | **string**| Document type (see Document type for values) | 
**file** | ***os.File*****os.File**| Document to be uploaded | 
 **optional** | ***UploadCustomerDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadCustomerDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 

### Return type

[**Document**](Document.md)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAccount

> ValidateAccount(ctx, customerID, accountID, optional)

Validate Account

Initiate or validatae account with availble validation strategies. Currently the only available strategy is micro-deposits. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| customerID of the Customer the accountID belongs to | 
**accountID** | **string**| accountID of the Account to validate | 
 **optional** | ***ValidateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ValidateAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional requestID allows application developer to trace requests through the systems logs | 
 **xUserID** | **optional.String**| Unique userID set by an auth proxy or client to identify and isolate objects. | 
 **updateValidation** | [**optional.Interface of UpdateValidation**](UpdateValidation.md)|  | 

### Return type

 (empty response body)

### Authorization

[GatewayAuth](../README.md#GatewayAuth), [LoginAuth](../README.md#LoginAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

