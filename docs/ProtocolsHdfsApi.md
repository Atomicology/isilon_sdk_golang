# \ProtocolsHdfsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProxyusersNameMember**](ProtocolsHdfsApi.md#CreateProxyusersNameMember) | **Post** /platform/1/protocols/hdfs/proxyusers/{Name}/members | 
[**DeleteProxyusersNameMember**](ProtocolsHdfsApi.md#DeleteProxyusersNameMember) | **Delete** /platform/1/protocols/hdfs/proxyusers/{Name}/members/{ProxyusersNameMemberId} | 
[**ListProxyusersNameMembers**](ProtocolsHdfsApi.md#ListProxyusersNameMembers) | **Get** /platform/1/protocols/hdfs/proxyusers/{Name}/members | 
[**UpdateProxyusersNameMember**](ProtocolsHdfsApi.md#UpdateProxyusersNameMember) | **Put** /platform/1/protocols/hdfs/proxyusers/{Name}/members/{ProxyusersNameMemberId} | 


# **CreateProxyusersNameMember**
> CreateResponse CreateProxyusersNameMember(ctx, proxyusersNameMember, name, optional)


Add a member to the HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **proxyusersNameMember** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyusersNameMember** | [**AuthAccessAccessItemFileGroup**](AuthAccessAccessItemFileGroup.md)|  | 
 **name** | **string**|  | 
 **zone** | **string**| Specifies which access zone or zones to use. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProxyusersNameMember**
> DeleteProxyusersNameMember(ctx, proxyusersNameMemberId, name, optional)


Remove a member from the HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **proxyusersNameMemberId** | **string**| Remove a member from the HDFS proxyuser. | 
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyusersNameMemberId** | **string**| Remove a member from the HDFS proxyuser. | 
 **name** | **string**|  | 
 **zone** | **string**| Specifies which access zone or zones to use. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProxyusersNameMembers**
> GroupMembers ListProxyusersNameMembers(ctx, name, optional)


List all the members of the HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**|  | 
 **zone** | **string**| Specifies which access zone or zones to use. | 

### Return type

[**GroupMembers**](GroupMembers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyusersNameMember**
> UpdateProxyusersNameMember(ctx, proxyusersNameMember, proxyusersNameMemberId, name, optional)


Create a new HDFS proxyuser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **proxyusersNameMember** | [**Empty**](Empty.md)|  | 
  **proxyusersNameMemberId** | **string**| Create a new HDFS proxyuser. | 
  **name** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proxyusersNameMember** | [**Empty**](Empty.md)|  | 
 **proxyusersNameMemberId** | **string**| Create a new HDFS proxyuser. | 
 **name** | **string**|  | 
 **zone** | **string**| Specifies which access zone or zones to use. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

