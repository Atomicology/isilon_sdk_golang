# \AuthRolesApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleMember**](AuthRolesApi.md#CreateRoleMember) | **Post** /platform/1/auth/roles/{Role}/members | 
[**CreateRolePrivilege**](AuthRolesApi.md#CreateRolePrivilege) | **Post** /platform/1/auth/roles/{Role}/privileges | 
[**DeleteRoleMember**](AuthRolesApi.md#DeleteRoleMember) | **Delete** /platform/1/auth/roles/{Role}/members/{RoleMemberId} | 
[**DeleteRolePrivilege**](AuthRolesApi.md#DeleteRolePrivilege) | **Delete** /platform/1/auth/roles/{Role}/privileges/{RolePrivilegeId} | 
[**ListRoleMembers**](AuthRolesApi.md#ListRoleMembers) | **Get** /platform/1/auth/roles/{Role}/members | 
[**ListRolePrivileges**](AuthRolesApi.md#ListRolePrivileges) | **Get** /platform/1/auth/roles/{Role}/privileges | 


# **CreateRoleMember**
> CreateResponse CreateRoleMember(ctx, roleMember, role)


Add a member to the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleMember** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
  **role** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRolePrivilege**
> CreateResponse CreateRolePrivilege(ctx, rolePrivilege, role)


Add a privilege to the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **rolePrivilege** | [**AuthIdNtokenPrivilegeItem**](AuthIdNtokenPrivilegeItem.md)|  | 
  **role** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRoleMember**
> DeleteRoleMember(ctx, roleMemberId, role)


Remove a member from the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **roleMemberId** | **string**| Remove a member from the role. | 
  **role** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRolePrivilege**
> DeleteRolePrivilege(ctx, rolePrivilegeId, role)


Remove a privilege from a role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **rolePrivilegeId** | **string**| Remove a privilege from a role. | 
  **role** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoleMembers**
> GroupMembers ListRoleMembers(ctx, role, optional)


List all the members of the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **role** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string**|  | 
 **resolveNames** | **bool**| Resolve names of personas. | 

### Return type

[**GroupMembers**](GroupMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRolePrivileges**
> RolePrivileges ListRolePrivileges(ctx, role)


List all privileges in the role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **role** | **string**|  | 

### Return type

[**RolePrivileges**](RolePrivileges.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

