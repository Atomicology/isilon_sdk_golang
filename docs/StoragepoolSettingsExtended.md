# StoragepoolSettingsExtended

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticallyManageIoOptimization** | **string** | Automatically manage IO optimization settings on files. | [optional] [default to null]
**AutomaticallyManageProtection** | **string** | Automatically manage protection settings on files. | [optional] [default to null]
**GlobalNamespaceAccelerationEnabled** | **bool** | Optimize namespace operations by storing metadata on SSDs. | [optional] [default to null]
**ProtectDirectoriesOneLevelHigher** | **bool** | Automatically add additional protection level to all directories. | [optional] [default to null]
**SpilloverEnabled** | **bool** | Spill writes into other pools as needed. | [optional] [default to null]
**SpilloverTarget** | [***StoragepoolSettingsSpilloverTarget**](StoragepoolSettingsSpilloverTarget.md) | Target pool for spilled writes. | [optional] [default to null]
**SsdL3CacheDefaultEnabled** | **bool** | The L3 Cache default enabled state. This specifies whether L3 Cache should be enabled on new node pools | [optional] [default to null]
**SsdQabMirrors** | **string** | Controls number of mirrors of QAB blocks to place on SSDs. | [optional] [default to null]
**SsdSystemBtreeMirrors** | **string** | Controls number of mirrors of system B-tree blocks to place on SSDs. | [optional] [default to null]
**SsdSystemDeltaMirrors** | **string** | Controls number of mirrors of system delta blocks to place on SSDs. | [optional] [default to null]
**VirtualHotSpareDenyWrites** | **bool** | Deny writes into reserved virtual hot spare space. | [optional] [default to null]
**VirtualHotSpareHideSpare** | **bool** | Hide reserved virtual hot spare space from free space counts. | [optional] [default to null]
**VirtualHotSpareLimitDrives** | **int32** | The number of drives to reserve for the virtual hot spare, from 0-4. | [optional] [default to null]
**VirtualHotSpareLimitPercent** | **int32** | The percent space to reserve for the virtual hot spare, from 0-20. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


