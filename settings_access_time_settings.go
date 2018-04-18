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

type SettingsAccessTimeSettings struct {

	// Enable access time tracking.
	Enabled bool `json:"enabled"`

	// Access time tracked in seconds for each cluster file if enabled.
	Precision int32 `json:"precision"`
}
