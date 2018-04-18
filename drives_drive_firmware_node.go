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

type DrivesDriveFirmwareNode struct {

	// List of the firmware revisions on the drives in this node.
	Drives []DrivesDriveFirmwareNodeDrive `json:"drives,omitempty"`

	// Error message, if the HTTP status returned from this node was not 200.
	Error_ string `json:"error,omitempty"`

	// Node ID of the node reporting this information.
	Id int32 `json:"id,omitempty"`

	// Logical node number of the node reporting this information.
	Lnn int32 `json:"lnn,omitempty"`

	// Status of the HTTP response from this node if not 200.  If 200, this field does not appear.
	Status int32 `json:"status,omitempty"`
}