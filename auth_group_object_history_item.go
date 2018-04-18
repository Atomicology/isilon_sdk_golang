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

type AuthGroupObjectHistoryItem struct {

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Gid *AuthAccessAccessItemFileGroup `json:"gid,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Sid *AuthAccessAccessItemFileGroup `json:"sid,omitempty"`

	// Specifies properties for a persona, which consists of either a 'type' and a 'name' or an 'ID'.
	Uid *AuthAccessAccessItemFileGroup `json:"uid,omitempty"`
}
