# \NetworkApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDnscacheFlushItem**](NetworkApi.md#CreateDnscacheFlushItem) | **Post** /platform/3/network/dnscache/flush | 
[**CreateNetworkGroupnet**](NetworkApi.md#CreateNetworkGroupnet) | **Post** /platform/3/network/groupnets | 
[**CreateNetworkScRebalanceAllItem**](NetworkApi.md#CreateNetworkScRebalanceAllItem) | **Post** /platform/3/network/sc-rebalance-all | 
[**DeleteNetworkGroupnet**](NetworkApi.md#DeleteNetworkGroupnet) | **Delete** /platform/3/network/groupnets/{NetworkGroupnetId} | 
[**GetNetworkDnscache**](NetworkApi.md#GetNetworkDnscache) | **Get** /platform/3/network/dnscache | 
[**GetNetworkExternal**](NetworkApi.md#GetNetworkExternal) | **Get** /platform/3/network/external | 
[**GetNetworkGroupnet**](NetworkApi.md#GetNetworkGroupnet) | **Get** /platform/3/network/groupnets/{NetworkGroupnetId} | 
[**GetNetworkInterfaces**](NetworkApi.md#GetNetworkInterfaces) | **Get** /platform/4/network/interfaces | 
[**GetNetworkPools**](NetworkApi.md#GetNetworkPools) | **Get** /platform/3/network/pools | 
[**GetNetworkRules**](NetworkApi.md#GetNetworkRules) | **Get** /platform/3/network/rules | 
[**GetNetworkSubnets**](NetworkApi.md#GetNetworkSubnets) | **Get** /platform/4/network/subnets | 
[**ListNetworkGroupnets**](NetworkApi.md#ListNetworkGroupnets) | **Get** /platform/3/network/groupnets | 
[**UpdateNetworkDnscache**](NetworkApi.md#UpdateNetworkDnscache) | **Put** /platform/3/network/dnscache | 
[**UpdateNetworkExternal**](NetworkApi.md#UpdateNetworkExternal) | **Put** /platform/3/network/external | 
[**UpdateNetworkGroupnet**](NetworkApi.md#UpdateNetworkGroupnet) | **Put** /platform/3/network/groupnets/{NetworkGroupnetId} | 


# **CreateDnscacheFlushItem**
> Empty CreateDnscacheFlushItem(ctx, dnscacheFlushItem)


Flush the DNSCache.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dnscacheFlushItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNetworkGroupnet**
> CreateResponse CreateNetworkGroupnet(ctx, networkGroupnet)


Create a new groupnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkGroupnet** | [**NetworkGroupnetCreateParams**](NetworkGroupnetCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNetworkScRebalanceAllItem**
> Empty CreateNetworkScRebalanceAllItem(ctx, networkScRebalanceAllItem)


Rebalance IP addresses in all pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkScRebalanceAllItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNetworkGroupnet**
> DeleteNetworkGroupnet(ctx, networkGroupnetId)


Delete a network groupnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkGroupnetId** | **string**| Delete a network groupnet. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkDnscache**
> NetworkDnscache GetNetworkDnscache(ctx, )


View network dns cache settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NetworkDnscache**](NetworkDnscache.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkExternal**
> NetworkExternal GetNetworkExternal(ctx, )


View external network settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NetworkExternal**](NetworkExternal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkGroupnet**
> NetworkGroupnets GetNetworkGroupnet(ctx, networkGroupnetId)


View a network groupnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkGroupnetId** | **string**| View a network groupnet. | 

### Return type

[**NetworkGroupnets**](NetworkGroupnets.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkInterfaces**
> PoolsPoolInterfaces GetNetworkInterfaces(ctx, optional)


Get a list of interfaces.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **network** | **string**| Show interfaces associated with external and/or internal networks. Default is &#39;external&#39; | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **lnns** | **string**| Get a list of interfaces for the specified lnn. | 
 **allocMethod** | **string**| Filter addresses and owners by pool address allocation method. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**PoolsPoolInterfaces**](PoolsPoolInterfaces.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkPools**
> NetworkPools GetNetworkPools(ctx, optional)


Get a list of flexnet pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **subnet** | **string**| If specified, only pools for this subnet will be returned. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **accessZone** | **string**| If specified, only pools with this zone name will be returned. | 
 **allocMethod** | **string**| If specified, only pools with this allocation type will be returned. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **groupnet** | **string**| If specified, only pools for this groupnet will be returned. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**NetworkPools**](NetworkPools.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkRules**
> PoolsPoolRulesExtended GetNetworkRules(ctx, optional)


Get a list of network rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **subnet** | **string**| Name of the subnet to list rules from. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **groupnet** | **string**| Name of the groupnet to list rules from. | 
 **pool** | **string**| Name of the pool to list rules from. | 

### Return type

[**PoolsPoolRulesExtended**](PoolsPoolRulesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNetworkSubnets**
> GroupnetSubnetsExtended GetNetworkSubnets(ctx, optional)


Get a list of subnets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **groupnet** | **string**| If specified, only subnets for this groupnet will be returned. | 
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

# **ListNetworkGroupnets**
> NetworkGroupnetsExtended ListNetworkGroupnets(ctx, optional)


Get a list of groupnets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**NetworkGroupnetsExtended**](NetworkGroupnetsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkDnscache**
> UpdateNetworkDnscache(ctx, networkDnscache)


Modify network dns cache settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkDnscache** | [**NetworkDnscacheExtended**](NetworkDnscacheExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkExternal**
> UpdateNetworkExternal(ctx, networkExternal)


Modify external network settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkExternal** | [**NetworkExternalExtended**](NetworkExternalExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNetworkGroupnet**
> UpdateNetworkGroupnet(ctx, networkGroupnet, networkGroupnetId)


Modify a network groupnet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **networkGroupnet** | [**NetworkGroupnet**](NetworkGroupnet.md)|  | 
  **networkGroupnetId** | **string**| Modify a network groupnet. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

