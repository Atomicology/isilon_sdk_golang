# \JobApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJobJob**](JobApi.md#CreateJobJob) | **Post** /platform/3/job/jobs | 
[**CreateJobPolicy**](JobApi.md#CreateJobPolicy) | **Post** /platform/1/job/policies | 
[**DeleteJobPolicy**](JobApi.md#DeleteJobPolicy) | **Delete** /platform/1/job/policies/{JobPolicyId} | 
[**GetJobEvents**](JobApi.md#GetJobEvents) | **Get** /platform/3/job/events | 
[**GetJobJob**](JobApi.md#GetJobJob) | **Get** /platform/3/job/jobs/{JobJobId} | 
[**GetJobJobSummary**](JobApi.md#GetJobJobSummary) | **Get** /platform/1/job/job-summary | 
[**GetJobPolicy**](JobApi.md#GetJobPolicy) | **Get** /platform/1/job/policies/{JobPolicyId} | 
[**GetJobRecent**](JobApi.md#GetJobRecent) | **Get** /platform/3/job/recent | 
[**GetJobReports**](JobApi.md#GetJobReports) | **Get** /platform/3/job/reports | 
[**GetJobStatistics**](JobApi.md#GetJobStatistics) | **Get** /platform/1/job/statistics | 
[**GetJobType**](JobApi.md#GetJobType) | **Get** /platform/1/job/types/{JobTypeId} | 
[**GetJobTypes**](JobApi.md#GetJobTypes) | **Get** /platform/1/job/types | 
[**ListJobJobs**](JobApi.md#ListJobJobs) | **Get** /platform/3/job/jobs | 
[**ListJobPolicies**](JobApi.md#ListJobPolicies) | **Get** /platform/1/job/policies | 
[**UpdateJobJob**](JobApi.md#UpdateJobJob) | **Put** /platform/3/job/jobs/{JobJobId} | 
[**UpdateJobPolicy**](JobApi.md#UpdateJobPolicy) | **Put** /platform/1/job/policies/{JobPolicyId} | 
[**UpdateJobType**](JobApi.md#UpdateJobType) | **Put** /platform/1/job/types/{JobTypeId} | 


# **CreateJobJob**
> CreateJobJobResponse CreateJobJob(ctx, jobJob)


Queue a new instance of a job type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobJob** | [**JobJobCreateParams**](JobJobCreateParams.md)|  | 

### Return type

[**CreateJobJobResponse**](CreateJobJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateJobPolicy**
> CreateResponse CreateJobPolicy(ctx, jobPolicy)


Create a new job impact policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobPolicy** | [**JobPolicyCreateParams**](JobPolicyCreateParams.md)|  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteJobPolicy**
> DeleteJobPolicy(ctx, jobPolicyId)


Delete a job impact policy.  System policies may not be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobPolicyId** | **string**| Delete a job impact policy.  System policies may not be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobEvents**
> JobEvents GetJobEvents(ctx, optional)


List job events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Restrict the query to events at or after the given time, in seconds since the Epoch. | 
 **end** | **int32**| Restrict the query to events before the given time, in seconds since the Epoch. | 
 **jobId** | **int32**| Restrict the query to the given job ID. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **jobType** | **string**| Restrict the query to the given job type. | 
 **timeoutMs** | **int32**| Query timeout in milliseconds. The default is 10000 ms. | 
 **state** | **string**| Restrict the query to events containing the given state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **key** | **string**| Restrict the query to the given key name. | 

### Return type

[**JobEvents**](JobEvents.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobJob**
> JobJobs GetJobJob(ctx, jobJobId)


View a single job instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobJobId** | **string**| View a single job instance. | 

### Return type

[**JobJobs**](JobJobs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobJobSummary**
> JobJobSummary GetJobJobSummary(ctx, )


View job engine status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**JobJobSummary**](JobJobSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobPolicy**
> JobPolicies GetJobPolicy(ctx, jobPolicyId)


View a single job impact policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobPolicyId** | **string**| View a single job impact policy. | 

### Return type

[**JobPolicies**](JobPolicies.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobRecent**
> JobRecent GetJobRecent(ctx, optional)


List recently completed jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeoutMs** | **int32**| Query timeout in milliseconds. The default is 10000 ms. | 
 **limit** | **int32**| Max number of recent jobs to return. The default is 8, the max is 100. | 

### Return type

[**JobRecent**](JobRecent.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobReports**
> JobReports GetJobReports(ctx, optional)


List job reports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **begin** | **int32**| Restrict the query to reports at or after the given time, in seconds since the Epoch. | 
 **end** | **int32**| Restrict the query to reports before the given time, in seconds since the Epoch. | 
 **jobId** | **int32**| Restrict the query to the given job ID. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **jobType** | **string**| Restrict the query to the given job type. | 
 **timeoutMs** | **int32**| Query timeout in milliseconds. The default is 10000 ms. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **key** | **string**| Restrict the query to the given report key. | 
 **verbose** | **bool**| Display more detailed information, including job engine framework statistics. | 

### Return type

[**JobReports**](JobReports.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobStatistics**
> JobStatistics GetJobStatistics(ctx, optional)


View job engine statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devid** | **int32**| Restrict the query to the given node. | 
 **jobId** | **int32**| Restrict the query to the given job ID. | 

### Return type

[**JobStatistics**](JobStatistics.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobType**
> JobTypes GetJobType(ctx, jobTypeId)


Retrieve job type information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobTypeId** | **string**| Retrieve job type information. | 

### Return type

[**JobTypes**](JobTypes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobTypes**
> JobTypesExtended GetJobTypes(ctx, optional)


List job types.

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
 **showAll** | **bool**| Whether to show all job types, including hidden ones.  Defaults to false. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**JobTypesExtended**](JobTypesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobJobs**
> JobJobsExtended ListJobJobs(ctx, optional)


List running and paused jobs.

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
 **batch** | **bool**| If true, other arguments are ignored, and the query will return all results, unsorted, as quickly as possible. | 
 **state** | **string**| Limit the results to jobs in the specified state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**JobJobsExtended**](JobJobsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListJobPolicies**
> JobPoliciesExtended ListJobPolicies(ctx, optional)


List job impact policies.

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

[**JobPoliciesExtended**](JobPoliciesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJobJob**
> UpdateJobJob(ctx, jobJob, jobJobId)


Modify a running or paused job instance.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobJob** | [**JobJob**](JobJob.md)|  | 
  **jobJobId** | **string**| Modify a running or paused job instance.  All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJobPolicy**
> UpdateJobPolicy(ctx, jobPolicy, jobPolicyId)


Modify a job impact policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobPolicy** | [**JobPolicy**](JobPolicy.md)|  | 
  **jobPolicyId** | **string**| Modify a job impact policy. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateJobType**
> UpdateJobType(ctx, jobType, jobTypeId)


Modify the job type.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **jobType** | [**JobType**](JobType.md)|  | 
  **jobTypeId** | **string**| Modify the job type.  All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

