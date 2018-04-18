# AuthAccessAccessItemFileFilePermissions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dacl** | **string** | Returns a status message if the Null ACL is set. | [optional] [default to null]
**DeleteChild** | **string** | Returns a status message if the parent directoryhas the delete_child property set for the user.If the delete_child property is set for a user,that user is able to delete the file.the delete_child for the user. | [optional] [default to null]
**Expected** | **string** | Specifies the Access Control Entry (ACE) for the user. | [optional] [default to null]
**Mode** | **string** | Specifies the mode bits on the file. | [optional] [default to null]
**Ownership** | **string** | Returns a status message if the user owns the file. | [optional] [default to null]
**RelevantAces** | [**[]AuthAccessAccessItemShareSharePermissionsShareRelevantAce**](AuthAccessAccessItemShareSharePermissionsShareRelevantAce.md) | Specifies a list of the relevant Access Control Entrieswith respect to the user in the share. | [optional] [default to null]
**RelevantMode** | **string** | Specifies the mode bits that are related to the user. | [optional] [default to null]
**Sticky** | **string** | Returns a status message if the user owns the file. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


