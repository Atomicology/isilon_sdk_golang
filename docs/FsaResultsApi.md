# \FsaResultsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistogramStatBy**](FsaResultsApi.md#GetHistogramStatBy) | **Get** /platform/3/fsa/results/{Id}/histogram/{Stat}/by | 
[**GetHistogramStatByBreakout**](FsaResultsApi.md#GetHistogramStatByBreakout) | **Get** /platform/3/fsa/results/{Id}/histogram/{Stat}/by/{HistogramStatByBreakout} | 
[**GetResultDirectories**](FsaResultsApi.md#GetResultDirectories) | **Get** /platform/3/fsa/results/{Id}/directories | 
[**GetResultDirectory**](FsaResultsApi.md#GetResultDirectory) | **Get** /platform/3/fsa/results/{Id}/directories/{ResultDirectoryId} | 
[**GetResultHistogram**](FsaResultsApi.md#GetResultHistogram) | **Get** /platform/3/fsa/results/{Id}/histogram | 
[**GetResultHistogramStat**](FsaResultsApi.md#GetResultHistogramStat) | **Get** /platform/3/fsa/results/{Id}/histogram/{ResultHistogramStat} | 
[**GetResultTopDir**](FsaResultsApi.md#GetResultTopDir) | **Get** /platform/3/fsa/results/{Id}/top-dirs/{ResultTopDirId} | 
[**GetResultTopDirs**](FsaResultsApi.md#GetResultTopDirs) | **Get** /platform/3/fsa/results/{Id}/top-dirs | 
[**GetResultTopFile**](FsaResultsApi.md#GetResultTopFile) | **Get** /platform/3/fsa/results/{Id}/top-files/{ResultTopFileId} | 
[**GetResultTopFiles**](FsaResultsApi.md#GetResultTopFiles) | **Get** /platform/3/fsa/results/{Id}/top-files | 


# **GetHistogramStatBy**
> HistogramStatBy GetHistogramStatBy(ctx, id, stat)


This resource retrieves a histogram breakout for an individual FSA result set. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
  **stat** | **string**|  | 

### Return type

[**HistogramStatBy**](HistogramStatBy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistogramStatByBreakout**
> HistogramStatBy GetHistogramStatByBreakout(ctx, histogramStatByBreakout, id, stat, optional)


This resource retrieves a histogram breakout for an individual FSA result set. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **histogramStatByBreakout** | **string**| This resource retrieves a histogram breakout for an individual FSA result set. ID in the resource path is the result set ID. | 
  **id** | **string**|  | 
  **stat** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **histogramStatByBreakout** | **string**| This resource retrieves a histogram breakout for an individual FSA result set. ID in the resource path is the result set ID. | 
 **id** | **string**|  | 
 **stat** | **string**|  | 
 **directoryFilter** | **string**| Filter according to a specific directory, which includes all of its subdirectories. | 
 **attributeFilter** | **string**| Filter according to the name of a file user attribute. | 
 **nodePoolFilter** | **string**| Filter according to the name of a node pool, which is a set of disk pools that belong to nodes of the same equivalence class. | 
 **diskPoolFilter** | **string**| Filter according to the name of a disk pool, which is a set of drives that represent an independent failure domain. | 
 **tierFilter** | **string**| Filter according to the name of a storage tier, which is a user-created set of node pools. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **logSizeFilter** | **int32**| Filter according to file logical size, where the filter value specifies the lower bound in bytes to a set of files that have been grouped by logical size. The list of valid log_size filter values may be found by performing a histogram breakout by log_size and viewing the resulting key values. | 
 **physSizeFilter** | **int32**| Filter according to file physical size, where the filter value specifies the lower bound in bytes to a set of files that have been grouped by physical size. The list of valid phys_size filter values may be found by performing a histogram breakout by phys_size and viewing the resulting key values. | 
 **limit** | **int32**| Limit the number of breakout results. | 
 **pathExtFilter** | **string**| Filter according to the name of a single file extension. | 
 **ctimeFilter** | **int32**| Filter according to file modified time, where the filter value specifies a negative number of seconds representing a time before the begin time of the report. The list of valid ctime filter values may be found by performing a histogram breakout by ctime and viewing the resulting key values. | 
 **atimeFilter** | **int32**| Filter according to file accessed time, where the filter value specifies a negative number of seconds representing a time before the begin time of the report. The list of valid atime filter values may be found by performing a histogram breakout by atime and viewing the resulting key values. | 

### Return type

[**HistogramStatBy**](HistogramStatBy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultDirectories**
> ResultDirectories GetResultDirectories(ctx, id, optional)


This resource retrieves directory information. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **path** | **string**| Primary directory path to report usage information, which may be specified instead of a LIN. | 
 **limit** | **int32**| Limit the number of reported subdirectories. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ResultDirectories**](ResultDirectories.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultDirectory**
> ResultDirectories GetResultDirectory(ctx, resultDirectoryId, id, optional)


This resource retrieves directory information. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resultDirectoryId** | **int32**| This resource retrieves directory information. ID in the resource path is the result set ID. | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultDirectoryId** | **int32**| This resource retrieves directory information. ID in the resource path is the result set ID. | 
 **id** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **limit** | **int32**| Limit the number of reported subdirectories. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ResultDirectories**](ResultDirectories.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultHistogram**
> ResultHistogram GetResultHistogram(ctx, id)


This resource retrieves a histogram of file counts for an individual FSA result set. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 

### Return type

[**ResultHistogram**](ResultHistogram.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultHistogramStat**
> ResultHistogram GetResultHistogramStat(ctx, resultHistogramStat, id, optional)


This resource retrieves a histogram of file counts for an individual FSA result set. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resultHistogramStat** | **string**| This resource retrieves a histogram of file counts for an individual FSA result set. ID in the resource path is the result set ID. | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultHistogramStat** | **string**| This resource retrieves a histogram of file counts for an individual FSA result set. ID in the resource path is the result set ID. | 
 **id** | **string**|  | 
 **directoryFilter** | **string**| Filter according to a specific directory, which includes all of its subdirectories. | 
 **attributeFilter** | **string**| Filter according to the name of a file user attribute. | 
 **nodePoolFilter** | **string**| Filter according to the name of a node pool, which is a set of disk pools that belong to nodes of the same equivalence class. | 
 **diskPoolFilter** | **string**| Filter according to the name of a disk pool, which is a set of drives that represent an independent failure domain. | 
 **tierFilter** | **string**| Filter according to the name of a storage tier, which is a user-created set of node pools. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **logSizeFilter** | **int32**| Filter according to file logical size, where the filter value specifies the lower bound in bytes to a set of files that have been grouped by logical size. The list of valid log_size filter values may be found by performing a histogram breakout by log_size and viewing the resulting key values. | 
 **physSizeFilter** | **int32**| Filter according to file physical size, where the filter value specifies the lower bound in bytes to a set of files that have been grouped by physical size. The list of valid phys_size filter values may be found by performing a histogram breakout by phys_size and viewing the resulting key values. | 
 **pathExtFilter** | **string**| Filter according to the name of a single file extension. | 
 **ctimeFilter** | **int32**| Filter according to file modified time, where the filter value specifies a negative number of seconds representing a time before the begin time of the report. The list of valid ctime filter values may be found by performing a histogram breakout by ctime and viewing the resulting key values. | 
 **atimeFilter** | **int32**| Filter according to file accessed time, where the filter value specifies a negative number of seconds representing a time before the begin time of the report. The list of valid atime filter values may be found by performing a histogram breakout by atime and viewing the resulting key values. | 

### Return type

[**ResultHistogram**](ResultHistogram.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultTopDir**
> ResultTopDirs GetResultTopDir(ctx, resultTopDirId, id, optional)


This resource retrieves the top directories. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resultTopDirId** | **string**| This resource retrieves the top directories. ID in the resource path is the result set ID. | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultTopDirId** | **string**| This resource retrieves the top directories. ID in the resource path is the result set ID. | 
 **id** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **start** | **int32**| Starting index for results. Default value of 0. | 
 **limit** | **int32**| Number of results from start index. Default value of 1000. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ResultTopDirs**](ResultTopDirs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultTopDirs**
> ResultTopDirs GetResultTopDirs(ctx, id)


This resource retrieves the top directories. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 

### Return type

[**ResultTopDirs**](ResultTopDirs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultTopFile**
> ResultTopFiles GetResultTopFile(ctx, resultTopFileId, id, optional)


This resource retrieves the top files. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **resultTopFileId** | **string**| This resource retrieves the top files. ID in the resource path is the result set ID. | 
  **id** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultTopFileId** | **string**| This resource retrieves the top files. ID in the resource path is the result set ID. | 
 **id** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting. | 
 **start** | **int32**| Starting index for results. Default value of 0. | 
 **limit** | **int32**| Number of results from start index. Default value of 1000. | 
 **compReport** | **int32**| Result set identifier for comparison of database results. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ResultTopFiles**](ResultTopFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResultTopFiles**
> ResultTopFiles GetResultTopFiles(ctx, id)


This resource retrieves the top files. ID in the resource path is the result set ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**|  | 

### Return type

[**ResultTopFiles**](ResultTopFiles.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

