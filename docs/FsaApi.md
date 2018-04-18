# \FsaApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFsaResult**](FsaApi.md#DeleteFsaResult) | **Delete** /platform/3/fsa/results/{FsaResultId} | 
[**DeleteFsaSettings**](FsaApi.md#DeleteFsaSettings) | **Delete** /platform/1/fsa/settings | 
[**GetFsaResult**](FsaApi.md#GetFsaResult) | **Get** /platform/3/fsa/results/{FsaResultId} | 
[**GetFsaResults**](FsaApi.md#GetFsaResults) | **Get** /platform/3/fsa/results | 
[**GetFsaSettings**](FsaApi.md#GetFsaSettings) | **Get** /platform/1/fsa/settings | 
[**UpdateFsaResult**](FsaApi.md#UpdateFsaResult) | **Put** /platform/3/fsa/results/{FsaResultId} | 
[**UpdateFsaSettings**](FsaApi.md#UpdateFsaSettings) | **Put** /platform/1/fsa/settings | 


# **DeleteFsaResult**
> DeleteFsaResult(ctx, fsaResultId)


Delete the result set.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsaResultId** | **string**| Delete the result set. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFsaSettings**
> DeleteFsaSettings(ctx, )


Revert all settings to their defaults.

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

# **GetFsaResult**
> FsaResults GetFsaResult(ctx, fsaResultId)


Retrieve result set information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsaResultId** | **string**| Retrieve result set information. | 

### Return type

[**FsaResults**](FsaResults.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFsaResults**
> FsaResultsExtended GetFsaResults(ctx, )


List all results.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FsaResultsExtended**](FsaResultsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFsaSettings**
> FsaSettings GetFsaSettings(ctx, optional)


List all settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string**| If specified as effective or not specified, all fields are returned.  If specified as user, only fields with non-default values are shown.  If specified as default, the original values are returned. | 

### Return type

[**FsaSettings**](FsaSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFsaResult**
> UpdateFsaResult(ctx, fsaResult, fsaResultId)


Modify result set. Only the pinned property can be changed at this time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsaResult** | [**FsaResult**](FsaResult.md)|  | 
  **fsaResultId** | **string**| Modify result set. Only the pinned property can be changed at this time. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFsaSettings**
> UpdateFsaSettings(ctx, fsaSettings)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fsaSettings** | [**FsaSettingsSettings**](FsaSettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

