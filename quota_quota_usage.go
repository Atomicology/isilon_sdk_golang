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

type QuotaQuotaUsage struct {

	// Number of inodes (filesystem entities) used by governed data.
	Inodes int32 `json:"inodes"`

	// Apparent bytes used by governed data.
	Logical int32 `json:"logical"`

	// Bytes used for governed data and filesystem overhead.
	Physical int32 `json:"physical"`
}