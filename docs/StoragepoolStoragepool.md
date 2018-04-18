# StoragepoolStoragepool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDisableL3** | **bool** | Indicates if disabling L3 is possible. | [optional] [default to null]
**CanEnableL3** | **bool** | Indicates if enabling L3 is possible. L3 cannot be enabled if there are unprovisioned drives. | [optional] [default to null]
**Children** | **[]string** | The names or IDs of the pool&#39;s children. | [optional] [default to null]
**HealthFlags** | **[]string** | An array of containing any health issues with this pool.  If the pool is healthy, the list is empty.  Only appears on &#39;nodepool&#39; type storagepools. | [optional] [default to null]
**Id** | **int32** | The system ID given to the pool. | [default to null]
**L3** | **bool** | Use SSDs in this node pool for L3 cache. | [optional] [default to null]
**L3Status** | **string** | &#39;storage&#39; if the &#39;l3&#39; option is disabled. If the l3 option is enabled, &#39;migrating&#39; if any SSDs in this node pool have not yet been migrated to L3. If all SSDs have been migrated, &#39;l3&#39;. | [optional] [default to null]
**Lnns** | **[]int32** | The nodes that are part of this pool. | [default to null]
**Manual** | **bool** | Whether or not the node pool is manually managed. | [optional] [default to null]
**Name** | **string** | The pool name. | [default to null]
**ProtectionPolicy** | **string** | The underlying protection policy. | [optional] [default to null]
**Type_** | **string** | The pool type. | [default to null]
**Usage** | [***StoragepoolTierUsage**](StoragepoolTierUsage.md) | Total pool usage. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


