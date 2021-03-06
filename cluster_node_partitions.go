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

type ClusterNodePartitions struct {

	// Count of how many partitions are included.
	Count int32 `json:"count,omitempty"`

	// Partition information.
	Partitions []NodePartitionsNodePartition `json:"partitions,omitempty"`
}
