# \ClusterApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterAddNodeItem**](ClusterApi.md#CreateClusterAddNodeItem) | **Post** /platform/3/cluster/add-node | 
[**CreateDiagnosticsGatherStartItem**](ClusterApi.md#CreateDiagnosticsGatherStartItem) | **Post** /platform/3/cluster/diagnostics/gather/start | 
[**CreateDiagnosticsGatherStopItem**](ClusterApi.md#CreateDiagnosticsGatherStopItem) | **Post** /platform/3/cluster/diagnostics/gather/stop | 
[**CreateDiagnosticsNetloggerStartItem**](ClusterApi.md#CreateDiagnosticsNetloggerStartItem) | **Post** /platform/3/cluster/diagnostics/netlogger/start | 
[**CreateDiagnosticsNetloggerStopItem**](ClusterApi.md#CreateDiagnosticsNetloggerStopItem) | **Post** /platform/3/cluster/diagnostics/netlogger/stop | 
[**GetClusterConfig**](ClusterApi.md#GetClusterConfig) | **Get** /platform/3/cluster/config | 
[**GetClusterEmail**](ClusterApi.md#GetClusterEmail) | **Get** /platform/1/cluster/email | 
[**GetClusterExternalIps**](ClusterApi.md#GetClusterExternalIps) | **Get** /platform/2/cluster/external-ips | 
[**GetClusterIdentity**](ClusterApi.md#GetClusterIdentity) | **Get** /platform/5/cluster/identity | 
[**GetClusterNode**](ClusterApi.md#GetClusterNode) | **Get** /platform/5/cluster/nodes/{ClusterNodeId} | 
[**GetClusterNodes**](ClusterApi.md#GetClusterNodes) | **Get** /platform/5/cluster/nodes | 
[**GetClusterNodesAvailable**](ClusterApi.md#GetClusterNodesAvailable) | **Get** /platform/3/cluster/nodes-available | 
[**GetClusterOwner**](ClusterApi.md#GetClusterOwner) | **Get** /platform/1/cluster/owner | 
[**GetClusterStatfs**](ClusterApi.md#GetClusterStatfs) | **Get** /platform/1/cluster/statfs | 
[**GetClusterTime**](ClusterApi.md#GetClusterTime) | **Get** /platform/3/cluster/time | 
[**GetClusterTimezone**](ClusterApi.md#GetClusterTimezone) | **Get** /platform/3/cluster/timezone | 
[**GetClusterVersion**](ClusterApi.md#GetClusterVersion) | **Get** /platform/3/cluster/version | 
[**GetDiagnosticsGather**](ClusterApi.md#GetDiagnosticsGather) | **Get** /platform/3/cluster/diagnostics/gather | 
[**GetDiagnosticsGatherSettings**](ClusterApi.md#GetDiagnosticsGatherSettings) | **Get** /platform/3/cluster/diagnostics/gather/settings | 
[**GetDiagnosticsGatherStatus**](ClusterApi.md#GetDiagnosticsGatherStatus) | **Get** /platform/3/cluster/diagnostics/gather/status | 
[**GetDiagnosticsNetlogger**](ClusterApi.md#GetDiagnosticsNetlogger) | **Get** /platform/3/cluster/diagnostics/netlogger | 
[**GetDiagnosticsNetloggerSettings**](ClusterApi.md#GetDiagnosticsNetloggerSettings) | **Get** /platform/3/cluster/diagnostics/netlogger/settings | 
[**GetDiagnosticsNetloggerStatus**](ClusterApi.md#GetDiagnosticsNetloggerStatus) | **Get** /platform/3/cluster/diagnostics/netlogger/status | 
[**GetTimezoneRegion**](ClusterApi.md#GetTimezoneRegion) | **Get** /platform/3/cluster/timezone/regions/{TimezoneRegionId} | 
[**GetTimezoneSettings**](ClusterApi.md#GetTimezoneSettings) | **Get** /platform/3/cluster/timezone/settings | 
[**UpdateClusterEmail**](ClusterApi.md#UpdateClusterEmail) | **Put** /platform/1/cluster/email | 
[**UpdateClusterNode**](ClusterApi.md#UpdateClusterNode) | **Put** /platform/5/cluster/nodes/{ClusterNodeId} | 
[**UpdateClusterOwner**](ClusterApi.md#UpdateClusterOwner) | **Put** /platform/1/cluster/owner | 
[**UpdateClusterTime**](ClusterApi.md#UpdateClusterTime) | **Put** /platform/3/cluster/time | 
[**UpdateClusterTimezone**](ClusterApi.md#UpdateClusterTimezone) | **Put** /platform/3/cluster/timezone | 
[**UpdateDiagnosticsGatherSettings**](ClusterApi.md#UpdateDiagnosticsGatherSettings) | **Put** /platform/3/cluster/diagnostics/gather/settings | 
[**UpdateDiagnosticsNetloggerSettings**](ClusterApi.md#UpdateDiagnosticsNetloggerSettings) | **Put** /platform/3/cluster/diagnostics/netlogger/settings | 
[**UpdateTimezoneSettings**](ClusterApi.md#UpdateTimezoneSettings) | **Put** /platform/3/cluster/timezone/settings | 


# **CreateClusterAddNodeItem**
> Empty CreateClusterAddNodeItem(ctx, clusterAddNodeItem)


Serial number and arguments of node to add.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterAddNodeItem** | [**ClusterAddNodeItem**](ClusterAddNodeItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiagnosticsGatherStartItem**
> Empty CreateDiagnosticsGatherStartItem(ctx, diagnosticsGatherStartItem)


Start a new gather

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsGatherStartItem** | [**DiagnosticsGatherSettingsExtended**](DiagnosticsGatherSettingsExtended.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiagnosticsGatherStopItem**
> Empty CreateDiagnosticsGatherStopItem(ctx, diagnosticsGatherStopItem)


Stop a running gather

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsGatherStopItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiagnosticsNetloggerStartItem**
> Empty CreateDiagnosticsNetloggerStartItem(ctx, diagnosticsNetloggerStartItem)


Start a new packet caputre

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsNetloggerStartItem** | [**DiagnosticsNetloggerSettings**](DiagnosticsNetloggerSettings.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiagnosticsNetloggerStopItem**
> Empty CreateDiagnosticsNetloggerStopItem(ctx, diagnosticsNetloggerStopItem)


Stop a running packet capture

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsNetloggerStopItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterConfig**
> ClusterConfig GetClusterConfig(ctx, )


Retrieve the cluster information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterConfig**](ClusterConfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterEmail**
> ClusterEmail GetClusterEmail(ctx, )


Get the cluster email notification settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterEmail**](ClusterEmail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterExternalIps**
> []string GetClusterExternalIps(ctx, )


Retrieve the cluster IP addresses including IPV4 and IPV6.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterIdentity**
> ClusterIdentity GetClusterIdentity(ctx, )


Retrieve the login information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterIdentity**](ClusterIdentity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNode**
> ClusterNodesExtendedExtended GetClusterNode(ctx, clusterNodeId, optional)


Retrieve node information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterNodeId** | **int32**| Retrieve node information. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterNodeId** | **int32**| Retrieve node information. | 
 **timeout** | **float32**| Request timeout | 

### Return type

[**ClusterNodesExtendedExtended**](ClusterNodesExtendedExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodes**
> ClusterNodesExtendedExtendedExtended GetClusterNodes(ctx, optional)


List the nodes on this cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeout** | **float32**| Request timeout | 

### Return type

[**ClusterNodesExtendedExtendedExtended**](ClusterNodesExtendedExtendedExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodesAvailable**
> ClusterNodesAvailable GetClusterNodesAvailable(ctx, )


List all nodes that are available to add to this cluster.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterNodesAvailable**](ClusterNodesAvailable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterOwner**
> ClusterOwner GetClusterOwner(ctx, )


Get the cluster contact info settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterOwner**](ClusterOwner.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterStatfs**
> ClusterStatfs GetClusterStatfs(ctx, )


Retrieve the filesystem statistics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterStatfs**](ClusterStatfs.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterTime**
> ClusterTime GetClusterTime(ctx, )


Retrieve the current time as reported by each node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterTime**](ClusterTime.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterTimezone**
> ClusterTimezone GetClusterTimezone(ctx, )


Get the cluster timezone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterTimezone**](ClusterTimezone.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterVersion**
> ClusterVersion GetClusterVersion(ctx, )


Retrieve the OneFS version as reported by each node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterVersion**](ClusterVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsGather**
> DiagnosticsGatherStatus GetDiagnosticsGather(ctx, )


Get the status of isi_gather_info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsGatherStatus**](DiagnosticsGatherStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsGatherSettings**
> DiagnosticsGatherSettings GetDiagnosticsGatherSettings(ctx, )


Get the default options for isi_gather_info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsGatherSettings**](DiagnosticsGatherSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsGatherStatus**
> DiagnosticsGatherStatus GetDiagnosticsGatherStatus(ctx, )


Get the status of isi_gather_info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsGatherStatus**](DiagnosticsGatherStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsNetlogger**
> DiagnosticsNetloggerStatus GetDiagnosticsNetlogger(ctx, )


Get the status of isi_netlogger.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsNetloggerStatus**](DiagnosticsNetloggerStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsNetloggerSettings**
> DiagnosticsNetloggerSettings GetDiagnosticsNetloggerSettings(ctx, )


Get the default options for isi_netlogger.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsNetloggerSettings**](DiagnosticsNetloggerSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDiagnosticsNetloggerStatus**
> DiagnosticsNetloggerStatus GetDiagnosticsNetloggerStatus(ctx, )


Get the status of isi_netlogger.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticsNetloggerStatus**](DiagnosticsNetloggerStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimezoneRegion**
> TimezoneRegions GetTimezoneRegion(ctx, timezoneRegionId, optional)


List timezone regions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **timezoneRegionId** | **string**| List timezone regions. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timezoneRegionId** | **string**| List timezone regions. | 
 **sort** | **string**| The field that will be used for sorting. | 
 **resume** | **string**| Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options). | 
 **showAll** | **bool**| Show all timezones within the region specified in the URI. | 
 **dstReset** | **bool**| This query arg is not needed in normal use cases. | 
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**TimezoneRegions**](TimezoneRegions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimezoneSettings**
> TimezoneSettings GetTimezoneSettings(ctx, )


Retrieve the cluster timezone.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TimezoneSettings**](TimezoneSettings.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterEmail**
> UpdateClusterEmail(ctx, clusterEmail)


Modify the cluster email notification settings.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterEmail** | [**ClusterEmailExtended**](ClusterEmailExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterNode**
> UpdateClusterNode(ctx, clusterNode, clusterNodeId)


Modify one or more node settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterNode** | [**ClusterNode**](ClusterNode.md)|  | 
  **clusterNodeId** | **int32**| Modify one or more node settings. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterOwner**
> UpdateClusterOwner(ctx, clusterOwner)


Modify the cluster contact info settings.  All input fields are optional, but one or more must be supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterOwner** | [**ClusterOwner**](ClusterOwner.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterTime**
> UpdateClusterTime(ctx, clusterTime)


Set cluster time.  Time will mostly be synchronized across nodes, but there may be slight drift.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterTime** | [**ClusterTimeExtended**](ClusterTimeExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterTimezone**
> UpdateClusterTimezone(ctx, clusterTimezone)


Set a new timezone for the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterTimezone** | [**ClusterTimezoneExtended**](ClusterTimezoneExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDiagnosticsGatherSettings**
> UpdateDiagnosticsGatherSettings(ctx, diagnosticsGatherSettings)


Set the default options for isi_gather_info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsGatherSettings** | [**DiagnosticsGatherSettingsExtended**](DiagnosticsGatherSettingsExtended.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDiagnosticsNetloggerSettings**
> UpdateDiagnosticsNetloggerSettings(ctx, diagnosticsNetloggerSettings)


Set the default options for isi_netlogger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **diagnosticsNetloggerSettings** | [**DiagnosticsNetloggerSettings**](DiagnosticsNetloggerSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTimezoneSettings**
> UpdateTimezoneSettings(ctx, timezoneSettings)


Modify the cluster timezone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **timezoneSettings** | [**TimezoneRegionTimezone**](TimezoneRegionTimezone.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

