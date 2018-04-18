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

// Dedupe settings.
type DedupeSettingsExtended struct {

	// The paths that will be assessed.
	AssessPaths []string `json:"assess_paths,omitempty"`

	// The paths that will be deduped.
	Paths []string `json:"paths,omitempty"`
}
