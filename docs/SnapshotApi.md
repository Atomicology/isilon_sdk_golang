# \SnapshotApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshotAlias**](SnapshotApi.md#CreateSnapshotAlias) | **Post** /platform/1/snapshot/aliases | 
[**CreateSnapshotChangelist**](SnapshotApi.md#CreateSnapshotChangelist) | **Post** /platform/1/snapshot/changelists | 
[**CreateSnapshotRepstate**](SnapshotApi.md#CreateSnapshotRepstate) | **Post** /platform/1/snapshot/repstates | 
[**CreateSnapshotSchedule**](SnapshotApi.md#CreateSnapshotSchedule) | **Post** /platform/3/snapshot/schedules | 
[**CreateSnapshotSnapshot**](SnapshotApi.md#CreateSnapshotSnapshot) | **Post** /platform/1/snapshot/snapshots | 
[**DeleteSnapshotAlias**](SnapshotApi.md#DeleteSnapshotAlias) | **Delete** /platform/1/snapshot/aliases/{SnapshotAliasId} | 
[**DeleteSnapshotAliases**](SnapshotApi.md#DeleteSnapshotAliases) | **Delete** /platform/1/snapshot/aliases | 
[**DeleteSnapshotChangelist**](SnapshotApi.md#DeleteSnapshotChangelist) | **Delete** /platform/1/snapshot/changelists/{SnapshotChangelistId} | 
[**DeleteSnapshotRepstate**](SnapshotApi.md#DeleteSnapshotRepstate) | **Delete** /platform/1/snapshot/repstates/{SnapshotRepstateId} | 
[**DeleteSnapshotSchedule**](SnapshotApi.md#DeleteSnapshotSchedule) | **Delete** /platform/3/snapshot/schedules/{SnapshotScheduleId} | 
[**DeleteSnapshotSchedules**](SnapshotApi.md#DeleteSnapshotSchedules) | **Delete** /platform/3/snapshot/schedules | 
[**DeleteSnapshotSnapshot**](SnapshotApi.md#DeleteSnapshotSnapshot) | **Delete** /platform/1/snapshot/snapshots/{SnapshotSnapshotId} | 
[**DeleteSnapshotSnapshots**](SnapshotApi.md#DeleteSnapshotSnapshots) | **Delete** /platform/1/snapshot/snapshots | 
[**GetSnapshotAlias**](SnapshotApi.md#GetSnapshotAlias) | **Get** /platform/1/snapshot/aliases/{SnapshotAliasId} | 
[**GetSnapshotChangelist**](SnapshotApi.md#GetSnapshotChangelist) | **Get** /platform/1/snapshot/changelists/{SnapshotChangelistId} | 
[**GetSnapshotLicense**](SnapshotApi.md#GetSnapshotLicense) | **Get** /platform/5/snapshot/license | 
[**GetSnapshotPending**](SnapshotApi.md#GetSnapshotPending) | **Get** /platform/1/snapshot/pending | 
[**GetSnapshotRepstate**](SnapshotApi.md#GetSnapshotRepstate) | **Get** /platform/1/snapshot/repstates/{SnapshotRepstateId} | 
[**GetSnapshotSchedule**](SnapshotApi.md#GetSnapshotSchedule) | **Get** /platform/3/snapshot/schedules/{SnapshotScheduleId} | 
[**GetSnapshotSettings**](SnapshotApi.md#GetSnapshotSettings) | **Get** /platform/1/snapshot/settings | 
[**GetSnapshotSnapshot**](SnapshotApi.md#GetSnapshotSnapshot) | **Get** /platform/1/snapshot/snapshots/{SnapshotSnapshotId} | 
[**GetSnapshotSnapshotsSummary**](SnapshotApi.md#GetSnapshotSnapshotsSummary) | **Get** /platform/1/snapshot/snapshots-summary | 
[**ListSnapshotAliases**](SnapshotApi.md#ListSnapshotAliases) | **Get** /platform/1/snapshot/aliases | 
[**ListSnapshotChangelists**](SnapshotApi.md#ListSnapshotChangelists) | **Get** /platform/1/snapshot/changelists | 
[**ListSnapshotRepstates**](SnapshotApi.md#ListSnapshotRepstates) | **Get** /platform/1/snapshot/repstates | 
[**ListSnapshotSchedules**](SnapshotApi.md#ListSnapshotSchedules) | **Get** /platform/3/snapshot/schedules | 
[**ListSnapshotSnapshots**](SnapshotApi.md#ListSnapshotSnapshots) | **Get** /platform/1/snapshot/snapshots | 
[**UpdateSnapshotAlias**](SnapshotApi.md#UpdateSnapshotAlias) | **Put** /platform/1/snapshot/aliases/{SnapshotAliasId} | 
[**UpdateSnapshotSchedule**](SnapshotApi.md#UpdateSnapshotSchedule) | **Put** /platform/3/snapshot/schedules/{SnapshotScheduleId} | 
[**UpdateSnapshotSettings**](SnapshotApi.md#UpdateSnapshotSettings) | **Put** /platform/1/snapshot/settings | 
[**UpdateSnapshotSnapshot**](SnapshotApi.md#UpdateSnapshotSnapshot) | **Put** /platform/1/snapshot/snapshots/{SnapshotSnapshotId} | 


# **CreateSnapshotAlias**
> CreateSnapshotAliasResponse CreateSnapshotAlias(ctx, snapshotAlias)


Create a new snapshot alias.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotAlias** | [**SnapshotAliasCreateParams**](SnapshotAliasCreateParams.md)|  | 

### Return type

[**CreateSnapshotAliasResponse**](CreateSnapshotAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSnapshotChangelist**
> CreateSnapshotChangelistResponse CreateSnapshotChangelist(ctx, snapshotChangelist)


Create a new changelist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotChangelist** | [**SnapshotChangelists**](SnapshotChangelists.md)|  | 

### Return type

[**CreateSnapshotChangelistResponse**](CreateSnapshotChangelistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSnapshotRepstate**
> CreateSnapshotRepstateResponse CreateSnapshotRepstate(ctx, snapshotRepstate)


Create a new repstates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotRepstate** | [**SnapshotRepstates**](SnapshotRepstates.md)|  | 

### Return type

[**CreateSnapshotRepstateResponse**](CreateSnapshotRepstateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSnapshotSchedule**
> CreateSnapshotScheduleResponse CreateSnapshotSchedule(ctx, snapshotSchedule)


Create a new schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSchedule** | [**SnapshotScheduleCreateParams**](SnapshotScheduleCreateParams.md)|  | 

### Return type

[**CreateSnapshotScheduleResponse**](CreateSnapshotScheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSnapshotSnapshot**
> SnapshotSnapshotExtended CreateSnapshotSnapshot(ctx, snapshotSnapshot)


Create a new snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSnapshot** | [**SnapshotSnapshotCreateParams**](SnapshotSnapshotCreateParams.md)|  | 

### Return type

[**SnapshotSnapshotExtended**](SnapshotSnapshotExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotAlias**
> DeleteSnapshotAlias(ctx, snapshotAliasId)


Delete the snapshot alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotAliasId** | **string**| Delete the snapshot alias | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotAliases**
> DeleteSnapshotAliases(ctx, )


Delete all or matching snapshot aliases.

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

# **DeleteSnapshotChangelist**
> DeleteSnapshotChangelist(ctx, snapshotChangelistId)


Delete the specified changelist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotChangelistId** | **string**| Delete the specified changelist. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotRepstate**
> DeleteSnapshotRepstate(ctx, snapshotRepstateId)


Delete the specified repstate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotRepstateId** | **string**| Delete the specified repstate. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotSchedule**
> DeleteSnapshotSchedule(ctx, snapshotScheduleId)


Delete the schedule. This does not affect already created snapshots.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotScheduleId** | **string**| Delete the schedule. This does not affect already created snapshots. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotSchedules**
> DeleteSnapshotSchedules(ctx, )


Delete all snapshot schedules.

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

# **DeleteSnapshotSnapshot**
> DeleteSnapshotSnapshot(ctx, snapshotSnapshotId)


Delete the snapshot. Deleted snapshots will be placed into a deleting state until the system can reclaim the space used by the snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSnapshotId** | **string**| Delete the snapshot. Deleted snapshots will be placed into a deleting state until the system can reclaim the space used by the snapshot. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnapshotSnapshots**
> DeleteSnapshotSnapshots(ctx, optional)


Delete all or matching snapshots.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| Only list snapshots matching this type. | 
 **schedule** | **string**| Only list snapshots created by this schedule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotAlias**
> SnapshotAliases GetSnapshotAlias(ctx, snapshotAliasId)


Retrieve snapshot alias information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotAliasId** | **string**| Retrieve snapshot alias information. | 

### Return type

[**SnapshotAliases**](SnapshotAliases.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotChangelist**
> SnapshotChangelists GetSnapshotChangelist(ctx, snapshotChangelistId, optional)


Retrieve basic information on a changelist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotChangelistId** | **string**| Retrieve basic information on a changelist. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotChangelistId** | **string**| Retrieve basic information on a changelist. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotChangelists**](SnapshotChangelists.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotLicense**
> LicenseLicense GetSnapshotLicense(ctx, )


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

# **GetSnapshotPending**
> SnapshotPending GetSnapshotPending(ctx, optional)


Return list of snapshots to be taken.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many result at once (see resume). | 
 **begin** | **int32**| Unix Epoch time to start generating matches. Default is now. | 
 **schedule** | **string**| Limit output only to the named schedule. | 
 **end** | **int32**| Unix Epoch time to end generating matches. Default is forever. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotPending**](SnapshotPending.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotRepstate**
> SnapshotRepstates GetSnapshotRepstate(ctx, snapshotRepstateId, optional)


Retrieve basic information on a repstate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotRepstateId** | **string**| Retrieve basic information on a repstate. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **snapshotRepstateId** | **string**| Retrieve basic information on a repstate. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotRepstates**](SnapshotRepstates.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotSchedule**
> SnapshotSchedules GetSnapshotSchedule(ctx, snapshotScheduleId)


Retrieve the schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotScheduleId** | **string**| Retrieve the schedule. | 

### Return type

[**SnapshotSchedules**](SnapshotSchedules.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotSettings**
> SnapshotSettings GetSnapshotSettings(ctx, )


List all settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SnapshotSettings**](SnapshotSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotSnapshot**
> SnapshotSnapshots GetSnapshotSnapshot(ctx, snapshotSnapshotId)


Retrieve snapshot information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSnapshotId** | **string**| Retrieve snapshot information. | 

### Return type

[**SnapshotSnapshots**](SnapshotSnapshots.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotSnapshotsSummary**
> SnapshotSnapshotsSummary GetSnapshotSnapshotsSummary(ctx, )


Return summary information about snapshots.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SnapshotSnapshotsSummary**](SnapshotSnapshotsSummary.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotAliases**
> SnapshotAliasesExtended ListSnapshotAliases(ctx, optional)


List all or matching snapshot aliases.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting.  Choices are id, name, snapshot, and created.  Default is id. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotAliasesExtended**](SnapshotAliasesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotChangelists**
> SnapshotChangelistsExtended ListSnapshotChangelists(ctx, optional)


List all changelists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotChangelistsExtended**](SnapshotChangelistsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotRepstates**
> SnapshotRepstatesExtended ListSnapshotRepstates(ctx, optional)


List all repstates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotRepstatesExtended**](SnapshotRepstatesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotSchedules**
> SnapshotSchedulesExtended ListSnapshotSchedules(ctx, optional)


List all or matching schedules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting.  Choices are id, name, path, pattern, schedule, duration, alias, next_run, and next_snapshot.  Default is id. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 

### Return type

[**SnapshotSchedulesExtended**](SnapshotSchedulesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotSnapshots**
> SnapshotSnapshotsExtended ListSnapshotSnapshots(ctx, optional)


List all or matching snapshots.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string**| The field that will be used for sorting.  Choices are id, name, path, created, expires, size, has_locks, schedule, alias_target, alias_target_name, pct_filesystem, pct_reserve, and state.  Default is id. | 
 **schedule** | **string**| Only list snapshots created by this schedule. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **state** | **string**| Only list snapshots matching this state. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **type_** | **string**| Only list snapshots matching this type. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**SnapshotSnapshotsExtended**](SnapshotSnapshotsExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshotAlias**
> UpdateSnapshotAlias(ctx, snapshotAlias, snapshotAliasId)


Modify snapshot alias. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotAlias** | [**SnapshotAlias**](SnapshotAlias.md)|  | 
  **snapshotAliasId** | **string**| Modify snapshot alias. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshotSchedule**
> UpdateSnapshotSchedule(ctx, snapshotSchedule, snapshotScheduleId)


Modify the schedule. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSchedule** | [**SnapshotSchedule**](SnapshotSchedule.md)|  | 
  **snapshotScheduleId** | **string**| Modify the schedule. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshotSettings**
> UpdateSnapshotSettings(ctx, snapshotSettings)


Modify one or more settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSettings** | [**SnapshotSettingsExtended**](SnapshotSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnapshotSnapshot**
> UpdateSnapshotSnapshot(ctx, snapshotSnapshot, snapshotSnapshotId)


Modify snapshot. All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **snapshotSnapshot** | [**SnapshotSnapshot**](SnapshotSnapshot.md)|  | 
  **snapshotSnapshotId** | **string**| Modify snapshot. All input fields are optional, but one or more must be supplied. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

