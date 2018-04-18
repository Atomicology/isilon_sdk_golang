# LicenseLicense

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysSinceExpiry** | **int32** | Number of days since a license expired. | [optional] [default to null]
**DaysToExpiry** | **int32** | Number of days before a license expires. | [optional] [default to null]
**Expiration** | **string** | Date of license expiry. Format is YYYY-MM-DD. It is not included if there is no expiration. Feature is considered expired at end of this day. The cluster time is used to determine expiry. | [optional] [default to null]
**ExpiredAlert** | **bool** | True when we are generating an alert that this feature has expired. | [default to null]
**ExpiringAlert** | **bool** | True when we are generating an alert that this feature is expiring. | [default to null]
**Id** | **string** | Name of the licensed feature. | [default to null]
**Name** | **string** | Name of the licensed feature. | [default to null]
**Status** | **string** | Current status of the license. | [default to null]
**Tiers** | [**[]LicenseLicenseTier**](LicenseLicenseTier.md) | Tiered License details. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


