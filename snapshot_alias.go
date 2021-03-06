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

type SnapshotAlias struct {

	// The user or system supplied snapshot alias name.
	Name string `json:"name,omitempty"`

	// Target snapshot for this snapshot alias.
	Target string `json:"target,omitempty"`
}
