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

type SettingsMappingExtendedExtended struct {

	// The FQDN of the source domain to map.
	Domain string `json:"domain"`

	// The FQDN of destination domain to map to.
	Mapping string `json:"mapping"`

	// The authentication provider type.
	Type_ string `json:"type"`
}
