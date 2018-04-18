# \HardeningApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHardeningApplyItem**](HardeningApi.md#CreateHardeningApplyItem) | **Post** /platform/3/hardening/apply | 
[**CreateHardeningResolveItem**](HardeningApi.md#CreateHardeningResolveItem) | **Post** /platform/3/hardening/resolve | 
[**CreateHardeningRevertItem**](HardeningApi.md#CreateHardeningRevertItem) | **Post** /platform/3/hardening/revert | 
[**GetHardeningState**](HardeningApi.md#GetHardeningState) | **Get** /platform/3/hardening/state | 
[**GetHardeningStatus**](HardeningApi.md#GetHardeningStatus) | **Get** /platform/3/hardening/status | 


# **CreateHardeningApplyItem**
> CreateHardeningApplyItemResponse CreateHardeningApplyItem(ctx, hardeningApplyItem)


Apply hardening on the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardeningApplyItem** | [**HardeningApplyItem**](HardeningApplyItem.md)|  | 

### Return type

[**CreateHardeningApplyItemResponse**](CreateHardeningApplyItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHardeningResolveItem**
> CreateHardeningResolveItemResponse CreateHardeningResolveItem(ctx, hardeningResolveItem, optional)


Resolve issues related to hardening, found in current cluster configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardeningResolveItem** | [**HardeningResolveItem**](HardeningResolveItem.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardeningResolveItem** | [**HardeningResolveItem**](HardeningResolveItem.md)|  | 
 **accept** | **bool**| If true, execution proceeds to resolve all issues. If false, executrion aborts. This is a required argument. | 

### Return type

[**CreateHardeningResolveItemResponse**](CreateHardeningResolveItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHardeningRevertItem**
> CreateHardeningRevertItemResponse CreateHardeningRevertItem(ctx, hardeningRevertItem, optional)


Revert hardening on the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardeningRevertItem** | [**Empty**](Empty.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hardeningRevertItem** | [**Empty**](Empty.md)|  | 
 **force** | **bool**| If specified, revert operation continues even in case of a failure. Default is false in which case revert stops at the first failure. | 

### Return type

[**CreateHardeningRevertItemResponse**](CreateHardeningRevertItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardeningState**
> HardeningState GetHardeningState(ctx, )


Get the state of the current hardening operation, if one is happening.  Note that this is different from the /status resource, which returns the overall hardening status of the cluster.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HardeningState**](HardeningState.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardeningStatus**
> HardeningStatus GetHardeningStatus(ctx, )


Get a message indicating whether or not the cluster is hardened. Note that this is different from the /state resource, which returns the state of a specific hardening operation (apply or revert).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HardeningStatus**](HardeningStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

