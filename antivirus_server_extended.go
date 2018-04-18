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

type AntivirusServerExtended struct {

	// A description for the server.
	Description string `json:"description,omitempty"`

	// Whether the server is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The icap url for the server.  This should have a format of: icap://host.domain:port/path
	Url string `json:"url,omitempty"`

	Definitions string `json:"definitions,omitempty"`

	// A unique identifier for the server.
	Id string `json:"id,omitempty"`

	// The status of the server.
	Status string `json:"status,omitempty"`
}
