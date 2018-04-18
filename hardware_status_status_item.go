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

type HardwareStatusStatusItem struct {

	// The ID of this upgrade.
	Id string `json:"id"`

	// Name of the upgrading pool.
	NodepoolName string `json:"nodepool_name,omitempty"`

	// The lnns of the nodes in the pool that haven't been upgraded yet.
	UnupgradedLnns []int32 `json:"unupgraded_lnns"`

	// The type of upgrade this is.
	UpgradeType string `json:"upgrade_type"`

	// The lnns of the nodes in the pool that have been successsfully upgraded.
	UpgradedLnns []int32 `json:"upgraded_lnns"`
}
