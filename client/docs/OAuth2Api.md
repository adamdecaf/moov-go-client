# \OAuth2Api

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOAuthClientCredentials**](OAuth2Api.md#CheckOAuthClientCredentials) | **Get** /v1/oauth2/authorize | Verify OAuth2 Bearer token
[**CreateOAuth2Client**](OAuth2Api.md#CreateOAuth2Client) | **Post** /v1/oauth2/client | Create OAuth2 client credentials
[**CreateOAuth2Token**](OAuth2Api.md#CreateOAuth2Token) | **Post** /v1/oauth2/token | Generate OAuth2 access token
[**GetClientsForUserId**](OAuth2Api.md#GetClientsForUserId) | **Get** /v1/oauth2/clients | List OAuth2 clients for the authenticated user



## CheckOAuthClientCredentials

> CheckOAuthClientCredentials(ctx, authorization, optional)
Verify OAuth2 Bearer token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| &lt;ignored&gt; | 
 **optional** | ***CheckOAuthClientCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CheckOAuthClientCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuth2Client

> []OAuth2Client CreateOAuth2Client(ctx, optional)
Create OAuth2 client credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOAuth2ClientOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOAuth2ClientOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOAuth2Token

> OAuth2Token CreateOAuth2Token(ctx, optional)
Generate OAuth2 access token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateOAuth2TokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOAuth2TokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 
 **grantType** | **optional.String**| OAuth2 grant type (must be &#39;client_credentials&#39;) | [default to client_credentials]
 **clientId** | **optional.String**| OAuth2 client ID | 
 **clientSecret** | **optional.String**| OAuth2 client secret | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsForUserId

> []OAuth2Client GetClientsForUserId(ctx, optional)
List OAuth2 clients for the authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetClientsForUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClientsForUserIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **optional.String**| Optional Request ID allows application developer to trace requests through the systems logs | 
 **xIdempotencyKey** | **optional.String**| Idempotent key in the header which expires after 24 hours. These strings should contain enough entropy for to not collide with each other in your requests. | 

### Return type

[**[]OAuth2Client**](OAuth2Client.md)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

