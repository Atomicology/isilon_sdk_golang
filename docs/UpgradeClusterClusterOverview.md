# UpgradeClusterClusterOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodesCurrent** | **int32** | Number of nodes running the current OneFS version. | [optional] [default to null]
**NodesTotal** | **int32** | Total number of nodes on the cluster. | [optional] [default to null]
**NodesTransitioning** | **int32** | Number of nodes transitioning between OneFS versions. Null if the cluster_state is &#39;committed&#39; or &#39;assessing.&#39; | [optional] [default to null]
**NodesUpgraded** | **int32** | Number of nodes running the upgraded OneFS version. Null if the cluster_state is &#39;committed&#39; or &#39;assessing.&#39; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


