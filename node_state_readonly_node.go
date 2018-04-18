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

type NodeStateReadonlyNode struct {

	// The current read-only mode allowed status for the node.
	Allowed bool `json:"allowed,omitempty"`

	// The current read-only user mode status for the node. NOTE: If read-only mode is currently disallowed for this node, it will remain read/write until read-only mode is allowed again. This value only sets or clears any user-specified requests for read-only mode. If the node has been placed into read-only mode by the system, it will remain in read-only mode until the system conditions which triggered read-only mode have cleared.
	Enabled bool `json:"enabled,omitempty"`

	// Error message, if the HTTP status returned from this node was not 200.
	Error_ string `json:"error,omitempty"`

	// Node ID of the node reporting this information.
	Id int32 `json:"id,omitempty"`

	// Logical node number of the node reporting this information.
	Lnn int32 `json:"lnn,omitempty"`

	// The current read-only mode status for the node.
	Mode bool `json:"mode,omitempty"`

	// The current read-only mode status description for the node.
	Status string `json:"status,omitempty"`

	// The read-only state values are valid (False = Error).
	Valid bool `json:"valid,omitempty"`

	// The current read-only value (enumerated bitfield) for the node.
	Value int32 `json:"value,omitempty"`
}
