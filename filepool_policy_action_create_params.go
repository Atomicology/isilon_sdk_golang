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

type FilepoolPolicyActionCreateParams struct {

	// Varies according to action_type
	ActionParam string `json:"action_param"`

	ActionType string `json:"action_type"`
}
