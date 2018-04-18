# \QuotaApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuotaQuota**](QuotaApi.md#CreateQuotaQuota) | **Post** /platform/1/quota/quotas | 
[**CreateQuotaReport**](QuotaApi.md#CreateQuotaReport) | **Post** /platform/1/quota/reports | 
[**CreateSettingsMapping**](QuotaApi.md#CreateSettingsMapping) | **Post** /platform/1/quota/settings/mappings | 
[**CreateSettingsNotification**](QuotaApi.md#CreateSettingsNotification) | **Post** /platform/1/quota/settings/notifications | 
[**DeleteQuotaQuota**](QuotaApi.md#DeleteQuotaQuota) | **Delete** /platform/1/quota/quotas/{QuotaQuotaId} | 
[**DeleteQuotaQuotas**](QuotaApi.md#DeleteQuotaQuotas) | **Delete** /platform/1/quota/quotas | 
[**DeleteQuotaReport**](QuotaApi.md#DeleteQuotaReport) | **Delete** /platform/1/quota/reports/{QuotaReportId} | 
[**DeleteSettingsMapping**](QuotaApi.md#DeleteSettingsMapping) | **Delete** /platform/1/quota/settings/mappings/{SettingsMappingId} | 
[**DeleteSettingsMappings**](QuotaApi.md#DeleteSettingsMappings) | **Delete** /platform/1/quota/settings/mappings | 
[**DeleteSettingsNotification**](QuotaApi.md#DeleteSettingsNotification) | **Delete** /platform/1/quota/settings/notifications/{SettingsNotificationId} | 
[**DeleteSettingsNotifications**](QuotaApi.md#DeleteSettingsNotifications) | **Delete** /platform/1/quota/settings/notifications | 
[**GetQuotaLicense**](QuotaApi.md#GetQuotaLicense) | **Get** /platform/5/quota/license | 
[**GetQuotaQuota**](QuotaApi.md#GetQuotaQuota) | **Get** /platform/1/quota/quotas/{QuotaQuotaId} | 
[**GetQuotaQuotasSummary**](QuotaApi.md#GetQuotaQuotasSummary) | **Get** /platform/1/quota/quotas-summary | 
[**GetQuotaReport**](QuotaApi.md#GetQuotaReport) | **Get** /platform/1/quota/reports/{QuotaReportId} | 
[**GetSettingsMapping**](QuotaApi.md#GetSettingsMapping) | **Get** /platform/1/quota/settings/mappings/{SettingsMappingId} | 
[**GetSettingsNotification**](QuotaApi.md#GetSettingsNotification) | **Get** /platform/1/quota/settings/notifications/{SettingsNotificationId} | 
[**GetSettingsReports**](QuotaApi.md#GetSettingsReports) | **Get** /platform/1/quota/settings/reports | 
[**ListQuotaQuotas**](QuotaApi.md#ListQuotaQuotas) | **Get** /platform/1/quota/quotas | 
[**ListQuotaReports**](QuotaApi.md#ListQuotaReports) | **Get** /platform/1/quota/reports | 
[**ListSettingsMappings**](QuotaApi.md#ListSettingsMappings) | **Get** /platform/1/quota/settings/mappings | 
[**ListSettingsNotifications**](QuotaApi.md#ListSettingsNotifications) | **Get** /platform/1/quota/settings/notifications | 
[**UpdateQuotaQuota**](QuotaApi.md#UpdateQuotaQuota) | **Put** /platform/1/quota/quotas/{QuotaQuotaId} | 
[**UpdateSettingsMapping**](QuotaApi.md#UpdateSettingsMapping) | **Put** /platform/1/quota/settings/mappings/{SettingsMappingId} | 
[**UpdateSettingsNotification**](QuotaApi.md#UpdateSettingsNotification) | **Put** /platform/1/quota/settings/notifications/{SettingsNotificationId} | 
[**UpdateSettingsReports**](QuotaApi.md#UpdateSettingsReports) | **Put** /platform/1/quota/settings/reports | 


# **CreateQuotaQuota**
> CreateResponse CreateQuotaQuota(ctx, quotaQuota, optional)


Create a new quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaQuota** | [**QuotaQuotaCreateParams**](QuotaQuotaCreateParams.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaQuota** | [**QuotaQuotaCreateParams**](QuotaQuotaCreateParams.md)|  | 
 **zone** | **string**| Optional named zone to use for user and group resolution. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuotaReport**
> CreateQuotaReportResponse CreateQuotaReport(ctx, quotaReport)


Create a new report. The type of this report is 'manual'; it is also sometimes called 'live' or 'ad-hoc'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaReport** | [**Empty**](Empty.md)|  | 

### Return type

[**CreateQuotaReportResponse**](CreateQuotaReportResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsMapping**
> CreateResponse CreateSettingsMapping(ctx, settingsMapping)


Create a new rule. The new rule must not conflict with an existing rule (e.g. match both the type and domain fields).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsMapping** | [**SettingsMappingExtendedExtended**](SettingsMappingExtendedExtended.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsNotification**
> CreateResponse CreateSettingsNotification(ctx, settingsNotification)


Create a new global notification rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsNotification** | [**QuotaNotificationCreateParams**](QuotaNotificationCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaQuota**
> DeleteQuotaQuota(ctx, quotaQuotaId)


Delete the quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaQuotaId** | **string**| Delete the quota. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaQuotas**
> DeleteQuotaQuotas(ctx, optional)


Delete all or matching quotas.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enforced** | **bool**| Only delete quotas with this enforcement (non-accounting). | 
 **includeSnapshots** | **bool**| Only delete quotas with this setting for include_snapshots. | 
 **zone** | **string**| Optional named zone to use for user and group resolution. | 
 **recursePathChildren** | **bool**| If used with the path argument, delete all quotas at that path or any descendent sub-directory. | 
 **recursePathParents** | **bool**| If used with the path argument, delete all quotas at that path or any parent directory. | 
 **persona** | **string**| Only delete user or group quotas matching this persona (must be used with the corresponding type argument).  Format is &lt;PERSONA_TYPE&gt;:&lt;string/integer&gt;, where PERSONA_TYPE is one of USER, GROUP, SID, ID, or GID. | 
 **path** | **string**| Only delete quotas matching this path (see also recurse_path_*). | 
 **type_** | **string**| Only delete quotas matching this type. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuotaReport**
> DeleteQuotaReport(ctx, quotaReportId)


Delete the report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaReportId** | **string**| Delete the report. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsMapping**
> DeleteSettingsMapping(ctx, settingsMappingId)


Delete the mapping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsMappingId** | **string**| Delete the mapping. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsMappings**
> DeleteSettingsMappings(ctx, )


Delete all rules.

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

# **DeleteSettingsNotification**
> DeleteSettingsNotification(ctx, settingsNotificationId)


Delete the notification rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsNotificationId** | **string**| Delete the notification rule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsNotifications**
> DeleteSettingsNotifications(ctx, )


Delete all rules.

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

# **GetQuotaLicense**
> LicenseLicense GetQuotaLicense(ctx, )


Retrieve license information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LicenseLicense**](LicenseLicense.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaQuota**
> QuotaQuotas GetQuotaQuota(ctx, quotaQuotaId, optional)


Retrieve quota information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaQuotaId** | **string**| Retrieve quota information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaQuotaId** | **string**| Retrieve quota information. | 
 **resolveNames** | **bool**| If true, resolve group and user names in personas. | 
 **zone** | **string**| Optional named zone to use for user and group resolution. | 

### Return type

[**QuotaQuotas**](QuotaQuotas.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaQuotasSummary**
> QuotaQuotasSummary GetQuotaQuotasSummary(ctx, )


Return summary information about quotas.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**QuotaQuotasSummary**](QuotaQuotasSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuotaReport**
> ReportAbout GetQuotaReport(ctx, quotaReportId, optional)


Retrieve report data (XML) or contents (meta-data).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaReportId** | **string**| Retrieve report data (XML) or contents (meta-data). | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaReportId** | **string**| Retrieve report data (XML) or contents (meta-data). | 
 **contents** | **bool**| Display JSON meta-data contents instead of report data. | 

### Return type

[**ReportAbout**](ReportAbout.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsMapping**
> SettingsMappings GetSettingsMapping(ctx, settingsMappingId)


Retrieve the mapping information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsMappingId** | **string**| Retrieve the mapping information. | 

### Return type

[**SettingsMappings**](SettingsMappings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsNotification**
> QuotaNotifications GetSettingsNotification(ctx, settingsNotificationId)


Retrieve notification rule information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsNotificationId** | **string**| Retrieve notification rule information. | 

### Return type

[**QuotaNotifications**](QuotaNotifications.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSettingsReports**
> SettingsReports GetSettingsReports(ctx, )


List all settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsReports**](SettingsReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaQuotas**
> QuotaQuotasExtended ListQuotaQuotas(ctx, optional)


List all or matching quotas. Can also be used to retrieve quota state from existing reports. For any query argument not supplied, the default behavior is return all.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enforced** | **bool**| Only list quotas with this enforcement (non-accounting). | 
 **includeSnapshots** | **bool**| Only list quotas with this setting for include_snapshots. | 
 **zone** | **string**| Optional named zone to use for user and group resolution. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **recursePathChildren** | **bool**| If used with the path argument, match all quotas at that path or any descendent sub-directory. | 
 **resolveNames** | **bool**| If true, resolve group and user names in personas. | 
 **recursePathParents** | **bool**| If used with the path argument, match all quotas at that path or any parent directory. | 
 **persona** | **string**| Only list user or group quotas matching this persona (must be used with the corresponding type argument).  Format is &lt;PERSONA_TYPE&gt;:&lt;string/integer&gt;, where PERSONA_TYPE is one of USER, GROUP, SID, ID, or GID. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **exceeded** | **bool**| Set to true to only list quotas which have exceeded one or more of their thresholds. | 
 **path** | **string**| Only list quotas matching this path (see also recurse_path_*). | 
 **type_** | **string**| Only list quotas matching this type. | 
 **reportId** | **string**| Use the named report as a source rather than the live quotas. See the /q/quota/reports resource for a list of valid reports. | 

### Return type

[**QuotaQuotasExtended**](QuotaQuotasExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListQuotaReports**
> QuotaReports ListQuotaReports(ctx, optional)


List all or matching reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| Order results by this field. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **generated** | **string**| Only list reports matching this source. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **type_** | **string**| Only list reports matching this type. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**QuotaReports**](QuotaReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSettingsMappings**
> SettingsMappings ListSettingsMappings(ctx, )


List all rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsMappings**](SettingsMappings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSettingsNotifications**
> QuotaNotificationsExtended ListSettingsNotifications(ctx, )


List all rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**QuotaNotificationsExtended**](QuotaNotificationsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuotaQuota**
> UpdateQuotaQuota(ctx, quotaQuota, quotaQuotaId)


Modify quota. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **quotaQuota** | [**QuotaQuota**](QuotaQuota.md)|  | 
  **quotaQuotaId** | **string**| Modify quota. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsMapping**
> UpdateSettingsMapping(ctx, settingsMapping, settingsMappingId)


Modify the mapping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsMapping** | [**SettingsMappingExtended**](SettingsMappingExtended.md)|  | 
  **settingsMappingId** | **string**| Modify the mapping. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsNotification**
> UpdateSettingsNotification(ctx, settingsNotification, settingsNotificationId)


Modify notification rule. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsNotification** | [**QuotaNotification**](QuotaNotification.md)|  | 
  **settingsNotificationId** | **string**| Modify notification rule. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSettingsReports**
> UpdateSettingsReports(ctx, settingsReports)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsReports** | [**SettingsReportsExtended**](SettingsReportsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

