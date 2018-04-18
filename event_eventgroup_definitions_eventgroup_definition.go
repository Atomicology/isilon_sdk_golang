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

type EventEventgroupDefinitionsEventgroupDefinition struct {

	// ID of eventgroup category: all, 100000000 (SYS_DISK_EVENTS), 200000000 (NODE_STATUS_EVENTS), 300000000 (REBOOT_EVENTS), 400000000 (SW_EVENTS), 500000000 (QUOTA_EVENTS), 600000000 (SNAP_EVENTS), 700000000 (WINNET_EVENTS), 800000000 (FILESYS_EVENTS), 900000000 (HW_EVENTS), 1100000000 (CPOOL_EVENTS)
	Category string `json:"category,omitempty"`

	// Human readable description - may contain value placeholders.
	Description string `json:"description,omitempty"`

	// Unique Identifier.
	Id string `json:"id,omitempty"`

	// Name for eventgroup.
	Name string `json:"name,omitempty"`
}
