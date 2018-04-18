# \FileFilterApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileFilterSettings**](FileFilterApi.md#GetFileFilterSettings) | **Get** /platform/3/file-filter/settings | 
[**UpdateFileFilterSettings**](FileFilterApi.md#UpdateFileFilterSettings) | **Put** /platform/3/file-filter/settings | 


# **GetFileFilterSettings**
> FileFilterSettings GetFileFilterSettings(ctx, optional)


View File Filtering settings of an access zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Specifies the access zones in which these settings apply. | 

### Return type

[**FileFilterSettings**](FileFilterSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFileFilterSettings**
> UpdateFileFilterSettings(ctx, fileFilterSettings, optional)


Modify one or more File Filtering settings for an access zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **fileFilterSettings** | [**FileFilterSettingsExtended**](FileFilterSettingsExtended.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileFilterSettings** | [**FileFilterSettingsExtended**](FileFilterSettingsExtended.md)|  | 
 **zone** | **string**| Specifies the access zones in which these settings apply. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

