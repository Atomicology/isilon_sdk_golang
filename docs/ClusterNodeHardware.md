# ClusterNodeHardware

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | **string** | Name of this node&#39;s chassis. | [optional] [default to null]
**ChassisCode** | **string** | Chassis code of this node (1U, 2U, etc.). | [optional] [default to null]
**ChassisCount** | **string** | Number of chassis making up this node. | [optional] [default to null]
**ChassisDepth** | **string** | Chassis depth for this node if applicable (Normal, Deep, Unknown). If not supported: Unknown. If Logic to determine chassis depth fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist. | [optional] [default to null]
**Class** | **string** | Class of this node (storage, accelerator, etc.). | [optional] [default to null]
**CodeName** | **string** | Code name of this node if applicable (Minnetonka, MiniHuron, Huron, Union, Tahoe, Superior, Unknown). If not supported: Unknown. If Logic to determine code name fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist. | [optional] [default to null]
**ComputeType** | **string** | Type of compute node if applicable (Low, Medium, High, Turbo, Ultra, Unknown). If not supported: Unknown. If Logic to determine compute type fails: Unknown. If PSI_Get fails: Unknown. PSI_Get can fail if PSI not initialized, or key does not exist. | [optional] [default to null]
**ConfigurationId** | **string** | Node configuration ID. | [optional] [default to null]
**Cpu** | **string** | Manufacturer and model of this node&#39;s CPU. | [optional] [default to null]
**DiskController** | **string** | Manufacturer and model of this node&#39;s disk controller. | [optional] [default to null]
**DiskExpander** | **string** | Manufacturer and model of this node&#39;s disk expander. | [optional] [default to null]
**FamilyCode** | **string** | Family code of this node (X, S, NL, etc.). | [optional] [default to null]
**FlashDrive** | **string** | Manufacturer, model, and device id of this node&#39;s flash drive. | [optional] [default to null]
**GenerationCode** | **string** | Generation code of this node. | [optional] [default to null]
**Hwgen** | **string** | Isilon hardware generation name. | [optional] [default to null]
**ImbVersion** | **string** | Version of this node&#39;s Isilon Management Board. | [optional] [default to null]
**Infiniband** | **string** | Infiniband card type. | [optional] [default to null]
**LcdVersion** | **string** | Version of the LCD panel. | [optional] [default to null]
**Model** | **string** | Isilon node model identifier string (S200, X410, Infinity-H500, etc.). | [optional] [default to null]
**ModelCode** | **string** | Isilon node model code string (S200, X410, H500, etc.). | [optional] [default to null]
**Motherboard** | **string** | Manufacturer and model of this node&#39;s motherboard. | [optional] [default to null]
**NetInterfaces** | **string** | Description of all this node&#39;s network interfaces. | [optional] [default to null]
**NodeSlotId** | **int32** | Position of node within chassis (e.g., 1-4 for Infinity chassis). -1 for error or not supported. | [optional] [default to null]
**Nvram** | **string** | Manufacturer and model of this node&#39;s NVRAM board. | [optional] [default to null]
**PeerSerialNumber** | **string** | Serial number of this node&#39;s peer/buddy node.(Infinity Only) | [optional] [default to null]
**PerformanceCode** | **string** | Performance code of this node, if applicable (2, 4, 5, etc.). | [optional] [default to null]
**Powersupplies** | **[]string** | Description strings for each power supply on this node. | [optional] [default to null]
**Processor** | **string** | Number of processors and cores on this node. | [optional] [default to null]
**Product** | **string** | Isilon product name. | [optional] [default to null]
**Ram** | **int32** | Size of RAM in bytes. | [optional] [default to null]
**SerialNumber** | **string** | Serial number of this node. | [optional] [default to null]
**Series** | **string** | Series of this node (X, I, NL, etc.). | [optional] [default to null]
**SledDriveCount** | **int32** | Size of drive sleds in node, if applicable. Expected values: 3, 4, 6. 0 if unable to determine sled size. -1 for error or not supported. If PSI_Get fails: -1. PSI_Get can fail if PSI not initialized, or key does not exist. | [optional] [default to null]
**StorageClass** | **string** | Storage class of this node (storage or diskless). | [optional] [default to null]
**Tier** | **int32** | Platform tier level of this node if applicable (1-4 are defined, 0 for unknown or not supported, -1 for error). If not supported: 0. If Logic to determine tier fails: 0 for unknown. If PSI_Get fails: -1 for error. PSI_Get can fail if PSI not initialized, or key does not exist. | [optional] [default to null]
**TopLevelAssemblySerialNumber** | **string** | Serial number of the top level assembly of this node.(Infinity Only) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


