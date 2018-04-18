# LicenseLicensesExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licenses** | [**[]LicenseLicense**](LicenseLicense.md) |  | [optional] [default to null]
**ActivationIncompleteAlert** | **bool** | True when we are generating an activation incomplete alert. An activation incomplete alert is generated if we do not have a signed license file 90 days after OneFS is upgraded. | [default to null]
**BaseOnlyLicenses** | **[]string** |  | [default to null]
**Evaluatable** | **[]string** |  | [default to null]
**Swid** | **string** | Software license identifier. SWID will be absent if not yet obtained from a license file. | [optional] [default to null]
**ValidSignature** | **bool** | True if license file contains a valid signature. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


