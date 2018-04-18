# LicenseLicenseTier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementsExceededAlerts** | [**[]LicenseLicenseTierEntitlementsExceededAlert**](LicenseLicenseTierEntitlementsExceededAlert.md) | List of alerts about exceeded entitlements: The following alerts appear when usage of a resource such as a node, an encryption node, or storage capacity exceeds the quantity licensed for that resource. | [optional] [default to null]
**LicensedDriveCapacity** | **int32** | Licensed terabyte (TB, 10^12 bytes) drive capacity allocated as storage associated with tier. Included if tier is not NONINF and license is not a base only license. | [optional] [default to null]
**LicensedNodeCount** | **int32** | Licensed number of nodes in this tier. | [optional] [default to null]
**LicensedNodesWithSedsCount** | **int32** | Licensed number of nodes of this tier that contain self-encrypting drives. Included only if license is ONEFS and tier is not NONINF. | [optional] [default to null]
**Tier** | **string** | OneFS hardware tier. Tier is a number, NONINF, or NO_TIER. NONINF indicates a non infinity tier. NO_TIER indicates a license that is not tier based. | [optional] [default to null]
**UsedDriveCapacity** | **int32** | Actual terabyte (TB, 10^12 bytes) drive capacity allocated as storage space associated with tier. Included if tier is not NONINF and license is not a base only license. | [optional] [default to null]
**UsedNodeCount** | **int32** | Actual number of nodes in this tier. | [optional] [default to null]
**UsedNodesWithSedsCount** | **int32** | Actual number of nodes of this tier that contain self-encrypting drives. Included only if license is ONEFS and if tier is not NONINF. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


