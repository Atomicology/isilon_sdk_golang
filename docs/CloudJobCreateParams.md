# CloudJobCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **[]string** | The names of accounts for COI restore | [optional] [default to null]
**Directories** | **[]string** | Directories addressed by this job | [optional] [default to null]
**ExpirationDate** | **int32** | The new expiration date in seconds | [optional] [default to null]
**FileMatchingPattern** | [***Empty**](Empty.md) | The file filtering logic to find files for this job. (Only applicable for &#39;recall&#39; jobs) | [optional] [default to null]
**Files** | **[]string** | Filenames addressed by this job | [optional] [default to null]
**Policy** | **string** | The name of an existing cloudpool policy to apply to this job. (Only applicable for &#39;archive&#39; jobs) | [optional] [default to null]
**Type_** | **string** | The type of cloud action to be performed by this job. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


