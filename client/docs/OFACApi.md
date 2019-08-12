# \OFACApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOFACCompanyNameWatch**](OFACApi.md#AddOFACCompanyNameWatch) | **Post** /v1/ofac/companies/watch | Add company watch by name. The match percentage will be included in the webhook&#39;s JSON payload.
[**AddOFACCompanyWatch**](OFACApi.md#AddOFACCompanyWatch) | **Post** /v1/ofac/companies/{companyID}/watch | Add OFAC watch on a Company
[**AddOFACCustomerNameWatch**](OFACApi.md#AddOFACCustomerNameWatch) | **Post** /v1/ofac/customers/watch | Add customer watch by name. The match percentage will be included in the webhook&#39;s JSON payload.
[**AddOFACCustomerWatch**](OFACApi.md#AddOFACCustomerWatch) | **Post** /v1/ofac/customers/{customerID}/watch | Add OFAC watch on a Customer
[**GetLatestDownloads**](OFACApi.md#GetLatestDownloads) | **Get** /v1/ofac/downloads | Return list of recent downloads of OFAC data
[**GetOFACCompany**](OFACApi.md#GetOFACCompany) | **Get** /v1/ofac/companies/{companyID} | Get information about a company, trust or organization such as addresses, alternate names, and remarks.
[**GetOFACCustomer**](OFACApi.md#GetOFACCustomer) | **Get** /v1/ofac/customers/{customerID} | Get information about a customer, addresses, alternate names, and their SDN metadata.
[**GetSDN**](OFACApi.md#GetSDN) | **Get** /v1/ofac/sdn/{sdnID} | Specially designated national
[**GetSDNAddresses**](OFACApi.md#GetSDNAddresses) | **Get** /v1/ofac/sdn/{sdnID}/addresses | Get addresses for a given SDN
[**GetSDNAltNames**](OFACApi.md#GetSDNAltNames) | **Get** /v1/ofac/sdn/{sdnID}/alts | Get alternate names for a given SDN
[**RemoveOFACCompanyNameWatch**](OFACApi.md#RemoveOFACCompanyNameWatch) | **Delete** /v1/ofac/companies/watch/{watchID} | Remove a Company name watch
[**RemoveOFACCompanyWatch**](OFACApi.md#RemoveOFACCompanyWatch) | **Delete** /v1/ofac/companies/{companyID}/watch/{watchID} | Remove company watch
[**RemoveOFACCustomerNameWatch**](OFACApi.md#RemoveOFACCustomerNameWatch) | **Delete** /v1/ofac/customers/watch/{watchID} | Remove a Customer name watch
[**RemoveOFACCustomerWatch**](OFACApi.md#RemoveOFACCustomerWatch) | **Delete** /v1/ofac/customers/{customerID}/watch/{watchID} | Remove customer watch
[**Search**](OFACApi.md#Search) | **Get** /v1/ofac/search | Search SDN names and metadata
[**UpdateOFACCompanyStatus**](OFACApi.md#UpdateOFACCompanyStatus) | **Put** /v1/ofac/companies/{companyID} | Update a Companies sanction status to always block or always allow transactions.
[**UpdateOFACCustomerStatus**](OFACApi.md#UpdateOFACCustomerStatus) | **Put** /v1/ofac/customers/{customerID} | Update a Customer&#39;s sanction status to always block or always allow transactions.



## AddOFACCompanyNameWatch

> Watch AddOFACCompanyNameWatch(ctx, name, watchRequest, optional)
Add company watch by name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Company name used to match and send watch notifications | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddOFACCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOFACCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## AddOFACCompanyWatch

> Watch AddOFACCompanyWatch(ctx, companyID, watchRequest, optional)
Add OFAC watch on a Company

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddOFACCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOFACCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## AddOFACCustomerNameWatch

> Watch AddOFACCustomerNameWatch(ctx, name, watchRequest, optional)
Add customer watch by name. The match percentage will be included in the webhook's JSON payload.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string**| Individual name used to match and send watch notifications | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddOFACCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOFACCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## AddOFACCustomerWatch

> Watch AddOFACCustomerWatch(ctx, customerID, watchRequest, optional)
Add OFAC watch on a Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**watchRequest** | [**WatchRequest**](WatchRequest.md)|  | 
 **optional** | ***AddOFACCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOFACCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## GetOFACCompany

> OfacCompany GetOFACCompany(ctx, companyID, optional)
Get information about a company, trust or organization such as addresses, alternate names, and remarks.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
 **optional** | ***GetOFACCompanyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOFACCompanyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## GetOFACCustomer

> OfacCustomer GetOFACCustomer(ctx, customerID, optional)
Get information about a customer, addresses, alternate names, and their SDN metadata.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
 **optional** | ***GetOFACCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOFACCustomerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## GetSDN

> Sdn GetSDN(ctx, sdnID, optional)
Specially designated national

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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

> []Address GetSDNAddresses(ctx, sdnID, optional)
Get addresses for a given SDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNAddressesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAddressesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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

> []Alt GetSDNAltNames(ctx, sdnID, optional)
Get alternate names for a given SDN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sdnID** | **string**| SDN ID | 
 **optional** | ***GetSDNAltNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSDNAltNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOFACCompanyNameWatch

> RemoveOFACCompanyNameWatch(ctx, watchID, name, optional)
Remove a Company name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchID** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Company name watch | 
 **optional** | ***RemoveOFACCompanyNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOFACCompanyNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOFACCompanyWatch

> RemoveOFACCompanyWatch(ctx, companyID, watchID, optional)
Remove company watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**watchID** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveOFACCompanyWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOFACCompanyWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOFACCustomerNameWatch

> RemoveOFACCustomerNameWatch(ctx, watchID, name, optional)
Remove a Customer name watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watchID** | **string**| Watch ID, used to identify a specific watch | 
**name** | **string**| Customer or Company name watch | 
 **optional** | ***RemoveOFACCustomerNameWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOFACCustomerNameWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## RemoveOFACCustomerWatch

> RemoveOFACCustomerWatch(ctx, customerID, watchID, optional)
Remove customer watch

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**watchID** | **string**| Watch ID, used to identify a specific watch | 
 **optional** | ***RemoveOFACCustomerWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveOFACCustomerWatchOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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
 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **q** | **optional.String**| Search across Name, Alt Names, and Address fields for all SDN entries. Entries may be returned in all response sub-objects. | 
 **name** | **optional.String**| Name which could correspond to a human on the SDN list. Only SDN results will be returned. | 
 **address** | **optional.String**| Phsical address which could correspond to a human on the SDN list. Only Address results will be returned. | 
 **city** | **optional.String**| City name as desginated by SDN guidelines. Only Address results will be returned. | 
 **state** | **optional.String**| State name as desginated by SDN guidelines. Only Address results will be returned. | 
 **providence** | **optional.String**| Providence name as desginated by SDN guidelines. Only Address results will be returned. | 
 **zip** | **optional.String**| Zip code as desginated by SDN guidelines. Only Address results will be returned. | 
 **country** | **optional.String**| Country name as desginated by SDN guidelines. Only Address results will be returned. | 
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


## UpdateOFACCompanyStatus

> UpdateOFACCompanyStatus(ctx, companyID, updateCompanyStatus, optional)
Update a Companies sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyID** | **string**| Company ID | 
**updateCompanyStatus** | [**UpdateCompanyStatus**](UpdateCompanyStatus.md)|  | 
 **optional** | ***UpdateOFACCompanyStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOFACCompanyStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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


## UpdateOFACCustomerStatus

> UpdateOFACCustomerStatus(ctx, customerID, updateCustomerStatus, optional)
Update a Customer's sanction status to always block or always allow transactions.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerID** | **string**| Customer ID | 
**updateCustomerStatus** | [**UpdateCustomerStatus**](UpdateCustomerStatus.md)|  | 
 **optional** | ***UpdateOFACCustomerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOFACCustomerStatusOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

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

