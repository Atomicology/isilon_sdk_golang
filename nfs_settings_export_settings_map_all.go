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

type NfsSettingsExportSettingsMapAll struct {

	// True if the user mapping is applied.
	Enabled bool `json:"enabled,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	PrimaryGroup *AuthAccessAccessItemFileGroup `json:"primary_group,omitempty"`

	// Specifies persona properties for the secondary user group. A persona consists of either a type and name, or an ID.
	SecondaryGroups []NfsSettingsExportSettingsMapAllSecondaryGroups `json:"secondary_groups,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	User *AuthAccessAccessItemFileGroup `json:"user,omitempty"`
}
