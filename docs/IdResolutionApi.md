# \IdResolutionApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdResolutionPath**](IdResolutionApi.md#GetIdResolutionPath) | **Get** /platform/4/id-resolution/paths/{IdResolutionPathId} | 
[**GetIdResolutionPaths**](IdResolutionApi.md#GetIdResolutionPaths) | **Get** /platform/4/id-resolution/paths | 


# **GetIdResolutionPath**
> IdResolutionPaths GetIdResolutionPath(ctx, idResolutionPathId)


List lin to path mappings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **idResolutionPathId** | **int32**| List lin to path mappings. | 

### Return type

[**IdResolutionPaths**](IdResolutionPaths.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdResolutionPaths**
> IdResolutionPathsExtended GetIdResolutionPaths(ctx, optional)


List lin to path mappings.

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
 **lins** | **string**| A comma separated list specifying the lins that will be mapped with a path. Only the lins specified in this list will be mapped. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**IdResolutionPathsExtended**](IdResolutionPathsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

