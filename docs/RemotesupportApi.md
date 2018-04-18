# \RemotesupportApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemotesupportConnectemc**](RemotesupportApi.md#GetRemotesupportConnectemc) | **Get** /platform/1/remotesupport/connectemc | 
[**UpdateRemotesupportConnectemc**](RemotesupportApi.md#UpdateRemotesupportConnectemc) | **Put** /platform/1/remotesupport/connectemc | 


# **GetRemotesupportConnectemc**
> RemotesupportConnectemc GetRemotesupportConnectemc(ctx, )


List all settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RemotesupportConnectemc**](RemotesupportConnectemc.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemotesupportConnectemc**
> UpdateRemotesupportConnectemc(ctx, remotesupportConnectemc)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **remotesupportConnectemc** | [**RemotesupportConnectemcConnectemc**](RemotesupportConnectemcConnectemc.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

