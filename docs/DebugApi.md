# \DebugApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDebugStats**](DebugApi.md#DeleteDebugStats) | **Delete** /platform/1/debug/stats | 
[**GetDebugStats**](DebugApi.md#GetDebugStats) | **Get** /platform/1/debug/stats | 


# **DeleteDebugStats**
> DeleteDebugStats(ctx, )


Clear per-resource statistics counters.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDebugStats**
> DebugStats GetDebugStats(ctx, )


List cumulative call statistics for each resource.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DebugStats**](DebugStats.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

