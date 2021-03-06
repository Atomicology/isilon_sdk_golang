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

type NodeStatusNodePowersupplies struct {

	// Count of how many power supplies are supported.
	Count int32 `json:"count,omitempty"`

	// Count of how many power supplies have failed.
	Failures int32 `json:"failures,omitempty"`

	// Does this node have a CFF power supply.
	HasCff bool `json:"has_cff,omitempty"`

	// A descriptive status string for this node's power supplies.
	Status string `json:"status,omitempty"`

	// List of this node's power supplies.
	Supplies []NodeStatusNodePowersuppliesSupply `json:"supplies,omitempty"`

	// Does this node support CFF power supplies.
	SupportsCff bool `json:"supports_cff,omitempty"`
}
