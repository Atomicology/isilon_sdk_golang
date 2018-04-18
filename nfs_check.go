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

type NfsCheck struct {

	// The ID of the export.
	Id int32 `json:"id,omitempty"`

	// The message about the export.
	Messages string `json:"messages"`
}