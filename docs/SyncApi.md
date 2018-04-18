# \SyncApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSyncJob**](SyncApi.md#CreateSyncJob) | **Post** /platform/3/sync/jobs | 
[**CreateSyncPolicy**](SyncApi.md#CreateSyncPolicy) | **Post** /platform/3/sync/policies | 
[**CreateSyncReportsRotateItem**](SyncApi.md#CreateSyncReportsRotateItem) | **Post** /platform/1/sync/reports-rotate | 
[**CreateSyncRule**](SyncApi.md#CreateSyncRule) | **Post** /platform/3/sync/rules | 
[**DeleteSyncPolicies**](SyncApi.md#DeleteSyncPolicies) | **Delete** /platform/3/sync/policies | 
[**DeleteSyncPolicy**](SyncApi.md#DeleteSyncPolicy) | **Delete** /platform/3/sync/policies/{SyncPolicyId} | 
[**DeleteSyncRule**](SyncApi.md#DeleteSyncRule) | **Delete** /platform/3/sync/rules/{SyncRuleId} | 
[**DeleteSyncRules**](SyncApi.md#DeleteSyncRules) | **Delete** /platform/3/sync/rules | 
[**DeleteTargetPolicy**](SyncApi.md#DeleteTargetPolicy) | **Delete** /platform/1/sync/target/policies/{TargetPolicyId} | 
[**GetHistoryCpu**](SyncApi.md#GetHistoryCpu) | **Get** /platform/3/sync/history/cpu | 
[**GetHistoryFile**](SyncApi.md#GetHistoryFile) | **Get** /platform/1/sync/history/file | 
[**GetHistoryNetwork**](SyncApi.md#GetHistoryNetwork) | **Get** /platform/1/sync/history/network | 
[**GetHistoryWorker**](SyncApi.md#GetHistoryWorker) | **Get** /platform/3/sync/history/worker | 
[**GetSyncJob**](SyncApi.md#GetSyncJob) | **Get** /platform/3/sync/jobs/{SyncJobId} | 
[**GetSyncLicense**](SyncApi.md#GetSyncLicense) | **Get** /platform/5/sync/license | 
[**GetSyncPolicy**](SyncApi.md#GetSyncPolicy) | **Get** /platform/3/sync/policies/{SyncPolicyId} | 
[**GetSyncReport**](SyncApi.md#GetSyncReport) | **Get** /platform/4/sync/reports/{SyncReportId} | 
[**GetSyncReports**](SyncApi.md#GetSyncReports) | **Get** /platform/4/sync/reports | 
[**GetSyncRule**](SyncApi.md#GetSyncRule) | **Get** /platform/3/sync/rules/{SyncRuleId} | 
[**GetSyncSettings**](SyncApi.md#GetSyncSettings) | **Get** /platform/3/sync/settings | 
[**GetTargetPolicies**](SyncApi.md#GetTargetPolicies) | **Get** /platform/1/sync/target/policies | 
[**GetTargetPolicy**](SyncApi.md#GetTargetPolicy) | **Get** /platform/1/sync/target/policies/{TargetPolicyId} | 
[**GetTargetReport**](SyncApi.md#GetTargetReport) | **Get** /platform/4/sync/target/reports/{TargetReportId} | 
[**GetTargetReports**](SyncApi.md#GetTargetReports) | **Get** /platform/4/sync/target/reports | 
[**ListSyncJobs**](SyncApi.md#ListSyncJobs) | **Get** /platform/3/sync/jobs | 
[**ListSyncPolicies**](SyncApi.md#ListSyncPolicies) | **Get** /platform/3/sync/policies | 
[**ListSyncReportsRotate**](SyncApi.md#ListSyncReportsRotate) | **Get** /platform/1/sync/reports-rotate | 
[**ListSyncRules**](SyncApi.md#ListSyncRules) | **Get** /platform/3/sync/rules | 
[**UpdateSyncJob**](SyncApi.md#UpdateSyncJob) | **Put** /platform/3/sync/jobs/{SyncJobId} | 
[**UpdateSyncPolicy**](SyncApi.md#UpdateSyncPolicy) | **Put** /platform/3/sync/policies/{SyncPolicyId} | 
[**UpdateSyncRule**](SyncApi.md#UpdateSyncRule) | **Put** /platform/3/sync/rules/{SyncRuleId} | 
[**UpdateSyncSettings**](SyncApi.md#UpdateSyncSettings) | **Put** /platform/3/sync/settings | 


# **CreateSyncJob**
> CreateResponse CreateSyncJob(ctx, syncJob)


Start a SyncIQ job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncJob** | [**SyncJobCreateParams**](SyncJobCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSyncPolicy**
> CreateResponse CreateSyncPolicy(ctx, syncPolicy)


Create a SyncIQ policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncPolicy** | [**SyncPolicyCreateParams**](SyncPolicyCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSyncReportsRotateItem**
> CreateSyncReportsRotateItemResponse CreateSyncReportsRotateItem(ctx, syncReportsRotateItem)


Rotate the records in the database(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncReportsRotateItem** | [**Empty**](Empty.md)|  | 

### Return type

[**CreateSyncReportsRotateItemResponse**](CreateSyncReportsRotateItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSyncRule**
> CreateResponse CreateSyncRule(ctx, syncRule)


Create a new SyncIQ performance rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncRule** | [**SyncRuleCreateParams**](SyncRuleCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSyncPolicies**
> DeleteSyncPolicies(ctx, optional)


Delete all SyncIQ policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **localOnly** | **bool**| Skip deleting the policy association on the target. | 
 **force** | **bool**| Ignore any running jobs when preparing to delete a policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSyncPolicy**
> DeleteSyncPolicy(ctx, syncPolicyId, optional)


Delete a single SyncIQ policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncPolicyId** | **string**| Delete a single SyncIQ policy. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncPolicyId** | **string**| Delete a single SyncIQ policy. | 
 **localOnly** | **bool**| Skip deleting the policy association on the target. | 
 **force** | **bool**| Ignore any running jobs when preparing to delete a policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSyncRule**
> DeleteSyncRule(ctx, syncRuleId)


Delete a single SyncIQ performance rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncRuleId** | **string**| Delete a single SyncIQ performance rule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSyncRules**
> DeleteSyncRules(ctx, optional)


Delete all SyncIQ performance rules or all rules of a specified type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| Delete all rules of the specified rule type only. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTargetPolicy**
> DeleteTargetPolicy(ctx, targetPolicyId, optional)


Break the target association with this cluster for this policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **targetPolicyId** | **string**| Break the target association with this cluster for this policy. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetPolicyId** | **string**| Break the target association with this cluster for this policy. | 
 **force** | **bool**| Ignore any running jobs when preparing to delete the policy target association. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryCpu**
> HistoryFile GetHistoryCpu(ctx, optional)


List cpu performance data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Begin timestamp for time-series report. | 
 **end** | **int32**| End timestamp for time-series report. | 

### Return type

[**HistoryFile**](HistoryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryFile**
> HistoryFile GetHistoryFile(ctx, optional)


List file operations performance data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Begin timestamp for time-series report. | 
 **end** | **int32**| End timestamp for time-series report. | 

### Return type

[**HistoryFile**](HistoryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryNetwork**
> HistoryFile GetHistoryNetwork(ctx, optional)


List network operations performance data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Begin timestamp for time-series report. | 
 **end** | **int32**| End timestamp for time-series report. | 

### Return type

[**HistoryFile**](HistoryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryWorker**
> HistoryFile GetHistoryWorker(ctx, optional)


List worker performance data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Begin timestamp for time-series report. | 
 **end** | **int32**| End timestamp for time-series report. | 

### Return type

[**HistoryFile**](HistoryFile.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncJob**
> SyncJobs GetSyncJob(ctx, syncJobId)


View a single SyncIQ job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncJobId** | **string**| View a single SyncIQ job. | 

### Return type

[**SyncJobs**](SyncJobs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncLicense**
> LicenseLicense GetSyncLicense(ctx, )


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

# **GetSyncPolicy**
> SyncPolicies GetSyncPolicy(ctx, syncPolicyId)


View a single SyncIQ policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncPolicyId** | **string**| View a single SyncIQ policy. | 

### Return type

[**SyncPolicies**](SyncPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncReport**
> SyncReports GetSyncReport(ctx, syncReportId)


View a single SyncIQ report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncReportId** | **string**| View a single SyncIQ report. | 

### Return type

[**SyncReports**](SyncReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncReports**
> SyncReportsExtended GetSyncReports(ctx, optional)


Get a list of SyncIQ reports.  By default 10 reports are returned per policy, unless otherwise specified by 'reports_per_policy'.

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
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **newerThan** | **int32**| Filter the returned reports to include only those whose jobs started more recently than the specified number of days ago. | 
 **policyName** | **string**| Filter the returned reports to include only those with this policy name. | 
 **state** | **string**| Filter the returned reports to include only those whose jobs are in this state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **reportsPerPolicy** | **int32**| If specified, only the N most recent reports will be returned per policy.  If no other query args are present this argument defaults to 10.  | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**SyncReportsExtended**](SyncReportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncRule**
> SyncRules GetSyncRule(ctx, syncRuleId)


View a single SyncIQ performance rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncRuleId** | **string**| View a single SyncIQ performance rule. | 

### Return type

[**SyncRules**](SyncRules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSyncSettings**
> SyncSettings GetSyncSettings(ctx, )


Retrieve the global SyncIQ settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SyncSettings**](SyncSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetPolicies**
> TargetPoliciesExtended GetTargetPolicies(ctx, optional)


List all SyncIQ target policies.

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
 **targetPath** | **string**| Filter the returned policies to include only those with this target path. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**TargetPoliciesExtended**](TargetPoliciesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetPolicy**
> TargetPolicies GetTargetPolicy(ctx, targetPolicyId)


View a single SyncIQ target policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **targetPolicyId** | **string**| View a single SyncIQ target policy. | 

### Return type

[**TargetPolicies**](TargetPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetReport**
> TargetReports GetTargetReport(ctx, targetReportId)


View a single SyncIQ target report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **targetReportId** | **string**| View a single SyncIQ target report. | 

### Return type

[**TargetReports**](TargetReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetReports**
> TargetReportsExtended GetTargetReports(ctx, optional)


Get a list of SyncIQ target reports.  By default 10 reports are returned per policy, unless otherwise specified by 'reports_per_policy'.

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
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **newerThan** | **int32**| Filter the returned reports to include only those whose jobs started more recently than the specified number of days ago. | 
 **policyName** | **string**| Filter the returned reports to include only those with this policy name. | 
 **state** | **string**| Filter the returned reports to include only those whose jobs are in this state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **reportsPerPolicy** | **int32**| If specified, only the N most recent reports will be returned per policy.  If no other query args are present this argument defaults to 10.  | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**TargetReportsExtended**](TargetReportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSyncJobs**
> SyncJobsExtended ListSyncJobs(ctx, optional)


Get a list of SyncIQ jobs.

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
 **state** | **string**| The state of the job. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SyncJobsExtended**](SyncJobsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSyncPolicies**
> SyncPoliciesExtended ListSyncPolicies(ctx, optional)


List all SyncIQ policies.

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
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **summary** | **bool**| Show only summary properties | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**SyncPoliciesExtended**](SyncPoliciesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSyncReportsRotate**
> SyncReportsRotate ListSyncReportsRotate(ctx, )


Whether the rotation is still running or not.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SyncReportsRotate**](SyncReportsRotate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSyncRules**
> SyncRulesExtended ListSyncRules(ctx, optional)


List all SyncIQ performance rules.

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
 **type_** | **string**| Filter the returned rules to include only those with this rule type. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SyncRulesExtended**](SyncRulesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSyncJob**
> UpdateSyncJob(ctx, syncJob, syncJobId)


Perform an action (pause, cancel, etc...) on a single job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncJob** | [**SyncJob**](SyncJob.md)|  | 
  **syncJobId** | **string**| Perform an action (pause, cancel, etc...) on a single job. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSyncPolicy**
> UpdateSyncPolicy(ctx, syncPolicy, syncPolicyId)


Modify a single SyncIQ policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncPolicy** | [**SyncPolicy**](SyncPolicy.md)|  | 
  **syncPolicyId** | **string**| Modify a single SyncIQ policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSyncRule**
> UpdateSyncRule(ctx, syncRule, syncRuleId)


Modify a single SyncIQ performance rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncRule** | [**SyncRule**](SyncRule.md)|  | 
  **syncRuleId** | **string**| Modify a single SyncIQ performance rule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSyncSettings**
> UpdateSyncSettings(ctx, syncSettings)


Modify the global SyncIQ settings.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **syncSettings** | [**SyncSettingsExtended**](SyncSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

