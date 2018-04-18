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

type ClusterNodeState struct {

	// Node readonly state.
	Readonly *Empty `json:"readonly,omitempty"`

	// Node service light state.
	Servicelight *ClusterNodeStateServicelight `json:"servicelight,omitempty"`

	// Node smartfail state.
	Smartfail *ClusterNodeStateSmartfail `json:"smartfail,omitempty"`
}
