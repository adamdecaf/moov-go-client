# \UserApi

All URIs are relative to *https://api.moov.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserLogin**](UserApi.md#CheckUserLogin) | **Get** /v1/users/login | Check if a cookie is valid and authentic for a user.
[**CreateUser**](UserApi.md#CreateUser) | **Post** /v1/users/create | Create a new user using an email address not seen before.
[**UserLogin**](UserApi.md#UserLogin) | **Post** /v1/users/login | Attempt to login with an email and password
[**UserLogout**](UserApi.md#UserLogout) | **Delete** /v1/users/login | Invalidat a user&#39;s cookie(s).


# **CheckUserLogin**
> CheckUserLogin(ctx, )
Check if a cookie is valid and authentic for a user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUser**
> User CreateUser(ctx, )
Create a new user using an email address not seen before.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogin**
> User UserLogin(ctx, )
Attempt to login with an email and password

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogout**
> UserLogout(ctx, )
Invalidat a user's cookie(s).

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

