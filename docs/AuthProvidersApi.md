# \AuthProvidersApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdsProviderControllers**](AuthProvidersApi.md#GetAdsProviderControllers) | **Get** /platform/3/auth/providers/ads/{Id}/controllers | 
[**GetAdsProviderDomain**](AuthProvidersApi.md#GetAdsProviderDomain) | **Get** /platform/3/auth/providers/ads/{Id}/domains/{AdsProviderDomainId} | 
[**GetAdsProviderDomains**](AuthProvidersApi.md#GetAdsProviderDomains) | **Get** /platform/3/auth/providers/ads/{Id}/domains | 
[**GetAdsProviderSearch**](AuthProvidersApi.md#GetAdsProviderSearch) | **Get** /platform/1/auth/providers/ads/{Id}/search | 


# **GetAdsProviderControllers**
> AdsProviderControllers GetAdsProviderControllers(ctx, id, optional)


List all ADS controllers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **groupnet** | **string**| Groupnet identifier | 

### Return type

[**AdsProviderControllers**](AdsProviderControllers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdsProviderDomain**
> AdsProviderDomains GetAdsProviderDomain(ctx, adsProviderDomainId, id)


Retrieve the ADS domain information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **adsProviderDomainId** | **string**| Retrieve the ADS domain information. | 
  **id** | **string**|  | 

### Return type

[**AdsProviderDomains**](AdsProviderDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdsProviderDomains**
> AdsProviderDomains GetAdsProviderDomains(ctx, id, optional)


List all ADS domains.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**AdsProviderDomains**](AdsProviderDomains.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdsProviderSearch**
> AdsProviderSearch GetAdsProviderSearch(ctx, id, optional)


Retrieve search results.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **domain** | **string**| The domain to search in. | 
 **description** | **string**| The user or group description to search for. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **searchUsers** | **bool**| If true, search for users. | 
 **filter** | **string**| The LDAP filter to apply to the search. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **user** | **string**| The user name for the domain if untrusted. | 
 **password** | **string**| The password for the domain if untrusted. | 
 **searchGroups** | **bool**| If true, search for groups. | 

### Return type

[**AdsProviderSearch**](AdsProviderSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

