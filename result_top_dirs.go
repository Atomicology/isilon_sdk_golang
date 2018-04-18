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

type ResultTopDirs struct {

	// Change in directory ranking from result set comparison.
	Change int32 `json:"change,omitempty"`

	// Directory access time enabled.
	DirAtimeEnabled bool `json:"dir_atime_enabled"`

	// Directory listing.
	Dirs []ResultTopDirsDir `json:"dirs"`

	// Limit on number of top results.
	TopNMax int32 `json:"top_n_max"`

	// Total count of directory results.
	TotalCount int32 `json:"total_count"`
}