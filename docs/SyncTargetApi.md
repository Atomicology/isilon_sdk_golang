# \SyncTargetApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePoliciesPolicyCancelItem**](SyncTargetApi.md#CreatePoliciesPolicyCancelItem) | **Post** /platform/1/sync/target/policies/{Policy}/cancel | 
[**GetReportsReportSubreport**](SyncTargetApi.md#GetReportsReportSubreport) | **Get** /platform/4/sync/target/reports/{Rid}/subreports/{ReportsReportSubreportId} | 
[**GetReportsReportSubreports**](SyncTargetApi.md#GetReportsReportSubreports) | **Get** /platform/4/sync/target/reports/{Rid}/subreports | 


# **CreatePoliciesPolicyCancelItem**
> CreateResponse CreatePoliciesPolicyCancelItem(ctx, policiesPolicyCancelItem, policy)


Cancel the most recent SyncIQ job for this policy from the target side.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **policiesPolicyCancelItem** | [**Empty**](Empty.md)|  | 
  **policy** | **string**|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsReportSubreport**
> ReportsReportSubreports GetReportsReportSubreport(ctx, reportsReportSubreportId, rid)


View a single SyncIQ target subreport.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **reportsReportSubreportId** | **string**| View a single SyncIQ target subreport. | 
  **rid** | **string**|  | 

### Return type

[**ReportsReportSubreports**](ReportsReportSubreports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsReportSubreports**
> ReportsReportSubreportsExtended GetReportsReportSubreports(ctx, rid, optional)


Get a list of SyncIQ target subreports for a report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **rid** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rid** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **newerThan** | **int32**| Filter the returned reports to include only those whose jobs started more recently than the specified number of days ago. | 
 **state** | **string**| Filter the returned reports to include only those whose jobs are in this state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ReportsReportSubreportsExtended**](ReportsReportSubreportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

