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

type SummaryProtocolStatsProtocolStatsDisk struct {

	// Disk iops
	Iops float32 `json:"iops"`

	// Disk reads
	Read float32 `json:"read"`

	// Disk writes
	Write float32 `json:"write"`
}
