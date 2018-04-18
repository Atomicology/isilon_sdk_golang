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

type AuthIdNtokenPrivilegeItem struct {

	// Specifies the ID of the privilege.
	Id string `json:"id"`

	// Specifies the name of the privilege.
	Name string `json:"name,omitempty"`

	// True, if the privilege is read-only.
	ReadOnly bool `json:"read_only,omitempty"`
}