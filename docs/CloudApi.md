# \CloudApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudAccessItem**](CloudApi.md#CreateCloudAccessItem) | **Post** /platform/3/cloud/access | 
[**CreateCloudAccount**](CloudApi.md#CreateCloudAccount) | **Post** /platform/4/cloud/accounts | 
[**CreateCloudJob**](CloudApi.md#CreateCloudJob) | **Post** /platform/3/cloud/jobs | 
[**CreateCloudPool**](CloudApi.md#CreateCloudPool) | **Post** /platform/3/cloud/pools | 
[**CreateCloudProxy**](CloudApi.md#CreateCloudProxy) | **Post** /platform/4/cloud/proxies | 
[**CreateSettingsEncryptionKeyItem**](CloudApi.md#CreateSettingsEncryptionKeyItem) | **Post** /platform/3/cloud/settings/encryption-key | 
[**CreateSettingsReportingEulaItem**](CloudApi.md#CreateSettingsReportingEulaItem) | **Post** /platform/3/cloud/settings/reporting-eula | 
[**DeleteCloudAccessGuid**](CloudApi.md#DeleteCloudAccessGuid) | **Delete** /platform/3/cloud/access/{CloudAccessGuid} | 
[**DeleteCloudAccount**](CloudApi.md#DeleteCloudAccount) | **Delete** /platform/4/cloud/accounts/{CloudAccountId} | 
[**DeleteCloudPool**](CloudApi.md#DeleteCloudPool) | **Delete** /platform/3/cloud/pools/{CloudPoolId} | 
[**DeleteCloudProxy**](CloudApi.md#DeleteCloudProxy) | **Delete** /platform/4/cloud/proxies/{CloudProxyId} | 
[**DeleteSettingsReportingEula**](CloudApi.md#DeleteSettingsReportingEula) | **Delete** /platform/3/cloud/settings/reporting-eula | 
[**GetCloudAccessGuid**](CloudApi.md#GetCloudAccessGuid) | **Get** /platform/3/cloud/access/{CloudAccessGuid} | 
[**GetCloudAccount**](CloudApi.md#GetCloudAccount) | **Get** /platform/4/cloud/accounts/{CloudAccountId} | 
[**GetCloudJob**](CloudApi.md#GetCloudJob) | **Get** /platform/3/cloud/jobs/{CloudJobId} | 
[**GetCloudJobsFile**](CloudApi.md#GetCloudJobsFile) | **Get** /platform/3/cloud/jobs-files/{CloudJobsFileId} | 
[**GetCloudPool**](CloudApi.md#GetCloudPool) | **Get** /platform/3/cloud/pools/{CloudPoolId} | 
[**GetCloudProxy**](CloudApi.md#GetCloudProxy) | **Get** /platform/4/cloud/proxies/{CloudProxyId} | 
[**GetCloudSettings**](CloudApi.md#GetCloudSettings) | **Get** /platform/3/cloud/settings | 
[**ListCloudAccess**](CloudApi.md#ListCloudAccess) | **Get** /platform/3/cloud/access | 
[**ListCloudAccounts**](CloudApi.md#ListCloudAccounts) | **Get** /platform/4/cloud/accounts | 
[**ListCloudJobs**](CloudApi.md#ListCloudJobs) | **Get** /platform/3/cloud/jobs | 
[**ListCloudPools**](CloudApi.md#ListCloudPools) | **Get** /platform/3/cloud/pools | 
[**ListCloudProxies**](CloudApi.md#ListCloudProxies) | **Get** /platform/4/cloud/proxies | 
[**ListSettingsReportingEula**](CloudApi.md#ListSettingsReportingEula) | **Get** /platform/3/cloud/settings/reporting-eula | 
[**UpdateCloudAccount**](CloudApi.md#UpdateCloudAccount) | **Put** /platform/4/cloud/accounts/{CloudAccountId} | 
[**UpdateCloudJob**](CloudApi.md#UpdateCloudJob) | **Put** /platform/3/cloud/jobs/{CloudJobId} | 
[**UpdateCloudPool**](CloudApi.md#UpdateCloudPool) | **Put** /platform/3/cloud/pools/{CloudPoolId} | 
[**UpdateCloudProxy**](CloudApi.md#UpdateCloudProxy) | **Put** /platform/4/cloud/proxies/{CloudProxyId} | 
[**UpdateCloudSettings**](CloudApi.md#UpdateCloudSettings) | **Put** /platform/3/cloud/settings | 


# **CreateCloudAccessItem**
> Empty CreateCloudAccessItem(ctx, cloudAccessItem)


Add a cluster identifier to access list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccessItem** | [**CloudAccessItem**](CloudAccessItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudAccount**
> CreateCloudAccountResponse CreateCloudAccount(ctx, cloudAccount)


Create a new account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccount** | [**CloudAccountCreateParams**](CloudAccountCreateParams.md)|  | 

### Return type

[**CreateCloudAccountResponse**](CreateCloudAccountResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudJob**
> CreateCloudJobResponse CreateCloudJob(ctx, cloudJob)


Create a new job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudJob** | [**CloudJobCreateParams**](CloudJobCreateParams.md)|  | 

### Return type

[**CreateCloudJobResponse**](CreateCloudJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudPool**
> CreateCloudPoolResponse CreateCloudPool(ctx, cloudPool)


Create a new pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPool** | [**CloudPoolCreateParams**](CloudPoolCreateParams.md)|  | 

### Return type

[**CreateCloudPoolResponse**](CreateCloudPoolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCloudProxy**
> CreateCloudProxyResponse CreateCloudProxy(ctx, cloudProxy)


Create a new proxy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProxy** | [**CloudProxyCreateParams**](CloudProxyCreateParams.md)|  | 

### Return type

[**CreateCloudProxyResponse**](CreateCloudProxyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsEncryptionKeyItem**
> Empty CreateSettingsEncryptionKeyItem(ctx, settingsEncryptionKeyItem)


Regenerate master encryption key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsEncryptionKeyItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSettingsReportingEulaItem**
> SettingsReportingEulaItem CreateSettingsReportingEulaItem(ctx, settingsReportingEulaItem)


Accept telemetry collection EULA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **settingsReportingEulaItem** | [**SettingsReportingEulaItem**](SettingsReportingEulaItem.md)|  | 

### Return type

[**SettingsReportingEulaItem**](SettingsReportingEulaItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudAccessGuid**
> DeleteCloudAccessGuid(ctx, cloudAccessGuid)


Delete cloud access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccessGuid** | **string**| Delete cloud access. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudAccount**
> DeleteCloudAccount(ctx, cloudAccountId, optional)


Delete cloud account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccountId** | **string**| Delete cloud account. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAccountId** | **string**| Delete cloud account. | 
 **acknowledgeForceDelete** | **string**| A value of 1 acknowledges that the user is deleting data. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudPool**
> DeleteCloudPool(ctx, cloudPoolId, optional)


Delete a cloud pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPoolId** | **string**| Delete a cloud pool. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPoolId** | **string**| Delete a cloud pool. | 
 **acknowledgeForceDelete** | **string**| A value of 1 acknowledges that the user is deleting data. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCloudProxy**
> DeleteCloudProxy(ctx, cloudProxyId)


Delete cloud account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProxyId** | **string**| Delete cloud account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSettingsReportingEula**
> DeleteSettingsReportingEula(ctx, )


Revoke acceptance of telemetry collection EULA.

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

# **GetCloudAccessGuid**
> CloudAccess GetCloudAccessGuid(ctx, cloudAccessGuid)


Retrieve cloud access information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccessGuid** | **string**| Retrieve cloud access information. | 

### Return type

[**CloudAccess**](CloudAccess.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudAccount**
> CloudAccounts GetCloudAccount(ctx, cloudAccountId)


Retrieve cloud account information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccountId** | **string**| Retrieve cloud account information. | 

### Return type

[**CloudAccounts**](CloudAccounts.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudJob**
> CloudJobs GetCloudJob(ctx, cloudJobId)


Retrieve cloudpool job information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudJobId** | **string**| Retrieve cloudpool job information. | 

### Return type

[**CloudJobs**](CloudJobs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudJobsFile**
> CloudJobsFiles GetCloudJobsFile(ctx, cloudJobsFileId, optional)


Retrieve files associated with a cloudpool job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudJobsFileId** | **string**| Retrieve files associated with a cloudpool job. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudJobsFileId** | **string**| Retrieve files associated with a cloudpool job. | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **batch** | **bool**| If true, only \&quot;limit\&quot; and \&quot;page\&quot; arguments are honored.  Query will return all results, unsorted, as quickly as possible. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **page** | **int32**| Works only when \&quot;batch\&quot; parameter and \&quot;limit\&quot; parameters are specified.  Indicates which the page index of results to be returned | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**CloudJobsFiles**](CloudJobsFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudPool**
> CloudPools GetCloudPool(ctx, cloudPoolId)


Retrieve cloud pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPoolId** | **string**| Retrieve cloud pool information | 

### Return type

[**CloudPools**](CloudPools.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudProxy**
> CloudProxies GetCloudProxy(ctx, cloudProxyId)


Retrieve cloud account information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProxyId** | **string**| Retrieve cloud account information. | 

### Return type

[**CloudProxies**](CloudProxies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudSettings**
> CloudSettings GetCloudSettings(ctx, )


List all cloud settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CloudSettings**](CloudSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudAccess**
> CloudAccessExtended ListCloudAccess(ctx, optional)


List all accessible cluster identifiers.

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

### Return type

[**CloudAccessExtended**](CloudAccessExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudAccounts**
> CloudAccountsExtended ListCloudAccounts(ctx, optional)


List all accounts.

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

### Return type

[**CloudAccountsExtended**](CloudAccountsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudJobs**
> CloudJobsExtended ListCloudJobs(ctx, optional)


List all cloudpools jobs.

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

### Return type

[**CloudJobsExtended**](CloudJobsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudPools**
> CloudPoolsExtended ListCloudPools(ctx, optional)


List all pools.

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

### Return type

[**CloudPoolsExtended**](CloudPoolsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloudProxies**
> CloudProxiesExtended ListCloudProxies(ctx, optional)


List all proxies.

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

### Return type

[**CloudProxiesExtended**](CloudProxiesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSettingsReportingEula**
> SettingsReportingEulaItem ListSettingsReportingEula(ctx, )


View telemetry collection EULA acceptance and content URI.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsReportingEulaItem**](SettingsReportingEulaItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudAccount**
> UpdateCloudAccount(ctx, cloudAccount, cloudAccountId)


Modify cloud account.  All fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudAccount** | [**CloudAccount**](CloudAccount.md)|  | 
  **cloudAccountId** | **string**| Modify cloud account.  All fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudJob**
> UpdateCloudJob(ctx, cloudJob, cloudJobId)


Modify a cloud job or operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudJob** | [**CloudJob**](CloudJob.md)|  | 
  **cloudJobId** | **string**| Modify a cloud job or operation. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudPool**
> UpdateCloudPool(ctx, cloudPool, cloudPoolId)


Modify a cloud pool.  All fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudPool** | [**CloudPool**](CloudPool.md)|  | 
  **cloudPoolId** | **string**| Modify a cloud pool.  All fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudProxy**
> UpdateCloudProxy(ctx, cloudProxy, cloudProxyId)


Modify cloud account.  All fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProxy** | [**CloudProxy**](CloudProxy.md)|  | 
  **cloudProxyId** | **string**| Modify cloud account.  All fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCloudSettings**
> UpdateCloudSettings(ctx, cloudSettings)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudSettings** | [**CloudSettingsSettings**](CloudSettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

