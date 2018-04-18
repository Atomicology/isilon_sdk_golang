# UpgradeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterOverview** | [***UpgradeClusterClusterOverview**](UpgradeClusterClusterOverview.md) | The cluster overview of an upgrade process. | [optional] [default to null]
**ClusterState** | **string** | The different states of an upgrade, rollback, or assessment. One of the following values: &#39;committed&#39;, &#39;upgraded&#39;, &#39;partially upgraded&#39;, &#39;upgrading&#39;, &#39;rolling back&#39;, &#39;assessing&#39;, &#39;error&#39; | [optional] [default to null]
**CurrentProcess** | **string** | The current upgrade activity. | [optional] [default to null]
**FinishTime** | **string** | The time when a rollback, assessment or upgrade has finished completely. Use ISO 8601 standard. Null if the cluster_state is not &#39;upgraded&#39;. | [optional] [default to null]
**InstallImagePath** | **string** | The location (path) of the upgrade image which must be within /ifs. Null if the cluster_state is &#39;committed&#39; or &#39;upgraded.&#39; | [optional] [default to null]
**NodeMedianTime** | **int32** | The median time (seconds) to complete each node so far during this upgrade. Before the first node in an upgrade has completed this key will have an associated null value. | [optional] [default to null]
**OnefsVersionCurrent** | [***ClusterNodesOnefsVersion**](ClusterNodesOnefsVersion.md) | The current OneFS version before upgrade. | [optional] [default to null]
**OnefsVersionUpgrade** | [***ClusterNodesOnefsVersion**](ClusterNodesOnefsVersion.md) | The OneFS version the user is attempting to upgrade to. Null if the cluster_state is &#39;committed&#39; or &#39;assessing.&#39; | [optional] [default to null]
**PatchAction** | **string** | The most recent patch action performed. | [optional] [default to null]
**PatchName** | **string** | The patch with the most recent patch action. | [optional] [default to null]
**StartTime** | **string** | The time when an upgrade, rollback, or assessment was started. Use ISO 8601 standard. Null if the cluster_state is &#39;committed&#39; or &#39;partially upgraded.&#39; | [optional] [default to null]
**UpgradeIsCommitted** | **bool** | True if upgrade is committed. | [optional] [default to null]
**UpgradeSettings** | [***UpgradeClusterUpgradeSettings**](UpgradeClusterUpgradeSettings.md) | The settings necessary when starting an upgrade. Null if the cluster_state is not &#39;upgrading&#39; or &#39;partially upgraded.&#39; or &#39;error&#39;. | [optional] [default to null]
**UpgradeTriggeredTime** | **string** | Time at which upgrade was originally requested. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


