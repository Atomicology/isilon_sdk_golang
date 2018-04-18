# ClusterPatchPatchesPatch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | A long comment about the patch. | [optional] [default to null]
**Conflicts** | **[]string** | Other patches that this patch conflicts with. | [optional] [default to null]
**Dependencies** | **[]string** | Other patches that this patch depends on. | [optional] [default to null]
**Description** | **string** | A short description of the patch. | [optional] [default to null]
**Files** | [**[]ClusterPatchPatchesPatchFile**](ClusterPatchPatchesPatchFile.md) | The files contained in this patch. | [optional] [default to null]
**Id** | **string** | A unique identifier for the patch. | [optional] [default to null]
**Name** | **string** | The name of the patch. | [optional] [default to null]
**Nodes** | **[]int32** | The nodes that this patch is installed on. | [optional] [default to null]
**Reboot** | **string** | Describes the reboot requirements | [optional] [default to null]
**Services** | [**[]ClusterPatchPatchesPatchService**](ClusterPatchPatchesPatchService.md) | The services affected during the patch deployment | [optional] [default to null]
**Status** | **string** | The intallation status of this patch on the cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


