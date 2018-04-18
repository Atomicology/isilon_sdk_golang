# \NetworkGroupnetsSubnetsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePoolsPoolRebalanceIp**](NetworkGroupnetsSubnetsApi.md#CreatePoolsPoolRebalanceIp) | **Post** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rebalance-ips | 
[**CreatePoolsPoolRule**](NetworkGroupnetsSubnetsApi.md#CreatePoolsPoolRule) | **Post** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rules | 
[**CreatePoolsPoolScResumeNode**](NetworkGroupnetsSubnetsApi.md#CreatePoolsPoolScResumeNode) | **Post** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/sc-resume-nodes | 
[**CreatePoolsPoolScSuspendNode**](NetworkGroupnetsSubnetsApi.md#CreatePoolsPoolScSuspendNode) | **Post** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/sc-suspend-nodes | 
[**DeletePoolsPoolRule**](NetworkGroupnetsSubnetsApi.md#DeletePoolsPoolRule) | **Delete** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rules/{PoolsPoolRuleId} | 
[**GetPoolsPoolInterfaces**](NetworkGroupnetsSubnetsApi.md#GetPoolsPoolInterfaces) | **Get** /platform/4/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/interfaces | 
[**GetPoolsPoolRule**](NetworkGroupnetsSubnetsApi.md#GetPoolsPoolRule) | **Get** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rules/{PoolsPoolRuleId} | 
[**ListPoolsPoolRules**](NetworkGroupnetsSubnetsApi.md#ListPoolsPoolRules) | **Get** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rules | 
[**UpdatePoolsPoolRule**](NetworkGroupnetsSubnetsApi.md#UpdatePoolsPoolRule) | **Put** /platform/3/network/groupnets/{Groupnet}/subnets/{Subnet}/pools/{Pool}/rules/{PoolsPoolRuleId} | 


# **CreatePoolsPoolRebalanceIp**
> Empty CreatePoolsPoolRebalanceIp(ctx, poolsPoolRebalanceIp, groupnet, subnet, pool)


Rebalance IP addresses in specified pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolRebalanceIp** | [**Empty**](Empty.md)|  | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePoolsPoolRule**
> CreateResponse CreatePoolsPoolRule(ctx, poolsPoolRule, groupnet, subnet, pool)


Create a new rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolRule** | [**PoolsPoolRuleCreateParams**](PoolsPoolRuleCreateParams.md)|  | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePoolsPoolScResumeNode**
> Empty CreatePoolsPoolScResumeNode(ctx, poolsPoolScResumeNode, groupnet, subnet, pool)


Resume suspended nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolScResumeNode** | [**PoolsPoolScResumeNode**](PoolsPoolScResumeNode.md)|  | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePoolsPoolScSuspendNode**
> Empty CreatePoolsPoolScSuspendNode(ctx, poolsPoolScSuspendNode, groupnet, subnet, pool)


Suspend nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolScSuspendNode** | [**PoolsPoolScResumeNode**](PoolsPoolScResumeNode.md)|  | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePoolsPoolRule**
> DeletePoolsPoolRule(ctx, poolsPoolRuleId, groupnet, subnet, pool)


Delete a network rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolRuleId** | **string**| Delete a network rule. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolsPoolInterfaces**
> PoolsPoolInterfaces GetPoolsPoolInterfaces(ctx, groupnet, subnet, pool, optional)


Get a list of interfaces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**|  | 
 **subnet** | **string**|  | 
 **pool** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **lnns** | **string**| Get a list of interfaces for the specified lnn. | 

### Return type

[**PoolsPoolInterfaces**](PoolsPoolInterfaces.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolsPoolRule**
> PoolsPoolRules GetPoolsPoolRule(ctx, poolsPoolRuleId, groupnet, subnet, pool)


View a single network rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolRuleId** | **string**| View a single network rule. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

[**PoolsPoolRules**](PoolsPoolRules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPoolsPoolRules**
> PoolsPoolRulesExtended ListPoolsPoolRules(ctx, groupnet, subnet, pool, optional)


Get a list of network rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupnet** | **string**|  | 
 **subnet** | **string**|  | 
 **pool** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**PoolsPoolRulesExtended**](PoolsPoolRulesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePoolsPoolRule**
> UpdatePoolsPoolRule(ctx, poolsPoolRule, poolsPoolRuleId, groupnet, subnet, pool)


Modify a network rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **poolsPoolRule** | [**PoolsPoolRule**](PoolsPoolRule.md)|  | 
  **poolsPoolRuleId** | **string**| Modify a network rule. | 
  **groupnet** | **string**|  | 
  **subnet** | **string**|  | 
  **pool** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

