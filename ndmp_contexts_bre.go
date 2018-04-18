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

// View a NDMP restartable backup Context
type NdmpContextsBre struct {

	// Backup type
	BackupType string `json:"backup type,omitempty"`

	// Backup Context ID
	BkpContextId string `json:"bkp_context_id,omitempty"`

	// Unique ID of NDMP BRE context
	BreContextId string `json:"bre_context_id,omitempty"`

	// Context creation time
	CreateTime int32 `json:"create_time,omitempty"`

	// List of environment variables for restartable backup
	EnvVariables []NdmpContextsBreEnvVariable `json:"env_variables,omitempty"`

	// Unique display id.
	Id string `json:"id,omitempty"`

	// Backup result
	Results int32 `json:"results,omitempty"`

	// Snapshot name of backup
	SnapName string `json:"snap_name,omitempty"`

	// Context status bits
	Status int32 `json:"status,omitempty"`

	// Backup Stream ID
	StreamId int32 `json:"stream_id,omitempty"`
}
