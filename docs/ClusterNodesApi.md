# \ClusterNodesApi

All URIs are relative to *https://YOUR_CLUSTER_HOSTNAME_OR_NODE_IP:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrivesDriveAddItem**](ClusterNodesApi.md#CreateDrivesDriveAddItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/add | 
[**CreateDrivesDriveFirmwareUpdateItem**](ClusterNodesApi.md#CreateDrivesDriveFirmwareUpdateItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/firmware/update | 
[**CreateDrivesDriveFormatItem**](ClusterNodesApi.md#CreateDrivesDriveFormatItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/format | 
[**CreateDrivesDrivePurposeItem**](ClusterNodesApi.md#CreateDrivesDrivePurposeItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/purpose | 
[**CreateDrivesDriveSmartfailItem**](ClusterNodesApi.md#CreateDrivesDriveSmartfailItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/smartfail | 
[**CreateDrivesDriveStopfailItem**](ClusterNodesApi.md#CreateDrivesDriveStopfailItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/stopfail | 
[**CreateDrivesDriveSuspendItem**](ClusterNodesApi.md#CreateDrivesDriveSuspendItem) | **Post** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/suspend | 
[**CreateNodeRebootItem**](ClusterNodesApi.md#CreateNodeRebootItem) | **Post** /platform/5/cluster/nodes/{Lnn}/reboot | 
[**CreateNodeShutdownItem**](ClusterNodesApi.md#CreateNodeShutdownItem) | **Post** /platform/5/cluster/nodes/{Lnn}/shutdown | 
[**GetDrivesDriveFirmware**](ClusterNodesApi.md#GetDrivesDriveFirmware) | **Get** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/firmware | 
[**GetNodeDrive**](ClusterNodesApi.md#GetNodeDrive) | **Get** /platform/5/cluster/nodes/{Lnn}/drives/{NodeDriveId} | 
[**GetNodeDriveconfig**](ClusterNodesApi.md#GetNodeDriveconfig) | **Get** /platform/5/cluster/nodes/{Lnn}/driveconfig | 
[**GetNodeDrives**](ClusterNodesApi.md#GetNodeDrives) | **Get** /platform/5/cluster/nodes/{Lnn}/drives | 
[**GetNodeDrivesPurposelist**](ClusterNodesApi.md#GetNodeDrivesPurposelist) | **Get** /platform/3/cluster/nodes/{Lnn}/drives-purposelist | 
[**GetNodeHardware**](ClusterNodesApi.md#GetNodeHardware) | **Get** /platform/5/cluster/nodes/{Lnn}/hardware | 
[**GetNodeHardwareFast**](ClusterNodesApi.md#GetNodeHardwareFast) | **Get** /platform/3/cluster/nodes/{Lnn}/hardware-fast | 
[**GetNodePartitions**](ClusterNodesApi.md#GetNodePartitions) | **Get** /platform/3/cluster/nodes/{Lnn}/partitions | 
[**GetNodeSensors**](ClusterNodesApi.md#GetNodeSensors) | **Get** /platform/3/cluster/nodes/{Lnn}/sensors | 
[**GetNodeSled**](ClusterNodesApi.md#GetNodeSled) | **Get** /platform/5/cluster/nodes/{Lnn}/sleds/{NodeSledId} | 
[**GetNodeSleds**](ClusterNodesApi.md#GetNodeSleds) | **Get** /platform/5/cluster/nodes/{Lnn}/sleds | 
[**GetNodeState**](ClusterNodesApi.md#GetNodeState) | **Get** /platform/3/cluster/nodes/{Lnn}/state | 
[**GetNodeStateReadonly**](ClusterNodesApi.md#GetNodeStateReadonly) | **Get** /platform/3/cluster/nodes/{Lnn}/state/readonly | 
[**GetNodeStateServicelight**](ClusterNodesApi.md#GetNodeStateServicelight) | **Get** /platform/3/cluster/nodes/{Lnn}/state/servicelight | 
[**GetNodeStateSmartfail**](ClusterNodesApi.md#GetNodeStateSmartfail) | **Get** /platform/3/cluster/nodes/{Lnn}/state/smartfail | 
[**GetNodeStatus**](ClusterNodesApi.md#GetNodeStatus) | **Get** /platform/3/cluster/nodes/{Lnn}/status | 
[**GetNodeStatusBatterystatus**](ClusterNodesApi.md#GetNodeStatusBatterystatus) | **Get** /platform/3/cluster/nodes/{Lnn}/status/batterystatus | 
[**ListDrivesDriveFirmwareUpdate**](ClusterNodesApi.md#ListDrivesDriveFirmwareUpdate) | **Get** /platform/3/cluster/nodes/{Lnn}/drives/{Driveid}/firmware/update | 
[**UpdateNodeDriveconfig**](ClusterNodesApi.md#UpdateNodeDriveconfig) | **Put** /platform/5/cluster/nodes/{Lnn}/driveconfig | 
[**UpdateNodeStateReadonly**](ClusterNodesApi.md#UpdateNodeStateReadonly) | **Put** /platform/3/cluster/nodes/{Lnn}/state/readonly | 
[**UpdateNodeStateServicelight**](ClusterNodesApi.md#UpdateNodeStateServicelight) | **Put** /platform/3/cluster/nodes/{Lnn}/state/servicelight | 
[**UpdateNodeStateSmartfail**](ClusterNodesApi.md#UpdateNodeStateSmartfail) | **Put** /platform/3/cluster/nodes/{Lnn}/state/smartfail | 


# **CreateDrivesDriveAddItem**
> Empty CreateDrivesDriveAddItem(ctx, drivesDriveAddItem, lnn, driveid)


Add a drive to a node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveAddItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDriveFirmwareUpdateItem**
> Empty CreateDrivesDriveFirmwareUpdateItem(ctx, drivesDriveFirmwareUpdateItem, lnn, driveid)


Start a drive firmware update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveFirmwareUpdateItem** | [**DrivesDriveFirmwareUpdateItem**](DrivesDriveFirmwareUpdateItem.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDriveFormatItem**
> Empty CreateDrivesDriveFormatItem(ctx, drivesDriveFormatItem, lnn, driveid)


Format a drive for use by OneFS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveFormatItem** | [**DrivesDriveFormatItem**](DrivesDriveFormatItem.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDrivePurposeItem**
> Empty CreateDrivesDrivePurposeItem(ctx, drivesDrivePurposeItem, lnn, driveid)


Assign a drive to a specific use case.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDrivePurposeItem** | [**DrivesDrivePurposeItem**](DrivesDrivePurposeItem.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDriveSmartfailItem**
> Empty CreateDrivesDriveSmartfailItem(ctx, drivesDriveSmartfailItem, lnn, driveid)


Remove a drive from use by OneFS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveSmartfailItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDriveStopfailItem**
> Empty CreateDrivesDriveStopfailItem(ctx, drivesDriveStopfailItem, lnn, driveid)


Stop restriping from a smartfailing drive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveStopfailItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDrivesDriveSuspendItem**
> Empty CreateDrivesDriveSuspendItem(ctx, drivesDriveSuspendItem, lnn, driveid)


Temporarily remove a drive from use by OneFS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **drivesDriveSuspendItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeRebootItem**
> Empty CreateNodeRebootItem(ctx, nodeRebootItem, lnn, optional)


Reboot the node specified by <LNN>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeRebootItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeRebootItem** | [**Empty**](Empty.md)|  | 
 **lnn** | **int32**|  | 
 **force** | **bool**| Force reboot on Infinity platform even if a drive sled is not present. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeShutdownItem**
> Empty CreateNodeShutdownItem(ctx, nodeShutdownItem, lnn, optional)


Shutdown the node specified by <LNN>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeShutdownItem** | [**Empty**](Empty.md)|  | 
  **lnn** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeShutdownItem** | [**Empty**](Empty.md)|  | 
 **lnn** | **int32**|  | 
 **force** | **bool**| Force shutdown on Infinity platform even if a drive sled is not present. | 

### Return type

[**Empty**](Empty.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDrivesDriveFirmware**
> DrivesDriveFirmware GetDrivesDriveFirmware(ctx, lnn, driveid)


Retrieve drive firmware information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**DrivesDriveFirmware**](DrivesDriveFirmware.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDrive**
> NodeDrives GetNodeDrive(ctx, nodeDriveId, lnn, optional)


Retrieve drive information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeDriveId** | **string**| Retrieve drive information. | 
  **lnn** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeDriveId** | **string**| Retrieve drive information. | 
 **lnn** | **int32**|  | 
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeDrives**](NodeDrives.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDriveconfig**
> NodeDriveconfig GetNodeDriveconfig(ctx, lnn, optional)


View a node's drive subsystem XML configuration file.

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
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeDriveconfig**](NodeDriveconfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDrives**
> NodeDrives GetNodeDrives(ctx, lnn, optional)


List the drives on this node.

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
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeDrives**](NodeDrives.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeDrivesPurposelist**
> NodeDrivesPurposelist GetNodeDrivesPurposelist(ctx, lnn)


Lists the available purposes for drives in this node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeDrivesPurposelist**](NodeDrivesPurposelist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeHardware**
> NodeHardware GetNodeHardware(ctx, lnn, optional)


Retrieve node hardware identity information.

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
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeHardware**](NodeHardware.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeHardwareFast**
> NodeHardwareFast GetNodeHardwareFast(ctx, lnn)


Quickly retrieve a subset of node hardware identity information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeHardwareFast**](NodeHardwareFast.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodePartitions**
> NodePartitions GetNodePartitions(ctx, lnn)


Retrieve node partition information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodePartitions**](NodePartitions.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeSensors**
> NodeSensors GetNodeSensors(ctx, lnn)


Retrieve node sensor information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeSensors**](NodeSensors.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeSled**
> NodeSleds GetNodeSled(ctx, nodeSledId, lnn, optional)


Get detailed information for the sled specified by <SLEDID>, or all sleds in the case where <SLEDID> is 'all', in the node specified by <LNN>.  Accepts <sledid> in either 'sled' or 'all' formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeSledId** | **string**| Get detailed information for the sled specified by &lt;SLEDID&gt;, or all sleds in the case where &lt;SLEDID&gt; is &#39;all&#39;, in the node specified by &lt;LNN&gt;.  Accepts &lt;sledid&gt; in either &#39;sled&#39; or &#39;all&#39; formats. | 
  **lnn** | **int32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeSledId** | **string**| Get detailed information for the sled specified by &lt;SLEDID&gt;, or all sleds in the case where &lt;SLEDID&gt; is &#39;all&#39;, in the node specified by &lt;LNN&gt;.  Accepts &lt;sledid&gt; in either &#39;sled&#39; or &#39;all&#39; formats. | 
 **lnn** | **int32**|  | 
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeSleds**](NodeSleds.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeSleds**
> NodeSleds GetNodeSleds(ctx, lnn, optional)


Get detailed information for all sleds in this node. Equivalent to /5/cluster/nodes/<lnn>/sleds/all.

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
 **timeout** | **float32**| Request timeout | 

### Return type

[**NodeSleds**](NodeSleds.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeState**
> NodeState GetNodeState(ctx, lnn)


Retrieve node state information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeState**](NodeState.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStateReadonly**
> NodeStateReadonly GetNodeStateReadonly(ctx, lnn)


Retrieve node readonly state information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeStateReadonly**](NodeStateReadonly.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStateServicelight**
> NodeStateServicelight GetNodeStateServicelight(ctx, lnn)


Retrieve node service light state information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeStateServicelight**](NodeStateServicelight.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStateSmartfail**
> NodeStateSmartfail GetNodeStateSmartfail(ctx, lnn)


Retrieve node smartfail state information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeStateSmartfail**](NodeStateSmartfail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStatus**
> NodeStatus GetNodeStatus(ctx, lnn)


Retrieve node status information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeStatus**](NodeStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStatusBatterystatus**
> NodeStatusBatterystatus GetNodeStatusBatterystatus(ctx, lnn)


Retrieve node battery status information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 

### Return type

[**NodeStatusBatterystatus**](NodeStatusBatterystatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDrivesDriveFirmwareUpdate**
> DrivesDriveFirmwareUpdate ListDrivesDriveFirmwareUpdate(ctx, lnn, driveid)


Retrieve firmware update information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **lnn** | **int32**|  | 
  **driveid** | **string**|  | 

### Return type

[**DrivesDriveFirmwareUpdate**](DrivesDriveFirmwareUpdate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeDriveconfig**
> UpdateNodeDriveconfig(ctx, nodeDriveconfig, lnn)


Modify a node's drive subsystem XML configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeDriveconfig** | [**NodeDriveconfigExtended**](NodeDriveconfigExtended.md)|  | 
  **lnn** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeStateReadonly**
> UpdateNodeStateReadonly(ctx, nodeStateReadonly, lnn)


Modify one or more node readonly state settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeStateReadonly** | [**NodeStateReadonlyExtended**](NodeStateReadonlyExtended.md)|  | 
  **lnn** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeStateServicelight**
> UpdateNodeStateServicelight(ctx, nodeStateServicelight, lnn)


Modify one or more node service light state settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeStateServicelight** | [**NodeStateServicelightExtended**](NodeStateServicelightExtended.md)|  | 
  **lnn** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeStateSmartfail**
> UpdateNodeStateSmartfail(ctx, nodeStateSmartfail, lnn)


Modify smartfail state of the node.  Only the 'smartfailed' body member has any effect on node smartfail state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **nodeStateSmartfail** | [**NodeStateSmartfailExtended**](NodeStateSmartfailExtended.md)|  | 
  **lnn** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

