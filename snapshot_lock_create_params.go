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

type SnapshotLockCreateParams struct {

	// The Unix Epoch time the snapshot lock will expire and be eligible for automatic deletion.
	Expires int32 `json:"expires,omitempty"`

	// Free form comment.
	Comment string `json:"comment,omitempty"`
}
