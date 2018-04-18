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

type FileFilterSettingsSettings struct {

	// List of file extensions to be filtered.
	FileFilterExtensions []string `json:"file_filter_extensions,omitempty"`

	// Specifies if filter list is for deny or allow. Default is deny.
	FileFilterType string `json:"file_filter_type,omitempty"`

	// Indicates whether file filtering is enabled on this zone.
	FileFilteringEnabled bool `json:"file_filtering_enabled,omitempty"`
}