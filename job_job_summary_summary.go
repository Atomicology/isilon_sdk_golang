/*
 * Isilon SDK
 *
 * Isilon SDK - Language bindings for the OneFS API
 *
 * API version: 5
 * Contact: sdk@isilon.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package isi_sdk_8_1_0

type JobJobSummarySummary struct {

	// Whether the cluster is in a degraded state.  Note this is from the perspective of the node handling the query, which might be different from another node.
	ClusterIsDegraded bool `json:"cluster_is_degraded"`

	// Whether isi_job_d instances on all up nodes in the cluster are connected to the coordinator.
	Connected bool `json:"connected"`

	// The devid of the job engine coordinator.
	Coordinator int32 `json:"coordinator"`

	// If connected=false, this is the set of devids that are not connected to the coordinator.
	DisconnectedNodes []int32 `json:"disconnected_nodes"`

	// Whether the cluster has any down or read-only nodes.  These nodes are not considered in the connected property.
	DownOrReadOnlyNodes bool `json:"down_or_read_only_nodes"`

	// The job ID to be assigned to the next job.
	NextJid int32 `json:"next_jid"`

	// Whether the job engine would allow most jobs to run even when the cluster is in a degraded state.
	RunDegraded bool `json:"run_degraded"`

	// Whether the coordinator has recently gathered statistics for all nodes in the cluster.
	StatsReady bool `json:"stats_ready"`
}
