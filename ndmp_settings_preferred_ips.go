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

// Get a list of preferred ip preferences.
type NdmpSettingsPreferredIps struct {

	Preferences []NdmpSettingsPreferredIpsPreference `json:"preferences,omitempty"`

	// Resume string returned by previous query.
	Resume string `json:"resume,omitempty"`

	// The number of preferences.
	Total int32 `json:"total,omitempty"`
}
