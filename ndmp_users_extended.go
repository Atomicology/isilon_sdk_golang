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

type NdmpUsersExtended struct {

	// Total number of items available.
	Total int32 `json:"total,omitempty"`

	Users []NdmpUsers `json:"users,omitempty"`
}
