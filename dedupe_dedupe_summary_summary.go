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

type DedupeDedupeSummarySummary struct {

	// Size in bytes of a filesystem block.
	BlockSize float32 `json:"block_size"`

	// Estimated number of physical blocks deduped.
	EstimatedPhysicalBlocks float32 `json:"estimated_physical_blocks"`

	// Estimated number of physical blocks saved by dedupe.
	EstimatedSavedBlocks float32 `json:"estimated_saved_blocks"`

	// Number of logical blocks deduped.
	LogicalBlocks float32 `json:"logical_blocks"`

	// Number of logical blocks saved by dedupe.
	SavedLogicalBlocks float32 `json:"saved_logical_blocks"`

	// Total physical blocks in the cluster.
	TotalBlocks float32 `json:"total_blocks"`

	// Total physical blocks used in the cluster.
	UsedBlocks float32 `json:"used_blocks"`
}