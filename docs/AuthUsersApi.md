# \AuthUsersApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserMemberOfItem**](AuthUsersApi.md#CreateUserMemberOfItem) | **Post** /platform/3/auth/users/{User}/member-of | 
[**DeleteUserMemberOfMemberOf**](AuthUsersApi.md#DeleteUserMemberOfMemberOf) | **Delete** /platform/3/auth/users/{User}/member-of/{UserMemberOfMemberOf} | 
[**ListUserMemberOf**](AuthUsersApi.md#ListUserMemberOf) | **Get** /platform/3/auth/users/{User}/member-of | 
[**UpdateUserChangePassword**](AuthUsersApi.md#UpdateUserChangePassword) | **Put** /platform/3/auth/users/{User}/change-password | 


# **CreateUserMemberOfItem**
> CreateResponse CreateUserMemberOfItem(ctx, userMemberOfItem, user, optional)


Add the user to a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userMemberOfItem** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
  **user** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userMemberOfItem** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
 **user** | **string**|  | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserMemberOfMemberOf**
> DeleteUserMemberOfMemberOf(ctx, userMemberOfMemberOf, user, optional)


Remove the user from the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userMemberOfMemberOf** | **string**| Remove the user from the group. | 
  **user** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userMemberOfMemberOf** | **string**| Remove the user from the group. | 
 **user** | **string**|  | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserMemberOf**
> UserMemberOf ListUserMemberOf(ctx, user, optional)


List all groups the user is a member of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **user** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string**|  | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **zone** | **string**| Filter groups by zone. | 
 **provider** | **string**| Filter groups by provider. | 

### Return type

[**UserMemberOf**](UserMemberOf.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserChangePassword**
> UpdateUserChangePassword(ctx, userChangePassword, user, optional)


Change the user's password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **userChangePassword** | [**UserChangePassword**](UserChangePassword.md)|  | 
  **user** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userChangePassword** | [**UserChangePassword**](UserChangePassword.md)|  | 
 **user** | **string**|  | 
 **zone** | **string**| Specifies access zone containing user. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

