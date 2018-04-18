# \AntivirusApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAntivirusPolicy**](AntivirusApi.md#CreateAntivirusPolicy) | **Post** /platform/3/antivirus/policies | 
[**CreateAntivirusScanItem**](AntivirusApi.md#CreateAntivirusScanItem) | **Post** /platform/3/antivirus/scan | 
[**CreateAntivirusServer**](AntivirusApi.md#CreateAntivirusServer) | **Post** /platform/3/antivirus/servers | 
[**DeleteAntivirusPolicies**](AntivirusApi.md#DeleteAntivirusPolicies) | **Delete** /platform/3/antivirus/policies | 
[**DeleteAntivirusPolicy**](AntivirusApi.md#DeleteAntivirusPolicy) | **Delete** /platform/3/antivirus/policies/{AntivirusPolicyId} | 
[**DeleteAntivirusServer**](AntivirusApi.md#DeleteAntivirusServer) | **Delete** /platform/3/antivirus/servers/{AntivirusServerId} | 
[**DeleteAntivirusServers**](AntivirusApi.md#DeleteAntivirusServers) | **Delete** /platform/3/antivirus/servers | 
[**DeleteReportsScan**](AntivirusApi.md#DeleteReportsScan) | **Delete** /platform/3/antivirus/reports/scans/{ReportsScanId} | 
[**DeleteReportsScans**](AntivirusApi.md#DeleteReportsScans) | **Delete** /platform/3/antivirus/reports/scans | 
[**GetAntivirusPolicy**](AntivirusApi.md#GetAntivirusPolicy) | **Get** /platform/3/antivirus/policies/{AntivirusPolicyId} | 
[**GetAntivirusQuarantinePath**](AntivirusApi.md#GetAntivirusQuarantinePath) | **Get** /platform/3/antivirus/quarantine/{AntivirusQuarantinePath} | 
[**GetAntivirusServer**](AntivirusApi.md#GetAntivirusServer) | **Get** /platform/3/antivirus/servers/{AntivirusServerId} | 
[**GetAntivirusSettings**](AntivirusApi.md#GetAntivirusSettings) | **Get** /platform/3/antivirus/settings | 
[**GetReportsScan**](AntivirusApi.md#GetReportsScan) | **Get** /platform/3/antivirus/reports/scans/{ReportsScanId} | 
[**GetReportsScans**](AntivirusApi.md#GetReportsScans) | **Get** /platform/3/antivirus/reports/scans | 
[**GetReportsThreat**](AntivirusApi.md#GetReportsThreat) | **Get** /platform/3/antivirus/reports/threats/{ReportsThreatId} | 
[**GetReportsThreats**](AntivirusApi.md#GetReportsThreats) | **Get** /platform/3/antivirus/reports/threats | 
[**ListAntivirusPolicies**](AntivirusApi.md#ListAntivirusPolicies) | **Get** /platform/3/antivirus/policies | 
[**ListAntivirusServers**](AntivirusApi.md#ListAntivirusServers) | **Get** /platform/3/antivirus/servers | 
[**UpdateAntivirusPolicy**](AntivirusApi.md#UpdateAntivirusPolicy) | **Put** /platform/3/antivirus/policies/{AntivirusPolicyId} | 
[**UpdateAntivirusQuarantinePath**](AntivirusApi.md#UpdateAntivirusQuarantinePath) | **Put** /platform/3/antivirus/quarantine/{AntivirusQuarantinePath} | 
[**UpdateAntivirusServer**](AntivirusApi.md#UpdateAntivirusServer) | **Put** /platform/3/antivirus/servers/{AntivirusServerId} | 
[**UpdateAntivirusSettings**](AntivirusApi.md#UpdateAntivirusSettings) | **Put** /platform/3/antivirus/settings | 


# **CreateAntivirusPolicy**
> CreateResponse CreateAntivirusPolicy(ctx, antivirusPolicy)


Create new antivirus scan policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusPolicy** | [**AntivirusPolicyCreateParams**](AntivirusPolicyCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAntivirusScanItem**
> CreateAntivirusScanItemResponse CreateAntivirusScanItem(ctx, antivirusScanItem)


Manually scan a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusScanItem** | [**AntivirusScanItem**](AntivirusScanItem.md)|  | 

### Return type

[**CreateAntivirusScanItemResponse**](CreateAntivirusScanItemResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAntivirusServer**
> CreateResponse CreateAntivirusServer(ctx, antivirusServer)


Create new antivirus servers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusServer** | [**AntivirusServerCreateParams**](AntivirusServerCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAntivirusPolicies**
> DeleteAntivirusPolicies(ctx, )


Delete all antivirus scan policies.

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

# **DeleteAntivirusPolicy**
> DeleteAntivirusPolicy(ctx, antivirusPolicyId)


Delete an antivirus scan policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusPolicyId** | **string**| Delete an antivirus scan policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAntivirusServer**
> DeleteAntivirusServer(ctx, antivirusServerId)


Delete an antivirus server entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusServerId** | **string**| Delete an antivirus server entry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAntivirusServers**
> DeleteAntivirusServers(ctx, )


Delete all antivirus servers.

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

# **DeleteReportsScan**
> DeleteReportsScan(ctx, reportsScanId)


Delete one antivirus scan report, and all of its associated threat reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **reportsScanId** | **string**| Delete one antivirus scan report, and all of its associated threat reports. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReportsScans**
> DeleteReportsScans(ctx, optional)


Delete antivirus scan reports, and any threat reports associated with those scans.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **age** | **int32**| An amount of time in seconds. If present, only reports older than this age are deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAntivirusPolicy**
> AntivirusPolicies GetAntivirusPolicy(ctx, antivirusPolicyId)


Retrieve one antivirus scan policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusPolicyId** | **string**| Retrieve one antivirus scan policy. | 

### Return type

[**AntivirusPolicies**](AntivirusPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAntivirusQuarantinePath**
> AntivirusQuarantine GetAntivirusQuarantinePath(ctx, antivirusQuarantinePath)


Retrieve the quarantine status of the file at the specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusQuarantinePath** | **string**| Retrieve the quarantine status of the file at the specified path. | 

### Return type

[**AntivirusQuarantine**](AntivirusQuarantine.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAntivirusServer**
> AntivirusServers GetAntivirusServer(ctx, antivirusServerId)


Retrieve one antivirus server entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusServerId** | **string**| Retrieve one antivirus server entry. | 

### Return type

[**AntivirusServers**](AntivirusServers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAntivirusSettings**
> AntivirusSettings GetAntivirusSettings(ctx, )


Retrieve the Antivirus settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AntivirusSettings**](AntivirusSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsScan**
> ReportsScans GetReportsScan(ctx, reportsScanId)


Retrieve one antivirus scan report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **reportsScanId** | **string**| Retrieve one antivirus scan report. | 

### Return type

[**ReportsScans**](ReportsScans.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsScans**
> ReportsScansExtended GetReportsScans(ctx, optional)


List antivirus scan reports.

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
 **status** | **string**| If present, only scan reports with this status will be returned. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **policyId** | **string**| If present, only reports for scans associated with this policy will be returned. | 

### Return type

[**ReportsScansExtended**](ReportsScansExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsThreat**
> ReportsThreats GetReportsThreat(ctx, reportsThreatId)


Retrieve one antivirus threat report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **reportsThreatId** | **string**| Retrieve one antivirus threat report. | 

### Return type

[**ReportsThreats**](ReportsThreats.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsThreats**
> ReportsThreatsExtended GetReportsThreats(ctx, optional)


List antivirus threat reports.

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
 **scanId** | **string**| If present, only returns threat reports associated with the given scan report. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **file** | **string**| If present, only returns threat reports for the given file. | 
 **remediation** | **string**| If present, only returns threat reports with the given remediation. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ReportsThreatsExtended**](ReportsThreatsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAntivirusPolicies**
> AntivirusPoliciesExtended ListAntivirusPolicies(ctx, optional)


List antivirus scan policies.

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

[**AntivirusPoliciesExtended**](AntivirusPoliciesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAntivirusServers**
> AntivirusServersExtended ListAntivirusServers(ctx, optional)


List antivirus servers.

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

[**AntivirusServersExtended**](AntivirusServersExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAntivirusPolicy**
> UpdateAntivirusPolicy(ctx, antivirusPolicy, antivirusPolicyId)


Modify an antivirus scan policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusPolicy** | [**AntivirusPolicy**](AntivirusPolicy.md)|  | 
  **antivirusPolicyId** | **string**| Modify an antivirus scan policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAntivirusQuarantinePath**
> UpdateAntivirusQuarantinePath(ctx, antivirusQuarantinePathParams, antivirusQuarantinePath)


Set the quarantine status of the file at the specified path.  Use either an empty object {} in the request body or {\"quarantined\":true} to quarantine the file, and {\"quarantined\":false} to unquarantine the file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusQuarantinePathParams** | [**AntivirusQuarantinePathParams**](AntivirusQuarantinePathParams.md)|  | 
  **antivirusQuarantinePath** | **string**| Set the quarantine status of the file at the specified path.  Use either an empty object {} in the request body or {\&quot;quarantined\&quot;:true} to quarantine the file, and {\&quot;quarantined\&quot;:false} to unquarantine the file. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAntivirusServer**
> UpdateAntivirusServer(ctx, antivirusServer, antivirusServerId)


Modify an antivirus server entry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusServer** | [**AntivirusServer**](AntivirusServer.md)|  | 
  **antivirusServerId** | **string**| Modify an antivirus server entry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAntivirusSettings**
> UpdateAntivirusSettings(ctx, antivirusSettings)


Modify the Antivirus settings. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **antivirusSettings** | [**AntivirusSettingsSettings**](AntivirusSettingsSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

