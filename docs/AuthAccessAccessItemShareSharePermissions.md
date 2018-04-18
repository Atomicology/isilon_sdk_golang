# AuthAccessAccessItemShareSharePermissions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedPermissions** | **string** | Returns Share level permissions for the user.{ &#39;read&#39; , &#39;write&#39; , &#39;full&#39; or &#39;none&#39; will be the values} | [optional] [default to null]
**ImpersonateGuest** | **bool** | Returns whether impersonate guest setting is enabled for the user on the share. | [optional] [default to null]
**ImpersonateUser** | **bool** | Returns whether impersonate user setting is enabled on the share | [optional] [default to null]
**RunAsRoot** | **bool** | Returns whether run as root is enabled for the user on the share | [optional] [default to null]
**ShareRelevantAces** | [**[]AuthAccessAccessItemShareSharePermissionsShareRelevantAce**](AuthAccessAccessItemShareSharePermissionsShareRelevantAce.md) | Specifies a list of the relevant Access Control Entries withrespect to the user in the share. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


