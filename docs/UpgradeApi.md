# \UpgradeApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterAddRemainingNode**](UpgradeApi.md#CreateClusterAddRemainingNode) | **Post** /platform/3/upgrade/cluster/add_remaining_nodes | 
[**CreateClusterArchiveItem**](UpgradeApi.md#CreateClusterArchiveItem) | **Post** /platform/3/upgrade/cluster/archive | 
[**CreateClusterAssessItem**](UpgradeApi.md#CreateClusterAssessItem) | **Post** /platform/5/upgrade/cluster/assess | 
[**CreateClusterCommitItem**](UpgradeApi.md#CreateClusterCommitItem) | **Post** /platform/3/upgrade/cluster/commit | 
[**CreateClusterFirmwareAssessItem**](UpgradeApi.md#CreateClusterFirmwareAssessItem) | **Post** /platform/3/upgrade/cluster/firmware/assess | 
[**CreateClusterFirmwareUpgradeItem**](UpgradeApi.md#CreateClusterFirmwareUpgradeItem) | **Post** /platform/3/upgrade/cluster/firmware/upgrade | 
[**CreateClusterPatchAbortItem**](UpgradeApi.md#CreateClusterPatchAbortItem) | **Post** /platform/3/upgrade/cluster/patch/abort | 
[**CreateClusterPatchPatch**](UpgradeApi.md#CreateClusterPatchPatch) | **Post** /platform/3/upgrade/cluster/patch/patches | 
[**CreateClusterRetryLastActionItem**](UpgradeApi.md#CreateClusterRetryLastActionItem) | **Post** /platform/3/upgrade/cluster/retry_last_action | 
[**CreateClusterRollbackItem**](UpgradeApi.md#CreateClusterRollbackItem) | **Post** /platform/3/upgrade/cluster/rollback | 
[**CreateClusterUpgradeItem**](UpgradeApi.md#CreateClusterUpgradeItem) | **Post** /platform/5/upgrade/cluster/upgrade | 
[**CreateHardwareStartItem**](UpgradeApi.md#CreateHardwareStartItem) | **Post** /platform/5/upgrade/hardware/start | 
[**CreateHardwareStopItem**](UpgradeApi.md#CreateHardwareStopItem) | **Post** /platform/5/upgrade/hardware/stop | 
[**DeleteClusterPatchPatch**](UpgradeApi.md#DeleteClusterPatchPatch) | **Delete** /platform/3/upgrade/cluster/patch/patches/{ClusterPatchPatchId} | 
[**GetClusterFirmwareProgress**](UpgradeApi.md#GetClusterFirmwareProgress) | **Get** /platform/3/upgrade/cluster/firmware/progress | 
[**GetClusterFirmwareStatus**](UpgradeApi.md#GetClusterFirmwareStatus) | **Get** /platform/3/upgrade/cluster/firmware/status | 
[**GetClusterNode**](UpgradeApi.md#GetClusterNode) | **Get** /platform/3/upgrade/cluster/nodes/{ClusterNodeId} | 
[**GetClusterNodes**](UpgradeApi.md#GetClusterNodes) | **Get** /platform/3/upgrade/cluster/nodes | 
[**GetClusterPatchPatch**](UpgradeApi.md#GetClusterPatchPatch) | **Get** /platform/3/upgrade/cluster/patch/patches/{ClusterPatchPatchId} | 
[**GetHardwareStatus**](UpgradeApi.md#GetHardwareStatus) | **Get** /platform/5/upgrade/hardware/status | 
[**GetUpgradeCluster**](UpgradeApi.md#GetUpgradeCluster) | **Get** /platform/4/upgrade/cluster | 
[**ListClusterPatchPatches**](UpgradeApi.md#ListClusterPatchPatches) | **Get** /platform/3/upgrade/cluster/patch/patches | 
[**UpdateClusterUpgrade**](UpgradeApi.md#UpdateClusterUpgrade) | **Put** /platform/5/upgrade/cluster/upgrade | 


# **CreateClusterAddRemainingNode**
> Empty CreateClusterAddRemainingNode(ctx, clusterAddRemainingNode)


Let system absorb any remaining or new nodes inside the existing upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterAddRemainingNode** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterArchiveItem**
> Empty CreateClusterArchiveItem(ctx, clusterArchiveItem)


Start an archive of an upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterArchiveItem** | [**ClusterArchiveItem**](ClusterArchiveItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterAssessItem**
> Empty CreateClusterAssessItem(ctx, clusterAssessItem)


Start upgrade assessment on cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterAssessItem** | [**ClusterAssessItem**](ClusterAssessItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterCommitItem**
> Empty CreateClusterCommitItem(ctx, clusterCommitItem)


Commit the upgrade of a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterCommitItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterFirmwareAssessItem**
> Empty CreateClusterFirmwareAssessItem(ctx, clusterFirmwareAssessItem)


Start firmware upgrade assessment on cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterFirmwareAssessItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterFirmwareUpgradeItem**
> Empty CreateClusterFirmwareUpgradeItem(ctx, clusterFirmwareUpgradeItem)


The settings necessary to start a firmware upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterFirmwareUpgradeItem** | [**ClusterFirmwareUpgradeItem**](ClusterFirmwareUpgradeItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterPatchAbortItem**
> Empty CreateClusterPatchAbortItem(ctx, clusterPatchAbortItem)


Abort the previous action performed by the patch system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterPatchAbortItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterPatchPatch**
> CreateResponse CreateClusterPatchPatch(ctx, clusterPatchPatch, optional)


Install a patch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterPatchPatch** | [**ClusterPatchPatch**](ClusterPatchPatch.md)|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterPatchPatch** | [**ClusterPatchPatch**](ClusterPatchPatch.md)|  | 
 **override** | **bool**| Whether to ignore patch system validation and force the installation. | 
 **rolling** | **bool**| Whether to install the patch on one node at a time. Defaults to true. | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterRetryLastActionItem**
> Empty CreateClusterRetryLastActionItem(ctx, clusterRetryLastActionItem)


Retry the last upgrade action, in-case the previous attempt failed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterRetryLastActionItem** | [**ClusterRetryLastActionItem**](ClusterRetryLastActionItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterRollbackItem**
> Empty CreateClusterRollbackItem(ctx, clusterRollbackItem)


Rollback the upgrade of a cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterRollbackItem** | [**Empty**](Empty.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterUpgradeItem**
> Empty CreateClusterUpgradeItem(ctx, clusterUpgradeItem)


The settings necessary to start an upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterUpgradeItem** | [**ClusterUpgradeItem**](ClusterUpgradeItem.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHardwareStartItem**
> CreateHardwareStartItem(ctx, hardwareStartItem)


Start a hardware upgrade

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareStartItem** | [**HardwareStartItem**](HardwareStartItem.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHardwareStopItem**
> CreateHardwareStopItem(ctx, hardwareStopItem)


Stop an in-progess hardware upgrade process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **hardwareStopItem** | [**HardwareStopItem**](HardwareStopItem.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterPatchPatch**
> DeleteClusterPatchPatch(ctx, clusterPatchPatchId, optional)


Uninstall a patch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterPatchPatchId** | **string**| Uninstall a patch. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterPatchPatchId** | **string**| Uninstall a patch. | 
 **override** | **bool**| Whether to ignore patch system validation and force the uninstallation. | 
 **rolling** | **bool**| Whether to uninstall the patch on one node at a time. Defaults to true. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterFirmwareProgress**
> ClusterFirmwareProgress GetClusterFirmwareProgress(ctx, )


Cluster wide firmware upgrade status info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterFirmwareProgress**](ClusterFirmwareProgress.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterFirmwareStatus**
> ClusterFirmwareStatus GetClusterFirmwareStatus(ctx, optional)


The firmware status for the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devices** | **bool**| Show devices. If false, this returns an empty list. Default is false. | 
 **package_** | **bool**| Show package. If false, this returns an empty list.Default is false. | 

### Return type

[**ClusterFirmwareStatus**](ClusterFirmwareStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNode**
> ClusterNodes GetClusterNode(ctx, clusterNodeId)


The node details useful during an upgrade or assessment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterNodeId** | **int32**| The node details useful during an upgrade or assessment. | 

### Return type

[**ClusterNodes**](ClusterNodes.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodes**
> ClusterNodesExtended GetClusterNodes(ctx, )


View information about nodes during an upgrade, rollback, or pre-upgrade assessment.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterNodesExtended**](ClusterNodesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterPatchPatch**
> ClusterPatchPatches GetClusterPatchPatch(ctx, clusterPatchPatchId, optional)


View a single patch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterPatchPatchId** | **string**| View a single patch. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterPatchPatchId** | **string**| View a single patch. | 
 **local** | **bool**| Only view patch information on the local node. | 
 **location** | **string**| Path location of patch file. | 

### Return type

[**ClusterPatchPatches**](ClusterPatchPatches.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHardwareStatus**
> HardwareStatus GetHardwareStatus(ctx, )


View the status of hardware upgrades in progress

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HardwareStatus**](HardwareStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUpgradeCluster**
> UpgradeCluster GetUpgradeCluster(ctx, )


Cluster wide upgrade status info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UpgradeCluster**](UpgradeCluster.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterPatchPatches**
> ClusterPatchPatchesExtended ListClusterPatchPatches(ctx, optional)


List all patches.

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
 **limit** | **int32**| Return no more than this many results at once (see resume). | 
 **location** | **string**| Path location of patch file. | 
 **local** | **bool**| Whether to view patches on the local node only. | 
 **dir** | **string**| The direction of the sort. | 

### Return type

[**ClusterPatchPatchesExtended**](ClusterPatchPatchesExtended.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterUpgrade**
> UpdateClusterUpgrade(ctx, clusterUpgrade)


Add nodes to a running upgrade.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **clusterUpgrade** | [**ClusterUpgrade**](ClusterUpgrade.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

