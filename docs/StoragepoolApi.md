# \StoragepoolApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompatibilitiesClassActiveItem**](StoragepoolApi.md#CreateCompatibilitiesClassActiveItem) | **Post** /platform/1/storagepool/compatibilities/class/active | 
[**CreateCompatibilitiesSsdActiveItem**](StoragepoolApi.md#CreateCompatibilitiesSsdActiveItem) | **Post** /platform/3/storagepool/compatibilities/ssd/active | 
[**CreateStoragepoolNodepool**](StoragepoolApi.md#CreateStoragepoolNodepool) | **Post** /platform/3/storagepool/nodepools | 
[**CreateStoragepoolTier**](StoragepoolApi.md#CreateStoragepoolTier) | **Post** /platform/1/storagepool/tiers | 
[**DeleteCompatibilitiesClassActiveById**](StoragepoolApi.md#DeleteCompatibilitiesClassActiveById) | **Delete** /platform/1/storagepool/compatibilities/class/active/{CompatibilitiesClassActiveId} | 
[**DeleteCompatibilitiesSsdActiveById**](StoragepoolApi.md#DeleteCompatibilitiesSsdActiveById) | **Delete** /platform/3/storagepool/compatibilities/ssd/active/{CompatibilitiesSsdActiveId} | 
[**DeleteStoragepoolNodepool**](StoragepoolApi.md#DeleteStoragepoolNodepool) | **Delete** /platform/3/storagepool/nodepools/{StoragepoolNodepoolId} | 
[**DeleteStoragepoolNodepools**](StoragepoolApi.md#DeleteStoragepoolNodepools) | **Delete** /platform/3/storagepool/nodepools | 
[**DeleteStoragepoolTier**](StoragepoolApi.md#DeleteStoragepoolTier) | **Delete** /platform/1/storagepool/tiers/{StoragepoolTierId} | 
[**DeleteStoragepoolTiers**](StoragepoolApi.md#DeleteStoragepoolTiers) | **Delete** /platform/1/storagepool/tiers | 
[**GetCompatibilitiesClassActiveById**](StoragepoolApi.md#GetCompatibilitiesClassActiveById) | **Get** /platform/1/storagepool/compatibilities/class/active/{CompatibilitiesClassActiveId} | 
[**GetCompatibilitiesClassAvailable**](StoragepoolApi.md#GetCompatibilitiesClassAvailable) | **Get** /platform/1/storagepool/compatibilities/class/available | 
[**GetCompatibilitiesSsdActiveById**](StoragepoolApi.md#GetCompatibilitiesSsdActiveById) | **Get** /platform/3/storagepool/compatibilities/ssd/active/{CompatibilitiesSsdActiveId} | 
[**GetCompatibilitiesSsdAvailable**](StoragepoolApi.md#GetCompatibilitiesSsdAvailable) | **Get** /platform/1/storagepool/compatibilities/ssd/available | 
[**GetStoragepoolNodepool**](StoragepoolApi.md#GetStoragepoolNodepool) | **Get** /platform/3/storagepool/nodepools/{StoragepoolNodepoolId} | 
[**GetStoragepoolSettings**](StoragepoolApi.md#GetStoragepoolSettings) | **Get** /platform/5/storagepool/settings | 
[**GetStoragepoolStatus**](StoragepoolApi.md#GetStoragepoolStatus) | **Get** /platform/1/storagepool/status | 
[**GetStoragepoolStoragepools**](StoragepoolApi.md#GetStoragepoolStoragepools) | **Get** /platform/3/storagepool/storagepools | 
[**GetStoragepoolSuggestedProtectionNid**](StoragepoolApi.md#GetStoragepoolSuggestedProtectionNid) | **Get** /platform/3/storagepool/suggested-protection/{StoragepoolSuggestedProtectionNid} | 
[**GetStoragepoolTier**](StoragepoolApi.md#GetStoragepoolTier) | **Get** /platform/1/storagepool/tiers/{StoragepoolTierId} | 
[**GetStoragepoolUnprovisioned**](StoragepoolApi.md#GetStoragepoolUnprovisioned) | **Get** /platform/1/storagepool/unprovisioned | 
[**ListCompatibilitiesClassActive**](StoragepoolApi.md#ListCompatibilitiesClassActive) | **Get** /platform/1/storagepool/compatibilities/class/active | 
[**ListCompatibilitiesSsdActive**](StoragepoolApi.md#ListCompatibilitiesSsdActive) | **Get** /platform/3/storagepool/compatibilities/ssd/active | 
[**ListStoragepoolNodepools**](StoragepoolApi.md#ListStoragepoolNodepools) | **Get** /platform/3/storagepool/nodepools | 
[**ListStoragepoolTiers**](StoragepoolApi.md#ListStoragepoolTiers) | **Get** /platform/1/storagepool/tiers | 
[**UpdateCompatibilitiesSsdActiveById**](StoragepoolApi.md#UpdateCompatibilitiesSsdActiveById) | **Put** /platform/3/storagepool/compatibilities/ssd/active/{CompatibilitiesSsdActiveId} | 
[**UpdateStoragepoolNodepool**](StoragepoolApi.md#UpdateStoragepoolNodepool) | **Put** /platform/3/storagepool/nodepools/{StoragepoolNodepoolId} | 
[**UpdateStoragepoolSettings**](StoragepoolApi.md#UpdateStoragepoolSettings) | **Put** /platform/5/storagepool/settings | 
[**UpdateStoragepoolTier**](StoragepoolApi.md#UpdateStoragepoolTier) | **Put** /platform/1/storagepool/tiers/{StoragepoolTierId} | 


# **CreateCompatibilitiesClassActiveItem**
> CreateCompatibilitiesClassActiveItemResponse CreateCompatibilitiesClassActiveItem(ctx, compatibilitiesClassActiveItem, optional)


Create a new compatibility

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesClassActiveItem** | [**CompatibilitiesClassActiveItem**](CompatibilitiesClassActiveItem.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibilitiesClassActiveItem** | [**CompatibilitiesClassActiveItem**](CompatibilitiesClassActiveItem.md)|  | 
 **assess** | **bool**| Do not perform action, only test that it is possible. | 

### Return type

[**CreateCompatibilitiesClassActiveItemResponse**](CreateCompatibilitiesClassActiveItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCompatibilitiesSsdActiveItem**
> CreateCompatibilitiesClassActiveItemResponse CreateCompatibilitiesSsdActiveItem(ctx, compatibilitiesSsdActiveItem, optional)


Create a new ssd compatibility

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesSsdActiveItem** | [**CompatibilitiesSsdActiveItem**](CompatibilitiesSsdActiveItem.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibilitiesSsdActiveItem** | [**CompatibilitiesSsdActiveItem**](CompatibilitiesSsdActiveItem.md)|  | 
 **assess** | **bool**| Do not perform action, only test that it is possible. | 

### Return type

[**CreateCompatibilitiesClassActiveItemResponse**](CreateCompatibilitiesClassActiveItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoragepoolNodepool**
> CreateStoragepoolTierResponse CreateStoragepoolNodepool(ctx, storagepoolNodepool)


Create a new node pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolNodepool** | [**StoragepoolNodepoolCreateParams**](StoragepoolNodepoolCreateParams.md)|  | 

### Return type

[**CreateStoragepoolTierResponse**](CreateStoragepoolTierResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStoragepoolTier**
> CreateStoragepoolTierResponse CreateStoragepoolTier(ctx, storagepoolTier)


Create a new tier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolTier** | [**StoragepoolTierCreateParams**](StoragepoolTierCreateParams.md)|  | 

### Return type

[**CreateStoragepoolTierResponse**](CreateStoragepoolTierResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCompatibilitiesClassActiveById**
> DeleteCompatibilitiesClassActiveById(ctx, compatibilitiesClassActiveId, optional)


Delete an active compatibility by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesClassActiveId** | **string**| Delete an active compatibility by id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibilitiesClassActiveId** | **string**| Delete an active compatibility by id | 
 **assess** | **bool**| Do not perform action, only test that it is possible. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCompatibilitiesSsdActiveById**
> DeleteCompatibilitiesSsdActiveById(ctx, compatibilitiesSsdActiveId, optional)


Delete an active ssd compatibility by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesSsdActiveId** | **string**| Delete an active ssd compatibility by id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibilitiesSsdActiveId** | **string**| Delete an active ssd compatibility by id | 
 **assess** | **bool**| Do not perform action, only test that it is possible. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoragepoolNodepool**
> DeleteStoragepoolNodepool(ctx, storagepoolNodepoolId)


Delete node pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolNodepoolId** | **string**| Delete node pool. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoragepoolNodepools**
> DeleteStoragepoolNodepools(ctx, )


Delete all node pools.

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

# **DeleteStoragepoolTier**
> DeleteStoragepoolTier(ctx, storagepoolTierId)


Delete tier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolTierId** | **string**| Delete tier. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStoragepoolTiers**
> DeleteStoragepoolTiers(ctx, )


Delete all tiers.

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

# **GetCompatibilitiesClassActiveById**
> CompatibilitiesClassActive GetCompatibilitiesClassActiveById(ctx, compatibilitiesClassActiveId)


Get an active compatibilities by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesClassActiveId** | **string**| Get an active compatibilities by id | 

### Return type

[**CompatibilitiesClassActive**](CompatibilitiesClassActive.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompatibilitiesClassAvailable**
> CompatibilitiesClassAvailable GetCompatibilitiesClassAvailable(ctx, )


Get a list of available compatibilities

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CompatibilitiesClassAvailable**](CompatibilitiesClassAvailable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompatibilitiesSsdActiveById**
> CompatibilitiesSsdActive GetCompatibilitiesSsdActiveById(ctx, compatibilitiesSsdActiveId)


Get a active ssd compatibilities by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesSsdActiveId** | **string**| Get a active ssd compatibilities by id | 

### Return type

[**CompatibilitiesSsdActive**](CompatibilitiesSsdActive.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCompatibilitiesSsdAvailable**
> CompatibilitiesSsdAvailable GetCompatibilitiesSsdAvailable(ctx, )


Get a list of available ssd compatibilities

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CompatibilitiesSsdAvailable**](CompatibilitiesSsdAvailable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolNodepool**
> StoragepoolNodepools GetStoragepoolNodepool(ctx, storagepoolNodepoolId)


Retrieve node pool information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolNodepoolId** | **string**| Retrieve node pool information. | 

### Return type

[**StoragepoolNodepools**](StoragepoolNodepools.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolSettings**
> StoragepoolSettings GetStoragepoolSettings(ctx, )


List all storagepool settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StoragepoolSettings**](StoragepoolSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolStatus**
> StoragepoolStatus GetStoragepoolStatus(ctx, )


List any health conditions detected.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StoragepoolStatus**](StoragepoolStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolStoragepools**
> StoragepoolStoragepools GetStoragepoolStoragepools(ctx, optional)


List all storage pools.

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
 **toplevels** | **string**| If true, node pools contained within tiers will be filtered out of results. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**StoragepoolStoragepools**](StoragepoolStoragepools.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolSuggestedProtectionNid**
> StoragepoolSuggestedProtection GetStoragepoolSuggestedProtectionNid(ctx, storagepoolSuggestedProtectionNid)


Retrieve the suggested protection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolSuggestedProtectionNid** | **string**| Retrieve the suggested protection policy. | 

### Return type

[**StoragepoolSuggestedProtection**](StoragepoolSuggestedProtection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolTier**
> StoragepoolTiers GetStoragepoolTier(ctx, storagepoolTierId)


Retrieve tier information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolTierId** | **string**| Retrieve tier information. | 

### Return type

[**StoragepoolTiers**](StoragepoolTiers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoragepoolUnprovisioned**
> StoragepoolUnprovisioned GetStoragepoolUnprovisioned(ctx, )


Get the unprovisioned nodes and drives

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StoragepoolUnprovisioned**](StoragepoolUnprovisioned.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompatibilitiesClassActive**
> CompatibilitiesClassActiveExtended ListCompatibilitiesClassActive(ctx, )


Get a list of active compatibilities

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CompatibilitiesClassActiveExtended**](CompatibilitiesClassActiveExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCompatibilitiesSsdActive**
> CompatibilitiesSsdActiveExtended ListCompatibilitiesSsdActive(ctx, )


Get a list of active ssd compatibilities

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CompatibilitiesSsdActiveExtended**](CompatibilitiesSsdActiveExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStoragepoolNodepools**
> StoragepoolNodepoolsExtended ListStoragepoolNodepools(ctx, )


List all node pools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StoragepoolNodepoolsExtended**](StoragepoolNodepoolsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStoragepoolTiers**
> StoragepoolTiersExtended ListStoragepoolTiers(ctx, )


List all tiers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StoragepoolTiersExtended**](StoragepoolTiersExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCompatibilitiesSsdActiveById**
> UpdateCompatibilitiesSsdActiveById(ctx, compatibilitiesSsdActiveIdParams, compatibilitiesSsdActiveId, optional)


Modify an ssd compatibility by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **compatibilitiesSsdActiveIdParams** | [**CompatibilitiesSsdActiveIdParams**](CompatibilitiesSsdActiveIdParams.md)|  | 
  **compatibilitiesSsdActiveId** | **string**| Modify an ssd compatibility by id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibilitiesSsdActiveIdParams** | [**CompatibilitiesSsdActiveIdParams**](CompatibilitiesSsdActiveIdParams.md)|  | 
 **compatibilitiesSsdActiveId** | **string**| Modify an ssd compatibility by id | 
 **assess** | **bool**| Do not perform action, only test that it is possible. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoragepoolNodepool**
> UpdateStoragepoolNodepool(ctx, storagepoolNodepool, storagepoolNodepoolId)


Modify node pool. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolNodepool** | [**StoragepoolNodepool**](StoragepoolNodepool.md)|  | 
  **storagepoolNodepoolId** | **string**| Modify node pool. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoragepoolSettings**
> UpdateStoragepoolSettings(ctx, storagepoolSettings)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolSettings** | [**StoragepoolSettingsExtended**](StoragepoolSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStoragepoolTier**
> UpdateStoragepoolTier(ctx, storagepoolTier, storagepoolTierId)


Modify tier. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **storagepoolTier** | [**StoragepoolTier**](StoragepoolTier.md)|  | 
  **storagepoolTierId** | **string**| Modify tier. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

