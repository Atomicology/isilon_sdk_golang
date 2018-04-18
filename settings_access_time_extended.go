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

type SettingsAccessTimeExtended struct {

	// Enable access time tracking.
	Enabled bool `json:"enabled,omitempty"`

	// Access time tracked on each cluster file accurate to this number of seconds.
	Precision int32 `json:"precision,omitempty"`
}