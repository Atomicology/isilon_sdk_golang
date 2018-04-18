# \FilesystemApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsAccessTime**](FilesystemApi.md#GetSettingsAccessTime) | **Get** /platform/1/filesystem/settings/access-time | 
[**GetSettingsCharacterEncodings**](FilesystemApi.md#GetSettingsCharacterEncodings) | **Get** /platform/1/filesystem/settings/character-encodings | 
[**UpdateSettingsAccessTime**](FilesystemApi.md#UpdateSettingsAccessTime) | **Put** /platform/1/filesystem/settings/access-time | 
[**UpdateSettingsCharacterEncodings**](FilesystemApi.md#UpdateSettingsCharacterEncodings) | **Put** /platform/1/filesystem/settings/character-encodings | 


# **GetSettingsAccessTime**
> SettingsAccessTime GetSettingsAccessTime(ctx, )


Retrieve settings for access time.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsAccessTime**](SettingsAccessTime.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsCharacterEncodings**
> SettingsCharacterEncodings GetSettingsCharacterEncodings(ctx, )


Retrieve current cluster character encoding settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsCharacterEncodings**](SettingsCharacterEncodings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsAccessTime**
> UpdateSettingsAccessTime(ctx, settingsAccessTime)


Set settings for access time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsAccessTime** | [**SettingsAccessTimeExtended**](SettingsAccessTimeExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsCharacterEncodings**
> UpdateSettingsCharacterEncodings(ctx, settingsCharacterEncodings)


Set current character encoding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsCharacterEncodings** | [**SettingsCharacterEncodingsExtended**](SettingsCharacterEncodingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

