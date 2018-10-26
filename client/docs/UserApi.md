# \UserApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserLogin**](UserApi.md#CheckUserLogin) | **Get** /v1/users/login | Check if a cookie is valid and authentic for a user.
[**CreateUser**](UserApi.md#CreateUser) | **Post** /v1/users/create | Create a new user using an email address not seen before.
[**UpdateUserProfile**](UserApi.md#UpdateUserProfile) | **Patch** /v1/users/{user_id} | Update a User&#39;s profile information
[**UserLogin**](UserApi.md#UserLogin) | **Post** /v1/users/login | Attempt to login with an email and password
[**UserLogout**](UserApi.md#UserLogout) | **Delete** /v1/users/login | Invalidat a user&#39;s cookie(s).


# **CheckUserLogin**
> CheckUserLogin(ctx, optional)
Check if a cookie is valid and authentic for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CheckUserLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckUserLoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, uNKNOWNBASETYPE)
Create a new user using an email address not seen before.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| A Moov API User object | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserProfile**
> UpdateUserProfile(ctx, userId, optional)
Update a User's profile information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| Moov API User ID | 
 **optional** | ***UpdateUserProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateUserProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.String**| moov_auth Cookie | 
 **uNKNOWNBASETYPE** | [**optional.Interface of UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| User profile information | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogin**
> User UserLogin(ctx, uNKNOWNBASETYPE)
Attempt to login with an email and password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Authenticating with an email and password | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogout**
> UserLogout(ctx, optional)
Invalidat a user's cookie(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserLogoutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserLogoutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.String**| moov_auth Cookie | 

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

