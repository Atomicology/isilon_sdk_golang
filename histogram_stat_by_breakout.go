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

type HistogramStatByBreakout struct {

	// List of bucket, file count pairs.
	Data [][]int32 `json:"data"`

	// Breakout key by which results are filtered.
	Key string `json:"key"`
}
