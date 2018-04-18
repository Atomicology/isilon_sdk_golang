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

type PoolsPoolRulesRule struct {

	// Description for the provisioning rule.
	Description string `json:"description"`

	// Name of the groupnet this rule belongs to
	Groupnet string `json:"groupnet"`

	// Unique rule ID.
	Id string `json:"id"`

	// Interface name the provisioning rule applies to.
	Iface string `json:"iface"`

	// Name of the provisioning rule.
	Name string `json:"name"`

	// Node type the provisioning rule applies to.
	NodeType string `json:"node_type"`

	// Name of the pool this rule belongs to.
	Pool string `json:"pool"`

	// Name of the subnet this rule belongs to.
	Subnet string `json:"subnet"`
}
