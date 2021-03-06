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

type HardwareStartItem struct {

	// The nodepool ID or name on which to start the upgrade.
	NodePool string `json:"node_pool"`

	// The type of upgrade to start.
	UpgradeType string `json:"upgrade_type"`
}
