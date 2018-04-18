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

type StoragepoolUnprovisionedUnprovisioned struct {

	// A list of unprovisioned drives that do not belong to an unprovisioned node.
	Drives []StoragepoolStatusUnprovisionedItem `json:"drives"`

	// A list of lnns whose drives are all unprovisioned
	Lnns []int32 `json:"lnns"`
}
