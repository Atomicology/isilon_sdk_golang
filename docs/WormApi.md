# \WormApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWormDomain**](WormApi.md#CreateWormDomain) | **Post** /platform/1/worm/domains | 
[**GetWormDomain**](WormApi.md#GetWormDomain) | **Get** /platform/1/worm/domains/{WormDomainId} | 
[**GetWormSettings**](WormApi.md#GetWormSettings) | **Get** /platform/1/worm/settings | 
[**ListWormDomains**](WormApi.md#ListWormDomains) | **Get** /platform/1/worm/domains | 
[**UpdateWormDomain**](WormApi.md#UpdateWormDomain) | **Put** /platform/1/worm/domains/{WormDomainId} | 
[**UpdateWormSettings**](WormApi.md#UpdateWormSettings) | **Put** /platform/1/worm/settings | 


# **CreateWormDomain**
> WormDomainExtended CreateWormDomain(ctx, wormDomain)


Create a WORM domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **wormDomain** | [**WormDomainCreateParams**](WormDomainCreateParams.md)|  | 

### Return type

[**WormDomainExtended**](WormDomainExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWormDomain**
> WormDomains GetWormDomain(ctx, wormDomainId)


View a single WORM domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **wormDomainId** | **string**| View a single WORM domain. | 

### Return type

[**WormDomains**](WormDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWormSettings**
> WormSettings GetWormSettings(ctx, )


Get the global WORM settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WormSettings**](WormSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWormDomains**
> WormDomainsExtended ListWormDomains(ctx, optional)


List all WORM domains.

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

[**WormDomainsExtended**](WormDomainsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWormDomain**
> UpdateWormDomain(ctx, wormDomain, wormDomainId)


Modify a single WORM domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **wormDomain** | [**WormDomain**](WormDomain.md)|  | 
  **wormDomainId** | **string**| Modify a single WORM domain. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWormSettings**
> UpdateWormSettings(ctx, wormSettings)


Modify the global WORM settings.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **wormSettings** | [**WormSettingsExtended**](WormSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

