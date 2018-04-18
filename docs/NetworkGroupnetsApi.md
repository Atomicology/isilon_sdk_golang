# \NetworkGroupnetsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupnetSubnet**](NetworkGroupnetsApi.md#CreateGroupnetSubnet) | **Post** /platform/4/network/groupnets/{Groupnet}/subnets | 
[**CreateSubnetsSubnetPool**](NetworkGroupnetsApi.md#CreateSubnetsSubnetPool) | **Post** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools | 
[**DeleteGroupnetSubnet**](NetworkGroupnetsApi.md#DeleteGroupnetSubnet) | **Delete** /platform/4/network/groupnets/{Groupnet}/subnets/{GroupnetSubnetId} | 
[**DeleteSubnetsSubnetPool**](NetworkGroupnetsApi.md#DeleteSubnetsSubnetPool) | **Delete** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{SubnetsSubnetPoolId} | 
[**GetGroupnetSubnet**](NetworkGroupnetsApi.md#GetGroupnetSubnet) | **Get** /platform/4/network/groupnets/{Groupnet}/subnets/{GroupnetSubnetId} | 
[**GetSubnetsSubnetPool**](NetworkGroupnetsApi.md#GetSubnetsSubnetPool) | **Get** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{SubnetsSubnetPoolId} | 
[**ListGroupnetSubnets**](NetworkGroupnetsApi.md#ListGroupnetSubnets) | **Get** /platform/4/network/groupnets/{Groupnet}/subnets | 
[**ListSubnetsSubnetPools**](NetworkGroupnetsApi.md#ListSubnetsSubnetPools) | **Get** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools | 
[**UpdateGroupnetSubnet**](NetworkGroupnetsApi.md#UpdateGroupnetSubnet) | **Put** /platform/4/network/groupnets/{Groupnet}/subnets/{GroupnetSubnetId} | 
[**UpdateSubnetsSubnetPool**](NetworkGroupnetsApi.md#UpdateSubnetsSubnetPool) | **Put** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{SubnetsSubnetPoolId} | 


# **CreateGroupnetSubnet**
> CreateResponse CreateGroupnetSubnet(ctx, groupnetSubnet, groupnet)


Create a new subnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnetSubnet** | [**GroupnetSubnetCreateParams**](GroupnetSubnetCreateParams.md)|  | 
  **groupnet** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubnetsSubnetPool**
> CreateResponse CreateSubnetsSubnetPool(ctx, subnetsSubnetPool, groupnet, subnet, optional)


Create a new pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subnetsSubnetPool** | [**SubnetsSubnetPoolCreateParams**](SubnetsSubnetPoolCreateParams.md)|  | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetsSubnetPool** | [**SubnetsSubnetPoolCreateParams**](SubnetsSubnetPoolCreateParams.md)|  | 
 **groupnet** | **string**|  | 
 **subnet** | **string**|  | 
 **force** | **bool**| Force creating this pool even if it causes an MTU conflict. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupnetSubnet**
> DeleteGroupnetSubnet(ctx, groupnetSubnetId, groupnet, optional)


Delete a network subnet..

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnetSubnetId** | **string**| Delete a network subnet.. | 
  **groupnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnetSubnetId** | **string**| Delete a network subnet.. | 
 **groupnet** | **string**|  | 
 **force** | **bool**| force deleting this subnet even if pools in other subnets rely on this subnet&#39;s SC VIP. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubnetsSubnetPool**
> DeleteSubnetsSubnetPool(ctx, subnetsSubnetPoolId, groupnet, subnet)


Delete a network pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subnetsSubnetPoolId** | **string**| Delete a network pool. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupnetSubnet**
> GroupnetSubnets GetGroupnetSubnet(ctx, groupnetSubnetId, groupnet)


View a network subnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnetSubnetId** | **string**| View a network subnet. | 
  **groupnet** | **string**|  | 

### Return type

[**GroupnetSubnets**](GroupnetSubnets.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubnetsSubnetPool**
> SubnetsSubnetPools GetSubnetsSubnetPool(ctx, subnetsSubnetPoolId, groupnet, subnet)


View a single network pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subnetsSubnetPoolId** | **string**| View a single network pool. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 

### Return type

[**SubnetsSubnetPools**](SubnetsSubnetPools.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGroupnetSubnets**
> GroupnetSubnetsExtended ListGroupnetSubnets(ctx, groupnet, optional)


Get a list of subnets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**GroupnetSubnetsExtended**](GroupnetSubnetsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubnetsSubnetPools**
> SubnetsSubnetPoolsExtended ListSubnetsSubnetPools(ctx, groupnet, subnet, optional)


Get a list of network pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**|  | 
 **subnet** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **accessZone** | **string**| If specified, only pools with this zone name will be returned. | 
 **allocMethod** | **string**| If specified, only pools with this allocation type will be returned. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**SubnetsSubnetPoolsExtended**](SubnetsSubnetPoolsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroupnetSubnet**
> UpdateGroupnetSubnet(ctx, groupnetSubnet, groupnetSubnetId, groupnet, optional)


Modify a network subnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnetSubnet** | [**GroupnetSubnet**](GroupnetSubnet.md)|  | 
  **groupnetSubnetId** | **string**| Modify a network subnet. | 
  **groupnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnetSubnet** | [**GroupnetSubnet**](GroupnetSubnet.md)|  | 
 **groupnetSubnetId** | **string**| Modify a network subnet. | 
 **groupnet** | **string**|  | 
 **force** | **bool**| force modifying this subnet even if it causes an MTU conflict. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubnetsSubnetPool**
> UpdateSubnetsSubnetPool(ctx, subnetsSubnetPool, subnetsSubnetPoolId, groupnet, subnet, optional)


Modify a network pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **subnetsSubnetPool** | [**SubnetsSubnetPool**](SubnetsSubnetPool.md)|  | 
  **subnetsSubnetPoolId** | **string**| Modify a network pool. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subnetsSubnetPool** | [**SubnetsSubnetPool**](SubnetsSubnetPool.md)|  | 
 **subnetsSubnetPoolId** | **string**| Modify a network pool. | 
 **groupnet** | **string**|  | 
 **subnet** | **string**|  | 
 **force** | **bool**| force creating this pool even if it causes an MTU conflict. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

