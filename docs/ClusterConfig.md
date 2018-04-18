# ClusterConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Customer configurable description. | [default to null]
**Devices** | [**[]ClusterConfigDevice**](ClusterConfigDevice.md) |  | [default to null]
**Encoding** | **string** | Default encoding. | [default to null]
**Guid** | **string** | Cluster GUID. | [default to null]
**HasQuorum** | **bool** | If true, the local node is in a group with quorum: It is communicating, storing, and protecting data normally with other nodes in its group.  Under normal circumstances, every node in the cluster is part of one group. | [default to null]
**IsCompliance** | **bool** | If true, the cluster is in compliance mode.  Compliance mode clusters do not allow root access and have stricter WORM (Write Once Read Many) features for retaining data in compliance with U.S. Securities and Exchange Commission rule 17a-4. | [default to null]
**IsVirtual** | **bool** | true if the cluster is deployed on a desktop VMWorkstation | [default to null]
**IsVonefs** | **bool** | true if this is a vOneFS cluster on an ESXi | [default to null]
**JoinMode** | **string** | Node join mode: &#39;manual&#39; or &#39;secure&#39;. | [default to null]
**LocalDevid** | **int32** | Device ID of the queried node. | [default to null]
**LocalLnn** | **int32** | Device logical node number of the queried node. | [default to null]
**LocalSerial** | **string** | Device serial number of the queried node. | [default to null]
**Name** | **string** | Cluster name. | [default to null]
**OnefsVersion** | [***ClusterConfigOnefsVersion**](ClusterConfigOnefsVersion.md) |  | [optional] [default to null]
**Timezone** | [***ClusterConfigTimezone**](ClusterConfigTimezone.md) | The cluster timezone settings. | [optional] [default to null]
**UpgradeType** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


