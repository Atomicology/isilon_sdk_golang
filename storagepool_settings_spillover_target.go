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

type StoragepoolSettingsSpilloverTarget struct {

	// Target pool ID if target specified, otherwise null.
	NameOrId string `json:"name_or_id,omitempty"`

	// Type of target pool.
	Type_ string `json:"type"`
}
