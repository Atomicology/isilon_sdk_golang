# ClusterFirmwareUpgradeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDevice** | **string** | Exclude the specified devices in the firmware upgrade. | [optional] [default to null]
**ExcludeType** | **string** | Exclude the specified device type in the firmware upgrade. | [optional] [default to null]
**IncludeDevice** | **string** | Include the specified devices in the firmware upgrade. | [optional] [default to null]
**IncludeType** | **string** | Include the specified device type in the firmware upgrade. | [optional] [default to null]
**NoBurn** | **bool** | Do not burn the firmware. | [optional] [default to null]
**NoReboot** | **bool** | Do not reboot the node after an upgrade | [optional] [default to null]
**NoVerify** | **bool** | Do not verify the firmware upgrade after an upgrade. | [optional] [default to null]
**NodesToUpgrade** | **[]int32** | The nodes scheduled for upgrade. Order in array determines queue position number. &#39;All&#39; and null option will upgrade all nodes in &lt;lnn&gt; order. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


