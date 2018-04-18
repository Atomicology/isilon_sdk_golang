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

type StoragepoolTierUsage struct {

	// Available free bytes remaining in the pool when virtual hot spare is taken into account.
	AvailBytes string `json:"avail_bytes"`

	// Available free bytes remaining in the pool on HDD drives when virtual hot spare is taken into account.
	AvailHddBytes string `json:"avail_hdd_bytes"`

	// Available free bytes remaining in the pool on SSD drives when virtual hot spare is taken into account.
	AvailSsdBytes string `json:"avail_ssd_bytes"`

	// Whether or not the pool usage is currently balanced.
	Balanced bool `json:"balanced"`

	// Free bytes remaining in the pool.
	FreeBytes string `json:"free_bytes"`

	// Free bytes remaining in the pool on HDD drives.
	FreeHddBytes string `json:"free_hdd_bytes"`

	// Free bytes remaining in the pool on SSD drives.
	FreeSsdBytes string `json:"free_ssd_bytes"`

	// Percentage of usable space in the pool which is used.
	PctUsed string `json:"pct_used"`

	// Percentage of usable space on HDD drives in the pool which is used.
	PctUsedHdd string `json:"pct_used_hdd"`

	// Percentage of usable space on SSD drives in the pool which is used.
	PctUsedSsd string `json:"pct_used_ssd"`

	// Total bytes in the pool.
	TotalBytes string `json:"total_bytes"`

	// Total bytes in the pool on HDD drives.
	TotalHddBytes string `json:"total_hdd_bytes"`

	// Total bytes in the pool on SSD drives.
	TotalSsdBytes string `json:"total_ssd_bytes"`

	// Total bytes in the pool drives when virtual hot spare is taken into account.
	UsableBytes string `json:"usable_bytes"`

	// Total bytes in the pool on HDD drives when virtual hot spare is taken into account.
	UsableHddBytes string `json:"usable_hdd_bytes"`

	// Total bytes in the pool on SSD drives when virtual hot spare is taken into account.
	UsableSsdBytes string `json:"usable_ssd_bytes"`

	// Used bytes in the pool.
	UsedBytes string `json:"used_bytes,omitempty"`

	// Used bytes in the pool on HDD drives.
	UsedHddBytes string `json:"used_hdd_bytes"`

	// Used bytes in the pool on SSD drives.
	UsedSsdBytes string `json:"used_ssd_bytes"`

	VirtualHotSpareBytes string `json:"virtual_hot_spare_bytes,omitempty"`
}