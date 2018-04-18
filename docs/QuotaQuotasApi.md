# \QuotaQuotasApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuotaNotification**](QuotaQuotasApi.md#CreateQuotaNotification) | **Post** /platform/1/quota/quotas/{Qid}/notifications | 
[**DeleteQuotaNotification**](QuotaQuotasApi.md#DeleteQuotaNotification) | **Delete** /platform/1/quota/quotas/{Qid}/notifications/{QuotaNotificationId} | 
[**DeleteQuotaNotifications**](QuotaQuotasApi.md#DeleteQuotaNotifications) | **Delete** /platform/1/quota/quotas/{Qid}/notifications | 
[**GetQuotaNotification**](QuotaQuotasApi.md#GetQuotaNotification) | **Get** /platform/1/quota/quotas/{Qid}/notifications/{QuotaNotificationId} | 
[**ListQuotaNotifications**](QuotaQuotasApi.md#ListQuotaNotifications) | **Get** /platform/1/quota/quotas/{Qid}/notifications | 
[**UpdateQuotaNotification**](QuotaQuotasApi.md#UpdateQuotaNotification) | **Put** /platform/1/quota/quotas/{Qid}/notifications/{QuotaNotificationId} | 
[**UpdateQuotaNotifications**](QuotaQuotasApi.md#UpdateQuotaNotifications) | **Put** /platform/1/quota/quotas/{Qid}/notifications | 


# **CreateQuotaNotification**
> CreateResponse CreateQuotaNotification(ctx, quotaNotification, qid)


Create a new notification rule specific to this quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaNotification** | [**QuotaNotificationCreateParams**](QuotaNotificationCreateParams.md)|  | 
  **qid** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaNotification**
> DeleteQuotaNotification(ctx, quotaNotificationId, qid)


Delete the notification rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaNotificationId** | **string**| Delete the notification rule. | 
  **qid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaNotifications**
> DeleteQuotaNotifications(ctx, qid)


Delete all quota specific rules. The quota will then use the global rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **qid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaNotification**
> QuotaNotifications GetQuotaNotification(ctx, quotaNotificationId, qid)


Retrieve notification rule information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaNotificationId** | **string**| Retrieve notification rule information. | 
  **qid** | **string**|  | 

### Return type

[**QuotaNotifications**](QuotaNotifications.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaNotifications**
> QuotaNotificationsExtended ListQuotaNotifications(ctx, qid)


List all rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **qid** | **string**|  | 

### Return type

[**QuotaNotificationsExtended**](QuotaNotificationsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuotaNotification**
> UpdateQuotaNotification(ctx, quotaNotification, quotaNotificationId, qid)


Modify notification rule. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaNotification** | [**QuotaNotification**](QuotaNotification.md)|  | 
  **quotaNotificationId** | **string**| Modify notification rule. All input fields are optional, but one or more must be supplied. | 
  **qid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuotaNotifications**
> UpdateQuotaNotifications(ctx, quotaNotifications, qid)


This method creates an empty set of rules so that the global rules are not used. The input must be an empty JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaNotifications** | [**Empty**](Empty.md)|  | 
  **qid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

