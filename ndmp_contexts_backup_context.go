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

type NdmpContextsBackupContext struct {

	// Context ID
	ContextId string `json:"context_id,omitempty"`

	// Unique display id.
	Id string `json:"id,omitempty"`

	// The directory of the backup. This is not applicable to restore contexts.
	Path string `json:"path,omitempty"`

	// Snapshot ID reserved for the context. This is not applicable to restore contexts.
	Snapid int32 `json:"snapid,omitempty"`

	// Context creation time
	StartTime int32 `json:"start_time,omitempty"`

	// Whether the context is active.
	Status string `json:"status,omitempty"`

	// The number of sessions in the context
	TotalSessions int32 `json:"total_sessions,omitempty"`

	// Type of context; It can be bre, backup, and restore
	Type_ string `json:"type,omitempty"`
}