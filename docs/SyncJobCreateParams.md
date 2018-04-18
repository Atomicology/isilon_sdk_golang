# SyncJobCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to be taken by this job. | [optional] [default to null]
**Id** | **string** | The ID or Name of the policy | [default to null]
**LogLevel** | **string** | Only valid for allow_write, and allow_write_revert; specify the desired logging level, will be stored in the logs for isi_migrate, defaults to &#39;info&#39;. | [optional] [default to null]
**SourceSnapshot** | **string** | An optional snapshot to copy/sync from. | [optional] [default to null]
**WorkersPerNode** | **int32** | Only valid for allow_write, and allow_write_revert; specify the desired workers per node, defaults to 3. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


