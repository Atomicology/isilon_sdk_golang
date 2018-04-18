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

type StoragepoolTier struct {

	// The names or IDs of the tier's children.
	Children []string `json:"children,omitempty"`

	// The tier name.
	Name string `json:"name,omitempty"`
}
