# \SnapshotSnapshotsApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshotLock**](SnapshotSnapshotsApi.md#CreateSnapshotLock) | **Post** /platform/1/snapshot/snapshots/{Sid}/locks | 
[**DeleteSnapshotLock**](SnapshotSnapshotsApi.md#DeleteSnapshotLock) | **Delete** /platform/1/snapshot/snapshots/{Sid}/locks/{SnapshotLockId} | 
[**DeleteSnapshotLocks**](SnapshotSnapshotsApi.md#DeleteSnapshotLocks) | **Delete** /platform/1/snapshot/snapshots/{Sid}/locks | 
[**GetSnapshotLock**](SnapshotSnapshotsApi.md#GetSnapshotLock) | **Get** /platform/1/snapshot/snapshots/{Sid}/locks/{SnapshotLockId} | 
[**ListSnapshotLocks**](SnapshotSnapshotsApi.md#ListSnapshotLocks) | **Get** /platform/1/snapshot/snapshots/{Sid}/locks | 
[**UpdateSnapshotLock**](SnapshotSnapshotsApi.md#UpdateSnapshotLock) | **Put** /platform/1/snapshot/snapshots/{Sid}/locks/{SnapshotLockId} | 


# **CreateSnapshotLock**
> CreateSnapshotLockResponse CreateSnapshotLock(ctx, snapshotLock, sid)


Create a new lock on this snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotLock** | [**SnapshotLockCreateParams**](SnapshotLockCreateParams.md)|  | 
  **sid** | **string**|  | 

### Return type

[**CreateSnapshotLockResponse**](CreateSnapshotLockResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotLock**
> DeleteSnapshotLock(ctx, snapshotLockId, sid)


Delete the snapshot lock.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotLockId** | **string**| Delete the snapshot lock. | 
  **sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotLocks**
> DeleteSnapshotLocks(ctx, sid)


Delete all locks. Will try to drain count of recursively held locks so that the snapshot can be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotLock**
> SnapshotLocks GetSnapshotLock(ctx, snapshotLockId, sid)


Retrieve lock information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotLockId** | **string**| Retrieve lock information. | 
  **sid** | **string**|  | 

### Return type

[**SnapshotLocks**](SnapshotLocks.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotLocks**
> SnapshotLocksExtended ListSnapshotLocks(ctx, sid, optional)


List all locks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sid** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sid** | **string**|  | 
 **sort** | **string**| The field that will be used for sorting.  Choices are id, expires, and comment.  Default is id. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotLocksExtended**](SnapshotLocksExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshotLock**
> UpdateSnapshotLock(ctx, snapshotLock, snapshotLockId, sid)


Modify lock. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotLock** | [**SnapshotLock**](SnapshotLock.md)|  | 
  **snapshotLockId** | **string**| Modify lock. All input fields are optional, but one or more must be supplied. | 
  **sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

