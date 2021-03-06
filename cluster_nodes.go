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

// The node details useful during an upgrade or assessment.
type ClusterNodes struct {

	// The current OneFS version before upgrade.
	Error_ *ClusterNodesError `json:"error,omitempty"`

	// The last action performed to completion/failure on this node.  Null if the node_state is 'committed' or 'assessing.' One of the following values: 'upgrade', 'rollback'.
	LastAction string `json:"last_action,omitempty"`

	// Did the node pass upgrade or rollback without failing? Null if the node_state is 'committed.' One of the following values: 'pass', 'fail', null
	LastActionResult string `json:"last_action_result,omitempty"`

	// The lnn of the node.
	Lnn int32 `json:"lnn,omitempty"`

	// \\e The state of the node during the upgrade, rollback, or assessment. One of the following values: 'committed', 'upgraded', 'upgrading', 'rolling back', 'assessing', 'error'
	NodeState string `json:"node_state,omitempty"`

	// The current OneFS version before upgrade.
	OnefsVersion *ClusterNodesOnefsVersion `json:"onefs_version,omitempty"`

	// What step is the upgrade, assessment, or rollback in? To show via progress indicator. NOTE: the value is an integer between 0 and 100 (percent)
	Progress int32 `json:"progress,omitempty"`
}
