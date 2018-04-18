# \AuthGroupsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupMember**](AuthGroupsApi.md#CreateGroupMember) | **Post** /platform/1/auth/groups/{Group}/members | 
[**DeleteGroupMember**](AuthGroupsApi.md#DeleteGroupMember) | **Delete** /platform/1/auth/groups/{Group}/members/{GroupMemberId} | 
[**ListGroupMembers**](AuthGroupsApi.md#ListGroupMembers) | **Get** /platform/1/auth/groups/{Group}/members | 


# **CreateGroupMember**
> CreateResponse CreateGroupMember(ctx, groupMember, group, optional)


Add a member to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupMember** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
  **group** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMember** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
 **group** | **string**|  | 
 **zone** | **string**| Filter group members by zone. | 
 **provider** | **string**| Filter group members by provider. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupMember**
> DeleteGroupMember(ctx, groupMemberId, group, optional)


Remove the member from the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupMemberId** | **string**| Remove the member from the group. | 
  **group** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupMemberId** | **string**| Remove the member from the group. | 
 **group** | **string**|  | 
 **zone** | **string**| Filter group members by zone. | 
 **provider** | **string**| Filter group members by provider. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupMembers**
> GroupMembers ListGroupMembers(ctx, group, optional)


List all the members of the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **group** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**|  | 
 **resolveNames** | **bool**| Resolve names of personas. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **zone** | **string**| Filter group members by zone. | 
 **provider** | **string**| Filter group members by provider. | 

### Return type

[**GroupMembers**](GroupMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

