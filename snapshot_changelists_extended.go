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

type SnapshotChangelistsExtended struct {

	Changelists []SnapshotChangelists `json:"changelists,omitempty"`

	// Continue returning results from previous call using this token (token should come from the previous call, resume cannot be used with other options).
	Resume string `json:"resume,omitempty"`

	// Total number of items available.
	Total int32 `json:"total,omitempty"`
}
