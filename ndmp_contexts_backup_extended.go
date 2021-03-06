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

// Get NDMP Context List
type NdmpContextsBackupExtended struct {

	Contexts []NdmpContextsBackupContext `json:"contexts,omitempty"`

	// Resume string returned by previous query.
	Resume string `json:"resume,omitempty"`

	// The number of ndmp contexts.
	Total int32 `json:"total,omitempty"`
}
