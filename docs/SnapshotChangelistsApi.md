# \SnapshotChangelistsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChangelistLin**](SnapshotChangelistsApi.md#GetChangelistLin) | **Get** /platform/1/snapshot/changelists/{Changelist}/lins/{ChangelistLinId} | 
[**GetChangelistLins**](SnapshotChangelistsApi.md#GetChangelistLins) | **Get** /platform/1/snapshot/changelists/{Changelist}/lins | 


# **GetChangelistLin**
> ChangelistLins GetChangelistLin(ctx, changelistLinId, changelist, optional)


Get a single entry from the changelist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changelistLinId** | **int32**| Get a single entry from the changelist. | 
  **changelist** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changelistLinId** | **int32**| Get a single entry from the changelist. | 
 **changelist** | **string**|  | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**ChangelistLins**](ChangelistLins.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChangelistLins**
> ChangelistLinsExtended GetChangelistLins(ctx, changelist, optional)


Get entries from a changelist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **changelist** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changelist** | **string**|  | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**ChangelistLinsExtended**](ChangelistLinsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

