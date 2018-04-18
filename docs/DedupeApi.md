# \DedupeApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDedupeDedupeSummary**](DedupeApi.md#GetDedupeDedupeSummary) | **Get** /platform/1/dedupe/dedupe-summary | 
[**GetDedupeReport**](DedupeApi.md#GetDedupeReport) | **Get** /platform/1/dedupe/reports/{DedupeReportId} | 
[**GetDedupeReports**](DedupeApi.md#GetDedupeReports) | **Get** /platform/1/dedupe/reports | 
[**GetDedupeSettings**](DedupeApi.md#GetDedupeSettings) | **Get** /platform/1/dedupe/settings | 
[**UpdateDedupeSettings**](DedupeApi.md#UpdateDedupeSettings) | **Put** /platform/1/dedupe/settings | 


# **GetDedupeDedupeSummary**
> DedupeDedupeSummary GetDedupeDedupeSummary(ctx, )


Return summary information about dedupe.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DedupeDedupeSummary**](DedupeDedupeSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedupeReport**
> DedupeReports GetDedupeReport(ctx, dedupeReportId, optional)


Retrieve a report for a single dedupe job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dedupeReportId** | **string**| Retrieve a report for a single dedupe job. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dedupeReportId** | **string**| Retrieve a report for a single dedupe job. | 
 **scope** | **string**| If specified as \&quot;effective\&quot; or not specified, all fields are returned.  If specified as \&quot;user\&quot;, only fields with non-default values are shown.  If specified as \&quot;default\&quot;, the original values are returned. | 

### Return type

[**DedupeReports**](DedupeReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedupeReports**
> DedupeReportsExtended GetDedupeReports(ctx, optional)


List dedupe reports.

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
 **begin** | **int32**| Restrict the query to reports at or after the given time, in seconds since the Epoch. | 
 **end** | **int32**| Restrict the query to reports at or before the given time, in seconds since the Epoch. | 
 **jobId** | **int32**| Restrict the query to the given job ID. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **jobType** | **string**| Restrict the query to the given job type. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**DedupeReportsExtended**](DedupeReportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedupeSettings**
> DedupeSettings GetDedupeSettings(ctx, )


Retrieve the dedupe settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DedupeSettings**](DedupeSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDedupeSettings**
> UpdateDedupeSettings(ctx, dedupeSettings)


Modify the dedupe settings. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **dedupeSettings** | [**DedupeSettingsExtended**](DedupeSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

