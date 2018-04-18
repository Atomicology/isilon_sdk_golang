# SubnetsSubnetPool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessZone** | **string** | Name of a valid access zone to map IP address pool to the zone. | [optional] [default to null]
**AggregationMode** | **string** | OneFS supports the following NIC aggregation modes. | [optional] [default to null]
**AllocMethod** | **string** | Specifies how IP address allocation is done among pool members. | [optional] [default to null]
**Description** | **string** | A description of the pool. | [optional] [default to null]
**Ifaces** | [**[]SubnetsSubnetPoolIface**](SubnetsSubnetPoolIface.md) | List of interface members in this pool. | [optional] [default to null]
**Name** | **string** | The name of the pool. It must be unique throughout the given subnet.It&#39;s a required field with POST method. | [optional] [default to null]
**Ranges** | [**[]SubnetsSubnetPoolRange**](SubnetsSubnetPoolRange.md) | List of IP address ranges in this pool. | [optional] [default to null]
**RebalancePolicy** | **string** | Rebalance policy.. | [optional] [default to null]
**ScAutoUnsuspendDelay** | **int32** | Time delay in seconds before a node which has been                 automatically unsuspended becomes usable in SmartConnect                responses for pool zones. | [optional] [default to null]
**ScConnectPolicy** | **string** | SmartConnect client connection balancing policy. | [optional] [default to null]
**ScDnsZone** | **string** | SmartConnect zone name for the pool. | [optional] [default to null]
**ScDnsZoneAliases** | **[]string** | List of SmartConnect zone aliases (DNS names) to the pool. | [optional] [default to null]
**ScFailoverPolicy** | **string** | SmartConnect IP failover policy. | [optional] [default to null]
**ScSubnet** | **string** | Name of SmartConnect service subnet for this pool. | [optional] [default to null]
**ScTtl** | **int32** | Time to live value for SmartConnect DNS query responses in seconds. | [optional] [default to null]
**StaticRoutes** | [**[]SubnetsSubnetPoolStaticRoute**](SubnetsSubnetPoolStaticRoute.md) | List of interface members in this pool. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


