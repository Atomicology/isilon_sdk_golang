# \AuditApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuditTopic**](AuditApi.md#CreateAuditTopic) | **Post** /platform/1/audit/topics | 
[**DeleteAuditTopic**](AuditApi.md#DeleteAuditTopic) | **Delete** /platform/1/audit/topics/{AuditTopicId} | 
[**GetAuditProgress**](AuditApi.md#GetAuditProgress) | **Get** /platform/4/audit/progress | 
[**GetAuditSettings**](AuditApi.md#GetAuditSettings) | **Get** /platform/3/audit/settings | 
[**GetAuditTopic**](AuditApi.md#GetAuditTopic) | **Get** /platform/1/audit/topics/{AuditTopicId} | 
[**GetProgressGlobal**](AuditApi.md#GetProgressGlobal) | **Get** /platform/4/audit/progress/global | 
[**GetSettingsGlobal**](AuditApi.md#GetSettingsGlobal) | **Get** /platform/3/audit/settings/global | 
[**ListAuditTopics**](AuditApi.md#ListAuditTopics) | **Get** /platform/1/audit/topics | 
[**UpdateAuditSettings**](AuditApi.md#UpdateAuditSettings) | **Put** /platform/3/audit/settings | 
[**UpdateAuditTopic**](AuditApi.md#UpdateAuditTopic) | **Put** /platform/1/audit/topics/{AuditTopicId} | 
[**UpdateSettingsGlobal**](AuditApi.md#UpdateSettingsGlobal) | **Put** /platform/3/audit/settings/global | 


# **CreateAuditTopic**
> CreateResponse CreateAuditTopic(ctx, auditTopic)


Create a new audit topic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **auditTopic** | [**AuditTopicCreateParams**](AuditTopicCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuditTopic**
> DeleteAuditTopic(ctx, auditTopicId)


Delete the audit topic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **auditTopicId** | **string**| Delete the audit topic. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditProgress**
> AuditProgress GetAuditProgress(ctx, optional)


View current audit log time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lnn** | **int32**| lnn of the node. | 

### Return type

[**AuditProgress**](AuditProgress.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditSettings**
> AuditSettings GetAuditSettings(ctx, optional)


View per-Access Zone Audit settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | **string**| Access zone which contains audit settings. | 

### Return type

[**AuditSettings**](AuditSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuditTopic**
> AuditTopics GetAuditTopic(ctx, auditTopicId)


Retrieve the audit topic information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **auditTopicId** | **string**| Retrieve the audit topic information. | 

### Return type

[**AuditTopics**](AuditTopics.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProgressGlobal**
> ProgressGlobal GetProgressGlobal(ctx, )


View the global audit log time.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ProgressGlobal**](ProgressGlobal.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsGlobal**
> SettingsGlobalExtended GetSettingsGlobal(ctx, )


View Global Audit settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsGlobalExtended**](SettingsGlobalExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAuditTopics**
> AuditTopicsExtended ListAuditTopics(ctx, )


Retrieve a list of audit topics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AuditTopicsExtended**](AuditTopicsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuditSettings**
> UpdateAuditSettings(ctx, auditSettings, optional)


Modify per-Access Zone Audit settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **auditSettings** | [**AuditSettingsSettings**](AuditSettingsSettings.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auditSettings** | [**AuditSettingsSettings**](AuditSettingsSettings.md)|  | 
 **zone** | **string**| Access zone which contains audit settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAuditTopic**
> UpdateAuditTopic(ctx, auditTopic, auditTopicId)


Modify the audit topic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **auditTopic** | [**AuditTopic**](AuditTopic.md)|  | 
  **auditTopicId** | **string**| Modify the audit topic. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsGlobal**
> UpdateSettingsGlobal(ctx, settingsGlobal)


Modify Global Audit settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsGlobal** | [**SettingsGlobalSettings**](SettingsGlobalSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

