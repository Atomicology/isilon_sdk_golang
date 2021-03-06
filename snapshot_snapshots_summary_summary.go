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

type SnapshotSnapshotsSummarySummary struct {

	// Total number of active snapshots.
	ActiveCount int32 `json:"active_count"`

	// Sum of sizes of active snapshots.
	ActiveSize int32 `json:"active_size"`

	// Total number of snapshot aliases.
	AliasesCount int32 `json:"aliases_count"`

	// Total number of snapshots.
	Count int32 `json:"count"`

	// Total number of delete-pending snapshots.
	DeletingCount int32 `json:"deleting_count"`

	// Sum of sizes of delete-pending snapshots.
	DeletingSize int32 `json:"deleting_size"`

	// Sum of shadow bytes of all snapshots.
	ShadowBytes int32 `json:"shadow_bytes"`

	// Sum of sizes in bytes of all snapshots.
	Size int32 `json:"size"`
}
