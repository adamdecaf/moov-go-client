# \OFACApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCompanyNameWatch**](OFACApi.md#AddCompanyNameWatch) | **Post** /v1/ofac/companies/watch | Add company watch by name. The match percentage will be included in the webhook&#39;s JSON payload.
[**AddCompanyWatch**](OFACApi.md#AddCompanyWatch) | **Post** /v1/ofac/companies/{companyId}/watch | Add OFAC watch on a Company
[**AddCustomerNameWatch**](OFACApi.md#AddCustomerNameWatch) | **Post** /v1/ofac/customers/watch | Add customer watch by name. The match percentage will be included in the webhook&#39;s JSON payload.
[**AddCustomerWatch**](OFACApi.md#AddCustomerWatch) | **Post** /v1/ofac/customers/{customerId}/watch | Add OFAC watch on a Customer
[**GetCompany**](OFACApi.md#GetCompany) | **Get** /v1/ofac/companies/{companyId} | Get information about a company, trust or organization such as addresses, alternate names, and remarks.
[**GetCustomer**](OFACApi.md#GetCustomer) | **Get** /v1/ofac/customers/{customerId} | Get information about a customer, addresses, alternate names, and their SDN metadata.
[**GetLatestDownloads**](OFACApi.md#GetLatestDownloads) | **Get** /v1/ofac/downloads | Return list of recent downloads of OFAC data
[**GetSDN**](OFACApi.md#GetSDN) | **Get** /v1/ofac/sdn/{sdnId} | Specially designated national
[**GetSDNAddresses**](OFACApi.md#GetSDNAddresses) | **Get** /v1/ofac/sdn/{sdnId}/addresses | Get addresses for a given SDN
[**GetSDNAltNames**](OFACApi.md#GetSDNAltNames) | **Get** /v1/ofac/sdn/{sdnId}/alts | Get alternate names for a given SDN
[**RemoveCompanyNameWatch**](OFACApi.md#RemoveCompanyNameWatch) | **Delete** /v1/ofac/companies/watch/{watchId} | Remove a Company name watch
[**RemoveCompanyWatch**](OFACApi.md#RemoveCompanyWatch) | **Delete** /v1/ofac/companies/{companyId}/watch/{watchId} | Remove company watch
[**RemoveCustomerNameWatch**](OFACApi.md#RemoveCustomerNameWatch) | **Delete** /v1/ofac/customers/watch/{watchId} | Remove a Customer name watch
[**RemoveCustomerWatch**](OFACApi.md#RemoveCustomerWatch) | **Delete** /v1/ofac/customers/{customerId}/watch/{watchId} | Remove customer watch
[**Search**](OFACApi.md#Search) | **Get** /v1/ofac/search | Search SDN names and metadata
[**UpdateCompanyStatus**](OFACApi.md#UpdateCompanyStatus) | **Put** /v1/ofac/companies/{companyId} | Update a Companies sanction status to always block or always allow transactions.
[**UpdateCustomerStatus**](OFACApi.md#UpdateCustomerStatus) | **Put** /v1/ofac/customers/{customerId} | Update a Customer&#39;s sanction status to always block or always allow transactions.



## AddCompanyNameWatch

> Watch AddCompanyNameWatch(ctx, name, watchRequest, optional)
Add company watch by name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Company name used to match and send watch notifications | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCompanyWatch

> Watch AddCompanyWatch(ctx, companyId, watchRequest, optional)
Add OFAC watch on a Company

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string**| Company ID | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCustomerNameWatch

> Watch AddCustomerNameWatch(ctx, name, watchRequest, optional)
Add customer watch by name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Individual name used to match and send watch notifications | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCustomerWatch

> Watch AddCustomerWatch(ctx, customerId, watchRequest, optional)
Add OFAC watch on a Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string**| Customer ID | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Watch**](Watch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> OfacCompany GetCompany(ctx, companyId, optional)
Get information about a company, trust or organization such as addresses, alternate names, and remarks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string**| Company ID | 
 **optional** | ***GetCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCompanyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacCompany**](OFACCompany.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomer

> OfacCustomer GetCustomer(ctx, customerId, optional)
Get information about a customer, addresses, alternate names, and their SDN metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string**| Customer ID | 
 **optional** | ***GetCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**OfacCustomer**](OFACCustomer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestDownloads

> []Download GetLatestDownloads(ctx, optional)
Return list of recent downloads of OFAC data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetLatestDownloadsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLatestDownloadsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Maximum results returned by a search | 

### Return type

[**[]Download**](Download.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDN

> Sdn GetSDN(ctx, sdnId, optional)
Specially designated national

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**Sdn**](SDN.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDNAddresses

> []Address GetSDNAddresses(ctx, sdnId, optional)
Get addresses for a given SDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAddressesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSDNAltNames

> []Alt GetSDNAltNames(ctx, sdnId, optional)
Get alternate names for a given SDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnId** | **string**| SDN ID | 
 **optional** | ***GetSDNAltNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAltNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

[**[]Alt**](Alt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyNameWatch

> RemoveCompanyNameWatch(ctx, watchId, name, optional)
Remove a Company name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchId** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Company name watch | 
 **optional** | ***RemoveCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCompanyWatch

> RemoveCompanyWatch(ctx, companyId, watchId, optional)
Remove company watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string**| Company ID | 
**watchId** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomerNameWatch

> RemoveCustomerNameWatch(ctx, watchId, name, optional)
Remove a Customer name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchId** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Customer or Company name watch | 
 **optional** | ***RemoveCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomerWatch

> RemoveCustomerWatch(ctx, customerId, watchId, optional)
Remove customer watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string**| Customer ID | 
**watchId** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> Search Search(ctx, optional)
Search SDN names and metadata

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **q** | **optional.String**| Search across Name, Alt Names, and Address fields for all SDN entries. Entries may be returned in all response sub-objects. | 
 **name** | **optional.String**| Name which could correspond to a human on the SDN list. Only SDN results will be returned. | 
 **address** | **optional.String**| Phsical address which could correspond to a human on the SDN list. Only Address results will be returned. | 
 **altName** | **optional.String**| Alternate name which could correspond to a human on the SDN list. Only Alt name results will be returned. | 
 **limit** | **optional.Int32**| Maximum results returned by a search | 

### Return type

[**Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyStatus

> UpdateCompanyStatus(ctx, companyId, updateCompanyStatus, optional)
Update a Companies sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string**| Company ID | 
**updateCompanyStatus** | [**UpdateCompanyStatus**](UpdateCompanyStatus.md)|  | 
 **optional** | ***UpdateCompanyStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCompanyStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerStatus

> UpdateCustomerStatus(ctx, customerId, updateCustomerStatus, optional)
Update a Customer's sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string**| Customer ID | 
**updateCustomerStatus** | [**UpdateCustomerStatus**](UpdateCustomerStatus.md)|  | 
 **optional** | ***UpdateCustomerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomerStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

