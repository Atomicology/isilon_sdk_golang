# JobJobSummarySummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIsDegraded** | **bool** | Whether the cluster is in a degraded state.  Note this is from the perspective of the node handling the query, which might be different from another node. | [default to null]
**Connected** | **bool** | Whether isi_job_d instances on all up nodes in the cluster are connected to the coordinator. | [default to null]
**Coordinator** | **int32** | The devid of the job engine coordinator. | [default to null]
**DisconnectedNodes** | **[]int32** | If connected&#x3D;false, this is the set of devids that are not connected to the coordinator. | [default to null]
**DownOrReadOnlyNodes** | **bool** | Whether the cluster has any down or read-only nodes.  These nodes are not considered in the connected property. | [default to null]
**NextJid** | **int32** | The job ID to be assigned to the next job. | [default to null]
**RunDegraded** | **bool** | Whether the job engine would allow most jobs to run even when the cluster is in a degraded state. | [default to null]
**StatsReady** | **bool** | Whether the coordinator has recently gathered statistics for all nodes in the cluster. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


