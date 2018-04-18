# \UpgradeClusterApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodesNodePatchSyncItem**](UpgradeClusterApi.md#CreateNodesNodePatchSyncItem) | **Post** /platform/4/upgrade/cluster/nodes/{Lnn}/patch/sync | 
[**GetNodesNodeFirmwareStatus**](UpgradeClusterApi.md#GetNodesNodeFirmwareStatus) | **Get** /platform/3/upgrade/cluster/nodes/{Lnn}/firmware/status | 


# **CreateNodesNodePatchSyncItem**
> Empty CreateNodesNodePatchSyncItem(ctx, nodesNodePatchSyncItem, lnn)


Retry any pending patch sync operations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodesNodePatchSyncItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodesNodeFirmwareStatus**
> NodesNodeFirmwareStatus GetNodesNodeFirmwareStatus(ctx, lnn, optional)


The firmware status for the node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lnn** | **int32**|  | 
 **devices** | **bool**| Show devices. If false, this returns an empty list. Default is false. | 
 **package_** | **bool**| Show package. If false, this returns an empty list.Default is false. | 

### Return type

[**NodesNodeFirmwareStatus**](NodesNodeFirmwareStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

