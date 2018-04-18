# \SyncReportsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportSubreport**](SyncReportsApi.md#GetReportSubreport) | **Get** /platform/4/sync/reports/{Rid}/subreports/{ReportSubreportId} | 
[**GetReportSubreports**](SyncReportsApi.md#GetReportSubreports) | **Get** /platform/4/sync/reports/{Rid}/subreports | 


# **GetReportSubreport**
> ReportSubreports GetReportSubreport(ctx, reportSubreportId, rid)


View a single SyncIQ subreport.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **reportSubreportId** | **string**| View a single SyncIQ subreport. | 
  **rid** | **string**|  | 

### Return type

[**ReportSubreports**](ReportSubreports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportSubreports**
> ReportSubreportsExtended GetReportSubreports(ctx, rid, optional)


Get a list of SyncIQ subreports for a report.

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

[**ReportSubreportsExtended**](ReportSubreportsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

