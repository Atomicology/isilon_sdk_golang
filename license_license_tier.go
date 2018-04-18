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

type LicenseLicenseTier struct {

	// List of alerts about exceeded entitlements: The following alerts appear when usage of a resource such as a node, an encryption node, or storage capacity exceeds the quantity licensed for that resource.
	EntitlementsExceededAlerts []LicenseLicenseTierEntitlementsExceededAlert `json:"entitlements_exceeded_alerts,omitempty"`

	// Licensed terabyte (TB, 10^12 bytes) drive capacity allocated as storage associated with tier. Included if tier is not NONINF and license is not a base only license.
	LicensedDriveCapacity int32 `json:"licensed_drive_capacity,omitempty"`

	// Licensed number of nodes in this tier.
	LicensedNodeCount int32 `json:"licensed_node_count,omitempty"`

	// Licensed number of nodes of this tier that contain self-encrypting drives. Included only if license is ONEFS and tier is not NONINF.
	LicensedNodesWithSedsCount int32 `json:"licensed_nodes_with_seds_count,omitempty"`

	// OneFS hardware tier. Tier is a number, NONINF, or NO_TIER. NONINF indicates a non infinity tier. NO_TIER indicates a license that is not tier based.
	Tier string `json:"tier,omitempty"`

	// Actual terabyte (TB, 10^12 bytes) drive capacity allocated as storage space associated with tier. Included if tier is not NONINF and license is not a base only license.
	UsedDriveCapacity int32 `json:"used_drive_capacity,omitempty"`

	// Actual number of nodes in this tier.
	UsedNodeCount int32 `json:"used_node_count,omitempty"`

	// Actual number of nodes of this tier that contain self-encrypting drives. Included only if license is ONEFS and if tier is not NONINF.
	UsedNodesWithSedsCount int32 `json:"used_nodes_with_seds_count,omitempty"`
}
